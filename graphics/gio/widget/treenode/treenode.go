package treenode

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

func New() TreeNode {
	return TreeNode{}
}

type TreeNode struct {
	Text     string
	Children []TreeNode
	component.DiscloserState
}

func (n *TreeNode) LayoutTreeNodeDefault(
	theme *material.Theme, gtx layout.Context,
	treeNode *TreeNode,
) layout.Dimensions {
	if len(treeNode.Children) == 0 {
		return layout.UniformInset(unit.Dp(2)).Layout(gtx,
			material.Body1(theme, treeNode.Text).Layout)
	}
	children := make([]layout.FlexChild, 0, len(treeNode.Children))
	for i := range treeNode.Children {
		child := &treeNode.Children[i]
		children = append(children, layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return n.LayoutTreeNodeDefault(theme, gtx, child)
			}))
	}
	return component.SimpleDiscloser(theme, &treeNode.DiscloserState).Layout(gtx,
		material.Body1(theme, treeNode.Text).Layout,
		func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
		})
}
