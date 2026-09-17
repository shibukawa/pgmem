package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecFindRowMark(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
			F_errmsg_internal(m, int32(_a_F_ExecFindRowMark_0), v6)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_ExecFindRowMark_1), int32(2571), int32(_a_F_ExecFindRowMark_2))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if base.Ui32(v10) < base.Ui32(l1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
				F_errmsg_internal(m, int32(_a_F_ExecFindRowMark_0), v6)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_ExecFindRowMark_1), int32(2571), int32(_a_F_ExecFindRowMark_2))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v12 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
					F_errmsg_internal(m, int32(_a_F_ExecFindRowMark_0), v6)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_ExecFindRowMark_1), int32(2571), int32(_a_F_ExecFindRowMark_2))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v12+l1<<(uint(int32(2))%32)-int32(4))))
				if v20 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
						F_errmsg_internal(m, int32(_a_F_ExecFindRowMark_0), v6)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ExecFindRowMark_1), int32(2571), int32(_a_F_ExecFindRowMark_2))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					m.G0 = v6 + int32(16)
					return v20
				}
			}
		}
	}
}
func F_ExpandRowReference(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
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
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v11 != int32(6) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v151
L2:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v78 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L3:
	;
	v75 = F_get_expr_result_tupdesc(m, l1, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L28
	}
L4:
	;
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	if v14 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v19 = F_GetNSItemByRangeTablePosn(m, l0, v17, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v69 != int32(2249) {
		goto L3
	} else {
		goto L26
	}
L8:
	;
	return int32(0)
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v25 = F_expandNSItemAttrs(m, l0, v19, v24, v23)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v31 = F_expandNSItemVars(m, l0, v19, v24, v23, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	return v25
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v33 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v36 | int64(2)
	goto L17
L16:
	;
	goto L17
L17:
	;
	if v31 == int32(0) {
		v151 = v4
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v42 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return v31
L20:
	;
	goto L21
L21:
	;
	v49 = int32(0)
	goto L22
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v49<<(uint(int32(2))%32))))
	F_markVarForSelectPriv(m, l0, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L24
	}
L23:
	;
	return v31
L24:
	;
	v65 = v49 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v65 < v66 {
		v49 = v65
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v72 = F_expandRecordVariable(m, l0, l1)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v77 = v72
	goto L2
L28:
	;
	v77 = v75
	goto L2
L29:
	;
	return int32(0)
L30:
	;
	goto L31
L31:
	;
	v88 = v4
	v89 = v4
	goto L32
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v99 = v77 + v93<<(uint(int32(4))%32) + v88*int32(100)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+111)))
	if v100 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v151 = v140
	goto L1
L34:
	;
	v104 = F_palloc0(m, int32(24))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L8
	} else {
		goto L37
	}
L35:
	;
	v140 = v89
	goto L36
L36:
	;
	v143 = v88 + int32(1)
	if v143 != v78 {
		v88 = v143
		v89 = v140
		goto L32
	} else {
		goto L45
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = int32(25)
	v108 = F_copyObjectImpl(m, l1)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v111 = v88 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v104)+8)) = uint16(v111)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v108
	v115 = v99 + int32(20)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v104)+16)) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v104)+20)) = v120
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v122 + int32(1)
	v129 = F_pstrdup(m, v99+int32(24))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	v135 = v104
	goto L41
L41:
	;
	v136 = F_lappend(m, v89, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L8
	} else {
		goto L44
	}
L42:
	;
	v132 = F_makeTargetEntry(m, v104, base.I32_extend16_s(v122), v129, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v135 = v132
	goto L41
L44:
	;
	v140 = v136
	goto L36
L45:
	;
	goto L33
}
func F_add_row_identity_columns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+119)))
	switch v8 - int32(102) {
	case 0:
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
		v24 = F_GetFdwRoutineForRelation(m, l3, int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
			if v26 != 0 {
				m.T0[v26].(func(*base.Module, int32, int32, int32, int32))(m, l0, l1, l2, l3)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					if v22 == int32(2) {
						v39 = int32(0)
						v44 = F_makeVar(m, l1, v39, int32(2249), int32(-1), v39, v39)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_add_row_identity_var(m, l0, v44, l1, int32(_a_F_add_row_identity_columns_0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
						if v31 == int32(0) {
							return
						} else {
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+19)))
							if v34 != 0 {
								v39 = int32(0)
								v44 = F_makeVar(m, l1, v39, int32(2249), int32(-1), v39, v39)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									F_add_row_identity_var(m, l0, v44, l1, int32(_a_F_add_row_identity_columns_0))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+18)))
								if v35 != int32(1) {
									return
								} else {
									v39 = int32(0)
									v44 = F_makeVar(m, l1, v39, int32(2249), int32(-1), v39, v39)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										F_add_row_identity_var(m, l0, v44, l1, int32(_a_F_add_row_identity_columns_0))
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				if v22 == int32(2) {
					v39 = int32(0)
					v44 = F_makeVar(m, l1, v39, int32(2249), int32(-1), v39, v39)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_add_row_identity_var(m, l0, v44, l1, int32(_a_F_add_row_identity_columns_0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
					if v31 == int32(0) {
						return
					} else {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+19)))
						if v34 != 0 {
							v39 = int32(0)
							v44 = F_makeVar(m, l1, v39, int32(2249), int32(-1), v39, v39)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_add_row_identity_var(m, l0, v44, l1, int32(_a_F_add_row_identity_columns_0))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+18)))
							if v35 != int32(1) {
								return
							} else {
								v39 = int32(0)
								v44 = F_makeVar(m, l1, v39, int32(2249), int32(-1), v39, v39)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									F_add_row_identity_var(m, l0, v44, l1, int32(_a_F_add_row_identity_columns_0))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	default:
		return
	case 7, 10, 12:
		v11 = int32(-1)
		v14 = int32(0)
		v16 = F_makeVar(m, l1, v11, int32(27), v11, v14, v14)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_add_row_identity_var(m, l0, v16, l1, int32(_a_F_add_row_identity_columns_1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_equalRowTypes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v7 != v8 {
		v84 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v84
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v10 != v11 {
		v84 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v7 <= int32(0) {
		v84 = int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v17 = v7 << (uint(int32(4)) % 32)
	v19 = int32(20)
	v25 = int32(0)
	goto L5
L5:
	;
	v32 = v25 * int32(100)
	v33 = l0 + v17 + v19 + v32
	v34 = int32(4)
	v35 = v33 + v34
	v36 = v32 + (l1 + v17 + v19)
	v38 = v36 + v34
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if base.B2i32(v41 == int32(0))|base.B2i32(v41 != v44) != 0 {
		v62 = v41
		v63 = v44
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v84 = int32(0)
	goto L1
L7:
	;
	goto L6
L8:
	;
	if v62-v63 != 0 {
		goto L7
	} else {
		goto L15
	}
L9:
	;
	goto L8
L10:
	;
	v47 = v35
	v48 = v38
	goto L11
L11:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v52 == int32(0) {
		v62 = v52
		v63 = v51
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v62 = v52
	v63 = v51
	goto L9
L13:
	;
	v55 = int32(1)
	if v52 == v51 {
		v47 = v47 + v55
		v48 = v48 + v55
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
	if v65 != v66 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	if v68 != v69 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v33)+96))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v36)+96))
	if v71 != v72 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+91)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+91)))
	if v74 != v75 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v77 = int32(1)
	v79 = v25 + v77
	if v7 != v79 {
		v25 = v79
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v84 = v77
	goto L1
}
func F_transformInsertRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v316 int32
	_ = v316
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	v7 = int32(0)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v12 = v11
	goto L3
L2:
	;
	v12 = v7
	goto L3
L3:
	;
	if l3 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v15 = v13
	goto L6
L5:
	;
	v15 = int32(0)
	goto L6
L6:
	;
	if v15 < v12 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if l2 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	return int32(0)
L11:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_errmsg(m, int32(_a_F_transformInsertRow_0), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if l3 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v33 = v31
	goto L16
L15:
	;
	v33 = int32(0)
	goto L16
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30+v33<<(uint(int32(2))%32))))
	v38 = F_exprLocation(m, v37)
	mBase = m.M
	F_parser_errposition(m, l0, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_transformInsertRow_1), int32(1074), int32(_a_F_transformInsertRow_2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v248 = int32(0)
	v255 = v7
	goto L86
L20:
	;
	if l1 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v50 = v49
	goto L23
L22:
	;
	v50 = int32(0)
	goto L23
L23:
	;
	if l3 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v53 = v51
	goto L26
L25:
	;
	v53 = int32(0)
	goto L26
L26:
	;
	if v53 <= v50 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	F_errmsg(m, int32(_a_F_transformInsertRow_3), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	if l1 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v66 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v229 = v7
	goto L33
L33:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231+v229<<(uint(int32(2))%32))))
	v236 = F_exprLocation(m, v235)
	mBase = m.M
	F_parser_errposition(m, l0, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L10
	} else {
		goto L84
	}
L34:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v229 = v225
	goto L33
L35:
	;
	v69 = int32(-1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v71 == int32(0) {
		v206 = v69
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if l3 != 0 {
		goto L79
	} else {
		goto L80
	}
L37:
	;
	v212 = v206
	goto L36
L38:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v74 != int32(6) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v206 = v205
	goto L37
L40:
	;
	if v74 != int32(36) {
		v206 = v69
		goto L37
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+8)))
	if v81 <= int32(0) {
		v206 = v69
		goto L37
	} else {
		goto L45
	}
L43:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v79 != 0 {
		v204 = v79
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v212 = int32(0)
	goto L36
L45:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	if v84 != int32(2249) {
		v206 = v69
		goto L37
	} else {
		goto L46
	}
L46:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v71)+28))
	v89 = int32(0)
	if v88 <= v89 {
		v136 = l0
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	if v150 != int32(1) {
		v206 = v69
		goto L37
	} else {
		goto L61
	}
L48:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143+v87<<(uint(int32(2))%32)-int32(4))))
	goto L47
L49:
	;
	v95 = v88 & int32(7)
	if v95 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if base.Ui32(v88) < base.Ui32(int32(8)) {
		v136 = v110
		goto L48
	} else {
		goto L57
	}
L51:
	;
	v110 = l0
	v113 = v88
	goto L50
L52:
	;
	goto L53
L53:
	;
	v98 = l0
	v101 = v88
	v103 = v89
	goto L54
L54:
	;
	v104 = int32(1)
	v105 = v101 - v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v108 = v103 + v104
	if v108 != v95 {
		v98 = v106
		v101 = v105
		v103 = v108
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v110 = v106
	v113 = v105
	goto L50
L56:
	;
	goto L55
L57:
	;
	v118 = v110
	v121 = v113
	goto L58
L58:
	;
	v124 = int32(8)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v124 < v121 {
		v118 = v133
		v121 = v121 - v124
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v136 = v133
	goto L48
L60:
	;
	goto L59
L61:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149)+36))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+76))
	if v154 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if v192 == int32(0) {
		v206 = v69
		goto L37
	} else {
		goto L75
	}
L63:
	;
	goto L62
L64:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v158 <= int32(0) {
		v192 = int32(0)
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v192 = int32(0)
	goto L63
L67:
	;
	v161 = int32(0)
	if v161 < v158 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v164 = v158
	goto L70
L69:
	;
	v164 = v161
	goto L70
L70:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v169 = int32(0)
	goto L71
L71:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v165+v169<<(uint(int32(2))%32))))
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177)+8)))
	if v178 == v81&int32(_a_F_transformInsertRow_4) {
		v192 = v177
		goto L63
	} else {
		goto L73
	}
L72:
	;
	goto L66
L73:
	;
	v181 = v169 + int32(1)
	if v181 != v164 {
		v169 = v181
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+26)))
	if v196 != 0 {
		v206 = v69
		goto L37
	} else {
		goto L76
	}
L76:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	if v198 != int32(36) {
		v206 = v69
		goto L37
	} else {
		goto L77
	}
L77:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v201 != 0 {
		v204 = v201
		goto L39
	} else {
		goto L78
	}
L78:
	;
	v212 = int32(0)
	goto L36
L79:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v215 = v213
	goto L81
L80:
	;
	v215 = int32(0)
	goto L81
L81:
	;
	if v212 != v215 {
		goto L34
	} else {
		goto L82
	}
L82:
	;
	F_errhint(m, int32(_a_F_transformInsertRow_5), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	goto L34
L84:
	;
	F_errfinish(m, int32(_a_F_transformInsertRow_1), int32(1096), int32(_a_F_transformInsertRow_2))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v256 = int32(0)
	if l1 == v256 {
		v266 = v256
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v267 = int32(0)
	if l3 == v267 {
		v277 = v267
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v260 <= v248 {
		v266 = int32(0)
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v266 = v262 + v248<<(uint(int32(2))%32)
	goto L88
L91:
	;
	if l4 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v271 <= v248 {
		v277 = int32(0)
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v277 = v273 + v248<<(uint(int32(2))%32)
	goto L91
L94:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v288+v248<<(uint(int32(2))%32))))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v296)+8))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v296)+16))
	v304 = F_transformAssignedExpr(m, l0, v294, int32(15), v297, v301, v302, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L10
	} else {
		goto L102
	}
L95:
	;
	v278 = int32(0)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.B2i32(v277 == v278)|(base.B2i32(v266 == v278)|base.B2i32(v282 <= v248)) == v278 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v292 = int32(0)
	goto L97
L97:
	;
	return v292
L98:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v288 != 0 {
		goto L94
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v292 = v255
	goto L97
L101:
	;
	goto L100
L102:
	;
	if l5 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v358 = F_lappend(m, v255, v353)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L10
	} else {
		goto L122
	}
L104:
	;
	v353 = v304
	goto L103
L105:
	;
	goto L106
L106:
	;
	if v304 == int32(0) {
		v353 = v304
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v316 = v304
	goto L108
L108:
	;
	v326 = v316
	goto L111
L109:
	;
	v353 = int32(0)
	goto L103
L110:
	;
	if v344 != 0 {
		v316 = v344
		goto L108
	} else {
		goto L121
	}
L111:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	if v330 != int32(55) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v326)+36))
	if v341 == int32(0) {
		v353 = v316
		goto L103
	} else {
		goto L120
	}
L113:
	;
	v334 = v330 - int32(14)
	if v334 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	v326 = v340
	goto L111
L115:
	;
	goto L112
L116:
	;
	if v334 != int32(12) {
		v353 = v316
		goto L103
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	goto L115
L119:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v326)+8))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	v344 = v339
	goto L110
L120:
	;
	v344 = v341
	goto L110
L121:
	;
	goto L109
L122:
	;
	v248 = v248 + int32(1)
	v255 = v358
	goto L86
}
