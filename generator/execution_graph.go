package generator

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	dag "github.com/laser/random-dag-generator-go"
	"github.com/mathnogueira/tracegen/generator/application"
	"github.com/mathnogueira/tracegen/generator/application/operation"
	"go.opentelemetry.io/otel/trace"
)

type ExecutionNode struct {
	Operation *application.Operation
	Next      []*ExecutionNode
	IsRoot    bool
}

func (n *ExecutionNode) IsLeaf() bool {
	return len(n.Next) == 0
}

type ExecutionGraph struct {
	EntryPoint ExecutionNode
}

func newExecutionGraph(services []*application.Service) *ExecutionGraph {
	serviceNodes := make([][]*ExecutionNode, len(services))
	leafNodes := make([][]*ExecutionNode, len(services))
	rootNodes := make([][]*ExecutionNode, len(services))
	for i, service := range services {
		serviceNodes[i] = newServiceExecutionNodes(service)
		leafNodes[i] = make([]*ExecutionNode, 0)
		rootNodes[i] = make([]*ExecutionNode, 0)
		for _, node := range serviceNodes[i] {
			if node.IsLeaf() {
				leafNodes[i] = append(leafNodes[i], node)
			}

			if node.IsRoot {
				rootNodes[i] = append(rootNodes[i], node)
			}
		}
	}

	for i := range services {
		linkNodes(leafNodes, rootNodes, i)
	}

	// promote one root to be the only root
	realRoot := rootNodes[0][rand.Intn(len(rootNodes[0]))]
	for _, root := range rootNodes[0] {
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

func linkNodes(source, target [][]*ExecutionNode, index int) {
	numberServices := len(source)
	if index+1 >= numberServices {
		return
	}

	totalNumberSpansNeedLinking := len(target[index+1])
	i := 0
	for totalNumberSpansNeedLinking > 0 {
		leaf := source[index][i%len(source[index])]
		root := target[index+1][i]

		fmt.Printf("[%s] %s --> [%s] %s\n", leaf.Operation.Entity, leaf.Operation.Name, root.Operation.Entity, root.Operation.Name)

		leaf.Next = append(leaf.Next, root)
		totalNumberSpansNeedLinking--
		i++
	}
}

func newServiceExecutionNodes(service *application.Service) []*ExecutionNode {
	graph := dag.Random(
		dag.WithNodeQty(len(service.Operations)),
		dag.WithMaxOutdegree(4),
		dag.WithEdgeFactor(0.5),
	)

	nodes := make(map[string]*ExecutionNode)
	for i, node := range graph.Nodes {
		nodes[node.GetID()] = &ExecutionNode{
			Operation: service.Operations[i],
			Next:      make([]*ExecutionNode, 0),
			IsRoot:    true,
		}
	}

	for _, edge := range graph.Edges {
		source := nodes[edge.GetSourceNodeID()]
		target := nodes[edge.GetTargetNodeID()]

		target.IsRoot = false
		source.Next = append(source.Next, target)
	}

	nodesSlice := make([]*ExecutionNode, 0)
	for _, node := range nodes {
		nodesSlice = append(nodesSlice, node)
	}

	return nodesSlice
}

func (g *ExecutionGraph) Execute(ctx context.Context, opts ...Option) {
	for _, opt := range opts {
		ctx = opt(ctx)
	}

	g.EntryPoint.Execute(ctx)
}

func (n *ExecutionNode) Execute(ctx context.Context) context.Context {
	fmt.Printf("Running [%s] %s\n", n.Operation.Entity, n.Operation.Name)
	ctx, span := n.Operation.CreateSpan(ctx)
	defer span.End()

	span.SetAttributes(operation.GenerateRandomAttributes(5)...)
	event := operation.NewEvent()
	span.AddEvent(event.Name, trace.WithAttributes(event.Attributes...))

	time.Sleep(100 * time.Millisecond)

	for _, next := range n.Next {
		next.Execute(ctx)
	}

	return ctx
}
