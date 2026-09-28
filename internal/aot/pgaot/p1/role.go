package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_role_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = int32(_a_F_check_role_1_0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_role_1[0])))
	if base.B2i32(v16 == v4)|base.B2i32(v16 != v19) != 0 {
		v37 = v16
		v38 = v19
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v10 - int32(-64)
	return v173
L2:
	;
	v162 = F_guc_malloc(m, int32(8))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L20
	} else {
		goto L49
	}
L3:
	;
	if v37-v38 == int32(0) {
		v158 = v4
		v159 = v4
		goto L2
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	v22 = v12
	v23 = v13
	goto L6
L6:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v27 == int32(0) {
		v37 = v27
		v38 = v26
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v37 = v27
	v38 = v26
	goto L4
L8:
	;
	v30 = int32(1)
	if v27 == v26 {
		v22 = v22 + v30
		v23 = v23 + v30
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_role_1[1])))
	if v43 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_check_role_1[2]))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_role_1[3])))
	if v50 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_check_role_1[4]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	goto L18
L14:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_role_1[5])))
	v158 = v51
	v159 = v53
	goto L2
L15:
	;
	v51 = v47
	goto L17
L16:
	;
	v51 = int32(0)
	goto L17
L17:
	;
	goto L14
L18:
	;
	if base.B2i32(v56 == int32(2)) == int32(0) {
		v173 = v4
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v62 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	v63 = F_SearchSysCache1(m, int32(10), v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	if v63 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if l2 == int32(12) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+22)))
	v107 = v105 + v106
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+68)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	F_ReleaseCatCache(m, v63)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L20
	} else {
		goto L35
	}
L25:
	;
	v71 = int32(1)
	v74 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L20
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_check_role_1[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_role_1[7])) = v93
	goto L33
L28:
	;
	if v74 == int32(0) {
		v173 = v71
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L20
	} else {
		goto L30
	}
L30:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v81
	F_errmsg(m, int32(_a_F_check_role_1_1), v10)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_check_role_1_2), int32(983), int32(_a_F_check_role_1_3))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	v173 = v71
	goto L1
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v96
	v102 = F_format_elog_string(m, int32(_a_F_check_role_1_1), v8+int32(-48))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_role_1[8])) = v102
	v173 = int32(0)
	goto L1
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_check_role_1[9]))
	v114 = F_member_can_set_role(m, v113, v109)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	if v114 != 0 {
		v158 = v109
		v159 = v108
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if l2 == int32(12) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v118 = int32(1)
	v121 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L20
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_role_1[10])) = int32(16797828)
	goto L46
L41:
	;
	if v121 == int32(0) {
		v173 = v118
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L20
	} else {
		goto L43
	}
L43:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v128
	F_errmsg(m, int32(_a_F_check_role_1_4), v8+int32(-32))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_check_role_1_2), int32(1004), int32(_a_F_check_role_1_3))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	v173 = v118
	goto L1
L46:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_check_role_1[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_role_1[7])) = v145
	goto L47
L47:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v148
	v154 = F_format_elog_string(m, int32(_a_F_check_role_1_5), v8+int32(-16))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_role_1[8])) = v154
	v173 = int32(0)
	goto L1
L49:
	;
	if v162 == int32(0) {
		v173 = int32(0)
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v158
	v167 = int32(1)
	v169 = v159 & v167
	*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)) = uint8(v169)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v162
	v173 = v167
	goto L1
}
func F_check_role_grantor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	if l2 == int32(0) {
		v13 = F_superuser_arg(m, l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v13 != 0 {
				v76 = int32(10)
				m.G0 = v8 + int32(96)
				return v76
			} else {
				v17 = F_select_best_admin(m, l0, l1)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v17 != 0 {
						v76 = v17
						m.G0 = v8 + int32(96)
						return v76
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_check_role_grantor_0), int32(0))
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_check_role_grantor_1), int32(2244), int32(_a_F_check_role_grantor_2))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
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
	} else {
		v32 = F_has_privs_of_role(m, l0, l2)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			if l3 != 0 {
				if v32 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16797828))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v89 = F_GetUserNameFromId(m, l2, int32(0))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v89
								F_errmsg(m, int32(_a_F_check_role_grantor_3), v8+int32(48))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									v98 = F_GetUserNameFromId(m, l2, int32(0))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v98
										v104 = F_errdetail(m, int32(_a_F_check_role_grantor_4), v8+int32(32))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_check_role_grantor_1), int32(2265), int32(_a_F_check_role_grantor_2))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
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
				} else {
					v36 = int32(10)
					if l2 == v36 {
						v76 = v36
						m.G0 = v8 + int32(96)
						return v76
					} else {
						v39 = F_select_best_admin(m, l2, l1)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 == l2 {
								v76 = l2
								m.G0 = v8 + int32(96)
								return v76
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v50 = F_GetUserNameFromId(m, l2, int32(0))
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v50
											F_errmsg(m, int32(_a_F_check_role_grantor_3), v8+int32(16))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												v59 = F_GetUserNameFromId(m, l1, int32(0))
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v59
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_check_role_grantor_5)
													v65 = F_errdetail(m, int32(_a_F_check_role_grantor_6), v8)
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_check_role_grantor_1), int32(2274), int32(_a_F_check_role_grantor_2))
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
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
					}
				}
			} else {
				if v32 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16797828))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int32(0)
						} else {
							v119 = F_GetUserNameFromId(m, l2, int32(0))
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v119
								F_errmsg(m, int32(_a_F_check_role_grantor_7), v8+int32(80))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									v128 = F_GetUserNameFromId(m, l2, int32(0))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v128
										v134 = F_errdetail(m, int32(_a_F_check_role_grantor_8), v8-int32(-64))
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_check_role_grantor_1), int32(2284), int32(_a_F_check_role_grantor_2))
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
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
				} else {
					v76 = l2
					m.G0 = v8 + int32(96)
					return v76
				}
			}
		}
	}
}
func F_get_role_password(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	v18 = F_SearchSysCache1(m, int32(10), base.I64_extend_i32_u(l0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v18 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
			v26 = F_psprintf(m, int32(_a_F_get_role_password_0), v14)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
				v178 = int32(0)
				m.G0 = v14 + int32(112)
				return v178
			}
		} else {
			v33 = F_SysCacheGetAttr(m, int32(10), v18, int32(11), v14+int32(111))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+111)))
				if v35 == int32(1) {
					F_ReleaseCatCache(m, v18)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l0
						v44 = F_psprintf(m, int32(_a_F_get_role_password_1), v14+int32(16))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v44
							v178 = int32(0)
							m.G0 = v14 + int32(112)
							return v178
						}
					}
				} else {
					v48 = F_text_to_cstring(m, base.I32_wrap_i64(v33))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v54 = F_SysCacheGetAttr(m, int32(10), v18, int32(12), v14+int32(111))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+111)))
							F_ReleaseCatCache(m, v18)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+111)))
								if v59 != 0 {
									v178 = v48
									m.G0 = v14 + int32(112)
									return v178
								} else {
									if v56&int32(1) != 0 {
										v63 = int64(0)
									} else {
										v63 = v54
									}
									v67 = m.G0
									v68 = int32(16)
									v69 = v67 - v68
									m.G0 = v69
									F_gettimeofday(m, v69)
									mBase = m.M
									v72 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
									v73 = int64(*(*int32)(unsafe.Add(mBase, uint32(v69)+8)))
									m.G0 = v69 + v68
									v81 = v73 + v72*int64(1000000) - int64(946684800000000)
									if v81 <= v63 {
										v84 = int64(*(*int32)(unsafe.Add(mBase, _c_F_get_role_password[0])))
										if v81 < v63 {
											v88 = v63 - v81
										} else {
											v88 = int64(0)
										}
										v90 = base.I64_div_u_s(v88, int64(1000000))
										if base.Ui64(v84) <= base.Ui64(v90) {
											v178 = v48
											m.G0 = v14 + int32(112)
											return v178
										} else {
											v92 = int32(_a_F_get_role_password_2)
											v93 = *(*int32)(unsafe.Add(mBase, _c_F_get_role_password[1]))
											v96 = *(*int32)(unsafe.Add(mBase, _c_F_get_role_password[2]))
											*(*int32)(unsafe.Add(mBase, _c_F_get_role_password[1])) = v96
											v98 = int64(3600000000)
											v99 = base.I64_rem_u_s(v88, v98)
											v102 = base.I32_div_u_s(base.I32_wrap_i64(v99), int32(60000000))
											v103 = int64(86400000000)
											v104 = base.I64_div_u_s(v88, v103)
											v107 = v88 - v104*v103
											v109 = base.I64_div_u_s(v107, v98)
											v111 = F_pstrdup(m, int32(_a_F_get_role_password_3))
											mBase = m.M
											v112 = m.ExcPending
											if v112 != 0 {
												return int32(0)
											} else {
												if base.Ui64(int64(86400000000)) <= base.Ui64(v88) {
													*(*uint32)(unsafe.Add(mBase, uint32(v14)+100)) = uint32(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = l0
													if base.Ui64(v88) < base.Ui64(int64(172800000000)) {
														v121 = int32(_a_F_get_role_password_4)
													} else {
														v121 = int32(_a_F_get_role_password_5)
													}
													v124 = F_psprintf(m, v121, v14+int32(96))
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return int32(0)
													} else {
														v158 = v124
														F_StoreConnectionWarning(m, v111, v158, int32(0))
														mBase = m.M
														v161 = m.ExcPending
														if v161 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_get_role_password[1])) = v93
															v178 = v48
															m.G0 = v14 + int32(112)
															return v178
														}
													}
												} else {
													if base.Ui64(int64(3600000000)) <= base.Ui64(v107) {
														*(*uint32)(unsafe.Add(mBase, uint32(v14)+84)) = uint32(v109)
														*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = l0
														if base.Ui64(v107) < base.Ui64(int64(7200000000)) {
															v134 = int32(_a_F_get_role_password_6)
														} else {
															v134 = int32(_a_F_get_role_password_7)
														}
														v137 = F_psprintf(m, v134, v14+int32(80))
														mBase = m.M
														v138 = m.ExcPending
														if v138 != 0 {
															return int32(0)
														} else {
															v158 = v137
															F_StoreConnectionWarning(m, v111, v158, int32(0))
															mBase = m.M
															v161 = m.ExcPending
															if v161 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_get_role_password[1])) = v93
																v178 = v48
																m.G0 = v14 + int32(112)
																return v178
															}
														}
													} else {
														if base.Ui64(int64(60000000)) <= base.Ui64(v99) {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v102
															*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = l0
															if base.Ui64(v99) < base.Ui64(int64(120000000)) {
																v147 = int32(_a_F_get_role_password_8)
															} else {
																v147 = int32(_a_F_get_role_password_9)
															}
															v150 = F_psprintf(m, v147, v14-int32(-64))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int32(0)
															} else {
																v158 = v150
																F_StoreConnectionWarning(m, v111, v158, int32(0))
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_get_role_password[1])) = v93
																	v178 = v48
																	m.G0 = v14 + int32(112)
																	return v178
																}
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l0
															v156 = F_psprintf(m, int32(_a_F_get_role_password_10), v14+int32(48))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																v158 = v156
																F_StoreConnectionWarning(m, v111, v158, int32(0))
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_get_role_password[1])) = v93
																	v178 = v48
																	m.G0 = v14 + int32(112)
																	return v178
																}
															}
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l0
										v168 = F_psprintf(m, int32(_a_F_get_role_password_11), v14+int32(32))
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = v168
											v178 = int32(0)
											m.G0 = v14 + int32(112)
											return v178
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
