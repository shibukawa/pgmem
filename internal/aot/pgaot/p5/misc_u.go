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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_UnlinkLockFiles[0]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_UnlinkLockFiles[0])) = int32(0)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_UnlinkLockFiles[1])))
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
			F_errmsg(m, int32(_a_F_UnlinkLockFiles_0), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_UnlinkLockFiles_1), int32(1215), int32(_a_F_UnlinkLockFiles_2))
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
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockBuffers[0]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a_F_UnlockBuffers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(_a_F_UnlockBuffers_1)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(_a_F_UnlockBuffers_2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v21 = int32(_a_F_UnlockBuffers_3)
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
	v45 = int32(_a_F_UnlockBuffers_4)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockBuffers[1]))
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
	v34 = int32(_a_F_UnlockBuffers_3)
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
	*(*int32)(unsafe.Add(mBase, _c_F_UnlockBuffers[1])) = v63
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
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockBuffers[2]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_UnlockBuffers[0])) = int32(0)
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
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[0])))
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[1]))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+160)))
	if v5 != v8 {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[2])))
		if v11 == int32(1) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+316))
			v17 = base.B2i32(v15 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[2])) = uint8(v17)
			v19 = v17
		} else {
			v19 = int32(0)
		}
		v20 = int32(_a_F_UpdateFullPageWrites_0)
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3]))
		*(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3])) = v22 + int32(1)
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
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[4]))
					v34 = int32(0)
					if v19|base.B2i32(v33 <= v34) == v34 {
						F_XLogBeginInsert(m)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_XLogRegisterData(m, int32(_a_F_UpdateFullPageWrites_1), int32(1))
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
									v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[0])))
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
												v59 = int32(_a_F_UpdateFullPageWrites_0)
												v61 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3]))
												*(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3])) = v61 - int32(1)
												return
											}
										}
									} else {
										v59 = int32(_a_F_UpdateFullPageWrites_0)
										v61 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3]))
										*(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3])) = v61 - int32(1)
										return
									}
								}
							}
						}
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[0])))
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
									v59 = int32(_a_F_UpdateFullPageWrites_0)
									v61 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3]))
									*(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3])) = v61 - int32(1)
									return
								}
							}
						} else {
							v59 = int32(_a_F_UpdateFullPageWrites_0)
							v61 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3]))
							*(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3])) = v61 - int32(1)
							return
						}
					}
				}
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[4]))
			v34 = int32(0)
			if v19|base.B2i32(v33 <= v34) == v34 {
				F_XLogBeginInsert(m)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_XLogRegisterData(m, int32(_a_F_UpdateFullPageWrites_1), int32(1))
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
							v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[0])))
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
										v59 = int32(_a_F_UpdateFullPageWrites_0)
										v61 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3]))
										*(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3])) = v61 - int32(1)
										return
									}
								}
							} else {
								v59 = int32(_a_F_UpdateFullPageWrites_0)
								v61 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3]))
								*(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3])) = v61 - int32(1)
								return
							}
						}
					}
				}
			} else {
				v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[0])))
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
							v59 = int32(_a_F_UpdateFullPageWrites_0)
							v61 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3]))
							*(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3])) = v61 - int32(1)
							return
						}
					}
				} else {
					v59 = int32(_a_F_UpdateFullPageWrites_0)
					v61 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3]))
					*(*int32)(unsafe.Add(mBase, _c_F_UpdateFullPageWrites[3])) = v61 - int32(1)
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
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v7 - int32(203) {
	case 0:
		goto L5
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v213 = v2
		goto L1
	case 10:
		goto L6
	default:
		goto L7
	}
L1:
	;
	return v213
L2:
	;
	if v7 != int32(159) {
		v213 = v2
		goto L1
	} else {
		goto L38
	}
L3:
	;
	v111 = F_ExplainResultDesc(m, l0)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L37
	}
L4:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v100 = F_FetchPreparedStatement(m, v98, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L31
	}
L5:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v88 != 0 {
		v213 = v2
		goto L1
	} else {
		goto L27
	}
L6:
	;
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = F_SearchSysCache1(m, int32(47), v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	switch v7 - int32(241) {
	case 0:
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v213 = v2
		goto L1
	case 12:
		goto L4
	default:
		goto L2
	}
L8:
	;
	return v23
L9:
	;
	return int32(0)
L10:
	;
	if v19 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v23 = F_build_function_result_tupdesc_t(m, v19)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
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
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L24
	}
L14:
	;
	F_ReleaseCatCache(m, v19)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	if v23 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	m.G0 = v14 + int32(16)
	goto L8
L17:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v29 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v33 = int32(0)
	v38 = v29
	goto L19
L19:
	;
	v40 = v33 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v33<<(uint(int32(2))%32))))
	v56 = F_exprType(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L21
	}
L20:
	;
	goto L16
L21:
	;
	F_TupleDescInitEntry(m, v23, base.I32_extend16_s(v40), v23+v38<<(uint(int32(4))%32)+v33*int32(100)+int32(24), v56, int32(-1), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v40 < v62 {
		v33 = v40
		v38 = v62
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v77
	F_errmsg_internal(m, int32(_a_F_UtilityTupleDescriptor_0), v14)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_UtilityTupleDescriptor_1), int32(2393), int32(_a_F_UtilityTupleDescriptor_2))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
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
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v90 = F_GetPortalByName(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	if v90 == int32(0) {
		v213 = v2
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+92))
	v95 = F_CreateTupleDescCopy(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	return v95
L31:
	;
	if v100 == int32(0) {
		v213 = v2
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+64))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+52))
	if v105 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v106 = F_CreateTupleDescCopy(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L36
	}
L34:
	;
	v109 = int32(0)
	goto L35
L35:
	;
	return v109
L36:
	;
	v109 = v106
	goto L35
L37:
	;
	return v111
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v117 = m.G0
	v119 = v117 - int32(16)
	m.G0 = v119
	v124 = v116
	v125 = int32(_a_F_UtilityTupleDescriptor_3)
	goto L41
L39:
	;
	m.G0 = v119 + int32(16)
	v213 = v208
	goto L1
L40:
	;
	if v166 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v129 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v166 = base.I32_extend8_s(v146) - base.I32_extend8_s(v155)
	goto L40
L43:
	;
	v134 = int32(1)
	if base.Ui32((v129-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L51
	} else {
		goto L52
	}
L44:
	;
	if v128 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if v128 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v166 = int32(1)
	goto L40
L48:
	;
	v133 = int32(-1)
	goto L50
L49:
	;
	v133 = int32(0)
	goto L50
L50:
	;
	v166 = v133
	goto L40
L51:
	;
	v146 = v129 | int32(32)
	goto L53
L52:
	;
	v146 = v129
	goto L53
L53:
	;
	if base.Ui32((v128-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v155 = v128 | int32(32)
	goto L56
L55:
	;
	v155 = v128
	goto L56
L56:
	;
	if v146 == v155&int32(255) {
		v124 = v124 + v134
		v125 = v125 + v134
		goto L41
	} else {
		goto L57
	}
L57:
	;
	goto L42
L58:
	;
	v170 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L9
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v196 = F_GetConfigOptionByName(m, v116, v119+int32(12), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L65
	}
L61:
	;
	F_TupleDescInitEntry(m, v170, int32(1), int32(_a_F_UtilityTupleDescriptor_4), int32(25), int32(-1), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	F_TupleDescInitEntry(m, v170, int32(2), int32(_a_F_UtilityTupleDescriptor_5), int32(25), int32(-1), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	F_TupleDescInitEntry(m, v170, int32(3), int32(_a_F_UtilityTupleDescriptor_6), int32(25), int32(-1), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v208 = v170
	goto L39
L65:
	;
	v199 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	F_TupleDescInitEntry(m, v199, int32(1), v202, int32(25), int32(-1), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	v208 = v199
	goto L39
}
func F___udivmodti4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v75 int64
	_ = v75
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v118 int64
	_ = v118
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v168 int64
	_ = v168
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v206 int64
	_ = v206
	var v210 int64
	_ = v210
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v227 int32
	_ = v227
	var v240 int64
	_ = v240
	var v248 int64
	_ = v248
	var v249 int64
	_ = v249
	var v253 int64
	_ = v253
	var v254 int64
	_ = v254
	var v256 int64
	_ = v256
	var __phi256 int64
	_ = __phi256
	var v257 int64
	_ = v257
	var __phi257 int64
	_ = __phi257
	var v258 int64
	_ = v258
	var __phi258 int64
	_ = __phi258
	var v259 int64
	_ = v259
	var __phi259 int64
	_ = __phi259
	var v260 int64
	_ = v260
	var __phi260 int64
	_ = __phi260
	var v266 int32
	_ = v266
	var __phi266 int32
	_ = __phi266
	var v268 int64
	_ = v268
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v281 int64
	_ = v281
	var v287 int64
	_ = v287
	var v294 int64
	_ = v294
	var v305 int64
	_ = v305
	var v318 int64
	_ = v318
	var v332 int64
	_ = v332
	var v333 int64
	_ = v333
	v6 = int64(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l2 == l4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v332
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v333
	m.G0 = v16 + int32(16)
	return
L2:
	;
	v21 = base.B2i32(base.Ui64(l3) <= base.Ui64(l1))
	goto L4
L3:
	;
	v21 = base.B2i32(base.Ui64(l4) <= base.Ui64(l2))
	goto L4
L4:
	;
	if v21 != 0 {
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
	v318 = v6
	goto L7
L7:
	;
	v332 = v318
	v333 = int64(0)
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
	v227 = base.I32_wrap_i64(base.I64_clz(l4)) - base.I32_wrap_i64(base.I64_clz(l2))
	if int32(0) <= v227 {
		goto L39
	} else {
		goto L40
	}
L11:
	;
	v332 = v210 + v211<<(uint(int64(32))%64)
	v333 = v213
	goto L1
L12:
	;
	v25 = base.I64_clz(l3)
	v32 = l2<<(uint(v25)%64) | int64(base.Ui64(int64(base.Ui64(l1)>>(uint(int64(1))%64)))>>(uint(v25^int64(-1))%64))
	v33 = l3 << (uint(v25) % 64)
	v34 = int64(32)
	v35 = int64(base.Ui64(v33) >> (uint(v34) % 64))
	v36 = base.I64_div_u_s(v32, v35)
	v39 = l1 << (uint(v25) % 64)
	v40 = int64(4294967295)
	v43 = int64(base.Ui64(v39) >> (uint(v34) % 64))
	v45 = v33 & v40
	v49 = v32 - v36*v35
	v50 = v36
	goto L15
L13:
	;
	goto L14
L14:
	;
	v115 = base.I64_div_u_s(l2, l3)
	v118 = base.I64_clz(l3)
	v125 = (l2-v115*l3)<<(uint(v118)%64) | int64(base.Ui64(int64(base.Ui64(l1)>>(uint(int64(1))%64)))>>(uint(v118^int64(-1))%64))
	v126 = l3 << (uint(v118) % 64)
	v127 = int64(32)
	v128 = int64(base.Ui64(v126) >> (uint(v127) % 64))
	v129 = base.I64_div_u_s(v125, v128)
	v132 = l1 << (uint(v118) % 64)
	v133 = int64(4294967295)
	v136 = int64(base.Ui64(v132) >> (uint(v127) % 64))
	v138 = v126 & v133
	v142 = v125 - v129*v128
	v143 = v129
	goto L27
L15:
	;
	if base.B2i32(base.Ui64(v50) <= base.Ui64(int64(4294967295)))&base.B2i32(base.Ui64(v50*v45) <= base.Ui64(v49<<(uint(int64(32))%64)|v43)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v80 = v43 | v32<<(uint(int64(32))%64) - v75*v33
	v81 = base.I64_div_u_s(v80, v35)
	v85 = v80 - v81*v35
	v87 = v81
	goto L21
L17:
	;
	v70 = v50 - int64(1)
	v71 = v35 + v49
	if base.Ui64(v71) < base.Ui64(int64(4294967296)) {
		v49 = v71
		v50 = v70
		goto L15
	} else {
		goto L20
	}
L18:
	;
	v75 = v50
	goto L19
L19:
	;
	goto L16
L20:
	;
	v75 = v70
	goto L19
L21:
	;
	if base.B2i32(base.Ui64(v87) <= base.Ui64(int64(4294967295)))&base.B2i32(base.Ui64(v87*v45) <= base.Ui64(v85<<(uint(int64(32))%64)|v39&v40)) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v210 = v113
	v211 = v75
	v213 = int64(0)
	goto L11
L23:
	;
	v108 = v87 - int64(1)
	v109 = v85 + v35
	if base.Ui64(v109) < base.Ui64(int64(4294967296)) {
		v85 = v109
		v87 = v108
		goto L21
	} else {
		goto L26
	}
L24:
	;
	v113 = v87
	goto L25
L25:
	;
	goto L22
L26:
	;
	v113 = v108
	goto L25
L27:
	;
	if base.B2i32(base.Ui64(v143) <= base.Ui64(int64(4294967295)))&base.B2i32(base.Ui64(v143*v138) <= base.Ui64(v142<<(uint(int64(32))%64)|v136)) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v173 = v125<<(uint(int64(32))%64) | v136 - v168*v126
	v174 = base.I64_div_u_s(v173, v128)
	v178 = v173 - v174*v128
	v180 = v174
	goto L33
L29:
	;
	v163 = v143 - int64(1)
	v164 = v128 + v142
	if base.Ui64(v164) < base.Ui64(int64(4294967296)) {
		v142 = v164
		v143 = v163
		goto L27
	} else {
		goto L32
	}
L30:
	;
	v168 = v143
	goto L31
L31:
	;
	goto L28
L32:
	;
	v168 = v163
	goto L31
L33:
	;
	if base.B2i32(base.Ui64(v180) <= base.Ui64(int64(4294967295)))&base.B2i32(base.Ui64(v180*v138) <= base.Ui64(v178<<(uint(int64(32))%64)|v132&v133)) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v210 = v206
	v211 = v168
	v213 = v115
	goto L11
L35:
	;
	v201 = v180 - int64(1)
	v202 = v178 + v128
	if base.Ui64(v202) < base.Ui64(int64(4294967296)) {
		v178 = v202
		v180 = v201
		goto L33
	} else {
		goto L38
	}
L36:
	;
	v206 = v180
	goto L37
L37:
	;
	goto L34
L38:
	;
	v206 = v201
	goto L37
L39:
	;
	if v227&int32(64) != 0 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v305 = v6
	goto L41
L41:
	;
	v318 = v305
	goto L7
L42:
	;
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	__phi256 = l1
	__phi257 = l2
	__phi258 = v254
	__phi259 = v253
	__phi260 = v6
	__phi266 = v227
	v256 = __phi256
	v257 = __phi257
	v258 = __phi258
	v259 = __phi259
	v260 = __phi260
	v266 = __phi266
	goto L48
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v248
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v249
	goto L42
L44:
	;
	v248 = int64(0)
	v249 = l3 << (uint(base.I64_extend_i32_u(v227+int32(-64))) % 64)
	goto L43
L45:
	;
	goto L46
L46:
	;
	if v227 == int32(0) {
		v248 = l3
		v249 = l4
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v240 = base.I64_extend_i32_u(v227)
	v248 = l3 << (uint(v240) % 64)
	v249 = l4<<(uint(v240)%64) | int64(base.Ui64(l3)>>(uint(base.I64_extend_i32_u(int32(64)-v227))%64))
	goto L43
L48:
	;
	v268 = int64(-1)
	v276 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v259+(v256^v268)) < base.Ui64(v259))) + (v258 + (v257 ^ v268))
	v277 = int64(63)
	v278 = v276 >> (uint(v277) % 64)
	v281 = v278 & v259
	v287 = int64(1)
	v294 = v260<<(uint(v287)%64) | int64(base.Ui64(v276)>>(uint(v277)%64))
	if v266 != 0 {
		__phi256 = v256 - v281
		__phi257 = v257 - v278&v258 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v256) < base.Ui64(v281)))
		__phi258 = int64(base.Ui64(v258) >> (uint(v287) % 64))
		__phi259 = v258<<(uint(v277)%64) | int64(base.Ui64(v259)>>(uint(v287)%64))
		__phi260 = v294
		__phi266 = v266 - int32(1)
		v256 = __phi256
		v257 = __phi257
		v258 = __phi258
		v259 = __phi259
		v260 = __phi260
		v266 = __phi266
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v305 = v294
	goto L41
L50:
	;
	goto L49
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F___uselocale[0]))
	if l0 != 0 {
		if l0 == int32(-1) {
			v9 = int32(_a_F___uselocale_0)
		} else {
			v9 = l0
		}
		*(*int32)(unsafe.Add(mBase, _c_F___uselocale[0])) = v9
	} else {
	}
	if v4 == int32(_a_F___uselocale_0) {
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
		*(*int32)(unsafe.Add(mBase, _c_F_unlink[0])) = int32(0) - v4
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
	v15 = int32(_a_F_update_grouptailpos_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_update_grouptailpos[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_update_grouptailpos[0])) = v20
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
	*(*int32)(unsafe.Add(mBase, _c_F_update_grouptailpos[0])) = v16
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
	v73 = int32(_a_F_update_grouptailpos_0)
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_update_grouptailpos[0]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_update_grouptailpos[0])) = v76
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
	*(*int32)(unsafe.Add(mBase, _c_F_update_grouptailpos[0])) = v74
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
