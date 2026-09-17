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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v148 int32
	_ = v148
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v148
L2:
	;
	v8 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v9 <= v8 {
		v148 = v8
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v148 = int32(0)
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
	v130 = v17 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v130 < v131 {
		v17 = v130
		goto L6
	} else {
		goto L44
	}
L9:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v76 != 0 {
		goto L28
	} else {
		goto L29
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
	v46 = int32(_a_F_check_role_2_0)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_role_2[0])))
	if base.B2i32(v49 == int32(0))|base.B2i32(v49 != v52) != 0 {
		v70 = v49
		v71 = v52
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
	if v70-v71 != 0 {
		goto L9
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v55 = v25
	v56 = v46
	goto L23
L23:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v60 == int32(0) {
		v70 = v60
		v71 = v59
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v70 = v60
	v71 = v59
	goto L21
L25:
	;
	v63 = int32(1)
	if v60 == v59 {
		v55 = v55 + v63
		v56 = v56 + v63
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	return int32(1)
L28:
	;
	v77 = F_strlen(m, l0)
	mBase = m.M
	v82 = F_palloc(m, v77<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L15
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v100 == int32(0))|base.B2i32(v100 != v103) != 0 {
		v121 = v100
		v122 = v103
		goto L37
	} else {
		goto L38
	}
L31:
	;
	v84 = F_strlen(m, l0)
	mBase = m.M
	v85 = F_pg_mb2wchar_with_len(m, l0, v82, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v88 = int32(0)
	v91 = F_pg_regexec(m, v87, v82, v85, v88, v88, v88)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	F_pfree(m, v82)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	if v91 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	return int32(1)
L36:
	;
	if v121-v122 != 0 {
		goto L8
	} else {
		goto L43
	}
L37:
	;
	goto L36
L38:
	;
	v106 = v97
	v107 = l0
	goto L39
L39:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v111 == int32(0) {
		v121 = v111
		v122 = v110
		goto L37
	} else {
		goto L41
	}
L40:
	;
	v121 = v111
	v122 = v110
	goto L37
L41:
	;
	v114 = int32(1)
	if v111 == v110 {
		v106 = v106 + v114
		v107 = v107 + v114
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	return int32(1)
L44:
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
	if l1 == int32(_a_F_check_role_membership_authorization_0) {
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
										F_errmsg(m, int32(_a_F_check_role_membership_authorization_1), v7+int32(32))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return
										} else {
											v112 = int32(_a_F_check_role_membership_authorization_2)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v112
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v112
											F_errdetail(m, int32(_a_F_check_role_membership_authorization_3), v7+int32(16))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_check_role_membership_authorization_4), int32(2140), int32(_a_F_check_role_membership_authorization_5))
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
										F_errmsg(m, int32(_a_F_check_role_membership_authorization_6), v7-int32(-64))
										mBase = m.M
										v34 = m.ExcPending
										if v34 != 0 {
											return
										} else {
											v35 = int32(_a_F_check_role_membership_authorization_2)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v35
											*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v35
											F_errdetail(m, int32(_a_F_check_role_membership_authorization_7), v7+int32(48))
											mBase = m.M
											v43 = m.ExcPending
											if v43 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_check_role_membership_authorization_4), int32(2147), int32(_a_F_check_role_membership_authorization_5))
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
										F_errmsg(m, int32(_a_F_check_role_membership_authorization_1), v7+int32(96))
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
												*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = int32(_a_F_check_role_membership_authorization_8)
												F_errdetail(m, int32(_a_F_check_role_membership_authorization_9), v7+int32(80))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_check_role_membership_authorization_4), int32(2163), int32(_a_F_check_role_membership_authorization_5))
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
										F_errmsg(m, int32(_a_F_check_role_membership_authorization_6), v7+int32(128))
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
												*(*int32)(unsafe.Add(mBase, uint32(v7)+112)) = int32(_a_F_check_role_membership_authorization_8)
												F_errdetail(m, int32(_a_F_check_role_membership_authorization_10), v7+int32(112))
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_check_role_membership_authorization_4), int32(2170), int32(_a_F_check_role_membership_authorization_5))
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
				v95 = F_GetUserNameFromId(m, int32(_a_F_check_role_membership_authorization_0), int32(0))
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v95
					F_errmsg(m, int32(_a_F_check_role_membership_authorization_11), v7)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_role_membership_authorization_4), int32(2127), int32(_a_F_check_role_membership_authorization_5))
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(_a_F_get_role_oid_or_public_0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_role_oid_or_public[0])))
	if base.B2i32(v11 == v2)|base.B2i32(v11 != v14) != 0 {
		v32 = v11
		v33 = v14
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return v61
L2:
	;
	if v32-v33 == int32(0) {
		v61 = v2
		goto L1
	} else {
		goto L9
	}
L3:
	;
	goto L2
L4:
	;
	v17 = l0
	v18 = v8
	goto L5
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v22
		v33 = v21
		goto L3
	} else {
		goto L7
	}
L6:
	;
	v32 = v22
	v33 = v21
	goto L3
L7:
	;
	v25 = int32(1)
	if v22 == v21 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v38 = int32(0)
	v41 = F_GetSysCacheOid(m, int32(10), l0, v38, v38, v38)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v41 != 0 {
		v61 = v41
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(_a_F_get_role_oid_or_public_1), v6)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_get_role_oid_or_public_2), int32(_a_F_get_role_oid_or_public_3), int32(_a_F_get_role_oid_or_public_4))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
