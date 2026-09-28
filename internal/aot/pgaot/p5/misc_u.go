package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UnlinkLockFiles(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_UnlinkLockFiles[0]))
	if v6 == v3 {
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
		if v9 <= int32(0) {
		} else {
			v12 = v3
			for {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v15+v12<<(uint(int32(2))%32))))
				v20 = F_unlink(m, v19)
				mBase = m.M
				v22 = v12 + int32(1)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
				if v22 < v23 {
					v12 = v22
					continue
				} else {
					break
				}
				break
			}
		}
	}
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
				F_errfinish(m, int32(_a_F_UnlinkLockFiles_1), int32(1147), int32(_a_F_UnlinkLockFiles_2))
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v19 int64
	_ = v19
	var v32 int64
	_ = v32
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v90 int64
	_ = v90
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v114 int64
	_ = v114
	var v118 int64
	_ = v118
	var v125 int64
	_ = v125
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockBuffers[0]))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = int64(4194304)
	v14 = base.AtomicRmwOr64(m, v11, int32(24), v12)
	if v14&v12 != int64(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v8 + int32(32)
	return
L4:
	;
	v19 = v14
	goto L7
L5:
	;
	v90 = v14
	goto L6
L6:
	;
	if v90&int64(536870912) != int64(0) {
		goto L29
	} else {
		goto L30
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(_a_F_UnlockBuffers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(_a_F_UnlockBuffers_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_UnlockBuffers_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v32
	if v19&int64(4194304) != v32 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v90 = v85
	goto L6
L9:
	;
	goto L12
L10:
	;
	goto L11
L11:
	;
	v63 = int32(_a_F_UnlockBuffers_3)
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockBuffers[1]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v66 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	return
L15:
	;
	v47 = int64(0)
	v50 = base.AtomicRmwCmpxchg64(m, v11, int32(24), v47, v47)
	if v50&int64(4194304) != v47 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v83 = int64(4194304)
	v85 = base.AtomicRmwOr64(m, v11, int32(24), v83)
	if v85&v83 != int64(0) {
		v19 = v85
		goto L7
	} else {
		goto L28
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_UnlockBuffers[1])) = v81
	goto L18
L20:
	;
	if int32(999) < v64 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v64 < int32(11) {
		goto L18
	} else {
		goto L27
	}
L23:
	;
	v71 = int32(900)
	if v71 <= v64 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v74 = v71
	goto L26
L25:
	;
	v74 = v64
	goto L26
L26:
	;
	v81 = v74 + int32(100)
	goto L19
L27:
	;
	v81 = v64 - int32(1)
	goto L19
L28:
	;
	goto L8
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockBuffers[2]))
	if v102 == v104 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v107 = int64(-1)
	goto L31
L31:
	;
	v109 = v90 | int64(4194304)
	v114 = base.AtomicRmwCmpxchg64(m, v11, int32(24), v109, v90&v107&int64(-4194305))
	if v114 != v109 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v106 = int64(-536870913)
	goto L34
L33:
	;
	v106 = int64(-1)
	goto L34
L34:
	;
	v107 = v106
	goto L31
L35:
	;
	v118 = v114
	goto L38
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_UnlockBuffers[0])) = int32(0)
	goto L3
L38:
	;
	v125 = base.AtomicRmwCmpxchg64(m, v11, int32(24), v118, v118&(v107&int64(-4194305)))
	if v118 != v125 {
		v118 = v125
		goto L38
	} else {
		goto L40
	}
L39:
	;
	goto L37
L40:
	;
	goto L39
}
func F_UnregisterExprContextCallback(m *base.Module, l0 int32, l1 int32, l2 int64) {
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
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = l0 + int32(80)
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
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
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
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+308))
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
	var v18 int64
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
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v7 - int32(241) {
	case 0:
		goto L4
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		v487 = v2
		goto L1
	case 12:
		goto L5
	case 26:
		goto L2
	default:
		goto L6
	}
L1:
	;
	return v487
L2:
	;
	v392 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L10
	} else {
		goto L110
	}
L3:
	;
	if v7 != int32(159) {
		v487 = v2
		goto L1
	} else {
		goto L61
	}
L4:
	;
	v202 = F_ExplainResultDesc(m, l0)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L10
	} else {
		goto L60
	}
L5:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v191 = F_FetchPreparedStatement(m, v189, int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L10
	} else {
		goto L54
	}
L6:
	;
	switch v7 - int32(203) {
	case 0:
		goto L7
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v487 = v2
		goto L1
	case 10:
		goto L8
	default:
		goto L3
	}
L7:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v179 != 0 {
		v487 = v2
		goto L1
	} else {
		goto L50
	}
L8:
	;
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+4)))
	v19 = F_SearchSysCache1(m, int32(47), v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return v23
L10:
	;
	return int32(0)
L11:
	;
	if v19 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v23 = F_build_function_result_tupdesc_t(m, v19)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L10
	} else {
		goto L47
	}
L15:
	;
	F_ReleaseCatCache(m, v19)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if v23 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if int32(0) < v27 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	m.G0 = v14 + int32(16)
	goto L9
L20:
	;
	v32 = int32(0)
	v36 = v27
	goto L23
L21:
	;
	goto L22
L22:
	;
	v68 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v68 < v77 {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	v38 = v32 + int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v32<<(uint(int32(2))%32))))
	v54 = F_exprType(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	F_TupleDescInitEntry(m, v23, base.I32_extend16_s(v38), v23+v36<<(uint(int32(3))%32)+v32*int32(100)+int32(32), v54, int32(-1), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v38 < v60 {
		v32 = v38
		v36 = v60
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	goto L19
L29:
	;
	v81 = v23 + int32(28)
	v88 = v68
	v89 = v77
	v91 = v68
	goto L33
L30:
	;
	v145 = v68
	v152 = v77
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v145
	goto L28
L32:
	;
	v145 = v139
	v152 = v118
	goto L31
L33:
	;
	v97 = v81 + v77<<(uint(int32(3))%32) + v88*int32(100)
	v100 = v81 + v88<<(uint(int32(3))%32)
	if v77 != v89 {
		v118 = v89
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v139 = v77
	goto L32
L35:
	;
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+2)))
	if v119 <= int32(0) {
		v139 = v88
		goto L32
	} else {
		goto L43
	}
L36:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+7)))
	if v102 != int32(118) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v118 = v88
	goto L35
L38:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
	if v105 != int32(1) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+6)))
	if v108&int32(6) != 0 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+2)))
	if v111 <= int32(0) {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+90)))
	if v114 != int32(118) {
		v118 = v77
		goto L35
	} else {
		goto L42
	}
L42:
	;
	goto L37
L43:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+90)))
	if v122 == int32(118) {
		v139 = v88
		goto L32
	} else {
		goto L44
	}
L44:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+5)))
	v131 = (v91 + v125 - int32(1)) & (int32(0) - v125)
	if int32(_a_F_UtilityTupleDescriptor_0) < v131 {
		v139 = v88
		goto L32
	} else {
		goto L45
	}
L45:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v100))) = uint16(v131)
	v137 = v88 + int32(1)
	if v137 != v77 {
		v88 = v137
		v89 = v118
		v91 = v131 + v119
		goto L33
	} else {
		goto L46
	}
L46:
	;
	goto L34
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v168
	F_errmsg_internal(m, int32(_a_F_UtilityTupleDescriptor_1), v14)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L10
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_UtilityTupleDescriptor_2), int32(2408), int32(_a_F_UtilityTupleDescriptor_3))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v181 = F_GetPortalByName(m, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	if v181 == int32(0) {
		v487 = v2
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181)+92))
	v186 = F_CreateTupleDescCopy(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	return v186
L54:
	;
	if v191 == int32(0) {
		v487 = v2
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191)+64))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	if v196 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v197 = F_CreateTupleDescCopy(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L10
	} else {
		goto L59
	}
L57:
	;
	v200 = int32(0)
	goto L58
L58:
	;
	return v200
L59:
	;
	v200 = v197
	goto L58
L60:
	;
	return v202
L61:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v208 = m.G0
	v210 = v208 - int32(16)
	m.G0 = v210
	v215 = v207
	v216 = int32(_a_F_UtilityTupleDescriptor_4)
	goto L64
L62:
	;
	v300 = int32(0)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	if v300 < v309 {
		goto L92
	} else {
		goto L93
	}
L63:
	;
	if v257 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L64:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v220 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v257 = base.I32_extend8_s(v237) - base.I32_extend8_s(v246)
	goto L63
L66:
	;
	v225 = int32(1)
	if base.Ui32((v220-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L67:
	;
	if v219 != 0 {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v219 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v257 = int32(1)
	goto L63
L71:
	;
	v224 = int32(-1)
	goto L73
L72:
	;
	v224 = int32(0)
	goto L73
L73:
	;
	v257 = v224
	goto L63
L74:
	;
	v237 = v220 | int32(32)
	goto L76
L75:
	;
	v237 = v220
	goto L76
L76:
	;
	if base.Ui32((v219-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v246 = v219 | int32(32)
	goto L79
L78:
	;
	v246 = v219
	goto L79
L79:
	;
	if v237 == v246&int32(255) {
		v215 = v215 + v225
		v216 = v216 + v225
		goto L64
	} else {
		goto L80
	}
L80:
	;
	goto L65
L81:
	;
	v261 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L10
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v287 = F_GetConfigOptionByName(m, v207, v210+int32(12), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L10
	} else {
		goto L88
	}
L84:
	;
	F_TupleDescInitEntry(m, v261, int32(1), int32(_a_F_UtilityTupleDescriptor_5), int32(25), int32(-1), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	F_TupleDescInitEntry(m, v261, int32(2), int32(_a_F_UtilityTupleDescriptor_6), int32(25), int32(-1), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	F_TupleDescInitEntry(m, v261, int32(3), int32(_a_F_UtilityTupleDescriptor_7), int32(25), int32(-1), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	v299 = v261
	goto L62
L88:
	;
	v290 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	F_TupleDescInitEntry(m, v290, int32(1), v293, int32(25), int32(-1), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	v299 = v290
	goto L62
L91:
	;
	m.G0 = v210 + int32(16)
	return v299
L92:
	;
	v313 = v299 + int32(28)
	v320 = v300
	v321 = v309
	v323 = v300
	goto L96
L93:
	;
	v377 = v300
	v384 = v309
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299)+20)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v299)+16)) = v377
	goto L91
L95:
	;
	v377 = v371
	v384 = v350
	goto L94
L96:
	;
	v329 = v313 + v309<<(uint(int32(3))%32) + v320*int32(100)
	v332 = v313 + v320<<(uint(int32(3))%32)
	if v309 != v321 {
		v350 = v321
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v371 = v309
	goto L95
L98:
	;
	v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v332)+2)))
	if v351 <= int32(0) {
		v371 = v320
		goto L95
	} else {
		goto L106
	}
L99:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+7)))
	if v334 != int32(118) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v350 = v320
	goto L98
L101:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+4)))
	if v337 != int32(1) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+6)))
	if v340&int32(6) != 0 {
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v343 = int32(*(*int16)(unsafe.Add(mBase, uint32(v332)+2)))
	if v343 <= int32(0) {
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+90)))
	if v346 != int32(118) {
		v350 = v309
		goto L98
	} else {
		goto L105
	}
L105:
	;
	goto L100
L106:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+90)))
	if v354 == int32(118) {
		v371 = v320
		goto L95
	} else {
		goto L107
	}
L107:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+5)))
	v363 = (v323 + v357 - int32(1)) & (int32(0) - v357)
	if int32(_a_F_UtilityTupleDescriptor_0) < v363 {
		v371 = v320
		goto L95
	} else {
		goto L108
	}
L108:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v332))) = uint16(v363)
	v369 = v320 + int32(1)
	if v369 != v309 {
		v320 = v369
		v321 = v350
		v323 = v363 + v351
		goto L96
	} else {
		goto L109
	}
L109:
	;
	goto L97
L110:
	;
	F_TupleDescInitBuiltinEntry(m, v392, int32(1), int32(_a_F_UtilityTupleDescriptor_8), int32(25))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L10
	} else {
		goto L111
	}
L111:
	;
	v399 = int32(0)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	if v399 < v408 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v487 = v392
	goto L1
L113:
	;
	v412 = v392 + int32(28)
	v419 = v399
	v420 = v408
	v422 = v399
	goto L117
L114:
	;
	v476 = v399
	v483 = v408
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v392)+20)) = v483
	*(*int32)(unsafe.Add(mBase, uint32(v392)+16)) = v476
	goto L112
L116:
	;
	v476 = v470
	v483 = v449
	goto L115
L117:
	;
	v428 = v412 + v408<<(uint(int32(3))%32) + v419*int32(100)
	v431 = v412 + v419<<(uint(int32(3))%32)
	if v408 != v420 {
		v449 = v420
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v470 = v408
	goto L116
L119:
	;
	v450 = int32(*(*int16)(unsafe.Add(mBase, uint32(v431)+2)))
	if v450 <= int32(0) {
		v470 = v419
		goto L116
	} else {
		goto L127
	}
L120:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431)+7)))
	if v433 != int32(118) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v449 = v419
	goto L119
L122:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431)+4)))
	if v436 != int32(1) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431)+6)))
	if v439&int32(6) != 0 {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v442 = int32(*(*int16)(unsafe.Add(mBase, uint32(v431)+2)))
	if v442 <= int32(0) {
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428)+90)))
	if v445 != int32(118) {
		v449 = v408
		goto L119
	} else {
		goto L126
	}
L126:
	;
	goto L121
L127:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428)+90)))
	if v453 == int32(118) {
		v470 = v419
		goto L116
	} else {
		goto L128
	}
L128:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431)+5)))
	v462 = (v422 + v456 - int32(1)) & (int32(0) - v456)
	if int32(_a_F_UtilityTupleDescriptor_0) < v462 {
		v470 = v419
		goto L116
	} else {
		goto L129
	}
L129:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v431))) = uint16(v462)
	v468 = v419 + int32(1)
	if v468 != v408 {
		v419 = v468
		v420 = v449
		v422 = v462 + v450
		goto L117
	} else {
		goto L130
	}
L130:
	;
	goto L118
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
	var v81 int64
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
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+390)))
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
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+390)) = uint8(v124)
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
	*(*int64)(unsafe.Add(mBase, uint32(l0)+360)) = v28
	goto L4
L10:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+360))
	v36 = v34 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+360)) = v36
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
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+412))
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
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+412))
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
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+360))
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
	v93 = *(*int64)(unsafe.Add(mBase, uint32(l0)+360))
	v95 = v93 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+360)) = v95
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
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+412))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
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
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	v81 = m.T0[v80].(func(*base.Module, int32, int32, int32) int64)(m, v67, v63, v10+int32(15))
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
	if v81 == int64(0) {
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
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+412))
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
func F_update_retention_status(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	v1 = l0
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_update_retention_status[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v11 = base.B2i32(v9 == int32(2))
	if v11 == int32(0) {
		F_StartTransactionCommand(m)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = F_GetTransactionSnapshot(m)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_PushActiveSnapshot(m, v18)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_update_retention_status[1]))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
					v25 = m.G0
					v27 = v25 - int32(272)
					m.G0 = v27
					v31 = F_table_open(m, int32(_a_F_update_retention_status_0), int32(3))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v36 = F_SearchSysCacheCopy(m, int32(67), base.I64_extend_i32_u(v24), int64(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							if v36 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v27))) = v24
									F_errmsg_internal(m, int32(_a_F_update_retention_status_1), v27)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_update_retention_status_2), int32(744), int32(_a_F_update_retention_status_3))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_LockSharedObject(m, int32(_a_F_update_retention_status_0), v24, int32(1))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v58 = v27 + int32(16)
									base.MemoryFill(m, v58, int32(0), int32(184))
									v62 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v27)+216)) = v62
									*(*int64)(unsafe.Add(mBase, uint32(v27)+223)) = v62
									*(*int64)(unsafe.Add(mBase, uint32(v27)+255)) = v62
									*(*int64)(unsafe.Add(mBase, uint32(v27)+248)) = v62
									*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v62
									*(*int64)(unsafe.Add(mBase, uint32(v27)+208)) = v62
									*(*int64)(unsafe.Add(mBase, uint32(v27)+136)) = base.I64_extend_i32_u(v1)
									v76 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v27)+223)) = uint8(v76)
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+52))
									v83 = F_heap_modify_tuple(m, v36, v78, v58, v27+int32(240), v27+int32(208))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										F_CatalogTupleUpdate(m, v31, v83+int32(4), v83)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v83)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v31, int32(0))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return int32(0)
												} else {
													m.G0 = v27 + int32(272)
													F_PopActiveSnapshot(m)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int32(0)
													} else {
														F_CommitTransactionCommand(m)
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return int32(0)
														} else {
															v103 = *(*int32)(unsafe.Add(mBase, _c_F_update_retention_status[2]))
															v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
															if v104 != 0 {
																v106 = F_pgmem_kill(m, v104, int32(10))
																mBase = m.M
															} else {
															}
															v108 = *(*int32)(unsafe.Add(mBase, _c_F_update_retention_status[1]))
															*(*uint8)(unsafe.Add(mBase, uint32(v108)+48)) = uint8(v1)
															return v11 ^ int32(1)
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
	} else {
		return v11 ^ int32(1)
	}
}
