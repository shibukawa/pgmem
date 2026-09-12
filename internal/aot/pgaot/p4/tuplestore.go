package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetTuplestoreDestReceiverParams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	return
}
func F_tuplestore_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int64
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v2 = l1
	v7 = F_palloc0(m, int32(120))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(0)
		v13 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+10)) = uint16(v13)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)) = uint8(v2)
		if l0 != 0 {
			v18 = int32(12)
		} else {
			v18 = int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18
		v20 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v20
		v26 = base.I64_extend_i32_s(l2) << (uint(int64(10)) % 64)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v26
		*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v26
		v30 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		v35 = F_GenerationContextCreate(m, v30, int32(154796), v20, int32(8192), int32(8388608))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v35
			v39 = *(*int32)(unsafe.Add(mBase, _consts[173]))
			v40 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+76)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v39
			v43 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+88)) = uint8(v43)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+84)) = int32(4096)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v40
			v50 = F_palloc(m, int32(16384))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+72)) = v50
				v53 = F_GetMemoryChunkSpace(m, v50)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+104)) = int32(8)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+96)) = int64(4294967296)
					v59 = *(*int64)(unsafe.Add(mBase, uint32(v7)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v59 - base.I64_extend_i32_u(v53)
					v64 = F_palloc(m, int32(192))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+92)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(v64))) = v18
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v7)+92))
						v69 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v68)+4)) = uint8(v69)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v7)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v69
						*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = int32(1866)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = int32(1867)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+60)) = int32(1868)
						return v7
					}
				}
			}
		}
	}
}
func F_tuplestore_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_tuplestore_gettuple(m, l0, l1, v8+int32(15))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
			if l2 == int32(0) {
				v25 = v16
				v26 = v12
				v29 = F_ExecStoreMinimalTuple(m, v26, l3, v25&int32(1))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return base.B2i32(v12 != int32(0))
				}
			} else {
				if v16&int32(1) != 0 {
					v25 = v16
					v26 = v12
					v29 = F_ExecStoreMinimalTuple(m, v26, l3, v25&int32(1))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return base.B2i32(v12 != int32(0))
					}
				} else {
					v22 = F_heap_copy_minimal_tuple(m, v12, int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = int32(1)
						v26 = v22
						v29 = F_ExecStoreMinimalTuple(m, v26, l3, v25&int32(1))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return base.B2i32(v12 != int32(0))
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
			m.T0[v32].(func(*base.Module, int32))(m, l3)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return base.B2i32(v12 != int32(0))
			}
		}
	}
}
func F_tuplestore_rescan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v7 = v3 + v4*int32(24)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v8 {
	case 0:
		v53 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v53
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v53)
		return
	case 1:
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(0)
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v11
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v11)
		return
	case 2:
		v15 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v15)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v21 = F_BufFileSeek(m, v17, v15, int64(0), v15)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v21 == int32(0) {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						F_errmsg(m, int32(368225), int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							F_errfinish(m, int32(475649), int32(1308), int32(270508))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
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
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(335882), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				F_errfinish(m, int32(475649), int32(1311), int32(270508))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
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
func F_tuplestore_skiptuples(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v98 int32
	_ = v98
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 <= int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v98
L2:
	;
	v98 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v98 = int32(0)
	goto L1
L6:
	;
	v17 = l1
	goto L9
L7:
	;
	goto L8
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v49 = v45 + v46*int32(24)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
	if l2 != 0 {
		goto L23
	} else {
		goto L24
	}
L9:
	;
	v25 = F_tuplestore_gettuple(m, l0, l2, v10+int32(15))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v98 = int32(1)
	goto L1
L11:
	;
	return int32(0)
L12:
	;
	if v25 == int32(0) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	if v31 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_pfree(m, v25)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v37 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v40 = int64(1)
	if base.Ui64(v40) < base.Ui64(v17) {
		v17 = v17 - v40
		goto L9
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	goto L10
L23:
	;
	if v50&int32(1) != 0 {
		v98 = v4
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v50&int32(1) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	if l1 <= base.I64_extend_i32_s(v53-v54) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v54 + base.I32_wrap_i64(l1)
	v98 = int32(1)
	goto L1
L28:
	;
	goto L29
L29:
	;
	v62 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)) = uint8(v62)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v53
	v98 = v4
	goto L1
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v76 < base.I64_extend_i32_s(v77-v78) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v76 = l1
	v77 = v69
	goto L30
L32:
	;
	goto L33
L33:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v70
	v76 = l1 - int64(1)
	v77 = v70
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v77 - base.I32_wrap_i64(v76)
	v98 = int32(1)
	goto L1
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v78
	goto L5
}
