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
		v8 = *(*int32)(unsafe.Add(mBase, _consts[29]))
		v12 = F_LWLockConditionalAcquire(m, v8+int32(512), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v12 != 0 {
				v14 = int32(0)
				v18 = *(*int32)(unsafe.Add(mBase, _consts[91]))
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
					v41 = *(*int32)(unsafe.Add(mBase, _consts[91]))
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
						v55 = int32(4485520)
						v56 = *(*int32)(unsafe.Add(mBase, _consts[91]))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
						v59 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v54+v57))) = uint8(v59)
						v62 = *(*int32)(unsafe.Add(mBase, _consts[91]))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
						*(*uint8)(unsafe.Add(mBase, uint32(v63+v54)+1)) = uint8(v59)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+276)) = uint16(v59)
					}
				} else {
					v54 = v20 << (uint(int32(1)) % 32)
					v55 = int32(4485520)
					v56 = *(*int32)(unsafe.Add(mBase, _consts[91]))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
					v59 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v54+v57))) = uint8(v59)
					v62 = *(*int32)(unsafe.Add(mBase, _consts[91]))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
					*(*uint8)(unsafe.Add(mBase, uint32(v63+v54)+1)) = uint8(v59)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+276)) = uint16(v59)
				}
				v70 = int32(4457112)
				v71 = *(*int32)(unsafe.Add(mBase, _consts[64]))
				v72 = *(*int64)(unsafe.Add(mBase, uint32(v71)+48))
				v73 = base.I32_wrap_i64(v72)
				v74 = F_TransactionIdPrecedes(m, v73, l1)
				mBase = m.M
				v76 = *(*int32)(unsafe.Add(mBase, _consts[64]))
				if v74 != 0 {
					*(*int64)(unsafe.Add(mBase, uint32(v76)+48)) = v72 + base.I64_extend_i32_s(l1-v73)
				} else {
				}
				v81 = *(*int64)(unsafe.Add(mBase, uint32(v76)+56))
				*(*int64)(unsafe.Add(mBase, uint32(v76)+56)) = v81 + int64(1)
				v317 = *(*int32)(unsafe.Add(mBase, _consts[29]))
				F_LWLockRelease(m, v317+int32(512))
				mBase = m.M
				v321 = m.ExcPending
				if v321 != 0 {
					return
				} else {
					return
				}
			} else {
				v86 = *(*int32)(unsafe.Add(mBase, _consts[91]))
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
					v124 = *(*int32)(unsafe.Add(mBase, _consts[162]))
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
					v138 = *(*int32)(unsafe.Add(mBase, _consts[162]))
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
					v155 = *(*int32)(unsafe.Add(mBase, _consts[29]))
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
							v317 = *(*int32)(unsafe.Add(mBase, _consts[29]))
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
								v173 = *(*int32)(unsafe.Add(mBase, _consts[602]))
								v176 = v173 + v166*int32(640)
								v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+544))
								v178 = int32(0)
								v182 = *(*int32)(unsafe.Add(mBase, _consts[91]))
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
									v205 = *(*int32)(unsafe.Add(mBase, _consts[91]))
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
										v219 = int32(4485520)
										v220 = *(*int32)(unsafe.Add(mBase, _consts[91]))
										v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
										v223 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v218+v221))) = uint8(v223)
										v226 = *(*int32)(unsafe.Add(mBase, _consts[91]))
										v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
										*(*uint8)(unsafe.Add(mBase, uint32(v227+v218)+1)) = uint8(v223)
										*(*uint16)(unsafe.Add(mBase, uint32(v176)+276)) = uint16(v223)
									}
								} else {
									v218 = v184 << (uint(int32(1)) % 32)
									v219 = int32(4485520)
									v220 = *(*int32)(unsafe.Add(mBase, _consts[91]))
									v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
									v223 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v218+v221))) = uint8(v223)
									v226 = *(*int32)(unsafe.Add(mBase, _consts[91]))
									v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
									*(*uint8)(unsafe.Add(mBase, uint32(v227+v218)+1)) = uint8(v223)
									*(*uint16)(unsafe.Add(mBase, uint32(v176)+276)) = uint16(v223)
								}
								v234 = int32(4457112)
								v235 = *(*int32)(unsafe.Add(mBase, _consts[64]))
								v236 = *(*int64)(unsafe.Add(mBase, uint32(v235)+48))
								v237 = base.I32_wrap_i64(v236)
								v238 = F_TransactionIdPrecedes(m, v237, v177)
								mBase = m.M
								v240 = *(*int32)(unsafe.Add(mBase, _consts[64]))
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
							v253 = *(*int32)(unsafe.Add(mBase, _consts[29]))
							F_LWLockRelease(m, v253+int32(512))
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return
							} else {
								v259 = v161
								for {
									v265 = *(*int32)(unsafe.Add(mBase, _consts[602]))
									v268 = v265 + v259*int32(640)
									v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+540))
									*(*int32)(unsafe.Add(mBase, uint32(v268)+540)) = int32(-1)
									v272 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v268)+536)) = uint8(v272)
									v275 = *(*int32)(unsafe.Add(mBase, _consts[88]))
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
			v294 = *(*int32)(unsafe.Add(mBase, _consts[29]))
			v298 = F_LWLockAcquire(m, v294+int32(512), int32(0))
			mBase = m.M
			v299 = m.ExcPending
			if v299 != 0 {
				return
			} else {
				v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)))
				v302 = v300 & int32(-15)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)) = uint8(v302)
				v305 = *(*int32)(unsafe.Add(mBase, _consts[91]))
				v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+12))
				v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*uint8)(unsafe.Add(mBase, uint32(v306+v307))) = uint8(v302)
				v317 = *(*int32)(unsafe.Add(mBase, _consts[29]))
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
	v88 = *(*int32)(unsafe.Add(mBase, _consts[88]))
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
		v6 = *(*int32)(unsafe.Add(mBase, _consts[91]))
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
				F_errmsg_internal(m, int32(421661), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					F_errfinish(m, int32(525187), int32(1989), int32(329151))
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
			F_errmsg_internal(m, int32(421661), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(525187), int32(1989), int32(329151))
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	v6 = F_WaitLatch(m, v3, int32(33), int32(0), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[77]))
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0)
		v13 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	v1 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v17 = v1
	v18 = v1
	v19 = int32(-1)
	v20 = v1
	v21 = v1
	v22 = v1
	v23 = v13
	v24 = int64(0)
	goto L1
L1:
	;
	goto L4
L2:
	;
	m.G0 = v13 + int32(32)
	return
L3:
	;
	goto L2
L4:
	;
	if v19 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v204 = int32(m.ExcTag)
	v205 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v204 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L7:
	;
	v180 = int32(4478868)
	v181 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	*(*int64)(unsafe.Add(mBase, uint32(v181)+104)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v175
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v171
	v189 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	F_ConditionVariableBroadcast(m, v189+int32(116))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		v201 = v177
		goto L6
	} else {
		goto L36
	}
L8:
	;
	v29 = v23 - int32(16)
	m.G0 = v29
	v32 = v29 - int32(160)
	m.G0 = v32
	v35 = *(*int32)(unsafe.Add(mBase, _consts[489]))
	if v35 == int32(0) {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	v69 = v17
	v70 = v18
	v72 = v20
	v73 = v21
	v74 = v22
	v75 = v23
	v76 = v24
	goto L10
L10:
	;
	if v70 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[489])) = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+104)) = v43
	v46 = *(*int32)(unsafe.Add(mBase, _consts[605]))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = v47
	if v47 == v43 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+112))
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v52
	if v52 == v53 {
		v171 = v29
		v174 = v20
		v175 = v21
		v176 = v32
		v177 = v32
		v178 = v47
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[74]))
	v63 = *(*int32)(unsafe.Add(mBase, _consts[291]))
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v13 + int32(4)
	goto L17
L15:
	;
	v69 = v29
	v70 = int32(0)
	v72 = v61
	v73 = v63
	v74 = v32
	v75 = v32
	v76 = v47
	goto L10
L17:
	;
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v72
	*(*int32)(unsafe.Add(mBase, _consts[291])) = v73
	v171 = v69
	v174 = v72
	v175 = v73
	v176 = v74
	v177 = v75
	v178 = v76
	goto L7
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[291])) = v74
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v83 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v72
	*(*int32)(unsafe.Add(mBase, _consts[291])) = v73
	v148 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+112))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+112)) = v149 | v150
	v154 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[489])) = v154
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v69
	F_pg_re_throw(m)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		v201 = v75
		goto L6
	} else {
		goto L35
	}
L22:
	;
	goto L23
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v97 = base.I32_ctz(v96)
	if v97 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v118 != 0 {
		goto L23
	} else {
		goto L32
	}
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v98 & base.I32_rotl(int32(-2), v97)
	goto L26
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v69
	F_smgrreleaseall(m)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		v201 = v75
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v111 & int32(-2)
	goto L31
L31:
	;
	goto L26
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v72
	*(*int32)(unsafe.Add(mBase, _consts[291])) = v73
	v171 = v69
	v174 = v72
	v175 = v73
	v176 = v74
	v177 = v75
	v178 = v76
	goto L7
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	goto L5
L37:
	;
	v209 = int32(v205)
	m.G0 = v201
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v13+int32(4) == v216 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	m.ExcPending = 1
	goto L46
L39:
	;
	if v219 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	v219 = v218
	goto L42
L41:
	;
	v219 = int32(0)
	goto L42
L42:
	;
	goto L39
L43:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v222 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v17 = v220
	v18 = v211
	v19 = v219
	v20 = v224
	v21 = v223
	v22 = v221
	v23 = v201
	v24 = v222
	goto L1
L44:
	;
	goto L45
L45:
	;
	F___wasm_longjmp(m, v212, v211)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	return
L47:
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
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	if l2 != int32(-1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(71)
	return int32(-1)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0)
	goto L1
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _consts[605]))
	v13 = v10 + l2<<(uint(int32(7))%32)
	v15 = v13 + int32(104)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
	v20 = v13 + int32(8)
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[606]))
	v44 = v42 + int32(37)
	if v44 < int32(0) {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	F_s_lock(m, v15, int32(522536), int32(293), int32(329166))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v28 != l0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+l1<<(uint(int32(2))%32))+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = int32(0)
	v38 = F_kill(m, l0, int32(10))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	return v38
L13:
	;
	v49 = v44
	goto L14
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[605]))
	v57 = v54 + v49<<(uint(int32(7))%32)
	v59 = v57 + int32(8)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if l0 == v60 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59+l1<<(uint(int32(2))%32))+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+96)) = int32(0)
	v92 = F_kill(m, l0, int32(10))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L26
	}
L16:
	;
	goto L15
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+96)) = int32(1)
	v66 = v57 + int32(104)
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v78 = int32(0)
	if base.B2i32(v49 <= v78) == v78 {
		v49 = v49 - int32(1)
		goto L14
	} else {
		goto L25
	}
L20:
	;
	F_s_lock(m, v66, int32(522536), int32(321), int32(329166))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v72 == l0 {
		goto L16
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = int32(0)
	goto L19
L25:
	;
	goto L1
L26:
	;
	return v92
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
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
					if l2 != v21 {
						v24 = v21
					} else {
						v24 = int32(0)
					}
					if v24 == int32(0) {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v27 == int32(0) {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
							if v51 != int32(2278) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_errmsg(m, int32(265783), int32(0))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(454496)
											F_errhint(m, int32(632419), v10+int32(16))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												F_errfinish(m, int32(518532), int32(1241), int32(172761))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
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
								v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
								if v54 != int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											F_errmsg(m, int32(265783), int32(0))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(454496)
												F_errhint(m, int32(632419), v10+int32(16))
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													F_errfinish(m, int32(518532), int32(1241), int32(172761))
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
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
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
									if v57 == int32(2281) {
										v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v244 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
											v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v248 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											} else {
											}
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v485 = m.ExcPending
												if v485 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v488 = m.ExcPending
													if v488 != 0 {
														return
													} else {
														F_errmsg(m, int32(264620), int32(0))
														mBase = m.M
														v492 = m.ExcPending
														if v492 != 0 {
															return
														} else {
															F_errfinish(m, int32(518532), int32(1407), int32(172761))
															mBase = m.M
															v497 = m.ExcPending
															if v497 != 0 {
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
												v256 = m.ExcPending
												if v256 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										} else {
											v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v251 != 0 {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v256 = m.ExcPending
												if v256 != 0 {
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
													v485 = m.ExcPending
													if v485 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v488 = m.ExcPending
														if v488 != 0 {
															return
														} else {
															F_errmsg(m, int32(264620), int32(0))
															mBase = m.M
															v492 = m.ExcPending
															if v492 != 0 {
																return
															} else {
																F_errfinish(m, int32(518532), int32(1407), int32(172761))
																mBase = m.M
																v497 = m.ExcPending
																if v497 != 0 {
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
													v256 = m.ExcPending
													if v256 != 0 {
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
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return
											} else {
												F_errmsg(m, int32(265783), int32(0))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(454496)
													F_errhint(m, int32(632419), v10+int32(16))
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return
													} else {
														F_errfinish(m, int32(518532), int32(1241), int32(172761))
														mBase = m.M
														v82 = m.ExcPending
														if v82 != 0 {
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
							if l2 == v27 {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
								if v51 != int32(2278) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											F_errmsg(m, int32(265783), int32(0))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(454496)
												F_errhint(m, int32(632419), v10+int32(16))
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													F_errfinish(m, int32(518532), int32(1241), int32(172761))
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
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
									v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
									if v54 != int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return
											} else {
												F_errmsg(m, int32(265783), int32(0))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(454496)
													F_errhint(m, int32(632419), v10+int32(16))
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return
													} else {
														F_errfinish(m, int32(518532), int32(1241), int32(172761))
														mBase = m.M
														v82 = m.ExcPending
														if v82 != 0 {
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
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										if v57 == int32(2281) {
											v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v244 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
												v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v248 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												} else {
												}
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v485 = m.ExcPending
													if v485 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v488 = m.ExcPending
														if v488 != 0 {
															return
														} else {
															F_errmsg(m, int32(264620), int32(0))
															mBase = m.M
															v492 = m.ExcPending
															if v492 != 0 {
																return
															} else {
																F_errfinish(m, int32(518532), int32(1407), int32(172761))
																mBase = m.M
																v497 = m.ExcPending
																if v497 != 0 {
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
													v256 = m.ExcPending
													if v256 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											} else {
												v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v251 != 0 {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
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
														v485 = m.ExcPending
														if v485 != 0 {
															return
														} else {
															F_errcode(m, int32(117833860))
															mBase = m.M
															v488 = m.ExcPending
															if v488 != 0 {
																return
															} else {
																F_errmsg(m, int32(264620), int32(0))
																mBase = m.M
																v492 = m.ExcPending
																if v492 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(518532), int32(1407), int32(172761))
																	mBase = m.M
																	v497 = m.ExcPending
																	if v497 != 0 {
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
														v256 = m.ExcPending
														if v256 != 0 {
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
											v63 = m.ExcPending
											if v63 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return
												} else {
													F_errmsg(m, int32(265783), int32(0))
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(454496)
														F_errhint(m, int32(632419), v10+int32(16))
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return
														} else {
															F_errfinish(m, int32(518532), int32(1241), int32(172761))
															mBase = m.M
															v82 = m.ExcPending
															if v82 != 0 {
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
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										F_errmsg(m, int32(385814), int32(0))
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											F_errfinish(m, int32(518532), int32(1224), int32(172761))
											mBase = m.M
											v47 = m.ExcPending
											if v47 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								F_errmsg(m, int32(385814), int32(0))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									F_errfinish(m, int32(518532), int32(1224), int32(172761))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
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
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v21 != v48 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v277 = m.ExcPending
						if v277 != 0 {
							return
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v280 = m.ExcPending
							if v280 != 0 {
								return
							} else {
								F_errmsg(m, int32(341255), int32(0))
								mBase = m.M
								v284 = m.ExcPending
								if v284 != 0 {
									return
								} else {
									F_errfinish(m, int32(518532), int32(1231), int32(172761))
									mBase = m.M
									v289 = m.ExcPending
									if v289 != 0 {
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
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
						if v51 != int32(2278) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									F_errmsg(m, int32(265783), int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(454496)
										F_errhint(m, int32(632419), v10+int32(16))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											F_errfinish(m, int32(518532), int32(1241), int32(172761))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
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
							v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v54 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_errmsg(m, int32(265783), int32(0))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(454496)
											F_errhint(m, int32(632419), v10+int32(16))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												F_errfinish(m, int32(518532), int32(1241), int32(172761))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
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
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
								if v57 == int32(2281) {
									v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v244 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
										v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v248 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										} else {
										}
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v485 = m.ExcPending
											if v485 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v488 = m.ExcPending
												if v488 != 0 {
													return
												} else {
													F_errmsg(m, int32(264620), int32(0))
													mBase = m.M
													v492 = m.ExcPending
													if v492 != 0 {
														return
													} else {
														F_errfinish(m, int32(518532), int32(1407), int32(172761))
														mBase = m.M
														v497 = m.ExcPending
														if v497 != 0 {
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
											v256 = m.ExcPending
											if v256 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									} else {
										v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v251 != 0 {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v256 = m.ExcPending
											if v256 != 0 {
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
												v485 = m.ExcPending
												if v485 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v488 = m.ExcPending
													if v488 != 0 {
														return
													} else {
														F_errmsg(m, int32(264620), int32(0))
														mBase = m.M
														v492 = m.ExcPending
														if v492 != 0 {
															return
														} else {
															F_errfinish(m, int32(518532), int32(1407), int32(172761))
															mBase = m.M
															v497 = m.ExcPending
															if v497 != 0 {
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
												v256 = m.ExcPending
												if v256 != 0 {
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
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											F_errmsg(m, int32(265783), int32(0))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(454496)
												F_errhint(m, int32(632419), v10+int32(16))
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													F_errfinish(m, int32(518532), int32(1241), int32(172761))
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
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
				v84 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+10)))
					if v86 == int32(1) {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						switch v89 - int32(1) {
						case 0:
							v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v92 != int32(2) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v293 = m.ExcPending
								if v293 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v296 = m.ExcPending
									if v296 != 0 {
										return
									} else {
										F_errmsg(m, int32(129341), int32(0))
										mBase = m.M
										v300 = m.ExcPending
										if v300 != 0 {
											return
										} else {
											F_errfinish(m, int32(518532), int32(1259), int32(172761))
											mBase = m.M
											v305 = m.ExcPending
											if v305 != 0 {
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
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
								if v95 != int32(23) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v309 = m.ExcPending
									if v309 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v312 = m.ExcPending
										if v312 != 0 {
											return
										} else {
											F_errmsg(m, int32(236051), int32(0))
											mBase = m.M
											v316 = m.ExcPending
											if v316 != 0 {
												return
											} else {
												F_errfinish(m, int32(518532), int32(1263), int32(172761))
												mBase = m.M
												v321 = m.ExcPending
												if v321 != 0 {
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
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v98 == int32(0) {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v101
									} else {
									}
									v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v103 != 0 {
									} else {
										v104 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v104
									}
									v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v244 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
										v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v248 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										} else {
										}
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v485 = m.ExcPending
											if v485 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v488 = m.ExcPending
												if v488 != 0 {
													return
												} else {
													F_errmsg(m, int32(264620), int32(0))
													mBase = m.M
													v492 = m.ExcPending
													if v492 != 0 {
														return
													} else {
														F_errfinish(m, int32(518532), int32(1407), int32(172761))
														mBase = m.M
														v497 = m.ExcPending
														if v497 != 0 {
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
											v256 = m.ExcPending
											if v256 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									} else {
										v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v251 != 0 {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v256 = m.ExcPending
											if v256 != 0 {
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
												v485 = m.ExcPending
												if v485 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v488 = m.ExcPending
													if v488 != 0 {
														return
													} else {
														F_errmsg(m, int32(264620), int32(0))
														mBase = m.M
														v492 = m.ExcPending
														if v492 != 0 {
															return
														} else {
															F_errfinish(m, int32(518532), int32(1407), int32(172761))
															mBase = m.M
															v497 = m.ExcPending
															if v497 != 0 {
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
												v256 = m.ExcPending
												if v256 != 0 {
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
							v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v106 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v325 = m.ExcPending
								if v325 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v328 = m.ExcPending
									if v328 != 0 {
										return
									} else {
										F_errmsg(m, int32(762855), int32(0))
										mBase = m.M
										v332 = m.ExcPending
										if v332 != 0 {
											return
										} else {
											F_errfinish(m, int32(518532), int32(1280), int32(172761))
											mBase = m.M
											v337 = m.ExcPending
											if v337 != 0 {
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
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
								if v109 != int32(2281) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v325 = m.ExcPending
									if v325 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v328 = m.ExcPending
										if v328 != 0 {
											return
										} else {
											F_errmsg(m, int32(762855), int32(0))
											mBase = m.M
											v332 = m.ExcPending
											if v332 != 0 {
												return
											} else {
												F_errfinish(m, int32(518532), int32(1280), int32(172761))
												mBase = m.M
												v337 = m.ExcPending
												if v337 != 0 {
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
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
									if v112 == int32(2278) {
										v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v244 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
											v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v248 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											} else {
											}
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v485 = m.ExcPending
												if v485 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v488 = m.ExcPending
													if v488 != 0 {
														return
													} else {
														F_errmsg(m, int32(264620), int32(0))
														mBase = m.M
														v492 = m.ExcPending
														if v492 != 0 {
															return
														} else {
															F_errfinish(m, int32(518532), int32(1407), int32(172761))
															mBase = m.M
															v497 = m.ExcPending
															if v497 != 0 {
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
												v256 = m.ExcPending
												if v256 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										} else {
											v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v251 != 0 {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v256 = m.ExcPending
												if v256 != 0 {
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
													v485 = m.ExcPending
													if v485 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v488 = m.ExcPending
														if v488 != 0 {
															return
														} else {
															F_errmsg(m, int32(264620), int32(0))
															mBase = m.M
															v492 = m.ExcPending
															if v492 != 0 {
																return
															} else {
																F_errfinish(m, int32(518532), int32(1407), int32(172761))
																mBase = m.M
																v497 = m.ExcPending
																if v497 != 0 {
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
													v256 = m.ExcPending
													if v256 != 0 {
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
										v118 = m.ExcPending
										if v118 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return
											} else {
												F_errmsg(m, int32(454343), int32(0))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return
												} else {
													F_errfinish(m, int32(518532), int32(1284), int32(172761))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
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
							v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v131 != int32(5) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v341 = m.ExcPending
								if v341 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v344 = m.ExcPending
									if v344 != 0 {
										return
									} else {
										F_errmsg(m, int32(129626), int32(0))
										mBase = m.M
										v348 = m.ExcPending
										if v348 != 0 {
											return
										} else {
											F_errfinish(m, int32(518532), int32(1295), int32(172761))
											mBase = m.M
											v353 = m.ExcPending
											if v353 != 0 {
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
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
								if v134 != int32(16) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v357 = m.ExcPending
									if v357 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v360 = m.ExcPending
										if v360 != 0 {
											return
										} else {
											F_errmsg(m, int32(297591), int32(0))
											mBase = m.M
											v364 = m.ExcPending
											if v364 != 0 {
												return
											} else {
												F_errfinish(m, int32(518532), int32(1299), int32(172761))
												mBase = m.M
												v369 = m.ExcPending
												if v369 != 0 {
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
									v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v137 == int32(0) {
										v140 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v140
									} else {
									}
									v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v142 != 0 {
									} else {
										v143 = *(*int32)(unsafe.Add(mBase, uint32(v18)+144))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v143
									}
									v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v244 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
										v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v248 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										} else {
										}
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v485 = m.ExcPending
											if v485 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v488 = m.ExcPending
												if v488 != 0 {
													return
												} else {
													F_errmsg(m, int32(264620), int32(0))
													mBase = m.M
													v492 = m.ExcPending
													if v492 != 0 {
														return
													} else {
														F_errfinish(m, int32(518532), int32(1407), int32(172761))
														mBase = m.M
														v497 = m.ExcPending
														if v497 != 0 {
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
											v256 = m.ExcPending
											if v256 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									} else {
										v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v251 != 0 {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v256 = m.ExcPending
											if v256 != 0 {
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
												v485 = m.ExcPending
												if v485 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v488 = m.ExcPending
													if v488 != 0 {
														return
													} else {
														F_errmsg(m, int32(264620), int32(0))
														mBase = m.M
														v492 = m.ExcPending
														if v492 != 0 {
															return
														} else {
															F_errfinish(m, int32(518532), int32(1407), int32(172761))
															mBase = m.M
															v497 = m.ExcPending
															if v497 != 0 {
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
												v256 = m.ExcPending
												if v256 != 0 {
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
							v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v145 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v373 = m.ExcPending
								if v373 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v376 = m.ExcPending
									if v376 != 0 {
										return
									} else {
										F_errmsg(m, int32(100660), int32(0))
										mBase = m.M
										v380 = m.ExcPending
										if v380 != 0 {
											return
										} else {
											F_errfinish(m, int32(518532), int32(1315), int32(172761))
											mBase = m.M
											v385 = m.ExcPending
											if v385 != 0 {
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
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
								if v148 != int32(16) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v389 = m.ExcPending
									if v389 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v392 = m.ExcPending
										if v392 != 0 {
											return
										} else {
											F_errmsg(m, int32(297639), int32(0))
											mBase = m.M
											v396 = m.ExcPending
											if v396 != 0 {
												return
											} else {
												F_errfinish(m, int32(518532), int32(1319), int32(172761))
												mBase = m.M
												v401 = m.ExcPending
												if v401 != 0 {
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
									v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v151 == v152 {
										v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v244 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
											v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v248 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											} else {
											}
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v485 = m.ExcPending
												if v485 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v488 = m.ExcPending
													if v488 != 0 {
														return
													} else {
														F_errmsg(m, int32(264620), int32(0))
														mBase = m.M
														v492 = m.ExcPending
														if v492 != 0 {
															return
														} else {
															F_errfinish(m, int32(518532), int32(1407), int32(172761))
															mBase = m.M
															v497 = m.ExcPending
															if v497 != 0 {
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
												v256 = m.ExcPending
												if v256 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										} else {
											v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v251 != 0 {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v256 = m.ExcPending
												if v256 != 0 {
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
													v485 = m.ExcPending
													if v485 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v488 = m.ExcPending
														if v488 != 0 {
															return
														} else {
															F_errmsg(m, int32(264620), int32(0))
															mBase = m.M
															v492 = m.ExcPending
															if v492 != 0 {
																return
															} else {
																F_errfinish(m, int32(518532), int32(1407), int32(172761))
																mBase = m.M
																v497 = m.ExcPending
																if v497 != 0 {
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
													v256 = m.ExcPending
													if v256 != 0 {
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
										v157 = m.ExcPending
										if v157 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return
											} else {
												F_errmsg(m, int32(385095), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return
												} else {
													F_errfinish(m, int32(518532), int32(1332), int32(172761))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
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
							v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v244 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
								v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v248 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
								} else {
								}
								if l2 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v485 = m.ExcPending
									if v485 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v488 = m.ExcPending
										if v488 != 0 {
											return
										} else {
											F_errmsg(m, int32(264620), int32(0))
											mBase = m.M
											v492 = m.ExcPending
											if v492 != 0 {
												return
											} else {
												F_errfinish(m, int32(518532), int32(1407), int32(172761))
												mBase = m.M
												v497 = m.ExcPending
												if v497 != 0 {
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
									v256 = m.ExcPending
									if v256 != 0 {
										return
									} else {
										m.G0 = v10 + int32(32)
										return
									}
								}
							} else {
								v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v251 != 0 {
									F_ReleaseCatCache(m, v14)
									mBase = m.M
									v256 = m.ExcPending
									if v256 != 0 {
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
										v485 = m.ExcPending
										if v485 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v488 = m.ExcPending
											if v488 != 0 {
												return
											} else {
												F_errmsg(m, int32(264620), int32(0))
												mBase = m.M
												v492 = m.ExcPending
												if v492 != 0 {
													return
												} else {
													F_errfinish(m, int32(518532), int32(1407), int32(172761))
													mBase = m.M
													v497 = m.ExcPending
													if v497 != 0 {
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
										v256 = m.ExcPending
										if v256 != 0 {
											return
										} else {
											m.G0 = v10 + int32(32)
											return
										}
									}
								}
							}
						case 5:
							v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v170 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v405 = m.ExcPending
								if v405 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v408 = m.ExcPending
									if v408 != 0 {
										return
									} else {
										F_errmsg(m, int32(762915), int32(0))
										mBase = m.M
										v412 = m.ExcPending
										if v412 != 0 {
											return
										} else {
											F_errfinish(m, int32(518532), int32(1340), int32(172761))
											mBase = m.M
											v417 = m.ExcPending
											if v417 != 0 {
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
								v173 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
								if v173 != int32(2281) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v405 = m.ExcPending
									if v405 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v408 = m.ExcPending
										if v408 != 0 {
											return
										} else {
											F_errmsg(m, int32(762915), int32(0))
											mBase = m.M
											v412 = m.ExcPending
											if v412 != 0 {
												return
											} else {
												F_errfinish(m, int32(518532), int32(1340), int32(172761))
												mBase = m.M
												v417 = m.ExcPending
												if v417 != 0 {
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
									v176 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
									if v176 != int32(2278) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v421 = m.ExcPending
										if v421 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v424 = m.ExcPending
											if v424 != 0 {
												return
											} else {
												F_errmsg(m, int32(454392), int32(0))
												mBase = m.M
												v428 = m.ExcPending
												if v428 != 0 {
													return
												} else {
													F_errfinish(m, int32(518532), int32(1344), int32(172761))
													mBase = m.M
													v433 = m.ExcPending
													if v433 != 0 {
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
										v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v179 == v180 {
											v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v244 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
												v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v248 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												} else {
												}
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v485 = m.ExcPending
													if v485 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v488 = m.ExcPending
														if v488 != 0 {
															return
														} else {
															F_errmsg(m, int32(264620), int32(0))
															mBase = m.M
															v492 = m.ExcPending
															if v492 != 0 {
																return
															} else {
																F_errfinish(m, int32(518532), int32(1407), int32(172761))
																mBase = m.M
																v497 = m.ExcPending
																if v497 != 0 {
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
													v256 = m.ExcPending
													if v256 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											} else {
												v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v251 != 0 {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
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
														v485 = m.ExcPending
														if v485 != 0 {
															return
														} else {
															F_errcode(m, int32(117833860))
															mBase = m.M
															v488 = m.ExcPending
															if v488 != 0 {
																return
															} else {
																F_errmsg(m, int32(264620), int32(0))
																mBase = m.M
																v492 = m.ExcPending
																if v492 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(518532), int32(1407), int32(172761))
																	mBase = m.M
																	v497 = m.ExcPending
																	if v497 != 0 {
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
														v256 = m.ExcPending
														if v256 != 0 {
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
											v185 = m.ExcPending
											if v185 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v188 = m.ExcPending
												if v188 != 0 {
													return
												} else {
													F_errmsg(m, int32(385043), int32(0))
													mBase = m.M
													v192 = m.ExcPending
													if v192 != 0 {
														return
													} else {
														F_errfinish(m, int32(518532), int32(1357), int32(172761))
														mBase = m.M
														v197 = m.ExcPending
														if v197 != 0 {
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
						v199 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
						mBase = m.M
						v200 = m.ExcPending
						if v200 != 0 {
							return
						} else {
							v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+12)))
							if v201 != int32(1) {
								v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v244 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
									v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v248 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
									} else {
									}
									if l2 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v485 = m.ExcPending
										if v485 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v488 = m.ExcPending
											if v488 != 0 {
												return
											} else {
												F_errmsg(m, int32(264620), int32(0))
												mBase = m.M
												v492 = m.ExcPending
												if v492 != 0 {
													return
												} else {
													F_errfinish(m, int32(518532), int32(1407), int32(172761))
													mBase = m.M
													v497 = m.ExcPending
													if v497 != 0 {
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
										v256 = m.ExcPending
										if v256 != 0 {
											return
										} else {
											m.G0 = v10 + int32(32)
											return
										}
									}
								} else {
									v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v251 != 0 {
										F_ReleaseCatCache(m, v14)
										mBase = m.M
										v256 = m.ExcPending
										if v256 != 0 {
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
											v485 = m.ExcPending
											if v485 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v488 = m.ExcPending
												if v488 != 0 {
													return
												} else {
													F_errmsg(m, int32(264620), int32(0))
													mBase = m.M
													v492 = m.ExcPending
													if v492 != 0 {
														return
													} else {
														F_errfinish(m, int32(518532), int32(1407), int32(172761))
														mBase = m.M
														v497 = m.ExcPending
														if v497 != 0 {
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
											v256 = m.ExcPending
											if v256 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								switch v204 - int32(1) {
								case 0:
									v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
									if v207 != int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v437 = m.ExcPending
										if v437 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v440 = m.ExcPending
											if v440 != 0 {
												return
											} else {
												F_errmsg(m, int32(100714), int32(0))
												mBase = m.M
												v444 = m.ExcPending
												if v444 != 0 {
													return
												} else {
													F_errfinish(m, int32(518532), int32(1367), int32(172761))
													mBase = m.M
													v449 = m.ExcPending
													if v449 != 0 {
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
										v210 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
										if v210 == int32(23) {
											v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v235 == int32(0) {
												v238 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v238
											} else {
											}
											v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v240 != 0 {
											} else {
												v241 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v241
											}
											v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v244 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
												v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v248 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												} else {
												}
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v485 = m.ExcPending
													if v485 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v488 = m.ExcPending
														if v488 != 0 {
															return
														} else {
															F_errmsg(m, int32(264620), int32(0))
															mBase = m.M
															v492 = m.ExcPending
															if v492 != 0 {
																return
															} else {
																F_errfinish(m, int32(518532), int32(1407), int32(172761))
																mBase = m.M
																v497 = m.ExcPending
																if v497 != 0 {
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
													v256 = m.ExcPending
													if v256 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											} else {
												v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v251 != 0 {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
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
														v485 = m.ExcPending
														if v485 != 0 {
															return
														} else {
															F_errcode(m, int32(117833860))
															mBase = m.M
															v488 = m.ExcPending
															if v488 != 0 {
																return
															} else {
																F_errmsg(m, int32(264620), int32(0))
																mBase = m.M
																v492 = m.ExcPending
																if v492 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(518532), int32(1407), int32(172761))
																	mBase = m.M
																	v497 = m.ExcPending
																	if v497 != 0 {
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
														v256 = m.ExcPending
														if v256 != 0 {
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
											v216 = m.ExcPending
											if v216 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v219 = m.ExcPending
												if v219 != 0 {
													return
												} else {
													F_errmsg(m, int32(236101), int32(0))
													mBase = m.M
													v223 = m.ExcPending
													if v223 != 0 {
														return
													} else {
														F_errfinish(m, int32(518532), int32(1371), int32(172761))
														mBase = m.M
														v228 = m.ExcPending
														if v228 != 0 {
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
									v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
									if v229 != int32(2) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v453 = m.ExcPending
										if v453 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v456 = m.ExcPending
											if v456 != 0 {
												return
											} else {
												F_errmsg(m, int32(129395), int32(0))
												mBase = m.M
												v460 = m.ExcPending
												if v460 != 0 {
													return
												} else {
													F_errfinish(m, int32(518532), int32(1378), int32(172761))
													mBase = m.M
													v465 = m.ExcPending
													if v465 != 0 {
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
										v232 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
										if v232 != int32(20) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v469 = m.ExcPending
											if v469 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v472 = m.ExcPending
												if v472 != 0 {
													return
												} else {
													F_errmsg(m, int32(95709), int32(0))
													mBase = m.M
													v476 = m.ExcPending
													if v476 != 0 {
														return
													} else {
														F_errfinish(m, int32(518532), int32(1382), int32(172761))
														mBase = m.M
														v481 = m.ExcPending
														if v481 != 0 {
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
											v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v235 == int32(0) {
												v238 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v238
											} else {
											}
											v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v240 != 0 {
											} else {
												v241 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v241
											}
											v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v244 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
												v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v248 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												} else {
												}
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v485 = m.ExcPending
													if v485 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v488 = m.ExcPending
														if v488 != 0 {
															return
														} else {
															F_errmsg(m, int32(264620), int32(0))
															mBase = m.M
															v492 = m.ExcPending
															if v492 != 0 {
																return
															} else {
																F_errfinish(m, int32(518532), int32(1407), int32(172761))
																mBase = m.M
																v497 = m.ExcPending
																if v497 != 0 {
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
													v256 = m.ExcPending
													if v256 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											} else {
												v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v251 != 0 {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
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
														v485 = m.ExcPending
														if v485 != 0 {
															return
														} else {
															F_errcode(m, int32(117833860))
															mBase = m.M
															v488 = m.ExcPending
															if v488 != 0 {
																return
															} else {
																F_errmsg(m, int32(264620), int32(0))
																mBase = m.M
																v492 = m.ExcPending
																if v492 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(518532), int32(1407), int32(172761))
																	mBase = m.M
																	v497 = m.ExcPending
																	if v497 != 0 {
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
														v256 = m.ExcPending
														if v256 != 0 {
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
									v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v235 == int32(0) {
										v238 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v238
									} else {
									}
									v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v240 != 0 {
									} else {
										v241 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v241
									}
									v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v244 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
										v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v248 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										} else {
										}
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v485 = m.ExcPending
											if v485 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v488 = m.ExcPending
												if v488 != 0 {
													return
												} else {
													F_errmsg(m, int32(264620), int32(0))
													mBase = m.M
													v492 = m.ExcPending
													if v492 != 0 {
														return
													} else {
														F_errfinish(m, int32(518532), int32(1407), int32(172761))
														mBase = m.M
														v497 = m.ExcPending
														if v497 != 0 {
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
											v256 = m.ExcPending
											if v256 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									} else {
										v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v251 != 0 {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v256 = m.ExcPending
											if v256 != 0 {
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
												v485 = m.ExcPending
												if v485 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v488 = m.ExcPending
													if v488 != 0 {
														return
													} else {
														F_errmsg(m, int32(264620), int32(0))
														mBase = m.M
														v492 = m.ExcPending
														if v492 != 0 {
															return
														} else {
															F_errfinish(m, int32(518532), int32(1407), int32(172761))
															mBase = m.M
															v497 = m.ExcPending
															if v497 != 0 {
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
												v256 = m.ExcPending
												if v256 != 0 {
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
			v263 = m.ExcPending
			if v263 != 0 {
				return
			} else {
				v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v264
				F_errmsg_internal(m, int32(48455), v10)
				mBase = m.M
				v268 = m.ExcPending
				if v268 != 0 {
					return
				} else {
					F_errfinish(m, int32(518532), int32(1212), int32(172761))
					mBase = m.M
					v273 = m.ExcPending
					if v273 != 0 {
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
