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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v135 int32
	_ = v135
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	if v14 <= v4 {
		if v12 <= int32(0) {
			F_populate_array_report_expected_array(m, v13, v12)
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int32(0)
			} else {
				v135 = int32(23)
				return v135
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v12
			v21 = v12 << (uint(int32(2)) % 32)
			v22 = F_palloc(m, v21)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v22
				v27 = F_palloc0(m, v21)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v27
					v31 = v12 & int32(3)
					v32 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v12) {
						v40 = v32
						v45 = v4
						for {
							v48 = v40 << (uint(int32(2)) % 32)
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
							v51 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v48+v49))) = v51
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v53+v48)+4)) = v51
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v57+v48)+8)) = v51
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v61+v48)+12)) = v51
							v65 = int32(4)
							v66 = v40 + v65
							v68 = v45 + v65
							if v68 != v12&int32(2147483644) {
								v40 = v66
								v45 = v68
								continue
							} else {
								break
							}
							break
						}
						if v31 == int32(0) {
						} else {
							v75 = v66
							v85 = v75
							v91 = v4
							for {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v92+v85<<(uint(int32(2))%32)))) = int32(-1)
								v98 = int32(1)
								v101 = v91 + v98
								if v101 != v31 {
									v85 = v85 + v98
									v91 = v101
									continue
								} else {
									break
								}
								break
							}
						}
					} else {
						v75 = v32
						v85 = v75
						v91 = v4
						for {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v92+v85<<(uint(int32(2))%32)))) = int32(-1)
							v98 = int32(1)
							v101 = v91 + v98
							if v101 != v31 {
								v85 = v85 + v98
								v91 = v101
								continue
							} else {
								break
							}
							break
						}
					}
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
					v121 = v113
					if v121 != v12 {
						v135 = int32(0)
						return v135
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
						return int32(0)
					}
				}
			}
		}
	} else {
		if v14 <= v12 {
			v121 = v14
			if v121 != v12 {
				v135 = int32(0)
				return v135
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
				return int32(0)
			}
		} else {
			F_populate_array_report_expected_array(m, v13, v12)
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int32(0)
			} else {
				v135 = int32(23)
				return v135
			}
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
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
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+26)) = uint8(v36)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+104)))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+27)) = uint8(base.B2i32(v38 != int32(112)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+108)))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)) = uint8(v42)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+111)))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+29)) = uint8(v44)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+110)))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+30)) = uint8(base.B2i32(v46 != v3))
	v51 = v33 + v24
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+106)))
	if v52 == int32(1) {
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
		if base.Ui32(v57) < base.Ui32(int32(_a_F_populate_compact_attribute_0)) {
			v60 = int32(118)
		} else {
			v60 = int32(117)
		}
		v62 = v60
	} else {
		v62 = int32(102)
	}
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+11)) = uint8(v62)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+83)))
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
			v75 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51)+83)))
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v75
			F_errmsg_internal(m, int32(_a_F_populate_compact_attribute_1), v11)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_populate_compact_attribute_2), int32(105), int32(_a_F_populate_compact_attribute_3))
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
				F_errmsg(m, int32(_a_F_populate_recordset_scalar_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_populate_recordset_scalar_1), int32(_a_F_populate_recordset_scalar_2), int32(_a_F_populate_recordset_scalar_3))
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
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
	v15 = int32(_a_F_populate_typ_list_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_populate_typ_list[0]))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_populate_typ_list[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_populate_typ_list[0])) = v19
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
	*(*int32)(unsafe.Add(mBase, _c_F_populate_typ_list[0])) = v16
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	m.T0[v59].(func(*base.Module, int32))(m, v13)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L14
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
	v34 = int32(_a_F_populate_typ_list_1)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_populate_typ_list[2]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_populate_typ_list[2])) = v37
	v40 = v29 + v30
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v41
	base.MemoryCopy(m, v32+int32(4), v40, int32(148))
	v47 = F_heap_getnext(m, v13)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v47 != 0 {
		v23 = v47
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	F_relation_close(m, v9, int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	return
}
