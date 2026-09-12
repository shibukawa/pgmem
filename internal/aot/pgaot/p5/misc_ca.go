package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CancelDBBackends(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	v3 = l2
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CancelDBBackends[0]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_CancelDBBackends[1]))
	v17 = F_LWLockAcquire(m, v13+int32(512), v4)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if int32(0) < v19 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_CancelDBBackends[2]))
	v31 = v4
	v32 = v25
	goto L6
L4:
	;
	goto L5
L5:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_CancelDBBackends[1]))
	F_LWLockRelease(m, v69+int32(512))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(36)+v31<<(uint(int32(2))%32))))
	v41 = v32 + v38*int32(640)
	if l0 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v56 = v31 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v56 < v57 {
		v31 = v56
		v32 = v53
		goto L6
	} else {
		goto L15
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+60))
	if v42 != l0 {
		v53 = v32
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+73)) = uint8(v3)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+44))
	if v45 == int32(0) {
		v53 = v32
		goto L8
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	v49 = F_SendProcSignal(m, v45, l1, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_CancelDBBackends[2]))
	v53 = v52
	goto L8
L15:
	;
	goto L7
L16:
	;
	return
}
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
	var v11 int32
	_ = v11
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
	var v23 int32
	_ = v23
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
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
	return v69 + v71
L2:
	;
	v9 = l0
	v10 = v2
	v11 = v5
	goto L5
L3:
	;
	v60 = v2
	v61 = v5
	goto L4
L4:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v69 = v60
	v71 = v63&int32(4095) + int32(1)
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
	v60 = v53
	v61 = v55
	goto L4
L7:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v50 == int32(1) {
		v69 = v10
		v71 = v49
		goto L1
	} else {
		goto L16
	}
L8:
	;
	v49 = v46 + v48
	goto L7
L9:
	;
	v21 = v13
	v22 = v14
	v23 = v17
	goto L12
L10:
	;
	v37 = v14
	v38 = v17
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v46 = v37
	v48 = v40&int32(4095) + int32(1)
	goto L8
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v26 = F_calcstrlen(m, v25)
	mBase = m.M
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v27 == int32(1) {
		v46 = v22
		v48 = v26
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v37 = v30
	v38 = v32
	goto L11
L14:
	;
	v30 = v22 + v26
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v33 != int32(1) {
		v21 = v31
		v22 = v30
		v23 = v32
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v53 = v10 + v49
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v56 != int32(1) {
		v9 = v54
		v10 = v53
		v11 = v55
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
	var v2 int64
	_ = v2
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
	var v23 int64
	_ = v23
	var v29 int32
	_ = v29
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
	var v66 int64
	_ = v66
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	v2 = int64(0)
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
	v76 = v2
	goto L3
L3:
	;
	return v76
L4:
	;
	F_list_free(m, v13)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L20
	}
L5:
	;
	return int64(0)
L6:
	;
	if v13 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v66 = v2
	goto L4
L8:
	;
	goto L9
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v19 <= int32(0) {
		v66 = v2
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v23 = v2
	v29 = int32(0)
	goto L11
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v29<<(uint(int32(2))%32))))
	v36 = F_relation_open(m, v34, int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v66 = v60
	goto L4
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v40 = F_calculate_relation_size(m, v36, v38, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v44 = F_calculate_relation_size(m, v36, v42, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v48 = F_calculate_relation_size(m, v36, v46, int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v52 = F_calculate_relation_size(m, v36, v50, int32(3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_relation_close(m, v36, int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v60 = v52 + (v48 + (v44 + (v23 + v40)))
	v62 = v29 + int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v62 < v63 {
		v23 = v60
		v29 = v62
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L12
L20:
	;
	v76 = v66
	goto L3
}
func F_casefold(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
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
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(1)
		v14 = v9 + v13
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		v19 = v17 & v13
		if v19 != 0 {
			v20 = v14
		} else {
			v20 = v9 + int32(4)
		}
		if v17 == int32(1) {
			v23 = int32(4)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v25&int32(254) == int32(2) {
				v34 = v23
			} else {
				v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
			}
			if v25 == int32(1) {
				v37 = v23
			} else {
				v37 = v34
			}
			v48 = v37
		} else {
			v38 = int32(1)
			if v19 != 0 {
				v48 = int32(base.Ui32(v17)>>(uint(v38)%32)) - v38
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v50 = int32(0)
		v51 = m.G0
		v53 = v51 - int32(16)
		m.G0 = v53
		if v20 == v50 {
			v109 = v50
			m.G0 = v53 + int32(16)
			v156 = F_cstring_to_text(m, v109)
			mBase = m.M
			v157 = m.ExcPending
			if v157 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v109)
				mBase = m.M
				v159 = m.ExcPending
				if v159 != 0 {
					return int32(0)
				} else {
					return v156
				}
			}
		} else {
			if v49 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(34209924))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v53))) = int32(_a_F_casefold_0)
						F_errmsg(m, int32(_a_F_casefold_1), v53)
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_casefold_2), int32(0))
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_casefold_3), int32(1847), int32(_a_F_casefold_4))
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
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
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, _c_F_casefold[0]))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
				if v61 != int32(6) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16801924))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_casefold_5), int32(0))
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_casefold_3), int32(1853), int32(_a_F_casefold_4))
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
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
					v64 = F_pg_newlocale_from_collation(m, v49)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+3)))
						if v66 == int32(1) {
							v69 = F_pnstrdup(m, v20, v48)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
								if v71 == int32(0) {
									v109 = v69
								} else {
									v75 = v71
									v77 = v69
									for {
										v81 = int32(255)
										v82 = v75 & v81
										if base.Ui32((v82-int32(65))&v81) < base.Ui32(int32(26)) {
											v91 = v82 | int32(32)
										} else {
											v91 = v82
										}
										*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v91)
										v94 = v77 + int32(1)
										v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
										if v95 != 0 {
											v75 = v95
											v77 = v94
											continue
										} else {
											break
										}
										break
									}
									v109 = v69
								}
								m.G0 = v53 + int32(16)
								v156 = F_cstring_to_text(m, v109)
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v109)
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return int32(0)
									} else {
										return v156
									}
								}
							}
						} else {
							v97 = v48 + int32(1)
							v98 = F_palloc(m, v97)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v100 = F_pg_strfold(m, v98, v97, v20, v48, v64)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									v103 = v100 + int32(1)
									if base.Ui32(v103) <= base.Ui32(v97) {
										v109 = v98
										m.G0 = v53 + int32(16)
										v156 = F_cstring_to_text(m, v109)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v109)
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int32(0)
											} else {
												return v156
											}
										}
									} else {
										v105 = F_repalloc(m, v98, v103)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											v107 = F_pg_strfold(m, v105, v103, v20, v48, v64)
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return int32(0)
											} else {
												v109 = v105
												m.G0 = v53 + int32(16)
												v156 = F_cstring_to_text(m, v109)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v109)
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														return v156
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
