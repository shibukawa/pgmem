package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CacheInvalidateRelSync(m *base.Module, l0 int32) {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_PrepareInvalidationState(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v18 = *(*int32)(unsafe.Add(mBase, _consts[1067]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v19 < v20 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v11 + int32(16)
	return
L4:
	;
	v23 = v19
	goto L7
L5:
	;
	goto L6
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[1068]))
	if v53 <= v20 {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v32 = v18 + v23<<(uint(int32(4))%32)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v33 == int32(250) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v36 == l0 {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v42 = v23 + int32(1)
	if v42 != v20 {
		v23 = v42
		goto L7
	} else {
		goto L14
	}
L12:
	;
	if v36 == int32(0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L8
L15:
	;
	if v18 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v75 = v18
	goto L17
L17:
	;
	v79 = v75 + v20<<(uint(int32(4))%32)
	v80 = int32(250)
	*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v80)
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+13)))
	*(*uint16)(unsafe.Add(mBase, uint32(v79)+1)) = uint16(v82)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v79)+3)) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v16
	v89 = v13 + int32(12)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v90 + int32(1)
	goto L3
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1068])) = v69
	*(*int32)(unsafe.Add(mBase, _consts[1067])) = v70
	v75 = v70
	goto L17
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v61 = F_MemoryContextAlloc(m, v59, int32(512))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v67 = F_repalloc(m, v18, v53<<(uint(int32(5))%32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v69 = int32(32)
	v70 = v61
	goto L18
L23:
	;
	v69 = v53 << (uint(int32(1)) % 32)
	v70 = v67
	goto L18
}
func F_CacheInvalidateRelcacheByRelid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(46249), v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(497347), int32(1697), int32(435048))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v30 = v28 + v29
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+117)))
			v33 = *(*int32)(unsafe.Add(mBase, _consts[128]))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
			v35 = F_PrepareInvalidationState(m)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				if v31 != 0 {
					v38 = int32(0)
				} else {
					v38 = v33
				}
				F_RegisterRelcacheInvalidation(m, v35, v38, v34)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_ReleaseCatCache(m, v11)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_CacheRegisterRelcacheCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	if int32(10) <= v5 {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(118573), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errfinish(m, int32(497347), int32(1862), int32(319134))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v22 = v5 << (uint(int32(3)) % 32)
		*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_consts[1077]))) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_consts[1078]))) = l0
		*(*int32)(unsafe.Add(mBase, _consts[1076])) = v5 + int32(1)
		return
	}
}
func F_CacheRegisterSyscacheCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v1 = l0
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if base.Ui32(v1) < base.Ui32(int32(85)) {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[1069]))
		if int32(64) <= v15 {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(118537), int32(0))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return
				} else {
					F_errfinish(m, int32(497347), int32(1823), int32(319104))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v19 = v1 << (uint(int32(1)) % 32)
			v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[1070]))))
			if v22 == int32(0) {
				v26 = v15 + int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[1070]))) = uint16(v26)
			} else {
				v31 = v22
				for {
					v36 = v31 * int32(12)
					v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[1071]))))
					if int32(0) < v39 {
						v31 = v39
						continue
					} else {
						break
					}
					break
				}
				v43 = v15 + int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[1071]))) = uint16(v43)
			}
			v53 = v15 * int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[1072]))) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[1073]))) = l1
			v62 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[1074]))) = uint16(v62)
			*(*uint16)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[1075]))) = uint16(v1)
			*(*int32)(unsafe.Add(mBase, _consts[1069])) = v15 + int32(1)
			m.G0 = v10 + int32(16)
			return
		}
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v1
			F_errmsg_internal(m, int32(487718), v10)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				F_errfinish(m, int32(497347), int32(1821), int32(319104))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
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
func F_cache_store_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	v8 = int32(4515248)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
	v15 = F_palloc(m, int32(8))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v22 = m.T0[v21].(func(*base.Module, int32, int32) int32)(m, l1, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v22
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v27 + base.I64_extend_i32_u(v28+int32(8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v34 == v24 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v15
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v9
	v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	if base.Ui64(v44) <= base.Ui64(v45) {
		v113 = int32(1)
		goto L8
	} else {
		goto L9
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v15
	goto L4
L6:
	;
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v15
	goto L4
L8:
	;
	return v113
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v49 = F_cache_reduce_memory(m, l0, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v49 == int32(0) {
		v113 = int32(0)
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v53 = int32(1)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
	if v54 == v53 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v57 == v48 {
		v113 = v53
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_prepare_probe_slot(m, l0, v48)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v62 = F_MemoizeHash_hash(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v66 = v62 & v65
	v69 = v64 + v66<<(uint(int32(4))%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+12)))
	if v70 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v104
	v113 = int32(1)
	goto L8
L19:
	;
	v73 = v69
	v74 = v66
	v75 = v64
	v77 = v65
	goto L22
L20:
	;
	goto L21
L21:
	;
	v104 = int32(0)
	goto L18
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	if v78 == v62 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v81 = F_MemoizeHash_equal(m, v61, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	v85 = v75
	v86 = v77
	goto L26
L26:
	;
	v89 = v86 & (v74 + int32(1))
	v92 = v85 + v89<<(uint(int32(4))%32)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+12)))
	if v93 != 0 {
		v73 = v92
		v74 = v89
		v75 = v85
		v77 = v86
		goto L22
	} else {
		goto L29
	}
L27:
	;
	if v81 != 0 {
		v104 = v73
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v85 = v84
	v86 = v83
	goto L26
L29:
	;
	goto L23
}
