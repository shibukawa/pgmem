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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = int32(391772)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[398])))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v17 == v4 {
		v36 = v16
		v37 = v17
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v10 - int32(-64)
	return v172
L2:
	;
	v161 = F_guc_malloc(m, int32(8))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L21
	} else {
		goto L50
	}
L3:
	;
	if v37-v36 == int32(0) {
		v157 = v4
		v158 = v4
		goto L2
	} else {
		goto L11
	}
L4:
	;
	goto L3
L5:
	;
	if v16 != v17 {
		v36 = v16
		v37 = v17
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v21 = v12
	v22 = v13
	goto L7
L7:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v26 == int32(0) {
		v36 = v25
		v37 = v26
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v36 = v25
	v37 = v26
	goto L4
L9:
	;
	v29 = int32(1)
	if v25 == v26 {
		v21 = v21 + v29
		v22 = v22 + v29
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[422])))
	if v42 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[348]))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[423])))
	if v49 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	goto L19
L15:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _consts[424])))
	v157 = v50
	v158 = v52
	goto L2
L16:
	;
	v50 = v46
	goto L18
L17:
	;
	v50 = int32(0)
	goto L18
L18:
	;
	goto L15
L19:
	;
	if base.B2i32(v55 == int32(2)) == int32(0) {
		v172 = v4
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v62 = F_SearchSysCache1(m, int32(10), v61)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	if v62 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if l2 == int32(12) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+22)))
	v106 = v104 + v105
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+68)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	F_ReleaseCatCache(m, v62)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L21
	} else {
		goto L36
	}
L26:
	;
	v70 = int32(1)
	v73 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L21
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, _consts[425])) = v92
	goto L34
L29:
	;
	if v73 == int32(0) {
		v172 = v70
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v80
	F_errmsg(m, int32(78385), v10)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L21
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(524951), int32(983), int32(405058))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v172 = v70
	goto L1
L34:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v95
	v101 = F_format_elog_string(m, int32(78385), v8+int32(-48))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v101
	v172 = int32(0)
	goto L1
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	v113 = F_member_can_set_role(m, v112, v108)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	if v113 != 0 {
		v157 = v108
		v158 = v107
		goto L2
	} else {
		goto L38
	}
L38:
	;
	if l2 == int32(12) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v117 = int32(1)
	v120 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L21
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[427])) = int32(16797828)
	goto L47
L42:
	;
	if v120 == int32(0) {
		v172 = v117
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v127
	F_errmsg(m, int32(751194), v8+int32(-32))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L21
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(524951), int32(1004), int32(405058))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L21
	} else {
		goto L46
	}
L46:
	;
	v172 = v117
	goto L1
L47:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, _consts[425])) = v144
	goto L48
L48:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v147
	v153 = F_format_elog_string(m, int32(751159), v8+int32(-16))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v153
	v172 = int32(0)
	goto L1
L50:
	;
	if v161 == int32(0) {
		v172 = int32(0)
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = v157
	v166 = int32(1)
	v168 = v158 & v166
	*(*uint8)(unsafe.Add(mBase, uint32(v161)+4)) = uint8(v168)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v161
	v172 = v166
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
							F_errmsg_internal(m, int32(140088), int32(0))
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(520594), int32(2231), int32(218974))
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
								F_errmsg(m, int32(751237), v8+int32(48))
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
										F_errdetail(m, int32(667375), v8+int32(32))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(520594), int32(2252), int32(218974))
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
											F_errmsg(m, int32(751237), v8+int32(16))
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
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(557325)
													F_errdetail(m, int32(698947), v8)
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(520594), int32(2261), int32(218974))
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
								F_errmsg(m, int32(750922), v8+int32(80))
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
										F_errdetail(m, int32(667227), v8-int32(-64))
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(520594), int32(2271), int32(218974))
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(10), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
			v20 = F_psprintf(m, int32(607668), v9)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v83 = v20
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v83
				v86 = int32(0)
				m.G0 = v9 + int32(48)
				return v86
			}
		} else {
			v26 = F_SysCacheGetAttr(m, int32(10), v12, int32(11), v9+int32(47))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+47)))
				if v28 == int32(1) {
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
						v37 = F_psprintf(m, int32(680200), v9+int32(16))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v83 = v37
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v83
							v86 = int32(0)
							m.G0 = v9 + int32(48)
							return v86
						}
					}
				} else {
					v39 = F_text_to_cstring(m, v26)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v45 = F_SysCacheGetAttr(m, int32(10), v12, int32(12), v9+int32(47))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+47)))
							if v47 == int32(0) {
								v50 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
								v51 = v50
							} else {
								v51 = int64(0)
							}
							F_ReleaseCatCache(m, v12)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+47)))
								if v54 != 0 {
									v86 = v39
									m.G0 = v9 + int32(48)
									return v86
								} else {
									v58 = m.G0
									v59 = int32(16)
									v60 = v58 - v59
									m.G0 = v60
									F___gettimeofday(m, v60)
									mBase = m.M
									v63 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
									v64 = int64(*(*int32)(unsafe.Add(mBase, uint32(v60)+8)))
									m.G0 = v60 + v59
									if v64+v63*int64(1000000)-int64(946684800000000) <= v51 {
										v86 = v39
										m.G0 = v9 + int32(48)
										return v86
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0
										v78 = F_psprintf(m, int32(675003), v9+int32(32))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											v83 = v78
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = v83
											v86 = int32(0)
											m.G0 = v9 + int32(48)
											return v86
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
