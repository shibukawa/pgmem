package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_column_alias_list(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v5 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_appendStringInfoChar(m, v13, int32(40))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v17 = F_quote_identifier(m, v12)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_appendStringInfoString(m, v13, v17)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(2) <= v21 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v26 = int32(1)
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_appendStringInfoChar(m, v13, int32(41))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L17
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v26<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v13, int32(746027))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v37 = F_quote_identifier(m, v33)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_appendStringInfoString(m, v13, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v42 = v26 + int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v42 < v43 {
		v26 = v42
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	goto L1
}
func F_has_column_privilege_name_id_attnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = F_get_role_oid_or_public(m, v14)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v23 = F_convert_any_priv_string(m, v16, int32(1656768))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v25)
				if v13&int32(65535) == v25 {
					v45 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
					v48 = int32(0)
					m.G0 = v10 + int32(16)
					return v48
				} else {
					v34 = F_pg_attribute_aclcheck_ext(m, v12, base.I32_extend16_s(v13), v20, v23, v10+int32(15))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 != 0 {
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v36 != 0 {
								v45 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
								v48 = int32(0)
								m.G0 = v10 + int32(16)
								return v48
							} else {
								v39 = F_pg_class_aclcheck_ext(m, v12, v20, v23, v10+int32(15))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									if v39 != 0 {
										v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										if v42 == int32(0) {
										} else {
											v45 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
										}
										v48 = int32(0)
									} else {
										v48 = int32(1)
									}
									m.G0 = v10 + int32(16)
									return v48
								}
							}
						} else {
							v48 = int32(1)
							m.G0 = v10 + int32(16)
							return v48
						}
					}
				}
			}
		}
	}
}
func F_has_column_privilege_name_id_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v21 = F_pg_detoast_datum_packed(m, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = F_get_role_oid_or_public(m, v14)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = F_convert_column_name(m, v13, v16)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v28 = F_convert_any_priv_string(m, v21, int32(1656768))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v30)
						if v25 == v30 {
							v47 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
							v50 = int32(0)
							m.G0 = v11 + int32(16)
							return v50
						} else {
							v36 = F_pg_attribute_aclcheck_ext(m, v13, v25, v23, v28, v11+int32(15))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								if v36 != 0 {
									v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
									if v38 != 0 {
										v47 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
										v50 = int32(0)
										m.G0 = v11 + int32(16)
										return v50
									} else {
										v41 = F_pg_class_aclcheck_ext(m, v13, v23, v28, v11+int32(15))
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return int32(0)
										} else {
											if v41 != 0 {
												v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
												if v44 == int32(0) {
												} else {
													v47 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
												}
												v50 = int32(0)
											} else {
												v50 = int32(1)
											}
											m.G0 = v11 + int32(16)
											return v50
										}
									}
								} else {
									v50 = int32(1)
									m.G0 = v11 + int32(16)
									return v50
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_has_column_privilege_name_name_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v23 = F_pg_detoast_datum_packed(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = F_get_role_oid_or_public(m, v13)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = F_textToQualifiedNameList(m, v15)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = F_makeRangeVarFromNameList(m, v27)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							v31 = int32(0)
							v35 = F_RangeVarGetRelidExtended(m, v29, v31, v31, v31, v31)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v37 = F_convert_column_name(m, v35, v20)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									v40 = F_convert_any_priv_string(m, v23, int32(1656768))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										v42 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v42)
										if v37 == v42 {
											v59 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
											v62 = int32(0)
											m.G0 = v11 + int32(16)
											return v62
										} else {
											v48 = F_pg_attribute_aclcheck_ext(m, v35, v37, v25, v40, v11+int32(15))
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												if v48 != 0 {
													v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
													if v50 != 0 {
														v59 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
														v62 = int32(0)
														m.G0 = v11 + int32(16)
														return v62
													} else {
														v53 = F_pg_class_aclcheck_ext(m, v35, v25, v40, v11+int32(15))
														mBase = m.M
														v54 = m.ExcPending
														if v54 != 0 {
															return int32(0)
														} else {
															if v53 != 0 {
																v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
																if v56 == int32(0) {
																} else {
																	v59 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
																}
																v62 = int32(0)
															} else {
																v62 = int32(1)
															}
															m.G0 = v11 + int32(16)
															return v62
														}
													}
												} else {
													v62 = int32(1)
													m.G0 = v11 + int32(16)
													return v62
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
func F_makeColumnRef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v19 = F_palloc0(m, int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(69)
	if l1 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	F_scanner_yyerror(m, int32(730412), l3)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L61
	}
L4:
	;
	F_scanner_yyerror(m, int32(730412), l3)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L60
	}
L5:
	;
	F_scanner_yyerror(m, int32(730412), l3)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L59
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+4)) = v230
	m.G0 = v16 + int32(16)
	return v229
L7:
	;
	v219 = F_makeString(m, l0)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L57
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v28 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v31 = int32(0)
	if v31 < v28 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v35 = v28
	goto L12
L11:
	;
	v35 = v31
	goto L12
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v42 = v31
	goto L13
L13:
	;
	v55 = v36 + v42<<(uint(int32(2))%32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	switch v57 - int32(77) {
	case 0:
		goto L16
	case 1:
		goto L17
	default:
		goto L15
	}
L14:
	;
	goto L7
L15:
	;
	v204 = v42 + int32(1)
	if v204 != v35 {
		v42 = v204
		goto L13
	} else {
		goto L56
	}
L16:
	;
	v198 = v55 + int32(4)
	if v198 == int32(0) {
		goto L15
	} else {
		goto L54
	}
L17:
	;
	v61 = F_palloc0(m, int32(12))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(79)
	if v42 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v67 = F_makeString(m, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v128 = F_list_copy_tail(m, l1, v42)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L35
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v67
	v74 = F_list_make1_impl(m, int32(1), v16+int32(8))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v77 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v87 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = l1
	v229 = v61
	v230 = v19
	goto L6
L27:
	;
	v100 = v80 + v87<<(uint(int32(2))%32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v102 != int32(77) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L26
L29:
	;
	v112 = v87 + int32(1)
	if v77 != v112 {
		v87 = v112
		goto L27
	} else {
		goto L33
	}
L30:
	;
	v106 = v100 + int32(4)
	if v106 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if base.Ui32(v106) < base.Ui32(v80+v77<<(uint(int32(2))%32)) {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	goto L28
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v128
	v183 = int32(0)
	if l1 == v183 {
		v191 = v183
		goto L46
	} else {
		goto L47
	}
L35:
	;
	if v128 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v132 <= int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v144 = int32(0)
	goto L38
L38:
	;
	v155 = v135 + v144<<(uint(int32(2))%32)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v157 != int32(77) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L34
L40:
	;
	v167 = v144 + int32(1)
	if v132 != v167 {
		v144 = v167
		goto L38
	} else {
		goto L44
	}
L41:
	;
	v161 = v155 + int32(4)
	if v161 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if base.Ui32(v161) < base.Ui32(v135+v132<<(uint(int32(2))%32)) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	goto L39
L45:
	;
	v192 = F_makeString(m, l0)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L52
	}
L46:
	;
	goto L45
L47:
	;
	if v42 <= int32(0) {
		v191 = v183
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v42 < v188 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v42
	goto L51
L50:
	;
	goto L51
L51:
	;
	v191 = l1
	goto L46
L52:
	;
	v194 = F_lcons(m, v192, v191)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v194
	v229 = v61
	v230 = v19
	goto L6
L54:
	;
	if base.Ui32(v198) < base.Ui32(v36+v28<<(uint(int32(2))%32)) {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	goto L15
L56:
	;
	goto L14
L57:
	;
	v221 = F_lcons(m, v219, l1)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v229 = v19
	v230 = v221
	goto L6
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
