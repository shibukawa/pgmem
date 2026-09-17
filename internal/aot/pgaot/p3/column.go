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
	F_appendStringInfoString(m, v13, int32(_a_F_get_column_alias_list_0))
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
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
			v23 = F_convert_any_priv_string(m, v16, int32(_a_F_has_column_privilege_name_id_attnum_0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v25)
				if v13&int32(_a_F_has_column_privilege_name_id_attnum_1) == v25 {
					v44 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
					v49 = int32(0)
					m.G0 = v10 + int32(16)
					return v49
				} else {
					v33 = v10 + int32(15)
					v34 = F_pg_attribute_aclcheck_ext(m, v12, base.I32_extend16_s(v13), v20, v23, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 != 0 {
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v36 != 0 {
								v44 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
								v49 = int32(0)
								m.G0 = v10 + int32(16)
								return v49
							} else {
								v37 = F_pg_class_aclcheck_ext(m, v12, v20, v23, v33)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									if v37 != 0 {
										v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										if v40 == int32(0) {
										} else {
											v44 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
										}
										v49 = int32(0)
									} else {
										v49 = int32(1)
									}
									m.G0 = v10 + int32(16)
									return v49
								}
							}
						} else {
							v49 = int32(1)
							m.G0 = v10 + int32(16)
							return v49
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
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = F_get_role_oid_or_public(m, v13)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_convert_column_name(m, v12, v15)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v27 = F_convert_any_priv_string(m, v20, int32(_a_F_has_column_privilege_name_id_name_0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v29)
						if v24 == v29 {
							v45 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
							v50 = int32(0)
							m.G0 = v10 + int32(16)
							return v50
						} else {
							v34 = v10 + int32(15)
							v35 = F_pg_attribute_aclcheck_ext(m, v12, v24, v22, v27, v34)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v35 != 0 {
									v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
									if v37 != 0 {
										v45 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
										v50 = int32(0)
										m.G0 = v10 + int32(16)
										return v50
									} else {
										v38 = F_pg_class_aclcheck_ext(m, v12, v22, v27, v34)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											if v38 != 0 {
												v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
												if v41 == int32(0) {
												} else {
													v45 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
												}
												v50 = int32(0)
											} else {
												v50 = int32(1)
											}
											m.G0 = v10 + int32(16)
											return v50
										}
									}
								} else {
									v50 = int32(1)
									m.G0 = v10 + int32(16)
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
	var v17 int32
	_ = v17
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v22 = F_pg_detoast_datum_packed(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_get_role_oid_or_public(m, v12)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = F_textToQualifiedNameList(m, v14)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = F_makeRangeVarFromNameList(m, v26)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = int32(0)
							v34 = F_RangeVarGetRelidExtended(m, v28, v30, v30, v30, v30)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v36 = F_convert_column_name(m, v34, v19)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									v39 = F_convert_any_priv_string(m, v22, int32(_a_F_has_column_privilege_name_name_name_0))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										v41 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v41)
										if v36 == v41 {
											v57 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
											v62 = int32(0)
											m.G0 = v10 + int32(16)
											return v62
										} else {
											v46 = v10 + int32(15)
											v47 = F_pg_attribute_aclcheck_ext(m, v34, v36, v24, v39, v46)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												if v47 != 0 {
													v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
													if v49 != 0 {
														v57 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
														v62 = int32(0)
														m.G0 = v10 + int32(16)
														return v62
													} else {
														v50 = F_pg_class_aclcheck_ext(m, v34, v24, v39, v46)
														mBase = m.M
														v51 = m.ExcPending
														if v51 != 0 {
															return int32(0)
														} else {
															if v50 != 0 {
																v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
																if v53 == int32(0) {
																} else {
																	v57 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
																}
																v62 = int32(0)
															} else {
																v62 = int32(1)
															}
															m.G0 = v10 + int32(16)
															return v62
														}
													}
												} else {
													v62 = int32(1)
													m.G0 = v10 + int32(16)
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v251 int32
	_ = v251
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
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_scanner_yyerror(m, int32(_a_F_makeColumnRef_0), l3)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L53
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+4)) = v223
	m.G0 = v16 + int32(16)
	return v222
L5:
	;
	v214 = F_makeString(m, l0)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L51
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v28 <= int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v31 = int32(0)
	if v31 < v28 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = v28
	goto L10
L9:
	;
	v35 = v31
	goto L10
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v41 = v31
	goto L11
L11:
	;
	v53 = v41 << (uint(int32(2)) % 32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v38+v53)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	switch v56 - int32(77) {
	case 0:
		goto L14
	case 1:
		goto L15
	default:
		goto L13
	}
L12:
	;
	goto L5
L13:
	;
	v199 = v41 + int32(1)
	if v199 != v35 {
		v41 = v199
		goto L11
	} else {
		goto L50
	}
L14:
	;
	if v53+int32(4) < v28<<(uint(int32(2))%32) {
		goto L3
	} else {
		goto L49
	}
L15:
	;
	v60 = F_palloc0(m, int32(12))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(79)
	if v41 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v66 = F_makeString(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v126 = F_list_copy_tail(m, l1, v41)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L31
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v66
	v73 = F_list_make1_impl(m, int32(1), v16+int32(8))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v73
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v76 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = l1
	v222 = v60
	v223 = v19
	goto L4
L23:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v85 = int32(0)
	goto L24
L24:
	;
	v97 = v85 << (uint(int32(2)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v81+v97)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if base.B2i32(v100 == int32(77))&base.B2i32(v97+int32(4) < v76<<(uint(int32(2))%32)) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L3
L26:
	;
	v110 = v85 + int32(1)
	if v110 != v76 {
		v85 = v110
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L22
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v126
	v179 = int32(0)
	if base.B2i32(l1 == v179)|base.B2i32(v41 <= v179) != 0 {
		goto L41
	} else {
		goto L42
	}
L31:
	;
	if v126 == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v130 <= int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v147 = int32(0)
	goto L34
L34:
	;
	v150 = v147 << (uint(int32(2)) % 32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v135+v150)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	if base.B2i32(v153 == int32(77))&base.B2i32(v150+int32(4) < v130<<(uint(int32(2))%32)) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L3
L36:
	;
	v163 = v147 + int32(1)
	if v163 != v130 {
		v147 = v163
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L35
L39:
	;
	goto L30
L40:
	;
	v190 = F_makeString(m, l0)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L47
	}
L41:
	;
	v189 = int32(0)
	goto L43
L42:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v41 < v186 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L40
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v41
	goto L46
L45:
	;
	goto L46
L46:
	;
	v189 = l1
	goto L43
L47:
	;
	v192 = F_lcons(m, v190, v189)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v192
	v222 = v60
	v223 = v19
	goto L4
L49:
	;
	goto L13
L50:
	;
	goto L12
L51:
	;
	v216 = F_lcons(m, v214, l1)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v222 = v19
	v223 = v216
	goto L4
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
