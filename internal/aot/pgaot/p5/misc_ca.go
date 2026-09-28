package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_calcstrlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v27 int32
	_ = v27
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v6 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v70 + v71
L2:
	;
	v9 = l0
	v10 = v5
	v12 = v2
	goto L5
L3:
	;
	v60 = v5
	v62 = v2
	goto L4
L4:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v70 = v63&int32(4095) + int32(1)
	v71 = v62
	goto L1
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v14 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v18 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v60 = v55
	v62 = v53
	goto L4
L7:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v50 == int32(1) {
		v70 = v49
		v71 = v12
		goto L1
	} else {
		goto L16
	}
L8:
	;
	v49 = v47 + v48
	goto L7
L9:
	;
	v21 = v13
	v22 = v17
	v24 = v14
	goto L12
L10:
	;
	v37 = v17
	v39 = v14
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v47 = v40&int32(4095) + int32(1)
	v48 = v39
	goto L8
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v26 = F_calcstrlen(m, v25)
	mBase = m.M
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v27 == int32(1) {
		v47 = v26
		v48 = v24
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v37 = v32
	v39 = v30
	goto L11
L14:
	;
	v30 = v26 + v24
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v33 != int32(1) {
		v21 = v31
		v22 = v32
		v24 = v30
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v53 = v49 + v12
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v56 != int32(1) {
		v9 = v54
		v10 = v55
		v12 = v53
		goto L5
	} else {
		goto L17
	}
L17:
	;
	goto L6
}
func F_calculate_indexes_size(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
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
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int64
	_ = v68
	var v74 int32
	_ = v74
	var v78 int64
	_ = v78
	v4 = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+116)))
	if v10 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v78 = v4
	goto L3
L3:
	;
	return v78
L4:
	;
	F_list_free(m, v13)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L18
	}
L5:
	;
	return int64(0)
L6:
	;
	if v13 == int32(0) {
		v68 = v4
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v19 <= int32(0) {
		v68 = v4
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v24 = int32(0)
	v25 = v4
	goto L9
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v24<<(uint(int32(2))%32))))
	v36 = F_relation_open(m, v34, int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	v68 = v60
	goto L4
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v40 = F_calculate_relation_size(m, v36, v38, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v44 = F_calculate_relation_size(m, v36, v42, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v48 = F_calculate_relation_size(m, v36, v46, int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v52 = F_calculate_relation_size(m, v36, v50, int32(3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	F_relation_close(m, v36, int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v60 = v52 + (v48 + (v44 + (v25 + v40)))
	v62 = v24 + int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v62 < v63 {
		v24 = v62
		v25 = v60
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L10
L18:
	;
	v78 = v68
	goto L3
}
func F_casefold(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
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
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		if v13 == int32(1) {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			if v19 == int32(18) {
				v22 = int32(16)
			} else {
				v22 = int32(0)
			}
			if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v29 = int32(4)
			} else {
				v29 = v22
			}
			v42 = v29
		} else {
			v30 = int32(1)
			if v13&v30 != 0 {
				v42 = int32(base.Ui32(v13)>>(uint(v30)%32)) - v30
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v42 = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v45 = m.G0
		v47 = v45 - int32(16)
		m.G0 = v47
		v49 = int32(1)
		if v13&v49 != 0 {
			v53 = v49
		} else {
			v53 = int32(4)
		}
		v54 = v9 + v53
		if v54 == int32(0) {
			v108 = int32(0)
			m.G0 = v47 + int32(16)
			v154 = F_cstring_to_text(m, v108)
			mBase = m.M
			v155 = m.ExcPending
			if v155 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v108)
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v154)
				}
			}
		} else {
			if v43 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(34209924))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(_a_F_casefold_0)
						F_errmsg(m, int32(_a_F_casefold_1), v47)
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return int64(0)
						} else {
							F_errhint(m, int32(_a_F_casefold_2), int32(0))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_casefold_3), int32(1831), int32(_a_F_casefold_4))
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, _c_F_casefold[0]))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
				if v61 != int32(6) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(16801924))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_casefold_5), int32(0))
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_casefold_3), int32(1837), int32(_a_F_casefold_4))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v64 = F_pg_newlocale_from_collation(m, v43)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int64(0)
					} else {
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+2)))
						if v66 == int32(1) {
							v69 = F_pnstrdup(m, v54, v42)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int64(0)
							} else {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
								if v71 == int32(0) {
									v108 = v69
								} else {
									v74 = v69
									v76 = v71
									for {
										if base.Ui32((v76-int32(65))&int32(255)) < base.Ui32(int32(26)) {
											v89 = v76 | int32(32)
										} else {
											v89 = v76
										}
										*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v89)
										v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
										if v91 != 0 {
											v74 = v74 + int32(1)
											v76 = v91
											continue
										} else {
											break
										}
										break
									}
									v108 = v69
								}
								m.G0 = v47 + int32(16)
								v154 = F_cstring_to_text(m, v108)
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
									return int64(0)
								} else {
									F_pfree(m, v108)
									mBase = m.M
									v157 = m.ExcPending
									if v157 != 0 {
										return int64(0)
									} else {
										return base.I64_extend_i32_u(v154)
									}
								}
							}
						} else {
							v95 = v42 + int32(1)
							v96 = F_palloc(m, v95)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int64(0)
							} else {
								v98 = F_pg_strfold(m, v96, v95, v54, v42, v64)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int64(0)
								} else {
									v101 = v98 + int32(1)
									if base.Ui32(v101) <= base.Ui32(v95) {
										v108 = v96
										m.G0 = v47 + int32(16)
										v154 = F_cstring_to_text(m, v108)
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return int64(0)
										} else {
											F_pfree(m, v108)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v154)
											}
										}
									} else {
										v103 = F_repalloc(m, v96, v101)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int64(0)
										} else {
											v105 = F_pg_strfold(m, v103, v101, v54, v42, v64)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int64(0)
											} else {
												v108 = v103
												m.G0 = v47 + int32(16)
												v154 = F_cstring_to_text(m, v108)
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return int64(0)
												} else {
													F_pfree(m, v108)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int64(0)
													} else {
														return base.I64_extend_i32_u(v154)
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
