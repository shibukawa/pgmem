package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	_ "unsafe"
)
//go:linkname F_bloom_get_procinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_bloom_get_procinfo
func F_bloom_get_procinfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_inclusion_get_procinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_inclusion_get_procinfo
func F_inclusion_get_procinfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_minmax_multi_get_strategy_procinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_minmax_multi_get_strategy_procinfo
func F_minmax_multi_get_strategy_procinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_brin_range_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_brin_range_deserialize
func F_brin_range_deserialize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_reduce_expanded_ranges github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_reduce_expanded_ranges
func F_reduce_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_brin_range_serialize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_range_serialize
func F_brin_range_serialize(m *base.Module, l0 int32) int32
//go:linkname F_compare_expanded_ranges github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_compare_expanded_ranges
func F_compare_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_brin_getinsertbuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_brin_getinsertbuffer
func F_brin_getinsertbuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_brin_initialize_empty_new_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_brin_initialize_empty_new_buffer
func F_brin_initialize_empty_new_buffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_brinRevmapTerminate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_brinRevmapTerminate
func F_brinRevmapTerminate(m *base.Module, l0 int32)
//go:linkname F_brinRevmapExtend github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brinRevmapExtend
func F_brinRevmapExtend(m *base.Module, l0 int32, l1 int32)
//go:linkname F_brinLockRevmapPageForUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brinLockRevmapPageForUpdate
func F_brinLockRevmapPageForUpdate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_brinGetTupleForHeapBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_brinGetTupleForHeapBlock
func F_brinGetTupleForHeapBlock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_brin_copy_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_copy_tuple
func F_brin_copy_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_brin_new_memtuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_new_memtuple
func F_brin_new_memtuple(m *base.Module, l0 int32) int32
//go:linkname F_brin_deform_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_brin_deform_tuple
func F_brin_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_attrmap_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_attrmap_by_name
func F_build_attrmap_by_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_attrmap_by_name_if_req github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_build_attrmap_by_name_if_req
func F_build_attrmap_by_name_if_req(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mask_unused_space github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mask_unused_space
func F_mask_unused_space(m *base.Module, l0 int32)
//go:linkname F_detoast_external_attr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_detoast_external_attr
func F_detoast_external_attr(m *base.Module, l0 int32) int32
//go:linkname F_toast_fetch_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_toast_fetch_datum
func F_toast_fetch_datum(m *base.Module, l0 int32) int32
//go:linkname F_toast_raw_datum_size github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_toast_raw_datum_size
func F_toast_raw_datum_size(m *base.Module, l0 int32) int32
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
//go:linkname F_nocachegetattr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_nocachegetattr
func F_nocachegetattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_getsysattr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_getsysattr
func F_heap_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_copytuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_copytuple
func F_heap_copytuple(m *base.Module, l0 int32) int32
//go:linkname F_heap_modify_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_modify_tuple
func F_heap_modify_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_heap_deform_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_deform_tuple
func F_heap_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_heap_form_minimal_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_form_minimal_tuple
func F_heap_form_minimal_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_index_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_form_tuple
func F_index_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_index_deform_tuple_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_deform_tuple_internal
func F_index_deform_tuple_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_CopyIndexTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CopyIndexTuple
func F_CopyIndexTuple(m *base.Module, l0 int32) int32
//go:linkname F_relation_open github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_relation_open
func F_relation_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_relation_close github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_relation_close
func F_relation_close(m *base.Module, l0 int32, l1 int32)
//go:linkname F_add_local_int_reloption github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_local_int_reloption
func F_add_local_int_reloption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_untransformRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_untransformRelOptions
func F_untransformRelOptions(m *base.Module, l0 int32) int32
//go:linkname F_extractRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_extractRelOptions
func F_extractRelOptions(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_reloptions
func F_build_reloptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ScanKeyEntryInitialize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ScanKeyEntryInitialize
func F_ScanKeyEntryInitialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_pglz_decompress_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pglz_decompress_datum
func F_pglz_decompress_datum(m *base.Module, l0 int32) int32
//go:linkname F_toast_compress_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_toast_compress_datum
func F_toast_compress_datum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_toast_open_indexes github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_toast_open_indexes
func F_toast_open_indexes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_toastrel_valueid_exists github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_toastrel_valueid_exists
func F_toastrel_valueid_exists(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_toast_delete_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_toast_delete_datum
func F_toast_delete_datum(m *base.Module, l0 int32, l1 int32)
//go:linkname F_toast_get_valid_index github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_toast_get_valid_index
func F_toast_get_valid_index(m *base.Module, l0 int32, l1 int32) int32
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
//go:linkname F_CreateTupleDescTruncatedCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateTupleDescTruncatedCopy
func F_CreateTupleDescTruncatedCopy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CreateTupleDescCopyConstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateTupleDescCopyConstr
func F_CreateTupleDescCopyConstr(m *base.Module, l0 int32) int32
//go:linkname F_TupleDescCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TupleDescCopy
func F_TupleDescCopy(m *base.Module, l0 int32, l1 int32)
//go:linkname F_TupleDescCopyEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TupleDescCopyEntry
func F_TupleDescCopyEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_FreeTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeTupleDesc
func F_FreeTupleDesc(m *base.Module, l0 int32)
//go:linkname F_DecrTupleDescRefCount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecrTupleDescRefCount
func F_DecrTupleDescRefCount(m *base.Module, l0 int32)
//go:linkname F_TupleDescInitEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TupleDescInitEntry
func F_TupleDescInitEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_BuildDescFromLists github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BuildDescFromLists
func F_BuildDescFromLists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_disassembleLeaf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_disassembleLeaf
func F_disassembleLeaf(m *base.Module, l0 int32) int32
//go:linkname F_computeLeafRecompressWALData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_computeLeafRecompressWALData
func F_computeLeafRecompressWALData(m *base.Module, l0 int32)
//go:linkname F_ginCompressPostingList github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ginCompressPostingList
func F_ginCompressPostingList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ginPostingListDecodeAllSegments github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ginPostingListDecodeAllSegments
func F_ginPostingListDecodeAllSegments(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ginMergeItemPointers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ginMergeItemPointers
func F_ginMergeItemPointers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_initGinState github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initGinState
func F_initGinState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_gintuple_get_attrnum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gintuple_get_attrnum
func F_gintuple_get_attrnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gintuple_get_key github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gintuple_get_key
func F_gintuple_get_key(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ginCompareEntries github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ginCompareEntries
func F_ginCompareEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ginCompareAttEntries github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ginCompareAttEntries
func F_ginCompareAttEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_gistfinishsplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistfinishsplit
func F_gistfinishsplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_gistplacetopage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistplacetopage
func F_gistplacetopage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32
//go:linkname F_gistScanPage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gistScanPage
func F_gistScanPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_index_getattr_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_getattr_2
func F_index_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_rtree_internal_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_rtree_internal_consistent
func F_rtree_internal_consistent(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_computeDistance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_computeDistance
func F_computeDistance(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_gistDeCompressAtt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistDeCompressAtt
func F_gistDeCompressAtt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F__hash_pgaddmultitup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__hash_pgaddmultitup
func F__hash_pgaddmultitup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F__hash_addovflpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_addovflpage
func F__hash_addovflpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__hash_getbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_getbuf
func F__hash_getbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__hash_getnewbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_getnewbuf
func F__hash_getnewbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__hash_getbuf_with_strategy github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__hash_getbuf_with_strategy
func F__hash_getbuf_with_strategy(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__hash_dropscanbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__hash_dropscanbuf
func F__hash_dropscanbuf(m *base.Module, l0 int32)
//go:linkname F__hash_getcachedmetap github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__hash_getcachedmetap
func F__hash_getcachedmetap(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__hash_next github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_next
func F__hash_next(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__hash_first github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__hash_first
func F__hash_first(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__hash_spareindex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__hash_spareindex
func F__hash_spareindex(m *base.Module, l0 int32) int32
//go:linkname F__hash_get_totalbuckets github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__hash_get_totalbuckets
func F__hash_get_totalbuckets(m *base.Module, l0 int32) int32
//go:linkname F__hash_checkpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_checkpage
func F__hash_checkpage(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F__hash_kill_items github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_kill_items
func F__hash_kill_items(m *base.Module, l0 int32)
//go:linkname F_heap_prepare_pagescan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_prepare_pagescan
func F_heap_prepare_pagescan(m *base.Module, l0 int32)
//go:linkname F_HeapCheckForSerializableConflictOut github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HeapCheckForSerializableConflictOut
func F_HeapCheckForSerializableConflictOut(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_heapgettup github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heapgettup
func F_heapgettup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_heap_fetch_next_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_fetch_next_buffer
func F_heap_fetch_next_buffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_heap_prepare_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_heap_prepare_insert
func F_heap_prepare_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_log_heap_new_cid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_log_heap_new_cid
func F_log_heap_new_cid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_simple_heap_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_simple_heap_insert
func F_simple_heap_insert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_HeapTupleGetUpdateXid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_HeapTupleGetUpdateXid
func F_HeapTupleGetUpdateXid(m *base.Module, l0 int32) int32
//go:linkname F_simple_heap_update github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_simple_heap_update
func F_simple_heap_update(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_HeapTupleSatisfiesVacuumHorizon github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_HeapTupleSatisfiesVacuumHorizon
func F_HeapTupleSatisfiesVacuumHorizon(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_HeapTupleSatisfiesVisibility github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_HeapTupleSatisfiesVisibility
func F_HeapTupleSatisfiesVisibility(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_toast_flatten_tuple_to_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_toast_flatten_tuple_to_datum
func F_toast_flatten_tuple_to_datum(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fastgetattr_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fastgetattr_1
func F_fastgetattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_RelationPutHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationPutHeapTuple
func F_RelationPutHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RelationGetBufferForTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetBufferForTuple
func F_RelationGetBufferForTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_raw_heap_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_raw_heap_insert
func F_raw_heap_insert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_logical_heap_rewrite_flush_mappings github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_logical_heap_rewrite_flush_mappings
func F_logical_heap_rewrite_flush_mappings(m *base.Module, l0 int32)
//go:linkname F_rewrite_heap_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_rewrite_heap_tuple
func F_rewrite_heap_tuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_vm_readbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_vm_readbuf
func F_vm_readbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_visibilitymap_prepare_truncate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_visibilitymap_prepare_truncate
func F_visibilitymap_prepare_truncate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetIndexAmRoutine github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetIndexAmRoutine
func F_GetIndexAmRoutine(m *base.Module, l0 int32) int32
//go:linkname F_GetIndexAmRoutineByAmId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetIndexAmRoutineByAmId
func F_GetIndexAmRoutineByAmId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_IndexAmTranslateCompareType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IndexAmTranslateCompareType
func F_IndexAmTranslateCompareType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_identify_opfamily_groups github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_identify_opfamily_groups
func F_identify_opfamily_groups(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_amproc_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_amproc_signature
func F_check_amproc_signature(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_check_amoptsproc_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_amoptsproc_signature
func F_check_amoptsproc_signature(m *base.Module, l0 int32) int32
//go:linkname F_check_amop_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_amop_signature
func F_check_amop_signature(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_opclass_for_family_datatype github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_opclass_for_family_datatype
func F_opclass_for_family_datatype(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RelationGetIndexScan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationGetIndexScan
func F_RelationGetIndexScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_systable_beginscan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_beginscan
func F_systable_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_systable_getnext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_systable_getnext
func F_systable_getnext(m *base.Module, l0 int32) int32
//go:linkname F_systable_recheck_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_recheck_tuple
func F_systable_recheck_tuple(m *base.Module, l0 int32) int32
//go:linkname F_systable_beginscan_ordered github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_systable_beginscan_ordered
func F_systable_beginscan_ordered(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_systable_getnext_ordered github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_systable_getnext_ordered
func F_systable_getnext_ordered(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_systable_endscan_ordered github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_endscan_ordered
func F_systable_endscan_ordered(m *base.Module, l0 int32)
//go:linkname F_index_open github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_open
func F_index_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_insert_cleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_insert_cleanup
func F_index_insert_cleanup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_index_beginscan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_beginscan
func F_index_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_index_rescan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_rescan
func F_index_rescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_index_endscan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_endscan
func F_index_endscan(m *base.Module, l0 int32)
//go:linkname F_index_parallelscan_initialize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_index_parallelscan_initialize
func F_index_parallelscan_initialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_index_beginscan_parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_beginscan_parallel
func F_index_beginscan_parallel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_index_getnext_slot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_getnext_slot
func F_index_getnext_slot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_index_bulk_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_bulk_delete
func F_index_bulk_delete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_index_can_return github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_index_can_return
func F_index_can_return(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_getprocinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_getprocinfo
func F_index_getprocinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_btoidvectorcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_btoidvectorcmp
func F_btoidvectorcmp(m *base.Module, l0 int32) int32
//go:linkname F__bt_insert_parent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_insert_parent
func F__bt_insert_parent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F__bt_checkpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_checkpage
func F__bt_checkpage(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_getbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_getbuf
func F__bt_getbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_relbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_relbuf
func F__bt_relbuf(m *base.Module, l0 int32)
//go:linkname F__bt_getmeta github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_getmeta
func F__bt_getmeta(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_allocbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_allocbuf
func F__bt_allocbuf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_relandgetbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_relandgetbuf
func F__bt_relandgetbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_saoparray_shrink github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_saoparray_shrink
func F__bt_saoparray_shrink(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F__bt_parallel_scan_and_sort github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_parallel_scan_and_sort
func F__bt_parallel_scan_and_sort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F__bt_mkscankey github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_mkscankey
func F__bt_mkscankey(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_check_compare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_check_compare
func F__bt_check_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F__bt_advance_array_keys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_advance_array_keys
func F__bt_advance_array_keys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F__bt_binsrch_skiparray_skey github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__bt_binsrch_skiparray_skey
func F__bt_binsrch_skiparray_skey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_infobits_desc github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_infobits_desc
func F_infobits_desc(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_array_desc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_array_desc
func F_array_desc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_standby_desc_invalidations github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_standby_desc_invalidations
func F_standby_desc_invalidations(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_spgPageIndexMultiDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_spgPageIndexMultiDelete
func F_spgPageIndexMultiDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_spg_key_orderbys_distances github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_spg_key_orderbys_distances
func F_spg_key_orderbys_distances(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_box_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_box_copy
func F_box_copy(m *base.Module, l0 int32) int32
//go:linkname F_getQuadrant github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_getQuadrant
func F_getQuadrant(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_open github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_table_open
func F_table_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_openrv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_table_openrv
func F_table_openrv(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_slot_callbacks github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_table_slot_callbacks
func F_table_slot_callbacks(m *base.Module, l0 int32) int32
//go:linkname F_table_slot_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_table_slot_create
func F_table_slot_create(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_beginscan_catalog github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_table_beginscan_catalog
func F_table_beginscan_catalog(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_table_parallelscan_initialize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_table_parallelscan_initialize
func F_table_parallelscan_initialize(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_table_beginscan_parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_table_beginscan_parallel
func F_table_beginscan_parallel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_tuple_get_latest_tid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_table_tuple_get_latest_tid
func F_table_tuple_get_latest_tid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_toast_tuple_try_compression github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_toast_tuple_try_compression
func F_toast_tuple_try_compression(m *base.Module, l0 int32, l1 int32)
//go:linkname F_TransactionIdSetTreeStatus github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TransactionIdSetTreeStatus
func F_TransactionIdSetTreeStatus(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64)
//go:linkname F_TransactionIdGetStatus github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TransactionIdGetStatus
func F_TransactionIdGetStatus(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_error_commit_ts_disabled github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_error_commit_ts_disabled
func F_error_commit_ts_disabled(m *base.Module)
//go:linkname F_GetMultiXactIdMembers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetMultiXactIdMembers
func F_GetMultiXactIdMembers(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReadNextMultiXactId github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReadNextMultiXactId
func F_ReadNextMultiXactId(m *base.Module) int32
//go:linkname F_SetMultiXactIdLimit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SetMultiXactIdLimit
func F_SetMultiXactIdLimit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_MultiXactMemberFreezeThreshold github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MultiXactMemberFreezeThreshold
func F_MultiXactMemberFreezeThreshold(m *base.Module) int32
//go:linkname F_WaitForParallelWorkersToExit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WaitForParallelWorkersToExit
func F_WaitForParallelWorkersToExit(m *base.Module, l0 int32)
//go:linkname F_RmgrNotFound github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RmgrNotFound
func F_RmgrNotFound(m *base.Module, l0 int32)
//go:linkname F_check_slru_buffers github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_slru_buffers
func F_check_slru_buffers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SlruSelectLRUPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SlruSelectLRUPage
func F_SlruSelectLRUPage(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_SimpleLruReadPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SimpleLruReadPage
func F_SimpleLruReadPage(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_SimpleLruReadPage_ReadOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SimpleLruReadPage_ReadOnly
func F_SimpleLruReadPage_ReadOnly(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_SimpleLruWritePage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SimpleLruWritePage
func F_SimpleLruWritePage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SimpleLruDoesPhysicalPageExist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SimpleLruDoesPhysicalPageExist
func F_SimpleLruDoesPhysicalPageExist(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_SimpleLruWriteAll github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SimpleLruWriteAll
func F_SimpleLruWriteAll(m *base.Module, l0 int32)
//go:linkname F_SimpleLruTruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SimpleLruTruncate
func F_SimpleLruTruncate(m *base.Module, l0 int32, l1 int64)
//go:linkname F_SlruSyncFileTag github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SlruSyncFileTag
func F_SlruSyncFileTag(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_readTimeLineHistory github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_readTimeLineHistory
func F_readTimeLineHistory(m *base.Module, l0 int32) int32
//go:linkname F_existsTimeLineHistory github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_existsTimeLineHistory
func F_existsTimeLineHistory(m *base.Module, l0 int32) int32
//go:linkname F_tliSwitchPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tliSwitchPoint
func F_tliSwitchPoint(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_TransactionIdDidCommit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransactionIdDidCommit
func F_TransactionIdDidCommit(m *base.Module, l0 int32) int32
//go:linkname F_TransactionIdPrecedes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TransactionIdPrecedes
func F_TransactionIdPrecedes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TwoPhaseGetDummyProc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TwoPhaseGetDummyProc
func F_TwoPhaseGetDummyProc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReadNextFullTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadNextFullTransactionId
func F_ReadNextFullTransactionId(m *base.Module) int64
//go:linkname F_AssignTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AssignTransactionId
func F_AssignTransactionId(m *base.Module, l0 int32)
//go:linkname F_GetCurrentTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetCurrentTransactionId
func F_GetCurrentTransactionId(m *base.Module) int32
//go:linkname F_GetCurrentSubTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetCurrentSubTransactionId
func F_GetCurrentSubTransactionId(m *base.Module) int32
//go:linkname F_GetCurrentCommandId github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetCurrentCommandId
func F_GetCurrentCommandId(m *base.Module, l0 int32) int32
//go:linkname F_StartTransactionCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_StartTransactionCommand
func F_StartTransactionCommand(m *base.Module)
//go:linkname F_StartTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_StartTransaction
func F_StartTransaction(m *base.Module)
//go:linkname F_CommitTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CommitTransaction
func F_CommitTransaction(m *base.Module)
//go:linkname F_CleanupTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CleanupTransaction
func F_CleanupTransaction(m *base.Module)
//go:linkname F_AbortTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AbortTransaction
func F_AbortTransaction(m *base.Module)
//go:linkname F_PrepareTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PrepareTransaction
func F_PrepareTransaction(m *base.Module)
//go:linkname F_StartSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_StartSubTransaction
func F_StartSubTransaction(m *base.Module)
//go:linkname F_CommitSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CommitSubTransaction
func F_CommitSubTransaction(m *base.Module)
//go:linkname F_AbortSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AbortSubTransaction
func F_AbortSubTransaction(m *base.Module)
//go:linkname F_CleanupSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CleanupSubTransaction
func F_CleanupSubTransaction(m *base.Module)
//go:linkname F_DefineSavepoint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DefineSavepoint
func F_DefineSavepoint(m *base.Module, l0 int32)
//go:linkname F_AbortCurrentTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AbortCurrentTransaction
func F_AbortCurrentTransaction(m *base.Module)
//go:linkname F_PreventInTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreventInTransactionBlock
func F_PreventInTransactionBlock(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BeginInternalSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BeginInternalSubTransaction
func F_BeginInternalSubTransaction(m *base.Module, l0 int32)
//go:linkname F_RollbackAndReleaseCurrentSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RollbackAndReleaseCurrentSubTransaction
func F_RollbackAndReleaseCurrentSubTransaction(m *base.Module)
//go:linkname F_AbortOutOfAnyTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AbortOutOfAnyTransaction
func F_AbortOutOfAnyTransaction(m *base.Module)
//go:linkname F_XLogFlush github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogFlush
func F_XLogFlush(m *base.Module, l0 int64)
//go:linkname F_XLogFileName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogFileName
func F_XLogFileName(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_XLogGetOldestSegno github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogGetOldestSegno
func F_XLogGetOldestSegno(m *base.Module, l0 int32) int64
//go:linkname F_LocalProcessControlFile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LocalProcessControlFile
func F_LocalProcessControlFile(m *base.Module)
//go:linkname F_GetRedoRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetRedoRecPtr
func F_GetRedoRecPtr(m *base.Module) int64
//go:linkname F_GetXLogInsertRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetXLogInsertRecPtr
func F_GetXLogInsertRecPtr(m *base.Module) int64
//go:linkname F_XLogArchiveCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogArchiveCleanup
func F_XLogArchiveCleanup(m *base.Module, l0 int32)
//go:linkname F_XLogEnsureRecordSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogEnsureRecordSpace
func F_XLogEnsureRecordSpace(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogRegisterBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogRegisterBuffer
func F_XLogRegisterBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogRegisterData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XLogRegisterData
func F_XLogRegisterData(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogRegisterBufData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogRegisterBufData
func F_XLogRegisterBufData(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogInsert
func F_XLogInsert(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_log_newpage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_log_newpage
func F_log_newpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_log_newpage_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_log_newpage_buffer
func F_log_newpage_buffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogReaderFree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogReaderFree
func F_XLogReaderFree(m *base.Module, l0 int32)
//go:linkname F_XLogBeginRead github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogBeginRead
func F_XLogBeginRead(m *base.Module, l0 int32, l1 int64)
//go:linkname F_XLogReadRecord github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogReadRecord
func F_XLogReadRecord(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReadPageInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReadPageInternal
func F_ReadPageInternal(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_GetRecoveryPauseState github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetRecoveryPauseState
func F_GetRecoveryPauseState(m *base.Module) int32
//go:linkname F_XLogInitBufferForRedo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogInitBufferForRedo
func F_XLogInitBufferForRedo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CreateFakeRelcacheEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateFakeRelcacheEntry
func F_CreateFakeRelcacheEntry(m *base.Module, l0 int32) int32
//go:linkname F_forget_invalid_pages github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_forget_invalid_pages
func F_forget_invalid_pages(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_bbsink_forward_archive_contents github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_bbsink_forward_archive_contents
func F_bbsink_forward_archive_contents(m *base.Module, l0 int32, l1 int32)
//go:linkname F_throttle github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_throttle
func F_throttle(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetWalSummaries github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetWalSummaries
func F_GetWalSummaries(m *base.Module, l0 int32, l1 int64, l2 int64) int32
//go:linkname F_heap_getattr_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_getattr_2
func F_heap_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_aclcheck_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_aclcheck_error
func F_aclcheck_error(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_class_aclmask_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_class_aclmask_ext
func F_pg_class_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int64
//go:linkname F_object_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_object_aclcheck
func F_object_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_object_aclmask_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_object_aclmask_ext
func F_object_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int64
//go:linkname F_object_aclcheck_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_object_aclcheck_ext
func F_object_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_pg_attribute_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_attribute_aclcheck
func F_pg_attribute_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_pg_attribute_aclmask_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_attribute_aclmask_ext
func F_pg_attribute_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int64
//go:linkname F_pg_attribute_aclcheck_all github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_attribute_aclcheck_all
func F_pg_attribute_aclcheck_all(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pg_class_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_class_aclcheck
func F_pg_class_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pg_class_aclcheck_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_class_aclcheck_ext
func F_pg_class_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pg_largeobject_aclmask_snapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_largeobject_aclmask_snapshot
func F_pg_largeobject_aclmask_snapshot(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int64
//go:linkname F_object_ownercheck github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_object_ownercheck
func F_object_ownercheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_has_createrole_privilege github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_has_createrole_privilege
func F_has_createrole_privilege(m *base.Module, l0 int32) int32
//go:linkname F_get_user_default_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_user_default_acl
func F_get_user_default_acl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_recordDependencyOnNewAcl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recordDependencyOnNewAcl
func F_recordDependencyOnNewAcl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetNewOidWithIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetNewOidWithIndex
func F_GetNewOidWithIndex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetNewRelFileNumber github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetNewRelFileNumber
func F_GetNewRelFileNumber(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_performDeletion github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_performDeletion
func F_performDeletion(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_new_object_addresses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_new_object_addresses
func F_new_object_addresses(m *base.Module) int32
//go:linkname F_free_object_addresses github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_free_object_addresses
func F_free_object_addresses(m *base.Module, l0 int32)
//go:linkname F_find_expr_references_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_expr_references_walker
func F_find_expr_references_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_record_object_address_dependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_record_object_address_dependencies
func F_record_object_address_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_SystemAttributeByName github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SystemAttributeByName
func F_SystemAttributeByName(m *base.Module, l0 int32) int32
//go:linkname F_heap_create github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_create
func F_heap_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32) int32
//go:linkname F_CheckAttributeNamesTypes github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CheckAttributeNamesTypes
func F_CheckAttributeNamesTypes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_InsertPgAttributeTuples github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InsertPgAttributeTuples
func F_InsertPgAttributeTuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_StoreRelCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_StoreRelCheck
func F_StoreRelCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_SetRelationNumChecks github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SetRelationNumChecks
func F_SetRelationNumChecks(m *base.Module, l0 int32, l1 int32)
//go:linkname F_index_build github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_build
func F_index_build(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_index_set_state_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_set_state_flags
func F_index_set_state_flags(m *base.Module, l0 int32, l1 int32)
//go:linkname F_IndexGetRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_IndexGetRelation
func F_IndexGetRelation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BuildSpeculativeIndexInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BuildSpeculativeIndexInfo
func F_BuildSpeculativeIndexInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_reindex_index github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_reindex_index
func F_reindex_index(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_reindex_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_reindex_relation
func F_reindex_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CatalogOpenIndexes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CatalogOpenIndexes
func F_CatalogOpenIndexes(m *base.Module, l0 int32) int32
//go:linkname F_CatalogCloseIndexes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CatalogCloseIndexes
func F_CatalogCloseIndexes(m *base.Module, l0 int32)
//go:linkname F_CatalogTupleInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CatalogTupleInsert
func F_CatalogTupleInsert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CatalogIndexInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CatalogIndexInsert
func F_CatalogIndexInsert(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CatalogTuplesMultiInsertWithInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CatalogTuplesMultiInsertWithInfo
func F_CatalogTuplesMultiInsertWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CatalogTupleDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CatalogTupleDelete
func F_CatalogTupleDelete(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RangeVarGetRelidExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RangeVarGetRelidExtended
func F_RangeVarGetRelidExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_LookupExplicitNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LookupExplicitNamespace
func F_LookupExplicitNamespace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RangeVarGetCreationNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RangeVarGetCreationNamespace
func F_RangeVarGetCreationNamespace(m *base.Module, l0 int32) int32
//go:linkname F_AccessTempTableNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AccessTempTableNamespace
func F_AccessTempTableNamespace(m *base.Module, l0 int32)
//go:linkname F_get_namespace_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_namespace_oid
func F_get_namespace_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_spcache_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_spcache_insert
func F_spcache_insert(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_isAnyTempNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isAnyTempNamespace
func F_isAnyTempNamespace(m *base.Module, l0 int32) int32
//go:linkname F_RelationIsVisible github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationIsVisible
func F_RelationIsVisible(m *base.Module, l0 int32) int32
//go:linkname F_RelationIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationIsVisibleExt
func F_RelationIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FuncnameGetCandidates github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FuncnameGetCandidates
func F_FuncnameGetCandidates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_DeconstructQualifiedName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DeconstructQualifiedName
func F_DeconstructQualifiedName(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_NameListToString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_NameListToString
func F_NameListToString(m *base.Module, l0 int32) int32
//go:linkname F_OpernameGetOprid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_OpernameGetOprid
func F_OpernameGetOprid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_OpernameGetCandidates github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_OpernameGetCandidates
func F_OpernameGetCandidates(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_OperatorIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OperatorIsVisibleExt
func F_OperatorIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_OpclassIsVisible github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OpclassIsVisible
func F_OpclassIsVisible(m *base.Module, l0 int32) int32
//go:linkname F_ConversionIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConversionIsVisibleExt
func F_ConversionIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TSTemplateIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TSTemplateIsVisibleExt
func F_TSTemplateIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TSConfigIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TSConfigIsVisibleExt
func F_TSConfigIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CheckSetNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CheckSetNamespace
func F_CheckSetNamespace(m *base.Module, l0 int32, l1 int32)
//go:linkname F_makeRangeVarFromNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeRangeVarFromNameList
func F_makeRangeVarFromNameList(m *base.Module, l0 int32) int32
//go:linkname F_checkTempNamespaceStatus github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_checkTempNamespaceStatus
func F_checkTempNamespaceStatus(m *base.Module, l0 int32) int32
//go:linkname F_get_collation_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_collation_oid
func F_get_collation_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RunObjectPostCreateHook github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RunObjectPostCreateHook
func F_RunObjectPostCreateHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RunObjectDropHook github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RunObjectDropHook
func F_RunObjectDropHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RunObjectPostAlterHook github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RunObjectPostAlterHook
func F_RunObjectPostAlterHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_RunNamespaceSearchHook github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RunNamespaceSearchHook
func F_RunNamespaceSearchHook(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_object_address github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_object_address
func F_get_object_address(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_getObjectDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getObjectDescription
func F_getObjectDescription(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_object_catcache_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_object_catcache_oid
func F_get_object_catcache_oid(m *base.Module, l0 int32) int32
//go:linkname F_get_object_catcache_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_object_catcache_name
func F_get_object_catcache_name(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_object_attnum_name
func F_get_object_attnum_name(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_namespace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_object_attnum_namespace
func F_get_object_attnum_namespace(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_owner github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_object_attnum_owner
func F_get_object_attnum_owner(m *base.Module, l0 int32) int32
//go:linkname F_get_object_type github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_object_type
func F_get_object_type(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getObjectDescriptionOids github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_getObjectDescriptionOids
func F_getObjectDescriptionOids(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getObjectTypeDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getObjectTypeDescription
func F_getObjectTypeDescription(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getObjectIdentityParts github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getObjectIdentityParts
func F_getObjectIdentityParts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_partition_parent_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_partition_parent_worker
func F_get_partition_parent_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_partition_ancestors github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_partition_ancestors
func F_get_partition_ancestors(m *base.Module, l0 int32) int32
//go:linkname F_index_get_partition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_get_partition
func F_index_get_partition(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_StoreAttrDefault github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StoreAttrDefault
func F_StoreAttrDefault(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_errdetail_relkind_not_supported github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errdetail_relkind_not_supported
func F_errdetail_relkind_not_supported(m *base.Module, l0 int32)
//go:linkname F_CollationCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CollationCreate
func F_CollationCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32
//go:linkname F_CreateConstraintEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateConstraintEntry
func F_CreateConstraintEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32, l21 int32, l22 int32, l23 int32, l24 int32, l25 int32, l26 int32, l27 int32, l28 int32, l29 int32, l30 int32, l31 int32, l32 int32) int32
//go:linkname F_ConstraintNameIsUsed github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ConstraintNameIsUsed
func F_ConstraintNameIsUsed(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ChooseConstraintName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ChooseConstraintName
func F_ChooseConstraintName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_extractNotNullColumn github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_extractNotNullColumn
func F_extractNotNullColumn(m *base.Module, l0 int32) int32
//go:linkname F_findNotNullConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_findNotNullConstraint
func F_findNotNullConstraint(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RenameConstraintById github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RenameConstraintById
func F_RenameConstraintById(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_relation_constraint_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_relation_constraint_oid
func F_get_relation_constraint_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_getattr_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_getattr_3
func F_heap_getattr_3(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_domain_constraint_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_domain_constraint_oid
func F_get_domain_constraint_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FindFKPeriodOpers github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FindFKPeriodOpers
func F_FindFKPeriodOpers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_heap_getattr_4 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_getattr_4
func F_heap_getattr_4(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_recordDependencyOn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recordDependencyOn
func F_recordDependencyOn(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_recordDependencyOnCurrentExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recordDependencyOnCurrentExtension
func F_recordDependencyOnCurrentExtension(m *base.Module, l0 int32, l1 int32)
//go:linkname F_checkMembershipInCurrentExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_checkMembershipInCurrentExtension
func F_checkMembershipInCurrentExtension(m *base.Module, l0 int32)
//go:linkname F_deleteDependencyRecordsFor github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_deleteDependencyRecordsFor
func F_deleteDependencyRecordsFor(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_deleteDependencyRecordsForClass github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_deleteDependencyRecordsForClass
func F_deleteDependencyRecordsForClass(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getIdentitySequence github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getIdentitySequence
func F_getIdentitySequence(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_index_constraint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_index_constraint
func F_get_index_constraint(m *base.Module, l0 int32) int32
//go:linkname F_find_inheritance_children github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_inheritance_children
func F_find_inheritance_children(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_all_inheritors github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_all_inheritors
func F_find_all_inheritors(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_StoreSingleInheritance github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_StoreSingleInheritance
func F_StoreSingleInheritance(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_NamespaceCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_NamespaceCreate
func F_NamespaceCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_function_parse_error_transpose github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_function_parse_error_transpose
func F_function_parse_error_transpose(m *base.Module, l0 int32) int32
//go:linkname F_GetPublication github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetPublication
func F_GetPublication(m *base.Module, l0 int32) int32
//go:linkname F_GetSchemaPublicationRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetSchemaPublicationRelations
func F_GetSchemaPublicationRelations(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_shdepLockAndCheckObject github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shdepLockAndCheckObject
func F_shdepLockAndCheckObject(m *base.Module, l0 int32, l1 int32)
//go:linkname F_recordDependencyOnOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_recordDependencyOnOwner
func F_recordDependencyOnOwner(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_changeDependencyOnOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_changeDependencyOnOwner
func F_changeDependencyOnOwner(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_shdepDropDependency github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shdepDropDependency
func F_shdepDropDependency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_storeObjectDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_storeObjectDescription
func F_storeObjectDescription(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_deleteSharedDependencyRecordsFor github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_deleteSharedDependencyRecordsFor
func F_deleteSharedDependencyRecordsFor(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RemoveSubscriptionRel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RemoveSubscriptionRel
func F_RemoveSubscriptionRel(m *base.Module, l0 int32, l1 int32)
//go:linkname F_heap_getattr_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_getattr_5
func F_heap_getattr_5(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_moveArrayTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_moveArrayTypeName
func F_moveArrayTypeName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RelFileLocatorSkippingWAL github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelFileLocatorSkippingWAL
func F_RelFileLocatorSkippingWAL(m *base.Module, l0 int32) int32
//go:linkname F_create_toast_table github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_toast_table
func F_create_toast_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_parse_sub_analyze github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parse_sub_analyze
func F_parse_sub_analyze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_transformReturningClause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_transformReturningClause
func F_transformReturningClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_transformLockingClause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_transformLockingClause
func F_transformLockingClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_transformInsertRow github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformInsertRow
func F_transformInsertRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_BuildOnConflictExcludedTargetlist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BuildOnConflictExcludedTargetlist
func F_BuildOnConflictExcludedTargetlist(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_transformSetOperationTree github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformSetOperationTree
func F_transformSetOperationTree(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_makeSortGroupClauseForSetOp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeSortGroupClauseForSetOp
func F_makeSortGroupClauseForSetOp(m *base.Module, l0 int32) int32
//go:linkname F_SystemFuncName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SystemFuncName
func F_SystemFuncName(m *base.Module, l0 int32) int32
//go:linkname F_transformAggregateCall github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformAggregateCall
func F_transformAggregateCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_check_agglevels_and_constraints github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_agglevels_and_constraints
func F_check_agglevels_and_constraints(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformWindowFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_transformWindowFuncCall
func F_transformWindowFuncCall(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_parseCheckAggregates github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_parseCheckAggregates
func F_parseCheckAggregates(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformFromClause github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformFromClause
func F_transformFromClause(m *base.Module, l0 int32, l1 int32)
//go:linkname F_setTargetTable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_setTargetTable
func F_setTargetTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) int32
//go:linkname F_transformWhereClause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_transformWhereClause
func F_transformWhereClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_transformLimitClause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_transformLimitClause
func F_transformLimitClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_flatten_grouping_sets github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_flatten_grouping_sets
func F_flatten_grouping_sets(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_transformGroupingSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformGroupingSet
func F_transformGroupingSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_transformGroupClauseExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformGroupClauseExpr
func F_transformGroupClauseExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_transformSortClause github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformSortClause
func F_transformSortClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_transformWindowDefinitions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_transformWindowDefinitions
func F_transformWindowDefinitions(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_transformDistinctClause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_transformDistinctClause
func F_transformDistinctClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_transformDistinctOnClause github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_transformDistinctOnClause
func F_transformDistinctOnClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_coerce_to_target_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_coerce_to_target_type
func F_coerce_to_target_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_can_coerce_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_can_coerce_type
func F_can_coerce_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_find_coercion_pathway github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_coercion_pathway
func F_find_coercion_pathway(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_typeIsOfTypedTable github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_typeIsOfTypedTable
func F_typeIsOfTypedTable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_coerce_to_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_coerce_to_domain
func F_coerce_to_domain(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_build_coercion_expression github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_build_coercion_expression
func F_build_coercion_expression(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_is_complex_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_is_complex_array
func F_is_complex_array(m *base.Module, l0 int32) int32
//go:linkname F_hide_coercion_node github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hide_coercion_node
func F_hide_coercion_node(m *base.Module, l0 int32)
//go:linkname F_parser_coercion_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parser_coercion_errposition
func F_parser_coercion_errposition(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_select_common_type_from_oids github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_select_common_type_from_oids
func F_select_common_type_from_oids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_coerce_to_specific_type github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_coerce_to_specific_type
func F_coerce_to_specific_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_coerce_null_to_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_coerce_null_to_domain
func F_coerce_null_to_domain(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_select_common_type github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_select_common_type
func F_select_common_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_coerce_to_common_type github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_coerce_to_common_type
func F_coerce_to_common_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_select_common_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_select_common_typmod
func F_select_common_typmod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TypeCategory github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TypeCategory
func F_TypeCategory(m *base.Module, l0 int32) int32
//go:linkname F_assign_query_collations github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_assign_query_collations
func F_assign_query_collations(m *base.Module, l0 int32, l1 int32)
//go:linkname F_assign_collations_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_assign_collations_walker
func F_assign_collations_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assign_list_collations github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_assign_list_collations
func F_assign_list_collations(m *base.Module, l0 int32, l1 int32)
//go:linkname F_select_common_collation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_select_common_collation
func F_select_common_collation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_transformWithClause github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformWithClause
func F_transformWithClause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_transformExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformExpr
func F_transformExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_transformArrayExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_transformArrayExpr
func F_transformArrayExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_transformAExprOp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformAExprOp
func F_transformAExprOp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_transformRowExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformRowExpr
func F_transformRowExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_row_comparison_op github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_row_comparison_op
func F_make_row_comparison_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_make_distinct_op github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_distinct_op
func F_make_distinct_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_transformJsonOutput github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformJsonOutput
func F_transformJsonOutput(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeJsonConstructorExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeJsonConstructorExpr
func F_makeJsonConstructorExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_transformJsonAggConstructor github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformJsonAggConstructor
func F_transformJsonAggConstructor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_transformJsonParseArg github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformJsonParseArg
func F_transformJsonParseArg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ValidJsonBehaviorDefaultExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ValidJsonBehaviorDefaultExpr
func F_ValidJsonBehaviorDefaultExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ParseComplexProjection github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ParseComplexProjection
func F_ParseComplexProjection(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_func_get_detail github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_func_get_detail
func F_func_get_detail(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32) int32
//go:linkname F_func_signature_string github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_func_signature_string
func F_func_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_check_srf_call_placement github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_srf_call_placement
func F_check_srf_call_placement(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_func_select_candidate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_func_select_candidate
func F_func_select_candidate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_fn_arguments github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_make_fn_arguments
func F_make_fn_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_LookupFuncName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LookupFuncName
func F_LookupFuncName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_parsestate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_parsestate
func F_make_parsestate(m *base.Module, l0 int32) int32
//go:linkname F_free_parsestate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_free_parsestate
func F_free_parsestate(m *base.Module, l0 int32)
//go:linkname F_parser_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parser_errposition
func F_parser_errposition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformContainerSubscripts github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformContainerSubscripts
func F_transformContainerSubscripts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_get_sort_group_operators github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_sort_group_operators
func F_get_sort_group_operators(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_make_oper_cache_key github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_oper_cache_key
func F_make_oper_cache_key(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_op_error github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_op_error
func F_op_error(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_left_oper github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_left_oper
func F_left_oper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_make_op github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_op
func F_make_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_make_scalar_array_op github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_scalar_array_op
func F_make_scalar_array_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_setup_parse_variable_parameters github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_setup_parse_variable_parameters
func F_setup_parse_variable_parameters(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_refnameNamespaceItem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_refnameNamespaceItem
func F_refnameNamespaceItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_scanRTEForColumn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_scanRTEForColumn
func F_scanRTEForColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_colNameToVar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_colNameToVar
func F_colNameToVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getRTEPermissionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getRTEPermissionInfo
func F_getRTEPermissionInfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addRTEPermissionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addRTEPermissionInfo
func F_addRTEPermissionInfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addRangeTableEntryForRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addRangeTableEntryForRelation
func F_addRangeTableEntryForRelation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_addRangeTableEntryForSubquery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_addRangeTableEntryForSubquery
func F_addRangeTableEntryForSubquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_addRangeTableEntryForValues github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_addRangeTableEntryForValues
func F_addRangeTableEntryForValues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_addRangeTableEntryForJoin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addRangeTableEntryForJoin
func F_addRangeTableEntryForJoin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32
//go:linkname F_expandRTE github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expandRTE
func F_expandRTE(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_expandNSItemAttrs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_expandNSItemAttrs
func F_expandNSItemAttrs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_attnumAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_attnumAttName
func F_attnumAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_attnumCollationId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_attnumCollationId
func F_attnumCollationId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_errorMissingRTE github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errorMissingRTE
func F_errorMissingRTE(m *base.Module, l0 int32, l1 int32)
//go:linkname F_isQueryUsingTempRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_isQueryUsingTempRelation
func F_isQueryUsingTempRelation(m *base.Module, l0 int32) int32
//go:linkname F_FigureColnameInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FigureColnameInternal
func F_FigureColnameInternal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FigureColname github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FigureColname
func F_FigureColname(m *base.Module, l0 int32) int32
//go:linkname F_transformTargetList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_transformTargetList
func F_transformTargetList(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExpandColumnRefStar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExpandColumnRefStar
func F_ExpandColumnRefStar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExpandRowReference github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExpandRowReference
func F_ExpandRowReference(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_resolveTargetListUnknowns github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_resolveTargetListUnknowns
func F_resolveTargetListUnknowns(m *base.Module, l0 int32, l1 int32)
//go:linkname F_markTargetListOrigins github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_markTargetListOrigins
func F_markTargetListOrigins(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformAssignedExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformAssignedExpr
func F_transformAssignedExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_transformAssignmentIndirection github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformAssignmentIndirection
func F_transformAssignmentIndirection(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32
//go:linkname F_checkInsertTargets github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_checkInsertTargets
func F_checkInsertTargets(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LookupTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LookupTypeName
func F_LookupTypeName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_appendTypeNameToBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendTypeNameToBuffer
func F_appendTypeNameToBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_TypeNameToString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TypeNameToString
func F_TypeNameToString(m *base.Module, l0 int32) int32
//go:linkname F_typenameType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_typenameType
func F_typenameType(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_typenameTypeIdAndMod github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_typenameTypeIdAndMod
func F_typenameTypeIdAndMod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_LookupCollation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LookupCollation
func F_LookupCollation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_typeidType github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_typeidType
func F_typeidType(m *base.Module, l0 int32) int32
//go:linkname F_typeTypeId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_typeTypeId
func F_typeTypeId(m *base.Module, l0 int32) int32
//go:linkname F_typeidTypeRelid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_typeidTypeRelid
func F_typeidTypeRelid(m *base.Module, l0 int32) int32
//go:linkname F_typeOrDomainTypeRelid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_typeOrDomainTypeRelid
func F_typeOrDomainTypeRelid(m *base.Module, l0 int32) int32
//go:linkname F_transformTableConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformTableConstraint
func F_transformTableConstraint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformColumnDefinition github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformColumnDefinition
func F_transformColumnDefinition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformPartitionBound github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_transformPartitionBound
func F_transformPartitionBound(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_scanner_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scanner_yyerror
func F_scanner_yyerror(m *base.Module, l0 int32, l1 int32)
//go:linkname F_scanner_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scanner_errposition
func F_scanner_errposition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_scanner_isspace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scanner_isspace
func F_scanner_isspace(m *base.Module, l0 int32) int32
//go:linkname F_report_namespace_conflict github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_report_namespace_conflict
func F_report_namespace_conflict(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AlterObjectOwner_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AlterObjectOwner_internal
func F_AlterObjectOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_index_am_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_index_am_oid
func F_get_index_am_oid(m *base.Module, l0 int32) int32
//go:linkname F_get_am_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_am_name
func F_get_am_name(m *base.Module, l0 int32) int32
//go:linkname F_NotifyMyFrontEnd github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_NotifyMyFrontEnd
func F_NotifyMyFrontEnd(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_check_index_is_clusterable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_index_is_clusterable
func F_check_index_is_clusterable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_finish_heap_swap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_finish_heap_swap
func F_finish_heap_swap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_IsThereCollationInNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_IsThereCollationInNamespace
func F_IsThereCollationInNamespace(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CopyLimitPrintoutLength github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopyLimitPrintoutLength
func F_CopyLimitPrintoutLength(m *base.Module, l0 int32) int32
//go:linkname F_CopyReadLine github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CopyReadLine
func F_CopyReadLine(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CopyReadAttributesText github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopyReadAttributesText
func F_CopyReadAttributesText(m *base.Module, l0 int32) int32
//go:linkname F_CopyAttributeOutText github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CopyAttributeOutText
func F_CopyAttributeOutText(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CopyAttributeOutCSV github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopyAttributeOutCSV
func F_CopyAttributeOutCSV(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CopySendTextLikeEndOfRow github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopySendTextLikeEndOfRow
func F_CopySendTextLikeEndOfRow(m *base.Module, l0 int32)
//go:linkname F_CopySendEndOfRow github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CopySendEndOfRow
func F_CopySendEndOfRow(m *base.Module, l0 int32)
//go:linkname F_get_database_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_database_name
func F_get_database_name(m *base.Module, l0 int32) int32
//go:linkname F_heap_getattr_6 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_getattr_6
func F_heap_getattr_6(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_defGetString github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_defGetString
func F_defGetString(m *base.Module, l0 int32) int32
//go:linkname F_defGetBoolean github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_defGetBoolean
func F_defGetBoolean(m *base.Module, l0 int32) int32
//go:linkname F_defGetInt32 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_defGetInt32
func F_defGetInt32(m *base.Module, l0 int32) int32
//go:linkname F_AlterEventTriggerOwner_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AlterEventTriggerOwner_internal
func F_AlterEventTriggerOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_EventTriggerSQLDropAddObject github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EventTriggerSQLDropAddObject
func F_EventTriggerSQLDropAddObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_EventTriggerCollectSimpleCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_EventTriggerCollectSimpleCommand
func F_EventTriggerCollectSimpleCommand(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainXMLTag github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainXMLTag
func F_ExplainXMLTag(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainIndentText github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExplainIndentText
func F_ExplainIndentText(m *base.Module, l0 int32)
//go:linkname F_ExplainPropertyText github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExplainPropertyText
func F_ExplainPropertyText(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainPropertyInteger github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainPropertyInteger
func F_ExplainPropertyInteger(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_ExplainOpenGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainOpenGroup
func F_ExplainOpenGroup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExplainCloseGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainCloseGroup
func F_ExplainCloseGroup(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_heap_getattr_7 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_heap_getattr_7
func F_heap_getattr_7(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_AlterForeignServerOwner_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AlterForeignServerOwner_internal
func F_AlterForeignServerOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ComputeIndexAttrs github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ComputeIndexAttrs
func F_ComputeIndexAttrs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32)
//go:linkname F_ChooseRelationName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ChooseRelationName
func F_ChooseRelationName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_GetDefaultOpClass github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetDefaultOpClass
func F_GetDefaultOpClass(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeObjectName github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeObjectName
func F_makeObjectName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReindexRelationConcurrently github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReindexRelationConcurrently
func F_ReindexRelationConcurrently(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LockViewRecurse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LockViewRecurse
func F_LockViewRecurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_opclass_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_opclass_oid
func F_get_opclass_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IsThereOpClassInNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_IsThereOpClassInNamespace
func F_IsThereOpClassInNamespace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_IsThereOpFamilyInNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_IsThereOpFamilyInNamespace
func F_IsThereOpFamilyInNamespace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RelationBuildRowSecurity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationBuildRowSecurity
func F_RelationBuildRowSecurity(m *base.Module, l0 int32)
//go:linkname F_InvalidatePublicationRels github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InvalidatePublicationRels
func F_InvalidatePublicationRels(m *base.Module, l0 int32)
//go:linkname F_AlterPublicationOwner_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AlterPublicationOwner_internal
func F_AlterPublicationOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_init_sequence github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_init_sequence
func F_init_sequence(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_read_seq_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_seq_tuple
func F_read_seq_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_duplicates_in_publist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_duplicates_in_publist
func F_check_duplicates_in_publist(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReplicationSlotDropAtPubNode github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReplicationSlotDropAtPubNode
func F_ReplicationSlotDropAtPubNode(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CheckTableNotInUse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CheckTableNotInUse
func F_CheckTableNotInUse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetForeignKeyCheckTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetForeignKeyCheckTriggers
func F_GetForeignKeyCheckTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_tryAttachPartitionForeignKey github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tryAttachPartitionForeignKey
func F_tryAttachPartitionForeignKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32
//go:linkname F_addFkConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addFkConstraint
func F_addFkConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32)
//go:linkname F_addFkRecurseReferencing github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addFkRecurseReferencing
func F_addFkRecurseReferencing(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32)
//go:linkname F_addFkRecurseReferenced github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_addFkRecurseReferenced
func F_addFkRecurseReferenced(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32)
//go:linkname F_ConstraintImpliedByRelConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ConstraintImpliedByRelConstraint
func F_ConstraintImpliedByRelConstraint(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_truncate_check_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_truncate_check_rel
func F_truncate_check_rel(m *base.Module, l0 int32, l1 int32)
//go:linkname F_renameatt_check github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_renameatt_check
func F_renameatt_check(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RenameRelationInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RenameRelationInternal
func F_RenameRelationInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ATExecChangeOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ATExecChangeOwner
func F_ATExecChangeOwner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ATSimplePermissions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ATSimplePermissions
func F_ATSimplePermissions(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ATExecAlterConstrEnforceability github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ATExecAlterConstrEnforceability
func F_ATExecAlterConstrEnforceability(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32
//go:linkname F_ATAddCheckNNConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ATAddCheckNNConstraint
func F_ATAddCheckNNConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_transformColumnNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformColumnNameList
func F_transformColumnNameList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_PrepareTempTablespaces github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PrepareTempTablespaces
func F_PrepareTempTablespaces(m *base.Module)
//go:linkname F_get_tablespace_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_tablespace_name
func F_get_tablespace_name(m *base.Module, l0 int32) int32
//go:linkname F_CreateTriggerFiringOn github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateTriggerFiringOn
func F_CreateTriggerFiringOn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32)
//go:linkname F_RelationBuildTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationBuildTriggers
func F_RelationBuildTriggers(m *base.Module, l0 int32)
//go:linkname F_before_stmt_triggers_fired github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_before_stmt_triggers_fired
func F_before_stmt_triggers_fired(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TriggerEnabled github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TriggerEnabled
func F_TriggerEnabled(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ExecCallTriggerFunc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecCallTriggerFunc
func F_ExecCallTriggerFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_AfterTriggerSaveEvent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AfterTriggerSaveEvent
func F_AfterTriggerSaveEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)
//go:linkname F_ExecBRDeleteTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecBRDeleteTriggers
func F_ExecBRDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_GetTupleForTrigger github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetTupleForTrigger
func F_GetTupleForTrigger(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_domainAddCheckConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_domainAddCheckConstraint
func F_domainAddCheckConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_checkDomainOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_checkDomainOwner
func F_checkDomainOwner(m *base.Module, l0 int32)
//go:linkname F_validateDomainNotNullConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_validateDomainNotNullConstraint
func F_validateDomainNotNullConstraint(m *base.Module, l0 int32)
//go:linkname F_validateDomainCheckConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_validateDomainCheckConstraint
func F_validateDomainCheckConstraint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AlterTypeOwnerInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AlterTypeOwnerInternal
func F_AlterTypeOwnerInternal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AlterTypeNamespaceInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AlterTypeNamespaceInternal
func F_AlterTypeNamespaceInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_vacuum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_vacuum
func F_vacuum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_vac_update_datfrozenxid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_vac_update_datfrozenxid
func F_vac_update_datfrozenxid(m *base.Module)
//go:linkname F_vacuum_get_cutoffs github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_vacuum_get_cutoffs
func F_vacuum_get_cutoffs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecReScan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecReScan
func F_ExecReScan(m *base.Module, l0 int32)
//go:linkname F_ExecInitExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitExpr
func F_ExecInitExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecInitSubPlanExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecInitSubPlanExpr
func F_ExecInitSubPlanExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecInitQual github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecInitQual
func F_ExecInitQual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecBuildProjectionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecBuildProjectionInfo
func F_ExecBuildProjectionInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ExecBuildUpdateProjection github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecBuildUpdateProjection
func F_ExecBuildUpdateProjection(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ExecPrepareExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecPrepareExpr
func F_ExecPrepareExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecPrepareExprList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecPrepareExprList
func F_ExecPrepareExprList(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecBuildHash32FromAttrs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecBuildHash32FromAttrs
func F_ExecBuildHash32FromAttrs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_ExecBuildGroupingEqual github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBuildGroupingEqual
func F_ExecBuildGroupingEqual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_ExecFilterJunk github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecFilterJunk
func F_ExecFilterJunk(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecutorFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecutorFinish
func F_ExecutorFinish(m *base.Module, l0 int32)
//go:linkname F_ExecutorEnd github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecutorEnd
func F_ExecutorEnd(m *base.Module, l0 int32)
//go:linkname F_CheckValidResultRel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckValidResultRel
func F_CheckValidResultRel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_InitResultRelInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InitResultRelInfo
func F_InitResultRelInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ExecPartitionCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecPartitionCheck
func F_ExecPartitionCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecBuildSlotValueDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBuildSlotValueDescription
func F_ExecBuildSlotValueDescription(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecParallelRetrieveInstrumentation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecParallelRetrieveInstrumentation
func F_ExecParallelRetrieveInstrumentation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecInitPartitionDispatchInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecInitPartitionDispatchInfo
func F_ExecInitPartitionDispatchInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ExecInitRoutingInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecInitRoutingInfo
func F_ExecInitRoutingInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_adjust_partition_colnos_using_map github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_adjust_partition_colnos_using_map
func F_adjust_partition_colnos_using_map(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecFindMatchingSubPlans github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecFindMatchingSubPlans
func F_ExecFindMatchingSubPlans(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecShutdownNode_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecShutdownNode_walker
func F_ExecShutdownNode_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CheckCmdReplicaIdentity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CheckCmdReplicaIdentity
func F_CheckCmdReplicaIdentity(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tupledesc_match github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tupledesc_match
func F_tupledesc_match(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecScan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecScan
func F_ExecScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecStoreMinimalTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecStoreMinimalTuple
func F_ExecStoreMinimalTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MakeTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MakeTupleTableSlot
func F_MakeTupleTableSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecDropSingleTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecDropSingleTupleTableSlot
func F_ExecDropSingleTupleTableSlot(m *base.Module, l0 int32)
//go:linkname F_ExecForceStoreHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecForceStoreHeapTuple
func F_ExecForceStoreHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecFetchSlotHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecFetchSlotHeapTuple
func F_ExecFetchSlotHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecTypeFromTL github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecTypeFromTL
func F_ExecTypeFromTL(m *base.Module, l0 int32) int32
//go:linkname F_ExecInitExtraTupleSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecInitExtraTupleSlot
func F_ExecInitExtraTupleSlot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_slot_getsomeattrs_int github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_slot_getsomeattrs_int
func F_slot_getsomeattrs_int(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BlessTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BlessTupleDesc
func F_BlessTupleDesc(m *base.Module, l0 int32) int32
//go:linkname F_TupleDescGetAttInMetadata github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TupleDescGetAttInMetadata
func F_TupleDescGetAttInMetadata(m *base.Module, l0 int32) int32
//go:linkname F_BuildTupleFromCStrings github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BuildTupleFromCStrings
func F_BuildTupleFromCStrings(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FreeExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FreeExprContext
func F_FreeExprContext(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateExprContext
func F_CreateExprContext(m *base.Module, l0 int32) int32
//go:linkname F_MakePerTupleExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MakePerTupleExprContext
func F_MakePerTupleExprContext(m *base.Module, l0 int32) int32
//go:linkname F_ExecConditionalAssignProjectionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecConditionalAssignProjectionInfo
func F_ExecConditionalAssignProjectionInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RegisterExprContextCallback github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RegisterExprContextCallback
func F_RegisterExprContextCallback(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecGetAllNullSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecGetAllNullSlot
func F_ExecGetAllNullSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetChildToRootMap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecGetChildToRootMap
func F_ExecGetChildToRootMap(m *base.Module, l0 int32) int32
//go:linkname F_ExecGetInsertedCols github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecGetInsertedCols
func F_ExecGetInsertedCols(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetUpdatedCols github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecGetUpdatedCols
func F_ExecGetUpdatedCols(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_prepare_sql_fn_parse_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_prepare_sql_fn_parse_info
func F_prepare_sql_fn_parse_info(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_sql_fn_retval github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_sql_fn_retval
func F_check_sql_fn_retval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_InstrStartNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_InstrStartNode
func F_InstrStartNode(m *base.Module, l0 int32)
//go:linkname F_BufferUsageAccumDiff github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BufferUsageAccumDiff
func F_BufferUsageAccumDiff(m *base.Module, l0 int32, l1 int32)
//go:linkname F_InstrEndLoop github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_InstrEndLoop
func F_InstrEndLoop(m *base.Module, l0 int32)
//go:linkname F_ExecAsyncAppendResponse github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecAsyncAppendResponse
func F_ExecAsyncAppendResponse(m *base.Module, l0 int32)
//go:linkname F_ExecParallelHashMergeCounters github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecParallelHashMergeCounters
func F_ExecParallelHashMergeCounters(m *base.Module, l0 int32)
//go:linkname F_dense_alloc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dense_alloc
func F_dense_alloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecHashIncreaseNumBatches github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecHashIncreaseNumBatches
func F_ExecHashIncreaseNumBatches(m *base.Module, l0 int32)
//go:linkname F_ExecChooseHashTableSize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecChooseHashTableSize
func F_ExecChooseHashTableSize(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_ExecParallelHashJoinSetUpBatches github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecParallelHashJoinSetUpBatches
func F_ExecParallelHashJoinSetUpBatches(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecParallelHashTupleAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecParallelHashTupleAlloc
func F_ExecParallelHashTupleAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecHashJoinSaveTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecHashJoinSaveTuple
func F_ExecHashJoinSaveTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecLookupResultRelByOid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecLookupResultRelByOid
func F_ExecLookupResultRelByOid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecBatchInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBatchInsert
func F_ExecBatchInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_ExecSetParamPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecSetParamPlan
func F_ExecSetParamPlan(m *base.Module, l0 int32, l1 int32)
//go:linkname F_TidListEval github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TidListEval
func F_TidListEval(m *base.Module, l0 int32)
//go:linkname F_begin_partition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_begin_partition
func F_begin_partition(m *base.Module, l0 int32)
//go:linkname F_spool_tuples github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_spool_tuples
func F_spool_tuples(m *base.Module, l0 int32, l1 int64)
//go:linkname F_release_partition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_release_partition
func F_release_partition(m *base.Module, l0 int32)
//go:linkname F_update_frameheadpos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_update_frameheadpos
func F_update_frameheadpos(m *base.Module, l0 int32)
//go:linkname F_window_gettupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_window_gettupleslot
func F_window_gettupleslot(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_WinSetMarkPosition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WinSetMarkPosition
func F_WinSetMarkPosition(m *base.Module, l0 int32, l1 int64)
//go:linkname F_row_is_in_frame github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_row_is_in_frame
func F_row_is_in_frame(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_update_frametailpos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_update_frametailpos
func F_update_frametailpos(m *base.Module, l0 int32)
//go:linkname F_update_grouptailpos github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_update_grouptailpos
func F_update_grouptailpos(m *base.Module, l0 int32)
//go:linkname F_WinGetFuncArgInPartition github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WinGetFuncArgInPartition
func F_WinGetFuncArgInPartition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_WinGetFuncArgCurrent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WinGetFuncArgCurrent
func F_WinGetFuncArgCurrent(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_connect_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SPI_connect_ext
func F_SPI_connect_ext(m *base.Module, l0 int32)
//go:linkname F_SPI_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_finish
func F_SPI_finish(m *base.Module) int32
//go:linkname F__SPI_prepare_oneshot_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__SPI_prepare_oneshot_plan
func F__SPI_prepare_oneshot_plan(m *base.Module, l0 int32, l1 int32)
//go:linkname F__SPI_execute_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__SPI_execute_plan
func F__SPI_execute_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SPI_freetuptable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_freetuptable
func F_SPI_freetuptable(m *base.Module, l0 int32)
//go:linkname F_SPI_execute_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_execute_plan
func F_SPI_execute_plan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_prepare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_prepare
func F_SPI_prepare(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_keepplan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SPI_keepplan
func F_SPI_keepplan(m *base.Module, l0 int32)
//go:linkname F_SPI_getvalue github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_getvalue
func F_SPI_getvalue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_getbinval github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SPI_getbinval
func F_SPI_getbinval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SPI_datumTransfer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SPI_datumTransfer
func F_SPI_datumTransfer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__SPI_cursor_operation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__SPI_cursor_operation
func F__SPI_cursor_operation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_SPI_result_code_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_result_code_string
func F_SPI_result_code_string(m *base.Module, l0 int32) int32
//go:linkname F_get_foreign_data_wrapper_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_foreign_data_wrapper_oid
func F_get_foreign_data_wrapper_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetFdwRoutineByServerId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetFdwRoutineByServerId
func F_GetFdwRoutineByServerId(m *base.Module, l0 int32) int32
//go:linkname F_dshash_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dshash_attach
func F_dshash_attach(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_dshash_find github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dshash_find
func F_dshash_find(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dshash_find_or_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dshash_find_or_insert
func F_dshash_find_or_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dshash_delete_key github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dshash_delete_key
func F_dshash_delete_key(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dshash_delete_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_dshash_delete_entry
func F_dshash_delete_entry(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dshash_release_lock github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dshash_release_lock
func F_dshash_release_lock(m *base.Module, l0 int32, l1 int32)
//go:linkname F_initHyperLogLog github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_initHyperLogLog
func F_initHyperLogLog(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pairingheap_add github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pairingheap_add
func F_pairingheap_add(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pairingheap_remove github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pairingheap_remove
func F_pairingheap_remove(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sanitize_char_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_sanitize_char_2
func F_sanitize_char_2(m *base.Module, l0 int32)
//go:linkname F_read_any_attr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_read_any_attr
func F_read_any_attr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lo_truncate_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_lo_truncate_internal
func F_lo_truncate_internal(m *base.Module, l0 int32, l1 int64)
//go:linkname F_lo_get_fragment_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lo_get_fragment_internal
func F_lo_get_fragment_internal(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_secure_write github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_secure_write
func F_secure_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pq_recvbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_recvbuf
func F_pq_recvbuf(m *base.Module) int32
//go:linkname F_pq_getbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pq_getbytes
func F_pq_getbytes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_startmsgread github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_startmsgread
func F_pq_startmsgread(m *base.Module)
//go:linkname F_pq_sendbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_sendbytes
func F_pq_sendbytes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_sendtext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_sendtext
func F_pq_sendtext(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_sendstring github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_sendstring
func F_pq_sendstring(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pq_getmsgbyte github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_getmsgbyte
func F_pq_getmsgbyte(m *base.Module, l0 int32) int32
//go:linkname F_pq_getmsgint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pq_getmsgint
func F_pq_getmsgint(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_copymsgbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_copymsgbytes
func F_pq_copymsgbytes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_getmsgint64 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pq_getmsgint64
func F_pq_getmsgint64(m *base.Module, l0 int32) int64
//go:linkname F_pq_getmsgfloat8 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pq_getmsgfloat8
func F_pq_getmsgfloat8(m *base.Module, l0 int32) float64
//go:linkname F_pq_getmsgstring github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_getmsgstring
func F_pq_getmsgstring(m *base.Module, l0 int32) int32
//go:linkname F_pq_redirect_to_shm_mq github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_redirect_to_shm_mq
func F_pq_redirect_to_shm_mq(m *base.Module, l0 int32, l1 int32)
//go:linkname F_main github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_main
func F_main(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_bms_equal
func F_bms_equal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_make_singleton github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bms_make_singleton
func F_bms_make_singleton(m *base.Module, l0 int32) int32
//go:linkname F_bms_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_free
func F_bms_free(m *base.Module, l0 int32)
//go:linkname F_bms_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_union
func F_bms_union(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_intersect github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_intersect
func F_bms_intersect(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_difference github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_difference
func F_bms_difference(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_overlap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bms_overlap
func F_bms_overlap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_add_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_add_member
func F_bms_add_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_del_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_del_member
func F_bms_del_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_add_members github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bms_add_members
func F_bms_add_members(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_del_members github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_del_members
func F_bms_del_members(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_join github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_join
func F_bms_join(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_next_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_next_member
func F_bms_next_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_copyObjectImpl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_copyObjectImpl
func F_copyObjectImpl(m *base.Module, l0 int32) int32
//go:linkname F__equalOpExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__equalOpExpr
func F__equalOpExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalSubLink github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__equalSubLink
func F__equalSubLink(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalCaseWhen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__equalCaseWhen
func F__equalCaseWhen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalJsonValueExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__equalJsonValueExpr
func F__equalJsonValueExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalJsonIsPredicate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__equalJsonIsPredicate
func F__equalJsonIsPredicate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalMergeAction github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__equalMergeAction
func F__equalMergeAction(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalResTarget github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__equalResTarget
func F__equalResTarget(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalRangeTableSample github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__equalRangeTableSample
func F__equalRangeTableSample(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalPartitionCmd github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__equalPartitionCmd
func F__equalPartitionCmd(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalJsonObjectConstructor github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__equalJsonObjectConstructor
func F__equalJsonObjectConstructor(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalJsonArrayQueryConstructor github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__equalJsonArrayQueryConstructor
func F__equalJsonArrayQueryConstructor(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalPLAssignStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__equalPLAssignStmt
func F__equalPLAssignStmt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalCreateSchemaStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__equalCreateSchemaStmt
func F__equalCreateSchemaStmt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalAccessPriv github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__equalAccessPriv
func F__equalAccessPriv(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalAlterTableSpaceOptionsStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__equalAlterTableSpaceOptionsStmt
func F__equalAlterTableSpaceOptionsStmt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalCreateExtensionStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__equalCreateExtensionStmt
func F__equalCreateExtensionStmt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalCreateUserMappingStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__equalCreateUserMappingStmt
func F__equalCreateUserMappingStmt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalAlterUserMappingStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__equalAlterUserMappingStmt
func F__equalAlterUserMappingStmt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalCreateEventTrigStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__equalCreateEventTrigStmt
func F__equalCreateEventTrigStmt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__equalCreateRoleStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__equalCreateRoleStmt
func F__equalCreateRoleStmt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetExtensibleNodeMethods github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetExtensibleNodeMethods
func F_GetExtensibleNodeMethods(m *base.Module, l0 int32) int32
//go:linkname F_list_make1_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_make1_impl
func F_list_make1_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_make2_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_list_make2_impl
func F_list_make2_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_make3_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_make3_impl
func F_list_make3_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lappend_int github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_lappend_int
func F_lappend_int(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_lappend_oid
func F_lappend_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lcons github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lcons
func F_lcons(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lcons_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lcons_oid
func F_lcons_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_concat
func F_list_concat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_concat_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_concat_copy
func F_list_concat_copy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_member github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_member
func F_list_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_member_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_member_ptr
func F_list_member_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_nth_cell github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_delete_nth_cell
func F_list_delete_nth_cell(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_free
func F_list_free(m *base.Module, l0 int32)
//go:linkname F_list_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_list_delete
func F_list_delete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_delete_ptr
func F_list_delete_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_first github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_delete_first
func F_list_delete_first(m *base.Module, l0 int32) int32
//go:linkname F_list_delete_last github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_delete_last
func F_list_delete_last(m *base.Module, l0 int32) int32
//go:linkname F_list_free_deep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_free_deep
func F_list_free_deep(m *base.Module, l0 int32)
//go:linkname F_list_copy_head github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_copy_head
func F_list_copy_head(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_copy_tail github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_copy_tail
func F_list_copy_tail(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_sort github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_sort
func F_list_sort(m *base.Module, l0 int32, l1 int32)
//go:linkname F_makeSimpleA_Expr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeSimpleA_Expr
func F_makeSimpleA_Expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeVar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeVar
func F_makeVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_makeVarFromTargetEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeVarFromTargetEntry
func F_makeVarFromTargetEntry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeWholeRowVar github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeWholeRowVar
func F_makeWholeRowVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_makeTargetEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeTargetEntry
func F_makeTargetEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_makeFromExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeFromExpr
func F_makeFromExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeConst github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeConst
func F_makeConst(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_makeBoolConst github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeBoolConst
func F_makeBoolConst(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeBoolExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeBoolExpr
func F_makeBoolExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeAlias github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeAlias
func F_makeAlias(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeRelabelType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeRelabelType
func F_makeRelabelType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeNotNullConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeNotNullConstraint
func F_makeNotNullConstraint(m *base.Module, l0 int32) int32
//go:linkname F_makeTypeNameFromOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeTypeNameFromOid
func F_makeTypeNameFromOid(m *base.Module, l0 int32) int32
//go:linkname F_makeFuncExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeFuncExpr
func F_makeFuncExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeDefElem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeDefElem
func F_makeDefElem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_opclause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_opclause
func F_make_opclause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_andclause github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_andclause
func F_make_andclause(m *base.Module, l0 int32) int32
//go:linkname F_make_orclause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_orclause
func F_make_orclause(m *base.Module, l0 int32) int32
//go:linkname F_makeGroupingSet github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeGroupingSet
func F_makeGroupingSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeVacuumRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeVacuumRelation
func F_makeVacuumRelation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeJsonFormat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeJsonFormat
func F_makeJsonFormat(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeJsonBehavior github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeJsonBehavior
func F_makeJsonBehavior(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeJsonIsPredicate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeJsonIsPredicate
func F_makeJsonIsPredicate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
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
//go:linkname F_expression_returns_set github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expression_returns_set
func F_expression_returns_set(m *base.Module, l0 int32) int32
//go:linkname F_expression_tree_walker_impl_x2especialized_x2e2 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_expression_tree_walker_impl_x2especialized_x2e2
func F_expression_tree_walker_impl_x2especialized_x2e2(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_exprLocation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprLocation
func F_exprLocation(m *base.Module, l0 int32) int32
//go:linkname F_fix_opfuncids github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fix_opfuncids
func F_fix_opfuncids(m *base.Module, l0 int32)
//go:linkname F_expression_tree_walker_impl_x2especialized_x2e1 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expression_tree_walker_impl_x2especialized_x2e1
func F_expression_tree_walker_impl_x2especialized_x2e1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_opfuncid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_opfuncid
func F_set_opfuncid(m *base.Module, l0 int32)
//go:linkname F_check_functions_in_node github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_functions_in_node
func F_check_functions_in_node(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_expression_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expression_tree_walker_impl
func F_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_query_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_query_tree_walker_impl
func F_query_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_expression_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_expression_tree_mutator_impl
func F_expression_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_query_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_query_tree_mutator_impl
func F_query_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_query_or_expression_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_query_or_expression_tree_mutator_impl
func F_query_or_expression_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_planstate_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_planstate_tree_walker_impl
func F_planstate_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__jumbleNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__jumbleNode
func F__jumbleNode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AppendJumble github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AppendJumble
func F_AppendJumble(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_stringToNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_stringToNode
func F_stringToNode(m *base.Module, l0 int32) int32
//go:linkname F_tbm_create_pagetable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tbm_create_pagetable
func F_tbm_create_pagetable(m *base.Module, l0 int32)
//go:linkname F_pagetable_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pagetable_insert
func F_pagetable_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tbm_add_page github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tbm_add_page
func F_tbm_add_page(m *base.Module, l0 int32, l1 int32)
//go:linkname F_makeString github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeString
func F_makeString(m *base.Module, l0 int32) int32
//go:linkname F_merge_clump github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_merge_clump
func F_merge_clump(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_generate_useful_gather_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_generate_useful_gather_paths
func F_generate_useful_gather_paths(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_subquery_is_pushdown_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_subquery_is_pushdown_safe
func F_subquery_is_pushdown_safe(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_clauselist_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clauselist_selectivity
func F_clauselist_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_clause_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_clause_selectivity
func F_clause_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_clamp_row_est github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clamp_row_est
func F_clamp_row_est(m *base.Module, l0 float64) float64
//go:linkname F_cost_qual_eval_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cost_qual_eval_walker
func F_cost_qual_eval_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cost_samplescan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cost_samplescan
func F_cost_samplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_cost_index github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_cost_index
func F_cost_index(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32)
//go:linkname F_cost_qual_eval github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cost_qual_eval
func F_cost_qual_eval(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_indexpath_pages github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_indexpath_pages
func F_get_indexpath_pages(m *base.Module, l0 int32) float64
//go:linkname F_cost_subqueryscan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cost_subqueryscan
func F_cost_subqueryscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_cost_qual_eval_node github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cost_qual_eval_node
func F_cost_qual_eval_node(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_cost_agg github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cost_agg
func F_cost_agg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64, l6 int32, l7 int32, l8 float64, l9 float64, l10 float64, l11 float64)
//go:linkname F_initial_cost_nestloop github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initial_cost_nestloop
func F_initial_cost_nestloop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_initial_cost_mergejoin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initial_cost_mergejoin
func F_initial_cost_mergejoin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_cost_subplan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cost_subplan
func F_cost_subplan(m *base.Module, l0 int32, l1 int32)
//go:linkname F_set_subquery_size_estimates github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_subquery_size_estimates
func F_set_subquery_size_estimates(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_eclass_for_sort_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_eclass_for_sort_expr
func F_get_eclass_for_sort_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_get_common_eclass_indexes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_common_eclass_indexes
func F_get_common_eclass_indexes(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_generate_join_implied_equalities_normal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_join_implied_equalities_normal
func F_generate_join_implied_equalities_normal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ec_search_derived_clause_for_ems github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ec_search_derived_clause_for_ems
func F_ec_search_derived_clause_for_ems(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_add_child_eq_member github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_child_eq_member
func F_add_child_eq_member(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_choose_bitmap_and github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_choose_bitmap_and
func F_choose_bitmap_and(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_loop_count github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_loop_count
func F_get_loop_count(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_match_index_to_operand github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_match_index_to_operand
func F_match_index_to_operand(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_paths_for_OR github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_build_paths_for_OR
func F_build_paths_for_OR(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_try_partial_mergejoin_path github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_try_partial_mergejoin_path
func F_try_partial_mergejoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_get_memoize_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_memoize_path
func F_get_memoize_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_generate_mergejoin_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_generate_mergejoin_paths
func F_generate_mergejoin_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_try_hashjoin_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_try_hashjoin_path
func F_try_hashjoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_try_partial_nestloop_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_try_partial_nestloop_path
func F_try_partial_nestloop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_add_outer_joins_to_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_add_outer_joins_to_relids
func F_add_outer_joins_to_relids(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_canonical_pathkey github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_canonical_pathkey
func F_make_canonical_pathkey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_compare_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_compare_pathkeys
func F_compare_pathkeys(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_build_index_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_index_pathkeys
func F_build_index_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_convert_subquery_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_subquery_pathkeys
func F_convert_subquery_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_build_join_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_join_pathkeys
func F_build_join_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_truncate_useless_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_truncate_useless_pathkeys
func F_truncate_useless_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_innerrel_is_unique_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_innerrel_is_unique_ext
func F_innerrel_is_unique_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_create_plan_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_plan_recurse
func F_create_plan_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_distribute_restrictinfo_to_rels github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_distribute_restrictinfo_to_rels
func F_distribute_restrictinfo_to_rels(m *base.Module, l0 int32, l1 int32)
//go:linkname F_expression_planner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expression_planner
func F_expression_planner(m *base.Module, l0 int32) int32
//go:linkname F_record_plan_function_dependency github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_record_plan_function_dependency
func F_record_plan_function_dependency(m *base.Module, l0 int32, l1 int32)
//go:linkname F_extract_query_dependencies_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_extract_query_dependencies_walker
func F_extract_query_dependencies_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fix_expr_common github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fix_expr_common
func F_fix_expr_common(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fix_param_node github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fix_param_node
func F_fix_param_node(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_convert_testexpr_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_convert_testexpr_mutator
func F_convert_testexpr_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_process_sublinks_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_process_sublinks_mutator
func F_process_sublinks_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hash_ok_operator github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hash_ok_operator
func F_hash_ok_operator(m *base.Module, l0 int32) int32
//go:linkname F_generate_subquery_params github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_generate_subquery_params
func F_generate_subquery_params(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_relids_for_join github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_relids_for_join
func F_get_relids_for_join(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_duplicate_ors github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_duplicate_ors
func F_find_duplicate_ors(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_adjust_appendrel_attrs_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_adjust_appendrel_attrs_mutator
func F_adjust_appendrel_attrs_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_appinfos_by_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_appinfos_by_relids
func F_find_appinfos_by_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_contain_subplans github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_contain_subplans
func F_contain_subplans(m *base.Module, l0 int32) int32
//go:linkname F_contain_volatile_functions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_contain_volatile_functions
func F_contain_volatile_functions(m *base.Module, l0 int32) int32
//go:linkname F_contain_volatile_functions_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_contain_volatile_functions_walker
func F_contain_volatile_functions_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_is_parallel_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_is_parallel_safe
func F_is_parallel_safe(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_contain_nonstrict_functions_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_contain_nonstrict_functions_walker
func F_contain_nonstrict_functions_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_contain_leaked_vars_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_contain_leaked_vars_walker
func F_contain_leaked_vars_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_is_strict_saop github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_is_strict_saop
func F_is_strict_saop(m *base.Module, l0 int32) int32
//go:linkname F_is_pseudo_constant_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_is_pseudo_constant_clause
func F_is_pseudo_constant_clause(m *base.Module, l0 int32) int32
//go:linkname F_eval_const_expressions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_eval_const_expressions
func F_eval_const_expressions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_eval_const_expressions_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_eval_const_expressions_mutator
func F_eval_const_expressions_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_evaluate_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_evaluate_expr
func F_evaluate_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_estimate_expression_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_estimate_expression_value
func F_estimate_expression_value(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_substitute_actual_parameters_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_substitute_actual_parameters_mutator
func F_substitute_actual_parameters_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generate_new_exec_param github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_new_exec_param
func F_generate_new_exec_param(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_add_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_path
func F_add_path(m *base.Module, l0 int32, l1 int32)
//go:linkname F_add_partial_path github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_partial_path
func F_add_partial_path(m *base.Module, l0 int32, l1 int32)
//go:linkname F_create_seqscan_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_create_seqscan_path
func F_create_seqscan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_index_path github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_create_index_path
func F_create_index_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 float64, l10 int32) int32
//go:linkname F_create_append_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_append_path
func F_create_append_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 float64) int32
//go:linkname F_create_material_path github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_material_path
func F_create_material_path(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_unique_path github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_unique_path
func F_create_unique_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_gather_merge_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_gather_merge_path
func F_create_gather_merge_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_create_subqueryscan_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_subqueryscan_path
func F_create_subqueryscan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_calc_non_nestloop_required_outer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_calc_non_nestloop_required_outer
func F_calc_non_nestloop_required_outer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_nestloop_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_nestloop_path
func F_create_nestloop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_create_mergejoin_path github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_create_mergejoin_path
func F_create_mergejoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) int32
//go:linkname F_create_hashjoin_path github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_create_hashjoin_path
func F_create_hashjoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32
//go:linkname F_create_incremental_sort_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_incremental_sort_path
func F_create_incremental_sort_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) int32
//go:linkname F_create_sort_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_sort_path
func F_create_sort_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32
//go:linkname F_path_is_reparameterizable_by_child github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_path_is_reparameterizable_by_child
func F_path_is_reparameterizable_by_child(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_placeholder_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_placeholder_expr
func F_make_placeholder_expr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_placeholder_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_placeholder_info
func F_find_placeholder_info(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_relation_data_width github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_relation_data_width
func F_get_relation_data_width(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_predicate_classify github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_predicate_classify
func F_predicate_classify(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clause_is_strict_for github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_clause_is_strict_for
func F_clause_is_strict_for(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_operator_predicate_proof github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_operator_predicate_proof
func F_operator_predicate_proof(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_setup_simple_rel_arrays github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_setup_simple_rel_arrays
func F_setup_simple_rel_arrays(m *base.Module, l0 int32)
//go:linkname F_build_simple_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_simple_rel
func F_build_simple_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_base_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_base_rel
func F_find_base_rel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_join_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_join_rel
func F_find_join_rel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fetch_upper_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fetch_upper_rel
func F_fetch_upper_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_baserel_parampathinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_baserel_parampathinfo
func F_get_baserel_parampathinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_restrictinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_restrictinfo
func F_make_restrictinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_get_sortgroupref_tle github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_sortgroupref_tle
func F_get_sortgroupref_tle(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_sortgroupclause_tle github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_sortgroupclause_tle
func F_get_sortgroupclause_tle(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_tlist_from_pathtarget github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_tlist_from_pathtarget
func F_make_tlist_from_pathtarget(m *base.Module, l0 int32) int32
//go:linkname F_split_pathtarget_at_srfs_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_split_pathtarget_at_srfs_extended
func F_split_pathtarget_at_srfs_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_pull_varnos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pull_varnos
func F_pull_varnos(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pull_varnos_of_level github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pull_varnos_of_level
func F_pull_varnos_of_level(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pull_varattnos github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pull_varattnos
func F_pull_varattnos(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_contain_var_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_contain_var_clause
func F_contain_var_clause(m *base.Module, l0 int32) int32
//go:linkname F_contain_vars_of_level github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_contain_vars_of_level
func F_contain_vars_of_level(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_locate_var_of_level github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_locate_var_of_level
func F_locate_var_of_level(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_flatten_join_alias_vars github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_flatten_join_alias_vars
func F_flatten_join_alias_vars(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_add_nullingrels_if_needed github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_nullingrels_if_needed
func F_add_nullingrels_if_needed(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_partition_rbound_datum_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_partition_rbound_datum_cmp
func F_partition_rbound_datum_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_partition_range_datum_bsearch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_partition_range_datum_bsearch
func F_partition_range_datum_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_compute_partition_hash_value github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_compute_partition_hash_value
func F_compute_partition_hash_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64
//go:linkname F_DestroyPartitionDirectory github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DestroyPartitionDirectory
func F_DestroyPartitionDirectory(m *base.Module, l0 int32)
//go:linkname F_get_steps_using_prefix_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_steps_using_prefix_recurse
func F_get_steps_using_prefix_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_relation_needs_vacanalyze github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_relation_needs_vacanalyze
func F_relation_needs_vacanalyze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_VacuumUpdateCosts github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_VacuumUpdateCosts
func F_VacuumUpdateCosts(m *base.Module)
//go:linkname F_perform_work_item github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_perform_work_item
func F_perform_work_item(m *base.Module, l0 int32)
//go:linkname F_AuxiliaryProcessMainCommon github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AuxiliaryProcessMainCommon
func F_AuxiliaryProcessMainCommon(m *base.Module)
//go:linkname F_ForgetBackgroundWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ForgetBackgroundWorker
func F_ForgetBackgroundWorker(m *base.Module, l0 int32)
//go:linkname F_BackgroundWorkerUnblockSignals github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BackgroundWorkerUnblockSignals
func F_BackgroundWorkerUnblockSignals(m *base.Module)
//go:linkname F_BackgroundWorkerInitializeConnectionByOid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BackgroundWorkerInitializeConnectionByOid
func F_BackgroundWorkerInitializeConnectionByOid(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetBackgroundWorkerPid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetBackgroundWorkerPid
func F_GetBackgroundWorkerPid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_postmaster_child_launch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_postmaster_child_launch
func F_postmaster_child_launch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ProcessPgArchInterrupts github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ProcessPgArchInterrupts
func F_ProcessPgArchInterrupts(m *base.Module)
//go:linkname F_pgarch_archiveXlog github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgarch_archiveXlog
func F_pgarch_archiveXlog(m *base.Module, l0 int32) int32
//go:linkname F_AssignPostmasterChildSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AssignPostmasterChildSlot
func F_AssignPostmasterChildSlot(m *base.Module, l0 int32) int32
//go:linkname F_ReleasePostmasterChildSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReleasePostmasterChildSlot
func F_ReleasePostmasterChildSlot(m *base.Module, l0 int32) int32
//go:linkname F_ExitPostmaster github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExitPostmaster
func F_ExitPostmaster(m *base.Module, l0 int32)
//go:linkname F_maybe_adjust_io_workers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_maybe_adjust_io_workers
func F_maybe_adjust_io_workers(m *base.Module)
//go:linkname F_StartChildProcess github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_StartChildProcess
func F_StartChildProcess(m *base.Module, l0 int32) int32
//go:linkname F_ConfigurePostmasterWaitSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ConfigurePostmasterWaitSet
func F_ConfigurePostmasterWaitSet(m *base.Module)
//go:linkname F_signal_child github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_signal_child
func F_signal_child(m *base.Module, l0 int32, l1 int32)
//go:linkname F_HandleFatalError github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HandleFatalError
func F_HandleFatalError(m *base.Module, l0 int32)
//go:linkname F_ProcessStartupProcInterrupts github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessStartupProcInterrupts
func F_ProcessStartupProcInterrupts(m *base.Module)
//go:linkname F_GetOldestUnsummarizedLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetOldestUnsummarizedLSN
func F_GetOldestUnsummarizedLSN(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_GetLatestLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetLatestLSN
func F_GetLatestLSN(m *base.Module, l0 int32) int64
//go:linkname F_pg_regcomp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_regcomp
func F_pg_regcomp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_reg_getcolor github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_reg_getcolor
func F_pg_reg_getcolor(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cleartraverse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cleartraverse
func F_cleartraverse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_chrnamed github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_chrnamed
func F_chrnamed(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_markreachable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_markreachable
func F_markreachable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_regerror github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_regerror
func F_pg_regerror(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_regexec github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_regexec
func F_pg_regexec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_newdfa github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_newdfa
func F_newdfa(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_shortest github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shortest
func F_shortest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_miss github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_miss
func F_miss(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_getsubdfa github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getsubdfa
func F_getsubdfa(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zaptreesubs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_zaptreesubs
func F_zaptreesubs(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getvacant github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getvacant
func F_getvacant(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_regfree github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_regfree
func F_pg_regfree(m *base.Module, l0 int32)
//go:linkname F_logicalrep_workers_find github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_logicalrep_workers_find
func F_logicalrep_workers_find(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_logicalrep_worker_wakeup github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_logicalrep_worker_wakeup
func F_logicalrep_worker_wakeup(m *base.Module, l0 int32)
//go:linkname F_ApplyLauncherForgetWorkerStartTime github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ApplyLauncherForgetWorkerStartTime
func F_ApplyLauncherForgetWorkerStartTime(m *base.Module, l0 int32)
//go:linkname F_StartupDecodingContext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_StartupDecodingContext
func F_StartupDecodingContext(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_LogicalConfirmReceivedLocation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LogicalConfirmReceivedLocation
func F_LogicalConfirmReceivedLocation(m *base.Module, l0 int64)
//go:linkname F_LogicalIncreaseRestartDecodingForSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LogicalIncreaseRestartDecodingForSlot
func F_LogicalIncreaseRestartDecodingForSlot(m *base.Module, l0 int64, l1 int64)
//go:linkname F_pg_logical_slot_get_changes_guts github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_logical_slot_get_changes_guts
func F_pg_logical_slot_get_changes_guts(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_replorigin_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_replorigin_by_name
func F_replorigin_by_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replorigin_drop_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_replorigin_drop_by_name
func F_replorigin_drop_by_name(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_logicalrep_relmap_free_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_logicalrep_relmap_free_entry
func F_logicalrep_relmap_free_entry(m *base.Module, l0 int32)
//go:linkname F_ReorderBufferFree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReorderBufferFree
func F_ReorderBufferFree(m *base.Module, l0 int32)
//go:linkname F_ReorderBufferFreeChange github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReorderBufferFreeChange
func F_ReorderBufferFreeChange(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReorderBufferSerializeTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReorderBufferSerializeTXN
func F_ReorderBufferSerializeTXN(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferAssignChild github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferAssignChild
func F_ReorderBufferAssignChild(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64)
//go:linkname F_ReorderBufferTruncateTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReorderBufferTruncateTXN
func F_ReorderBufferTruncateTXN(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReorderBufferCleanupTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReorderBufferCleanupTXN
func F_ReorderBufferCleanupTXN(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferToastReset github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReorderBufferToastReset
func F_ReorderBufferToastReset(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferProcessXid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReorderBufferProcessXid
func F_ReorderBufferProcessXid(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_ReorderBufferCopySnap github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferCopySnap
func F_ReorderBufferCopySnap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ReorderBufferIterTXNFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferIterTXNFinish
func F_ReorderBufferIterTXNFinish(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ValidateSlotSyncParams github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ValidateSlotSyncParams
func F_ValidateSlotSyncParams(m *base.Module, l0 int32) int32
//go:linkname F_check_and_set_sync_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_check_and_set_sync_info
func F_check_and_set_sync_info(m *base.Module, l0 int32)
//go:linkname F_validate_remote_info github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_validate_remote_info
func F_validate_remote_info(m *base.Module, l0 int32)
//go:linkname F_synchronize_slots github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_synchronize_slots
func F_synchronize_slots(m *base.Module, l0 int32) int32
//go:linkname F_slotsync_failure_callback github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_slotsync_failure_callback
func F_slotsync_failure_callback(m *base.Module, l0 int32, l1 int32)
//go:linkname F_FreeSnapshotBuilder github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeSnapshotBuilder
func F_FreeSnapshotBuilder(m *base.Module, l0 int32)
//go:linkname F_SnapBuildSnapDecRefcount github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SnapBuildSnapDecRefcount
func F_SnapBuildSnapDecRefcount(m *base.Module, l0 int32)
//go:linkname F_SnapBuildWaitSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SnapBuildWaitSnapshot
func F_SnapBuildWaitSnapshot(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SnapBuildRestore github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SnapBuildRestore
func F_SnapBuildRestore(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_ReplicationSlotNameForTablesync github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotNameForTablesync
func F_ReplicationSlotNameForTablesync(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReplicationOriginNameForLogicalRep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReplicationOriginNameForLogicalRep
func F_ReplicationOriginNameForLogicalRep(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_stream_start_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_stream_start_internal
func F_stream_start_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_subxact_info_write github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_subxact_info_write
func F_subxact_info_write(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LogicalRepWorkersWakeupAtCommit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LogicalRepWorkersWakeupAtCommit
func F_LogicalRepWorkersWakeupAtCommit(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplicationSlotCleanup
func F_ReplicationSlotCleanup(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotValidateName github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReplicationSlotValidateName
func F_ReplicationSlotValidateName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SaveSlotToPath github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SaveSlotToPath
func F_SaveSlotToPath(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReplicationSlotMarkDirty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReplicationSlotMarkDirty
func F_ReplicationSlotMarkDirty(m *base.Module)
//go:linkname F_ReplicationSlotSave github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotSave
func F_ReplicationSlotSave(m *base.Module)
//go:linkname F_CheckSlotPermissions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckSlotPermissions
func F_CheckSlotPermissions(m *base.Module)
//go:linkname F_StandbySlotsHaveCaughtup github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StandbySlotsHaveCaughtup
func F_StandbySlotsHaveCaughtup(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_setRuleCheckAsUser github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_setRuleCheckAsUser
func F_setRuleCheckAsUser(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AcquireRewriteLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AcquireRewriteLocks
func F_AcquireRewriteLocks(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_build_column_default github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_column_default
func F_build_column_default(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_view_query github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_view_query
func F_get_view_query(m *base.Module, l0 int32) int32
//go:linkname F_view_has_instead_trigger github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_view_has_instead_trigger
func F_view_has_instead_trigger(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_relation_is_updatable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_relation_is_updatable
func F_relation_is_updatable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_adjust_view_column_set github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_adjust_view_column_set
func F_adjust_view_column_set(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_error_view_not_updatable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_error_view_not_updatable
func F_error_view_not_updatable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_build_generation_expression github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_generation_expression
func F_build_generation_expression(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rewriteTargetListIU github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_rewriteTargetListIU
func F_rewriteTargetListIU(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_matchLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_matchLocks
func F_matchLocks(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_markQueryForLocking github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_markQueryForLocking
func F_markQueryForLocking(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_checkExprHasSubLink github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_checkExprHasSubLink
func F_checkExprHasSubLink(m *base.Module, l0 int32) int32
//go:linkname F_CombineRangeTables github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CombineRangeTables
func F_CombineRangeTables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_OffsetVarNodes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_OffsetVarNodes
func F_OffsetVarNodes(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ChangeVarNodes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ChangeVarNodes
func F_ChangeVarNodes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_IncrementVarSublevelsUp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_IncrementVarSublevelsUp
func F_IncrementVarSublevelsUp(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_rangeTableEntry_used github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_rangeTableEntry_used
func F_rangeTableEntry_used(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getInsertSelectQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getInsertSelectQuery
func F_getInsertSelectQuery(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AddQual github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AddQual
func F_AddQual(m *base.Module, l0 int32, l1 int32)
//go:linkname F_add_nulling_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_nulling_relids
func F_add_nulling_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_map_variable_attnos github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_map_variable_attnos
func F_map_variable_attnos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ReplaceVarFromTargetList github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplaceVarFromTargetList
func F_ReplaceVarFromTargetList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ReplaceVarsFromTargetList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplaceVarsFromTargetList
func F_ReplaceVarsFromTargetList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_add_security_quals github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_security_quals
func F_add_security_quals(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_add_with_check_options github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_with_check_options
func F_add_with_check_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_statext_is_kind_built github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_statext_is_kind_built
func F_statext_is_kind_built(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgaio_io_update_state github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgaio_io_update_state
func F_pgaio_io_update_state(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgaio_io_process_completion github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgaio_io_process_completion
func F_pgaio_io_process_completion(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgaio_error_cleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgaio_error_cleanup
func F_pgaio_error_cleanup(m *base.Module)
//go:linkname F_pgaio_closing_fd github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgaio_closing_fd
func F_pgaio_closing_fd(m *base.Module, l0 int32)
//go:linkname F_read_stream_begin_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_read_stream_begin_impl
func F_read_stream_begin_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_read_stream_next_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_read_stream_next_buffer
func F_read_stream_next_buffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_read_stream_end github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_stream_end
func F_read_stream_end(m *base.Module, l0 int32)
//go:linkname F_GetPrivateRefCountEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetPrivateRefCountEntry
func F_GetPrivateRefCountEntry(m *base.Module, l0 int32) int32
//go:linkname F_ReadBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReadBuffer
func F_ReadBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReadBufferWithoutRelcache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReadBufferWithoutRelcache
func F_ReadBufferWithoutRelcache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_StartBufferIO github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_StartBufferIO
func F_StartBufferIO(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_TerminateBufferIO github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TerminateBufferIO
func F_TerminateBufferIO(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ReleaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReleaseBuffer
func F_ReleaseBuffer(m *base.Module, l0 int32)
//go:linkname F_MarkBufferDirty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MarkBufferDirty
func F_MarkBufferDirty(m *base.Module, l0 int32)
//go:linkname F_ReleaseAndReadBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReleaseAndReadBuffer
func F_ReleaseAndReadBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RelationGetNumberOfBlocksInFork github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetNumberOfBlocksInFork
func F_RelationGetNumberOfBlocksInFork(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_UnlockReleaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UnlockReleaseBuffer
func F_UnlockReleaseBuffer(m *base.Module, l0 int32)
//go:linkname F_LockBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockBuffer
func F_LockBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_MarkBufferDirtyHint github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MarkBufferDirtyHint
func F_MarkBufferDirtyHint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LockBufferForCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockBufferForCleanup
func F_LockBufferForCleanup(m *base.Module, l0 int32)
//go:linkname F_IsBufferCleanupOK github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_IsBufferCleanupOK
func F_IsBufferCleanupOK(m *base.Module, l0 int32) int32
//go:linkname F_GetAccessStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetAccessStrategy
func F_GetAccessStrategy(m *base.Module, l0 int32) int32
//go:linkname F_GetAccessStrategyWithSize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetAccessStrategyWithSize
func F_GetAccessStrategyWithSize(m *base.Module, l0 int32) int32
//go:linkname F_InitLocalBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InitLocalBuffers
func F_InitLocalBuffers(m *base.Module)
//go:linkname F_GetLocalVictimBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetLocalVictimBuffer
func F_GetLocalVictimBuffer(m *base.Module) int32
//go:linkname F_StartLocalBufferIO github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StartLocalBufferIO
func F_StartLocalBufferIO(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BufFileOpenFileSet github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BufFileOpenFileSet
func F_BufFileOpenFileSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_BufFileDumpBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BufFileDumpBuffer
func F_BufFileDumpBuffer(m *base.Module, l0 int32)
//go:linkname F_BufFileClose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BufFileClose
func F_BufFileClose(m *base.Module, l0 int32)
//go:linkname F_BufFileWrite github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufFileWrite
func F_BufFileWrite(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_BufFileSeekBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufFileSeekBlock
func F_BufFileSeekBlock(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_fsync_fname github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fsync_fname
func F_fsync_fname(m *base.Module, l0 int32, l1 int32)
//go:linkname F_OpenTransientFilePerm github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_OpenTransientFilePerm
func F_OpenTransientFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_durable_rename github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_durable_rename
func F_durable_rename(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FreeDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeDesc
func F_FreeDesc(m *base.Module, l0 int32) int32
//go:linkname F_durable_unlink github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_durable_unlink
func F_durable_unlink(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BasicOpenFilePerm github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BasicOpenFilePerm
func F_BasicOpenFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LruDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LruDelete
func F_LruDelete(m *base.Module, l0 int32)
//go:linkname F_ReserveExternalFD github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReserveExternalFD
func F_ReserveExternalFD(m *base.Module)
//go:linkname F_PathNameOpenFilePerm github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PathNameOpenFilePerm
func F_PathNameOpenFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_walkdir github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_walkdir
func F_walkdir(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ReadDirExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReadDirExtended
func F_ReadDirExtended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FreeDir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeDir
func F_FreeDir(m *base.Module, l0 int32)
//go:linkname F_TempTablespacePath github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TempTablespacePath
func F_TempTablespacePath(m *base.Module, l0 int32, l1 int32)
//go:linkname F_FileWriteV github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FileWriteV
func F_FileWriteV(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_FileSync github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FileSync
func F_FileSync(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FileSize github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FileSize
func F_FileSize(m *base.Module, l0 int32) int64
//go:linkname F_OpenPipeStream github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_OpenPipeStream
func F_OpenPipeStream(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReadDir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadDir
func F_ReadDir(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ClosePipeStream github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ClosePipeStream
func F_ClosePipeStream(m *base.Module, l0 int32) int32
//go:linkname F_AtEOXact_Files github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AtEOXact_Files
func F_AtEOXact_Files(m *base.Module, l0 int32)
//go:linkname F_RemovePgTempFiles github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RemovePgTempFiles
func F_RemovePgTempFiles(m *base.Module)
//go:linkname F_FileSetDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FileSetDelete
func F_FileSetDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SharedFileSetInit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SharedFileSetInit
func F_SharedFileSetInit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SharedFileSetAttach github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SharedFileSetAttach
func F_SharedFileSetAttach(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fsm_readbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fsm_readbuf
func F_fsm_readbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FreeSpaceMapPrepareTruncateRel github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FreeSpaceMapPrepareTruncateRel
func F_FreeSpaceMapPrepareTruncateRel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FreeSpaceMapVacuumRange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreeSpaceMapVacuumRange
func F_FreeSpaceMapVacuumRange(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_BarrierArriveAndWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BarrierArriveAndWait
func F_BarrierArriveAndWait(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsm_pin_mapping github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dsm_pin_mapping
func F_dsm_pin_mapping(m *base.Module, l0 int32)
//go:linkname F_dsm_impl_op github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dsm_impl_op
func F_dsm_impl_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_proc_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_proc_exit
func F_proc_exit(m *base.Module, l0 int32)
//go:linkname F_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shmem_exit
func F_shmem_exit(m *base.Module, l0 int32)
//go:linkname F_before_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_before_shmem_exit
func F_before_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_on_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_on_shmem_exit
func F_on_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_cancel_before_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cancel_before_shmem_exit
func F_cancel_before_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CalculateShmemSize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CalculateShmemSize
func F_CalculateShmemSize(m *base.Module, l0 int32) int32
//go:linkname F_CreateSharedMemoryAndSemaphores github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CreateSharedMemoryAndSemaphores
func F_CreateSharedMemoryAndSemaphores(m *base.Module)
//go:linkname F_InitializeLatchWaitSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_InitializeLatchWaitSet
func F_InitializeLatchWaitSet(m *base.Module)
//go:linkname F_WaitLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_WaitLatch
func F_WaitLatch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SetLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SetLatch
func F_SetLatch(m *base.Module, l0 int32)
//go:linkname F_SendPostmasterSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SendPostmasterSignal
func F_SendPostmasterSignal(m *base.Module, l0 int32)
//go:linkname F_PostmasterIsAliveInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PostmasterIsAliveInternal
func F_PostmasterIsAliveInternal(m *base.Module) int32
//go:linkname F_KnownAssignedXidsDisplay github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_KnownAssignedXidsDisplay
func F_KnownAssignedXidsDisplay(m *base.Module, l0 int32)
//go:linkname F_KnownAssignedXidsCompress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_KnownAssignedXidsCompress
func F_KnownAssignedXidsCompress(m *base.Module, l0 int32, l1 int32)
//go:linkname F_KnownAssignedXidsRemoveTree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_KnownAssignedXidsRemoveTree
func F_KnownAssignedXidsRemoveTree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetSnapshotData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetSnapshotData
func F_GetSnapshotData(m *base.Module, l0 int32) int32
//go:linkname F_ProcArrayInstallImportedXmin github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcArrayInstallImportedXmin
func F_ProcArrayInstallImportedXmin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BackendPidGetProc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BackendPidGetProc
func F_BackendPidGetProc(m *base.Module, l0 int32) int32
//go:linkname F_GlobalVisTestIsRemovableXid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GlobalVisTestIsRemovableXid
func F_GlobalVisTestIsRemovableXid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_shm_mq_get_sender github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shm_mq_get_sender
func F_shm_mq_get_sender(m *base.Module, l0 int32) int32
//go:linkname F_shm_mq_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_mq_attach
func F_shm_mq_attach(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_shm_mq_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_mq_detach
func F_shm_mq_detach(m *base.Module, l0 int32)
//go:linkname F_shm_toc_allocate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_shm_toc_allocate
func F_shm_toc_allocate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_shm_toc_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_toc_insert
func F_shm_toc_insert(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_ShmemInitStruct github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ShmemInitStruct
func F_ShmemInitStruct(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_add_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_size
func F_add_size(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReceiveSharedInvalidMessages github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReceiveSharedInvalidMessages
func F_ReceiveSharedInvalidMessages(m *base.Module)
//go:linkname F_CreateWaitEventSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateWaitEventSet
func F_CreateWaitEventSet(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FreeWaitEventSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeWaitEventSet
func F_FreeWaitEventSet(m *base.Module, l0 int32)
//go:linkname F_AddWaitEventToSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AddWaitEventToSet
func F_AddWaitEventToSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ModifyWaitEvent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ModifyWaitEvent
func F_ModifyWaitEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_WaitEventSetWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WaitEventSetWait
func F_WaitEventSetWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_inv_seek github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_inv_seek
func F_inv_seek(m *base.Module, l0 int32, l1 int64, l2 int32) int64
//go:linkname F_inv_write github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_inv_write
func F_inv_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ConditionVariableInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ConditionVariableInit
func F_ConditionVariableInit(m *base.Module, l0 int32)
//go:linkname F_ConditionVariableCancelSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ConditionVariableCancelSleep
func F_ConditionVariableCancelSleep(m *base.Module)
//go:linkname F_ConditionVariableTimedSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ConditionVariableTimedSleep
func F_ConditionVariableTimedSleep(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockRelationOid
func F_LockRelationOid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ConditionalLockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConditionalLockRelationOid
func F_ConditionalLockRelationOid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LockTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockTuple
func F_LockTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockDatabaseObject github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LockDatabaseObject
func F_LockDatabaseObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockSharedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockSharedObject
func F_LockSharedObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_UnlockSharedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_UnlockSharedObject
func F_UnlockSharedObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockAcquire
func F_LockAcquire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_LockAcquireExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockAcquireExtended
func F_LockAcquireExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_LockRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LockRelease
func F_LockRelease(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetSingleProcBlockerStatusData github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetSingleProcBlockerStatusData
func F_GetSingleProcBlockerStatusData(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LWLockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockAcquire
func F_LWLockAcquire(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LWLockConditionalAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockConditionalAcquire
func F_LWLockConditionalAcquire(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LWLockRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LWLockRelease
func F_LWLockRelease(m *base.Module, l0 int32)
//go:linkname F_LWLockReleaseAll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockReleaseAll
func F_LWLockReleaseAll(m *base.Module)
//go:linkname F_GetSerializableTransactionSnapshotInt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetSerializableTransactionSnapshotInt
func F_GetSerializableTransactionSnapshotInt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReleasePredicateLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReleasePredicateLocks
func F_ReleasePredicateLocks(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreatePredicateLock github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreatePredicateLock
func F_CreatePredicateLock(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_DecrementParentLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecrementParentLocks
func F_DecrementParentLocks(m *base.Module, l0 int32)
//go:linkname F_PredicateLockTID github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PredicateLockTID
func F_PredicateLockTID(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_TransferPredicateLocksToHeapRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransferPredicateLocksToHeapRelation
func F_TransferPredicateLocksToHeapRelation(m *base.Module, l0 int32)
//go:linkname F_PredicateLockPageSplit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PredicateLockPageSplit
func F_PredicateLockPageSplit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CheckForSerializableConflictIn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckForSerializableConflictIn
func F_CheckForSerializableConflictIn(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_InitProcess github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InitProcess
func F_InitProcess(m *base.Module)
//go:linkname F_s_lock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_s_lock
func F_s_lock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_perform_spin_delay github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_perform_spin_delay
func F_perform_spin_delay(m *base.Module, l0 int32)
//go:linkname F_PageInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PageInit
func F_PageInit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PageIsVerified github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PageIsVerified
func F_PageIsVerified(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_PageAddItemExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PageAddItemExtended
func F_PageAddItemExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_PageIndexMultiDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PageIndexMultiDelete
func F_PageIndexMultiDelete(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PageIndexTupleDeleteNoCompact github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PageIndexTupleDeleteNoCompact
func F_PageIndexTupleDeleteNoCompact(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_checksum_page github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_checksum_page
func F_pg_checksum_page(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgr_bulk_start_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgr_bulk_start_rel
func F_smgr_bulk_start_rel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgr_bulk_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_smgr_bulk_finish
func F_smgr_bulk_finish(m *base.Module, l0 int32)
//go:linkname F_smgr_bulk_flush github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_smgr_bulk_flush
func F_smgr_bulk_flush(m *base.Module, l0 int32)
//go:linkname F_smgr_bulk_get_buf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgr_bulk_get_buf
func F_smgr_bulk_get_buf(m *base.Module, l0 int32) int32
//go:linkname F__mdfd_getseg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__mdfd_getseg
func F__mdfd_getseg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_smgropen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgropen
func F_smgropen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgrclose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrclose
func F_smgrclose(m *base.Module, l0 int32)
//go:linkname F_smgrreleaseall github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrreleaseall
func F_smgrreleaseall(m *base.Module)
//go:linkname F_smgrexists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrexists
func F_smgrexists(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgrcreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgrcreate
func F_smgrcreate(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_smgrextend github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrextend
func F_smgrextend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_smgrwritev github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_smgrwritev
func F_smgrwritev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_smgrnblocks github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_smgrnblocks
func F_smgrnblocks(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgrtruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrtruncate
func F_smgrtruncate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_RememberSyncRequest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RememberSyncRequest
func F_RememberSyncRequest(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RegisterSyncRequest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RegisterSyncRequest
func F_RegisterSyncRequest(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CreateDestReceiver github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CreateDestReceiver
func F_CreateDestReceiver(m *base.Module, l0 int32) int32
//go:linkname F_ProcessClientReadInterrupt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ProcessClientReadInterrupt
func F_ProcessClientReadInterrupt(m *base.Module, l0 int32)
//go:linkname F_ProcessInterrupts github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ProcessInterrupts
func F_ProcessInterrupts(m *base.Module)
//go:linkname F_pg_parse_query github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_parse_query
func F_pg_parse_query(m *base.Module, l0 int32) int32
//go:linkname F_pg_rewrite_query github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_rewrite_query
func F_pg_rewrite_query(m *base.Module, l0 int32) int32
//go:linkname F_pg_analyze_and_rewrite_withcb github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_analyze_and_rewrite_withcb
func F_pg_analyze_and_rewrite_withcb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_set_debug_options github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_set_debug_options
func F_set_debug_options(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_FreeQueryDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeQueryDesc
func F_FreeQueryDesc(m *base.Module, l0 int32)
//go:linkname F_PreventCommandIfReadOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreventCommandIfReadOnly
func F_PreventCommandIfReadOnly(m *base.Module, l0 int32)
//go:linkname F_PreventCommandDuringRecovery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PreventCommandDuringRecovery
func F_PreventCommandDuringRecovery(m *base.Module, l0 int32)
//go:linkname F_ProcessUtility github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ProcessUtility
func F_ProcessUtility(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_findwrd github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_findwrd
func F_findwrd(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_tsvector github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_tsvector
func F_make_tsvector(m *base.Module, l0 int32) int32
//go:linkname F_add_to_tsvector github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_to_tsvector
func F_add_to_tsvector(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tsearch_readline_begin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tsearch_readline_begin
func F_tsearch_readline_begin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tsearch_readline github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tsearch_readline
func F_tsearch_readline(m *base.Module, l0 int32) int32
//go:linkname F_tsearch_readline_end github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tsearch_readline_end
func F_tsearch_readline_end(m *base.Module, l0 int32)
//go:linkname F_searchstoplist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_searchstoplist
func F_searchstoplist(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_progress_parallel_incr_param github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_progress_parallel_incr_param
func F_pgstat_progress_parallel_incr_param(m *base.Module, l0 int32, l1 int64)
//go:linkname F_pgstat_clear_backend_activity_snapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_clear_backend_activity_snapshot
func F_pgstat_clear_backend_activity_snapshot(m *base.Module)
//go:linkname F_pgstat_report_activity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_report_activity
func F_pgstat_report_activity(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_clip_activity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_clip_activity
func F_pgstat_clip_activity(m *base.Module, l0 int32) int32
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
//go:linkname F_pgstat_fetch_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_fetch_entry
func F_pgstat_fetch_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pgstat_snapshot_fixed github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_snapshot_fixed
func F_pgstat_snapshot_fixed(m *base.Module, l0 int32)
//go:linkname F_pgstat_prep_pending_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_prep_pending_entry
func F_pgstat_prep_pending_entry(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pgstat_count_backend_io_op github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_count_backend_io_op
func F_pgstat_count_backend_io_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64)
//go:linkname F_pgstat_fetch_stat_checkpointer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_fetch_stat_checkpointer
func F_pgstat_fetch_stat_checkpointer(m *base.Module) int32
//go:linkname F_pgstat_prepare_report_checksum_failure github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_prepare_report_checksum_failure
func F_pgstat_prepare_report_checksum_failure(m *base.Module, l0 int32)
//go:linkname F_pgstat_report_checksum_failures_in_db github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_report_checksum_failures_in_db
func F_pgstat_report_checksum_failures_in_db(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_report_tempfile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_report_tempfile
func F_pgstat_report_tempfile(m *base.Module, l0 int32)
//go:linkname F_pgstat_fetch_stat_dbentry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_fetch_stat_dbentry
func F_pgstat_fetch_stat_dbentry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_init_function_usage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_init_function_usage
func F_pgstat_init_function_usage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_fetch_stat_funcentry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_fetch_stat_funcentry
func F_pgstat_fetch_stat_funcentry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_fetch_stat_tabentry_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_fetch_stat_tabentry_ext
func F_pgstat_fetch_stat_tabentry_ext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_count_heap_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_count_heap_insert
func F_pgstat_count_heap_insert(m *base.Module, l0 int32, l1 int64)
//go:linkname F_pgstat_fetch_stat_tabentry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_fetch_stat_tabentry
func F_pgstat_fetch_stat_tabentry(m *base.Module, l0 int32) int32
//go:linkname F_find_tabstat_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_tabstat_entry
func F_find_tabstat_entry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_unlock_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_unlock_entry
func F_pgstat_unlock_entry(m *base.Module, l0 int32)
//go:linkname F_pgstat_get_entry_ref_locked github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_get_entry_ref_locked
func F_pgstat_get_entry_ref_locked(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pgstat_drop_transactional github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_drop_transactional
func F_pgstat_drop_transactional(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_aclmask github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_aclmask
func F_aclmask(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int64
//go:linkname F_getid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getid
func F_getid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_acldefault github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_acldefault
func F_acldefault(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_aclnewowner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_aclnewowner
func F_aclnewowner(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_has_privs_of_role github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_has_privs_of_role
func F_has_privs_of_role(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_convert_any_priv_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_any_priv_string
func F_convert_any_priv_string(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_get_role_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_role_oid
func F_get_role_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_member_can_set_role github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_member_can_set_role
func F_member_can_set_role(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_can_set_role github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_can_set_role
func F_check_can_set_role(m *base.Module, l0 int32, l1 int32)
//go:linkname F_is_member_of_role_nosuper github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_is_member_of_role_nosuper
func F_is_member_of_role_nosuper(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_is_admin_of_role github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_is_admin_of_role
func F_is_admin_of_role(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rolespec_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_rolespec_oid
func F_get_rolespec_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rolespec_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_rolespec_tuple
func F_get_rolespec_tuple(m *base.Module, l0 int32) int32
//go:linkname F_DatumGetAnyArrayP github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DatumGetAnyArrayP
func F_DatumGetAnyArrayP(m *base.Module, l0 int32) int32
//go:linkname F_fetch_array_arg_replace_nulls github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fetch_array_arg_replace_nulls
func F_fetch_array_arg_replace_nulls(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CopyArrayEls github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CopyArrayEls
func F_CopyArrayEls(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_construct_empty_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_empty_array
func F_construct_empty_array(m *base.Module, l0 int32) int32
//go:linkname F_array_iter_next github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_iter_next
func F_array_iter_next(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_array_set_element github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_array_set_element
func F_array_set_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_construct_md_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_md_array
func F_construct_md_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_array_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_ref
func F_array_ref(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_array_set github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_set
func F_array_set(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_construct_array_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_array_builtin
func F_construct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_array_create_iterator github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_create_iterator
func F_array_create_iterator(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_array_iterate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_iterate
func F_array_iterate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_array_free_iterator github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_free_iterator
func F_array_free_iterator(m *base.Module, l0 int32)
//go:linkname F_accumArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_accumArrayResult
func F_accumArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_accumArrayResultAny github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_accumArrayResultAny
func F_accumArrayResultAny(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeArrayResultAny github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeArrayResultAny
func F_makeArrayResultAny(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_array_fill_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_fill_internal
func F_array_fill_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ArrayGetNItems github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ArrayGetNItems
func F_ArrayGetNItems(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ArrayCheckBounds github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ArrayCheckBounds
func F_ArrayCheckBounds(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ArrayGetIntegerTypmods github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ArrayGetIntegerTypmods
func F_ArrayGetIntegerTypmods(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_bool_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_bool_with_len
func F_parse_bool_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_cash_div_float8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cash_div_float8
func F_cash_div_float8(m *base.Module, l0 int64, l1 float64) int64
//go:linkname F_cryptohash_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cryptohash_internal
func F_cryptohash_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_anytime_typmod_check github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_anytime_typmod_check
func F_anytime_typmod_check(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_j2date github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_j2date
func F_j2date(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_DetermineTimeZoneOffsetInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DetermineTimeZoneOffsetInternal
func F_DetermineTimeZoneOffsetInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DetermineTimeZoneOffset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DetermineTimeZoneOffset
func F_DetermineTimeZoneOffset(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DecodeTimezoneNameToTz github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecodeTimezoneNameToTz
func F_DecodeTimezoneNameToTz(m *base.Module, l0 int32) int32
//go:linkname F_AppendSeconds github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AppendSeconds
func F_AppendSeconds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_EncodeTimezone github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_EncodeTimezone
func F_EncodeTimezone(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datumCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_datumCopy
func F_datumCopy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datumTransfer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_datumTransfer
func F_datumTransfer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datumIsEqual github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_datumIsEqual
func F_datumIsEqual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_datumEstimateSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_datumEstimateSpace
func F_datumEstimateSpace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_datumSerialize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_datumSerialize
func F_datumSerialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_domain_check github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_domain_check
func F_domain_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_domain_check_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_domain_check_safe
func F_domain_check_safe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_enum_cmp_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_enum_cmp_internal
func F_enum_cmp_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EOH_get_flat_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EOH_get_flat_size
func F_EOH_get_flat_size(m *base.Module, l0 int32) int32
//go:linkname F_DeleteExpandedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DeleteExpandedObject
func F_DeleteExpandedObject(m *base.Module, l0 int32)
//go:linkname F_expanded_record_fetch_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expanded_record_fetch_tupdesc
func F_expanded_record_fetch_tupdesc(m *base.Module, l0 int32) int32
//go:linkname F_deconstruct_expanded_record github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_deconstruct_expanded_record
func F_deconstruct_expanded_record(m *base.Module, l0 int32)
//go:linkname F_expanded_record_set_field_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expanded_record_set_field_internal
func F_expanded_record_set_field_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_float_overflow_error github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float_overflow_error
func F_float_overflow_error(m *base.Module)
//go:linkname F_float8in_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_float8in_internal
func F_float8in_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_float8out_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float8out_internal
func F_float8out_internal(m *base.Module, l0 float64) int32
//go:linkname F_init_degree_constants github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_init_degree_constants
func F_init_degree_constants(m *base.Module)
//go:linkname F_format_type_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_format_type_extended
func F_format_type_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_format_type_be_qualified github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_format_type_be_qualified
func F_format_type_be_qualified(m *base.Module, l0 int32) int32
//go:linkname F_format_type_with_typemod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_format_type_with_typemod
func F_format_type_with_typemod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_str_tolower github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_str_tolower
func F_str_tolower(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_convert_and_check_filename github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_and_check_filename
func F_convert_and_check_filename(m *base.Module, l0 int32) int32
//go:linkname F_read_binary_file github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_binary_file
func F_read_binary_file(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32
//go:linkname F_pg_ls_tmpdir github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_ls_tmpdir
func F_pg_ls_tmpdir(m *base.Module, l0 int32, l1 int32)
//go:linkname F_path_decode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_path_decode
func F_path_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_pair_decode github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pair_decode
func F_pair_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_box_ar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_box_ar
func F_box_ar(m *base.Module, l0 int32) float64
//go:linkname F_box_cn github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_box_cn
func F_box_cn(m *base.Module, l0 int32, l1 int32)
//go:linkname F_point_dt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_point_dt
func F_point_dt(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_point_sl github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_point_sl
func F_point_sl(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_line_construct github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_line_construct
func F_line_construct(m *base.Module, l0 int32, l1 int32, l2 float64)
//go:linkname F_line_interpt_line github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_line_interpt_line
func F_line_interpt_line(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lseg_interpt_line github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_lseg_interpt_line
func F_lseg_interpt_line(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lseg_closept_lseg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lseg_closept_lseg
func F_lseg_closept_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_lseg_closept_point github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lseg_closept_point
func F_lseg_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_line_closept_point github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_line_closept_point
func F_line_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_box_closept_lseg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_box_closept_lseg
func F_box_closept_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_point_inside github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_point_inside
func F_point_inside(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_point_div_point github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_point_div_point
func F_point_div_point(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_inet_cidr_pton_ipv6 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_inet_cidr_pton_ipv6
func F_inet_cidr_pton_ipv6(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_int4recv github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_int4recv
func F_int4recv(m *base.Module, l0 int32) int32
//go:linkname F_JsonEncodeDateTime github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_JsonEncodeDateTime
func F_JsonEncodeDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_escape_json github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_escape_json
func F_escape_json(m *base.Module, l0 int32, l1 int32)
//go:linkname F_datum_to_json_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_datum_to_json_internal
func F_datum_to_json_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_escape_json_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_escape_json_with_len
func F_escape_json_with_len(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_add_json github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_json
func F_add_json(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_JsonbToCString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonbToCString
func F_JsonbToCString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_JsonbExtractScalar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonbExtractScalar
func F_JsonbExtractScalar(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_array_dim_to_jsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_dim_to_jsonb
func F_array_dim_to_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_add_jsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_jsonb
func F_add_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_cannotCastJsonbValue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cannotCastJsonbValue
func F_cannotCastJsonbValue(m *base.Module, l0 int32, l1 int32)
//go:linkname F_extract_jsp_query github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_extract_jsp_query
func F_extract_jsp_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_extract_jsp_path_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_extract_jsp_path_expr
func F_extract_jsp_path_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_execute_jsp_gin_node github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_execute_jsp_gin_node
func F_execute_jsp_gin_node(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonbValueToJsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_JsonbValueToJsonb
func F_JsonbValueToJsonb(m *base.Module, l0 int32) int32
//go:linkname F_pushJsonbValue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pushJsonbValue
func F_pushJsonbValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_JsonbIteratorNext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonbIteratorNext
func F_JsonbIteratorNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_compareJsonbContainers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_compareJsonbContainers
func F_compareJsonbContainers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonbIteratorInit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_JsonbIteratorInit
func F_JsonbIteratorInit(m *base.Module, l0 int32) int32
//go:linkname F_getKeyJsonValueFromContainer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getKeyJsonValueFromContainer
func F_getKeyJsonValueFromContainer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_JsonbHashScalarValue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_JsonbHashScalarValue
func F_JsonbHashScalarValue(m *base.Module, l0 int32, l1 int32)
//go:linkname F_convertJsonbScalar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_convertJsonbScalar
func F_convertJsonbScalar(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_parse_json_or_errsave github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_parse_json_or_errsave
func F_pg_parse_json_or_errsave(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_path_all github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_path_all
func F_get_path_all(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_jsonb_path_all github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_jsonb_path_all
func F_get_jsonb_path_all(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_setPath github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_setPath
func F_setPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_each_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_each_worker
func F_each_worker(m *base.Module, l0 int32, l1 int32)
//go:linkname F_each_worker_jsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_each_worker_jsonb
func F_each_worker_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_elements_worker_jsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_elements_worker_jsonb
func F_elements_worker_jsonb(m *base.Module, l0 int32, l1 int32)
//go:linkname F_populate_record_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_populate_record_worker
func F_populate_record_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_populate_record_field github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_populate_record_field
func F_populate_record_field(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_parse_jsonb_index_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_jsonb_index_flags
func F_parse_jsonb_index_flags(m *base.Module, l0 int32) int32
//go:linkname F_iterate_json_values github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_iterate_json_values
func F_iterate_json_values(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_json_categorize_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_json_categorize_type
func F_json_categorize_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_jspInitByBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_jspInitByBuffer
func F_jspInitByBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_printJsonPathItem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_printJsonPathItem
func F_printJsonPathItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_jspOperationName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jspOperationName
func F_jspOperationName(m *base.Module, l0 int32) int32
//go:linkname F_jspGetArg github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_jspGetArg
func F_jspGetArg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_jspGetNext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_jspGetNext
func F_jspGetNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_jspGetRightArg github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_jspGetRightArg
func F_jspGetRightArg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_jspIsMutableWalker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_jspIsMutableWalker
func F_jspIsMutableWalker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_jsonb_path_exists_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_jsonb_path_exists_internal
func F_jsonb_path_exists_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_executeJsonPath github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_executeJsonPath
func F_executeJsonPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_jsonb_path_query_array_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jsonb_path_query_array_internal
func F_jsonb_path_query_array_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_jsonb_path_query_first_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_jsonb_path_query_first_internal
func F_jsonb_path_query_first_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_executeItemOptUnwrapTarget github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_executeItemOptUnwrapTarget
func F_executeItemOptUnwrapTarget(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_executePredicate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_executePredicate
func F_executePredicate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_JsonTableResetNestedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_JsonTableResetNestedPlan
func F_JsonTableResetNestedPlan(m *base.Module, l0 int32)
//go:linkname F_GenericMatchText github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GenericMatchText
func F_GenericMatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_Generic_Text_IC_like github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_Generic_Text_IC_like
func F_Generic_Text_IC_like(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_like_regex_support github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_like_regex_support
func F_like_regex_support(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_patternsel_common github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_patternsel_common
func F_patternsel_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64
//go:linkname F_like_fixed_prefix github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_like_fixed_prefix
func F_like_fixed_prefix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_regex_fixed_prefix github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_regex_fixed_prefix
func F_regex_fixed_prefix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_get_multirange_io_data github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_multirange_io_data
func F_get_multirange_io_data(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_multirange
func F_make_multirange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_multirange_get_range github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_multirange_get_range
func F_multirange_get_range(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_multirange_get_typcache github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_multirange_get_typcache
func F_multirange_get_typcache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_multirange_get_bounds github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_multirange_get_bounds
func F_multirange_get_bounds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_multirange_eq_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_multirange_eq_internal
func F_multirange_eq_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_multirange_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_multirange_cmp
func F_multirange_cmp(m *base.Module, l0 int32) int32
//go:linkname F_calc_hist_selectivity_scalar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_calc_hist_selectivity_scalar
func F_calc_hist_selectivity_scalar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_calc_hist_selectivity_contained github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_calc_hist_selectivity_contained
func F_calc_hist_selectivity_contained(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) float64
//go:linkname F_bitncmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bitncmp
func F_bitncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_internal_inetpl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_internal_inetpl
func F_internal_inetpl(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_make_result_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_result_opt_error
func F_make_result_opt_error(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_var_from_str github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_set_var_from_str
func F_set_var_from_str(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_mul_var github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_mul_var
func F_mul_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_add_var github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_var
func F_add_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_div_var github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_div_var
func F_div_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_set_var_from_num github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_var_from_num
func F_set_var_from_num(m *base.Module, l0 int32, l1 int32)
//go:linkname F_cmp_abs_common github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cmp_abs_common
func F_cmp_abs_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_add_abs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_add_abs
func F_add_abs(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_sub_abs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sub_abs
func F_sub_abs(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_numericvar_to_int64 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_numericvar_to_int64
func F_numericvar_to_int64(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_numeric_sub_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_numeric_sub_opt_error
func F_numeric_sub_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ln_var github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ln_var
func F_ln_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_accum_sum_add github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_accum_sum_add
func F_accum_sum_add(m *base.Module, l0 int32, l1 int32)
//go:linkname F_numericvar_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_numericvar_deserialize
func F_numericvar_deserialize(m *base.Module, l0 int32, l1 int32)
//go:linkname F_numericvar_to_int128 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_numericvar_to_int128
func F_numericvar_to_int128(m *base.Module, l0 int32, l1 int32)
//go:linkname F_numeric_stddev_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_numeric_stddev_internal
func F_numeric_stddev_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_strtoint32_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_strtoint32_safe
func F_pg_strtoint32_safe(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_strtoint64_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_strtoint64_safe
func F_pg_strtoint64_safe(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_uint32in_subr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_uint32in_subr
func F_uint32in_subr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_ultostr_zeropad github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_ultostr_zeropad
func F_pg_ultostr_zeropad(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_ultostr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_ultostr
func F_pg_ultostr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dotrim github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dotrim
func F_dotrim(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_percentile_cont_final_common github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_percentile_cont_final_common
func F_percentile_cont_final_common(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hypothetical_rank_common github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hypothetical_rank_common
func F_hypothetical_rank_common(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_check_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_locale
func F_check_locale(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_newlocale_from_collation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_newlocale_from_collation
func F_pg_newlocale_from_collation(m *base.Module, l0 int32) int32
//go:linkname F_get_collation_actual_version github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_collation_actual_version
func F_get_collation_actual_version(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_strtitle github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_strtitle
func F_pg_strtitle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_strncoll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_strncoll
func F_pg_strncoll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_strnxfrm github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_strnxfrm
func F_pg_strnxfrm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_strlower_libc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strlower_libc
func F_strlower_libc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_quote_literal_cstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_quote_literal_cstr
func F_quote_literal_cstr(m *base.Module, l0 int32) int32
//go:linkname F_make_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_range
func F_make_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_range_out github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_out
func F_range_out(m *base.Module, l0 int32) int32
//go:linkname F_range_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_deserialize
func F_range_deserialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_range_get_typcache github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_get_typcache
func F_range_get_typcache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_range_contains_elem_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_range_contains_elem_internal
func F_range_contains_elem_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_eq_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_eq_internal
func F_range_eq_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_after_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_after_internal
func F_range_after_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_adjacent_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_adjacent_internal
func F_range_adjacent_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overlaps_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_overlaps_internal
func F_range_overlaps_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_union_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_union_internal
func F_range_union_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_range_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_cmp
func F_range_cmp(m *base.Module, l0 int32) int32
//go:linkname F_datum_write github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_datum_write
func F_datum_write(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_range_gist_consistent_leaf_range github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_gist_consistent_leaf_range
func F_range_gist_consistent_leaf_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_range_gist_consistent_leaf_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_gist_consistent_leaf_multirange
func F_range_gist_consistent_leaf_multirange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_range_gist_consistent_int_range github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_gist_consistent_int_range
func F_range_gist_consistent_int_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_range_gist_consistent_int_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_gist_consistent_int_multirange
func F_range_gist_consistent_int_multirange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_adjacent_inner_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_adjacent_inner_consistent
func F_adjacent_inner_consistent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_RE_wchar_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RE_wchar_execute
func F_RE_wchar_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_parse_re_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parse_re_flags
func F_parse_re_flags(m *base.Module, l0 int32, l1 int32)
//go:linkname F_regexp_count github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_regexp_count
func F_regexp_count(m *base.Module, l0 int32) int32
//go:linkname F_setup_regexp_matches github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_setup_regexp_matches
func F_setup_regexp_matches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_build_regexp_split_result github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_regexp_split_result
func F_build_regexp_split_result(m *base.Module, l0 int32) int32
//go:linkname F_stringToQualifiedNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_stringToQualifiedNameList
func F_stringToQualifiedNameList(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_format_procedure github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_format_procedure
func F_format_procedure(m *base.Module, l0 int32) int32
//go:linkname F_format_operator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_format_operator
func F_format_operator(m *base.Module, l0 int32) int32
//go:linkname F_ri_FetchConstraintInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ri_FetchConstraintInfo
func F_ri_FetchConstraintInfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ri_FetchPreparedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ri_FetchPreparedPlan
func F_ri_FetchPreparedPlan(m *base.Module, l0 int32) int32
//go:linkname F_ri_PlanCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ri_PlanCheck
func F_ri_PlanCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ri_PerformCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ri_PerformCheck
func F_ri_PerformCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_record_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_record_cmp
func F_record_cmp(m *base.Module, l0 int32) int32
//go:linkname F_pg_get_ruledef_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_get_ruledef_worker
func F_pg_get_ruledef_worker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rule_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_rule_expr
func F_get_rule_expr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_query_def github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_query_def
func F_get_query_def(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_pg_get_triggerdef_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_get_triggerdef_worker
func F_pg_get_triggerdef_worker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fastgetattr_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fastgetattr_3
func F_fastgetattr_3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_generate_collation_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_generate_collation_name
func F_generate_collation_name(m *base.Module, l0 int32) int32
//go:linkname F_get_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_variable
func F_get_variable(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_get_partkeydef_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_get_partkeydef_worker
func F_pg_get_partkeydef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_get_expr_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_get_expr_worker
func F_pg_get_expr_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_quote_qualified_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_quote_qualified_identifier
func F_quote_qualified_identifier(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generate_operator_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generate_operator_clause
func F_generate_operator_clause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_agg_expr_helper github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_agg_expr_helper
func F_get_agg_expr_helper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_json_expr_options github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_json_expr_options
func F_get_json_expr_options(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_json_table_nested_columns github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_json_table_nested_columns
func F_get_json_table_nested_columns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_get_restriction_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_restriction_variable
func F_get_restriction_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_statistic_proc_security_check github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_statistic_proc_security_check
func F_statistic_proc_security_check(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generic_restriction_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generic_restriction_selectivity
func F_generic_restriction_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) float64
//go:linkname F_scalarineqsel_wrapper github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scalarineqsel_wrapper
func F_scalarineqsel_wrapper(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_join_variables github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_join_variables
func F_get_join_variables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_eqjoinsel_semi github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_eqjoinsel_semi
func F_eqjoinsel_semi(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 float64, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) float64
//go:linkname F_estimate_num_groups github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_estimate_num_groups
func F_estimate_num_groups(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32) float64
//go:linkname F_index_other_operands_eval_cost github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_other_operands_eval_cost
func F_index_other_operands_eval_cost(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_genericcostestimate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_genericcostestimate
func F_genericcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32)
//go:linkname F_anytimestamp_typmod_check github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_anytimestamp_typmod_check
func F_anytimestamp_typmod_check(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_timestamp2tm github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_timestamp2tm
func F_timestamp2tm(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_timestamp2timestamptz_opt_overflow github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamp2timestamptz_opt_overflow
func F_timestamp2timestamptz_opt_overflow(m *base.Module, l0 int64, l1 int32) int64
//go:linkname F_GetCurrentTimestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetCurrentTimestamp
func F_GetCurrentTimestamp(m *base.Module) int64
//go:linkname F_timestamptz_to_time_t github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_timestamptz_to_time_t
func F_timestamptz_to_time_t(m *base.Module, l0 int64) int64
//go:linkname F_timestamp_cmp_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamp_cmp_internal
func F_timestamp_cmp_internal(m *base.Module, l0 int64, l1 int64) int32
//go:linkname F_timestamp_cmp_timestamptz_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_timestamp_cmp_timestamptz_internal
func F_timestamp_cmp_timestamptz_internal(m *base.Module, l0 int64, l1 int64) int32
//go:linkname F_timestamptz_pl_interval_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamptz_pl_interval_internal
func F_timestamptz_pl_interval_internal(m *base.Module, l0 int64, l1 int32, l2 int32) int64
//go:linkname F_finite_interval_pl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_finite_interval_pl
func F_finite_interval_pl(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_finite_interval_mi github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_finite_interval_mi
func F_finite_interval_mi(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_timestamp_part_common github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_timestamp_part_common
func F_timestamp_part_common(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_timestamptz_part_common github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_timestamptz_part_common
func F_timestamptz_part_common(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_tsquery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_parse_tsquery
func F_parse_tsquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_QTNFree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_QTNFree
func F_QTNFree(m *base.Module, l0 int32)
//go:linkname F_QTNodeCompare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_QTNodeCompare
func F_QTNodeCompare(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_QTNSort github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_QTNSort
func F_QTNSort(m *base.Module, l0 int32)
//go:linkname F_QTNEq github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_QTNEq
func F_QTNEq(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fillQT github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fillQT
func F_fillQT(m *base.Module, l0 int32, l1 int32)
//go:linkname F_QTNCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_QTNCopy
func F_QTNCopy(m *base.Module, l0 int32) int32
//go:linkname F_tsvector_delete_by_indices github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tsvector_delete_by_indices
func F_tsvector_delete_by_indices(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tsCompareString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tsCompareString
func F_tsCompareString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_TS_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TS_execute
func F_TS_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_insertStatEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_insertStatEntry
func F_insertStatEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_gettoken_tsvector github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gettoken_tsvector
func F_gettoken_tsvector(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_bitsubstring github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bitsubstring
func F_bitsubstring(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_cstring_to_text_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cstring_to_text_with_len
func F_cstring_to_text_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_text_to_cstring github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_text_to_cstring
func F_text_to_cstring(m *base.Module, l0 int32) int32
//go:linkname F_text_to_cstring_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_text_to_cstring_buffer
func F_text_to_cstring_buffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_byteaout github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_byteaout
func F_byteaout(m *base.Module, l0 int32) int32
//go:linkname F_textout github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_textout
func F_textout(m *base.Module, l0 int32) int32
//go:linkname F_text_substring github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text_substring
func F_text_substring(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_varstr_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_varstr_cmp
func F_varstr_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_varstr_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_varstr_sortsupport
func F_varstr_sortsupport(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_varstrfastcmp_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_varstrfastcmp_locale
func F_varstrfastcmp_locale(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_textToQualifiedNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_textToQualifiedNameList
func F_textToQualifiedNameList(m *base.Module, l0 int32) int32
//go:linkname F_SplitIdentifierString github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SplitIdentifierString
func F_SplitIdentifierString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_split_text github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_split_text
func F_split_text(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_concat_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_concat_internal
func F_concat_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_varstr_levenshtein_less_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_varstr_levenshtein_less_equal
func F_varstr_levenshtein_less_equal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_map_sql_identifier_to_xml_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_map_sql_identifier_to_xml_name
func F_map_sql_identifier_to_xml_name(m *base.Module) int32
//go:linkname F_escape_xml github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_escape_xml
func F_escape_xml(m *base.Module, l0 int32) int32
//go:linkname F_map_sql_typecoll_to_xmlschema_types github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_map_sql_typecoll_to_xmlschema_types
func F_map_sql_typecoll_to_xmlschema_types(m *base.Module, l0 int32) int32
//go:linkname F_query_to_oid_list github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_query_to_oid_list
func F_query_to_oid_list(m *base.Module, l0 int32) int32
//go:linkname F_CatCacheRemoveCTup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CatCacheRemoveCTup
func F_CatCacheRemoveCTup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_InitCatCachePhase2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_InitCatCachePhase2
func F_InitCatCachePhase2(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CatalogCacheInitializeCache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CatalogCacheInitializeCache
func F_CatalogCacheInitializeCache(m *base.Module, l0 int32)
//go:linkname F_SearchCatCache github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SearchCatCache
func F_SearchCatCache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SearchCatCache1 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SearchCatCache1
func F_SearchCatCache1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReleaseCatCache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReleaseCatCache
func F_ReleaseCatCache(m *base.Module, l0 int32)
//go:linkname F_CatalogCacheComputeTupleHashValue github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CatalogCacheComputeTupleHashValue
func F_CatalogCacheComputeTupleHashValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReleaseCatCacheList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReleaseCatCacheList
func F_ReleaseCatCacheList(m *base.Module, l0 int32)
//go:linkname F_cfunc_resolve_polymorphic_argtypes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cfunc_resolve_polymorphic_argtypes
func F_cfunc_resolve_polymorphic_argtypes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_InvalidateSystemCachesExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InvalidateSystemCachesExtended
func F_InvalidateSystemCachesExtended(m *base.Module)
//go:linkname F_LocalExecuteInvalidationMessage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LocalExecuteInvalidationMessage
func F_LocalExecuteInvalidationMessage(m *base.Module, l0 int32)
//go:linkname F_LogLogicalInvalidations github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LogLogicalInvalidations
func F_LogLogicalInvalidations(m *base.Module)
//go:linkname F_CacheInvalidateHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CacheInvalidateHeapTuple
func F_CacheInvalidateHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RegisterCatcacheInvalidation github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RegisterCatcacheInvalidation
func F_RegisterCatcacheInvalidation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RegisterRelcacheInvalidation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RegisterRelcacheInvalidation
func F_RegisterRelcacheInvalidation(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CacheInvalidateRelcacheByTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CacheInvalidateRelcacheByTuple
func F_CacheInvalidateRelcacheByTuple(m *base.Module, l0 int32)
//go:linkname F_CacheRegisterSyscacheCallback github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CacheRegisterSyscacheCallback
func F_CacheRegisterSyscacheCallback(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CacheRegisterRelcacheCallback github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CacheRegisterRelcacheCallback
func F_CacheRegisterRelcacheCallback(m *base.Module, l0 int32)
//go:linkname F_get_op_opfamily_strategy github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_op_opfamily_strategy
func F_get_op_opfamily_strategy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_op_opfamily_sortfamily github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_op_opfamily_sortfamily
func F_get_op_opfamily_sortfamily(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_opfamily_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_opfamily_member
func F_get_opfamily_member(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_op_hash_functions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_op_hash_functions
func F_get_op_hash_functions(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_opfamily_proc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opfamily_proc
func F_get_opfamily_proc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_op_index_interpretation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_op_index_interpretation
func F_get_op_index_interpretation(m *base.Module, l0 int32) int32
//go:linkname F_get_negator github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_negator
func F_get_negator(m *base.Module, l0 int32) int32
//go:linkname F_equality_ops_are_compatible github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_equality_ops_are_compatible
func F_equality_ops_are_compatible(m *base.Module, l0 int32, l1 int32) int32
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
//go:linkname F_get_opclass_family github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opclass_family
func F_get_opclass_family(m *base.Module, l0 int32) int32
//go:linkname F_get_opclass_input_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_opclass_input_type
func F_get_opclass_input_type(m *base.Module, l0 int32) int32
//go:linkname F_get_opfamily_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_opfamily_name
func F_get_opfamily_name(m *base.Module, l0 int32) int32
//go:linkname F_get_opcode github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opcode
func F_get_opcode(m *base.Module, l0 int32) int32
//go:linkname F_op_input_types github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_op_input_types
func F_op_input_types(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_func_strict github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_func_strict
func F_func_strict(m *base.Module, l0 int32) int32
//go:linkname F_op_volatile github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_op_volatile
func F_op_volatile(m *base.Module, l0 int32) int32
//go:linkname F_func_volatile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_func_volatile
func F_func_volatile(m *base.Module, l0 int32) int32
//go:linkname F_get_commutator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_commutator
func F_get_commutator(m *base.Module, l0 int32) int32
//go:linkname F_get_oprrest github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_oprrest
func F_get_oprrest(m *base.Module, l0 int32) int32
//go:linkname F_get_func_name github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_name
func F_get_func_name(m *base.Module, l0 int32) int32
//go:linkname F_get_func_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_rettype
func F_get_func_rettype(m *base.Module, l0 int32) int32
//go:linkname F_get_func_leakproof github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_leakproof
func F_get_func_leakproof(m *base.Module, l0 int32) int32
//go:linkname F_get_func_support github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_func_support
func F_get_func_support(m *base.Module, l0 int32) int32
//go:linkname F_get_relname_relid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_relname_relid
func F_get_relname_relid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rel_namespace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_rel_namespace
func F_get_rel_namespace(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_relkind github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_rel_relkind
func F_get_rel_relkind(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_persistence github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_rel_persistence
func F_get_rel_persistence(m *base.Module, l0 int32) int32
//go:linkname F_get_typlen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_typlen
func F_get_typlen(m *base.Module, l0 int32) int32
//go:linkname F_get_typlenbyval github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_typlenbyval
func F_get_typlenbyval(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_typlenbyvalalign github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typlenbyvalalign
func F_get_typlenbyvalalign(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_type_io_data github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_type_io_data
func F_get_type_io_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_get_typstorage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_typstorage
func F_get_typstorage(m *base.Module, l0 int32) int32
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
//go:linkname F_type_is_enum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_type_is_enum
func F_type_is_enum(m *base.Module, l0 int32) int32
//go:linkname F_type_is_range github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_type_is_range
func F_type_is_range(m *base.Module, l0 int32) int32
//go:linkname F_type_is_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_type_is_multirange
func F_type_is_multirange(m *base.Module, l0 int32) int32
//go:linkname F_get_type_category_preferred github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_type_category_preferred
func F_get_type_category_preferred(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_element_type github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_element_type
func F_get_element_type(m *base.Module, l0 int32) int32
//go:linkname F_get_array_type github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_array_type
func F_get_array_type(m *base.Module, l0 int32) int32
//go:linkname F_get_base_element_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_base_element_type
func F_get_base_element_type(m *base.Module, l0 int32) int32
//go:linkname F_getTypeInputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getTypeInputInfo
func F_getTypeInputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_getTypeOutputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getTypeOutputInfo
func F_getTypeOutputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_getTypeBinaryInputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getTypeBinaryInputInfo
func F_getTypeBinaryInputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_typcollation github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_typcollation
func F_get_typcollation(m *base.Module, l0 int32) int32
//go:linkname F_type_is_collatable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_type_is_collatable
func F_type_is_collatable(m *base.Module, l0 int32) int32
//go:linkname F_get_attavgwidth github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_attavgwidth
func F_get_attavgwidth(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_attstatsslot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_attstatsslot
func F_get_attstatsslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_free_attstatsslot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_free_attstatsslot
func F_free_attstatsslot(m *base.Module, l0 int32)
//go:linkname F_get_namespace_name_or_temp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_namespace_name_or_temp
func F_get_namespace_name_or_temp(m *base.Module, l0 int32) int32
//go:linkname F_get_range_subtype github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_range_subtype
func F_get_range_subtype(m *base.Module, l0 int32) int32
//go:linkname F_get_range_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_range_multirange
func F_get_range_multirange(m *base.Module, l0 int32) int32
//go:linkname F_get_multirange_range github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_multirange_range
func F_get_multirange_range(m *base.Module, l0 int32) int32
//go:linkname F_get_index_isclustered github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_index_isclustered
func F_get_index_isclustered(m *base.Module, l0 int32) int32
//go:linkname F_get_publication_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_publication_oid
func F_get_publication_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationGetPartitionKey github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetPartitionKey
func F_RelationGetPartitionKey(m *base.Module, l0 int32) int32
//go:linkname F_CachedPlanGetTargetList github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CachedPlanGetTargetList
func F_CachedPlanGetTargetList(m *base.Module, l0 int32) int32
//go:linkname F_ReleaseAllPlanCacheRefsInOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReleaseAllPlanCacheRefsInOwner
func F_ReleaseAllPlanCacheRefsInOwner(m *base.Module, l0 int32)
//go:linkname F_RelationInitTableAccessMethod github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationInitTableAccessMethod
func F_RelationInitTableAccessMethod(m *base.Module, l0 int32)
//go:linkname F_RelationIdGetRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationIdGetRelation
func F_RelationIdGetRelation(m *base.Module, l0 int32) int32
//go:linkname F_RelationParseRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationParseRelOptions
func F_RelationParseRelOptions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationBuildRuleLock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationBuildRuleLock
func F_RelationBuildRuleLock(m *base.Module, l0 int32)
//go:linkname F_RelationClose github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RelationClose
func F_RelationClose(m *base.Module, l0 int32)
//go:linkname F_load_relcache_init_file github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_load_relcache_init_file
func F_load_relcache_init_file(m *base.Module, l0 int32) int32
//go:linkname F_formrdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_formrdesc
func F_formrdesc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_load_critical_index github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_load_critical_index
func F_load_critical_index(m *base.Module, l0 int32, l1 int32)
//go:linkname F_write_relcache_init_file github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_write_relcache_init_file
func F_write_relcache_init_file(m *base.Module, l0 int32)
//go:linkname F_RelationGetFKeyList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationGetFKeyList
func F_RelationGetFKeyList(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexList github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationGetIndexList
func F_RelationGetIndexList(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetPrimaryKeyIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetPrimaryKeyIndex
func F_RelationGetPrimaryKeyIndex(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationGetIndexExpressions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetIndexExpressions
func F_RelationGetIndexExpressions(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexPredicate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetIndexPredicate
func F_RelationGetIndexPredicate(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexAttrBitmap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RelationGetIndexAttrBitmap
func F_RelationGetIndexAttrBitmap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_errtable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_errtable
func F_errtable(m *base.Module, l0 int32)
//go:linkname F_RelidByRelfilenumber github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelidByRelfilenumber
func F_RelidByRelfilenumber(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_read_relmap_file github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_read_relmap_file
func F_read_relmap_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_write_relmap_file github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_write_relmap_file
func F_write_relmap_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_RelationMapUpdateMap github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationMapUpdateMap
func F_RelationMapUpdateMap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_merge_map_updates github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_merge_map_updates
func F_merge_map_updates(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_tablespace_page_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_tablespace_page_costs
func F_get_tablespace_page_costs(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_SearchSysCache2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCache2
func F_SearchSysCache2(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchSysCache3 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SearchSysCache3
func F_SearchSysCache3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SearchSysCache4 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCache4
func F_SearchSysCache4(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SearchSysCacheLockedCopy1 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCacheLockedCopy1
func F_SearchSysCacheLockedCopy1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheExists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCacheExists
func F_SearchSysCacheExists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_GetSysCacheOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetSysCacheOid
func F_GetSysCacheOid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SearchSysCacheAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SearchSysCacheAttName
func F_SearchSysCacheAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheCopyAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SearchSysCacheCopyAttName
func F_SearchSysCacheCopyAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheCopyAttNum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SearchSysCacheCopyAttNum
func F_SearchSysCacheCopyAttNum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SysCacheGetAttrNotNull github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SysCacheGetAttrNotNull
func F_SysCacheGetAttrNotNull(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchSysCacheList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SearchSysCacheList
func F_SearchSysCacheList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_lookup_ts_parser_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lookup_ts_parser_cache
func F_lookup_ts_parser_cache(m *base.Module, l0 int32) int32
//go:linkname F_lookup_ts_dictionary_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lookup_ts_dictionary_cache
func F_lookup_ts_dictionary_cache(m *base.Module, l0 int32) int32
//go:linkname F_getTSCurrentConfig github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_getTSCurrentConfig
func F_getTSCurrentConfig(m *base.Module) int32
//go:linkname F_cache_record_field_properties github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cache_record_field_properties
func F_cache_record_field_properties(m *base.Module, l0 int32)
//go:linkname F_load_rangetype_info github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_load_rangetype_info
func F_load_rangetype_info(m *base.Module, l0 int32)
//go:linkname F_load_typcache_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_load_typcache_tupdesc
func F_load_typcache_tupdesc(m *base.Module, l0 int32)
//go:linkname F_load_domaintype_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_load_domaintype_info
func F_load_domaintype_info(m *base.Module, l0 int32)
//go:linkname F_InitDomainConstraintRef github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InitDomainConstraintRef
func F_InitDomainConstraintRef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_lookup_rowtype_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lookup_rowtype_tupdesc
func F_lookup_rowtype_tupdesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookup_rowtype_tupdesc_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_lookup_rowtype_tupdesc_copy
func F_lookup_rowtype_tupdesc_copy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assign_record_type_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_assign_record_type_typmod
func F_assign_record_type_typmod(m *base.Module, l0 int32)
//go:linkname F_assign_record_type_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_assign_record_type_identifier
func F_assign_record_type_identifier(m *base.Module, l0 int32, l1 int32) int64
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
//go:linkname F_FreeErrorDataContents github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreeErrorDataContents
func F_FreeErrorDataContents(m *base.Module, l0 int32)
//go:linkname F_pg_re_throw github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_re_throw
func F_pg_re_throw(m *base.Module)
//go:linkname F_errsave_start github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errsave_start
func F_errsave_start(m *base.Module, l0 int32) int32
//go:linkname F_errsave_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errsave_finish
func F_errsave_finish(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errcode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errcode
func F_errcode(m *base.Module, l0 int32)
//go:linkname F_errcode_for_socket_access github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errcode_for_socket_access
func F_errcode_for_socket_access(m *base.Module)
//go:linkname F_errmsg_plural github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errmsg_plural
func F_errmsg_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errdetail github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errdetail
func F_errdetail(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errhint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errhint
func F_errhint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errhint_plural github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errhint_plural
func F_errhint_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errcontext_msg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errcontext_msg
func F_errcontext_msg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errhidestmt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errhidestmt
func F_errhidestmt(m *base.Module)
//go:linkname F_errhidecontext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errhidecontext
func F_errhidecontext(m *base.Module)
//go:linkname F_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errposition
func F_errposition(m *base.Module, l0 int32) int32
//go:linkname F_internalerrposition github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_internalerrposition
func F_internalerrposition(m *base.Module, l0 int32)
//go:linkname F_internalerrquery github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_internalerrquery
func F_internalerrquery(m *base.Module, l0 int32) int32
//go:linkname F_geterrposition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_geterrposition
func F_geterrposition(m *base.Module) int32
//go:linkname F_format_elog_string github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_format_elog_string
func F_format_elog_string(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CopyErrorData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopyErrorData
func F_CopyErrorData(m *base.Module) int32
//go:linkname F_FlushErrorState github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FlushErrorState
func F_FlushErrorState(m *base.Module)
//go:linkname F_ReThrowError github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReThrowError
func F_ReThrowError(m *base.Module, l0 int32)
//go:linkname F_load_external_function github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_load_external_function
func F_load_external_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
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
//go:linkname F_OidFunctionCall0Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_OidFunctionCall0Coll
func F_OidFunctionCall0Coll(m *base.Module, l0 int32) int32
//go:linkname F_OidFunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_OidFunctionCall1Coll
func F_OidFunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_OidFunctionCall2Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_OidFunctionCall2Coll
func F_OidFunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_InputFunctionCallSafe github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InputFunctionCallSafe
func F_InputFunctionCallSafe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_OutputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OutputFunctionCall
func F_OutputFunctionCall(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReceiveFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReceiveFunctionCall
func F_ReceiveFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_OidInputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_OidInputFunctionCall
func F_OidInputFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_OidOutputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OidOutputFunctionCall
func F_OidOutputFunctionCall(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_Float8GetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_Float8GetDatum
func F_Float8GetDatum(m *base.Module, l0 float64) int32
//go:linkname F_pg_detoast_datum_packed github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_detoast_datum_packed
func F_pg_detoast_datum_packed(m *base.Module, l0 int32) int32
//go:linkname F_get_fn_expr_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_fn_expr_rettype
func F_get_fn_expr_rettype(m *base.Module, l0 int32) int32
//go:linkname F_internal_get_result_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_internal_get_result_type
func F_internal_get_result_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_init_MultiFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_init_MultiFuncCall
func F_init_MultiFuncCall(m *base.Module, l0 int32) int32
//go:linkname F_end_MultiFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_end_MultiFuncCall
func F_end_MultiFuncCall(m *base.Module, l0 int32)
//go:linkname F_get_expr_result_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_expr_result_type
func F_get_expr_result_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_func_arg_info github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_func_arg_info
func F_get_func_arg_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_create github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_create
func F_hash_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
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
//go:linkname F_hash_seq_init_with_hash_value github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_seq_init_with_hash_value
func F_hash_seq_init_with_hash_value(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_hash_seq_search github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_seq_search
func F_hash_seq_search(m *base.Module, l0 int32) int32
//go:linkname F_hash_seq_term github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_seq_term
func F_hash_seq_term(m *base.Module, l0 int32)
//go:linkname F_AtEOXact_HashTables github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AtEOXact_HashTables
func F_AtEOXact_HashTables(m *base.Module, l0 int32)
//go:linkname F_GetBackendTypeDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetBackendTypeDesc
func F_GetBackendTypeDesc(m *base.Module, l0 int32) int32
//go:linkname F_SetSessionAuthorization github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SetSessionAuthorization
func F_SetSessionAuthorization(m *base.Module, l0 int32, l1 int32)
//go:linkname F_InitializeSystemUser github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InitializeSystemUser
func F_InitializeSystemUser(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetUserNameFromId github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetUserNameFromId
func F_GetUserNameFromId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BaseInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BaseInit
func F_BaseInit(m *base.Module)
//go:linkname F_InitPostgres github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InitPostgres
func F_InitPostgres(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_local2local github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_local2local
func F_local2local(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_latin2mic github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_latin2mic
func F_latin2mic(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_mic2latin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_mic2latin
func F_mic2latin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_latin2mic_with_table github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_latin2mic_with_table
func F_latin2mic_with_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_UtfToLocal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UtfToLocal
func F_UtfToLocal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_LocalToUtf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LocalToUtf
func F_LocalToUtf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_SetClientEncoding github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SetClientEncoding
func F_SetClientEncoding(m *base.Module, l0 int32) int32
//go:linkname F_report_invalid_encoding github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_invalid_encoding
func F_report_invalid_encoding(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_any_to_server github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_any_to_server
func F_pg_any_to_server(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_server_to_any github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_server_to_any
func F_pg_server_to_any(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_unicode_to_server github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_unicode_to_server
func F_pg_unicode_to_server(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_mb2wchar_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_mb2wchar_with_len
func F_pg_mb2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_mblen_cstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_mblen_cstr
func F_pg_mblen_cstr(m *base.Module, l0 int32) int32
//go:linkname F_report_invalid_encoding_db github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_report_invalid_encoding_db
func F_report_invalid_encoding_db(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_mblen_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mblen_range
func F_pg_mblen_range(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mblen_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_mblen_with_len
func F_pg_mblen_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mbstrlen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_mbstrlen
func F_pg_mbstrlen(m *base.Module, l0 int32) int32
//go:linkname F_pg_mbstrlen_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mbstrlen_with_len
func F_pg_mbstrlen_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mbcliplen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mbcliplen
func F_pg_mbcliplen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_verifymbstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_verifymbstr
func F_pg_verifymbstr(m *base.Module, l0 int32, l1 int32)
//go:linkname F_check_encoding_conversion_args github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_encoding_conversion_args
func F_check_encoding_conversion_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_appendStringInfoStringQuoted github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_appendStringInfoStringQuoted
func F_appendStringInfoStringQuoted(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_find_option github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_option
func F_find_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_set_config_with_handle github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_set_config_with_handle
func F_set_config_with_handle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_guc_strdup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_guc_strdup
func F_guc_strdup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_config_option github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_config_option
func F_set_config_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_SetConfigOption github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SetConfigOption
func F_SetConfigOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_convert_GUC_name_for_parameter_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_convert_GUC_name_for_parameter_acl
func F_convert_GUC_name_for_parameter_acl(m *base.Module, l0 int32) int32
//go:linkname F_call_bool_check_hook github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_call_bool_check_hook
func F_call_bool_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_call_int_check_hook github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_call_int_check_hook
func F_call_int_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_call_real_check_hook github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_call_real_check_hook
func F_call_real_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_call_string_check_hook github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_call_string_check_hook
func F_call_string_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_call_enum_check_hook github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_call_enum_check_hook
func F_call_enum_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_discard_stack_value github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_discard_stack_value
func F_discard_stack_value(m *base.Module, l0 int32, l1 int32)
//go:linkname F_set_stack_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_stack_value
func F_set_stack_value(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RestrictSearchPath github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RestrictSearchPath
func F_RestrictSearchPath(m *base.Module)
//go:linkname F_AtEOXact_GUC github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AtEOXact_GUC
func F_AtEOXact_GUC(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ShowGUCOption github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ShowGUCOption
func F_ShowGUCOption(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_config_option_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_set_config_option_ext
func F_set_config_option_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_reapply_stacked_values github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_reapply_stacked_values
func F_reapply_stacked_values(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_DefineCustomRealVariable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DefineCustomRealVariable
func F_DefineCustomRealVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64)
//go:linkname F_DefineCustomEnumVariable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_DefineCustomEnumVariable
func F_DefineCustomEnumVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_MarkGUCPrefixReserved github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MarkGUCPrefixReserved
func F_MarkGUCPrefixReserved(m *base.Module, l0 int32)
//go:linkname F_read_gucstate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_gucstate
func F_read_gucstate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ParseLongOption github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ParseLongOption
func F_ParseLongOption(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GUCArrayAdd github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GUCArrayAdd
func F_GUCArrayAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_validate_option_array_item github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_validate_option_array_item
func F_validate_option_array_item(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GUCArrayDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GUCArrayDelete
func F_GUCArrayDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ProcessConfigFile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ProcessConfigFile
func F_ProcessConfigFile(m *base.Module, l0 int32)
//go:linkname F_ExtractSetVariableArgs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExtractSetVariableArgs
func F_ExtractSetVariableArgs(m *base.Module, l0 int32) int32
//go:linkname F_pg_rusage_show github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_rusage_show
func F_pg_rusage_show(m *base.Module, l0 int32) int32
//go:linkname F_check_stack_depth github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_stack_depth
func F_check_stack_depth(m *base.Module)
//go:linkname F_superuser_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_superuser_arg
func F_superuser_arg(m *base.Module, l0 int32) int32
//go:linkname F_enable_timeout github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_enable_timeout
func F_enable_timeout(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32)
//go:linkname F_schedule_alarm github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_schedule_alarm
func F_schedule_alarm(m *base.Module, l0 int64)
//go:linkname F_disable_timeout github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_disable_timeout
func F_disable_timeout(m *base.Module, l0 int32)
//go:linkname F_ParseTzFile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ParseTzFile
func F_ParseTzFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_AllocSetContextCreateInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AllocSetContextCreateInternal
func F_AllocSetContextCreateInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_AllocSetAllocLarge github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AllocSetAllocLarge
func F_AllocSetAllocLarge(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_AllocSetAllocFromNewBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AllocSetAllocFromNewBlock
func F_AllocSetAllocFromNewBlock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_dsa_create_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dsa_create_ext
func F_dsa_create_ext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dsa_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_dsa_attach
func F_dsa_attach(m *base.Module, l0 int32) int32
//go:linkname F_dsa_attach_in_place github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_dsa_attach_in_place
func F_dsa_attach_in_place(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsa_pin_mapping github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dsa_pin_mapping
func F_dsa_pin_mapping(m *base.Module, l0 int32)
//go:linkname F_alloc_object github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_alloc_object
func F_alloc_object(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_best_segment github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_best_segment
func F_get_best_segment(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_new_segment github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_new_segment
func F_make_new_segment(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsa_free github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dsa_free
func F_dsa_free(m *base.Module, l0 int32, l1 int32)
//go:linkname F_init_span github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_init_span
func F_init_span(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_check_for_freed_segments_locked github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_check_for_freed_segments_locked
func F_check_for_freed_segments_locked(m *base.Module, l0 int32)
//go:linkname F_get_segment_by_index github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_segment_by_index
func F_get_segment_by_index(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsa_pin github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_dsa_pin
func F_dsa_pin(m *base.Module, l0 int32)
//go:linkname F_dsa_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_dsa_detach
func F_dsa_detach(m *base.Module, l0 int32)
//go:linkname F_FreePageManagerGet github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreePageManagerGet
func F_FreePageManagerGet(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FreePageManagerPut github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreePageManagerPut
func F_FreePageManagerPut(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GenerationAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GenerationAlloc
func F_GenerationAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GenerationFree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GenerationFree
func F_GenerationFree(m *base.Module, l0 int32)
//go:linkname F_MemoryContextReset github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextReset
func F_MemoryContextReset(m *base.Module, l0 int32)
//go:linkname F_MemoryContextDeleteChildren github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MemoryContextDeleteChildren
func F_MemoryContextDeleteChildren(m *base.Module, l0 int32)
//go:linkname F_MemoryContextDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextDelete
func F_MemoryContextDelete(m *base.Module, l0 int32)
//go:linkname F_MemoryContextSetParent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextSetParent
func F_MemoryContextSetParent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetMemoryChunkContext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetMemoryChunkContext
func F_GetMemoryChunkContext(m *base.Module, l0 int32) int32
//go:linkname F_GetMemoryChunkSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetMemoryChunkSpace
func F_GetMemoryChunkSpace(m *base.Module, l0 int32) int32
//go:linkname F_MemoryContextAllocationFailure github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextAllocationFailure
func F_MemoryContextAllocationFailure(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MemoryContextAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextAlloc
func F_MemoryContextAlloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextAllocZero github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoryContextAllocZero
func F_MemoryContextAllocZero(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextAllocExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextAllocExtended
func F_MemoryContextAllocExtended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_palloc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_palloc
func F_palloc(m *base.Module, l0 int32) int32
//go:linkname F_palloc0 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_palloc0
func F_palloc0(m *base.Module, l0 int32) int32
//go:linkname F_palloc_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_palloc_extended
func F_palloc_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextAllocAligned github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoryContextAllocAligned
func F_MemoryContextAllocAligned(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pfree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pfree
func F_pfree(m *base.Module, l0 int32)
//go:linkname F_repalloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_repalloc
func F_repalloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextAllocHuge github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextAllocHuge
func F_MemoryContextAllocHuge(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pstrdup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pstrdup
func F_pstrdup(m *base.Module, l0 int32) int32
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
//go:linkname F_ReleaseAuxProcessResources github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReleaseAuxProcessResources
func F_ReleaseAuxProcessResources(m *base.Module, l0 int32)
//go:linkname F_ResourceOwnerForgetLock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResourceOwnerForgetLock
func F_ResourceOwnerForgetLock(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LogicalTapeSetCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalTapeSetCreate
func F_LogicalTapeSetCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ltsWriteBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltsWriteBlock
func F_ltsWriteBlock(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_ltsReadFillBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltsReadFillBuffer
func F_ltsReadFillBuffer(m *base.Module, l0 int32) int32
//go:linkname F_qsort_interruptible github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_qsort_interruptible
func F_qsort_interruptible(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_sts_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sts_attach
func F_sts_attach(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sts_end_write github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sts_end_write
func F_sts_end_write(m *base.Module, l0 int32)
//go:linkname F_sts_puttuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sts_puttuple
func F_sts_puttuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PrepareSortSupportComparisonShim github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PrepareSortSupportComparisonShim
func F_PrepareSortSupportComparisonShim(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PrepareSortSupportFromOrderingOp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PrepareSortSupportFromOrderingOp
func F_PrepareSortSupportFromOrderingOp(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PrepareSortSupportFromIndexRel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PrepareSortSupportFromIndexRel
func F_PrepareSortSupportFromIndexRel(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_begin_common github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_begin_common
func F_tuplesort_begin_common(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplesort_begin_batch github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_begin_batch
func F_tuplesort_begin_batch(m *base.Module, l0 int32)
//go:linkname F_tuplesort_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_free
func F_tuplesort_free(m *base.Module, l0 int32)
//go:linkname F_tuplesort_puttuple_common github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_puttuple_common
func F_tuplesort_puttuple_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_dumptuples github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dumptuples
func F_dumptuples(m *base.Module, l0 int32, l1 int32)
//go:linkname F_inittapes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_inittapes
func F_inittapes(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplesort_sort_memtuples github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_sort_memtuples
func F_tuplesort_sort_memtuples(m *base.Module, l0 int32)
//go:linkname F_mergeruns github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mergeruns
func F_mergeruns(m *base.Module, l0 int32)
//go:linkname F_tuplesort_gettuple_common github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_gettuple_common
func F_tuplesort_gettuple_common(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplesort_rescan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_rescan
func F_tuplesort_rescan(m *base.Module, l0 int32)
//go:linkname F_tuplesort_begin_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_begin_heap
func F_tuplesort_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_tuplesort_begin_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_begin_datum
func F_tuplesort_begin_datum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_tuplesort_putindextuplevalues github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplesort_putindextuplevalues
func F_tuplesort_putindextuplevalues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_tuplesort_putdatum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_putdatum
func F_tuplesort_putdatum(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_getheaptuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_getheaptuple
func F_tuplesort_getheaptuple(m *base.Module, l0 int32) int32
//go:linkname F_tuplestore_begin_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_begin_heap
func F_tuplestore_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplestore_set_eflags github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplestore_set_eflags
func F_tuplestore_set_eflags(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_alloc_read_pointer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplestore_alloc_read_pointer
func F_tuplestore_alloc_read_pointer(m *base.Module, l0 int32, l1 int32) int32
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
//go:linkname F_tuplestore_gettupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_gettupleslot
func F_tuplestore_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_tuplestore_advance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplestore_advance
func F_tuplestore_advance(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetComboCommandId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetComboCommandId
func F_GetComboCommandId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetLatestSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetLatestSnapshot
func F_GetLatestSnapshot(m *base.Module) int32
//go:linkname F_GetNonHistoricCatalogSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetNonHistoricCatalogSnapshot
func F_GetNonHistoricCatalogSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_PushActiveSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PushActiveSnapshot
func F_PushActiveSnapshot(m *base.Module, l0 int32)
//go:linkname F_PopActiveSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PopActiveSnapshot
func F_PopActiveSnapshot(m *base.Module)
//go:linkname F_RegisterSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RegisterSnapshot
func F_RegisterSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_UnregisterSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_UnregisterSnapshot
func F_UnregisterSnapshot(m *base.Module, l0 int32)
//go:linkname F_UnregisterSnapshotFromOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_UnregisterSnapshotFromOwner
func F_UnregisterSnapshotFromOwner(m *base.Module, l0 int32, l1 int32)
//go:linkname F_EstimateSnapshotSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EstimateSnapshotSpace
func F_EstimateSnapshotSpace(m *base.Module, l0 int32) int32
//go:linkname F_RestoreTransactionSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RestoreTransactionSnapshot
func F_RestoreTransactionSnapshot(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XidInMVCCSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XidInMVCCSnapshot
func F_XidInMVCCSnapshot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_tzset github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_tzset
func F_pg_tzset(m *base.Module, l0 int32) int32
//go:linkname F_pg_tzenumerate_next github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_tzenumerate_next
func F_pg_tzenumerate_next(m *base.Module, l0 int32) int32
//go:linkname F__yconv github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__yconv
func F__yconv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_binaryheap_allocate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_binaryheap_allocate
func F_binaryheap_allocate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_binaryheap_add_unordered github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_binaryheap_add_unordered
func F_binaryheap_add_unordered(m *base.Module, l0 int32, l1 int32)
//go:linkname F_binaryheap_build github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_binaryheap_build
func F_binaryheap_build(m *base.Module, l0 int32)
//go:linkname F_binaryheap_remove_first github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_binaryheap_remove_first
func F_binaryheap_remove_first(m *base.Module, l0 int32) int32
//go:linkname F_binaryheap_replace_first github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_binaryheap_replace_first
func F_binaryheap_replace_first(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateEmptyBlockRefTable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateEmptyBlockRefTable
func F_CreateEmptyBlockRefTable(m *base.Module) int32
//go:linkname F_BlockRefTableMarkBlockModified github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BlockRefTableMarkBlockModified
func F_BlockRefTableMarkBlockModified(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_controlfile github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_controlfile
func F_get_controlfile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_my_exec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_find_my_exec
func F_find_my_exec(m *base.Module, l0 int32, l1 int32) int32
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
//go:linkname F_json_errdetail github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_json_errdetail
func F_json_errdetail(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ScanKeywordLookup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ScanKeywordLookup
func F_ScanKeywordLookup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_md5_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_md5_hash
func F_pg_md5_hash(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_prng_seed github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_prng_seed
func F_pg_prng_seed(m *base.Module, l0 int32, l1 int64)
//go:linkname F_pg_prng_seed_check github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_prng_seed_check
func F_pg_prng_seed_check(m *base.Module, l0 int32) int32
//go:linkname F_pg_prng_uint32 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_prng_uint32
func F_pg_prng_uint32(m *base.Module) int32
//go:linkname F_psprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_psprintf
func F_psprintf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetDatabasePath github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetDatabasePath
func F_GetDatabasePath(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetRelationPath github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetRelationPath
func F_GetRelationPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_rmtree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_rmtree
func F_rmtree(m *base.Module, l0 int32) int32
//go:linkname F_is_code_in_table github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_is_code_in_table
func F_is_code_in_table(m *base.Module, l0 int32) int32
//go:linkname F_scram_H github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scram_H
func F_scram_H(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_clean_ascii github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_clean_ascii
func F_pg_clean_ascii(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeStringInfo
func F_makeStringInfo(m *base.Module) int32
//go:linkname F_initStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initStringInfo
func F_initStringInfo(m *base.Module, l0 int32)
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
//go:linkname F_case_index github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_case_index
func F_case_index(m *base.Module, l0 int32) int32
//go:linkname F_unicode_normalize github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_unicode_normalize
func F_unicode_normalize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_ascii_verifystr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_ascii_verifystr
func F_pg_ascii_verifystr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_encoding_verifymbchar github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_encoding_verifymbchar
func F_pg_encoding_verifymbchar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_cryptohash_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_cryptohash_create
func F_pg_cryptohash_create(m *base.Module, l0 int32) int32
//go:linkname F_pg_cryptohash_update github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_cryptohash_update
func F_pg_cryptohash_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_cryptohash_final github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_cryptohash_final
func F_pg_cryptohash_final(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_cryptohash_free github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_cryptohash_free
func F_pg_cryptohash_free(m *base.Module, l0 int32)
//go:linkname F_pg_hmac_init github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_hmac_init
func F_pg_hmac_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_get_encoding_from_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_get_encoding_from_locale
func F_pg_get_encoding_from_locale(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_inet_net_ntop github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_inet_net_ntop
func F_pg_inet_net_ntop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_relative_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_relative_path
func F_make_relative_path(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_etc_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_etc_path
func F_get_etc_path(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_pkglib_path github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_pkglib_path
func F_get_pkglib_path(m *base.Module, l0 int32)
//go:linkname F_get_doc_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_doc_path
func F_get_doc_path(m *base.Module, l0 int32)
//go:linkname F_pg_strong_random github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_strong_random
func F_pg_strong_random(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_strcasecmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_strcasecmp
func F_pg_strcasecmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_toupper github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_toupper
func F_pg_toupper(m *base.Module, l0 int32) int32
//go:linkname F_pg_tolower github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_tolower
func F_pg_tolower(m *base.Module, l0 int32) int32
//go:linkname F_pqsignal_be github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pqsignal_be
func F_pqsignal_be(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_qsort github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_qsort
func F_pg_qsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_qsort_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_qsort_arg
func F_qsort_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_pg_vsnprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_vsnprintf
func F_pg_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_dopr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dopr
func F_dopr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_sprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_sprintf
func F_pg_sprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_fprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_fprintf
func F_pg_fprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgl_longjmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgl_longjmp
func F_pgl_longjmp(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgl_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgl_recv
func F_pgl_recv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pgmem_dlsym github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgmem_dlsym
func F_pgmem_dlsym(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_compile github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plpgsql_compile
func F_plpgsql_compile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_build_datatype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_datatype
func F_build_datatype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_plpgsql_compile_inline github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_plpgsql_compile_inline
func F_plpgsql_compile_inline(m *base.Module, l0 int32) int32
//go:linkname F_plpgsql_estate_setup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_estate_setup
func F_plpgsql_estate_setup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_assign_simple_var github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_assign_simple_var
func F_assign_simple_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_exec_move_row_from_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_move_row_from_datum
func F_exec_move_row_from_datum(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_exec_move_row github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_exec_move_row
func F_exec_move_row(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_coerce_function_result_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_coerce_function_result_tuple
func F_coerce_function_result_tuple(m *base.Module, l0 int32, l1 int32)
//go:linkname F_exec_cast_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_cast_value
func F_exec_cast_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_plpgsql_create_econtext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plpgsql_create_econtext
func F_plpgsql_create_econtext(m *base.Module, l0 int32)
//go:linkname F_exec_eval_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_exec_eval_datum
func F_exec_eval_datum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_exec_assign_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_assign_value
func F_exec_assign_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_exec_assign_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_assign_expr
func F_exec_assign_expr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_exec_stmts github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_exec_stmts
func F_exec_stmts(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_free_function_memory github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plpgsql_free_function_memory
func F_plpgsql_free_function_memory(m *base.Module, l0 int32)
//go:linkname F_plpgsql_scanner_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_plpgsql_scanner_errposition
func F_plpgsql_scanner_errposition(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SN_create_env github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SN_create_env
func F_SN_create_env(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_r_fix_chdz github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_r_fix_chdz
func F_r_fix_chdz(m *base.Module, l0 int32) int32
//go:linkname F_r_fix_ending github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_r_fix_ending
func F_r_fix_ending(m *base.Module, l0 int32) int32
//go:linkname F_r_fix_va_start github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_r_fix_va_start
func F_r_fix_va_start(m *base.Module, l0 int32) int32
//go:linkname F_lose_s github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lose_s
func F_lose_s(m *base.Module, l0 int32)
//go:linkname F_find_among github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_find_among
func F_find_among(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_among_b github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_among_b
func F_find_among_b(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_replace_s github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_replace_s
func F_replace_s(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_slice_from_s github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_slice_from_s
func F_slice_from_s(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_slice_to github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_slice_to
func F_slice_to(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_des_init github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_des_init
func F_des_init(m *base.Module)
//go:linkname F_des_setkey github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_des_setkey
func F_des_setkey(m *base.Module, l0 int32)
//go:linkname F_do_des github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_do_des
func F_do_des(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F__crypt_gensalt_sha github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__crypt_gensalt_sha
func F__crypt_gensalt_sha(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pullf_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pullf_create
func F_pullf_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pushf_write github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pushf_write
func F_pushf_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_px_find_digest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_px_find_digest
func F_px_find_digest(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_new_len github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parse_new_len
func F_parse_new_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_mpi_check github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mpi_check
func F_mpi_check(m *base.Module, l0 int32) int32
//go:linkname F_bytes_to_mpi github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bytes_to_mpi
func F_bytes_to_mpi(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgp_mpi_free github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgp_mpi_free
func F_pgp_mpi_free(m *base.Module, l0 int32) int32
//go:linkname F_pgp_load_digest github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_load_digest
func F_pgp_load_digest(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_px_crypt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_px_crypt
func F_px_crypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_citextcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_citextcmp
func F_citextcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_generate_trgm_only github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_generate_trgm_only
func F_generate_trgm_only(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_calc_word_similarity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_calc_word_similarity
func F_calc_word_similarity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float32
//go:linkname F_hstoreUpgrade github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstoreUpgrade
func F_hstoreUpgrade(m *base.Module, l0 int32) int32
//go:linkname F_hstoreUniquePairs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstoreUniquePairs
func F_hstoreUniquePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hstorePairs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hstorePairs
func F_hstorePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ltree_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ltree_concat
func F_ltree_concat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_infix_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_infix_2
func F_infix_2(m *base.Module, l0 int32, l1 int32)
//go:linkname F_gbt_num_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_num_compress
func F_gbt_num_compress(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gbt_num_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_num_fetch
func F_gbt_num_fetch(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gbt_num_union github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_num_union
func F_gbt_num_union(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gbt_num_same github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_num_same
func F_gbt_num_same(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gbt_num_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_num_consistent
func F_gbt_num_consistent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_gbt_num_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_num_distance
func F_gbt_num_distance(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_gbt_num_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_num_picksplit
func F_gbt_num_picksplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gin_btree_extract_query github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gin_btree_extract_query
func F_gin_btree_extract_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F___memcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___memcpy
func F___memcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__emscripten_memcpy_bulkmem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_memcpy_bulkmem
func F__emscripten_memcpy_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___memset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___memset
func F___memset(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__emscripten_memset_bulkmem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_memset_bulkmem
func F__emscripten_memset_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_atan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_atan
func F_atan(m *base.Module, l0 float64) float64
//go:linkname F___isspace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___isspace
func F___isspace(m *base.Module, l0 int32) int32
//go:linkname F_bsearch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bsearch
func F_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_close github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_close
func F_close(m *base.Module, l0 int32) int32
//go:linkname F___cos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___cos
func F___cos(m *base.Module, l0 float64, l1 float64) float64
//go:linkname F___rem_pio2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___rem_pio2
func F___rem_pio2(m *base.Module, l0 float64, l1 int32) int32
//go:linkname F___sin github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___sin
func F___sin(m *base.Module, l0 float64, l1 float64, l2 int32) float64
//go:linkname F___time github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___time
func F___time(m *base.Module) int64
//go:linkname F___gettimeofday github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___gettimeofday
func F___gettimeofday(m *base.Module, l0 int32)
//go:linkname F_fgets github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fgets
func F_fgets(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fopen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fopen
func F_fopen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___fstatat github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___fstatat
func F___fstatat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_fsync github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fsync
func F_fsync(m *base.Module, l0 int32) int32
//go:linkname F_ftruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ftruncate
func F_ftruncate(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_fwrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fwrite
func F_fwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getenv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getenv
func F_getenv(m *base.Module, l0 int32) int32
//go:linkname F_getopt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getopt
func F_getopt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getrusage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getrusage
func F_getrusage(m *base.Module, l0 int32)
//go:linkname F_isalnum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isalnum
func F_isalnum(m *base.Module, l0 int32) int32
//go:linkname F_iswalpha github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_iswalpha
func F_iswalpha(m *base.Module, l0 int32) int32
//go:linkname F_ldexp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ldexp
func F_ldexp(m *base.Module, l0 float64, l1 int32) float64
//go:linkname F_log github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_log
func F_log(m *base.Module, l0 float64) float64
//go:linkname F_memcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_memcmp
func F_memcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mkdir github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mkdir
func F_mkdir(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_open github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_open
func F_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_opendir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_opendir
func F_opendir(m *base.Module, l0 int32) int32
//go:linkname F_pipe github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pipe
func F_pipe(m *base.Module, l0 int32) int32
//go:linkname F_pow github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pow
func F_pow(m *base.Module, l0 float64, l1 float64) float64
//go:linkname F_pread github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pread
func F_pread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_pwrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pwrite
func F_pwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_pwritev github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pwritev
func F_pwritev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_raise github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_raise
func F_raise(m *base.Module, l0 int32) int32
//go:linkname F_read github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_read
func F_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_readdir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_readdir
func F_readdir(m *base.Module, l0 int32) int32
//go:linkname F_readlink github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_readlink
func F_readlink(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_scalbn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_scalbn
func F_scalbn(m *base.Module, l0 float64, l1 int32) float64
//go:linkname F__emscripten_timeout github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_timeout
func F__emscripten_timeout(m *base.Module, l0 int32, l1 float64)
//go:linkname F_sigdelset github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sigdelset
func F_sigdelset(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sigemptyset github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sigemptyset
func F_sigemptyset(m *base.Module, l0 int32)
//go:linkname F_sigprocmask github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sigprocmask
func F_sigprocmask(m *base.Module, l0 int32, l1 int32)
//go:linkname F_snprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_snprintf
func F_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_sscanf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sscanf
func F_sscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___strchrnul github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___strchrnul
func F___strchrnul(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strcmp
func F_strcmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strcpy
func F_strcpy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strlcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strlcpy
func F_strlcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strlen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strlen
func F_strlen(m *base.Module, l0 int32) int32
//go:linkname F_strncmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strncmp
func F_strncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strncpy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strncpy
func F_strncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strstr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strstr
func F_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strtod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strtod
func F_strtod(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_strtox_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strtox_2
func F_strtox_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64
//go:linkname F_strtol github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strtol
func F_strtol(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___syscall_ret github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___syscall_ret
func F___syscall_ret(m *base.Module, l0 int32) int32
//go:linkname F___tan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___tan
func F___tan(m *base.Module, l0 float64, l1 float64, l2 int32) float64
//go:linkname F_tolower github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tolower
func F_tolower(m *base.Module, l0 int32) int32
//go:linkname F_toupper github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_toupper
func F_toupper(m *base.Module, l0 int32) int32
//go:linkname F_towlower github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_towlower
func F_towlower(m *base.Module, l0 int32) int32
//go:linkname F_towupper github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_towupper
func F_towupper(m *base.Module, l0 int32) int32
//go:linkname F_unlink github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_unlink
func F_unlink(m *base.Module, l0 int32) int32
//go:linkname F___vfprintf_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___vfprintf_internal
func F___vfprintf_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_vfprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_vfprintf
func F_vfprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_write github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_write
func F_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_emscripten_builtin_malloc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_emscripten_builtin_malloc
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32
//go:linkname F_emscripten_builtin_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_emscripten_builtin_free
func F_emscripten_builtin_free(m *base.Module, l0 int32)
//go:linkname F_emscripten_builtin_realloc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_emscripten_builtin_realloc
func F_emscripten_builtin_realloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___wasm_longjmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___wasm_longjmp
func F___wasm_longjmp(m *base.Module, l0 int32, l1 int32)
//go:linkname F___udivmodti4 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___udivmodti4
func F___udivmodti4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F_connect github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_connect
func F_connect(m *base.Module, l0 int32) int32
//go:linkname F_socket github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_socket
func F_socket(m *base.Module, l0 int32, l1 int32, l2 int32) int32
