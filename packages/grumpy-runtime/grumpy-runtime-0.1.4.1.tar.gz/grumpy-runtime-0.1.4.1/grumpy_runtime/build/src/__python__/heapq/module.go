package heapq
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/heapq.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAttributeError := πg.InternStr("AttributeError")
		ßNone := πg.InternStr("None")
		ßStopIteration := πg.InternStr("StopIteration")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ß__about__ := πg.InternStr("__about__")
		ß__all__ := πg.InternStr("__all__")
		ß__lt__ := πg.InternStr("__lt__")
		ß__self__ := πg.InternStr("__self__")
		ß_heapify_max := πg.InternStr("_heapify_max")
		ß_heappushpop_max := πg.InternStr("_heappushpop_max")
		ß_nlargest := πg.InternStr("_nlargest")
		ß_nsmallest := πg.InternStr("_nsmallest")
		ß_siftdown := πg.InternStr("_siftdown")
		ß_siftdown_max := πg.InternStr("_siftdown_max")
		ß_siftup := πg.InternStr("_siftup")
		ß_siftup_max := πg.InternStr("_siftup_max")
		ßappend := πg.InternStr("append")
		ßchain := πg.InternStr("chain")
		ßcmp_lt := πg.InternStr("cmp_lt")
		ßcount := πg.InternStr("count")
		ßenumerate := πg.InternStr("enumerate")
		ßhasattr := πg.InternStr("hasattr")
		ßheapify := πg.InternStr("heapify")
		ßheappop := πg.InternStr("heappop")
		ßheappush := πg.InternStr("heappush")
		ßheappushpop := πg.InternStr("heappushpop")
		ßheapreplace := πg.InternStr("heapreplace")
		ßimap := πg.InternStr("imap")
		ßislice := πg.InternStr("islice")
		ßitemgetter := πg.InternStr("itemgetter")
		ßiter := πg.InternStr("iter")
		ßitertools := πg.InternStr("itertools")
		ßizip := πg.InternStr("izip")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßmap := πg.InternStr("map")
		ßmax := πg.InternStr("max")
		ßmerge := πg.InternStr("merge")
		ßmin := πg.InternStr("min")
		ßnext := πg.InternStr("next")
		ßnlargest := πg.InternStr("nlargest")
		ßnsmallest := πg.InternStr("nsmallest")
		ßoperator := πg.InternStr("operator")
		ßpop := πg.InternStr("pop")
		ßrange := πg.InternStr("range")
		ßreversed := πg.InternStr("reversed")
		ßsort := πg.InternStr("sort")
		ßsorted := πg.InternStr("sorted")
		ßtee := πg.InternStr("tee")
		ßxrange := πg.InternStr("xrange")
		var πTemp001 []*πg.Object
		_ = πTemp001
		var πTemp002 *πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 []πg.Param
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		var πTemp017 *πg.Object
		_ = πTemp017
		var πTemp018 *πg.Object
		_ = πTemp018
		var πTemp019 *πg.Object
		_ = πTemp019
		var πTemp020 *πg.Object
		_ = πTemp020
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 3: """Heap queue algorithm (a.k.a. priority queue).
			πF.SetLineno(3)
			// line 35: __about__ = """Heap queues
			πF.SetLineno(35)
			if πE = πF.Globals().SetItem(πF, ß__about__.ToObject(), πg.NewStr("Heap queues\n\n[explanation by Fran\xe7ois Pinard]\n\nHeaps are arrays for which a[k] <= a[2*k+1] and a[k] <= a[2*k+2] for\nall k, counting elements from 0.  For the sake of comparison,\nnon-existing elements are considered to be infinite.  The interesting\nproperty of a heap is that a[0] is always its smallest element.\n\nThe strange invariant above is meant to be an efficient memory\nrepresentation for a tournament.  The numbers below are `k', not a[k]:\n\n                                   0\n\n                  1                                 2\n\n          3               4                5               6\n\n      7       8       9       10      11      12      13      14\n\n    15 16   17 18   19 20   21 22   23 24   25 26   27 28   29 30\n\n\nIn the tree above, each cell `k' is topping `2*k+1' and `2*k+2'.  In\na usual binary tournament we see in sports, each cell is the winner\nover the two cells it tops, and we can trace the winner down the tree\nto see all opponents s/he had.  However, in many computer applications\nof such tournaments, we do not need to trace the history of a winner.\nTo be more memory efficient, when a winner is promoted, we try to\nreplace it by something else at a lower level, and the rule becomes\nthat a cell and the two cells it tops contain three different items,\nbut the top cell \"wins\" over the two topped cells.\n\nIf this heap invariant is protected at all time, index 0 is clearly\nthe overall winner.  The simplest algorithmic way to remove it and\nfind the \"next\" winner is to move some loser (let's say cell 30 in the\ndiagram above) into the 0 position, and then percolate this new 0 down\nthe tree, exchanging values, until the invariant is re-established.\nThis is clearly logarithmic on the total number of items in the tree.\nBy iterating over all items, you get an O(n ln n) sort.\n\nA nice feature of this sort is that you can efficiently insert new\nitems while the sort is going on, provided that the inserted items are\nnot \"better\" than the last 0'th element you extracted.  This is\nespecially useful in simulation contexts, where the tree holds all\nincoming events, and the \"win\" condition means the smallest scheduled\ntime.  When an event schedule other events for execution, they are\nscheduled into the future, so they can easily go into the heap.  So, a\nheap is a good structure for implementing schedulers (this is what I\nused for my MIDI sequencer :-).\n\nVarious structures for implementing schedulers have been extensively\nstudied, and heaps are good for this, as they are reasonably speedy,\nthe speed is almost constant, and the worst case is not much different\nthan the average case.  However, there are other representations which\nare more efficient overall, yet the worst cases might be terrible.\n\nHeaps are also very useful in big disk sorts.  You most probably all\nknow that a big sort implies producing \"runs\" (which are pre-sorted\nsequences, which size is usually related to the amount of CPU memory),\nfollowed by a merging passes for these runs, which merging is often\nvery cleverly organised[1].  It is very important that the initial\nsort produces the longest runs possible.  Tournaments are a good way\nto that.  If, using all the memory available to hold a tournament, you\nreplace and percolate items that happen to fit the current run, you'll\nproduce runs which are twice the size of the memory for random input,\nand much better for input fuzzily ordered.\n\nMoreover, if you output the 0'th item on disk and get an input which\nmay not fit in the current tournament (because the value \"wins\" over\nthe last output value), it cannot fit in the heap, so the size of the\nheap decreases.  The freed memory could be cleverly reused immediately\nfor progressively building a second heap, which grows at exactly the\nsame rate the first heap is melting.  When the first heap completely\nvanishes, you switch heaps and start a new run.  Clever and quite\neffective!\n\nIn a word, heaps are useful memory structures to know.  I use them in\na few applications, and I think it is good to keep a `heap' module\naround. :-)\n\n--------------------\n[1] The disk balancing algorithms which are current, nowadays, are\nmore annoying than clever, and this is a consequence of the seeking\ncapabilities of the disks.  On devices which cannot seek, like big\ntape drives, the story was quite different, and one had to be very\nclever to ensure (far in advance) that each tape movement will be the\nmost effective possible (that is, will best participate at\n\"progressing\" the merge).  Some tapes were even able to read\nbackwards, and this was also used to avoid the rewinding time.\nBelieve me, real good tape sorts were quite spectacular to watch!\nFrom all times, sorting has always been a Great Art! :-)\n").ToObject()); πE != nil {
				continue
			}
			// line 129: __all__ = ['heappush', 'heappop', 'heapify', 'heapreplace', 'merge',
			πF.SetLineno(129)
			πTemp001 = make([]*πg.Object, 8)
			πTemp001[0] = ßheappush.ToObject()
			πTemp001[1] = ßheappop.ToObject()
			πTemp001[2] = ßheapify.ToObject()
			πTemp001[3] = ßheapreplace.ToObject()
			πTemp001[4] = ßmerge.ToObject()
			πTemp001[5] = ßnlargest.ToObject()
			πTemp001[6] = ßnsmallest.ToObject()
			πTemp001[7] = ßheappushpop.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 132: import itertools
			πF.SetLineno(132)
			if πTemp001, πE = πg.ImportModule(πF, "itertools"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßitertools.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 133: islice = itertools.islice
			πF.SetLineno(133)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßislice, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßislice.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 134: count = itertools.count
			πF.SetLineno(134)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcount, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcount.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 135: imap = itertools.imap
			πF.SetLineno(135)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßimap, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßimap.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 136: izip = itertools.izip
			πF.SetLineno(136)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßizip, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßizip.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 137: tee = itertools.tee
			πF.SetLineno(137)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtee, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtee.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 138: chain = itertools.chain
			πF.SetLineno(138)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßchain, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßchain.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 139: import operator
			πF.SetLineno(139)
			if πTemp001, πE = πg.ImportModule(πF, "operator"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßoperator.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 140: itemgetter = operator.itemgetter
			πF.SetLineno(140)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßitemgetter.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 142: def cmp_lt(x, y):
			πF.SetLineno(142)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp004[1] = πg.Param{Name: "y", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("cmp_lt", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µy *πg.Object = πArgs[1]; _ = µy
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 145: return (x < y) if hasattr(x, '__lt__') else (not y <= x)
					πF.SetLineno(145)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					πTemp002[1] = ß__lt__.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µx, µy); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					goto Label2
				Label1:
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LE(πF, µy, µx); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp005).ToObject()
					πTemp001 = πTemp003
				Label2:
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcmp_lt.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 147: def heappush(heap, item):
			πF.SetLineno(147)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "heap", Def: nil}
			πTemp004[1] = πg.Param{Name: "item", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("heappush", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µheap *πg.Object = πArgs[0]; _ = µheap
				var µitem *πg.Object = πArgs[1]; _ = µitem
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 148: """Push item onto heap, maintaining the heap invariant."""
					πF.SetLineno(148)
					// line 149: heap.append(item)
					πF.SetLineno(149)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µheap, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 150: _siftdown(heap, 0, len(heap)-1)
					πF.SetLineno(150)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp001[0] = µheap
					πTemp001[1] = πg.NewInt(0).ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp004[0] = µheap
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_siftdown); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßheappush.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 152: def heappop(heap):
			πF.SetLineno(152)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "heap", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("heappop", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µheap *πg.Object = πArgs[0]; _ = µheap
				var µlastelt *πg.Object = πg.UnboundLocal; _ = µlastelt
				var µreturnitem *πg.Object = πg.UnboundLocal; _ = µreturnitem
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 153: """Pop the smallest item off the heap, maintaining the heap invariant."""
					πF.SetLineno(153)
					// line 154: lastelt = heap.pop()    # raises appropriate IndexError if heap is empty
					πF.SetLineno(154)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µheap, ßpop, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µlastelt = πTemp002
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µheap); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 155: if heap:
					πF.SetLineno(155)
				Label1:
					// line 156: returnitem = heap[0]
					πF.SetLineno(156)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µheap, πTemp001); πE != nil {
						continue
					}
					µreturnitem = πTemp002
					// line 157: heap[0] = lastelt
					πF.SetLineno(157)
					if πE = πg.CheckLocal(πF, µlastelt, "lastelt"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlastelt); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µheap, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 158: _siftup(heap, 0)
					πF.SetLineno(158)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp004[0] = µheap
					πTemp004[1] = πg.NewInt(0).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_siftup); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label3
				Label2:
					// line 160: returnitem = lastelt
					πF.SetLineno(160)
					if πE = πg.CheckLocal(πF, µlastelt, "lastelt"); πE != nil {
						continue
					}
					µreturnitem = µlastelt
					goto Label3
				Label3:
					// line 161: return returnitem
					πF.SetLineno(161)
					if πE = πg.CheckLocal(πF, µreturnitem, "returnitem"); πE != nil {
						continue
					}
					πR = µreturnitem
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßheappop.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 163: def heapreplace(heap, item):
			πF.SetLineno(163)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "heap", Def: nil}
			πTemp004[1] = πg.Param{Name: "item", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("heapreplace", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µheap *πg.Object = πArgs[0]; _ = µheap
				var µitem *πg.Object = πArgs[1]; _ = µitem
				var µreturnitem *πg.Object = πg.UnboundLocal; _ = µreturnitem
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 164: """Pop and return the current smallest value, and add the new item.
					πF.SetLineno(164)
					// line 174: returnitem = heap[0]    # raises appropriate IndexError if heap is empty
					πF.SetLineno(174)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µheap, πTemp001); πE != nil {
						continue
					}
					µreturnitem = πTemp002
					// line 175: heap[0] = item
					πF.SetLineno(175)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µitem); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µheap, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 176: _siftup(heap, 0)
					πF.SetLineno(176)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp003[0] = µheap
					πTemp003[1] = πg.NewInt(0).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_siftup); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 177: return returnitem
					πF.SetLineno(177)
					if πE = πg.CheckLocal(πF, µreturnitem, "returnitem"); πE != nil {
						continue
					}
					πR = µreturnitem
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßheapreplace.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 179: def heappushpop(heap, item):
			πF.SetLineno(179)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "heap", Def: nil}
			πTemp004[1] = πg.Param{Name: "item", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("heappushpop", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µheap *πg.Object = πArgs[0]; _ = µheap
				var µitem *πg.Object = πArgs[1]; _ = µitem
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 180: """Fast version of a heappush followed by a heappop."""
					πF.SetLineno(180)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp001 = µheap
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(2)
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µheap, πTemp004); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp003[1] = µitem
					if πTemp004, πE = πg.ResolveGlobal(πF, ßcmp_lt); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 181: if heap and cmp_lt(heap[0], item):
					πF.SetLineno(181)
				Label2:
					// line 182: item, heap[0] = heap[0], item
					πF.SetLineno(182)
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µheap, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp005, µitem).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
						continue
					}
					µitem = πTemp004
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µheap, πTemp004, πTemp005); πE != nil {
						continue
					}
					// line 183: _siftup(heap, 0)
					πF.SetLineno(183)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp003[0] = µheap
					πTemp003[1] = πg.NewInt(0).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_siftup); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label3
				Label3:
					// line 184: return item
					πF.SetLineno(184)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πR = µitem
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßheappushpop.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 186: def heapify(x):
			πF.SetLineno(186)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("heapify", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 187: """Transform list into a heap, in-place, in O(len(x)) time."""
					πF.SetLineno(187)
					// line 188: n = len(x)
					πF.SetLineno(188)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µn = πTemp003
					πTemp001 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.FloorDiv(πF, µn, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßreversed); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp006 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µi = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 195: _siftup(x, i)
					πF.SetLineno(195)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp001[1] = µi
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_siftup); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßheapify.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 197: def _heappushpop_max(heap, item):
			πF.SetLineno(197)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "heap", Def: nil}
			πTemp004[1] = πg.Param{Name: "item", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("_heappushpop_max", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µheap *πg.Object = πArgs[0]; _ = µheap
				var µitem *πg.Object = πArgs[1]; _ = µitem
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 198: """Maxheap version of a heappush followed by a heappop."""
					πF.SetLineno(198)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp001 = µheap
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp003[0] = µitem
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µheap, πTemp004); πE != nil {
						continue
					}
					πTemp003[1] = πTemp005
					if πTemp004, πE = πg.ResolveGlobal(πF, ßcmp_lt); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 199: if heap and cmp_lt(item, heap[0]):
					πF.SetLineno(199)
				Label2:
					// line 200: item, heap[0] = heap[0], item
					πF.SetLineno(200)
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µheap, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp005, µitem).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
						continue
					}
					µitem = πTemp004
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µheap, πTemp004, πTemp005); πE != nil {
						continue
					}
					// line 201: _siftup_max(heap, 0)
					πF.SetLineno(201)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp003[0] = µheap
					πTemp003[1] = πg.NewInt(0).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_siftup_max); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label3
				Label3:
					// line 202: return item
					πF.SetLineno(202)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πR = µitem
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_heappushpop_max.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 204: def _heapify_max(x):
			πF.SetLineno(204)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("_heapify_max", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 205: """Transform list into a maxheap, in-place, in O(len(x)) time."""
					πF.SetLineno(205)
					// line 206: n = len(x)
					πF.SetLineno(206)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µn = πTemp003
					πTemp001 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.FloorDiv(πF, µn, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßreversed); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp006 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µi = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 208: _siftup_max(x, i)
					πF.SetLineno(208)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp001[1] = µi
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_siftup_max); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_heapify_max.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 210: def nlargest(n, iterable):
			πF.SetLineno(210)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "n", Def: nil}
			πTemp004[1] = πg.Param{Name: "iterable", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("nlargest", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
				var µiterable *πg.Object = πArgs[1]; _ = µiterable
				var µit *πg.Object = πg.UnboundLocal; _ = µit
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µ_heappushpop *πg.Object = πg.UnboundLocal; _ = µ_heappushpop
				var µelem *πg.Object = πg.UnboundLocal; _ = µelem
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 211: """Find the n largest elements in a dataset.
					πF.SetLineno(211)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 215: if n < 0:
					πF.SetLineno(215)
				Label1:
					// line 216: return []
					πF.SetLineno(216)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 217: it = iter(iterable)
					πF.SetLineno(217)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp003[0] = µiterable
					if πTemp001, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µit = πTemp004
					// line 218: result = list(islice(it, n))
					πF.SetLineno(218)
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp005[0] = µit
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp005[1] = µn
					if πTemp001, πE = πg.ResolveGlobal(πF, ßislice); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µresult = πTemp004
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µresult); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 219: if not result:
					πF.SetLineno(219)
				Label3:
					// line 220: return result
					πF.SetLineno(220)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πR = µresult
					continue
					goto Label4
				Label4:
					// line 221: heapify(result)
					πF.SetLineno(221)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003[0] = µresult
					if πTemp001, πE = πg.ResolveGlobal(πF, ßheapify); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 222: _heappushpop = heappushpop
					πF.SetLineno(222)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßheappushpop); πE != nil {
						continue
					}
					µ_heappushpop = πTemp001
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µit); πE != nil {
						continue
					}
					πF.PushCheckpoint(6)
					πTemp002 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label7
					}
					if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						µelem = πTemp004
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 224: _heappushpop(result, elem)
					πF.SetLineno(224)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003[0] = µresult
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					πTemp003[1] = µelem
					if πE = πg.CheckLocal(πF, µ_heappushpop, "_heappushpop"); πE != nil {
						continue
					}
					if πTemp004, πE = µ_heappushpop.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					// line 225: result.sort(reverse=True)
					πF.SetLineno(225)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"reverse", πTemp001},
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µresult, ßsort, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, πTemp007); πE != nil {
						continue
					}
					// line 226: return result
					πF.SetLineno(226)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πR = µresult
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnlargest.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 228: def nsmallest(n, iterable):
			πF.SetLineno(228)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "n", Def: nil}
			πTemp004[1] = πg.Param{Name: "iterable", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("nsmallest", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
				var µiterable *πg.Object = πArgs[1]; _ = µiterable
				var µit *πg.Object = πg.UnboundLocal; _ = µit
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µ_heappushpop *πg.Object = πg.UnboundLocal; _ = µ_heappushpop
				var µelem *πg.Object = πg.UnboundLocal; _ = µelem
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 229: """Find the n smallest elements in a dataset.
					πF.SetLineno(229)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 233: if n < 0:
					πF.SetLineno(233)
				Label1:
					// line 234: return []
					πF.SetLineno(234)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 235: it = iter(iterable)
					πF.SetLineno(235)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp003[0] = µiterable
					if πTemp001, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µit = πTemp004
					// line 236: result = list(islice(it, n))
					πF.SetLineno(236)
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp005[0] = µit
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp005[1] = µn
					if πTemp001, πE = πg.ResolveGlobal(πF, ßislice); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µresult = πTemp004
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µresult); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 237: if not result:
					πF.SetLineno(237)
				Label3:
					// line 238: return result
					πF.SetLineno(238)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πR = µresult
					continue
					goto Label4
				Label4:
					// line 239: _heapify_max(result)
					πF.SetLineno(239)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003[0] = µresult
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_heapify_max); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 240: _heappushpop = _heappushpop_max
					πF.SetLineno(240)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_heappushpop_max); πE != nil {
						continue
					}
					µ_heappushpop = πTemp001
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µit); πE != nil {
						continue
					}
					πF.PushCheckpoint(6)
					πTemp002 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label7
					}
					if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						µelem = πTemp004
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 242: _heappushpop(result, elem)
					πF.SetLineno(242)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003[0] = µresult
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					πTemp003[1] = µelem
					if πE = πg.CheckLocal(πF, µ_heappushpop, "_heappushpop"); πE != nil {
						continue
					}
					if πTemp004, πE = µ_heappushpop.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					// line 243: result.sort()
					πF.SetLineno(243)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µresult, ßsort, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 244: return result
					πF.SetLineno(244)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πR = µresult
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnsmallest.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 249: def _siftdown(heap, startpos, pos):
			πF.SetLineno(249)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "heap", Def: nil}
			πTemp004[1] = πg.Param{Name: "startpos", Def: nil}
			πTemp004[2] = πg.Param{Name: "pos", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("_siftdown", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µheap *πg.Object = πArgs[0]; _ = µheap
				var µstartpos *πg.Object = πArgs[1]; _ = µstartpos
				var µpos *πg.Object = πArgs[2]; _ = µpos
				var µnewitem *πg.Object = πg.UnboundLocal; _ = µnewitem
				var µparentpos *πg.Object = πg.UnboundLocal; _ = µparentpos
				var µparent *πg.Object = πg.UnboundLocal; _ = µparent
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 250: newitem = heap[pos]
					πF.SetLineno(250)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp001 = µpos
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µheap, πTemp001); πE != nil {
						continue
					}
					µnewitem = πTemp002
					// line 253: while pos > startpos:
					πF.SetLineno(253)
					πF.PushCheckpoint(2)
					πTemp003 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstartpos, "startpos"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µpos, µstartpos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 254: parentpos = (pos - 1) >> 1
					πF.SetLineno(254)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µpos, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.RShift(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µparentpos = πTemp001
					// line 255: parent = heap[parentpos]
					πF.SetLineno(255)
					if πE = πg.CheckLocal(πF, µparentpos, "parentpos"); πE != nil {
						continue
					}
					πTemp001 = µparentpos
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µheap, πTemp001); πE != nil {
						continue
					}
					µparent = πTemp002
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µnewitem, "newitem"); πE != nil {
						continue
					}
					πTemp005[0] = µnewitem
					if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
						continue
					}
					πTemp005[1] = µparent
					if πTemp001, πE = πg.ResolveGlobal(πF, ßcmp_lt); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 256: if cmp_lt(newitem, parent):
					πF.SetLineno(256)
				Label4:
					// line 257: heap[pos] = parent
					πF.SetLineno(257)
					if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µparent); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp002 = µpos
					if πE = πg.SetItem(πF, µheap, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 258: pos = parentpos
					πF.SetLineno(258)
					if πE = πg.CheckLocal(πF, µparentpos, "parentpos"); πE != nil {
						continue
					}
					µpos = µparentpos
					// line 259: continue
					πF.SetLineno(259)
					continue
					goto Label5
				Label5:
					// line 260: break
					πF.SetLineno(260)
					πTemp003 = true
					continue
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 261: heap[pos] = newitem
					πF.SetLineno(261)
					if πE = πg.CheckLocal(πF, µnewitem, "newitem"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnewitem); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp002 = µpos
					if πE = πg.SetItem(πF, µheap, πTemp002, πTemp001); πE != nil {
						continue
					}
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_siftdown.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 302: def _siftup(heap, pos):
			πF.SetLineno(302)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "heap", Def: nil}
			πTemp004[1] = πg.Param{Name: "pos", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("_siftup", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µheap *πg.Object = πArgs[0]; _ = µheap
				var µpos *πg.Object = πArgs[1]; _ = µpos
				var µendpos *πg.Object = πg.UnboundLocal; _ = µendpos
				var µstartpos *πg.Object = πg.UnboundLocal; _ = µstartpos
				var µnewitem *πg.Object = πg.UnboundLocal; _ = µnewitem
				var µchildpos *πg.Object = πg.UnboundLocal; _ = µchildpos
				var µrightpos *πg.Object = πg.UnboundLocal; _ = µrightpos
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 303: endpos = len(heap)
					πF.SetLineno(303)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp001[0] = µheap
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µendpos = πTemp003
					// line 304: startpos = pos
					πF.SetLineno(304)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					µstartpos = µpos
					// line 305: newitem = heap[pos]
					πF.SetLineno(305)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp002 = µpos
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µheap, πTemp002); πE != nil {
						continue
					}
					µnewitem = πTemp003
					// line 307: childpos = 2*pos + 1    # leftmost child position
					πF.SetLineno(307)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewInt(2).ToObject(), µpos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µchildpos = πTemp002
					// line 308: while childpos < endpos:
					πF.SetLineno(308)
					πF.PushCheckpoint(2)
					πTemp004 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µchildpos, "childpos"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µendpos, "endpos"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µchildpos, µendpos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 310: rightpos = childpos + 1
					πF.SetLineno(310)
					if πE = πg.CheckLocal(πF, µchildpos, "childpos"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µchildpos, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µrightpos = πTemp002
					if πE = πg.CheckLocal(πF, µrightpos, "rightpos"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µendpos, "endpos"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µrightpos, µendpos); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label4
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µchildpos, "childpos"); πE != nil {
						continue
					}
					πTemp006 = µchildpos
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µheap, πTemp006); πE != nil {
						continue
					}
					πTemp001[0] = πTemp007
					if πE = πg.CheckLocal(πF, µrightpos, "rightpos"); πE != nil {
						continue
					}
					πTemp006 = µrightpos
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µheap, πTemp006); πE != nil {
						continue
					}
					πTemp001[1] = πTemp007
					if πTemp006, πE = πg.ResolveGlobal(πF, ßcmp_lt); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp008).ToObject()
					πTemp002 = πTemp003
				Label4:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label5
					}
					goto Label6
					// line 311: if rightpos < endpos and not cmp_lt(heap[childpos], heap[rightpos]):
					πF.SetLineno(311)
				Label5:
					// line 312: childpos = rightpos
					πF.SetLineno(312)
					if πE = πg.CheckLocal(πF, µrightpos, "rightpos"); πE != nil {
						continue
					}
					µchildpos = µrightpos
					goto Label6
				Label6:
					// line 314: heap[pos] = heap[childpos]
					πF.SetLineno(314)
					if πE = πg.CheckLocal(πF, µchildpos, "childpos"); πE != nil {
						continue
					}
					πTemp002 = µchildpos
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µheap, πTemp002); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp006 = µpos
					if πE = πg.SetItem(πF, µheap, πTemp006, πTemp002); πE != nil {
						continue
					}
					// line 315: pos = childpos
					πF.SetLineno(315)
					if πE = πg.CheckLocal(πF, µchildpos, "childpos"); πE != nil {
						continue
					}
					µpos = µchildpos
					// line 316: childpos = 2*pos + 1
					πF.SetLineno(316)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewInt(2).ToObject(), µpos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µchildpos = πTemp002
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 319: heap[pos] = newitem
					πF.SetLineno(319)
					if πE = πg.CheckLocal(πF, µnewitem, "newitem"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µnewitem); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp003 = µpos
					if πE = πg.SetItem(πF, µheap, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 320: _siftdown(heap, startpos, pos)
					πF.SetLineno(320)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp001[0] = µheap
					if πE = πg.CheckLocal(πF, µstartpos, "startpos"); πE != nil {
						continue
					}
					πTemp001[1] = µstartpos
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp001[2] = µpos
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_siftdown); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_siftup.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 322: def _siftdown_max(heap, startpos, pos):
			πF.SetLineno(322)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "heap", Def: nil}
			πTemp004[1] = πg.Param{Name: "startpos", Def: nil}
			πTemp004[2] = πg.Param{Name: "pos", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("_siftdown_max", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µheap *πg.Object = πArgs[0]; _ = µheap
				var µstartpos *πg.Object = πArgs[1]; _ = µstartpos
				var µpos *πg.Object = πArgs[2]; _ = µpos
				var µnewitem *πg.Object = πg.UnboundLocal; _ = µnewitem
				var µparentpos *πg.Object = πg.UnboundLocal; _ = µparentpos
				var µparent *πg.Object = πg.UnboundLocal; _ = µparent
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 323: 'Maxheap variant of _siftdown'
					πF.SetLineno(323)
					// line 324: newitem = heap[pos]
					πF.SetLineno(324)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp001 = µpos
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µheap, πTemp001); πE != nil {
						continue
					}
					µnewitem = πTemp002
					// line 327: while pos > startpos:
					πF.SetLineno(327)
					πF.PushCheckpoint(2)
					πTemp003 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstartpos, "startpos"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µpos, µstartpos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 328: parentpos = (pos - 1) >> 1
					πF.SetLineno(328)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µpos, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.RShift(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µparentpos = πTemp001
					// line 329: parent = heap[parentpos]
					πF.SetLineno(329)
					if πE = πg.CheckLocal(πF, µparentpos, "parentpos"); πE != nil {
						continue
					}
					πTemp001 = µparentpos
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µheap, πTemp001); πE != nil {
						continue
					}
					µparent = πTemp002
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
						continue
					}
					πTemp005[0] = µparent
					if πE = πg.CheckLocal(πF, µnewitem, "newitem"); πE != nil {
						continue
					}
					πTemp005[1] = µnewitem
					if πTemp001, πE = πg.ResolveGlobal(πF, ßcmp_lt); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 330: if cmp_lt(parent, newitem):
					πF.SetLineno(330)
				Label4:
					// line 331: heap[pos] = parent
					πF.SetLineno(331)
					if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µparent); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp002 = µpos
					if πE = πg.SetItem(πF, µheap, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 332: pos = parentpos
					πF.SetLineno(332)
					if πE = πg.CheckLocal(πF, µparentpos, "parentpos"); πE != nil {
						continue
					}
					µpos = µparentpos
					// line 333: continue
					πF.SetLineno(333)
					continue
					goto Label5
				Label5:
					// line 334: break
					πF.SetLineno(334)
					πTemp003 = true
					continue
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 335: heap[pos] = newitem
					πF.SetLineno(335)
					if πE = πg.CheckLocal(πF, µnewitem, "newitem"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnewitem); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp002 = µpos
					if πE = πg.SetItem(πF, µheap, πTemp002, πTemp001); πE != nil {
						continue
					}
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_siftdown_max.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 337: def _siftup_max(heap, pos):
			πF.SetLineno(337)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "heap", Def: nil}
			πTemp004[1] = πg.Param{Name: "pos", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("_siftup_max", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µheap *πg.Object = πArgs[0]; _ = µheap
				var µpos *πg.Object = πArgs[1]; _ = µpos
				var µendpos *πg.Object = πg.UnboundLocal; _ = µendpos
				var µstartpos *πg.Object = πg.UnboundLocal; _ = µstartpos
				var µnewitem *πg.Object = πg.UnboundLocal; _ = µnewitem
				var µchildpos *πg.Object = πg.UnboundLocal; _ = µchildpos
				var µrightpos *πg.Object = πg.UnboundLocal; _ = µrightpos
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 338: 'Maxheap variant of _siftup'
					πF.SetLineno(338)
					// line 339: endpos = len(heap)
					πF.SetLineno(339)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp001[0] = µheap
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µendpos = πTemp003
					// line 340: startpos = pos
					πF.SetLineno(340)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					µstartpos = µpos
					// line 341: newitem = heap[pos]
					πF.SetLineno(341)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp002 = µpos
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µheap, πTemp002); πE != nil {
						continue
					}
					µnewitem = πTemp003
					// line 343: childpos = 2*pos + 1    # leftmost child position
					πF.SetLineno(343)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewInt(2).ToObject(), µpos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µchildpos = πTemp002
					// line 344: while childpos < endpos:
					πF.SetLineno(344)
					πF.PushCheckpoint(2)
					πTemp004 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µchildpos, "childpos"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µendpos, "endpos"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µchildpos, µendpos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 346: rightpos = childpos + 1
					πF.SetLineno(346)
					if πE = πg.CheckLocal(πF, µchildpos, "childpos"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µchildpos, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µrightpos = πTemp002
					if πE = πg.CheckLocal(πF, µrightpos, "rightpos"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µendpos, "endpos"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µrightpos, µendpos); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label4
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrightpos, "rightpos"); πE != nil {
						continue
					}
					πTemp006 = µrightpos
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µheap, πTemp006); πE != nil {
						continue
					}
					πTemp001[0] = πTemp007
					if πE = πg.CheckLocal(πF, µchildpos, "childpos"); πE != nil {
						continue
					}
					πTemp006 = µchildpos
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µheap, πTemp006); πE != nil {
						continue
					}
					πTemp001[1] = πTemp007
					if πTemp006, πE = πg.ResolveGlobal(πF, ßcmp_lt); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp008).ToObject()
					πTemp002 = πTemp003
				Label4:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label5
					}
					goto Label6
					// line 347: if rightpos < endpos and not cmp_lt(heap[rightpos], heap[childpos]):
					πF.SetLineno(347)
				Label5:
					// line 348: childpos = rightpos
					πF.SetLineno(348)
					if πE = πg.CheckLocal(πF, µrightpos, "rightpos"); πE != nil {
						continue
					}
					µchildpos = µrightpos
					goto Label6
				Label6:
					// line 350: heap[pos] = heap[childpos]
					πF.SetLineno(350)
					if πE = πg.CheckLocal(πF, µchildpos, "childpos"); πE != nil {
						continue
					}
					πTemp002 = µchildpos
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µheap, πTemp002); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp006 = µpos
					if πE = πg.SetItem(πF, µheap, πTemp006, πTemp002); πE != nil {
						continue
					}
					// line 351: pos = childpos
					πF.SetLineno(351)
					if πE = πg.CheckLocal(πF, µchildpos, "childpos"); πE != nil {
						continue
					}
					µpos = µchildpos
					// line 352: childpos = 2*pos + 1
					πF.SetLineno(352)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewInt(2).ToObject(), µpos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µchildpos = πTemp002
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 355: heap[pos] = newitem
					πF.SetLineno(355)
					if πE = πg.CheckLocal(πF, µnewitem, "newitem"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µnewitem); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp003 = µpos
					if πE = πg.SetItem(πF, µheap, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 356: _siftdown_max(heap, startpos, pos)
					πF.SetLineno(356)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µheap, "heap"); πE != nil {
						continue
					}
					πTemp001[0] = µheap
					if πE = πg.CheckLocal(πF, µstartpos, "startpos"); πE != nil {
						continue
					}
					πTemp001[1] = µstartpos
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp001[2] = µpos
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_siftdown_max); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_siftup_max.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 364: def merge(*iterables):
			πF.SetLineno(364)
			πTemp004 = make([]πg.Param, 0)
			πTemp017 = πg.NewFunction(πg.NewCode("merge", "build/src/__python__/heapq.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µiterables *πg.Object = πArgs[0]; _ = µiterables
				var µ_heappop *πg.Object = πg.UnboundLocal; _ = µ_heappop
				var µ_heapreplace *πg.Object = πg.UnboundLocal; _ = µ_heapreplace
				var µ_StopIteration *πg.Object = πg.UnboundLocal; _ = µ_StopIteration
				var µ_len *πg.Object = πg.UnboundLocal; _ = µ_len
				var µh *πg.Object = πg.UnboundLocal; _ = µh
				var µh_append *πg.Object = πg.UnboundLocal; _ = µh_append
				var µitnum *πg.Object = πg.UnboundLocal; _ = µitnum
				var µit *πg.Object = πg.UnboundLocal; _ = µit
				var µnext *πg.Object = πg.UnboundLocal; _ = µnext
				var µv *πg.Object = πg.UnboundLocal; _ = µv
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.BaseException
				_ = πTemp009
				var πTemp010 *πg.Traceback
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						case 2: goto Label2
						case 5: goto Label5
						case 7: goto Label7
						case 8: goto Label8
						case 11: goto Label11
						case 12: goto Label12
						case 13: goto Label13
						case 15: goto Label15
						case 19: goto Label19
						case 20: goto Label20
						case 21: goto Label21
						case 23: goto Label23
						default: panic("unexpected function state")
						}
						// line 365: '''Merge multiple sorted inputs into a single sorted output.
						πF.SetLineno(365)
						// line 375: _heappop, _heapreplace, _StopIteration = heappop, heapreplace, StopIteration
						πF.SetLineno(375)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßheappop); πE != nil {
							continue
						}
						if πTemp003, πE = πg.ResolveGlobal(πF, ßheapreplace); πE != nil {
							continue
						}
						if πTemp004, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
							continue
						}
						πTemp001 = πg.NewTuple3(πTemp002, πTemp003, πTemp004).ToObject()
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
							continue
						}
						µ_heappop = πTemp002
						µ_heapreplace = πTemp003
						µ_StopIteration = πTemp004
						// line 376: _len = len
						πF.SetLineno(376)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
							continue
						}
						µ_len = πTemp001
						// line 378: h = []
						πF.SetLineno(378)
						πTemp005 = make([]*πg.Object, 0)
						πTemp001 = πg.NewList(πTemp005...).ToObject()
						µh = πTemp001
						// line 379: h_append = h.append
						πF.SetLineno(379)
						if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.GetAttr(πF, µh, ßappend, nil); πE != nil {
							continue
						}
						µh_append = πTemp001
						πTemp005 = πF.MakeArgs(1)
						πTemp006 = πF.MakeArgs(2)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
							continue
						}
						πTemp006[0] = πTemp002
						if πE = πg.CheckLocal(πF, µiterables, "iterables"); πE != nil {
							continue
						}
						πTemp006[1] = µiterables
						if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp006)
						πTemp005[0] = πTemp003
						if πTemp002, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
							continue
						}
						πF.PushCheckpoint(2)
						πTemp007 = false
					Label1:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp007 {
							πF.PopCheckpoint()
							goto Label3
						}
						if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp008 = !isStop
						} else {
							πTemp008 = true
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
								continue
							}
							µitnum = πTemp003
							µit = πTemp004
						}
						if πE != nil || !πTemp008 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 381: try:
						πF.SetLineno(381)
						πF.PushCheckpoint(5)
						// line 382: next = it.next
						πF.SetLineno(382)
						if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µit, ßnext, nil); πE != nil {
							continue
						}
						µnext = πTemp002
						// line 383: h_append([next(), itnum, next])
						πF.SetLineno(383)
						πTemp005 = πF.MakeArgs(1)
						πTemp006 = make([]*πg.Object, 3)
						if πE = πg.CheckLocal(πF, µnext, "next"); πE != nil {
							continue
						}
						if πTemp002, πE = µnext.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp006[0] = πTemp002
						if πE = πg.CheckLocal(πF, µitnum, "itnum"); πE != nil {
							continue
						}
						πTemp006[1] = µitnum
						if πE = πg.CheckLocal(πF, µnext, "next"); πE != nil {
							continue
						}
						πTemp006[2] = µnext
						πTemp002 = πg.NewList(πTemp006...).ToObject()
						πTemp005[0] = πTemp002
						if πE = πg.CheckLocal(πF, µh_append, "h_append"); πE != nil {
							continue
						}
						if πTemp002, πE = µh_append.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						πF.PopCheckpoint()
						goto Label4
					Label5:
						if πE == nil {
						  continue
						}
						πE = nil
						πTemp009, πTemp010 = πF.ExcInfo()
						if πE = πg.CheckLocal(πF, µ_StopIteration, "_StopIteration"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.IsInstance(πF, πTemp009.ToObject(), µ_StopIteration); πE != nil {
							continue
						}
						if πTemp008 {
							goto Label6
						}
						πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
						continue
						// line 384: except _StopIteration:
						πF.SetLineno(384)
					Label6:
						// line 385: pass
						πF.SetLineno(385)
						πF.RestoreExc(nil, nil)
						goto Label4
					Label4:
						continue
					Label2:
						if πE != nil || πR != nil {
							continue
						}
					Label3:
						// line 386: heapify(h)
						πF.SetLineno(386)
						πTemp005 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
							continue
						}
						πTemp005[0] = µh
						if πTemp001, πE = πg.ResolveGlobal(πF, ßheapify); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						// line 388: while _len(h) > 1:
						πF.SetLineno(388)
						πF.PushCheckpoint(8)
						πTemp007 = false
					Label7:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp007 {
							πF.PopCheckpoint()
							goto Label9
						}
						πTemp005 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
							continue
						}
						πTemp005[0] = µh
						if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
							continue
						}
						if πTemp002, πE = µ_len.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						if πTemp001, πE = πg.GT(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πE != nil || !πTemp008 {
							continue
						}
						πF.PushCheckpoint(7)            
						// line 389: try:
						πF.SetLineno(389)
						πF.PushCheckpoint(11)
						// line 390: while 1:
						πF.SetLineno(390)
						πF.PushCheckpoint(13)
						πTemp008 = false
					Label12:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp008 {
							πF.PopCheckpoint()
							goto Label14
						}
						if πTemp011, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						if πE != nil || !πTemp011 {
							continue
						}
						πF.PushCheckpoint(12)            
						// line 391: v, itnum, next = s = h[0]
						πF.SetLineno(391)
						πTemp001 = πg.NewInt(0).ToObject()
						if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetItem(πF, µh, πTemp001); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
							continue
						}
						µv = πTemp001
						µitnum = πTemp003
						µnext = πTemp004
						µs = πTemp002
						// line 392: yield v
						πF.SetLineno(392)
						if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
							continue
						}
						πF.PushCheckpoint(15)
						return µv, nil
					Label15:
						πTemp001 = πSent
						// line 393: s[0] = next()               # raises StopIteration when exhausted
						πF.SetLineno(393)
						if πE = πg.CheckLocal(πF, µnext, "next"); πE != nil {
							continue
						}
						if πTemp001, πE = µnext.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
							continue
						}
						πTemp003 = πg.NewInt(0).ToObject()
						if πE = πg.SetItem(πF, µs, πTemp003, πTemp002); πE != nil {
							continue
						}
						// line 394: _heapreplace(h, s)          # restore heap condition
						πF.SetLineno(394)
						πTemp005 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
							continue
						}
						πTemp005[0] = µh
						if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
							continue
						}
						πTemp005[1] = µs
						if πE = πg.CheckLocal(πF, µ_heapreplace, "_heapreplace"); πE != nil {
							continue
						}
						if πTemp001, πE = µ_heapreplace.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						continue
					Label13:
						if πE != nil || πR != nil {
							continue
						}
					Label14:
						πF.PopCheckpoint()
						goto Label10
					Label11:
						if πE == nil {
						  continue
						}
						πE = nil
						πTemp009, πTemp010 = πF.ExcInfo()
						if πE = πg.CheckLocal(πF, µ_StopIteration, "_StopIteration"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.IsInstance(πF, πTemp009.ToObject(), µ_StopIteration); πE != nil {
							continue
						}
						if πTemp008 {
							goto Label16
						}
						πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
						continue
						// line 395: except _StopIteration:
						πF.SetLineno(395)
					Label16:
						// line 396: _heappop(h)                     # remove empty iterator
						πF.SetLineno(396)
						πTemp005 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
							continue
						}
						πTemp005[0] = µh
						if πE = πg.CheckLocal(πF, µ_heappop, "_heappop"); πE != nil {
							continue
						}
						if πTemp001, πE = µ_heappop.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						πF.RestoreExc(nil, nil)
						goto Label10
					Label10:
						continue
					Label8:
						if πE != nil || πR != nil {
							continue
						}
					Label9:
						if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
							continue
						}
						if πTemp007, πE = πg.IsTrue(πF, µh); πE != nil {
							continue
						}
						if πTemp007 {
							goto Label17
						}
						goto Label18
						// line 397: if h:
						πF.SetLineno(397)
					Label17:
						// line 399: v, itnum, next = h[0]
						πF.SetLineno(399)
						πTemp001 = πg.NewInt(0).ToObject()
						if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetItem(πF, µh, πTemp001); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
							continue
						}
						µv = πTemp001
						µitnum = πTemp003
						µnext = πTemp004
						// line 400: yield v
						πF.SetLineno(400)
						if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
							continue
						}
						πF.PushCheckpoint(19)
						return µv, nil
					Label19:
						πTemp001 = πSent
						if πE = πg.CheckLocal(πF, µnext, "next"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µnext, ß__self__, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
							continue
						}
						πF.PushCheckpoint(21)
						πTemp007 = false
					Label20:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp007 {
							πF.PopCheckpoint()
							goto Label22
						}
						if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp008 = !isStop
						} else {
							πTemp008 = true
							µv = πTemp002
						}
						if πE != nil || !πTemp008 {
							continue
						}
						πF.PushCheckpoint(20)            
						// line 402: yield v
						πF.SetLineno(402)
						if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
							continue
						}
						πF.PushCheckpoint(23)
						return µv, nil
					Label23:
						πTemp002 = πSent
						continue
					Label21:
						if πE != nil || πR != nil {
							continue
						}
					Label22:
						goto Label18
					Label18:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßmerge.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 405: _nsmallest = nsmallest
			πF.SetLineno(405)
			if πTemp018, πE = πg.ResolveGlobal(πF, ßnsmallest); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_nsmallest.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 406: def nsmallest(n, iterable, key=None):
			πF.SetLineno(406)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "n", Def: nil}
			πTemp004[1] = πg.Param{Name: "iterable", Def: nil}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "key", Def: πTemp019}
			πTemp018 = πg.NewFunction(πg.NewCode("nsmallest", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
				var µiterable *πg.Object = πArgs[1]; _ = µiterable
				var µkey *πg.Object = πArgs[2]; _ = µkey
				var µit *πg.Object = πg.UnboundLocal; _ = µit
				var µhead *πg.Object = πg.UnboundLocal; _ = µhead
				var µsize *πg.Object = πg.UnboundLocal; _ = µsize
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µin1 *πg.Object = πg.UnboundLocal; _ = µin1
				var µin2 *πg.Object = πg.UnboundLocal; _ = µin2
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.BaseException
				_ = πTemp010
				var πTemp011 *πg.Traceback
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8: goto Label8
					default: panic("unexpected function state")
					}
					// line 407: """Find the n smallest elements in a dataset.
					πF.SetLineno(407)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 412: if n == 1:
					πF.SetLineno(412)
				Label1:
					// line 413: it = iter(iterable)
					πF.SetLineno(413)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp003[0] = µiterable
					if πTemp001, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µit = πTemp004
					// line 414: head = list(islice(it, 1))
					πF.SetLineno(414)
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp005[0] = µit
					πTemp005[1] = πg.NewInt(1).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßislice); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µhead = πTemp004
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µhead); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 415: if not head:
					πF.SetLineno(415)
				Label3:
					// line 416: return []
					πF.SetLineno(416)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µkey == πTemp004).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					goto Label6
					// line 417: if key is None:
					πF.SetLineno(417)
				Label5:
					// line 418: return [min(chain(head, it))]
					πF.SetLineno(418)
					πTemp003 = make([]*πg.Object, 1)
					πTemp005 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp006[0] = µhead
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp006[1] = µit
					if πTemp001, πE = πg.ResolveGlobal(πF, ßchain); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[0] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label6
				Label6:
					// line 419: return [min(chain(head, it), key=key)]
					πF.SetLineno(419)
					πTemp003 = make([]*πg.Object, 1)
					πTemp005 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp006[0] = µhead
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp006[1] = µit
					if πTemp001, πE = πg.ResolveGlobal(πF, ßchain); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[0] = πTemp004
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"key", µkey},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 422: try:
					πF.SetLineno(422)
					πF.PushCheckpoint(8)
					// line 423: size = len(iterable)
					πF.SetLineno(423)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp003[0] = µiterable
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µsize = πTemp004
					πF.PopCheckpoint()
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GE(πF, µn, µsize); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 427: if n >= size:
					πF.SetLineno(427)
				Label9:
					// line 428: return sorted(iterable, key=key)[:n]
					πF.SetLineno(428)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µn, πg.None}, nil); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp003[0] = µiterable
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"key", µkey},
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp003, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.GetItem(πF, πTemp009, πTemp001); πE != nil {
						continue
					}
					πR = πTemp004
					continue
					goto Label10
				Label10:
					goto Label7
				Label8:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp010, πTemp011 = πF.ExcInfo()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp004, πTemp008).ToObject()
					if πTemp002, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label11
					}
					πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
					continue
					// line 424: except (TypeError, AttributeError):
					πF.SetLineno(424)
				Label11:
					// line 425: pass
					πF.SetLineno(425)
					πF.RestoreExc(nil, nil)
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µkey == πTemp004).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					goto Label13
					// line 431: if key is None:
					πF.SetLineno(431)
				Label12:
					// line 432: it = izip(iterable, count())                        # decorate
					πF.SetLineno(432)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp003[0] = µiterable
					if πTemp001, πE = πg.ResolveGlobal(πF, ßcount); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßizip); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µit = πTemp004
					// line 433: result = _nsmallest(n, it)
					πF.SetLineno(433)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp003[0] = µn
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp003[1] = µit
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_nsmallest); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µresult = πTemp004
					// line 434: return map(itemgetter(0), result)                   # undecorate
					πF.SetLineno(434)
					πTemp003 = πF.MakeArgs(2)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewInt(0).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßitemgetter); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003[1] = µresult
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label13
				Label13:
					// line 437: in1, in2 = tee(iterable)
					πF.SetLineno(437)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp003[0] = µiterable
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtee); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
						continue
					}
					µin1 = πTemp001
					µin2 = πTemp008
					// line 438: it = izip(imap(key, in1), count(), in2)                 # decorate
					πF.SetLineno(438)
					πTemp003 = πF.MakeArgs(3)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp005[0] = µkey
					if πE = πg.CheckLocal(πF, µin1, "in1"); πE != nil {
						continue
					}
					πTemp005[1] = µin1
					if πTemp001, πE = πg.ResolveGlobal(πF, ßimap); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßcount); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
					if πE = πg.CheckLocal(πF, µin2, "in2"); πE != nil {
						continue
					}
					πTemp003[2] = µin2
					if πTemp001, πE = πg.ResolveGlobal(πF, ßizip); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µit = πTemp004
					// line 439: result = _nsmallest(n, it)
					πF.SetLineno(439)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp003[0] = µn
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp003[1] = µit
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_nsmallest); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µresult = πTemp004
					// line 440: return map(itemgetter(2), result)                       # undecorate
					πF.SetLineno(440)
					πTemp003 = πF.MakeArgs(2)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewInt(2).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßitemgetter); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003[1] = µresult
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnsmallest.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 442: _nlargest = nlargest
			πF.SetLineno(442)
			if πTemp019, πE = πg.ResolveGlobal(πF, ßnlargest); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_nlargest.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 443: def nlargest(n, iterable, key=None):
			πF.SetLineno(443)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "n", Def: nil}
			πTemp004[1] = πg.Param{Name: "iterable", Def: nil}
			if πTemp020, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "key", Def: πTemp020}
			πTemp019 = πg.NewFunction(πg.NewCode("nlargest", "build/src/__python__/heapq.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
				var µiterable *πg.Object = πArgs[1]; _ = µiterable
				var µkey *πg.Object = πArgs[2]; _ = µkey
				var µit *πg.Object = πg.UnboundLocal; _ = µit
				var µhead *πg.Object = πg.UnboundLocal; _ = µhead
				var µsize *πg.Object = πg.UnboundLocal; _ = µsize
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µin1 *πg.Object = πg.UnboundLocal; _ = µin1
				var µin2 *πg.Object = πg.UnboundLocal; _ = µin2
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.BaseException
				_ = πTemp010
				var πTemp011 *πg.Traceback
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8: goto Label8
					default: panic("unexpected function state")
					}
					// line 444: """Find the n largest elements in a dataset.
					πF.SetLineno(444)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 450: if n == 1:
					πF.SetLineno(450)
				Label1:
					// line 451: it = iter(iterable)
					πF.SetLineno(451)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp003[0] = µiterable
					if πTemp001, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µit = πTemp004
					// line 452: head = list(islice(it, 1))
					πF.SetLineno(452)
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp005[0] = µit
					πTemp005[1] = πg.NewInt(1).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßislice); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µhead = πTemp004
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µhead); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 453: if not head:
					πF.SetLineno(453)
				Label3:
					// line 454: return []
					πF.SetLineno(454)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µkey == πTemp004).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					goto Label6
					// line 455: if key is None:
					πF.SetLineno(455)
				Label5:
					// line 456: return [max(chain(head, it))]
					πF.SetLineno(456)
					πTemp003 = make([]*πg.Object, 1)
					πTemp005 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp006[0] = µhead
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp006[1] = µit
					if πTemp001, πE = πg.ResolveGlobal(πF, ßchain); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[0] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label6
				Label6:
					// line 457: return [max(chain(head, it), key=key)]
					πF.SetLineno(457)
					πTemp003 = make([]*πg.Object, 1)
					πTemp005 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp006[0] = µhead
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp006[1] = µit
					if πTemp001, πE = πg.ResolveGlobal(πF, ßchain); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[0] = πTemp004
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"key", µkey},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 460: try:
					πF.SetLineno(460)
					πF.PushCheckpoint(8)
					// line 461: size = len(iterable)
					πF.SetLineno(461)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp003[0] = µiterable
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µsize = πTemp004
					πF.PopCheckpoint()
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GE(πF, µn, µsize); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 465: if n >= size:
					πF.SetLineno(465)
				Label9:
					// line 466: return sorted(iterable, key=key, reverse=True)[:n]
					πF.SetLineno(466)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µn, πg.None}, nil); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp003[0] = µiterable
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"key", µkey},
						{"reverse", πTemp008},
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp003, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.GetItem(πF, πTemp009, πTemp001); πE != nil {
						continue
					}
					πR = πTemp004
					continue
					goto Label10
				Label10:
					goto Label7
				Label8:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp010, πTemp011 = πF.ExcInfo()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp004, πTemp008).ToObject()
					if πTemp002, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label11
					}
					πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
					continue
					// line 462: except (TypeError, AttributeError):
					πF.SetLineno(462)
				Label11:
					// line 463: pass
					πF.SetLineno(463)
					πF.RestoreExc(nil, nil)
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µkey == πTemp004).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					goto Label13
					// line 469: if key is None:
					πF.SetLineno(469)
				Label12:
					// line 470: it = izip(iterable, count(0,-1))                    # decorate
					πF.SetLineno(470)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp003[0] = µiterable
					πTemp005 = πF.MakeArgs(2)
					πTemp005[0] = πg.NewInt(0).ToObject()
					if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßcount); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[1] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßizip); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µit = πTemp004
					// line 471: result = _nlargest(n, it)
					πF.SetLineno(471)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp003[0] = µn
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp003[1] = µit
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_nlargest); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µresult = πTemp004
					// line 472: return map(itemgetter(0), result)                   # undecorate
					πF.SetLineno(472)
					πTemp003 = πF.MakeArgs(2)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewInt(0).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßitemgetter); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003[1] = µresult
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label13
				Label13:
					// line 475: in1, in2 = tee(iterable)
					πF.SetLineno(475)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp003[0] = µiterable
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtee); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
						continue
					}
					µin1 = πTemp001
					µin2 = πTemp008
					// line 476: it = izip(imap(key, in1), count(0,-1), in2)             # decorate
					πF.SetLineno(476)
					πTemp003 = πF.MakeArgs(3)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp005[0] = µkey
					if πE = πg.CheckLocal(πF, µin1, "in1"); πE != nil {
						continue
					}
					πTemp005[1] = µin1
					if πTemp001, πE = πg.ResolveGlobal(πF, ßimap); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					πTemp005 = πF.MakeArgs(2)
					πTemp005[0] = πg.NewInt(0).ToObject()
					if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßcount); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[1] = πTemp004
					if πE = πg.CheckLocal(πF, µin2, "in2"); πE != nil {
						continue
					}
					πTemp003[2] = µin2
					if πTemp001, πE = πg.ResolveGlobal(πF, ßizip); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µit = πTemp004
					// line 477: result = _nlargest(n, it)
					πF.SetLineno(477)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp003[0] = µn
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp003[1] = µit
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_nlargest); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µresult = πTemp004
					// line 478: return map(itemgetter(2), result)                       # undecorate
					πF.SetLineno(478)
					πTemp003 = πF.MakeArgs(2)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewInt(2).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßitemgetter); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003[1] = µresult
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnlargest.ToObject(), πTemp019); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("heapq", Code)
}
