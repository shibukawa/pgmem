package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ChangeVarNodes_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v3 = int32(0)
	if l0 == v3 {
		v157 = v3
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v172
	return int32(0)
L2:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v159 + int32(1)
	v165 = F_query_tree_walker_impl(m, l0, int32(1127), l1, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L20
	} else {
		goto L68
	}
L3:
	;
	return v157
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v9 - int32(58) {
	case 0:
		goto L10
	case 1, 2, 3, 4:
		v112 = v9
		goto L6
	case 5:
		goto L9
	case 6:
		goto L8
	default:
		goto L11
	}
L5:
	;
	v151 = F_expression_tree_walker_impl(m, l0, int32(1127), l1)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L20
	} else {
		goto L67
	}
L6:
	;
	if v112 == int32(67) {
		goto L2
	} else {
		goto L52
	}
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v68 != v69 {
		goto L5
	} else {
		goto L35
	}
L8:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v62 != 0 {
		goto L5
	} else {
		goto L33
	}
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v58 != 0 {
		v157 = v3
		goto L3
	} else {
		goto L31
	}
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v54 != 0 {
		v157 = v3
		goto L3
	} else {
		goto L29
	}
L11:
	;
	if v9 == int32(321) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v9 != int32(6) {
		v112 = v9
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v16 != v17 {
		v157 = v3
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 == v21 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v25 = v24
	goto L17
L16:
	;
	v25 = v20
	goto L17
L17:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v25 < int32(0) {
		v43 = v26
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v45 != v46 {
		v157 = v3
		goto L3
	} else {
		goto L27
	}
L19:
	;
	v29 = F_bms_is_member(m, v25, v26)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	if v29 == int32(0) {
		v43 = v26
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v35 = F_bms_copy(m, v26)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v37 = F_bms_del_member(m, v35, v25)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	if v19 < int32(0) {
		v43 = v37
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v41 = F_bms_add_member(m, v37, v19)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v43 = v41
	goto L18
L27:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v48 == int32(-5) {
		v157 = v3
		goto L3
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v48
	return int32(0)
L29:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v55 != v56 {
		v157 = v3
		goto L3
	} else {
		goto L30
	}
L30:
	;
	goto L1
L31:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v59 != v60 {
		v157 = v3
		goto L3
	} else {
		goto L32
	}
L32:
	;
	goto L1
L33:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v63 != v64 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v66
	goto L5
L35:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v72 < int32(0) {
		v88 = v71
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v88
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v92 < int32(0) {
		v108 = v91
		goto L44
	} else {
		goto L45
	}
L37:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v76 = F_bms_is_member(m, v72, v71)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	if v76 == int32(0) {
		v88 = v71
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v80 = F_bms_copy(m, v71)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	v82 = F_bms_del_member(m, v80, v72)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	if v75 < int32(0) {
		v88 = v82
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v86 = F_bms_add_member(m, v82, v75)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L20
	} else {
		goto L43
	}
L43:
	;
	v88 = v86
	goto L36
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v108
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v112 = v111
	goto L6
L45:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v96 = F_bms_is_member(m, v92, v91)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	if v96 == int32(0) {
		v108 = v91
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v100 = F_bms_copy(m, v91)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	v102 = F_bms_del_member(m, v100, v92)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	if v95 < int32(0) {
		v108 = v102
		goto L44
	} else {
		goto L50
	}
L50:
	;
	v106 = F_bms_add_member(m, v102, v95)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	v108 = v106
	goto L44
L52:
	;
	if v112 != int32(324) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v112 != int32(378) {
		goto L5
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v135 != 0 {
		goto L5
	} else {
		goto L62
	}
L56:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v121 != 0 {
		v157 = v3
		goto L3
	} else {
		goto L57
	}
L57:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v122 == v123 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v128 = v127
	goto L60
L59:
	;
	v128 = v122
	goto L60
L60:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v128 != v129 {
		v157 = v3
		goto L3
	} else {
		goto L61
	}
L61:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v131
	return int32(0)
L62:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v136 == v137 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v142 = v141
	goto L65
L64:
	;
	v142 = v136
	goto L65
L65:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v142 != v143 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v145
	goto L5
L67:
	;
	v157 = v151
	goto L3
L68:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v167 - int32(1)
	return v165
}
func F_CheckPubDeadTupleRetention(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v4 = m.G0
	v5 = int32(16)
	v6 = v4 - v5
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v5
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPubDeadTupleRetention[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v13 = m.T0[v12].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if int32(_a_F_CheckPubDeadTupleRetention_0) < v13 {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPubDeadTupleRetention[0]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
			v24 = m.T0[v23].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, int32(_a_F_CheckPubDeadTupleRetention_1), int32(1), v6+int32(12))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				if v26 != int32(2) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						F_errcode(m, int32(100663808))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v91
							F_errmsg(m, int32(_a_F_CheckPubDeadTupleRetention_2), v6)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_CheckPubDeadTupleRetention_3), int32(3237), int32(_a_F_CheckPubDeadTupleRetention_4))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
					v31 = F_MakeSingleTupleTableSlot(m, v29, int32(_a_F_CheckPubDeadTupleRetention_5))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
						v36 = F_tuplestore_gettupleslot(m, v33, int32(1), int32(0), v31)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							if v36 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(_a_F_CheckPubDeadTupleRetention_6), int32(0))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CheckPubDeadTupleRetention_3), int32(3241), int32(_a_F_CheckPubDeadTupleRetention_4))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31)+6)))
								if v40 <= int32(0) {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
									m.T0[v45].(func(*base.Module, int32, int32))(m, v31, int32(1))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
										v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
										if v49 != int64(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_CheckPubDeadTupleRetention_7), int32(0))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_CheckPubDeadTupleRetention_3), int32(3248), int32(_a_F_CheckPubDeadTupleRetention_4))
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											F_ExecDropSingleTupleTableSlot(m, v31)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return
											} else {
												v54 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
												if v54 != 0 {
													F_pfree(m, v54)
													mBase = m.M
													v56 = m.ExcPending
													if v56 != 0 {
														return
													} else {
														v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
														if v57 != 0 {
															F_tuplestore_end(m, v57)
															mBase = m.M
															v59 = m.ExcPending
															if v59 != 0 {
																return
															} else {
																v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
																if v60 != 0 {
																	F_FreeTupleDesc(m, v60)
																	mBase = m.M
																	v62 = m.ExcPending
																	if v62 != 0 {
																		return
																	} else {
																		F_pfree(m, v24)
																		mBase = m.M
																		v64 = m.ExcPending
																		if v64 != 0 {
																			return
																		} else {
																			m.G0 = v6 + int32(16)
																			return
																		}
																	}
																} else {
																	F_pfree(m, v24)
																	mBase = m.M
																	v64 = m.ExcPending
																	if v64 != 0 {
																		return
																	} else {
																		m.G0 = v6 + int32(16)
																		return
																	}
																}
															}
														} else {
															v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
															if v60 != 0 {
																F_FreeTupleDesc(m, v60)
																mBase = m.M
																v62 = m.ExcPending
																if v62 != 0 {
																	return
																} else {
																	F_pfree(m, v24)
																	mBase = m.M
																	v64 = m.ExcPending
																	if v64 != 0 {
																		return
																	} else {
																		m.G0 = v6 + int32(16)
																		return
																	}
																}
															} else {
																F_pfree(m, v24)
																mBase = m.M
																v64 = m.ExcPending
																if v64 != 0 {
																	return
																} else {
																	m.G0 = v6 + int32(16)
																	return
																}
															}
														}
													}
												} else {
													v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
													if v57 != 0 {
														F_tuplestore_end(m, v57)
														mBase = m.M
														v59 = m.ExcPending
														if v59 != 0 {
															return
														} else {
															v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
															if v60 != 0 {
																F_FreeTupleDesc(m, v60)
																mBase = m.M
																v62 = m.ExcPending
																if v62 != 0 {
																	return
																} else {
																	F_pfree(m, v24)
																	mBase = m.M
																	v64 = m.ExcPending
																	if v64 != 0 {
																		return
																	} else {
																		m.G0 = v6 + int32(16)
																		return
																	}
																}
															} else {
																F_pfree(m, v24)
																mBase = m.M
																v64 = m.ExcPending
																if v64 != 0 {
																	return
																} else {
																	m.G0 = v6 + int32(16)
																	return
																}
															}
														}
													} else {
														v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
														if v60 != 0 {
															F_FreeTupleDesc(m, v60)
															mBase = m.M
															v62 = m.ExcPending
															if v62 != 0 {
																return
															} else {
																F_pfree(m, v24)
																mBase = m.M
																v64 = m.ExcPending
																if v64 != 0 {
																	return
																} else {
																	m.G0 = v6 + int32(16)
																	return
																}
															}
														} else {
															F_pfree(m, v24)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return
															} else {
																m.G0 = v6 + int32(16)
																return
															}
														}
													}
												}
											}
										}
									}
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
									v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
									if v49 != int64(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return
										} else {
											F_errcode(m, int32(1088))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_CheckPubDeadTupleRetention_7), int32(0))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_CheckPubDeadTupleRetention_3), int32(3248), int32(_a_F_CheckPubDeadTupleRetention_4))
													mBase = m.M
													v129 = m.ExcPending
													if v129 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										F_ExecDropSingleTupleTableSlot(m, v31)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
											if v54 != 0 {
												F_pfree(m, v54)
												mBase = m.M
												v56 = m.ExcPending
												if v56 != 0 {
													return
												} else {
													v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
													if v57 != 0 {
														F_tuplestore_end(m, v57)
														mBase = m.M
														v59 = m.ExcPending
														if v59 != 0 {
															return
														} else {
															v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
															if v60 != 0 {
																F_FreeTupleDesc(m, v60)
																mBase = m.M
																v62 = m.ExcPending
																if v62 != 0 {
																	return
																} else {
																	F_pfree(m, v24)
																	mBase = m.M
																	v64 = m.ExcPending
																	if v64 != 0 {
																		return
																	} else {
																		m.G0 = v6 + int32(16)
																		return
																	}
																}
															} else {
																F_pfree(m, v24)
																mBase = m.M
																v64 = m.ExcPending
																if v64 != 0 {
																	return
																} else {
																	m.G0 = v6 + int32(16)
																	return
																}
															}
														}
													} else {
														v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
														if v60 != 0 {
															F_FreeTupleDesc(m, v60)
															mBase = m.M
															v62 = m.ExcPending
															if v62 != 0 {
																return
															} else {
																F_pfree(m, v24)
																mBase = m.M
																v64 = m.ExcPending
																if v64 != 0 {
																	return
																} else {
																	m.G0 = v6 + int32(16)
																	return
																}
															}
														} else {
															F_pfree(m, v24)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return
															} else {
																m.G0 = v6 + int32(16)
																return
															}
														}
													}
												}
											} else {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
												if v57 != 0 {
													F_tuplestore_end(m, v57)
													mBase = m.M
													v59 = m.ExcPending
													if v59 != 0 {
														return
													} else {
														v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
														if v60 != 0 {
															F_FreeTupleDesc(m, v60)
															mBase = m.M
															v62 = m.ExcPending
															if v62 != 0 {
																return
															} else {
																F_pfree(m, v24)
																mBase = m.M
																v64 = m.ExcPending
																if v64 != 0 {
																	return
																} else {
																	m.G0 = v6 + int32(16)
																	return
																}
															}
														} else {
															F_pfree(m, v24)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return
															} else {
																m.G0 = v6 + int32(16)
																return
															}
														}
													}
												} else {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
													if v60 != 0 {
														F_FreeTupleDesc(m, v60)
														mBase = m.M
														v62 = m.ExcPending
														if v62 != 0 {
															return
														} else {
															F_pfree(m, v24)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return
															} else {
																m.G0 = v6 + int32(16)
																return
															}
														}
													} else {
														F_pfree(m, v24)
														mBase = m.M
														v64 = m.ExcPending
														if v64 != 0 {
															return
														} else {
															m.G0 = v6 + int32(16)
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_CheckPubDeadTupleRetention_8), int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CheckPubDeadTupleRetention_3), int32(3229), int32(_a_F_CheckPubDeadTupleRetention_4))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
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
}
func F_CheckUsageOnTypesInExpr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	v11 = m.G0
	v12 = int32(16)
	v13 = v11 - v12
	m.G0 = v13
	v16 = F_palloc(m, v12)
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
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(137438953472)
	v22 = F_palloc_mul(m, int32(12), int32(32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v16
	v30 = int32(1)
	v32 = F_list_make1_impl(m, v30, v13)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v32
	v37 = F_find_expr_references_walker(m, l0, v13+int32(8))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v39 < int32(2) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if int32(0) < v110 {
		goto L23
	} else {
		goto L24
	}
L7:
	;
	v110 = v39
	goto L6
L8:
	;
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	F_pg_qsort(m, v42, v39, int32(12), int32(497))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if int32(2) <= v47 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v53 = v50
	v57 = v30
	v58 = int32(1)
	goto L14
L12:
	;
	v99 = v30
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v99
	v110 = v99
	goto L6
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v66 = v63 + v58*int32(12)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v62 != v67 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v99 = v87
	goto L13
L16:
	;
	v91 = v58 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v91 < v92 {
		v53 = v86
		v57 = v87
		v58 = v91
		goto L14
	} else {
		goto L22
	}
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+12)) = v80
	v86 = v53 + int32(12)
	v87 = v57 + int32(1)
	goto L16
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v69 != v70 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if v72 == v73 {
		v86 = v53
		v87 = v57
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if v72 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v73
	v86 = v53
	v87 = v57
	goto L16
L22:
	;
	goto L15
L23:
	;
	v124 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	F_pfree(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L36
	}
L26:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v130 = v127 + v124*int32(12)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v131 != int32(1247) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v165 = v124 + int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v165 < v166 {
		v124 = v165
		goto L26
	} else {
		goto L35
	}
L29:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v145 = int32(1)
	goto L30
L30:
	;
	if base.B2i32(int32(0)|base.B2i32(base.Ui32(int32(_a_F_CheckUsageOnTypesInExpr_0)) < base.Ui32(v135)) == int32(0))&((v145|base.B2i32(v135 != int32(2200)))&v145) != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v156 = F_object_aclcheck(m, v153, v154, l2, int64(256))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v156 == int32(0) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	F_aclcheck_error_type(m, v156, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L28
L35:
	;
	goto L27
L36:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v181 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_pfree(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_pfree(m, v16)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	m.G0 = v13 + int32(16)
	return
}
func F_CheckpointerShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerShmemInit[0]))
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4)+4)), uint32(v2))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerShmemInit[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v4)+56)) = int64(0)
	v12 = int32(_a_F_CheckpointerShmemInit_0)
	if v12 <= v9 {
		v15 = v12
	} else {
		v15 = v9
	}
	*(*int32)(unsafe.Add(mBase, uint32(v4)+52)) = v15
	v18 = v4 + int32(24)
	v19 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18))), uint32(v19))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+4)) = int64(-1)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerShmemInit[0]))
	v27 = v25 + int32(36)
	v28 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v27))), uint32(v28))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+4)) = int64(-1)
	return
}
func F_char_bpchar(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_palloc(m, int32(5))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v3)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(20)
		return base.I64_extend_i32_u(v5)
	}
}
func F_chargt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(base.Ui32(v3) < base.Ui32(v2)))
}
func F_charle(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(base.Ui32(v2) <= base.Ui32(v3)))
}
func F_checkMembershipInCurrentExtension(m *base.Module, l0 int32) {
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
	var v16 int32
	_ = v16
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_checkMembershipInCurrentExtension[0])))
	if v8 != int32(1) {
		m.G0 = v5 + int32(16)
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v13 = F_getExtensionOfObject(m, v11, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_checkMembershipInCurrentExtension[1]))
			if v13 == v16 {
				m.G0 = v5 + int32(16)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v26 = F_getObjectDescription(m, l0, int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, _c_F_checkMembershipInCurrentExtension[1]))
							v30 = F_get_extension_name(m, v29)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v30
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v26
								F_errmsg(m, int32(_a_F_checkMembershipInCurrentExtension_0), v5)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									v39 = F_errdetail(m, int32(_a_F_checkMembershipInCurrentExtension_1), int32(0))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_checkMembershipInCurrentExtension_2), int32(297), int32(_a_F_checkMembershipInCurrentExtension_3))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
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
				}
			}
		}
	}
}
func F_check_application_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_pg_clean_ascii(m, v4, int32(2))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v11 = F_guc_strdup(m, int32(15), v6)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				if v11 == int32(0) {
					F_pfree(m, v6)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_bms_free(m, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v6)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11
							v28 = int32(1)
							return v28
						}
					}
				}
			}
		} else {
			v28 = int32(0)
			return v28
		}
	}
}
func F_check_canonical_path(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 != 0 {
		F_canonicalize_path_enc(m, v4)
		mBase = m.M
	} else {
	}
	return int32(1)
}
func F_check_db_file_conflict(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
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
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v14 = F_table_open(m, int32(1213), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	m.T0[v59].(func(*base.Module, int32))(m, v20)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L17
	}
L2:
	;
	return int32(0)
L3:
	;
	v18 = int32(0)
	v20 = F_table_beginscan_catalog(m, v14, v18, v18)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = F_heap_getnext(m, v20)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v22 == int32(0) {
		v56 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v27 = v22
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33+v34)))
	if v36 == int32(1664) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v56 = v2
	goto L1
L9:
	;
	v48 = F_heap_getnext(m, v20)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L15
	}
L10:
	;
	v39 = F_GetDatabasePath(m, l0, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v43 = F___fstatat(m, int32(-100), v39, v10, int32(256))
	mBase = m.M
	goto L12
L12:
	;
	F_pfree(m, v39)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if v43 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v56 = int32(1)
	goto L1
L15:
	;
	if v48 != 0 {
		v27 = v48
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L8
L17:
	;
	F_relation_close(m, v14, int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	m.G0 = v10 + int32(96)
	return v56
}
func F_check_encoding_locale_matches(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v12 = F_pg_get_encoding_from_locale(m, l2, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = F_pg_get_encoding_from_locale(m, l1, int32(1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if base.B2i32(l0 == v12)|base.B2i32(base.Ui32(v12+int32(1)) < base.Ui32(int32(2))) == int32(0) {
				if l0 == int32(0) {
					v27 = F_superuser(m)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						if v27 != 0 {
							if base.Ui32(v15+int32(1)) < base.Ui32(int32(2)) {
								m.G0 = v9 - int32(-64)
								return
							} else {
								v88 = F_superuser(m)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									if v88 != 0 {
										m.G0 = v9 - int32(-64)
										return
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
													v108 = int32(_a_F_check_encoding_locale_matches_0)
												} else {
													v107 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
													v108 = v107
												}
												*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v108
												F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-48))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													if base.B2i32(v15 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v15)) != 0 {
														v127 = int32(_a_F_check_encoding_locale_matches_0)
													} else {
														v126 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
														v127 = v126
													}
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = v127
													v130 = F_errdetail(m, int32(_a_F_check_encoding_locale_matches_2), v9)
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1641), int32(_a_F_check_encoding_locale_matches_4))
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
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
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
										v47 = int32(_a_F_check_encoding_locale_matches_0)
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
										v47 = v46
									}
									*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v47
									F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-16))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										if base.B2i32(v12 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v12)) != 0 {
											v66 = int32(_a_F_check_encoding_locale_matches_0)
										} else {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
											v66 = v65
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v66
										v71 = F_errdetail(m, int32(_a_F_check_encoding_locale_matches_5), v7+int32(-32))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1626), int32(_a_F_check_encoding_locale_matches_4))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
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
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
								v47 = int32(_a_F_check_encoding_locale_matches_0)
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
								v47 = v46
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v47
							F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-16))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								if base.B2i32(v12 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v12)) != 0 {
									v66 = int32(_a_F_check_encoding_locale_matches_0)
								} else {
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
									v66 = v65
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v66
								v71 = F_errdetail(m, int32(_a_F_check_encoding_locale_matches_5), v7+int32(-32))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1626), int32(_a_F_check_encoding_locale_matches_4))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
			} else {
				if base.B2i32(l0 == v15)|base.B2i32(base.Ui32(v15+int32(1)) < base.Ui32(int32(2))) != 0 {
					m.G0 = v9 - int32(-64)
					return
				} else {
					if l0 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
									v108 = int32(_a_F_check_encoding_locale_matches_0)
								} else {
									v107 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
									v108 = v107
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v108
								F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-48))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									if base.B2i32(v15 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v15)) != 0 {
										v127 = int32(_a_F_check_encoding_locale_matches_0)
									} else {
										v126 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
										v127 = v126
									}
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v127
									v130 = F_errdetail(m, int32(_a_F_check_encoding_locale_matches_2), v9)
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1641), int32(_a_F_check_encoding_locale_matches_4))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
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
					} else {
						v88 = F_superuser(m)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							if v88 != 0 {
								m.G0 = v9 - int32(-64)
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
											v108 = int32(_a_F_check_encoding_locale_matches_0)
										} else {
											v107 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
											v108 = v107
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v108
										F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-48))
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											if base.B2i32(v15 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v15)) != 0 {
												v127 = int32(_a_F_check_encoding_locale_matches_0)
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
												v127 = v126
											}
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v127
											v130 = F_errdetail(m, int32(_a_F_check_encoding_locale_matches_2), v9)
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1641), int32(_a_F_check_encoding_locale_matches_4))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
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
						}
					}
				}
			}
		}
	}
}
func F_check_of_type(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
	v11 = v9 + v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
	if v12 == int32(99) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
		v17 = F_relation_open(m, v15, int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+119)))
			F_relation_close(m, v17, int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if v20 != int32(99) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							v58 = F_format_type_be(m, v57)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v58
								F_errmsg(m, int32(_a_F_check_of_type_0), v7)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									v66 = F_errdetail(m, int32(_a_F_check_of_type_1), int32(0))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_check_of_type_2), int32(_a_F_check_of_type_3), int32(_a_F_check_of_type_4))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
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
				} else {
					m.G0 = v7 + int32(32)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v37 = F_format_type_be(m, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v37
					F_errmsg(m, int32(_a_F_check_of_type_5), v7+int32(16))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_of_type_2), int32(_a_F_check_of_type_6), int32(_a_F_check_of_type_4))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
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
}
func F_check_serial_buffers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_check_slru_buffers(m, int32(_a_F_check_serial_buffers_0), l0)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_check_subtrans_buffers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_check_slru_buffers(m, int32(_a_F_check_subtrans_buffers_0), l0)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_choose_next_subplan_for_leader(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v7 = F_LWLockAcquire(m, v5, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v11 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v105 = v5 + int32(20)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+v106))))
	if v108 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14+v11)+20)) = uint8(v16)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v18 - int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)))
	if v22 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v27 = F_ExecFindMatchingSubPlans(m, v24, v23, v23)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v29)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v27
	v33 = int64(0)
	if v27 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if base.B2i32(v77 == v78)|base.B2i32(v78 <= int32(0)) != 0 {
		goto L3
	} else {
		goto L24
	}
L10:
	;
	v77 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v38 = v27 + int32(8)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v39 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v77 = base.I32_popcnt(v42)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v45 = v39 << (uint(int32(2)) % 32)
	if v45 <= int32(7) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v77 = base.I32_wrap_i64(v72)
	goto L9
L17:
	;
	if v45 == int32(0) {
		v72 = v33
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v69 = F_pg_popcount_optimized(m, v38, v45)
	mBase = m.M
	v72 = v69
	goto L16
L20:
	;
	v50 = v45
	v51 = v38
	v52 = v33
	goto L21
L21:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+3)))
	v54 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_choose_next_subplan_for_leader[0]))))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+2)))
	v56 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F_choose_next_subplan_for_leader[0]))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	v58 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v57)+uint32(_c_F_choose_next_subplan_for_leader[0]))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v60 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v59)+uint32(_c_F_choose_next_subplan_for_leader[0]))))
	v64 = v54 + (v56 + (v58 + (v52 + v60)))
	v65 = int32(4)
	v68 = v50 - v65
	if v68 != 0 {
		v50 = v68
		v51 = v51 + v65
		v52 = v64
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v72 = v64
	goto L16
L23:
	;
	goto L22
L24:
	;
	v84 = v23
	goto L25
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v88 = F_bms_is_member(m, v84, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L3
L27:
	;
	if v88 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v94 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v92+v84)+20)) = uint8(v94)
	goto L30
L29:
	;
	goto L30
L30:
	;
	v97 = v84 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v97 < v98 {
		v84 = v97
		goto L25
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	v112 = v106
	goto L35
L33:
	;
	v131 = v106
	goto L34
L34:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v131 < v134 {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	if v112 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v131 = v126
	goto L34
L37:
	;
	v117 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v117
	F_LWLockRelease(m, v5)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v126 = v112 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v126
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v105))))
	if v129 != 0 {
		v112 = v126
		goto L35
	} else {
		goto L41
	}
L40:
	;
	return int32(0)
L41:
	;
	goto L36
L42:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v138 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v136+v131)+20)) = uint8(v138)
	goto L44
L43:
	;
	goto L44
L44:
	;
	F_LWLockRelease(m, v5)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	return int32(1)
}
