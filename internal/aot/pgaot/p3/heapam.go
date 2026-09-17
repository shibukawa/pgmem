package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_heapam_estimate_rel_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 float32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v67 float64
	_ = v67
	var v71 float64
	_ = v71
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v85 float64
	_ = v85
	v14 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
		v19 = *(*float32)(unsafe.Add(mBase, uint32(v16)+100))
		if base.B2i32(base.F32_lt(v19, float32(0)) == int32(0))|base.B2i32(base.Ui32(int32(9)) < base.Ui32(v14)) != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v14
			if v14 != 0 {
				v36 = v14
				v39 = int32(0)
				if base.B2i32(base.F32_ge(v19, float32(0)) == v39)|base.B2i32(v18 == v39) != 0 {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
					if v44 != 0 {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
						v47 = v45
					} else {
						v47 = int32(100)
					}
					v51 = base.I32_div_u_s(v47*int32(_a_F_heapam_estimate_rel_size_0), int32(100))
					v52 = F_get_rel_data_width(m, l0, l1)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v56 = base.I32_div_u_s(v51, v52+int32(28))
						v57 = base.F64_convert_i32_u(v56)
						v58 = float64(1e+100)
						if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)))|base.F64_gt(v57, v58) != 0 {
							v71 = v58
						} else {
							v67 = float64(1)
							if base.F64_le(v57, v67) != 0 {
								v71 = v67
							} else {
								v71 = base.F64_nearest(v57)
							}
						}
						v76 = v71
						v77 = base.F64_convert_i32_u(v36)
						*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_nearest(base.F64_mul(v76, v77))
						if v17 == int32(0) {
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(0)
						} else {
							v85 = base.F64_convert_i32_u(v17)
							if base.F64_le(v77, v85) != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(1)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_div(v85, v77)
							}
						}
						return
					}
				} else {
					v76 = base.F64_div(base.F64_promote_f32(v19), base.F64_convert_i32_u(v18))
					v77 = base.F64_convert_i32_u(v36)
					*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_nearest(base.F64_mul(v76, v77))
					if v17 == int32(0) {
						*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(0)
					} else {
						v85 = base.F64_convert_i32_u(v17)
						if base.F64_le(v77, v85) != 0 {
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(1)
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_div(v85, v77)
						}
					}
					return
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
				*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(0)
				return
			}
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+126)))
			if v27 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v14
				if v14 != 0 {
					v36 = v14
					v39 = int32(0)
					if base.B2i32(base.F32_ge(v19, float32(0)) == v39)|base.B2i32(v18 == v39) != 0 {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
						if v44 != 0 {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
							v47 = v45
						} else {
							v47 = int32(100)
						}
						v51 = base.I32_div_u_s(v47*int32(_a_F_heapam_estimate_rel_size_0), int32(100))
						v52 = F_get_rel_data_width(m, l0, l1)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v56 = base.I32_div_u_s(v51, v52+int32(28))
							v57 = base.F64_convert_i32_u(v56)
							v58 = float64(1e+100)
							if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)))|base.F64_gt(v57, v58) != 0 {
								v71 = v58
							} else {
								v67 = float64(1)
								if base.F64_le(v57, v67) != 0 {
									v71 = v67
								} else {
									v71 = base.F64_nearest(v57)
								}
							}
							v76 = v71
							v77 = base.F64_convert_i32_u(v36)
							*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_nearest(base.F64_mul(v76, v77))
							if v17 == int32(0) {
								*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(0)
							} else {
								v85 = base.F64_convert_i32_u(v17)
								if base.F64_le(v77, v85) != 0 {
									*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(1)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_div(v85, v77)
								}
							}
							return
						}
					} else {
						v76 = base.F64_div(base.F64_promote_f32(v19), base.F64_convert_i32_u(v18))
						v77 = base.F64_convert_i32_u(v36)
						*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_nearest(base.F64_mul(v76, v77))
						if v17 == int32(0) {
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(0)
						} else {
							v85 = base.F64_convert_i32_u(v17)
							if base.F64_le(v77, v85) != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(1)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_div(v85, v77)
							}
						}
						return
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
					*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(0)
					return
				}
			} else {
				v28 = int32(10)
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v28
				v36 = v28
				v39 = int32(0)
				if base.B2i32(base.F32_ge(v19, float32(0)) == v39)|base.B2i32(v18 == v39) != 0 {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
					if v44 != 0 {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
						v47 = v45
					} else {
						v47 = int32(100)
					}
					v51 = base.I32_div_u_s(v47*int32(_a_F_heapam_estimate_rel_size_0), int32(100))
					v52 = F_get_rel_data_width(m, l0, l1)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v56 = base.I32_div_u_s(v51, v52+int32(28))
						v57 = base.F64_convert_i32_u(v56)
						v58 = float64(1e+100)
						if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)))|base.F64_gt(v57, v58) != 0 {
							v71 = v58
						} else {
							v67 = float64(1)
							if base.F64_le(v57, v67) != 0 {
								v71 = v67
							} else {
								v71 = base.F64_nearest(v57)
							}
						}
						v76 = v71
						v77 = base.F64_convert_i32_u(v36)
						*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_nearest(base.F64_mul(v76, v77))
						if v17 == int32(0) {
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(0)
						} else {
							v85 = base.F64_convert_i32_u(v17)
							if base.F64_le(v77, v85) != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(1)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_div(v85, v77)
							}
						}
						return
					}
				} else {
					v76 = base.F64_div(base.F64_promote_f32(v19), base.F64_convert_i32_u(v18))
					v77 = base.F64_convert_i32_u(v36)
					*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_nearest(base.F64_mul(v76, v77))
					if v17 == int32(0) {
						*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(0)
					} else {
						v85 = base.F64_convert_i32_u(v17)
						if base.F64_le(v77, v85) != 0 {
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = float64(1)
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_div(v85, v77)
						}
					}
					return
				}
			}
		}
	}
}
func F_heapam_index_fetch_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v9 != 0 {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_LockBuffer(m, v28, int32(1))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v35 = l3 + int32(48)
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
			v41 = F_heap_hot_search_buffer(m, l1, v32, v33, l2, v35, l5, (v36^int32(-1))&int32(1))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(l3)+56)) = uint16(v43)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v45
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_LockBuffer(m, v47, int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					if v41 != 0 {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v56 = base.B2i32(v51 != int32(0)) & base.B2i32(v51 != int32(5))
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v56)
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+56))
						*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v59
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_ExecStoreBufferHeapTuple(m, v35, l3, v61)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							return v41
						}
					} else {
						v65 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v65)
						return v41
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
		v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
		v17 = F_ReleaseAndReadBuffer(m, v10, v11, v12|v13<<(uint(int32(16))%32))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17
			if v17 == v10 {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_LockBuffer(m, v28, int32(1))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v35 = l3 + int32(48)
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
					v41 = F_heap_hot_search_buffer(m, l1, v32, v33, l2, v35, l5, (v36^int32(-1))&int32(1))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(l3)+56)) = uint16(v43)
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v45
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_LockBuffer(m, v47, int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v41 != 0 {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v56 = base.B2i32(v51 != int32(0)) & base.B2i32(v51 != int32(5))
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v56)
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+56))
								*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v59
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								F_ExecStoreBufferHeapTuple(m, v35, l3, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									return v41
								}
							} else {
								v65 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v65)
								return v41
							}
						}
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				F_heap_page_prune_opt(m, v23, v17)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_LockBuffer(m, v28, int32(1))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v35 = l3 + int32(48)
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
						v41 = F_heap_hot_search_buffer(m, l1, v32, v33, l2, v35, l5, (v36^int32(-1))&int32(1))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
							*(*uint16)(unsafe.Add(mBase, uint32(l3)+56)) = uint16(v43)
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v45
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							F_LockBuffer(m, v47, int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								if v41 != 0 {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v56 = base.B2i32(v51 != int32(0)) & base.B2i32(v51 != int32(5))
									*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v56)
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+56))
									*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v59
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									F_ExecStoreBufferHeapTuple(m, v35, l3, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										return v41
									}
								} else {
									v65 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v65)
									return v41
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_heapam_tuple_delete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_heap_delete(m, l0, l1, l2, l4, l5, l6, l7)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_heapam_tuple_insert_speculative(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v6 = l5
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v14)
	v19 = F_ExecFetchSlotHeapTuple(m, l1, v14, v12+int32(15))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v21
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
		v25 = int32(_a_F_heapam_tuple_insert_speculative_0)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+16)) = uint16(v25)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+14)) = uint16(v6)
		v28 = int32(16)
		v29 = int32(base.Ui32(v6) >> (uint(v28) % 32))
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+12)) = uint16(v29)
		F_heap_insert(m, l0, v19, l2, l3|v28, l4)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v35)
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v37
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
			if v39 == int32(1) {
				F_pfree(m, v19)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					m.G0 = v12 + int32(16)
					return
				}
			} else {
				m.G0 = v12 + int32(16)
				return
			}
		}
	}
}
