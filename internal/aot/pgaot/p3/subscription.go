package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetSubscription(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(67), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v153
L2:
	;
	return int32(0)
L3:
	;
	if v12 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		v153 = int32(0)
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	v34 = F_palloc(m, int32(56))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L11
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg_internal(m, int32(_a_F_GetSubscription_0), v9)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_GetSubscription_1), int32(87), int32(_a_F_GetSubscription_2))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = l0
	v37 = v31 + v32
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v38
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v40
	v44 = F_pstrdup(m, v37+int32(16))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v44
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v47
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+84)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+25)) = uint8(v49)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+85)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+26)) = uint8(v51)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+86)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+27)) = uint8(v53)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+87)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+28)) = uint8(v55)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+88)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+29)) = uint8(v57)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+89)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+30)) = uint8(v59)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+90)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+31)) = uint8(v61)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+91)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+32)) = uint8(v63)
	v67 = F_SysCacheGetAttrNotNull(m, int32(67), v12, int32(14))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v69 = F_text_to_cstring(m, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v69
	v72 = int32(0)
	v77 = F_SysCacheGetAttr(m, int32(67), v12, int32(15), v9+int32(7))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)))
	if v80 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v84 = int32(0)
	goto L18
L17:
	;
	v82 = F_pstrdup(m, v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v84
	v88 = F_SysCacheGetAttrNotNull(m, int32(67), v12, int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L20
	}
L19:
	;
	v84 = v82
	goto L18
L20:
	;
	v90 = F_text_to_cstring(m, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v90
	v95 = F_SysCacheGetAttrNotNull(m, int32(67), v12, int32(17))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v97 = F_pg_detoast_datum(m, v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	F_deconstruct_array_builtin(m, v97, int32(25), v9+int32(12), int32(0), v9+int32(8))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if int32(0) < v107 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v110 = int32(0)
	v111 = v72
	goto L28
L26:
	;
	v132 = v72
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v132
	v140 = F_SysCacheGetAttrNotNull(m, int32(67), v12, int32(18))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L34
	}
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+v110<<(uint(int32(2))%32))))
	v121 = F_text_to_cstring(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L30
	}
L29:
	;
	v132 = v125
	goto L27
L30:
	;
	v123 = F_makeString(m, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v125 = F_lappend(m, v111, v123)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v128 = v110 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v128 < v129 {
		v110 = v128
		v111 = v125
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v142 = F_text_to_cstring(m, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v142
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v146 = F_superuser_arg(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)) = uint8(v146)
	F_ReleaseCatCache(m, v12)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v153 = v34
	goto L1
}
func F_GetSubscriptionRelState(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_table_open(m, int32(_a_F_GetSubscriptionRelState_0), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = F_SearchSysCache2(m, int32(68), l1, l0)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 == int32(0) {
				F_relation_close(m, v13, int32(1))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
					v48 = int32(0)
					m.G0 = v9 + int32(16)
					return base.I32_extend8_s(v48)
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v29)+8)))
				v36 = F_SysCacheGetAttr(m, int32(68), v18, int32(4), v9+int32(15))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v38 != 0 {
						v41 = int64(0)
					} else {
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
						v41 = v40
					}
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = v41
					F_ReleaseCatCache(m, v18)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_relation_close(m, v13, int32(1))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = v31
							m.G0 = v9 + int32(16)
							return base.I32_extend8_s(v48)
						}
					}
				}
			}
		}
	}
}
