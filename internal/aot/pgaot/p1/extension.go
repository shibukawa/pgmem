package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_extension_file_exists(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	v8 = F_get_extension_control_directories(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v8 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if int32(0) < v16 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = int32(0)
	goto L10
L7:
	;
	v198 = int32(0)
	goto L8
L8:
	;
	return v198
L9:
	;
	F_FreeDir(m, v32)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L60
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v23<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v32 = F_AllocateDir(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	return int32(0)
L12:
	;
	v182 = v23 + int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v182 < v183 {
		v23 = v182
		goto L10
	} else {
		goto L59
	}
L13:
	;
	if v32 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_extension_file_exists[0]))
	if v37 == int32(44) {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v41 = F_ReadDir(m, v32, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	if v41 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v44 = v41
	goto L22
L20:
	;
	goto L21
L21:
	;
	F_FreeDir(m, v32)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L58
	}
L22:
	;
	v51 = v44 + int32(19)
	v55 = F_strlen(m, v51)
	mBase = m.M
	v62 = v55 + int32(1)
	goto L27
L23:
	;
	goto L21
L24:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v163 = F_ReadDir(m, v32, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L56
	}
L25:
	;
	if v74 == int32(0) {
		goto L24
	} else {
		goto L31
	}
L26:
	;
	goto L25
L27:
	;
	v64 = int32(0)
	if v62 == v64 {
		v74 = v64
		goto L26
	} else {
		goto L29
	}
L28:
	;
	v74 = v69
	goto L26
L29:
	;
	v68 = v62 - int32(1)
	v69 = v51 + v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v70 != int32(46) {
		v62 = v68
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v77 = int32(_a_F_extension_file_exists_0)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_extension_file_exists[1])))
	if base.B2i32(v80 == int32(0))|base.B2i32(v80 != v83) != 0 {
		v101 = v80
		v102 = v83
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v101-v102 != 0 {
		goto L24
	} else {
		goto L39
	}
L33:
	;
	goto L32
L34:
	;
	v86 = v74
	v87 = v77
	goto L35
L35:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if v91 == int32(0) {
		v101 = v91
		v102 = v90
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v101 = v91
	v102 = v90
	goto L33
L37:
	;
	v94 = int32(1)
	if v91 == v90 {
		v86 = v86 + v94
		v87 = v87 + v94
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v104 = F_pstrdup(m, v51)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v109 = F_strlen(m, v104)
	mBase = m.M
	v116 = v109 + int32(1)
	goto L43
L41:
	;
	v129 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v129)
	v132 = F_strstr(m, v104, int32(_a_F_extension_file_exists_1))
	mBase = m.M
	if v132 != 0 {
		goto L24
	} else {
		goto L47
	}
L42:
	;
	goto L41
L43:
	;
	v118 = int32(0)
	if v116 == v118 {
		v128 = v118
		goto L42
	} else {
		goto L45
	}
L44:
	;
	v128 = v123
	goto L42
L45:
	;
	v122 = v116 - int32(1)
	v123 = v104 + v122
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v124 != int32(46) {
		v116 = v122
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v135 == int32(0))|base.B2i32(v135 != v138) != 0 {
		v156 = v135
		v157 = v138
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v156-v157 == int32(0) {
		goto L9
	} else {
		goto L55
	}
L49:
	;
	goto L48
L50:
	;
	v141 = v104
	v142 = l0
	goto L51
L51:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)))
	if v146 == int32(0) {
		v156 = v146
		v157 = v145
		goto L49
	} else {
		goto L53
	}
L52:
	;
	v156 = v146
	v157 = v145
	goto L49
L53:
	;
	v149 = int32(1)
	if v146 == v145 {
		v141 = v141 + v149
		v142 = v142 + v149
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L24
L56:
	;
	if v163 != 0 {
		v44 = v163
		goto L22
	} else {
		goto L57
	}
L57:
	;
	goto L23
L58:
	;
	goto L12
L59:
	;
	goto L11
L60:
	;
	v198 = int32(1)
	goto L8
}
func F_getExtensionOfObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
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
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v12 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v8, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v8+int32(56), int32(2), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v34 = F_systable_beginscan(m, v12, int32(2673), int32(1), int32(0), int32(2), v8)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_systable_endscan(m, v34)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v36 = F_systable_getnext(m, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v36 == int32(0) {
		v61 = v3
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v40 = v36
	goto L9
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+22)))
	v47 = v45 + v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if v48 != int32(3079) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v61 = v3
	goto L5
L11:
	;
	v55 = F_systable_getnext(m, v34)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+24)))
	if v51 != int32(101) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v61 = v54
	goto L5
L14:
	;
	if v55 != 0 {
		v40 = v55
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	F_relation_close(m, v12, int32(1))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	m.G0 = v8 + int32(112)
	return v61
}
func F_recordExtensionInitPrivWorker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	v13 = m.G0
	v15 = v13 - int32(256)
	m.G0 = v15
	v19 = F_aclmembers(m, l3, v15+int32(72))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v23 = F_table_open(m, int32(3394), int32(3))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v26 = v15 + int32(80)
			v30 = base.I64_extend_i32_u(l0)
			F_ScanKeyInit(m, v26, int32(1), int32(3), int32(184), v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v38 = base.I64_extend_i32_u(l1)
				F_ScanKeyInit(m, v15+int32(136), int32(2), int32(3), int32(184), v38)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v43 = int32(3)
					v46 = base.I64_extend_i32_s(l2)
					F_ScanKeyInit(m, v15+int32(192), v43, v43, int32(65), v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						v53 = F_systable_beginscan(m, v23, int32(3395), int32(1), int32(0), int32(3), v26)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v55 = F_systable_getnext(m, v53)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								if v55 != 0 {
									v57 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v57
									*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v57
									*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v57
									*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v57
									v65 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)) = uint8(v65)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v65
									*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v65
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
									v75 = F_heap_getattr_2(m, v55, int32(5), v72, v15+int32(15))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										v78 = F_pg_detoast_datum(m, base.I32_wrap_i64(v75))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											v82 = F_aclmembers(m, v78, v15+int32(76))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
												F_updateInitAclDependencies(m, l1, l0, l2, v82, v84, v19, v85)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													if l3 == int32(0) {
														F_simple_heap_delete(m, v23, v55+int32(4))
														mBase = m.M
														v113 = m.ExcPending
														if v113 != 0 {
															return
														} else {
															F_systable_endscan(m, v53)
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
																return
															} else {
																F_CommandCounterIncrement(m)
																mBase = m.M
																v150 = m.ExcPending
																if v150 != 0 {
																	return
																} else {
																	F_relation_close(m, v23, int32(3))
																	mBase = m.M
																	v153 = m.ExcPending
																	if v153 != 0 {
																		return
																	} else {
																		m.G0 = v15 + int32(256)
																		return
																	}
																}
															}
														}
													} else {
														v90 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
														if v90 == int32(0) {
															F_simple_heap_delete(m, v23, v55+int32(4))
															mBase = m.M
															v113 = m.ExcPending
															if v113 != 0 {
																return
															} else {
																F_systable_endscan(m, v53)
																mBase = m.M
																v148 = m.ExcPending
																if v148 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v150 = m.ExcPending
																	if v150 != 0 {
																		return
																	} else {
																		F_relation_close(m, v23, int32(3))
																		mBase = m.M
																		v153 = m.ExcPending
																		if v153 != 0 {
																			return
																		} else {
																			m.G0 = v15 + int32(256)
																			return
																		}
																	}
																}
															}
														} else {
															v93 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)) = uint8(v93)
															*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = base.I64_extend_i32_u(l3)
															v97 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
															v104 = F_heap_modify_tuple(m, v55, v97, v15+int32(32), v15+int32(24), v15+int32(16))
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																F_CatalogTupleUpdate(m, v23, v104+int32(4), v104)
																mBase = m.M
																v109 = m.ExcPending
																if v109 != 0 {
																	return
																} else {
																	F_systable_endscan(m, v53)
																	mBase = m.M
																	v148 = m.ExcPending
																	if v148 != 0 {
																		return
																	} else {
																		F_CommandCounterIncrement(m)
																		mBase = m.M
																		v150 = m.ExcPending
																		if v150 != 0 {
																			return
																		} else {
																			F_relation_close(m, v23, int32(3))
																			mBase = m.M
																			v153 = m.ExcPending
																			if v153 != 0 {
																				return
																			} else {
																				m.G0 = v15 + int32(256)
																				return
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
									v114 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)) = uint8(v114)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v114
									if l3 == v114 {
										F_systable_endscan(m, v53)
										mBase = m.M
										v148 = m.ExcPending
										if v148 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return
											} else {
												F_relation_close(m, v23, int32(3))
												mBase = m.M
												v153 = m.ExcPending
												if v153 != 0 {
													return
												} else {
													m.G0 = v15 + int32(256)
													return
												}
											}
										}
									} else {
										v120 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
										if v120 == int32(0) {
											F_systable_endscan(m, v53)
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return
											} else {
												F_CommandCounterIncrement(m)
												mBase = m.M
												v150 = m.ExcPending
												if v150 != 0 {
													return
												} else {
													F_relation_close(m, v23, int32(3))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return
													} else {
														m.G0 = v15 + int32(256)
														return
													}
												}
											}
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = int64(101)
											*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v46
											*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v38
											*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v30
											*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = base.I64_extend_i32_u(l3)
											v130 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
											v135 = F_heap_form_tuple(m, v130, v15+int32(32), v15+int32(24))
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return
											} else {
												F_CatalogTupleInsert(m, v23, v135)
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return
												} else {
													v139 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v139
													v143 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
													F_updateInitAclDependencies(m, l1, l0, l2, v139, v139, v19, v143)
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
														return
													} else {
														F_systable_endscan(m, v53)
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v150 = m.ExcPending
															if v150 != 0 {
																return
															} else {
																F_relation_close(m, v23, int32(3))
																mBase = m.M
																v153 = m.ExcPending
																if v153 != 0 {
																	return
																} else {
																	m.G0 = v15 + int32(256)
																	return
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
			}
		}
	}
}
