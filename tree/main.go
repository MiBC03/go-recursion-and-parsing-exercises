package main

import (
	"fmt"
	"strings"
)

// Estrutura que representa um nó da árvore binária
type TreeNode struct {
	ID        string
	Value     int
	Left      *TreeNode
	Right     *TreeNode
	PosX      float64
	PosY      float64
}

// Função auxiliar para criar um nó
func NewNode(id string, value int, left, right *TreeNode) *TreeNode {
	return &TreeNode{
		ID:    id,
		Value: value,
		Left:  left,
		Right: right,
	}
}

// Função recursiva para calcular as posições dos nós
func calculatePositions(
	node *TreeNode,
	level int,
	scale float64,
	leftLimit float64,
) (*TreeNode, float64) {

	if node == nil {
		return nil, leftLimit
	}

	posY := scale * float64(level)

	// Nó folha
	if node.Left == nil && node.Right == nil {
		node.PosX = leftLimit
		node.PosY = posY
		return node, leftLimit + scale
	}

	// Apenas filho esquerdo
	if node.Left != nil && node.Right == nil {
		leftChild, rightLimit := calculatePositions(node.Left, level+1, scale, leftLimit)
		node.Left = leftChild
		node.PosX = leftChild.PosX
		node.PosY = posY
		return node, rightLimit
	}

	// Apenas filho direito
	if node.Left == nil && node.Right != nil {
		rightChild, rightLimit := calculatePositions(node.Right, level+1, scale, leftLimit)
		node.Right = rightChild
		node.PosX = rightChild.PosX
		node.PosY = posY
		return node, rightLimit
	}

	// Dois filhos
	leftChild, midLimit := calculatePositions(node.Left, level+1, scale, leftLimit)
	rightChild, rightLimit := calculatePositions(node.Right, level+1, scale, midLimit+scale)

	node.Left = leftChild
	node.Right = rightChild
	node.PosX = (leftChild.PosX + rightChild.PosX) / 2
	node.PosY = posY

	return node, rightLimit
}

// Impressão da árvore
func printTree(node *TreeNode, indent int) {
	if node == nil {
		return
	}

	fmt.Printf("%s%s (%d) - x: %.1f, y: %.1f\n",
		strings.Repeat("  ", indent),
		node.ID,
		node.Value,
		node.PosX,
		node.PosY,
	)

	printTree(node.Left, indent+1)
	printTree(node.Right, indent+1)
}

func main() {
	tree :=
		NewNode("raiz", 50,
			NewNode("esquerdo", 30,
				NewNode("esq-esq", 20, nil, nil),
				NewNode("esq-dir", 40, nil, nil),
			),
			NewNode("direito", 70,
				NewNode("dir-esq", 60, nil, nil),
				NewNode("dir-dir", 80, nil, nil),
			),
		)

	scale := 50.0
	tree, _ = calculatePositions(tree, 0, scale, 0)
	printTree(tree, 0)
}
