package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	_ "unsafe"
)
//go:linkname F_add_values_to_range github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_values_to_range
func F_add_values_to_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_brinsummarize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brinsummarize
func F_brinsummarize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_minmax_get_strategy_procinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_minmax_get_strategy_procinfo
func F_minmax_get_strategy_procinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_brin_doinsert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_doinsert
func F_brin_doinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_brin_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_brin_form_tuple
func F_brin_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_brin_new_memtuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_new_memtuple
func F_brin_new_memtuple(m *base.Module, l0 int32) int32
//go:linkname F_brin_memtuple_initialize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_memtuple_initialize
func F_brin_memtuple_initialize(m *base.Module, l0 int32, l1 int32)
//go:linkname F_brin_xlog_insert_update github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_xlog_insert_update
func F_brin_xlog_insert_update(m *base.Module, l0 int32, l1 int32)
//go:linkname F_free_attrmap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_free_attrmap
func F_free_attrmap(m *base.Module, l0 int32)
//go:linkname F_build_attrmap_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_attrmap_by_name
func F_build_attrmap_by_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_attrmap_by_name_if_req github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_build_attrmap_by_name_if_req
func F_build_attrmap_by_name_if_req(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mask_unused_space github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mask_unused_space
func F_mask_unused_space(m *base.Module, l0 int32)
//go:linkname F_detoast_attr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_detoast_attr
func F_detoast_attr(m *base.Module, l0 int32) int32
//go:linkname F_getmissingattr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getmissingattr
func F_getmissingattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_compute_data_size github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_compute_data_size
func F_heap_compute_data_size(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_fill_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_fill_tuple
func F_heap_fill_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_fill_val github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fill_val
func F_fill_val(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_heap_attisnull github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_attisnull
func F_heap_attisnull(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_form_tuple
func F_heap_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_deform_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_deform_tuple
func F_heap_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_heap_copy_minimal_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_copy_minimal_tuple
func F_heap_copy_minimal_tuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_minimal_tuple_from_heap_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_minimal_tuple_from_heap_tuple
func F_minimal_tuple_from_heap_tuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_form_tuple_context github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_index_form_tuple_context
func F_index_form_tuple_context(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_index_deform_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_deform_tuple
func F_index_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CopyIndexTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CopyIndexTuple
func F_CopyIndexTuple(m *base.Module, l0 int32) int32
//go:linkname F_SendRowDescriptionMessage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SendRowDescriptionMessage
func F_SendRowDescriptionMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_relation_open github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_relation_open
func F_relation_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_try_relation_open github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_try_relation_open
func F_try_relation_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_relation_close github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_relation_close
func F_relation_close(m *base.Module, l0 int32, l1 int32)
//go:linkname F_add_local_int_reloption github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_local_int_reloption
func F_add_local_int_reloption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_transformRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_transformRelOptions
func F_transformRelOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_build_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_reloptions
func F_build_reloptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_heap_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_reloptions
func F_heap_reloptions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_attribute_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_attribute_reloptions
func F_attribute_reloptions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ScanKeyInit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ScanKeyInit
func F_ScanKeyInit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_TidStoreCreateLocal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TidStoreCreateLocal
func F_TidStoreCreateLocal(m *base.Module, l0 int32) int32
//go:linkname F_shared_ts_free_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_shared_ts_free_recurse
func F_shared_ts_free_recurse(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_TidStoreMemoryUsage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TidStoreMemoryUsage
func F_TidStoreMemoryUsage(m *base.Module, l0 int32) int32
//go:linkname F_toast_open_indexes github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_toast_open_indexes
func F_toast_open_indexes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_toast_get_valid_index github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_toast_get_valid_index
func F_toast_get_valid_index(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_convert_tuples_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_convert_tuples_by_name
func F_convert_tuples_by_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_convert_tuples_by_name_attrmap github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_convert_tuples_by_name_attrmap
func F_convert_tuples_by_name_attrmap(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_execute_attr_map_slot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_execute_attr_map_slot
func F_execute_attr_map_slot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_populate_compact_attribute github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_populate_compact_attribute
func F_populate_compact_attribute(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateTemplateTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateTemplateTupleDesc
func F_CreateTemplateTupleDesc(m *base.Module, l0 int32) int32
//go:linkname F_CreateTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateTupleDesc
func F_CreateTupleDesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CreateTupleDescCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateTupleDescCopy
func F_CreateTupleDescCopy(m *base.Module, l0 int32) int32
//go:linkname F_CreateTupleDescCopyConstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateTupleDescCopyConstr
func F_CreateTupleDescCopyConstr(m *base.Module, l0 int32) int32
//go:linkname F_FreeTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeTupleDesc
func F_FreeTupleDesc(m *base.Module, l0 int32)
//go:linkname F_IncrTupleDescRefCount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_IncrTupleDescRefCount
func F_IncrTupleDescRefCount(m *base.Module, l0 int32)
//go:linkname F_DecrTupleDescRefCount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecrTupleDescRefCount
func F_DecrTupleDescRefCount(m *base.Module, l0 int32)
//go:linkname F_TupleDescInitEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TupleDescInitEntry
func F_TupleDescInitEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ginStepRight github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ginStepRight
func F_ginStepRight(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ginFinishSplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ginFinishSplit
func F_ginFinishSplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ginInsertValue github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ginInsertValue
func F_ginInsertValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ginInitBA github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ginInitBA
func F_ginInitBA(m *base.Module, l0 int32)
//go:linkname F_ginGetBAEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ginGetBAEntry
func F_ginGetBAEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ginEntryInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ginEntryInsert
func F_ginEntryInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F__gin_parallel_scan_and_build github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__gin_parallel_scan_and_build
func F__gin_parallel_scan_and_build(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_GinBufferInit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GinBufferInit
func F_GinBufferInit(m *base.Module, l0 int32) int32
//go:linkname F_GinBufferStoreTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GinBufferStoreTuple
func F_GinBufferStoreTuple(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ginFlushBuildState github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ginFlushBuildState
func F_ginFlushBuildState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ginPostingListDecode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ginPostingListDecode
func F_ginPostingListDecode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_initGinState github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initGinState
func F_initGinState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_gintuple_get_attrnum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gintuple_get_attrnum
func F_gintuple_get_attrnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gintuple_get_key github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gintuple_get_key
func F_gintuple_get_key(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GinNewBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GinNewBuffer
func F_GinNewBuffer(m *base.Module, l0 int32) int32
//go:linkname F_ginCompareAttEntries github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ginCompareAttEntries
func F_ginCompareAttEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ginExtractEntries github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ginExtractEntries
func F_ginExtractEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_initGISTstate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_initGISTstate
func F_initGISTstate(m *base.Module, l0 int32) int32
//go:linkname F_gistdoinsert github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gistdoinsert
func F_gistdoinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_createTempGistContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_createTempGistContext
func F_createTempGistContext(m *base.Module) int32
//go:linkname F_gist_indexsortbuild_levelstate_flush github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gist_indexsortbuild_levelstate_flush
func F_gist_indexsortbuild_levelstate_flush(m *base.Module, l0 int32, l1 int32)
//go:linkname F_gistProcessEmptyingQueue github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistProcessEmptyingQueue
func F_gistProcessEmptyingQueue(m *base.Module, l0 int32)
//go:linkname F_gistbufferinginserttuples github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gistbufferinginserttuples
func F_gistbufferinginserttuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_gistGetNodeBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gistGetNodeBuffer
func F_gistGetNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_gistPushItupToNodeBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistPushItupToNodeBuffer
func F_gistPushItupToNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_index_getattr_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_getattr_2
func F_index_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gistfillbuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gistfillbuffer
func F_gistfillbuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_gistFormTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gistFormTuple
func F_gistFormTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_gistgetadjusted github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistgetadjusted
func F_gistgetadjusted(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gistchoose github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gistchoose
func F_gistchoose(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gistcheckpage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistcheckpage
func F_gistcheckpage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_gistvacuumscan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gistvacuumscan
func F_gistvacuumscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_hashint8extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hashint8extended
func F_hashint8extended(m *base.Module, l0 int32) int32
//go:linkname F__hash_doinsert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__hash_doinsert
func F__hash_doinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__hash_dropscanbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__hash_dropscanbuf
func F__hash_dropscanbuf(m *base.Module, l0 int32)
//go:linkname F__hash_init github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__hash_init
func F__hash_init(m *base.Module, l0 int32, l1 float64, l2 int32) int32
//go:linkname F__hash_kill_items github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_kill_items
func F__hash_kill_items(m *base.Module, l0 int32)
//go:linkname F_heap_getnext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_getnext
func F_heap_getnext(m *base.Module, l0 int32) int32
//go:linkname F_heap_getattr_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_getattr_1
func F_heap_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_FreeBulkInsertState github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FreeBulkInsertState
func F_FreeBulkInsertState(m *base.Module, l0 int32)
//go:linkname F_log_heap_new_cid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_log_heap_new_cid
func F_log_heap_new_cid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_simple_heap_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_simple_heap_insert
func F_simple_heap_insert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_DoesMultiXactIdConflict github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DoesMultiXactIdConflict
func F_DoesMultiXactIdConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_Do_MultiXactIdWait github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Do_MultiXactIdWait
func F_Do_MultiXactIdWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_ExtractReplicaIdentity github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExtractReplicaIdentity
func F_ExtractReplicaIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_simple_heap_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_simple_heap_delete
func F_simple_heap_delete(m *base.Module, l0 int32, l1 int32)
//go:linkname F_HeapTupleGetUpdateXid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_HeapTupleGetUpdateXid
func F_HeapTupleGetUpdateXid(m *base.Module, l0 int32) int32
//go:linkname F_HeapTupleSetHintBits github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_HeapTupleSetHintBits
func F_HeapTupleSetHintBits(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_HeapTupleSatisfiesVacuum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HeapTupleSatisfiesVacuum
func F_HeapTupleSatisfiesVacuum(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_HeapTupleHeaderIsOnlyLocked github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_HeapTupleHeaderIsOnlyLocked
func F_HeapTupleHeaderIsOnlyLocked(m *base.Module, l0 int32) int32
//go:linkname F_HeapTupleSatisfiesVisibility github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_HeapTupleSatisfiesVisibility
func F_HeapTupleSatisfiesVisibility(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_toast_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_heap_toast_delete
func F_heap_toast_delete(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_toast_flatten_tuple_to_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_toast_flatten_tuple_to_datum
func F_toast_flatten_tuple_to_datum(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lazy_check_wraparound_failsafe github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lazy_check_wraparound_failsafe
func F_lazy_check_wraparound_failsafe(m *base.Module, l0 int32) int32
//go:linkname F_visibilitymap_clear github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_visibilitymap_clear
func F_visibilitymap_clear(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_vm_readbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_vm_readbuf
func F_vm_readbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_visibilitymap_set github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_visibilitymap_set
func F_visibilitymap_set(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_visibilitymap_get_status github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_visibilitymap_get_status
func F_visibilitymap_get_status(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_visibilitymap_count github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_visibilitymap_count
func F_visibilitymap_count(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetIndexAmRoutineByAmId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetIndexAmRoutineByAmId
func F_GetIndexAmRoutineByAmId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_IndexAmTranslateStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_IndexAmTranslateStrategy
func F_IndexAmTranslateStrategy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_opclass_for_family_datatype github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_opclass_for_family_datatype
func F_opclass_for_family_datatype(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RelationGetIndexScan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationGetIndexScan
func F_RelationGetIndexScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_BuildIndexValueDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BuildIndexValueDescription
func F_BuildIndexValueDescription(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_systable_getnext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_systable_getnext
func F_systable_getnext(m *base.Module, l0 int32) int32
//go:linkname F_HandleConcurrentAbort github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_HandleConcurrentAbort
func F_HandleConcurrentAbort(m *base.Module)
//go:linkname F_systable_endscan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_systable_endscan
func F_systable_endscan(m *base.Module, l0 int32)
//go:linkname F_systable_beginscan_ordered github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_systable_beginscan_ordered
func F_systable_beginscan_ordered(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_systable_getnext_ordered github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_systable_getnext_ordered
func F_systable_getnext_ordered(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_systable_inplace_update_cancel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_systable_inplace_update_cancel
func F_systable_inplace_update_cancel(m *base.Module, l0 int32)
//go:linkname F_index_open github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_open
func F_index_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_beginscan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_beginscan
func F_index_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_index_rescan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_rescan
func F_index_rescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_index_endscan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_endscan
func F_index_endscan(m *base.Module, l0 int32)
//go:linkname F_index_beginscan_parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_beginscan_parallel
func F_index_beginscan_parallel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_index_getnext_tid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_index_getnext_tid
func F_index_getnext_tid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_fetch_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_index_fetch_heap
func F_index_fetch_heap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_getprocinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_getprocinfo
func F_index_getprocinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_btoidvectorcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_btoidvectorcmp
func F_btoidvectorcmp(m *base.Module, l0 int32) int32
//go:linkname F__bt_update_posting github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_update_posting
func F__bt_update_posting(m *base.Module, l0 int32)
//go:linkname F__bt_swap_posting github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_swap_posting
func F__bt_swap_posting(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_insertonpg github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_insertonpg
func F__bt_insertonpg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32)
//go:linkname F__bt_getstackbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_getstackbuf
func F__bt_getstackbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_checkpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_checkpage
func F__bt_checkpage(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_getbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_getbuf
func F__bt_getbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_getroot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_getroot
func F__bt_getroot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_relandgetbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_relandgetbuf
func F__bt_relandgetbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_metaversion github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__bt_metaversion
func F__bt_metaversion(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F__bt_setup_array_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_setup_array_cmp
func F__bt_setup_array_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F__bt_find_extreme_element github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_find_extreme_element
func F__bt_find_extreme_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F__bt_mark_scankey_required github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_mark_scankey_required
func F__bt_mark_scankey_required(m *base.Module, l0 int32)
//go:linkname F__bt_compare_scankey_args github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_compare_scankey_args
func F__bt_compare_scankey_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F__bt_parallel_seize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_parallel_seize
func F__bt_parallel_seize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_search github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_search
func F__bt_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F__bt_compare github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_compare
func F__bt_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_readnextpage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_readnextpage
func F__bt_readnextpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F__bt_readfirstpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_readfirstpage
func F__bt_readfirstpage(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_next github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_next
func F__bt_next(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_mkscankey github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_mkscankey
func F__bt_mkscankey(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_binsrch_array_skey github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_binsrch_array_skey
func F__bt_binsrch_array_skey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F__bt_start_array_keys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_start_array_keys
func F__bt_start_array_keys(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_start_prim_scan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_start_prim_scan
func F__bt_start_prim_scan(m *base.Module, l0 int32) int32
//go:linkname F__bt_advance_array_keys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_advance_array_keys
func F__bt_advance_array_keys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F__bt_killitems github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_killitems
func F__bt_killitems(m *base.Module, l0 int32)
//go:linkname F__bt_restore_page github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_restore_page
func F__bt_restore_page(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_desc_recompress_leaf github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_desc_recompress_leaf
func F_desc_recompress_leaf(m *base.Module, l0 int32, l1 int32)
//go:linkname F_array_desc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_array_desc
func F_array_desc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_spg_key_orderbys_distances github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_spg_key_orderbys_distances
func F_spg_key_orderbys_distances(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_spgWalk github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_spgWalk
func F_spgWalk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_fillTypeDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fillTypeDesc
func F_fillTypeDesc(m *base.Module, l0 int32, l1 int32)
//go:linkname F_initSpGistState github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_initSpGistState
func F_initSpGistState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SpGistUpdateMetaPage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SpGistUpdateMetaPage
func F_SpGistUpdateMetaPage(m *base.Module, l0 int32)
//go:linkname F_allocNewBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_allocNewBuffer
func F_allocNewBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sequence_close github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sequence_close
func F_sequence_close(m *base.Module, l0 int32, l1 int32)
//go:linkname F_table_open github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_table_open
func F_table_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_try_table_open github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_try_table_open
func F_try_table_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_openrv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_table_openrv
func F_table_openrv(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_beginscan_catalog github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_table_beginscan_catalog
func F_table_beginscan_catalog(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_table_parallelscan_estimate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_table_parallelscan_estimate
func F_table_parallelscan_estimate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_beginscan_parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_table_beginscan_parallel
func F_table_beginscan_parallel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TransactionIdSetTreeStatus github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TransactionIdSetTreeStatus
func F_TransactionIdSetTreeStatus(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64)
//go:linkname F_TransactionTreeSetCommitTsData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TransactionTreeSetCommitTsData
func F_TransactionTreeSetCommitTsData(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32)
//go:linkname F_TransactionIdGetCommitTsData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransactionIdGetCommitTsData
func F_TransactionIdGetCommitTsData(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SetCommitTsLimit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SetCommitTsLimit
func F_SetCommitTsLimit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_MultiXactIdCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MultiXactIdCreate
func F_MultiXactIdCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_MultiXactIdCreateFromMembers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MultiXactIdCreateFromMembers
func F_MultiXactIdCreateFromMembers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetMultiXactIdMembers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetMultiXactIdMembers
func F_GetMultiXactIdMembers(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MultiXactIdIsRunning github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MultiXactIdIsRunning
func F_MultiXactIdIsRunning(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MultiXactIdSetOldestMember github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MultiXactIdSetOldestMember
func F_MultiXactIdSetOldestMember(m *base.Module)
//go:linkname F_SetMultiXactIdLimit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SetMultiXactIdLimit
func F_SetMultiXactIdLimit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_MultiXactSetNextMXact github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MultiXactSetNextMXact
func F_MultiXactSetNextMXact(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetOldestMultiXactId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetOldestMultiXactId
func F_GetOldestMultiXactId(m *base.Module) int32
//go:linkname F_CreateParallelContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateParallelContext
func F_CreateParallelContext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_InitializeParallelDSM github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InitializeParallelDSM
func F_InitializeParallelDSM(m *base.Module, l0 int32)
//go:linkname F_WaitForParallelWorkersToFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WaitForParallelWorkersToFinish
func F_WaitForParallelWorkersToFinish(m *base.Module, l0 int32)
//go:linkname F_WaitForParallelWorkersToExit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WaitForParallelWorkersToExit
func F_WaitForParallelWorkersToExit(m *base.Module, l0 int32)
//go:linkname F_LaunchParallelWorkers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LaunchParallelWorkers
func F_LaunchParallelWorkers(m *base.Module, l0 int32)
//go:linkname F_WaitForParallelWorkersToAttach github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WaitForParallelWorkersToAttach
func F_WaitForParallelWorkersToAttach(m *base.Module, l0 int32)
//go:linkname F_DestroyParallelContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DestroyParallelContext
func F_DestroyParallelContext(m *base.Module, l0 int32)
//go:linkname F_AtEOSubXact_Parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtEOSubXact_Parallel
func F_AtEOSubXact_Parallel(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AtEOXact_Parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtEOXact_Parallel
func F_AtEOXact_Parallel(m *base.Module, l0 int32)
//go:linkname F_SimpleLruInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SimpleLruInit
func F_SimpleLruInit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_check_slru_buffers github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_slru_buffers
func F_check_slru_buffers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SimpleLruZeroPage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SimpleLruZeroPage
func F_SimpleLruZeroPage(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_SimpleLruWaitIO github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SimpleLruWaitIO
func F_SimpleLruWaitIO(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SlruInternalWritePage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SlruInternalWritePage
func F_SlruInternalWritePage(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_SimpleLruReadPage_ReadOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SimpleLruReadPage_ReadOnly
func F_SimpleLruReadPage_ReadOnly(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_SimpleLruWritePage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SimpleLruWritePage
func F_SimpleLruWritePage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SimpleLruTruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SimpleLruTruncate
func F_SimpleLruTruncate(m *base.Module, l0 int32, l1 int64)
//go:linkname F_SlruScanDirectory github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SlruScanDirectory
func F_SlruScanDirectory(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SlruSyncFileTag github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SlruSyncFileTag
func F_SlruSyncFileTag(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExtendSUBTRANS github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExtendSUBTRANS
func F_ExtendSUBTRANS(m *base.Module, l0 int32)
//go:linkname F_TransactionIdDidCommit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransactionIdDidCommit
func F_TransactionIdDidCommit(m *base.Module, l0 int32) int32
//go:linkname F_TransactionIdPrecedes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TransactionIdPrecedes
func F_TransactionIdPrecedes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TransactionIdDidAbort github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TransactionIdDidAbort
func F_TransactionIdDidAbort(m *base.Module, l0 int32) int32
//go:linkname F_TransactionIdCommitTree github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TransactionIdCommitTree
func F_TransactionIdCommitTree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_TransactionIdAbortTree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransactionIdAbortTree
func F_TransactionIdAbortTree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AtAbort_Twophase github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtAbort_Twophase
func F_AtAbort_Twophase(m *base.Module)
//go:linkname F_TwoPhaseGetXidByVirtualXID github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TwoPhaseGetXidByVirtualXID
func F_TwoPhaseGetXidByVirtualXID(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TwoPhaseGetDummyProcNumber github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TwoPhaseGetDummyProcNumber
func F_TwoPhaseGetDummyProcNumber(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReadTwoPhaseFile github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadTwoPhaseFile
func F_ReadTwoPhaseFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_XlogReadTwoPhaseData github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XlogReadTwoPhaseData
func F_XlogReadTwoPhaseData(m *base.Module, l0 int64, l1 int32, l2 int32)
//go:linkname F_RemoveTwoPhaseFile github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RemoveTwoPhaseFile
func F_RemoveTwoPhaseFile(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PrepareRedoRemove github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PrepareRedoRemove
func F_PrepareRedoRemove(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetNewTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetNewTransactionId
func F_GetNewTransactionId(m *base.Module, l0 int32) int64
//go:linkname F_AdvanceOldestClogXid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AdvanceOldestClogXid
func F_AdvanceOldestClogXid(m *base.Module, l0 int32)
//go:linkname F_SetTransactionIdLimit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SetTransactionIdLimit
func F_SetTransactionIdLimit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetNewObjectId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetNewObjectId
func F_GetNewObjectId(m *base.Module) int32
//go:linkname F_GetTopTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetTopTransactionId
func F_GetTopTransactionId(m *base.Module) int32
//go:linkname F_GetCurrentCommandId github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetCurrentCommandId
func F_GetCurrentCommandId(m *base.Module, l0 int32) int32
//go:linkname F_CommandCounterIncrement github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CommandCounterIncrement
func F_CommandCounterIncrement(m *base.Module)
//go:linkname F_StartTransactionCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_StartTransactionCommand
func F_StartTransactionCommand(m *base.Module)
//go:linkname F_ShowTransactionStateRec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ShowTransactionStateRec
func F_ShowTransactionStateRec(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CommitTransactionCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CommitTransactionCommand
func F_CommitTransactionCommand(m *base.Module)
//go:linkname F_AbortCurrentTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AbortCurrentTransaction
func F_AbortCurrentTransaction(m *base.Module)
//go:linkname F_EndTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EndTransactionBlock
func F_EndTransactionBlock(m *base.Module, l0 int32) int32
//go:linkname F_PopTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PopTransaction
func F_PopTransaction(m *base.Module)
//go:linkname F_RecordTransactionAbort github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RecordTransactionAbort
func F_RecordTransactionAbort(m *base.Module, l0 int32) int32
//go:linkname F_AbortOutOfAnyTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AbortOutOfAnyTransaction
func F_AbortOutOfAnyTransaction(m *base.Module)
//go:linkname F_XactLogCommitRecord github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XactLogCommitRecord
func F_XactLogCommitRecord(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int64
//go:linkname F_XactLogAbortRecord github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XactLogAbortRecord
func F_XactLogAbortRecord(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int64
//go:linkname F_XLogFlush github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogFlush
func F_XLogFlush(m *base.Module, l0 int64)
//go:linkname F_WaitXLogInsertionsToFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WaitXLogInsertionsToFinish
func F_WaitXLogInsertionsToFinish(m *base.Module, l0 int64) int64
//go:linkname F_XLogWrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogWrite
func F_XLogWrite(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RecoveryInProgress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RecoveryInProgress
func F_RecoveryInProgress(m *base.Module) int32
//go:linkname F_XLogSetAsyncXactLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogSetAsyncXactLSN
func F_XLogSetAsyncXactLSN(m *base.Module, l0 int64)
//go:linkname F_XLogFileName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogFileName
func F_XLogFileName(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_AdvanceXLInsertBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AdvanceXLInsertBuffer
func F_AdvanceXLInsertBuffer(m *base.Module, l0 int64, l1 int32, l2 int32)
//go:linkname F_XLogNeedsFlush github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogNeedsFlush
func F_XLogNeedsFlush(m *base.Module, l0 int64) int32
//go:linkname F_XLogFileInitInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogFileInitInternal
func F_XLogFileInitInternal(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ReadControlFile github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReadControlFile
func F_ReadControlFile(m *base.Module)
//go:linkname F_XLOGShmemSize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XLOGShmemSize
func F_XLOGShmemSize(m *base.Module) int32
//go:linkname F_UpdateFullPageWrites github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UpdateFullPageWrites
func F_UpdateFullPageWrites(m *base.Module)
//go:linkname F_GetInsertRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetInsertRecPtr
func F_GetInsertRecPtr(m *base.Module) int64
//go:linkname F_CreateRestartPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CreateRestartPoint
func F_CreateRestartPoint(m *base.Module, l0 int32) int32
//go:linkname F_CreateCheckPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateCheckPoint
func F_CreateCheckPoint(m *base.Module, l0 int32) int32
//go:linkname F_GetXLogInsertRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetXLogInsertRecPtr
func F_GetXLogInsertRecPtr(m *base.Module) int64
//go:linkname F_SetWalWriterSleeping github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SetWalWriterSleeping
func F_SetWalWriterSleeping(m *base.Module, l0 int32)
//go:linkname F_RestoreArchivedFile github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RestoreArchivedFile
func F_RestoreArchivedFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_KeepFileRestoredFromArchive github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_KeepFileRestoredFromArchive
func F_KeepFileRestoredFromArchive(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogArchiveForceDone github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogArchiveForceDone
func F_XLogArchiveForceDone(m *base.Module, l0 int32)
//go:linkname F_XLogArchiveNotify github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogArchiveNotify
func F_XLogArchiveNotify(m *base.Module, l0 int32)
//go:linkname F_XLogBeginInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_XLogBeginInsert
func F_XLogBeginInsert(m *base.Module)
//go:linkname F_XLogRegisterBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogRegisterBuffer
func F_XLogRegisterBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogRegisterData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XLogRegisterData
func F_XLogRegisterData(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogRegisterBufData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogRegisterBufData
func F_XLogRegisterBufData(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogInsert
func F_XLogInsert(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_log_newpage_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_log_newpage_buffer
func F_log_newpage_buffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_log_newpage_range github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_log_newpage_range
func F_log_newpage_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_XLogReadAhead github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogReadAhead
func F_XLogReadAhead(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_report_invalid_record github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_invalid_record
func F_report_invalid_record(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogRecGetBlockTag github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XLogRecGetBlockTag
func F_XLogRecGetBlockTag(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_RestoreBlockImage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RestoreBlockImage
func F_RestoreBlockImage(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetLatestXTime github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetLatestXTime
func F_GetLatestXTime(m *base.Module) int64
//go:linkname F_WakeupRecovery github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WakeupRecovery
func F_WakeupRecovery(m *base.Module)
//go:linkname F_GetXLogReplayRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetXLogReplayRecPtr
func F_GetXLogReplayRecPtr(m *base.Module, l0 int32) int64
//go:linkname F_RecoveryRequiresIntParameter github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RecoveryRequiresIntParameter
func F_RecoveryRequiresIntParameter(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogReadBufferForRedo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogReadBufferForRedo
func F_XLogReadBufferForRedo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_XLogReadBufferExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogReadBufferExtended
func F_XLogReadBufferExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_XLogInitBufferForRedo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogInitBufferForRedo
func F_XLogInitBufferForRedo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_XLogDropRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogDropRelation
func F_XLogDropRelation(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AddFileToBackupManifest github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AddFileToBackupManifest
func F_AddFileToBackupManifest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32)
//go:linkname F_sendDir github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sendDir
func F_sendDir(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int64
//go:linkname F__tarWriteHeader github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__tarWriteHeader
func F__tarWriteHeader(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_push_to_sink github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_push_to_sink
func F_push_to_sink(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_read_file_data_into_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_read_file_data_into_buffer
func F_read_file_data_into_buffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32) int64
//go:linkname F_SendXlogRecPtrResult github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SendXlogRecPtrResult
func F_SendXlogRecPtrResult(m *base.Module, l0 int64, l1 int32)
//go:linkname F_bbsink_forward_begin_archive github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bbsink_forward_begin_archive
func F_bbsink_forward_begin_archive(m *base.Module, l0 int32, l1 int32)
//go:linkname F_bbsink_forward_end_archive github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bbsink_forward_end_archive
func F_bbsink_forward_end_archive(m *base.Module, l0 int32)
//go:linkname F_boot_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_boot_yyensure_buffer_stack
func F_boot_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_boot_yy_create_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_boot_yy_create_buffer
func F_boot_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_yy_fatal_error_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_yy_fatal_error_1
func F_yy_fatal_error_1(m *base.Module, l0 int32)
//go:linkname F_boot_yylex_init github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_boot_yylex_init
func F_boot_yylex_init(m *base.Module, l0 int32) int32
//go:linkname F_boot_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_boot_yyerror
func F_boot_yyerror(m *base.Module, l0 int32, l1 int32)
//go:linkname F_closerel github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_closerel
func F_closerel(m *base.Module, l0 int32)
//go:linkname F_populate_typ_list github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_populate_typ_list
func F_populate_typ_list(m *base.Module)
//go:linkname F_boot_get_type_io_data github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_boot_get_type_io_data
func F_boot_get_type_io_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_heap_getattr_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_getattr_2
func F_heap_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_aclcheck_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_aclcheck_error
func F_aclcheck_error(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_aclcheck_error_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_aclcheck_error_type
func F_aclcheck_error_type(m *base.Module, l0 int32, l1 int32)
//go:linkname F_object_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_object_aclcheck
func F_object_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_object_aclcheck_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_object_aclcheck_ext
func F_object_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_pg_attribute_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_attribute_aclcheck
func F_pg_attribute_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_pg_attribute_aclcheck_all github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_attribute_aclcheck_all
func F_pg_attribute_aclcheck_all(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pg_attribute_aclcheck_all_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_attribute_aclcheck_all_ext
func F_pg_attribute_aclcheck_all_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32
//go:linkname F_pg_class_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_class_aclcheck
func F_pg_class_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pg_class_aclcheck_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_class_aclcheck_ext
func F_pg_class_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pg_parameter_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_parameter_aclcheck
func F_pg_parameter_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_object_ownercheck github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_object_ownercheck
func F_object_ownercheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IsCatalogRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IsCatalogRelation
func F_IsCatalogRelation(m *base.Module, l0 int32) int32
//go:linkname F_GetNewOidWithIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetNewOidWithIndex
func F_GetNewOidWithIndex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_performDeletion github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_performDeletion
func F_performDeletion(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_new_object_addresses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_new_object_addresses
func F_new_object_addresses(m *base.Module) int32
//go:linkname F_performMultipleDeletions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_performMultipleDeletions
func F_performMultipleDeletions(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_add_exact_object_address github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_exact_object_address
func F_add_exact_object_address(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SystemAttributeDefinition github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SystemAttributeDefinition
func F_SystemAttributeDefinition(m *base.Module, l0 int32) int32
//go:linkname F_heap_create github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_create
func F_heap_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32) int32
//go:linkname F_CheckAttributeNamesTypes github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CheckAttributeNamesTypes
func F_CheckAttributeNamesTypes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_heap_create_with_catalog github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_create_with_catalog
func F_heap_create_with_catalog(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32) int32
//go:linkname F_RelationClearMissing github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationClearMissing
func F_RelationClearMissing(m *base.Module, l0 int32)
//go:linkname F_heap_truncate_check_FKs github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_truncate_check_FKs
func F_heap_truncate_check_FKs(m *base.Module, l0 int32, l1 int32)
//go:linkname F_heap_truncate_find_FKs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_heap_truncate_find_FKs
func F_heap_truncate_find_FKs(m *base.Module, l0 int32) int32
//go:linkname F_index_build github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_build
func F_index_build(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_BuildIndexInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BuildIndexInfo
func F_BuildIndexInfo(m *base.Module, l0 int32) int32
//go:linkname F_reindex_index github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_reindex_index
func F_reindex_index(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_CatalogOpenIndexes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CatalogOpenIndexes
func F_CatalogOpenIndexes(m *base.Module, l0 int32) int32
//go:linkname F_CatalogCloseIndexes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CatalogCloseIndexes
func F_CatalogCloseIndexes(m *base.Module, l0 int32)
//go:linkname F_CatalogIndexInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CatalogIndexInsert
func F_CatalogIndexInsert(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CatalogTupleInsertWithInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CatalogTupleInsertWithInfo
func F_CatalogTupleInsertWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CatalogTupleUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CatalogTupleUpdate
func F_CatalogTupleUpdate(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CatalogTupleUpdateWithInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CatalogTupleUpdateWithInfo
func F_CatalogTupleUpdateWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RangeVarGetRelidExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RangeVarGetRelidExtended
func F_RangeVarGetRelidExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_LookupExplicitNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LookupExplicitNamespace
func F_LookupExplicitNamespace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_recomputeNamespacePath github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recomputeNamespacePath
func F_recomputeNamespacePath(m *base.Module)
//go:linkname F_AccessTempTableNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AccessTempTableNamespace
func F_AccessTempTableNamespace(m *base.Module, l0 int32)
//go:linkname F_get_namespace_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_namespace_oid
func F_get_namespace_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationIsVisible github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationIsVisible
func F_RelationIsVisible(m *base.Module, l0 int32) int32
//go:linkname F_TypeIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TypeIsVisibleExt
func F_TypeIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FuncnameGetCandidates github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FuncnameGetCandidates
func F_FuncnameGetCandidates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_DeconstructQualifiedName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DeconstructQualifiedName
func F_DeconstructQualifiedName(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_NameListToString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_NameListToString
func F_NameListToString(m *base.Module, l0 int32) int32
//go:linkname F_FunctionIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FunctionIsVisibleExt
func F_FunctionIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_OperatorIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OperatorIsVisibleExt
func F_OperatorIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_ts_config_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_ts_config_oid
func F_get_ts_config_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TSConfigIsVisible github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TSConfigIsVisible
func F_TSConfigIsVisible(m *base.Module, l0 int32) int32
//go:linkname F_CheckSetNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CheckSetNamespace
func F_CheckSetNamespace(m *base.Module, l0 int32, l1 int32)
//go:linkname F_QualifiedNameGetCreationNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_QualifiedNameGetCreationNamespace
func F_QualifiedNameGetCreationNamespace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeRangeVarFromNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeRangeVarFromNameList
func F_makeRangeVarFromNameList(m *base.Module, l0 int32) int32
//go:linkname F_isTempToastNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_isTempToastNamespace
func F_isTempToastNamespace(m *base.Module, l0 int32) int32
//go:linkname F_GetTempNamespaceProcNumber github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetTempNamespaceProcNumber
func F_GetTempNamespaceProcNumber(m *base.Module, l0 int32) int32
//go:linkname F_GetSearchPathMatcher github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetSearchPathMatcher
func F_GetSearchPathMatcher(m *base.Module, l0 int32) int32
//go:linkname F_get_collation_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_collation_oid
func F_get_collation_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FindDefaultConversionProc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FindDefaultConversionProc
func F_FindDefaultConversionProc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AtEOXact_Namespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtEOXact_Namespace
func F_AtEOXact_Namespace(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fetch_search_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fetch_search_path
func F_fetch_search_path(m *base.Module, l0 int32) int32
//go:linkname F_RunObjectPostCreateHook github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RunObjectPostCreateHook
func F_RunObjectPostCreateHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RunObjectPostAlterHook github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RunObjectPostAlterHook
func F_RunObjectPostAlterHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_RunFunctionExecuteHook github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RunFunctionExecuteHook
func F_RunFunctionExecuteHook(m *base.Module, l0 int32)
//go:linkname F_get_object_address github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_object_address
func F_get_object_address(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_getObjectDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getObjectDescription
func F_getObjectDescription(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_object_ownership github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_object_ownership
func F_check_object_ownership(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_get_object_class_descr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_object_class_descr
func F_get_object_class_descr(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_owner github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_object_attnum_owner
func F_get_object_attnum_owner(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_object_attnum_acl
func F_get_object_attnum_acl(m *base.Module, l0 int32) int32
//go:linkname F_get_object_type github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_object_type
func F_get_object_type(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_catalog_object_by_oid_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_catalog_object_by_oid_extended
func F_get_catalog_object_by_oid_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getObjectTypeDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getObjectTypeDescription
func F_getObjectTypeDescription(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getObjectIdentityParts github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getObjectIdentityParts
func F_getObjectIdentityParts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_partition_parent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_partition_parent
func F_get_partition_parent(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_partition_ancestors_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_partition_ancestors_worker
func F_get_partition_ancestors_worker(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_has_partition_attrs github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_has_partition_attrs
func F_has_partition_attrs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_errdetail_relkind_not_supported github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errdetail_relkind_not_supported
func F_errdetail_relkind_not_supported(m *base.Module, l0 int32)
//go:linkname F_CreateConstraintEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateConstraintEntry
func F_CreateConstraintEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32, l21 int32, l22 int32, l23 int32, l24 int32, l25 int32, l26 int32, l27 int32, l28 int32, l29 int32, l30 int32, l31 int32, l32 int32) int32
//go:linkname F_findNotNullConstraintAttnum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_findNotNullConstraintAttnum
func F_findNotNullConstraintAttnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_extractNotNullColumn github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_extractNotNullColumn
func F_extractNotNullColumn(m *base.Module, l0 int32) int32
//go:linkname F_AlterConstraintNamespaces github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AlterConstraintNamespaces
func F_AlterConstraintNamespaces(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_heap_getattr_4 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_getattr_4
func F_heap_getattr_4(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_recordDependencyOn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recordDependencyOn
func F_recordDependencyOn(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_getExtensionOfObject github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getExtensionOfObject
func F_getExtensionOfObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_deleteDependencyRecordsForClass github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_deleteDependencyRecordsForClass
func F_deleteDependencyRecordsForClass(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_changeDependencyFor github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_changeDependencyFor
func F_changeDependencyFor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_getOwnedSequences github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getOwnedSequences
func F_getOwnedSequences(m *base.Module, l0 int32) int32
//go:linkname F_find_all_inheritors github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_all_inheritors
func F_find_all_inheritors(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_StoreSingleInheritance github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_StoreSingleInheritance
func F_StoreSingleInheritance(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_DeleteInheritsTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DeleteInheritsTuple
func F_DeleteInheritsTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_makeOperatorDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeOperatorDependencies
func F_makeOperatorDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetPubPartitionOptionRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetPubPartitionOptionRelations
func F_GetPubPartitionOptionRelations(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetRelationPublications github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetRelationPublications
func F_GetRelationPublications(m *base.Module, l0 int32) int32
//go:linkname F_changeDependencyOnOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_changeDependencyOnOwner
func F_changeDependencyOnOwner(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_shdepChangeDep github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shdepChangeDep
func F_shdepChangeDep(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_updateAclDependenciesWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_updateAclDependenciesWorker
func F_updateAclDependenciesWorker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_GetPublicationsStr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetPublicationsStr
func F_GetPublicationsStr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetSubscription github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetSubscription
func F_GetSubscription(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationCreateStorage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationCreateStorage
func F_RelationCreateStorage(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RelationDropStorage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationDropStorage
func F_RelationDropStorage(m *base.Module, l0 int32)
//go:linkname F_RelationPreserveStorage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationPreserveStorage
func F_RelationPreserveStorage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationTruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationTruncate
func F_RelationTruncate(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationCopyStorage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationCopyStorage
func F_RelationCopyStorage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RelFileLocatorSkippingWAL github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelFileLocatorSkippingWAL
func F_RelFileLocatorSkippingWAL(m *base.Module, l0 int32) int32
//go:linkname F_smgrDoPendingDeletes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgrDoPendingDeletes
func F_smgrDoPendingDeletes(m *base.Module, l0 int32)
//go:linkname F_smgrDoPendingSyncs github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrDoPendingSyncs
func F_smgrDoPendingSyncs(m *base.Module, l0 int32, l1 int32)
//go:linkname F_create_toast_table github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_toast_table
func F_create_toast_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_NewRelationCreateToastTable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_NewRelationCreateToastTable
func F_NewRelationCreateToastTable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_parse_analyze_fixedparams github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_analyze_fixedparams
func F_parse_analyze_fixedparams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_parse_sub_analyze github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parse_sub_analyze
func F_parse_sub_analyze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_stmt_requires_parse_analysis github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_stmt_requires_parse_analysis
func F_stmt_requires_parse_analysis(m *base.Module, l0 int32) int32
//go:linkname F_query_requires_rewrite_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_query_requires_rewrite_plan
func F_query_requires_rewrite_plan(m *base.Module, l0 int32) int32
//go:linkname F_CheckSelectLocking github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckSelectLocking
func F_CheckSelectLocking(m *base.Module, l0 int32, l1 int32)
//go:linkname F_expand_groupingset_node github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_expand_groupingset_node
func F_expand_groupingset_node(m *base.Module, l0 int32) int32
//go:linkname F_markRelsAsNulledBy github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_markRelsAsNulledBy
func F_markRelsAsNulledBy(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_extractRemainingColumns github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_extractRemainingColumns
func F_extractRemainingColumns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_checkExprIsVarFree github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_checkExprIsVarFree
func F_checkExprIsVarFree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_transformGroupClause github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformGroupClause
func F_transformGroupClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_addTargetToSortList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_addTargetToSortList
func F_addTargetToSortList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_transformFrameOffset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformFrameOffset
func F_transformFrameOffset(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_can_coerce_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_can_coerce_type
func F_can_coerce_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_coerce_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_coerce_type
func F_coerce_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_coerce_type_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_coerce_type_typmod
func F_coerce_type_typmod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_find_coercion_pathway github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_coercion_pathway
func F_find_coercion_pathway(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_coerce_to_boolean github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_coerce_to_boolean
func F_coerce_to_boolean(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_coerce_to_specific_type_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_coerce_to_specific_type_typmod
func F_coerce_to_specific_type_typmod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_coerce_to_specific_type github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_coerce_to_specific_type
func F_coerce_to_specific_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_select_common_type github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_select_common_type
func F_select_common_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_coerce_to_common_type github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_coerce_to_common_type
func F_coerce_to_common_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_select_common_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_select_common_typmod
func F_select_common_typmod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_enforce_generic_type_consistency github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_enforce_generic_type_consistency
func F_enforce_generic_type_consistency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_TypeCategory github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TypeCategory
func F_TypeCategory(m *base.Module, l0 int32) int32
//go:linkname F_IsPreferredType github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_IsPreferredType
func F_IsPreferredType(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_IsBinaryCoercible github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IsBinaryCoercible
func F_IsBinaryCoercible(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assign_collations_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_assign_collations_walker
func F_assign_collations_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assign_list_collations github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_assign_list_collations
func F_assign_list_collations(m *base.Module, l0 int32, l1 int32)
//go:linkname F_assign_expr_collations github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_assign_expr_collations
func F_assign_expr_collations(m *base.Module, l0 int32, l1 int32)
//go:linkname F_analyzeCTETargetList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_analyzeCTETargetList
func F_analyzeCTETargetList(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_name_matches_visible_ENR github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_name_matches_visible_ENR
func F_name_matches_visible_ENR(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_transformExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformExpr
func F_transformExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_transformExprRecurse github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformExprRecurse
func F_transformExprRecurse(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_func_signature_string github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_func_signature_string
func F_func_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_check_srf_call_placement github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_srf_call_placement
func F_check_srf_call_placement(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_make_fn_arguments github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_make_fn_arguments
func F_make_fn_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_LookupFuncName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LookupFuncName
func F_LookupFuncName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CheckDuplicateColumnOrPathNames github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckDuplicateColumnOrPathNames
func F_CheckDuplicateColumnOrPathNames(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformJsonTableColumns github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformJsonTableColumns
func F_transformJsonTableColumns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_parsestate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_parsestate
func F_make_parsestate(m *base.Module, l0 int32) int32
//go:linkname F_free_parsestate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_free_parsestate
func F_free_parsestate(m *base.Module, l0 int32)
//go:linkname F_LookupOperName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LookupOperName
func F_LookupOperName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_op_signature_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_op_signature_string
func F_op_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_oper github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_oper
func F_oper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_left_oper github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_left_oper
func F_left_oper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_GetNSItemByRangeTablePosn github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetNSItemByRangeTablePosn
func F_GetNSItemByRangeTablePosn(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_scanNSItemForColumn github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_scanNSItemForColumn
func F_scanNSItemForColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_markRTEForSelectPriv github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_markRTEForSelectPriv
func F_markRTEForSelectPriv(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_markNullableIfNeeded github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_markNullableIfNeeded
func F_markNullableIfNeeded(m *base.Module, l0 int32, l1 int32)
//go:linkname F_markVarForSelectPriv github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_markVarForSelectPriv
func F_markVarForSelectPriv(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getRTEPermissionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getRTEPermissionInfo
func F_getRTEPermissionInfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parserOpenTable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parserOpenTable
func F_parserOpenTable(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_buildRelationAliases github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_buildRelationAliases
func F_buildRelationAliases(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addRangeTableEntryForRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addRangeTableEntryForRelation
func F_addRangeTableEntryForRelation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_addRangeTableEntryForSubquery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_addRangeTableEntryForSubquery
func F_addRangeTableEntryForSubquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_buildNSItemFromLists github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_buildNSItemFromLists
func F_buildNSItemFromLists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_addRangeTableEntryForTableFunc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addRangeTableEntryForTableFunc
func F_addRangeTableEntryForTableFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_addRangeTableEntryForJoin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addRangeTableEntryForJoin
func F_addRangeTableEntryForJoin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32
//go:linkname F_addNSItemToQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_addNSItemToQuery
func F_addNSItemToQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_expandRTE github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expandRTE
func F_expandRTE(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_expandNSItemVars github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_expandNSItemVars
func F_expandNSItemVars(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_get_rte_attribute_name github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_rte_attribute_name
func F_get_rte_attribute_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_attnumAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_attnumAttName
func F_attnumAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_attnumTypeId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_attnumTypeId
func F_attnumTypeId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_attnumCollationId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_attnumCollationId
func F_attnumCollationId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_isQueryUsingTempRelation_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isQueryUsingTempRelation_walker
func F_isQueryUsingTempRelation_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FigureColname github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FigureColname
func F_FigureColname(m *base.Module, l0 int32) int32
//go:linkname F_expandRecordVariable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expandRecordVariable
func F_expandRecordVariable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TypeNameToString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TypeNameToString
func F_TypeNameToString(m *base.Module, l0 int32) int32
//go:linkname F_typenameTypeIdAndMod github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_typenameTypeIdAndMod
func F_typenameTypeIdAndMod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetColumnDefCollation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetColumnDefCollation
func F_GetColumnDefCollation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_typeidType github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_typeidType
func F_typeidType(m *base.Module, l0 int32) int32
//go:linkname F_parseTypeString github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_parseTypeString
func F_parseTypeString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_transformIndexStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformIndexStmt
func F_transformIndexStmt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_transformStatsStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformStatsStmt
func F_transformStatsStmt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_transformAlterTableStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformAlterTableStmt
func F_transformAlterTableStmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_raw_parser github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_raw_parser
func F_raw_parser(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_yy_fatal_error_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_yy_fatal_error_2
func F_yy_fatal_error_2(m *base.Module, l0 int32)
//go:linkname F_scanner_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_scanner_finish
func F_scanner_finish(m *base.Module, l0 int32)
//go:linkname F_downcase_truncate_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_downcase_truncate_identifier
func F_downcase_truncate_identifier(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_table_am_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_table_am_oid
func F_get_table_am_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_am_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_am_name
func F_get_am_name(m *base.Module, l0 int32) int32
//go:linkname F_std_typanalyze github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_std_typanalyze
func F_std_typanalyze(m *base.Module, l0 int32) int32
//go:linkname F_asyncQueueUnregister github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_asyncQueueUnregister
func F_asyncQueueUnregister(m *base.Module)
//go:linkname F_swap_relation_files github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_swap_relation_files
func F_swap_relation_files(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_CreateComments github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateComments
func F_CreateComments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetComment github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetComment
func F_GetComment(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CopyGetData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopyGetData
func F_CopyGetData(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_database_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_database_oid
func F_get_database_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CreateDirAndVersionFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateDirAndVersionFile
func F_CreateDirAndVersionFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_database_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_database_name
func F_get_database_name(m *base.Module, l0 int32) int32
//go:linkname F_heap_getattr_6 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_getattr_6
func F_heap_getattr_6(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_recovery_create_dbdir github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_recovery_create_dbdir
func F_recovery_create_dbdir(m *base.Module, l0 int32, l1 int32)
//go:linkname F_defGetBoolean github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_defGetBoolean
func F_defGetBoolean(m *base.Module, l0 int32) int32
//go:linkname F_defGetInt64 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_defGetInt64
func F_defGetInt64(m *base.Module, l0 int32) int64
//go:linkname F_errorConflictingDefElem github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errorConflictingDefElem
func F_errorConflictingDefElem(m *base.Module, l0 int32, l1 int32)
//go:linkname F_EventTriggerCommonSetup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EventTriggerCommonSetup
func F_EventTriggerCommonSetup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_EventTriggerInvoke github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EventTriggerInvoke
func F_EventTriggerInvoke(m *base.Module, l0 int32, l1 int32)
//go:linkname F_show_buffer_usage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_show_buffer_usage
func F_show_buffer_usage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExplainPrintJIT github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExplainPrintJIT
func F_ExplainPrintJIT(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainPreScanNode github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExplainPreScanNode
func F_ExplainPreScanNode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExplainNode github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExplainNode
func F_ExplainNode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_report_triggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_report_triggers
func F_report_triggers(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainPropertyList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExplainPropertyList
func F_ExplainPropertyList(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainPropertyText github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExplainPropertyText
func F_ExplainPropertyText(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainPropertyInteger github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainPropertyInteger
func F_ExplainPropertyInteger(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_ExplainPropertyUInteger github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExplainPropertyUInteger
func F_ExplainPropertyUInteger(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_ExplainPropertyFloat github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExplainPropertyFloat
func F_ExplainPropertyFloat(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32)
//go:linkname F_ExplainOpenGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainOpenGroup
func F_ExplainOpenGroup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExplainCloseGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainCloseGroup
func F_ExplainCloseGroup(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_check_valid_extension_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_valid_extension_name
func F_check_valid_extension_name(m *base.Module, l0 int32)
//go:linkname F_get_ext_ver_list github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_ext_ver_list
func F_get_ext_ver_list(m *base.Module, l0 int32) int32
//go:linkname F_find_update_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_find_update_path
func F_find_update_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_heap_getattr_7 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_heap_getattr_7
func F_heap_getattr_7(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CheckIndexCompatible github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckIndexCompatible
func F_CheckIndexCompatible(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_GetOperatorFromCompareType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetOperatorFromCompareType
func F_GetOperatorFromCompareType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_DefineIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_DefineIndex
func F_DefineIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)
//go:linkname F_makeObjectName github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeObjectName
func F_makeObjectName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_simple_rowfilter_expr_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_simple_rowfilter_expr_walker
func F_check_simple_rowfilter_expr_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fill_seq_with_data github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fill_seq_with_data
func F_fill_seq_with_data(m *base.Module, l0 int32, l1 int32)
//go:linkname F_init_sequence github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_init_sequence
func F_init_sequence(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_read_seq_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_seq_tuple
func F_read_seq_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_do_setval github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_do_setval
func F_do_setval(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_DefineRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DefineRelation
func F_DefineRelation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_CheckTableNotInUse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CheckTableNotInUse
func F_CheckTableNotInUse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SetRelationHasSubclass github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SetRelationHasSubclass
func F_SetRelationHasSubclass(m *base.Module, l0 int32, l1 int32)
//go:linkname F_truncate_check_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_truncate_check_rel
func F_truncate_check_rel(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RenameRelationInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RenameRelationInternal
func F_RenameRelationInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_find_composite_type_dependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_find_composite_type_dependencies
func F_find_composite_type_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AtEOXact_on_commit_actions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtEOXact_on_commit_actions
func F_AtEOXact_on_commit_actions(m *base.Module, l0 int32)
//go:linkname F_AtEOSubXact_on_commit_actions github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AtEOSubXact_on_commit_actions
func F_AtEOSubXact_on_commit_actions(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ATSimplePermissions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ATSimplePermissions
func F_ATSimplePermissions(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ATPrepAddColumn github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ATPrepAddColumn
func F_ATPrepAddColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_ATSimpleRecursion github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ATSimpleRecursion
func F_ATSimpleRecursion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ATParseTransformCmd github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ATParseTransformCmd
func F_ATParseTransformCmd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_drop_parent_dependency github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_drop_parent_dependency
func F_drop_parent_dependency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ATTypedTableRecursion github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ATTypedTableRecursion
func F_ATTypedTableRecursion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_RebuildConstraintComment github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RebuildConstraintComment
func F_RebuildConstraintComment(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_tablespace_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_tablespace_oid
func F_get_tablespace_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_tablespace_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_tablespace_name
func F_get_tablespace_name(m *base.Module, l0 int32) int32
//go:linkname F_TriggerEnabled github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TriggerEnabled
func F_TriggerEnabled(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ExecCallTriggerFunc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecCallTriggerFunc
func F_ExecCallTriggerFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_AfterTriggerSaveEvent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AfterTriggerSaveEvent
func F_AfterTriggerSaveEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)
//go:linkname F_ExecBRInsertTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecBRInsertTriggers
func F_ExecBRInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecARInsertTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecARInsertTriggers
func F_ExecARInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ExecBRDeleteTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecBRDeleteTriggers
func F_ExecBRDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_ExecARDeleteTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecARDeleteTriggers
func F_ExecARDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ExecIRDeleteTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecIRDeleteTriggers
func F_ExecIRDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecBRUpdateTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecBRUpdateTriggers
func F_ExecBRUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_ExecARUpdateTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecARUpdateTriggers
func F_ExecARUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_ExecBSTruncateTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecBSTruncateTriggers
func F_ExecBSTruncateTriggers(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AfterTriggerEndQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AfterTriggerEndQuery
func F_AfterTriggerEndQuery(m *base.Module, l0 int32)
//go:linkname F_AfterTriggerEndXact github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AfterTriggerEndXact
func F_AfterTriggerEndXact(m *base.Module)
//go:linkname F_AfterTriggerEndSubXact github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AfterTriggerEndSubXact
func F_AfterTriggerEndSubXact(m *base.Module, l0 int32)
//go:linkname F_deserialize_deflist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_deserialize_deflist
func F_deserialize_deflist(m *base.Module, l0 int32) int32
//go:linkname F_vacuum_delay_point github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vacuum_delay_point
func F_vacuum_delay_point(m *base.Module, l0 int32)
//go:linkname F_vac_bulkdel_one_index github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vac_bulkdel_one_index
func F_vac_bulkdel_one_index(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_parallel_vacuum_reset_dead_items github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parallel_vacuum_reset_dead_items
func F_parallel_vacuum_reset_dead_items(m *base.Module, l0 int32)
//go:linkname F_parallel_vacuum_process_one_index github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parallel_vacuum_process_one_index
func F_parallel_vacuum_process_one_index(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecReScan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecReScan
func F_ExecReScan(m *base.Module, l0 int32)
//go:linkname F_expr_setup_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_expr_setup_walker
func F_expr_setup_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecPushExprSetupSteps github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecPushExprSetupSteps
func F_ExecPushExprSetupSteps(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExprEvalPushStep github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExprEvalPushStep
func F_ExprEvalPushStep(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecInitFunc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecInitFunc
func F_ExecInitFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ExecInitSubPlanExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecInitSubPlanExpr
func F_ExecInitSubPlanExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecCheck
func F_ExecCheck(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecBuildAggTrans github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBuildAggTrans
func F_ExecBuildAggTrans(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ExecBuildGroupingEqual github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBuildGroupingEqual
func F_ExecBuildGroupingEqual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_ExecReadyInterpretedExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecReadyInterpretedExpr
func F_ExecReadyInterpretedExpr(m *base.Module, l0 int32)
//go:linkname F_tuplehash_insert_hash_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplehash_insert_hash_internal
func F_tuplehash_insert_hash_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplehash_lookup_hash_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplehash_lookup_hash_internal
func F_tuplehash_lookup_hash_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecOpenIndices github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecOpenIndices
func F_ExecOpenIndices(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecCloseIndices github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecCloseIndices
func F_ExecCloseIndices(m *base.Module, l0 int32)
//go:linkname F_ExecInsertIndexTuples github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecInsertIndexTuples
func F_ExecInsertIndexTuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_ExecInitJunkFilter github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecInitJunkFilter
func F_ExecInitJunkFilter(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecutorStart github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecutorStart
func F_ExecutorStart(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecutorRun github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecutorRun
func F_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_ExecutorFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecutorFinish
func F_ExecutorFinish(m *base.Module, l0 int32)
//go:linkname F_ExecCloseResultRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecCloseResultRelations
func F_ExecCloseResultRelations(m *base.Module, l0 int32)
//go:linkname F_ExecPartitionCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecPartitionCheck
func F_ExecPartitionCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecPartitionCheckEmitError github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecPartitionCheckEmitError
func F_ExecPartitionCheckEmitError(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecBuildSlotValueDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBuildSlotValueDescription
func F_ExecBuildSlotValueDescription(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecConstraints
func F_ExecConstraints(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecWithCheckOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecWithCheckOptions
func F_ExecWithCheckOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_EvalPlanQualBegin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_EvalPlanQualBegin
func F_EvalPlanQualBegin(m *base.Module, l0 int32)
//go:linkname F_EvalPlanQualSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EvalPlanQualSlot
func F_EvalPlanQualSlot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EvalPlanQualEnd github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EvalPlanQualEnd
func F_EvalPlanQualEnd(m *base.Module, l0 int32)
//go:linkname F_ExecParallelFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecParallelFinish
func F_ExecParallelFinish(m *base.Module, l0 int32)
//go:linkname F_ExecParallelCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecParallelCleanup
func F_ExecParallelCleanup(m *base.Module, l0 int32)
//go:linkname F_ExecInitPartitionDispatchInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecInitPartitionDispatchInfo
func F_ExecInitPartitionDispatchInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ExecCleanupTupleRouting github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecCleanupTupleRouting
func F_ExecCleanupTupleRouting(m *base.Module, l0 int32, l1 int32)
//go:linkname F_MultiExecProcNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MultiExecProcNode
func F_MultiExecProcNode(m *base.Module, l0 int32) int32
//go:linkname F_ExecSetTupleBound github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecSetTupleBound
func F_ExecSetTupleBound(m *base.Module, l0 int64, l1 int32)
//go:linkname F_CheckAndReportConflict github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CheckAndReportConflict
func F_CheckAndReportConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_CheckCmdReplicaIdentity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CheckCmdReplicaIdentity
func F_CheckCmdReplicaIdentity(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecScan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecScan
func F_ExecScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecStoreMinimalTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecStoreMinimalTuple
func F_ExecStoreMinimalTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MakeSingleTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MakeSingleTupleTableSlot
func F_MakeSingleTupleTableSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecDropSingleTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecDropSingleTupleTableSlot
func F_ExecDropSingleTupleTableSlot(m *base.Module, l0 int32)
//go:linkname F_ExecForceStoreHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecForceStoreHeapTuple
func F_ExecForceStoreHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecForceStoreMinimalTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecForceStoreMinimalTuple
func F_ExecForceStoreMinimalTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecStoreAllNullTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecStoreAllNullTuple
func F_ExecStoreAllNullTuple(m *base.Module, l0 int32)
//go:linkname F_ExecFetchSlotHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecFetchSlotHeapTuple
func F_ExecFetchSlotHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecFetchSlotMinimalTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecFetchSlotMinimalTuple
func F_ExecFetchSlotMinimalTuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecTypeFromTLInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecTypeFromTLInternal
func F_ExecTypeFromTLInternal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecInitExtraTupleSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecInitExtraTupleSlot
func F_ExecInitExtraTupleSlot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_slot_getsomeattrs_int github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_slot_getsomeattrs_int
func F_slot_getsomeattrs_int(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecCleanTypeFromTL github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecCleanTypeFromTL
func F_ExecCleanTypeFromTL(m *base.Module, l0 int32) int32
//go:linkname F_ExecTypeFromExprList github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecTypeFromExprList
func F_ExecTypeFromExprList(m *base.Module, l0 int32) int32
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
//go:linkname F_CreateStandaloneExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateStandaloneExprContext
func F_CreateStandaloneExprContext(m *base.Module) int32
//go:linkname F_ReScanExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReScanExprContext
func F_ReScanExprContext(m *base.Module, l0 int32)
//go:linkname F_MakePerTupleExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MakePerTupleExprContext
func F_MakePerTupleExprContext(m *base.Module, l0 int32) int32
//go:linkname F_ExecGetResultSlotOps github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetResultSlotOps
func F_ExecGetResultSlotOps(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecConditionalAssignProjectionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecConditionalAssignProjectionInfo
func F_ExecConditionalAssignProjectionInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_executor_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_executor_errposition
func F_executor_errposition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecGetTriggerOldSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetTriggerOldSlot
func F_ExecGetTriggerOldSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetReturningSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecGetReturningSlot
func F_ExecGetReturningSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetAllNullSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecGetAllNullSlot
func F_ExecGetAllNullSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetUpdatedCols github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecGetUpdatedCols
func F_ExecGetUpdatedCols(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetAllUpdatedCols github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetAllUpdatedCols
func F_ExecGetAllUpdatedCols(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_InstrAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_InstrAlloc
func F_InstrAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_InstrStartNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_InstrStartNode
func F_InstrStartNode(m *base.Module, l0 int32)
//go:linkname F_InstrStopNode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InstrStopNode
func F_InstrStopNode(m *base.Module, l0 int32, l1 float64)
//go:linkname F_BufferUsageAccumDiff github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BufferUsageAccumDiff
func F_BufferUsageAccumDiff(m *base.Module, l0 int32, l1 int32)
//go:linkname F_InstrEndLoop github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_InstrEndLoop
func F_InstrEndLoop(m *base.Module, l0 int32)
//go:linkname F_initialize_phase github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_initialize_phase
func F_initialize_phase(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fetch_input_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fetch_input_tuple
func F_fetch_input_tuple(m *base.Module, l0 int32) int32
//go:linkname F_lookup_hash_entries github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lookup_hash_entries
func F_lookup_hash_entries(m *base.Module, l0 int32)
//go:linkname F_hashagg_spill_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hashagg_spill_finish
func F_hashagg_spill_finish(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_initialize_aggregate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_initialize_aggregate
func F_initialize_aggregate(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_finalize_aggregates github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_finalize_aggregates
func F_finalize_aggregates(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_hashagg_reset_spill_state github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hashagg_reset_spill_state
func F_hashagg_reset_spill_state(m *base.Module, l0 int32)
//go:linkname F_hashagg_spill_init github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hashagg_spill_init
func F_hashagg_spill_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 float64)
//go:linkname F_hashagg_spill_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hashagg_spill_tuple
func F_hashagg_spill_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecAsyncAppendResponse github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecAsyncAppendResponse
func F_ExecAsyncAppendResponse(m *base.Module, l0 int32)
//go:linkname F_load_tuple_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_load_tuple_array
func F_load_tuple_array(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecEndGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecEndGroup
func F_ExecEndGroup(m *base.Module, l0 int32)
//go:linkname F_ExecHashTableCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecHashTableCreate
func F_ExecHashTableCreate(m *base.Module, l0 int32) int32
//go:linkname F_get_hash_memory_limit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_hash_memory_limit
func F_get_hash_memory_limit(m *base.Module) int32
//go:linkname F_ExecHashTableDestroy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecHashTableDestroy
func F_ExecHashTableDestroy(m *base.Module, l0 int32)
//go:linkname F_ExecParallelHashTupleAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecParallelHashTupleAlloc
func F_ExecParallelHashTupleAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecParallelHashTableSetCurrentBatch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecParallelHashTableSetCurrentBatch
func F_ExecParallelHashTableSetCurrentBatch(m *base.Module, l0 int32, l1 int32)
//go:linkname F_isCurrentGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_isCurrentGroup
func F_isCurrentGroup(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_recompute_limits github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recompute_limits
func F_recompute_limits(m *base.Module, l0 int32)
//go:linkname F_ExecInitGenerated github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecInitGenerated
func F_ExecInitGenerated(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecProcessReturning github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecProcessReturning
func F_ExecProcessReturning(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ExecInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecInsert
func F_ExecInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ExecInitUpdateProjection github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecInitUpdateProjection
func F_ExecInitUpdateProjection(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecPendingInserts github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecPendingInserts
func F_ExecPendingInserts(m *base.Module, l0 int32)
//go:linkname F_ExecUpdateEpilogue github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecUpdateEpilogue
func F_ExecUpdateEpilogue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ExecProjectSRF github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecProjectSRF
func F_ExecProjectSRF(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecEndSeqScan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecEndSeqScan
func F_ExecEndSeqScan(m *base.Module, l0 int32)
//go:linkname F_setop_compare_slots github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_setop_compare_slots
func F_setop_compare_slots(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_window_gettupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_window_gettupleslot
func F_window_gettupleslot(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_WinSetMarkPosition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WinSetMarkPosition
func F_WinSetMarkPosition(m *base.Module, l0 int32, l1 int64)
//go:linkname F_WinGetPartitionLocalMemory github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WinGetPartitionLocalMemory
func F_WinGetPartitionLocalMemory(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_WinGetPartitionRowCount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WinGetPartitionRowCount
func F_WinGetPartitionRowCount(m *base.Module, l0 int32) int64
//go:linkname F_WinGetFuncArgInPartition github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WinGetFuncArgInPartition
func F_WinGetFuncArgInPartition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_WinGetFuncArgCurrent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WinGetFuncArgCurrent
func F_WinGetFuncArgCurrent(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_connect_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SPI_connect_ext
func F_SPI_connect_ext(m *base.Module, l0 int32)
//go:linkname F_SPI_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_finish
func F_SPI_finish(m *base.Module) int32
//go:linkname F_AtEOXact_SPI github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AtEOXact_SPI
func F_AtEOXact_SPI(m *base.Module, l0 int32)
//go:linkname F__SPI_execute_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__SPI_execute_plan
func F__SPI_execute_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SPI_execute_plan_with_paramlist github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SPI_execute_plan_with_paramlist
func F_SPI_execute_plan_with_paramlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__SPI_prepare_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__SPI_prepare_plan
func F__SPI_prepare_plan(m *base.Module, l0 int32, l1 int32)
//go:linkname F__SPI_make_plan_non_temp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__SPI_make_plan_non_temp
func F__SPI_make_plan_non_temp(m *base.Module, l0 int32) int32
//go:linkname F_SPI_cursor_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SPI_cursor_fetch
func F_SPI_cursor_fetch(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SPI_result_code_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_result_code_string
func F_SPI_result_code_string(m *base.Module, l0 int32) int32
//go:linkname F_CreateTupleQueueDestReceiver github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateTupleQueueDestReceiver
func F_CreateTupleQueueDestReceiver(m *base.Module, l0 int32) int32
//go:linkname F_TupleQueueReaderNext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TupleQueueReaderNext
func F_TupleQueueReaderNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetForeignDataWrapper github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetForeignDataWrapper
func F_GetForeignDataWrapper(m *base.Module, l0 int32) int32
//go:linkname F_get_foreign_server_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_foreign_server_oid
func F_get_foreign_server_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetForeignServerIdByRelId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetForeignServerIdByRelId
func F_GetForeignServerIdByRelId(m *base.Module, l0 int32) int32
//go:linkname F_GetFdwRoutineByServerId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetFdwRoutineByServerId
func F_GetFdwRoutineByServerId(m *base.Module, l0 int32) int32
//go:linkname F_GetFdwRoutineByRelId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetFdwRoutineByRelId
func F_GetFdwRoutineByRelId(m *base.Module, l0 int32) int32
//go:linkname F_dshash_create github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dshash_create
func F_dshash_create(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dshash_seq_next github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dshash_seq_next
func F_dshash_seq_next(m *base.Module, l0 int32) int32
//go:linkname F_dshash_seq_term github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_dshash_seq_term
func F_dshash_seq_term(m *base.Module, l0 int32)
//go:linkname F_initHyperLogLog github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_initHyperLogLog
func F_initHyperLogLog(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pairingheap_remove github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pairingheap_remove
func F_pairingheap_remove(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AtEOXact_LargeObject github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AtEOXact_LargeObject
func F_AtEOXact_LargeObject(m *base.Module, l0 int32)
//go:linkname F_AtEOSubXact_LargeObject github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AtEOSubXact_LargeObject
func F_AtEOSubXact_LargeObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_open_auth_file github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_open_auth_file
func F_open_auth_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_tokenize_auth_file github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tokenize_auth_file
func F_tokenize_auth_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_regcomp_auth_token github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_regcomp_auth_token
func F_regcomp_auth_token(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pq_getkeepalivesidle github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_getkeepalivesidle
func F_pq_getkeepalivesidle(m *base.Module, l0 int32) int32
//go:linkname F_pq_getkeepalivescount github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_getkeepalivescount
func F_pq_getkeepalivescount(m *base.Module, l0 int32) int32
//go:linkname F_pq_gettcpusertimeout github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_gettcpusertimeout
func F_pq_gettcpusertimeout(m *base.Module, l0 int32) int32
//go:linkname F_pq_beginmessage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_beginmessage
func F_pq_beginmessage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pq_sendbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_sendbytes
func F_pq_sendbytes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_sendcountedtext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_sendcountedtext
func F_pq_sendcountedtext(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_sendtext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_sendtext
func F_pq_sendtext(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_sendfloat8 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_sendfloat8
func F_pq_sendfloat8(m *base.Module, l0 int32, l1 float64)
//go:linkname F_pq_endmessage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_endmessage
func F_pq_endmessage(m *base.Module, l0 int32)
//go:linkname F_pq_endmessage_reuse github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_endmessage_reuse
func F_pq_endmessage_reuse(m *base.Module, l0 int32)
//go:linkname F_pq_begintypsend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_begintypsend
func F_pq_begintypsend(m *base.Module, l0 int32)
//go:linkname F_pq_putemptymessage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_putemptymessage
func F_pq_putemptymessage(m *base.Module, l0 int32)
//go:linkname F_pq_getmsgbyte github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_getmsgbyte
func F_pq_getmsgbyte(m *base.Module, l0 int32) int32
//go:linkname F_pq_getmsgint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pq_getmsgint
func F_pq_getmsgint(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_getmsgint64 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pq_getmsgint64
func F_pq_getmsgint64(m *base.Module, l0 int32) int64
//go:linkname F_pq_getmsgrawstring github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_getmsgrawstring
func F_pq_getmsgrawstring(m *base.Module, l0 int32) int32
//go:linkname F_pq_getmsgend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_getmsgend
func F_pq_getmsgend(m *base.Module, l0 int32)
//go:linkname F_bms_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_copy
func F_bms_copy(m *base.Module, l0 int32) int32
//go:linkname F_bms_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_bms_equal
func F_bms_equal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_make_singleton github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bms_make_singleton
func F_bms_make_singleton(m *base.Module, l0 int32) int32
//go:linkname F_bms_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_free
func F_bms_free(m *base.Module, l0 int32)
//go:linkname F_bms_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_union
func F_bms_union(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_difference github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_difference
func F_bms_difference(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_is_subset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_is_subset
func F_bms_is_subset(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_is_member github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_is_member
func F_bms_is_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_singleton_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_singleton_member
func F_bms_singleton_member(m *base.Module, l0 int32) int32
//go:linkname F_bms_add_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_add_member
func F_bms_add_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_del_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_del_member
func F_bms_del_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_add_range github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_add_range
func F_bms_add_range(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_bms_int_members github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_int_members
func F_bms_int_members(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_del_members github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_del_members
func F_bms_del_members(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_join github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_join
func F_bms_join(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_copyObjectImpl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_copyObjectImpl
func F_copyObjectImpl(m *base.Module, l0 int32) int32
//go:linkname F_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_equal
func F_equal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetExtensibleNodeMethods github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetExtensibleNodeMethods
func F_GetExtensibleNodeMethods(m *base.Module, l0 int32) int32
//go:linkname F_list_make1_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_make1_impl
func F_list_make1_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_make2_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_list_make2_impl
func F_list_make2_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lappend
func F_lappend(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_lappend_oid
func F_lappend_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend_xid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lappend_xid
func F_lappend_xid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_insert_nth github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_insert_nth
func F_list_insert_nth(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lcons github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lcons
func F_lcons(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lcons_int github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lcons_int
func F_lcons_int(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lcons_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lcons_oid
func F_lcons_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_concat
func F_list_concat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_list_copy
func F_list_copy(m *base.Module, l0 int32) int32
//go:linkname F_list_concat_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_concat_copy
func F_list_concat_copy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_member github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_member
func F_list_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_free
func F_list_free(m *base.Module, l0 int32)
//go:linkname F_list_delete_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_delete_ptr
func F_list_delete_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_first github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_delete_first
func F_list_delete_first(m *base.Module, l0 int32) int32
//go:linkname F_list_delete_last github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_delete_last
func F_list_delete_last(m *base.Module, l0 int32) int32
//go:linkname F_list_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_list_union
func F_list_union(m *base.Module, l0 int32) int32
//go:linkname F_list_difference_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_difference_ptr
func F_list_difference_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_append_unique github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_append_unique
func F_list_append_unique(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_concat_unique github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_concat_unique
func F_list_concat_unique(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_sort github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_sort
func F_list_sort(m *base.Module, l0 int32, l1 int32)
//go:linkname F_makeSimpleA_Expr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeSimpleA_Expr
func F_makeSimpleA_Expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeVar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeVar
func F_makeVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_makeTargetEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeTargetEntry
func F_makeTargetEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_makeFromExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeFromExpr
func F_makeFromExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeConst github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeConst
func F_makeConst(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_makeNullConst github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeNullConst
func F_makeNullConst(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeBoolConst github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeBoolConst
func F_makeBoolConst(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeBoolExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeBoolExpr
func F_makeBoolExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeAlias github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeAlias
func F_makeAlias(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeRelabelType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeRelabelType
func F_makeRelabelType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeRangeVar github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeRangeVar
func F_makeRangeVar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeNotNullConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeNotNullConstraint
func F_makeNotNullConstraint(m *base.Module, l0 int32) int32
//go:linkname F_makeTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeTypeName
func F_makeTypeName(m *base.Module, l0 int32) int32
//go:linkname F_makeFuncExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeFuncExpr
func F_makeFuncExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeDefElem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeDefElem
func F_makeDefElem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeFuncCall
func F_makeFuncCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_opclause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_opclause
func F_make_opclause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_andclause github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_andclause
func F_make_andclause(m *base.Module, l0 int32) int32
//go:linkname F_make_orclause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_orclause
func F_make_orclause(m *base.Module, l0 int32) int32
//go:linkname F_make_ands_explicit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_ands_explicit
func F_make_ands_explicit(m *base.Module, l0 int32) int32
//go:linkname F_make_ands_implicit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_ands_implicit
func F_make_ands_implicit(m *base.Module, l0 int32) int32
//go:linkname F_makeIndexInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeIndexInfo
func F_makeIndexInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32
//go:linkname F_mbms_add_member github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mbms_add_member
func F_mbms_add_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_mbms_add_members github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_mbms_add_members
func F_mbms_add_members(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_exprType github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_exprType
func F_exprType(m *base.Module, l0 int32) int32
//go:linkname F_exprTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprTypmod
func F_exprTypmod(m *base.Module, l0 int32) int32
//go:linkname F_applyRelabelType github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_applyRelabelType
func F_applyRelabelType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_exprCollation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprCollation
func F_exprCollation(m *base.Module, l0 int32) int32
//go:linkname F_relabel_to_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_relabel_to_typmod
func F_relabel_to_typmod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expression_returns_set github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expression_returns_set
func F_expression_returns_set(m *base.Module, l0 int32) int32
//go:linkname F_expression_returns_set_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_expression_returns_set_walker
func F_expression_returns_set_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_exprLocation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprLocation
func F_exprLocation(m *base.Module, l0 int32) int32
//go:linkname F_set_opfuncid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_opfuncid
func F_set_opfuncid(m *base.Module, l0 int32)
//go:linkname F_check_functions_in_node github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_functions_in_node
func F_check_functions_in_node(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_expression_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expression_tree_walker_impl
func F_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_query_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_query_tree_walker_impl
func F_query_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_range_table_entry_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_table_entry_walker_impl
func F_range_table_entry_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_expression_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_expression_tree_mutator_impl
func F_expression_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_query_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_query_tree_mutator_impl
func F_query_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_query_or_expression_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_query_or_expression_tree_walker_impl
func F_query_or_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_planstate_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_planstate_tree_walker_impl
func F_planstate_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_outBitmapset github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_outBitmapset
func F_outBitmapset(m *base.Module, l0 int32, l1 int32)
//go:linkname F_nodeToString github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_nodeToString
func F_nodeToString(m *base.Module, l0 int32) int32
//go:linkname F_makeParamList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeParamList
func F_makeParamList(m *base.Module, l0 int32) int32
//go:linkname F_elog_node_display github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_elog_node_display
func F_elog_node_display(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F__jumbleNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__jumbleNode
func F__jumbleNode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AppendJumble github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AppendJumble
func F_AppendJumble(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_stringToNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_stringToNode
func F_stringToNode(m *base.Module, l0 int32) int32
//go:linkname F_pg_strtok github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_strtok
func F_pg_strtok(m *base.Module, l0 int32) int32
//go:linkname F_debackslash github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_debackslash
func F_debackslash(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_readAttrNumberCols github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_readAttrNumberCols
func F_readAttrNumberCols(m *base.Module, l0 int32) int32
//go:linkname F_readOidCols github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_readOidCols
func F_readOidCols(m *base.Module, l0 int32) int32
//go:linkname F_readBoolCols github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_readBoolCols
func F_readBoolCols(m *base.Module, l0 int32) int32
//go:linkname F_readIntCols github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_readIntCols
func F_readIntCols(m *base.Module, l0 int32) int32
//go:linkname F_tbm_free github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tbm_free
func F_tbm_free(m *base.Module, l0 int32)
//go:linkname F_tbm_add_tuples github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tbm_add_tuples
func F_tbm_add_tuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_tbm_create_pagetable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tbm_create_pagetable
func F_tbm_create_pagetable(m *base.Module, l0 int32)
//go:linkname F_pagetable_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pagetable_insert
func F_pagetable_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeInteger github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeInteger
func F_makeInteger(m *base.Module, l0 int32) int32
//go:linkname F_makeFloat github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeFloat
func F_makeFloat(m *base.Module, l0 int32) int32
//go:linkname F_makeString github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeString
func F_makeString(m *base.Module, l0 int32) int32
//go:linkname F_geqo_eval github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_geqo_eval
func F_geqo_eval(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_gimme_tree github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gimme_tree
func F_gimme_tree(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_add_paths_to_append_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_paths_to_append_rel
func F_add_paths_to_append_rel(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_generate_useful_gather_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_generate_useful_gather_paths
func F_generate_useful_gather_paths(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_generate_partitionwise_join_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_partitionwise_join_paths
func F_generate_partitionwise_join_paths(m *base.Module, l0 int32, l1 int32)
//go:linkname F_subquery_is_pushdown_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_subquery_is_pushdown_safe
func F_subquery_is_pushdown_safe(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_window_run_conditions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_find_window_run_conditions
func F_find_window_run_conditions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_subquery_push_qual github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_subquery_push_qual
func F_subquery_push_qual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_clauselist_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clauselist_selectivity
func F_clauselist_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_clamp_row_est github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clamp_row_est
func F_clamp_row_est(m *base.Module, l0 float64) float64
//go:linkname F_cost_qual_eval_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cost_qual_eval_walker
func F_cost_qual_eval_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cost_samplescan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cost_samplescan
func F_cost_samplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_cost_qual_eval github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cost_qual_eval
func F_cost_qual_eval(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_compute_bitmap_pages github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_compute_bitmap_pages
func F_compute_bitmap_pages(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 int32, l5 int32) float64
//go:linkname F_cost_bitmap_tree_node github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cost_bitmap_tree_node
func F_cost_bitmap_tree_node(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_cost_ctescan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cost_ctescan
func F_cost_ctescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_cost_resultscan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cost_resultscan
func F_cost_resultscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_cost_tuplesort github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cost_tuplesort
func F_cost_tuplesort(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 float64, l5 int32, l6 float64)
//go:linkname F_cost_agg github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cost_agg
func F_cost_agg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64, l6 int32, l7 int32, l8 float64, l9 float64, l10 float64, l11 float64)
//go:linkname F_get_expr_width github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_expr_width
func F_get_expr_width(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_initial_cost_mergejoin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initial_cost_mergejoin
func F_initial_cost_mergejoin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_set_baserel_size_estimates github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_baserel_size_estimates
func F_set_baserel_size_estimates(m *base.Module, l0 int32, l1 int32)
//go:linkname F_set_rel_width github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_rel_width
func F_set_rel_width(m *base.Module, l0 int32, l1 int32)
//go:linkname F_set_subquery_size_estimates github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_subquery_size_estimates
func F_set_subquery_size_estimates(m *base.Module, l0 int32, l1 int32)
//go:linkname F_set_cte_size_estimates github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_cte_size_estimates
func F_set_cte_size_estimates(m *base.Module, l0 int32, l1 int32, l2 float64)
//go:linkname F_generate_join_implied_equalities github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_join_implied_equalities
func F_generate_join_implied_equalities(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_get_common_eclass_indexes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_common_eclass_indexes
func F_get_common_eclass_indexes(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_add_child_eq_member github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_child_eq_member
func F_add_child_eq_member(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_generate_implied_equalities_for_column github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_implied_equalities_for_column
func F_generate_implied_equalities_for_column(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_match_clause_to_index github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_match_clause_to_index
func F_match_clause_to_index(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_index_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_index_paths
func F_get_index_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_consider_index_join_outer_rels github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_consider_index_join_outer_rels
func F_consider_index_join_outer_rels(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_generate_bitmap_or_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_bitmap_or_paths
func F_generate_bitmap_or_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_choose_bitmap_and github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_choose_bitmap_and
func F_choose_bitmap_and(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_loop_count github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_loop_count
func F_get_loop_count(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_relation_has_unique_index_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_relation_has_unique_index_ext
func F_relation_has_unique_index_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_have_join_order_restriction github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_have_join_order_restriction
func F_have_join_order_restriction(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_join_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_join_rel
func F_make_join_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_compare_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_compare_pathkeys
func F_compare_pathkeys(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_useful_group_keys_orderings github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_useful_group_keys_orderings
func F_get_useful_group_keys_orderings(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_pathkey_from_sortinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_pathkey_from_sortinfo
func F_make_pathkey_from_sortinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_convert_subquery_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_subquery_pathkeys
func F_convert_subquery_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_pathkeys_for_sortclauses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_pathkeys_for_sortclauses
func F_make_pathkeys_for_sortclauses(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IsBinaryTidClause github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_IsBinaryTidClause
func F_IsBinaryTidClause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BuildParameterizedTidPaths github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BuildParameterizedTidPaths
func F_BuildParameterizedTidPaths(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RestrictInfoIsTidQual github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RestrictInfoIsTidQual
func F_RestrictInfoIsTidQual(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rel_is_distinct_for github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_rel_is_distinct_for
func F_rel_is_distinct_for(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_query_is_distinct_for github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_query_is_distinct_for
func F_query_is_distinct_for(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_create_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_plan
func F_create_plan(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_plan_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_plan_recurse
func F_create_plan_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_replace_nestloop_params_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_replace_nestloop_params_mutator
func F_replace_nestloop_params_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_order_qual_clauses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_order_qual_clauses
func F_order_qual_clauses(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_gating_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_gating_plan
func F_create_gating_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_use_physical_tlist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_use_physical_tlist
func F_use_physical_tlist(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_create_indexscan_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_indexscan_plan
func F_create_indexscan_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_planner github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_planner
func F_planner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_subquery_planner github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_subquery_planner
func F_subquery_planner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 int32) int32
//go:linkname F_expression_planner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expression_planner
func F_expression_planner(m *base.Module, l0 int32) int32
//go:linkname F_get_number_of_groups github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_number_of_groups
func F_get_number_of_groups(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32) float64
//go:linkname F_gather_grouping_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gather_grouping_paths
func F_gather_grouping_paths(m *base.Module, l0 int32, l1 int32)
//go:linkname F_consider_groupingsets_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_consider_groupingsets_paths
func F_consider_groupingsets_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 float64)
//go:linkname F_add_rte_to_flat_rtable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_rte_to_flat_rtable
func F_add_rte_to_flat_rtable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_extract_query_dependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_extract_query_dependencies
func F_extract_query_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_agg_clause_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_agg_clause_costs
func F_get_agg_clause_costs(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_replace_empty_jointree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_replace_empty_jointree
func F_replace_empty_jointree(m *base.Module, l0 int32)
//go:linkname F_pull_up_sublinks_qual_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pull_up_sublinks_qual_recurse
func F_pull_up_sublinks_qual_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_expand_virtual_generated_columns github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_expand_virtual_generated_columns
func F_expand_virtual_generated_columns(m *base.Module, l0 int32) int32
//go:linkname F_is_simple_subquery github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_is_simple_subquery
func F_is_simple_subquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_is_simple_union_all_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_is_simple_union_all_recurse
func F_is_simple_union_all_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pull_up_union_leaf_queries github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pull_up_union_leaf_queries
func F_pull_up_union_leaf_queries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_perform_pullup_replace_vars github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_perform_pullup_replace_vars
func F_perform_pullup_replace_vars(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_relids_in_jointree github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_relids_in_jointree
func F_get_relids_in_jointree(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_nullingrels_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_nullingrels_recurse
func F_get_nullingrels_recurse(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_negate_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_negate_clause
func F_negate_clause(m *base.Module, l0 int32) int32
//go:linkname F_canonicalize_qual github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_canonicalize_qual
func F_canonicalize_qual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pull_ors github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pull_ors
func F_pull_ors(m *base.Module, l0 int32) int32
//go:linkname F_pull_ands github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pull_ands
func F_pull_ands(m *base.Module, l0 int32) int32
//go:linkname F_adjust_appendrel_attrs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_adjust_appendrel_attrs
func F_adjust_appendrel_attrs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_adjust_appendrel_attrs_multilevel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_adjust_appendrel_attrs_multilevel
func F_adjust_appendrel_attrs_multilevel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_expression_returns_set_rows github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expression_returns_set_rows
func F_expression_returns_set_rows(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_contain_subplans github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_contain_subplans
func F_contain_subplans(m *base.Module, l0 int32) int32
//go:linkname F_contain_mutable_functions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_contain_mutable_functions
func F_contain_mutable_functions(m *base.Module, l0 int32) int32
//go:linkname F_contain_volatile_functions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_contain_volatile_functions
func F_contain_volatile_functions(m *base.Module, l0 int32) int32
//go:linkname F_is_parallel_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_is_parallel_safe
func F_is_parallel_safe(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_contain_leaked_vars github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_contain_leaked_vars
func F_contain_leaked_vars(m *base.Module, l0 int32) int32
//go:linkname F_find_nonnullable_rels github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_nonnullable_rels
func F_find_nonnullable_rels(m *base.Module, l0 int32) int32
//go:linkname F_is_strict_saop github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_is_strict_saop
func F_is_strict_saop(m *base.Module, l0 int32) int32
//go:linkname F_find_nonnullable_vars_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_find_nonnullable_vars_walker
func F_find_nonnullable_vars_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_eval_const_expressions_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_eval_const_expressions_mutator
func F_eval_const_expressions_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_estimate_expression_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_estimate_expression_value
func F_estimate_expression_value(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_inline_set_returning_function github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inline_set_returning_function
func F_inline_set_returning_function(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expand_partitioned_rtentry github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expand_partitioned_rtentry
func F_expand_partitioned_rtentry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_expand_single_inheritance_child github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expand_single_inheritance_child
func F_expand_single_inheritance_child(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_set_cheapest github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_cheapest
func F_set_cheapest(m *base.Module, l0 int32)
//go:linkname F_add_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_path
func F_add_path(m *base.Module, l0 int32, l1 int32)
//go:linkname F_add_partial_path github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_partial_path
func F_add_partial_path(m *base.Module, l0 int32, l1 int32)
//go:linkname F_create_seqscan_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_create_seqscan_path
func F_create_seqscan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_bitmap_heap_path github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_create_bitmap_heap_path
func F_create_bitmap_heap_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 int32) int32
//go:linkname F_create_tidscan_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_create_tidscan_path
func F_create_tidscan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_append_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_append_path
func F_create_append_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 float64) int32
//go:linkname F_create_subqueryscan_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_subqueryscan_path
func F_create_subqueryscan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_create_mergejoin_path github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_create_mergejoin_path
func F_create_mergejoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) int32
//go:linkname F_create_incremental_sort_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_incremental_sort_path
func F_create_incremental_sort_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) int32
//go:linkname F_create_sort_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_sort_path
func F_create_sort_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32
//go:linkname F_create_group_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_group_path
func F_create_group_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) int32
//go:linkname F_create_agg_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_agg_path
func F_create_agg_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 float64) int32
//go:linkname F_estimate_rel_size github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_estimate_rel_size
func F_estimate_rel_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_predicate_implied_by github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_predicate_implied_by
func F_predicate_implied_by(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_predicate_refuted_by github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_predicate_refuted_by
func F_predicate_refuted_by(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_expand_planner_arrays github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_expand_planner_arrays
func F_expand_planner_arrays(m *base.Module, l0 int32, l1 int32)
//go:linkname F_build_simple_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_simple_rel
func F_build_simple_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_base_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_base_rel
func F_find_base_rel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_base_rel_ignore_join github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_find_base_rel_ignore_join
func F_find_base_rel_ignore_join(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fetch_upper_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fetch_upper_rel
func F_fetch_upper_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_childrel_parents github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_find_childrel_parents
func F_find_childrel_parents(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_baserel_parampathinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_baserel_parampathinfo
func F_get_baserel_parampathinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_actual_clauses github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_actual_clauses
func F_get_actual_clauses(m *base.Module, l0 int32) int32
//go:linkname F_extract_actual_clauses github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_extract_actual_clauses
func F_extract_actual_clauses(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_sortgroupclause_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_sortgroupclause_expr
func F_get_sortgroupclause_expr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_sortgrouplist_exprs github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_sortgrouplist_exprs
func F_get_sortgrouplist_exprs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_tlist_from_pathtarget github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_tlist_from_pathtarget
func F_make_tlist_from_pathtarget(m *base.Module, l0 int32) int32
//go:linkname F_copy_pathtarget github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_copy_pathtarget
func F_copy_pathtarget(m *base.Module, l0 int32) int32
//go:linkname F_create_empty_pathtarget github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_empty_pathtarget
func F_create_empty_pathtarget(m *base.Module) int32
//go:linkname F_add_column_to_pathtarget github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_column_to_pathtarget
func F_add_column_to_pathtarget(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_add_new_columns_to_pathtarget github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_new_columns_to_pathtarget
func F_add_new_columns_to_pathtarget(m *base.Module, l0 int32, l1 int32)
//go:linkname F_apply_pathtarget_labeling_to_tlist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_apply_pathtarget_labeling_to_tlist
func F_apply_pathtarget_labeling_to_tlist(m *base.Module, l0 int32, l1 int32)
//go:linkname F_split_pathtarget_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_split_pathtarget_walker
func F_split_pathtarget_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_add_sp_item_to_pathtarget github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_sp_item_to_pathtarget
func F_add_sp_item_to_pathtarget(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pull_varnos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pull_varnos
func F_pull_varnos(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pull_varattnos github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pull_varattnos
func F_pull_varattnos(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_contain_vars_of_level github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_contain_vars_of_level
func F_contain_vars_of_level(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pull_var_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pull_var_clause
func F_pull_var_clause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_flatten_join_alias_vars github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_flatten_join_alias_vars
func F_flatten_join_alias_vars(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_partition_list_bsearch github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_partition_list_bsearch
func F_partition_list_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_partition_range_datum_bsearch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_partition_range_datum_bsearch
func F_partition_range_datum_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_compute_partition_hash_value github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_compute_partition_hash_value
func F_compute_partition_hash_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64
//go:linkname F_get_steps_using_prefix github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_steps_using_prefix
func F_get_steps_using_prefix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_AutoVacLauncherShutdown github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AutoVacLauncherShutdown
func F_AutoVacLauncherShutdown(m *base.Module)
//go:linkname F_do_start_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_do_start_worker
func F_do_start_worker(m *base.Module) int32
//go:linkname F_launcher_determine_sleep github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_launcher_determine_sleep
func F_launcher_determine_sleep(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_database_list github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_database_list
func F_get_database_list(m *base.Module) int32
//go:linkname F_AuxiliaryProcessMainCommon github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AuxiliaryProcessMainCommon
func F_AuxiliaryProcessMainCommon(m *base.Module)
//go:linkname F_SanityCheckBackgroundWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SanityCheckBackgroundWorker
func F_SanityCheckBackgroundWorker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AbsorbSyncRequests github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AbsorbSyncRequests
func F_AbsorbSyncRequests(m *base.Module)
//go:linkname F_CheckArchiveTimeout github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckArchiveTimeout
func F_CheckArchiveTimeout(m *base.Module)
//go:linkname F_ProcessMainLoopInterrupts github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessMainLoopInterrupts
func F_ProcessMainLoopInterrupts(m *base.Module)
//go:linkname F_MaxLivePostmasterChildren github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MaxLivePostmasterChildren
func F_MaxLivePostmasterChildren(m *base.Module) int32
//go:linkname F_InitPostmasterChildSlots github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InitPostmasterChildSlots
func F_InitPostmasterChildSlots(m *base.Module)
//go:linkname F_signal_child github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_signal_child
func F_signal_child(m *base.Module, l0 int32, l1 int32)
//go:linkname F_update_metainfo_datafile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_update_metainfo_datafile
func F_update_metainfo_datafile(m *base.Module)
//go:linkname F_logfile_rotate_dest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_logfile_rotate_dest
func F_logfile_rotate_dest(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_GetOldestUnsummarizedLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetOldestUnsummarizedLSN
func F_GetOldestUnsummarizedLSN(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_freesubre github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_freesubre
func F_freesubre(m *base.Module, l0 int32, l1 int32)
//go:linkname F_newstate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_newstate
func F_newstate(m *base.Module, l0 int32) int32
//go:linkname F_rainbow github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_rainbow
func F_rainbow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_createarc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_createarc
func F_createarc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_next github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_next
func F_next(m *base.Module, l0 int32) int32
//go:linkname F_cparc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cparc
func F_cparc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pg_reg_getcolor github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_reg_getcolor
func F_pg_reg_getcolor(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_subcoloronechr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_subcoloronechr
func F_subcoloronechr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_allcases github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_allcases
func F_allcases(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_subcolorcvec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_subcolorcvec
func F_subcolorcvec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_cclasscvec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cclasscvec
func F_cclasscvec(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_charclasscomplement github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_charclasscomplement
func F_charclasscomplement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_cleartraverse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cleartraverse
func F_cleartraverse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_duptraverse github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_duptraverse
func F_duptraverse(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_newdfa github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_newdfa
func F_newdfa(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getvacant github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getvacant
func F_getvacant(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_traverse_lacons github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_traverse_lacons
func F_traverse_lacons(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_pa_free_worker_info github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pa_free_worker_info
func F_pa_free_worker_info(m *base.Module, l0 int32)
//go:linkname F_build_index_value_desc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_build_index_value_desc
func F_build_index_value_desc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_logicalrep_worker_stop_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_logicalrep_worker_stop_internal
func F_logicalrep_worker_stop_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ApplyLauncherForgetWorkerStartTime github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ApplyLauncherForgetWorkerStartTime
func F_ApplyLauncherForgetWorkerStartTime(m *base.Module, l0 int32)
//go:linkname F_AtEOXact_ApplyLauncher github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AtEOXact_ApplyLauncher
func F_AtEOXact_ApplyLauncher(m *base.Module, l0 int32)
//go:linkname F_CreateInitDecodingContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateInitDecodingContext
func F_CreateInitDecodingContext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_DecodingContextFindStartpoint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_DecodingContextFindStartpoint
func F_DecodingContextFindStartpoint(m *base.Module, l0 int32)
//go:linkname F_FreeDecodingContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FreeDecodingContext
func F_FreeDecodingContext(m *base.Module, l0 int32)
//go:linkname F_replorigin_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_replorigin_by_name
func F_replorigin_by_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replorigin_by_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replorigin_by_oid
func F_replorigin_by_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replorigin_session_advance github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_replorigin_session_advance
func F_replorigin_session_advance(m *base.Module, l0 int64, l1 int64)
//go:linkname F_replorigin_check_prerequisites github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replorigin_check_prerequisites
func F_replorigin_check_prerequisites(m *base.Module)
//go:linkname F_GetRelationIdentityOrPK github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetRelationIdentityOrPK
func F_GetRelationIdentityOrPK(m *base.Module, l0 int32) int32
//go:linkname F_ReorderBufferFreeChange github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReorderBufferFreeChange
func F_ReorderBufferFreeChange(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReorderBufferTXNByXid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReorderBufferTXNByXid
func F_ReorderBufferTXNByXid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_ReorderBufferStreamTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReorderBufferStreamTXN
func F_ReorderBufferStreamTXN(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferCheckAndTruncateAbortedTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReorderBufferCheckAndTruncateAbortedTXN
func F_ReorderBufferCheckAndTruncateAbortedTXN(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReorderBufferSerializeTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReorderBufferSerializeTXN
func F_ReorderBufferSerializeTXN(m *base.Module, l0 int32, l1 int32)
//go:linkname F_update_local_synced_slot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_update_local_synced_slot
func F_update_local_synced_slot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_store_flush_position github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_store_flush_position
func F_store_flush_position(m *base.Module, l0 int64, l1 int64)
//go:linkname F_AtEOXact_LogicalRepWorkers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AtEOXact_LogicalRepWorkers
func F_AtEOXact_LogicalRepWorkers(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReplicationSlotRelease
func F_ReplicationSlotRelease(m *base.Module)
//go:linkname F_ReplicationSlotCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplicationSlotCleanup
func F_ReplicationSlotCleanup(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReplicationSlotCreate
func F_ReplicationSlotCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_SearchNamedReplicationSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SearchNamedReplicationSlot
func F_SearchNamedReplicationSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReplicationSlotAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplicationSlotAcquire
func F_ReplicationSlotAcquire(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReplicationSlotDropPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReplicationSlotDropPtr
func F_ReplicationSlotDropPtr(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotsComputeRequiredXmin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotsComputeRequiredXmin
func F_ReplicationSlotsComputeRequiredXmin(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotMarkDirty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReplicationSlotMarkDirty
func F_ReplicationSlotMarkDirty(m *base.Module)
//go:linkname F_ReplicationSlotSave github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotSave
func F_ReplicationSlotSave(m *base.Module)
//go:linkname F_ReplicationSlotPersist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplicationSlotPersist
func F_ReplicationSlotPersist(m *base.Module)
//go:linkname F_ReplicationSlotsDropDBSlots github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotsDropDBSlots
func F_ReplicationSlotsDropDBSlots(m *base.Module, l0 int32)
//go:linkname F_CheckSlotRequirements github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckSlotRequirements
func F_CheckSlotRequirements(m *base.Module)
//go:linkname F_SyncRepWaitForLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SyncRepWaitForLSN
func F_SyncRepWaitForLSN(m *base.Module, l0 int64, l1 int32)
//go:linkname F_SyncRepInitConfig github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SyncRepInitConfig
func F_SyncRepInitConfig(m *base.Module)
//go:linkname F_SyncRepUpdateSyncStandbysDefined github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SyncRepUpdateSyncStandbysDefined
func F_SyncRepUpdateSyncStandbysDefined(m *base.Module)
//go:linkname F_WalRcvFetchTimeLineHistoryFiles github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_WalRcvFetchTimeLineHistoryFiles
func F_WalRcvFetchTimeLineHistoryFiles(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogWalRcvSendReply github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogWalRcvSendReply
func F_XLogWalRcvSendReply(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogWalRcvSendHSFeedback github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogWalRcvSendHSFeedback
func F_XLogWalRcvSendHSFeedback(m *base.Module, l0 int32)
//go:linkname F_ProcessWalSndrMessage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ProcessWalSndrMessage
func F_ProcessWalSndrMessage(m *base.Module, l0 int64, l1 int64)
//go:linkname F_XLogWalRcvClose github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XLogWalRcvClose
func F_XLogWalRcvClose(m *base.Module, l0 int32)
//go:linkname F_XLogWalRcvFlush github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogWalRcvFlush
func F_XLogWalRcvFlush(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetWalRcvFlushRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetWalRcvFlushRecPtr
func F_GetWalRcvFlushRecPtr(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_WalSndWakeup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WalSndWakeup
func F_WalSndWakeup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ProcessRepliesIfAny github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessRepliesIfAny
func F_ProcessRepliesIfAny(m *base.Module)
//go:linkname F_WalSndKeepalive github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WalSndKeepalive
func F_WalSndKeepalive(m *base.Module, l0 int32, l1 int64)
//go:linkname F_WalSndWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WalSndWait
func F_WalSndWait(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_WalSndShutdown github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WalSndShutdown
func F_WalSndShutdown(m *base.Module)
//go:linkname F_DefineQueryRewrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DefineQueryRewrite
func F_DefineQueryRewrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_get_view_query github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_view_query
func F_get_view_query(m *base.Module, l0 int32) int32
//go:linkname F_adjust_view_column_set github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_adjust_view_column_set
func F_adjust_view_column_set(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_error_view_not_updatable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_error_view_not_updatable
func F_error_view_not_updatable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_expand_generated_columns_in_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_expand_generated_columns_in_expr
func F_expand_generated_columns_in_expr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fireRIRrules github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fireRIRrules
func F_fireRIRrules(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CombineRangeTables github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CombineRangeTables
func F_CombineRangeTables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_OffsetVarNodes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_OffsetVarNodes
func F_OffsetVarNodes(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ChangeVarNodes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ChangeVarNodes
func F_ChangeVarNodes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_IncrementVarSublevelsUp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_IncrementVarSublevelsUp
func F_IncrementVarSublevelsUp(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_map_variable_attnos github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_map_variable_attnos
func F_map_variable_attnos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_statext_dependencies_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_statext_dependencies_deserialize
func F_statext_dependencies_deserialize(m *base.Module, l0 int32) int32
//go:linkname F_stats_check_required_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_stats_check_required_arg
func F_stats_check_required_arg(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pgaio_io_reclaim github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgaio_io_reclaim
func F_pgaio_io_reclaim(m *base.Module, l0 int32)
//go:linkname F_pgaio_submit_staged github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgaio_submit_staged
func F_pgaio_submit_staged(m *base.Module)
//go:linkname F_pgaio_io_update_state github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgaio_io_update_state
func F_pgaio_io_update_state(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgaio_wref_wait github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgaio_wref_wait
func F_pgaio_wref_wait(m *base.Module, l0 int32)
//go:linkname F_AtEOXact_Aio github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtEOXact_Aio
func F_AtEOXact_Aio(m *base.Module)
//go:linkname F_pgaio_result_report github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgaio_result_report
func F_pgaio_result_report(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_read_stream_begin_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_read_stream_begin_relation
func F_read_stream_begin_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_read_stream_look_ahead github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_stream_look_ahead
func F_read_stream_look_ahead(m *base.Module, l0 int32)
//go:linkname F_read_stream_reset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_read_stream_reset
func F_read_stream_reset(m *base.Module, l0 int32)
//go:linkname F_read_stream_end github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_stream_end
func F_read_stream_end(m *base.Module, l0 int32)
//go:linkname F_BufTableHashCode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BufTableHashCode
func F_BufTableHashCode(m *base.Module, l0 int32) int32
//go:linkname F_BufTableLookup github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BufTableLookup
func F_BufTableLookup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BufTableDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BufTableDelete
func F_BufTableDelete(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnpinBufferNoOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnpinBufferNoOwner
func F_UnpinBufferNoOwner(m *base.Module, l0 int32)
//go:linkname F_ReservePrivateRefCountEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReservePrivateRefCountEntry
func F_ReservePrivateRefCountEntry(m *base.Module)
//go:linkname F_LockBufHdr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LockBufHdr
func F_LockBufHdr(m *base.Module, l0 int32) int32
//go:linkname F_ReadBufferExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadBufferExtended
func F_ReadBufferExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_StartReadBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_StartReadBuffer
func F_StartReadBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_WaitReadBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WaitReadBuffers
func F_WaitReadBuffers(m *base.Module, l0 int32)
//go:linkname F_ExtendBufferedRelCommon github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExtendBufferedRelCommon
func F_ExtendBufferedRelCommon(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_ExtendBufferedRelTo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExtendBufferedRelTo
func F_ExtendBufferedRelTo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ReleaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReleaseBuffer
func F_ReleaseBuffer(m *base.Module, l0 int32)
//go:linkname F_MarkBufferDirty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MarkBufferDirty
func F_MarkBufferDirty(m *base.Module, l0 int32)
//go:linkname F_SyncOneBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SyncOneBuffer
func F_SyncOneBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FlushBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FlushBuffer
func F_FlushBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_IssuePendingWritebacks github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_IssuePendingWritebacks
func F_IssuePendingWritebacks(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationGetNumberOfBlocksInFork github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetNumberOfBlocksInFork
func F_RelationGetNumberOfBlocksInFork(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_InvalidateBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_InvalidateBuffer
func F_InvalidateBuffer(m *base.Module, l0 int32)
//go:linkname F_WaitIO github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WaitIO
func F_WaitIO(m *base.Module, l0 int32)
//go:linkname F_DropDatabaseBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DropDatabaseBuffers
func F_DropDatabaseBuffers(m *base.Module, l0 int32)
//go:linkname F_FlushRelationBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FlushRelationBuffers
func F_FlushRelationBuffers(m *base.Module, l0 int32)
//go:linkname F_UnlockReleaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UnlockReleaseBuffer
func F_UnlockReleaseBuffer(m *base.Module, l0 int32)
//go:linkname F_LockBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockBuffer
func F_LockBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnlockBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UnlockBuffers
func F_UnlockBuffers(m *base.Module)
//go:linkname F_ConditionalLockBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConditionalLockBuffer
func F_ConditionalLockBuffer(m *base.Module, l0 int32) int32
//go:linkname F_CheckBufferIsPinnedOnce github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckBufferIsPinnedOnce
func F_CheckBufferIsPinnedOnce(m *base.Module, l0 int32)
//go:linkname F_StrategyNotifyBgWriter github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_StrategyNotifyBgWriter
func F_StrategyNotifyBgWriter(m *base.Module, l0 int32)
//go:linkname F_UnpinLocalBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UnpinLocalBuffer
func F_UnpinLocalBuffer(m *base.Module, l0 int32)
//go:linkname F_BufFileCreateTemp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BufFileCreateTemp
func F_BufFileCreateTemp(m *base.Module, l0 int32) int32
//go:linkname F_BufFileDumpBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BufFileDumpBuffer
func F_BufFileDumpBuffer(m *base.Module, l0 int32)
//go:linkname F_BufFileSeek github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BufFileSeek
func F_BufFileSeek(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_copydir github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_copydir
func F_copydir(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_file_exists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_file_exists
func F_pg_file_exists(m *base.Module, l0 int32) int32
//go:linkname F_fsync_fname github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fsync_fname
func F_fsync_fname(m *base.Module, l0 int32, l1 int32)
//go:linkname F_OpenTransientFilePerm github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_OpenTransientFilePerm
func F_OpenTransientFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32
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
//go:linkname F_set_max_safe_fds github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_set_max_safe_fds
func F_set_max_safe_fds(m *base.Module)
//go:linkname F_BasicOpenFilePerm github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BasicOpenFilePerm
func F_BasicOpenFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LruDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LruDelete
func F_LruDelete(m *base.Module, l0 int32)
//go:linkname F_PathNameOpenFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PathNameOpenFile
func F_PathNameOpenFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_PathNameOpenFilePerm github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PathNameOpenFilePerm
func F_PathNameOpenFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_AllocateDir github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AllocateDir
func F_AllocateDir(m *base.Module, l0 int32) int32
//go:linkname F_ReadDirExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReadDirExtended
func F_ReadDirExtended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FreeDir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeDir
func F_FreeDir(m *base.Module, l0 int32)
//go:linkname F_FileWriteV github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FileWriteV
func F_FileWriteV(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_FileSync github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FileSync
func F_FileSync(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AllocateFile github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AllocateFile
func F_AllocateFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_reserveAllocatedDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_reserveAllocatedDesc
func F_reserveAllocatedDesc(m *base.Module) int32
//go:linkname F_FreeFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FreeFile
func F_FreeFile(m *base.Module, l0 int32) int32
//go:linkname F_ReadDir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadDir
func F_ReadDir(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AtEOSubXact_Files github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AtEOSubXact_Files
func F_AtEOSubXact_Files(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AtEOXact_Files github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AtEOXact_Files
func F_AtEOXact_Files(m *base.Module, l0 int32)
//go:linkname F_RecordPageWithFreeSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RecordPageWithFreeSpace
func F_RecordPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_FreeSpaceMapVacuum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeSpaceMapVacuum
func F_FreeSpaceMapVacuum(m *base.Module, l0 int32)
//go:linkname F_fsm_vacuum_page github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fsm_vacuum_page
func F_fsm_vacuum_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_GetFreeIndexPage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetFreeIndexPage
func F_GetFreeIndexPage(m *base.Module, l0 int32) int32
//go:linkname F_BarrierArriveAndWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BarrierArriveAndWait
func F_BarrierArriveAndWait(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BarrierArriveAndDetachExceptLast github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BarrierArriveAndDetachExceptLast
func F_BarrierArriveAndDetachExceptLast(m *base.Module, l0 int32) int32
//go:linkname F_BarrierDetach github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BarrierDetach
func F_BarrierDetach(m *base.Module, l0 int32)
//go:linkname F_dsm_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsm_attach
func F_dsm_attach(m *base.Module, l0 int32) int32
//go:linkname F_dsm_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsm_detach
func F_dsm_detach(m *base.Module, l0 int32)
//go:linkname F_dsm_unpin_segment github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dsm_unpin_segment
func F_dsm_unpin_segment(m *base.Module, l0 int32)
//go:linkname F_on_dsm_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_on_dsm_detach
func F_on_dsm_detach(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dsm_impl_op github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dsm_impl_op
func F_dsm_impl_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_proc_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_proc_exit
func F_proc_exit(m *base.Module, l0 int32)
//go:linkname F_on_proc_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_on_proc_exit
func F_on_proc_exit(m *base.Module, l0 int32)
//go:linkname F_before_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_before_shmem_exit
func F_before_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_on_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_on_shmem_exit
func F_on_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CalculateShmemSize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CalculateShmemSize
func F_CalculateShmemSize(m *base.Module, l0 int32) int32
//go:linkname F_WaitLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_WaitLatch
func F_WaitLatch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_WaitLatchOrSocket github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WaitLatchOrSocket
func F_WaitLatchOrSocket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SetLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SetLatch
func F_SetLatch(m *base.Module, l0 int32)
//go:linkname F_ProcArrayRemove github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ProcArrayRemove
func F_ProcArrayRemove(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ProcArrayEndTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcArrayEndTransaction
func F_ProcArrayEndTransaction(m *base.Module, l0 int32, l1 int32)
//go:linkname F_KnownAssignedXidsAdd github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_KnownAssignedXidsAdd
func F_KnownAssignedXidsAdd(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_TransactionIdIsInProgress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TransactionIdIsInProgress
func F_TransactionIdIsInProgress(m *base.Module, l0 int32) int32
//go:linkname F_ComputeXidHorizons github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ComputeXidHorizons
func F_ComputeXidHorizons(m *base.Module, l0 int32)
//go:linkname F_GlobalVisHorizonKindForRel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GlobalVisHorizonKindForRel
func F_GlobalVisHorizonKindForRel(m *base.Module, l0 int32) int32
//go:linkname F_GetOldestSafeDecodingTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetOldestSafeDecodingTransactionId
func F_GetOldestSafeDecodingTransactionId(m *base.Module, l0 int32) int32
//go:linkname F_BackendPidGetProc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BackendPidGetProc
func F_BackendPidGetProc(m *base.Module, l0 int32) int32
//go:linkname F_CountDBBackends github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CountDBBackends
func F_CountDBBackends(m *base.Module, l0 int32) int32
//go:linkname F_CancelDBBackends github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CancelDBBackends
func F_CancelDBBackends(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExpireTreeKnownAssignedTransactionIds github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExpireTreeKnownAssignedTransactionIds
func F_ExpireTreeKnownAssignedTransactionIds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_SendProcSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SendProcSignal
func F_SendProcSignal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EmitProcSignalBarrier github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EmitProcSignalBarrier
func F_EmitProcSignalBarrier(m *base.Module) int64
//go:linkname F_WaitForProcSignalBarrier github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WaitForProcSignalBarrier
func F_WaitForProcSignalBarrier(m *base.Module, l0 int64)
//go:linkname F_ProcessProcSignalBarrier github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcessProcSignalBarrier
func F_ProcessProcSignalBarrier(m *base.Module)
//go:linkname F_shm_mq_set_sender github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_shm_mq_set_sender
func F_shm_mq_set_sender(m *base.Module, l0 int32, l1 int32)
//go:linkname F_shm_mq_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_mq_attach
func F_shm_mq_attach(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_shm_mq_send_bytes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_shm_mq_send_bytes
func F_shm_mq_send_bytes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_shm_mq_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_mq_detach
func F_shm_mq_detach(m *base.Module, l0 int32)
//go:linkname F_shm_toc_allocate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_shm_toc_allocate
func F_shm_toc_allocate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_shm_toc_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_toc_insert
func F_shm_toc_insert(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_shm_toc_lookup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_shm_toc_lookup
func F_shm_toc_lookup(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_ShmemAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ShmemAlloc
func F_ShmemAlloc(m *base.Module, l0 int32) int32
//go:linkname F_ShmemInitStruct github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ShmemInitStruct
func F_ShmemInitStruct(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ShmemInitHash github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ShmemInitHash
func F_ShmemInitHash(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_add_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_size
func F_add_size(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_mul_size github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_mul_size
func F_mul_size(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_signal_backend github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_signal_backend
func F_pg_signal_backend(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SendSharedInvalidMessages github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SendSharedInvalidMessages
func F_SendSharedInvalidMessages(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReceiveSharedInvalidMessages github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReceiveSharedInvalidMessages
func F_ReceiveSharedInvalidMessages(m *base.Module)
//go:linkname F_ProcessCatchupInterrupt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessCatchupInterrupt
func F_ProcessCatchupInterrupt(m *base.Module)
//go:linkname F_LogRecoveryConflict github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogRecoveryConflict
func F_LogRecoveryConflict(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32)
//go:linkname F_StandbyReleaseLockTree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StandbyReleaseLockTree
func F_StandbyReleaseLockTree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LogStandbySnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogStandbySnapshot
func F_LogStandbySnapshot(m *base.Module) int64
//go:linkname F_CreateWaitEventSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateWaitEventSet
func F_CreateWaitEventSet(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FreeWaitEventSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeWaitEventSet
func F_FreeWaitEventSet(m *base.Module, l0 int32)
//go:linkname F_AddWaitEventToSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AddWaitEventToSet
func F_AddWaitEventToSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_WaitEventSetWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WaitEventSetWait
func F_WaitEventSetWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_inv_open github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inv_open
func F_inv_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_inv_read github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inv_read
func F_inv_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_inv_write github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_inv_write
func F_inv_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ConditionVariableCancelSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ConditionVariableCancelSleep
func F_ConditionVariableCancelSleep(m *base.Module)
//go:linkname F_ConditionVariableTimedSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ConditionVariableTimedSleep
func F_ConditionVariableTimedSleep(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ConditionVariableSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConditionVariableSignal
func F_ConditionVariableSignal(m *base.Module, l0 int32)
//go:linkname F_ConditionVariableBroadcast github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ConditionVariableBroadcast
func F_ConditionVariableBroadcast(m *base.Module, l0 int32)
//go:linkname F_FindLockCycleRecurseMember github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FindLockCycleRecurseMember
func F_FindLockCycleRecurseMember(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_LockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockRelationOid
func F_LockRelationOid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnlockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockRelationOid
func F_UnlockRelationOid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LockRelationForExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LockRelationForExtension
func F_LockRelationForExtension(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnlockRelationForExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockRelationForExtension
func F_UnlockRelationForExtension(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LockTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockTuple
func F_LockTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_UnlockTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockTuple
func F_UnlockTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XactLockTableWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XactLockTableWait
func F_XactLockTableWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_LockSharedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockSharedObject
func F_LockSharedObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_UnlockSharedObjectForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_UnlockSharedObjectForSession
func F_UnlockSharedObjectForSession(m *base.Module, l0 int32)
//go:linkname F_UnlockApplyTransactionForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_UnlockApplyTransactionForSession
func F_UnlockApplyTransactionForSession(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_DescribeLockTag github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DescribeLockTag
func F_DescribeLockTag(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RemoveLocalLock github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RemoveLocalLock
func F_RemoveLocalLock(m *base.Module, l0 int32)
//go:linkname F_LockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockAcquire
func F_LockAcquire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SetupLockInTable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SetupLockInTable
func F_SetupLockInTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_AbortStrongLockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AbortStrongLockAcquire
func F_AbortStrongLockAcquire(m *base.Module)
//go:linkname F_LockCheckConflicts github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LockCheckConflicts
func F_LockCheckConflicts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_WaitOnLock github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WaitOnLock
func F_WaitOnLock(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LockRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LockRelease
func F_LockRelease(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReleaseLockIfHeld github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReleaseLockIfHeld
func F_ReleaseLockIfHeld(m *base.Module, l0 int32, l1 int32)
//go:linkname F_VirtualXactLockTableInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_VirtualXactLockTableInsert
func F_VirtualXactLockTableInsert(m *base.Module, l0 int32)
//go:linkname F_LWLockShmemSize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LWLockShmemSize
func F_LWLockShmemSize(m *base.Module) int32
//go:linkname F_LWLockQueueSelf github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LWLockQueueSelf
func F_LWLockQueueSelf(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LWLockDequeueSelf github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LWLockDequeueSelf
func F_LWLockDequeueSelf(m *base.Module, l0 int32)
//go:linkname F_LWLockDisownInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LWLockDisownInternal
func F_LWLockDisownInternal(m *base.Module, l0 int32) int32
//go:linkname F_LWLockRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LWLockRelease
func F_LWLockRelease(m *base.Module, l0 int32)
//go:linkname F_ReleasePredicateLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReleasePredicateLocks
func F_ReleasePredicateLocks(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PredicateLockRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PredicateLockRelation
func F_PredicateLockRelation(m *base.Module, l0 int32, l1 int32)
//go:linkname F_DecrementParentLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecrementParentLocks
func F_DecrementParentLocks(m *base.Module, l0 int32)
//go:linkname F_PredicateLockPageSplit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PredicateLockPageSplit
func F_PredicateLockPageSplit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CheckForSerializableConflictOutNeeded github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckForSerializableConflictOutNeeded
func F_CheckForSerializableConflictOutNeeded(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FlagRWConflict github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FlagRWConflict
func F_FlagRWConflict(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CheckForSerializableConflictIn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckForSerializableConflictIn
func F_CheckForSerializableConflictIn(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CheckTableForSerializableConflictIn github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CheckTableForSerializableConflictIn
func F_CheckTableForSerializableConflictIn(m *base.Module, l0 int32)
//go:linkname F_PGProcShmemSize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PGProcShmemSize
func F_PGProcShmemSize(m *base.Module) int32
//go:linkname F_InitProcess github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InitProcess
func F_InitProcess(m *base.Module)
//go:linkname F_LockErrorCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockErrorCleanup
func F_LockErrorCleanup(m *base.Module)
//go:linkname F_ProcWaitForSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcWaitForSignal
func F_ProcWaitForSignal(m *base.Module, l0 int32)
//go:linkname F_s_lock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_s_lock
func F_s_lock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_perform_spin_delay github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_perform_spin_delay
func F_perform_spin_delay(m *base.Module, l0 int32)
//go:linkname F_PageInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PageInit
func F_PageInit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PageAddItemExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PageAddItemExtended
func F_PageAddItemExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_PageGetTempPageCopySpecial github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageGetTempPageCopySpecial
func F_PageGetTempPageCopySpecial(m *base.Module, l0 int32) int32
//go:linkname F_PageRestoreTempPage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PageRestoreTempPage
func F_PageRestoreTempPage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PageIndexTupleDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageIndexTupleDelete
func F_PageIndexTupleDelete(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PageIndexTupleOverwrite github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageIndexTupleOverwrite
func F_PageIndexTupleOverwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ForgetDatabaseSyncRequests github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ForgetDatabaseSyncRequests
func F_ForgetDatabaseSyncRequests(m *base.Module, l0 int32)
//go:linkname F_smgropen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgropen
func F_smgropen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgrclose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrclose
func F_smgrclose(m *base.Module, l0 int32)
//go:linkname F_smgrexists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrexists
func F_smgrexists(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgrcreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgrcreate
func F_smgrcreate(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_smgrdounlinkall github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_smgrdounlinkall
func F_smgrdounlinkall(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_smgrnblocks github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_smgrnblocks
func F_smgrnblocks(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ProcessInterrupts github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ProcessInterrupts
func F_ProcessInterrupts(m *base.Module)
//go:linkname F_ShowUsage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ShowUsage
func F_ShowUsage(m *base.Module, l0 int32)
//go:linkname F_pg_rewrite_query github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_rewrite_query
func F_pg_rewrite_query(m *base.Module, l0 int32) int32
//go:linkname F_CreateQueryDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CreateQueryDesc
func F_CreateQueryDesc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_FreeQueryDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeQueryDesc
func F_FreeQueryDesc(m *base.Module, l0 int32)
//go:linkname F_FetchPortalTargetList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FetchPortalTargetList
func F_FetchPortalTargetList(m *base.Module, l0 int32) int32
//go:linkname F_PortalRunUtility github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PortalRunUtility
func F_PortalRunUtility(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ProcessQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ProcessQuery
func F_ProcessQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_CreateCommandTag github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateCommandTag
func F_CreateCommandTag(m *base.Module, l0 int32) int32
//go:linkname F_UtilityReturnsTuples github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_UtilityReturnsTuples
func F_UtilityReturnsTuples(m *base.Module, l0 int32) int32
//go:linkname F_UtilityTupleDescriptor github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UtilityTupleDescriptor
func F_UtilityTupleDescriptor(m *base.Module, l0 int32) int32
//go:linkname F_getNextFlagFromString github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_getNextFlagFromString
func F_getNextFlagFromString(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_t_isalnum_cstr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_t_isalnum_cstr
func F_t_isalnum_cstr(m *base.Module, l0 int32) int32
//go:linkname F_tsearch_readline_begin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tsearch_readline_begin
func F_tsearch_readline_begin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tsearch_readline github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tsearch_readline
func F_tsearch_readline(m *base.Module, l0 int32) int32
//go:linkname F_tsearch_readline_end github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tsearch_readline_end
func F_tsearch_readline_end(m *base.Module, l0 int32)
//go:linkname F_LexizeExec github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LexizeExec
func F_LexizeExec(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addHLParsedLex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addHLParsedLex
func F_addHLParsedLex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_tsearch_config_filename github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_tsearch_config_filename
func F_get_tsearch_config_filename(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_prs_setup_firstcall github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_prs_setup_firstcall
func F_prs_setup_firstcall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_prs_process_call github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_prs_process_call
func F_prs_process_call(m *base.Module, l0 int32) int32
//go:linkname F_headline_json_value github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_headline_json_value
func F_headline_json_value(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgstat_clear_backend_activity_snapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_clear_backend_activity_snapshot
func F_pgstat_clear_backend_activity_snapshot(m *base.Module)
//go:linkname F_pgstat_get_beentry_by_proc_number github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_get_beentry_by_proc_number
func F_pgstat_get_beentry_by_proc_number(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_get_local_beentry_by_index github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_get_local_beentry_by_index
func F_pgstat_get_local_beentry_by_index(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_fetch_stat_numbackends github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_fetch_stat_numbackends
func F_pgstat_fetch_stat_numbackends(m *base.Module) int32
//go:linkname F_pgstat_report_stat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_report_stat
func F_pgstat_report_stat(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_reset github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_reset
func F_pgstat_reset(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_pgstat_reset_of_kind github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_reset_of_kind
func F_pgstat_reset_of_kind(m *base.Module, l0 int32)
//go:linkname F_pgstat_snapshot_fixed github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_snapshot_fixed
func F_pgstat_snapshot_fixed(m *base.Module, l0 int32)
//go:linkname F_pgstat_fetch_pending_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_fetch_pending_entry
func F_pgstat_fetch_pending_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pgstat_count_backend_io_op github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_count_backend_io_op
func F_pgstat_count_backend_io_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64)
//go:linkname F_pgstat_fetch_stat_bgwriter github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_fetch_stat_bgwriter
func F_pgstat_fetch_stat_bgwriter(m *base.Module) int32
//go:linkname F_pgstat_report_checkpointer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_report_checkpointer
func F_pgstat_report_checkpointer(m *base.Module)
//go:linkname F_pgstat_prepare_report_checksum_failure github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_prepare_report_checksum_failure
func F_pgstat_prepare_report_checksum_failure(m *base.Module, l0 int32)
//go:linkname F_pgstat_report_checksum_failures_in_db github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_report_checksum_failures_in_db
func F_pgstat_report_checksum_failures_in_db(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_report_tempfile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_report_tempfile
func F_pgstat_report_tempfile(m *base.Module, l0 int32)
//go:linkname F_pgstat_report_connect github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_report_connect
func F_pgstat_report_connect(m *base.Module)
//go:linkname F_pgstat_fetch_stat_dbentry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_fetch_stat_dbentry
func F_pgstat_fetch_stat_dbentry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_flush_io github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_flush_io
func F_pgstat_flush_io(m *base.Module, l0 int32)
//go:linkname F_pgstat_assoc_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_assoc_relation
func F_pgstat_assoc_relation(m *base.Module, l0 int32)
//go:linkname F_pgstat_count_heap_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_count_heap_delete
func F_pgstat_count_heap_delete(m *base.Module, l0 int32)
//go:linkname F_pgstat_fetch_stat_tabentry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_fetch_stat_tabentry
func F_pgstat_fetch_stat_tabentry(m *base.Module, l0 int32) int32
//go:linkname F_StatsShmemSize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_StatsShmemSize
func F_StatsShmemSize(m *base.Module) int32
//go:linkname F_pgstat_release_entry_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_release_entry_ref
func F_pgstat_release_entry_ref(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pgstat_get_entry_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_get_entry_ref
func F_pgstat_get_entry_ref(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32
//go:linkname F_pgstat_unlock_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_unlock_entry
func F_pgstat_unlock_entry(m *base.Module, l0 int32)
//go:linkname F_pgstat_drop_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_drop_entry
func F_pgstat_drop_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pgstat_report_subscription_error github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_report_subscription_error
func F_pgstat_report_subscription_error(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_report_wal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_report_wal
func F_pgstat_report_wal(m *base.Module, l0 int32)
//go:linkname F_AtEOXact_PgStat github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AtEOXact_PgStat
func F_AtEOXact_PgStat(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AtEOSubXact_PgStat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AtEOSubXact_PgStat
func F_AtEOSubXact_PgStat(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_get_xact_stack_level github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_get_xact_stack_level
func F_pgstat_get_xact_stack_level(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_execute_transactional_drops github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_execute_transactional_drops
func F_pgstat_execute_transactional_drops(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_drop_transactional github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_drop_transactional
func F_pgstat_drop_transactional(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_check_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_acl
func F_check_acl(m *base.Module, l0 int32)
//go:linkname F_acldefault github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_acldefault
func F_acldefault(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_aclnewowner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_aclnewowner
func F_aclnewowner(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_roles_is_member_of github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_roles_is_member_of
func F_roles_is_member_of(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_has_privs_of_role github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_has_privs_of_role
func F_has_privs_of_role(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_convert_any_priv_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_any_priv_string
func F_convert_any_priv_string(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_get_role_oid_or_public github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_role_oid_or_public
func F_get_role_oid_or_public(m *base.Module, l0 int32) int32
//go:linkname F_check_can_set_role github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_can_set_role
func F_check_can_set_role(m *base.Module, l0 int32, l1 int32)
//go:linkname F_DatumGetAnyArrayP github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DatumGetAnyArrayP
func F_DatumGetAnyArrayP(m *base.Module, l0 int32) int32
//go:linkname F_fetch_array_arg_replace_nulls github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fetch_array_arg_replace_nulls
func F_fetch_array_arg_replace_nulls(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_array_shuffle_n github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_shuffle_n
func F_array_shuffle_n(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_CopyArrayEls github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CopyArrayEls
func F_CopyArrayEls(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_construct_empty_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_empty_array
func F_construct_empty_array(m *base.Module, l0 int32) int32
//go:linkname F_array_iter_next github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_iter_next
func F_array_iter_next(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_array_send github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_send
func F_array_send(m *base.Module, l0 int32) int32
//go:linkname F_array_get_element github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_get_element
func F_array_get_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_array_seek github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_seek
func F_array_seek(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_array_get_slice github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_array_get_slice
func F_array_get_slice(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_array_set_element github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_array_set_element
func F_array_set_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_construct_md_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_md_array
func F_construct_md_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_deconstruct_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_array
func F_deconstruct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_array_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_ref
func F_array_ref(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_construct_array github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_construct_array
func F_construct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_construct_array_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_array_builtin
func F_construct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_deconstruct_array_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_array_builtin
func F_deconstruct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_array_contains_nulls github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_contains_nulls
func F_array_contains_nulls(m *base.Module, l0 int32) int32
//go:linkname F_array_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_cmp
func F_array_cmp(m *base.Module, l0 int32) int32
//go:linkname F_initArrayResultWithSize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_initArrayResultWithSize
func F_initArrayResultWithSize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_accumArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_accumArrayResult
func F_accumArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeArrayResult
func F_makeArrayResult(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeMdArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeMdArrayResult
func F_makeMdArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_makeArrayResultArr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeArrayResultArr
func F_makeArrayResultArr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_array_replace_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_replace_internal
func F_array_replace_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_ArrayGetNItems github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ArrayGetNItems
func F_ArrayGetNItems(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ArrayCheckBounds github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ArrayCheckBounds
func F_ArrayCheckBounds(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ArrayGetIntegerTypmods github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ArrayGetIntegerTypmods
func F_ArrayGetIntegerTypmods(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_bool_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_bool_with_len
func F_parse_bool_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EncodeSpecialDate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_EncodeSpecialDate
func F_EncodeSpecialDate(m *base.Module, l0 int32, l1 int32)
//go:linkname F_date_send github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_date_send
func F_date_send(m *base.Module, l0 int32) int32
//go:linkname F_date2j github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_date2j
func F_date2j(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DecodeDateTime github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_DecodeDateTime
func F_DecodeDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_DetermineTimeZoneOffsetInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DetermineTimeZoneOffsetInternal
func F_DetermineTimeZoneOffsetInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DetermineTimeZoneOffset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DetermineTimeZoneOffset
func F_DetermineTimeZoneOffset(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DecodeTimezoneNameToTz github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecodeTimezoneNameToTz
func F_DecodeTimezoneNameToTz(m *base.Module, l0 int32) int32
//go:linkname F_EncodeDateTime github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_EncodeDateTime
func F_EncodeDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_datumCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_datumCopy
func F_datumCopy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datum_image_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_datum_image_eq
func F_datum_image_eq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_datumRestore github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_datumRestore
func F_datumRestore(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_calculate_tablespace_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_calculate_tablespace_size
func F_calculate_tablespace_size(m *base.Module, l0 int32) int64
//go:linkname F_calculate_indexes_size github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_calculate_indexes_size
func F_calculate_indexes_size(m *base.Module, l0 int32) int64
//go:linkname F_errdatatype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errdatatype
func F_errdatatype(m *base.Module, l0 int32)
//go:linkname F_hex_decode_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hex_decode_safe
func F_hex_decode_safe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64
//go:linkname F_enum_cmp_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_enum_cmp_internal
func F_enum_cmp_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_enum_endpoint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_enum_endpoint
func F_enum_endpoint(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_EOH_get_flat_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EOH_get_flat_size
func F_EOH_get_flat_size(m *base.Module, l0 int32) int32
//go:linkname F_DeleteExpandedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DeleteExpandedObject
func F_DeleteExpandedObject(m *base.Module, l0 int32)
//go:linkname F_make_expanded_record_from_typeid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_expanded_record_from_typeid
func F_make_expanded_record_from_typeid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_expanded_record_from_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_expanded_record_from_tupdesc
func F_make_expanded_record_from_tupdesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expanded_record_fetch_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expanded_record_fetch_tupdesc
func F_expanded_record_fetch_tupdesc(m *base.Module, l0 int32) int32
//go:linkname F_expanded_record_set_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_expanded_record_set_tuple
func F_expanded_record_set_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_deconstruct_expanded_record github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_deconstruct_expanded_record
func F_deconstruct_expanded_record(m *base.Module, l0 int32)
//go:linkname F_expanded_record_lookup_field github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expanded_record_lookup_field
func F_expanded_record_lookup_field(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_expanded_record_fetch_field github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expanded_record_fetch_field
func F_expanded_record_fetch_field(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_float_underflow_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_float_underflow_error
func F_float_underflow_error(m *base.Module)
//go:linkname F_float8in_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_float8in_internal
func F_float8in_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_format_type_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_format_type_extended
func F_format_type_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_format_type_be github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_format_type_be
func F_format_type_be(m *base.Module, l0 int32) int32
//go:linkname F_str_toupper github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_str_toupper
func F_str_toupper(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_str_initcap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_str_initcap
func F_str_initcap(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datetime_to_char_body github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_datetime_to_char_body
func F_datetime_to_char_body(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_DCH_cache_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DCH_cache_fetch
func F_DCH_cache_fetch(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_format github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_format
func F_parse_format(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_from_char_parse_int_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_from_char_parse_int_len
func F_from_char_parse_int_len(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_NUM_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_NUM_cache
func F_NUM_cache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_NUM_processor github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_NUM_processor
func F_NUM_processor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_get_th github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_th
func F_get_th(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_int_to_roman github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_int_to_roman
func F_int_to_roman(m *base.Module, l0 int32) int32
//go:linkname F_pg_read_file_common github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_read_file_common
func F_pg_read_file_common(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32
//go:linkname F_pg_ls_dir_files github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_ls_dir_files
func F_pg_ls_dir_files(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_path_decode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_path_decode
func F_path_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_pair_decode github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pair_decode
func F_pair_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_path_encode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_path_encode
func F_path_encode(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_point_dt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_point_dt
func F_point_dt(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_line_interpt_line github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_line_interpt_line
func F_line_interpt_line(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lseg_interpt_lseg github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lseg_interpt_lseg
func F_lseg_interpt_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lseg_interpt_line github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_lseg_interpt_line
func F_lseg_interpt_line(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lseg_closept_point github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lseg_closept_point
func F_lseg_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_line_closept_point github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_line_closept_point
func F_line_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_box_interpt_lseg github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_box_interpt_lseg
func F_box_interpt_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dist_ppoly_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dist_ppoly_internal
func F_dist_ppoly_internal(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_point_inside github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_point_inside
func F_point_inside(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_line_contain_point github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_line_contain_point
func F_line_contain_point(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lseg_inside_poly github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lseg_inside_poly
func F_lseg_inside_poly(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_point_div_point github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_point_div_point
func F_point_div_point(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_int8inc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_int8inc
func F_int8inc(m *base.Module, l0 int32) int32
//go:linkname F_json_send github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_json_send
func F_json_send(m *base.Module, l0 int32) int32
//go:linkname F_JsonEncodeDateTime github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_JsonEncodeDateTime
func F_JsonEncodeDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_array_dim_to_json github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_dim_to_json
func F_array_dim_to_json(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_composite_to_json github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_composite_to_json
func F_composite_to_json(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_escape_json_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_escape_json_with_len
func F_escape_json_with_len(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_JsonbToCString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonbToCString
func F_JsonbToCString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_JsonbExtractScalar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonbExtractScalar
func F_JsonbExtractScalar(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_add_jsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_jsonb
func F_add_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_cannotCastJsonbValue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cannotCastJsonbValue
func F_cannotCastJsonbValue(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pushJsonbValue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pushJsonbValue
func F_pushJsonbValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_appendElement github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_appendElement
func F_appendElement(m *base.Module, l0 int32, l1 int32)
//go:linkname F_convertJsonbValue github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_convertJsonbValue
func F_convertJsonbValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_JsonbIteratorNext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonbIteratorNext
func F_JsonbIteratorNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_compareJsonbContainers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_compareJsonbContainers
func F_compareJsonbContainers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonbIteratorInit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_JsonbIteratorInit
func F_JsonbIteratorInit(m *base.Module, l0 int32) int32
//go:linkname F_getKeyJsonValueFromContainer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getKeyJsonValueFromContainer
func F_getKeyJsonValueFromContainer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getIthJsonbValueFromContainer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getIthJsonbValueFromContainer
func F_getIthJsonbValueFromContainer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonbDeepContains github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_JsonbDeepContains
func F_JsonbDeepContains(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_parse_json_or_errsave github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_parse_json_or_errsave
func F_pg_parse_json_or_errsave(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_json_errsave_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_json_errsave_error
func F_json_errsave_error(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_makeJsonLexContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeJsonLexContext
func F_makeJsonLexContext(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_JsonbValueAsText github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_JsonbValueAsText
func F_JsonbValueAsText(m *base.Module, l0 int32) int32
//go:linkname F_setPath github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_setPath
func F_setPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_each_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_each_worker
func F_each_worker(m *base.Module, l0 int32, l1 int32)
//go:linkname F_elements_worker_jsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_elements_worker_jsonb
func F_elements_worker_jsonb(m *base.Module, l0 int32, l1 int32)
//go:linkname F_populate_record_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_populate_record_worker
func F_populate_record_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_populate_record_field github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_populate_record_field
func F_populate_record_field(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_populate_recordset_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_populate_recordset_worker
func F_populate_recordset_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_populate_recordset_record github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_populate_recordset_record
func F_populate_recordset_record(m *base.Module, l0 int32, l1 int32)
//go:linkname F_json_categorize_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_json_categorize_type
func F_json_categorize_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_populate_array_report_expected_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_populate_array_report_expected_array
func F_populate_array_report_expected_array(m *base.Module, l0 int32, l1 int32)
//go:linkname F_populate_array_check_dimension github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_populate_array_check_dimension
func F_populate_array_check_dimension(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_jspOperationName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jspOperationName
func F_jspOperationName(m *base.Module, l0 int32) int32
//go:linkname F_jspGetArg github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_jspGetArg
func F_jspGetArg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_jspGetNext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_jspGetNext
func F_jspGetNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_jspGetRightArg github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_jspGetRightArg
func F_jspGetRightArg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetJsonTableExecContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetJsonTableExecContext
func F_GetJsonTableExecContext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonTablePlanNextRow github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_JsonTablePlanNextRow
func F_JsonTablePlanNextRow(m *base.Module, l0 int32) int32
//go:linkname F_jsonb_path_match_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_jsonb_path_match_internal
func F_jsonb_path_match_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_executeItemOptUnwrapTarget github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_executeItemOptUnwrapTarget
func F_executeItemOptUnwrapTarget(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_executeNextItem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_executeNextItem
func F_executeNextItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_executeItemOptUnwrapResult github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_executeItemOptUnwrapResult
func F_executeItemOptUnwrapResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SB_MatchText github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SB_MatchText
func F_SB_MatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_UTF8_MatchText github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_UTF8_MatchText
func F_UTF8_MatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_MB_MatchText github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MB_MatchText
func F_MB_MatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SB_IMatchText github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SB_IMatchText
func F_SB_IMatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_like_regex_support github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_like_regex_support
func F_like_regex_support(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pattern_fixed_prefix github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pattern_fixed_prefix
func F_pattern_fixed_prefix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_make_greater_string github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_greater_string
func F_make_greater_string(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_count_nulls github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_count_nulls
func F_count_nulls(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_input_is_valid_common github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_input_is_valid_common
func F_pg_input_is_valid_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_current_logfile github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_current_logfile
func F_pg_current_logfile(m *base.Module, l0 int32) int32
//go:linkname F_make_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_multirange
func F_make_multirange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_multirange_get_bounds github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_multirange_get_bounds
func F_multirange_get_bounds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_multirange_contains_elem_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_multirange_contains_elem_internal
func F_multirange_contains_elem_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_contains_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_contains_multirange_internal
func F_range_contains_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_position github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_position
func F_get_position(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64
//go:linkname F_namestrcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_namestrcpy
func F_namestrcpy(m *base.Module, l0 int32, l1 int32)
//go:linkname F_network_in github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_network_in
func F_network_in(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_inet_hist_value_sel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inet_hist_value_sel
func F_inet_hist_value_sel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64
//go:linkname F_networkjoinsel_semi github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_networkjoinsel_semi
func F_networkjoinsel_semi(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64
//go:linkname F_make_result_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_result_opt_error
func F_make_result_opt_error(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_numeric_out_sci github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_numeric_out_sci
func F_numeric_out_sci(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_div_var github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_div_var
func F_div_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_cmp_abs_common github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cmp_abs_common
func F_cmp_abs_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_sub_var github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sub_var
func F_sub_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_numeric_add_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_numeric_add_opt_error
func F_numeric_add_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_numeric_sub_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_numeric_sub_opt_error
func F_numeric_sub_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_numeric_div_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_numeric_div_opt_error
func F_numeric_div_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sqrt_var github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sqrt_var
func F_sqrt_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_estimate_ln_dweight github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_estimate_ln_dweight
func F_estimate_ln_dweight(m *base.Module, l0 int32) int32
//go:linkname F_ln_var github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ln_var
func F_ln_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_int64_to_numeric github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_int64_to_numeric
func F_int64_to_numeric(m *base.Module, l0 int64) int32
//go:linkname F_numeric_int4_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_numeric_int4_opt_error
func F_numeric_int4_opt_error(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_accum_sum_final github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_accum_sum_final
func F_accum_sum_final(m *base.Module, l0 int32, l1 int32)
//go:linkname F_numeric_poly_stddev_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_numeric_poly_stddev_internal
func F_numeric_poly_stddev_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_strtoint32 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_strtoint32
func F_pg_strtoint32(m *base.Module, l0 int32) int32
//go:linkname F_pg_strtoint64_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_strtoint64_safe
func F_pg_strtoint64_safe(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_pg_ultostr_zeropad github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_ultostr_zeropad
func F_pg_ultostr_zeropad(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_valid_oidvector github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_valid_oidvector
func F_check_valid_oidvector(m *base.Module, l0 int32)
//go:linkname F_dotrim github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dotrim
func F_dotrim(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ordered_set_startup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ordered_set_startup
func F_ordered_set_startup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_setup_pct_info github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_setup_pct_info
func F_setup_pct_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_cache_locale_time github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cache_locale_time
func F_cache_locale_time(m *base.Module)
//go:linkname F_pg_newlocale_from_collation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_newlocale_from_collation
func F_pg_newlocale_from_collation(m *base.Module, l0 int32) int32
//go:linkname F_pg_strlower github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_strlower
func F_pg_strlower(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_strlower_libc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strlower_libc
func F_strlower_libc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_report_newlocale_failure github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_newlocale_failure
func F_report_newlocale_failure(m *base.Module, l0 int32)
//go:linkname F_quote_literal_cstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_quote_literal_cstr
func F_quote_literal_cstr(m *base.Module, l0 int32) int32
//go:linkname F_get_range_io_data github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_range_io_data
func F_get_range_io_data(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_serialize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_serialize
func F_range_serialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_range_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_deserialize
func F_range_deserialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_range_cmp_bounds github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_cmp_bounds
func F_range_cmp_bounds(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_before_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_range_before_internal
func F_range_before_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_bounds_adjacent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bounds_adjacent
func F_bounds_adjacent(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overlaps_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_overlaps_internal
func F_range_overlaps_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overleft_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_overleft_internal
func F_range_overleft_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_intersect_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_intersect_internal
func F_range_intersect_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RE_compile_and_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RE_compile_and_cache
func F_RE_compile_and_cache(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RE_wchar_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RE_wchar_execute
func F_RE_wchar_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_parse_re_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parse_re_flags
func F_parse_re_flags(m *base.Module, l0 int32, l1 int32)
//go:linkname F_regexp_substr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_regexp_substr
func F_regexp_substr(m *base.Module, l0 int32) int32
//go:linkname F_stringToQualifiedNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_stringToQualifiedNameList
func F_stringToQualifiedNameList(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_format_operator_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_format_operator_extended
func F_format_operator_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ri_CheckTrigger github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ri_CheckTrigger
func F_ri_CheckTrigger(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RI_FKey_check github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RI_FKey_check
func F_RI_FKey_check(m *base.Module, l0 int32)
//go:linkname F_ri_FetchConstraintInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ri_FetchConstraintInfo
func F_ri_FetchConstraintInfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ri_FetchPreparedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ri_FetchPreparedPlan
func F_ri_FetchPreparedPlan(m *base.Module, l0 int32) int32
//go:linkname F_ri_PlanCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ri_PlanCheck
func F_ri_PlanCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ri_PerformCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ri_PerformCheck
func F_ri_PerformCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_ri_restrict github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ri_restrict
func F_ri_restrict(m *base.Module, l0 int32, l1 int32)
//go:linkname F_record_image_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_record_image_cmp
func F_record_image_cmp(m *base.Module, l0 int32) int32
//go:linkname F_quote_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_quote_identifier
func F_quote_identifier(m *base.Module, l0 int32) int32
//go:linkname F_generate_qualified_relation_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_generate_qualified_relation_name
func F_generate_qualified_relation_name(m *base.Module, l0 int32) int32
//go:linkname F_get_rule_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_rule_expr
func F_get_rule_expr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_fastgetattr_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fastgetattr_3
func F_fastgetattr_3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_set_simple_column_names github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_simple_column_names
func F_set_simple_column_names(m *base.Module, l0 int32)
//go:linkname F_generate_function_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_function_name
func F_generate_function_name(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_pg_get_indexdef_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_get_indexdef_worker
func F_pg_get_indexdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_flatten_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_flatten_reloptions
func F_flatten_reloptions(m *base.Module, l0 int32) int32
//go:linkname F_appendContextKeyword github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_appendContextKeyword
func F_appendContextKeyword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_decompile_column_index_array github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_decompile_column_index_array
func F_decompile_column_index_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_get_expr_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_get_expr_worker
func F_pg_get_expr_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_print_function_arguments github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_print_function_arguments
func F_print_function_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_print_function_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_print_function_rettype
func F_print_function_rettype(m *base.Module, l0 int32, l1 int32)
//go:linkname F_print_function_sqlbody github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_print_function_sqlbody
func F_print_function_sqlbody(m *base.Module, l0 int32, l1 int32)
//go:linkname F_quote_qualified_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_quote_qualified_identifier
func F_quote_qualified_identifier(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_relation_column_names github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_relation_column_names
func F_set_relation_column_names(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_add_cast_to github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_cast_to
func F_add_cast_to(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_const_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_const_expr
func F_get_const_expr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_has_dangerous_join_using github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_has_dangerous_join_using
func F_has_dangerous_join_using(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_build_colinfo_names_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_build_colinfo_names_hash
func F_build_colinfo_names_hash(m *base.Module, l0 int32)
//go:linkname F_make_colname_unique github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_make_colname_unique
func F_make_colname_unique(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_agg_expr_helper github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_agg_expr_helper
func F_get_agg_expr_helper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_windowfunc_expr_helper github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_windowfunc_expr_helper
func F_get_windowfunc_expr_helper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_get_json_returning github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_json_returning
func F_get_json_returning(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_json_table_columns github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_json_table_columns
func F_get_json_table_columns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_var_eq_const github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_var_eq_const
func F_var_eq_const(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) float64
//go:linkname F_mcv_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_mcv_selectivity
func F_mcv_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64
//go:linkname F_examine_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_examine_variable
func F_examine_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ineq_histogram_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ineq_histogram_selectivity
func F_ineq_histogram_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) float64
//go:linkname F_scalarineqsel_wrapper github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scalarineqsel_wrapper
func F_scalarineqsel_wrapper(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_estimate_array_length github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_estimate_array_length
func F_estimate_array_length(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_get_join_variables github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_join_variables
func F_get_join_variables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_add_unique_group_var github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_add_unique_group_var
func F_add_unique_group_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_estimate_multivariate_ndistinct github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_estimate_multivariate_ndistinct
func F_estimate_multivariate_ndistinct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_index_other_operands_eval_cost github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_other_operands_eval_cost
func F_index_other_operands_eval_cost(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_anytimestamp_typmod_check github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_anytimestamp_typmod_check
func F_anytimestamp_typmod_check(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SetEpochTimestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SetEpochTimestamp
func F_SetEpochTimestamp(m *base.Module) int64
//go:linkname F_AdjustTimestampForTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AdjustTimestampForTypmod
func F_AdjustTimestampForTypmod(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_timestamp2tm github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_timestamp2tm
func F_timestamp2tm(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_timestamp2timestamptz_opt_overflow github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamp2timestamptz_opt_overflow
func F_timestamp2timestamptz_opt_overflow(m *base.Module, l0 int64, l1 int32) int64
//go:linkname F_AdjustIntervalForTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AdjustIntervalForTypmod
func F_AdjustIntervalForTypmod(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_timestamp_cmp_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamp_cmp_internal
func F_timestamp_cmp_internal(m *base.Module, l0 int64, l1 int64) int32
//go:linkname F_timestamptz_pl_interval_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamptz_pl_interval_internal
func F_timestamptz_pl_interval_internal(m *base.Module, l0 int64, l1 int32, l2 int32) int64
//go:linkname F_finite_interval_pl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_finite_interval_pl
func F_finite_interval_pl(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_finite_interval_mi github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_finite_interval_mi
func F_finite_interval_mi(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_timestamptz_trunc_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamptz_trunc_internal
func F_timestamptz_trunc_internal(m *base.Module, l0 int32, l1 int64, l2 int32) int64
//go:linkname F_NonFiniteTimestampTzPart github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_NonFiniteTimestampTzPart
func F_NonFiniteTimestampTzPart(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_pushOperator github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pushOperator
func F_pushOperator(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pushValue github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pushValue
func F_pushValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_pushStop github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pushStop
func F_pushStop(m *base.Module, l0 int32)
//go:linkname F_parse_tsquery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_parse_tsquery
func F_parse_tsquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_freetree github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_freetree
func F_freetree(m *base.Module, l0 int32)
//go:linkname F_CompareTSQ github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CompareTSQ
func F_CompareTSQ(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getWeights github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getWeights
func F_getWeights(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SortAndUniqItems github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SortAndUniqItems
func F_SortAndUniqItems(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tsCompareString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tsCompareString
func F_tsCompareString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_TS_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TS_execute
func F_TS_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_TS_execute_ternary github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TS_execute_ternary
func F_TS_execute_ternary(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tsquery_requires_match github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tsquery_requires_match
func F_tsquery_requires_match(m *base.Module, l0 int32) int32
//go:linkname F_ts_stat_sql github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ts_stat_sql
func F_ts_stat_sql(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ts_setup_firstcall github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ts_setup_firstcall
func F_ts_setup_firstcall(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ts_process_call github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ts_process_call
func F_ts_process_call(m *base.Module, l0 int32) int32
//go:linkname F_varbit_out github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_varbit_out
func F_varbit_out(m *base.Module, l0 int32) int32
//go:linkname F_cstring_to_text github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cstring_to_text
func F_cstring_to_text(m *base.Module, l0 int32) int32
//go:linkname F_cstring_to_text_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cstring_to_text_with_len
func F_cstring_to_text_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_text_to_cstring_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_text_to_cstring_buffer
func F_text_to_cstring_buffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_text_catenate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text_catenate
func F_text_catenate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_text_position_setup github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text_position_setup
func F_text_position_setup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_text_position_next github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text_position_next
func F_text_position_next(m *base.Module, l0 int32) int32
//go:linkname F_varstr_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_varstr_cmp
func F_varstr_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_bytea_substring github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bytea_substring
func F_bytea_substring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_textToQualifiedNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_textToQualifiedNameList
func F_textToQualifiedNameList(m *base.Module, l0 int32) int32
//go:linkname F_SplitDirectoriesString github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SplitDirectoriesString
func F_SplitDirectoriesString(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SplitGUCList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SplitGUCList
func F_SplitGUCList(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replace_text_regexp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_replace_text_regexp
func F_replace_text_regexp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_array_to_text_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_to_text_internal
func F_array_to_text_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_varstr_levenshtein_less_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_varstr_levenshtein_less_equal
func F_varstr_levenshtein_less_equal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_unicode_norm_form_from_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_unicode_norm_form_from_string
func F_unicode_norm_form_from_string(m *base.Module, l0 int32) int32
//go:linkname F_map_sql_identifier_to_xml_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_map_sql_identifier_to_xml_name
func F_map_sql_identifier_to_xml_name(m *base.Module) int32
//go:linkname F_SPI_sql_row_to_xmlelement github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_sql_row_to_xmlelement
func F_SPI_sql_row_to_xmlelement(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_schema_to_xmlschema_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_schema_to_xmlschema_internal
func F_schema_to_xmlschema_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CatCacheRemoveCList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CatCacheRemoveCList
func F_CatCacheRemoveCList(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateCacheMemoryContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateCacheMemoryContext
func F_CreateCacheMemoryContext(m *base.Module)
//go:linkname F_CatalogCacheInitializeCache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CatalogCacheInitializeCache
func F_CatalogCacheInitializeCache(m *base.Module, l0 int32)
//go:linkname F_SearchCatCacheInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SearchCatCacheInternal
func F_SearchCatCacheInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ReleaseCatCache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReleaseCatCache
func F_ReleaseCatCache(m *base.Module, l0 int32)
//go:linkname F_CatalogCacheComputeTupleHashValue github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CatalogCacheComputeTupleHashValue
func F_CatalogCacheComputeTupleHashValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CatalogCacheCreateEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CatalogCacheCreateEntry
func F_CatalogCacheCreateEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ReleaseCatCacheList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReleaseCatCacheList
func F_ReleaseCatCacheList(m *base.Module, l0 int32)
//go:linkname F_AtEOXact_Inval github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AtEOXact_Inval
func F_AtEOXact_Inval(m *base.Module, l0 int32)
//go:linkname F_AtEOSubXact_Inval github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AtEOSubXact_Inval
func F_AtEOSubXact_Inval(m *base.Module, l0 int32)
//go:linkname F_CacheInvalidateHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CacheInvalidateHeapTuple
func F_CacheInvalidateHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PrepareInvalidationState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PrepareInvalidationState
func F_PrepareInvalidationState(m *base.Module) int32
//go:linkname F_CacheInvalidateHeapTupleCommon github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CacheInvalidateHeapTupleCommon
func F_CacheInvalidateHeapTupleCommon(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CacheInvalidateRelcacheByTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CacheInvalidateRelcacheByTuple
func F_CacheInvalidateRelcacheByTuple(m *base.Module, l0 int32)
//go:linkname F_CacheRegisterSyscacheCallback github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CacheRegisterSyscacheCallback
func F_CacheRegisterSyscacheCallback(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_op_in_opfamily github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_op_in_opfamily
func F_op_in_opfamily(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_op_opfamily_properties github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_op_opfamily_properties
func F_get_op_opfamily_properties(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_opfamily_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_opfamily_member
func F_get_opfamily_member(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_opfamily_member_for_cmptype github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_opfamily_member_for_cmptype
func F_get_opfamily_member_for_cmptype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_ordering_op_properties github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_ordering_op_properties
func F_get_ordering_op_properties(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_ordering_op_for_equality_op github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_ordering_op_for_equality_op
func F_get_ordering_op_for_equality_op(m *base.Module, l0 int32) int32
//go:linkname F_get_op_hash_functions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_op_hash_functions
func F_get_op_hash_functions(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_opfamily_proc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opfamily_proc
func F_get_opfamily_proc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_negator github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_negator
func F_get_negator(m *base.Module, l0 int32) int32
//go:linkname F_get_attname github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_attname
func F_get_attname(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_atttype github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_atttype
func F_get_atttype(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_collation_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_collation_name
func F_get_collation_name(m *base.Module, l0 int32) int32
//go:linkname F_get_constraint_index github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_constraint_index
func F_get_constraint_index(m *base.Module, l0 int32) int32
//go:linkname F_get_constraint_type github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_constraint_type
func F_get_constraint_type(m *base.Module, l0 int32) int32
//go:linkname F_get_language_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_language_name
func F_get_language_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_opclass_opfamily_and_input_type github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_opclass_opfamily_and_input_type
func F_get_opclass_opfamily_and_input_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_opcode github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opcode
func F_get_opcode(m *base.Module, l0 int32) int32
//go:linkname F_op_input_types github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_op_input_types
func F_op_input_types(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_op_strict github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_op_strict
func F_op_strict(m *base.Module, l0 int32) int32
//go:linkname F_func_volatile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_func_volatile
func F_func_volatile(m *base.Module, l0 int32) int32
//go:linkname F_get_func_name github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_name
func F_get_func_name(m *base.Module, l0 int32) int32
//go:linkname F_get_func_namespace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_func_namespace
func F_get_func_namespace(m *base.Module, l0 int32) int32
//go:linkname F_get_func_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_rettype
func F_get_func_rettype(m *base.Module, l0 int32) int32
//go:linkname F_get_func_retset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_retset
func F_get_func_retset(m *base.Module, l0 int32) int32
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
//go:linkname F_get_rel_type_id github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_rel_type_id
func F_get_rel_type_id(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_relkind github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_rel_relkind
func F_get_rel_relkind(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_tablespace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_rel_tablespace
func F_get_rel_tablespace(m *base.Module, l0 int32) int32
//go:linkname F_get_typlen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_typlen
func F_get_typlen(m *base.Module, l0 int32) int32
//go:linkname F_get_typlenbyval github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_typlenbyval
func F_get_typlenbyval(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_type_io_data github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_type_io_data
func F_get_type_io_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_getBaseType github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getBaseType
func F_getBaseType(m *base.Module, l0 int32) int32
//go:linkname F_getBaseTypeAndTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getBaseTypeAndTypmod
func F_getBaseTypeAndTypmod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_typavgwidth github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_typavgwidth
func F_get_typavgwidth(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_type_is_range github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_type_is_range
func F_type_is_range(m *base.Module, l0 int32) int32
//go:linkname F_get_element_type github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_element_type
func F_get_element_type(m *base.Module, l0 int32) int32
//go:linkname F_get_array_type github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_array_type
func F_get_array_type(m *base.Module, l0 int32) int32
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
//go:linkname F_type_is_collatable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_type_is_collatable
func F_type_is_collatable(m *base.Module, l0 int32) int32
//go:linkname F_getSubscriptingRoutines github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_getSubscriptingRoutines
func F_getSubscriptingRoutines(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_attstatsslot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_attstatsslot
func F_get_attstatsslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_free_attstatsslot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_free_attstatsslot
func F_free_attstatsslot(m *base.Module, l0 int32)
//go:linkname F_get_namespace_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_namespace_name
func F_get_namespace_name(m *base.Module, l0 int32) int32
//go:linkname F_get_namespace_name_or_temp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_namespace_name_or_temp
func F_get_namespace_name_or_temp(m *base.Module, l0 int32) int32
//go:linkname F_get_range_subtype github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_range_subtype
func F_get_range_subtype(m *base.Module, l0 int32) int32
//go:linkname F_get_multirange_range github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_multirange_range
func F_get_multirange_range(m *base.Module, l0 int32) int32
//go:linkname F_get_index_column_opclass github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_index_column_opclass
func F_get_index_column_opclass(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_index_isreplident github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_index_isreplident
func F_get_index_isreplident(m *base.Module, l0 int32) int32
//go:linkname F_get_index_isvalid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_index_isvalid
func F_get_index_isvalid(m *base.Module, l0 int32) int32
//go:linkname F_get_index_isclustered github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_index_isclustered
func F_get_index_isclustered(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetPartitionQual github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetPartitionQual
func F_RelationGetPartitionQual(m *base.Module, l0 int32) int32
//go:linkname F_generate_partition_qual github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_partition_qual
func F_generate_partition_qual(m *base.Module, l0 int32) int32
//go:linkname F_DropCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DropCachedPlan
func F_DropCachedPlan(m *base.Module, l0 int32)
//go:linkname F_GetCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetCachedPlan
func F_GetCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_RelationGetIndexAttOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetIndexAttOptions
func F_RelationGetIndexAttOptions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationClearRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationClearRelation
func F_RelationClearRelation(m *base.Module, l0 int32)
//go:linkname F_RelationCacheInvalidate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationCacheInvalidate
func F_RelationCacheInvalidate(m *base.Module)
//go:linkname F_AtEOSubXact_RelationCache github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AtEOSubXact_RelationCache
func F_AtEOSubXact_RelationCache(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RelationSetNewRelfilenumber github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationSetNewRelfilenumber
func F_RelationSetNewRelfilenumber(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationGetIndexList github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationGetIndexList
func F_RelationGetIndexList(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexExpressions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetIndexExpressions
func F_RelationGetIndexExpressions(m *base.Module, l0 int32) int32
//go:linkname F_errtable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_errtable
func F_errtable(m *base.Module, l0 int32)
//go:linkname F_errtableconstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errtableconstraint
func F_errtableconstraint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationCacheInitFilePostInvalidate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationCacheInitFilePostInvalidate
func F_RelationCacheInitFilePostInvalidate(m *base.Module)
//go:linkname F_read_relmap_file github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_read_relmap_file
func F_read_relmap_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_AtEOXact_RelationMap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AtEOXact_RelationMap
func F_AtEOXact_RelationMap(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_tablespace_page_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_tablespace_page_costs
func F_get_tablespace_page_costs(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_tablespace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_tablespace
func F_get_tablespace(m *base.Module, l0 int32) int32
//go:linkname F_SearchSysCache1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchSysCache1
func F_SearchSysCache1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCache2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCache2
func F_SearchSysCache2(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchSysCache3 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SearchSysCache3
func F_SearchSysCache3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SearchSysCacheCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchSysCacheCopy
func F_SearchSysCacheCopy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchSysCacheLockedCopy1 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCacheLockedCopy1
func F_SearchSysCacheLockedCopy1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheExists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCacheExists
func F_SearchSysCacheExists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_GetSysCacheOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetSysCacheOid
func F_GetSysCacheOid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SearchSysCacheAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SearchSysCacheAttName
func F_SearchSysCacheAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheExistsAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SearchSysCacheExistsAttName
func F_SearchSysCacheExistsAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheCopyAttNum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SearchSysCacheCopyAttNum
func F_SearchSysCacheCopyAttNum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SysCacheGetAttr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SysCacheGetAttr
func F_SysCacheGetAttr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
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
//go:linkname F_lookup_type_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lookup_type_cache
func F_lookup_type_cache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DomainHasConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DomainHasConstraints
func F_DomainHasConstraints(m *base.Module, l0 int32) int32
//go:linkname F_lookup_rowtype_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lookup_rowtype_tupdesc
func F_lookup_rowtype_tupdesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookup_rowtype_tupdesc_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lookup_rowtype_tupdesc_internal
func F_lookup_rowtype_tupdesc_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookup_rowtype_tupdesc_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lookup_rowtype_tupdesc_domain
func F_lookup_rowtype_tupdesc_domain(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_or_make_matching_shared_tupledesc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_find_or_make_matching_shared_tupledesc
func F_find_or_make_matching_shared_tupledesc(m *base.Module, l0 int32) int32
//go:linkname F_errstart_cold github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errstart_cold
func F_errstart_cold(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errstart github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_errstart
func F_errstart(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_write_stderr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_write_stderr
func F_write_stderr(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errfinish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errfinish
func F_errfinish(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_EmitErrorReport github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_EmitErrorReport
func F_EmitErrorReport(m *base.Module)
//go:linkname F_errsave_start github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errsave_start
func F_errsave_start(m *base.Module, l0 int32) int32
//go:linkname F_errcode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errcode
func F_errcode(m *base.Module, l0 int32)
//go:linkname F_errcode_for_file_access github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errcode_for_file_access
func F_errcode_for_file_access(m *base.Module)
//go:linkname F_errcode_for_socket_access github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errcode_for_socket_access
func F_errcode_for_socket_access(m *base.Module)
//go:linkname F_errmsg github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errmsg
func F_errmsg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errmsg_plural github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errmsg_plural
func F_errmsg_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errdetail github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errdetail
func F_errdetail(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errdetail_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errdetail_internal
func F_errdetail_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errdetail_log github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errdetail_log
func F_errdetail_log(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errhint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errhint
func F_errhint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errcontext_msg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errcontext_msg
func F_errcontext_msg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_set_errcontext_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_errcontext_domain
func F_set_errcontext_domain(m *base.Module, l0 int32)
//go:linkname F_errhidestmt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errhidestmt
func F_errhidestmt(m *base.Module)
//go:linkname F_internalerrposition github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_internalerrposition
func F_internalerrposition(m *base.Module, l0 int32)
//go:linkname F_internalerrquery github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_internalerrquery
func F_internalerrquery(m *base.Module, l0 int32) int32
//go:linkname F_err_generic_string github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_err_generic_string
func F_err_generic_string(m *base.Module, l0 int32, l1 int32)
//go:linkname F_geterrcode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_geterrcode
func F_geterrcode(m *base.Module) int32
//go:linkname F_geterrposition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_geterrposition
func F_geterrposition(m *base.Module) int32
//go:linkname F_format_elog_string github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_format_elog_string
func F_format_elog_string(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CopyErrorData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopyErrorData
func F_CopyErrorData(m *base.Module) int32
//go:linkname F_FlushErrorState github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FlushErrorState
func F_FlushErrorState(m *base.Module)
//go:linkname F_ThrowErrorData github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ThrowErrorData
func F_ThrowErrorData(m *base.Module, l0 int32)
//go:linkname F_load_external_function github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_load_external_function
func F_load_external_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_substitute_path_macro github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_substitute_path_macro
func F_substitute_path_macro(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_load_file github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_load_file
func F_load_file(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fmgr_info github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fmgr_info
func F_fmgr_info(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fmgr_info_cxt_security github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fmgr_info_cxt_security
func F_fmgr_info_cxt_security(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_fmgr_info_cxt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fmgr_info_cxt
func F_fmgr_info_cxt(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_fmgr_info_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fmgr_info_copy
func F_fmgr_info_copy(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_DirectFunctionCall2Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DirectFunctionCall2Coll
func F_DirectFunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CallerFInfoFunctionCall2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CallerFInfoFunctionCall2
func F_CallerFInfoFunctionCall2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_FunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FunctionCall1Coll
func F_FunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FunctionCall2Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FunctionCall2Coll
func F_FunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_FunctionCall3Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FunctionCall3Coll
func F_FunctionCall3Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_FunctionCall7Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FunctionCall7Coll
func F_FunctionCall7Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_OidFunctionCall0Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_OidFunctionCall0Coll
func F_OidFunctionCall0Coll(m *base.Module, l0 int32) int32
//go:linkname F_OidFunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_OidFunctionCall1Coll
func F_OidFunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_OidFunctionCall6Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_OidFunctionCall6Coll
func F_OidFunctionCall6Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
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
//go:linkname F_OidReceiveFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OidReceiveFunctionCall
func F_OidReceiveFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_Int64GetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Int64GetDatum
func F_Int64GetDatum(m *base.Module, l0 int64) int32
//go:linkname F_Float8GetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_Float8GetDatum
func F_Float8GetDatum(m *base.Module, l0 float64) int32
//go:linkname F_pg_detoast_datum_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_detoast_datum_copy
func F_pg_detoast_datum_copy(m *base.Module, l0 int32) int32
//go:linkname F_pg_detoast_datum_slice github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_detoast_datum_slice
func F_pg_detoast_datum_slice(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_detoast_datum_packed github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_detoast_datum_packed
func F_pg_detoast_datum_packed(m *base.Module, l0 int32) int32
//go:linkname F_get_fn_expr_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_fn_expr_rettype
func F_get_fn_expr_rettype(m *base.Module, l0 int32) int32
//go:linkname F_get_fn_expr_argtype github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_fn_expr_argtype
func F_get_fn_expr_argtype(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_fn_opclass_options github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_fn_opclass_options
func F_get_fn_opclass_options(m *base.Module, l0 int32) int32
//go:linkname F_InitMaterializedSRF github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InitMaterializedSRF
func F_InitMaterializedSRF(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_call_result_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_call_result_type
func F_get_call_result_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_end_MultiFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_end_MultiFuncCall
func F_end_MultiFuncCall(m *base.Module, l0 int32)
//go:linkname F_get_expr_result_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_expr_result_type
func F_get_expr_result_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_expr_result_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_expr_result_tupdesc
func F_get_expr_result_tupdesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_func_arg_info github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_func_arg_info
func F_get_func_arg_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_func_input_arg_names github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_func_input_arg_names
func F_get_func_input_arg_names(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hash_create github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_create
func F_hash_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_my_log2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_my_log2
func F_my_log2(m *base.Module, l0 int32) int32
//go:linkname F_hash_destroy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hash_destroy
func F_hash_destroy(m *base.Module, l0 int32)
//go:linkname F_get_hash_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_hash_value
func F_get_hash_value(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hash_search github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_search
func F_hash_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_search_with_hash_value github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_search_with_hash_value
func F_hash_search_with_hash_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_hash_seq_init github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_seq_init
func F_hash_seq_init(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hash_seq_search github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_seq_search
func F_hash_seq_search(m *base.Module, l0 int32) int32
//go:linkname F_hash_seq_term github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_seq_term
func F_hash_seq_term(m *base.Module, l0 int32)
//go:linkname F_AtEOXact_HashTables github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AtEOXact_HashTables
func F_AtEOXact_HashTables(m *base.Module, l0 int32)
//go:linkname F_AtEOSubXact_HashTables github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AtEOSubXact_HashTables
func F_AtEOSubXact_HashTables(m *base.Module, l0 int32, l1 int32)
//go:linkname F_InitStandaloneProcess github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InitStandaloneProcess
func F_InitStandaloneProcess(m *base.Module, l0 int32)
//go:linkname F_SwitchBackToLocalLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SwitchBackToLocalLatch
func F_SwitchBackToLocalLatch(m *base.Module)
//go:linkname F_GetBackendTypeDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetBackendTypeDesc
func F_GetBackendTypeDesc(m *base.Module, l0 int32) int32
//go:linkname F_checkDataDir github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_checkDataDir
func F_checkDataDir(m *base.Module)
//go:linkname F_ChangeToDataDir github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ChangeToDataDir
func F_ChangeToDataDir(m *base.Module)
//go:linkname F_has_rolreplication github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_has_rolreplication
func F_has_rolreplication(m *base.Module, l0 int32) int32
//go:linkname F_SetCurrentRoleId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SetCurrentRoleId
func F_SetCurrentRoleId(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateDataDirLockFile github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateDataDirLockFile
func F_CreateDataDirLockFile(m *base.Module, l0 int32)
//go:linkname F_AddToDataDirLockFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AddToDataDirLockFile
func F_AddToDataDirLockFile(m *base.Module, l0 int32, l1 int32)
//go:linkname F_InitializeMaxBackends github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InitializeMaxBackends
func F_InitializeMaxBackends(m *base.Module)
//go:linkname F_BaseInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BaseInit
func F_BaseInit(m *base.Module)
//go:linkname F_InitPostgres github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InitPostgres
func F_InitPostgres(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_SwitchToUntrustedUser github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SwitchToUntrustedUser
func F_SwitchToUntrustedUser(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RestoreUserContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RestoreUserContext
func F_RestoreUserContext(m *base.Module, l0 int32)
//go:linkname F_mic2latin_with_table github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_mic2latin_with_table
func F_mic2latin_with_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_UtfToLocal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UtfToLocal
func F_UtfToLocal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_LocalToUtf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LocalToUtf
func F_LocalToUtf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_PrepareClientEncoding github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PrepareClientEncoding
func F_PrepareClientEncoding(m *base.Module, l0 int32) int32
//go:linkname F_report_invalid_encoding github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_invalid_encoding
func F_report_invalid_encoding(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_client_to_server github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_client_to_server
func F_pg_client_to_server(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_any_to_server github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_any_to_server
func F_pg_any_to_server(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_server_to_any github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_server_to_any
func F_pg_server_to_any(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_unicode_to_server github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_unicode_to_server
func F_pg_unicode_to_server(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_unicode_to_server_noerror github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_unicode_to_server_noerror
func F_pg_unicode_to_server_noerror(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mb2wchar_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_mb2wchar_with_len
func F_pg_mb2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_wchar2mb_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_wchar2mb_with_len
func F_pg_wchar2mb_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_mblen_cstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_mblen_cstr
func F_pg_mblen_cstr(m *base.Module, l0 int32) int32
//go:linkname F_report_invalid_encoding_db github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_report_invalid_encoding_db
func F_report_invalid_encoding_db(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_mblen_unbounded github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_mblen_unbounded
func F_pg_mblen_unbounded(m *base.Module, l0 int32) int32
//go:linkname F_pg_mbcharcliplen github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_mbcharcliplen
func F_pg_mbcharcliplen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_encoding_conversion_args github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_encoding_conversion_args
func F_check_encoding_conversion_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_report_untranslatable_char github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_untranslatable_char
func F_report_untranslatable_char(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_find_option github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_option
func F_find_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_set_config_with_handle github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_set_config_with_handle
func F_set_config_with_handle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_SetConfigOption github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SetConfigOption
func F_SetConfigOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_guc_malloc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_guc_malloc
func F_guc_malloc(m *base.Module, l0 int32) int32
//go:linkname F_InitializeGUCOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_InitializeGUCOptions
func F_InitializeGUCOptions(m *base.Module)
//go:linkname F_call_bool_check_hook github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_call_bool_check_hook
func F_call_bool_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_call_int_check_hook github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_call_int_check_hook
func F_call_int_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_call_string_check_hook github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_call_string_check_hook
func F_call_string_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_call_enum_check_hook github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_call_enum_check_hook
func F_call_enum_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SelectConfigFiles github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SelectConfigFiles
func F_SelectConfigFiles(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_push_old_value github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_push_old_value
func F_push_old_value(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AtEOXact_GUC github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AtEOXact_GUC
func F_AtEOXact_GUC(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BeginReportingGUCOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BeginReportingGUCOptions
func F_BeginReportingGUCOptions(m *base.Module)
//go:linkname F_get_config_unit_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_config_unit_name
func F_get_config_unit_name(m *base.Module, l0 int32) int32
//go:linkname F_parse_int github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_int
func F_parse_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_parse_real github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_parse_real
func F_parse_real(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_config_enum_get_options github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_config_enum_get_options
func F_config_enum_get_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GetConfigOptionFlags github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetConfigOptionFlags
func F_GetConfigOptionFlags(m *base.Module, l0 int32) int32
//go:linkname F_init_custom_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_init_custom_variable
func F_init_custom_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_define_custom_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_define_custom_variable
func F_define_custom_variable(m *base.Module, l0 int32)
//go:linkname F_GetConfigOptionByName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetConfigOptionByName
func F_GetConfigOptionByName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ProcessGUCArray github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ProcessGUCArray
func F_ProcessGUCArray(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GUC_flex_fatal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GUC_flex_fatal
func F_GUC_flex_fatal(m *base.Module, l0 int32)
//go:linkname F_ProcessConfigFile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ProcessConfigFile
func F_ProcessConfigFile(m *base.Module, l0 int32)
//go:linkname F_ParseConfigFp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ParseConfigFp
func F_ParseConfigFp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_DeescapeQuotedString github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DeescapeQuotedString
func F_DeescapeQuotedString(m *base.Module, l0 int32) int32
//go:linkname F_FreeConfigVariables github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeConfigVariables
func F_FreeConfigVariables(m *base.Module, l0 int32)
//go:linkname F_ConfigOptionIsVisible github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ConfigOptionIsVisible
func F_ConfigOptionIsVisible(m *base.Module, l0 int32) int32
//go:linkname F_get_visible_ENR_metadata github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_visible_ENR_metadata
func F_get_visible_ENR_metadata(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ENRMetadataGetTupDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ENRMetadataGetTupDesc
func F_ENRMetadataGetTupDesc(m *base.Module, l0 int32) int32
//go:linkname F_check_stack_depth github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_stack_depth
func F_check_stack_depth(m *base.Module)
//go:linkname F_superuser github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_superuser
func F_superuser(m *base.Module) int32
//go:linkname F_superuser_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_superuser_arg
func F_superuser_arg(m *base.Module, l0 int32) int32
//go:linkname F_enable_timeout github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_enable_timeout
func F_enable_timeout(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32)
//go:linkname F_schedule_alarm github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_schedule_alarm
func F_schedule_alarm(m *base.Module, l0 int64)
//go:linkname F_reschedule_timeouts github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_reschedule_timeouts
func F_reschedule_timeouts(m *base.Module)
//go:linkname F_enable_timeout_after github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_enable_timeout_after
func F_enable_timeout_after(m *base.Module, l0 int32, l1 int32)
//go:linkname F_disable_timeout github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_disable_timeout
func F_disable_timeout(m *base.Module, l0 int32)
//go:linkname F_AllocSetContextCreateInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AllocSetContextCreateInternal
func F_AllocSetContextCreateInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_dsa_create_in_place_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dsa_create_in_place_ext
func F_dsa_create_in_place_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_attach_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_attach_internal
func F_attach_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsa_attach_in_place github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_dsa_attach_in_place
func F_dsa_attach_in_place(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsa_allocate_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsa_allocate_extended
func F_dsa_allocate_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dsa_free github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dsa_free
func F_dsa_free(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dsa_get_address github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsa_get_address
func F_dsa_get_address(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsa_pin github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_dsa_pin
func F_dsa_pin(m *base.Module, l0 int32)
//go:linkname F_dsa_set_size_limit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dsa_set_size_limit
func F_dsa_set_size_limit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_FreePageManagerGetInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FreePageManagerGetInternal
func F_FreePageManagerGetInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FreePageBtreeCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreePageBtreeCleanup
func F_FreePageBtreeCleanup(m *base.Module, l0 int32) int32
//go:linkname F_FreePageManagerPut github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreePageManagerPut
func F_FreePageManagerPut(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GenerationContextCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GenerationContextCreate
func F_GenerationContextCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_MemoryContextReset github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextReset
func F_MemoryContextReset(m *base.Module, l0 int32)
//go:linkname F_MemoryContextResetOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoryContextResetOnly
func F_MemoryContextResetOnly(m *base.Module, l0 int32)
//go:linkname F_MemoryContextDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextDelete
func F_MemoryContextDelete(m *base.Module, l0 int32)
//go:linkname F_MemoryContextSetParent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextSetParent
func F_MemoryContextSetParent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetMemoryChunkContext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetMemoryChunkContext
func F_GetMemoryChunkContext(m *base.Module, l0 int32) int32
//go:linkname F_MemoryContextMemAllocated github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MemoryContextMemAllocated
func F_MemoryContextMemAllocated(m *base.Module, l0 int32) int32
//go:linkname F_MemoryContextAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextAlloc
func F_MemoryContextAlloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextAllocZero github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoryContextAllocZero
func F_MemoryContextAllocZero(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextAllocExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextAllocExtended
func F_MemoryContextAllocExtended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ProcessLogMemoryContextInterrupt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcessLogMemoryContextInterrupt
func F_ProcessLogMemoryContextInterrupt(m *base.Module)
//go:linkname F_palloc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_palloc
func F_palloc(m *base.Module, l0 int32) int32
//go:linkname F_palloc0 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_palloc0
func F_palloc0(m *base.Module, l0 int32) int32
//go:linkname F_palloc_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_palloc_extended
func F_palloc_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_repalloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_repalloc
func F_repalloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_repalloc_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_repalloc_extended
func F_repalloc_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_repalloc0 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_repalloc0
func F_repalloc0(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MemoryContextStrdup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MemoryContextStrdup
func F_MemoryContextStrdup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pnstrdup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pnstrdup
func F_pnstrdup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetPortalByName github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetPortalByName
func F_GetPortalByName(m *base.Module, l0 int32) int32
//go:linkname F_CreatePortal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreatePortal
func F_CreatePortal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_PortalDrop github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PortalDrop
func F_PortalDrop(m *base.Module, l0 int32, l1 int32)
//go:linkname F_HoldPortal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_HoldPortal
func F_HoldPortal(m *base.Module, l0 int32)
//go:linkname F_AtAbort_Portals github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AtAbort_Portals
func F_AtAbort_Portals(m *base.Module)
//go:linkname F_AtSubAbort_Portals github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AtSubAbort_Portals
func F_AtSubAbort_Portals(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ResourceOwnerCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResourceOwnerCreate
func F_ResourceOwnerCreate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ResourceOwnerEnlarge github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResourceOwnerEnlarge
func F_ResourceOwnerEnlarge(m *base.Module, l0 int32)
//go:linkname F_ResourceOwnerReleaseInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResourceOwnerReleaseInternal
func F_ResourceOwnerReleaseInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ResourceOwnerDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ResourceOwnerDelete
func F_ResourceOwnerDelete(m *base.Module, l0 int32)
//go:linkname F_ReleaseAuxProcessResources github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReleaseAuxProcessResources
func F_ReleaseAuxProcessResources(m *base.Module, l0 int32)
//go:linkname F_LogicalTapeSetCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalTapeSetCreate
func F_LogicalTapeSetCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LogicalTapeClose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LogicalTapeClose
func F_LogicalTapeClose(m *base.Module, l0 int32)
//go:linkname F_LogicalTapeRead github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalTapeRead
func F_LogicalTapeRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ltsReadFillBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltsReadFillBuffer
func F_ltsReadFillBuffer(m *base.Module, l0 int32) int32
//go:linkname F_LogicalTapeSetBlocks github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LogicalTapeSetBlocks
func F_LogicalTapeSetBlocks(m *base.Module, l0 int32) int64
//go:linkname F_sts_end_write github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sts_end_write
func F_sts_end_write(m *base.Module, l0 int32)
//go:linkname F_sts_begin_parallel_scan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sts_begin_parallel_scan
func F_sts_begin_parallel_scan(m *base.Module, l0 int32)
//go:linkname F_sts_end_parallel_scan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sts_end_parallel_scan
func F_sts_end_parallel_scan(m *base.Module, l0 int32)
//go:linkname F_sts_puttuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sts_puttuple
func F_sts_puttuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_sts_parallel_scan_next github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sts_parallel_scan_next
func F_sts_parallel_scan_next(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_PrepareSortSupportFromIndexRel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PrepareSortSupportFromIndexRel
func F_PrepareSortSupportFromIndexRel(m *base.Module, l0 int32, l1 int32, l2 int32)
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
//go:linkname F_tuplesort_skiptuples github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_skiptuples
func F_tuplesort_skiptuples(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_tuplesort_rescan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_rescan
func F_tuplesort_rescan(m *base.Module, l0 int32)
//go:linkname F_tuplesort_readtup_alloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_readtup_alloc
func F_tuplesort_readtup_alloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tuplesort_estimate_shared github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_estimate_shared
func F_tuplesort_estimate_shared(m *base.Module, l0 int32) int32
//go:linkname F_tuplesort_initialize_shared github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_initialize_shared
func F_tuplesort_initialize_shared(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_begin_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_begin_heap
func F_tuplesort_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_tuplesort_begin_index_gin github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_begin_index_gin
func F_tuplesort_begin_index_gin(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplesort_puttupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_puttupleslot
func F_tuplesort_puttupleslot(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplesort_putdatum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_putdatum
func F_tuplesort_putdatum(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_gettupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_gettupleslot
func F_tuplesort_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_tuplesort_getheaptuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_getheaptuple
func F_tuplesort_getheaptuple(m *base.Module, l0 int32) int32
//go:linkname F_tuplesort_getgintuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_getgintuple
func F_tuplesort_getgintuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tuplesort_getdatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_getdatum
func F_tuplesort_getdatum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_tuplestore_clear github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tuplestore_clear
func F_tuplestore_clear(m *base.Module, l0 int32)
//go:linkname F_tuplestore_end github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplestore_end
func F_tuplestore_end(m *base.Module, l0 int32)
//go:linkname F_tuplestore_select_read_pointer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplestore_select_read_pointer
func F_tuplestore_select_read_pointer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_puttupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplestore_puttupleslot
func F_tuplestore_puttupleslot(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_puttuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplestore_puttuple
func F_tuplestore_puttuple(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_putvalues github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplestore_putvalues
func F_tuplestore_putvalues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_tuplestore_gettuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplestore_gettuple
func F_tuplestore_gettuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplestore_copy_read_pointer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplestore_copy_read_pointer
func F_tuplestore_copy_read_pointer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplestore_trim github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplestore_trim
func F_tuplestore_trim(m *base.Module, l0 int32)
//go:linkname F_HeapTupleHeaderAdjustCmax github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_HeapTupleHeaderAdjustCmax
func F_HeapTupleHeaderAdjustCmax(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetTransactionSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetTransactionSnapshot
func F_GetTransactionSnapshot(m *base.Module) int32
//go:linkname F_GetCatalogSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetCatalogSnapshot
func F_GetCatalogSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_PushActiveSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PushActiveSnapshot
func F_PushActiveSnapshot(m *base.Module, l0 int32)
//go:linkname F_PushActiveSnapshotWithLevel github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PushActiveSnapshotWithLevel
func F_PushActiveSnapshotWithLevel(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UpdateActiveSnapshotCommandId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UpdateActiveSnapshotCommandId
func F_UpdateActiveSnapshotCommandId(m *base.Module)
//go:linkname F_PopActiveSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PopActiveSnapshot
func F_PopActiveSnapshot(m *base.Module)
//go:linkname F_RegisterSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RegisterSnapshot
func F_RegisterSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_RegisterSnapshotOnOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RegisterSnapshotOnOwner
func F_RegisterSnapshotOnOwner(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_UnregisterSnapshotFromOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_UnregisterSnapshotFromOwner
func F_UnregisterSnapshotFromOwner(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SetTransactionSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SetTransactionSnapshot
func F_SetTransactionSnapshot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_EstimateSnapshotSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EstimateSnapshotSpace
func F_EstimateSnapshotSpace(m *base.Module, l0 int32) int32
//go:linkname F_XidInMVCCSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XidInMVCCSnapshot
func F_XidInMVCCSnapshot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_localsub github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_localsub
func F_localsub(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_tzset github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_tzset
func F_pg_tzset(m *base.Module, l0 int32) int32
//go:linkname F_pg_strftime github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_strftime
func F_pg_strftime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_jit_compile_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jit_compile_expr
func F_jit_compile_expr(m *base.Module, l0 int32) int32
//go:linkname F_pg_checksum_init github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_checksum_init
func F_pg_checksum_init(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_checksum_update github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_checksum_update
func F_pg_checksum_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_double_to_shortest_decimal_buf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_double_to_shortest_decimal_buf
func F_double_to_shortest_decimal_buf(m *base.Module, l0 float64, l1 int32)
//go:linkname F_pg_char_to_encoding_private github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_char_to_encoding_private
func F_pg_char_to_encoding_private(m *base.Module, l0 int32) int32
//go:linkname F_hash_bytes_uint32 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hash_bytes_uint32
func F_hash_bytes_uint32(m *base.Module, l0 int32) int32
//go:linkname F_pg_getaddrinfo_all github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_getaddrinfo_all
func F_pg_getaddrinfo_all(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_getnameinfo_all github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_getnameinfo_all
func F_pg_getnameinfo_all(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_makeJsonLexContextCstringLen github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeJsonLexContextCstringLen
func F_makeJsonLexContextCstringLen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_freeJsonLexContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_freeJsonLexContext
func F_freeJsonLexContext(m *base.Module, l0 int32)
//go:linkname F_pg_parse_json github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_parse_json
func F_pg_parse_json(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_json_lex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_json_lex
func F_json_lex(m *base.Module, l0 int32) int32
//go:linkname F_parse_object github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_object
func F_parse_object(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_parse_array
func F_parse_array(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_scalar github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_parse_scalar
func F_parse_scalar(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_md5_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_md5_hash
func F_pg_md5_hash(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_prng_uint64_range github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_prng_uint64_range
func F_pg_prng_uint64_range(m *base.Module, l0 int32, l1 int64, l2 int64) int64
//go:linkname F_pg_prng_double github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_prng_double
func F_pg_prng_double(m *base.Module, l0 int32) float64
//go:linkname F_GetRelationPath github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetRelationPath
func F_GetRelationPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_rmtree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_rmtree
func F_rmtree(m *base.Module, l0 int32) int32
//go:linkname F_makeStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeStringInfo
func F_makeStringInfo(m *base.Module) int32
//go:linkname F_initStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initStringInfo
func F_initStringInfo(m *base.Module, l0 int32)
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
//go:linkname F_case_index github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_case_index
func F_case_index(m *base.Module, l0 int32) int32
//go:linkname F_convert_case github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_convert_case
func F_convert_case(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_unicode_normalize github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_unicode_normalize
func F_unicode_normalize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_canonicalize_path_enc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_canonicalize_path_enc
func F_canonicalize_path_enc(m *base.Module, l0 int32)
//go:linkname F_get_share_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_share_path
func F_get_share_path(m *base.Module, l0 int32)
//go:linkname F_make_relative_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_relative_path
func F_make_relative_path(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_usleep github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_usleep
func F_pg_usleep(m *base.Module, l0 int32)
//go:linkname F_pg_strcasecmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_strcasecmp
func F_pg_strcasecmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_toupper github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_toupper
func F_pg_toupper(m *base.Module, l0 int32) int32
//go:linkname F_pqsignal_be github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pqsignal_be
func F_pqsignal_be(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_qsort github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_qsort
func F_pg_qsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_qsort_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_qsort_arg
func F_qsort_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_dopr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dopr
func F_dopr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dopr_outchmulti github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dopr_outchmulti
func F_dopr_outchmulti(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_snprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_snprintf
func F_pg_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_fprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_fprintf
func F_pg_fprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_printf github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_printf
func F_pg_printf(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_strfromd github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_strfromd
func F_pg_strfromd(m *base.Module, l0 int32, l1 int32, l2 float64)
//go:linkname F_pgl_longjmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgl_longjmp
func F_pgl_longjmp(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgl_popen github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgl_popen
func F_pgl_popen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgl_pclose github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgl_pclose
func F_pgl_pclose(m *base.Module, l0 int32) int32
//go:linkname F_pgl_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgl_exit
func F_pgl_exit(m *base.Module, l0 int32)
//go:linkname F_pgl_shmctl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgl_shmctl
func F_pgl_shmctl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_datatype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_datatype
func F_build_datatype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_resolve_column_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_resolve_column_ref
func F_resolve_column_ref(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_plpgsql_build_recfield github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plpgsql_build_recfield
func F_plpgsql_build_recfield(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assign_simple_var github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_assign_simple_var
func F_assign_simple_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_revalidate_rectypeid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_revalidate_rectypeid
func F_revalidate_rectypeid(m *base.Module, l0 int32)
//go:linkname F_make_expanded_record_for_rec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_expanded_record_for_rec
func F_make_expanded_record_for_rec(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_exec_move_row_from_fields github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_exec_move_row_from_fields
func F_exec_move_row_from_fields(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_instantiate_empty_record_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_instantiate_empty_record_variable
func F_instantiate_empty_record_variable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_exec_stmts github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_exec_stmts
func F_exec_stmts(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_exec_prepare_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_prepare_plan
func F_exec_prepare_plan(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_exec_eval_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exec_eval_expr
func F_exec_eval_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_make_tuple_from_row github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_tuple_from_row
func F_make_tuple_from_row(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_assign_text_var github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_assign_text_var
func F_assign_text_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_format_expr_params github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_format_expr_params
func F_format_expr_params(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_ns_additem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_ns_additem
func F_plpgsql_ns_additem(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_mark_stmt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mark_stmt
func F_mark_stmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dump_stmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dump_stmt
func F_dump_stmt(m *base.Module, l0 int32)
//go:linkname F_plpgsql_scanner_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_plpgsql_scanner_errposition
func F_plpgsql_scanner_errposition(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SN_set_current github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SN_set_current
func F_SN_set_current(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_r_consonant_pair_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_consonant_pair_2
func F_r_consonant_pair_2(m *base.Module, l0 int32) int32
//go:linkname F_r_check_vowel_harmony github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_check_vowel_harmony
func F_r_check_vowel_harmony(m *base.Module, l0 int32) int32
//go:linkname F_create_s github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_s
func F_create_s(m *base.Module) int32
//go:linkname F_lose_s github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lose_s
func F_lose_s(m *base.Module, l0 int32)
//go:linkname F_skip_b_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_skip_b_utf8
func F_skip_b_utf8(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_find_among github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_find_among
func F_find_among(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_among_b github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_among_b
func F_find_among_b(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_slice_from_s github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_slice_from_s
func F_slice_from_s(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_slice_del github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_slice_del
func F_slice_del(m *base.Module, l0 int32) int32
//go:linkname F_des_init github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_des_init
func F_des_init(m *base.Module)
//go:linkname F_pullf_read github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pullf_read
func F_pullf_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pullf_read_fixed github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pullf_read_fixed
func F_pullf_read_fixed(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pushf_create github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pushf_create
func F_pushf_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pushf_write github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pushf_write
func F_pushf_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_px_find_digest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_px_find_digest
func F_px_find_digest(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_px_find_cipher github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_px_find_cipher
func F_px_find_cipher(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CheckBuiltinCryptoMode github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CheckBuiltinCryptoMode
func F_CheckBuiltinCryptoMode(m *base.Module)
//go:linkname F_pgp_cfb_free github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgp_cfb_free
func F_pgp_cfb_free(m *base.Module, l0 int32)
//go:linkname F_pgp_cfb_decrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgp_cfb_decrypt
func F_pgp_cfb_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pgp_key_free github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgp_key_free
func F_pgp_key_free(m *base.Module, l0 int32)
//go:linkname F_pgp_get_digest_code github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgp_get_digest_code
func F_pgp_get_digest_code(m *base.Module, l0 int32) int32
//go:linkname F_pgp_get_cipher_code github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgp_get_cipher_code
func F_pgp_get_cipher_code(m *base.Module, l0 int32) int32
//go:linkname F_pgp_init github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgp_init
func F_pgp_init(m *base.Module, l0 int32) int32
//go:linkname F_px_THROW_ERROR github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_px_THROW_ERROR
func F_px_THROW_ERROR(m *base.Module, l0 int32)
//go:linkname F_px_debug github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_px_debug
func F_px_debug(m *base.Module, l0 int32, l1 int32)
//go:linkname F_citextcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_citextcmp
func F_citextcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_index_strategy_get_limit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_strategy_get_limit
func F_index_strategy_get_limit(m *base.Module, l0 int32) float64
//go:linkname F_generate_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_trgm
func F_generate_trgm(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generate_wildcard_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_generate_wildcard_trgm
func F_generate_wildcard_trgm(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_trgm_contained_by github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_trgm_contained_by
func F_trgm_contained_by(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_trgm_presence_map github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_trgm_presence_map
func F_trgm_presence_map(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_calc_word_similarity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_calc_word_similarity
func F_calc_word_similarity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float32
//go:linkname F_createTrgmNFA github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_createTrgmNFA
func F_createTrgmNFA(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hstoreUpgrade github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstoreUpgrade
func F_hstoreUpgrade(m *base.Module, l0 int32) int32
//go:linkname F_hstoreUniquePairs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstoreUniquePairs
func F_hstoreUniquePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hstorePairs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hstorePairs
func F_hstorePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hstore_exists github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hstore_exists
func F_hstore_exists(m *base.Module, l0 int32) int32
//go:linkname F_array_iterator github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_iterator
func F_array_iterator(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ltree_gist_alloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltree_gist_alloc
func F_ltree_gist_alloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_finish_nodeitem github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_finish_nodeitem
func F_finish_nodeitem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_parse_lquery github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_parse_lquery
func F_parse_lquery(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_inner_subltree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_inner_subltree
func F_inner_subltree(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_infix_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_infix_2
func F_infix_2(m *base.Module, l0 int32, l1 int32)
//go:linkname F_abs_interval github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_abs_interval
func F_abs_interval(m *base.Module, l0 int32) int32
//go:linkname F_gbt_num_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_num_compress
func F_gbt_num_compress(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gbt_num_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_num_fetch
func F_gbt_num_fetch(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gbt_num_same github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_num_same
func F_gbt_num_same(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gbt_num_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_num_consistent
func F_gbt_num_consistent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_gbt_num_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_num_distance
func F_gbt_num_distance(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_gbt_var_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_var_consistent
func F_gbt_var_consistent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_gin_btree_extract_query github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gin_btree_extract_query
func F_gin_btree_extract_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_placeChar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_placeChar
func F_placeChar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_contains_required_value github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_contains_required_value
func F_contains_required_value(m *base.Module, l0 int32) int32
//go:linkname F_int_query_opr_selec github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_int_query_opr_selec
func F_int_query_opr_selec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float32) float64
//go:linkname F_inner_int_overlap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inner_int_overlap
func F_inner_int_overlap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_copy_intArrayType github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_copy_intArrayType
func F_copy_intArrayType(m *base.Module, l0 int32) int32
//go:linkname F_resize_intArrayType github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_resize_intArrayType
func F_resize_intArrayType(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_isort github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isort
func F_isort(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_restore github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_restore
func F_restore(m *base.Module, l0 int32, l1 float32, l2 int32) int32
//go:linkname F_seg_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_seg_yyensure_buffer_stack
func F_seg_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_seg_yy_scan_bytes github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_seg_yy_scan_bytes
func F_seg_yy_scan_bytes(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___memcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___memcpy
func F___memcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__emscripten_memcpy_bulkmem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_memcpy_bulkmem
func F__emscripten_memcpy_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___memset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___memset
func F___memset(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__emscripten_memset_bulkmem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_memset_bulkmem
func F__emscripten_memset_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__Exit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__Exit
func F__Exit(m *base.Module, l0 int32)
//go:linkname F_abort github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_abort
func F_abort(m *base.Module)
//go:linkname F_access github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_access
func F_access(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_atof github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_atof
func F_atof(m *base.Module, l0 int32) float64
//go:linkname F_atoi github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_atoi
func F_atoi(m *base.Module, l0 int32) int32
//go:linkname F_bsearch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bsearch
func F_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F___clock_gettime github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F___clock_gettime
func F___clock_gettime(m *base.Module, l0 int32, l1 int32)
//go:linkname F_close github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_close
func F_close(m *base.Module, l0 int32) int32
//go:linkname F_dup2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dup2
func F_dup2(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_memmove github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_memmove
func F_memmove(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___time github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___time
func F___time(m *base.Module) int64
//go:linkname F___math_xflow github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___math_xflow
func F___math_xflow(m *base.Module, l0 int32, l1 float64) float64
//go:linkname F_fp_barrier_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fp_barrier_1
func F_fp_barrier_1(m *base.Module, l0 float64) float64
//go:linkname F_exp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exp
func F_exp(m *base.Module, l0 float64) float64
//go:linkname F_fflush github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fflush
func F_fflush(m *base.Module, l0 int32) int32
//go:linkname F___toread github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___toread
func F___toread(m *base.Module, l0 int32) int32
//go:linkname F_do_getc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_do_getc
func F_do_getc(m *base.Module, l0 int32) int32
//go:linkname F___overflow github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F___overflow
func F___overflow(m *base.Module, l0 int32, l1 int32)
//go:linkname F___fstatat github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___fstatat
func F___fstatat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_fsync github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fsync
func F_fsync(m *base.Module, l0 int32) int32
//go:linkname F___ftello_unlocked github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F___ftello_unlocked
func F___ftello_unlocked(m *base.Module, l0 int32) int64
//go:linkname F___fwritex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___fwritex
func F___fwritex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fwrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fwrite
func F_fwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getcwd github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getcwd
func F_getcwd(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fputs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fputs
func F_fputs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getopt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getopt
func F_getopt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getrusage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getrusage
func F_getrusage(m *base.Module, l0 int32)
//go:linkname F_isalnum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isalnum
func F_isalnum(m *base.Module, l0 int32) int32
//go:linkname F_iswalpha github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_iswalpha
func F_iswalpha(m *base.Module, l0 int32) int32
//go:linkname F_kill github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_kill
func F_kill(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ldexp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ldexp
func F_ldexp(m *base.Module, l0 float64, l1 int32) float64
//go:linkname F_log github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_log
func F_log(m *base.Module, l0 float64) float64
//go:linkname F___lseek github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___lseek
func F___lseek(m *base.Module, l0 int32, l1 int64, l2 int32) int64
//go:linkname F_memchr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_memchr
func F_memchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_memcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_memcmp
func F_memcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mkdir github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mkdir
func F_mkdir(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_do_tzset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_do_tzset
func F_do_tzset(m *base.Module)
//go:linkname F___mmap github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___mmap
func F___mmap(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___get_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___get_locale
func F___get_locale(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___loc_is_allocated github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___loc_is_allocated
func F___loc_is_allocated(m *base.Module, l0 int32) int32
//go:linkname F_open github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_open
func F_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_posix_fadvise github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_posix_fadvise
func F_posix_fadvise(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32
//go:linkname F_pwrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pwrite
func F_pwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_rename github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_rename
func F_rename(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rmdir github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_rmdir
func F_rmdir(m *base.Module, l0 int32) int32
//go:linkname F_scalbn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_scalbn
func F_scalbn(m *base.Module, l0 float64, l1 int32) float64
//go:linkname F_sigprocmask github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sigprocmask
func F_sigprocmask(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sscanf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sscanf
func F_sscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strchr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strchr
func F_strchr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___strchrnul github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___strchrnul
func F___strchrnul(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strcmp
func F_strcmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strcpy
func F_strcpy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strlcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strlcpy
func F_strlcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strncmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strncmp
func F_strncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strncpy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strncpy
func F_strncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strspn github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_strspn
func F_strspn(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strstr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strstr
func F_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___floatscan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___floatscan
func F___floatscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_strtoull github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_strtoull
func F_strtoull(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_strtox_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strtox_2
func F_strtox_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64
//go:linkname F_strtoul github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strtoul
func F_strtoul(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strtol github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strtol
func F_strtol(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___syscall_ret github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___syscall_ret
func F___syscall_ret(m *base.Module, l0 int32) int32
//go:linkname F_tolower github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tolower
func F_tolower(m *base.Module, l0 int32) int32
//go:linkname F_toupper github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_toupper
func F_toupper(m *base.Module, l0 int32) int32
//go:linkname F_towlower github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_towlower
func F_towlower(m *base.Module, l0 int32) int32
//go:linkname F_unlink github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_unlink
func F_unlink(m *base.Module, l0 int32) int32
//go:linkname F_frexp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_frexp
func F_frexp(m *base.Module, l0 float64, l1 int32) float64
//go:linkname F_pad github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pad
func F_pad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_write github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_write
func F_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_emscripten_builtin_malloc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_emscripten_builtin_malloc
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32
//go:linkname F_emscripten_builtin_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_emscripten_builtin_free
func F_emscripten_builtin_free(m *base.Module, l0 int32)
//go:linkname F_emscripten_builtin_realloc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_emscripten_builtin_realloc
func F_emscripten_builtin_realloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___ashlti3 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___ashlti3
func F___ashlti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32)
//go:linkname F___wasm_longjmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___wasm_longjmp
func F___wasm_longjmp(m *base.Module, l0 int32, l1 int32)
//go:linkname F___lshrti3 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___lshrti3
func F___lshrti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32)
//go:linkname F___udivmodti4 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___udivmodti4
func F___udivmodti4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F_freeaddrinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_freeaddrinfo
func F_freeaddrinfo(m *base.Module, l0 int32)
//go:linkname F_recvfrom github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recvfrom
func F_recvfrom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_sendto github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sendto
func F_sendto(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
