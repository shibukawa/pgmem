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
			F_errmsg_internal(m, int32(40103), v6)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493487), int32(2571), int32(313350))
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
				F_errmsg_internal(m, int32(40103), v6)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493487), int32(2571), int32(313350))
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
					F_errmsg_internal(m, int32(40103), v6)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493487), int32(2571), int32(313350))
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
						F_errmsg_internal(m, int32(40103), v6)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(493487), int32(2571), int32(313350))
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v12 != int32(6) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v153
L2:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v80 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L3:
	;
	v77 = F_get_expr_result_tupdesc(m, l1, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L28
	}
L4:
	;
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	if v15 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v20 = F_GetNSItemByRangeTablePosn(m, l0, v18, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v71 != int32(2249) {
		goto L3
	} else {
		goto L26
	}
L8:
	;
	return int32(0)
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v26 = F_expandNSItemAttrs(m, l0, v20, v25, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v32 = F_expandNSItemVars(m, l0, v20, v25, v24, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	return v26
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v34 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v29)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v37 | int64(2)
	goto L17
L16:
	;
	goto L17
L17:
	;
	if v32 == int32(0) {
		v153 = v4
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v43 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return v32
L20:
	;
	goto L21
L21:
	;
	v51 = int32(0)
	goto L22
L22:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v51<<(uint(int32(2))%32))))
	F_markVarForSelectPriv(m, l0, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L24
	}
L23:
	;
	return v32
L24:
	;
	v67 = v51 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v67 < v68 {
		v51 = v67
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v74 = F_expandRecordVariable(m, l0, l1)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v79 = v74
	goto L2
L28:
	;
	v79 = v77
	goto L2
L29:
	;
	return int32(0)
L30:
	;
	goto L31
L31:
	;
	v91 = v4
	v93 = v4
	goto L32
L32:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v104 = v79 + int32(20) + v98<<(uint(int32(4))%32) + v91*int32(100)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+91)))
	if v105 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v153 = v142
	goto L1
L34:
	;
	v109 = F_palloc0(m, int32(24))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L37
	}
L35:
	;
	v142 = v93
	goto L36
L36:
	;
	v145 = v91 + int32(1)
	if v145 != v80 {
		v91 = v145
		v93 = v142
		goto L32
	} else {
		goto L45
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(25)
	v113 = F_copyObjectImpl(m, l1)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v116 = v91 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)) = uint16(v116)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v113
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v104)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v104)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v104)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v123
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v125 + int32(1)
	v132 = F_pstrdup(m, v104+int32(4))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	v138 = v109
	goto L41
L41:
	;
	v139 = F_lappend(m, v93, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L8
	} else {
		goto L44
	}
L42:
	;
	v135 = F_makeTargetEntry(m, v109, base.I32_extend16_s(v125), v132, int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v138 = v135
	goto L41
L44:
	;
	v142 = v139
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
							F_add_row_identity_var(m, l0, v44, l1, int32(30013))
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
									F_add_row_identity_var(m, l0, v44, l1, int32(30013))
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
										F_add_row_identity_var(m, l0, v44, l1, int32(30013))
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
						F_add_row_identity_var(m, l0, v44, l1, int32(30013))
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
								F_add_row_identity_var(m, l0, v44, l1, int32(30013))
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
									F_add_row_identity_var(m, l0, v44, l1, int32(30013))
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
			F_add_row_identity_var(m, l0, v16, l1, int32(430426))
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
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v7 != v8 {
		v83 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v83
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v10 != v11 {
		v83 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v7 <= int32(0) {
		v83 = int32(1)
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
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v42 == int32(0) {
		v61 = v41
		v62 = v42
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v83 = int32(0)
	goto L1
L7:
	;
	goto L6
L8:
	;
	if v62-v61 != 0 {
		goto L7
	} else {
		goto L16
	}
L9:
	;
	goto L8
L10:
	;
	if v41 != v42 {
		v61 = v41
		v62 = v42
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v46 = v35
	v47 = v38
	goto L12
L12:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
	if v51 == int32(0) {
		v61 = v50
		v62 = v51
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v61 = v50
	v62 = v51
	goto L9
L14:
	;
	v54 = int32(1)
	if v50 == v51 {
		v46 = v46 + v54
		v47 = v47 + v54
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
	if v64 != v65 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	if v67 != v68 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v33)+96))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v36)+96))
	if v70 != v71 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+91)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+91)))
	if v73 != v74 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v76 = int32(1)
	v78 = v25 + v76
	if v7 != v78 {
		v25 = v78
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v83 = v76
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
	var v102 int32
	_ = v102
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
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
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
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
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
	F_errmsg(m, int32(146768), int32(0))
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
	F_errfinish(m, int32(495167), int32(1074), int32(31750))
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
	F_errmsg(m, int32(145352), int32(0))
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
	v230 = v7
	goto L33
L33:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231+v230<<(uint(int32(2))%32))))
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
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v230 = v226
	goto L33
L35:
	;
	v69 = int32(-1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v71 == int32(0) {
		v208 = v69
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
	v214 = v208
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
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v208 = v207
	goto L37
L40:
	;
	if v74 != int32(36) {
		v208 = v69
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
		v208 = v69
		goto L37
	} else {
		goto L45
	}
L43:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v79 != 0 {
		v206 = v79
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v214 = int32(0)
	goto L36
L45:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	if v84 != int32(2249) {
		v208 = v69
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
		v138 = l0
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	if v152 != int32(1) {
		v208 = v69
		goto L37
	} else {
		goto L61
	}
L48:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v145+v87<<(uint(int32(2))%32)-int32(4))))
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
		v138 = v110
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
	v102 = v89
	goto L54
L54:
	;
	v104 = int32(1)
	v105 = v101 - v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v108 = v102 + v104
	if v108 != v95 {
		v98 = v106
		v101 = v105
		v102 = v108
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
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if base.Ui32(v121-int32(9)) < base.Ui32(int32(-2)) {
		v118 = v133
		v121 = v121 - int32(8)
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v138 = v133
	goto L48
L60:
	;
	goto L59
L61:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)+36))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+76))
	if v156 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if v194 == int32(0) {
		v208 = v69
		goto L37
	} else {
		goto L75
	}
L63:
	;
	goto L62
L64:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v160 <= int32(0) {
		v194 = int32(0)
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v194 = int32(0)
	goto L63
L67:
	;
	v163 = int32(0)
	if v163 < v160 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v166 = v160
	goto L70
L69:
	;
	v166 = v163
	goto L70
L70:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v171 = int32(0)
	goto L71
L71:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v167+v171<<(uint(int32(2))%32))))
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179)+8)))
	if v180 == v81&int32(65535) {
		v194 = v179
		goto L63
	} else {
		goto L73
	}
L72:
	;
	goto L66
L73:
	;
	v183 = v171 + int32(1)
	if v183 != v166 {
		v171 = v183
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+26)))
	if v198 != 0 {
		v208 = v69
		goto L37
	} else {
		goto L76
	}
L76:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	if v200 != int32(36) {
		v208 = v69
		goto L37
	} else {
		goto L77
	}
L77:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v203 != 0 {
		v206 = v203
		goto L39
	} else {
		goto L78
	}
L78:
	;
	v214 = int32(0)
	goto L36
L79:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v217 = v215
	goto L81
L80:
	;
	v217 = int32(0)
	goto L81
L81:
	;
	if v214 != v217 {
		goto L34
	} else {
		goto L82
	}
L82:
	;
	F_errhint(m, int32(542077), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	goto L34
L84:
	;
	F_errfinish(m, int32(495167), int32(1096), int32(31750))
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
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v295)+8))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
	v300 = F_transformAssignedExpr(m, l0, v293, int32(15), v296, v297, v298, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L10
	} else {
		goto L103
	}
L95:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v278 <= v248 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v291 = int32(0)
	goto L97
L97:
	;
	return v291
L98:
	;
	v291 = v255
	goto L97
L99:
	;
	if v266 == int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	if v277 == int32(0) {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v287 = v284 + v248<<(uint(int32(2))%32)
	if v287 != 0 {
		goto L94
	} else {
		goto L102
	}
L102:
	;
	goto L98
L103:
	;
	if l5 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v352 = F_lappend(m, v255, v348)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L10
	} else {
		goto L121
	}
L105:
	;
	v348 = v300
	goto L104
L106:
	;
	goto L107
L107:
	;
	if v300 == int32(0) {
		v348 = v300
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v312 = v300
	goto L109
L109:
	;
	v322 = v312
	goto L112
L110:
	;
	v348 = int32(0)
	goto L104
L111:
	;
	if v338 != 0 {
		v312 = v338
		goto L109
	} else {
		goto L120
	}
L112:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	if v326 != int32(55) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v322)+36))
	if v335 == int32(0) {
		v348 = v312
		goto L104
	} else {
		goto L119
	}
L114:
	;
	switch v326 - int32(14) {
	case 0:
		goto L117
	default:
		v348 = v312
		goto L104
	case 12:
		goto L118
	}
L115:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	v322 = v334
	goto L112
L116:
	;
	goto L113
L117:
	;
	goto L116
L118:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v338 = v333
	goto L111
L119:
	;
	v338 = v335
	goto L111
L120:
	;
	goto L110
L121:
	;
	v248 = v248 + int32(1)
	v255 = v352
	goto L86
}
