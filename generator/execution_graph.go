package generator

import (
	"context"
	"math/rand"
	"time"

	"github.com/kubeshop/tracegen/generator/application"
	dag "github.com/laser/random-dag-generator-go"
)

type ExecutionNode struct {
	Operation *application.Operation
	Next      []*ExecutionNode
	IsRoot    bool
}

type ExecutionGraph struct {
	EntryPoint ExecutionNode
}

func newExecutionGraph(operations []*application.Operation) *ExecutionGraph {
	graph := dag.Random(
		dag.WithNodeQty(len(operations)),
		dag.WithMaxOutdegree(4),
		dag.WithEdgeFactor(0.5),
	)

	nodes := make(map[string]*ExecutionNode)
	for i, node := range graph.Nodes {
		nodes[node.GetID()] = &ExecutionNode{
			Operation: operations[i],
			Next:      make([]*ExecutionNode, 0),
			IsRoot:    true,
		}
	}

	referencedNodes := make(map[string]bool)

	for _, edge := range graph.Edges {
		source := nodes[edge.GetSourceNodeID()]
		target := nodes[edge.GetTargetNodeID()]

		if _, ok := referencedNodes[edge.GetTargetNodeID()]; ok {
			// an operation can only be referenced once
			continue
		}

		referencedNodes[edge.GetTargetNodeID()] = true

		source.Next = append(source.Next, target)
		target.IsRoot = false
	}

	rootNodes := make([]*ExecutionNode, 0)
	for _, node := range nodes {
		if node.IsRoot {
			rootNodes = append(rootNodes, node)
		}
	}

	// promote one root to be the only root
	realRoot := rootNodes[rand.Intn(len(rootNodes))]
	for _, root := range rootNodes {
		if realRoot == root {
			continue
		}

		realRoot.Next = append(realRoot.Next, root)
		root.IsRoot = false
	}

	return &ExecutionGraph{
		EntryPoint: *realRoot,
	}
}

func (g *ExecutionGraph) Execute(ctx context.Context) {
	g.EntryPoint.Execute(ctx)
}

func (n *ExecutionNode) Execute(ctx context.Context) context.Context {
	ctx, span := n.Operation.CreateSpan(ctx)
	defer span.End()

	time.Sleep(100 * time.Millisecond)

	for _, next := range n.Next {
		next.Execute(ctx)
	}

	return ctx
}
