package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_domain_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	v2 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v6 == v2 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = v9
	} else {
		v10 = v2
	}
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v11 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
		if v16 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v17 == v14 {
				v27 = v16
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
				v32 = F_ReceiveFunctionCall(m, v27+int32(16), v10, v30, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = int32(0)
					F_domain_check_input(m, v32, base.B2i32(v10 == v34), v27, v34)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v10 != 0 {
							v46 = v32
						} else {
							v42 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v42)
							v46 = int32(0)
						}
						return v46
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
				v21 = F_domain_state_setup(m, v14, int32(1), v20)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v21
					v27 = v21
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
					v32 = F_ReceiveFunctionCall(m, v27+int32(16), v10, v30, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = int32(0)
						F_domain_check_input(m, v32, base.B2i32(v10 == v34), v27, v34)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v10 != 0 {
								v46 = v32
							} else {
								v42 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v42)
								v46 = int32(0)
							}
							return v46
						}
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
			v21 = F_domain_state_setup(m, v14, int32(1), v20)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v21
				v27 = v21
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
				v32 = F_ReceiveFunctionCall(m, v27+int32(16), v10, v30, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = int32(0)
					F_domain_check_input(m, v32, base.B2i32(v10 == v34), v27, v34)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v10 != 0 {
							v46 = v32
						} else {
							v42 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v42)
							v46 = int32(0)
						}
						return v46
					}
				}
			}
		}
	} else {
		v42 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v42)
		v46 = int32(0)
		return v46
	}
}
func F_replace_domain_constraint_value(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v5 == v3 {
		v47 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v47
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v8 != int32(1) {
		v47 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = int32(334839)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[510])))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v18 == int32(0) {
		v37 = v17
		v38 = v18
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v38-v37 != 0 {
		v47 = v3
		goto L1
	} else {
		goto L12
	}
L5:
	;
	goto L4
L6:
	;
	if v17 != v18 {
		v37 = v17
		v38 = v18
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v22 = v13
	v23 = v14
	goto L8
L8:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v27 == int32(0) {
		v37 = v26
		v38 = v27
		goto L5
	} else {
		goto L10
	}
L9:
	;
	v37 = v26
	v38 = v27
	goto L5
L10:
	;
	v30 = int32(1)
	if v26 == v27 {
		v22 = v22 + v30
		v23 = v23 + v30
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v41 = F_copyObjectImpl(m, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v45
	v47 = v41
	goto L1
}
func F_validateDomainNotNullConstraint(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = F_get_rels_with_domain(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L3
	} else {
		goto L47
	}
L2:
	;
	m.G0 = v15 + int32(16)
	return
L3:
	;
	return
L4:
	;
	if v17 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v21 <= int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v31 = int32(0)
	goto L7
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v31<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	v43 = F_GetLatestSnapshot(m)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L1
L9:
	;
	v45 = F_RegisterSnapshot(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v47 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v41)+188))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v53 = m.T0[v52].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v41, v45, v47, v47, v47, int32(449))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v56 = F_table_slot_create(m, v41, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+36)) = v59
	v62 = *(*int32)(unsafe.Add(mBase, _consts[452]))
	if v62 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[453])))
	if v64&int32(1) == int32(0) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L17
L16:
	;
	goto L15
L17:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+188))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	v85 = m.T0[v84].(func(*base.Module, int32, int32, int32) int32)(m, v53, int32(1), v56)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L3
	} else {
		goto L20
	}
L18:
	;
	goto L8
L19:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+36)) = v183
	v186 = *(*int32)(unsafe.Add(mBase, _consts[452]))
	if v186 == int32(0) {
		goto L17
	} else {
		goto L45
	}
L20:
	;
	if v85 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v87 = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v88 <= v87 {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_ExecDropSingleTupleTableSlot(m, v56)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L40
	}
L24:
	;
	v91 = v87
	goto L25
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v91<<(uint(int32(2))%32))))
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+6)))
	if v109 < v108 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L3
	} else {
		goto L35
	}
L27:
	;
	F_slot_getsomeattrs_int(m, v56, v108)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	v115 = int32(1)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v108-v115))))
	if v117 != v115 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v121 = v91 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v122 <= v121 {
		goto L19
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L26
L34:
	;
	v91 = v121
	goto L25
L35:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	v132 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v131 + v132
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v42 + v103<<(uint(v132)%32) + v108*int32(100) - int32(76)
	F_errmsg(m, int32(151143), v15)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	F_errtablecol(m, v41, v108)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(477716), int32(3178), int32(87441))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+188))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	m.T0[v158].(func(*base.Module, int32))(m, v53)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_UnregisterSnapshot(m, v45)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	F_sequence_close(m, v41, int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v167 = v31 + int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v167 < v168 {
		v31 = v167
		goto L7
	} else {
		goto L44
	}
L44:
	;
	goto L2
L45:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, _consts[453])))
	if v190&int32(1) != 0 {
		goto L17
	} else {
		goto L46
	}
L46:
	;
	goto L18
L47:
	;
	F_errmsg_internal(m, int32(323972), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(314654), int32(1034), int32(81516))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
