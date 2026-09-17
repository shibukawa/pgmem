package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_seg_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 float32
	_ = v31
	var v32 float32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 float32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_palloc(m, int32(40))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	switch v17 - int32(60) {
	case 0, 2:
		goto L4
	case 1:
		v29 = v13
		goto L3
	default:
		goto L5
	}
L3:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	v31 = *(*float32)(unsafe.Add(mBase, uint32(v11)))
	v32 = *(*float32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.F32_ne(v31, v32) != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
	v26 = F_pg_sprintf(m, v13, int32(_a_F_seg_out_0), v9+int32(16))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	if v17 != int32(126) {
		v29 = v13
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v29 = v26 + v13
	goto L3
L8:
	;
	m.G0 = v9 + int32(32)
	return v13
L9:
	;
	if v30 != int32(45) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	if v34 != v30 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+8)))
	v37 = F_restore(m, v29, v31, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+8)))
	v42 = F_restore(m, v29, v31, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v50 = v29
	goto L15
L15:
	;
	v53 = F_pg_sprintf(m, v50, int32(_a_F_seg_out_2), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v44 = v42 + v29
	v47 = F_pg_sprintf(m, v44, int32(_a_F_seg_out_1), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v50 = v47 + v44
	goto L15
L18:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	if v55 == int32(45) {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v58 = v50 + v53
	v61 = F_pg_sprintf(m, v58, int32(_a_F_seg_out_1), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v63 = v61 + v58
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	switch v64 - int32(60) {
	case 0, 2:
		goto L22
	default:
		goto L23
	}
L21:
	;
	v77 = *(*float32)(unsafe.Add(mBase, uint32(v11)+4))
	v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+9)))
	v79 = F_restore(m, v76, v77, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L26
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = base.I32_extend8_s(v64)
	v73 = F_pg_sprintf(m, v63, int32(_a_F_seg_out_0), v9)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	if v67 != int32(126) {
		v76 = v63
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v76 = v73 + v63
	goto L21
L26:
	;
	goto L8
}
func F_seg_right(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float32)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(v4)+4))
	return base.F32_gt(v3, v5)
}
func F_seg_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_seg_same_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int32(0))
	}
}
func F_seg_yy_scan_string(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	v3 = int32(0)
	v12 = F_strlen(m, l0)
	mBase = m.M
	v14 = v12 + int32(2)
	v15 = F_palloc(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 != 0 {
			if v12 <= int32(0) {
			} else {
				v22 = v12 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v12) {
					v29 = v3
					v36 = v3
					for {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v29))))
						*(*uint8)(unsafe.Add(mBase, uint32(v29+v15))) = uint8(v40)
						v43 = v29 | int32(1)
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v43))))
						*(*uint8)(unsafe.Add(mBase, uint32(v15+v43))) = uint8(v46)
						v49 = v29 | int32(2)
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v49))))
						*(*uint8)(unsafe.Add(mBase, uint32(v15+v49))) = uint8(v52)
						v55 = v29 | int32(3)
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v55))))
						*(*uint8)(unsafe.Add(mBase, uint32(v15+v55))) = uint8(v58)
						v60 = int32(4)
						v61 = v29 + v60
						v63 = v36 + v60
						if v63 != v12&int32(2147483644) {
							v29 = v61
							v36 = v63
							continue
						} else {
							break
						}
						break
					}
					if v22 == int32(0) {
					} else {
						v69 = v61
						v80 = v69
						v88 = v3
						for {
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v80))))
							*(*uint8)(unsafe.Add(mBase, uint32(v80+v15))) = uint8(v91)
							v93 = int32(1)
							v96 = v88 + v93
							if v96 != v22 {
								v80 = v80 + v93
								v88 = v96
								continue
							} else {
								break
							}
							break
						}
					}
				} else {
					v69 = v3
					v80 = v69
					v88 = v3
					for {
						v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v80))))
						*(*uint8)(unsafe.Add(mBase, uint32(v80+v15))) = uint8(v91)
						v93 = int32(1)
						v96 = v88 + v93
						if v96 != v22 {
							v80 = v80 + v93
							v88 = v96
							continue
						} else {
							break
						}
						break
					}
				}
			}
			v110 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v15+v12))) = uint16(v110)
			v112 = F_seg_yy_scan_buffer(m, v15, v14, l1)
			mBase = m.M
			v113 = m.ExcPending
			if v113 != 0 {
				return int32(0)
			} else {
				if v112 == int32(0) {
					F_yy_fatal_error_7(m, int32(_a_F_seg_yy_scan_string_0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = int32(1)
					return v112
				}
			}
		} else {
			F_yy_fatal_error_7(m, int32(_a_F_seg_yy_scan_string_1))
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_seg_yy_switch_to_buffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	F_seg_yyensure_buffer_stack(m, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		if v8 == int32(0) {
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8+v11<<(uint(int32(2))%32))))
			if v15 == l0 {
			} else {
				if v15 != 0 {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
					*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v18)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v20+v21<<(uint(int32(2))%32))))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v26
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v28
					v30 = v20
					v32 = v21
				} else {
					v30 = v8
					v32 = v11
				}
				*(*int32)(unsafe.Add(mBase, uint32(v30+v32<<(uint(int32(2))%32)))) = l0
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v39
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v42
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v44)
			}
		}
		return
	}
}
