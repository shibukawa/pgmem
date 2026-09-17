package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecRelGenVirtualNotNull(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v11 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = int32(_a_F_ExecRelGenVirtualNotNull_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRelGenVirtualNotNull[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecRelGenVirtualNotNull[0])) = v18
	if l3 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v95 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecRelGenVirtualNotNull[0])) = v16
	goto L3
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v23 = F_palloc0(m, v20<<(uint(int32(2))%32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v70 = F_palloc0(m, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L17
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v23
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v28 <= int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v36 = v5
	goto L11
L11:
	;
	v42 = v36 << (uint(int32(2)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42+v43)))
	v47 = F_palloc0(m, int32(20))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L13
	}
L12:
	;
	goto L4
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(52)
	v51 = F_build_generation_expression(m, v14, v45)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = int32(-1)
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+12)) = uint8(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v51
	v60 = F_ExecPrepareExpr(m, v47, l2)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v62+v42))) = v60
	v66 = v36 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v66 < v67 {
		v36 = v66
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v70
	goto L4
L18:
	;
	v98 = F_MakePerTupleExprContext(m, l2)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	v100 = v95
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = l1
	if l3 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v100 = v98
	goto L20
L22:
	;
	return int32(0)
L23:
	;
	goto L24
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v106 <= int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	goto L27
L27:
	;
	v113 = int32(0)
	goto L29
L28:
	;
	return base.I32_extend16_s(v126)
L29:
	;
	v123 = v113 << (uint(int32(2)) % 32)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123+v124)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v127+v123)))
	v130 = F_ExecCheck(m, v129, v100)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	if v130 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v135 = v113 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v135 < v136 {
		v113 = v135
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
}
func F_rel_is_distinct_for(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v127 int32
	_ = v127
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 != 0 {
		v127 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v127
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	switch v12 {
	case 0:
		goto L5
	case 1:
		goto L4
	default:
		v127 = v5
		goto L1
	}
L3:
	;
	v127 = int32(1)
	goto L1
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19+v20<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if l2 == int32(0) {
		v99 = v5
		v100 = v5
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v13 = int32(0)
	v15 = F_relation_has_unique_index_ext(m, l0, l1, l2, v13, v13, l3)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v15 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v127 = v5
	goto L1
L9:
	;
	v103 = F_query_is_distinct_for(m, v25, v99, v100)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L6
	} else {
		goto L33
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v28 <= int32(0) {
		v99 = v5
		v100 = v5
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v33 = int32(0)
	v38 = v5
	v39 = v5
	goto L12
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v33<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+120)))
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v99 = v87
	v100 = v88
	goto L9
L14:
	;
	v90 = v33 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v90 < v91 {
		v33 = v90
		v38 = v87
		v39 = v88
		goto L12
	} else {
		goto L32
	}
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 == int32(0) {
		v87 = v38
		v88 = v39
		goto L14
	} else {
		goto L22
	}
L16:
	;
	if v48 == int32(0) {
		v87 = v38
		v88 = v39
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v48 == int32(0) {
		v87 = v38
		v88 = v39
		goto L14
	} else {
		goto L21
	}
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v53 < int32(2) {
		v87 = v38
		v88 = v39
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v62 = v56 + int32(4)
	goto L15
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v62 = v61
	goto L15
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v66 == int32(27) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v69 == int32(0) {
		v87 = v38
		v88 = v39
		goto L14
	} else {
		goto L26
	}
L24:
	;
	v73 = v63
	v74 = v66
	goto L25
L25:
	;
	if v74 != int32(6) {
		v87 = v38
		v88 = v39
		goto L14
	} else {
		goto L27
	}
L26:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v73 = v69
	v74 = v72
	goto L25
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v77 != v20 {
		v87 = v38
		v88 = v39
		goto L14
	} else {
		goto L28
	}
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
	if v79 != 0 {
		v87 = v38
		v88 = v39
		goto L14
	} else {
		goto L29
	}
L29:
	;
	v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+8)))
	v81 = F_lappend_int(m, v38, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v83 = F_lappend_oid(m, v39, v49)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v87 = v81
	v88 = v83
	goto L14
L32:
	;
	goto L13
L33:
	;
	if v103 == int32(0) {
		v127 = v5
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L3
}
func F_truncate_check_rel(m *base.Module, l0 int32, l1 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = l1 + int32(4)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+119)))
	switch v11 - int32(102) {
	case 0:
		v14 = F_GetForeignServerIdByRelId(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = F_GetFdwRoutineByServerId(m, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+136))
				if v18 != 0 {
					v54 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_truncate_check_rel[0])))
					if v54 != 0 {
						v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
						if v76 != 0 {
							v79 = int32(0)
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
							m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						} else {
							m.G0 = v7 + int32(48)
							return
						}
					} else {
						v56 = int32(1)
						if base.Ui32(l0) < base.Ui32(int32(_a_F_truncate_check_rel_0)) {
							v64 = v56
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
							if v59 == int32(99) {
								v64 = v56
							} else {
								v62 = F_isTempToastNamespace(m, v59)
								mBase = m.M
								v64 = v62
							}
						}
						if v64 == int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
							if v76 != 0 {
								v79 = int32(0)
								v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
								m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									m.G0 = v7 + int32(48)
									return
								}
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						} else {
							if l0 != int32(2613) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v10
										F_errmsg(m, int32(_a_F_truncate_check_rel_1), v7+int32(32))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_truncate_check_rel_2), int32(2411), int32(_a_F_truncate_check_rel_3))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
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
								v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_truncate_check_rel[2])))
								if v70&int32(1) == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										F_errcode(m, int32(16797828))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v10
											F_errmsg(m, int32(_a_F_truncate_check_rel_1), v7+int32(32))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_truncate_check_rel_2), int32(2411), int32(_a_F_truncate_check_rel_3))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
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
									v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
									if v76 != 0 {
										v79 = int32(0)
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
										m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return
										} else {
											m.G0 = v7 + int32(48)
											return
										}
									} else {
										m.G0 = v7 + int32(48)
										return
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v10
							F_errmsg(m, int32(_a_F_truncate_check_rel_4), v7+int32(16))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_truncate_check_rel_2), int32(2391), int32(_a_F_truncate_check_rel_3))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
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
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
				F_errmsg(m, int32(_a_F_truncate_check_rel_5), v7)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_truncate_check_rel_2), int32(2397), int32(_a_F_truncate_check_rel_3))
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
	case 10, 12:
		v54 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_truncate_check_rel[0])))
		if v54 != 0 {
			v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
			if v76 != 0 {
				v79 = int32(0)
				v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
				m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					m.G0 = v7 + int32(48)
					return
				}
			} else {
				m.G0 = v7 + int32(48)
				return
			}
		} else {
			v56 = int32(1)
			if base.Ui32(l0) < base.Ui32(int32(_a_F_truncate_check_rel_0)) {
				v64 = v56
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
				if v59 == int32(99) {
					v64 = v56
				} else {
					v62 = F_isTempToastNamespace(m, v59)
					mBase = m.M
					v64 = v62
				}
			}
			if v64 == int32(0) {
				v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
				if v76 != 0 {
					v79 = int32(0)
					v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
					m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						m.G0 = v7 + int32(48)
						return
					}
				} else {
					m.G0 = v7 + int32(48)
					return
				}
			} else {
				if l0 != int32(2613) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						F_errcode(m, int32(16797828))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v10
							F_errmsg(m, int32(_a_F_truncate_check_rel_1), v7+int32(32))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_truncate_check_rel_2), int32(2411), int32(_a_F_truncate_check_rel_3))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
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
					v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_truncate_check_rel[2])))
					if v70&int32(1) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v10
								F_errmsg(m, int32(_a_F_truncate_check_rel_1), v7+int32(32))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_truncate_check_rel_2), int32(2411), int32(_a_F_truncate_check_rel_3))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
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
						v76 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
						if v76 != 0 {
							v79 = int32(0)
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_check_rel[1]))
							m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(5), int32(1259), l0, v79, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						} else {
							m.G0 = v7 + int32(48)
							return
						}
					}
				}
			}
		}
	}
}
