package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_replorigin_by_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(58), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v19 = F_text_to_cstring(m, v14+v15+int32(4))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
				F_ReleaseCatCache(m, v10)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return base.B2i32(v10 != int32(0))
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
			m.G0 = v6 + int32(16)
			return base.B2i32(v10 != int32(0))
		}
	}
}
func F_replorigin_check_prerequisites(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_check_prerequisites[0]))
	if v2 != 0 {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_replorigin_check_prerequisites[1])))
		if v5 == int32(1) {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_check_prerequisites[2]))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+308))
			v13 = base.B2i32(v11 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _c_F_replorigin_check_prerequisites[1])) = uint8(v13)
			v15 = v13
		} else {
			v15 = int32(0)
		}
		if v15 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				F_errcode(m, int32(100663618))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_replorigin_check_prerequisites_0), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_replorigin_check_prerequisites_1), int32(217), int32(_a_F_replorigin_check_prerequisites_2))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_replorigin_check_prerequisites_3), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_replorigin_check_prerequisites_1), int32(212), int32(_a_F_replorigin_check_prerequisites_2))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
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
func F_replorigin_create(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v7 = m.G0
	v9 = v7 - int32(176)
	m.G0 = v9
	v11 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(v11) < base.Ui32(int32(513)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L33
	}
L2:
	;
	v14 = F_cstring_to_text(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L28
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+104)) = int32(4)
	v23 = F_table_open(m, int32(_a_F_replorigin_create_0), int32(7))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v30 = int64(1)
	goto L9
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = base.I64_extend_i32_u(v14)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v30
	v65 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+46)) = uint16(v65)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v72 = F_heap_form_tuple(m, v67, v9+int32(16), v9+int32(46))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L22
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_create[0]))
	if v33 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_relation_close(m, v23, int32(7))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L21
	}
L11:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v37 = v9 + int32(48)
	F_ScanKeyInit(m, v37, int32(1), int32(3), int32(184), v30)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v44 = int32(1)
	v48 = F_systable_beginscan(m, v23, int32(_a_F_replorigin_create_1), v44, v9+int32(104), v44, v37)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v50 = F_systable_getnext(m, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_systable_endscan(m, v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	if v50 == int32(0) {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v57 = v30 + int64(1)
	if v57 != int64(65535) {
		v30 = v57
		goto L9
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	goto L1
L22:
	;
	F_CatalogTupleInsert(m, v23, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	F_relation_close(m, v23, int32(7))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	if v72 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_pfree(m, v72)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	m.G0 = v9 + int32(176)
	return base.I32_wrap_i64(v30) & int32(_a_F_replorigin_create_2)
L28:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_errmsg(m, int32(_a_F_replorigin_create_3), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(512)
	v106 = F_errdetail(m, int32(_a_F_replorigin_create_4), v9)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_replorigin_create_5), int32(294), int32(_a_F_replorigin_create_6))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(_a_F_replorigin_create_7), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_replorigin_create_5), int32(376), int32(_a_F_replorigin_create_6))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_replorigin_identify(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if l0 == int32(16) {
		v6 = int32(_a_F_replorigin_identify_0)
	} else {
		v6 = int32(0)
	}
	if l0 != 0 {
		v8 = v6
	} else {
		v8 = int32(_a_F_replorigin_identify_1)
	}
	return v8
}
func F_replorigin_session_reset(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[0]))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[1]))
		if v9 == v11 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
			if int32(2) <= v13 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[0]))
						v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75))))
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v76
						F_errmsg(m, int32(_a_F_replorigin_session_reset_0), v5)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v83 = F_errdetail(m, int32(_a_F_replorigin_session_reset_1), int32(0))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								F_errhint(m, int32(_a_F_replorigin_session_reset_2), int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_replorigin_session_reset_3), int32(1323), int32(_a_F_replorigin_session_reset_4))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[2]))
				v21 = F_LWLockAcquire(m, v17+int32(_a_F_replorigin_session_reset_5), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[0]))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
					v27 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[1]))
					if v25 == v27 {
						*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = int32(0)
					} else {
					}
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v31 - int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[0])) = int32(0)
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[2]))
					F_LWLockRelease(m, v39+int32(_a_F_replorigin_session_reset_5))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_ConditionVariableBroadcast(m, v24+int32(32))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							m.G0 = v5 + int32(16)
							return
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[2]))
			v21 = F_LWLockAcquire(m, v17+int32(_a_F_replorigin_session_reset_5), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[0]))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[1]))
				if v25 == v27 {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = int32(0)
				} else {
				}
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v31 - int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[0])) = int32(0)
				v39 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_reset[2]))
				F_LWLockRelease(m, v39+int32(_a_F_replorigin_session_reset_5))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_ConditionVariableBroadcast(m, v24+int32(32))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_replorigin_session_reset_6), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_replorigin_session_reset_3), int32(1308), int32(_a_F_replorigin_session_reset_4))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
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
func F_replorigin_session_setup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	v1 = l0
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_replorigin_session_setup[0])))
	if v16 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_on_shmem_exit(m, int32(1081), int64(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_setup[1]))
	if v27 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	return
L5:
	;
	v24 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_replorigin_session_setup[0])) = uint8(v24)
	goto L3
L6:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v241)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v241)+28)) = v248 + int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_setup[2]))
	F_LWLockRelease(m, v253+int32(_a_F_replorigin_session_setup_0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L66
	}
L7:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_setup[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+24)) = v236
	v241 = v228
	goto L6
L8:
	;
	if l1 != 0 {
		v241 = v218
		goto L6
	} else {
		goto L65
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L60
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L56
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L52
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L48
	}
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_setup[2]))
	v35 = F_LWLockAcquire(m, v31+int32(_a_F_replorigin_session_setup_0), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L44
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_setup[4]))
	if v38 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_setup[1]))
	if v108 != 0 {
		v218 = v108
		goto L8
	} else {
		goto L41
	}
L18:
	;
	v101 = int32(-1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_setup[5]))
	v50 = int32(-1)
	v51 = v3
	goto L21
L21:
	;
	v57 = v43 + v51<<(uint(int32(6))%32)
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57))))
	if v58 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v101 = v93
	goto L17
L23:
	;
	v95 = v51 + int32(1)
	if v95 != v38 {
		v50 = v93
		v51 = v95
		goto L21
	} else {
		goto L40
	}
L24:
	;
	if v50 == int32(-1) {
		v93 = v51
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v1 != v58 {
		v93 = v50
		goto L23
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	if l1 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_setup[1])) = v57
	v218 = v57
	goto L8
L30:
	;
	if v65 != 0 {
		goto L12
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v65 != l1 {
		goto L11
	} else {
		goto L39
	}
L33:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	if v68 <= int32(0) {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v78
	F_errmsg(m, int32(_a_F_replorigin_session_setup_1), v13+int32(32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_replorigin_session_setup_2), int32(1220), int32(_a_F_replorigin_session_setup_3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
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
	goto L29
L40:
	;
	goto L22
L41:
	;
	if l1 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	if v101 == int32(-1) {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_setup[5]))
	v116 = v113 + v101<<(uint(int32(6))%32)
	*(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_setup[1])) = v116
	*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v1)
	v228 = v116
	goto L7
L44:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(_a_F_replorigin_session_setup_4), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_replorigin_session_setup_2), int32(1173), int32(_a_F_replorigin_session_setup_3))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57))))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v142
	F_errmsg(m, int32(_a_F_replorigin_session_setup_5), v13+int32(48))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_replorigin_session_setup_2), int32(1206), int32(_a_F_replorigin_session_setup_3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
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
	F_errcode(m, int32(100663621))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v163
	F_errmsg(m, int32(_a_F_replorigin_session_setup_6), v13-int32(-64))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_replorigin_session_setup_2), int32(1233), int32(_a_F_replorigin_session_setup_3))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	F_errmsg(m, int32(_a_F_replorigin_session_setup_7), v13+int32(16))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_replorigin_session_setup_2), int32(1252), int32(_a_F_replorigin_session_setup_3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(_a_F_replorigin_session_setup_8))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v1
	F_errmsg(m, int32(_a_F_replorigin_session_setup_9), v13)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errhint(m, int32(_a_F_replorigin_session_setup_10), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_replorigin_session_setup_2), int32(1260), int32(_a_F_replorigin_session_setup_3))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	v228 = v218
	goto L7
L66:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_setup[1]))
	F_ConditionVariableBroadcast(m, v259+int32(32))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	m.G0 = v13 + int32(80)
	return
}
