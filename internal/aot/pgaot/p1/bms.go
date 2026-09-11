package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bms_add_member(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	if int32(0) <= l1 {
		v9 = l1 & int32(31)
		v11 = int32(base.Ui32(l1) >> (uint(int32(5)) % 32))
		if l0 == int32(0) {
			v15 = v11 + int32(1)
			v20 = F_palloc0(m, v15<<(uint(int32(2))%32)+int32(8))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v15
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(445)
				*(*int32)(unsafe.Add(mBase, uint32(v20+v11<<(uint(int32(2))%32))+8)) = int32(1) << (uint(v9) % 32)
				return v20
			}
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v34 <= v11 {
				v37 = v11 + int32(1)
				v42 = F_repalloc(m, l0, v37<<(uint(int32(2))%32)+int32(8))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v37
					v48 = v34
					for {
						*(*int32)(unsafe.Add(mBase, uint32(v42+int32(8)+v48<<(uint(int32(2))%32)))) = int32(0)
						v58 = v48 + int32(1)
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
						if v58 < v59 {
							v48 = v58
							continue
						} else {
							break
						}
						break
					}
					v61 = v42
					v70 = v61 + v11<<(uint(int32(2))%32) + int32(8)
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
					*(*int32)(unsafe.Add(mBase, uint32(v70))) = v71 | int32(1)<<(uint(v9)%32)
					return v61
				}
			} else {
				v61 = l0
				v70 = v61 + v11<<(uint(int32(2))%32) + int32(8)
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
				*(*int32)(unsafe.Add(mBase, uint32(v70))) = v71 | int32(1)<<(uint(v9)%32)
				return v61
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v80 = m.ExcPending
		if v80 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_bms_add_member_0), int32(0))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_bms_add_member_1), int32(823), int32(_a_F_bms_add_member_2))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
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
func F_bms_del_member(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	if int32(0) <= l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L16
	} else {
		goto L18
	}
L4:
	;
	return v55
L5:
	;
	v10 = int32(base.Ui32(l1) >> (uint(int32(5)) % 32))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 <= v10 {
		v55 = l0
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v55 = int32(0)
	goto L4
L8:
	;
	v14 = l0 + int32(8)
	v17 = v14 + v10<<(uint(int32(2))%32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v21 = v18 & base.I32_rotl(int32(-2), l1)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v21
	if v21 != 0 {
		v55 = l0
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v10 != v11-int32(1) {
		v55 = l0
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v28 = v10
	goto L12
L11:
	;
	F_pfree(m, l0)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	if v28 <= int32(0) {
		goto L11
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28
	return l0
L14:
	;
	v35 = v28 - int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14+v35<<(uint(int32(2))%32))))
	if v39 == int32(0) {
		v28 = v35
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	return int32(0)
L17:
	;
	goto L7
L18:
	;
	F_errmsg_internal(m, int32(_a_F_bms_del_member_0), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_bms_del_member_1), int32(876), int32(_a_F_bms_del_member_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bms_free(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	if l0 != 0 {
		F_pfree(m, l0)
		v3 = m.ExcPending
		if v3 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_bms_join(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v153 int32
	_ = v153
	v3 = int32(0)
	if l0 == v3 {
		return l1
	} else {
		if l1 == int32(0) {
			return l0
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v24 = base.B2i32(v22 < v23)
			if v22 < v23 {
				v25 = l0
			} else {
				v25 = l1
			}
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			if v26 <= int32(1) {
				v29 = int32(1)
			} else {
				v29 = v26
			}
			v31 = v29 & int32(3)
			if v22 < v23 {
				v32 = l1
			} else {
				v32 = l0
			}
			v33 = int32(8)
			v34 = v32 + v33
			v36 = v25 + v33
			v37 = int32(0)
			if int32(4) <= v26 {
				v46 = v37
				v48 = int32(0)
				for {
					v58 = v46 << (uint(int32(2)) % 32)
					v59 = v34 + v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v36)))
					*(*int32)(unsafe.Add(mBase, uint32(v59))) = v60 | v62
					v65 = int32(4)
					v66 = v58 | v65
					v67 = v34 + v66
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v36)))
					*(*int32)(unsafe.Add(mBase, uint32(v67))) = v68 | v70
					v74 = v58 | int32(8)
					v75 = v34 + v74
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v36)))
					*(*int32)(unsafe.Add(mBase, uint32(v75))) = v76 | v78
					v82 = v58 | int32(12)
					v83 = v34 + v82
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v36)))
					*(*int32)(unsafe.Add(mBase, uint32(v83))) = v84 | v86
					v90 = v46 + v65
					v92 = v48 + v65
					if v92 != v29&int32(2147483644) {
						v46 = v90
						v48 = v92
						continue
					} else {
						break
					}
					break
				}
				v97 = v90
			} else {
				v97 = v37
			}
			if v31 != 0 {
				v111 = v97
				v119 = v3
				for {
					v123 = v111 << (uint(int32(2)) % 32)
					v124 = v34 + v123
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
					v127 = *(*int32)(unsafe.Add(mBase, uint32(v123+v36)))
					*(*int32)(unsafe.Add(mBase, uint32(v124))) = v125 | v127
					v130 = int32(1)
					v133 = v119 + v130
					if v133 != v31 {
						v111 = v111 + v130
						v119 = v133
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			if l0 != l1 {
				F_pfree(m, v25)
				mBase = m.M
				v153 = m.ExcPending
				if v153 != 0 {
					return int32(0)
				} else {
					return v32
				}
			} else {
				return v32
			}
		}
	}
}
func F_bms_next_member(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.I32_ctz(v45) | v46<<(uint(int32(5))%32)
L2:
	;
	return int32(-2)
L3:
	;
	v9 = l1 + int32(1)
	v11 = base.I32_div_s(v9, int32(32))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 <= v11 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v15 = l0 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15+v11<<(uint(int32(2))%32))))
	v22 = v19 & (int32(-1) << (uint(v9) % 32))
	if v22 != 0 {
		v45 = v22
		v46 = v11
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v24 = v11 + int32(1)
	if v24 == v12 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v27 = v24
	goto L7
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v15+v27<<(uint(int32(2))%32))))
	if v34 != 0 {
		v45 = v34
		v46 = v27
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L2
L9:
	;
	v36 = v27 + int32(1)
	if v36 != v12 {
		v27 = v36
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
}
func F_bms_singleton_member(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L16
	} else {
		goto L20
	}
L2:
	;
	v6 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 <= v6 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v10 = v6
	goto L7
L6:
	;
	v10 = v7
	goto L7
L7:
	;
	v15 = int32(0)
	v17 = int32(-1)
	goto L8
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v15<<(uint(int32(2))%32))))
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return v33
L10:
	;
	if int32(0) <= v17 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v33 = v17
	goto L12
L12:
	;
	v35 = v15 + int32(1)
	if v35 != v10 {
		v15 = v35
		v17 = v33
		goto L8
	} else {
		goto L15
	}
L13:
	;
	if base.Ui32(int32(2)) <= base.Ui32(base.I32_popcnt(v23)) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v33 = base.I32_ctz(v23) | v15<<(uint(int32(5))%32)
	goto L12
L15:
	;
	goto L9
L16:
	;
	return int32(0)
L17:
	;
	F_errmsg_internal(m, int32(_a_F_bms_singleton_member_0), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_bms_singleton_member_1), int32(681), int32(_a_F_bms_singleton_member_2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F_errmsg_internal(m, int32(_a_F_bms_singleton_member_3), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_bms_singleton_member_1), int32(692), int32(_a_F_bms_singleton_member_2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
