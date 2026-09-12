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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v186 int32
	_ = v186
	v8 = F_get_extension_control_directories(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v186
L2:
	;
	return int32(0)
L3:
	;
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v12 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v13 <= v12 {
		v186 = v12
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v186 = int32(0)
	goto L1
L7:
	;
	v20 = int32(0)
	goto L8
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v20<<(uint(int32(2))%32))))
	v28 = F_AllocateDir(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L6
L10:
	;
	v168 = v20 + int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v168 < v169 {
		v20 = v168
		goto L8
	} else {
		goto L56
	}
L11:
	;
	if v28 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v33 == int32(44) {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L17
L15:
	;
	goto L14
L16:
	;
	F_FreeDir(m, v28)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L55
	}
L17:
	;
	v43 = F_ReadDir(m, v28, v27)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L19
	}
L18:
	;
	F_FreeDir(m, v28)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L54
	}
L19:
	;
	if v43 == int32(0) {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v48 = v43 + int32(19)
	v52 = F_strlen(m, v48)
	mBase = m.M
	v59 = v52 + int32(1)
	goto L23
L21:
	;
	if v71 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	goto L21
L23:
	;
	v61 = int32(0)
	if v59 == v61 {
		v71 = v61
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v71 = v66
	goto L22
L25:
	;
	v65 = v59 - int32(1)
	v66 = v48 + v65
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v67 != int32(46) {
		v59 = v65
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v74 = int32(301363)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _consts[383])))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v78 == int32(0) {
		v97 = v77
		v98 = v78
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v98-v97 != 0 {
		goto L17
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	if v77 != v78 {
		v97 = v77
		v98 = v78
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v82 = v71
	v83 = v74
	goto L32
L32:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v87 == int32(0) {
		v97 = v86
		v98 = v87
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v97 = v86
	v98 = v87
	goto L29
L34:
	;
	v90 = int32(1)
	if v86 == v87 {
		v82 = v82 + v90
		v83 = v83 + v90
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v100 = F_pstrdup(m, v48)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v105 = F_strlen(m, v100)
	mBase = m.M
	v112 = v105 + int32(1)
	goto L40
L38:
	;
	v125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v125)
	v128 = F_strstr(m, v100, int32(669827))
	mBase = m.M
	if v128 != 0 {
		goto L17
	} else {
		goto L44
	}
L39:
	;
	goto L38
L40:
	;
	v114 = int32(0)
	if v112 == v114 {
		v124 = v114
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v124 = v119
	goto L39
L42:
	;
	v118 = v112 - int32(1)
	v119 = v100 + v118
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v120 != int32(46) {
		v112 = v118
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v132 == int32(0) {
		v151 = v131
		v152 = v132
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v152-v151 != 0 {
		goto L17
	} else {
		goto L53
	}
L46:
	;
	goto L45
L47:
	;
	if v131 != v132 {
		v151 = v131
		v152 = v132
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v136 = v100
	v137 = l0
	goto L49
L49:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	if v141 == int32(0) {
		v151 = v140
		v152 = v141
		goto L46
	} else {
		goto L51
	}
L50:
	;
	v151 = v140
	v152 = v141
	goto L46
L51:
	;
	v144 = int32(1)
	if v140 == v141 {
		v136 = v136 + v144
		v137 = v137 + v144
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	goto L18
L54:
	;
	return int32(1)
L55:
	;
	goto L10
L56:
	;
	goto L9
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
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
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
	var v43 int32
	_ = v43
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
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(96)
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
	F_ScanKeyInit(m, v8, int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v8+int32(48), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = F_systable_beginscan(m, v12, int32(2673), int32(1), int32(0), int32(2), v8)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_systable_endscan(m, v32)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v34 = F_systable_getnext(m, v32)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v34 == int32(0) {
		v59 = v3
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v38 = v34
	goto L9
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v45 = v43 + v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if v46 != int32(3079) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v59 = v3
	goto L5
L11:
	;
	v53 = F_systable_getnext(m, v32)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+24)))
	if v49 != int32(101) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	v59 = v52
	goto L5
L14:
	;
	if v53 != 0 {
		v38 = v53
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	F_sequence_close(m, v12, int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	m.G0 = v8 + int32(96)
	return v59
}
func F_recordExtensionInitPrivWorker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	v10 = m.G0
	v12 = v10 - int32(208)
	m.G0 = v12
	v16 = F_aclmembers(m, l3, v12+int32(56))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v20 = F_table_open(m, int32(3394), int32(3))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			F_ScanKeyInit(m, v12-int32(-64), int32(1), int32(3), int32(184), l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_ScanKeyInit(m, v12+int32(112), int32(2), int32(3), int32(184), l1)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v38 = int32(3)
					F_ScanKeyInit(m, v12+int32(160), v38, v38, int32(65), l2)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v49 = F_systable_beginscan(m, v20, int32(3395), int32(1), int32(0), int32(3), v12-int32(-64))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v51 = F_systable_getnext(m, v49)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								if v51 != 0 {
									v53 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v53
									*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v53
									v57 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)) = uint8(v57)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v57
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v57
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
									v67 = F_heap_getattr_2(m, v51, int32(5), v64, v12+int32(15))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v69 = F_pg_detoast_datum(m, v67)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											v73 = F_aclmembers(m, v69, v12+int32(60))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
												v76 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
												F_updateInitAclDependencies(m, l1, l0, l2, v73, v75, v16, v76)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													if l3 == int32(0) {
														F_CatalogTupleDelete(m, v20, v51+int32(4))
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return
														} else {
															F_systable_endscan(m, v49)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return
															} else {
																F_CommandCounterIncrement(m)
																mBase = m.M
																v139 = m.ExcPending
																if v139 != 0 {
																	return
																} else {
																	F_sequence_close(m, v20, int32(3))
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return
																	} else {
																		m.G0 = v12 + int32(208)
																		return
																	}
																}
															}
														}
													} else {
														v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
														if v81 == int32(0) {
															F_CatalogTupleDelete(m, v20, v51+int32(4))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return
															} else {
																F_systable_endscan(m, v49)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v139 = m.ExcPending
																	if v139 != 0 {
																		return
																	} else {
																		F_sequence_close(m, v20, int32(3))
																		mBase = m.M
																		v142 = m.ExcPending
																		if v142 != 0 {
																			return
																		} else {
																			m.G0 = v12 + int32(208)
																			return
																		}
																	}
																}
															}
														} else {
															v84 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)) = uint8(v84)
															*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
															v94 = F_heap_modify_tuple(m, v51, v87, v12+int32(32), v12+int32(24), v12+int32(16))
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return
															} else {
																F_CatalogTupleUpdate(m, v20, v94+int32(4), v94)
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return
																} else {
																	F_systable_endscan(m, v49)
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return
																	} else {
																		F_CommandCounterIncrement(m)
																		mBase = m.M
																		v139 = m.ExcPending
																		if v139 != 0 {
																			return
																		} else {
																			F_sequence_close(m, v20, int32(3))
																			mBase = m.M
																			v142 = m.ExcPending
																			if v142 != 0 {
																				return
																			} else {
																				m.G0 = v12 + int32(208)
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
									v104 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)) = uint8(v104)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v104
									if l3 == v104 {
										F_systable_endscan(m, v49)
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v139 = m.ExcPending
											if v139 != 0 {
												return
											} else {
												F_sequence_close(m, v20, int32(3))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return
												} else {
													m.G0 = v12 + int32(208)
													return
												}
											}
										}
									} else {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
										if v110 == int32(0) {
											F_systable_endscan(m, v49)
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return
											} else {
												F_CommandCounterIncrement(m)
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
													return
												} else {
													F_sequence_close(m, v20, int32(3))
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return
													} else {
														m.G0 = v12 + int32(208)
														return
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
											*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(101)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
											v124 = F_heap_form_tuple(m, v119, v12+int32(32), v12+int32(24))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return
											} else {
												F_CatalogTupleInsert(m, v20, v124)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return
												} else {
													v128 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v128
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
													F_updateInitAclDependencies(m, l1, l0, l2, v128, v128, v16, v132)
													mBase = m.M
													v134 = m.ExcPending
													if v134 != 0 {
														return
													} else {
														F_systable_endscan(m, v49)
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v139 = m.ExcPending
															if v139 != 0 {
																return
															} else {
																F_sequence_close(m, v20, int32(3))
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return
																} else {
																	m.G0 = v12 + int32(208)
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
