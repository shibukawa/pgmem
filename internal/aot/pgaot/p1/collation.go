package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_collation_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(16), l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v17 = v15 + v16
			v18 = F_CollationIsVisible(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 != 0 {
					F_initStringInfo(m, v8+int32(32))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v47 = F_quote_identifier(m, v17+int32(4))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_appendStringInfoString(m, v8+int32(32), v47)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
								F_ReleaseCatCache(m, v11)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v51
								}
							}
						}
					}
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
					v25 = F_get_namespace_name_or_temp(m, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v28 = v8 + int32(32)
						F_initStringInfo(m, v28)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							if v25 == int32(0) {
								v47 = F_quote_identifier(m, v17+int32(4))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									F_appendStringInfoString(m, v8+int32(32), v47)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
										F_ReleaseCatCache(m, v11)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											m.G0 = v8 + int32(48)
											return v51
										}
									}
								}
							} else {
								v33 = F_quote_identifier(m, v25)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v33
									F_appendStringInfo(m, v28, int32(_a_F_generate_collation_name_0), v8+int32(16))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										v47 = F_quote_identifier(m, v17+int32(4))
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return int32(0)
										} else {
											F_appendStringInfoString(m, v8+int32(32), v47)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return int32(0)
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
												F_ReleaseCatCache(m, v11)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(48)
													return v51
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
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_generate_collation_name_1), v8)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_generate_collation_name_2), int32(_a_F_generate_collation_name_3), int32(_a_F_generate_collation_name_4))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
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
func F_get_collation_actual_version_libc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	v5 = int32(_a_F_get_collation_actual_version_libc_0)
	v6 = l0
	goto L3
L1:
	;
	return int32(0)
L2:
	;
	if v43 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v9 == v10 {
		v32 = v9
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v43 = int32(0)
	goto L2
L5:
	;
	v34 = int32(1)
	if v32 != 0 {
		v5 = v5 + v34
		v6 = v6 + v34
		goto L3
	} else {
		goto L14
	}
L6:
	;
	if base.Ui32((v9-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v20 = v9 | int32(32)
	goto L9
L8:
	;
	v20 = v9
	goto L9
L9:
	;
	if base.Ui32((v10-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v29 = v10 | int32(32)
	goto L12
L11:
	;
	v29 = v10
	goto L12
L12:
	;
	if v20 == v29 {
		v32 = v20
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v43 = v20 - v29
	goto L2
L14:
	;
	goto L4
L15:
	;
	v50 = int32(_a_F_get_collation_actual_version_libc_1)
	v51 = l0
	v52 = int32(2)
	goto L17
L16:
	;
	if v97 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L17:
	;
	if v52 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v97 = int32(0)
	goto L16
L19:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 == v56 {
		v78 = v55
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	v80 = int32(1)
	if v78 != 0 {
		v50 = v50 + v80
		v51 = v51 + v80
		v52 = v52 - v80
		goto L17
	} else {
		goto L31
	}
L23:
	;
	if base.Ui32((v55-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v66 = v55 | int32(32)
	goto L26
L25:
	;
	v66 = v55
	goto L26
L26:
	;
	if base.Ui32((v56-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v75 = v56 | int32(32)
	goto L29
L28:
	;
	v75 = v56
	goto L29
L29:
	;
	if v66 == v75 {
		v78 = v66
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v97 = v66 - v75
	goto L16
L31:
	;
	goto L21
L32:
	;
	v103 = int32(_a_F_get_collation_actual_version_libc_2)
	v104 = l0
	goto L34
L33:
	;
	goto L1
L34:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v107 == v108 {
		v130 = v107
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v132 = int32(1)
	if v130 != 0 {
		v103 = v103 + v132
		v104 = v104 + v132
		goto L34
	} else {
		goto L45
	}
L37:
	;
	if base.Ui32((v107-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v118 = v107 | int32(32)
	goto L40
L39:
	;
	v118 = v107
	goto L40
L40:
	;
	if base.Ui32((v108-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v127 = v108 | int32(32)
	goto L43
L42:
	;
	v127 = v108
	goto L43
L43:
	;
	if v118 == v127 {
		v130 = v118
		goto L36
	} else {
		goto L44
	}
L44:
	;
	goto L33
L45:
	;
	goto L35
}
