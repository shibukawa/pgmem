package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	_ "unsafe"
)
//go:linkname F_brin_build_desc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_brin_build_desc
func F_brin_build_desc(m *base.Module, l0 int32) int32
//go:linkname F__brin_end_parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__brin_end_parallel
func F__brin_end_parallel(m *base.Module, l0 int32)
//go:linkname F__brin_parallel_scan_and_build github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__brin_parallel_scan_and_build
func F__brin_parallel_scan_and_build(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_union_tuples github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_union_tuples
func F_union_tuples(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_brin_fill_empty_ranges github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_fill_empty_ranges
func F_brin_fill_empty_ranges(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_brin_free_desc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_free_desc
func F_brin_free_desc(m *base.Module, l0 int32)
//go:linkname F_minmax_get_strategy_procinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_minmax_get_strategy_procinfo
func F_minmax_get_strategy_procinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_compare_expanded_ranges github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_compare_expanded_ranges
func F_compare_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_brin_doinsert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_doinsert
func F_brin_doinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_brinRevmapInitialize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_brinRevmapInitialize
func F_brinRevmapInitialize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_brinRevmapTerminate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_brinRevmapTerminate
func F_brinRevmapTerminate(m *base.Module, l0 int32)
//go:linkname F_brin_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_brin_form_tuple
func F_brin_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_brin_new_memtuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_new_memtuple
func F_brin_new_memtuple(m *base.Module, l0 int32) int32
//go:linkname F_brin_deform_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_brin_deform_tuple
func F_brin_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_attrmap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_attrmap
func F_make_attrmap(m *base.Module, l0 int32) int32
//go:linkname F_free_attrmap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_free_attrmap
func F_free_attrmap(m *base.Module, l0 int32)
//go:linkname F_build_attrmap_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_attrmap_by_name
func F_build_attrmap_by_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_attrmap_by_name_if_req github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_build_attrmap_by_name_if_req
func F_build_attrmap_by_name_if_req(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mask_unused_space github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mask_unused_space
func F_mask_unused_space(m *base.Module, l0 int32)
//go:linkname F_detoast_external_attr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_detoast_external_attr
func F_detoast_external_attr(m *base.Module, l0 int32) int32
//go:linkname F_detoast_attr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_detoast_attr
func F_detoast_attr(m *base.Module, l0 int32) int32
//go:linkname F_detoast_attr_slice github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_detoast_attr_slice
func F_detoast_attr_slice(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_toast_raw_datum_size github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_toast_raw_datum_size
func F_toast_raw_datum_size(m *base.Module, l0 int32) int32
//go:linkname F_getmissingattr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getmissingattr
func F_getmissingattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_compute_data_size github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_compute_data_size
func F_heap_compute_data_size(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_nocachegetattr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_nocachegetattr
func F_nocachegetattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_getsysattr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_getsysattr
func F_heap_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_copytuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_copytuple
func F_heap_copytuple(m *base.Module, l0 int32) int32
//go:linkname F_heap_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_form_tuple
func F_heap_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_modify_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_modify_tuple
func F_heap_modify_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_heap_deform_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_deform_tuple
func F_heap_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_heap_form_minimal_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_form_minimal_tuple
func F_heap_form_minimal_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_index_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_form_tuple
func F_index_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_nocache_index_getattr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_nocache_index_getattr
func F_nocache_index_getattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_index_deform_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_deform_tuple
func F_index_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CopyIndexTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CopyIndexTuple
func F_CopyIndexTuple(m *base.Module, l0 int32) int32
//go:linkname F_relation_open github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_relation_open
func F_relation_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_relation_openrv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_relation_openrv
func F_relation_openrv(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_relation_openrv_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_relation_openrv_extended
func F_relation_openrv_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_add_reloption_kind github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_reloption_kind
func F_add_reloption_kind(m *base.Module) int32
//go:linkname F_add_int_reloption github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_int_reloption
func F_add_int_reloption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_add_local_int_reloption github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_local_int_reloption
func F_add_local_int_reloption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_transformRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_transformRelOptions
func F_transformRelOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_build_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_reloptions
func F_build_reloptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_index_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_reloptions
func F_index_reloptions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_initialize_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_initialize_reloptions
func F_initialize_reloptions(m *base.Module)
//go:linkname F_parseRelOptionsInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parseRelOptionsInternal
func F_parseRelOptionsInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_fillRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fillRelOptions
func F_fillRelOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_tablespace_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tablespace_reloptions
func F_tablespace_reloptions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ScanKeyInit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ScanKeyInit
func F_ScanKeyInit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ss_report_location github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ss_report_location
func F_ss_report_location(m *base.Module, l0 int32, l1 int32)
//go:linkname F_TidStoreCreateShared github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TidStoreCreateShared
func F_TidStoreCreateShared(m *base.Module, l0 int32) int32
//go:linkname F_TidStoreDestroy github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TidStoreDestroy
func F_TidStoreDestroy(m *base.Module, l0 int32)
//go:linkname F_shared_ts_extend_down github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shared_ts_extend_down
func F_shared_ts_extend_down(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_populate_compact_attribute github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_populate_compact_attribute
func F_populate_compact_attribute(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateTemplateTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateTemplateTupleDesc
func F_CreateTemplateTupleDesc(m *base.Module, l0 int32) int32
//go:linkname F_CreateTupleDescCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateTupleDescCopy
func F_CreateTupleDescCopy(m *base.Module, l0 int32) int32
//go:linkname F_CreateTupleDescCopyConstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateTupleDescCopyConstr
func F_CreateTupleDescCopyConstr(m *base.Module, l0 int32) int32
//go:linkname F_DecrTupleDescRefCount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecrTupleDescRefCount
func F_DecrTupleDescRefCount(m *base.Module, l0 int32)
//go:linkname F_ginFindLeafPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ginFindLeafPage
func F_ginFindLeafPage(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ginStepRight github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ginStepRight
func F_ginStepRight(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_freeGinBtreeStack github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_freeGinBtreeStack
func F_freeGinBtreeStack(m *base.Module, l0 int32)
//go:linkname F_ginPlaceToPage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ginPlaceToPage
func F_ginPlaceToPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ginInitBA github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ginInitBA
func F_ginInitBA(m *base.Module, l0 int32)
//go:linkname F_ginInsertBAEntries github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ginInsertBAEntries
func F_ginInsertBAEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ginGetBAEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ginGetBAEntry
func F_ginGetBAEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_GinDataLeafPageGetItems github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GinDataLeafPageGetItems
func F_GinDataLeafPageGetItems(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ginScanBeginPostingTree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ginScanBeginPostingTree
func F_ginScanBeginPostingTree(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ginReadTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ginReadTuple
func F_ginReadTuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ginInsertCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ginInsertCleanup
func F_ginInsertCleanup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_scanGetCandidate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_scanGetCandidate
func F_scanGetCandidate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_matchPartialInPendingList github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_matchPartialInPendingList
func F_matchPartialInPendingList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_ginEntryInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ginEntryInsert
func F_ginEntryInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F__gin_build_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__gin_build_tuple
func F__gin_build_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_ginFreeScanKeys github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ginFreeScanKeys
func F_ginFreeScanKeys(m *base.Module, l0 int32)
//go:linkname F_ginFillScanEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ginFillScanEntry
func F_ginFillScanEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_gintuple_get_attrnum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gintuple_get_attrnum
func F_gintuple_get_attrnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gintuple_get_key github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gintuple_get_key
func F_gintuple_get_key(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ginGetStats github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ginGetStats
func F_ginGetStats(m *base.Module, l0 int32, l1 int32)
//go:linkname F_initGISTstate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_initGISTstate
func F_initGISTstate(m *base.Module, l0 int32) int32
//go:linkname F_createTempGistContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_createTempGistContext
func F_createTempGistContext(m *base.Module) int32
//go:linkname F_gistfinishsplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistfinishsplit
func F_gistfinishsplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_gistplacetopage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistplacetopage
func F_gistplacetopage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32
//go:linkname F_gist_indexsortbuild_levelstate_add github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gist_indexsortbuild_levelstate_add
func F_gist_indexsortbuild_levelstate_add(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_gistProcessEmptyingQueue github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistProcessEmptyingQueue
func F_gistProcessEmptyingQueue(m *base.Module, l0 int32)
//go:linkname F_gistScanPage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gistScanPage
func F_gistScanPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_index_getattr_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_getattr_2
func F_index_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_computeDistance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_computeDistance
func F_computeDistance(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_gistSplitByKey github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gistSplitByKey
func F_gistSplitByKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_gistextractpage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gistextractpage
func F_gistextractpage(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gistjoinvector github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gistjoinvector
func F_gistjoinvector(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gistfillitupvec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistfillitupvec
func F_gistfillitupvec(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_gistMakeUnionItVec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gistMakeUnionItVec
func F_gistMakeUnionItVec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_gistFormTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gistFormTuple
func F_gistFormTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_gistcheckpage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistcheckpage
func F_gistcheckpage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_gistNewBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gistNewBuffer
func F_gistNewBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__hash_doinsert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__hash_doinsert
func F__hash_doinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__hash_dropscanbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__hash_dropscanbuf
func F__hash_dropscanbuf(m *base.Module, l0 int32)
//go:linkname F__hash_readnext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__hash_readnext
func F__hash_readnext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__hash_convert_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__hash_convert_tuple
func F__hash_convert_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_HeapCheckForSerializableConflictOut github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HeapCheckForSerializableConflictOut
func F_HeapCheckForSerializableConflictOut(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_heap_getnext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_getnext
func F_heap_getnext(m *base.Module, l0 int32) int32
//go:linkname F_heap_getattr_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_getattr_1
func F_heap_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_heap_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_fetch
func F_heap_fetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_heap_hot_search_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_hot_search_buffer
func F_heap_hot_search_buffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_GetBulkInsertState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetBulkInsertState
func F_GetBulkInsertState(m *base.Module) int32
//go:linkname F_heap_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_insert
func F_heap_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_heap_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_delete
func F_heap_delete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_DoesMultiXactIdConflict github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DoesMultiXactIdConflict
func F_DoesMultiXactIdConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_compute_new_xmax_infomask github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_compute_new_xmax_infomask
func F_compute_new_xmax_infomask(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_heap_update github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_heap_update
func F_heap_update(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_HeapTupleGetUpdateXid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_HeapTupleGetUpdateXid
func F_HeapTupleGetUpdateXid(m *base.Module, l0 int32) int32
//go:linkname F_heap_lock_updated_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_lock_updated_tuple
func F_heap_lock_updated_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_heap_abort_speculative github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_abort_speculative
func F_heap_abort_speculative(m *base.Module, l0 int32, l1 int32)
//go:linkname F_HeapTupleHeaderAdvanceConflictHorizon github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HeapTupleHeaderAdvanceConflictHorizon
func F_HeapTupleHeaderAdvanceConflictHorizon(m *base.Module, l0 int32, l1 int32)
//go:linkname F_HeapTupleSetHintBits github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_HeapTupleSetHintBits
func F_HeapTupleSetHintBits(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_HeapTupleSatisfiesUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HeapTupleSatisfiesUpdate
func F_HeapTupleSatisfiesUpdate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_toast_insert_or_update github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_toast_insert_or_update
func F_heap_toast_insert_or_update(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_heap_page_prune_opt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_page_prune_opt
func F_heap_page_prune_opt(m *base.Module, l0 int32, l1 int32)
//go:linkname F_heap_prune_record_unchanged_lp_normal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_heap_prune_record_unchanged_lp_normal
func F_heap_prune_record_unchanged_lp_normal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_log_heap_prune_and_freeze github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_log_heap_prune_and_freeze
func F_log_heap_prune_and_freeze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32)
//go:linkname F_visibilitymap_clear github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_visibilitymap_clear
func F_visibilitymap_clear(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_visibilitymap_pin github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_visibilitymap_pin
func F_visibilitymap_pin(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_visibilitymap_get_status github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_visibilitymap_get_status
func F_visibilitymap_get_status(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetIndexAmRoutine github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetIndexAmRoutine
func F_GetIndexAmRoutine(m *base.Module, l0 int32) int32
//go:linkname F_GetIndexAmRoutineByAmId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetIndexAmRoutineByAmId
func F_GetIndexAmRoutineByAmId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_IndexAmTranslateStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_IndexAmTranslateStrategy
func F_IndexAmTranslateStrategy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_identify_opfamily_groups github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_identify_opfamily_groups
func F_identify_opfamily_groups(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_amproc_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_amproc_signature
func F_check_amproc_signature(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_check_amoptsproc_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_amoptsproc_signature
func F_check_amoptsproc_signature(m *base.Module, l0 int32) int32
//go:linkname F_check_amop_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_amop_signature
func F_check_amop_signature(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_opfamily_can_sort_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_opfamily_can_sort_type
func F_opfamily_can_sort_type(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationGetIndexScan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationGetIndexScan
func F_RelationGetIndexScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_systable_beginscan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_beginscan
func F_systable_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_systable_endscan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_systable_endscan
func F_systable_endscan(m *base.Module, l0 int32)
//go:linkname F_systable_beginscan_ordered github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_systable_beginscan_ordered
func F_systable_beginscan_ordered(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_systable_endscan_ordered github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_endscan_ordered
func F_systable_endscan_ordered(m *base.Module, l0 int32)
//go:linkname F_index_insert_cleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_insert_cleanup
func F_index_insert_cleanup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_index_beginscan_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_beginscan_internal
func F_index_beginscan_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_index_rescan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_rescan
func F_index_rescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_index_getnext_tid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_index_getnext_tid
func F_index_getnext_tid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_fetch_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_index_fetch_heap
func F_index_fetch_heap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_getnext_slot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_getnext_slot
func F_index_getnext_slot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_index_bulk_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_bulk_delete
func F_index_bulk_delete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_index_can_return github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_index_can_return
func F_index_can_return(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_store_float8_orderby_distances github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_store_float8_orderby_distances
func F_index_store_float8_orderby_distances(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__bt_finish_split github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_finish_split
func F__bt_finish_split(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__bt_getbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_getbuf
func F__bt_getbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_relbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_relbuf
func F__bt_relbuf(m *base.Module, l0 int32)
//go:linkname F__bt_unlockbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_unlockbuf
func F__bt_unlockbuf(m *base.Module, l0 int32)
//go:linkname F__bt_parallel_primscan_schedule github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_parallel_primscan_schedule
func F__bt_parallel_primscan_schedule(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_readpage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__bt_readpage
func F__bt_readpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_steppage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_steppage
func F__bt_steppage(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_binsrch_array_skey github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_binsrch_array_skey
func F__bt_binsrch_array_skey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F__bt_check_compare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_check_compare
func F__bt_check_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F__bt_tuple_before_array_skeys github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_tuple_before_array_skeys
func F__bt_tuple_before_array_skeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F__bt_binsrch_skiparray_skey github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__bt_binsrch_skiparray_skey
func F__bt_binsrch_skiparray_skey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F__bt_truncate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_truncate
func F__bt_truncate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_check_third_page github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_check_third_page
func F__bt_check_third_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_array_desc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_array_desc
func F_array_desc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_standby_desc_invalidations github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_standby_desc_invalidations
func F_standby_desc_invalidations(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_spgdoinsert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_spgdoinsert
func F_spgdoinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_spgWalk github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_spgWalk
func F_spgWalk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_spgvacuumscan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_spgvacuumscan
func F_spgvacuumscan(m *base.Module, l0 int32)
//go:linkname F_table_open github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_table_open
func F_table_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_openrv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_table_openrv
func F_table_openrv(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_openrv_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_table_openrv_extended
func F_table_openrv_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_table_slot_callbacks github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_table_slot_callbacks
func F_table_slot_callbacks(m *base.Module, l0 int32) int32
//go:linkname F_table_slot_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_table_slot_create
func F_table_slot_create(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_beginscan_catalog github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_table_beginscan_catalog
func F_table_beginscan_catalog(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_table_parallelscan_estimate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_table_parallelscan_estimate
func F_table_parallelscan_estimate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_parallelscan_initialize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_table_parallelscan_initialize
func F_table_parallelscan_initialize(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_TransactionIdSetPageStatusInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TransactionIdSetPageStatusInternal
func F_TransactionIdSetPageStatusInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64)
//go:linkname F_TransactionTreeSetCommitTsData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TransactionTreeSetCommitTsData
func F_TransactionTreeSetCommitTsData(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32)
//go:linkname F_ActivateCommitTs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ActivateCommitTs
func F_ActivateCommitTs(m *base.Module)
//go:linkname F_GenericXLogStart github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GenericXLogStart
func F_GenericXLogStart(m *base.Module, l0 int32) int32
//go:linkname F_GenericXLogRegisterBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GenericXLogRegisterBuffer
func F_GenericXLogRegisterBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GenericXLogFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GenericXLogFinish
func F_GenericXLogFinish(m *base.Module, l0 int32)
//go:linkname F_GetMultiXactIdMembers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetMultiXactIdMembers
func F_GetMultiXactIdMembers(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReadNextMultiXactId github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReadNextMultiXactId
func F_ReadNextMultiXactId(m *base.Module) int32
//go:linkname F_find_multixact_start github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_find_multixact_start
func F_find_multixact_start(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MultiXactSetNextMXact github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MultiXactSetNextMXact
func F_MultiXactSetNextMXact(m *base.Module, l0 int32, l1 int32)
//go:linkname F_MultiXactAdvanceOldest github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MultiXactAdvanceOldest
func F_MultiXactAdvanceOldest(m *base.Module, l0 int32, l1 int32)
//go:linkname F_MultiXactMemberFreezeThreshold github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MultiXactMemberFreezeThreshold
func F_MultiXactMemberFreezeThreshold(m *base.Module) int32
//go:linkname F_CreateParallelContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateParallelContext
func F_CreateParallelContext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_InitializeParallelDSM github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InitializeParallelDSM
func F_InitializeParallelDSM(m *base.Module, l0 int32)
//go:linkname F_ReinitializeParallelDSM github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReinitializeParallelDSM
func F_ReinitializeParallelDSM(m *base.Module, l0 int32)
//go:linkname F_WaitForParallelWorkersToFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WaitForParallelWorkersToFinish
func F_WaitForParallelWorkersToFinish(m *base.Module, l0 int32)
//go:linkname F_LaunchParallelWorkers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LaunchParallelWorkers
func F_LaunchParallelWorkers(m *base.Module, l0 int32)
//go:linkname F_WaitForParallelWorkersToAttach github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WaitForParallelWorkersToAttach
func F_WaitForParallelWorkersToAttach(m *base.Module, l0 int32)
//go:linkname F_DestroyParallelContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DestroyParallelContext
func F_DestroyParallelContext(m *base.Module, l0 int32)
//go:linkname F_AtEOXact_Parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtEOXact_Parallel
func F_AtEOXact_Parallel(m *base.Module, l0 int32)
//go:linkname F_check_slru_buffers github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_slru_buffers
func F_check_slru_buffers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SimpleLruZeroPage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SimpleLruZeroPage
func F_SimpleLruZeroPage(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_SimpleLruReadPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SimpleLruReadPage
func F_SimpleLruReadPage(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_SimpleLruTruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SimpleLruTruncate
func F_SimpleLruTruncate(m *base.Module, l0 int32, l1 int64)
//go:linkname F_SlruScanDirectory github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SlruScanDirectory
func F_SlruScanDirectory(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SlruSyncFileTag github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SlruSyncFileTag
func F_SlruSyncFileTag(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SubTransGetTopmostTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SubTransGetTopmostTransaction
func F_SubTransGetTopmostTransaction(m *base.Module, l0 int32) int32
//go:linkname F_TransactionIdDidCommit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransactionIdDidCommit
func F_TransactionIdDidCommit(m *base.Module, l0 int32) int32
//go:linkname F_TransactionIdPrecedes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TransactionIdPrecedes
func F_TransactionIdPrecedes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TransactionIdDidAbort github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TransactionIdDidAbort
func F_TransactionIdDidAbort(m *base.Module, l0 int32) int32
//go:linkname F_TransactionIdCommitTree github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TransactionIdCommitTree
func F_TransactionIdCommitTree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_TransactionIdAsyncCommitTree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TransactionIdAsyncCommitTree
func F_TransactionIdAsyncCommitTree(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64)
//go:linkname F_TransactionIdAbortTree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransactionIdAbortTree
func F_TransactionIdAbortTree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AtAbort_Twophase github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtAbort_Twophase
func F_AtAbort_Twophase(m *base.Module)
//go:linkname F_MarkAsPreparingGuts github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MarkAsPreparingGuts
func F_MarkAsPreparingGuts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32)
//go:linkname F_TwoPhaseGetDummyProcNumber github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TwoPhaseGetDummyProcNumber
func F_TwoPhaseGetDummyProcNumber(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TwoPhaseGetDummyProc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TwoPhaseGetDummyProc
func F_TwoPhaseGetDummyProc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RegisterTwoPhaseRecord github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RegisterTwoPhaseRecord
func F_RegisterTwoPhaseRecord(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ProcessTwoPhaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ProcessTwoPhaseBuffer
func F_ProcessTwoPhaseBuffer(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_PrescanPreparedTransactions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PrescanPreparedTransactions
func F_PrescanPreparedTransactions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SetTransactionIdLimit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SetTransactionIdLimit
func F_SetTransactionIdLimit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetTopTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetTopTransactionId
func F_GetTopTransactionId(m *base.Module) int32
//go:linkname F_AssignTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AssignTransactionId
func F_AssignTransactionId(m *base.Module, l0 int32)
//go:linkname F_GetCurrentTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetCurrentTransactionId
func F_GetCurrentTransactionId(m *base.Module) int32
//go:linkname F_GetCurrentCommandId github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetCurrentCommandId
func F_GetCurrentCommandId(m *base.Module, l0 int32) int32
//go:linkname F_CommandCounterIncrement github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CommandCounterIncrement
func F_CommandCounterIncrement(m *base.Module)
//go:linkname F_StartTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_StartTransaction
func F_StartTransaction(m *base.Module)
//go:linkname F_ShowTransactionStateRec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ShowTransactionStateRec
func F_ShowTransactionStateRec(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CommitTransactionCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CommitTransactionCommand
func F_CommitTransactionCommand(m *base.Module)
//go:linkname F_CleanupTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CleanupTransaction
func F_CleanupTransaction(m *base.Module)
//go:linkname F_AbortTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AbortTransaction
func F_AbortTransaction(m *base.Module)
//go:linkname F_AbortSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AbortSubTransaction
func F_AbortSubTransaction(m *base.Module)
//go:linkname F_CleanupSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CleanupSubTransaction
func F_CleanupSubTransaction(m *base.Module)
//go:linkname F_AbortCurrentTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AbortCurrentTransaction
func F_AbortCurrentTransaction(m *base.Module)
//go:linkname F_PreventInTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreventInTransactionBlock
func F_PreventInTransactionBlock(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XactLogCommitRecord github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XactLogCommitRecord
func F_XactLogCommitRecord(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int64
//go:linkname F_XactLogAbortRecord github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XactLogAbortRecord
func F_XactLogAbortRecord(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int64
//go:linkname F_XLogFlush github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogFlush
func F_XLogFlush(m *base.Module, l0 int64)
//go:linkname F_XLogSetAsyncXactLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogSetAsyncXactLSN
func F_XLogSetAsyncXactLSN(m *base.Module, l0 int64)
//go:linkname F_issue_xlog_fsync github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_issue_xlog_fsync
func F_issue_xlog_fsync(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_DataChecksumsEnabled github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DataChecksumsEnabled
func F_DataChecksumsEnabled(m *base.Module) int32
//go:linkname F_CheckRequiredParameterValues github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckRequiredParameterValues
func F_CheckRequiredParameterValues(m *base.Module)
//go:linkname F_GetRedoRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetRedoRecPtr
func F_GetRedoRecPtr(m *base.Module) int64
//go:linkname F_GetFullPageWriteInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetFullPageWriteInfo
func F_GetFullPageWriteInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateCheckPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateCheckPoint
func F_CreateCheckPoint(m *base.Module, l0 int32) int32
//go:linkname F_KeepLogSeg github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_KeepLogSeg
func F_KeepLogSeg(m *base.Module, l0 int64, l1 int32)
//go:linkname F_RequestXLogSwitch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RequestXLogSwitch
func F_RequestXLogSwitch(m *base.Module, l0 int32) int64
//go:linkname F_RecoveryRestartPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RecoveryRestartPoint
func F_RecoveryRestartPoint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetXLogInsertRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetXLogInsertRecPtr
func F_GetXLogInsertRecPtr(m *base.Module) int64
//go:linkname F_GetOldestRestartPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetOldestRestartPoint
func F_GetOldestRestartPoint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogBeginInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_XLogBeginInsert
func F_XLogBeginInsert(m *base.Module)
//go:linkname F_XLogRegisterData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XLogRegisterData
func F_XLogRegisterData(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogRegisterBufData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogRegisterBufData
func F_XLogRegisterBufData(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogInsert
func F_XLogInsert(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_log_newpage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_log_newpage
func F_log_newpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_log_newpage_range github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_log_newpage_range
func F_log_newpage_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_XLogReaderAllocate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_XLogReaderAllocate
func F_XLogReaderAllocate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_XLogReaderFree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogReaderFree
func F_XLogReaderFree(m *base.Module, l0 int32)
//go:linkname F_XLogReadAhead github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogReadAhead
func F_XLogReadAhead(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_XLogRecGetBlockTag github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XLogRecGetBlockTag
func F_XLogRecGetBlockTag(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_CheckForStandbyTrigger github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckForStandbyTrigger
func F_CheckForStandbyTrigger(m *base.Module) int32
//go:linkname F_SetRecoveryPause github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SetRecoveryPause
func F_SetRecoveryPause(m *base.Module, l0 int32)
//go:linkname F_PromoteIsTriggered github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PromoteIsTriggered
func F_PromoteIsTriggered(m *base.Module) int32
//go:linkname F_WakeupRecovery github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WakeupRecovery
func F_WakeupRecovery(m *base.Module)
//go:linkname F_GetCurrentReplayRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetCurrentReplayRecPtr
func F_GetCurrentReplayRecPtr(m *base.Module, l0 int32) int64
//go:linkname F_error_multiple_recovery_targets github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_error_multiple_recovery_targets
func F_error_multiple_recovery_targets(m *base.Module)
//go:linkname F_XLogReadBufferForRedo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogReadBufferForRedo
func F_XLogReadBufferForRedo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_XLogReadBufferForRedoExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_XLogReadBufferForRedoExtended
func F_XLogReadBufferForRedoExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_log_invalid_page github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_log_invalid_page
func F_log_invalid_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CreateFakeRelcacheEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateFakeRelcacheEntry
func F_CreateFakeRelcacheEntry(m *base.Module, l0 int32) int32
//go:linkname F_bbsink_forward_archive_contents github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_bbsink_forward_archive_contents
func F_bbsink_forward_archive_contents(m *base.Module, l0 int32, l1 int32)
//go:linkname F_bbsink_forward_end_manifest github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bbsink_forward_end_manifest
func F_bbsink_forward_end_manifest(m *base.Module, l0 int32)
//go:linkname F_ReportWalSummaryError github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReportWalSummaryError
func F_ReportWalSummaryError(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_restrict_and_check_grant github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_restrict_and_check_grant
func F_restrict_and_check_grant(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int64
//go:linkname F_merge_acl_with_grant github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_merge_acl_with_grant
func F_merge_acl_with_grant(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int32, l7 int32) int32
//go:linkname F_heap_getattr_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_getattr_2
func F_heap_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_aclcheck_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_aclcheck_error
func F_aclcheck_error(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_aclcheck_error_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_aclcheck_error_type
func F_aclcheck_error_type(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_class_aclmask_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_class_aclmask_ext
func F_pg_class_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int64
//go:linkname F_object_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_object_aclcheck
func F_object_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_pg_attribute_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_attribute_aclcheck
func F_pg_attribute_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_pg_attribute_aclcheck_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_attribute_aclcheck_ext
func F_pg_attribute_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_pg_attribute_aclcheck_all github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_attribute_aclcheck_all
func F_pg_attribute_aclcheck_all(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pg_class_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_class_aclcheck
func F_pg_class_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pg_class_aclcheck_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_class_aclcheck_ext
func F_pg_class_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pg_largeobject_aclcheck_snapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_largeobject_aclcheck_snapshot
func F_pg_largeobject_aclcheck_snapshot(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_get_user_default_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_user_default_acl
func F_get_user_default_acl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_recordDependencyOnNewAcl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recordDependencyOnNewAcl
func F_recordDependencyOnNewAcl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetNewOidWithIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetNewOidWithIndex
func F_GetNewOidWithIndex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_performDeletion github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_performDeletion
func F_performDeletion(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_add_exact_object_address github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_exact_object_address
func F_add_exact_object_address(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SystemAttributeDefinition github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SystemAttributeDefinition
func F_SystemAttributeDefinition(m *base.Module, l0 int32) int32
//go:linkname F_CheckAttributeType github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckAttributeType
func F_CheckAttributeType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_StoreRelCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_StoreRelCheck
func F_StoreRelCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_check_nested_generated_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_nested_generated_walker
func F_check_nested_generated_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_virtual_generated_security_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_check_virtual_generated_security_walker
func F_check_virtual_generated_security_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_heap_truncate_find_FKs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_heap_truncate_find_FKs
func F_heap_truncate_find_FKs(m *base.Module, l0 int32) int32
//go:linkname F_index_check_primary_key github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_index_check_primary_key
func F_index_check_primary_key(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_index_create github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_create
func F_index_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32) int32
//go:linkname F_BuildIndexInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BuildIndexInfo
func F_BuildIndexInfo(m *base.Module, l0 int32) int32
//go:linkname F_index_concurrently_build github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_index_concurrently_build
func F_index_concurrently_build(m *base.Module, l0 int32, l1 int32)
//go:linkname F_index_set_state_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_set_state_flags
func F_index_set_state_flags(m *base.Module, l0 int32, l1 int32)
//go:linkname F_IndexGetRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_IndexGetRelation
func F_IndexGetRelation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CompareIndexInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CompareIndexInfo
func F_CompareIndexInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_CatalogOpenIndexes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CatalogOpenIndexes
func F_CatalogOpenIndexes(m *base.Module, l0 int32) int32
//go:linkname F_CatalogCloseIndexes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CatalogCloseIndexes
func F_CatalogCloseIndexes(m *base.Module, l0 int32)
//go:linkname F_CatalogTupleInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CatalogTupleInsert
func F_CatalogTupleInsert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CatalogTupleInsertWithInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CatalogTupleInsertWithInfo
func F_CatalogTupleInsertWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CatalogTupleUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CatalogTupleUpdate
func F_CatalogTupleUpdate(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CatalogTupleUpdateWithInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CatalogTupleUpdateWithInfo
func F_CatalogTupleUpdateWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_LookupExplicitNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LookupExplicitNamespace
func F_LookupExplicitNamespace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_recomputeNamespacePath github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recomputeNamespacePath
func F_recomputeNamespacePath(m *base.Module)
//go:linkname F_RangeVarGetCreationNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RangeVarGetCreationNamespace
func F_RangeVarGetCreationNamespace(m *base.Module, l0 int32) int32
//go:linkname F_AccessTempTableNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AccessTempTableNamespace
func F_AccessTempTableNamespace(m *base.Module, l0 int32)
//go:linkname F_get_namespace_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_namespace_oid
func F_get_namespace_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RangeVarAdjustRelationPersistence github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RangeVarAdjustRelationPersistence
func F_RangeVarAdjustRelationPersistence(m *base.Module, l0 int32, l1 int32)
//go:linkname F_isAnyTempNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isAnyTempNamespace
func F_isAnyTempNamespace(m *base.Module, l0 int32) int32
//go:linkname F_FuncnameGetCandidates github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FuncnameGetCandidates
func F_FuncnameGetCandidates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_DeconstructQualifiedName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DeconstructQualifiedName
func F_DeconstructQualifiedName(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_FunctionIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FunctionIsVisibleExt
func F_FunctionIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_OpernameGetOprid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_OpernameGetOprid
func F_OpernameGetOprid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_OpclassnameGetOpcid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_OpclassnameGetOpcid
func F_OpclassnameGetOpcid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CollationIsVisible github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CollationIsVisible
func F_CollationIsVisible(m *base.Module, l0 int32) int32
//go:linkname F_get_ts_dict_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_ts_dict_oid
func F_get_ts_dict_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TSDictionaryIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TSDictionaryIsVisibleExt
func F_TSDictionaryIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_ts_template_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_ts_template_oid
func F_get_ts_template_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_ts_config_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_ts_config_oid
func F_get_ts_config_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeRangeVarFromNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeRangeVarFromNameList
func F_makeRangeVarFromNameList(m *base.Module, l0 int32) int32
//go:linkname F_isTempToastNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_isTempToastNamespace
func F_isTempToastNamespace(m *base.Module, l0 int32) int32
//go:linkname F_GetSearchPathMatcher github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetSearchPathMatcher
func F_GetSearchPathMatcher(m *base.Module, l0 int32) int32
//go:linkname F_SearchPathMatchesCurrentEnvironment github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchPathMatchesCurrentEnvironment
func F_SearchPathMatchesCurrentEnvironment(m *base.Module, l0 int32) int32
//go:linkname F_get_collation_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_collation_oid
func F_get_collation_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AtEOXact_Namespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtEOXact_Namespace
func F_AtEOXact_Namespace(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RunObjectPostCreateHook github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RunObjectPostCreateHook
func F_RunObjectPostCreateHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RunObjectPostAlterHook github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RunObjectPostAlterHook
func F_RunObjectPostAlterHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_get_object_address_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_object_address_type
func F_get_object_address_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_getObjectDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getObjectDescription
func F_getObjectDescription(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_object_class_descr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_object_class_descr
func F_get_object_class_descr(m *base.Module, l0 int32) int32
//go:linkname F_get_object_oid_index github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_object_oid_index
func F_get_object_oid_index(m *base.Module, l0 int32) int32
//go:linkname F_get_object_catcache_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_object_catcache_oid
func F_get_object_catcache_oid(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_object_attnum_name
func F_get_object_attnum_name(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_namespace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_object_attnum_namespace
func F_get_object_attnum_namespace(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_owner github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_object_attnum_owner
func F_get_object_attnum_owner(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_object_attnum_acl
func F_get_object_attnum_acl(m *base.Module, l0 int32) int32
//go:linkname F_get_catalog_object_by_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_catalog_object_by_oid
func F_get_catalog_object_by_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getObjectDescriptionOids github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_getObjectDescriptionOids
func F_getObjectDescriptionOids(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_partition_parent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_partition_parent
func F_get_partition_parent(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_partition_ancestors github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_partition_ancestors
func F_get_partition_ancestors(m *base.Module, l0 int32) int32
//go:linkname F_map_partition_varattnos github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_map_partition_varattnos
func F_map_partition_varattnos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_StoreAttrDefault github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StoreAttrDefault
func F_StoreAttrDefault(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GetAttrDefaultOid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetAttrDefaultOid
func F_GetAttrDefaultOid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CreateConstraintEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateConstraintEntry
func F_CreateConstraintEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32, l21 int32, l22 int32, l23 int32, l24 int32, l25 int32, l26 int32, l27 int32, l28 int32, l29 int32, l30 int32, l31 int32, l32 int32) int32
//go:linkname F_extractNotNullColumn github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_extractNotNullColumn
func F_extractNotNullColumn(m *base.Module, l0 int32) int32
//go:linkname F_ConstraintSetParentConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ConstraintSetParentConstraint
func F_ConstraintSetParentConstraint(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_relation_constraint_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_relation_constraint_oid
func F_get_relation_constraint_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_relation_idx_constraint_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_relation_idx_constraint_oid
func F_get_relation_idx_constraint_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_domain_constraint_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_domain_constraint_oid
func F_get_domain_constraint_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DeconstructFkConstraintRow github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DeconstructFkConstraintRow
func F_DeconstructFkConstraintRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_FindFKPeriodOpers github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FindFKPeriodOpers
func F_FindFKPeriodOpers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_recordDependencyOn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recordDependencyOn
func F_recordDependencyOn(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_recordDependencyOnCurrentExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recordDependencyOnCurrentExtension
func F_recordDependencyOnCurrentExtension(m *base.Module, l0 int32, l1 int32)
//go:linkname F_deleteDependencyRecordsForClass github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_deleteDependencyRecordsForClass
func F_deleteDependencyRecordsForClass(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_sequenceIsOwned github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sequenceIsOwned
func F_sequenceIsOwned(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_find_inheritance_children github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_inheritance_children
func F_find_inheritance_children(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_all_inheritors github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_all_inheritors
func F_find_all_inheritors(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_has_superclass github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_has_superclass
func F_has_superclass(m *base.Module, l0 int32) int32
//go:linkname F_LargeObjectExists github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LargeObjectExists
func F_LargeObjectExists(m *base.Module, l0 int32) int32
//go:linkname F_LargeObjectExistsWithSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LargeObjectExistsWithSnapshot
func F_LargeObjectExistsWithSnapshot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetPubPartitionOptionRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetPubPartitionOptionRelations
func F_GetPubPartitionOptionRelations(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetPublication github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetPublication
func F_GetPublication(m *base.Module, l0 int32) int32
//go:linkname F_GetSchemaPublicationRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetSchemaPublicationRelations
func F_GetSchemaPublicationRelations(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetPublicationSchemas github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetPublicationSchemas
func F_GetPublicationSchemas(m *base.Module, l0 int32) int32
//go:linkname F_GetAllSchemaPublicationRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetAllSchemaPublicationRelations
func F_GetAllSchemaPublicationRelations(m *base.Module, l0 int32) int32
//go:linkname F_GetPublicationByName github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetPublicationByName
func F_GetPublicationByName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_recordSharedDependencyOn github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_recordSharedDependencyOn
func F_recordSharedDependencyOn(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_recordDependencyOnOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_recordDependencyOnOwner
func F_recordDependencyOnOwner(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_updateAclDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_updateAclDependencies
func F_updateAclDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_updateInitAclDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_updateInitAclDependencies
func F_updateInitAclDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_AddSubscriptionRelState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AddSubscriptionRelState
func F_AddSubscriptionRelState(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32)
//go:linkname F_GenerateTypeDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GenerateTypeDependencies
func F_GenerateTypeDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_RenameTypeInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RenameTypeInternal
func F_RenameTypeInternal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RelationCreateStorage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationCreateStorage
func F_RelationCreateStorage(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_smgrDoPendingDeletes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgrDoPendingDeletes
func F_smgrDoPendingDeletes(m *base.Module, l0 int32)
//go:linkname F_smgrGetPendingDeletes github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_smgrGetPendingDeletes
func F_smgrGetPendingDeletes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_toast_table github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_toast_table
func F_create_toast_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_transformStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformStmt
func F_transformStmt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SystemTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SystemTypeName
func F_SystemTypeName(m *base.Module, l0 int32) int32
//go:linkname F_SystemFuncName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SystemFuncName
func F_SystemFuncName(m *base.Module, l0 int32) int32
//go:linkname F_check_agglevels_and_constraints github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_agglevels_and_constraints
func F_check_agglevels_and_constraints(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformWindowFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_transformWindowFuncCall
func F_transformWindowFuncCall(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_transformFromClauseItem github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_transformFromClauseItem
func F_transformFromClauseItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_transformWhereClause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_transformWhereClause
func F_transformWhereClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_findTargetlistEntrySQL99 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_findTargetlistEntrySQL99
func F_findTargetlistEntrySQL99(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_findTargetlistEntrySQL92 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_findTargetlistEntrySQL92
func F_findTargetlistEntrySQL92(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_addTargetToGroupList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_addTargetToGroupList
func F_addTargetToGroupList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_transformSortClause github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformSortClause
func F_transformSortClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_addTargetToSortList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_addTargetToSortList
func F_addTargetToSortList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_transformDistinctClause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_transformDistinctClause
func F_transformDistinctClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_coerce_to_target_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_coerce_to_target_type
func F_coerce_to_target_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_can_coerce_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_can_coerce_type
func F_can_coerce_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_coerce_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_coerce_type
func F_coerce_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_coerce_to_boolean github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_coerce_to_boolean
func F_coerce_to_boolean(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_coerce_to_specific_type github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_coerce_to_specific_type
func F_coerce_to_specific_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_coerce_to_common_type github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_coerce_to_common_type
func F_coerce_to_common_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_select_common_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_select_common_typmod
func F_select_common_typmod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_IsBinaryCoercibleWithCast github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_IsBinaryCoercibleWithCast
func F_IsBinaryCoercibleWithCast(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_assign_expr_collations github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_assign_expr_collations
func F_assign_expr_collations(m *base.Module, l0 int32, l1 int32)
//go:linkname F_merge_collation_state github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_merge_collation_state
func F_merge_collation_state(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_select_common_collation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_select_common_collation
func F_select_common_collation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_checkWellFormedRecursionWalker github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_checkWellFormedRecursionWalker
func F_checkWellFormedRecursionWalker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_analyzeCTE github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_analyzeCTE
func F_analyzeCTE(m *base.Module, l0 int32, l1 int32)
//go:linkname F_analyzeCTETargetList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_analyzeCTETargetList
func F_analyzeCTETargetList(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_transformExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformExpr
func F_transformExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getJsonEncodingConst github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_getJsonEncodingConst
func F_getJsonEncodingConst(m *base.Module, l0 int32) int32
//go:linkname F_ParseFuncOrColumn github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ParseFuncOrColumn
func F_ParseFuncOrColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_func_signature_string github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_func_signature_string
func F_func_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_LookupFuncWithArgs github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LookupFuncWithArgs
func F_LookupFuncWithArgs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_parsestate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_parsestate
func F_make_parsestate(m *base.Module, l0 int32) int32
//go:linkname F_free_parsestate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_free_parsestate
func F_free_parsestate(m *base.Module, l0 int32)
//go:linkname F_parser_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parser_errposition
func F_parser_errposition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformContainerSubscripts github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformContainerSubscripts
func F_transformContainerSubscripts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_op_signature_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_op_signature_string
func F_op_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LookupOperWithArgs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LookupOperWithArgs
func F_LookupOperWithArgs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_sort_group_operators github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_sort_group_operators
func F_get_sort_group_operators(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_check_variable_parameters github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_variable_parameters
func F_check_variable_parameters(m *base.Module, l0 int32, l1 int32)
//go:linkname F_refnameNamespaceItem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_refnameNamespaceItem
func F_refnameNamespaceItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_checkNameSpaceConflicts github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_checkNameSpaceConflicts
func F_checkNameSpaceConflicts(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetNSItemByRangeTablePosn github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetNSItemByRangeTablePosn
func F_GetNSItemByRangeTablePosn(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetCTEForRTE github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetCTEForRTE
func F_GetCTEForRTE(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_markVarForSelectPriv github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_markVarForSelectPriv
func F_markVarForSelectPriv(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getRTEPermissionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getRTEPermissionInfo
func F_getRTEPermissionInfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parserOpenTable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parserOpenTable
func F_parserOpenTable(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_buildRelationAliases github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_buildRelationAliases
func F_buildRelationAliases(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addRangeTableEntryForSubquery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_addRangeTableEntryForSubquery
func F_addRangeTableEntryForSubquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_buildNSItemFromLists github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_buildNSItemFromLists
func F_buildNSItemFromLists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_addNSItemToQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_addNSItemToQuery
func F_addNSItemToQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_expandNSItemVars github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_expandNSItemVars
func F_expandNSItemVars(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_expandNSItemAttrs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_expandNSItemAttrs
func F_expandNSItemAttrs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_rte_attribute_name github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_rte_attribute_name
func F_get_rte_attribute_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_attnumTypeId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_attnumTypeId
func F_attnumTypeId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_errorMissingRTE github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errorMissingRTE
func F_errorMissingRTE(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformAssignmentIndirection github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformAssignmentIndirection
func F_transformAssignmentIndirection(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32
//go:linkname F_expandRecordVariable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expandRecordVariable
func F_expandRecordVariable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LookupTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LookupTypeName
func F_LookupTypeName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LookupTypeNameExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LookupTypeNameExtended
func F_LookupTypeNameExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_appendTypeNameToBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendTypeNameToBuffer
func F_appendTypeNameToBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LookupTypeNameOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LookupTypeNameOid
func F_LookupTypeNameOid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_typenameTypeId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_typenameTypeId
func F_typenameTypeId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_typenameTypeIdAndMod github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_typenameTypeIdAndMod
func F_typenameTypeIdAndMod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_LookupCollation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LookupCollation
func F_LookupCollation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_generateSerialExtraStmts github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generateSerialExtraStmts
func F_generateSerialExtraStmts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_generateClonedIndexStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generateClonedIndexStmt
func F_generateClonedIndexStmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_raw_parser github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_raw_parser
func F_raw_parser(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_core_yylex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_core_yylex
func F_core_yylex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_scanner_init github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_scanner_init
func F_scanner_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_scanner_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_scanner_finish
func F_scanner_finish(m *base.Module, l0 int32)
//go:linkname F_downcase_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_downcase_identifier
func F_downcase_identifier(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_report_namespace_conflict github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_report_namespace_conflict
func F_report_namespace_conflict(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_index_am_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_index_am_oid
func F_get_index_am_oid(m *base.Module, l0 int32) int32
//go:linkname F_get_am_type_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_am_type_oid
func F_get_am_type_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_am_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_am_name
func F_get_am_name(m *base.Module, l0 int32) int32
//go:linkname F_do_analyze_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_do_analyze_rel
func F_do_analyze_rel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_Async_Notify github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Async_Notify
func F_Async_Notify(m *base.Module, l0 int32, l1 int32)
//go:linkname F_asyncQueueReadAllNotifications github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_asyncQueueReadAllNotifications
func F_asyncQueueReadAllNotifications(m *base.Module)
//go:linkname F_ProcessNotifyInterrupt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ProcessNotifyInterrupt
func F_ProcessNotifyInterrupt(m *base.Module, l0 int32)
//go:linkname F_cluster_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cluster_rel
func F_cluster_rel(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_make_new_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_new_heap
func F_make_new_heap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_finish_heap_swap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_finish_heap_swap
func F_finish_heap_swap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_IsThereCollationInNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_IsThereCollationInNamespace
func F_IsThereCollationInNamespace(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateComments github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateComments
func F_CreateComments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CopyGetAttnums github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopyGetAttnums
func F_CopyGetAttnums(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ProcessCopyOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ProcessCopyOptions
func F_ProcessCopyOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CopyLoadRawBuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CopyLoadRawBuf
func F_CopyLoadRawBuf(m *base.Module, l0 int32)
//go:linkname F_CopyGetData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopyGetData
func F_CopyGetData(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CopySendEndOfRow github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CopySendEndOfRow
func F_CopySendEndOfRow(m *base.Module, l0 int32)
//go:linkname F_get_db_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_db_info
func F_get_db_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32) int32
//go:linkname F_errdetail_busy_db github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errdetail_busy_db
func F_errdetail_busy_db(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_database_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_database_name
func F_get_database_name(m *base.Module, l0 int32) int32
//go:linkname F_remove_dbtablespaces github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_remove_dbtablespaces
func F_remove_dbtablespaces(m *base.Module, l0 int32)
//go:linkname F_movedb_failure_callback github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_movedb_failure_callback
func F_movedb_failure_callback(m *base.Module, l0 int32, l1 int32)
//go:linkname F_defGetString github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_defGetString
func F_defGetString(m *base.Module, l0 int32) int32
//go:linkname F_defGetInt64 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_defGetInt64
func F_defGetInt64(m *base.Module, l0 int32) int64
//go:linkname F_defGetQualifiedName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_defGetQualifiedName
func F_defGetQualifiedName(m *base.Module, l0 int32) int32
//go:linkname F_defGetTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_defGetTypeName
func F_defGetTypeName(m *base.Module, l0 int32) int32
//go:linkname F_errorConflictingDefElem github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errorConflictingDefElem
func F_errorConflictingDefElem(m *base.Module, l0 int32, l1 int32)
//go:linkname F_DiscardCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DiscardCommand
func F_DiscardCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_EventTriggerCommonSetup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EventTriggerCommonSetup
func F_EventTriggerCommonSetup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_EventTriggerInvoke github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EventTriggerInvoke
func F_EventTriggerInvoke(m *base.Module, l0 int32, l1 int32)
//go:linkname F_standard_ExplainOneQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_standard_ExplainOneQuery
func F_standard_ExplainOneQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_ExplainOnePlan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExplainOnePlan
func F_ExplainOnePlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_ExplainPreScanNode github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExplainPreScanNode
func F_ExplainPreScanNode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExplainNode github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExplainNode
func F_ExplainNode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_report_triggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_report_triggers
func F_report_triggers(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainPropertyList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExplainPropertyList
func F_ExplainPropertyList(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainProperty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExplainProperty
func F_ExplainProperty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ExplainPropertyInteger github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainPropertyInteger
func F_ExplainPropertyInteger(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_ExplainOpenGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainOpenGroup
func F_ExplainOpenGroup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExplainCloseGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainCloseGroup
func F_ExplainCloseGroup(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainDummyGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExplainDummyGroup
func F_ExplainDummyGroup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExplainSeparatePlans github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExplainSeparatePlans
func F_ExplainSeparatePlans(m *base.Module, l0 int32)
//go:linkname F_get_extension_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_extension_oid
func F_get_extension_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_extension_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_extension_name
func F_get_extension_name(m *base.Module, l0 int32) int32
//go:linkname F_parse_extension_control_file github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parse_extension_control_file
func F_parse_extension_control_file(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_ext_ver_list github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_ext_ver_list
func F_get_ext_ver_list(m *base.Module, l0 int32) int32
//go:linkname F_InsertExtensionTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InsertExtensionTuple
func F_InsertExtensionTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_get_extension_control_directories github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_extension_control_directories
func F_get_extension_control_directories(m *base.Module) int32
//go:linkname F_heap_getattr_7 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_heap_getattr_7
func F_heap_getattr_7(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_transform_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_transform_oid
func F_get_transform_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IsThereFunctionInNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_IsThereFunctionInNamespace
func F_IsThereFunctionInNamespace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ComputeIndexAttrs github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ComputeIndexAttrs
func F_ComputeIndexAttrs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32)
//go:linkname F_ChooseRelationName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ChooseRelationName
func F_ChooseRelationName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_IndexSetParentIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_IndexSetParentIndex
func F_IndexSetParentIndex(m *base.Module, l0 int32, l1 int32)
//go:linkname F_makeObjectName github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeObjectName
func F_makeObjectName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LockViewRecurse_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LockViewRecurse_walker
func F_LockViewRecurse_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SetMatViewPopulatedState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SetMatViewPopulatedState
func F_SetMatViewPopulatedState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_opfamily_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_opfamily_oid
func F_get_opfamily_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IsThereOpClassInNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_IsThereOpClassInNamespace
func F_IsThereOpClassInNamespace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_IsThereOpFamilyInNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_IsThereOpFamilyInNamespace
func F_IsThereOpFamilyInNamespace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_EvaluateParams github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EvaluateParams
func F_EvaluateParams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_language_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_language_oid
func F_get_language_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tryAttachPartitionForeignKey github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tryAttachPartitionForeignKey
func F_tryAttachPartitionForeignKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32
//go:linkname F_renameatt_check github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_renameatt_check
func F_renameatt_check(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_find_typed_table_dependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_typed_table_dependencies
func F_find_typed_table_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_for_column_name_collision github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_check_for_column_name_collision
func F_check_for_column_name_collision(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rename_constraint_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_rename_constraint_internal
func F_rename_constraint_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_RenameRelationInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RenameRelationInternal
func F_RenameRelationInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ATPrepCmd github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ATPrepCmd
func F_ATPrepCmd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_find_composite_type_dependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_find_composite_type_dependencies
func F_find_composite_type_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PreCommit_on_commit_actions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreCommit_on_commit_actions
func F_PreCommit_on_commit_actions(m *base.Module)
//go:linkname F_AtEOXact_on_commit_actions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtEOXact_on_commit_actions
func F_AtEOXact_on_commit_actions(m *base.Module, l0 int32)
//go:linkname F_AlterConstrUpdateConstraintEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AlterConstrUpdateConstraintEntry
func F_AlterConstrUpdateConstraintEntry(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_createForeignKeyActionTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_createForeignKeyActionTriggers
func F_createForeignKeyActionTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_createForeignKeyCheckTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_createForeignKeyCheckTriggers
func F_createForeignKeyCheckTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_DropForeignKeyConstraintTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DropForeignKeyConstraintTriggers
func F_DropForeignKeyConstraintTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_AlterConstrEnforceabilityRecurse github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AlterConstrEnforceabilityRecurse
func F_AlterConstrEnforceabilityRecurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)
//go:linkname F_get_tablespace_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_tablespace_oid
func F_get_tablespace_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_directory_is_empty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_directory_is_empty
func F_directory_is_empty(m *base.Module, l0 int32) int32
//go:linkname F_GetDefaultTablespace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetDefaultTablespace
func F_GetDefaultTablespace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_PrepareTempTablespaces github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PrepareTempTablespaces
func F_PrepareTempTablespaces(m *base.Module)
//go:linkname F_get_tablespace_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_tablespace_name
func F_get_tablespace_name(m *base.Module, l0 int32) int32
//go:linkname F_CreateTriggerFiringOn github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateTriggerFiringOn
func F_CreateTriggerFiringOn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32)
//go:linkname F_renametrig_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_renametrig_internal
func F_renametrig_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_AfterTriggerSaveEvent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AfterTriggerSaveEvent
func F_AfterTriggerSaveEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)
//go:linkname F_afterTriggerMarkEvents github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_afterTriggerMarkEvents
func F_afterTriggerMarkEvents(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_afterTriggerInvokeEvents github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_afterTriggerInvokeEvents
func F_afterTriggerInvokeEvents(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_AfterTriggerFreeQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AfterTriggerFreeQuery
func F_AfterTriggerFreeQuery(m *base.Module, l0 int32)
//go:linkname F_AfterTriggerFireDeferred github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AfterTriggerFireDeferred
func F_AfterTriggerFireDeferred(m *base.Module)
//go:linkname F_AfterTriggerEndXact github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AfterTriggerEndXact
func F_AfterTriggerEndXact(m *base.Module)
//go:linkname F_deserialize_deflist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_deserialize_deflist
func F_deserialize_deflist(m *base.Module, l0 int32) int32
//go:linkname F_vacuum_is_permitted_for_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_vacuum_is_permitted_for_relation
func F_vacuum_is_permitted_for_relation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_vac_update_datfrozenxid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_vac_update_datfrozenxid
func F_vac_update_datfrozenxid(m *base.Module)
//go:linkname F_vacuum_open_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_vacuum_open_relation
func F_vacuum_open_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_vac_open_indexes github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vac_open_indexes
func F_vac_open_indexes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_vac_close_indexes github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_vac_close_indexes
func F_vac_close_indexes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_vacuum_delay_point github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vacuum_delay_point
func F_vacuum_delay_point(m *base.Module, l0 int32)
//go:linkname F_parallel_vacuum_process_one_index github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parallel_vacuum_process_one_index
func F_parallel_vacuum_process_one_index(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecReScan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecReScan
func F_ExecReScan(m *base.Module, l0 int32)
//go:linkname F_IndexSupportsBackwardScan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_IndexSupportsBackwardScan
func F_IndexSupportsBackwardScan(m *base.Module, l0 int32) int32
//go:linkname F_ExecAsyncRequest github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecAsyncRequest
func F_ExecAsyncRequest(m *base.Module, l0 int32)
//go:linkname F_ExecInitExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitExpr
func F_ExecInitExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecInitQual github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecInitQual
func F_ExecInitQual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecBuildProjectionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecBuildProjectionInfo
func F_ExecBuildProjectionInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ExecBuildUpdateProjection github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecBuildUpdateProjection
func F_ExecBuildUpdateProjection(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ExecInterpExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecInterpExpr
func F_ExecInterpExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecInitJunkFilter github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecInitJunkFilter
func F_ExecInitJunkFilter(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecutorStart github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecutorStart
func F_ExecutorStart(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecCheckPermissions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecCheckPermissions
func F_ExecCheckPermissions(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecutorRun github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecutorRun
func F_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_standard_ExecutorRun github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_standard_ExecutorRun
func F_standard_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_ExecutorFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecutorFinish
func F_ExecutorFinish(m *base.Module, l0 int32)
//go:linkname F_ExecutorEnd github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecutorEnd
func F_ExecutorEnd(m *base.Module, l0 int32)
//go:linkname F_ExecCheckPermissionsModified github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecCheckPermissionsModified
func F_ExecCheckPermissionsModified(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_InitResultRelInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InitResultRelInfo
func F_InitResultRelInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ExecPartitionCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecPartitionCheck
func F_ExecPartitionCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_EvalPlanQualInit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EvalPlanQualInit
func F_EvalPlanQualInit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_EvalPlanQualEnd github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EvalPlanQualEnd
func F_EvalPlanQualEnd(m *base.Module, l0 int32)
//go:linkname F_ExecParallelEstimate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecParallelEstimate
func F_ExecParallelEstimate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecParallelSetupTupleQueues github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecParallelSetupTupleQueues
func F_ExecParallelSetupTupleQueues(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SerializeParamExecParams github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SerializeParamExecParams
func F_SerializeParamExecParams(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecParallelInitializeDSM github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecParallelInitializeDSM
func F_ExecParallelInitializeDSM(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecParallelReInitializeDSM github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecParallelReInitializeDSM
func F_ExecParallelReInitializeDSM(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecSetupPartitionTupleRouting github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecSetupPartitionTupleRouting
func F_ExecSetupPartitionTupleRouting(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecFindPartition github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecFindPartition
func F_ExecFindPartition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_InitPartitionPruneContext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InitPartitionPruneContext
func F_InitPartitionPruneContext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ExecFindMatchingSubPlans github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecFindMatchingSubPlans
func F_ExecFindMatchingSubPlans(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecInitNode github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecInitNode
func F_ExecInitNode(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MultiExecProcNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MultiExecProcNode
func F_MultiExecProcNode(m *base.Module, l0 int32) int32
//go:linkname F_RelationFindReplTupleByIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationFindReplTupleByIndex
func F_RelationFindReplTupleByIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_RelationFindReplTupleSeq github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationFindReplTupleSeq
func F_RelationFindReplTupleSeq(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecSimpleRelationInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecSimpleRelationInsert
func F_ExecSimpleRelationInsert(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecSimpleRelationUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecSimpleRelationUpdate
func F_ExecSimpleRelationUpdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ExecSimpleRelationDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecSimpleRelationDelete
func F_ExecSimpleRelationDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CheckSubscriptionRelkind github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckSubscriptionRelkind
func F_CheckSubscriptionRelkind(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecScan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecScan
func F_ExecScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MakeTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MakeTupleTableSlot
func F_MakeTupleTableSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecDropSingleTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecDropSingleTupleTableSlot
func F_ExecDropSingleTupleTableSlot(m *base.Module, l0 int32)
//go:linkname F_ExecStoreHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecStoreHeapTuple
func F_ExecStoreHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecStoreBufferHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecStoreBufferHeapTuple
func F_ExecStoreBufferHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecStorePinnedBufferHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecStorePinnedBufferHeapTuple
func F_ExecStorePinnedBufferHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecStoreAllNullTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecStoreAllNullTuple
func F_ExecStoreAllNullTuple(m *base.Module, l0 int32)
//go:linkname F_ExecFetchSlotHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecFetchSlotHeapTuple
func F_ExecFetchSlotHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecFetchSlotMinimalTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecFetchSlotMinimalTuple
func F_ExecFetchSlotMinimalTuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecInitExtraTupleSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecInitExtraTupleSlot
func F_ExecInitExtraTupleSlot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_slot_getsomeattrs_int github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_slot_getsomeattrs_int
func F_slot_getsomeattrs_int(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecCleanTypeFromTL github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecCleanTypeFromTL
func F_ExecCleanTypeFromTL(m *base.Module, l0 int32) int32
//go:linkname F_BlessTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BlessTupleDesc
func F_BlessTupleDesc(m *base.Module, l0 int32) int32
//go:linkname F_TupleDescGetAttInMetadata github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TupleDescGetAttInMetadata
func F_TupleDescGetAttInMetadata(m *base.Module, l0 int32) int32
//go:linkname F_BuildTupleFromCStrings github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BuildTupleFromCStrings
func F_BuildTupleFromCStrings(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_HeapTupleHeaderGetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HeapTupleHeaderGetDatum
func F_HeapTupleHeaderGetDatum(m *base.Module, l0 int32) int32
//go:linkname F_CreateExecutorState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateExecutorState
func F_CreateExecutorState(m *base.Module) int32
//go:linkname F_FreeExecutorState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FreeExecutorState
func F_FreeExecutorState(m *base.Module, l0 int32)
//go:linkname F_FreeExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FreeExprContext
func F_FreeExprContext(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateExprContext
func F_CreateExprContext(m *base.Module, l0 int32) int32
//go:linkname F_MakePerTupleExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MakePerTupleExprContext
func F_MakePerTupleExprContext(m *base.Module, l0 int32) int32
//go:linkname F_ExecAssignExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecAssignExprContext
func F_ExecAssignExprContext(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecGetRangeTableRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecGetRangeTableRelation
func F_ExecGetRangeTableRelation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecInitRangeTable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitRangeTable
func F_ExecInitRangeTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_UpdateChangedParamSet github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UpdateChangedParamSet
func F_UpdateChangedParamSet(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecGetRootToChildMap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetRootToChildMap
func F_ExecGetRootToChildMap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sql_fn_resolve_param_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sql_fn_resolve_param_name
func F_sql_fn_resolve_param_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_BufferUsageAccumDiff github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BufferUsageAccumDiff
func F_BufferUsageAccumDiff(m *base.Module, l0 int32, l1 int32)
//go:linkname F_advance_transition_function github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_advance_transition_function
func F_advance_transition_function(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecAppendAsyncEventWait github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecAppendAsyncEventWait
func F_ExecAppendAsyncEventWait(m *base.Module, l0 int32)
//go:linkname F_ExecHashTableInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecHashTableInsert
func F_ExecHashTableInsert(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecParallelHashJoinSetUpBatches github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecParallelHashJoinSetUpBatches
func F_ExecParallelHashJoinSetUpBatches(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecHashJoinGetSavedTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecHashJoinGetSavedTuple
func F_ExecHashJoinGetSavedTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MemoizeHash_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoizeHash_hash
func F_MemoizeHash_hash(m *base.Module, l0 int32) int32
//go:linkname F_MemoizeHash_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoizeHash_equal
func F_MemoizeHash_equal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecSetParamPlanMulti github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecSetParamPlanMulti
func F_ExecSetParamPlanMulti(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tfuncFetchRows github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tfuncFetchRows
func F_tfuncFetchRows(m *base.Module, l0 int32, l1 int32)
//go:linkname F_TidRangeEval github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TidRangeEval
func F_TidRangeEval(m *base.Module, l0 int32) int32
//go:linkname F_WinGetPartitionLocalMemory github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WinGetPartitionLocalMemory
func F_WinGetPartitionLocalMemory(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_WinGetPartitionRowCount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WinGetPartitionRowCount
func F_WinGetPartitionRowCount(m *base.Module, l0 int32) int64
//go:linkname F_WinGetFuncArgInPartition github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WinGetFuncArgInPartition
func F_WinGetFuncArgInPartition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SPI_connect_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SPI_connect_ext
func F_SPI_connect_ext(m *base.Module, l0 int32)
//go:linkname F_SPI_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_finish
func F_SPI_finish(m *base.Module) int32
//go:linkname F_AtEOXact_SPI github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AtEOXact_SPI
func F_AtEOXact_SPI(m *base.Module, l0 int32)
//go:linkname F_SPI_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_execute
func F_SPI_execute(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__SPI_execute_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__SPI_execute_plan
func F__SPI_execute_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SPI_freetuptable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_freetuptable
func F_SPI_freetuptable(m *base.Module, l0 int32)
//go:linkname F_SPI_exec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_exec
func F_SPI_exec(m *base.Module, l0 int32) int32
//go:linkname F_SPI_prepare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_prepare
func F_SPI_prepare(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_freeplan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_freeplan
func F_SPI_freeplan(m *base.Module, l0 int32)
//go:linkname F_SPI_getvalue github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_getvalue
func F_SPI_getvalue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_getbinval github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SPI_getbinval
func F_SPI_getbinval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SPI_gettypeid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SPI_gettypeid
func F_SPI_gettypeid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SPI_palloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_palloc
func F_SPI_palloc(m *base.Module, l0 int32) int32
//go:linkname F_SPI_cursor_open github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_cursor_open
func F_SPI_cursor_open(m *base.Module, l0 int32) int32
//go:linkname F_SPI_cursor_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SPI_cursor_fetch
func F_SPI_cursor_fetch(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SPI_cursor_close github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_cursor_close
func F_SPI_cursor_close(m *base.Module, l0 int32)
//go:linkname F_get_foreign_data_wrapper_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_foreign_data_wrapper_oid
func F_get_foreign_data_wrapper_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetForeignServerExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetForeignServerExtended
func F_GetForeignServerExtended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_foreign_server_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_foreign_server_oid
func F_get_foreign_server_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetFdwRoutineForRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetFdwRoutineForRelation
func F_GetFdwRoutineForRelation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dshash_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dshash_attach
func F_dshash_attach(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_dshash_seq_term github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_dshash_seq_term
func F_dshash_seq_term(m *base.Module, l0 int32)
//go:linkname F_initHyperLogLog github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_initHyperLogLog
func F_initHyperLogLog(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pairingheap_add github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pairingheap_add
func F_pairingheap_add(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pairingheap_remove_first github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pairingheap_remove_first
func F_pairingheap_remove_first(m *base.Module, l0 int32) int32
//go:linkname F_pairingheap_remove github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pairingheap_remove
func F_pairingheap_remove(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AtEOXact_LargeObject github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AtEOXact_LargeObject
func F_AtEOXact_LargeObject(m *base.Module, l0 int32)
//go:linkname F_lo_get_fragment_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lo_get_fragment_internal
func F_lo_get_fragment_internal(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_secure_read github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_secure_read
func F_secure_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_password_type github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_password_type
func F_get_password_type(m *base.Module, l0 int32) int32
//go:linkname F_open_auth_file github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_open_auth_file
func F_open_auth_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_tokenize_include_file github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tokenize_include_file
func F_tokenize_include_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_parse_ident_line github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_parse_ident_line
func F_parse_ident_line(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_getkeepalivescount github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_getkeepalivescount
func F_pq_getkeepalivescount(m *base.Module, l0 int32) int32
//go:linkname F_internal_flush_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_internal_flush_buffer
func F_internal_flush_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pq_sendfloat8 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_sendfloat8
func F_pq_sendfloat8(m *base.Module, l0 int32, l1 float64)
//go:linkname F_pq_begintypsend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_begintypsend
func F_pq_begintypsend(m *base.Module, l0 int32)
//go:linkname F_pq_getmsgint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pq_getmsgint
func F_pq_getmsgint(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_copymsgbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_copymsgbytes
func F_pq_copymsgbytes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_getmsgint64 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pq_getmsgint64
func F_pq_getmsgint64(m *base.Module, l0 int32) int64
//go:linkname F_pq_getmsgfloat8 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pq_getmsgfloat8
func F_pq_getmsgfloat8(m *base.Module, l0 int32) float64
//go:linkname F_pq_getmsgbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_getmsgbytes
func F_pq_getmsgbytes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_getmsgtext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_getmsgtext
func F_pq_getmsgtext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pq_getmsgend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_getmsgend
func F_pq_getmsgend(m *base.Module, l0 int32)
//go:linkname F_bms_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_copy
func F_bms_copy(m *base.Module, l0 int32) int32
//go:linkname F_bms_make_singleton github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bms_make_singleton
func F_bms_make_singleton(m *base.Module, l0 int32) int32
//go:linkname F_bms_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_union
func F_bms_union(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_intersect github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_intersect
func F_bms_intersect(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_is_subset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_is_subset
func F_bms_is_subset(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_is_member github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_is_member
func F_bms_is_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_member_index github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_member_index
func F_bms_member_index(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_overlap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bms_overlap
func F_bms_overlap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_add_members github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bms_add_members
func F_bms_add_members(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_del_members github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_del_members
func F_bms_del_members(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_copyObjectImpl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_copyObjectImpl
func F_copyObjectImpl(m *base.Module, l0 int32) int32
//go:linkname F_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_equal
func F_equal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_make1_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_make1_impl
func F_list_make1_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_make2_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_list_make2_impl
func F_list_make2_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lappend
func F_lappend(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend_int github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_lappend_int
func F_lappend_int(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_lappend_oid
func F_lappend_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_new_head_cell github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_new_head_cell
func F_new_head_cell(m *base.Module, l0 int32)
//go:linkname F_list_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_concat
func F_list_concat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_list_copy
func F_list_copy(m *base.Module, l0 int32) int32
//go:linkname F_list_concat_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_concat_copy
func F_list_concat_copy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_member github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_member
func F_list_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_member_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_member_ptr
func F_list_member_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_nth_cell github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_delete_nth_cell
func F_list_delete_nth_cell(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_delete_ptr
func F_list_delete_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_first github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_delete_first
func F_list_delete_first(m *base.Module, l0 int32) int32
//go:linkname F_list_delete_last github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_delete_last
func F_list_delete_last(m *base.Module, l0 int32) int32
//go:linkname F_list_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_list_union
func F_list_union(m *base.Module, l0 int32) int32
//go:linkname F_list_append_unique_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_append_unique_ptr
func F_list_append_unique_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_concat_unique_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_concat_unique_oid
func F_list_concat_unique_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_free_deep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_free_deep
func F_list_free_deep(m *base.Module, l0 int32)
//go:linkname F_list_copy_head github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_copy_head
func F_list_copy_head(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeSimpleA_Expr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeSimpleA_Expr
func F_makeSimpleA_Expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeVar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeVar
func F_makeVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_makeVarFromTargetEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeVarFromTargetEntry
func F_makeVarFromTargetEntry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeConst github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeConst
func F_makeConst(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_makeNullConst github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeNullConst
func F_makeNullConst(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeAlias github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeAlias
func F_makeAlias(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeRangeVar github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeRangeVar
func F_makeRangeVar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeNotNullConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeNotNullConstraint
func F_makeNotNullConstraint(m *base.Module, l0 int32) int32
//go:linkname F_makeTypeNameFromNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeTypeNameFromNameList
func F_makeTypeNameFromNameList(m *base.Module, l0 int32) int32
//go:linkname F_makeFuncExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeFuncExpr
func F_makeFuncExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeStringConst github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeStringConst
func F_makeStringConst(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeFuncCall
func F_makeFuncCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_opclause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_opclause
func F_make_opclause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_and_qual github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_and_qual
func F_make_and_qual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_ands_explicit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_ands_explicit
func F_make_ands_explicit(m *base.Module, l0 int32) int32
//go:linkname F_make_ands_implicit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_ands_implicit
func F_make_ands_implicit(m *base.Module, l0 int32) int32
//go:linkname F_makeIndexInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeIndexInfo
func F_makeIndexInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32
//go:linkname F_makeGroupingSet github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeGroupingSet
func F_makeGroupingSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeVacuumRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeVacuumRelation
func F_makeVacuumRelation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeJsonValueExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeJsonValueExpr
func F_makeJsonValueExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_exprType github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_exprType
func F_exprType(m *base.Module, l0 int32) int32
//go:linkname F_exprTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprTypmod
func F_exprTypmod(m *base.Module, l0 int32) int32
//go:linkname F_exprCollation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprCollation
func F_exprCollation(m *base.Module, l0 int32) int32
//go:linkname F_relabel_to_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_relabel_to_typmod
func F_relabel_to_typmod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expression_returns_set github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expression_returns_set
func F_expression_returns_set(m *base.Module, l0 int32) int32
//go:linkname F_exprSetCollation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprSetCollation
func F_exprSetCollation(m *base.Module, l0 int32, l1 int32)
//go:linkname F_exprLocation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprLocation
func F_exprLocation(m *base.Module, l0 int32) int32
//go:linkname F_fix_opfuncids github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fix_opfuncids
func F_fix_opfuncids(m *base.Module, l0 int32)
//go:linkname F_set_opfuncid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_opfuncid
func F_set_opfuncid(m *base.Module, l0 int32)
//go:linkname F_check_functions_in_node github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_functions_in_node
func F_check_functions_in_node(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_expression_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expression_tree_walker_impl
func F_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_query_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_query_tree_walker_impl
func F_query_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_query_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_query_tree_mutator_impl
func F_query_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_query_or_expression_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_query_or_expression_tree_walker_impl
func F_query_or_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_raw_expression_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_raw_expression_tree_walker_impl
func F_raw_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_nodeToString github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_nodeToString
func F_nodeToString(m *base.Module, l0 int32) int32
//go:linkname F_elog_node_display github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_elog_node_display
func F_elog_node_display(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_JumbleQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_JumbleQuery
func F_JumbleQuery(m *base.Module, l0 int32) int32
//go:linkname F__jumbleNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__jumbleNode
func F__jumbleNode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AppendJumble32 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AppendJumble32
func F_AppendJumble32(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AppendJumble8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AppendJumble8
func F_AppendJumble8(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AppendJumble github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AppendJumble
func F_AppendJumble(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_stringToNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_stringToNode
func F_stringToNode(m *base.Module, l0 int32) int32
//go:linkname F_pg_strtok github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_strtok
func F_pg_strtok(m *base.Module, l0 int32) int32
//go:linkname F_tbm_create github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tbm_create
func F_tbm_create(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tbm_free github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tbm_free
func F_tbm_free(m *base.Module, l0 int32)
//go:linkname F_tbm_add_tuples github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tbm_add_tuples
func F_tbm_add_tuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_tbm_add_page github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tbm_add_page
func F_tbm_add_page(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tbm_begin_private_iterate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tbm_begin_private_iterate
func F_tbm_begin_private_iterate(m *base.Module, l0 int32) int32
//go:linkname F_makeString github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeString
func F_makeString(m *base.Module, l0 int32) int32
//go:linkname F_gimme_tree github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gimme_tree
func F_gimme_tree(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_add_paths_to_append_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_paths_to_append_rel
func F_add_paths_to_append_rel(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_generate_useful_gather_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_generate_useful_gather_paths
func F_generate_useful_gather_paths(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_recurse_push_qual github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_recurse_push_qual
func F_recurse_push_qual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_clause_selectivity_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_clause_selectivity_ext
func F_clause_selectivity_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64
//go:linkname F_find_single_rel_for_clauses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_find_single_rel_for_clauses
func F_find_single_rel_for_clauses(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cost_seqscan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cost_seqscan
func F_cost_seqscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_cost_qual_eval_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cost_qual_eval_walker
func F_cost_qual_eval_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cost_bitmap_heap_scan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cost_bitmap_heap_scan
func F_cost_bitmap_heap_scan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64)
//go:linkname F_cost_bitmap_tree_node github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cost_bitmap_tree_node
func F_cost_bitmap_tree_node(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_cost_incremental_sort github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cost_incremental_sort
func F_cost_incremental_sort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64, l6 float64, l7 float64, l8 int32, l9 int32, l10 float64)
//go:linkname F_set_joinrel_size_estimates github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_joinrel_size_estimates
func F_set_joinrel_size_estimates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_set_pathtarget_cost_width github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_pathtarget_cost_width
func F_set_pathtarget_cost_width(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ec_clear_derived_clauses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ec_clear_derived_clauses
func F_ec_clear_derived_clauses(m *base.Module, l0 int32)
//go:linkname F_get_eclass_for_sort_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_eclass_for_sort_expr
func F_get_eclass_for_sort_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_generate_join_implied_equalities github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_join_implied_equalities
func F_generate_join_implied_equalities(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_rebuild_eclass_attr_needed github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_rebuild_eclass_attr_needed
func F_rebuild_eclass_attr_needed(m *base.Module, l0 int32)
//go:linkname F_get_loop_count github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_loop_count
func F_get_loop_count(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_find_indexpath_quals github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_indexpath_quals
func F_find_indexpath_quals(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_match_boolean_index_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_match_boolean_index_clause
func F_match_boolean_index_clause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_try_mergejoin_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_try_mergejoin_path
func F_try_mergejoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32)
//go:linkname F_has_legal_joinclause github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_has_legal_joinclause
func F_has_legal_joinclause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_join_is_legal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_join_is_legal
func F_join_is_legal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_populate_joinrel_with_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_populate_joinrel_with_paths
func F_populate_joinrel_with_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_mark_dummy_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mark_dummy_rel
func F_mark_dummy_rel(m *base.Module, l0 int32)
//go:linkname F_make_canonical_pathkey github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_canonical_pathkey
func F_make_canonical_pathkey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_append_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_append_pathkeys
func F_append_pathkeys(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_truncate_useless_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_truncate_useless_pathkeys
func F_truncate_useless_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_pathkeys_for_sortclauses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_pathkeys_for_sortclauses
func F_make_pathkeys_for_sortclauses(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_pathkeys_for_sortclauses_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_pathkeys_for_sortclauses_extended
func F_make_pathkeys_for_sortclauses_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_find_mergeclauses_for_outer_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_find_mergeclauses_for_outer_pathkeys
func F_find_mergeclauses_for_outer_pathkeys(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_inner_pathkeys_for_merge github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_inner_pathkeys_for_merge
func F_make_inner_pathkeys_for_merge(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_trim_mergeclauses_for_inner_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_trim_mergeclauses_for_inner_pathkeys
func F_trim_mergeclauses_for_inner_pathkeys(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_remove_rel_from_query github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_remove_rel_from_query
func F_remove_rel_from_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_innerrel_is_unique_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_innerrel_is_unique_ext
func F_innerrel_is_unique_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_add_non_redundant_clauses github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_non_redundant_clauses
func F_add_non_redundant_clauses(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_rebuild_lateral_attr_needed github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_rebuild_lateral_attr_needed
func F_rebuild_lateral_attr_needed(m *base.Module, l0 int32)
//go:linkname F_restriction_is_always_true github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_restriction_is_always_true
func F_restriction_is_always_true(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_restriction_is_always_false github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_restriction_is_always_false
func F_restriction_is_always_false(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rebuild_joinclause_attr_needed github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_rebuild_joinclause_attr_needed
func F_rebuild_joinclause_attr_needed(m *base.Module, l0 int32)
//go:linkname F_query_planner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_query_planner
func F_query_planner(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_adjust_paths_for_srfs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_adjust_paths_for_srfs
func F_adjust_paths_for_srfs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_make_pathkeys_for_window github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_pathkeys_for_window
func F_make_pathkeys_for_window(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_expression_planner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expression_planner
func F_expression_planner(m *base.Module, l0 int32) int32
//go:linkname F_set_plan_refs github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_set_plan_refs
func F_set_plan_refs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_add_rte_to_flat_rtable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_rte_to_flat_rtable
func F_add_rte_to_flat_rtable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_trivial_subqueryscan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_trivial_subqueryscan
func F_trivial_subqueryscan(m *base.Module, l0 int32) int32
//go:linkname F_extract_query_dependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_extract_query_dependencies
func F_extract_query_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_fix_scan_expr_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fix_scan_expr_mutator
func F_fix_scan_expr_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fix_scan_expr_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fix_scan_expr_walker
func F_fix_scan_expr_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SS_identify_outer_params github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SS_identify_outer_params
func F_SS_identify_outer_params(m *base.Module, l0 int32)
//go:linkname F_canonicalize_qual github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_canonicalize_qual
func F_canonicalize_qual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_adjust_appendrel_attrs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_adjust_appendrel_attrs
func F_adjust_appendrel_attrs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_find_appinfos_by_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_appinfos_by_relids
func F_find_appinfos_by_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_add_row_identity_var github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_row_identity_var
func F_add_row_identity_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_contain_mutable_functions_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_contain_mutable_functions_walker
func F_contain_mutable_functions_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_contain_mutable_functions_after_planning github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_contain_mutable_functions_after_planning
func F_contain_mutable_functions_after_planning(m *base.Module, l0 int32) int32
//go:linkname F_contain_volatile_functions_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_contain_volatile_functions_walker
func F_contain_volatile_functions_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_is_parallel_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_is_parallel_safe
func F_is_parallel_safe(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_is_pseudo_constant_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_is_pseudo_constant_clause
func F_is_pseudo_constant_clause(m *base.Module, l0 int32) int32
//go:linkname F_is_pseudo_constant_clause_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_is_pseudo_constant_clause_relids
func F_is_pseudo_constant_clause_relids(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_NumRelids github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_NumRelids
func F_NumRelids(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_eval_const_expressions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_eval_const_expressions
func F_eval_const_expressions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generate_new_exec_param github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_new_exec_param
func F_generate_new_exec_param(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_compare_path_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_compare_path_costs
func F_compare_path_costs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_compare_fractional_path_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_compare_fractional_path_costs
func F_compare_fractional_path_costs(m *base.Module, l0 int32, l1 int32, l2 float64) int32
//go:linkname F_set_cheapest github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_cheapest
func F_set_cheapest(m *base.Module, l0 int32)
//go:linkname F_create_bitmap_and_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_bitmap_and_path
func F_create_bitmap_and_path(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_create_projection_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_projection_path
func F_create_projection_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_apply_projection_to_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_apply_projection_to_path
func F_apply_projection_to_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_find_placeholder_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_placeholder_info
func F_find_placeholder_info(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rebuild_placeholder_attr_needed github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_rebuild_placeholder_attr_needed
func F_rebuild_placeholder_attr_needed(m *base.Module, l0 int32)
//go:linkname F_estimate_rel_size github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_estimate_rel_size
func F_estimate_rel_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_get_relation_statistics_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_relation_statistics_worker
func F_get_relation_statistics_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_predicate_implied_by github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_predicate_implied_by
func F_predicate_implied_by(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_predicate_refuted_by_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_predicate_refuted_by_recurse
func F_predicate_refuted_by_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lookup_proof_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lookup_proof_cache
func F_lookup_proof_cache(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_setup_simple_rel_arrays github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_setup_simple_rel_arrays
func F_setup_simple_rel_arrays(m *base.Module, l0 int32)
//go:linkname F_find_base_rel_ignore_join github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_find_base_rel_ignore_join
func F_find_base_rel_ignore_join(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_join_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_join_rel
func F_find_join_rel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_build_joinrel_restrictlist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_joinrel_restrictlist
func F_build_joinrel_restrictlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_build_joinrel_tlist github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_build_joinrel_tlist
func F_build_joinrel_tlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_build_joinrel_partition_info github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_joinrel_partition_info
func F_build_joinrel_partition_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_baserel_parampathinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_baserel_parampathinfo
func F_get_baserel_parampathinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_restrictinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_restrictinfo
func F_make_restrictinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_tlist_same_datatypes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tlist_same_datatypes
func F_tlist_same_datatypes(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_copy_pathtarget github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_copy_pathtarget
func F_copy_pathtarget(m *base.Module, l0 int32) int32
//go:linkname F_create_empty_pathtarget github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_empty_pathtarget
func F_create_empty_pathtarget(m *base.Module) int32
//go:linkname F_pull_varattnos github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pull_varattnos
func F_pull_varattnos(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_contain_var_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_contain_var_clause
func F_contain_var_clause(m *base.Module, l0 int32) int32
//go:linkname F_contain_vars_of_level github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_contain_vars_of_level
func F_contain_vars_of_level(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_locate_var_of_level github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_locate_var_of_level
func F_locate_var_of_level(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pull_var_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pull_var_clause
func F_pull_var_clause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_flatten_join_alias_vars_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_flatten_join_alias_vars_mutator
func F_flatten_join_alias_vars_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_flatten_group_exprs_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_flatten_group_exprs_mutator
func F_flatten_group_exprs_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_partition_op_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_partition_op_expr
func F_make_partition_op_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_RelationGetPartitionDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetPartitionDesc
func F_RelationGetPartitionDesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CreatePartitionDirectory github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CreatePartitionDirectory
func F_CreatePartitionDirectory(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rebuild_database_list github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_rebuild_database_list
func F_rebuild_database_list(m *base.Module, l0 int32)
//go:linkname F_VacuumUpdateCosts github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_VacuumUpdateCosts
func F_VacuumUpdateCosts(m *base.Module)
//go:linkname F_BackgroundWorkerUnblockSignals github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BackgroundWorkerUnblockSignals
func F_BackgroundWorkerUnblockSignals(m *base.Module)
//go:linkname F_SanityCheckBackgroundWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SanityCheckBackgroundWorker
func F_SanityCheckBackgroundWorker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RegisterDynamicBackgroundWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RegisterDynamicBackgroundWorker
func F_RegisterDynamicBackgroundWorker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetBackgroundWorkerPid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetBackgroundWorkerPid
func F_GetBackgroundWorkerPid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_postmaster_child_launch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_postmaster_child_launch
func F_postmaster_child_launch(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AssignPostmasterChildSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AssignPostmasterChildSlot
func F_AssignPostmasterChildSlot(m *base.Module, l0 int32) int32
//go:linkname F_ReleasePostmasterChildSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReleasePostmasterChildSlot
func F_ReleasePostmasterChildSlot(m *base.Module, l0 int32) int32
//go:linkname F_ExitPostmaster github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExitPostmaster
func F_ExitPostmaster(m *base.Module, l0 int32)
//go:linkname F_ProcessStartupProcInterrupts github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessStartupProcInterrupts
func F_ProcessStartupProcInterrupts(m *base.Module)
//go:linkname F_pg_set_regex_collation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_set_regex_collation
func F_pg_set_regex_collation(m *base.Module, l0 int32)
//go:linkname F_newnfa github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_newnfa
func F_newnfa(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_subcolor github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_subcolor
func F_subcolor(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_okcolors github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_okcolors
func F_okcolors(m *base.Module, l0 int32, l1 int32)
//go:linkname F_parse github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_parse
func F_parse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_specialcolors github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_specialcolors
func F_specialcolors(m *base.Module, l0 int32)
//go:linkname F_numst github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_numst
func F_numst(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_markst github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_markst
func F_markst(m *base.Module, l0 int32)
//go:linkname F_nfatree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_nfatree
func F_nfatree(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_nfanode github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_nfanode
func F_nfanode(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_optimize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_optimize
func F_optimize(m *base.Module, l0 int32) int32
//go:linkname F_makesearch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makesearch
func F_makesearch(m *base.Module, l0 int32, l1 int32)
//go:linkname F_compact github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_compact
func F_compact(m *base.Module, l0 int32, l1 int32)
//go:linkname F_rfree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_rfree
func F_rfree(m *base.Module, l0 int32)
//go:linkname F_createarc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_createarc
func F_createarc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_next github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_next
func F_next(m *base.Module, l0 int32) int32
//go:linkname F_pg_regerror github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_regerror
func F_pg_regerror(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_regexec github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_regexec
func F_pg_regexec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_pg_reg_getnumoutarcs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_reg_getnumoutarcs
func F_pg_reg_getnumoutarcs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_reg_getoutarcs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_reg_getoutarcs
func F_pg_reg_getoutarcs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetTupleTransactionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetTupleTransactionInfo
func F_GetTupleTransactionInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ReportApplyConflict github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReportApplyConflict
func F_ReportApplyConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_InitConflictIndexes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_InitConflictIndexes
func F_InitConflictIndexes(m *base.Module, l0 int32)
//go:linkname F_LogicalDecodingProcessRecord github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LogicalDecodingProcessRecord
func F_LogicalDecodingProcessRecord(m *base.Module, l0 int32, l1 int32)
//go:linkname F_logicalrep_workers_find github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_logicalrep_workers_find
func F_logicalrep_workers_find(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_logicalrep_worker_wakeup_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_logicalrep_worker_wakeup_ptr
func F_logicalrep_worker_wakeup_ptr(m *base.Module, l0 int32)
//go:linkname F_logicalrep_worker_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_logicalrep_worker_attach
func F_logicalrep_worker_attach(m *base.Module, l0 int32)
//go:linkname F_AtEOXact_ApplyLauncher github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AtEOXact_ApplyLauncher
func F_AtEOXact_ApplyLauncher(m *base.Module, l0 int32)
//go:linkname F_pg_logical_emit_message_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_logical_emit_message_bytea
func F_pg_logical_emit_message_bytea(m *base.Module, l0 int32) int32
//go:linkname F_replorigin_session_advance github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_replorigin_session_advance
func F_replorigin_session_advance(m *base.Module, l0 int64, l1 int64)
//go:linkname F_logicalrep_rel_mark_updatable github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_logicalrep_rel_mark_updatable
func F_logicalrep_rel_mark_updatable(m *base.Module, l0 int32)
//go:linkname F_FindLogicalRepLocalIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FindLogicalRepLocalIndex
func F_FindLogicalRepLocalIndex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReorderBufferAllocChange github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferAllocChange
func F_ReorderBufferAllocChange(m *base.Module, l0 int32) int32
//go:linkname F_ReorderBufferAllocTupleBuf github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReorderBufferAllocTupleBuf
func F_ReorderBufferAllocTupleBuf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReorderBufferQueueChange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReorderBufferQueueChange
func F_ReorderBufferQueueChange(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32)
//go:linkname F_ReorderBufferTXNByXid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReorderBufferTXNByXid
func F_ReorderBufferTXNByXid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_ReorderBufferRestoreCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReorderBufferRestoreCleanup
func F_ReorderBufferRestoreCleanup(m *base.Module, l0 int32)
//go:linkname F_ReorderBufferProcessXid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReorderBufferProcessXid
func F_ReorderBufferProcessXid(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_ResolveCminCmaxDuringDecoding github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ResolveCminCmaxDuringDecoding
func F_ResolveCminCmaxDuringDecoding(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_SnapBuildSnapDecRefcount github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SnapBuildSnapDecRefcount
func F_SnapBuildSnapDecRefcount(m *base.Module, l0 int32)
//go:linkname F_SnapBuildProcessChange github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SnapBuildProcessChange
func F_SnapBuildProcessChange(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_SnapBuildRestore github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SnapBuildRestore
func F_SnapBuildRestore(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_SnapBuildSerialize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SnapBuildSerialize
func F_SnapBuildSerialize(m *base.Module, l0 int32, l1 int64)
//go:linkname F_slot_store_data github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_slot_store_data
func F_slot_store_data(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_TargetPrivilegesCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TargetPrivilegesCheck
func F_TargetPrivilegesCheck(m *base.Module, l0 int32, l1 int64)
//go:linkname F_check_relation_updatable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_relation_updatable
func F_check_relation_updatable(m *base.Module, l0 int32)
//go:linkname F_slot_modify_data github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_slot_modify_data
func F_slot_modify_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_apply_handle_delete_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_apply_handle_delete_internal
func F_apply_handle_delete_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_InitializeLogRepWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InitializeLogRepWorker
func F_InitializeLogRepWorker(m *base.Module)
//go:linkname F_LogicalRepWorkersWakeupAtCommit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LogicalRepWorkersWakeupAtCommit
func F_LogicalRepWorkersWakeupAtCommit(m *base.Module, l0 int32)
//go:linkname F_yy_fatal_error_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_yy_fatal_error_3
func F_yy_fatal_error_3(m *base.Module, l0 int32)
//go:linkname F_replication_yylex_destroy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_replication_yylex_destroy
func F_replication_yylex_destroy(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReplicationSlotRelease
func F_ReplicationSlotRelease(m *base.Module)
//go:linkname F_ReplicationSlotCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplicationSlotCleanup
func F_ReplicationSlotCleanup(m *base.Module, l0 int32)
//go:linkname F_SaveSlotToPath github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SaveSlotToPath
func F_SaveSlotToPath(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReplicationSlotAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplicationSlotAcquire
func F_ReplicationSlotAcquire(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReplicationSlotDropPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReplicationSlotDropPtr
func F_ReplicationSlotDropPtr(m *base.Module, l0 int32)
//go:linkname F_InvalidateObsoleteReplicationSlots github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InvalidateObsoleteReplicationSlots
func F_InvalidateObsoleteReplicationSlots(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_SyncRepWaitForLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SyncRepWaitForLSN
func F_SyncRepWaitForLSN(m *base.Module, l0 int64, l1 int32)
//go:linkname F_SyncRepInitConfig github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SyncRepInitConfig
func F_SyncRepInitConfig(m *base.Module)
//go:linkname F_XLogWalRcvSendReply github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogWalRcvSendReply
func F_XLogWalRcvSendReply(m *base.Module, l0 int32, l1 int32)
//go:linkname F_WalRcvRunning github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WalRcvRunning
func F_WalRcvRunning(m *base.Module) int32
//go:linkname F_WalSndWakeup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WalSndWakeup
func F_WalSndWakeup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ProcessRepliesIfAny github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessRepliesIfAny
func F_ProcessRepliesIfAny(m *base.Module)
//go:linkname F_WalSndWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WalSndWait
func F_WalSndWait(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ProcessPendingWrites github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ProcessPendingWrites
func F_ProcessPendingWrites(m *base.Module)
//go:linkname F_AcquireRewriteLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AcquireRewriteLocks
func F_AcquireRewriteLocks(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_view_query github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_view_query
func F_get_view_query(m *base.Module, l0 int32) int32
//go:linkname F_QueryRewrite github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_QueryRewrite
func F_QueryRewrite(m *base.Module, l0 int32) int32
//go:linkname F_ChangeVarNodesExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ChangeVarNodesExtended
func F_ChangeVarNodesExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ChangeVarNodes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ChangeVarNodes
func F_ChangeVarNodes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_adjust_relid_set github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_adjust_relid_set
func F_adjust_relid_set(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IncrementVarSublevelsUp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_IncrementVarSublevelsUp
func F_IncrementVarSublevelsUp(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_remove_nulling_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_remove_nulling_relids
func F_remove_nulling_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_replace_rte_variables github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replace_rte_variables
func F_replace_rte_variables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ReplaceVarFromTargetList github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplaceVarFromTargetList
func F_ReplaceVarFromTargetList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_statext_dependencies_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_statext_dependencies_deserialize
func F_statext_dependencies_deserialize(m *base.Module, l0 int32) int32
//go:linkname F_dependency_is_compatible_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dependency_is_compatible_clause
func F_dependency_is_compatible_clause(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dependency_is_compatible_expression github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dependency_is_compatible_expression
func F_dependency_is_compatible_expression(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_choose_best_statistics github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_choose_best_statistics
func F_choose_best_statistics(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_statext_is_compatible_clause_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_statext_is_compatible_clause_internal
func F_statext_is_compatible_clause_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_statext_mcv_load github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_statext_mcv_load
func F_statext_mcv_load(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_relation_statistics_update github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_relation_statistics_update
func F_relation_statistics_update(m *base.Module, l0 int32) int32
//go:linkname F_stats_fill_fcinfo_from_arg_pairs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_stats_fill_fcinfo_from_arg_pairs
func F_stats_fill_fcinfo_from_arg_pairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgaio_io_acquire_nb github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgaio_io_acquire_nb
func F_pgaio_io_acquire_nb(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgaio_io_reclaim github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgaio_io_reclaim
func F_pgaio_io_reclaim(m *base.Module, l0 int32)
//go:linkname F_pgaio_submit_staged github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgaio_submit_staged
func F_pgaio_submit_staged(m *base.Module)
//go:linkname F_pgaio_io_wait github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgaio_io_wait
func F_pgaio_io_wait(m *base.Module, l0 int32, l1 int64)
//go:linkname F_pgaio_wref_clear github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgaio_wref_clear
func F_pgaio_wref_clear(m *base.Module, l0 int32)
//go:linkname F_pgaio_wref_wait github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgaio_wref_wait
func F_pgaio_wref_wait(m *base.Module, l0 int32)
//go:linkname F_pgaio_error_cleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgaio_error_cleanup
func F_pgaio_error_cleanup(m *base.Module)
//go:linkname F_AtEOXact_Aio github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtEOXact_Aio
func F_AtEOXact_Aio(m *base.Module)
//go:linkname F_pgaio_io_register_callbacks github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgaio_io_register_callbacks
func F_pgaio_io_register_callbacks(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pgaio_result_report github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgaio_result_report
func F_pgaio_result_report(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pgaio_io_perform_synchronously github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgaio_io_perform_synchronously
func F_pgaio_io_perform_synchronously(m *base.Module, l0 int32)
//go:linkname F_read_stream_begin_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_read_stream_begin_relation
func F_read_stream_begin_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_read_stream_next_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_read_stream_next_buffer
func F_read_stream_next_buffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_read_stream_start_pending_read github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_read_stream_start_pending_read
func F_read_stream_start_pending_read(m *base.Module, l0 int32) int32
//go:linkname F_read_stream_reset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_read_stream_reset
func F_read_stream_reset(m *base.Module, l0 int32)
//go:linkname F_BufTableHashCode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BufTableHashCode
func F_BufTableHashCode(m *base.Module, l0 int32) int32
//go:linkname F_BufTableInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufTableInsert
func F_BufTableInsert(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_BufTableDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BufTableDelete
func F_BufTableDelete(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnpinBufferNoOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnpinBufferNoOwner
func F_UnpinBufferNoOwner(m *base.Module, l0 int32)
//go:linkname F_PrefetchBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PrefetchBuffer
func F_PrefetchBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ReservePrivateRefCountEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReservePrivateRefCountEntry
func F_ReservePrivateRefCountEntry(m *base.Module)
//go:linkname F_LockBufHdr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LockBufHdr
func F_LockBufHdr(m *base.Module, l0 int32) int32
//go:linkname F_GetPrivateRefCountEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetPrivateRefCountEntry
func F_GetPrivateRefCountEntry(m *base.Module, l0 int32) int32
//go:linkname F_ReadBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReadBuffer
func F_ReadBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExtendBufferedRel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExtendBufferedRel
func F_ExtendBufferedRel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GetVictimBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetVictimBuffer
func F_GetVictimBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_UnpinBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_UnpinBuffer
func F_UnpinBuffer(m *base.Module, l0 int32)
//go:linkname F_ZeroAndLockBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ZeroAndLockBuffer
func F_ZeroAndLockBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_StartReadBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_StartReadBuffer
func F_StartReadBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_WaitReadBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WaitReadBuffers
func F_WaitReadBuffers(m *base.Module, l0 int32)
//go:linkname F_ReadBufferWithoutRelcache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReadBufferWithoutRelcache
func F_ReadBufferWithoutRelcache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ExtendBufferedRelCommon github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExtendBufferedRelCommon
func F_ExtendBufferedRelCommon(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_StartBufferIO github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_StartBufferIO
func F_StartBufferIO(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReleaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReleaseBuffer
func F_ReleaseBuffer(m *base.Module, l0 int32)
//go:linkname F_MarkBufferDirty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MarkBufferDirty
func F_MarkBufferDirty(m *base.Module, l0 int32)
//go:linkname F_FlushBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FlushBuffer
func F_FlushBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_BufferGetLSNAtomic github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufferGetLSNAtomic
func F_BufferGetLSNAtomic(m *base.Module, l0 int32) int64
//go:linkname F_FindAndDropRelationBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FindAndDropRelationBuffers
func F_FindAndDropRelationBuffers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_DropDatabaseBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DropDatabaseBuffers
func F_DropDatabaseBuffers(m *base.Module, l0 int32)
//go:linkname F_UnlockReleaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UnlockReleaseBuffer
func F_UnlockReleaseBuffer(m *base.Module, l0 int32)
//go:linkname F_LockBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockBuffer
func F_LockBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_IncrBufferRefCount github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_IncrBufferRefCount
func F_IncrBufferRefCount(m *base.Module, l0 int32)
//go:linkname F_MarkBufferDirtyHint github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MarkBufferDirtyHint
func F_MarkBufferDirtyHint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ConditionalLockBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConditionalLockBuffer
func F_ConditionalLockBuffer(m *base.Module, l0 int32) int32
//go:linkname F_InvalidateVictimBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InvalidateVictimBuffer
func F_InvalidateVictimBuffer(m *base.Module, l0 int32) int32
//go:linkname F_StrategyFreeBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StrategyFreeBuffer
func F_StrategyFreeBuffer(m *base.Module, l0 int32)
//go:linkname F_GetAccessStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetAccessStrategy
func F_GetAccessStrategy(m *base.Module, l0 int32) int32
//go:linkname F_GetAccessStrategyWithSize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetAccessStrategyWithSize
func F_GetAccessStrategyWithSize(m *base.Module, l0 int32) int32
//go:linkname F_LocalBufferAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LocalBufferAlloc
func F_LocalBufferAlloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_PinLocalBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PinLocalBuffer
func F_PinLocalBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_InvalidateLocalBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InvalidateLocalBuffer
func F_InvalidateLocalBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_StartLocalBufferIO github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StartLocalBufferIO
func F_StartLocalBufferIO(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_UnpinLocalBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UnpinLocalBuffer
func F_UnpinLocalBuffer(m *base.Module, l0 int32)
//go:linkname F_BufFileCreateTemp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BufFileCreateTemp
func F_BufFileCreateTemp(m *base.Module, l0 int32) int32
//go:linkname F_BufFileOpenFileSet github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BufFileOpenFileSet
func F_BufFileOpenFileSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_BufFileDeleteFileSet github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BufFileDeleteFileSet
func F_BufFileDeleteFileSet(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_BufFileDumpBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BufFileDumpBuffer
func F_BufFileDumpBuffer(m *base.Module, l0 int32)
//go:linkname F_BufFileWrite github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufFileWrite
func F_BufFileWrite(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_copydir github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_copydir
func F_copydir(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_fsync_fname_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fsync_fname_ext
func F_fsync_fname_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CloseTransientFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CloseTransientFile
func F_CloseTransientFile(m *base.Module, l0 int32) int32
//go:linkname F_durable_rename github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_durable_rename
func F_durable_rename(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_OpenTransientFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_OpenTransientFile
func F_OpenTransientFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FreeDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeDesc
func F_FreeDesc(m *base.Module, l0 int32) int32
//go:linkname F_FileClose github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileClose
func F_FileClose(m *base.Module, l0 int32)
//go:linkname F_LruDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LruDelete
func F_LruDelete(m *base.Module, l0 int32)
//go:linkname F_AllocateDir github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AllocateDir
func F_AllocateDir(m *base.Module, l0 int32) int32
//go:linkname F_ReadDirExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReadDirExtended
func F_ReadDirExtended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FileAccess github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileAccess
func F_FileAccess(m *base.Module, l0 int32) int32
//go:linkname F_FileTruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileTruncate
func F_FileTruncate(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_reserveAllocatedDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_reserveAllocatedDesc
func F_reserveAllocatedDesc(m *base.Module) int32
//go:linkname F_OpenPipeStream github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_OpenPipeStream
func F_OpenPipeStream(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FreeFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FreeFile
func F_FreeFile(m *base.Module, l0 int32) int32
//go:linkname F_FileSetCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FileSetCreate
func F_FileSetCreate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FileSetDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FileSetDelete
func F_FileSetDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fsm_search github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fsm_search
func F_fsm_search(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RecordPageWithFreeSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RecordPageWithFreeSpace
func F_RecordPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogRecordPageWithFreeSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogRecordPageWithFreeSpace
func F_XLogRecordPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_FreeSpaceMapVacuum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeSpaceMapVacuum
func F_FreeSpaceMapVacuum(m *base.Module, l0 int32)
//go:linkname F_FreeSpaceMapVacuumRange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreeSpaceMapVacuumRange
func F_FreeSpaceMapVacuumRange(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RecordFreeIndexPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RecordFreeIndexPage
func F_RecordFreeIndexPage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BarrierArriveAndWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BarrierArriveAndWait
func F_BarrierArriveAndWait(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BarrierAttach github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BarrierAttach
func F_BarrierAttach(m *base.Module, l0 int32) int32
//go:linkname F_dsm_create github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dsm_create
func F_dsm_create(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsm_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsm_detach
func F_dsm_detach(m *base.Module, l0 int32)
//go:linkname F_dsm_pin_segment github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsm_pin_segment
func F_dsm_pin_segment(m *base.Module, l0 int32)
//go:linkname F_proc_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_proc_exit
func F_proc_exit(m *base.Module, l0 int32)
//go:linkname F_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shmem_exit
func F_shmem_exit(m *base.Module, l0 int32)
//go:linkname F_on_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_on_shmem_exit
func F_on_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_cancel_before_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cancel_before_shmem_exit
func F_cancel_before_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_WaitLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_WaitLatch
func F_WaitLatch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SetLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SetLatch
func F_SetLatch(m *base.Module, l0 int32)
//go:linkname F_SendPostmasterSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SendPostmasterSignal
func F_SendPostmasterSignal(m *base.Module, l0 int32)
//go:linkname F_ProcArrayAdd github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ProcArrayAdd
func F_ProcArrayAdd(m *base.Module, l0 int32)
//go:linkname F_ProcArrayRemove github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ProcArrayRemove
func F_ProcArrayRemove(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ProcArrayEndTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcArrayEndTransaction
func F_ProcArrayEndTransaction(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ProcArrayApplyRecoveryInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcArrayApplyRecoveryInfo
func F_ProcArrayApplyRecoveryInfo(m *base.Module, l0 int32)
//go:linkname F_GetOldestNonRemovableTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetOldestNonRemovableTransactionId
func F_GetOldestNonRemovableTransactionId(m *base.Module, l0 int32) int32
//go:linkname F_ComputeXidHorizons github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ComputeXidHorizons
func F_ComputeXidHorizons(m *base.Module, l0 int32)
//go:linkname F_GlobalVisHorizonKindForRel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GlobalVisHorizonKindForRel
func F_GlobalVisHorizonKindForRel(m *base.Module, l0 int32) int32
//go:linkname F_GetRunningTransactionData github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetRunningTransactionData
func F_GetRunningTransactionData(m *base.Module) int32
//go:linkname F_BackendPidGetProc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BackendPidGetProc
func F_BackendPidGetProc(m *base.Module, l0 int32) int32
//go:linkname F_GetCurrentVirtualXIDs github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetCurrentVirtualXIDs
func F_GetCurrentVirtualXIDs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CountOtherDBBackends github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CountOtherDBBackends
func F_CountOtherDBBackends(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SendProcSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SendProcSignal
func F_SendProcSignal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EmitProcSignalBarrier github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EmitProcSignalBarrier
func F_EmitProcSignalBarrier(m *base.Module) int64
//go:linkname F_WaitForProcSignalBarrier github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WaitForProcSignalBarrier
func F_WaitForProcSignalBarrier(m *base.Module, l0 int64)
//go:linkname F_shm_mq_send github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shm_mq_send
func F_shm_mq_send(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_shm_mq_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_mq_detach
func F_shm_mq_detach(m *base.Module, l0 int32)
//go:linkname F_shm_toc_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_toc_insert
func F_shm_toc_insert(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_shm_toc_lookup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_shm_toc_lookup
func F_shm_toc_lookup(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_ShmemInitStruct github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ShmemInitStruct
func F_ShmemInitStruct(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_add_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_size
func F_add_size(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_mul_size github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_mul_size
func F_mul_size(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReceiveSharedInvalidMessages github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReceiveSharedInvalidMessages
func F_ReceiveSharedInvalidMessages(m *base.Module)
//go:linkname F_ProcessCatchupInterrupt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessCatchupInterrupt
func F_ProcessCatchupInterrupt(m *base.Module)
//go:linkname F_SICleanupQueue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SICleanupQueue
func F_SICleanupQueue(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ResolveRecoveryConflictWithVirtualXIDs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResolveRecoveryConflictWithVirtualXIDs
func F_ResolveRecoveryConflictWithVirtualXIDs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_WaitEventSetWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WaitEventSetWait
func F_WaitEventSetWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_inv_create github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inv_create
func F_inv_create(m *base.Module, l0 int32) int32
//go:linkname F_inv_open github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inv_open
func F_inv_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ConditionVariablePrepareToSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ConditionVariablePrepareToSleep
func F_ConditionVariablePrepareToSleep(m *base.Module, l0 int32)
//go:linkname F_ConditionVariableCancelSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ConditionVariableCancelSleep
func F_ConditionVariableCancelSleep(m *base.Module)
//go:linkname F_ConditionVariableSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ConditionVariableSleep
func F_ConditionVariableSleep(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ConditionVariableTimedSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ConditionVariableTimedSleep
func F_ConditionVariableTimedSleep(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ConditionVariableBroadcast github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ConditionVariableBroadcast
func F_ConditionVariableBroadcast(m *base.Module, l0 int32)
//go:linkname F_FindLockCycleRecurseMember github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FindLockCycleRecurseMember
func F_FindLockCycleRecurseMember(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_LockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockRelationOid
func F_LockRelationOid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ConditionalLockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConditionalLockRelationOid
func F_ConditionalLockRelationOid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_UnlockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockRelationOid
func F_UnlockRelationOid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnlockRelationIdForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockRelationIdForSession
func F_UnlockRelationIdForSession(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnlockRelationForExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockRelationForExtension
func F_UnlockRelationForExtension(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LockTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockTuple
func F_LockTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ConditionalLockTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConditionalLockTuple
func F_ConditionalLockTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_UnlockTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockTuple
func F_UnlockTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XactLockTableWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XactLockTableWait
func F_XactLockTableWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ConditionalXactLockTableWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ConditionalXactLockTableWait
func F_ConditionalXactLockTableWait(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_WaitForLockersMultiple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WaitForLockersMultiple
func F_WaitForLockersMultiple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_UnlockDatabaseObject github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockDatabaseObject
func F_UnlockDatabaseObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockSharedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockSharedObject
func F_LockSharedObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_UnlockSharedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_UnlockSharedObject
func F_UnlockSharedObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockSharedObjectForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockSharedObjectForSession
func F_LockSharedObjectForSession(m *base.Module, l0 int32)
//go:linkname F_UnlockSharedObjectForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_UnlockSharedObjectForSession
func F_UnlockSharedObjectForSession(m *base.Module, l0 int32)
//go:linkname F_LockApplyTransactionForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockApplyTransactionForSession
func F_LockApplyTransactionForSession(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_UnlockApplyTransactionForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_UnlockApplyTransactionForSession
func F_UnlockApplyTransactionForSession(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RemoveLocalLock github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RemoveLocalLock
func F_RemoveLocalLock(m *base.Module, l0 int32)
//go:linkname F_LockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockAcquire
func F_LockAcquire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_LockAcquireExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockAcquireExtended
func F_LockAcquireExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_SetupLockInTable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SetupLockInTable
func F_SetupLockInTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_LockRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LockRelease
func F_LockRelease(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LockRefindAndRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockRefindAndRelease
func F_LockRefindAndRelease(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_VirtualXactLock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_VirtualXactLock
func F_VirtualXactLock(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LWLockNewTrancheId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LWLockNewTrancheId
func F_LWLockNewTrancheId(m *base.Module) int32
//go:linkname F_LWLockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockAcquire
func F_LWLockAcquire(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LWLockQueueSelf github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LWLockQueueSelf
func F_LWLockQueueSelf(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LWLockDequeueSelf github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LWLockDequeueSelf
func F_LWLockDequeueSelf(m *base.Module, l0 int32)
//go:linkname F_LWLockConditionalAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockConditionalAcquire
func F_LWLockConditionalAcquire(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LWLockRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LWLockRelease
func F_LWLockRelease(m *base.Module, l0 int32)
//go:linkname F_LWLockReleaseInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockReleaseInternal
func F_LWLockReleaseInternal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LWLockReleaseAll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockReleaseAll
func F_LWLockReleaseAll(m *base.Module)
//go:linkname F_PredicateLockPage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PredicateLockPage
func F_PredicateLockPage(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PredicateLockTID github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PredicateLockTID
func F_PredicateLockTID(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_PredicateLockPageSplit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PredicateLockPageSplit
func F_PredicateLockPageSplit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CheckForSerializableConflictIn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckForSerializableConflictIn
func F_CheckForSerializableConflictIn(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PreCommit_CheckForSerializationFailure github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreCommit_CheckForSerializationFailure
func F_PreCommit_CheckForSerializationFailure(m *base.Module)
//go:linkname F_ProcLockWakeup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcLockWakeup
func F_ProcLockWakeup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_s_lock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_s_lock
func F_s_lock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_perform_spin_delay github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_perform_spin_delay
func F_perform_spin_delay(m *base.Module, l0 int32)
//go:linkname F_PageAddItemExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PageAddItemExtended
func F_PageAddItemExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_PageRepairFragmentation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PageRepairFragmentation
func F_PageRepairFragmentation(m *base.Module, l0 int32)
//go:linkname F_PageIndexTupleDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageIndexTupleDelete
func F_PageIndexTupleDelete(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PageIndexTupleOverwrite github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageIndexTupleOverwrite
func F_PageIndexTupleOverwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_smgr_bulk_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_smgr_bulk_finish
func F_smgr_bulk_finish(m *base.Module, l0 int32)
//go:linkname F_smgr_bulk_write github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_smgr_bulk_write
func F_smgr_bulk_write(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_mdopenfork github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mdopenfork
func F_mdopenfork(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_register_dirty_segment github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_register_dirty_segment
func F_register_dirty_segment(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_mdextend github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mdextend
func F_mdextend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F__mdfd_getseg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__mdfd_getseg
func F__mdfd_getseg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_mdimmedsync github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mdimmedsync
func F_mdimmedsync(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ForgetDatabaseSyncRequests github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ForgetDatabaseSyncRequests
func F_ForgetDatabaseSyncRequests(m *base.Module, l0 int32)
//go:linkname F_smgropen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgropen
func F_smgropen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgrdestroyall github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_smgrdestroyall
func F_smgrdestroyall(m *base.Module)
//go:linkname F_smgrcreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgrcreate
func F_smgrcreate(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_smgrwritev github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_smgrwritev
func F_smgrwritev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_smgrnblocks github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_smgrnblocks
func F_smgrnblocks(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RegisterSyncRequest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RegisterSyncRequest
func F_RegisterSyncRequest(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CreateDestReceiver github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CreateDestReceiver
func F_CreateDestReceiver(m *base.Module, l0 int32) int32
//go:linkname F_EndCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EndCommand
func F_EndCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ProcessInterrupts github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ProcessInterrupts
func F_ProcessInterrupts(m *base.Module)
//go:linkname F_ShowUsage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ShowUsage
func F_ShowUsage(m *base.Module, l0 int32)
//go:linkname F_pg_analyze_and_rewrite_fixedparams github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_analyze_and_rewrite_fixedparams
func F_pg_analyze_and_rewrite_fixedparams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_plan_query github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_plan_query
func F_pg_plan_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CreateQueryDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CreateQueryDesc
func F_CreateQueryDesc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_FreeQueryDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeQueryDesc
func F_FreeQueryDesc(m *base.Module, l0 int32)
//go:linkname F_ChoosePortalStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ChoosePortalStrategy
func F_ChoosePortalStrategy(m *base.Module, l0 int32) int32
//go:linkname F_FillPortalStore github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FillPortalStore
func F_FillPortalStore(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PortalRunSelect github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PortalRunSelect
func F_PortalRunSelect(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64
//go:linkname F_PortalRunMulti github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PortalRunMulti
func F_PortalRunMulti(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_PreventCommandIfReadOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreventCommandIfReadOnly
func F_PreventCommandIfReadOnly(m *base.Module, l0 int32)
//go:linkname F_PreventCommandIfParallelMode github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreventCommandIfParallelMode
func F_PreventCommandIfParallelMode(m *base.Module, l0 int32)
//go:linkname F_PreventCommandDuringRecovery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PreventCommandDuringRecovery
func F_PreventCommandDuringRecovery(m *base.Module, l0 int32)
//go:linkname F_UtilityTupleDescriptor github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UtilityTupleDescriptor
func F_UtilityTupleDescriptor(m *base.Module, l0 int32) int32
//go:linkname F_newLexeme github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_newLexeme
func F_newLexeme(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_addWrd github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_addWrd
func F_addWrd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_setCompoundAffixFlagValue github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_setCompoundAffixFlagValue
func F_setCompoundAffixFlagValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_make_tsvector github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_tsvector
func F_make_tsvector(m *base.Module, l0 int32) int32
//go:linkname F_tsearch_readline_begin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tsearch_readline_begin
func F_tsearch_readline_begin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tsearch_readline github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tsearch_readline
func F_tsearch_readline(m *base.Module, l0 int32) int32
//go:linkname F_tsearch_readline_end github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tsearch_readline_end
func F_tsearch_readline_end(m *base.Module, l0 int32)
//go:linkname F_parsetext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parsetext
func F_parsetext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_hlparsetext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hlparsetext
func F_hlparsetext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_generateHeadline github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_generateHeadline
func F_generateHeadline(m *base.Module, l0 int32) int32
//go:linkname F_get_tsearch_config_filename github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_tsearch_config_filename
func F_get_tsearch_config_filename(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_readstoplist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_readstoplist
func F_readstoplist(m *base.Module, l0 int32, l1 int32)
//go:linkname F_TParserGet github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TParserGet
func F_TParserGet(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_progress_parallel_incr_param github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_progress_parallel_incr_param
func F_pgstat_progress_parallel_incr_param(m *base.Module, l0 int32, l1 int64)
//go:linkname F_pgstat_report_activity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_report_activity
func F_pgstat_report_activity(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_clip_activity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_clip_activity
func F_pgstat_clip_activity(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_get_beentry_by_proc_number github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_get_beentry_by_proc_number
func F_pgstat_get_beentry_by_proc_number(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_read_current_status github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_read_current_status
func F_pgstat_read_current_status(m *base.Module)
//go:linkname F_pgstat_get_local_beentry_by_index github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_get_local_beentry_by_index
func F_pgstat_get_local_beentry_by_index(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_clear_snapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_clear_snapshot
func F_pgstat_clear_snapshot(m *base.Module)
//go:linkname F_pgstat_prep_snapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_prep_snapshot
func F_pgstat_prep_snapshot(m *base.Module)
//go:linkname F_pgstat_snapshot_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_snapshot_insert
func F_pgstat_snapshot_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgstat_prep_pending_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_prep_pending_entry
func F_pgstat_prep_pending_entry(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pgstat_fetch_pending_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_fetch_pending_entry
func F_pgstat_fetch_pending_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pgstat_count_backend_io_op github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_count_backend_io_op
func F_pgstat_count_backend_io_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64)
//go:linkname F_pgstat_fetch_stat_checkpointer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_fetch_stat_checkpointer
func F_pgstat_fetch_stat_checkpointer(m *base.Module) int32
//go:linkname F_pgstat_prepare_report_checksum_failure github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_prepare_report_checksum_failure
func F_pgstat_prepare_report_checksum_failure(m *base.Module, l0 int32)
//go:linkname F_pgstat_report_checksum_failures_in_db github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_report_checksum_failures_in_db
func F_pgstat_report_checksum_failures_in_db(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_fetch_stat_dbentry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_fetch_stat_dbentry
func F_pgstat_fetch_stat_dbentry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_fetch_stat_funcentry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_fetch_stat_funcentry
func F_pgstat_fetch_stat_funcentry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_flush_io github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_flush_io
func F_pgstat_flush_io(m *base.Module, l0 int32)
//go:linkname F_pgstat_assoc_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_assoc_relation
func F_pgstat_assoc_relation(m *base.Module, l0 int32)
//go:linkname F_pgstat_count_heap_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_count_heap_insert
func F_pgstat_count_heap_insert(m *base.Module, l0 int32, l1 int64)
//go:linkname F_pgstat_count_truncate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_count_truncate
func F_pgstat_count_truncate(m *base.Module, l0 int32)
//go:linkname F_find_tabstat_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_tabstat_entry
func F_find_tabstat_entry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_get_entry_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_get_entry_ref
func F_pgstat_get_entry_ref(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32
//go:linkname F_pgstat_lock_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_lock_entry
func F_pgstat_lock_entry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_get_entry_ref_locked github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_get_entry_ref_locked
func F_pgstat_get_entry_ref_locked(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pgstat_drop_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_drop_entry
func F_pgstat_drop_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pgstat_reset_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_reset_entry
func F_pgstat_reset_entry(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64)
//go:linkname F_AtEOXact_PgStat github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtEOXact_PgStat
func F_AtEOXact_PgStat(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_get_xact_stack_level github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_get_xact_stack_level
func F_pgstat_get_xact_stack_level(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_get_transactional_drops github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_get_transactional_drops
func F_pgstat_get_transactional_drops(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_create_transactional github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_create_transactional
func F_pgstat_create_transactional(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_pgstat_get_wait_event github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_get_wait_event
func F_pgstat_get_wait_event(m *base.Module, l0 int32) int32
//go:linkname F_check_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_acl
func F_check_acl(m *base.Module, l0 int32)
//go:linkname F_aclmask github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_aclmask
func F_aclmask(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int64
//go:linkname F_aclitemsort github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_aclitemsort
func F_aclitemsort(m *base.Module, l0 int32)
//go:linkname F_acldefault github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_acldefault
func F_acldefault(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_roles_is_member_of github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_roles_is_member_of
func F_roles_is_member_of(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_has_privs_of_role github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_has_privs_of_role
func F_has_privs_of_role(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_convert_any_priv_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_any_priv_string
func F_convert_any_priv_string(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_convert_column_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_column_name
func F_convert_column_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_role_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_role_oid
func F_get_role_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DatumGetAnyArrayP github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DatumGetAnyArrayP
func F_DatumGetAnyArrayP(m *base.Module, l0 int32) int32
//go:linkname F_mcelem_array_selec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mcelem_array_selec
func F_mcelem_array_selec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) float64
//go:linkname F_array_shuffle_n github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_shuffle_n
func F_array_shuffle_n(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_array_sort_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_sort_internal
func F_array_sort_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_construct_empty_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_empty_array
func F_construct_empty_array(m *base.Module, l0 int32) int32
//go:linkname F_array_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_recv
func F_array_recv(m *base.Module, l0 int32) int32
//go:linkname F_array_seek github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_seek
func F_array_seek(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_array_slice_size github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_array_slice_size
func F_array_slice_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_array_set_element github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_array_set_element
func F_array_set_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_construct_md_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_md_array
func F_construct_md_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_deconstruct_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_array
func F_deconstruct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_array_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_ref
func F_array_ref(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_array_set github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_set
func F_array_set(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_construct_array_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_array_builtin
func F_construct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_deconstruct_array_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_array_builtin
func F_deconstruct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_array_contains_nulls github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_contains_nulls
func F_array_contains_nulls(m *base.Module, l0 int32) int32
//go:linkname F_array_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_cmp
func F_array_cmp(m *base.Module, l0 int32) int32
//go:linkname F_array_create_iterator github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_create_iterator
func F_array_create_iterator(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_array_iterate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_iterate
func F_array_iterate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_array_free_iterator github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_free_iterator
func F_array_free_iterator(m *base.Module, l0 int32)
//go:linkname F_initArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_initArrayResult
func F_initArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_initArrayResultWithSize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_initArrayResultWithSize
func F_initArrayResultWithSize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_makeArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeArrayResult
func F_makeArrayResult(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_initArrayResultAny github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_initArrayResultAny
func F_initArrayResultAny(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_accumArrayResultAny github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_accumArrayResultAny
func F_accumArrayResultAny(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeArrayResultAny github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeArrayResultAny
func F_makeArrayResultAny(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_array_fill_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_fill_internal
func F_array_fill_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ArrayGetNItemsSafe github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ArrayGetNItemsSafe
func F_ArrayGetNItemsSafe(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ArrayCheckBounds github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ArrayCheckBounds
func F_ArrayCheckBounds(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ArrayGetIntegerTypmods github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ArrayGetIntegerTypmods
func F_ArrayGetIntegerTypmods(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_encode_to_ascii github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_encode_to_ascii
func F_encode_to_ascii(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cash_send github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_cash_send
func F_cash_send(m *base.Module, l0 int32) int32
//go:linkname F_date2j github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_date2j
func F_date2j(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_j2date github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_j2date
func F_j2date(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_j2day github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_j2day
func F_j2day(m *base.Module, l0 int32) int32
//go:linkname F_GetCurrentTimeUsec github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetCurrentTimeUsec
func F_GetCurrentTimeUsec(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ParseDateTime github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ParseDateTime
func F_ParseDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_DecodeNumberField github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecodeNumberField
func F_DecodeNumberField(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_DecodeDate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DecodeDate
func F_DecodeDate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_DecodeTimeCommon github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DecodeTimeCommon
func F_DecodeTimeCommon(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_DecodeNumber github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DecodeNumber
func F_DecodeNumber(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_DecodeTimezoneAbbrev github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DecodeTimezoneAbbrev
func F_DecodeTimezoneAbbrev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_DetermineTimeZoneOffsetInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DetermineTimeZoneOffsetInternal
func F_DetermineTimeZoneOffsetInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DetermineTimeZoneOffset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DetermineTimeZoneOffset
func F_DetermineTimeZoneOffset(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DecodeTimeOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecodeTimeOnly
func F_DecodeTimeOnly(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_DecodeTimezoneName github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DecodeTimezoneName
func F_DecodeTimezoneName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DateTimeParseError github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DateTimeParseError
func F_DateTimeParseError(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_datumGetSize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_datumGetSize
func F_datumGetSize(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datumCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_datumCopy
func F_datumCopy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datumTransfer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_datumTransfer
func F_datumTransfer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datum_image_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_datum_image_eq
func F_datum_image_eq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_datumSerialize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_datumSerialize
func F_datumSerialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_calculate_database_size github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_calculate_database_size
func F_calculate_database_size(m *base.Module, l0 int32) int64
//go:linkname F_domain_check github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_domain_check
func F_domain_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_check_safe_enum_use github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_safe_enum_use
func F_check_safe_enum_use(m *base.Module, l0 int32)
//go:linkname F_enum_cmp_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_enum_cmp_internal
func F_enum_cmp_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EOH_get_flat_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EOH_get_flat_size
func F_EOH_get_flat_size(m *base.Module, l0 int32) int32
//go:linkname F_EOH_flatten_into github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EOH_flatten_into
func F_EOH_flatten_into(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_DeleteExpandedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DeleteExpandedObject
func F_DeleteExpandedObject(m *base.Module, l0 int32)
//go:linkname F_make_expanded_record_from_typeid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_expanded_record_from_typeid
func F_make_expanded_record_from_typeid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_expanded_record_fetch_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expanded_record_fetch_tupdesc
func F_expanded_record_fetch_tupdesc(m *base.Module, l0 int32) int32
//go:linkname F_build_dummy_expanded_header github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_dummy_expanded_header
func F_build_dummy_expanded_header(m *base.Module, l0 int32)
//go:linkname F_deconstruct_expanded_record github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_deconstruct_expanded_record
func F_deconstruct_expanded_record(m *base.Module, l0 int32)
//go:linkname F_float_overflow_error github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float_overflow_error
func F_float_overflow_error(m *base.Module)
//go:linkname F_float_underflow_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_float_underflow_error
func F_float_underflow_error(m *base.Module)
//go:linkname F_float_zero_divide_error github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float_zero_divide_error
func F_float_zero_divide_error(m *base.Module)
//go:linkname F_float8in_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_float8in_internal
func F_float8in_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_float8out_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float8out_internal
func F_float8out_internal(m *base.Module, l0 float64) int32
//go:linkname F_format_type_be github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_format_type_be
func F_format_type_be(m *base.Module, l0 int32) int32
//go:linkname F_format_type_be_qualified github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_format_type_be_qualified
func F_format_type_be_qualified(m *base.Module, l0 int32) int32
//go:linkname F_format_type_with_typemod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_format_type_with_typemod
func F_format_type_with_typemod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_str_tolower github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_str_tolower
func F_str_tolower(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_NUM_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_NUM_cache
func F_NUM_cache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_NUM_processor github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_NUM_processor
func F_NUM_processor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_pg_read_file_common github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_read_file_common
func F_pg_read_file_common(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32
//go:linkname F_convert_and_check_filename github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_and_check_filename
func F_convert_and_check_filename(m *base.Module, l0 int32) int32
//go:linkname F_box_cn github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_box_cn
func F_box_cn(m *base.Module, l0 int32, l1 int32)
//go:linkname F_point_dt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_point_dt
func F_point_dt(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_point_sl github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_point_sl
func F_point_sl(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_line_construct github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_line_construct
func F_line_construct(m *base.Module, l0 int32, l1 int32, l2 float64)
//go:linkname F_lseg_interpt_lseg github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lseg_interpt_lseg
func F_lseg_interpt_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_box_closept_point github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_box_closept_point
func F_box_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_dist_ppoly_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dist_ppoly_internal
func F_dist_ppoly_internal(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_point_inside github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_point_inside
func F_point_inside(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_touched_lseg_inside_poly github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_touched_lseg_inside_poly
func F_touched_lseg_inside_poly(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_buildint2vector github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_buildint2vector
func F_buildint2vector(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_escape_json github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_escape_json
func F_escape_json(m *base.Module, l0 int32, l1 int32)
//go:linkname F_escape_json_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_escape_json_with_len
func F_escape_json_with_len(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_json_object_agg_transfn_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_json_object_agg_transfn_worker
func F_json_object_agg_transfn_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_json_build_object_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_json_build_object_worker
func F_json_build_object_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_JsonbToCStringWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_JsonbToCStringWorker
func F_JsonbToCStringWorker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_jsonb_object_agg_transfn_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_jsonb_object_agg_transfn_worker
func F_jsonb_object_agg_transfn_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_extract_jsp_bool_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_extract_jsp_bool_expr
func F_extract_jsp_bool_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_execute_jsp_gin_node github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_execute_jsp_gin_node
func F_execute_jsp_gin_node(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonbValueToJsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_JsonbValueToJsonb
func F_JsonbValueToJsonb(m *base.Module, l0 int32) int32
//go:linkname F_pushJsonbValue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pushJsonbValue
func F_pushJsonbValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_iteratorFromContainer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_iteratorFromContainer
func F_iteratorFromContainer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_compareJsonbContainers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_compareJsonbContainers
func F_compareJsonbContainers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonbIteratorInit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_JsonbIteratorInit
func F_JsonbIteratorInit(m *base.Module, l0 int32) int32
//go:linkname F_getIthJsonbValueFromContainer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getIthJsonbValueFromContainer
func F_getIthJsonbValueFromContainer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonbHashScalarValue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_JsonbHashScalarValue
func F_JsonbHashScalarValue(m *base.Module, l0 int32, l1 int32)
//go:linkname F_json_errsave_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_json_errsave_error
func F_json_errsave_error(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_JsonbValueAsText github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_JsonbValueAsText
func F_JsonbValueAsText(m *base.Module, l0 int32) int32
//go:linkname F_jsonb_get_element github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_jsonb_get_element
func F_jsonb_get_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_populate_record_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_populate_record_worker
func F_populate_record_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_populate_recordset_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_populate_recordset_worker
func F_populate_recordset_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_iterate_json_values github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_iterate_json_values
func F_iterate_json_values(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_jspInitByBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_jspInitByBuffer
func F_jspInitByBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_jspGetArg github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_jspGetArg
func F_jspGetArg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_jspGetNext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_jspGetNext
func F_jspGetNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_executeJsonPath github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_executeJsonPath
func F_executeJsonPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_jsonb_path_query_first_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_jsonb_path_query_first_internal
func F_jsonb_path_query_first_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_executeItemOptUnwrapTarget github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_executeItemOptUnwrapTarget
func F_executeItemOptUnwrapTarget(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_executeAnyItem github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_executeAnyItem
func F_executeAnyItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_makeItemLikeRegex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeItemLikeRegex
func F_makeItemLikeRegex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeItemUnary github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeItemUnary
func F_makeItemUnary(m *base.Module, l0 int32) int32
//go:linkname F_jsonpath_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_jsonpath_yyensure_buffer_stack
func F_jsonpath_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_jsonpath_yy_create_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_jsonpath_yy_create_buffer
func F_jsonpath_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_yy_fatal_error_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_yy_fatal_error_5
func F_yy_fatal_error_5(m *base.Module, l0 int32)
//go:linkname F_jsonpath_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_jsonpath_yyerror
func F_jsonpath_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_parseUnicode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_parseUnicode
func F_parseUnicode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_addUnicodeChar github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_addUnicodeChar
func F_addUnicodeChar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SB_MatchText github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SB_MatchText
func F_SB_MatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_UTF8_MatchText github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_UTF8_MatchText
func F_UTF8_MatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_MB_MatchText github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MB_MatchText
func F_MB_MatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pattern_fixed_prefix github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pattern_fixed_prefix
func F_pattern_fixed_prefix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_make_greater_string github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_greater_string
func F_make_greater_string(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_count_nulls github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_count_nulls
func F_count_nulls(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_multirange_out github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_multirange_out
func F_multirange_out(m *base.Module, l0 int32) int32
//go:linkname F_multirange_get_range github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_multirange_get_range
func F_multirange_get_range(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_multirange_get_typcache github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_multirange_get_typcache
func F_multirange_get_typcache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_multirange_get_bounds github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_multirange_get_bounds
func F_multirange_get_bounds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_multirange_intersect_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_multirange_intersect_internal
func F_multirange_intersect_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_range_contains_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_contains_multirange_internal
func F_range_contains_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overleft_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_overleft_multirange_internal
func F_range_overleft_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overright_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_overright_multirange_internal
func F_range_overright_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_multirange_contains_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_multirange_contains_multirange_internal
func F_multirange_contains_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_before_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_before_multirange_internal
func F_range_before_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_after_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_range_after_multirange_internal
func F_range_after_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_adjacent_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_adjacent_multirange_internal
func F_range_adjacent_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_multirange_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_multirange_cmp
func F_multirange_cmp(m *base.Module, l0 int32) int32
//go:linkname F_get_position github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_position
func F_get_position(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64
//go:linkname F_network_out github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_network_out
func F_network_out(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cidr_set_masklen_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cidr_set_masklen_internal
func F_cidr_set_masklen_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_match_network_subset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_match_network_subset
func F_match_network_subset(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_convert_network_to_scalar github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_convert_network_to_scalar
func F_convert_network_to_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_inet_spg_consistent_bitmap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_inet_spg_consistent_bitmap
func F_inet_spg_consistent_bitmap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_result_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_result_opt_error
func F_make_result_opt_error(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_mul_var github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_mul_var
func F_mul_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_add_var github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_var
func F_add_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_str_from_var github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_str_from_var
func F_get_str_from_var(m *base.Module, l0 int32) int32
//go:linkname F_div_var github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_div_var
func F_div_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_sub_var github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sub_var
func F_sub_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_div_var_int github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_div_var_int
func F_div_var_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_numeric_add_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_numeric_add_opt_error
func F_numeric_add_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_numeric_div_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_numeric_div_opt_error
func F_numeric_div_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_int64_to_numeric github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_int64_to_numeric
func F_int64_to_numeric(m *base.Module, l0 int64) int32
//go:linkname F_int64_div_fast_to_numeric github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_int64_div_fast_to_numeric
func F_int64_div_fast_to_numeric(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_numeric_int4_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_numeric_int4_opt_error
func F_numeric_int4_opt_error(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_do_numeric_accum github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_do_numeric_accum
func F_do_numeric_accum(m *base.Module, l0 int32, l1 int32)
//go:linkname F_accum_sum_add github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_accum_sum_add
func F_accum_sum_add(m *base.Module, l0 int32, l1 int32)
//go:linkname F_accum_sum_final github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_accum_sum_final
func F_accum_sum_final(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_strtoint32 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_strtoint32
func F_pg_strtoint32(m *base.Module, l0 int32) int32
//go:linkname F_uint32in_subr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_uint32in_subr
func F_uint32in_subr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_ultoa_n github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_ultoa_n
func F_pg_ultoa_n(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_ultostr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_ultostr
func F_pg_ultostr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_buildoidvector github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_buildoidvector
func F_buildoidvector(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_valid_oidvector github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_valid_oidvector
func F_check_valid_oidvector(m *base.Module, l0 int32)
//go:linkname F_dobyteatrim github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_dobyteatrim
func F_dobyteatrim(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_percentile_cont_multi_final_common github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_percentile_cont_multi_final_common
func F_percentile_cont_multi_final_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_check_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_locale
func F_check_locale(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_PGLC_localeconv github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PGLC_localeconv
func F_PGLC_localeconv(m *base.Module) int32
//go:linkname F_pg_newlocale_from_collation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_newlocale_from_collation
func F_pg_newlocale_from_collation(m *base.Module, l0 int32) int32
//go:linkname F_get_collation_actual_version github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_collation_actual_version
func F_get_collation_actual_version(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_strupper github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_strupper
func F_pg_strupper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_strnxfrm github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_strnxfrm
func F_pg_strnxfrm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_builtin_locale_encoding github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_builtin_locale_encoding
func F_builtin_locale_encoding(m *base.Module, l0 int32) int32
//go:linkname F_char2wchar github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_char2wchar
func F_char2wchar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_report_newlocale_failure github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_newlocale_failure
func F_report_newlocale_failure(m *base.Module, l0 int32)
//go:linkname F_make_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_range
func F_make_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_range_serialize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_serialize
func F_range_serialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_range_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_deserialize
func F_range_deserialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_range_get_typcache github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_get_typcache
func F_range_get_typcache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_range_contains_elem_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_range_contains_elem_internal
func F_range_contains_elem_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_eq_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_eq_internal
func F_range_eq_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_cmp_bounds github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_cmp_bounds
func F_range_cmp_bounds(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_before_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_range_before_internal
func F_range_before_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_after_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_after_internal
func F_range_after_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_adjacent_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_adjacent_internal
func F_range_adjacent_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overlaps_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_overlaps_internal
func F_range_overlaps_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overright_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_overright_internal
func F_range_overright_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_empty_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_empty_range
func F_make_empty_range(m *base.Module, l0 int32) int32
//go:linkname F_range_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_cmp
func F_range_cmp(m *base.Module, l0 int32) int32
//go:linkname F_range_gist_consistent_leaf_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_gist_consistent_leaf_multirange
func F_range_gist_consistent_leaf_multirange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_range_gist_consistent_int_range github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_gist_consistent_int_range
func F_range_gist_consistent_int_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_range_super_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_super_union
func F_range_super_union(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RE_compile_and_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RE_compile_and_cache
func F_RE_compile_and_cache(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_parse_re_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parse_re_flags
func F_parse_re_flags(m *base.Module, l0 int32, l1 int32)
//go:linkname F_regexp_instr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_regexp_instr
func F_regexp_instr(m *base.Module, l0 int32) int32
//go:linkname F_regexp_matches github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_regexp_matches
func F_regexp_matches(m *base.Module, l0 int32) int32
//go:linkname F_regexp_split_to_table github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_regexp_split_to_table
func F_regexp_split_to_table(m *base.Module, l0 int32) int32
//go:linkname F_parseNameAndArgTypes github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parseNameAndArgTypes
func F_parseNameAndArgTypes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_format_operator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_format_operator
func F_format_operator(m *base.Module, l0 int32) int32
//go:linkname F_ri_CheckTrigger github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ri_CheckTrigger
func F_ri_CheckTrigger(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RI_FKey_check github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RI_FKey_check
func F_RI_FKey_check(m *base.Module, l0 int32)
//go:linkname F_ri_restrict github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ri_restrict
func F_ri_restrict(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ri_set github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ri_set
func F_ri_set(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_record_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_record_eq
func F_record_eq(m *base.Module, l0 int32) int32
//go:linkname F_record_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_record_cmp
func F_record_cmp(m *base.Module, l0 int32) int32
//go:linkname F_quote_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_quote_identifier
func F_quote_identifier(m *base.Module, l0 int32) int32
//go:linkname F_generate_relation_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generate_relation_name
func F_generate_relation_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_deparse_for_query github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_deparse_for_query
func F_set_deparse_for_query(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_query_def github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_query_def
func F_get_query_def(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_set_rtable_names github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_rtable_names
func F_set_rtable_names(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_set_simple_column_names github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_simple_column_names
func F_set_simple_column_names(m *base.Module, l0 int32)
//go:linkname F_generate_function_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_function_name
func F_generate_function_name(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_get_opclass_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_opclass_name
func F_get_opclass_name(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_reloptions
func F_get_reloptions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_generate_operator_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_operator_name
func F_generate_operator_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_flatten_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_flatten_reloptions
func F_flatten_reloptions(m *base.Module, l0 int32) int32
//go:linkname F_get_rule_sortgroupclause github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rule_sortgroupclause
func F_get_rule_sortgroupclause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_appendContextKeyword github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_appendContextKeyword
func F_appendContextKeyword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_processIndirection github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_processIndirection
func F_processIndirection(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_variable
func F_get_variable(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_rule_list_toplevel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rule_list_toplevel
func F_get_rule_list_toplevel(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_deparse_expression github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_deparse_expression
func F_deparse_expression(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_get_constraintdef_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_get_constraintdef_worker
func F_pg_get_constraintdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_generate_operator_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generate_operator_clause
func F_generate_operator_clause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_range_partbound_string github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_range_partbound_string
func F_get_range_partbound_string(m *base.Module, l0 int32) int32
//go:linkname F_get_const_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_const_expr
func F_get_const_expr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_agg_expr_helper github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_agg_expr_helper
func F_get_agg_expr_helper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_windowfunc_expr_helper github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_windowfunc_expr_helper
func F_get_windowfunc_expr_helper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_get_rule_expr_paren github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_rule_expr_paren
func F_get_rule_expr_paren(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_isSimpleNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_isSimpleNode
func F_isSimpleNode(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_json_returning github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_json_returning
func F_get_json_returning(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_json_expr_options github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_json_expr_options
func F_get_json_expr_options(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_json_agg_constructor github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_json_agg_constructor
func F_get_json_agg_constructor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_xmltable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_xmltable
func F_get_xmltable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_json_table github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_json_table
func F_get_json_table(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_from_clause_item github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_from_clause_item
func F_get_from_clause_item(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_restriction_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_restriction_variable
func F_get_restriction_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_statistic_proc_security_check github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_statistic_proc_security_check
func F_statistic_proc_security_check(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_examine_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_examine_variable
func F_examine_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_scalarineqsel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_scalarineqsel
func F_scalarineqsel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64
//go:linkname F_all_rows_selectable github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_all_rows_selectable
func F_all_rows_selectable(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_nulltestsel github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_nulltestsel
func F_nulltestsel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64
//go:linkname F_get_variable_range github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_variable_range
func F_get_variable_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_estimate_num_groups github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_estimate_num_groups
func F_estimate_num_groups(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32) float64
//go:linkname F_genericcostestimate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_genericcostestimate
func F_genericcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32)
//go:linkname F_currtid_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_currtid_internal
func F_currtid_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_timestamp2tm github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_timestamp2tm
func F_timestamp2tm(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_make_timestamp_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_timestamp_internal
func F_make_timestamp_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) int64
//go:linkname F_timestamp2timestamptz_opt_overflow github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamp2timestamptz_opt_overflow
func F_timestamp2timestamptz_opt_overflow(m *base.Module, l0 int64, l1 int32) int64
//go:linkname F_GetCurrentTimestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetCurrentTimestamp
func F_GetCurrentTimestamp(m *base.Module) int64
//go:linkname F_timestamp_cmp_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamp_cmp_internal
func F_timestamp_cmp_internal(m *base.Module, l0 int64, l1 int64) int32
//go:linkname F_interval_um_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_interval_um_internal
func F_interval_um_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_timestamptz_trunc_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamptz_trunc_internal
func F_timestamptz_trunc_internal(m *base.Module, l0 int32, l1 int64, l2 int32) int64
//go:linkname F_interval_part_common github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_interval_part_common
func F_interval_part_common(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gin_extract_tsquery github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gin_extract_tsquery
func F_gin_extract_tsquery(m *base.Module, l0 int32) int32
//go:linkname F_infix_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_infix_1
func F_infix_1(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_QT2QTN github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_QT2QTN
func F_QT2QTN(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_QTNFree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_QTNFree
func F_QTNFree(m *base.Module, l0 int32)
//go:linkname F_QTNodeCompare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_QTNodeCompare
func F_QTNodeCompare(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_QTN2QT github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_QTN2QT
func F_QTN2QT(m *base.Module, l0 int32) int32
//go:linkname F_getWeights github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getWeights
func F_getWeights(m *base.Module, l0 int32, l1 int32)
//go:linkname F_calc_rank_cd github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_calc_rank_cd
func F_calc_rank_cd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32
//go:linkname F_TS_execute_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TS_execute_recurse
func F_TS_execute_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ts_setup_firstcall github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ts_setup_firstcall
func F_ts_setup_firstcall(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ts_process_call github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ts_process_call
func F_ts_process_call(m *base.Module, l0 int32) int32
//go:linkname F_chooseNextStatEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_chooseNextStatEntry
func F_chooseNextStatEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_init_tsvector_parser github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_init_tsvector_parser
func F_init_tsvector_parser(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_close_tsvector_parser github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_close_tsvector_parser
func F_close_tsvector_parser(m *base.Module, l0 int32)
//go:linkname F_gettoken_tsvector github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gettoken_tsvector
func F_gettoken_tsvector(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_anychar_typmodin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_anychar_typmodin
func F_anychar_typmodin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_varchar_input github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_varchar_input
func F_varchar_input(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_cstring_to_text github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cstring_to_text
func F_cstring_to_text(m *base.Module, l0 int32) int32
//go:linkname F_cstring_to_text_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cstring_to_text_with_len
func F_cstring_to_text_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_text_to_cstring github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_text_to_cstring
func F_text_to_cstring(m *base.Module, l0 int32) int32
//go:linkname F_text_catenate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text_catenate
func F_text_catenate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_text_overlay github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_text_overlay
func F_text_overlay(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_text_position_setup github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text_position_setup
func F_text_position_setup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_text_position_next github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text_position_next
func F_text_position_next(m *base.Module, l0 int32) int32
//go:linkname F_varstr_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_varstr_cmp
func F_varstr_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_bytea_overlay github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bytea_overlay
func F_bytea_overlay(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_textToQualifiedNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_textToQualifiedNameList
func F_textToQualifiedNameList(m *base.Module, l0 int32) int32
//go:linkname F_SplitIdentifierString github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SplitIdentifierString
func F_SplitIdentifierString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_map_sql_table_to_xmlschema github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_map_sql_table_to_xmlschema
func F_map_sql_table_to_xmlschema(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CatCacheRemoveCTup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CatCacheRemoveCTup
func F_CatCacheRemoveCTup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateCacheMemoryContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateCacheMemoryContext
func F_CreateCacheMemoryContext(m *base.Module)
//go:linkname F_ResetCatalogCache github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResetCatalogCache
func F_ResetCatalogCache(m *base.Module, l0 int32)
//go:linkname F_CatalogCacheInitializeCache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CatalogCacheInitializeCache
func F_CatalogCacheInitializeCache(m *base.Module, l0 int32)
//go:linkname F_SearchCatCache github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SearchCatCache
func F_SearchCatCache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SearchCatCacheInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SearchCatCacheInternal
func F_SearchCatCacheInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_SearchCatCache2 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SearchCatCache2
func F_SearchCatCache2(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReleaseCatCache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReleaseCatCache
func F_ReleaseCatCache(m *base.Module, l0 int32)
//go:linkname F_fastgetattr_4 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fastgetattr_4
func F_fastgetattr_4(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_cached_function_compile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cached_function_compile
func F_cached_function_compile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_AtEOXact_Inval github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AtEOXact_Inval
func F_AtEOXact_Inval(m *base.Module, l0 int32)
//go:linkname F_xactGetCommittedInvalidationMessages github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_xactGetCommittedInvalidationMessages
func F_xactGetCommittedInvalidationMessages(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ProcessCommittedInvalidationMessages github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ProcessCommittedInvalidationMessages
func F_ProcessCommittedInvalidationMessages(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_LogLogicalInvalidations github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LogLogicalInvalidations
func F_LogLogicalInvalidations(m *base.Module)
//go:linkname F_CacheInvalidateHeapTupleCommon github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CacheInvalidateHeapTupleCommon
func F_CacheInvalidateHeapTupleCommon(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CacheInvalidateRelcache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CacheInvalidateRelcache
func F_CacheInvalidateRelcache(m *base.Module, l0 int32)
//go:linkname F_CacheInvalidateRelcacheByRelid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CacheInvalidateRelcacheByRelid
func F_CacheInvalidateRelcacheByRelid(m *base.Module, l0 int32)
//go:linkname F_CacheInvalidateRelSync github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CacheInvalidateRelSync
func F_CacheInvalidateRelSync(m *base.Module, l0 int32)
//go:linkname F_CacheInvalidateSmgr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CacheInvalidateSmgr
func F_CacheInvalidateSmgr(m *base.Module, l0 int32)
//go:linkname F_CacheRegisterSyscacheCallback github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CacheRegisterSyscacheCallback
func F_CacheRegisterSyscacheCallback(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CacheRegisterRelcacheCallback github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CacheRegisterRelcacheCallback
func F_CacheRegisterRelcacheCallback(m *base.Module, l0 int32)
//go:linkname F_op_in_opfamily github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_op_in_opfamily
func F_op_in_opfamily(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_op_opfamily_properties github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_op_opfamily_properties
func F_get_op_opfamily_properties(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_opfamily_member_for_cmptype github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_opfamily_member_for_cmptype
func F_get_opfamily_member_for_cmptype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_opfamily_method github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_opfamily_method
func F_get_opfamily_method(m *base.Module, l0 int32) int32
//go:linkname F_get_equality_op_for_ordering_op github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_equality_op_for_ordering_op
func F_get_equality_op_for_ordering_op(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_opfamily_proc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opfamily_proc
func F_get_opfamily_proc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_attname github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_attname
func F_get_attname(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_attnum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_attnum
func F_get_attnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_attoptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_attoptions
func F_get_attoptions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_collation_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_collation_name
func F_get_collation_name(m *base.Module, l0 int32) int32
//go:linkname F_get_collation_isdeterministic github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_collation_isdeterministic
func F_get_collation_isdeterministic(m *base.Module, l0 int32) int32
//go:linkname F_get_constraint_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_constraint_name
func F_get_constraint_name(m *base.Module, l0 int32) int32
//go:linkname F_get_opclass_input_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_opclass_input_type
func F_get_opclass_input_type(m *base.Module, l0 int32) int32
//go:linkname F_get_opclass_opfamily_and_input_type github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_opclass_opfamily_and_input_type
func F_get_opclass_opfamily_and_input_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_opfamily_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_opfamily_name
func F_get_opfamily_name(m *base.Module, l0 int32) int32
//go:linkname F_get_opcode github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opcode
func F_get_opcode(m *base.Module, l0 int32) int32
//go:linkname F_get_op_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_op_rettype
func F_get_op_rettype(m *base.Module, l0 int32) int32
//go:linkname F_func_strict github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_func_strict
func F_func_strict(m *base.Module, l0 int32) int32
//go:linkname F_func_volatile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_func_volatile
func F_func_volatile(m *base.Module, l0 int32) int32
//go:linkname F_get_commutator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_commutator
func F_get_commutator(m *base.Module, l0 int32) int32
//go:linkname F_get_oprrest github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_oprrest
func F_get_oprrest(m *base.Module, l0 int32) int32
//go:linkname F_get_func_name github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_name
func F_get_func_name(m *base.Module, l0 int32) int32
//go:linkname F_get_func_variadictype github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_func_variadictype
func F_get_func_variadictype(m *base.Module, l0 int32) int32
//go:linkname F_func_parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_func_parallel
func F_func_parallel(m *base.Module, l0 int32) int32
//go:linkname F_get_func_prokind github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_prokind
func F_get_func_prokind(m *base.Module, l0 int32) int32
//go:linkname F_get_func_leakproof github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_leakproof
func F_get_func_leakproof(m *base.Module, l0 int32) int32
//go:linkname F_get_relname_relid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_relname_relid
func F_get_relname_relid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rel_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rel_name
func F_get_rel_name(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_namespace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_rel_namespace
func F_get_rel_namespace(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_relkind github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_rel_relkind
func F_get_rel_relkind(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_relispartition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_rel_relispartition
func F_get_rel_relispartition(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_persistence github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_rel_persistence
func F_get_rel_persistence(m *base.Module, l0 int32) int32
//go:linkname F_get_typlen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_typlen
func F_get_typlen(m *base.Module, l0 int32) int32
//go:linkname F_get_typbyval github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typbyval
func F_get_typbyval(m *base.Module, l0 int32) int32
//go:linkname F_get_typlenbyvalalign github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typlenbyvalalign
func F_get_typlenbyvalalign(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_getBaseType github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getBaseType
func F_getBaseType(m *base.Module, l0 int32) int32
//go:linkname F_getBaseTypeAndTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getBaseTypeAndTypmod
func F_getBaseTypeAndTypmod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_typavgwidth github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_typavgwidth
func F_get_typavgwidth(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_typtype github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typtype
func F_get_typtype(m *base.Module, l0 int32) int32
//go:linkname F_type_is_rowtype github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_type_is_rowtype
func F_type_is_rowtype(m *base.Module, l0 int32) int32
//go:linkname F_type_is_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_type_is_multirange
func F_type_is_multirange(m *base.Module, l0 int32) int32
//go:linkname F_get_type_category_preferred github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_type_category_preferred
func F_get_type_category_preferred(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_element_type github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_element_type
func F_get_element_type(m *base.Module, l0 int32) int32
//go:linkname F_get_base_element_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_base_element_type
func F_get_base_element_type(m *base.Module, l0 int32) int32
//go:linkname F_getTypeInputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getTypeInputInfo
func F_getTypeInputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_getTypeOutputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getTypeOutputInfo
func F_getTypeOutputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_getTypeBinaryInputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getTypeBinaryInputInfo
func F_getTypeBinaryInputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_getTypeBinaryOutputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getTypeBinaryOutputInfo
func F_getTypeBinaryOutputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_typcollation github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_typcollation
func F_get_typcollation(m *base.Module, l0 int32) int32
//go:linkname F_getSubscriptingRoutines github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_getSubscriptingRoutines
func F_getSubscriptingRoutines(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_namespace_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_namespace_name
func F_get_namespace_name(m *base.Module, l0 int32) int32
//go:linkname F_get_namespace_name_or_temp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_namespace_name_or_temp
func F_get_namespace_name_or_temp(m *base.Module, l0 int32) int32
//go:linkname F_get_index_column_opclass github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_index_column_opclass
func F_get_index_column_opclass(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_index_isvalid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_index_isvalid
func F_get_index_isvalid(m *base.Module, l0 int32) int32
//go:linkname F_get_publication_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_publication_oid
func F_get_publication_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_subscription_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_subscription_oid
func F_get_subscription_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generate_partition_qual github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_partition_qual
func F_generate_partition_qual(m *base.Module, l0 int32) int32
//go:linkname F_AcquireExecutorLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AcquireExecutorLocks
func F_AcquireExecutorLocks(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BuildCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BuildCachedPlan
func F_BuildCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ScanQueryForLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ScanQueryForLocks
func F_ScanQueryForLocks(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReleaseCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReleaseCachedPlan
func F_ReleaseCachedPlan(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationInitTableAccessMethod github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationInitTableAccessMethod
func F_RelationInitTableAccessMethod(m *base.Module, l0 int32)
//go:linkname F_RelationIdGetRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationIdGetRelation
func F_RelationIdGetRelation(m *base.Module, l0 int32) int32
//go:linkname F_RelationBuildDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationBuildDesc
func F_RelationBuildDesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationIncrementReferenceCount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationIncrementReferenceCount
func F_RelationIncrementReferenceCount(m *base.Module, l0 int32)
//go:linkname F_RelationInitPhysicalAddr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationInitPhysicalAddr
func F_RelationInitPhysicalAddr(m *base.Module, l0 int32)
//go:linkname F_ScanPgRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ScanPgRelation
func F_ScanPgRelation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RelationParseRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationParseRelOptions
func F_RelationParseRelOptions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationDestroyRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationDestroyRelation
func F_RelationDestroyRelation(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationClose github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RelationClose
func F_RelationClose(m *base.Module, l0 int32)
//go:linkname F_RelationClearRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationClearRelation
func F_RelationClearRelation(m *base.Module, l0 int32)
//go:linkname F_RelationCacheInvalidate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationCacheInvalidate
func F_RelationCacheInvalidate(m *base.Module)
//go:linkname F_AtEOXact_RelationCache github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AtEOXact_RelationCache
func F_AtEOXact_RelationCache(m *base.Module, l0 int32)
//go:linkname F_RelationGetFKeyList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationGetFKeyList
func F_RelationGetFKeyList(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexList github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationGetIndexList
func F_RelationGetIndexList(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexExpressions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetIndexExpressions
func F_RelationGetIndexExpressions(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexPredicate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetIndexPredicate
func F_RelationGetIndexPredicate(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexAttrBitmap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RelationGetIndexAttrBitmap
func F_RelationGetIndexAttrBitmap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationMapInvalidate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationMapInvalidate
func F_RelationMapInvalidate(m *base.Module, l0 int32)
//go:linkname F_AtEOXact_RelationMap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AtEOXact_RelationMap
func F_AtEOXact_RelationMap(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_tablespace_page_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_tablespace_page_costs
func F_get_tablespace_page_costs(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_tablespace_maintenance_io_concurrency github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_tablespace_maintenance_io_concurrency
func F_get_tablespace_maintenance_io_concurrency(m *base.Module, l0 int32) int32
//go:linkname F_SearchSysCache1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchSysCache1
func F_SearchSysCache1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCache3 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SearchSysCache3
func F_SearchSysCache3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SearchSysCacheLocked1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchSysCacheLocked1
func F_SearchSysCacheLocked1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchSysCacheCopy
func F_SearchSysCacheCopy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetSysCacheOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetSysCacheOid
func F_GetSysCacheOid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SearchSysCacheAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SearchSysCacheAttName
func F_SearchSysCacheAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheCopyAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SearchSysCacheCopyAttName
func F_SearchSysCacheCopyAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SysCacheGetAttr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SysCacheGetAttr
func F_SysCacheGetAttr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SysCacheGetAttrNotNull github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SysCacheGetAttrNotNull
func F_SysCacheGetAttrNotNull(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetSysCacheHashValue github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetSysCacheHashValue
func F_GetSysCacheHashValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchSysCacheList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SearchSysCacheList
func F_SearchSysCacheList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_lookup_ts_parser_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lookup_ts_parser_cache
func F_lookup_ts_parser_cache(m *base.Module, l0 int32) int32
//go:linkname F_lookup_ts_dictionary_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lookup_ts_dictionary_cache
func F_lookup_ts_dictionary_cache(m *base.Module, l0 int32) int32
//go:linkname F_lookup_ts_config_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lookup_ts_config_cache
func F_lookup_ts_config_cache(m *base.Module, l0 int32) int32
//go:linkname F_getTSCurrentConfig github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_getTSCurrentConfig
func F_getTSCurrentConfig(m *base.Module) int32
//go:linkname F_lookup_type_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lookup_type_cache
func F_lookup_type_cache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookup_rowtype_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lookup_rowtype_tupdesc
func F_lookup_rowtype_tupdesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookup_rowtype_tupdesc_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lookup_rowtype_tupdesc_domain
func F_lookup_rowtype_tupdesc_domain(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AtEOXact_TypeCache github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AtEOXact_TypeCache
func F_AtEOXact_TypeCache(m *base.Module)
//go:linkname F_errstart_cold github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errstart_cold
func F_errstart_cold(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errstart github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_errstart
func F_errstart(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_errmsg_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errmsg_internal
func F_errmsg_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errfinish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errfinish
func F_errfinish(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_EmitErrorReport github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_EmitErrorReport
func F_EmitErrorReport(m *base.Module)
//go:linkname F_pg_re_throw github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_re_throw
func F_pg_re_throw(m *base.Module)
//go:linkname F_errsave_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errsave_finish
func F_errsave_finish(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errcode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errcode
func F_errcode(m *base.Module, l0 int32)
//go:linkname F_errcode_for_file_access github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errcode_for_file_access
func F_errcode_for_file_access(m *base.Module)
//go:linkname F_errcode_for_socket_access github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errcode_for_socket_access
func F_errcode_for_socket_access(m *base.Module)
//go:linkname F_errmsg github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errmsg
func F_errmsg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errdetail github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errdetail
func F_errdetail(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errdetail_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errdetail_internal
func F_errdetail_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errdetail_plural github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errdetail_plural
func F_errdetail_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errhint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errhint
func F_errhint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errcontext_msg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errcontext_msg
func F_errcontext_msg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_set_errcontext_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_errcontext_domain
func F_set_errcontext_domain(m *base.Module, l0 int32)
//go:linkname F_errhidecontext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errhidecontext
func F_errhidecontext(m *base.Module)
//go:linkname F_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errposition
func F_errposition(m *base.Module, l0 int32) int32
//go:linkname F_internalerrposition github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_internalerrposition
func F_internalerrposition(m *base.Module, l0 int32)
//go:linkname F_internalerrquery github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_internalerrquery
func F_internalerrquery(m *base.Module, l0 int32) int32
//go:linkname F_geterrcode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_geterrcode
func F_geterrcode(m *base.Module) int32
//go:linkname F_geterrposition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_geterrposition
func F_geterrposition(m *base.Module) int32
//go:linkname F_CopyErrorData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopyErrorData
func F_CopyErrorData(m *base.Module) int32
//go:linkname F_FlushErrorState github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FlushErrorState
func F_FlushErrorState(m *base.Module)
//go:linkname F_ReThrowError github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReThrowError
func F_ReThrowError(m *base.Module, l0 int32)
//go:linkname F_load_file github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_load_file
func F_load_file(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fmgr_info github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fmgr_info
func F_fmgr_info(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fmgr_info_cxt_security github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fmgr_info_cxt_security
func F_fmgr_info_cxt_security(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_fmgr_info_cxt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fmgr_info_cxt
func F_fmgr_info_cxt(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_detoast_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_detoast_datum
func F_pg_detoast_datum(m *base.Module, l0 int32) int32
//go:linkname F_DirectFunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DirectFunctionCall1Coll
func F_DirectFunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DirectFunctionCall2Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DirectFunctionCall2Coll
func F_DirectFunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_DirectFunctionCall3Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DirectFunctionCall3Coll
func F_DirectFunctionCall3Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_DirectFunctionCall4Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DirectFunctionCall4Coll
func F_DirectFunctionCall4Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_DirectFunctionCall5Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DirectFunctionCall5Coll
func F_DirectFunctionCall5Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_CallerFInfoFunctionCall2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CallerFInfoFunctionCall2
func F_CallerFInfoFunctionCall2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_FunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FunctionCall1Coll
func F_FunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FunctionCall2Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FunctionCall2Coll
func F_FunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_FunctionCall3Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FunctionCall3Coll
func F_FunctionCall3Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_FunctionCall4Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FunctionCall4Coll
func F_FunctionCall4Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_FunctionCall5Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FunctionCall5Coll
func F_FunctionCall5Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_FunctionCall6Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FunctionCall6Coll
func F_FunctionCall6Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_FunctionCall7Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FunctionCall7Coll
func F_FunctionCall7Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_InputFunctionCallSafe github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InputFunctionCallSafe
func F_InputFunctionCallSafe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_DirectInputFunctionCallSafe github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DirectInputFunctionCallSafe
func F_DirectInputFunctionCallSafe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_OutputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OutputFunctionCall
func F_OutputFunctionCall(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReceiveFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReceiveFunctionCall
func F_ReceiveFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SendFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SendFunctionCall
func F_SendFunctionCall(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_OidOutputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OidOutputFunctionCall
func F_OidOutputFunctionCall(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_Int64GetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Int64GetDatum
func F_Int64GetDatum(m *base.Module, l0 int64) int32
//go:linkname F_Float8GetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_Float8GetDatum
func F_Float8GetDatum(m *base.Module, l0 float64) int32
//go:linkname F_pg_detoast_datum_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_detoast_datum_copy
func F_pg_detoast_datum_copy(m *base.Module, l0 int32) int32
//go:linkname F_get_fn_expr_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_fn_expr_rettype
func F_get_fn_expr_rettype(m *base.Module, l0 int32) int32
//go:linkname F_get_fn_expr_argtype github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_fn_expr_argtype
func F_get_fn_expr_argtype(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_fn_opclass_options github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_fn_opclass_options
func F_set_fn_opclass_options(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_fn_opclass_options github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_fn_opclass_options
func F_get_fn_opclass_options(m *base.Module, l0 int32) int32
//go:linkname F_InitMaterializedSRF github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InitMaterializedSRF
func F_InitMaterializedSRF(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_call_result_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_call_result_type
func F_get_call_result_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_init_MultiFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_init_MultiFuncCall
func F_init_MultiFuncCall(m *base.Module, l0 int32) int32
//go:linkname F_end_MultiFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_end_MultiFuncCall
func F_end_MultiFuncCall(m *base.Module, l0 int32)
//go:linkname F_get_expr_result_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_expr_result_type
func F_get_expr_result_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_function_result_tupdesc_d github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_function_result_tupdesc_d
func F_build_function_result_tupdesc_d(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_extract_variadic_args github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_extract_variadic_args
func F_extract_variadic_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_create github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_create
func F_hash_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_my_log2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_my_log2
func F_my_log2(m *base.Module, l0 int32) int32
//go:linkname F_hash_destroy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hash_destroy
func F_hash_destroy(m *base.Module, l0 int32)
//go:linkname F_get_hash_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_hash_value
func F_get_hash_value(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hash_search_with_hash_value github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_search_with_hash_value
func F_hash_search_with_hash_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_hash_seq_init github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_seq_init
func F_hash_seq_init(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetUserNameFromId github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetUserNameFromId
func F_GetUserNameFromId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_InitPostgres github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InitPostgres
func F_InitPostgres(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_local2local github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_local2local
func F_local2local(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_LocalToUtf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LocalToUtf
func F_LocalToUtf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_SetClientEncoding github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SetClientEncoding
func F_SetClientEncoding(m *base.Module, l0 int32) int32
//go:linkname F_pg_do_encoding_conversion github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_do_encoding_conversion
func F_pg_do_encoding_conversion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_report_invalid_encoding github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_invalid_encoding
func F_report_invalid_encoding(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_do_encoding_conversion_buf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_do_encoding_conversion_buf
func F_pg_do_encoding_conversion_buf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_pg_any_to_server github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_any_to_server
func F_pg_any_to_server(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_server_to_client github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_server_to_client
func F_pg_server_to_client(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_server_to_any github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_server_to_any
func F_pg_server_to_any(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_mblen_cstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_mblen_cstr
func F_pg_mblen_cstr(m *base.Module, l0 int32) int32
//go:linkname F_report_invalid_encoding_int github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_report_invalid_encoding_int
func F_report_invalid_encoding_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pg_mblen_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mblen_range
func F_pg_mblen_range(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mblen_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_mblen_with_len
func F_pg_mblen_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mbstrlen_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mbstrlen_with_len
func F_pg_mbstrlen_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mbcliplen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mbcliplen
func F_pg_mbcliplen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_verifymbstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_verifymbstr
func F_pg_verifymbstr(m *base.Module, l0 int32, l1 int32)
//go:linkname F_check_encoding_conversion_args github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_encoding_conversion_args
func F_check_encoding_conversion_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_report_untranslatable_char github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_untranslatable_char
func F_report_untranslatable_char(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_AbsoluteConfigLocation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AbsoluteConfigLocation
func F_AbsoluteConfigLocation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetConfFilesInDir github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetConfFilesInDir
func F_GetConfFilesInDir(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_find_option github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_option
func F_find_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_set_config_with_handle github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_set_config_with_handle
func F_set_config_with_handle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_InitializeGUCOptionsFromEnvironment github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_InitializeGUCOptionsFromEnvironment
func F_InitializeGUCOptionsFromEnvironment(m *base.Module)
//go:linkname F_guc_strdup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_guc_strdup
func F_guc_strdup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_config_option github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_config_option
func F_set_config_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_SetConfigOption github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SetConfigOption
func F_SetConfigOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_guc_malloc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_guc_malloc
func F_guc_malloc(m *base.Module, l0 int32) int32
//go:linkname F_build_guc_variables github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_build_guc_variables
func F_build_guc_variables(m *base.Module)
//go:linkname F_convert_GUC_name_for_parameter_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_convert_GUC_name_for_parameter_acl
func F_convert_GUC_name_for_parameter_acl(m *base.Module, l0 int32) int32
//go:linkname F_InitializeOneGUCOption github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InitializeOneGUCOption
func F_InitializeOneGUCOption(m *base.Module, l0 int32)
//go:linkname F_RestrictSearchPath github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RestrictSearchPath
func F_RestrictSearchPath(m *base.Module)
//go:linkname F_AtEOXact_GUC github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AtEOXact_GUC
func F_AtEOXact_GUC(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReportGUCOption github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReportGUCOption
func F_ReportGUCOption(m *base.Module, l0 int32)
//go:linkname F_ShowGUCOption github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ShowGUCOption
func F_ShowGUCOption(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DefineCustomBoolVariable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DefineCustomBoolVariable
func F_DefineCustomBoolVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_define_custom_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_define_custom_variable
func F_define_custom_variable(m *base.Module, l0 int32)
//go:linkname F_DefineCustomIntVariable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DefineCustomIntVariable
func F_DefineCustomIntVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_MarkGUCPrefixReserved github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MarkGUCPrefixReserved
func F_MarkGUCPrefixReserved(m *base.Module, l0 int32)
//go:linkname F_validate_option_array_item github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_validate_option_array_item
func F_validate_option_array_item(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GUC_yylex github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GUC_yylex
func F_GUC_yylex(m *base.Module, l0 int32) int32
//go:linkname F_GUC_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GUC_yyensure_buffer_stack
func F_GUC_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_GUC_yy_create_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GUC_yy_create_buffer
func F_GUC_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ProcessConfigFile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ProcessConfigFile
func F_ProcessConfigFile(m *base.Module, l0 int32)
//go:linkname F_ParseConfigFile github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ParseConfigFile
func F_ParseConfigFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_DeescapeQuotedString github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DeescapeQuotedString
func F_DeescapeQuotedString(m *base.Module, l0 int32) int32
//go:linkname F_flatten_set_variable_args github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_flatten_set_variable_args
func F_flatten_set_variable_args(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_visible_ENR_metadata github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_visible_ENR_metadata
func F_get_visible_ENR_metadata(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_stack_depth github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_stack_depth
func F_check_stack_depth(m *base.Module)
//go:linkname F_superuser github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_superuser
func F_superuser(m *base.Module) int32
//go:linkname F_schedule_alarm github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_schedule_alarm
func F_schedule_alarm(m *base.Module, l0 int64)
//go:linkname F_disable_timeout github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_disable_timeout
func F_disable_timeout(m *base.Module, l0 int32)
//go:linkname F_AllocSetContextCreateInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AllocSetContextCreateInternal
func F_AllocSetContextCreateInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_AllocSetAllocLarge github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AllocSetAllocLarge
func F_AllocSetAllocLarge(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_AllocSetAllocFromNewBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AllocSetAllocFromNewBlock
func F_AllocSetAllocFromNewBlock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_AllocSetFree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AllocSetFree
func F_AllocSetFree(m *base.Module, l0 int32)
//go:linkname F_BumpContextCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BumpContextCreate
func F_BumpContextCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dsa_create_in_place_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dsa_create_in_place_ext
func F_dsa_create_in_place_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_dsa_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_dsa_attach
func F_dsa_attach(m *base.Module, l0 int32) int32
//go:linkname F_dsa_attach_in_place github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_dsa_attach_in_place
func F_dsa_attach_in_place(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsa_pin_mapping github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dsa_pin_mapping
func F_dsa_pin_mapping(m *base.Module, l0 int32)
//go:linkname F_dsa_allocate_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsa_allocate_extended
func F_dsa_allocate_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_best_segment github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_best_segment
func F_get_best_segment(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsa_free github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dsa_free
func F_dsa_free(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_segment_by_index github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_segment_by_index
func F_get_segment_by_index(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_transfer_first_span github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_transfer_first_span
func F_transfer_first_span(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_dsa_get_address github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsa_get_address
func F_dsa_get_address(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsa_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_dsa_detach
func F_dsa_detach(m *base.Module, l0 int32)
//go:linkname F_FreePageManagerGet github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreePageManagerGet
func F_FreePageManagerGet(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FreePageManagerPutInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreePageManagerPutInternal
func F_FreePageManagerPutInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_FreePageManagerPut github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreePageManagerPut
func F_FreePageManagerPut(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_FreePageBtreeRemovePage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreePageBtreeRemovePage
func F_FreePageBtreeRemovePage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_FreePageBtreeConsolidate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreePageBtreeConsolidate
func F_FreePageBtreeConsolidate(m *base.Module, l0 int32, l1 int32)
//go:linkname F_MemoryContextDeleteChildren github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MemoryContextDeleteChildren
func F_MemoryContextDeleteChildren(m *base.Module, l0 int32)
//go:linkname F_MemoryContextResetOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoryContextResetOnly
func F_MemoryContextResetOnly(m *base.Module, l0 int32)
//go:linkname F_MemoryContextSetParent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextSetParent
func F_MemoryContextSetParent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetMemoryChunkSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetMemoryChunkSpace
func F_GetMemoryChunkSpace(m *base.Module, l0 int32) int32
//go:linkname F_MemoryContextStatsInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MemoryContextStatsInternal
func F_MemoryContextStatsInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_MemoryContextAllocationFailure github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextAllocationFailure
func F_MemoryContextAllocationFailure(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MemoryContextAllocZero github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoryContextAllocZero
func F_MemoryContextAllocZero(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextAllocExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextAllocExtended
func F_MemoryContextAllocExtended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_palloc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_palloc
func F_palloc(m *base.Module, l0 int32) int32
//go:linkname F_palloc0 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_palloc0
func F_palloc0(m *base.Module, l0 int32) int32
//go:linkname F_MemoryContextAllocAligned github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoryContextAllocAligned
func F_MemoryContextAllocAligned(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pfree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pfree
func F_pfree(m *base.Module, l0 int32)
//go:linkname F_repalloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_repalloc
func F_repalloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextStrdup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MemoryContextStrdup
func F_MemoryContextStrdup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pstrdup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pstrdup
func F_pstrdup(m *base.Module, l0 int32) int32
//go:linkname F_pnstrdup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pnstrdup
func F_pnstrdup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_PortalDrop github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PortalDrop
func F_PortalDrop(m *base.Module, l0 int32, l1 int32)
//go:linkname F_MarkPortalActive github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MarkPortalActive
func F_MarkPortalActive(m *base.Module, l0 int32)
//go:linkname F_MarkPortalFailed github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MarkPortalFailed
func F_MarkPortalFailed(m *base.Module, l0 int32)
//go:linkname F_PreCommit_Portals github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreCommit_Portals
func F_PreCommit_Portals(m *base.Module, l0 int32) int32
//go:linkname F_AtSubAbort_Portals github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AtSubAbort_Portals
func F_AtSubAbort_Portals(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_HoldPinnedPortals github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_HoldPinnedPortals
func F_HoldPinnedPortals(m *base.Module)
//go:linkname F_ForgetPortalSnapshots github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ForgetPortalSnapshots
func F_ForgetPortalSnapshots(m *base.Module)
//go:linkname F_ResourceOwnerCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResourceOwnerCreate
func F_ResourceOwnerCreate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ResourceOwnerEnlarge github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResourceOwnerEnlarge
func F_ResourceOwnerEnlarge(m *base.Module, l0 int32)
//go:linkname F_ResourceOwnerRemember github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResourceOwnerRemember
func F_ResourceOwnerRemember(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ResourceOwnerForget github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResourceOwnerForget
func F_ResourceOwnerForget(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ResourceOwnerReleaseInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResourceOwnerReleaseInternal
func F_ResourceOwnerReleaseInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ResourceOwnerDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ResourceOwnerDelete
func F_ResourceOwnerDelete(m *base.Module, l0 int32)
//go:linkname F_ResourceOwnerForgetLock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResourceOwnerForgetLock
func F_ResourceOwnerForgetLock(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LogicalTapeSetClose github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalTapeSetClose
func F_LogicalTapeSetClose(m *base.Module, l0 int32)
//go:linkname F_LogicalTapeCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalTapeCreate
func F_LogicalTapeCreate(m *base.Module, l0 int32) int32
//go:linkname F_LogicalTapeWrite github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LogicalTapeWrite
func F_LogicalTapeWrite(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ltsWriteBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltsWriteBlock
func F_ltsWriteBlock(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_LogicalTapeRead github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalTapeRead
func F_LogicalTapeRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_qsort_interruptible github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_qsort_interruptible
func F_qsort_interruptible(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_FinishSortSupportFunction github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FinishSortSupportFunction
func F_FinishSortSupportFunction(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PrepareSortSupportFromGistIndexRel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PrepareSortSupportFromGistIndexRel
func F_PrepareSortSupportFromGistIndexRel(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplesort_begin_common github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_begin_common
func F_tuplesort_begin_common(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplesort_end github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_end
func F_tuplesort_end(m *base.Module, l0 int32)
//go:linkname F_tuplesort_reset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_reset
func F_tuplesort_reset(m *base.Module, l0 int32)
//go:linkname F_tuplesort_puttuple_common github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_puttuple_common
func F_tuplesort_puttuple_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_tuplesort_performsort github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_performsort
func F_tuplesort_performsort(m *base.Module, l0 int32)
//go:linkname F_tuplesort_gettuple_common github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_gettuple_common
func F_tuplesort_gettuple_common(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplesort_estimate_shared github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_estimate_shared
func F_tuplesort_estimate_shared(m *base.Module, l0 int32) int32
//go:linkname F_tuplesort_initialize_shared github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_initialize_shared
func F_tuplesort_initialize_shared(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_attach_shared github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_attach_shared
func F_tuplesort_attach_shared(m *base.Module, l0 int32, l1 int32)
//go:linkname F_qsort_ssup_med3 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_qsort_ssup_med3
func F_qsort_ssup_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_tuplesort_begin_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_begin_datum
func F_tuplesort_begin_datum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_tuplesort_puttupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_puttupleslot
func F_tuplesort_puttupleslot(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplesort_putgintuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_putgintuple
func F_tuplesort_putgintuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_gettupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_gettupleslot
func F_tuplesort_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_tuplesort_getheaptuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_getheaptuple
func F_tuplesort_getheaptuple(m *base.Module, l0 int32) int32
//go:linkname F_tuplesort_getbrintuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_getbrintuple
func F_tuplesort_getbrintuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tuplesort_getdatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_getdatum
func F_tuplesort_getdatum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_tuplestore_puttupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplestore_puttupleslot
func F_tuplestore_puttupleslot(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_puttuple_common github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tuplestore_puttuple_common
func F_tuplestore_puttuple_common(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_gettupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_gettupleslot
func F_tuplestore_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GetTransactionSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetTransactionSnapshot
func F_GetTransactionSnapshot(m *base.Module) int32
//go:linkname F_InvalidateCatalogSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InvalidateCatalogSnapshot
func F_InvalidateCatalogSnapshot(m *base.Module)
//go:linkname F_PushActiveSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PushActiveSnapshot
func F_PushActiveSnapshot(m *base.Module, l0 int32)
//go:linkname F_PushCopiedSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PushCopiedSnapshot
func F_PushCopiedSnapshot(m *base.Module, l0 int32)
//go:linkname F_PopActiveSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PopActiveSnapshot
func F_PopActiveSnapshot(m *base.Module)
//go:linkname F_RegisterSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RegisterSnapshot
func F_RegisterSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_UnregisterSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_UnregisterSnapshot
func F_UnregisterSnapshot(m *base.Module, l0 int32)
//go:linkname F_AtEOXact_Snapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AtEOXact_Snapshot
func F_AtEOXact_Snapshot(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RestoreSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RestoreSnapshot
func F_RestoreSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_localsub github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_localsub
func F_localsub(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_gmtime github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_gmtime
func F_pg_gmtime(m *base.Module, l0 int32) int32
//go:linkname F_pg_next_dst_boundary github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_next_dst_boundary
func F_pg_next_dst_boundary(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_pg_interpret_timezone_abbrev github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_interpret_timezone_abbrev
func F_pg_interpret_timezone_abbrev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_tzset github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_tzset
func F_pg_tzset(m *base.Module, l0 int32) int32
//go:linkname F_provider_init github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_provider_init
func F_provider_init(m *base.Module) int32
//go:linkname F_binaryheap_add_unordered github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_binaryheap_add_unordered
func F_binaryheap_add_unordered(m *base.Module, l0 int32, l1 int32)
//go:linkname F_binaryheap_build github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_binaryheap_build
func F_binaryheap_build(m *base.Module, l0 int32)
//go:linkname F_binaryheap_replace_first github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_binaryheap_replace_first
func F_binaryheap_replace_first(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BlockRefTableRead github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BlockRefTableRead
func F_BlockRefTableRead(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_controlfile github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_controlfile
func F_get_controlfile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_update_controlfile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_update_controlfile
func F_update_controlfile(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hash_bytes_uint32 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hash_bytes_uint32
func F_hash_bytes_uint32(m *base.Module, l0 int32) int32
//go:linkname F_pg_getnameinfo_all github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_getnameinfo_all
func F_pg_getnameinfo_all(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_IsValidJsonNumber github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_IsValidJsonNumber
func F_IsValidJsonNumber(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeJsonLexContextCstringLen github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeJsonLexContextCstringLen
func F_makeJsonLexContextCstringLen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_freeJsonLexContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_freeJsonLexContext
func F_freeJsonLexContext(m *base.Module, l0 int32)
//go:linkname F_pg_parse_json github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_parse_json
func F_pg_parse_json(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_json_lex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_json_lex
func F_json_lex(m *base.Module, l0 int32) int32
//go:linkname F_parse_object_field github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parse_object_field
func F_parse_object_field(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_array_element github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_parse_array_element
func F_parse_array_element(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replace_percent_placeholders github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_replace_percent_placeholders
func F_replace_percent_placeholders(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_get_line_append github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_get_line_append
func F_pg_get_line_append(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_prng_seed github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_prng_seed
func F_pg_prng_seed(m *base.Module, l0 int32, l1 int64)
//go:linkname F_pg_prng_uint32 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_prng_uint32
func F_pg_prng_uint32(m *base.Module) int32
//go:linkname F_pg_prng_double github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_prng_double
func F_pg_prng_double(m *base.Module, l0 int32) float64
//go:linkname F_psprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_psprintf
func F_psprintf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetDatabasePath github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetDatabasePath
func F_GetDatabasePath(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetRelationPath github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetRelationPath
func F_GetRelationPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_rmtree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_rmtree
func F_rmtree(m *base.Module, l0 int32) int32
//go:linkname F_strtoint github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_strtoint
func F_strtoint(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_clean_ascii github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_clean_ascii
func F_pg_clean_ascii(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeStringInfo
func F_makeStringInfo(m *base.Module) int32
//go:linkname F_resetStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_resetStringInfo
func F_resetStringInfo(m *base.Module, l0 int32)
//go:linkname F_appendStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_appendStringInfo
func F_appendStringInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_enlargeStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_enlargeStringInfo
func F_enlargeStringInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_appendStringInfoVA github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendStringInfoVA
func F_appendStringInfoVA(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_appendStringInfoString github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendStringInfoString
func F_appendStringInfoString(m *base.Module, l0 int32, l1 int32)
//go:linkname F_appendBinaryStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendBinaryStringInfo
func F_appendBinaryStringInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_appendStringInfoChar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_appendStringInfoChar
func F_appendStringInfoChar(m *base.Module, l0 int32, l1 int32)
//go:linkname F_appendStringInfoSpaces github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendStringInfoSpaces
func F_appendStringInfoSpaces(m *base.Module, l0 int32, l1 int32)
//go:linkname F_appendBinaryStringInfoNT github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_appendBinaryStringInfoNT
func F_appendBinaryStringInfoNT(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_convert_case github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_convert_case
func F_convert_case(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_get_decomposed_size github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_decomposed_size
func F_get_decomposed_size(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_decompose_code github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_decompose_code
func F_decompose_code(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pg_utf_mblen_private github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_utf_mblen_private
func F_pg_utf_mblen_private(m *base.Module, l0 int32) int32
//go:linkname F_pg_encoding_verifymbchar github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_encoding_verifymbchar
func F_pg_encoding_verifymbchar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_cryptohash_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_cryptohash_create
func F_pg_cryptohash_create(m *base.Module, l0 int32) int32
//go:linkname F_pg_cryptohash_init github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_cryptohash_init
func F_pg_cryptohash_init(m *base.Module, l0 int32) int32
//go:linkname F_pg_cryptohash_update github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_cryptohash_update
func F_pg_cryptohash_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_cryptohash_final github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_cryptohash_final
func F_pg_cryptohash_final(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_cryptohash_free github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_cryptohash_free
func F_pg_cryptohash_free(m *base.Module, l0 int32)
//go:linkname F_pg_get_encoding_from_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_get_encoding_from_locale
func F_pg_get_encoding_from_locale(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_inet_net_ntop github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_inet_net_ntop
func F_pg_inet_net_ntop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_canonicalize_path_enc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_canonicalize_path_enc
func F_canonicalize_path_enc(m *base.Module, l0 int32)
//go:linkname F_pg_strcasecmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_strcasecmp
func F_pg_strcasecmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_qsort github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_qsort
func F_pg_qsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_qsort_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_qsort_arg
func F_qsort_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_dostr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dostr
func F_dostr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_snprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_snprintf
func F_pg_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_sprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_sprintf
func F_pg_sprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_fprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_fprintf
func F_pg_fprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgmem_dlsym github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgmem_dlsym
func F_pgmem_dlsym(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_estate_setup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_estate_setup
func F_plpgsql_estate_setup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_copy_plpgsql_datums github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_copy_plpgsql_datums
func F_copy_plpgsql_datums(m *base.Module, l0 int32, l1 int32)
//go:linkname F_exec_stmt_block github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_exec_stmt_block
func F_exec_stmt_block(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_exec_cast_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_cast_value
func F_exec_cast_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_exec_assign_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_assign_value
func F_exec_assign_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_revalidate_rectypeid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_revalidate_rectypeid
func F_revalidate_rectypeid(m *base.Module, l0 int32)
//go:linkname F_read_sql_construct github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_read_sql_construct
func F_read_sql_construct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32
//go:linkname F_plpgsql_yylex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_plpgsql_yylex
func F_plpgsql_yylex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_plpgsql_push_back_token github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_plpgsql_push_back_token
func F_plpgsql_push_back_token(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_plpgsql_scanner_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_plpgsql_scanner_errposition
func F_plpgsql_scanner_errposition(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_yyerror
func F_plpgsql_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_r_undouble_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_undouble_2
func F_r_undouble_2(m *base.Module, l0 int32) int32
//go:linkname F_r_Suffix_Verb_Step2a github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_Suffix_Verb_Step2a
func F_r_Suffix_Verb_Step2a(m *base.Module, l0 int32) int32
//go:linkname F_r_Suffix_Noun_Step2a github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_r_Suffix_Noun_Step2a
func F_r_Suffix_Noun_Step2a(m *base.Module, l0 int32) int32
//go:linkname F_r_Suffix_Noun_Step2b github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_Suffix_Noun_Step2b
func F_r_Suffix_Noun_Step2b(m *base.Module, l0 int32) int32
//go:linkname F_r_Suffix_Noun_Step2c1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_Suffix_Noun_Step2c1
func F_r_Suffix_Noun_Step2c1(m *base.Module, l0 int32) int32
//go:linkname F_r_check_vowel_harmony github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_check_vowel_harmony
func F_r_check_vowel_harmony(m *base.Module, l0 int32) int32
//go:linkname F_r_mark_lAr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_mark_lAr
func F_r_mark_lAr(m *base.Module, l0 int32) int32
//go:linkname F_r_mark_possessives github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_mark_possessives
func F_r_mark_possessives(m *base.Module, l0 int32) int32
//go:linkname F_skip_b_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_skip_b_utf8
func F_skip_b_utf8(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_in_grouping_b_U github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_in_grouping_b_U
func F_in_grouping_b_U(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_out_grouping_b_U github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_out_grouping_b_U
func F_out_grouping_b_U(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_find_among github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_find_among
func F_find_among(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_among_b github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_among_b
func F_find_among_b(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_replace_s github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_replace_s
func F_replace_s(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_slice_from_s github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_slice_from_s
func F_slice_from_s(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_slice_del github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_slice_del
func F_slice_del(m *base.Module, l0 int32) int32
//go:linkname F_px_crypt_shacrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_px_crypt_shacrypt
func F_px_crypt_shacrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pullf_read_fixed github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pullf_read_fixed
func F_pullf_read_fixed(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pushf_write github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pushf_write
func F_pushf_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CheckBuiltinCryptoMode github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CheckBuiltinCryptoMode
func F_CheckBuiltinCryptoMode(m *base.Module)
//go:linkname F_pgp_mpi_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_mpi_hash
func F_pgp_mpi_hash(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgp_load_digest github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_load_digest
func F_pgp_load_digest(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_px_gen_salt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_px_gen_salt
func F_px_gen_salt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_px_debug github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_px_debug
func F_px_debug(m *base.Module, l0 int32, l1 int32)
//go:linkname F_px_find_combo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_px_find_combo
func F_px_find_combo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addKey github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_addKey
func F_addKey(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addArc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_addArc
func F_addArc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_hstoreUpgrade github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstoreUpgrade
func F_hstoreUpgrade(m *base.Module, l0 int32) int32
//go:linkname F_hstore_from_text github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hstore_from_text
func F_hstore_from_text(m *base.Module, l0 int32) int32
//go:linkname F_hstoreUniquePairs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstoreUniquePairs
func F_hstoreUniquePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hstorePairs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hstorePairs
func F_hstorePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_val github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_val
func F_get_val(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hstore_akeys github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hstore_akeys
func F_hstore_akeys(m *base.Module, l0 int32) int32
//go:linkname F_hstore_each github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_each
func F_hstore_each(m *base.Module, l0 int32) int32
//go:linkname F_hstoreArrayToPairs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstoreArrayToPairs
func F_hstoreArrayToPairs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hstore_to_array_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_to_array_internal
func F_hstore_to_array_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_array_iterator github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_iterator
func F_array_iterator(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_checkCond github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_checkCond
func F_checkCond(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ltree_gist_alloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltree_gist_alloc
func F_ltree_gist_alloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_parse_lquery github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_parse_lquery
func F_parse_lquery(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_deparse_lquery github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_deparse_lquery
func F_deparse_lquery(m *base.Module, l0 int32) int32
//go:linkname F_inner_subltree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_inner_subltree
func F_inner_subltree(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ltree_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltree_execute
func F_ltree_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gbt_num_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_num_compress
func F_gbt_num_compress(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gbt_num_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_num_fetch
func F_gbt_num_fetch(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gbt_num_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_num_consistent
func F_gbt_num_consistent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_gbt_num_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_num_distance
func F_gbt_num_distance(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_gbt_num_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_num_picksplit
func F_gbt_num_picksplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gbt_var_key_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_var_key_copy
func F_gbt_var_key_copy(m *base.Module, l0 int32) int32
//go:linkname F_gbt_var_union github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_var_union
func F_gbt_var_union(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_gbt_var_node_cp_len github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_var_node_cp_len
func F_gbt_var_node_cp_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gbt_var_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_var_penalty
func F_gbt_var_penalty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_gbt_var_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_var_consistent
func F_gbt_var_consistent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_gin_btree_extract_query github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gin_btree_extract_query
func F_gin_btree_extract_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_query_has_required_values github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_query_has_required_values
func F_query_has_required_values(m *base.Module, l0 int32) int32
//go:linkname F__int_unique github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__int_unique
func F__int_unique(m *base.Module, l0 int32) int32
//go:linkname F_inner_int_inter github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_inner_int_inter
func F_inner_int_inter(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_isort github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isort
func F_isort(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_intarray_add_elem github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_intarray_add_elem
func F_intarray_add_elem(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_update_node github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_update_node
func F_update_node(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_cube_union_v0 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_union_v0
func F_cube_union_v0(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cube_yyparse github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cube_yyparse
func F_cube_yyparse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_cube_yy_scan_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cube_yy_scan_buffer
func F_cube_yy_scan_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_cube_scanner_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_scanner_finish
func F_cube_scanner_finish(m *base.Module, l0 int32)
//go:linkname F_seg_yy_scan_string github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_seg_yy_scan_string
func F_seg_yy_scan_string(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_initBloomState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_initBloomState
func F_initBloomState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_signValue github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_signValue
func F_signValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_BloomNewBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BloomNewBuffer
func F_BloomNewBuffer(m *base.Module, l0 int32) int32
//go:linkname F_pgstatindex_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstatindex_impl
func F_pgstatindex_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_relpages_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_relpages_impl
func F_pg_relpages_impl(m *base.Module, l0 int32) int64
//go:linkname F_pgstat_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_relation
func F_pgstat_relation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_uuid_generate_time github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_uuid_generate_time
func F_uuid_generate_time(m *base.Module, l0 int32)
//go:linkname F_uuid_unparse github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_uuid_unparse
func F_uuid_unparse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_collect_visibility_data github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_collect_visibility_data
func F_collect_visibility_data(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_verify_brin_page github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_verify_brin_page
func F_verify_brin_page(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_bt_page_items_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bt_page_items_internal
func F_bt_page_items_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_verify_hash_page github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_verify_hash_page
func F_verify_hash_page(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_page_from_raw github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_page_from_raw
func F_get_page_from_raw(m *base.Module, l0 int32) int32
//go:linkname F_entry_alloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_entry_alloc
func F_entry_alloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_stat_statements_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_stat_statements_internal
func F_pg_stat_statements_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_qtext_load_file github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_qtext_load_file
func F_qtext_load_file(m *base.Module, l0 int32) int32
//go:linkname F_qtext_store github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_qtext_store
func F_qtext_store(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_InitBitVector github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InitBitVector
func F_InitBitVector(m *base.Module, l0 int32) int32
//go:linkname F_HnswInit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HnswInit
func F_HnswInit(m *base.Module)
//go:linkname F_HnswUpdateNeighborsOnDisk github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HnswUpdateNeighborsOnDisk
func F_HnswUpdateNeighborsOnDisk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_HnswInsertTupleOnDisk github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_HnswInsertTupleOnDisk
func F_HnswInsertTupleOnDisk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pointerhash_insert_hash_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pointerhash_insert_hash_internal
func F_pointerhash_insert_hash_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_HnswInitSupport github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_HnswInitSupport
func F_HnswInitSupport(m *base.Module, l0 int32, l1 int32)
//go:linkname F_HnswInitNeighbors github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_HnswInitNeighbors
func F_HnswInitNeighbors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_HnswLoadElementImpl github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HnswLoadElementImpl
func F_HnswLoadElementImpl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_HnswFindElementNeighbors github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HnswFindElementNeighbors
func F_HnswFindElementNeighbors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_HnswGetTypeInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HnswGetTypeInfo
func F_HnswGetTypeInfo(m *base.Module, l0 int32) int32
//go:linkname F_BuildIndex_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BuildIndex_2
func F_BuildIndex_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_IvfflatInit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_IvfflatInit
func F_IvfflatInit(m *base.Module)
//go:linkname F_IvfflatGetMetaPageInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_IvfflatGetMetaPageInfo
func F_IvfflatGetMetaPageInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CheckElement_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckElement_3
func F_CheckElement_3(m *base.Module, l0 float32)
//go:linkname F_CheckDim_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckDim_3
func F_CheckDim_3(m *base.Module, l0 int32)
//go:linkname F__Exit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__Exit
func F__Exit(m *base.Module, l0 int32)
//go:linkname F_R github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_R
func F_R(m *base.Module, l0 float64) float64
//go:linkname F___isspace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___isspace
func F___isspace(m *base.Module, l0 int32) int32
//go:linkname F_bsearch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bsearch
func F_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F___clock_gettime github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F___clock_gettime
func F___clock_gettime(m *base.Module, l0 int32, l1 int32)
//go:linkname F_close github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_close
func F_close(m *base.Module, l0 int32) int32
//go:linkname F___rem_pio2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___rem_pio2
func F___rem_pio2(m *base.Module, l0 float64, l1 int32) int32
//go:linkname F___sin github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___sin
func F___sin(m *base.Module, l0 float64, l1 float64, l2 int32) float64
//go:linkname F___memcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___memcpy
func F___memcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_time github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_time
func F_time(m *base.Module) int64
//go:linkname F_gettimeofday github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gettimeofday
func F_gettimeofday(m *base.Module, l0 int32)
//go:linkname F_erfc2 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_erfc2
func F_erfc2(m *base.Module, l0 int32, l1 float64) float64
//go:linkname F_exp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exp
func F_exp(m *base.Module, l0 float64) float64
//go:linkname F___memset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___memset
func F___memset(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_fflush github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fflush
func F_fflush(m *base.Module, l0 int32) int32
//go:linkname F_fopen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fopen
func F_fopen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fiprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fiprintf
func F_fiprintf(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_fread github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fread
func F_fread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_freopen github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_freopen
func F_freopen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___fseeko_unlocked github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___fseeko_unlocked
func F___fseeko_unlocked(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F___fstatat github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___fstatat
func F___fstatat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ftruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ftruncate
func F_ftruncate(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_fwrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fwrite
func F_fwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getrusage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getrusage
func F_getrusage(m *base.Module, l0 int32)
//go:linkname F_isalnum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isalnum
func F_isalnum(m *base.Module, l0 int32) int32
//go:linkname F_isatty github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_isatty
func F_isatty(m *base.Module, l0 int32) int32
//go:linkname F_kill github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_kill
func F_kill(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_log github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_log
func F_log(m *base.Module, l0 float64) float64
//go:linkname F___lseek github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___lseek
func F___lseek(m *base.Module, l0 int32, l1 int64, l2 int32) int64
//go:linkname F_memchr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_memchr
func F_memchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_memcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_memcmp
func F_memcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___gmtime_r github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___gmtime_r
func F___gmtime_r(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___mmap github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___mmap
func F___mmap(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___get_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___get_locale
func F___get_locale(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___loc_is_allocated github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___loc_is_allocated
func F___loc_is_allocated(m *base.Module, l0 int32) int32
//go:linkname F_open github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_open
func F_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_read github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_read
func F_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rmdir github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_rmdir
func F_rmdir(m *base.Module, l0 int32) int32
//go:linkname F_scalbn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_scalbn
func F_scalbn(m *base.Module, l0 float64, l1 int32) float64
//go:linkname F___shm_mapname github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___shm_mapname
func F___shm_mapname(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sigemptyset github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sigemptyset
func F_sigemptyset(m *base.Module, l0 int32)
//go:linkname F_sigprocmask github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sigprocmask
func F_sigprocmask(m *base.Module, l0 int32, l1 int32)
//go:linkname F_snprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_snprintf
func F_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_sscanf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sscanf
func F_sscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strchr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strchr
func F_strchr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strcmp
func F_strcmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strcpy
func F_strcpy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___strftime_l github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___strftime_l
func F___strftime_l(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_strlen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strlen
func F_strlen(m *base.Module, l0 int32) int32
//go:linkname F_strncmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strncmp
func F_strncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strncpy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strncpy
func F_strncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strstr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strstr
func F_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strtof github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strtof
func F_strtof(m *base.Module, l0 int32, l1 int32) float32
//go:linkname F_strtod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strtod
func F_strtod(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_strtol github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strtol
func F_strtol(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___syscall_ret github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___syscall_ret
func F___syscall_ret(m *base.Module, l0 int32) int32
//go:linkname F_casemap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_casemap
func F_casemap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_umask github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_umask
func F_umask(m *base.Module, l0 int32) int32
//go:linkname F_unlink github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_unlink
func F_unlink(m *base.Module, l0 int32) int32
//go:linkname F_out github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_out
func F_out(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pop_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pop_arg
func F_pop_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pad github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pad
func F_pad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_wcslen github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_wcslen
func F_wcslen(m *base.Module, l0 int32) int32
//go:linkname F_wcrtomb github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_wcrtomb
func F_wcrtomb(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_emscripten_builtin_malloc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_emscripten_builtin_malloc
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32
//go:linkname F_emscripten_builtin_realloc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_emscripten_builtin_realloc
func F_emscripten_builtin_realloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___wasm_longjmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___wasm_longjmp
func F___wasm_longjmp(m *base.Module, l0 int32, l1 int32)
//go:linkname F___multf3 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___multf3
func F___multf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___udivmodti4 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___udivmodti4
func F___udivmodti4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname Fn13825 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13825
func Fn13825(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname Fn13828 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13828
func Fn13828(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname Fn13832 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13832
func Fn13832(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13838 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13838
func Fn13838(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname Fn13839 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13839
func Fn13839(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname Fn13840 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13840
func Fn13840(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13841 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13841
func Fn13841(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname Fn13845 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13845
func Fn13845(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname Fn13847 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13847
func Fn13847(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname Fn13849 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13849
func Fn13849(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13852 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13852
func Fn13852(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname Fn13856 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13856
func Fn13856(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13859 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13859
func Fn13859(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname Fn13860 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13860
func Fn13860(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname Fn13863 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13863
func Fn13863(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname Fn13867 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13867
func Fn13867(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13869 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13869
func Fn13869(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13870 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13870
func Fn13870(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13876 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13876
func Fn13876(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname Fn13883 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13883
func Fn13883(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13884 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13884
func Fn13884(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13888 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13888
func Fn13888(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13890 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13890
func Fn13890(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13891 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13891
func Fn13891(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname Fn13896 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13896
func Fn13896(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13898 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13898
func Fn13898(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13900 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13900
func Fn13900(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13902 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13902
func Fn13902(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname Fn13903 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13903
func Fn13903(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname Fn13905 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13905
func Fn13905(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13906 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13906
func Fn13906(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname Fn13907 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13907
func Fn13907(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13909 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13909
func Fn13909(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname Fn13912 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13912
func Fn13912(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname Fn13913 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13913
func Fn13913(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13914 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13914
func Fn13914(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13918 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13918
func Fn13918(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname Fn13919 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13919
func Fn13919(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13920 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13920
func Fn13920(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13926 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13926
func Fn13926(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13930 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13930
func Fn13930(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13931 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13931
func Fn13931(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13934 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13934
func Fn13934(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13935 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13935
func Fn13935(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13943 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13943
func Fn13943(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13948 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13948
func Fn13948(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13949 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13949
func Fn13949(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13955 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13955
func Fn13955(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13959 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13959
func Fn13959(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13964 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13964
func Fn13964(m *base.Module, l0 int64) int32
//go:linkname Fn13966 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13966
func Fn13966(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13967 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13967
func Fn13967(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13969 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13969
func Fn13969(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname Fn13977 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13977
func Fn13977(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13984 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13984
func Fn13984(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13985 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13985
func Fn13985(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13987 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13987
func Fn13987(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13988 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13988
func Fn13988(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13991 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13991
func Fn13991(m *base.Module, l0 int32, l1 int32)
//go:linkname Fn13992 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn13992
func Fn13992(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13993 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13993
func Fn13993(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn14003 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn14003
func Fn14003(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32
//go:linkname Fn14004 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn14004
func Fn14004(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32
//go:linkname Fn14006 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn14006
func Fn14006(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn14007 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn14007
func Fn14007(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname Fn14010 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn14010
func Fn14010(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn14013 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn14013
func Fn14013(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname Fn14015 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn14015
func Fn14015(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn14016 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn14016
func Fn14016(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname Fn14017 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn14017
func Fn14017(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn14021 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.Fn14021
func Fn14021(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn14022 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn14022
func Fn14022(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn14024 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn14024
func Fn14024(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
