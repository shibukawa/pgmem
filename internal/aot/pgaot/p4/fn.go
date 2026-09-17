package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_set_fn_opclass_options(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = int32(-1)
	v5 = int32(0)
	v10 = F_makeConst(m, int32(17), v4, v5, v4, l1, base.B2i32(l1 == v5), v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v10
		return
	}
}
func Fn13825(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v148
	v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v148)+11)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v155
	return v154
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v67 = l4
	v69 = l3
	goto L19
L5:
	;
	if v49-v50 == int32(0) {
		v148 = v11
		goto L1
	} else {
		goto L18
	}
L7:
	;
	goto L8
L8:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v19 = l0
	v20 = v11
	v21 = int32(10)
	v22 = v18
	goto L13
L10:
	;
	v45 = v11
	v49 = int32(0)
	goto L11
L11:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	goto L5
L12:
	;
	v45 = v40
	v49 = v42
	goto L11
L13:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v22 != v24)|base.B2i32(v24 == int32(0)) != 0 {
		v40 = v20
		v42 = v22
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v40 = v34
	v42 = int32(0)
	goto L12
L15:
	;
	v30 = v21 - int32(1)
	if v30 == int32(0) {
		v40 = v20
		v42 = v22
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v33 = int32(1)
	v34 = v20 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v35 != 0 {
		v19 = v19 + v33
		v20 = v34
		v21 = v30
		v22 = v35
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L4
L19:
	;
	v76 = v67 + (v69-v67)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76))))
	v78 = v60 - v77
	if v78 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(31)
L21:
	;
	goto L26
L22:
	;
	v129 = v78
	goto L23
L23:
	;
	v133 = base.B2i32(v129 < int32(0))
	if v129 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L24:
	;
	if v120 == int32(0) {
		v148 = v76
		goto L1
	} else {
		goto L37
	}
L26:
	;
	goto L27
L27:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v87 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v88 = l0
	v89 = v76
	v90 = int32(10)
	v91 = v87
	goto L32
L29:
	;
	v114 = v76
	v118 = int32(0)
	goto L30
L30:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v120 = v118 - v119
	goto L24
L31:
	;
	v114 = v109
	v118 = v111
	goto L30
L32:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if base.B2i32(v91 != v93)|base.B2i32(v93 == int32(0)) != 0 {
		v109 = v89
		v111 = v91
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v109 = v103
	v111 = int32(0)
	goto L31
L34:
	;
	v99 = v90 - int32(1)
	if v99 == int32(0) {
		v109 = v89
		v111 = v91
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v102 = int32(1)
	v103 = v89 + v102
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	if v104 != 0 {
		v88 = v88 + v102
		v89 = v103
		v90 = v99
		v91 = v104
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v129 = v120
	goto L23
L38:
	;
	v134 = v76 - int32(16)
	goto L40
L39:
	;
	v134 = v69
	goto L40
L40:
	;
	if v129 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v137 = v67
	goto L43
L42:
	;
	v137 = v76 + int32(16)
	goto L43
L43:
	;
	if base.Ui32(v137) <= base.Ui32(v134) {
		v67 = v137
		v69 = v134
		goto L19
	} else {
		goto L44
	}
L44:
	;
	goto L20
}
func Fn13829(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v3 = int32(0)
	v9 = F_SearchSysCacheList(m, l1, int32(1), l0, v3, v3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if int32(0) < v13 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = int32(0)
	v21 = v3
	goto L6
L4:
	;
	v40 = v3
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v9)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(48)+v19<<(uint(int32(2))%32))))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29)+4))
	v32 = F_lappend_oid(m, v21, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v40 = v32
	goto L5
L8:
	;
	v35 = v19 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if v35 < v36 {
		v19 = v35
		v21 = v32
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	return v40
}
func Fn13836(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	v9 = F_psprintf(m, l1, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v9
	}
}
func Fn13841(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	F_errstart_cold(m, int32(21), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		F_errcode(m, int32(1088))
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_Fn13841_0), int32(0))
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errdetail(m, int32(_a_Fn13841_1), int32(0))
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_Fn13841_2), l3, l2)
					v22 = m.ExcPending
					if v22 != 0 {
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
func Fn13847(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v11 < int32(20) {
		v15 = v11 << (uint(int32(3)) % 32)
		*(*int32)(unsafe.Add(mBase, uint32(v15+l7))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v15+l6))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(l5))) = v11 + int32(1)
		v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn13847[0])))
		if v24 == int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, _c_Fn13847[1]))
			if v29 <= int32(31) {
				*(*int32)(unsafe.Add(mBase, _c_Fn13847[1])) = v29 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v29<<(uint(int32(2))%32))+uint32(_c_Fn13847[2]))) = int32(1103)
			} else {
			}
			v46 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_Fn13847[0])) = uint8(v46)
		} else {
		}
		return
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				F_errmsg_internal(m, l4, int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_Fn13847_0), l3, l2)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
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
func Fn13852(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l5
			F_errmsg(m, l4, v9)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, l3, l2, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func Fn13858(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v13 = v12 - v8
	v16 = int32(0)
	if base.B2i32(base.B2i32(int64(0) < v8)^base.B2i32(v13 < v12) == v16)&base.B2i32(v13 != int64(-9223372036854775807-1)) == v16 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l3, int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l2, int32(107), l1)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v39 = v13 >> (uint(int64(63)) % 64)
		v42 = F_Int64GetDatum(m, v13^v39-v39)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			return v42
		}
	}
}
func Fn13861(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v37 int64
	_ = v37
	var v44 int64
	_ = v44
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v14 = int64(63)
	v16 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	v23 = int64(32)
	v24 = int64(base.Ui64(v16) >> (uint(v23) % 64))
	v26 = int64(base.Ui64(v13) >> (uint(v23) % 64))
	v29 = int64(4294967295)
	v30 = v16 & v29
	v32 = v13 & v29
	v33 = v30 * v32
	v37 = int64(base.Ui64(v33)>>(uint(v23)%64)) + v30*v26
	v44 = v32*v24 + v37&v29
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v13*(v16>>(uint(v14)%64)) + v13>>(uint(v14)%64)*v16 + v24*v26 + int64(base.Ui64(v37)>>(uint(v23)%64)) + int64(base.Ui64(v44)>>(uint(v23)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v33&v29 | v44<<(uint(v23)%64)
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if v55 != v56>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l4, int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v74 = F_Int64GetDatum(m, v56)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			m.G0 = v10 + int32(16)
			return v74
		}
	}
}
func Fn13870(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	if v11 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v23
				F_errmsg(m, l3, v8)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13870_0), l2, l1)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v8 + int32(16)
		return int32(0)
	}
}
func Fn13876(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v6)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	if v16&int32(1) == v6 {
		v21 = int32(4)
		v25 = l2 + l1<<(uint(v21)%32) + v21
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
		if v26 < int32(0) {
			v73 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				v79 = v73
				m.G0 = v11 + int32(16)
				return v79
			}
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v31 = v15 + v29 + v26
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+6)))
			if v32 != int32(1) {
				v79 = v31
				m.G0 = v11 + int32(16)
				return v79
			} else {
				v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+4)))
				switch v35&int32(_a_Fn13876_0) - int32(1) {
				case 0:
					v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31))))
					v79 = v40
					m.G0 = v11 + int32(16)
					return v79
				case 1:
					v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31))))
					v79 = v41
					m.G0 = v11 + int32(16)
					return v79
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v35
						F_errmsg_internal(m, int32(_a_Fn13876_1), v11)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l4, int32(70), int32(_a_Fn13876_2))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					v79 = v42
					m.G0 = v11 + int32(16)
					return v79
				}
			}
		}
	} else {
		v57 = int32(1)
		v58 = l1 - v57
		v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(base.Ui32(v58)>>(uint(int32(3))%32)))+23)))
		if int32(base.Ui32(v62)>>(uint(v58&int32(7))%32))&v57 != 0 {
			v73 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				v79 = v73
				m.G0 = v11 + int32(16)
				return v79
			}
		} else {
			v68 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v68)
			v79 = int32(0)
			m.G0 = v11 + int32(16)
			return v79
		}
	}
}
func Fn13883(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_num_same(m, v5, v6, l1, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v4))) = uint8(v8)
		return v4
	}
}
func Fn13885(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v22 = v11 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v25 = int32(4)
		v26 = v23 + v25
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v26
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v29 = int32(2)
		v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		if base.Ui32(v30+v25) < base.Ui32(int32(base.Ui32(v38)>>(uint(v29)%32))) {
			v42 = v26 + (v30+int32(3))&int32(2147483644)
		} else {
			v42 = v26
		}
		*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v42
		v44 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v44)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+16)))
		v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+v48)+12)))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v54 = F_gbt_var_consistent(m, v22, v15, v19, v46, v50&int32(1), l1, v53)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			m.G0 = v11 + int32(16)
			return v54
		}
	}
}
func Fn13889(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 float64
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v14 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v13 + v14
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v13
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+16)))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+v23)+12)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = F_gbt_num_distance(m, v8+v14, v8+int32(12), v25&int32(1), l1, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int32(0)
	} else {
		v33 = F_Float8GetDatum(m, v29)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v33
		}
	}
}
func Fn13896(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, l1, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+92))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func Fn13898(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, l1, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+68))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func Fn13906(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_palloc(m, int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v26 = F_palloc(m, int32(16))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v28
			v31 = F_palloc(m, v28)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v31
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = l2
				*(*uint8)(unsafe.Add(mBase, uint32(v26)+8)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v17
				*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v16)
				v42 = F_palloc(m, int32(4))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v42
					*(*int32)(unsafe.Add(mBase, uint32(v42))) = v26
					v47 = v16 & int32(_a_Fn13906_0)
					switch v47 - int32(1) {
					case 0, 1:
						v66 = F_Int64GetDatum(m, l1)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v66
							v69 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v69)
							m.G0 = v13 + int32(16)
							return v21
						}
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v17
						m.G0 = v13 + int32(16)
						return v21
					case 3, 4:
						v50 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v50)
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v17
						m.G0 = v13 + int32(16)
						return v21
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = v47
							F_errmsg_internal(m, int32(_a_Fn13906_1), v13)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13906_2), int32(97), int32(_a_Fn13906_3))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
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
		}
	}
}
func Fn13913(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v18)
		v21 = *(*int32)(unsafe.Add(mBase, _c_Fn13913[0]))
		v22 = F_convert_any_priv_string(m, v14, l1)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, l2, v12, v21, v22, v10+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v10 + int32(16)
				return v35
			}
		}
	}
}
func Fn13926(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	F_ean2isn(m, v9, v6+int32(8), l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
		v17 = F_Int64GetDatum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v17
		}
	}
}
func Fn13933(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v9
	v14 = F_pushJsonbValue(m, v7, l2, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v19 = F_pushJsonbValue(m, v7, l1, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v19
			v22 = F_JsonbValueToJsonb(m, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v22
			}
		}
	}
}
func Fn13935(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v10, v11, v12, l1, int32(7))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v20 = F_latin2mic(m, v8, v7, v12, l2, l1, base.B2i32(v9 != int32(0)))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			return v20
		}
	}
}
func Fn13940(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 float64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 float64
	_ = v56
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v64 int32
	_ = v64
	var v65 float64
	_ = v65
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v103 float64
	_ = v103
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v202 float64
	_ = v202
	var v204 float64
	_ = v204
	var v206 float64
	_ = v206
	var v219 float64
	_ = v219
	var v229 float64
	_ = v229
	var v240 float64
	_ = v240
	var v252 float64
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int64
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int64
	_ = v298
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int64
	_ = v324
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	v11 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	if l0 < int32(_a_Fn13940_0) {
		v336 = v11
		m.G0 = v18 + int32(96)
		return v336
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
		if v23 < int64(10000) {
			v336 = v11
			m.G0 = v18 + int32(96)
			return v336
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
			if v26 != int32(1) {
				v336 = v11
				m.G0 = v18 + int32(96)
				return v336
			} else {
				v30 = v22 + int32(16)
				v31 = int32(0)
				v38 = float64(0)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
				if v40 != 0 {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
					if v40 != int32(1) {
						v50 = v31
						v51 = v31
						v56 = v38
						for {
							v58 = float64(1)
							v59 = v50 + v41
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
							v61 = F_scalbn(m, v58, v60)
							mBase = m.M
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
							v65 = F_scalbn(m, v58, v64)
							mBase = m.M
							v70 = base.F64_add(base.F64_add(v56, base.F64_div(v58, v65)), base.F64_div(v58, v61))
							v71 = int32(2)
							v72 = v50 + v71
							v74 = v51 + v71
							if v74 != v40&int32(-2) {
								v50 = v72
								v51 = v74
								v56 = v70
								continue
							} else {
								break
							}
							break
						}
						if v40&int32(1) == int32(0) {
							v103 = v70
						} else {
							v80 = v72
							v86 = v70
							v88 = float64(1)
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v41))))
							v92 = F_scalbn(m, v88, v91)
							mBase = m.M
							v103 = base.F64_add(v86, base.F64_div(v88, v92))
						}
					} else {
						v80 = v31
						v86 = v38
						v88 = float64(1)
						v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v41))))
						v92 = F_scalbn(m, v88, v91)
						mBase = m.M
						v103 = base.F64_add(v86, base.F64_div(v88, v92))
					}
					v105 = *(*float64)(unsafe.Add(mBase, uint32(v30)+8))
					v106 = base.F64_div(v105, v103)
					v107 = base.F64_convert_i32_u(v40)
					if base.F64_le(v106, base.F64_mul(v107, float64(2.5))) == int32(0) {
						v219 = v106
						if base.F64_gt(v219, float64(1.4316557653333333e+08)) == int32(0) {
							v240 = v219
						} else {
							v229 = F_log(m, base.F64_add(base.F64_mul(v219, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v240 = base.F64_mul(v229, float64(-4.294967296e+09))
						}
						v252 = v240
					} else {
						v114 = v40 & int32(3)
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
						v116 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v40) {
							v125 = int32(0)
							v126 = v116
							v127 = v116
							for {
								v134 = v126 + v115
								v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
								v136 = int32(0)
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+2)))
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+3)))
								v150 = v127 + base.B2i32(v135 == v136) + base.B2i32(v139 == v136) + base.B2i32(v143 == v136) + base.B2i32(v147 == v136)
								v151 = int32(4)
								v152 = v126 + v151
								v154 = v125 + v151
								if v154 != v40&int32(-4) {
									v125 = v154
									v126 = v152
									v127 = v150
									continue
								} else {
									break
								}
								break
							}
							if v114 == int32(0) {
								v191 = v150
							} else {
								v160 = v152
								v161 = v150
								v170 = v160
								v171 = v161
								v174 = v116
								for {
									v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v115))))
									v182 = v171 + base.B2i32(v179 == int32(0))
									v183 = int32(1)
									v186 = v174 + v183
									if v186 != v114 {
										v170 = v170 + v183
										v171 = v182
										v174 = v186
										continue
									} else {
										break
									}
									break
								}
								v191 = v182
							}
						} else {
							v160 = v116
							v161 = v116
							v170 = v160
							v171 = v161
							v174 = v116
							for {
								v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v115))))
								v182 = v171 + base.B2i32(v179 == int32(0))
								v183 = int32(1)
								v186 = v174 + v183
								if v186 != v114 {
									v170 = v170 + v183
									v171 = v182
									v174 = v186
									continue
								} else {
									break
								}
								break
							}
							v191 = v182
						}
						if v191 == int32(0) {
							v240 = v106
							v252 = v240
						} else {
							v202 = F_log(m, base.F64_div(v107, base.F64_convert_i32_s(v191)))
							mBase = m.M
							v252 = base.F64_mul(v202, v107)
						}
					}
				} else {
					v204 = *(*float64)(unsafe.Add(mBase, uint32(v30)+8))
					v206 = base.F64_div(v204, float64(0))
					if base.F64_le(v206, base.F64_mul(base.F64_convert_i32_u(v40), float64(2.5))) != 0 {
						v240 = v206
					} else {
						v219 = v206
						if base.F64_gt(v219, float64(1.4316557653333333e+08)) == int32(0) {
							v240 = v219
						} else {
							v229 = F_log(m, base.F64_add(base.F64_mul(v219, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v240 = base.F64_mul(v229, float64(-4.294967296e+09))
						}
					}
					v252 = v240
				}
				if base.F64_gt(v252, float64(100000)) != 0 {
					v256 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn13940[0])))
					if v256 != int32(1) {
						v276 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v276)
						v336 = v11
						m.G0 = v18 + int32(96)
						return v336
					} else {
						v261 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v264 = m.ExcPending
						if v264 != 0 {
							return int32(0)
						} else {
							if v261 == int32(0) {
								v276 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v276)
								v336 = v11
								m.G0 = v18 + int32(96)
								return v336
							} else {
								v267 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
								*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l0
								*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v267
								*(*float64)(unsafe.Add(mBase, uint32(v18))) = v252
								F_errmsg_internal(m, l9, v18)
								mBase = m.M
								v272 = m.ExcPending
								if v272 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, l4, l8, l2)
									mBase = m.M
									v274 = m.ExcPending
									if v274 != 0 {
										return int32(0)
									} else {
										v276 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v276)
										v336 = v11
										m.G0 = v18 + int32(96)
										return v336
									}
								}
							}
						}
					}
				} else {
					v279 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn13940[0])))
					v280 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
					if base.F64_gt(base.F64_add(base.F64_div(base.F64_convert_i64_s(v280), float64(2000)), float64(0.5)), v252) != 0 {
						v287 = int32(1)
						if v279&v287 == int32(0) {
							v336 = v287
							m.G0 = v18 + int32(96)
							return v336
						} else {
							v294 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v295 = m.ExcPending
							if v295 != 0 {
								return int32(0)
							} else {
								if v294 == int32(0) {
									v336 = v287
									m.G0 = v18 + int32(96)
									return v336
								} else {
									v298 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v298
									*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v252
									*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = base.F64_add(base.F64_div(base.F64_convert_i64_s(v298), float64(2000)), float64(0.5))
									F_errmsg_internal(m, l7, v18+int32(32))
									mBase = m.M
									v311 = m.ExcPending
									if v311 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, l4, l6, l2)
										mBase = m.M
										v313 = m.ExcPending
										if v313 != 0 {
											return int32(0)
										} else {
											v336 = v287
											m.G0 = v18 + int32(96)
											return v336
										}
									}
								}
							}
						}
					} else {
						if v279&int32(1) == int32(0) {
							v336 = v11
							m.G0 = v18 + int32(96)
							return v336
						} else {
							v320 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v321 = m.ExcPending
							if v321 != 0 {
								return int32(0)
							} else {
								if v320 == int32(0) {
									v336 = v11
									m.G0 = v18 + int32(96)
									return v336
								} else {
									v324 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v324
									*(*float64)(unsafe.Add(mBase, uint32(v18)+64)) = v252
									F_errmsg_internal(m, l5, v18-int32(-64))
									mBase = m.M
									v331 = m.ExcPending
									if v331 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, l4, l3, l2)
										mBase = m.M
										v333 = m.ExcPending
										if v333 != 0 {
											return int32(0)
										} else {
											v336 = v11
											m.G0 = v18 + int32(96)
											return v336
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
func Fn13942(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v32 int32
	_ = v32
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l3
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
	if v11 == int32(1) {
		v14 = int32(_a_Fn13942_0)
		v15 = *(*int32)(unsafe.Add(mBase, _c_Fn13942[0]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		*(*int32)(unsafe.Add(mBase, _c_Fn13942[0])) = v17
		v20 = F_palloc(m, int32(40))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v24)
			*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(0)
			F_initHyperLogLog(m, v20+int32(16), int32(10))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(116)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v20
				*(*int32)(unsafe.Add(mBase, _c_Fn13942[0])) = v15
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func Fn13957(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(34209794)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v9
	v16 = *(*int32)(unsafe.Add(mBase, _c_Fn13957[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v16
	v19 = F_LockRelease(m, v7, l1, int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v19
	}
}
func Fn13960(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v4 - int32(142) {
	case 0:
		v15 = l1
		return v15
	case 1:
		return int32(3)
	default:
		if int32(0) <= base.I32_extend8_s(v4) {
			v14 = int32(1)
		} else {
			v14 = int32(2)
		}
		v15 = v14
		return v15
	}
}
func Fn13966(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(34209794)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v10
	v17 = *(*int32)(unsafe.Add(mBase, _c_Fn13966[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17
	v20 = F_LockAcquire(m, v8, l2, l1, int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return base.B2i32(v20 != int32(0))
	}
}
func Fn13971(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(176)
	m.G0 = v12
	v19 = v4
	v20 = int32(-1)
	v21 = v4
	v22 = v4
	goto L1
L1:
	;
	if v20 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn13971[0])) = v43
	*(*int32)(unsafe.Add(mBase, _c_Fn13971[1])) = v44
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v94 - int32(1)
	m.G0 = v12 + int32(176)
	return
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v27 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26 + v27
	v31 = *(*int32)(unsafe.Add(mBase, _c_Fn13971[0]))
	v33 = *(*int32)(unsafe.Add(mBase, _c_Fn13971[1]))
	v35 = v12 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v12 + int32(12)
	goto L6
L4:
	;
	v42 = v19
	v43 = v21
	v44 = v22
	goto L5
L5:
	;
	goto L8
L6:
	;
	v42 = int32(0)
	v43 = v31
	v44 = v33
	goto L5
L7:
	;
	goto L2
L8:
	;
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v69 = int32(m.ExcTag)
	v70 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v69 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_standard_ExecutorFinish(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L10
	} else {
		goto L18
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn13971[0])) = v12 + int32(16)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v51 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn13971[1])) = v44
	*(*int32)(unsafe.Add(mBase, _c_Fn13971[0])) = v43
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v60 - int32(1)
	F_pg_re_throw(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	m.T0[v51].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	goto L9
L19:
	;
	v74 = int32(v70)
	m.G0 = v12
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v12+int32(12) == v80 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	m.ExcPending = 1
	goto L28
L21:
	;
	if v84 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v84 = v82
	goto L24
L23:
	;
	v84 = int32(0)
	goto L24
L24:
	;
	goto L21
L25:
	;
	F___wasm_longjmp(m, v77, v76)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v19 = v76
	v20 = v84
	v21 = v43
	v22 = v44
	goto L1
L28:
	;
	return
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn13977(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v4
	v17 = F_query_or_expression_tree_walker_impl(m, l1, int32(896), v7+int32(4), v4)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		m.G0 = v7 + int32(16)
		return v21
	}
}
func Fn13986(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = l1
	return v4
}
func Fn13988(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = F_patternsel_common(m, v3, v4, v5, v6, v7, v8, l1, v5)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_Float8GetDatum(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func Fn13993(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v2 = l1
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v9 <= v6+int32(1) {
		F_appendStringInfoChar(m, v5, v2)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		*(*uint8)(unsafe.Add(mBase, uint32(v17+v6))) = uint8(v2)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
		v23 = v21 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v23
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v27 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v25+v23))) = uint8(v27)
		return v27
	}
}
func Fn13999(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v66 int64
	_ = v66
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	if base.Ui64(int64(2)) <= base.Ui64(v15-int64(9223372036854775807)) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
		if v21 != 0 {
			if v21 != int32(2147483647) {
				if v21 != int32(-2147483648) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v103 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
								F_errmsg(m, int32(_a_Fn13999_0), v12)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn13999_1), l2, l1)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
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
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
					if v26 != int32(-2147483648) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v103 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
									F_errmsg(m, int32(_a_Fn13999_0), v12)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_Fn13999_1), l2, l1)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
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
					} else {
						v29 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
						if v29 == int64(-9223372036854775807-1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									v49 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
										F_errmsg(m, int32(_a_Fn13999_2), v12+int32(16))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_Fn13999_1), l5, l1)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
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
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v103 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
										F_errmsg(m, int32(_a_Fn13999_0), v12)
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_Fn13999_1), l2, l1)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
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
					}
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
				if v32 != int32(2147483647) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v103 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
								F_errmsg(m, int32(_a_Fn13999_0), v12)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn13999_1), l2, l1)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
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
				} else {
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
					if v35 != int64(9223372036854775807) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v103 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
									F_errmsg(m, int32(_a_Fn13999_0), v12)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_Fn13999_1), l2, l1)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v49 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
									F_errmsg(m, int32(_a_Fn13999_2), v12+int32(16))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_Fn13999_1), l5, l1)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
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
				}
			}
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
			if v60 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						v103 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
							F_errmsg(m, int32(_a_Fn13999_0), v12)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13999_1), l2, l1)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
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
			} else {
				v61 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
				v62 = base.I64_div_s(v61, l4)
				v66 = base.I64_extend32_s(v62)*int64(-1000000) + v15
				if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v66+int64(211813488000000000)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_Fn13999_3), int32(0))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13999_1), l3, l1)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v73 = v66
					v74 = F_Int64GetDatum(m, v73)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						m.G0 = v12 + int32(32)
						return v74
					}
				}
			}
		}
	} else {
		v73 = v15
		v74 = F_Int64GetDatum(m, v73)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			m.G0 = v12 + int32(32)
			return v74
		}
	}
}
func Fn14002(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) <= v8 {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
		v14 = F_psprintf(m, int32(_a_Fn14002_0), v6)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = v14
			m.G0 = v6 + int32(16)
			return v20
		}
	} else {
		v18 = F_pstrdup(m, l1)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = v18
			m.G0 = v6 + int32(16)
			return v20
		}
	}
}
func Fn14017(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_WinGetFuncArgInFrame(m, v9, int32(0), l1, int32(1), v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v24 = int32(0)
		} else {
			v24 = v14
		}
		m.G0 = v7 + int32(16)
		return v24
	}
}
func Fn14020(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 float32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 float64
	_ = v98
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = int32(1)
			v21 = v16 + v20
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			v24 = v22 & v20
			if v24 != 0 {
				v25 = v21
			} else {
				v25 = v16 + int32(4)
			}
			if v22 == int32(1) {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				if v31 == int32(18) {
					v34 = int32(16)
				} else {
					v34 = int32(0)
				}
				if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v41 = int32(4)
				} else {
					v41 = v34
				}
				v52 = v41
			} else {
				v42 = int32(1)
				if v24 != 0 {
					v52 = int32(base.Ui32(v22)>>(uint(v42)%32)) - v42
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v53 = int32(1)
			v54 = v11 + v53
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			v59 = v57 & v53
			if v59 != 0 {
				v60 = v54
			} else {
				v60 = v11 + int32(4)
			}
			if v57 == int32(1) {
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				if v66 == int32(18) {
					v69 = int32(16)
				} else {
					v69 = int32(0)
				}
				if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v76 = int32(4)
				} else {
					v76 = v69
				}
				v87 = v76
			} else {
				v77 = int32(1)
				if v59 != 0 {
					v87 = int32(base.Ui32(v57)>>(uint(v77)%32)) - v77
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v87 = int32(base.Ui32(v81)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v88 = F_calc_word_similarity(m, v25, v52, v60, v87, l2)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v90 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v94 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								return base.F64_le(v98, base.F64_promote_f32(v88))
							}
						} else {
							v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							return base.F64_le(v98, base.F64_promote_f32(v88))
						}
					}
				} else {
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v94 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							return base.F64_le(v98, base.F64_promote_f32(v88))
						}
					} else {
						v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						return base.F64_le(v98, base.F64_promote_f32(v88))
					}
				}
			}
		}
	}
}
