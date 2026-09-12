package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_role_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v146 int32
	_ = v146
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v146
L2:
	;
	v8 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v9 <= v8 {
		v146 = v8
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v146 = int32(0)
	goto L1
L5:
	;
	v17 = int32(0)
	goto L6
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+v17<<(uint(int32(2))%32))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L4
L8:
	;
	v128 = v17 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v128 < v129 {
		v17 = v128
		goto L6
	} else {
		goto L46
	}
L9:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v75 != 0 {
		goto L29
	} else {
		goto L30
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v26 == int32(43) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if l1 == int32(0) {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v46 = int32(305615)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[369])))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v50 == int32(0) {
		v69 = v49
		v70 = v50
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v31 = int32(1)
	v34 = F_get_role_oid(m, v25+v31, v31)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	if v34 == int32(0) {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v40 = F_is_member_of_role_nosuper(m, l1, v34)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	if v40 == int32(0) {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	return int32(1)
L20:
	;
	if v70-v69 != 0 {
		goto L9
	} else {
		goto L28
	}
L21:
	;
	goto L20
L22:
	;
	if v49 != v50 {
		v69 = v49
		v70 = v50
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v54 = v25
	v55 = v46
	goto L24
L24:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v59 == int32(0) {
		v69 = v58
		v70 = v59
		goto L21
	} else {
		goto L26
	}
L25:
	;
	v69 = v58
	v70 = v59
	goto L21
L26:
	;
	v62 = int32(1)
	if v58 == v59 {
		v54 = v54 + v62
		v55 = v55 + v62
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	return int32(1)
L29:
	;
	v76 = F_strlen(m, l0)
	mBase = m.M
	v81 = F_palloc(m, v76<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L15
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v100 == int32(0) {
		v119 = v99
		v120 = v100
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v83 = F_strlen(m, l0)
	mBase = m.M
	v84 = F_pg_mb2wchar_with_len(m, l0, v81, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v87 = int32(0)
	v90 = F_pg_regexec(m, v86, v81, v84, v87, v87, v87)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	F_pfree(m, v81)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	if v90 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	return int32(1)
L37:
	;
	if v120-v119 != 0 {
		goto L8
	} else {
		goto L45
	}
L38:
	;
	goto L37
L39:
	;
	if v99 != v100 {
		v119 = v99
		v120 = v100
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v104 = v96
	v105 = l0
	goto L41
L41:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	if v109 == int32(0) {
		v119 = v108
		v120 = v109
		goto L38
	} else {
		goto L43
	}
L42:
	;
	v119 = v108
	v120 = v109
	goto L38
L43:
	;
	v112 = int32(1)
	if v108 == v109 {
		v104 = v104 + v112
		v105 = v105 + v112
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	return int32(1)
L46:
	;
	goto L7
}
func F_check_role_membership_authorization(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	v5 = m.G0
	v7 = v5 - int32(144)
	m.G0 = v7
	if l1 == int32(6171) {
		v12 = l2
	} else {
		v12 = int32(0)
	}
	if v12 == int32(0) {
		v15 = F_superuser_arg(m, l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 != 0 {
				v17 = F_superuser_arg(m, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					if v17 != 0 {
						m.G0 = v7 + int32(144)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								v27 = F_GetUserNameFromId(m, l1, int32(0))
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
									if l2 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v27
										F_errmsg(m, int32(715559), v7+int32(32))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return
										} else {
											v112 = int32(525854)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v112
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v112
											F_errdetail(m, int32(629712), v7+int32(16))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												F_errfinish(m, int32(495664), int32(2140), int32(258696))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v27
										F_errmsg(m, int32(715795), v7-int32(-64))
										mBase = m.M
										v34 = m.ExcPending
										if v34 != 0 {
											return
										} else {
											v35 = int32(525854)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v35
											*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v35
											F_errdetail(m, int32(630246), v7+int32(48))
											mBase = m.M
											v43 = m.ExcPending
											if v43 != 0 {
												return
											} else {
												F_errfinish(m, int32(495664), int32(2147), int32(258696))
												mBase = m.M
												v48 = m.ExcPending
												if v48 != 0 {
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
			} else {
				v49 = F_is_admin_of_role(m, l0, l1)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					if v49 != 0 {
						m.G0 = v7 + int32(144)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								v59 = F_GetUserNameFromId(m, l1, int32(0))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									if l2 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = v59
										F_errmsg(m, int32(715559), v7+int32(96))
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return
										} else {
											v133 = F_GetUserNameFromId(m, l1, int32(0))
											mBase = m.M
											v134 = m.ExcPending
											if v134 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+84)) = v133
												*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = int32(530786)
												F_errdetail(m, int32(637144), v7+int32(80))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return
												} else {
													F_errfinish(m, int32(495664), int32(2163), int32(258696))
													mBase = m.M
													v147 = m.ExcPending
													if v147 != 0 {
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
										*(*int32)(unsafe.Add(mBase, uint32(v7)+128)) = v59
										F_errmsg(m, int32(715795), v7+int32(128))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											v68 = F_GetUserNameFromId(m, l1, int32(0))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+116)) = v68
												*(*int32)(unsafe.Add(mBase, uint32(v7)+112)) = int32(530786)
												F_errdetail(m, int32(637538), v7+int32(112))
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													F_errfinish(m, int32(495664), int32(2170), int32(258696))
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v89 = m.ExcPending
		if v89 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return
			} else {
				v95 = F_GetUserNameFromId(m, int32(6171), int32(0))
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v95
					F_errmsg(m, int32(135902), v7)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return
					} else {
						F_errfinish(m, int32(495664), int32(2127), int32(258696))
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
		}
	}
}
func F_get_role_oid_or_public(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(491806)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[768])))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v12 == v2 {
		v31 = v11
		v32 = v12
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return v60
L2:
	;
	if v32-v31 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	goto L2
L4:
	;
	if v11 != v12 {
		v31 = v11
		v32 = v12
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v16 = l0
	v17 = v8
	goto L6
L6:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v21 == int32(0) {
		v31 = v20
		v32 = v21
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v31 = v20
	v32 = v21
	goto L3
L8:
	;
	v24 = int32(1)
	if v20 == v21 {
		v16 = v16 + v24
		v17 = v17 + v24
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v60 = v2
	goto L1
L11:
	;
	goto L12
L12:
	;
	v37 = int32(0)
	v40 = F_GetSysCacheOid(m, int32(10), l0, v37, v37, v37)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v40 != 0 {
		v60 = v40
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(72320), v6)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(497936), int32(5562), int32(435314))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
