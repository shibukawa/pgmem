package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_HeapTupleGetUpdateXid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_GetMultiXactIdMembers(m, v10, v8+int32(12), v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if int32(0) < v14 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v22 = v2
	goto L8
L4:
	;
	v42 = v2
	goto L5
L5:
	;
	m.G0 = v8 + int32(16)
	return v42
L6:
	;
	F_pfree(m, v20)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v38 = v36
	goto L6
L8:
	;
	v28 = v20 + v22<<(uint(int32(3))%32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v29) {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v38 = int32(0)
	goto L6
L10:
	;
	v33 = v22 + int32(1)
	if v33 != v14 {
		v22 = v33
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v42 = v38
	goto L5
}
func F_HeapTupleHeaderGetCmax(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v4&int32(32) != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderGetCmax[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8+v3<<(uint(int32(3))%32))+4))
		v13 = v12
	} else {
		v13 = v3
	}
	return v13
}
func F_HeapTupleHeaderGetCmin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v4&int32(32) != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderGetCmin[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8+v3<<(uint(int32(3))%32))))
		v13 = v12
	} else {
		v13 = v3
	}
	return v13
}
func F_HeapTupleSatisfiesMVCCBatch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v6
	if l2 <= v6 {
		v68 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v68
L2:
	;
	v28 = v6
	v29 = v6
	goto L3
L3:
	;
	v36 = l3 + v29*int32(20)
	v39 = F_HeapTupleSatisfiesMVCC(m, v36, l0, l1, v14+int32(12))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v55 != int32(2) {
		v68 = v51
		goto L1
	} else {
		goto L11
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v29+(l3+int32(_a_F_HeapTupleSatisfiesMVCCBatch_0))))) = uint8(v39)
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v44 = int32(1)
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(l4+v28<<(uint(v44)%32)))) = uint16(v47)
	v51 = v28 + v44
	goto L9
L8:
	;
	v51 = v28
	goto L9
L9:
	;
	v53 = v29 + int32(1)
	if v53 != l2 {
		v28 = v51
		v29 = v53
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	v58 = int32(1)
	F_BufferFinishSetHintBits(m, l1, v58, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v68 = v51
	goto L1
}
func F_heap_getattr_7(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v85 int64
	_ = v85
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+18)))
	if base.Ui32(v14&int32(2047)) < base.Ui32(l1) {
		v18 = F_getmissingattr(m, l2, l1, l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v85 = v18
			m.G0 = v11 + int32(16)
			return v85
		}
	} else {
		v22 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v22)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+20)))
		if v25&int32(1) == v22 {
			v34 = l2 + l1<<(uint(int32(3))%32) + int32(20)
			v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34))))
			if v35 < int32(0) {
				v78 = F_nocachegetattr(m, l0, l1, l2)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int64(0)
				} else {
					v85 = v78
					m.G0 = v11 + int32(16)
					return v85
				}
			} else {
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
				v40 = v24 + v38 + v35
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
				if v41 == int32(1) {
					v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+2)))
					if base.I32_popcnt(v44) != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v44
							F_errmsg_internal(m, int32(_a_F_heap_getattr_7_0), v11)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_heap_getattr_7_1), int32(123), int32(_a_F_heap_getattr_7_2))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						switch base.I32_ctz(v44) {
						case 0:
							v49 = int64(*(*int8)(unsafe.Add(mBase, uint32(v40))))
							v85 = v49
							m.G0 = v11 + int32(16)
							return v85
						case 1:
							v50 = int64(*(*int16)(unsafe.Add(mBase, uint32(v40))))
							v85 = v50
							m.G0 = v11 + int32(16)
							return v85
						case 2:
							v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40))))
							v85 = v51
							m.G0 = v11 + int32(16)
							return v85
						case 3:
							v52 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
							v85 = v52
							m.G0 = v11 + int32(16)
							return v85
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v44
								F_errmsg_internal(m, int32(_a_F_heap_getattr_7_0), v11)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_heap_getattr_7_1), int32(123), int32(_a_F_heap_getattr_7_2))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					v85 = base.I64_extend_i32_u(v40)
					m.G0 = v11 + int32(16)
					return v85
				}
			}
		} else {
			v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+23)))
			v68 = int32(1)
			if int32(base.Ui32(v67)>>(uint(l1-v68)%32))&v68 != 0 {
				v78 = F_nocachegetattr(m, l0, l1, l2)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int64(0)
				} else {
					v85 = v78
					m.G0 = v11 + int32(16)
					return v85
				}
			} else {
				v73 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v73)
				v85 = int64(0)
				m.G0 = v11 + int32(16)
				return v85
			}
		}
	}
}
func F_heap_getnextslot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v9&int32(1) != 0 {
		F_heapgettup_pagemode(m, l0, l1, v8, v7)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			if v18 == int32(0) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
				m.T0[v22].(func(*base.Module, int32))(m, l2)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return base.B2i32(v18 != int32(0))
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+272))
				if v28 == int32(0) {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+268)))
					if v31 != int32(1) {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
						F_ExecStoreBufferHeapTuple(m, l0+int32(68), l2, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v18 != int32(0))
						}
					} else {
						F_pgstat_assoc_relation(m, v27)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+272))
							v38 = v37
							v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39 + int64(1)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
							F_ExecStoreBufferHeapTuple(m, l0+int32(68), l2, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v18 != int32(0))
							}
						}
					}
				} else {
					v38 = v28
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39 + int64(1)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					F_ExecStoreBufferHeapTuple(m, l0+int32(68), l2, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v18 != int32(0))
					}
				}
			}
		}
	} else {
		F_heapgettup(m, l0, l1, v8, v7)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			if v18 == int32(0) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
				m.T0[v22].(func(*base.Module, int32))(m, l2)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return base.B2i32(v18 != int32(0))
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+272))
				if v28 == int32(0) {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+268)))
					if v31 != int32(1) {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
						F_ExecStoreBufferHeapTuple(m, l0+int32(68), l2, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v18 != int32(0))
						}
					} else {
						F_pgstat_assoc_relation(m, v27)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+272))
							v38 = v37
							v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39 + int64(1)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
							F_ExecStoreBufferHeapTuple(m, l0+int32(68), l2, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v18 != int32(0))
							}
						}
					}
				} else {
					v38 = v28
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39 + int64(1)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					F_ExecStoreBufferHeapTuple(m, l0+int32(68), l2, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v18 != int32(0))
					}
				}
			}
		}
	}
}
func F_heap_multi_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v372 int32
	_ = v372
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v425 int32
	_ = v425
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v595 int32
	_ = v595
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v651 int32
	_ = v651
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int64
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v782 int32
	_ = v782
	var v789 int32
	_ = v789
	var v815 int32
	_ = v815
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v857 int32
	_ = v857
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v932 int64
	_ = v932
	var v933 int32
	_ = v933
	var v935 int64
	_ = v935
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v969 int64
	_ = v969
	var v970 int32
	_ = v970
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1091 int32
	_ = v1091
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1171 int32
	_ = v1171
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1249 int32
	_ = v1249
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1329 int32
	_ = v1329
	v7 = int32(0)
	v38 = m.G0
	v40 = v38 - int32(_a_F_heap_multi_insert_0)
	m.G0 = v40
	v42 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v44
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_multi_insert[0])))
	v48 = int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[1]))
	if v47&v48|base.B2i32(v48 < v51) == v44 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v82 = int32(1)
	if base.B2i32(v78&v82 == int32(0))&base.B2i32(v79 <= v82) != 0 {
		v114 = v7
		goto L15
	} else {
		goto L16
	}
L4:
	;
	v78 = int32(0)
	v79 = v51
	v81 = v7
	goto L3
L5:
	;
	goto L6
L6:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+118)))
	if v59 != int32(112) {
		v78 = v47
		v79 = v51
		v81 = v7
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v51 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v64 != 0 {
		v78 = v47
		v79 = v51
		v81 = v7
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+119)))
	if v66 == int32(102) {
		v78 = v47
		v79 = v51
		v81 = v7
		goto L3
	} else {
		goto L13
	}
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v65 != 0 {
		v78 = v47
		v79 = v51
		v81 = v7
		goto L3
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L14
L14:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_multi_insert[0])))
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[1]))
	v78 = v75
	v79 = v77
	v81 = base.B2i32(base.Ui32(v69) < base.Ui32(int32(_a_F_heap_multi_insert_1))) ^ int32(1)
	goto L3
L15:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+118)))
	if v117 != int32(112) {
		v130 = int32(0)
		goto L28
	} else {
		goto L29
	}
L16:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+118)))
	if v90 != int32(112) {
		v114 = v7
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if int32(0) < v79 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L22
L19:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v95 != 0 {
		v114 = v7
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v96 == int32(0) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v114 = v7
	goto L15
L22:
	;
	if base.Ui32(v100) < base.Ui32(int32(_a_F_heap_multi_insert_1)) {
		v114 = int32(1)
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v103 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v114 = int32(0)
	goto L15
L25:
	;
	goto L26
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+119)))
	switch v109 - int32(109) {
	case 0, 5:
		goto L27
	default:
		v114 = int32(0)
		goto L15
	}
L27:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+112)))
	v114 = v112
	goto L15
L28:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v131 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[1]))
	if int32(0) < v122 {
		v130 = int32(1)
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v126 != 0 {
		v130 = int32(0)
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v130 = base.B2i32(v127 == int32(0))
	goto L28
L32:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v138 = base.I32_div_s(int32(_a_F_heap_multi_insert_2)-v133<<(uint(int32(13))%32), int32(100))
	v139 = v138
	goto L34
L33:
	;
	v139 = v7
	goto L34
L34:
	;
	v142 = F_palloc(m, l2<<(uint(int32(2))%32))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if int32(0) < l2 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	if v1067 != 0 {
		goto L199
	} else {
		goto L200
	}
L37:
	;
	v153 = int32(0)
	goto L40
L38:
	;
	goto L39
L39:
	;
	F_CheckForSerializableConflictIn(m, l0, int32(0), int32(-1))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L198
	}
L40:
	;
	v185 = v153 << (uint(int32(2)) % 32)
	v186 = l1 + v185
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v190 = F_ExecFetchSlotHeapTuple(m, v187, int32(1), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	F_CheckForSerializableConflictIn(m, l0, int32(0), int32(-1))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L45
	}
L42:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+40)) = v193
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+12)) = v196
	v199 = F_heap_prepare_insert(m, l0, v190, v42, l3, l4)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185+v142))) = v199
	v203 = v153 + int32(1)
	if v203 != l2 {
		v153 = v203
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	if v81 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v211 = int32(24)
	goto L48
L47:
	;
	v211 = int32(8)
	goto L48
L48:
	;
	v212 = int32(4)
	v213 = l4 & v212
	v215 = base.B2i32(v213 == int32(0))
	v219 = v114 & v130
	v221 = int32(_a_F_heap_multi_insert_3) - v139
	v225 = v40 + int32(16) | v212
	v237 = v7
	v242 = v7
	v244 = v7
	v253 = v7
	goto L49
L49:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[2]))
	if v264 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L36
L51:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v267 = int32(0)
	if v242&base.B2i32(v237 != v267) == v267 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L53
L55:
	;
	v447 = v142 + v237<<(uint(int32(2))%32)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	v450 = int32(0)
	v455 = F_RelationGetBufferForTuple(m, l0, v449, v450, l4, l5, v40+int32(12), v450, v425-v444)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L75
	}
L56:
	;
	v272 = int32(1)
	if v237+v272 == l2 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	goto L58
L58:
	;
	v425 = v244
	v444 = v253 + int32(1)
	goto L55
L59:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v142+v361<<(uint(int32(2))%32))))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	v425 = v372 + base.B2i32(base.Ui32(v360) < base.Ui32((v395+int32(7))&int32(-8)|int32(4)))
	v444 = int32(0)
	goto L55
L60:
	;
	v360 = v221
	v361 = v237
	v372 = v272
	goto L59
L61:
	;
	goto L62
L62:
	;
	v276 = l2 - v237
	v288 = v221
	v289 = v237
	v291 = int32(0)
	v300 = v272
	goto L63
L63:
	;
	v321 = v142 + v289<<(uint(int32(2))%32)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v329 = (v323+int32(7))&int32(-8) | int32(4)
	v330 = base.B2i32(base.Ui32(v288) < base.Ui32(v329))
	if base.Ui32(v288) < base.Ui32(v329) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v351 = int32(0)
	if v276&int32(1) == v351 {
		v425 = v345
		v444 = v351
		goto L55
	} else {
		goto L72
	}
L65:
	;
	v331 = v221
	goto L67
L66:
	;
	v331 = v288
	goto L67
L67:
	;
	v332 = v331 - v329
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v340 = (v334+int32(7))&int32(-8) | int32(4)
	v341 = base.B2i32(base.Ui32(v332) < base.Ui32(v340))
	if base.Ui32(v332) < base.Ui32(v340) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v342 = v221
	goto L70
L69:
	;
	v342 = v332
	goto L70
L70:
	;
	v343 = v342 - v340
	v345 = v330 + v300 + v341
	v346 = int32(2)
	v347 = v289 + v346
	v349 = v291 + v346
	if v349 != v276&int32(-2) {
		v288 = v343
		v289 = v347
		v291 = v349
		v300 = v345
		goto L63
	} else {
		goto L71
	}
L71:
	;
	goto L64
L72:
	;
	v360 = v343
	v361 = v347
	v372 = v345
	goto L59
L73:
	;
	v500 = int32(_a_F_heap_multi_insert_4)
	v502 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[3])) = v502 + int32(1)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	F_RelationPutHeapTuple(m, v455, v506, int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L85
	}
L74:
	;
	v475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474)+12)))
	v482 = int32(0)
	v484 = base.B2i32(base.Ui32(v475) < base.Ui32(int32(25))) | base.B2i32((v475+int32(_a_F_heap_multi_insert_5))&int32(_a_F_heap_multi_insert_6) == v482)
	v487 = v215 | base.B2i32(v484 == v482)
	if v487 != 0 {
		goto L79
	} else {
		goto L80
	}
L75:
	;
	if v455 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v460 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[4]))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v460+(v455^int32(-1))<<(uint(int32(2))%32))))
	v474 = v466
	goto L74
L77:
	;
	goto L78
L78:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[5]))
	v474 = v468 + v455<<(uint(int32(13))%32) + int32(-8192)
	goto L74
L79:
	;
	v488 = int32(0)
	if v213 != 0 {
		v499 = v488
		goto L73
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	F_LockBufferInternal(m, v495, int32(3))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	v489 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474)+10)))
	if v489&int32(4) == int32(0) {
		v499 = v488
		goto L73
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v499 = v487
	goto L73
L85:
	;
	if v219 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	F_log_heap_new_cid(m, l0, v510)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v513 = int32(1)
	v515 = v237 + v513
	if l2 <= v515 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L88
L90:
	;
	v681 = v487 ^ int32(1)
	if v499 != 0 {
		goto L123
	} else {
		goto L124
	}
L91:
	;
	v651 = v513
	v679 = v515
	goto L90
L92:
	;
	goto L93
L93:
	;
	v517 = l2 - v237
	v525 = v515
	v527 = v513
	goto L94
L94:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v142+v525<<(uint(int32(2))%32))))
	v562 = int32(4)
	v563 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474)+14)))
	v564 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474)+12)))
	v565 = v563 - v564
	if v565 <= v562 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v651 = v517
	v679 = l2
	goto L90
L96:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	if base.Ui32(v625) < base.Ui32((v626+int32(7))&int32(-8)+v139) {
		v651 = v527
		v679 = v525
		goto L90
	} else {
		goto L115
	}
L97:
	;
	v568 = v562
	goto L99
L98:
	;
	v568 = v565
	goto L99
L99:
	;
	v570 = v568 - int32(4)
	if v570 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v625 = int32(0)
	goto L96
L101:
	;
	goto L102
L102:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v564) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v625 = v570
	goto L96
L104:
	;
	v581 = int32(base.Ui32(v564+int32(_a_F_heap_multi_insert_5)) >> (uint(int32(2)) % 32))
	goto L106
L105:
	;
	v581 = int32(0)
	goto L106
L106:
	;
	if base.Ui32(v581&int32(_a_F_heap_multi_insert_7)) < base.Ui32(int32(291)) {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+10)))
	if v586&int32(1) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v625 = int32(0)
	goto L96
L109:
	;
	goto L110
L110:
	;
	v595 = int32(1)
	goto L111
L111:
	;
	v604 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474+int32(20)+v595&int32(_a_F_heap_multi_insert_7)<<(uint(int32(2))%32))+1)))
	if v604&int32(384) == int32(0) {
		goto L103
	} else {
		goto L113
	}
L112:
	;
	v625 = int32(0)
	goto L96
L113:
	;
	v610 = v595 + int32(1)
	v611 = int32(_a_F_heap_multi_insert_7)
	if base.Ui32(v610&v611) <= base.Ui32(v581&v611) {
		v595 = v610
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	F_RelationPutHeapTuple(m, v455, v558, int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v219 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_log_heap_new_cid(m, l0, v558)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v639 = v527 + int32(1)
	if v639 != v517 {
		v525 = v639 + v237
		v527 = v639
		goto L94
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	goto L95
L122:
	;
	if v215&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v42)) == int32(0) {
		goto L137
	} else {
		goto L138
	}
L123:
	;
	if v455 < int32(0) {
		goto L127
	} else {
		goto L128
	}
L124:
	;
	goto L125
L125:
	;
	if v487 != 0 {
		v745 = int32(0)
		goto L122
	} else {
		goto L131
	}
L126:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v703 = F_visibilitymap_clear(m, v700, v701, int32(3))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L130
	}
L127:
	;
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[6]))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v685+(v455^int32(-1))*int32(56))+16))
	v700 = v691
	goto L126
L128:
	;
	goto L129
L129:
	;
	v693 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[7]))
	v694 = int32(56)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v693+v455*v694-v694)+16))
	v700 = v699
	goto L126
L130:
	;
	v705 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474)+10)))
	v707 = v705 & int32(_a_F_heap_multi_insert_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v474)+10)) = uint16(v707)
	v745 = v703 | v681
	goto L122
L131:
	;
	v711 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+20)) = v711
	v713 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474)+10)))
	v715 = v713 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v474)+10)) = uint16(v715)
	if v455 < v711 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v736 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = v736
	v738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v738
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v742 = F_visibilitymap_set(m, v735, v740, int32(3))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L1
	} else {
		goto L136
	}
L133:
	;
	v720 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[6]))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v720+(v455^int32(-1))*int32(56))+16))
	v735 = v726
	goto L132
L134:
	;
	goto L135
L135:
	;
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[7]))
	v729 = int32(56)
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v728+v455*v729-v729)+16))
	v735 = v734
	goto L132
L136:
	;
	v745 = int32(1)
	goto L122
L137:
	;
	F_MarkBufferDirty(m, v455)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L143
	}
L138:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v474)+20))
	v749 = int32(0)
	if base.B2i32(base.Ui32(v748) < base.Ui32(int32(3)))|base.B2i32(v749 <= v42-v748) != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v756 = v748
	goto L141
L140:
	;
	v756 = v749
	goto L141
L141:
	;
	if v756 != 0 {
		goto L137
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474)+20)) = v42
	goto L137
L143:
	;
	if v130 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v1011 = int32(_a_F_heap_multi_insert_4)
	v1013 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[3]))
	v1014 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[3])) = v1013 - v1014
	if v499|v681 == v1014 {
		goto L192
	} else {
		goto L193
	}
L145:
	;
	if v487 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v764 = v499
	goto L148
L147:
	;
	v764 = int32(32)
	goto L148
L148:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+16)) = uint8(v764)
	*(*uint16)(unsafe.Add(mBase, uint32(v40)+18)) = uint16(v651)
	v767 = int32(0)
	if v484 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v771 = v767
	goto L151
L150:
	;
	v771 = v651 << (uint(int32(1)) % 32)
	goto L151
L151:
	;
	v772 = v225 + v771
	if int32(0) < v651 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v782 = v772
	v789 = v767
	goto L155
L153:
	;
	v857 = v772
	goto L154
L154:
	;
	if v81 != 0 {
		goto L164
	} else {
		goto L165
	}
L155:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v447+v789<<(uint(int32(2))%32))))
	if v484 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v857 = v846
	goto L154
L157:
	;
	v821 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v815)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v225+v789<<(uint(int32(1))%32)))) = uint16(v821)
	goto L159
L158:
	;
	goto L159
L159:
	;
	v826 = (v782 + int32(1)) & int32(-2)
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v815)+16))
	v828 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v827)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v826)+2)) = uint16(v828)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v815)+16))
	v831 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v830)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v826)+4)) = uint16(v831)
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v815)+16))
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v833)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v826)+6)) = uint8(v834)
	v837 = v826 + int32(7)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v815)))
	v840 = v838 - int32(23)
	if v840 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v815)+16))
	base.MemoryCopy(m, v837, v841+int32(23), v840)
	goto L162
L161:
	;
	goto L162
L162:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v826))) = uint16(v840)
	v846 = v840 + v837
	v848 = v789 + int32(1)
	if v848 != v651 {
		v782 = v846
		v789 = v848
		goto L155
	} else {
		goto L163
	}
L163:
	;
	goto L156
L164:
	;
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+16)))
	v889 = v887 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+16)) = uint8(v889)
	goto L166
L165:
	;
	goto L166
L166:
	;
	if l2 == v679 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+16)))
	v894 = v892 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+16)) = uint8(v894)
	goto L169
L168:
	;
	goto L169
L169:
	;
	v896 = v857 - v772
	F_XLogBeginInsert(m)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	F_XLogRegisterData(m, v40+int32(16), v771+int32(4))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v905 = int32(0)
	if v484 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v908 = int32(6)
	goto L174
L173:
	;
	v908 = v905
	goto L174
L174:
	;
	F_XLogRegisterBuffer(m, v905, v455, v908|v211)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	if v484 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v914 = int32(-48)
	goto L178
L177:
	;
	v914 = int32(80)
	goto L178
L178:
	;
	if v745 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	F_XLogRegisterBuffer(m, int32(1), v916, int32(0))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	F_XLogRegisterBufData(m, int32(0), v772, v896)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L189
	}
L182:
	;
	F_XLogRegisterBufData(m, int32(0), v772, v896)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v924 = int32(_a_F_heap_multi_insert_9)
	v926 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_multi_insert[8])))
	v927 = v926 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_multi_insert[8])) = uint8(v927)
	goto L184
L184:
	;
	v932 = F_XLogInsert(m, int32(9), v914&int32(255))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v935 = base.I64_rotl(v932, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v474))) = v935
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	if v937 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v941 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[4]))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v941+(v937^int32(-1))<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v947))) = v935
	goto L144
L187:
	;
	goto L188
L188:
	;
	v950 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v950+v937<<(uint(int32(13))%32))+uint32(_c_F_heap_multi_insert[9]))) = v935
	goto L144
L189:
	;
	v961 = int32(_a_F_heap_multi_insert_9)
	v963 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_multi_insert[8])))
	v964 = v963 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_multi_insert[8])) = uint8(v964)
	goto L190
L190:
	;
	v969 = F_XLogInsert(m, int32(9), v914&int32(255))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v474))) = base.I64_rotl(v969, int64(32))
	goto L144
L192:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	F_UnlockBuffer(m, v1020)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	F_UnlockReleaseBuffer(m, v455)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L1
	} else {
		goto L196
	}
L195:
	;
	goto L194
L196:
	;
	if v679 < l2 {
		v237 = v679
		v242 = v484
		v244 = v425
		v253 = v444
		goto L49
	} else {
		goto L197
	}
L197:
	;
	goto L50
L198:
	;
	goto L36
L199:
	;
	F_ReleaseBuffer(m, v1067)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v1070 = int32(0)
	F_CheckForSerializableConflictIn(m, l0, v1070, int32(-1))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L1
	} else {
		goto L203
	}
L202:
	;
	goto L201
L203:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L206
L204:
	;
	F_pgstat_count_heap_insert(m, l0, base.I64_extend_i32_s(l2))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L1
	} else {
		goto L222
	}
L205:
	;
	v1171 = int32(0)
	if l2 != int32(1) {
		goto L215
	} else {
		goto L216
	}
L206:
	;
	v1078 = int32(0)
	if base.B2i32(base.B2i32(base.Ui32(v1075) < base.Ui32(int32(_a_F_heap_multi_insert_1))) == v1078)|base.B2i32(l2 <= v1078) == v1078 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1091 = v1070
	goto L210
L208:
	;
	goto L209
L209:
	;
	if l2 <= int32(0) {
		goto L204
	} else {
		goto L214
	}
L210:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v142+v1091<<(uint(int32(2))%32))))
	F_CacheInvalidateHeapTuple(m, l0, v1125, int32(0))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L212
	}
L211:
	;
	goto L205
L212:
	;
	v1130 = v1091 + int32(1)
	if v1130 != l2 {
		v1091 = v1130
		goto L210
	} else {
		goto L213
	}
L213:
	;
	goto L211
L214:
	;
	goto L205
L215:
	;
	v1185 = v1171
	v1188 = int32(0)
	goto L218
L216:
	;
	v1249 = v1171
	goto L217
L217:
	;
	v1281 = v1249 << (uint(int32(2)) % 32)
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(l1+v1281)))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1281+v142)))
	v1286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1285)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1283)+36)) = uint16(v1286)
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1285)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1283)+32)) = v1288
	goto L204
L218:
	;
	v1216 = int32(2)
	v1217 = v1185 << (uint(v1216) % 32)
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l1+v1217)))
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1217+v142)))
	v1222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1221)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1219)+36)) = uint16(v1222)
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1219)+32)) = v1224
	v1227 = v1217 | int32(4)
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(l1+v1227)))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1227+v142)))
	v1232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1231)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1229)+36)) = uint16(v1232)
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+32)) = v1234
	v1237 = v1185 + v1216
	v1239 = v1188 + v1216
	if v1239 != l2&int32(-2) {
		v1185 = v1237
		v1188 = v1239
		goto L218
	} else {
		goto L220
	}
L219:
	;
	if l2&int32(1) == int32(0) {
		goto L204
	} else {
		goto L221
	}
L220:
	;
	goto L219
L221:
	;
	v1249 = v1237
	goto L217
L222:
	;
	m.G0 = v40 + int32(_a_F_heap_multi_insert_0)
	return
}
func F_heap_page_fix_vm_corruption(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v14 = v12 + int32(4)
	v17 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		switch l2 - int32(1) {
		case 0:
			if v17 == int32(0) {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+10)))
				v81 = v79 & int32(_a_F_heap_page_fix_vm_corruption_0)
				*(*uint16)(unsafe.Add(mBase, uint32(v78)+10)) = uint16(v81)
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				F_MarkBufferDirtyHint(m, v83, int32(1))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
					F_LockBufferInternal(m, v90, int32(3))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
						v98 = F_visibilitymap_clear(m, v95, v96, int32(3))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
							F_UnlockBuffer(m, v100)
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return
							} else {
								v103 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[1]))) = uint8(v103)
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			} else {
				v54 = int32(893)
				v55 = int32(_a_F_heap_page_fix_vm_corruption_1)
				F_errcode(m, int32(16779816))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					F_errmsg(m, v55, int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v65
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
							F_errcontext_msg(m, int32(_a_F_heap_page_fix_vm_corruption_2), v9)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_heap_page_fix_vm_corruption_3), v54, int32(_a_F_heap_page_fix_vm_corruption_4))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+10)))
									v81 = v79 & int32(_a_F_heap_page_fix_vm_corruption_0)
									*(*uint16)(unsafe.Add(mBase, uint32(v78)+10)) = uint16(v81)
									v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									F_MarkBufferDirtyHint(m, v83, int32(1))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
										F_LockBufferInternal(m, v90, int32(3))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
											v98 = F_visibilitymap_clear(m, v95, v96, int32(3))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
												F_UnlockBuffer(m, v100)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return
												} else {
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[1]))) = uint8(v103)
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		case 1:
			if v17 == int32(0) {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+10)))
				v81 = v79 & int32(_a_F_heap_page_fix_vm_corruption_0)
				*(*uint16)(unsafe.Add(mBase, uint32(v78)+10)) = uint16(v81)
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				F_MarkBufferDirtyHint(m, v83, int32(1))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
					F_LockBufferInternal(m, v90, int32(3))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
						v98 = F_visibilitymap_clear(m, v95, v96, int32(3))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
							F_UnlockBuffer(m, v100)
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return
							} else {
								v103 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[1]))) = uint8(v103)
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			} else {
				v54 = int32(917)
				v55 = int32(_a_F_heap_page_fix_vm_corruption_5)
				F_errcode(m, int32(16779816))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					F_errmsg(m, v55, int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v65
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
							F_errcontext_msg(m, int32(_a_F_heap_page_fix_vm_corruption_2), v9)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_heap_page_fix_vm_corruption_3), v54, int32(_a_F_heap_page_fix_vm_corruption_4))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+10)))
									v81 = v79 & int32(_a_F_heap_page_fix_vm_corruption_0)
									*(*uint16)(unsafe.Add(mBase, uint32(v78)+10)) = uint16(v81)
									v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									F_MarkBufferDirtyHint(m, v83, int32(1))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
										F_LockBufferInternal(m, v90, int32(3))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
											v98 = F_visibilitymap_clear(m, v95, v96, int32(3))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
												F_UnlockBuffer(m, v100)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return
												} else {
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[1]))) = uint8(v103)
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		default:
			if v17 == int32(0) {
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
				F_LockBufferInternal(m, v90, int32(3))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
					v98 = F_visibilitymap_clear(m, v95, v96, int32(3))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return
					} else {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
						F_UnlockBuffer(m, v100)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							v103 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[1]))) = uint8(v103)
							m.G0 = v9 + int32(32)
							return
						}
					}
				}
			} else {
				F_errcode(m, int32(16779816))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_heap_page_fix_vm_corruption_6), int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v41
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v14
							F_errcontext_msg(m, int32(_a_F_heap_page_fix_vm_corruption_7), v9+int32(16))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_heap_page_fix_vm_corruption_3), int32(938), int32(_a_F_heap_page_fix_vm_corruption_4))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
									F_LockBufferInternal(m, v90, int32(3))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
										v98 = F_visibilitymap_clear(m, v95, v96, int32(3))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[0])))
											F_UnlockBuffer(m, v100)
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												v103 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_page_fix_vm_corruption[1]))) = uint8(v103)
												m.G0 = v9 + int32(32)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_heap_prepare_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_insert[0]))
	if v8 < int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)))
		v14 = v12 & int32(15)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v14)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+18)))
		v19 = v17 & int32(_a_F_heap_prepare_insert_0)
		*(*uint16)(unsafe.Add(mBase, uint32(v16)+18)) = uint16(v19)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+20)))
		v24 = v22 | int32(2048)
		*(*uint16)(unsafe.Add(mBase, uint32(v21)+20)) = uint16(v24)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v26))) = l2
		if l4&int32(4) != 0 {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+20)))
			v33 = v31 | int32(768)
			*(*uint16)(unsafe.Add(mBase, uint32(v30)+20)) = uint16(v33)
		} else {
		}
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = l3
		v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+20)))
		v40 = v38 & int32(_a_F_heap_prepare_insert_1)
		*(*uint16)(unsafe.Add(mBase, uint32(v36)+20)) = uint16(v40)
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = int32(0)
		v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v45
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+119)))
		switch v48 - int32(109) {
		case 0, 5:
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+20)))
			if v52&int32(4) == int32(0) {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if base.Ui32(v57) < base.Ui32(int32(2033)) {
					v65 = l1
					return v65
				} else {
					v61 = F_heap_toast_insert_or_update(m, l0, l1, int32(0), l4)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v65 = v61
						return v65
					}
				}
			} else {
				v61 = F_heap_toast_insert_or_update(m, l0, l1, int32(0), l4)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					v65 = v61
					return v65
				}
			}
		default:
			v65 = l1
			return v65
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v70 = m.ExcPending
		if v70 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(322))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_heap_prepare_insert_2), int32(0))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_heap_prepare_insert_3), int32(2240), int32(_a_F_heap_prepare_insert_4))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_heap_prune_record_unchanged_lp_normal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	v2 = l1
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = l0 + v2
	v14 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[0]))) = uint8(v14)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[1]))) = uint8(v14)
	v18 = int32(2)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12+v2<<(uint(v18)%32))+20))
	v25 = v12 + v22&int32(_a_F_heap_prune_record_unchanged_lp_normal_0)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[2]))))
	switch v26 - v14 {
	case 0:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[3])))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[3]))) = v29 + int32(1)
		v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+20)))
		if v33&int32(256) == int32(0) {
			v38 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[4]))) = uint16(v38)
		} else {
			v40 = int32(768)
			if v33&v40 == v40 {
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v45 = int32(3)
				v46 = base.B2i32(base.Ui32(v44) < base.Ui32(v45))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[5])))
				if v46|base.B2i32(base.Ui32(v47) < base.Ui32(v45)) == int32(0) {
					if int32(0) < v44-v47 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[5]))) = v44
					} else {
					}
				} else {
					if v46|base.B2i32(base.Ui32(v44) <= base.Ui32(v47)) != 0 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[5]))) = v44
					}
				}
			}
		}
		v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
		if v172 != int32(1) {
			m.G0 = v10 + int32(16)
			return
		} else {
			v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v179 = l0 + int32(2380)
			v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v186 = F_heap_prepare_freeze_tuple(m, v25, v175, l0+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v179+v180*int32(12), v10+int32(15))
			mBase = m.M
			v187 = m.ExcPending
			if v187 != 0 {
				return
			} else {
				if v186 != 0 {
					v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v188 + int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v179+v188*int32(12))+10)) = uint16(v2)
				} else {
				}
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
				if v197 != 0 {
				} else {
					v198 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v198)
				}
				m.G0 = v10 + int32(16)
				return
			}
		}
	case 1:
		v59 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[4]))) = uint16(v59)
		v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[7])))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[7]))) = v61 + int32(1)
		v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+20)))
		if v65&int32(_a_F_heap_prune_record_unchanged_lp_normal_2) == int32(_a_F_heap_prune_record_unchanged_lp_normal_3) {
			v70 = F_HeapTupleGetUpdateXid(m, v25)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				v73 = v70
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v74 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v73
				} else {
					v77 = int32(3)
					if base.B2i32(base.Ui32(v73) < base.Ui32(v77))|base.B2i32(base.Ui32(v74) < base.Ui32(v77)) == int32(0) {
						if v73-v74 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v73
						} else {
						}
					} else {
						if base.Ui32(v74) <= base.Ui32(v73) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v73
						}
					}
				}
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)))
				if v90&int32(4) == int32(0) {
					v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
					if v172 != int32(1) {
						m.G0 = v10 + int32(16)
						return
					} else {
						v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v179 = l0 + int32(2380)
						v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v186 = F_heap_prepare_freeze_tuple(m, v25, v175, l0+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v179+v180*int32(12), v10+int32(15))
						mBase = m.M
						v187 = m.ExcPending
						if v187 != 0 {
							return
						} else {
							if v186 != 0 {
								v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v188 + int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v179+v188*int32(12))+10)) = uint16(v2)
							} else {
							}
							v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v197 != 0 {
							} else {
								v198 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v198)
							}
							m.G0 = v10 + int32(16)
							return
						}
					}
				} else {
					F_heap_page_fix_vm_corruption(m, l0, v2, int32(2))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return
					} else {
						v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
						if v172 != int32(1) {
							m.G0 = v10 + int32(16)
							return
						} else {
							v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v179 = l0 + int32(2380)
							v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v186 = F_heap_prepare_freeze_tuple(m, v25, v175, l0+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v179+v180*int32(12), v10+int32(15))
							mBase = m.M
							v187 = m.ExcPending
							if v187 != 0 {
								return
							} else {
								if v186 != 0 {
									v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v188 + int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v179+v188*int32(12))+10)) = uint16(v2)
								} else {
								}
								v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
								if v197 != 0 {
								} else {
									v198 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v198)
								}
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				}
			}
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			v73 = v72
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v74 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v73
			} else {
				v77 = int32(3)
				if base.B2i32(base.Ui32(v73) < base.Ui32(v77))|base.B2i32(base.Ui32(v74) < base.Ui32(v77)) == int32(0) {
					if v73-v74 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v73
					} else {
					}
				} else {
					if base.Ui32(v74) <= base.Ui32(v73) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v73
					}
				}
			}
			v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)))
			if v90&int32(4) == int32(0) {
				v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
				if v172 != int32(1) {
					m.G0 = v10 + int32(16)
					return
				} else {
					v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v179 = l0 + int32(2380)
					v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v186 = F_heap_prepare_freeze_tuple(m, v25, v175, l0+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v179+v180*int32(12), v10+int32(15))
					mBase = m.M
					v187 = m.ExcPending
					if v187 != 0 {
						return
					} else {
						if v186 != 0 {
							v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v188 + int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v179+v188*int32(12))+10)) = uint16(v2)
						} else {
						}
						v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v197 != 0 {
						} else {
							v198 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v198)
						}
						m.G0 = v10 + int32(16)
						return
					}
				}
			} else {
				F_heap_page_fix_vm_corruption(m, l0, v2, int32(2))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return
				} else {
					v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
					if v172 != int32(1) {
						m.G0 = v10 + int32(16)
						return
					} else {
						v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v179 = l0 + int32(2380)
						v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v186 = F_heap_prepare_freeze_tuple(m, v25, v175, l0+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v179+v180*int32(12), v10+int32(15))
						mBase = m.M
						v187 = m.ExcPending
						if v187 != 0 {
							return
						} else {
							if v186 != 0 {
								v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v188 + int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v179+v188*int32(12))+10)) = uint16(v2)
							} else {
							}
							v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v197 != 0 {
							} else {
								v198 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v198)
							}
							m.G0 = v10 + int32(16)
							return
						}
					}
				}
			}
		}
	case 2:
		v98 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[4]))) = uint16(v98)
		v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+20)))
		v101 = int32(768)
		if v100&v101 != v101 {
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
			v106 = v105
		} else {
			v106 = v18
		}
		v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v107 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v106
		} else {
			v110 = int32(3)
			if base.B2i32(base.Ui32(v106) < base.Ui32(v110))|base.B2i32(base.Ui32(v107) < base.Ui32(v110)) == int32(0) {
				if v106-v107 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v106
				} else {
				}
			} else {
				if base.Ui32(v107) <= base.Ui32(v106) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v106
				}
			}
		}
		v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)))
		if v122&int32(4) == int32(0) {
			v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
			if v172 != int32(1) {
				m.G0 = v10 + int32(16)
				return
			} else {
				v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v179 = l0 + int32(2380)
				v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v186 = F_heap_prepare_freeze_tuple(m, v25, v175, l0+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v179+v180*int32(12), v10+int32(15))
				mBase = m.M
				v187 = m.ExcPending
				if v187 != 0 {
					return
				} else {
					if v186 != 0 {
						v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v188 + int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v179+v188*int32(12))+10)) = uint16(v2)
					} else {
					}
					v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
					if v197 != 0 {
					} else {
						v198 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v198)
					}
					m.G0 = v10 + int32(16)
					return
				}
			}
		} else {
			F_heap_page_fix_vm_corruption(m, l0, v2, int32(2))
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return
			} else {
				v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
				if v172 != int32(1) {
					m.G0 = v10 + int32(16)
					return
				} else {
					v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v179 = l0 + int32(2380)
					v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v186 = F_heap_prepare_freeze_tuple(m, v25, v175, l0+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v179+v180*int32(12), v10+int32(15))
					mBase = m.M
					v187 = m.ExcPending
					if v187 != 0 {
						return
					} else {
						if v186 != 0 {
							v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v188 + int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v179+v188*int32(12))+10)) = uint16(v2)
						} else {
						}
						v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v197 != 0 {
						} else {
							v198 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v198)
						}
						m.G0 = v10 + int32(16)
						return
					}
				}
			}
		}
	case 3:
		v130 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[4]))) = uint16(v130)
		v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[3])))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[3]))) = v132 + int32(1)
		v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+20)))
		if v136&int32(_a_F_heap_prune_record_unchanged_lp_normal_2) == int32(_a_F_heap_prune_record_unchanged_lp_normal_3) {
			v141 = F_HeapTupleGetUpdateXid(m, v25)
			mBase = m.M
			v142 = m.ExcPending
			if v142 != 0 {
				return
			} else {
				v144 = v141
				v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v145 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v144
				} else {
					v148 = int32(3)
					if base.B2i32(base.Ui32(v144) < base.Ui32(v148))|base.B2i32(base.Ui32(v145) < base.Ui32(v148)) == int32(0) {
						if v144-v145 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v144
						} else {
						}
					} else {
						if base.Ui32(v145) <= base.Ui32(v144) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v144
						}
					}
				}
				v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+10)))
				if v161&int32(4) == int32(0) {
					v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
					if v172 != int32(1) {
						m.G0 = v10 + int32(16)
						return
					} else {
						v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v179 = l0 + int32(2380)
						v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v186 = F_heap_prepare_freeze_tuple(m, v25, v175, l0+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v179+v180*int32(12), v10+int32(15))
						mBase = m.M
						v187 = m.ExcPending
						if v187 != 0 {
							return
						} else {
							if v186 != 0 {
								v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v188 + int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v179+v188*int32(12))+10)) = uint16(v2)
							} else {
							}
							v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v197 != 0 {
							} else {
								v198 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v198)
							}
							m.G0 = v10 + int32(16)
							return
						}
					}
				} else {
					F_heap_page_fix_vm_corruption(m, l0, v2, int32(2))
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return
					} else {
						v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
						if v172 != int32(1) {
							m.G0 = v10 + int32(16)
							return
						} else {
							v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v179 = l0 + int32(2380)
							v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v186 = F_heap_prepare_freeze_tuple(m, v25, v175, l0+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v179+v180*int32(12), v10+int32(15))
							mBase = m.M
							v187 = m.ExcPending
							if v187 != 0 {
								return
							} else {
								if v186 != 0 {
									v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v188 + int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v179+v188*int32(12))+10)) = uint16(v2)
								} else {
								}
								v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
								if v197 != 0 {
								} else {
									v198 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v198)
								}
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				}
			}
		} else {
			v143 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			v144 = v143
			v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v145 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v144
			} else {
				v148 = int32(3)
				if base.B2i32(base.Ui32(v144) < base.Ui32(v148))|base.B2i32(base.Ui32(v145) < base.Ui32(v148)) == int32(0) {
					if v144-v145 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v144
					} else {
					}
				} else {
					if base.Ui32(v145) <= base.Ui32(v144) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v144
					}
				}
			}
			v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+10)))
			if v161&int32(4) == int32(0) {
				v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
				if v172 != int32(1) {
					m.G0 = v10 + int32(16)
					return
				} else {
					v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v179 = l0 + int32(2380)
					v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v186 = F_heap_prepare_freeze_tuple(m, v25, v175, l0+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v179+v180*int32(12), v10+int32(15))
					mBase = m.M
					v187 = m.ExcPending
					if v187 != 0 {
						return
					} else {
						if v186 != 0 {
							v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v188 + int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v179+v188*int32(12))+10)) = uint16(v2)
						} else {
						}
						v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v197 != 0 {
						} else {
							v198 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v198)
						}
						m.G0 = v10 + int32(16)
						return
					}
				}
			} else {
				F_heap_page_fix_vm_corruption(m, l0, v2, int32(2))
				mBase = m.M
				v168 = m.ExcPending
				if v168 != 0 {
					return
				} else {
					v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
					if v172 != int32(1) {
						m.G0 = v10 + int32(16)
						return
					} else {
						v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v179 = l0 + int32(2380)
						v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v186 = F_heap_prepare_freeze_tuple(m, v25, v175, l0+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v179+v180*int32(12), v10+int32(15))
						mBase = m.M
						v187 = m.ExcPending
						if v187 != 0 {
							return
						} else {
							if v186 != 0 {
								v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v188 + int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v179+v188*int32(12))+10)) = uint16(v2)
							} else {
							}
							v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v197 != 0 {
							} else {
								v198 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v198)
							}
							m.G0 = v10 + int32(16)
							return
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v208 = m.ExcPending
		if v208 != 0 {
			return
		} else {
			v211 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[2]))))
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v211
			F_errmsg_internal(m, int32(_a_F_heap_prune_record_unchanged_lp_normal_4), v10)
			mBase = m.M
			v215 = m.ExcPending
			if v215 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_heap_prune_record_unchanged_lp_normal_5), int32(2000), int32(_a_F_heap_prune_record_unchanged_lp_normal_6))
				mBase = m.M
				v220 = m.ExcPending
				if v220 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_heap_scan_stream_read_next_parallel(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+52)))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v17 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v17
	v22 = v9 + int32(24)
	v24 = int32(-1)
	goto L5
L2:
	;
	goto L3
L3:
	;
	v103 = F_table_block_parallelscan_nextpage(m, v11, v10, v9)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L38
	}
L4:
	;
	v66 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9)+24)), uint32(v66))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	if v69 == int32(-1) {
		goto L25
	} else {
		goto L26
	}
L5:
	;
	v34 = base.AtomicRmwXchg32(m, v22, int32(0), int32(1))
	if v34 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v63
	goto L4
L7:
	;
	F_s_lock(m, v22, int32(_a_F_heap_scan_stream_read_next_parallel_0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if v16 == int32(-1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L9
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if v46 != int32(-1) {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	if v42 != int32(-1) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v16
	goto L12
L15:
	;
	if v15 != int32(-1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L6
L17:
	;
	v63 = v15
	goto L16
L18:
	;
	goto L19
L19:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)))
	if v51 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v63 = int32(0)
	goto L16
L21:
	;
	goto L22
L22:
	;
	if v24 != int32(-1) {
		v63 = v24
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v57 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9)+24)), uint32(v57))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v61 = F_ss_get_location(m, v11, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v24 = v61
	goto L5
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v73 = v72
	goto L27
L26:
	;
	v73 = v69
	goto L27
L27:
	;
	v76 = int32(4095)
	if base.Ui32(v73) <= base.Ui32(v76) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v79 = v76
	goto L30
L29:
	;
	v79 = v73
	goto L30
L30:
	;
	v81 = int32(base.Ui32(v79) >> (uint(int32(11)) % 32))
	if v81&(v81-int32(1)) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v89 = int32(2) << (uint(base.I32_clz(v81)^int32(31)) % 32)
	goto L33
L32:
	;
	v89 = v81
	goto L33
L33:
	;
	if base.Ui32(int32(_a_F_heap_scan_stream_read_next_parallel_1)) <= base.Ui32(v89) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v92 = int32(_a_F_heap_scan_stream_read_next_parallel_1)
	goto L36
L35:
	;
	v92 = v89
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v97 = F_table_block_parallelscan_nextpage(m, v94, v95, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	v99 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+52)) = uint8(v99)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v97
	return v97
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v103
	return v103
}
func F_heap_toast_delete(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	v9 = m.G0
	v11 = v9 - int32(_a_F_heap_toast_delete_0)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v15 = v11 + int32(1600)
	F_heap_deform_tuple(m, l1, v13, v15, v11)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v18 < v20 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = v18
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v11 + int32(_a_F_heap_toast_delete_0)
	return
L6:
	;
	v32 = v24 << (uint(int32(3)) % 32)
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v32)+30)))
	if v34 != int32(_a_F_heap_toast_delete_1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v53 = v24 + int32(1)
	if v53 != v20 {
		v24 = v53
		goto L6
	} else {
		goto L14
	}
L9:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v11))))
	if v38 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v32+v15)))
	v41 = base.I32_wrap_i64(v40)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v42 != int32(1) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v45 != int32(18) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_toast_delete_datum(m, v40, l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L8
L14:
	;
	goto L7
}
func F_heap_truncate_find_FKs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int64
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = F_list_copy(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = v17
	v32 = int32(0)
	goto L4
L4:
	;
	v37 = int32(0)
	v42 = F_systable_beginscan(m, v23, v37, v37, v37, v37, v37)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	F_relation_close(m, v23, int32(1))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L102
	}
L6:
	;
	goto L5
L7:
	;
	F_list_free(m, v325)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L100
	}
L8:
	;
	v230 = v196
	v232 = v205
	v233 = v31
	v237 = v205
	goto L72
L9:
	;
	F_list_free(m, int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L71
	}
L10:
	;
	v44 = F_systable_getnext(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v49 = int32(0)
	v52 = v44
	v54 = v32
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_systable_endscan(m, v42)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L70
	}
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
	v61 = v59 + v60
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+72)))
	if v62 != int32(102) {
		v196 = v49
		v198 = v54
		goto L17
	} else {
		goto L18
	}
L16:
	;
	F_systable_endscan(m, v42)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L67
	}
L17:
	;
	v199 = F_systable_getnext(m, v42)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L65
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+96))
	v66 = int32(0)
	if v31 == v66 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v104 == int32(0) {
		v196 = v49
		v198 = v54
		goto L17
	} else {
		goto L32
	}
L20:
	;
	v104 = int32(0)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v72 <= int32(0) {
		v98 = v66
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v104 = v98
	goto L19
L24:
	;
	v75 = int32(0)
	if v75 < v72 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = v72
	goto L27
L26:
	;
	v78 = v75
	goto L27
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v81 = int32(0)
	goto L28
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v79+v81<<(uint(int32(2))%32))))
	v90 = base.B2i32(v89 == v65)
	if v89 == v65 {
		v98 = v90
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v98 = v90
	goto L23
L30:
	;
	v92 = v81 + int32(1)
	if v92 != v78 {
		v81 = v92
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v61)+92))
	if v107 == int32(0) {
		v152 = v49
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v61)+80))
	v154 = int32(0)
	if l0 == v154 {
		goto L51
	} else {
		goto L52
	}
L34:
	;
	v110 = int32(0)
	if v49 == v110 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v148 != 0 {
		v152 = v49
		goto L33
	} else {
		goto L48
	}
L36:
	;
	v148 = int32(0)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v116 <= int32(0) {
		v142 = v110
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v148 = v142
	goto L35
L40:
	;
	v119 = int32(0)
	if v119 < v116 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v122 = v116
	goto L43
L42:
	;
	v122 = v119
	goto L43
L43:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v125 = int32(0)
	goto L44
L44:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v123+v125<<(uint(int32(2))%32))))
	v134 = base.B2i32(v133 == v107)
	if v133 == v107 {
		v142 = v134
		goto L39
	} else {
		goto L46
	}
L45:
	;
	v142 = v134
	goto L39
L46:
	;
	v136 = v125 + int32(1)
	if v136 != v122 {
		v125 = v136
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v61)+92))
	v150 = F_lappend_oid(m, v49, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v152 = v150
	goto L33
L50:
	;
	if v192 != 0 {
		v196 = v152
		v198 = v54
		goto L17
	} else {
		goto L63
	}
L51:
	;
	v192 = int32(0)
	goto L50
L52:
	;
	goto L53
L53:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v160 <= int32(0) {
		v186 = v154
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v192 = v186
	goto L50
L55:
	;
	v163 = int32(0)
	if v163 < v160 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v166 = v160
	goto L58
L57:
	;
	v166 = v163
	goto L58
L58:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v169 = int32(0)
	goto L59
L59:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v167+v169<<(uint(int32(2))%32))))
	v178 = base.B2i32(v177 == v153)
	if v177 == v153 {
		v186 = v178
		goto L54
	} else {
		goto L61
	}
L60:
	;
	v186 = v178
	goto L54
L61:
	;
	v180 = v169 + int32(1)
	if v180 != v166 {
		v169 = v180
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v61)+80))
	v194 = F_lappend_oid(m, v54, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v196 = v152
	v198 = v194
	goto L17
L65:
	;
	if v199 != 0 {
		v49 = v196
		v52 = v199
		v54 = v198
		goto L15
	} else {
		goto L66
	}
L66:
	;
	goto L16
L67:
	;
	if v196 == int32(0) {
		v219 = v198
		goto L9
	} else {
		goto L68
	}
L68:
	;
	v205 = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v205 < v207 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	v325 = v196
	v328 = v31
	v332 = v205
	goto L7
L70:
	;
	v219 = v32
	goto L9
L71:
	;
	v342 = v31
	v343 = v219
	goto L6
L72:
	;
	v240 = v13 + int32(-56)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	v248 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v244+v232<<(uint(int32(2))%32)))))
	F_ScanKeyInit(m, v240, int32(1), int32(3), int32(184), v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	v325 = v312
	v328 = v313
	v332 = v314
	goto L7
L74:
	;
	v252 = int32(1)
	v255 = F_systable_beginscan(m, v23, int32(2667), v252, int32(0), v252, v240)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L76
	}
L75:
	;
	F_systable_endscan(m, v255)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L98
	}
L76:
	;
	v257 = F_systable_getnext(m, v255)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v257 == int32(0) {
		v312 = v230
		v313 = v233
		v314 = v237
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257)+16))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+22)))
	v263 = v261 + v262
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+92))
	if v264 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v265 = F_list_append_unique_oid(m, v230, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v263)+96))
	v268 = int32(0)
	if v233 == v268 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v312 = v265
	v313 = v233
	v314 = v237
	goto L75
L83:
	;
	if v306 != 0 {
		v312 = v230
		v313 = v233
		v314 = v237
		goto L75
	} else {
		goto L96
	}
L84:
	;
	v306 = int32(0)
	goto L83
L85:
	;
	goto L86
L86:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v274 <= int32(0) {
		v300 = v268
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v306 = v300
	goto L83
L88:
	;
	v277 = int32(0)
	if v277 < v274 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v280 = v274
	goto L91
L90:
	;
	v280 = v277
	goto L91
L91:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v283 = int32(0)
	goto L92
L92:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v281+v283<<(uint(int32(2))%32))))
	v292 = base.B2i32(v291 == v267)
	if v291 == v267 {
		v300 = v292
		goto L87
	} else {
		goto L94
	}
L93:
	;
	v300 = v292
	goto L87
L94:
	;
	v294 = v283 + int32(1)
	if v294 != v280 {
		v283 = v294
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v263)+96))
	v309 = F_lappend_oid(m, v233, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v312 = v230
	v313 = v309
	v314 = int32(1)
	goto L75
L98:
	;
	v319 = v232 + int32(1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v319 < v320 {
		v230 = v312
		v232 = v319
		v233 = v313
		v237 = v314
		goto L72
	} else {
		goto L99
	}
L99:
	;
	goto L73
L100:
	;
	if v332 != 0 {
		v31 = v328
		v32 = v198
		goto L4
	} else {
		goto L101
	}
L101:
	;
	v342 = v328
	v343 = v198
	goto L6
L102:
	;
	F_list_free(m, v342)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_list_sort(m, v343, int32(502))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v356 = int32(0)
	if v343 == v356 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	m.G0 = v15 - int32(-64)
	return v343
L106:
	;
	goto L105
L107:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	if v366 < int32(2) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v369 = int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v343)+12))
	if v366 != int32(2) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v343)+4)) = v451 + int32(1)
	goto L106
L110:
	;
	v373 = int32(1)
	v374 = v366 - v373
	v379 = int32(0)
	v382 = v379
	v384 = v369
	v385 = v379
	goto L113
L111:
	;
	v427 = v356
	v429 = v369
	goto L112
L112:
	;
	v435 = int32(2)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v370+v429<<(uint(v435)%32))))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v370+v427<<(uint(v435)%32))))
	if v438 == v442 {
		v451 = v427
		goto L109
	} else {
		goto L123
	}
L113:
	;
	v390 = int32(2)
	v392 = v370 + v384<<(uint(v390)%32)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v370+v382<<(uint(v390)%32))))
	if v393 != v397 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v374&v373 == int32(0) {
		v451 = v418
		goto L109
	} else {
		goto L122
	}
L115:
	;
	v400 = v382 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v370+v400<<(uint(int32(2))%32)))) = v393
	v405 = v400
	goto L117
L116:
	;
	v405 = v382
	goto L117
L117:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v370+v405<<(uint(int32(2))%32))))
	if v406 != v410 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v413 = v405 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v370+v413<<(uint(int32(2))%32)))) = v406
	v418 = v413
	goto L120
L119:
	;
	v418 = v405
	goto L120
L120:
	;
	v419 = int32(2)
	v420 = v384 + v419
	v422 = v385 + v419
	if v422 != v374&int32(-2) {
		v382 = v418
		v384 = v420
		v385 = v422
		goto L113
	} else {
		goto L121
	}
L121:
	;
	goto L114
L122:
	;
	v427 = v418
	v429 = v420
	goto L112
L123:
	;
	v445 = v427 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v370+v445<<(uint(int32(2))%32)))) = v438
	v451 = v445
	goto L109
}
func F_heap_update(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v293 int64
	_ = v293
	var v294 int32
	_ = v294
	var v297 int64
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v413 int32
	_ = v413
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v461 int32
	_ = v461
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v555 int32
	_ = v555
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v895 int32
	_ = v895
	var v905 int32
	_ = v905
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v976 int32
	_ = v976
	var v987 int32
	_ = v987
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1054 int32
	_ = v1054
	var v1076 int32
	_ = v1076
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1111 int32
	_ = v1111
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1165 int32
	_ = v1165
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1225 int32
	_ = v1225
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1332 int32
	_ = v1332
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1418 int32
	_ = v1418
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1483 int32
	_ = v1483
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1504 int32
	_ = v1504
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1575 int32
	_ = v1575
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1591 int32
	_ = v1591
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1634 int32
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1648 int32
	_ = v1648
	var v1653 int32
	_ = v1653
	var v1682 int32
	_ = v1682
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1703 int32
	_ = v1703
	var v1747 int32
	_ = v1747
	var v1751 int32
	_ = v1751
	var v1756 int32
	_ = v1756
	var v1761 int32
	_ = v1761
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1870 int32
	_ = v1870
	var v1875 int32
	_ = v1875
	var v1884 int32
	_ = v1884
	var v1893 int32
	_ = v1893
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1919 int32
	_ = v1919
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1948 int32
	_ = v1948
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1989 int32
	_ = v1989
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1996 int32
	_ = v1996
	var v1999 int32
	_ = v1999
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2011 int32
	_ = v2011
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2033 int32
	_ = v2033
	var v2035 int32
	_ = v2035
	var v2037 int32
	_ = v2037
	var v2052 int32
	_ = v2052
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2063 int32
	_ = v2063
	var v2066 int64
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2069 int64
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2093 int64
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2101 int32
	_ = v2101
	var v2103 int32
	_ = v2103
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2180 int32
	_ = v2180
	var v2183 int32
	_ = v2183
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2206 int32
	_ = v2206
	var v2211 int32
	_ = v2211
	var v2220 int32
	_ = v2220
	var v2229 int32
	_ = v2229
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2250 int32
	_ = v2250
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2257 int32
	_ = v2257
	var v2261 int32
	_ = v2261
	var v2265 int32
	_ = v2265
	var v2268 int32
	_ = v2268
	var v2272 int32
	_ = v2272
	var v2277 int32
	_ = v2277
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2334 int32
	_ = v2334
	var v2378 int32
	_ = v2378
	var v2382 int32
	_ = v2382
	var v2388 int32
	_ = v2388
	var v2390 int32
	_ = v2390
	var v2396 int32
	_ = v2396
	var v2400 int32
	_ = v2400
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2421 int32
	_ = v2421
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2436 int32
	_ = v2436
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2445 int32
	_ = v2445
	var v2452 int32
	_ = v2452
	var v2454 int32
	_ = v2454
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2484 int32
	_ = v2484
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2493 int32
	_ = v2493
	var v2500 int32
	_ = v2500
	var v2502 int32
	_ = v2502
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2507 int32
	_ = v2507
	var v2509 int32
	_ = v2509
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2519 int32
	_ = v2519
	var v2521 int32
	_ = v2521
	var v2523 int32
	_ = v2523
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2529 int32
	_ = v2529
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2542 int32
	_ = v2542
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2551 int32
	_ = v2551
	var v2558 int32
	_ = v2558
	var v2560 int32
	_ = v2560
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2565 int32
	_ = v2565
	var v2567 int32
	_ = v2567
	var v2574 int32
	_ = v2574
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2584 int32
	_ = v2584
	var v2587 int32
	_ = v2587
	var v2591 int32
	_ = v2591
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	var v2604 int32
	_ = v2604
	var v2608 int32
	_ = v2608
	var v2610 int32
	_ = v2610
	var v2614 int32
	_ = v2614
	var v2620 int32
	_ = v2620
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2633 int32
	_ = v2633
	var v2639 int32
	_ = v2639
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2662 int32
	_ = v2662
	var v2664 int32
	_ = v2664
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2680 int32
	_ = v2680
	var v2683 int32
	_ = v2683
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2689 int32
	_ = v2689
	var v2693 int32
	_ = v2693
	var v2696 int32
	_ = v2696
	var v2708 int32
	_ = v2708
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2725 int32
	_ = v2725
	var v2726 int32
	_ = v2726
	var v2728 int32
	_ = v2728
	var v2730 int32
	_ = v2730
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2733 int32
	_ = v2733
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2738 int32
	_ = v2738
	var v2741 int32
	_ = v2741
	var v2743 int32
	_ = v2743
	var v2744 int32
	_ = v2744
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2751 int32
	_ = v2751
	var v2754 int32
	_ = v2754
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2765 int32
	_ = v2765
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2772 int32
	_ = v2772
	var v2773 int32
	_ = v2773
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2779 int32
	_ = v2779
	var v2780 int32
	_ = v2780
	var v2781 int32
	_ = v2781
	var v2783 int32
	_ = v2783
	var v2790 int32
	_ = v2790
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2795 int32
	_ = v2795
	var v2797 int32
	_ = v2797
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2803 int32
	_ = v2803
	var v2805 int32
	_ = v2805
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2813 int32
	_ = v2813
	var v2817 int32
	_ = v2817
	var v2819 int32
	_ = v2819
	var v2821 int32
	_ = v2821
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2833 int32
	_ = v2833
	var v2839 int32
	_ = v2839
	var v2841 int32
	_ = v2841
	var v2842 int32
	_ = v2842
	var v2847 int32
	_ = v2847
	var v2848 int32
	_ = v2848
	var v2849 int32
	_ = v2849
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2855 int32
	_ = v2855
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
	var v2863 int32
	_ = v2863
	var v2867 int32
	_ = v2867
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2875 int32
	_ = v2875
	var v2878 int32
	_ = v2878
	var v2879 int32
	_ = v2879
	var v2881 int32
	_ = v2881
	var v2887 int32
	_ = v2887
	var v2892 int32
	_ = v2892
	var v2897 int32
	_ = v2897
	var v2900 int32
	_ = v2900
	var v2901 int32
	_ = v2901
	var v2904 int32
	_ = v2904
	var v2911 int32
	_ = v2911
	var v2913 int32
	_ = v2913
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2924 int32
	_ = v2924
	var v2930 int32
	_ = v2930
	var v2932 int32
	_ = v2932
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2941 int32
	_ = v2941
	var v2945 int32
	_ = v2945
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2956 int32
	_ = v2956
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2961 int32
	_ = v2961
	var v2968 int32
	_ = v2968
	var v2970 int32
	_ = v2970
	var v2971 int32
	_ = v2971
	var v2972 int32
	_ = v2972
	var v2976 int32
	_ = v2976
	var v2978 int32
	_ = v2978
	var v2988 int32
	_ = v2988
	var v2994 int32
	_ = v2994
	var v2996 int32
	_ = v2996
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3004 int32
	_ = v3004
	var v3007 int64
	_ = v3007
	var v3008 int64
	_ = v3008
	var v3013 int32
	_ = v3013
	var v3017 int32
	_ = v3017
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3029 int32
	_ = v3029
	var v3039 int32
	_ = v3039
	var v3041 int32
	_ = v3041
	var v3081 int32
	_ = v3081
	var v3083 int32
	_ = v3083
	var v3086 int32
	_ = v3086
	var v3089 int32
	_ = v3089
	var v3091 int32
	_ = v3091
	var v3142 int32
	_ = v3142
	var v3149 int32
	_ = v3149
	var v3190 int32
	_ = v3190
	var v3194 int32
	_ = v3194
	var v3199 int32
	_ = v3199
	var v3207 int32
	_ = v3207
	var v3217 int32
	_ = v3217
	var v3247 int32
	_ = v3247
	var v3249 int32
	_ = v3249
	var v3251 int32
	_ = v3251
	var v3254 int32
	_ = v3254
	var v3257 int32
	_ = v3257
	var v3259 int32
	_ = v3259
	var v3310 int32
	_ = v3310
	var v3319 int32
	_ = v3319
	var v3358 int32
	_ = v3358
	var v3359 int32
	_ = v3359
	var v3362 int32
	_ = v3362
	var v3367 int32
	_ = v3367
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3375 int32
	_ = v3375
	var v3376 int32
	_ = v3376
	var v3379 int32
	_ = v3379
	var v3384 int32
	_ = v3384
	var v3386 int32
	_ = v3386
	var v3425 int32
	_ = v3425
	var v3426 int32
	_ = v3426
	var v3432 int32
	_ = v3432
	var v3438 int32
	_ = v3438
	var v3439 int32
	_ = v3439
	var v3442 int32
	_ = v3442
	var v3443 int32
	_ = v3443
	var v3447 int32
	_ = v3447
	var v3448 int32
	_ = v3448
	var v3454 int32
	_ = v3454
	var v3463 int32
	_ = v3463
	var v3464 int32
	_ = v3464
	var v3467 int32
	_ = v3467
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3471 int32
	_ = v3471
	var v3473 int32
	_ = v3473
	var v3474 int32
	_ = v3474
	var v3476 int32
	_ = v3476
	var v3477 int32
	_ = v3477
	var v3481 int32
	_ = v3481
	var v3483 int32
	_ = v3483
	var v3487 int32
	_ = v3487
	var v3502 int32
	_ = v3502
	var v3504 int32
	_ = v3504
	var v3505 int32
	_ = v3505
	var v3510 int32
	_ = v3510
	var v3512 int32
	_ = v3512
	var v3518 int32
	_ = v3518
	var v3523 int32
	_ = v3523
	var v3529 int32
	_ = v3529
	var v3531 int32
	_ = v3531
	var v3547 int32
	_ = v3547
	var v3555 int32
	_ = v3555
	var v3561 int32
	_ = v3561
	var v3562 int32
	_ = v3562
	var v3563 int32
	_ = v3563
	var v3565 int32
	_ = v3565
	var v3567 int32
	_ = v3567
	var v3574 int32
	_ = v3574
	var v3575 int32
	_ = v3575
	var v3576 int32
	_ = v3576
	var v3580 int32
	_ = v3580
	var v3582 int32
	_ = v3582
	var v3583 int32
	_ = v3583
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3591 int32
	_ = v3591
	var v3596 int32
	_ = v3596
	var v3597 int32
	_ = v3597
	var v3598 int32
	_ = v3598
	var v3599 int32
	_ = v3599
	var v3601 int32
	_ = v3601
	var v3602 int32
	_ = v3602
	var v3603 int32
	_ = v3603
	var v3607 int32
	_ = v3607
	var v3609 int32
	_ = v3609
	var v3610 int32
	_ = v3610
	var v3614 int32
	_ = v3614
	var v3618 int32
	_ = v3618
	var v3625 int32
	_ = v3625
	var v3626 int32
	_ = v3626
	var v3628 int32
	_ = v3628
	var v3630 int32
	_ = v3630
	var v3636 int32
	_ = v3636
	var v3637 int32
	_ = v3637
	var v3638 int32
	_ = v3638
	var v3640 int32
	_ = v3640
	var v3644 int32
	_ = v3644
	var v3647 int32
	_ = v3647
	var v3651 int32
	_ = v3651
	var v3653 int32
	_ = v3653
	var v3657 int32
	_ = v3657
	var v3659 int32
	_ = v3659
	var v3661 int32
	_ = v3661
	var v3662 int32
	_ = v3662
	var v3667 int64
	_ = v3667
	var v3668 int32
	_ = v3668
	var v3670 int64
	_ = v3670
	var v3675 int32
	_ = v3675
	var v3679 int32
	_ = v3679
	var v3685 int32
	_ = v3685
	var v3687 int32
	_ = v3687
	var v3693 int32
	_ = v3693
	var v3698 int32
	_ = v3698
	var v3702 int32
	_ = v3702
	var v3708 int32
	_ = v3708
	var v3710 int32
	_ = v3710
	var v3716 int32
	_ = v3716
	var v3763 int32
	_ = v3763
	var v3765 int32
	_ = v3765
	var v3769 int32
	_ = v3769
	var v3771 int32
	_ = v3771
	var v3772 int32
	_ = v3772
	var v3774 int32
	_ = v3774
	var v3778 int32
	_ = v3778
	var v3780 int32
	_ = v3780
	var v3784 int32
	_ = v3784
	var v3786 int32
	_ = v3786
	var v3788 int32
	_ = v3788
	var v3792 int32
	_ = v3792
	var v3794 int32
	_ = v3794
	var v3795 int32
	_ = v3795
	var v3797 int32
	_ = v3797
	var v3798 int32
	_ = v3798
	var v3800 int32
	_ = v3800
	var v3803 int32
	_ = v3803
	var v3806 int32
	_ = v3806
	var v3808 int32
	_ = v3808
	var v3809 int32
	_ = v3809
	var v3812 int32
	_ = v3812
	var v3816 int32
	_ = v3816
	var v3817 int32
	_ = v3817
	var v3818 int32
	_ = v3818
	var v3820 int32
	_ = v3820
	var v3821 int32
	_ = v3821
	var v3822 int32
	_ = v3822
	var v3823 int32
	_ = v3823
	var v3825 int32
	_ = v3825
	var v3826 int32
	_ = v3826
	var v3828 int32
	_ = v3828
	var v3830 int32
	_ = v3830
	var v3831 int32
	_ = v3831
	var v3833 int32
	_ = v3833
	var v3836 int32
	_ = v3836
	var v3840 int32
	_ = v3840
	var v3843 int64
	_ = v3843
	var v3852 int32
	_ = v3852
	var v3853 int32
	_ = v3853
	var v3854 int64
	_ = v3854
	var v3864 int32
	_ = v3864
	var v3865 int32
	_ = v3865
	var v3867 int32
	_ = v3867
	var v3870 int32
	_ = v3870
	var v3873 int32
	_ = v3873
	var v3877 int32
	_ = v3877
	var v3883 int32
	_ = v3883
	var v3885 int32
	_ = v3885
	var v3888 int32
	_ = v3888
	var v3933 int32
	_ = v3933
	var v3935 int32
	_ = v3935
	var v3959 int32
	_ = v3959
	var v3961 int32
	_ = v3961
	var v3964 int32
	_ = v3964
	var v3978 int32
	_ = v3978
	var v3980 int32
	_ = v3980
	var v3982 int32
	_ = v3982
	var v3984 int32
	_ = v3984
	var v3986 int32
	_ = v3986
	v11 = int32(0)
	v46 = m.G0
	v48 = v46 - int32(96)
	m.G0 = v48
	*(*int32)(unsafe.Add(mBase, uint32(v48)+52)) = l3
	v51 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+31)) = uint8(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[0]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+72))
	if v64 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_bms_free(m, v3961)
	mBase = m.M
	v3978 = m.ExcPending
	if v3978 != 0 {
		goto L1
	} else {
		goto L818
	}
L4:
	;
	v3933 = v3888
	v3935 = v79
	v3959 = v443
	v3961 = v76
	v3964 = v82
	goto L3
L5:
	;
	if v98 < int32(0) {
		goto L439
	} else {
		goto L440
	}
L6:
	;
	if v2378 < int32(0) {
		goto L435
	} else {
		goto L436
	}
L7:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v2123)))
	v2324 = int32(0)
	v2331 = F_RelationGetBufferForTuple(m, l0, v2323, v98, v2324, v2324, v48+int32(20), v48+int32(24), v2324)
	mBase = m.M
	v2332 = m.ExcPending
	if v2332 != 0 {
		goto L1
	} else {
		goto L434
	}
L8:
	;
	if v67&int32(1) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v67 = int32(1)
	goto L11
L10:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+76)))
	v67 = v66
	goto L11
L11:
	;
	goto L8
L12:
	;
	v73 = F_RelationGetIndexAttrBitmap(m, l0, int32(3))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L1
	} else {
		goto L430
	}
L15:
	;
	v76 = F_RelationGetIndexAttrBitmap(m, l0, int32(4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v79 = F_RelationGetIndexAttrBitmap(m, l0, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v82 = F_RelationGetIndexAttrBitmap(m, l0, int32(2))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v85 = F_bms_add_members(m, int32(0), v73)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v87 = F_bms_add_members(m, v85, v76)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v89 = F_bms_add_members(m, v87, v79)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v91 = F_bms_add_members(m, v89, v82)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v97 = v93 | v94<<(uint(int32(16))%32)
	v98 = F_ReadBuffer(m, l0, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+10)))
	if v118&int32(4) != 0 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if v98 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v103+(v98^int32(-1))<<(uint(int32(2))%32))))
	v117 = v109
	goto L23
L26:
	;
	goto L27
L27:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v117 = v111 + v98<<(uint(int32(13))%32) + int32(-8192)
	goto L23
L28:
	;
	F_visibilitymap_pin(m, l0, v97, v48+int32(24))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_LockBufferInternal(m, v98, int32(3))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v129 = v117 + int32(20)
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v133 = v129 + v130<<(uint(int32(2))%32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v134&int32(_a_F_heap_update_0) != int32(_a_F_heap_update_1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_UnlockReleaseBuffer(m, v98)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+44)) = v153
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+48)) = v117 + v155&int32(_a_F_heap_update_2)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = int32(base.Ui32(v160) >> (uint(int32(17)) % 32))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+36)) = v164
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+40)) = uint16(v166)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v153
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v91 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v141 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_ReleaseBuffer(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l7)+4)) = uint16(v144)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v146
	*(*int64)(unsafe.Add(mBase, uint32(l7)+8)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = int32(0)
	v3933 = int32(4)
	v3935 = v76
	v3959 = v82
	v3961 = v73
	v3964 = v79
	goto L3
L40:
	;
	goto L39
L41:
	;
	if int32(0) <= v226 {
		goto L52
	} else {
		goto L53
	}
L42:
	;
	v226 = base.I32_ctz(v212) | v213<<(uint(int32(5))%32)
	goto L41
L43:
	;
	v226 = int32(-2)
	goto L41
L44:
	;
	v177 = int32(0)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v180 <= v177 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v183 = v91 + int32(8)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v190 = v187 & int32(-1)
	if v190 != 0 {
		v212 = v190
		v213 = v177
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v191 = int32(1)
	if v191 == v180 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v195 = v191
	goto L48
L48:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v183+v195<<(uint(int32(2))%32))))
	if v202 != 0 {
		v212 = v202
		v213 = v195
		goto L42
	} else {
		goto L50
	}
L49:
	;
	goto L43
L50:
	;
	v204 = v195 + int32(1)
	if v204 != v180 {
		v195 = v204
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v232 = v226
	v251 = v11
	v258 = v11
	goto L55
L53:
	;
	v436 = v11
	v443 = v11
	goto L54
L54:
	;
	v461 = int32(0)
	if base.B2i32(v443 == v461)|base.B2i32(v79 == v461) != 0 {
		v506 = v461
		goto L93
	} else {
		goto L94
	}
L55:
	;
	v279 = v232<<(uint(int32(16))%32) - int32(_a_F_heap_update_3)
	if v279 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v436 = v353
	v443 = v355
	goto L54
L57:
	;
	if v91 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L58:
	;
	if v332&int32(1)|base.B2i32(v285 < int32(0)) != 0 {
		v353 = v251
		v355 = v258
		goto L57
	} else {
		goto L75
	}
L59:
	;
	v330 = F_bms_add_member(m, v258, v232)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L74
	}
L60:
	;
	v285 = v279 >> (uint(int32(16)) % 32)
	if base.B2i32(v279 != int32(-393216))&base.B2i32(v285 < int32(0)) != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v293 = F_heap_getattr_1(m, v48+int32(32), v285, v169, v48+int32(80))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v297 = F_heap_getattr_1(m, l2, v285, v169, v48+int32(72))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+72)))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+80)))
	if (v299|v300)&int32(1) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if v285 <= int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	if v300&int32(255) == v299 {
		v332 = v300
		goto L58
	} else {
		goto L73
	}
L67:
	;
	if base.I32_wrap_i64(v293) == base.I32_wrap_i64(v297) {
		v332 = int32(0)
		goto L58
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v314 = v169 + int32(20) + v285<<(uint(int32(3))%32)
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+4)))
	v316 = int32(*(*int16)(unsafe.Add(mBase, uint32(v314)+2)))
	v317 = F_datumIsEqual(m, v293, v297, v315, v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	goto L59
L71:
	;
	if v317 == int32(0) {
		goto L59
	} else {
		goto L72
	}
L72:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+80)))
	v332 = v321
	goto L58
L73:
	;
	goto L59
L74:
	;
	v353 = v251
	v355 = v330
	goto L57
L75:
	;
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169+v285<<(uint(int32(3))%32))+22)))
	if v341 != int32(_a_F_heap_update_4) {
		v353 = v251
		v355 = v258
		goto L57
	} else {
		goto L76
	}
L76:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v293)))))
	if v345 != int32(1) {
		v353 = v251
		v355 = v258
		goto L57
	} else {
		goto L77
	}
L77:
	;
	v348 = F_bms_is_member(m, v232, v82)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v353 = v348 | v251
	v355 = v258
	goto L57
L79:
	;
	if int32(0) <= v413 {
		v232 = v413
		v251 = v353
		v258 = v355
		goto L55
	} else {
		goto L90
	}
L80:
	;
	v413 = base.I32_ctz(v399) | v400<<(uint(int32(5))%32)
	goto L79
L81:
	;
	v413 = int32(-2)
	goto L79
L82:
	;
	v364 = v232 + int32(1)
	v366 = int32(base.Ui32(v364) >> (uint(int32(5)) % 32))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v367 <= v366 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v370 = v91 + int32(8)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v370+v366<<(uint(int32(2))%32))))
	v377 = v374 & (int32(-1) << (uint(v364) % 32))
	if v377 != 0 {
		v399 = v377
		v400 = v366
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v379 = v366 + int32(1)
	if v379 == v367 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v382 = v379
	goto L86
L86:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v370+v382<<(uint(int32(2))%32))))
	if v389 != 0 {
		v399 = v389
		v400 = v382
		goto L80
	} else {
		goto L88
	}
L87:
	;
	goto L81
L88:
	;
	v391 = v382 + int32(1)
	if v391 != v367 {
		v382 = v391
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	goto L56
L91:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v521 = F_HeapTupleSatisfiesUpdate(m, v48+int32(32), v520, v98)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L110
	}
L92:
	;
	if v506 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L93:
	;
	goto L92
L94:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v471 < v472 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v474 = v471
	goto L97
L96:
	;
	v474 = v472
	goto L97
L97:
	;
	if v474 <= int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v477 = int32(1)
	goto L100
L99:
	;
	v477 = v474
	goto L100
L100:
	;
	v478 = int32(8)
	v483 = int32(0)
	goto L101
L101:
	;
	v490 = v483 << (uint(int32(2)) % 32)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v79+v478+v490)))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v443+v478+v490)))
	v495 = v492 & v494
	v497 = base.B2i32(v495 != int32(0))
	if v495 != 0 {
		v506 = v497
		goto L93
	} else {
		goto L103
	}
L102:
	;
	v506 = v497
	goto L93
L103:
	;
	v499 = v483 + int32(1)
	if v499 != v477 {
		v483 = v499
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = int32(2)
	F_MultiXactIdSetOldestMember(m)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = int32(3)
	v517 = int32(5)
	goto L91
L108:
	;
	v517 = int32(4)
	goto L91
L109:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+4))
	v1426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1424)+20)))
	v1427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1424)+18)))
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	F_compute_new_xmax_infomask(m, v1425, v1426, v1427, v51, v1428, int32(1), v48+int32(12), v48+int32(10), v48+int32(8))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L271
	}
L110:
	;
	if v521 != int32(1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v526 = l4 & int32(1)
	v528 = v48 + int32(36)
	v533 = v521
	v555 = int32(0)
	goto L114
L112:
	;
	goto L113
L113:
	;
	F_UnlockReleaseBuffer(m, v98)
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L1
	} else {
		goto L266
	}
L114:
	;
	v577 = int32(1)
	if base.B2i32(l6 == int32(0))|base.B2i32(v533 != int32(5)) != 0 {
		v1032 = v533
		v1035 = v577
		v1054 = v555
		goto L119
	} else {
		goto L120
	}
L115:
	;
	goto L113
L116:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v1357 = F_HeapTupleSatisfiesUpdate(m, v48+int32(32), v1356, v98)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L264
	}
L117:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v1294 != 0 {
		goto L109
	} else {
		goto L259
	}
L118:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v1134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1133)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(l7)+4)) = uint16(v1134)
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v1136
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+4))
	v1139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1133)+20)))
	if v1139&int32(_a_F_heap_update_5) != int32(_a_F_heap_update_6) {
		goto L230
	} else {
		goto L231
	}
L119:
	;
	v1076 = int32(0)
	if base.B2i32(l5 == v1076)|v1032 == v1076 {
		goto L223
	} else {
		goto L224
	}
L120:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)+4))
	v585 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583)+20)))
	if v585&int32(_a_F_heap_update_6) != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v1012 = v976 + int32(12)
	v1013 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v528)+2)))
	v1014 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v528))))
	v1015 = int32(16)
	v1018 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1012)+2)))
	v1019 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1012))))
	if v1013|v1014<<(uint(v1015)%32) == v1018|v1019<<(uint(v1015)%32) {
		goto L216
	} else {
		goto L217
	}
L122:
	;
	v588 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+66)) = uint8(v588)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v593 = F_DoesMultiXactIdConflict(m, v584, v585, v590, v48+int32(66))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	v773 = int32(0)
	if base.Ui32(v584) < base.Ui32(int32(3)) {
		goto L156
	} else {
		goto L157
	}
L125:
	;
	if v637&int32(128)|base.B2i32(v637&int32(_a_F_heap_update_7) == int32(64)) != 0 {
		goto L141
	} else {
		goto L142
	}
L126:
	;
	if v593 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v598 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v597)+20)))
	v637 = v598
	v638 = v597
	v639 = v555
	v640 = int32(0)
	goto L125
L128:
	;
	goto L129
L129:
	;
	F_UnlockBuffer(m, v98)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+66)))
	if (v602|v555)&int32(1) != 0 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v617 = int32(0)
	v622 = F_Do_MultiXactIdWait(m, v584, v517, v585, v617, l0, v528, int32(1), v48+int32(72), v617)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L136
	}
L132:
	;
	v616 = v602 ^ int32(1) | v555
	goto L131
L133:
	;
	goto L134
L134:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v609*int32(12))+uint32(_c_F_heap_update[3])))
	F_LockTuple(m, l0, v528, v612)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v616 = int32(1)
	goto L131
L136:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v48)+72))
	F_LockBufferInternal(m, v98, int32(3))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v629 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v628)+20)))
	if (v629^v585)&int32(_a_F_heap_update_8) != 0 {
		v1332 = v616
		goto L116
	} else {
		goto L138
	}
L138:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v628)+4))
	if v633 != v584 {
		v1332 = v616
		goto L116
	} else {
		goto L139
	}
L139:
	;
	v637 = v629
	v638 = v628
	v639 = v616
	v640 = base.B2i32(v624 != int32(0))
	goto L125
L140:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v976 = v772
	v987 = v639
	goto L121
L141:
	;
	v1032 = int32(0)
	v1035 = v593 ^ int32(1) | v640
	v1054 = v639
	goto L119
L142:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v638)+4))
	v652 = F_GetMultiXactIdMembers(m, v648, v48+int32(80), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	if v652 <= int32(0) {
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v48)+80))
	v659 = int32(0)
	goto L146
L145:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	F_pfree(m, v657)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L151
	}
L146:
	;
	v705 = v657 + v659<<(uint(int32(3))%32)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v706) {
		goto L145
	} else {
		goto L148
	}
L147:
	;
	F_pfree(m, v657)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	v710 = v659 + int32(1)
	if v710 != v652 {
		v659 = v710
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	goto L141
L151:
	;
	if v714 == int32(0) {
		goto L141
	} else {
		goto L152
	}
L152:
	;
	v719 = F_TransactionIdDidAbort(m, v714)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	if v719 == int32(0) {
		goto L140
	} else {
		goto L154
	}
L154:
	;
	goto L141
L155:
	;
	if v905|(v506^int32(-1))&base.B2i32(v585&int32(80) == int32(16)) != 0 {
		v1032 = v773
		v1035 = v577
		v1054 = v555
		goto L119
	} else {
		goto L195
	}
L156:
	;
	v905 = int32(0)
	goto L155
L157:
	;
	goto L158
L158:
	;
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[4]))
	if v785 == v584 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v905 = int32(1)
	goto L155
L160:
	;
	goto L161
L161:
	;
	v789 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[5]))
	if v789 <= int32(0) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v905 = v895
	goto L155
L163:
	;
	v793 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[0]))
	if v793 == int32(0) {
		v895 = v773
		goto L162
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v863 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[6]))
	v865 = int32(0)
	v868 = v789 - int32(1)
	goto L185
L166:
	;
	v798 = v793
	goto L167
L167:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v798)+20))
	if v804 == int32(4) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v895 = int32(0)
	goto L162
L169:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v798)+80))
	if v858 != 0 {
		v798 = v858
		goto L167
	} else {
		goto L184
	}
L170:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v798)))
	if v807 == int32(0) {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v810 = int32(1)
	if v584 == v807 {
		v895 = v810
		goto L162
	} else {
		goto L172
	}
L172:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v798)+52))
	v814 = v812 - int32(1)
	if v814 < int32(0) {
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v798)+48))
	v820 = int32(0)
	v823 = v814
	goto L174
L174:
	;
	v828 = int32(2)
	v829 = base.I32_div_s(v823-v820, v828)
	v830 = v829 + v820
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v817+v830<<(uint(v828)%32))))
	if v834 == v584 {
		v895 = v810
		goto L162
	} else {
		goto L176
	}
L175:
	;
	goto L169
L176:
	;
	v843 = base.B2i32(v834-v584 < int32(0)) | base.B2i32(base.Ui32(v834) < base.Ui32(int32(3)))
	if v843 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v844 = v830 + int32(1)
	goto L179
L178:
	;
	v844 = v820
	goto L179
L179:
	;
	if v843 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v847 = v823
	goto L182
L181:
	;
	v847 = v830 - int32(1)
	goto L182
L182:
	;
	if v844 <= v847 {
		v820 = v844
		v823 = v847
		goto L174
	} else {
		goto L183
	}
L183:
	;
	goto L175
L184:
	;
	goto L168
L185:
	;
	v873 = int32(2)
	v874 = base.I32_div_s(v868-v865, v873)
	v875 = v874 + v865
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v863+v875<<(uint(v873)%32))))
	v880 = base.B2i32(v879 == v584)
	if v879 == v584 {
		v895 = v880
		goto L162
	} else {
		goto L187
	}
L186:
	;
	v895 = v880
	goto L162
L187:
	;
	v883 = base.B2i32(base.Ui32(v879) < base.Ui32(v584))
	if base.Ui32(v879) < base.Ui32(v584) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v884 = v875 + int32(1)
	goto L190
L189:
	;
	v884 = v865
	goto L190
L190:
	;
	if base.Ui32(v879) < base.Ui32(v584) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v887 = v868
	goto L193
L192:
	;
	v887 = v875 - int32(1)
	goto L193
L193:
	;
	if v884 <= v887 {
		v865 = v884
		v868 = v887
		goto L185
	} else {
		goto L194
	}
L194:
	;
	goto L186
L195:
	;
	F_UnlockBuffer(m, v98)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	if v555&int32(1) == int32(0) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v918*int32(12))+uint32(_c_F_heap_update[3])))
	F_LockTuple(m, l0, v528, v921)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L1
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v924 = int32(1)
	F_XactLockTableWait(m, v584, l0, v528, v924)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L201
	}
L200:
	;
	goto L199
L201:
	;
	F_LockBufferInternal(m, v98, int32(3))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v932 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v931)+20)))
	if (v932^v585)&int32(_a_F_heap_update_8) != 0 {
		v1332 = v924
		goto L116
	} else {
		goto L203
	}
L203:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v931)+4))
	if v584 != v936 {
		v1332 = v924
		goto L116
	} else {
		goto L204
	}
L204:
	;
	if v932&int32(3072) != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v958 = int32(0)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v960)+21)))
	if v961&int32(8) != 0 {
		v1032 = v958
		v1035 = v958
		v1054 = v924
		goto L119
	} else {
		goto L213
	}
L206:
	;
	if v932&int32(128)|base.B2i32(v932&int32(_a_F_heap_update_7) == int32(64)) != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	F_HeapTupleSetHintBits(m, v931, v98, int32(2048), int32(0))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L212
	}
L208:
	;
	v947 = F_TransactionIdDidCommit(m, v584)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	if v947 == int32(0) {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	F_HeapTupleSetHintBits(m, v931, v98, int32(1024), v584)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	goto L205
L212:
	;
	goto L205
L213:
	;
	v976 = v960
	v987 = v924
	goto L121
L214:
	;
	if v1029 != 0 {
		goto L220
	} else {
		goto L221
	}
L215:
	;
	goto L214
L216:
	;
	v1025 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v528)+4)))
	v1026 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1012)+4)))
	if v1025 == v1026 {
		v1029 = int32(1)
		goto L215
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v1029 = int32(0)
	goto L215
L219:
	;
	goto L218
L220:
	;
	v1030 = int32(4)
	goto L222
L221:
	;
	v1030 = int32(3)
	goto L222
L222:
	;
	v1089 = v1030
	v1111 = v987
	goto L118
L223:
	;
	v1083 = F_HeapTupleSatisfiesVisibility(m, v48+int32(32), l5, v98)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	if v1032 == int32(0) {
		goto L117
	} else {
		goto L228
	}
L226:
	;
	if v1083 != 0 {
		goto L117
	} else {
		goto L227
	}
L227:
	;
	v1089 = int32(3)
	v1111 = v1054
	goto L118
L228:
	;
	v1089 = v1032
	v1111 = v1054
	goto L118
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7)+8)) = v1225
	if v1089 == int32(2) {
		goto L242
	} else {
		goto L243
	}
L230:
	;
	v1225 = v1138
	goto L229
L231:
	;
	goto L232
L232:
	;
	v1144 = int32(0)
	v1148 = F_GetMultiXactIdMembers(m, v1138, v48+int32(80), v1144)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	if v1148 <= int32(0) {
		v1225 = v1144
		goto L229
	} else {
		goto L234
	}
L234:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v48)+80))
	v1165 = v1144
	goto L237
L235:
	;
	F_pfree(m, v1152)
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L1
	} else {
		goto L241
	}
L236:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1200)))
	v1210 = v1208
	goto L235
L237:
	;
	v1200 = v1152 + v1165<<(uint(int32(3))%32)
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v1201) {
		goto L236
	} else {
		goto L239
	}
L238:
	;
	v1210 = int32(0)
	goto L235
L239:
	;
	v1205 = v1165 + int32(1)
	if v1205 != v1148 {
		v1165 = v1205
		goto L237
	} else {
		goto L240
	}
L240:
	;
	goto L238
L241:
	;
	v1225 = v1210
	goto L229
L242:
	;
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+8))
	v1264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1261)+20)))
	if v1264&int32(32) != 0 {
		goto L246
	} else {
		goto L247
	}
L243:
	;
	v1275 = int32(-1)
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7)+12)) = v1275
	F_UnlockReleaseBuffer(m, v98)
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L1
	} else {
		goto L249
	}
L245:
	;
	v1275 = v1273
	goto L244
L246:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[7]))
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1268+v1263<<(uint(int32(3))%32))+4))
	v1273 = v1272
	goto L248
L247:
	;
	v1273 = v1263
	goto L248
L248:
	;
	goto L245
L249:
	;
	if v1111&int32(1) != 0 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1281*int32(12))+uint32(_c_F_heap_update[3])))
	F_UnlockTuple(m, l0, v528, v1284)
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L1
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v1287 != 0 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	goto L252
L254:
	;
	F_ReleaseBuffer(m, v1287)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L1
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = int32(0)
	F_bms_free(m, v73)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L1
	} else {
		goto L258
	}
L257:
	;
	goto L256
L258:
	;
	v3888 = v1089
	goto L4
L259:
	;
	v1295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+10)))
	if v1295&int32(4) == int32(0) {
		goto L109
	} else {
		goto L260
	}
L260:
	;
	F_UnlockBuffer(m, v98)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	F_visibilitymap_pin(m, l0, v97, v48+int32(24))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	F_LockBufferInternal(m, v98, int32(3))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	v1332 = v1054
	goto L116
L264:
	;
	if v1357 != int32(1) {
		v533 = v1357
		v555 = v1332
		goto L114
	} else {
		goto L265
	}
L265:
	;
	goto L115
L266:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	F_errmsg(m, int32(_a_F_heap_update_9), int32(0))
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(_a_F_heap_update_10), int32(3511), int32(_a_F_heap_update_11))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L271:
	;
	v1439 = int32(_a_F_heap_update_12)
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v1441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1440)+20)))
	if v1441&int32(2048) != 0 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1796 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1795)+20)))
	v1798 = v1796 & int32(15)
	*(*uint16)(unsafe.Add(mBase, uint32(v1795)+20)) = uint16(v1798)
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1801 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1800)+18)))
	v1803 = v1801 & int32(_a_F_heap_update_13)
	*(*uint16)(unsafe.Add(mBase, uint32(v1800)+18)) = uint16(v1803)
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1805))) = v51
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v1807)+8)) = v1808
	v1810 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1807)+20)))
	v1812 = v1810 & int32(_a_F_heap_update_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v1807)+20)) = uint16(v1812)
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1815 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1814)+20)))
	v1816 = v1815 | v1761
	*(*uint16)(unsafe.Add(mBase, uint32(v1814)+20)) = uint16(v1816)
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1819 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1818)+18)))
	v1820 = v1819 | v1751
	*(*uint16)(unsafe.Add(mBase, uint32(v1818)+18)) = uint16(v1820)
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1822)+4)) = v1756
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	F_HeapTupleHeaderAdjustCmax(m, v1824, v48+int32(52), v48+int32(19))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L1
	} else {
		goto L327
	}
L273:
	;
	v1444 = int32(0)
	v1751 = v1444
	v1756 = v1444
	v1761 = v1439
	goto L272
L274:
	;
	goto L275
L275:
	;
	v1446 = int32(0)
	if v1035&base.B2i32(v1441&int32(_a_F_heap_update_8) != int32(_a_F_heap_update_15)) == v1446 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1751 = int32(0)
	v1756 = v1446
	v1761 = v1439
	goto L272
L277:
	;
	goto L278
L278:
	;
	v1455 = int32(0)
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+4))
	if v1456 == v1455 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1751 = v1455
	v1756 = int32(0)
	v1761 = v1439
	goto L272
L280:
	;
	goto L281
L281:
	;
	if v1441&int32(_a_F_heap_update_6) == int32(0) {
		v1751 = v1455
		v1756 = v1456
		v1761 = int32(_a_F_heap_update_16)
		goto L272
	} else {
		goto L282
	}
L282:
	;
	v1469 = F_GetMultiXactIdMembers(m, v1456, v48+int32(80), int32(0))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L1
	} else {
		goto L284
	}
L283:
	;
	v1751 = v1703
	v1756 = v1456
	v1761 = v1747 | int32(_a_F_heap_update_17)
	goto L272
L284:
	;
	if v1469 <= int32(0) {
		v1703 = v1455
		v1747 = int32(_a_F_heap_update_18)
		goto L283
	} else {
		goto L285
	}
L285:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v48)+80))
	if v1469 == int32(1) {
		goto L288
	} else {
		goto L289
	}
L286:
	;
	F_pfree(m, v1473)
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L1
	} else {
		goto L315
	}
L287:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1473+v1585<<(uint(int32(3))%32))+4))
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1622<<(uint(int32(2))%32))+uint32(_c_F_heap_update[8])))
	if base.Ui32(v1586) < base.Ui32(v1625) {
		goto L309
	} else {
		goto L310
	}
L288:
	;
	v1476 = int32(0)
	v1575 = v1455
	v1585 = v1476
	v1586 = v1476
	v1591 = v1476
	goto L287
L289:
	;
	goto L290
L290:
	;
	v1483 = int32(0)
	v1488 = v1455
	v1491 = v1483
	v1498 = v1483
	v1499 = v1483
	v1504 = v1483
	goto L291
L291:
	;
	v1534 = v1473 + v1498<<(uint(int32(3))%32)
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1534)+4))
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1535<<(uint(int32(2))%32))+uint32(_c_F_heap_update[8])))
	if base.Ui32(v1499) < base.Ui32(v1538) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	if v1469&int32(1) == int32(0) {
		v1637 = v1564
		v1648 = v1566
		v1653 = v1565
		goto L286
	} else {
		goto L308
	}
L293:
	;
	v1540 = v1538
	goto L295
L294:
	;
	v1540 = v1499
	goto L295
L295:
	;
	switch v1535 - int32(3) {
	case 0:
		goto L299
	case 1:
		v1547 = v1488
		goto L297
	case 2:
		goto L298
	default:
		v1549 = v1488
		v1550 = v1504
		goto L296
	}
L296:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1534)+12))
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1551<<(uint(int32(2))%32))+uint32(_c_F_heap_update[8])))
	switch v1551 - int32(3) {
	case 0:
		goto L303
	case 1:
		v1562 = v1549
		goto L301
	case 2:
		goto L302
	default:
		v1564 = v1549
		v1565 = v1550
		goto L300
	}
L297:
	;
	v1549 = v1547
	v1550 = int32(1)
	goto L296
L298:
	;
	v1547 = v1488 | int32(_a_F_heap_update_17)
	goto L297
L299:
	;
	v1549 = v1488 | int32(_a_F_heap_update_17)
	v1550 = v1504
	goto L296
L300:
	;
	if base.Ui32(v1540) < base.Ui32(v1554) {
		goto L304
	} else {
		goto L305
	}
L301:
	;
	v1564 = v1562
	v1565 = int32(1)
	goto L300
L302:
	;
	v1562 = v1549 | int32(_a_F_heap_update_17)
	goto L301
L303:
	;
	v1564 = v1549 | int32(_a_F_heap_update_17)
	v1565 = v1550
	goto L300
L304:
	;
	v1566 = v1554
	goto L306
L305:
	;
	v1566 = v1540
	goto L306
L306:
	;
	v1567 = int32(2)
	v1568 = v1498 + v1567
	v1570 = v1491 + v1567
	if v1570 != v1469&int32(2147483646) {
		v1488 = v1564
		v1491 = v1570
		v1498 = v1568
		v1499 = v1566
		v1504 = v1565
		goto L291
	} else {
		goto L307
	}
L307:
	;
	goto L292
L308:
	;
	v1575 = v1564
	v1585 = v1568
	v1586 = v1566
	v1591 = v1565
	goto L287
L309:
	;
	v1627 = v1625
	goto L311
L310:
	;
	v1627 = v1586
	goto L311
L311:
	;
	switch v1622 - int32(3) {
	case 0:
		goto L314
	case 1:
		v1634 = v1575
		goto L312
	case 2:
		goto L313
	default:
		v1637 = v1575
		v1648 = v1627
		v1653 = v1591
		goto L286
	}
L312:
	;
	v1637 = v1634
	v1648 = v1627
	v1653 = int32(1)
	goto L286
L313:
	;
	v1634 = v1575 | int32(_a_F_heap_update_17)
	goto L312
L314:
	;
	v1637 = v1575 | int32(_a_F_heap_update_17)
	v1648 = v1627
	v1653 = v1591
	goto L286
L315:
	;
	if v1648&int32(-2) == int32(2) {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	if v1653&int32(1) != 0 {
		v1703 = v1637
		v1747 = int32(_a_F_heap_update_19)
		goto L283
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	if v1648 != 0 {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	v1703 = v1637
	v1747 = int32(_a_F_heap_update_20)
	goto L283
L320:
	;
	v1694 = int32(_a_F_heap_update_6)
	goto L322
L321:
	;
	v1694 = int32(_a_F_heap_update_21)
	goto L322
L322:
	;
	if v1648 == int32(1) {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1697 = int32(_a_F_heap_update_7)
	goto L325
L324:
	;
	v1697 = v1694
	goto L325
L325:
	;
	if v1653&int32(1) != 0 {
		v1703 = v1637
		v1747 = v1697
		goto L283
	} else {
		goto L326
	}
L326:
	;
	v1703 = v1637
	v1747 = v1697 | int32(128)
	goto L283
L327:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1831)+119)))
	switch v1832 - int32(109) {
	case 0, 5:
		goto L329
	default:
		v1847 = int32(0)
		goto L328
	}
L328:
	;
	v1851 = int32(4)
	v1852 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+14)))
	v1853 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+12)))
	v1854 = v1852 - v1853
	if v1854 <= v1851 {
		goto L333
	} else {
		goto L334
	}
L329:
	;
	v1835 = int32(1)
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v1837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1836)+20)))
	if v1837&int32(4) != 0 {
		v1847 = v1835
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1840)+20)))
	if v1841&int32(4) != 0 {
		v1847 = v1835
		goto L328
	} else {
		goto L331
	}
L331:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1847 = base.B2i32(base.Ui32(int32(2032)) < base.Ui32(v1844))
	goto L328
L332:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1919 = (v1915 + int32(7)) & int32(-8)
	if v1847 == int32(0) {
		goto L351
	} else {
		goto L352
	}
L333:
	;
	v1857 = v1851
	goto L335
L334:
	;
	v1857 = v1854
	goto L335
L335:
	;
	v1859 = v1857 - int32(4)
	if v1859 == int32(0) {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v1914 = int32(0)
	goto L332
L337:
	;
	goto L338
L338:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1853) {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	v1914 = v1859
	goto L332
L340:
	;
	v1870 = int32(base.Ui32(v1853+int32(_a_F_heap_update_22)) >> (uint(int32(2)) % 32))
	goto L342
L341:
	;
	v1870 = int32(0)
	goto L342
L342:
	;
	if base.Ui32(v1870&int32(_a_F_heap_update_4)) < base.Ui32(int32(291)) {
		goto L339
	} else {
		goto L343
	}
L343:
	;
	v1875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+10)))
	if v1875&int32(1) == int32(0) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1914 = int32(0)
	goto L332
L345:
	;
	goto L346
L346:
	;
	v1884 = int32(1)
	goto L347
L347:
	;
	v1893 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117+int32(20)+v1884&int32(_a_F_heap_update_4)<<(uint(int32(2))%32))+1)))
	if v1893&int32(384) == int32(0) {
		goto L339
	} else {
		goto L349
	}
L348:
	;
	v1914 = int32(0)
	goto L332
L349:
	;
	v1899 = v1884 + int32(1)
	v1900 = int32(_a_F_heap_update_4)
	if base.Ui32(v1899&v1900) <= base.Ui32(v1870&v1900) {
		v1884 = v1899
		goto L347
	} else {
		goto L350
	}
L350:
	;
	goto L348
L351:
	;
	if base.Ui32(v1919) <= base.Ui32(v1914) {
		v2334 = l2
		v2378 = v98
		goto L6
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	v1924 = int32(0)
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+4))
	v1927 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1925)+20)))
	v1928 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1925)+18)))
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	F_compute_new_xmax_infomask(m, v1926, v1927, v1928, v51, v1929, v1924, v48+int32(72), v48+int32(66), v48+int32(62))
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L1
	} else {
		goto L355
	}
L354:
	;
	goto L353
L355:
	;
	v1939 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+10)))
	v1941 = v1939 & int32(4)
	if v1941 != 0 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	F_LockBufferInternal(m, v1942, int32(3))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L1
	} else {
		goto L359
	}
L357:
	;
	goto L358
L358:
	;
	v1946 = int32(_a_F_heap_update_23)
	v1948 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_update[9])) = v1948 + int32(1)
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v1953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1952)+20)))
	v1955 = v1953 & int32(_a_F_heap_update_24)
	*(*uint16)(unsafe.Add(mBase, uint32(v1952)+20)) = uint16(v1955)
	v1957 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1952)+18)))
	v1959 = v1957 & int32(_a_F_heap_update_25)
	*(*uint16)(unsafe.Add(mBase, uint32(v1952)+18)) = uint16(v1959)
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v1962 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1961)+18)))
	v1964 = v1962 & int32(_a_F_heap_update_26)
	*(*uint16)(unsafe.Add(mBase, uint32(v1961)+18)) = uint16(v1964)
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v48)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1966)+4)) = v1967
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v1970 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1969)+20)))
	v1971 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+66)))
	v1972 = v1970 | v1971
	*(*uint16)(unsafe.Add(mBase, uint32(v1969)+20)) = uint16(v1972)
	v1974 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1969)+18)))
	v1975 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+62)))
	v1976 = v1974 | v1975
	*(*uint16)(unsafe.Add(mBase, uint32(v1969)+18)) = uint16(v1976)
	v1978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+19)))
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v1979)+8)) = v1980
	v1982 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1979)+20)))
	v1989 = v1982&int32(_a_F_heap_update_14) | v1978<<(uint(int32(5))%32)&int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1979)+20)) = uint16(v1989)
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v1992 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v528)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1991)+16)) = uint16(v1992)
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	*(*int32)(unsafe.Add(mBase, uint32(v1991)+12)) = v1994
	v1996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+10)))
	if v1996&int32(4) != 0 {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	goto L358
L360:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	v2001 = F_visibilitymap_clear(m, v97, v1999, int32(2))
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L1
	} else {
		goto L363
	}
L361:
	;
	v2003 = v1924
	goto L362
L362:
	;
	F_MarkBufferDirty(m, v98)
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L1
	} else {
		goto L364
	}
L363:
	;
	v2003 = v2001
	goto L362
L364:
	;
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2006)+118)))
	if v2007 != int32(112) {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v2101 = int32(_a_F_heap_update_23)
	v2103 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_update[9])) = v2103 - int32(1)
	if v1941 != 0 {
		goto L384
	} else {
		goto L385
	}
L366:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[10]))
	if v2011 <= int32(0) {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2014 != 0 {
		goto L365
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
		goto L1
	} else {
		goto L372
	}
L370:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2015 != 0 {
		goto L365
	} else {
		goto L371
	}
L371:
	;
	goto L369
L372:
	;
	F_XLogRegisterBuffer(m, int32(0), v98, int32(8))
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L1
	} else {
		goto L373
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+80)) = v1967
	v2023 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+40)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+84)) = uint16(v2023)
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v2026 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2025)+18)))
	v2027 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2025)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+87)) = uint8(v2003)
	v2033 = int32(1)
	v2035 = int32(8)
	v2037 = int32(4)
	v2052 = int32(base.Ui32(v2026)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v2027)>>(uint(v2033)%32))&v2035 | (int32(base.Ui32(v2027)>>(uint(v2037)%32))&v2037 | (int32(base.Ui32(v2027)>>(uint(int32(12))%32))&v2033 | int32(base.Ui32(v2027)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+86)) = uint8(v2052)
	F_XLogRegisterData(m, v48+int32(80), v2035)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L1
	} else {
		goto L374
	}
L374:
	;
	if v2003 != 0 {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	F_XLogRegisterBuffer(m, int32(1), v2060, int32(0))
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L1
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	v2093 = F_XLogInsert(m, int32(10), int32(96))
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L1
	} else {
		goto L383
	}
L378:
	;
	v2066 = F_XLogInsert(m, int32(10), int32(96))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v2069 = base.I64_rotl(v2066, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = v2069
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v2071 < int32(0) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v2075 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v2075+(v2071^int32(-1))<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2081))) = v2069
	goto L365
L381:
	;
	goto L382
L382:
	;
	v2084 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v2084+v2071<<(uint(int32(13))%32))+uint32(_c_F_heap_update[11]))) = v2069
	goto L365
L383:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = base.I64_rotl(v2093, int64(32))
	goto L365
L384:
	;
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	F_UnlockBuffer(m, v2107)
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L1
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	F_UnlockBuffer(m, v98)
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L1
	} else {
		goto L388
	}
L387:
	;
	goto L386
L388:
	;
	if v1847 != 0 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v2116 = F_heap_toast_insert_or_update(m, l0, l2, v48+int32(32), v526<<(uint(int32(3))%32))
	mBase = m.M
	v2117 = m.ExcPending
	if v2117 != 0 {
		goto L1
	} else {
		goto L392
	}
L390:
	;
	v2123 = l2
	v2124 = v1919
	goto L391
L391:
	;
	if base.Ui32(v1914) < base.Ui32(v2124) {
		goto L7
	} else {
		goto L393
	}
L392:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v2116)))
	v2123 = v2116
	v2124 = (v2118 + int32(7)) & int32(-8)
	goto L391
L393:
	;
	goto L394
L394:
	;
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v2171 != 0 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	F_LockBufferInternal(m, v98, int32(3))
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L1
	} else {
		goto L400
	}
L397:
	;
	v2172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+10)))
	if v2172&int32(4) == int32(0) {
		goto L396
	} else {
		goto L398
	}
L398:
	;
	F_visibilitymap_pin(m, l0, v97, v48+int32(24))
	mBase = m.M
	v2180 = m.ExcPending
	if v2180 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	goto L396
L400:
	;
	v2187 = int32(4)
	v2188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+14)))
	v2189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+12)))
	v2190 = v2188 - v2189
	if v2190 <= v2187 {
		goto L402
	} else {
		goto L403
	}
L401:
	;
	if base.Ui32(v2250) < base.Ui32(v2124) {
		goto L420
	} else {
		goto L421
	}
L402:
	;
	v2193 = v2187
	goto L404
L403:
	;
	v2193 = v2190
	goto L404
L404:
	;
	v2195 = v2193 - int32(4)
	if v2195 == int32(0) {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v2250 = int32(0)
	goto L401
L406:
	;
	goto L407
L407:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v2189) {
		goto L409
	} else {
		goto L410
	}
L408:
	;
	v2250 = v2195
	goto L401
L409:
	;
	v2206 = int32(base.Ui32(v2189+int32(_a_F_heap_update_22)) >> (uint(int32(2)) % 32))
	goto L411
L410:
	;
	v2206 = int32(0)
	goto L411
L411:
	;
	if base.Ui32(v2206&int32(_a_F_heap_update_4)) < base.Ui32(int32(291)) {
		goto L408
	} else {
		goto L412
	}
L412:
	;
	v2211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+10)))
	if v2211&int32(1) == int32(0) {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v2250 = int32(0)
	goto L401
L414:
	;
	goto L415
L415:
	;
	v2220 = int32(1)
	goto L416
L416:
	;
	v2229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117+int32(20)+v2220&int32(_a_F_heap_update_4)<<(uint(int32(2))%32))+1)))
	if v2229&int32(384) == int32(0) {
		goto L408
	} else {
		goto L418
	}
L417:
	;
	v2250 = int32(0)
	goto L401
L418:
	;
	v2235 = v2220 + int32(1)
	v2236 = int32(_a_F_heap_update_4)
	if base.Ui32(v2235&v2236) <= base.Ui32(v2206&v2236) {
		v2220 = v2235
		goto L416
	} else {
		goto L419
	}
L419:
	;
	goto L417
L420:
	;
	F_UnlockBuffer(m, v98)
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L1
	} else {
		goto L423
	}
L421:
	;
	goto L422
L422:
	;
	v2254 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v2254 == int32(0) {
		goto L425
	} else {
		goto L426
	}
L423:
	;
	goto L7
L424:
	;
	F_UnlockBuffer(m, v98)
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L1
	} else {
		goto L429
	}
L425:
	;
	v2257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+10)))
	if v2257&int32(4) != 0 {
		goto L424
	} else {
		goto L428
	}
L426:
	;
	goto L427
L427:
	;
	v2334 = v2123
	v2378 = v98
	goto L6
L428:
	;
	goto L427
L429:
	;
	goto L394
L430:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	F_errmsg(m, int32(_a_F_heap_update_27), int32(0))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L1
	} else {
		goto L432
	}
L432:
	;
	F_errfinish(m, int32(_a_F_heap_update_10), int32(3334), int32(_a_F_heap_update_11))
	mBase = m.M
	v2277 = m.ExcPending
	if v2277 != 0 {
		goto L1
	} else {
		goto L433
	}
L433:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L434:
	;
	v2334 = v2123
	v2378 = v2331
	goto L6
L435:
	;
	v2382 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v2382+(v2378^int32(-1))<<(uint(int32(2))%32))))
	v2396 = v2388
	goto L5
L436:
	;
	goto L437
L437:
	;
	v2390 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v2396 = v2390 + v2378<<(uint(int32(13))%32) + int32(-8192)
	goto L5
L438:
	;
	F_CheckForSerializableConflictIn(m, l0, v528, v2415)
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L1
	} else {
		goto L442
	}
L439:
	;
	v2400 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[12]))
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v2400+(v98^int32(-1))*int32(56))+16))
	v2415 = v2406
	goto L438
L440:
	;
	goto L441
L441:
	;
	v2408 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[13]))
	v2409 = int32(56)
	v2414 = *(*int32)(unsafe.Add(mBase, uint32(v2408+v98*v2409-v2409)+16))
	v2415 = v2414
	goto L438
L442:
	;
	v2418 = base.B2i32(v2378 != v98)
	if v2418 == int32(0) {
		goto L444
	} else {
		goto L445
	}
L443:
	;
	v2529 = int32(0)
	if base.B2i32(v443 == v2529)|base.B2i32(v82 == v2529) != 0 {
		v2574 = v2529
		goto L478
	} else {
		goto L479
	}
L444:
	;
	v2421 = int32(0)
	if base.B2i32(v443 == v2421)|base.B2i32(v73 == v2421) != 0 {
		v2468 = v2421
		goto L448
	} else {
		goto L449
	}
L445:
	;
	goto L446
L446:
	;
	v2519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+10)))
	v2521 = v2519 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+10)) = uint16(v2521)
	v2523 = int32(0)
	v2525 = v2523
	v2526 = v2523
	goto L443
L447:
	;
	if v2468 != 0 {
		v2525 = v2421
		v2526 = v2421
		goto L443
	} else {
		goto L460
	}
L448:
	;
	goto L447
L449:
	;
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	v2434 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v2433 < v2434 {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v2436 = v2433
	goto L452
L451:
	;
	v2436 = v2434
	goto L452
L452:
	;
	if v2436 <= int32(1) {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	v2439 = int32(1)
	goto L455
L454:
	;
	v2439 = v2436
	goto L455
L455:
	;
	v2440 = int32(8)
	v2445 = int32(0)
	goto L456
L456:
	;
	v2452 = v2445 << (uint(int32(2)) % 32)
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v73+v2440+v2452)))
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v443+v2440+v2452)))
	v2457 = v2454 & v2456
	v2459 = base.B2i32(v2457 != int32(0))
	if v2457 != 0 {
		v2468 = v2459
		goto L448
	} else {
		goto L458
	}
L457:
	;
	v2468 = v2459
	goto L448
L458:
	;
	v2461 = v2445 + int32(1)
	if v2461 != v2439 {
		v2445 = v2461
		goto L456
	} else {
		goto L459
	}
L459:
	;
	goto L457
L460:
	;
	v2470 = int32(0)
	if base.B2i32(v443 == v2470)|base.B2i32(v76 == v2470) != 0 {
		v2516 = v2470
		goto L462
	} else {
		goto L463
	}
L461:
	;
	if v2516 != 0 {
		goto L474
	} else {
		goto L475
	}
L462:
	;
	goto L461
L463:
	;
	v2481 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v2481 < v2482 {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v2484 = v2481
	goto L466
L465:
	;
	v2484 = v2482
	goto L466
L466:
	;
	if v2484 <= int32(1) {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	v2487 = int32(1)
	goto L469
L468:
	;
	v2487 = v2484
	goto L469
L469:
	;
	v2488 = int32(8)
	v2493 = int32(0)
	goto L470
L470:
	;
	v2500 = v2493 << (uint(int32(2)) % 32)
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(v76+v2488+v2500)))
	v2504 = *(*int32)(unsafe.Add(mBase, uint32(v443+v2488+v2500)))
	v2505 = v2502 & v2504
	v2507 = base.B2i32(v2505 != int32(0))
	if v2505 != 0 {
		v2516 = v2507
		goto L462
	} else {
		goto L472
	}
L471:
	;
	v2516 = v2507
	goto L462
L472:
	;
	v2509 = v2493 + int32(1)
	if v2509 != v2487 {
		v2493 = v2509
		goto L470
	} else {
		goto L473
	}
L473:
	;
	goto L471
L474:
	;
	v2517 = int32(2)
	goto L476
L475:
	;
	v2517 = v2470
	goto L476
L476:
	;
	v2525 = int32(1)
	v2526 = v2517
	goto L443
L477:
	;
	v2580 = F_ExtractReplicaIdentity(m, l0, v48+int32(32), (v2574|v436)&int32(1), v48+int32(31))
	mBase = m.M
	v2581 = m.ExcPending
	if v2581 != 0 {
		goto L1
	} else {
		goto L490
	}
L478:
	;
	goto L477
L479:
	;
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v2539 < v2540 {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	v2542 = v2539
	goto L482
L481:
	;
	v2542 = v2540
	goto L482
L482:
	;
	if v2542 <= int32(1) {
		goto L483
	} else {
		goto L484
	}
L483:
	;
	v2545 = int32(1)
	goto L485
L484:
	;
	v2545 = v2542
	goto L485
L485:
	;
	v2546 = int32(8)
	v2551 = int32(0)
	goto L486
L486:
	;
	v2558 = v2551 << (uint(int32(2)) % 32)
	v2560 = *(*int32)(unsafe.Add(mBase, uint32(v82+v2546+v2558)))
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(v443+v2546+v2558)))
	v2563 = v2560 & v2562
	v2565 = base.B2i32(v2563 != int32(0))
	if v2563 != 0 {
		v2574 = v2565
		goto L478
	} else {
		goto L488
	}
L487:
	;
	v2574 = v2565
	goto L478
L488:
	;
	v2567 = v2551 + int32(1)
	if v2567 != v2545 {
		v2551 = v2567
		goto L486
	} else {
		goto L489
	}
L489:
	;
	goto L487
L490:
	;
	v2582 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+10)))
	v2584 = v2582 & int32(4)
	if v2418 == int32(0) {
		goto L495
	} else {
		goto L496
	}
L491:
	;
	v2687 = int32(_a_F_heap_update_23)
	v2689 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_update[9])) = v2689 + int32(1)
	v2693 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if v2693 == int32(0) {
		goto L536
	} else {
		goto L537
	}
L492:
	;
	F_LockBufferInternal(m, v2673, int32(3))
	mBase = m.M
	v2680 = m.ExcPending
	if v2680 != 0 {
		goto L1
	} else {
		goto L534
	}
L493:
	;
	v2669 = int32(base.Ui32(v2584) >> (uint(int32(2)) % 32))
	if v2664 == int32(0) {
		v2683 = v2667
		v2685 = v2667
		v2686 = v2669
		goto L491
	} else {
		goto L533
	}
L494:
	;
	if v2652 == int32(0) {
		v2664 = v2657
		v2667 = v2656
		goto L493
	} else {
		goto L531
	}
L495:
	;
	v2587 = int32(0)
	if v2584 == v2587 {
		goto L498
	} else {
		goto L499
	}
L496:
	;
	goto L497
L497:
	;
	v2593 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2396)+10)))
	v2595 = v2593 & int32(4)
	v2597 = int32(base.Ui32(v2595) >> (uint(int32(2)) % 32))
	v2598 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v2599 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v2584&v2593 == int32(0) {
		goto L501
	} else {
		goto L502
	}
L498:
	;
	v2683 = int32(0)
	v2685 = v2587
	v2686 = v11
	goto L491
L499:
	;
	goto L500
L500:
	;
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	v2652 = v2591
	v2656 = v2587
	v2657 = int32(0)
	goto L494
L501:
	;
	if v2595 != 0 {
		goto L504
	} else {
		goto L505
	}
L502:
	;
	goto L503
L503:
	;
	if v2599 == v2598 {
		goto L508
	} else {
		goto L509
	}
L504:
	;
	v2604 = v2598
	goto L506
L505:
	;
	v2604 = int32(0)
	goto L506
L506:
	;
	if v2584 != 0 {
		v2652 = v2599
		v2656 = v2597
		v2657 = v2604
		goto L494
	} else {
		goto L507
	}
L507:
	;
	v2664 = v2604
	v2667 = v2597
	goto L493
L508:
	;
	v2673 = v2599
	v2674 = int32(1)
	v2676 = v2597
	v2677 = v11
	goto L492
L509:
	;
	goto L510
L510:
	;
	if v2595 != 0 {
		goto L511
	} else {
		goto L512
	}
L511:
	;
	v2608 = v2598
	goto L513
L512:
	;
	v2608 = int32(0)
	goto L513
L513:
	;
	if v2584 != 0 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v2610 = v2599
	goto L516
L515:
	;
	v2610 = int32(0)
	goto L516
L516:
	;
	if v2610 < int32(0) {
		goto L518
	} else {
		goto L519
	}
L517:
	;
	if v2608 < int32(0) {
		goto L522
	} else {
		goto L523
	}
L518:
	;
	v2614 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[12]))
	v2620 = *(*int32)(unsafe.Add(mBase, uint32(v2614+(v2610^int32(-1))*int32(56))+16))
	v2629 = v2620
	goto L517
L519:
	;
	goto L520
L520:
	;
	v2622 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[13]))
	v2623 = int32(56)
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v2622+v2610*v2623-v2623)+16))
	v2629 = v2628
	goto L517
L521:
	;
	v2649 = base.B2i32(base.Ui32(v2648) < base.Ui32(v2629))
	if base.Ui32(v2648) < base.Ui32(v2629) {
		goto L525
	} else {
		goto L526
	}
L522:
	;
	v2633 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[12]))
	v2639 = *(*int32)(unsafe.Add(mBase, uint32(v2633+(v2608^int32(-1))*int32(56))+16))
	v2648 = v2639
	goto L521
L523:
	;
	goto L524
L524:
	;
	v2641 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[13]))
	v2642 = int32(56)
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(v2641+v2608*v2642-v2642)+16))
	v2648 = v2647
	goto L521
L525:
	;
	v2650 = v2608
	goto L527
L526:
	;
	v2650 = v2610
	goto L527
L527:
	;
	if base.Ui32(v2648) < base.Ui32(v2629) {
		goto L528
	} else {
		goto L529
	}
L528:
	;
	v2651 = v2610
	goto L530
L529:
	;
	v2651 = v2608
	goto L530
L530:
	;
	v2652 = v2650
	v2656 = v2597
	v2657 = v2651
	goto L494
L531:
	;
	F_LockBufferInternal(m, v2652, int32(3))
	mBase = m.M
	v2662 = m.ExcPending
	if v2662 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	v2664 = v2657
	v2667 = v2656
	goto L493
L533:
	;
	v2673 = v2664
	v2674 = v2667
	v2676 = v2667
	v2677 = v2669
	goto L492
L534:
	;
	v2683 = v2674
	v2685 = v2676
	v2686 = v2677
	goto L491
L535:
	;
	v2708 = base.B2i32(v2378 == v98)
	if v2378 == v98 {
		goto L543
	} else {
		goto L544
	}
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v51
	goto L535
L537:
	;
	v2696 = int32(3)
	if base.B2i32(base.Ui32(v51) < base.Ui32(v2696))|base.B2i32(base.Ui32(v2693) < base.Ui32(v2696)) == int32(0) {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	if v51-v2693 < int32(0) {
		goto L536
	} else {
		goto L541
	}
L539:
	;
	goto L540
L540:
	;
	if base.Ui32(v2693) <= base.Ui32(v51) {
		goto L535
	} else {
		goto L542
	}
L541:
	;
	goto L535
L542:
	;
	goto L536
L543:
	;
	v2725 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v2726 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2725)+18)))
	if v2525 != 0 {
		goto L553
	} else {
		goto L554
	}
L544:
	;
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(v2396)+20))
	if v2709 == int32(0) {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2396)+20)) = v51
	goto L543
L546:
	;
	v2712 = int32(3)
	if base.B2i32(base.Ui32(v51) < base.Ui32(v2712))|base.B2i32(base.Ui32(v2709) < base.Ui32(v2712)) == int32(0) {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	if v51-v2709 < int32(0) {
		goto L545
	} else {
		goto L550
	}
L548:
	;
	goto L549
L549:
	;
	if base.Ui32(v2709) <= base.Ui32(v51) {
		goto L543
	} else {
		goto L551
	}
L550:
	;
	goto L543
L551:
	;
	goto L545
L552:
	;
	v2754 = int32(0)
	F_RelationPutHeapTuple(m, v2378, v2334, v2754)
	mBase = m.M
	v2757 = m.ExcPending
	if v2757 != 0 {
		goto L1
	} else {
		goto L556
	}
L553:
	;
	v2728 = v2726 | int32(_a_F_heap_update_28)
	*(*uint16)(unsafe.Add(mBase, uint32(v2725)+18)) = uint16(v2728)
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v2334)+16))
	v2731 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2730)+18)))
	v2732 = int32(_a_F_heap_update_1)
	v2733 = v2731 | v2732
	*(*uint16)(unsafe.Add(mBase, uint32(v2730)+18)) = uint16(v2733)
	v2735 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2735)+18)))
	v2738 = v2736 | v2732
	*(*uint16)(unsafe.Add(mBase, uint32(v2735)+18)) = uint16(v2738)
	goto L552
L554:
	;
	goto L555
L555:
	;
	v2741 = v2726 & int32(_a_F_heap_update_26)
	*(*uint16)(unsafe.Add(mBase, uint32(v2725)+18)) = uint16(v2741)
	v2743 = *(*int32)(unsafe.Add(mBase, uint32(v2334)+16))
	v2744 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2743)+18)))
	v2745 = int32(_a_F_heap_update_2)
	v2746 = v2744 & v2745
	*(*uint16)(unsafe.Add(mBase, uint32(v2743)+18)) = uint16(v2746)
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2749 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2748)+18)))
	v2751 = v2749 & v2745
	*(*uint16)(unsafe.Add(mBase, uint32(v2748)+18)) = uint16(v2751)
	goto L552
L556:
	;
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v2759 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2758)+20)))
	v2761 = v2759 & int32(_a_F_heap_update_24)
	*(*uint16)(unsafe.Add(mBase, uint32(v2758)+20)) = uint16(v2761)
	v2763 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2758)+18)))
	v2765 = v2763 & int32(_a_F_heap_update_25)
	*(*uint16)(unsafe.Add(mBase, uint32(v2758)+18)) = uint16(v2765)
	v2767 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v2768 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2767)+4)) = v2768
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v2771 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2770)+20)))
	v2772 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+10)))
	v2773 = v2771 | v2772
	*(*uint16)(unsafe.Add(mBase, uint32(v2770)+20)) = uint16(v2773)
	v2775 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2770)+18)))
	v2776 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+8)))
	v2777 = v2775 | v2776
	*(*uint16)(unsafe.Add(mBase, uint32(v2770)+18)) = uint16(v2777)
	v2779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+19)))
	v2780 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v2781 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v2780)+8)) = v2781
	v2783 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2780)+20)))
	v2790 = v2783&int32(_a_F_heap_update_14) | v2779<<(uint(int32(5))%32)&int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v2780)+20)) = uint16(v2790)
	v2792 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v2793 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2334)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2792)+16)) = uint16(v2793)
	v2795 = *(*int32)(unsafe.Add(mBase, uint32(v2334)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2792)+12)) = v2795
	if v2584 != 0 {
		goto L559
	} else {
		goto L560
	}
L557:
	;
	if v2708 == int32(0) {
		goto L575
	} else {
		goto L576
	}
L558:
	;
	if v2378 < int32(0) {
		goto L571
	} else {
		goto L572
	}
L559:
	;
	v2797 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	v2799 = F_visibilitymap_clear(m, v97, v2797, int32(3))
	mBase = m.M
	v2800 = m.ExcPending
	if v2800 != 0 {
		goto L1
	} else {
		goto L562
	}
L560:
	;
	goto L561
L561:
	;
	v2821 = int32(0)
	if v2685 == v2821 {
		v2861 = v2754
		v2862 = v2821
		v2863 = v2821
		goto L557
	} else {
		goto L569
	}
L562:
	;
	if v2799 != 0 {
		goto L563
	} else {
		goto L564
	}
L563:
	;
	if v2685 == int32(0) {
		goto L566
	} else {
		goto L567
	}
L564:
	;
	goto L565
L565:
	;
	v2817 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+10)))
	v2819 = v2817 & int32(_a_F_heap_update_29)
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+10)) = uint16(v2819)
	goto L561
L566:
	;
	v2803 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+10)))
	v2805 = v2803 & int32(_a_F_heap_update_29)
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+10)) = uint16(v2805)
	v2861 = v2754
	v2862 = int32(1)
	v2863 = int32(0)
	goto L557
L567:
	;
	goto L568
L568:
	;
	v2809 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v2810 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	v2811 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+10)))
	v2813 = v2811 & int32(_a_F_heap_update_29)
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+10)) = uint16(v2813)
	v2827 = base.B2i32(v2809 == v2810)
	v2828 = base.B2i32(v2809 != v2810)
	goto L558
L569:
	;
	v2827 = v2754
	v2828 = v2821
	goto L558
L570:
	;
	v2849 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v2851 = F_visibilitymap_clear(m, v2848, v2849, int32(3))
	mBase = m.M
	v2852 = m.ExcPending
	if v2852 != 0 {
		goto L1
	} else {
		goto L574
	}
L571:
	;
	v2833 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[12]))
	v2839 = *(*int32)(unsafe.Add(mBase, uint32(v2833+(v2378^int32(-1))*int32(56))+16))
	v2848 = v2839
	goto L570
L572:
	;
	goto L573
L573:
	;
	v2841 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[13]))
	v2842 = int32(56)
	v2847 = *(*int32)(unsafe.Add(mBase, uint32(v2841+v2378*v2842-v2842)+16))
	v2848 = v2847
	goto L570
L574:
	;
	v2853 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2396)+10)))
	v2855 = v2853 & int32(_a_F_heap_update_29)
	*(*uint16)(unsafe.Add(mBase, uint32(v2396)+10)) = uint16(v2855)
	v2861 = int32(1)
	v2862 = v2828
	v2863 = v2851 | v2827
	goto L557
L575:
	;
	F_MarkBufferDirty(m, v2378)
	mBase = m.M
	v2867 = m.ExcPending
	if v2867 != 0 {
		goto L1
	} else {
		goto L578
	}
L576:
	;
	goto L577
L577:
	;
	F_MarkBufferDirty(m, v98)
	mBase = m.M
	v2869 = m.ExcPending
	if v2869 != 0 {
		goto L1
	} else {
		goto L579
	}
L578:
	;
	goto L577
L579:
	;
	v2870 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2870)+118)))
	if v2871 != int32(112) {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	v3763 = int32(_a_F_heap_update_23)
	v3765 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_update[9])) = v3765 - int32(1)
	if v2686 != 0 {
		goto L757
	} else {
		goto L758
	}
L581:
	;
	v2875 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[10]))
	if v2875 <= int32(0) {
		goto L584
	} else {
		goto L585
	}
L582:
	;
	v2915 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	v2916 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v2917 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+60)) = uint16(v2917)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+58)) = uint16(v2917)
	if v2378 < v2917 {
		goto L602
	} else {
		goto L603
	}
L583:
	;
	v2892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L592
L584:
	;
	v2878 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2878 != 0 {
		goto L580
	} else {
		goto L587
	}
L585:
	;
	goto L586
L586:
	;
	if v2875 != int32(1) {
		goto L583
	} else {
		goto L590
	}
L587:
	;
	v2879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2879 != 0 {
		goto L580
	} else {
		goto L588
	}
L588:
	;
	v2881 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_update[14])))
	if v2881 == int32(1) {
		goto L583
	} else {
		goto L589
	}
L589:
	;
	goto L582
L590:
	;
	v2887 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_update[14])))
	if v2887&int32(1) == int32(0) {
		goto L582
	} else {
		goto L591
	}
L591:
	;
	goto L583
L592:
	;
	if base.B2i32(base.Ui32(v2892) < base.Ui32(int32(_a_F_heap_update_30))) == int32(0) {
		goto L593
	} else {
		goto L594
	}
L593:
	;
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v2897 == int32(0) {
		goto L582
	} else {
		goto L596
	}
L594:
	;
	goto L595
L595:
	;
	F_log_heap_new_cid(m, l0, v48+int32(32))
	mBase = m.M
	v2911 = m.ExcPending
	if v2911 != 0 {
		goto L1
	} else {
		goto L599
	}
L596:
	;
	v2900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2900)+119)))
	switch v2901 - int32(109) {
	case 0, 5:
		goto L597
	default:
		goto L582
	}
L597:
	;
	v2904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2897)+112)))
	if v2904 != int32(1) {
		goto L582
	} else {
		goto L598
	}
L598:
	;
	goto L595
L599:
	;
	F_log_heap_new_cid(m, l0, v2334)
	mBase = m.M
	v2913 = m.ExcPending
	if v2913 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	goto L582
L601:
	;
	v2939 = int32(0)
	if v526 != 0 {
		v2968 = v2939
		goto L605
	} else {
		goto L606
	}
L602:
	;
	v2924 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v2930 = *(*int32)(unsafe.Add(mBase, uint32(v2924+(v2378^int32(-1))<<(uint(int32(2))%32))))
	v2938 = v2930
	goto L601
L603:
	;
	goto L604
L604:
	;
	v2932 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v2938 = v2932 + v2378<<(uint(int32(13))%32) + int32(-8192)
	goto L601
L605:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v2970 = m.ExcPending
	if v2970 != 0 {
		goto L1
	} else {
		goto L619
	}
L606:
	;
	v2941 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[10]))
	if v2941 <= int32(1) {
		goto L607
	} else {
		goto L608
	}
L607:
	;
	v2945 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_update[14])))
	if v2945&int32(1) == int32(0) {
		v2968 = v2939
		goto L605
	} else {
		goto L610
	}
L608:
	;
	goto L609
L609:
	;
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2950)+118)))
	if v2951 != int32(112) {
		v2968 = v2939
		goto L605
	} else {
		goto L611
	}
L610:
	;
	goto L609
L611:
	;
	if v2941 <= int32(0) {
		goto L612
	} else {
		goto L613
	}
L612:
	;
	v2956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2956 != 0 {
		v2968 = v2939
		goto L605
	} else {
		goto L615
	}
L613:
	;
	goto L614
L614:
	;
	v2958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2950)+119)))
	if v2958 == int32(102) {
		v2968 = v2939
		goto L605
	} else {
		goto L617
	}
L615:
	;
	v2957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2957 != 0 {
		v2968 = v2939
		goto L605
	} else {
		goto L616
	}
L616:
	;
	goto L614
L617:
	;
	v2961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L618
L618:
	;
	v2968 = base.B2i32(base.Ui32(v2961) < base.Ui32(int32(_a_F_heap_update_30))) ^ int32(1)
	goto L605
L619:
	;
	v2971 = *(*int32)(unsafe.Add(mBase, uint32(v2334)+16))
	v2972 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2971)+18)))
	if v2968|v2418 != 0 {
		goto L621
	} else {
		goto L622
	}
L620:
	;
	v3426 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+87)) = uint8(v3425)
	if v2968 == v3426 {
		goto L671
	} else {
		goto L672
	}
L621:
	;
	v3373 = int32(0)
	v3375 = int32(2)
	v3376 = int32(base.Ui32(v2584) >> (uint(v3375) % 32))
	if v2861 != 0 {
		goto L668
	} else {
		goto L669
	}
L622:
	;
	v2976 = m.G0
	v2978 = v2976 - int32(16)
	m.G0 = v2978
	F_GetFullPageWriteInfo(m, v2978+int32(8), v2978+int32(7))
	mBase = m.M
	if v98 < int32(0) {
		goto L625
	} else {
		goto L626
	}
L623:
	;
	if v3013 != 0 {
		goto L621
	} else {
		goto L633
	}
L624:
	;
	v3003 = int32(1)
	v3004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2978)+7)))
	if v3004 == v3003 {
		goto L629
	} else {
		goto L630
	}
L625:
	;
	v2988 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v2994 = *(*int32)(unsafe.Add(mBase, uint32(v2988+(v98^int32(-1))<<(uint(int32(2))%32))))
	v3002 = v2994
	goto L624
L626:
	;
	goto L627
L627:
	;
	v2996 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v3002 = v2996 + v98<<(uint(int32(13))%32) + int32(-8192)
	goto L624
L628:
	;
	m.G0 = v2978 + int32(16)
	goto L623
L629:
	;
	v3007 = *(*int64)(unsafe.Add(mBase, uint32(v2978)+8))
	v3008 = *(*int64)(unsafe.Add(mBase, uint32(v3002)))
	if base.Ui64(base.I64_rotl(v3008, int64(32))) <= base.Ui64(v3007) {
		v3013 = v3003
		goto L628
	} else {
		goto L632
	}
L630:
	;
	goto L631
L631:
	;
	v3013 = int32(0)
	goto L628
L632:
	;
	goto L631
L633:
	;
	v3017 = *(*int32)(unsafe.Add(mBase, uint32(v2334)))
	v3018 = *(*int32)(unsafe.Add(mBase, uint32(v2334)+16))
	v3019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3018)+22)))
	v3020 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v3021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3020)+22)))
	v3022 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+60)) = uint16(v3022)
	v3025 = *(*int32)(unsafe.Add(mBase, uint32(v48)+32))
	v3026 = v3025 - v3021
	v3027 = v3017 - v3019
	if v3026 < v3027 {
		goto L635
	} else {
		goto L636
	}
L634:
	;
	v3190 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+58)) = uint16(v3190)
	v3194 = v3029 - v3149&int32(_a_F_heap_update_4)
	if v3190 < v3194 {
		goto L649
	} else {
		goto L650
	}
L635:
	;
	v3029 = v3026
	goto L637
L636:
	;
	v3029 = v3027
	goto L637
L637:
	;
	if int32(0) < v3029 {
		goto L638
	} else {
		goto L639
	}
L638:
	;
	v3039 = v3022
	v3041 = int32(0)
	goto L641
L639:
	;
	goto L640
L640:
	;
	v3142 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+60)) = uint16(v3142)
	v3149 = v3142
	goto L634
L641:
	;
	v3081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3041+(v3019+v3018)))))
	v3083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3041+(v3020+v3021)))))
	if v3081 == v3083 {
		goto L643
	} else {
		goto L644
	}
L642:
	;
	if base.Ui32(int32(2)) < base.Ui32(v3091&int32(_a_F_heap_update_4)) {
		v3149 = v3091
		goto L634
	} else {
		goto L647
	}
L643:
	;
	v3086 = v3039 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+60)) = uint16(v3086)
	v3089 = v3086 & int32(_a_F_heap_update_4)
	if base.Ui32(v3089) < base.Ui32(v3029) {
		v3039 = v3086
		v3041 = v3089
		goto L641
	} else {
		goto L646
	}
L644:
	;
	v3091 = v3039
	goto L645
L645:
	;
	goto L642
L646:
	;
	v3091 = v3086
	goto L645
L647:
	;
	goto L640
L648:
	;
	v3358 = int32(2)
	v3359 = int32(base.Ui32(v2584) >> (uint(v3358) % 32))
	if v2861 != 0 {
		goto L659
	} else {
		goto L660
	}
L649:
	;
	v3199 = int32(0)
	v3207 = v3199
	v3217 = v3199
	goto L652
L650:
	;
	goto L651
L651:
	;
	v3310 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+58)) = uint16(v3310)
	v3319 = v3310
	goto L648
L652:
	;
	v3247 = v3217 ^ int32(-1)
	v3249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3017+v3018+v3247))))
	v3251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3247+(v3020+v3025)))))
	if v3249 == v3251 {
		goto L654
	} else {
		goto L655
	}
L653:
	;
	if base.Ui32(int32(2)) < base.Ui32(v3259&int32(_a_F_heap_update_4)) {
		v3319 = v3259
		goto L648
	} else {
		goto L658
	}
L654:
	;
	v3254 = v3207 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+58)) = uint16(v3254)
	v3257 = v3254 & int32(_a_F_heap_update_4)
	if base.Ui32(v3257) < base.Ui32(v3194) {
		v3207 = v3254
		v3217 = v3257
		goto L652
	} else {
		goto L657
	}
L655:
	;
	v3259 = v3207
	goto L656
L656:
	;
	goto L653
L657:
	;
	v3259 = v3254
	goto L656
L658:
	;
	goto L651
L659:
	;
	v3362 = v3359 | v3358
	goto L661
L660:
	;
	v3362 = v3359
	goto L661
L661:
	;
	if v3149&int32(_a_F_heap_update_4) != 0 {
		goto L662
	} else {
		goto L663
	}
L662:
	;
	v3367 = v3362 | int32(32)
	goto L664
L663:
	;
	v3367 = v3362
	goto L664
L664:
	;
	if v3319&int32(_a_F_heap_update_4) != 0 {
		goto L665
	} else {
		goto L666
	}
L665:
	;
	v3372 = v3367 | int32(64)
	goto L667
L666:
	;
	v3372 = v3367
	goto L667
L667:
	;
	v3384 = v3149
	v3386 = v3319
	v3425 = v3372
	goto L620
L668:
	;
	v3379 = v3376 | v3375
	goto L670
L669:
	;
	v3379 = v3376
	goto L670
L670:
	;
	v3384 = v3373
	v3386 = v3373
	v3425 = v3379
	goto L620
L671:
	;
	if v2972 < v3426 {
		goto L677
	} else {
		goto L678
	}
L672:
	;
	v3432 = v3425 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+87)) = uint8(v3432)
	if v2580 == int32(0) {
		goto L671
	} else {
		goto L673
	}
L673:
	;
	v3438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3438)+130)))
	if v3439 == int32(102) {
		goto L674
	} else {
		goto L675
	}
L674:
	;
	v3442 = int32(20)
	goto L676
L675:
	;
	v3442 = int32(24)
	goto L676
L676:
	;
	v3443 = v3442 | v3425
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+87)) = uint8(v3443)
	goto L671
L677:
	;
	v3447 = int32(64)
	goto L679
L678:
	;
	v3447 = int32(32)
	goto L679
L679:
	;
	v3448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2334)+8)))
	if v3448 != int32(1) {
		goto L681
	} else {
		goto L682
	}
L680:
	;
	v3471 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+40)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+84)) = uint16(v3471)
	v3473 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v3474 = *(*int32)(unsafe.Add(mBase, uint32(v3473)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+80)) = v3474
	v3476 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3473)+18)))
	v3477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3473)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+92)) = uint16(v3448)
	v3481 = int32(16)
	v3483 = int32(1)
	v3487 = int32(4)
	v3502 = int32(base.Ui32(v3476)>>(uint(int32(9))%32))&v3481 | (int32(base.Ui32(v3477)>>(uint(v3483)%32))&int32(8) | (int32(base.Ui32(v3477)>>(uint(v3487)%32))&v3487 | (int32(base.Ui32(v3477)>>(uint(int32(12))%32))&v3483 | int32(base.Ui32(v3477)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+86)) = uint8(v3502)
	v3504 = *(*int32)(unsafe.Add(mBase, uint32(v2334)+16))
	v3505 = *(*int32)(unsafe.Add(mBase, uint32(v3504)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+88)) = v3505
	if v2968 != 0 {
		goto L690
	} else {
		goto L691
	}
L681:
	;
	v3469 = v3447
	v3470 = int32(8)
	goto L680
L682:
	;
	goto L683
L683:
	;
	v3454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2938)+12)))
	v3463 = base.B2i32(base.Ui32(int32(24)) < base.Ui32(v3454)) & base.B2i32((v3454+int32(_a_F_heap_update_22))&int32(_a_F_heap_update_31) == int32(4))
	if v3463 != 0 {
		goto L684
	} else {
		goto L685
	}
L684:
	;
	v3464 = int32(14)
	goto L686
L685:
	;
	v3464 = int32(8)
	goto L686
L686:
	;
	if v3463 != 0 {
		goto L687
	} else {
		goto L688
	}
L687:
	;
	v3467 = v3447 | int32(-128)
	goto L689
L688:
	;
	v3467 = v3447
	goto L689
L689:
	;
	v3469 = v3467
	v3470 = v3464
	goto L680
L690:
	;
	v3510 = v3470 | v3481
	goto L692
L691:
	;
	v3510 = v3470
	goto L692
L692:
	;
	F_XLogRegisterBuffer(m, int32(0), v2378, v3510)
	mBase = m.M
	v3512 = m.ExcPending
	if v3512 != 0 {
		goto L1
	} else {
		goto L693
	}
L693:
	;
	if v2708 == int32(0) {
		goto L694
	} else {
		goto L695
	}
L694:
	;
	F_XLogRegisterBuffer(m, int32(1), v98, int32(8))
	mBase = m.M
	v3518 = m.ExcPending
	if v3518 != 0 {
		goto L1
	} else {
		goto L697
	}
L695:
	;
	goto L696
L696:
	;
	F_XLogRegisterData(m, v48+int32(80), int32(14))
	mBase = m.M
	v3523 = m.ExcPending
	if v3523 != 0 {
		goto L1
	} else {
		goto L698
	}
L697:
	;
	goto L696
L698:
	;
	if (v3384|v3386)&int32(_a_F_heap_update_4) == int32(0) {
		goto L699
	} else {
		goto L700
	}
L699:
	;
	v3562 = *(*int32)(unsafe.Add(mBase, uint32(v2334)+16))
	v3563 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3562)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+72)) = uint16(v3563)
	v3565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3562)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+74)) = uint16(v3565)
	v3567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3562)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+76)) = uint8(v3567)
	F_XLogRegisterBufData(m, int32(0), v48+int32(72), int32(5))
	mBase = m.M
	v3574 = m.ExcPending
	if v3574 != 0 {
		goto L1
	} else {
		goto L710
	}
L700:
	;
	v3529 = int32(_a_F_heap_update_4)
	v3531 = int32(0)
	if base.B2i32(v3386&v3529 == v3531)|base.B2i32(v3384&v3529 == v3531) == v3531 {
		goto L701
	} else {
		goto L702
	}
L701:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+64)) = uint16(v3386)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+62)) = uint16(v3384)
	F_XLogRegisterBufData(m, int32(0), v48+int32(62), int32(4))
	mBase = m.M
	v3547 = m.ExcPending
	if v3547 != 0 {
		goto L1
	} else {
		goto L704
	}
L702:
	;
	goto L703
L703:
	;
	if v3384&int32(_a_F_heap_update_4) != 0 {
		goto L705
	} else {
		goto L706
	}
L704:
	;
	goto L699
L705:
	;
	F_XLogRegisterBufData(m, int32(0), v48+int32(60), int32(2))
	mBase = m.M
	v3555 = m.ExcPending
	if v3555 != 0 {
		goto L1
	} else {
		goto L708
	}
L706:
	;
	goto L707
L707:
	;
	F_XLogRegisterBufData(m, int32(0), v48+int32(58), int32(2))
	mBase = m.M
	v3561 = m.ExcPending
	if v3561 != 0 {
		goto L1
	} else {
		goto L709
	}
L708:
	;
	goto L699
L709:
	;
	goto L699
L710:
	;
	v3575 = *(*int32)(unsafe.Add(mBase, uint32(v2334)+16))
	v3576 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+60)))
	if v3576 == int32(0) {
		goto L712
	} else {
		goto L713
	}
L711:
	;
	v3618 = int32(0)
	if base.B2i32(v2580 == v3618)|(v2968^int32(1)) == v3618 {
		goto L721
	} else {
		goto L722
	}
L712:
	;
	v3580 = int32(23)
	v3582 = *(*int32)(unsafe.Add(mBase, uint32(v2334)))
	v3583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+58)))
	F_XLogRegisterBufData(m, int32(0), v3575+v3580, v3582-v3583-v3580)
	mBase = m.M
	v3588 = m.ExcPending
	if v3588 != 0 {
		goto L1
	} else {
		goto L715
	}
L713:
	;
	goto L714
L714:
	;
	v3589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3575)+22)))
	v3591 = v3589 - int32(23)
	if v3591 != 0 {
		goto L716
	} else {
		goto L717
	}
L715:
	;
	goto L711
L716:
	;
	F_XLogRegisterBufData(m, int32(0), v3575+int32(23), v3591)
	mBase = m.M
	v3596 = m.ExcPending
	if v3596 != 0 {
		goto L1
	} else {
		goto L719
	}
L717:
	;
	v3601 = v3576
	v3602 = v3575
	v3603 = int32(23)
	goto L718
L718:
	;
	v3607 = v3603 + v3601&int32(_a_F_heap_update_4)
	v3609 = *(*int32)(unsafe.Add(mBase, uint32(v2334)))
	v3610 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+58)))
	F_XLogRegisterBufData(m, int32(0), v3607+v3602, v3609-(v3607+v3610))
	mBase = m.M
	v3614 = m.ExcPending
	if v3614 != 0 {
		goto L1
	} else {
		goto L720
	}
L719:
	;
	v3597 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+60)))
	v3598 = *(*int32)(unsafe.Add(mBase, uint32(v2334)+16))
	v3599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3598)+22)))
	v3601 = v3597
	v3602 = v3598
	v3603 = v3599
	goto L718
L720:
	;
	goto L711
L721:
	;
	v3625 = *(*int32)(unsafe.Add(mBase, uint32(v2580)+16))
	v3626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3625)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+66)) = uint16(v3626)
	v3628 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3625)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+68)) = uint16(v3628)
	v3630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3625)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+70)) = uint8(v3630)
	F_XLogRegisterData(m, v48+int32(66), int32(5))
	mBase = m.M
	v3636 = m.ExcPending
	if v3636 != 0 {
		goto L1
	} else {
		goto L724
	}
L722:
	;
	goto L723
L723:
	;
	if v2863 != 0 {
		goto L726
	} else {
		goto L727
	}
L724:
	;
	v3637 = *(*int32)(unsafe.Add(mBase, uint32(v2580)+16))
	v3638 = int32(23)
	v3640 = *(*int32)(unsafe.Add(mBase, uint32(v2580)))
	F_XLogRegisterData(m, v3637+v3638, v3640-v3638)
	mBase = m.M
	v3644 = m.ExcPending
	if v3644 != 0 {
		goto L1
	} else {
		goto L725
	}
L725:
	;
	goto L723
L726:
	;
	v3647 = v2916
	goto L728
L727:
	;
	v3647 = int32(0)
	goto L728
L728:
	;
	if v3647 != 0 {
		goto L729
	} else {
		goto L730
	}
L729:
	;
	F_XLogRegisterBuffer(m, int32(2), v3647, int32(0))
	mBase = m.M
	v3651 = m.ExcPending
	if v3651 != 0 {
		goto L1
	} else {
		goto L732
	}
L730:
	;
	goto L731
L731:
	;
	if v2862 != 0 {
		goto L733
	} else {
		goto L734
	}
L732:
	;
	goto L731
L733:
	;
	v3653 = v2915
	goto L735
L734:
	;
	v3653 = int32(0)
	goto L735
L735:
	;
	if v3653 != 0 {
		goto L736
	} else {
		goto L737
	}
L736:
	;
	F_XLogRegisterBuffer(m, int32(3), v3653, int32(0))
	mBase = m.M
	v3657 = m.ExcPending
	if v3657 != 0 {
		goto L1
	} else {
		goto L739
	}
L737:
	;
	goto L738
L738:
	;
	v3659 = int32(_a_F_heap_update_32)
	v3661 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_update[15])))
	v3662 = v3661 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_update[15])) = uint8(v3662)
	goto L740
L739:
	;
	goto L738
L740:
	;
	v3667 = F_XLogInsert(m, int32(10), v3469&int32(255))
	mBase = m.M
	v3668 = m.ExcPending
	if v3668 != 0 {
		goto L1
	} else {
		goto L741
	}
L741:
	;
	v3670 = base.I64_rotl(v3667, int64(32))
	if v2708 == int32(0) {
		goto L742
	} else {
		goto L743
	}
L742:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2396))) = v3670
	goto L744
L743:
	;
	goto L744
L744:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = v3670
	if v2862 != 0 {
		goto L745
	} else {
		goto L746
	}
L745:
	;
	v3675 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v3675 < int32(0) {
		goto L749
	} else {
		goto L750
	}
L746:
	;
	goto L747
L747:
	;
	if v2863 == int32(0) {
		goto L580
	} else {
		goto L752
	}
L748:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3693))) = v3670
	goto L747
L749:
	;
	v3679 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v3685 = *(*int32)(unsafe.Add(mBase, uint32(v3679+(v3675^int32(-1))<<(uint(int32(2))%32))))
	v3693 = v3685
	goto L748
L750:
	;
	goto L751
L751:
	;
	v3687 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v3693 = v3687 + v3675<<(uint(int32(13))%32) + int32(-8192)
	goto L748
L752:
	;
	v3698 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v3698 < int32(0) {
		goto L754
	} else {
		goto L755
	}
L753:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3716))) = v3670
	goto L580
L754:
	;
	v3702 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v3708 = *(*int32)(unsafe.Add(mBase, uint32(v3702+(v3698^int32(-1))<<(uint(int32(2))%32))))
	v3716 = v3708
	goto L753
L755:
	;
	goto L756
L756:
	;
	v3710 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v3716 = v3710 + v3698<<(uint(int32(13))%32) + int32(-8192)
	goto L753
L757:
	;
	v3769 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	F_UnlockBuffer(m, v3769)
	mBase = m.M
	v3771 = m.ExcPending
	if v3771 != 0 {
		goto L1
	} else {
		goto L760
	}
L758:
	;
	goto L759
L759:
	;
	if v2683 != 0 {
		goto L761
	} else {
		goto L762
	}
L760:
	;
	goto L759
L761:
	;
	v3772 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	F_UnlockBuffer(m, v3772)
	mBase = m.M
	v3774 = m.ExcPending
	if v3774 != 0 {
		goto L1
	} else {
		goto L764
	}
L762:
	;
	goto L763
L763:
	;
	if v2708 == int32(0) {
		goto L766
	} else {
		goto L767
	}
L764:
	;
	goto L763
L765:
	;
	F_ReleaseBuffer(m, v98)
	mBase = m.M
	v3794 = m.ExcPending
	if v3794 != 0 {
		goto L1
	} else {
		goto L775
	}
L766:
	;
	F_UnlockBuffer(m, v2378)
	mBase = m.M
	v3778 = m.ExcPending
	if v3778 != 0 {
		goto L1
	} else {
		goto L769
	}
L767:
	;
	goto L768
L768:
	;
	F_UnlockBuffer(m, v98)
	mBase = m.M
	v3788 = m.ExcPending
	if v3788 != 0 {
		goto L1
	} else {
		goto L773
	}
L769:
	;
	F_UnlockBuffer(m, v98)
	mBase = m.M
	v3780 = m.ExcPending
	if v3780 != 0 {
		goto L1
	} else {
		goto L770
	}
L770:
	;
	F_CacheInvalidateHeapTuple(m, l0, v48+int32(32), v2334)
	mBase = m.M
	v3784 = m.ExcPending
	if v3784 != 0 {
		goto L1
	} else {
		goto L771
	}
L771:
	;
	F_ReleaseBuffer(m, v2378)
	mBase = m.M
	v3786 = m.ExcPending
	if v3786 != 0 {
		goto L1
	} else {
		goto L772
	}
L772:
	;
	goto L765
L773:
	;
	F_CacheInvalidateHeapTuple(m, l0, v48+int32(32), v2334)
	mBase = m.M
	v3792 = m.ExcPending
	if v3792 != 0 {
		goto L1
	} else {
		goto L774
	}
L774:
	;
	goto L765
L775:
	;
	v3795 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v3795 != 0 {
		goto L776
	} else {
		goto L777
	}
L776:
	;
	F_ReleaseBuffer(m, v3795)
	mBase = m.M
	v3797 = m.ExcPending
	if v3797 != 0 {
		goto L1
	} else {
		goto L779
	}
L777:
	;
	goto L778
L778:
	;
	v3798 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v3798 != 0 {
		goto L780
	} else {
		goto L781
	}
L779:
	;
	goto L778
L780:
	;
	F_ReleaseBuffer(m, v3798)
	mBase = m.M
	v3800 = m.ExcPending
	if v3800 != 0 {
		goto L1
	} else {
		goto L783
	}
L781:
	;
	goto L782
L782:
	;
	if v1054&int32(1) != 0 {
		goto L784
	} else {
		goto L785
	}
L783:
	;
	goto L782
L784:
	;
	v3803 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v3806 = *(*int32)(unsafe.Add(mBase, uint32(v3803*int32(12))+uint32(_c_F_heap_update[3])))
	F_UnlockTuple(m, l0, v528, v3806)
	mBase = m.M
	v3808 = m.ExcPending
	if v3808 != 0 {
		goto L1
	} else {
		goto L787
	}
L785:
	;
	goto L786
L786:
	;
	v3809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v3809 == int32(0) {
		goto L789
	} else {
		goto L790
	}
L787:
	;
	goto L786
L788:
	;
	if v2334 != l2 {
		goto L806
	} else {
		goto L807
	}
L789:
	;
	v3812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v3812 != int32(1) {
		goto L788
	} else {
		goto L792
	}
L790:
	;
	v3818 = v3809
	goto L791
L791:
	;
	v3820 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[0]))
	v3821 = *(*int32)(unsafe.Add(mBase, uint32(v3820)+28))
	goto L794
L792:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v3816 = m.ExcPending
	if v3816 != 0 {
		goto L1
	} else {
		goto L793
	}
L793:
	;
	v3817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v3818 = v3817
	goto L791
L794:
	;
	v3822 = *(*int32)(unsafe.Add(mBase, uint32(v3818)+8))
	if v3822 != 0 {
		goto L796
	} else {
		goto L797
	}
L795:
	;
	v3843 = *(*int64)(unsafe.Add(mBase, uint32(v3840)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3840)+8)) = v3843 + int64(1)
	if v2525|v2418 == int32(0) {
		goto L788
	} else {
		goto L802
	}
L796:
	;
	v3823 = *(*int32)(unsafe.Add(mBase, uint32(v3822)+56))
	if v3823 == v3821 {
		v3840 = v3822
		goto L795
	} else {
		goto L799
	}
L797:
	;
	goto L798
L798:
	;
	v3825 = F_pgstat_get_xact_stack_level(m, v3821)
	mBase = m.M
	v3826 = m.ExcPending
	if v3826 != 0 {
		goto L1
	} else {
		goto L800
	}
L799:
	;
	goto L798
L800:
	;
	v3828 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[16]))
	v3830 = F_MemoryContextAllocZero(m, v3828, int32(72))
	mBase = m.M
	v3831 = m.ExcPending
	if v3831 != 0 {
		goto L1
	} else {
		goto L801
	}
L801:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+56)) = v3821
	v3833 = *(*int32)(unsafe.Add(mBase, uint32(v3818)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+64)) = v3818
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+60)) = v3833
	v3836 = *(*int32)(unsafe.Add(mBase, uint32(v3825)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+68)) = v3836
	*(*int32)(unsafe.Add(mBase, uint32(v3825)+20)) = v3830
	*(*int32)(unsafe.Add(mBase, uint32(v3818)+8)) = v3830
	v3840 = v3830
	goto L795
L802:
	;
	if v2525 != 0 {
		goto L803
	} else {
		goto L804
	}
L803:
	;
	v3852 = int32(64)
	goto L805
L804:
	;
	v3852 = int32(72)
	goto L805
L805:
	;
	v3853 = v3818 + v3852
	v3854 = *(*int64)(unsafe.Add(mBase, uint32(v3853)))
	*(*int64)(unsafe.Add(mBase, uint32(v3853))) = v3854 + int64(1)
	goto L788
L806:
	;
	v3864 = v2334 + int32(4)
	v3865 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3864)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)) = uint16(v3865)
	v3867 = *(*int32)(unsafe.Add(mBase, uint32(v3864)))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v3867
	F_pfree(m, v2334)
	mBase = m.M
	v3870 = m.ExcPending
	if v3870 != 0 {
		goto L1
	} else {
		goto L809
	}
L807:
	;
	goto L808
L808:
	;
	if v2525 != 0 {
		goto L810
	} else {
		goto L811
	}
L809:
	;
	goto L808
L810:
	;
	v3873 = v2526
	goto L812
L811:
	;
	v3873 = int32(1)
	goto L812
L812:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v3873
	if v2580 == int32(0) {
		goto L813
	} else {
		goto L814
	}
L813:
	;
	F_bms_free(m, v73)
	mBase = m.M
	v3885 = m.ExcPending
	if v3885 != 0 {
		goto L1
	} else {
		goto L817
	}
L814:
	;
	v3877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+31)))
	if v3877&int32(1) == int32(0) {
		goto L813
	} else {
		goto L815
	}
L815:
	;
	F_pfree(m, v2580)
	mBase = m.M
	v3883 = m.ExcPending
	if v3883 != 0 {
		goto L1
	} else {
		goto L816
	}
L816:
	;
	goto L813
L817:
	;
	v3888 = int32(0)
	goto L4
L818:
	;
	F_bms_free(m, v3935)
	mBase = m.M
	v3980 = m.ExcPending
	if v3980 != 0 {
		goto L1
	} else {
		goto L819
	}
L819:
	;
	F_bms_free(m, v3964)
	mBase = m.M
	v3982 = m.ExcPending
	if v3982 != 0 {
		goto L1
	} else {
		goto L820
	}
L820:
	;
	F_bms_free(m, v3959)
	mBase = m.M
	v3984 = m.ExcPending
	if v3984 != 0 {
		goto L1
	} else {
		goto L821
	}
L821:
	;
	F_bms_free(m, v91)
	mBase = m.M
	v3986 = m.ExcPending
	if v3986 != 0 {
		goto L1
	} else {
		goto L822
	}
L822:
	;
	m.G0 = v48 + int32(96)
	return v3933
}
func F_rewrite_heap_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int64
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int64
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v422 int32
	_ = v422
	v16 = m.G0
	v18 = v16 - int32(80)
	m.G0 = v18
	v20 = int32(_a_F_rewrite_heap_tuple_0)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_rewrite_heap_tuple[0]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_rewrite_heap_tuple[0])) = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v27
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+20)))
	v34 = v32 & int32(15)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+20)) = uint16(v34)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+18)))
	v39 = v37 & int32(_a_F_rewrite_heap_tuple_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+18)) = uint16(v39)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+20)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+20)))
	v47 = v42 | v44&int32(_a_F_rewrite_heap_tuple_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+20)) = uint16(v47)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+136))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+140))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v56 = m.G0
	v58 = v56 + int32(-64)
	m.G0 = v58
	*(*int32)(unsafe.Add(mBase, uint32(v58)+44)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v58)+40)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v58)+36)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v58)+32)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v54
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v72)
	v80 = F_heap_prepare_freeze_tuple(m, v49, v56+int32(-40), v58, v56+int32(-12), v56+int32(-13))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v80 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v82
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+60)))
	if v84&int32(2) != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	m.G0 = v58 - int32(-64)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v102 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v101)+16)) = uint16(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v101)+12)) = int32(-1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+21)))
	if v107&int32(8) != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(2)
	goto L8
L7:
	;
	goto L8
L8:
	;
	if v84&int32(4) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(0)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+58)))
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+20)) = uint16(v93)
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+18)) = uint16(v95)
	goto L5
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_rewrite_heap_tuple[0])) = v21
	m.G0 = v18 + int32(80)
	return
L13:
	;
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+28)) = uint16(v201)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v203
	v208 = v18 + int32(8) | int32(4)
	v210 = v18 + int32(62)
	v212 = v18 + int32(44)
	v214 = v18 + int32(56)
	v218 = l2
	v219 = int32(0)
	goto L41
L14:
	;
	v110 = F_HeapTupleHeaderIsOnlyLocked(m, v106)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v110 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+16)))
	if v113 == int32(_a_F_rewrite_heap_tuple_3) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+12)))
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+14)))
	if v116&v117 == int32(_a_F_rewrite_heap_tuple_4) {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v122 = l1 + int32(4)
	v124 = v112 + int32(12)
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)))
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122))))
	v127 = int32(16)
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+2)))
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	if v125|v126<<(uint(v127)%32) == v130|v131<<(uint(v127)%32) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L19
L21:
	;
	if v141 != 0 {
		goto L13
	} else {
		goto L27
	}
L22:
	;
	goto L21
L23:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+4)))
	if v137 == v138 {
		v141 = int32(1)
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v141 = int32(0)
	goto L22
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = int64(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+20)))
	if v147&int32(_a_F_rewrite_heap_tuple_5) == int32(_a_F_rewrite_heap_tuple_6) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v159
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)) = uint16(v161)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v165 = v18 + int32(8)
	v166 = int32(0)
	v168 = F_hash_search(m, v163, v165, v166, v166)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L33
	}
L29:
	;
	v152 = F_HeapTupleGetUpdateXid(m, v146)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v156 = v146
	v157 = v155
	goto L28
L32:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v156 = v154
	v157 = v152
	goto L28
L33:
	;
	if v168 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v176 = F_hash_search(m, v172, v165, int32(1), v18+int32(7))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v185)+16)) = uint16(v186)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+12)) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v196 = F_hash_search(m, v190, v18+int32(8), int32(2), v18+int32(7))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v176)+16)) = uint16(v178)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+12)) = v180
	v182 = F_heap_copytuple(m, l2)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+20)) = v182
	goto L12
L39:
	;
	goto L13
L40:
	;
	if v219&int32(1) == int32(0) {
		goto L12
	} else {
		goto L85
	}
L41:
	;
	F_raw_heap_insert(m, l0, v218)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v407 = F_hash_search(m, v401, v18+int32(8), int32(1), v18+int32(7))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L84
	}
L43:
	;
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)) = uint16(v233)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v235
	v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+76)) = uint16(v237)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v239
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v241 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v218)+16))
	v332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v331)+20)))
	if v332&int32(_a_F_rewrite_heap_tuple_7) == int32(0) {
		goto L40
	} else {
		goto L64
	}
L45:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v218)+16))
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v245)+20)))
	v247 = int32(768)
	if v246&v247 != v247 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v252 = v251
	goto L48
L47:
	;
	v252 = int32(2)
	goto L48
L48:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v246&int32(_a_F_rewrite_heap_tuple_5) == int32(_a_F_rewrite_heap_tuple_6) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v262 = int32(0)
	v264 = int32(3)
	v272 = (base.B2i32(base.Ui32(v253) < base.Ui32(v264)) | base.B2i32(v262 <= v252-v253)) & base.B2i32(base.Ui32(v264) <= base.Ui32(v252))
	if base.Ui32(v261) < base.Ui32(v264) {
		v290 = v262
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v258 = F_HeapTupleGetUpdateXid(m, v245)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v261 = v260
	goto L49
L53:
	;
	v261 = v258
	goto L49
L54:
	;
	if v272|v290 != int32(1) {
		goto L44
	} else {
		goto L57
	}
L55:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v218)+16))
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275)+20)))
	if v276&int32(128)|base.B2i32(v276&int32(_a_F_rewrite_heap_tuple_8) == int32(64)) != 0 {
		v290 = v262
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v290 = base.B2i32(base.Ui32(v253) < base.Ui32(int32(3))) | base.B2i32(int32(0) <= v261-v253)
	goto L54
L57:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v296
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v295)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v298
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v214)+4)) = uint16(v300)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v212)+8)) = v305
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
	*(*int64)(unsafe.Add(mBase, uint32(v212))) = v307
	v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v210)+4)) = uint16(v309)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v311
	if v272 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_logical_rewrite_log_mapping(m, l0, v252, v18+int32(32))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v290&base.B2i32(v252 != v261) == int32(0) {
		goto L44
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	F_logical_rewrite_log_mapping(m, l0, v261, v18+int32(32))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L44
L64:
	;
	v337 = int32(768)
	if v332&v337 != v337 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v365
	v367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v208)+4)) = uint16(v367)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v369
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v374 = int32(0)
	v376 = F_hash_search(m, v371, v18+int32(8), v374, v374)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L75
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = int64(0)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v365 = v362
	goto L65
L67:
	;
	if base.Ui32(v341) < base.Ui32(v344) {
		goto L40
	} else {
		goto L74
	}
L68:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v342 = int32(3)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.B2i32(base.Ui32(v341) < base.Ui32(v342))|base.B2i32(base.Ui32(v344) < base.Ui32(v342)) != 0 {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(int32(2)) < base.Ui32(v351) {
		goto L40
	} else {
		goto L73
	}
L71:
	;
	if int32(0) <= v341-v344 {
		goto L66
	} else {
		goto L72
	}
L72:
	;
	goto L40
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(0)
	v365 = int32(2)
	goto L65
L74:
	;
	goto L66
L75:
	;
	if v376 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v219&int32(1) != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	goto L42
L79:
	;
	F_pfree(m, v218)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v376)+20))
	v383 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v376)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+28)) = uint16(v383)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v382)+16))
	v388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v387)+16)) = uint16(v388)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v387)+12)) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v398 = F_hash_search(m, v392, v18+int32(8), int32(2), v18+int32(7))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v218 = v382
	v219 = int32(1)
	goto L41
L84:
	;
	v409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v407)+16)) = uint16(v409)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v407)+12)) = v411
	goto L40
L85:
	;
	F_pfree(m, v218)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	goto L12
}
