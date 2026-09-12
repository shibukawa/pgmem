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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
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
	v15 = int32(4489152)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+100))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v18
	if l3 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v97 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v16
	goto L3
L5:
	;
	v23 = F_palloc0(m, int32(0))
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v31 = F_palloc0(m, v28<<(uint(int32(2))%32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v23
	goto L4
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v31
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v34 <= int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v42 = v5
	goto L12
L12:
	;
	v48 = v42 << (uint(int32(2)) % 32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48+v49)))
	v53 = F_palloc0(m, int32(20))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L4
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = int32(52)
	v57 = F_build_generation_expression(m, v14, v51)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = int32(-1)
	v61 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)) = uint8(v61)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v57
	v66 = F_ExecPrepareExpr(m, v53, l2)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v68+v48))) = v66
	v72 = v42 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v72 < v73 {
		v42 = v72
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	v100 = F_MakePerTupleExprContext(m, l2)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	v102 = v97
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+4)) = l1
	if l3 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v102 = v100
	goto L20
L22:
	;
	return int32(0)
L23:
	;
	goto L24
L24:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v108 <= int32(0) {
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
	v118 = int32(0)
	goto L29
L28:
	;
	return base.I32_extend16_s(v128)
L29:
	;
	v125 = v118 << (uint(int32(2)) % 32)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125+v126)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129+v125)))
	v132 = F_ExecCheck(m, v131, v102)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	if v132 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v137 = v118 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v137 < v138 {
		v118 = v137
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
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
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v128 int32
	_ = v128
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 != 0 {
		v128 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v128
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	switch v12 {
	case 0:
		goto L5
	case 1:
		goto L4
	default:
		v128 = v5
		goto L1
	}
L3:
	;
	v128 = int32(1)
	goto L1
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19+v20<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if l2 == int32(0) {
		goto L10
	} else {
		goto L11
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
	v128 = v5
	goto L1
L9:
	;
	v106 = F_query_is_distinct_for(m, v25, v101, v100)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L34
	}
L10:
	;
	v100 = v5
	v101 = v5
	goto L9
L11:
	;
	v28 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v29 <= v28 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v33 = v28
	v36 = v5
	v37 = v5
	goto L13
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v33<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+120)))
	if v50 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v100 = v89
	v101 = v90
	goto L9
L15:
	;
	v92 = v33 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v92 < v93 {
		v33 = v92
		v36 = v89
		v37 = v90
		goto L13
	} else {
		goto L33
	}
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v65 == int32(0) {
		v89 = v36
		v90 = v37
		goto L15
	} else {
		goto L23
	}
L17:
	;
	if v48 == int32(0) {
		v89 = v36
		v90 = v37
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v48 == int32(0) {
		v89 = v36
		v90 = v37
		goto L15
	} else {
		goto L22
	}
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v55 < int32(2) {
		v89 = v36
		v90 = v37
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v64 = v58 + int32(4)
	goto L16
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v64 = v63
	goto L16
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v68 == int32(27) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v71 == int32(0) {
		v89 = v36
		v90 = v37
		goto L15
	} else {
		goto L27
	}
L25:
	;
	v75 = v65
	v76 = v68
	goto L26
L26:
	;
	if v76 != int32(6) {
		v89 = v36
		v90 = v37
		goto L15
	} else {
		goto L28
	}
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v75 = v71
	v76 = v74
	goto L26
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v79 != v20 {
		v89 = v36
		v90 = v37
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+28))
	if v81 != 0 {
		v89 = v36
		v90 = v37
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v75)+8)))
	v83 = F_lappend_int(m, v37, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v85 = F_lappend_oid(m, v36, v49)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v89 = v85
	v90 = v83
	goto L15
L33:
	;
	goto L14
L34:
	;
	if v106 == int32(0) {
		v128 = v5
		goto L1
	} else {
		goto L35
	}
L35:
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
					v54 = int32(*(*uint8)(unsafe.Add(mBase, _consts[222])))
					if v54 != 0 {
						v76 = *(*int32)(unsafe.Add(mBase, _consts[230]))
						if v76 != 0 {
							v79 = int32(0)
							v82 = *(*int32)(unsafe.Add(mBase, _consts[230]))
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
						if base.Ui32(l0) < base.Ui32(int32(12000)) {
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
							v76 = *(*int32)(unsafe.Add(mBase, _consts[230]))
							if v76 != 0 {
								v79 = int32(0)
								v82 = *(*int32)(unsafe.Add(mBase, _consts[230]))
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
										F_errmsg(m, int32(326345), v7+int32(32))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											F_errfinish(m, int32(492711), int32(2411), int32(306712))
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
								v70 = int32(*(*uint8)(unsafe.Add(mBase, _consts[79])))
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
											F_errmsg(m, int32(326345), v7+int32(32))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												F_errfinish(m, int32(492711), int32(2411), int32(306712))
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
									v76 = *(*int32)(unsafe.Add(mBase, _consts[230]))
									if v76 != 0 {
										v79 = int32(0)
										v82 = *(*int32)(unsafe.Add(mBase, _consts[230]))
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
							F_errmsg(m, int32(702398), v7+int32(16))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_errfinish(m, int32(492711), int32(2391), int32(306712))
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
				F_errmsg(m, int32(393932), v7)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_errfinish(m, int32(492711), int32(2397), int32(306712))
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
		v54 = int32(*(*uint8)(unsafe.Add(mBase, _consts[222])))
		if v54 != 0 {
			v76 = *(*int32)(unsafe.Add(mBase, _consts[230]))
			if v76 != 0 {
				v79 = int32(0)
				v82 = *(*int32)(unsafe.Add(mBase, _consts[230]))
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
			if base.Ui32(l0) < base.Ui32(int32(12000)) {
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
				v76 = *(*int32)(unsafe.Add(mBase, _consts[230]))
				if v76 != 0 {
					v79 = int32(0)
					v82 = *(*int32)(unsafe.Add(mBase, _consts[230]))
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
							F_errmsg(m, int32(326345), v7+int32(32))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return
							} else {
								F_errfinish(m, int32(492711), int32(2411), int32(306712))
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
					v70 = int32(*(*uint8)(unsafe.Add(mBase, _consts[79])))
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
								F_errmsg(m, int32(326345), v7+int32(32))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									F_errfinish(m, int32(492711), int32(2411), int32(306712))
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
						v76 = *(*int32)(unsafe.Add(mBase, _consts[230]))
						if v76 != 0 {
							v79 = int32(0)
							v82 = *(*int32)(unsafe.Add(mBase, _consts[230]))
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
