package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_populate_array_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v13 <= v4 {
		if v11 <= int32(0) {
			F_populate_array_report_expected_array(m, v12, v11)
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return int32(0)
			} else {
				return int32(23)
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v11
			v20 = v11 << (uint(int32(2)) % 32)
			v21 = F_palloc(m, v20)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v21
				v26 = F_palloc0(m, v20)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v26
					v29 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v11) {
						v36 = v29
						v41 = v4
						for {
							v44 = v36 << (uint(int32(2)) % 32)
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
							v47 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v44+v45))) = v47
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v49+v44)+4)) = v47
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v53+v44)+8)) = v47
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v57+v44)+12)) = v47
							v61 = int32(4)
							v62 = v36 + v61
							v64 = v41 + v61
							if v64 != v11&int32(2147483644) {
								v36 = v62
								v41 = v64
								continue
							} else {
								break
							}
							break
						}
						v68 = v62
					} else {
						v68 = v29
					}
					v76 = v11 & int32(3)
					if v76 != 0 {
						v79 = v68
						v83 = v4
						for {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v86+v79<<(uint(int32(2))%32)))) = int32(-1)
							v92 = int32(1)
							v95 = v83 + v92
							if v95 != v76 {
								v79 = v79 + v92
								v83 = v95
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
					v110 = v106
					if v110 == v11 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
					} else {
					}
					return int32(0)
				}
			}
		}
	} else {
		if v11 < v13 {
			F_populate_array_report_expected_array(m, v12, v11)
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return int32(0)
			} else {
				return int32(23)
			}
		} else {
			v110 = v13
			if v110 == v11 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
			} else {
			}
			return int32(0)
		}
	}
}
func F_populate_compact_attribute(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(4)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = l0 + l1<<(uint(v13)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v3
	v21 = v17 + int32(24)
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = int64(0)
	v24 = int32(20)
	v25 = v17 + v24
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(-1)
	v33 = l0 + v14<<(uint(v13)%32) + l1*int32(100)
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+92)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21))) = uint16(v34)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+26)) = uint8(v37)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+104)))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+27)) = uint8(base.B2i32(v39 != int32(112)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+108)))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)) = uint8(v43)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+111)))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+29)) = uint8(v45)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+110)))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+30)) = uint8(base.B2i32(v47 != v3))
	v52 = v33 + v24
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+106)))
	if v53 == int32(1) {
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
		if base.Ui32(v58) < base.Ui32(int32(12000)) {
			v61 = int32(118)
		} else {
			v61 = int32(117)
		}
		v62 = v61
	} else {
		v62 = int32(102)
	}
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+11)) = uint8(v62)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+83)))
	switch v64 - int32(99) {
	case 0:
		v86 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)) = uint8(v86)
		m.G0 = v11 + int32(16)
		return
	case 1:
		v86 = int32(8)
		*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)) = uint8(v86)
		m.G0 = v11 + int32(16)
		return
	default:
		v69 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)) = uint8(v69)
		F_errstart_cold(m, int32(21), v69)
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return
		} else {
			v75 = int32(*(*int8)(unsafe.Add(mBase, uint32(v52)+83)))
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v75
			F_errmsg_internal(m, int32(483084), v11)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return
			} else {
				F_errfinish(m, int32(480491), int32(105), int32(299341))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 6:
		v86 = v13
		*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)) = uint8(v86)
		m.G0 = v11 + int32(16)
		return
	case 16:
		v86 = int32(2)
		*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)) = uint8(v86)
		m.G0 = v11 + int32(16)
		return
	}
}
func F_populate_recordset_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	switch v9 {
	case 0:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v19
				F_errmsg(m, int32(220118), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(475576), int32(4298), int32(219856))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	default:
		m.G0 = v6 + int32(16)
		return int32(0)
	case 2:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
		m.G0 = v6 + int32(16)
		return int32(0)
	}
}
func F_populate_typ_list(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	v9 = F_table_open(m, int32(1247), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = int32(0)
	v13 = F_table_beginscan_catalog(m, v9, v11, v11)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(4449520)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v19 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v19
	v21 = F_heap_getnext(m, v13)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = v21
	goto L8
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+188))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	m.T0[v60].(func(*base.Module, int32))(m, v13)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
	v32 = F_palloc(m, int32(152))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v34 = int32(4346024)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[416]))
	v37 = F_lappend(m, v36, v32)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[416])) = v37
	v40 = v29 + v30
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v41
	goto L13
L12:
	;
	v48 = F_heap_getnext(m, v13)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v46 = F__emscripten_memcpy_bulkmem(m, v32+int32(4), v40, int32(148))
	mBase = m.M
	goto L15
L15:
	;
	goto L12
L16:
	;
	if v48 != 0 {
		v23 = v48
		goto L8
	} else {
		goto L17
	}
L17:
	;
	goto L9
L18:
	;
	F_sequence_close(m, v9, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	return
}
