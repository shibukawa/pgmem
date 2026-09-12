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
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[526])))
	if v8 != 0 {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v9 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(429391), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(481816), int32(338), int32(343894))
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
					F_errmsg_internal(m, int32(429391), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(481816), int32(338), int32(343894))
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
						F_errmsg_internal(m, int32(429391), int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(481816), int32(338), int32(343894))
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
							v31 = F_table_open(m, int32(6100), int32(3))
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
												F_sequence_close(m, v31, int32(3))
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
				F_errmsg(m, int32(403318), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(481816), int32(334), int32(343894))
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[526])))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
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
	v119 = m.ExcPending
	if v119 != 0 {
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
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0-int32(-64)))))
	if v42 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v46 = v45
	goto L18
L17:
	;
	v46 = v2
	goto L18
L18:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v47 != 0 {
		v90 = v2
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v96 = F_text_to_cstring(m, v24)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L30
	}
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v49 = F_pg_detoast_datum(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	F_deconstruct_array_builtin(m, v49, int32(25), v13, int32(0), v13+int32(12))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v57 <= int32(0) {
		v90 = v2
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v64 = v2
	v65 = v2
	goto L24
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v65<<(uint(int32(2))%32))))
	v75 = F_text_to_cstring(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L26
	}
L25:
	;
	v90 = v80
	goto L19
L26:
	;
	v78 = F_get_extension_oid(m, v75, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v80 = F_lappend_oid(m, v64, v78)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v83 = v65 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v83 < v84 {
		v64 = v80
		v65 = v83
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v100 = F_text_to_cstring(m, v29)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	v103 = F_get_namespace_oid(m, v100, int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v107 = F_text_to_cstring(m, v33)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	F_InsertExtensionTuple(m, v13, v96, v99, v103, base.B2i32(v31 != int32(0)), v107, v39, v46, v90)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
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
	v122 = m.ExcPending
	if v122 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(403318), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(481816), int32(194), int32(265793))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
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
	F_errmsg_internal(m, int32(429321), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(481816), int32(201), int32(265793))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
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
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[526])))
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
				F_errmsg(m, int32(403318), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(481816), int32(68), int32(425125))
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
		*(*int32)(unsafe.Add(mBase, _consts[1119])) = v25
		return int32(0)
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
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[526])))
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
				F_errmsg(m, int32(403318), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(481816), int32(253), int32(111601))
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
		*(*uint8)(unsafe.Add(mBase, _consts[296])) = uint8(base.B2i32(v25 != v26))
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
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
	v242 = m.ExcPending
	if v242 != 0 {
		goto L7
	} else {
		goto L84
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L7
	} else {
		goto L80
	}
L3:
	;
	v15 = F_AllocateFile(m, l0, int32(225897))
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
	v209 = m.ExcPending
	if v209 != 0 {
		goto L7
	} else {
		goto L76
	}
L6:
	;
	m.G0 = v10 - int32(-64)
	return v201
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
	v45 = F___fseeko(m, v15, l1, base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(62))%64)))&int32(2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L20
	}
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v22 == int32(44) {
		v201 = int32(0)
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
	F_errmsg(m, int32(287790), v10)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(487752), int32(124), int32(377531))
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
	if v45 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	if int64(0) <= l2 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	if v178 < int32(0) {
		goto L71
	} else {
		goto L72
	}
L23:
	;
	v49 = base.I32_wrap_i64(l2)
	v52 = F_palloc(m, v49+int32(4))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_initStringInfo(m, v8+int32(-16))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L28
	}
L26:
	;
	v57 = F_fread(m, v52+int32(4), int32(1), v49, v15)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v174 = v57
	v177 = v52
	goto L22
L28:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v63 + int32(4)
	v67 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	if v68 < v67 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v174 = v166
	v177 = v170
	goto L22
L30:
	;
	if int32(base.Ui32(v73)>>(uint(int32(4))%32))&int32(1) != 0 {
		v166 = v67
		goto L29
	} else {
		goto L35
	}
L31:
	;
	goto L30
L32:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v73 = v71
	goto L31
L33:
	;
	goto L34
L34:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v73 = v72
	goto L31
L35:
	;
	v81 = v67
	goto L36
L36:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	if v85 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v166 = v150
	goto L29
L38:
	;
	if int32(base.Ui32(v90)>>(uint(int32(5))%32))&int32(1) != 0 {
		v166 = v81
		goto L29
	} else {
		goto L43
	}
L39:
	;
	goto L38
L40:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v90 = v88
	goto L39
L41:
	;
	goto L42
L42:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v90 = v89
	goto L39
L43:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	if v95 == int32(1073741822) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v100 = int32(1)
	v102 = F_fread(m, v8+int32(-17), v100, v100, v15)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_enlargeStringInfo(m, v8+int32(-16), int32(4096))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L61
	}
L47:
	;
	if v102 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	if v106 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	goto L50
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L57
	}
L51:
	;
	if int32(base.Ui32(v111)>>(uint(int32(4))%32))&int32(1) != 0 {
		v166 = v81
		goto L29
	} else {
		goto L56
	}
L52:
	;
	goto L51
L53:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v111 = v109
	goto L52
L54:
	;
	goto L55
L55:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v111 = v110
	goto L52
L56:
	;
	goto L50
L57:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(391670), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(487752), int32(171), int32(377531))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	v145 = F_fread(m, v137+v138, int32(1), v141+(v137^int32(-1)), v15)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v145 + v147
	v150 = v81 + v145
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	if v151 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	if int32(base.Ui32(v156)>>(uint(int32(4))%32))&int32(1) == int32(0) {
		v81 = v150
		goto L36
	} else {
		goto L68
	}
L64:
	;
	goto L63
L65:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v156 = v154
	goto L64
L66:
	;
	goto L67
L67:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v156 = v155
	goto L64
L68:
	;
	goto L37
L69:
	;
	if int32(base.Ui32(v183)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L1
	} else {
		goto L74
	}
L70:
	;
	goto L69
L71:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v183 = v181
	goto L70
L72:
	;
	goto L73
L73:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v183 = v182
	goto L70
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v174<<(uint(int32(2))%32) + int32(16)
	v193 = F_FreeFile(m, v15)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	v201 = v177
	goto L6
L76:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	F_errmsg(m, int32(391692), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(487752), int32(114), int32(377531))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L7
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
	F_errmsg(m, int32(292118), v8+int32(-32))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(487752), int32(131), int32(377531))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
	F_errmsg(m, int32(293069), v8+int32(-48))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(487752), int32(197), int32(377531))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
