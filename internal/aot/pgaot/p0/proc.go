package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcArrayEndTransaction(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int64
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
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
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v245 int64
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	if l1 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
		v12 = F_LWLockConditionalAcquire(m, v8+int32(512), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v12 != 0 {
				v14 = int32(0)
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v19+v20<<(uint(int32(2))%32)))) = v14
				*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v14
				*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v14
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+73)) = uint8(v14)
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)))
				if v34&int32(14) != 0 {
					v38 = v34 & int32(241)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)) = uint8(v38)
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*uint8)(unsafe.Add(mBase, uint32(v42+v43))) = uint8(v38)
				} else {
				}
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+276)))
				if v47 == int32(0) {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)))
					if v50 != int32(1) {
					} else {
						v54 = v20 << (uint(int32(1)) % 32)
						v55 = int32(_a_F_ProcArrayEndTransaction_0)
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
						v59 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v54+v57))) = uint8(v59)
						v62 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
						*(*uint8)(unsafe.Add(mBase, uint32(v63+v54)+1)) = uint8(v59)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+276)) = uint16(v59)
					}
				} else {
					v54 = v20 << (uint(int32(1)) % 32)
					v55 = int32(_a_F_ProcArrayEndTransaction_0)
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
					v59 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v54+v57))) = uint8(v59)
					v62 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
					*(*uint8)(unsafe.Add(mBase, uint32(v63+v54)+1)) = uint8(v59)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+276)) = uint16(v59)
				}
				v70 = int32(_a_F_ProcArrayEndTransaction_1)
				v71 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[2]))
				v72 = *(*int64)(unsafe.Add(mBase, uint32(v71)+48))
				v73 = base.I32_wrap_i64(v72)
				v74 = F_TransactionIdPrecedes(m, v73, l1)
				mBase = m.M
				v76 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[2]))
				if v74 != 0 {
					*(*int64)(unsafe.Add(mBase, uint32(v76)+48)) = v72 + base.I64_extend_i32_s(l1-v73)
				} else {
				}
				v81 = *(*int64)(unsafe.Add(mBase, uint32(v76)+56))
				*(*int64)(unsafe.Add(mBase, uint32(v76)+56)) = v81 + int64(1)
				v317 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
				F_LWLockRelease(m, v317+int32(512))
				mBase = m.M
				v321 = m.ExcPending
				if v321 != 0 {
					return
				} else {
					return
				}
			} else {
				v86 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+544)) = l1
				v89 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+536)) = uint8(v89)
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)+52))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+540)) = v91
				v95 = base.I32_div_s(l0-v87, int32(640))
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v86)+52))
				v97 = base.B2i32(v96 == v91)
				if v96 == v91 {
					v98 = v95
				} else {
					v98 = v96
				}
				*(*int32)(unsafe.Add(mBase, uint32(v86)+52)) = v98
				if v97 == int32(0) {
					v103 = v96
					for {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+540)) = v103
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v86)+52))
						if v109 == v103 {
							v111 = v95
						} else {
							v111 = v109
						}
						*(*int32)(unsafe.Add(mBase, uint32(v86)+52)) = v111
						if v109 != v103 {
							v103 = v109
							continue
						} else {
							break
						}
						break
					}
					v116 = v103
				} else {
					v116 = v91
				}
				if v116 != int32(-1) {
					v124 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[3]))
					*(*int32)(unsafe.Add(mBase, uint32(v124))) = int32(134217769)
					v129 = int32(0)
					for {
						v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+536)))
						if v136 != 0 {
							v129 = v129 + int32(1)
							continue
						} else {
							break
						}
						break
					}
					v138 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[3]))
					v139 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v138))) = v139
					if v129 <= v139 {
					} else {
						v144 = v129
						for {
							v150 = int32(1)
							if base.Ui32(v150) < base.Ui32(v144) {
								v144 = v144 - v150
								continue
							} else {
								break
							}
							break
						}
					}
					return
				} else {
					v155 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
					v159 = F_LWLockAcquire(m, v155+int32(512), int32(0))
					mBase = m.M
					v160 = m.ExcPending
					if v160 != 0 {
						return
					} else {
						v161 = *(*int32)(unsafe.Add(mBase, uint32(v86)+52))
						v162 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v86)+52)) = v162
						if v161 == v162 {
							v317 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
							F_LWLockRelease(m, v317+int32(512))
							mBase = m.M
							v321 = m.ExcPending
							if v321 != 0 {
								return
							} else {
								return
							}
						} else {
							v166 = v161
							for {
								v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[4]))
								v176 = v173 + v166*int32(640)
								v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+544))
								v178 = int32(0)
								v182 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
								v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
								v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v183+v184<<(uint(int32(2))%32)))) = v178
								*(*int32)(unsafe.Add(mBase, uint32(v176)+56)) = v178
								*(*int64)(unsafe.Add(mBase, uint32(v176)+36)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v176)+120)) = v178
								*(*uint8)(unsafe.Add(mBase, uint32(v176)+73)) = uint8(v178)
								v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+124)))
								if v198&int32(14) != 0 {
									v202 = v198 & int32(241)
									*(*uint8)(unsafe.Add(mBase, uint32(v176)+124)) = uint8(v202)
									v205 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
									v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+12))
									v207 = *(*int32)(unsafe.Add(mBase, uint32(v176)+48))
									*(*uint8)(unsafe.Add(mBase, uint32(v206+v207))) = uint8(v202)
								} else {
								}
								v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+276)))
								if v211 == int32(0) {
									v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+277)))
									if v214 != int32(1) {
									} else {
										v218 = v184 << (uint(int32(1)) % 32)
										v219 = int32(_a_F_ProcArrayEndTransaction_0)
										v220 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
										v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
										v223 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v218+v221))) = uint8(v223)
										v226 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
										v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
										*(*uint8)(unsafe.Add(mBase, uint32(v227+v218)+1)) = uint8(v223)
										*(*uint16)(unsafe.Add(mBase, uint32(v176)+276)) = uint16(v223)
									}
								} else {
									v218 = v184 << (uint(int32(1)) % 32)
									v219 = int32(_a_F_ProcArrayEndTransaction_0)
									v220 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
									v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
									v223 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v218+v221))) = uint8(v223)
									v226 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
									v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
									*(*uint8)(unsafe.Add(mBase, uint32(v227+v218)+1)) = uint8(v223)
									*(*uint16)(unsafe.Add(mBase, uint32(v176)+276)) = uint16(v223)
								}
								v234 = int32(_a_F_ProcArrayEndTransaction_1)
								v235 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[2]))
								v236 = *(*int64)(unsafe.Add(mBase, uint32(v235)+48))
								v237 = base.I32_wrap_i64(v236)
								v238 = F_TransactionIdPrecedes(m, v237, v177)
								mBase = m.M
								v240 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[2]))
								if v238 != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(v240)+48)) = v236 + base.I64_extend_i32_s(v177-v237)
								} else {
								}
								v245 = *(*int64)(unsafe.Add(mBase, uint32(v240)+56))
								*(*int64)(unsafe.Add(mBase, uint32(v240)+56)) = v245 + int64(1)
								v249 = *(*int32)(unsafe.Add(mBase, uint32(v176)+540))
								if v249 != int32(-1) {
									v166 = v249
									continue
								} else {
									break
								}
								break
							}
							v253 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
							F_LWLockRelease(m, v253+int32(512))
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return
							} else {
								v259 = v161
								for {
									v265 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[4]))
									v268 = v265 + v259*int32(640)
									v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+540))
									*(*int32)(unsafe.Add(mBase, uint32(v268)+540)) = int32(-1)
									v272 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v268)+536)) = uint8(v272)
									v275 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[5]))
									if v275 != v268 {
									} else {
									}
									if v269 != int32(-1) {
										v259 = v269
										continue
									} else {
										break
									}
									break
								}
								return
							}
						}
					}
				}
			}
		}
	} else {
		v280 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v280
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v280
		*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v280
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+73)) = uint8(v280)
		v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)))
		if v288&int32(14) == v280 {
			return
		} else {
			v294 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
			v298 = F_LWLockAcquire(m, v294+int32(512), int32(0))
			mBase = m.M
			v299 = m.ExcPending
			if v299 != 0 {
				return
			} else {
				v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)))
				v302 = v300 & int32(-15)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)) = uint8(v302)
				v305 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
				v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+12))
				v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*uint8)(unsafe.Add(mBase, uint32(v306+v307))) = uint8(v302)
				v317 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
				F_LWLockRelease(m, v317+int32(512))
				mBase = m.M
				v321 = m.ExcPending
				if v321 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_ProcLockWakeup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
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
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v9 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = l1 + int32(32)
	if v12 == v16 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = v12
	v23 = v3
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32))))
	if v32&v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L1
L7:
	;
	if v26 != v16 {
		v22 = v26
		v23 = v99
		goto L5
	} else {
		goto L19
	}
L8:
	;
	v99 = int32(1)<<(uint(v28)%32) | v23
	goto L7
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
	v35 = F_LockCheckConflicts(m, l0, v28, l1, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v35 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	v41 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+128)) = v40 + v41
	v46 = l1 + v28<<(uint(int32(2))%32)
	v48 = v46 + int32(88)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v49 + v41
	v54 = v41 << (uint(v28) % 32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v54 | v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v58 == v59 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v69 == int32(0) {
		v99 = v23
		goto L7
	} else {
		goto L17
	}
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v61 & (v54 ^ int32(-1))
	goto L16
L15:
	;
	goto L16
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v66 | v54
	goto L13
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v69
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v75
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+40)) = v79 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+92)) = v77
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_ProcLockWakeup[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v88)+112)) = v77
	F_SetLatch(m, v22+int32(20))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v99 = v23
	goto L7
L19:
	;
	goto L6
}
func F_ProcSendSignal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	if int32(0) <= l0 {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSendSignal[0]))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
		if base.Ui32(l0) < base.Ui32(v7) {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			F_SetLatch(m, v23+l0*int32(640)+int32(20))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_ProcSendSignal_0), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ProcSendSignal_1), int32(1989), int32(_a_F_ProcSendSignal_2))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_ProcSendSignal_0), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ProcSendSignal_1), int32(1989), int32(_a_F_ProcSendSignal_2))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_ProcWaitForSignal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ProcWaitForSignal[0]))
	v6 = F_WaitLatch(m, v3, int32(33), int32(0), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_ProcWaitForSignal[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_ProcWaitForSignal[1]))
		if v13 != 0 {
			F_ProcessInterrupts(m)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_ProcessProcSignalBarrier(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	v1 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(192)
	m.G0 = v10
	v14 = int32(-1)
	v15 = v1
	v16 = v1
	v17 = v1
	v18 = int64(0)
	goto L1
L1:
	;
	goto L4
L2:
	;
	m.G0 = v10 + int32(192)
	return
L3:
	;
	goto L2
L4:
	;
	if v14 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v156 = int32(m.ExcTag)
	v157 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v156 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L7:
	;
	v137 = int32(_a_F_ProcessProcSignalBarrier_0)
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v138)+104)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v10)+176)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v10)+184)) = v135
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0]))
	F_ConditionVariableBroadcast(m, v144+int32(116))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L32
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[1]))
	if v23 == int32(0) {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	v60 = v15
	v61 = v16
	v62 = v17
	v63 = v18
	goto L10
L10:
	;
	if v61 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[1])) = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0]))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+104)) = v31
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[2]))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
	*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
	if v35 == v31 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0]))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+112))
	v41 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+112)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v40
	if v40 == v41 {
		v132 = v15
		v134 = v17
		v135 = v35
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[3]))
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[4]))
	goto L14
L14:
	;
	v53 = v10 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v10 + int32(12)
	goto L17
L15:
	;
	v60 = v49
	v61 = int32(0)
	v62 = v51
	v63 = v35
	goto L10
L17:
	;
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[3])) = v60
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[4])) = v62
	v132 = v60
	v134 = v62
	v135 = v63
	goto L7
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[4])) = v10 + int32(16)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v10)+172))
	if v72 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[3])) = v60
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[4])) = v62
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+172))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+112)) = v111 | v112
	v116 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[1])) = v116
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[5])) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v10)+176)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v10)+184)) = v63
	F_pg_re_throw(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L31
	}
L22:
	;
	goto L23
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+172))
	v83 = base.I32_ctz(v82)
	if v83 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[3])) = v60
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[4])) = v62
	v132 = v60
	v134 = v62
	v135 = v63
	goto L7
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)+172))
	if v100 != 0 {
		goto L23
	} else {
		goto L30
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v10)+176)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v10)+184)) = v63
	F_smgrreleaseall(m)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v10)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v95 & base.I32_rotl(int32(-2), v83)
	goto L25
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v10)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v91 & int32(-2)
	goto L25
L30:
	;
	goto L24
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	goto L5
L33:
	;
	v161 = int32(v157)
	m.G0 = v10
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v10+int32(12) == v167 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	m.ExcPending = 1
	goto L42
L35:
	;
	if v171 != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v171 = v169
	goto L38
L37:
	;
	v171 = int32(0)
	goto L38
L38:
	;
	goto L35
L39:
	;
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v10)+184))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v10)+180))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v10)+176))
	v14 = v171
	v15 = v174
	v16 = v163
	v17 = v173
	v18 = v172
	goto L1
L40:
	;
	goto L41
L41:
	;
	F___wasm_longjmp(m, v164, v163)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	return
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SendProcSignal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	if l2 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+l1<<(uint(int32(2))%32))+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+96)) = int32(0)
	v100 = F_kill(m, l0, int32(10))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L10
	} else {
		goto L26
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[0])) = int32(71)
	return int32(-1)
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0)
	goto L2
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[1]))
	v13 = v10 + l2<<(uint(int32(7))%32)
	v15 = v13 + int32(104)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[2]))
	v33 = v31 + int32(37)
	if v33 < int32(0) {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	F_s_lock(m, v15, int32(_a_F_SendProcSignal_0), int32(293), int32(_a_F_SendProcSignal_1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v27 = v13 + int32(8)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v28 != l0 {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L9
L12:
	;
	v89 = v27
	goto L1
L13:
	;
	v38 = v33
	goto L14
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[1]))
	v46 = v43 + v38<<(uint(int32(7))%32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if l0 == v47 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v89 = v50
	goto L1
L16:
	;
	goto L15
L17:
	;
	v50 = v46 + int32(8)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+96)) = int32(1)
	v55 = v46 + int32(104)
	if v51 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v67 = int32(0)
	if base.B2i32(v38 <= v67) == v67 {
		v38 = v38 - int32(1)
		goto L14
	} else {
		goto L25
	}
L20:
	;
	F_s_lock(m, v55, int32(_a_F_SendProcSignal_0), int32(321), int32(_a_F_SendProcSignal_1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v61 == l0 {
		goto L16
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(0)
	goto L19
L25:
	;
	goto L2
L26:
	;
	return v100
}
func F_assignProcTypes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_SearchSysCache1(m, int32(47), v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v18 = v16 + v17
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if l3 == v19 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if l2 != 0 {
					if v21 != l2 {
						v24 = v21
					} else {
						v24 = int32(0)
					}
					if v24 == int32(0) {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if base.B2i32(v27 == int32(0))|base.B2i32(v27 == l2) != 0 {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
							if v52 != int32(2278) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_0), int32(0))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_assignProcTypes_1)
											F_errhint(m, int32(_a_F_assignProcTypes_2), v10+int32(16))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1241), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
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
								v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
								if v55 != int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_0), int32(0))
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_assignProcTypes_1)
												F_errhint(m, int32(_a_F_assignProcTypes_2), v10+int32(16))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1241), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return
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
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
									if v58 == int32(2281) {
										v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v245 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
											v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v249 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											} else {
											}
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										} else {
											v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v252 != 0 {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_0), int32(0))
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_assignProcTypes_1)
													F_errhint(m, int32(_a_F_assignProcTypes_2), v10+int32(16))
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1241), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return
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
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_assignProcTypes_6), int32(0))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1224), int32(_a_F_assignProcTypes_4))
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
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
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_assignProcTypes_6), int32(0))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1224), int32(_a_F_assignProcTypes_4))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
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
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v21 != v49 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v278 = m.ExcPending
						if v278 != 0 {
							return
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v281 = m.ExcPending
							if v281 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_assignProcTypes_7), int32(0))
								mBase = m.M
								v285 = m.ExcPending
								if v285 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1231), int32(_a_F_assignProcTypes_4))
									mBase = m.M
									v290 = m.ExcPending
									if v290 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
						if v52 != int32(2278) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_assignProcTypes_0), int32(0))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_assignProcTypes_1)
										F_errhint(m, int32(_a_F_assignProcTypes_2), v10+int32(16))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1241), int32(_a_F_assignProcTypes_4))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
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
							v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v55 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_0), int32(0))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_assignProcTypes_1)
											F_errhint(m, int32(_a_F_assignProcTypes_2), v10+int32(16))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1241), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
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
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
								if v58 == int32(2281) {
									v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v245 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
										v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v249 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										} else {
										}
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									} else {
										v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v252 != 0 {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_0), int32(0))
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_assignProcTypes_1)
												F_errhint(m, int32(_a_F_assignProcTypes_2), v10+int32(16))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1241), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return
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
							}
						}
					}
				}
			} else {
				v85 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+10)))
					if v87 == int32(1) {
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						switch v90 - int32(1) {
						case 0:
							v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v93 != int32(2) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v294 = m.ExcPending
								if v294 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v297 = m.ExcPending
									if v297 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_8), int32(0))
										mBase = m.M
										v301 = m.ExcPending
										if v301 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1259), int32(_a_F_assignProcTypes_4))
											mBase = m.M
											v306 = m.ExcPending
											if v306 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
								if v96 != int32(23) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v310 = m.ExcPending
									if v310 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v313 = m.ExcPending
										if v313 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_9), int32(0))
											mBase = m.M
											v317 = m.ExcPending
											if v317 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1263), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v322 = m.ExcPending
												if v322 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v99 == int32(0) {
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v102
									} else {
									}
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v104 != 0 {
									} else {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v105
									}
									v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v245 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
										v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v249 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										} else {
										}
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									} else {
										v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v252 != 0 {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						case 1:
							v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v107 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v326 = m.ExcPending
								if v326 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v329 = m.ExcPending
									if v329 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_10), int32(0))
										mBase = m.M
										v333 = m.ExcPending
										if v333 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1280), int32(_a_F_assignProcTypes_4))
											mBase = m.M
											v338 = m.ExcPending
											if v338 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
								if v110 != int32(2281) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v326 = m.ExcPending
									if v326 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v329 = m.ExcPending
										if v329 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_10), int32(0))
											mBase = m.M
											v333 = m.ExcPending
											if v333 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1280), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v338 = m.ExcPending
												if v338 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
									if v113 == int32(2278) {
										v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v245 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
											v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v249 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											} else {
											}
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										} else {
											v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v252 != 0 {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_11), int32(0))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1284), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return
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
							}
						case 2:
							v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v132 != int32(5) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v342 = m.ExcPending
								if v342 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v345 = m.ExcPending
									if v345 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_12), int32(0))
										mBase = m.M
										v349 = m.ExcPending
										if v349 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1295), int32(_a_F_assignProcTypes_4))
											mBase = m.M
											v354 = m.ExcPending
											if v354 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
								if v135 != int32(16) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v358 = m.ExcPending
									if v358 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v361 = m.ExcPending
										if v361 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_13), int32(0))
											mBase = m.M
											v365 = m.ExcPending
											if v365 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1299), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v370 = m.ExcPending
												if v370 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v138 == int32(0) {
										v141 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v141
									} else {
									}
									v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v143 != 0 {
									} else {
										v144 = *(*int32)(unsafe.Add(mBase, uint32(v18)+144))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v144
									}
									v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v245 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
										v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v249 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										} else {
										}
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									} else {
										v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v252 != 0 {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						case 3:
							v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v146 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v374 = m.ExcPending
								if v374 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v377 = m.ExcPending
									if v377 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_14), int32(0))
										mBase = m.M
										v381 = m.ExcPending
										if v381 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1315), int32(_a_F_assignProcTypes_4))
											mBase = m.M
											v386 = m.ExcPending
											if v386 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v149 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
								if v149 != int32(16) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v390 = m.ExcPending
									if v390 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v393 = m.ExcPending
										if v393 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_15), int32(0))
											mBase = m.M
											v397 = m.ExcPending
											if v397 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1319), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v402 = m.ExcPending
												if v402 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v152 == v153 {
										v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v245 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
											v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v249 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											} else {
											}
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										} else {
											v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v252 != 0 {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v158 = m.ExcPending
										if v158 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_16), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1332), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v170 = m.ExcPending
													if v170 != 0 {
														return
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
							}
						default:
							v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v245 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
								v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v249 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
								} else {
								}
								if l2 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v486 = m.ExcPending
									if v486 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v489 = m.ExcPending
										if v489 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
											mBase = m.M
											v493 = m.ExcPending
											if v493 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v498 = m.ExcPending
												if v498 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									F_ReleaseCatCache(m, v14)
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return
									} else {
										m.G0 = v10 + int32(32)
										return
									}
								}
							} else {
								v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v252 != 0 {
									F_ReleaseCatCache(m, v14)
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return
									} else {
										m.G0 = v10 + int32(32)
										return
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
									if l2 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v486 = m.ExcPending
										if v486 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v489 = m.ExcPending
											if v489 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
												mBase = m.M
												v493 = m.ExcPending
												if v493 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v498 = m.ExcPending
													if v498 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										F_ReleaseCatCache(m, v14)
										mBase = m.M
										v257 = m.ExcPending
										if v257 != 0 {
											return
										} else {
											m.G0 = v10 + int32(32)
											return
										}
									}
								}
							}
						case 5:
							v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v171 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v406 = m.ExcPending
								if v406 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v409 = m.ExcPending
									if v409 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_17), int32(0))
										mBase = m.M
										v413 = m.ExcPending
										if v413 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1340), int32(_a_F_assignProcTypes_4))
											mBase = m.M
											v418 = m.ExcPending
											if v418 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v174 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
								if v174 != int32(2281) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v406 = m.ExcPending
									if v406 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v409 = m.ExcPending
										if v409 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_17), int32(0))
											mBase = m.M
											v413 = m.ExcPending
											if v413 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1340), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v418 = m.ExcPending
												if v418 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v177 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
									if v177 != int32(2278) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v422 = m.ExcPending
										if v422 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v425 = m.ExcPending
											if v425 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_18), int32(0))
												mBase = m.M
												v429 = m.ExcPending
												if v429 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1344), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v434 = m.ExcPending
													if v434 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v180 == v181 {
											v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v245 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
												v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v249 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												} else {
												}
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											} else {
												v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v252 != 0 {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
													if l2 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v486 = m.ExcPending
														if v486 != 0 {
															return
														} else {
															F_errcode(m, int32(117833860))
															mBase = m.M
															v489 = m.ExcPending
															if v489 != 0 {
																return
															} else {
																F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
																mBase = m.M
																v493 = m.ExcPending
																if v493 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
																	mBase = m.M
																	v498 = m.ExcPending
																	if v498 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														F_ReleaseCatCache(m, v14)
														mBase = m.M
														v257 = m.ExcPending
														if v257 != 0 {
															return
														} else {
															m.G0 = v10 + int32(32)
															return
														}
													}
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v186 = m.ExcPending
											if v186 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v189 = m.ExcPending
												if v189 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_19), int32(0))
													mBase = m.M
													v193 = m.ExcPending
													if v193 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1357), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v198 = m.ExcPending
														if v198 != 0 {
															return
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
								}
							}
						}
					} else {
						v200 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+12)))
							if v202 != int32(1) {
								v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v245 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
									v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v249 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
									} else {
									}
									if l2 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v486 = m.ExcPending
										if v486 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v489 = m.ExcPending
											if v489 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
												mBase = m.M
												v493 = m.ExcPending
												if v493 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v498 = m.ExcPending
													if v498 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										F_ReleaseCatCache(m, v14)
										mBase = m.M
										v257 = m.ExcPending
										if v257 != 0 {
											return
										} else {
											m.G0 = v10 + int32(32)
											return
										}
									}
								} else {
									v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v252 != 0 {
										F_ReleaseCatCache(m, v14)
										mBase = m.M
										v257 = m.ExcPending
										if v257 != 0 {
											return
										} else {
											m.G0 = v10 + int32(32)
											return
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								switch v205 - int32(1) {
								case 0:
									v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
									if v208 != int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v438 = m.ExcPending
										if v438 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v441 = m.ExcPending
											if v441 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_20), int32(0))
												mBase = m.M
												v445 = m.ExcPending
												if v445 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1367), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v450 = m.ExcPending
													if v450 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v211 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
										if v211 == int32(23) {
											v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v236 == int32(0) {
												v239 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v239
											} else {
											}
											v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v241 != 0 {
											} else {
												v242 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v242
											}
											v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v245 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
												v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v249 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												} else {
												}
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											} else {
												v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v252 != 0 {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
													if l2 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v486 = m.ExcPending
														if v486 != 0 {
															return
														} else {
															F_errcode(m, int32(117833860))
															mBase = m.M
															v489 = m.ExcPending
															if v489 != 0 {
																return
															} else {
																F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
																mBase = m.M
																v493 = m.ExcPending
																if v493 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
																	mBase = m.M
																	v498 = m.ExcPending
																	if v498 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														F_ReleaseCatCache(m, v14)
														mBase = m.M
														v257 = m.ExcPending
														if v257 != 0 {
															return
														} else {
															m.G0 = v10 + int32(32)
															return
														}
													}
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v217 = m.ExcPending
											if v217 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v220 = m.ExcPending
												if v220 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_21), int32(0))
													mBase = m.M
													v224 = m.ExcPending
													if v224 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1371), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v229 = m.ExcPending
														if v229 != 0 {
															return
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
								case 1:
									v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
									if v230 != int32(2) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v454 = m.ExcPending
										if v454 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v457 = m.ExcPending
											if v457 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_22), int32(0))
												mBase = m.M
												v461 = m.ExcPending
												if v461 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1378), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v466 = m.ExcPending
													if v466 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v233 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
										if v233 != int32(20) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v470 = m.ExcPending
											if v470 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v473 = m.ExcPending
												if v473 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_23), int32(0))
													mBase = m.M
													v477 = m.ExcPending
													if v477 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1382), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v482 = m.ExcPending
														if v482 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v236 == int32(0) {
												v239 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v239
											} else {
											}
											v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v241 != 0 {
											} else {
												v242 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v242
											}
											v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v245 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
												v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v249 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												} else {
												}
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											} else {
												v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v252 != 0 {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
													if l2 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v486 = m.ExcPending
														if v486 != 0 {
															return
														} else {
															F_errcode(m, int32(117833860))
															mBase = m.M
															v489 = m.ExcPending
															if v489 != 0 {
																return
															} else {
																F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
																mBase = m.M
																v493 = m.ExcPending
																if v493 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
																	mBase = m.M
																	v498 = m.ExcPending
																	if v498 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														F_ReleaseCatCache(m, v14)
														mBase = m.M
														v257 = m.ExcPending
														if v257 != 0 {
															return
														} else {
															m.G0 = v10 + int32(32)
															return
														}
													}
												}
											}
										}
									}
								default:
									v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v236 == int32(0) {
										v239 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v239
									} else {
									}
									v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v241 != 0 {
									} else {
										v242 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v242
									}
									v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v245 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
										v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v249 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										} else {
										}
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									} else {
										v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v252 != 0 {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v264 = m.ExcPending
			if v264 != 0 {
				return
			} else {
				v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v265
				F_errmsg_internal(m, int32(_a_F_assignProcTypes_24), v10)
				mBase = m.M
				v269 = m.ExcPending
				if v269 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1212), int32(_a_F_assignProcTypes_4))
					mBase = m.M
					v274 = m.ExcPending
					if v274 != 0 {
						return
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
