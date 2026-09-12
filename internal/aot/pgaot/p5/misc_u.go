package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UnlinkLockFiles(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	if v5 == v3 {
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		if v8 <= int32(0) {
		} else {
			v11 = v3
			for {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+v11<<(uint(int32(2))%32))))
				v18 = F_unlink(m, v17)
				mBase = m.M
				v20 = v11 + int32(1)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				if v20 < v21 {
					v11 = v20
					continue
				} else {
					break
				}
				break
			}
		}
	}
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = int32(0)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[800])))
	if v31 != 0 {
		v32 = int32(15)
	} else {
		v32 = int32(18)
	}
	v34 = F_errstart(m, v32, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		if v34 != 0 {
			F_errmsg(m, int32(255996), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				F_errfinish(m, int32(515589), int32(1198), int32(175120))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			return
		}
	}
}
func F_UnlockBuffers(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[631]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(240560)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(518165)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v21 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v20 | v21
	if v20&v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v6 + int32(32)
	return
L4:
	;
	goto L7
L5:
	;
	v40 = v20
	goto L6
L6:
	;
	v45 = int32(4160188)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[602]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(8))+8))
	if v48 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	F_perform_spin_delay(m, v6+int32(8))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v40 = v33
	goto L6
L9:
	;
	return
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v34 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v33 | v34
	if v33&v34 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	if v40&int32(536870912) != 0 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[602])) = v63
	goto L13
L15:
	;
	if int32(999) < v46 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v46 < int32(11) {
		goto L13
	} else {
		goto L22
	}
L18:
	;
	v53 = int32(900)
	if v53 <= v46 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v56 = v53
	goto L21
L20:
	;
	v56 = v46
	goto L21
L21:
	;
	v63 = v56 + int32(100)
	goto L14
L22:
	;
	v63 = v46 - int32(1)
	goto L14
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v71 = *(*int32)(unsafe.Add(mBase, _consts[74]))
	if v69 == v71 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v74 = v40
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v74 & int32(-4194305)
	*(*int32)(unsafe.Add(mBase, _consts[631])) = int32(0)
	goto L3
L26:
	;
	v73 = v40 & int32(-536870913)
	goto L28
L27:
	;
	v73 = v40
	goto L28
L28:
	;
	v74 = v73
	goto L25
}
func F_UnregisterExprContextCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = l0 + int32(68)
	v11 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if l1 != v12 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v21 != 0 {
		v8 = v20
		v11 = v21
		goto L4
	} else {
		goto L15
	}
L7:
	;
	v20 = v11
	goto L6
L8:
	;
	goto L9
L9:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if l2 != v14 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v20 = v11
	goto L6
L11:
	;
	goto L12
L12:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
	F_pfree(m, v11)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	v20 = v8
	goto L6
L15:
	;
	goto L5
}
func F_UpdateFullPageWrites(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+160)))
	if v5 != v8 {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[113])))
		if v11 == int32(1) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+316))
			v17 = base.B2i32(v15 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _consts[113])) = uint8(v17)
			v19 = v17
		} else {
			v19 = int32(0)
		}
		v20 = int32(4548788)
		v22 = *(*int32)(unsafe.Add(mBase, _consts[14]))
		*(*int32)(unsafe.Add(mBase, _consts[14])) = v22 + int32(1)
		if v5 != 0 {
			F_WALInsertLockAcquireExclusive(m)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+160)) = uint8(v28)
				F_WALInsertLockRelease(m)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _consts[15]))
					v34 = int32(0)
					if v19|base.B2i32(v33 <= v34) == v34 {
						F_XLogBeginInsert(m)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_XLogRegisterData(m, int32(4159308), int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v47 = F_XLogInsert(m, int32(0), int32(128))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									v50 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
									if v50 == int32(0) {
										F_WALInsertLockAcquireExclusive(m)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											v55 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v7)+160)) = uint8(v55)
											F_WALInsertLockRelease(m)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return
											} else {
												v59 = int32(4548788)
												v61 = *(*int32)(unsafe.Add(mBase, _consts[14]))
												*(*int32)(unsafe.Add(mBase, _consts[14])) = v61 - int32(1)
												return
											}
										}
									} else {
										v59 = int32(4548788)
										v61 = *(*int32)(unsafe.Add(mBase, _consts[14]))
										*(*int32)(unsafe.Add(mBase, _consts[14])) = v61 - int32(1)
										return
									}
								}
							}
						}
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
						if v50 == int32(0) {
							F_WALInsertLockAcquireExclusive(m)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v55 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v7)+160)) = uint8(v55)
								F_WALInsertLockRelease(m)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									v59 = int32(4548788)
									v61 = *(*int32)(unsafe.Add(mBase, _consts[14]))
									*(*int32)(unsafe.Add(mBase, _consts[14])) = v61 - int32(1)
									return
								}
							}
						} else {
							v59 = int32(4548788)
							v61 = *(*int32)(unsafe.Add(mBase, _consts[14]))
							*(*int32)(unsafe.Add(mBase, _consts[14])) = v61 - int32(1)
							return
						}
					}
				}
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			v34 = int32(0)
			if v19|base.B2i32(v33 <= v34) == v34 {
				F_XLogBeginInsert(m)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_XLogRegisterData(m, int32(4159308), int32(1))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v47 = F_XLogInsert(m, int32(0), int32(128))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
							if v50 == int32(0) {
								F_WALInsertLockAcquireExclusive(m)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									v55 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v7)+160)) = uint8(v55)
									F_WALInsertLockRelease(m)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										v59 = int32(4548788)
										v61 = *(*int32)(unsafe.Add(mBase, _consts[14]))
										*(*int32)(unsafe.Add(mBase, _consts[14])) = v61 - int32(1)
										return
									}
								}
							} else {
								v59 = int32(4548788)
								v61 = *(*int32)(unsafe.Add(mBase, _consts[14]))
								*(*int32)(unsafe.Add(mBase, _consts[14])) = v61 - int32(1)
								return
							}
						}
					}
				}
			} else {
				v50 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
				if v50 == int32(0) {
					F_WALInsertLockAcquireExclusive(m)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						v55 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+160)) = uint8(v55)
						F_WALInsertLockRelease(m)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							v59 = int32(4548788)
							v61 = *(*int32)(unsafe.Add(mBase, _consts[14]))
							*(*int32)(unsafe.Add(mBase, _consts[14])) = v61 - int32(1)
							return
						}
					}
				} else {
					v59 = int32(4548788)
					v61 = *(*int32)(unsafe.Add(mBase, _consts[14]))
					*(*int32)(unsafe.Add(mBase, _consts[14])) = v61 - int32(1)
					return
				}
			}
		}
	} else {
		return
	}
}
func F_UtilityTupleDescriptor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v8 - int32(203) {
	case 0:
		goto L5
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v221 = v2
		goto L1
	case 10:
		goto L6
	default:
		goto L7
	}
L1:
	;
	return v221
L2:
	;
	if v8 != int32(159) {
		v221 = v2
		goto L1
	} else {
		goto L38
	}
L3:
	;
	v114 = F_ExplainResultDesc(m, l0)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L9
	} else {
		goto L37
	}
L4:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v103 = F_FetchPreparedStatement(m, v101, int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L31
	}
L5:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v91 != 0 {
		v221 = v2
		goto L1
	} else {
		goto L27
	}
L6:
	;
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = F_SearchSysCache1(m, int32(47), v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	switch v8 - int32(241) {
	case 0:
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v221 = v2
		goto L1
	case 12:
		goto L4
	default:
		goto L2
	}
L8:
	;
	return v24
L9:
	;
	return int32(0)
L10:
	;
	if v20 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v24 = F_build_function_result_tupdesc_t(m, v20)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L24
	}
L14:
	;
	F_ReleaseCatCache(m, v20)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	if v24 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	m.G0 = v15 + int32(16)
	goto L8
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v30 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v36 = int32(0)
	v41 = v30
	goto L19
L19:
	;
	v44 = v36 + int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v36<<(uint(int32(2))%32))))
	v58 = F_exprType(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L21
	}
L20:
	;
	goto L16
L21:
	;
	F_TupleDescInitEntry(m, v24, base.I32_extend16_s(v44), v24+int32(24)+v41<<(uint(int32(4))%32)+v36*int32(100), v58, int32(-1), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v44 < v64 {
		v36 = v44
		v41 = v64
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v80
	F_errmsg_internal(m, int32(54566), v15)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(517157), int32(2393), int32(510134))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v93 = F_GetPortalByName(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	if v93 == int32(0) {
		v221 = v2
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+92))
	v98 = F_CreateTupleDescCopy(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	return v98
L31:
	;
	if v103 == int32(0) {
		v221 = v2
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+64))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	if v108 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v109 = F_CreateTupleDescCopy(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L9
	} else {
		goto L36
	}
L34:
	;
	v112 = int32(0)
	goto L35
L35:
	;
	return v112
L36:
	;
	v112 = v109
	goto L35
L37:
	;
	return v114
L38:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v120 = m.G0
	v122 = v120 - int32(16)
	m.G0 = v122
	v127 = v119
	v128 = int32(319286)
	goto L41
L39:
	;
	m.G0 = v122 + int32(16)
	v221 = v215
	goto L1
L40:
	;
	if v173 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v132 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v173 = base.I32_extend8_s(v153) - base.I32_extend8_s(v162)
	goto L40
L43:
	;
	v141 = int32(1)
	if base.Ui32((v132-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L51
	} else {
		goto L52
	}
L44:
	;
	if v131&int32(255) != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if v131&int32(255) != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v173 = int32(1)
	goto L40
L48:
	;
	v140 = int32(-1)
	goto L50
L49:
	;
	v140 = int32(0)
	goto L50
L50:
	;
	v173 = v140
	goto L40
L51:
	;
	v153 = v132 | int32(32)
	goto L53
L52:
	;
	v153 = v132
	goto L53
L53:
	;
	if base.Ui32((v131-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v162 = v131 | int32(32)
	goto L56
L55:
	;
	v162 = v131
	goto L56
L56:
	;
	if v153 == v162&int32(255) {
		v127 = v127 + v141
		v128 = v128 + v141
		goto L41
	} else {
		goto L57
	}
L57:
	;
	goto L42
L58:
	;
	v177 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L9
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v203 = F_GetConfigOptionByName(m, v119, v122+int32(12), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L65
	}
L61:
	;
	F_TupleDescInitEntry(m, v177, int32(1), int32(399600), int32(25), int32(-1), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	F_TupleDescInitEntry(m, v177, int32(2), int32(344768), int32(25), int32(-1), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	F_TupleDescInitEntry(m, v177, int32(3), int32(258980), int32(25), int32(-1), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v215 = v177
	goto L39
L65:
	;
	v206 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	F_TupleDescInitEntry(m, v206, int32(1), v209, int32(25), int32(-1), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	v215 = v206
	goto L39
}
func F___udivmodti4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v74 int64
	_ = v74
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v163 int64
	_ = v163
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v199 int64
	_ = v199
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v208 int64
	_ = v208
	var v221 int32
	_ = v221
	var v234 int64
	_ = v234
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v247 int64
	_ = v247
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var __phi250 int64
	_ = __phi250
	var v251 int64
	_ = v251
	var __phi251 int64
	_ = __phi251
	var v252 int64
	_ = v252
	var __phi252 int64
	_ = __phi252
	var v253 int64
	_ = v253
	var __phi253 int64
	_ = __phi253
	var v255 int64
	_ = v255
	var __phi255 int64
	_ = __phi255
	var v261 int32
	_ = v261
	var __phi261 int32
	_ = __phi261
	var v263 int64
	_ = v263
	var v271 int64
	_ = v271
	var v272 int64
	_ = v272
	var v273 int64
	_ = v273
	var v276 int64
	_ = v276
	var v282 int64
	_ = v282
	var v289 int64
	_ = v289
	var v301 int64
	_ = v301
	var v315 int64
	_ = v315
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	v6 = int64(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if l2 == l4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v330
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v332
	m.G0 = v17 + int32(16)
	return
L2:
	;
	v22 = base.B2i32(base.Ui64(l3) <= base.Ui64(l1))
	goto L4
L3:
	;
	v22 = base.B2i32(base.Ui64(l4) <= base.Ui64(l2))
	goto L4
L4:
	;
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l4 == int64(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v315 = v6
	goto L7
L7:
	;
	v330 = v315
	v332 = int64(0)
	goto L1
L8:
	;
	if base.Ui64(l2) < base.Ui64(l3) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	v221 = base.I32_wrap_i64(base.I64_clz(l4)) - base.I32_wrap_i64(base.I64_clz(l2))
	if int32(0) <= v221 {
		goto L47
	} else {
		goto L48
	}
L11:
	;
	v330 = v203 + v204<<(uint(int64(32))%64)
	v332 = v208
	goto L1
L12:
	;
	v26 = base.I64_clz(l3)
	v33 = l2<<(uint(v26)%64) | int64(base.Ui64(int64(base.Ui64(l1)>>(uint(int64(1))%64)))>>(uint(v26^int64(-1))%64))
	v34 = l3 << (uint(v26) % 64)
	v35 = int64(32)
	v36 = int64(base.Ui64(v34) >> (uint(v35) % 64))
	v37 = base.I64_div_u_s(v33, v36)
	v40 = l1 << (uint(v26) % 64)
	v41 = int64(4294967295)
	v44 = int64(base.Ui64(v40) >> (uint(v35) % 64))
	v46 = v34 & v41
	v50 = v33 - v37*v36
	v51 = v37
	goto L15
L13:
	;
	goto L14
L14:
	;
	v112 = base.I64_div_u_s(l2, l3)
	v115 = base.I64_clz(l3)
	v122 = (l2-v112*l3)<<(uint(v115)%64) | int64(base.Ui64(int64(base.Ui64(l1)>>(uint(int64(1))%64)))>>(uint(v115^int64(-1))%64))
	v123 = l3 << (uint(v115) % 64)
	v124 = int64(32)
	v125 = int64(base.Ui64(v123) >> (uint(v124) % 64))
	v126 = base.I64_div_u_s(v122, v125)
	v129 = l1 << (uint(v115) % 64)
	v130 = int64(4294967295)
	v133 = int64(base.Ui64(v129) >> (uint(v124) % 64))
	v135 = v123 & v130
	v139 = v122 - v126*v125
	v140 = v126
	goto L31
L15:
	;
	if base.Ui64(v51) <= base.Ui64(int64(4294967295)) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v79 = v44 | v33<<(uint(int64(32))%64) - v74*v34
	v80 = base.I64_div_u_s(v79, v36)
	v84 = v79 - v80*v36
	v86 = v80
	goto L23
L17:
	;
	goto L16
L18:
	;
	if base.Ui64(v51*v46) <= base.Ui64(v50<<(uint(int64(32))%64)|v44) {
		v74 = v51
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v69 = v51 - int64(1)
	v70 = v36 + v50
	if base.Ui64(v70) < base.Ui64(int64(4294967296)) {
		v50 = v70
		v51 = v69
		goto L15
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v74 = v69
	goto L17
L23:
	;
	if base.Ui64(v86) <= base.Ui64(int64(4294967295)) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v203 = v110
	v204 = v74
	v208 = int64(0)
	goto L11
L25:
	;
	goto L24
L26:
	;
	if base.Ui64(v86*v46) <= base.Ui64(v84<<(uint(int64(32))%64)|v40&v41) {
		v110 = v86
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v105 = v86 - int64(1)
	v106 = v84 + v36
	if base.Ui64(v106) < base.Ui64(int64(4294967296)) {
		v84 = v106
		v86 = v105
		goto L23
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v110 = v105
	goto L25
L31:
	;
	if base.Ui64(v140) <= base.Ui64(int64(4294967295)) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v168 = v122<<(uint(int64(32))%64) | v133 - v163*v123
	v169 = base.I64_div_u_s(v168, v125)
	v173 = v168 - v169*v125
	v175 = v169
	goto L39
L33:
	;
	goto L32
L34:
	;
	if base.Ui64(v140*v135) <= base.Ui64(v139<<(uint(int64(32))%64)|v133) {
		v163 = v140
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v158 = v140 - int64(1)
	v159 = v125 + v139
	if base.Ui64(v159) < base.Ui64(int64(4294967296)) {
		v139 = v159
		v140 = v158
		goto L31
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v163 = v158
	goto L33
L39:
	;
	if base.Ui64(v175) <= base.Ui64(int64(4294967295)) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v203 = v199
	v204 = v163
	v208 = v112
	goto L11
L41:
	;
	goto L40
L42:
	;
	if base.Ui64(v175*v135) <= base.Ui64(v173<<(uint(int64(32))%64)|v129&v130) {
		v199 = v175
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v194 = v175 - int64(1)
	v195 = v173 + v125
	if base.Ui64(v195) < base.Ui64(int64(4294967296)) {
		v173 = v195
		v175 = v194
		goto L39
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	v199 = v194
	goto L41
L47:
	;
	if v221&int32(64) != 0 {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	v301 = v6
	goto L49
L49:
	;
	v315 = v301
	goto L7
L50:
	;
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	__phi250 = l1
	__phi251 = l2
	__phi252 = v248
	__phi253 = v247
	__phi255 = v6
	__phi261 = v221
	v250 = __phi250
	v251 = __phi251
	v252 = __phi252
	v253 = __phi253
	v255 = __phi255
	v261 = __phi261
	goto L56
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v242
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v243
	goto L50
L52:
	;
	v242 = int64(0)
	v243 = l3 << (uint(base.I64_extend_i32_u(v221+int32(-64))) % 64)
	goto L51
L53:
	;
	goto L54
L54:
	;
	if v221 == int32(0) {
		v242 = l3
		v243 = l4
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v234 = base.I64_extend_i32_u(v221)
	v242 = l3 << (uint(v234) % 64)
	v243 = l4<<(uint(v234)%64) | int64(base.Ui64(l3)>>(uint(base.I64_extend_i32_u(int32(64)-v221))%64))
	goto L51
L56:
	;
	v263 = int64(-1)
	v271 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v252+(v250^v263)) < base.Ui64(v252))) + (v253 + (v251 ^ v263))
	v272 = int64(63)
	v273 = v271 >> (uint(v272) % 64)
	v276 = v252 & v273
	v282 = int64(1)
	v289 = v255<<(uint(v282)%64) | int64(base.Ui64(v271)>>(uint(v272)%64))
	if v261 != 0 {
		__phi250 = v250 - v276
		__phi251 = v251 - v273&v253 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v250) < base.Ui64(v276)))
		__phi252 = v253<<(uint(v272)%64) | int64(base.Ui64(v252)>>(uint(v282)%64))
		__phi253 = int64(base.Ui64(v253) >> (uint(v282) % 64))
		__phi255 = v289
		__phi261 = v261 - int32(1)
		v250 = __phi250
		v251 = __phi251
		v252 = __phi252
		v253 = __phi253
		v255 = __phi255
		v261 = __phi261
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v301 = v289
	goto L49
L58:
	;
	goto L57
}
func F___uselocale(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if l0 != 0 {
		if l0 == int32(-1) {
			v9 = int32(4718872)
		} else {
			v9 = l0
		}
		*(*int32)(unsafe.Add(mBase, _consts[1])) = v9
	} else {
	}
	if v4 == int32(4718872) {
		v14 = int32(-1)
	} else {
		v14 = v4
	}
	return v14
}
func F_unlink(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	v4 = m.Env.X__syscall_unlinkat(m, int32(-100), l0, int32(0))
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v4) {
		*(*int32)(unsafe.Add(mBase, _consts[87])) = int32(0) - v4
		v12 = int32(-1)
	} else {
		v12 = v4
	}
	return v12
}
func F_update_grouptailpos(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+382)))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = int32(4554128)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+96))
	if v22 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(16)
	return
L4:
	;
	v124 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+382)) = uint8(v124)
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v16
	goto L3
L5:
	;
	F_spool_tuples(m, l0, int64(-1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	F_tuplestore_select_read_pointer(m, v30, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return
L9:
	;
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+352)) = v28
	goto L4
L10:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
	v36 = v34 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+352)) = v36
	F_spool_tuples(m, l0, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v41 = int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
	v44 = F_tuplestore_gettupleslot(m, v40, v41, v41, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	m.T0[v114].(func(*base.Module, int32))(m, v112)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L8
	} else {
		goto L30
	}
L13:
	;
	if v44 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	goto L15
L15:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	if v55 <= v56 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L12
L17:
	;
	v93 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
	v95 = v93 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+352)) = v95
	F_spool_tuples(m, l0, v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L27
	}
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+96))
	if v59 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v62
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v67 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	F_MemoryContextReset(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v73 = int32(4554128)
	v74 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v76
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v81 = m.T0[v80].(func(*base.Module, int32, int32, int32) int32)(m, v67, v63, v10+int32(15))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L24
	}
L23:
	;
	goto L17
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v74
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	F_MemoryContextReset(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	if v81 == int32(0) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	goto L17
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v100 = int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
	v103 = F_tuplestore_gettupleslot(m, v99, v100, v100, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	if v103 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	goto L16
L30:
	;
	goto L4
}
