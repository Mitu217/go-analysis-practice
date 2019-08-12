package practice

import (
	"log"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const Doc = `analysis practice command`

var Analyzer = &analysis.Analyzer{
	Name: "analysis practice",
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	funcs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	for _, fn := range funcs {
		for _, block := range fn.Blocks {
			for _, instr := range block.Instrs {
				switch instr := instr.(type) {
				case *ssa.Alloc:
					log.Printf("ssa.Alloc: %v", instr)
				case *ssa.BinOp:
					log.Printf("ssa.BinOp: %v", instr)
				case *ssa.Call:
					log.Printf("ssa.Call: %v", instr)
				case *ssa.Defer:
					log.Printf("ssa.Defer: %v", instr)
				case *ssa.Go:
					log.Printf("ssa.Go: %v", instr)
				case *ssa.ChangeInterface:
					log.Printf("ssa.ChangeInterface: %v", instr)
				case *ssa.ChangeType:
					log.Printf("ssa.ChangeType: %v", instr)
				case *ssa.Convert:
					log.Printf("ssa.Convert: %v", instr)
				case *ssa.Field:
					log.Printf("ssa.Field: %v", instr)
				case *ssa.FieldAddr:
					log.Printf("ssa.FieldAddr: %v", instr)
				case *ssa.IndexAddr:
					log.Printf("ssa.IndexAddr: %v", instr)
				case *ssa.MapUpdate:
					log.Printf("ssa.MapUpdate: %v", instr)
				case *ssa.Slice:
					log.Printf("ssa.Slice: %v", instr)
				case *ssa.Store:
					log.Printf("ssa.Store: %v", instr)
				case *ssa.Extract:
					log.Printf("ssa.Extract: %v", instr)
				case *ssa.TypeAssert:
					log.Printf("ssa.TypeAssert: %v", instr)
				case *ssa.UnOp:
					log.Printf("ssa.UnOp: %v", instr)
				case *ssa.DebugRef:
					log.Printf("ssa.DebugRef: %v", instr)
				case *ssa.If:
					log.Printf("ssa.If: %v", instr)
				case *ssa.Index:
					log.Printf("ssa.Index: %v", instr)
				case *ssa.Jump:
					log.Printf("ssa.Jump: %v", instr)
				case *ssa.Lookup:
					log.Printf("ssa.Lookup: %v", instr)
				case *ssa.MakeChan:
					log.Printf("ssa.MakeChan: %v", instr)
				case *ssa.MakeClosure:
					log.Printf("ssa.MakeClosure: %v", instr)
				case *ssa.MakeInterface:
					log.Printf("ssa.MakeInterface: %v", instr)
				case *ssa.MakeSlice:
					log.Printf("ssa.MakeSlice: %v", instr)
				case *ssa.Next:
					log.Printf("ssa.Next: %v", instr)
				case *ssa.Panic:
					log.Printf("ssa.Panic: %v", instr)
				case *ssa.Phi:
					log.Printf("ssa.Phi: %v", instr)
				case *ssa.Range:
					log.Printf("ssa.Range: %v", instr)
				case *ssa.Return:
					log.Printf("ssa.Return: %v", instr)
				case *ssa.RunDefers:
					log.Printf("ssa.RunDefers: %v", instr)
				case *ssa.Select:
					log.Printf("ssa.Select: %v", instr)
				case *ssa.Send:
					log.Printf("ssa.Send: %v", instr)
				}
			}
		}
	}
	return nil, nil
}
