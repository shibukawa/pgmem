package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsBinaryCoercible(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_IsBinaryCoercibleWithCast(m, l0, l1, v6+int32(12))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v10
	}
}
func F_binary_upgrade_add_sub_rel_state(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_add_sub_rel_state[0])))
	if v8 != 0 {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v9 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_binary_upgrade_add_sub_rel_state_0), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_binary_upgrade_add_sub_rel_state_1), int32(338), int32(_a_F_binary_upgrade_add_sub_rel_state_2))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v10 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_binary_upgrade_add_sub_rel_state_0), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_binary_upgrade_add_sub_rel_state_1), int32(338), int32(_a_F_binary_upgrade_add_sub_rel_state_2))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
				if v11 == int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_binary_upgrade_add_sub_rel_state_0), int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_binary_upgrade_add_sub_rel_state_1), int32(338), int32(_a_F_binary_upgrade_add_sub_rel_state_2))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v15 = F_pg_detoast_datum_packed(m, v14)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = F_text_to_cstring(m, v15)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return int32(0)
						} else {
							v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
							if v23 == int32(0) {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
								v28 = v27
							} else {
								v28 = int64(0)
							}
							v31 = F_table_open(m, int32(_a_F_binary_upgrade_add_sub_rel_state_3), int32(3))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v34 = F_get_subscription_oid(m, v19, int32(0))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									v37 = F_relation_open(m, v22, int32(1))
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int32(0)
									} else {
										F_AddSubscriptionRelState(m, v34, v22, base.I32_extend8_s(v21), v28, int32(0))
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return int32(0)
										} else {
											F_relation_close(m, v37, int32(1))
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v31, int32(3))
												mBase = m.M
												v48 = m.ExcPending
												if v48 != 0 {
													return int32(0)
												} else {
													return int32(0)
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
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_binary_upgrade_add_sub_rel_state_4), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_binary_upgrade_add_sub_rel_state_1), int32(334), int32(_a_F_binary_upgrade_add_sub_rel_state_2))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
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
func F_binary_upgrade_create_empty_extension(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_create_empty_extension[0])))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L39
	}
L2:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v17 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L35
	}
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v18 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v19 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v20 == int32(1) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = F_pg_detoast_datum_packed(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v29 = F_pg_detoast_datum_packed(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v33 = F_pg_detoast_datum_packed(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v35 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v39 = v38
	goto L15
L14:
	;
	v39 = v2
	goto L15
L15:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
	if v40 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v44 = v43
	goto L18
L17:
	;
	v44 = v2
	goto L18
L18:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v45 != 0 {
		v88 = v2
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v94 = F_text_to_cstring(m, v24)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L30
	}
L20:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v47 = F_pg_detoast_datum(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	F_deconstruct_array_builtin(m, v47, int32(25), v13, int32(0), v13+int32(12))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v55 <= int32(0) {
		v88 = v2
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v62 = v2
	v63 = v2
	goto L24
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v63<<(uint(int32(2))%32))))
	v73 = F_text_to_cstring(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L26
	}
L25:
	;
	v88 = v78
	goto L19
L26:
	;
	v76 = F_get_extension_oid(m, v73, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v78 = F_lappend_oid(m, v62, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v81 = v63 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v81 < v82 {
		v62 = v78
		v63 = v81
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_create_empty_extension[1]))
	v98 = F_text_to_cstring(m, v29)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	v101 = F_get_namespace_oid(m, v98, int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v105 = F_text_to_cstring(m, v33)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	F_InsertExtensionTuple(m, v13, v94, v97, v101, base.B2i32(v31 != int32(0)), v105, v39, v44, v88)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	m.G0 = v13 + int32(16)
	return int32(0)
L35:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(_a_F_binary_upgrade_create_empty_extension_0), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_binary_upgrade_create_empty_extension_1), int32(194), int32(_a_F_binary_upgrade_create_empty_extension_2))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	F_errmsg_internal(m, int32(_a_F_binary_upgrade_create_empty_extension_3), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_binary_upgrade_create_empty_extension_1), int32(201), int32(_a_F_binary_upgrade_create_empty_extension_2))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_binary_upgrade_set_next_array_pg_type_oid(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13871(m, l0, int32(_a_F_binary_upgrade_set_next_array_pg_type_oid_0), int32(_a_F_binary_upgrade_set_next_array_pg_type_oid_1), int32(68))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_binary_upgrade_set_record_init_privs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_set_record_init_privs[0])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_binary_upgrade_set_record_init_privs_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_binary_upgrade_set_record_init_privs_1), int32(253), int32(_a_F_binary_upgrade_set_record_init_privs_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v26 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_set_record_init_privs[1])) = uint8(base.B2i32(v25 != v26))
		return v26
	}
}
func F_read_binary_file(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	if l2 < int64(1073741820) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L7
	} else {
		goto L67
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L7
	} else {
		goto L63
	}
L3:
	;
	v15 = F_AllocateFile(m, l0, int32(_a_F_read_binary_file_0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L7
	} else {
		goto L59
	}
L6:
	;
	m.G0 = v10 - int32(-64)
	return v176
L7:
	;
	return int32(0)
L8:
	;
	if v15 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if l3 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if l1 < int64(0) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_read_binary_file[0]))
	if v22 == int32(44) {
		v176 = int32(0)
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L7
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	F_errmsg(m, int32(_a_F_read_binary_file_1), v10)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_read_binary_file_2), int32(124), int32(_a_F_read_binary_file_3))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
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
	v44 = int32(2)
	goto L22
L21:
	;
	v44 = int32(0)
	goto L22
L22:
	;
	v45 = F___fseeko_unlocked(m, v15, l1, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	if v45 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	if int64(0) <= l2 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	goto L56
L26:
	;
	v49 = base.I32_wrap_i64(l2)
	v52 = F_palloc(m, v49+int32(4))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_initStringInfo(m, v8+int32(-16))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L31
	}
L29:
	;
	v57 = F_fread(m, v52+int32(4), int32(1), v49, v15)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v154 = v57
	v157 = v52
	goto L25
L31:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v64 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v63 + v64
	v67 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	goto L33
L32:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v154 = v146
	v157 = v150
	goto L25
L33:
	;
	if int32(base.Ui32(v68)>>(uint(v64)%32))&int32(1) != 0 {
		v146 = v67
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v76 = v67
	goto L35
L35:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	goto L37
L36:
	;
	v146 = v135
	goto L32
L37:
	;
	if int32(base.Ui32(v80)>>(uint(int32(5))%32))&int32(1) != 0 {
		v146 = v76
		goto L32
	} else {
		goto L38
	}
L38:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	if v85 == int32(1073741822) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v90 = int32(1)
	v92 = F_fread(m, v8+int32(-17), v90, v90, v15)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_enlargeStringInfo(m, v8+int32(-16), int32(_a_F_read_binary_file_4))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L52
	}
L42:
	;
	if v92 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	goto L46
L44:
	;
	goto L45
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L7
	} else {
		goto L48
	}
L46:
	;
	if int32(base.Ui32(v96)>>(uint(int32(4))%32))&int32(1) != 0 {
		v146 = v76
		goto L32
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(_a_F_read_binary_file_5), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_read_binary_file_2), int32(171), int32(_a_F_read_binary_file_3))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	v130 = F_fread(m, v122+v123, int32(1), v126+(v122^int32(-1)), v15)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v130 + v132
	v135 = v76 + v130
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	goto L54
L54:
	;
	if int32(base.Ui32(v136)>>(uint(int32(4))%32))&int32(1) == int32(0) {
		v76 = v135
		goto L35
	} else {
		goto L55
	}
L55:
	;
	goto L36
L56:
	;
	if int32(base.Ui32(v158)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v154<<(uint(int32(2))%32) + int32(16)
	v168 = F_FreeFile(m, v15)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	v176 = v157
	goto L6
L59:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(_a_F_read_binary_file_6), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_read_binary_file_2), int32(114), int32(_a_F_read_binary_file_3))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
	F_errmsg(m, int32(_a_F_read_binary_file_7), v8+int32(-32))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_read_binary_file_2), int32(131), int32(_a_F_read_binary_file_3))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
	F_errmsg(m, int32(_a_F_read_binary_file_8), v8+int32(-48))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_read_binary_file_2), int32(197), int32(_a_F_read_binary_file_3))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
