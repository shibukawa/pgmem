package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_SICleanupQueue(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
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
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[0]))
	if l0 == v3 {
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
		v28 = F_LWLockAcquire(m, v24+int32(768), int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
			v35 = F_LWLockAcquire(m, v31+int32(640), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[2])))
				if v38 <= int32(0) {
					v95 = v38
					v97 = v37
					v102 = v3
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[3])))
					v51 = int32(0)
					v53 = v38
					v55 = v37
					v59 = v37 - int32(2048)
					v60 = v3
					for {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v48+v51<<(uint(int32(2))%32))))
						v70 = v20 + int32(_a_F_SICleanupQueue_0) + v67<<(uint(int32(4))%32)
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)))
						if v71 != 0 {
							v84 = v53
							v85 = v55
							v87 = v59
							v88 = v60
						} else {
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+11)))
							if v72 != 0 {
								v84 = v53
								v85 = v55
								v87 = v59
								v88 = v60
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
								if v73 < l1+v37-int32(_a_F_SICleanupQueue_1) {
									v75 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)) = uint8(v75)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[2])))
									v84 = v77
									v85 = v55
									v87 = v59
									v88 = v60
								} else {
									if v73 < v55 {
										v79 = v73
									} else {
										v79 = v55
									}
									if v59 <= v73 {
										v84 = v53
										v85 = v79
										v87 = v59
										v88 = v60
									} else {
										v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+9)))
										if v81 != 0 {
											v82 = v59
										} else {
											v82 = v73
										}
										if v81 != 0 {
											v83 = v60
										} else {
											v83 = v70
										}
										v84 = v53
										v85 = v79
										v87 = v82
										v88 = v83
									}
								}
							}
						}
						v90 = v51 + int32(1)
						if v90 < v84 {
							v51 = v90
							v53 = v84
							v55 = v85
							v59 = v87
							v60 = v88
							continue
						} else {
							break
						}
						break
					}
					v95 = v84
					v97 = v85
					v102 = v88
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v97
				if v97 < int32(1073741824) {
					v242 = v97
					v244 = v37
				} else {
					v109 = int32(1073741824)
					v110 = v37 - v109
					*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v110
					v113 = v97 - v109
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v113
					if v95 <= int32(0) {
						v242 = v113
						v244 = v110
					} else {
						v118 = v95 & int32(3)
						v119 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[3])))
						v120 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v95) {
							v128 = v120
							v130 = int32(0)
							for {
								v143 = v119 + v128<<(uint(int32(2))%32)
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
								v145 = int32(4)
								v147 = v20 + v144<<(uint(v145)%32)
								v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)+uint32(_c_F_SICleanupQueue[4])))
								v151 = int32(1073741824)
								*(*int32)(unsafe.Add(mBase, uint32(v147)+uint32(_c_F_SICleanupQueue[4]))) = v150 - v151
								v154 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
								v157 = v20 + v154<<(uint(v145)%32)
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v157)+uint32(_c_F_SICleanupQueue[4])))
								*(*int32)(unsafe.Add(mBase, uint32(v157)+uint32(_c_F_SICleanupQueue[4]))) = v160 - v151
								v164 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
								v167 = v20 + v164<<(uint(v145)%32)
								v170 = *(*int32)(unsafe.Add(mBase, uint32(v167)+uint32(_c_F_SICleanupQueue[4])))
								*(*int32)(unsafe.Add(mBase, uint32(v167)+uint32(_c_F_SICleanupQueue[4]))) = v170 - v151
								v174 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
								v177 = v20 + v174<<(uint(v145)%32)
								v180 = *(*int32)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_SICleanupQueue[4])))
								*(*int32)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_SICleanupQueue[4]))) = v180 - v151
								v185 = v128 + v145
								v187 = v130 + v145
								if v187 != v95&int32(2147483644) {
									v128 = v185
									v130 = v187
									continue
								} else {
									break
								}
								break
							}
							if v118 == int32(0) {
								v242 = v113
								v244 = v110
							} else {
								v192 = v185
								v206 = v192
								v216 = v120
								for {
									v222 = *(*int32)(unsafe.Add(mBase, uint32(v119+v206<<(uint(int32(2))%32))))
									v225 = v20 + v222<<(uint(int32(4))%32)
									v228 = *(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_SICleanupQueue[4])))
									*(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_SICleanupQueue[4]))) = v228 - int32(1073741824)
									v232 = int32(1)
									v235 = v216 + v232
									if v235 != v118 {
										v206 = v206 + v232
										v216 = v235
										continue
									} else {
										break
									}
									break
								}
								v242 = v113
								v244 = v110
							}
						} else {
							v192 = v120
							v206 = v192
							v216 = v120
							for {
								v222 = *(*int32)(unsafe.Add(mBase, uint32(v119+v206<<(uint(int32(2))%32))))
								v225 = v20 + v222<<(uint(int32(4))%32)
								v228 = *(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_SICleanupQueue[4])))
								*(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_SICleanupQueue[4]))) = v228 - int32(1073741824)
								v232 = int32(1)
								v235 = v216 + v232
								if v235 != v118 {
									v206 = v206 + v232
									v216 = v235
									continue
								} else {
									break
								}
								break
							}
							v242 = v113
							v244 = v110
						}
					}
				}
				v251 = int32(2048)
				v252 = v244 - v242
				if v252 < v251 {
					v259 = v251
				} else {
					v259 = v252&int32(2147483392) + int32(256)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v259
				if v102 != 0 {
					v261 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v102)+9)) = uint8(v261)
					v263 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
					v265 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
					F_LWLockRelease(m, v265+int32(640))
					mBase = m.M
					v269 = m.ExcPending
					if v269 != 0 {
						return
					} else {
						v271 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
						F_LWLockRelease(m, v271+int32(768))
						mBase = m.M
						v275 = m.ExcPending
						if v275 != 0 {
							return
						} else {
							v278 = F_errstart(m, int32(11), int32(0))
							mBase = m.M
							v279 = m.ExcPending
							if v279 != 0 {
								return
							} else {
								if v278 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v17))) = v263
									F_errmsg_internal(m, int32(_a_F_SICleanupQueue_2), v17)
									mBase = m.M
									v283 = m.ExcPending
									if v283 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SICleanupQueue_3), int32(675), int32(_a_F_SICleanupQueue_4))
										mBase = m.M
										v288 = m.ExcPending
										if v288 != 0 {
											return
										} else {
											v295 = F_SendProcSignal(m, v263, int32(0), (v102-v20-int32(_a_F_SICleanupQueue_0))>>(uint(int32(4))%32))
											mBase = m.M
											v296 = m.ExcPending
											if v296 != 0 {
												return
											} else {
												if l0 == int32(0) {
													m.G0 = v17 + int32(16)
													return
												} else {
													v300 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
													v304 = F_LWLockAcquire(m, v300+int32(768), int32(0))
													mBase = m.M
													v305 = m.ExcPending
													if v305 != 0 {
														return
													} else {
														m.G0 = v17 + int32(16)
														return
													}
												}
											}
										}
									}
								} else {
									v295 = F_SendProcSignal(m, v263, int32(0), (v102-v20-int32(_a_F_SICleanupQueue_0))>>(uint(int32(4))%32))
									mBase = m.M
									v296 = m.ExcPending
									if v296 != 0 {
										return
									} else {
										if l0 == int32(0) {
											m.G0 = v17 + int32(16)
											return
										} else {
											v300 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
											v304 = F_LWLockAcquire(m, v300+int32(768), int32(0))
											mBase = m.M
											v305 = m.ExcPending
											if v305 != 0 {
												return
											} else {
												m.G0 = v17 + int32(16)
												return
											}
										}
									}
								}
							}
						}
					}
				} else {
					v307 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
					F_LWLockRelease(m, v307+int32(640))
					mBase = m.M
					v311 = m.ExcPending
					if v311 != 0 {
						return
					} else {
						if l0 != 0 {
							m.G0 = v17 + int32(16)
							return
						} else {
							v313 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
							F_LWLockRelease(m, v313+int32(768))
							mBase = m.M
							v317 = m.ExcPending
							if v317 != 0 {
								return
							} else {
								m.G0 = v17 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
		v35 = F_LWLockAcquire(m, v31+int32(640), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[2])))
			if v38 <= int32(0) {
				v95 = v38
				v97 = v37
				v102 = v3
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[3])))
				v51 = int32(0)
				v53 = v38
				v55 = v37
				v59 = v37 - int32(2048)
				v60 = v3
				for {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v48+v51<<(uint(int32(2))%32))))
					v70 = v20 + int32(_a_F_SICleanupQueue_0) + v67<<(uint(int32(4))%32)
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)))
					if v71 != 0 {
						v84 = v53
						v85 = v55
						v87 = v59
						v88 = v60
					} else {
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+11)))
						if v72 != 0 {
							v84 = v53
							v85 = v55
							v87 = v59
							v88 = v60
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
							if v73 < l1+v37-int32(_a_F_SICleanupQueue_1) {
								v75 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)) = uint8(v75)
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[2])))
								v84 = v77
								v85 = v55
								v87 = v59
								v88 = v60
							} else {
								if v73 < v55 {
									v79 = v73
								} else {
									v79 = v55
								}
								if v59 <= v73 {
									v84 = v53
									v85 = v79
									v87 = v59
									v88 = v60
								} else {
									v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+9)))
									if v81 != 0 {
										v82 = v59
									} else {
										v82 = v73
									}
									if v81 != 0 {
										v83 = v60
									} else {
										v83 = v70
									}
									v84 = v53
									v85 = v79
									v87 = v82
									v88 = v83
								}
							}
						}
					}
					v90 = v51 + int32(1)
					if v90 < v84 {
						v51 = v90
						v53 = v84
						v55 = v85
						v59 = v87
						v60 = v88
						continue
					} else {
						break
					}
					break
				}
				v95 = v84
				v97 = v85
				v102 = v88
			}
			*(*int32)(unsafe.Add(mBase, uint32(v20))) = v97
			if v97 < int32(1073741824) {
				v242 = v97
				v244 = v37
			} else {
				v109 = int32(1073741824)
				v110 = v37 - v109
				*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v110
				v113 = v97 - v109
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v113
				if v95 <= int32(0) {
					v242 = v113
					v244 = v110
				} else {
					v118 = v95 & int32(3)
					v119 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[3])))
					v120 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v95) {
						v128 = v120
						v130 = int32(0)
						for {
							v143 = v119 + v128<<(uint(int32(2))%32)
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
							v145 = int32(4)
							v147 = v20 + v144<<(uint(v145)%32)
							v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)+uint32(_c_F_SICleanupQueue[4])))
							v151 = int32(1073741824)
							*(*int32)(unsafe.Add(mBase, uint32(v147)+uint32(_c_F_SICleanupQueue[4]))) = v150 - v151
							v154 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
							v157 = v20 + v154<<(uint(v145)%32)
							v160 = *(*int32)(unsafe.Add(mBase, uint32(v157)+uint32(_c_F_SICleanupQueue[4])))
							*(*int32)(unsafe.Add(mBase, uint32(v157)+uint32(_c_F_SICleanupQueue[4]))) = v160 - v151
							v164 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
							v167 = v20 + v164<<(uint(v145)%32)
							v170 = *(*int32)(unsafe.Add(mBase, uint32(v167)+uint32(_c_F_SICleanupQueue[4])))
							*(*int32)(unsafe.Add(mBase, uint32(v167)+uint32(_c_F_SICleanupQueue[4]))) = v170 - v151
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
							v177 = v20 + v174<<(uint(v145)%32)
							v180 = *(*int32)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_SICleanupQueue[4])))
							*(*int32)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_SICleanupQueue[4]))) = v180 - v151
							v185 = v128 + v145
							v187 = v130 + v145
							if v187 != v95&int32(2147483644) {
								v128 = v185
								v130 = v187
								continue
							} else {
								break
							}
							break
						}
						if v118 == int32(0) {
							v242 = v113
							v244 = v110
						} else {
							v192 = v185
							v206 = v192
							v216 = v120
							for {
								v222 = *(*int32)(unsafe.Add(mBase, uint32(v119+v206<<(uint(int32(2))%32))))
								v225 = v20 + v222<<(uint(int32(4))%32)
								v228 = *(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_SICleanupQueue[4])))
								*(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_SICleanupQueue[4]))) = v228 - int32(1073741824)
								v232 = int32(1)
								v235 = v216 + v232
								if v235 != v118 {
									v206 = v206 + v232
									v216 = v235
									continue
								} else {
									break
								}
								break
							}
							v242 = v113
							v244 = v110
						}
					} else {
						v192 = v120
						v206 = v192
						v216 = v120
						for {
							v222 = *(*int32)(unsafe.Add(mBase, uint32(v119+v206<<(uint(int32(2))%32))))
							v225 = v20 + v222<<(uint(int32(4))%32)
							v228 = *(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_SICleanupQueue[4])))
							*(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_SICleanupQueue[4]))) = v228 - int32(1073741824)
							v232 = int32(1)
							v235 = v216 + v232
							if v235 != v118 {
								v206 = v206 + v232
								v216 = v235
								continue
							} else {
								break
							}
							break
						}
						v242 = v113
						v244 = v110
					}
				}
			}
			v251 = int32(2048)
			v252 = v244 - v242
			if v252 < v251 {
				v259 = v251
			} else {
				v259 = v252&int32(2147483392) + int32(256)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v259
			if v102 != 0 {
				v261 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v102)+9)) = uint8(v261)
				v263 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
				v265 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
				F_LWLockRelease(m, v265+int32(640))
				mBase = m.M
				v269 = m.ExcPending
				if v269 != 0 {
					return
				} else {
					v271 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
					F_LWLockRelease(m, v271+int32(768))
					mBase = m.M
					v275 = m.ExcPending
					if v275 != 0 {
						return
					} else {
						v278 = F_errstart(m, int32(11), int32(0))
						mBase = m.M
						v279 = m.ExcPending
						if v279 != 0 {
							return
						} else {
							if v278 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v17))) = v263
								F_errmsg_internal(m, int32(_a_F_SICleanupQueue_2), v17)
								mBase = m.M
								v283 = m.ExcPending
								if v283 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SICleanupQueue_3), int32(675), int32(_a_F_SICleanupQueue_4))
									mBase = m.M
									v288 = m.ExcPending
									if v288 != 0 {
										return
									} else {
										v295 = F_SendProcSignal(m, v263, int32(0), (v102-v20-int32(_a_F_SICleanupQueue_0))>>(uint(int32(4))%32))
										mBase = m.M
										v296 = m.ExcPending
										if v296 != 0 {
											return
										} else {
											if l0 == int32(0) {
												m.G0 = v17 + int32(16)
												return
											} else {
												v300 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
												v304 = F_LWLockAcquire(m, v300+int32(768), int32(0))
												mBase = m.M
												v305 = m.ExcPending
												if v305 != 0 {
													return
												} else {
													m.G0 = v17 + int32(16)
													return
												}
											}
										}
									}
								}
							} else {
								v295 = F_SendProcSignal(m, v263, int32(0), (v102-v20-int32(_a_F_SICleanupQueue_0))>>(uint(int32(4))%32))
								mBase = m.M
								v296 = m.ExcPending
								if v296 != 0 {
									return
								} else {
									if l0 == int32(0) {
										m.G0 = v17 + int32(16)
										return
									} else {
										v300 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
										v304 = F_LWLockAcquire(m, v300+int32(768), int32(0))
										mBase = m.M
										v305 = m.ExcPending
										if v305 != 0 {
											return
										} else {
											m.G0 = v17 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v307 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
				F_LWLockRelease(m, v307+int32(640))
				mBase = m.M
				v311 = m.ExcPending
				if v311 != 0 {
					return
				} else {
					if l0 != 0 {
						m.G0 = v17 + int32(16)
						return
					} else {
						v313 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
						F_LWLockRelease(m, v313+int32(768))
						mBase = m.M
						v317 = m.ExcPending
						if v317 != 0 {
							return
						} else {
							m.G0 = v17 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_SampleCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 float64
	_ = v36
	var v40 int32
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v54 float64
	_ = v54
	var v75 float64
	_ = v75
	var v79 float64
	_ = v79
	var v81 float64
	_ = v81
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v92 float64
	_ = v92
	var v100 float64
	_ = v100
	var v101 float64
	_ = v101
	var v103 float64
	_ = v103
	var v106 float64
	_ = v106
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v120 float64
	_ = v120
	var v132 float64
	_ = v132
	var v137 float64
	_ = v137
	var v138 float64
	_ = v138
	var v139 float64
	_ = v139
	var v143 float64
	_ = v143
	var v145 float64
	_ = v145
	var v147 float64
	_ = v147
	var v150 float64
	_ = v150
	var v153 float64
	_ = v153
	var v159 float64
	_ = v159
	var v160 int32
	_ = v160
	var v161 float64
	_ = v161
	var v164 float64
	_ = v164
	var v168 float64
	_ = v168
	var v169 float64
	_ = v169
	var v173 float64
	_ = v173
	var v181 float64
	_ = v181
	var v182 float64
	_ = v182
	var v185 float64
	_ = v185
	var v195 float64
	_ = v195
	var v217 float64
	_ = v217
	var v220 float64
	_ = v220
	var v222 float64
	_ = v222
	var v223 float64
	_ = v223
	var v226 float64
	_ = v226
	var v234 float64
	_ = v234
	var v254 float64
	_ = v254
	var v263 float64
	_ = v263
	var v271 float64
	_ = v271
	var v278 int32
	_ = v278
	var v279 float64
	_ = v279
	var v280 float64
	_ = v280
	var v285 float64
	_ = v285
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = int32(_a_F_SampleCallback_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_SampleCallback[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l5)+168))
	*(*int32)(unsafe.Add(mBase, _c_F_SampleCallback[0])) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l5)+64))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	return
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SampleCallback[0])) = v12
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l5)+168))
	F_MemoryContextReset(m, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L63
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l5)+60))
	v24 = F_IvfflatCheckNorm(m, v21, v22, base.I64_extend_i32_u(v19))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v28 < v17 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if v24 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v285 = *(*float64)(unsafe.Add(mBase, uint32(l5)+136))
	*(*float64)(unsafe.Add(mBase, uint32(l5)+136)) = base.F64_add(v285, float64(1))
	goto L6
L13:
	;
	F_VectorArraySet(m, v16, v28, v19)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v36 = *(*float64)(unsafe.Add(mBase, uint32(l5)+144))
	if base.F64_lt(v36, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v32 + int32(1)
	goto L12
L17:
	;
	v40 = l5 + int32(112)
	v41 = *(*float64)(unsafe.Add(mBase, uint32(l5)+136))
	v42 = float64(0)
	v54 = base.F64_convert_i32_s(v17)
	if base.F64_ge(base.F64_mul(v54, float64(22)), v41) != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v263 = v36
	goto L19
L19:
	;
	if base.F64_le(v263, float64(0)) != 0 {
		goto L55
	} else {
		goto L56
	}
L20:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5)+144)) = v254
	v263 = v254
	goto L19
L21:
	;
	goto L20
L22:
	;
	goto L25
L23:
	;
	goto L24
L24:
	;
	v108 = float64(1)
	v109 = base.F64_add(v41, v108)
	v110 = base.F64_sub(v41, v54)
	v112 = base.F64_add(v110, v108)
	v113 = base.F64_div(v109, v112)
	v115 = l5 + int32(120)
	v116 = *(*float64)(unsafe.Add(mBase, uint32(v40)))
	v120 = v116
	goto L32
L25:
	;
	v75 = F_pg_prng_double(m, l5+int32(120))
	mBase = m.M
	if base.F64_eq(v75, float64(0)) != 0 {
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v79 = base.F64_add(v41, float64(1))
	v81 = base.F64_div(base.F64_sub(v79, v54), v79)
	if base.F64_gt(v81, v75) == int32(0) {
		v254 = v42
		goto L21
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v88 = v79
	v89 = v81
	v92 = v42
	goto L29
L29:
	;
	v100 = float64(1)
	v101 = base.F64_add(v92, v100)
	v103 = base.F64_add(v88, v100)
	v106 = base.F64_mul(v89, base.F64_div(base.F64_sub(v103, v54), v103))
	if base.F64_gt(v106, v75) != 0 {
		v88 = v103
		v89 = v106
		v92 = v101
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v254 = v101
	goto L21
L31:
	;
	goto L30
L32:
	;
	v132 = F_pg_prng_double(m, v115)
	mBase = m.M
	if base.F64_eq(v132, float64(0)) != 0 {
		goto L32
	} else {
		goto L34
	}
L33:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v40))) = v234
	v254 = v138
	goto L21
L34:
	;
	v137 = base.F64_mul(v41, base.F64_add(v120, float64(-1)))
	v138 = base.F64_floor(v137)
	v139 = base.F64_add(v112, v138)
	v143 = base.F64_add(v41, v137)
	v145 = F_log(m, base.F64_div(base.F64_mul(v139, base.F64_mul(v113, base.F64_mul(v113, v132))), v143))
	mBase = m.M
	v147 = F_exp(m, base.F64_div(v145, v54))
	mBase = m.M
	v150 = base.F64_div(base.F64_mul(v112, base.F64_div(v143, v139)), v41)
	if base.F64_le(v147, v150) != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v234 = base.F64_div(v150, v147)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v153 = base.F64_add(v41, v138)
	v159 = base.F64_div(base.F64_mul(base.F64_add(v153, float64(1)), base.F64_div(base.F64_mul(v109, v132), v112)), v143)
	v160 = base.F64_lt(v54, v138)
	if v160 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v161 = v139
	goto L41
L40:
	;
	v161 = v109
	goto L41
L41:
	;
	if base.F64_le(v161, v153) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v160 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v195 = v159
	goto L44
L44:
	;
	goto L51
L45:
	;
	v164 = v41
	goto L47
L46:
	;
	v164 = base.F64_add(v110, v138)
	goto L47
L47:
	;
	v168 = v153
	v169 = v164
	v173 = v159
	goto L48
L48:
	;
	v181 = base.F64_mul(v173, base.F64_div(v168, v169))
	v182 = float64(-1)
	v185 = base.F64_add(v168, v182)
	if base.F64_ge(v185, v161) != 0 {
		v168 = v185
		v169 = base.F64_add(v169, v182)
		v173 = v181
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v195 = v181
	goto L44
L50:
	;
	goto L49
L51:
	;
	v217 = F_pg_prng_double(m, v115)
	mBase = m.M
	if base.F64_eq(v217, float64(0)) != 0 {
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v220 = F_log(m, v195)
	mBase = m.M
	v222 = F_exp(m, base.F64_div(v220, v54))
	mBase = m.M
	v223 = F_log(m, v217)
	mBase = m.M
	v226 = F_exp(m, base.F64_div(base.F64_neg(v223), v54))
	mBase = m.M
	if base.F64_le(v222, base.F64_div(v143, v41)) == int32(0) {
		v120 = v226
		goto L32
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v234 = v226
	goto L35
L55:
	;
	goto L59
L56:
	;
	v280 = v263
	goto L57
L57:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5)+144)) = base.F64_add(v280, float64(-1))
	goto L12
L58:
	;
	F_VectorArraySet(m, v16, base.I32_trunc_sat_f64_s(base.F64_mul(v271, base.F64_convert_i32_s(v17))), v19)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L62
	}
L59:
	;
	v271 = F_pg_prng_double(m, l5+int32(120))
	mBase = m.M
	if base.F64_eq(v271, float64(0)) != 0 {
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L58
L61:
	;
	goto L60
L62:
	;
	v279 = *(*float64)(unsafe.Add(mBase, uint32(l5)+144))
	v280 = v279
	goto L57
L63:
	;
	goto L3
}
func F_SetQuitSignalReason(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_SetQuitSignalReason[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+44)) = l0
	return
}
func F_SetSessionAuthorization(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	v2 = l1
	*(*uint8)(unsafe.Add(mBase, _c_F_SetSessionAuthorization[0])) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, _c_F_SetSessionAuthorization[1])) = l0
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetSessionAuthorization[2])))
	if v8 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_SetSessionAuthorization[3])) = l0
		*(*int32)(unsafe.Add(mBase, _c_F_SetSessionAuthorization[4])) = l0
		if v2 != 0 {
			v18 = int32(_a_F_SetSessionAuthorization_0)
		} else {
			v18 = int32(_a_F_SetSessionAuthorization_1)
		}
		F_SetConfigOption(m, int32(_a_F_SetSessionAuthorization_2), v18, int32(0), int32(1))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_SpeculativeInsertionWait(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(74027918874902528)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	v15 = F_LockAcquire(m, v6, int32(5), v3, v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v19 = F_LockRelease(m, v6, int32(5), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_SplitColQualList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v5
	if l0 == v5 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L16
	} else {
		goto L22
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L19
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v44
	m.G0 = v11 + int32(16)
	return
L4:
	;
	v44 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v18 = l0
	v22 = v5
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v26 <= v22 {
		v44 = v18
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v44 = v42
	goto L3
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v22<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v33 != int32(74) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if v33 != int32(161) {
		goto L2
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v40 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	if v18 != 0 {
		v22 = v22 + int32(1)
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v44 = v18
	goto L3
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v32
	v42 = F_list_delete_nth_cell(m, v18, v22)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	if v42 != 0 {
		v18 = v42
		goto L7
	} else {
		goto L18
	}
L18:
	;
	goto L8
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v60
	F_errmsg_internal(m, int32(_a_F_SplitColQualList_0), v11)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_SplitColQualList_1), int32(_a_F_SplitColQualList_2), int32(_a_F_SplitColQualList_3))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(_a_F_SplitColQualList_4), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	F_scanner_errposition(m, v81, l3)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_SplitColQualList_1), int32(_a_F_SplitColQualList_5), int32(_a_F_SplitColQualList_3))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F___sigaction(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	if base.Ui32(int32(65)) <= base.Ui32(l0) {
		*(*int32)(unsafe.Add(mBase, _c_F___sigaction[0])) = int32(28)
		return int32(-1)
	} else {
		if l2 != 0 {
			v13 = l0 * int32(20)
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F___sigaction[1])))
			*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v14
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F___sigaction[2])))
			*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v16
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F___sigaction[3])))
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v18
		} else {
		}
		if l1 != 0 {
			v22 = l0 * int32(20)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F___sigaction[1]))) = v23
			v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F___sigaction[2]))) = v25
			v27 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F___sigaction[3]))) = v27
		} else {
		}
		return int32(0)
	}
}
func F___stdio_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = l1
	v20 = v17 - v15
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v20
	v22 = v20 + l2
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v25 = v13 + int32(16)
	v28 = base.B2i32(v15 == v17)
	if v15 == v17 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v152
L2:
	;
	v129 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v133 | int32(32)
	if v126 == int32(2) {
		v152 = v129
		goto L1
	} else {
		goto L39
	}
L3:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v113
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113 + v116
	v152 = l2
	goto L1
L4:
	;
	if v97 != int32(-1) {
		v120 = v92
		v126 = v98
		goto L2
	} else {
		goto L38
	}
L5:
	;
	v29 = v25 | int32(8)
	goto L7
L6:
	;
	v29 = v25
	goto L7
L7:
	;
	if v15 == v17 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v32 = int32(1)
	goto L10
L9:
	;
	v32 = int32(2)
	goto L10
L10:
	;
	v35 = m.Wasi_snapshot_preview1.Fd_write(m, v23, v29, v32, v13+int32(12))
	mBase = m.M
	if v35 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v42 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v42 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___stdio_write[0])) = v35
	v42 = int32(-1)
	goto L11
L15:
	;
	v92 = v29
	v97 = v22
	v98 = v32
	goto L4
L16:
	;
	goto L17
L17:
	;
	v46 = v29
	v49 = v22
	v50 = v32
	goto L18
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v49 == v53 {
		goto L3
	} else {
		goto L20
	}
L19:
	;
	v92 = v62
	v97 = v76
	v98 = v78
	goto L4
L20:
	;
	if v53 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v120 = v46
	v126 = v50
	goto L2
L22:
	;
	goto L23
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v60 = base.B2i32(base.Ui32(v59) < base.Ui32(v53))
	if base.Ui32(v59) < base.Ui32(v53) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v61 = int32(8)
	goto L26
L25:
	;
	v61 = int32(0)
	goto L26
L26:
	;
	v62 = v46 + v61
	if base.Ui32(v59) < base.Ui32(v53) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v64 = v59
	goto L29
L28:
	;
	v64 = int32(0)
	goto L29
L29:
	;
	v65 = v53 - v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v65 + v66
	if base.Ui32(v59) < base.Ui32(v53) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v71 = int32(12)
	goto L32
L31:
	;
	v71 = int32(4)
	goto L32
L32:
	;
	v72 = v46 + v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v73 - v65
	v76 = v49 - v53
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v78 = v50 - v60
	v81 = m.Wasi_snapshot_preview1.Fd_write(m, v77, v62, v78, v13+int32(12))
	mBase = m.M
	if v81 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v88 == int32(0) {
		v46 = v62
		v49 = v76
		v50 = v78
		goto L18
	} else {
		goto L37
	}
L34:
	;
	v88 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___stdio_write[0])) = v81
	v88 = int32(-1)
	goto L33
L37:
	;
	goto L19
L38:
	;
	goto L3
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v152 = l2 - v140
	goto L1
}
func F___strftime_l(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v157 int64
	_ = v157
	var v161 int64
	_ = v161
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int64
	_ = v229
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int64
	_ = v289
	var v290 int64
	_ = v290
	var v293 int64
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int64
	_ = v299
	var v304 int64
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v314 int64
	_ = v314
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int64
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v353 int64
	_ = v353
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v390 int64
	_ = v390
	var v391 int64
	_ = v391
	var v392 int64
	_ = v392
	var v395 int64
	_ = v395
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v466 int64
	_ = v466
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int64
	_ = v480
	var v481 int64
	_ = v481
	var v482 int64
	_ = v482
	var v501 int64
	_ = v501
	var v503 int64
	_ = v503
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int64
	_ = v586
	var v597 int64
	_ = v597
	var v601 int64
	_ = v601
	var v605 int64
	_ = v605
	var v607 int64
	_ = v607
	var v610 int64
	_ = v610
	var v612 int64
	_ = v612
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v664 int64
	_ = v664
	var v668 int32
	_ = v668
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int64
	_ = v801
	var v805 int64
	_ = v805
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v816 int64
	_ = v816
	var v819 int32
	_ = v819
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v960 int32
	_ = v960
	var v982 int32
	_ = v982
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1088 int32
	_ = v1088
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1278 int32
	_ = v1278
	var v1290 int32
	_ = v1290
	var v1310 int32
	_ = v1310
	var v1320 int32
	_ = v1320
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1347 int32
	_ = v1347
	v6 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(128)
	m.G0 = v28
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = l2
	v38 = v6
	goto L5
L2:
	;
	v1347 = v6
	goto L3
L3:
	;
	m.G0 = v28 + int32(128)
	return v1347
L4:
	;
	v1339 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v1320))) = uint8(v1339)
	v1347 = v1337
	goto L3
L5:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v55 != int32(37) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	if l1 == v1290 {
		goto L347
	} else {
		goto L348
	}
L7:
	;
	goto L6
L8:
	;
	if base.Ui32(v1278) < base.Ui32(l1) {
		v32 = v1255 + int32(1)
		v38 = v1278
		goto L5
	} else {
		goto L346
	}
L9:
	;
	v79 = v74 & int32(255)
	v82 = v32 + v75 + base.B2i32(v79 == int32(43))
	v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(v82))))
	if base.Ui32(v83-int32(48)) <= base.Ui32(int32(9)) {
		goto L20
	} else {
		goto L21
	}
L10:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+2)))
	v74 = v72
	v75 = int32(2)
	v76 = v60
	goto L9
L11:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v38))) = uint8(v55)
	v1255 = v32
	v1278 = v38 + int32(1)
	goto L8
L12:
	;
	if v55 != 0 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v58 = int32(0)
	v59 = int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	switch v60 - int32(45) {
	case 0, 3:
		goto L10
	case 1, 2:
		v74 = v60
		v75 = v59
		v76 = v58
		goto L9
	default:
		goto L16
	}
L15:
	;
	v1320 = v38
	v1337 = v38
	goto L4
L16:
	;
	if v60 == int32(95) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	if v60 != 0 {
		v74 = v60
		v75 = v59
		v76 = v58
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	v99 = int32(0)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v102 = v100 - int32(67)
	if base.B2i32(base.Ui32(int32(22)) < base.Ui32(v102))|base.B2i32(int32(1)<<(uint(v102)%32)&int32(_a_F___strftime_l_0) == v99) != 0 {
		v113 = v99
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v92 = F_strtox_2(m, v82, v28+int32(12), int32(10), int64(4294967295))
	mBase = m.M
	goto L23
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v82
	v97 = int32(0)
	v98 = v82
	goto L19
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v97 = base.I32_wrap_i64(v92)
	v98 = v94
	goto L19
L24:
	;
	if base.B2i32(v100 == int32(79))|base.B2i32(v100 == int32(69)) != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if v97 != 0 {
		v113 = v97
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v113 = base.B2i32(v98 != v82)
	goto L24
L27:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	v122 = v119
	v123 = v98 + int32(1)
	goto L29
L28:
	;
	v122 = v100
	v123 = v98
	goto L29
L29:
	;
	v125 = v28 + int32(16)
	v127 = m.G0
	v129 = v127 - int32(80)
	m.G0 = v129
	v132 = int32(48)
	v135 = v28 + int32(124)
	v136 = base.I32_extend8_s(v122)
	switch v136 - int32(37) {
	case 0:
		goto L42
	default:
		v854 = int32(0)
		goto L30
	case 28:
		goto L74
	case 29:
		goto L72
	case 30:
		goto L71
	case 31:
		v792 = int32(_a_F___strftime_l_1)
		goto L36
	case 33:
		goto L68
	case 34, 66:
		goto L67
	case 35:
		goto L66
	case 36:
		goto L65
	case 40:
		goto L62
	case 45:
		goto L59
	case 46:
		goto L57
	case 47:
		goto L55
	case 48:
		goto L53
	case 49:
		goto L51
	case 50:
		goto L52
	case 51:
		goto L47
	case 52:
		goto L45
	case 53:
		goto L43
	case 60:
		goto L75
	case 61, 67:
		goto L73
	case 62:
		v730 = int32(_a_F___strftime_l_2)
		goto L37
	case 63:
		v163 = v132
		goto L69
	case 64:
		goto L70
	case 69:
		goto L64
	case 72:
		goto L63
	case 73:
		goto L61
	case 75:
		goto L60
	case 77:
		goto L38
	case 78:
		goto L58
	case 79:
		goto L56
	case 80:
		goto L54
	case 82:
		goto L50
	case 83:
		goto L48
	case 84:
		goto L46
	case 85:
		goto L44
	}
L30:
	;
	m.G0 = v129 + int32(80)
	if v854 == int32(0) {
		v1290 = v38
		goto L7
	} else {
		goto L262
	}
L31:
	;
	v852 = F_strlen(m, v851)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v852
	v854 = v851
	goto L30
L32:
	;
	v851 = int32(_a_F___strftime_l_3)
	goto L31
L33:
	;
	if v76 != 0 {
		goto L252
	} else {
		goto L253
	}
L34:
	;
	v808 = int32(4)
	v813 = v132
	v816 = v805
	goto L33
L35:
	;
	v808 = int32(2)
	v813 = v800
	v816 = v801
	goto L33
L36:
	;
	v794 = F___strftime_l(m, v125, int32(100), v792, l3, l4)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L186
	} else {
		goto L247
	}
L37:
	;
	if v730 == int32(14) {
		goto L222
	} else {
		goto L223
	}
L38:
	;
	v730 = int32(_a_F___strftime_l_4)
	goto L37
L39:
	;
	if v668 == int32(14) {
		goto L196
	} else {
		goto L197
	}
L40:
	;
	v668 = v139 | int32(_a_F___strftime_l_5)
	goto L39
L41:
	;
	v664 = base.I64_rem_s(v290, int64(100))
	v800 = v132
	v801 = v664
	goto L35
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(1)
	v854 = int32(_a_F___strftime_l_6)
	goto L30
L43:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v652 < int32(0) {
		goto L192
	} else {
		goto L193
	}
L44:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v625 < int32(0) {
		goto L188
	} else {
		goto L189
	}
L45:
	;
	v610 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v612 = v610 + int64(1900)
	if v610 < int64(8100) {
		v805 = v612
		goto L34
	} else {
		goto L185
	}
L46:
	;
	v601 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v605 = base.I64_rem_s(v601+int64(1900), int64(100))
	v607 = v605 >> (uint(int64(63)) % 64)
	v800 = v132
	v801 = v605 ^ v607 - v607
	goto L35
L47:
	;
	v730 = int32(_a_F___strftime_l_7)
	goto L37
L48:
	;
	v730 = int32(_a_F___strftime_l_8)
	goto L37
L49:
	;
	v808 = int32(1)
	v813 = v132
	v816 = v597
	goto L33
L50:
	;
	v586 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+24)))
	v597 = v586
	goto L49
L51:
	;
	v535 = int32(53)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v540 = int32(7)
	v541 = base.I32_rem_u_s(v537+int32(6), v540)
	v546 = base.I32_div_u_s(v536-v541+v540, v540)
	v547 = v537 - v536
	v551 = base.I32_rem_u_s(v547+int32(369), v540)
	v554 = v546 + base.B2i32(base.Ui32(v551) < base.Ui32(int32(3)))
	if v554 != v535 {
		goto L175
	} else {
		goto L176
	}
L52:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v524 = int32(7)
	v525 = base.I32_rem_u_s(v521+int32(6), v524)
	v530 = base.I32_div_u_s(v520-v525+v524, v524)
	v800 = v132
	v801 = base.I64_extend_i32_u(v530)
	goto L35
L53:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v515 = int32(7)
	v518 = base.I32_div_u_s(v512-v513+v515, v515)
	v800 = v132
	v801 = base.I64_extend_i32_u(v518)
	goto L35
L54:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v508 != 0 {
		goto L170
	} else {
		goto L171
	}
L55:
	;
	v792 = int32(_a_F___strftime_l_9)
	goto L36
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(1)
	v854 = int32(_a_F___strftime_l_10)
	goto L30
L57:
	;
	v503 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3))))
	v800 = v132
	v801 = v503
	goto L35
L58:
	;
	v325 = int32(0)
	v327 = m.G0
	v329 = v327 - int32(16)
	m.G0 = v329
	v331 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui32(int32(12)) <= base.Ui32(v332) {
		goto L126
	} else {
		goto L127
	}
L59:
	;
	v792 = int32(_a_F___strftime_l_11)
	goto L36
L60:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if int32(11) < v320 {
		goto L122
	} else {
		goto L123
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(1)
	v854 = int32(_a_F___strftime_l_12)
	goto L30
L62:
	;
	v314 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+4)))
	v800 = v132
	v801 = v314
	goto L35
L63:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v800 = v132
	v801 = base.I64_extend_i32_s(v310 + int32(1))
	goto L35
L64:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v808 = int32(3)
	v813 = v132
	v816 = base.I64_extend_i32_s(v305 + int32(1))
	goto L33
L65:
	;
	v294 = int32(2)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v295 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L66:
	;
	v293 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+8)))
	v800 = v132
	v801 = v293
	goto L35
L67:
	;
	v166 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v168 = v166 + int64(1900)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	if v169 <= int32(2) {
		goto L81
	} else {
		goto L82
	}
L68:
	;
	v792 = int32(_a_F___strftime_l_13)
	goto L36
L69:
	;
	v164 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+12)))
	v800 = v163
	v801 = v164
	goto L35
L70:
	;
	v163 = int32(95)
	goto L69
L71:
	;
	v157 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v161 = base.I64_div_s(v157+int64(1900), int64(100))
	v800 = v132
	v801 = v161
	goto L35
L72:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui32(int32(11)) < base.Ui32(v152) {
		goto L32
	} else {
		goto L79
	}
L73:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui32(int32(11)) < base.Ui32(v147) {
		goto L32
	} else {
		goto L78
	}
L74:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if base.Ui32(int32(6)) < base.Ui32(v142) {
		goto L32
	} else {
		goto L77
	}
L75:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if base.Ui32(v139) <= base.Ui32(int32(6)) {
		goto L40
	} else {
		goto L76
	}
L76:
	;
	goto L32
L77:
	;
	v668 = v142 + int32(_a_F___strftime_l_14)
	goto L39
L78:
	;
	v668 = v147 + int32(_a_F___strftime_l_15)
	goto L39
L79:
	;
	v668 = v152 + int32(_a_F___strftime_l_16)
	goto L39
L80:
	;
	if v136 == int32(103) {
		goto L41
	} else {
		goto L115
	}
L81:
	;
	v177 = int32(53)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v182 = int32(7)
	v183 = base.I32_rem_u_s(v179+int32(6), v182)
	v188 = base.I32_div_u_s(v178-v183+v182, v182)
	v189 = v179 - v178
	v193 = base.I32_rem_u_s(v189+int32(369), v182)
	v196 = v188 + base.B2i32(base.Ui32(v193) < base.Ui32(int32(3)))
	if v196 != v177 {
		goto L86
	} else {
		goto L87
	}
L82:
	;
	goto L83
L83:
	;
	if base.Ui32(v169) < base.Ui32(int32(361)) {
		v290 = v168
		goto L80
	} else {
		goto L99
	}
L84:
	;
	if v226 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L85:
	;
	v226 = v224
	goto L84
L86:
	;
	if v196 != 0 {
		v224 = v196
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v218 = base.I32_rem_u_s(v189+int32(371), int32(7))
	switch v218 - int32(3) {
	case 0:
		goto L94
	case 1:
		v224 = v177
		goto L85
	default:
		goto L93
	}
L89:
	;
	v199 = int32(52)
	v203 = base.I32_rem_u_s(v189+int32(6), int32(7))
	switch v203 - int32(4) {
	case 0:
		goto L90
	case 1:
		goto L91
	default:
		v224 = v199
		goto L85
	}
L90:
	;
	v226 = int32(53)
	goto L84
L91:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v208 = base.I32_rem_s(v206, int32(400))
	v211 = F_is_leap(m, v208-int32(1))
	mBase = m.M
	if v211 == int32(0) {
		v224 = v199
		goto L85
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v224 = int32(1)
	goto L85
L94:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v222 = F_is_leap(m, v221)
	mBase = m.M
	if v222 != 0 {
		v224 = v177
		goto L85
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v229 = v168
	goto L98
L97:
	;
	v229 = v166 + int64(1899)
	goto L98
L98:
	;
	v290 = v229
	goto L80
L99:
	;
	v237 = int32(53)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v242 = int32(7)
	v243 = base.I32_rem_u_s(v239+int32(6), v242)
	v248 = base.I32_div_u_s(v238-v243+v242, v242)
	v249 = v239 - v238
	v253 = base.I32_rem_u_s(v249+int32(369), v242)
	v256 = v248 + base.B2i32(base.Ui32(v253) < base.Ui32(int32(3)))
	if v256 != v237 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	if v286 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L101:
	;
	v286 = v284
	goto L100
L102:
	;
	if v256 != 0 {
		v284 = v256
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v278 = base.I32_rem_u_s(v249+int32(371), int32(7))
	switch v278 - int32(3) {
	case 0:
		goto L110
	case 1:
		v284 = v237
		goto L101
	default:
		goto L109
	}
L105:
	;
	v259 = int32(52)
	v263 = base.I32_rem_u_s(v249+int32(6), int32(7))
	switch v263 - int32(4) {
	case 0:
		goto L106
	case 1:
		goto L107
	default:
		v284 = v259
		goto L101
	}
L106:
	;
	v286 = int32(53)
	goto L100
L107:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v268 = base.I32_rem_s(v266, int32(400))
	v271 = F_is_leap(m, v268-int32(1))
	mBase = m.M
	if v271 == int32(0) {
		v284 = v259
		goto L101
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v284 = int32(1)
	goto L101
L110:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v282 = F_is_leap(m, v281)
	mBase = m.M
	if v282 != 0 {
		v284 = v237
		goto L101
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v289 = v166 + int64(1901)
	goto L114
L113:
	;
	v289 = v168
	goto L114
L114:
	;
	v290 = v289
	goto L80
L115:
	;
	v805 = v290
	goto L34
L116:
	;
	v808 = v294
	v813 = v132
	v816 = int64(12)
	goto L33
L117:
	;
	goto L118
L118:
	;
	v299 = base.I64_extend_i32_s(v295)
	if int32(12) < v295 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v304 = v299 - int64(12)
	goto L121
L120:
	;
	v304 = v299
	goto L121
L121:
	;
	v808 = v294
	v813 = v132
	v816 = v304
	goto L33
L122:
	;
	v323 = int32(_a_F___strftime_l_17)
	goto L124
L123:
	;
	v323 = int32(_a_F___strftime_l_18)
	goto L124
L124:
	;
	v668 = v323
	goto L39
L125:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v352<<(uint(int32(2))%32))+uint32(_c_F___strftime_l[0])))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	if v474 != 0 {
		goto L164
	} else {
		goto L165
	}
L126:
	;
	v335 = int32(12)
	v336 = base.I32_div_s(v332, v335)
	v339 = v332 - v336*v335
	if v339 < int32(0) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v352 = v332
	v353 = v331
	goto L128
L128:
	;
	v355 = v329 + int32(12)
	if base.Ui64(v353-int64(2)) <= base.Ui64(int64(136)) {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	v344 = v339 + v335
	goto L131
L130:
	;
	v344 = v339
	goto L131
L131:
	;
	v352 = v344
	v353 = base.I64_extend_i32_s(v336+v339>>(uint(int32(31))%32)) + v331
	goto L128
L132:
	;
	v360 = base.I32_wrap_i64(v353)
	v364 = (v360 - int32(68)) >> (uint(int32(2)) % 32)
	if v360&int32(3) == int32(0) {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	goto L134
L134:
	;
	v390 = v353 - int64(100)
	v391 = int64(400)
	v392 = base.I64_div_s(v390, v391)
	v395 = v390 - v392*v391
	v401 = base.I32_wrap_i64(v395)
	if v395 < int64(0) {
		goto L145
	} else {
		goto L146
	}
L135:
	;
	v466 = base.I64_extend_i32_s(v360*int32(31536000) + v380*int32(_a_F___strftime_l_19) + int32(2087447296))
	goto L125
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355))) = v378
	v380 = v377
	goto L135
L137:
	;
	v370 = v364 - int32(1)
	if v355 == int32(0) {
		v380 = v370
		goto L135
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	if v355 == int32(0) {
		v380 = v364
		goto L135
	} else {
		goto L141
	}
L140:
	;
	v377 = v370
	v378 = int32(1)
	goto L136
L141:
	;
	v377 = v364
	v378 = int32(0)
	goto L136
L142:
	;
	v466 = v390*int64(31536000) + base.I64_extend_i32_s(v443+(v442*int32(24)+(base.I32_wrap_i64(v395>>(uint(int64(63))%64))+base.I32_wrap_i64(v392))*int32(97))-v441)*int64(86400) + int64(946771200)
	goto L125
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355))) = v437
	v441 = v437
	v442 = v438
	v443 = v439
	goto L142
L144:
	;
	v430 = int32(base.Ui32(v423) >> (uint(int32(2)) % 32))
	v433 = int32(0)
	v434 = base.B2i32(v423&int32(3) == v433)
	if v355 == v433 {
		v441 = v434
		v442 = v422
		v443 = v430
		goto L142
	} else {
		goto L163
	}
L145:
	;
	v406 = v401 + int32(400)
	goto L147
L146:
	;
	v406 = v401
	goto L147
L147:
	;
	if v406 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	if int32(200) <= v406 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v427 = v325
	v428 = int32(1)
	goto L150
L150:
	;
	if v355 != 0 {
		v437 = v428
		v438 = v427
		v439 = v325
		goto L143
	} else {
		goto L162
	}
L151:
	;
	if v423 != 0 {
		goto L144
	} else {
		goto L161
	}
L152:
	;
	if base.Ui32(int32(300)) <= base.Ui32(v406) {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	goto L154
L154:
	;
	v420 = base.B2i32(int32(99) < v406)
	if int32(99) < v406 {
		goto L158
	} else {
		goto L159
	}
L155:
	;
	v422 = int32(3)
	v423 = v406 - int32(300)
	goto L151
L156:
	;
	goto L157
L157:
	;
	v422 = int32(2)
	v423 = v406 - int32(200)
	goto L151
L158:
	;
	v421 = v406 - int32(100)
	goto L160
L159:
	;
	v421 = v406
	goto L160
L160:
	;
	v422 = v420
	v423 = v421
	goto L151
L161:
	;
	v427 = v422
	v428 = int32(0)
	goto L150
L162:
	;
	v441 = v428
	v442 = v427
	v443 = v325
	goto L142
L163:
	;
	v437 = v434
	v438 = v422
	v439 = v430
	goto L143
L164:
	;
	v475 = v471 + int32(_a_F___strftime_l_19)
	goto L166
L165:
	;
	v475 = v471
	goto L166
L166:
	;
	if int32(1) < v352 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v478 = v475
	goto L169
L168:
	;
	v478 = v471
	goto L169
L169:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v480 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+8)))
	v481 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+4)))
	v482 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3))))
	m.G0 = v329 + int32(16)
	v501 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+36)))
	v597 = v482 + (v466 + base.I64_extend_i32_s(v478) + base.I64_extend_i32_s(v479-int32(1))*int64(86400) + v480*int64(3600) + v481*int64(60)) - v501
	goto L49
L170:
	;
	v510 = v508
	goto L172
L171:
	;
	v510 = int32(7)
	goto L172
L172:
	;
	v597 = base.I64_extend_i32_s(v510)
	goto L49
L173:
	;
	v800 = v132
	v801 = base.I64_extend_i32_u(v584)
	goto L35
L174:
	;
	v584 = v582
	goto L173
L175:
	;
	if v554 != 0 {
		v582 = v554
		goto L174
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v576 = base.I32_rem_u_s(v547+int32(371), int32(7))
	switch v576 - int32(3) {
	case 0:
		goto L183
	case 1:
		v582 = v535
		goto L174
	default:
		goto L182
	}
L178:
	;
	v557 = int32(52)
	v561 = base.I32_rem_u_s(v547+int32(6), int32(7))
	switch v561 - int32(4) {
	case 0:
		goto L179
	case 1:
		goto L180
	default:
		v582 = v557
		goto L174
	}
L179:
	;
	v584 = int32(53)
	goto L173
L180:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v566 = base.I32_rem_s(v564, int32(400))
	v569 = F_is_leap(m, v566-int32(1))
	mBase = m.M
	if v569 == int32(0) {
		v582 = v557
		goto L174
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	v582 = int32(1)
	goto L174
L183:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v580 = F_is_leap(m, v579)
	mBase = m.M
	if v580 != 0 {
		v582 = v535
		goto L174
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v129)+48)) = v612
	v620 = F_snprintf(m, v125, int32(100), int32(_a_F___strftime_l_20), v129+int32(48))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	return int32(0)
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v620
	v854 = v125
	goto L30
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(0)
	v854 = int32(_a_F___strftime_l_21)
	goto L30
L189:
	;
	goto L190
L190:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	v632 = int32(3600)
	v633 = base.I32_div_s(v631, v632)
	v634 = int32(100)
	v641 = base.I32_div_s(base.I32_extend16_s(v631-v633*v632), int32(60))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+64)) = v633*v634 + base.I32_extend16_s(v641)
	v649 = F_snprintf(m, v125, v634, int32(_a_F___strftime_l_22), v129-int32(-64))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L186
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v649
	v854 = v125
	goto L30
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(0)
	v854 = int32(_a_F___strftime_l_21)
	goto L30
L193:
	;
	goto L194
L194:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	F_do_tzset(m)
	mBase = m.M
	v851 = v658
	goto L31
L195:
	;
	v851 = v728
	goto L31
L196:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v675 != 0 {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	goto L198
L198:
	;
	v677 = int32(_a_F___strftime_l_23)
	v678 = v668 & v677
	v682 = v668 >> (uint(int32(16)) % 32)
	if base.B2i32(v678 != v677)|base.B2i32(int32(5) < v682) == int32(0) {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	v676 = int32(_a_F___strftime_l_24)
	goto L201
L200:
	;
	v676 = int32(_a_F___strftime_l_25)
	goto L201
L201:
	;
	v728 = v676
	goto L195
L202:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l4+v682<<(uint(int32(2))%32))))
	if v691 != 0 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	goto L204
L204:
	;
	v696 = int32(_a_F___strftime_l_21)
	switch v682 - int32(1) {
	case 0:
		goto L212
	case 1:
		goto L211
	default:
		v720 = v696
		goto L208
	case 4:
		goto L210
	}
L205:
	;
	v695 = v691 + int32(8)
	goto L207
L206:
	;
	v695 = int32(_a_F___strftime_l_26)
	goto L207
L207:
	;
	v728 = v695
	goto L195
L208:
	;
	v728 = v720
	goto L195
L209:
	;
	if v678 == int32(0) {
		v720 = v708
		goto L208
	} else {
		goto L216
	}
L210:
	;
	if base.Ui32(int32(3)) < base.Ui32(v678) {
		v720 = v696
		goto L208
	} else {
		goto L215
	}
L211:
	;
	if base.Ui32(int32(49)) < base.Ui32(v678) {
		v720 = v696
		goto L208
	} else {
		goto L214
	}
L212:
	;
	if base.Ui32(int32(1)) < base.Ui32(v678) {
		v720 = v696
		goto L208
	} else {
		goto L213
	}
L213:
	;
	v708 = int32(_a_F___strftime_l_27)
	goto L209
L214:
	;
	v708 = int32(_a_F___strftime_l_28)
	goto L209
L215:
	;
	v708 = int32(_a_F___strftime_l_29)
	goto L209
L216:
	;
	v711 = v708
	v713 = v678
	goto L217
L217:
	;
	v716 = v711 + int32(1)
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711))))
	if v717 != 0 {
		v711 = v716
		goto L217
	} else {
		goto L219
	}
L218:
	;
	v720 = v716
	goto L208
L219:
	;
	v719 = v713 - int32(1)
	if v719 != 0 {
		v711 = v716
		v713 = v719
		goto L217
	} else {
		goto L220
	}
L220:
	;
	goto L218
L221:
	;
	v792 = v790
	goto L36
L222:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v737 != 0 {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	v739 = int32(_a_F___strftime_l_23)
	v740 = v730 & v739
	v744 = v730 >> (uint(int32(16)) % 32)
	if base.B2i32(v740 != v739)|base.B2i32(int32(5) < v744) == int32(0) {
		goto L228
	} else {
		goto L229
	}
L225:
	;
	v738 = int32(_a_F___strftime_l_24)
	goto L227
L226:
	;
	v738 = int32(_a_F___strftime_l_25)
	goto L227
L227:
	;
	v790 = v738
	goto L221
L228:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l4+v744<<(uint(int32(2))%32))))
	if v753 != 0 {
		goto L231
	} else {
		goto L232
	}
L229:
	;
	goto L230
L230:
	;
	v758 = int32(_a_F___strftime_l_21)
	switch v744 - int32(1) {
	case 0:
		goto L238
	case 1:
		goto L237
	default:
		v782 = v758
		goto L234
	case 4:
		goto L236
	}
L231:
	;
	v757 = v753 + int32(8)
	goto L233
L232:
	;
	v757 = int32(_a_F___strftime_l_26)
	goto L233
L233:
	;
	v790 = v757
	goto L221
L234:
	;
	v790 = v782
	goto L221
L235:
	;
	if v740 == int32(0) {
		v782 = v770
		goto L234
	} else {
		goto L242
	}
L236:
	;
	if base.Ui32(int32(3)) < base.Ui32(v740) {
		v782 = v758
		goto L234
	} else {
		goto L241
	}
L237:
	;
	if base.Ui32(int32(49)) < base.Ui32(v740) {
		v782 = v758
		goto L234
	} else {
		goto L240
	}
L238:
	;
	if base.Ui32(int32(1)) < base.Ui32(v740) {
		v782 = v758
		goto L234
	} else {
		goto L239
	}
L239:
	;
	v770 = int32(_a_F___strftime_l_27)
	goto L235
L240:
	;
	v770 = int32(_a_F___strftime_l_28)
	goto L235
L241:
	;
	v770 = int32(_a_F___strftime_l_29)
	goto L235
L242:
	;
	v773 = v770
	v775 = v740
	goto L243
L243:
	;
	v778 = v773 + int32(1)
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773))))
	if v779 != 0 {
		v773 = v778
		goto L243
	} else {
		goto L245
	}
L244:
	;
	v782 = v778
	goto L234
L245:
	;
	v781 = v775 - int32(1)
	if v781 != 0 {
		v773 = v778
		v775 = v781
		goto L243
	} else {
		goto L246
	}
L246:
	;
	goto L244
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v794
	if v794 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v798 = v125
	goto L250
L249:
	;
	v798 = int32(0)
	goto L250
L250:
	;
	v854 = v798
	goto L30
L251:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v129)+8)) = v816
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v808
	v845 = F_snprintf(m, v125, int32(100), int32(_a_F___strftime_l_30), v129)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L186
	} else {
		goto L261
	}
L252:
	;
	v819 = v76
	goto L254
L253:
	;
	v819 = v813
	goto L254
L254:
	;
	if v819 != int32(95) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	if v819 != int32(45) {
		goto L251
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v129)+40)) = v816
	*(*int32)(unsafe.Add(mBase, uint32(v129)+32)) = v808
	v838 = F_snprintf(m, v125, int32(100), int32(_a_F___strftime_l_31), v129+int32(32))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L186
	} else {
		goto L260
	}
L258:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v129)+16)) = v816
	v829 = F_snprintf(m, v125, int32(100), int32(_a_F___strftime_l_32), v129+int32(16))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L186
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v829
	v854 = v125
	goto L30
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v838
	v854 = v125
	goto L30
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v845
	v854 = v125
	goto L30
L262:
	;
	if v113 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1077 = l1 - v1060
	if base.Ui32(v1059) < base.Ui32(v1077) {
		goto L296
	} else {
		goto L297
	}
L264:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	v1059 = v873
	v1060 = v38
	v1061 = v854
	goto L263
L265:
	;
	goto L266
L266:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v854))))
	switch v874 - int32(43) {
	case 0, 2:
		goto L268
	default:
		goto L269
	}
L267:
	;
	if v884&int32(255) != int32(48) {
		v934 = v886
		v936 = v885
		goto L270
	} else {
		goto L271
	}
L268:
	;
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v854)+1)))
	v879 = int32(1)
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	v884 = v878
	v885 = v854 + v879
	v886 = v881 - v879
	goto L267
L269:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	v884 = v874
	v885 = v854
	v886 = v877
	goto L267
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v934
	v960 = int32(0)
	goto L276
L271:
	;
	v898 = v886
	v900 = v885
	goto L272
L272:
	;
	v916 = int32(*(*int8)(unsafe.Add(mBase, uint32(v900)+1)))
	if base.Ui32(int32(9)) < base.Ui32(v916-int32(48)) {
		v934 = v898
		v936 = v900
		goto L270
	} else {
		goto L274
	}
L273:
	;
	v934 = v924
	v936 = v922
	goto L270
L274:
	;
	v921 = int32(1)
	v922 = v900 + v921
	v924 = v898 - v921
	if v916 == int32(48) {
		v898 = v924
		v900 = v922
		goto L272
	} else {
		goto L275
	}
L275:
	;
	goto L273
L276:
	;
	v982 = int32(*(*int8)(unsafe.Add(mBase, uint32(v960+v936))))
	if base.Ui32(v982-int32(48)) < base.Ui32(int32(10)) {
		v960 = v960 + int32(1)
		goto L276
	} else {
		goto L278
	}
L277:
	;
	if base.Ui32(v934) < base.Ui32(v113) {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	goto L277
L279:
	;
	v988 = v113
	goto L281
L280:
	;
	v988 = v934
	goto L281
L281:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v990 < int32(-1900) {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	if base.B2i32(base.Ui32(v1013) <= base.Ui32(v934))|base.B2i32(base.Ui32(l1) <= base.Ui32(v1014)) != 0 {
		v1059 = v934
		v1060 = v1014
		v1061 = v936
		goto L263
	} else {
		goto L291
	}
L283:
	;
	v1007 = int32(45)
	goto L285
L284:
	;
	if v79 != int32(43) {
		v1013 = v988
		v1014 = v38
		goto L282
	} else {
		goto L286
	}
L285:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v38))) = uint8(v1007)
	v1009 = int32(1)
	v1013 = v988 - v1009
	v1014 = v38 + v1009
	goto L282
L286:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1000))))
	if v1001 == int32(67) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1004 = int32(3)
	goto L289
L288:
	;
	v1004 = int32(5)
	goto L289
L289:
	;
	if base.Ui32(v988-v934+v960) < base.Ui32(v1004) {
		v1013 = v988
		v1014 = v38
		goto L282
	} else {
		goto L290
	}
L290:
	;
	v1007 = int32(43)
	goto L285
L291:
	;
	v1024 = v1013
	v1026 = v1014
	goto L292
L292:
	;
	v1044 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v1026))) = uint8(v1044)
	v1046 = int32(1)
	v1047 = v1026 + v1046
	v1049 = v1024 - v1046
	if base.Ui32(v1049) <= base.Ui32(v934) {
		v1059 = v934
		v1060 = v1047
		v1061 = v936
		goto L263
	} else {
		goto L294
	}
L293:
	;
	v1059 = v934
	v1060 = v1047
	v1061 = v936
	goto L263
L294:
	;
	if base.Ui32(v1047) < base.Ui32(l1) {
		v1024 = v1049
		v1026 = v1047
		goto L292
	} else {
		goto L295
	}
L295:
	;
	goto L293
L296:
	;
	v1079 = v1059
	goto L298
L297:
	;
	v1079 = v1077
	goto L298
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v1079
	v1081 = l0 + v1060
	if base.Ui32(int32(512)) <= base.Ui32(v1079) {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	v1255 = v123
	v1278 = v1251 + v1060
	goto L8
L300:
	;
	if v1079 != 0 {
		goto L303
	} else {
		goto L304
	}
L301:
	;
	goto L302
L302:
	;
	v1088 = v1081 + v1079
	if (v1081^v1061)&int32(3) == int32(0) {
		goto L307
	} else {
		goto L308
	}
L303:
	;
	base.MemoryCopy(m, v1081, v1061, v1079)
	goto L305
L304:
	;
	goto L305
L305:
	;
	goto L299
L306:
	;
	if base.Ui32(v1220) < base.Ui32(v1088) {
		goto L340
	} else {
		goto L341
	}
L307:
	;
	if v1081&int32(3) == int32(0) {
		goto L311
	} else {
		goto L312
	}
L308:
	;
	goto L309
L309:
	;
	if base.Ui32(v1088) < base.Ui32(int32(4)) {
		goto L331
	} else {
		goto L332
	}
L310:
	;
	v1124 = v1088 & int32(-4)
	if base.Ui32(v1088) < base.Ui32(int32(64)) {
		v1174 = v1118
		v1175 = v1119
		goto L321
	} else {
		goto L322
	}
L311:
	;
	v1118 = v1061
	v1119 = v1081
	goto L310
L312:
	;
	goto L313
L313:
	;
	if v1079 == int32(0) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1118 = v1061
	v1119 = v1081
	goto L310
L315:
	;
	goto L316
L316:
	;
	v1101 = v1061
	v1102 = v1081
	goto L317
L317:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1102))) = uint8(v1106)
	v1108 = int32(1)
	v1109 = v1101 + v1108
	v1111 = v1102 + v1108
	if v1111&int32(3) == int32(0) {
		v1118 = v1109
		v1119 = v1111
		goto L310
	} else {
		goto L319
	}
L318:
	;
	v1118 = v1109
	v1119 = v1111
	goto L310
L319:
	;
	if base.Ui32(v1111) < base.Ui32(v1088) {
		v1101 = v1109
		v1102 = v1111
		goto L317
	} else {
		goto L320
	}
L320:
	;
	goto L318
L321:
	;
	if base.Ui32(v1124) <= base.Ui32(v1175) {
		v1219 = v1174
		v1220 = v1175
		goto L306
	} else {
		goto L327
	}
L322:
	;
	v1128 = v1124 + int32(-64)
	if base.Ui32(v1128) < base.Ui32(v1119) {
		v1174 = v1118
		v1175 = v1119
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v1131 = v1118
	v1132 = v1119
	goto L324
L324:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1131)))
	*(*int32)(unsafe.Add(mBase, uint32(v1132))) = v1136
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+4)) = v1138
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+8)) = v1140
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+12)) = v1142
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+16)) = v1144
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+20)) = v1146
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+24)) = v1148
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+28)) = v1150
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+32)) = v1152
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+36)) = v1154
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+40)) = v1156
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+44)) = v1158
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+48)) = v1160
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+52)) = v1162
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+56)) = v1164
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+60)) = v1166
	v1168 = int32(-64)
	v1169 = v1131 - v1168
	v1171 = v1132 - v1168
	if base.Ui32(v1171) <= base.Ui32(v1128) {
		v1131 = v1169
		v1132 = v1171
		goto L324
	} else {
		goto L326
	}
L325:
	;
	v1174 = v1169
	v1175 = v1171
	goto L321
L326:
	;
	goto L325
L327:
	;
	v1181 = v1174
	v1182 = v1175
	goto L328
L328:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1181)))
	*(*int32)(unsafe.Add(mBase, uint32(v1182))) = v1186
	v1188 = int32(4)
	v1189 = v1181 + v1188
	v1191 = v1182 + v1188
	if base.Ui32(v1191) < base.Ui32(v1124) {
		v1181 = v1189
		v1182 = v1191
		goto L328
	} else {
		goto L330
	}
L329:
	;
	v1219 = v1189
	v1220 = v1191
	goto L306
L330:
	;
	goto L329
L331:
	;
	v1219 = v1061
	v1220 = v1081
	goto L306
L332:
	;
	goto L333
L333:
	;
	if base.Ui32(v1079) < base.Ui32(int32(4)) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1219 = v1061
	v1220 = v1081
	goto L306
L335:
	;
	goto L336
L336:
	;
	v1200 = v1061
	v1201 = v1081
	goto L337
L337:
	;
	v1205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1201))) = uint8(v1205)
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1201)+1)) = uint8(v1207)
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1201)+2)) = uint8(v1209)
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1201)+3)) = uint8(v1211)
	v1213 = int32(4)
	v1214 = v1200 + v1213
	v1216 = v1201 + v1213
	if base.Ui32(v1216) <= base.Ui32(v1088-int32(4)) {
		v1200 = v1214
		v1201 = v1216
		goto L337
	} else {
		goto L339
	}
L338:
	;
	v1219 = v1214
	v1220 = v1216
	goto L306
L339:
	;
	goto L338
L340:
	;
	v1226 = v1219
	v1227 = v1220
	goto L343
L341:
	;
	goto L342
L342:
	;
	goto L299
L343:
	;
	v1231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1227))) = uint8(v1231)
	v1233 = int32(1)
	v1236 = v1227 + v1233
	if v1236 != v1088 {
		v1226 = v1226 + v1233
		v1227 = v1236
		goto L343
	} else {
		goto L345
	}
L344:
	;
	goto L342
L345:
	;
	goto L344
L346:
	;
	v1290 = v1278
	goto L7
L347:
	;
	v1310 = l1 - int32(1)
	goto L349
L348:
	;
	v1310 = v1290
	goto L349
L349:
	;
	v1320 = v1310
	v1337 = int32(0)
	goto L4
}
func F___syscall_ret(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	if base.Ui32(int32(-4095)) <= base.Ui32(l0) {
		*(*int32)(unsafe.Add(mBase, _c_F___syscall_ret[0])) = int32(0) - l0
		v9 = int32(-1)
	} else {
		v9 = l0
	}
	return v9
}
func F__soundex(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var __phi51 int32
	_ = __phi51
	var v52 int32
	_ = v52
	var __phi52 int32
	_ = __phi52
	var v54 int32
	_ = v54
	var __phi54 int32
	_ = __phi54
	var v55 int32
	_ = v55
	var __phi55 int32
	_ = __phi55
	var v56 int32
	_ = v56
	var __phi56 int32
	_ = __phi56
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if base.Ui32((v10-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v8 = l0
	v10 = v7
	goto L5
L3:
	;
	goto L4
L4:
	;
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v31)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
	return
L5:
	;
	if base.Ui32((v10&int32(223)-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L4
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	if v22 != 0 {
		v8 = v8 + int32(1)
		v10 = v22
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v43 = v10 - int32(32)
	goto L11
L10:
	;
	v43 = v10
	goto L11
L11:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v43)
	v45 = int32(1)
	v47 = l1 + v45
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v161 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v161)
	return
L13:
	;
	__phi51 = v8
	__phi52 = v48
	__phi54 = v47
	__phi55 = v45
	__phi56 = v8 + int32(1)
	v51 = __phi51
	v52 = __phi52
	v54 = __phi54
	v55 = __phi55
	v56 = __phi56
	goto L16
L14:
	;
	v147 = v47
	v148 = v45
	goto L15
L15:
	;
	v151 = int32(4) - v148
	if v151 != 0 {
		goto L42
	} else {
		goto L43
	}
L16:
	;
	if base.Ui32(int32(25)) < base.Ui32((v52&int32(223)-int32(65))&int32(255)) {
		v132 = v54
		v133 = v55
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if int32(3) < v133 {
		v158 = v132
		goto L12
	} else {
		goto L41
	}
L18:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v135 != 0 {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	if base.Ui32((v52-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v73 = v52 - int32(32)
	goto L22
L21:
	;
	v73 = v52
	goto L22
L22:
	;
	v79 = base.B2i32(base.Ui32(int32(25)) < base.Ui32((v73-int32(65))&int32(255)))
	if base.Ui32(int32(25)) < base.Ui32((v73-int32(65))&int32(255)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v85 = v73
	goto L25
L24:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73&int32(255))+uint32(_c_F__soundex[0]))))
	v85 = v84
	goto L25
L25:
	;
	v86 = int32(255)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if base.Ui32((v88-int32(97))&v86) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v97 = v88 - int32(32)
	goto L28
L27:
	;
	v97 = v88
	goto L28
L28:
	;
	if base.Ui32((v97-int32(65))&int32(255)) <= base.Ui32(int32(25)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97&int32(255))+uint32(_c_F__soundex[0]))))
	v109 = v108
	goto L31
L30:
	;
	v109 = v97
	goto L31
L31:
	;
	if v85&v86 == v109&int32(255) {
		v132 = v54
		v133 = v55
		goto L18
	} else {
		goto L32
	}
L32:
	;
	if v79 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73&int32(255))+uint32(_c_F__soundex[0]))))
	v120 = v119
	goto L35
L34:
	;
	v120 = v73
	goto L35
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v120)
	if v120&int32(255) == int32(48) {
		v132 = v54
		v133 = v55
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v126 = int32(1)
	v132 = v54 + v126
	v133 = v55 + v126
	goto L18
L37:
	;
	if v133 < int32(4) {
		__phi51 = v56
		__phi52 = v135
		__phi54 = v132
		__phi55 = v133
		__phi56 = v56 + int32(1)
		v51 = __phi51
		v52 = __phi52
		v54 = __phi54
		v55 = __phi55
		v56 = __phi56
		goto L16
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	goto L17
L40:
	;
	goto L39
L41:
	;
	v147 = v132
	v148 = v133
	goto L15
L42:
	;
	base.MemoryFill(m, v147, int32(48), v151)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v158 = v151 + v147
	goto L12
}
func F_s_lock(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v64 float64
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	v3 = int32(0)
	v9 = base.AtomicRmwXchg32(m, l0, v3, int32(1))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_s_lock[0])) = v111
	goto L1
L3:
	;
	if int32(999) < v95 {
		goto L1
	} else {
		goto L26
	}
L4:
	;
	F_s_lock_stuck(m, l1, int32(58), int32(_a_F_s_lock_0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L24
	} else {
		goto L25
	}
L5:
	;
	v12 = v3
	v13 = v3
	v14 = v3
	goto L8
L6:
	;
	goto L7
L7:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[0]))
	v95 = v88
	goto L3
L8:
	;
	v16 = v12 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[0]))
	if v18 <= v16 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[0]))
	if v74 == int32(0) {
		v95 = v80
		goto L3
	} else {
		goto L22
	}
L10:
	;
	v21 = v14 + int32(1)
	if int32(1001) <= v21 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	v73 = v16
	v74 = v13
	v75 = v14
	goto L12
L12:
	;
	v78 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v78 != 0 {
		v12 = v73
		v13 = v74
		v14 = v75
		goto L8
	} else {
		goto L21
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(150994951)
	if v13 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v30 = v13
	goto L16
L15:
	;
	v30 = int32(1000)
	goto L16
L16:
	;
	F_pg_usleep(m, v30)
	mBase = m.M
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(0)
	v38 = int32(_a_F_s_lock_1)
	v41 = *(*int64)(unsafe.Add(mBase, _c_F_s_lock[2]))
	v42 = *(*int64)(unsafe.Add(mBase, _c_F_s_lock[3]))
	v43 = v41 ^ v42
	*(*int64)(unsafe.Add(mBase, _c_F_s_lock[3])) = base.I64_rotl(v43, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_s_lock[2])) = v43<<(uint(int64(16))%64) ^ base.I64_rotl(v41, int64(24)) ^ v43
	v64 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v41*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L17
L17:
	;
	v69 = v30 + base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v30), v64), float64(0.5)))
	if int32(_a_F_s_lock_2) < v69 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v72 = int32(1000)
	goto L20
L19:
	;
	v72 = v69
	goto L20
L20:
	;
	v73 = int32(0)
	v74 = v72
	v75 = v21
	goto L12
L21:
	;
	goto L9
L22:
	;
	if v80 < int32(11) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v111 = v80 - int32(1)
	goto L2
L24:
	;
	return
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v100 = int32(900)
	if v100 <= v95 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v103 = v100
	goto L29
L28:
	;
	v103 = v95
	goto L29
L29:
	;
	v111 = v103 + int32(100)
	goto L2
}
func F_satisfies_hash_partition(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v219 int32
	_ = v219
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v277 int64
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v293 int64
	_ = v293
	var v295 int64
	_ = v295
	var v296 int32
	_ = v296
	var v306 int64
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v366 int64
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int64
	_ = v379
	var v381 int64
	_ = v381
	var v382 int32
	_ = v382
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int64
	_ = v394
	var v396 int32
	_ = v396
	var v407 int64
	_ = v407
	var v410 int64
	_ = v410
	var v412 int64
	_ = v412
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
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
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v578 int64
	_ = v578
	v10 = int64(0)
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v17 != 0 {
		v578 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(96)
	return v578
L2:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v18 != 0 {
		v578 = v10
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
	if v19 != 0 {
		v578 = v10
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v21 = base.I32_wrap_i64(v20)
	if int32(0) < v21 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	F_relation_close(m, v35, int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L22
	} else {
		goto L129
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L22
	} else {
		goto L123
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L22
	} else {
		goto L117
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L22
	} else {
		goto L113
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L22
	} else {
		goto L108
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L22
	} else {
		goto L104
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L22
	} else {
		goto L100
	}
L12:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v25 = base.I32_wrap_i64(v24)
	if v25 < int32(0) {
		goto L11
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
	v420 = m.ExcPending
	if v420 != 0 {
		goto L22
	} else {
		goto L96
	}
L15:
	;
	if base.Ui32(v21) <= base.Ui32(v25) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	if v31 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	if v255 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 == v29 {
		v244 = v31
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v35 = F_relation_open(m, v29, int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	return int64(0)
L23:
	;
	v39 = F_RelationGetPartitionKey(m, v35)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v39 == int32(0) {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v43 != int32(104) {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = int32(0)
	if v46 == v47 {
		v58 = v47
		goto L29
	} else {
		goto L30
	}
L27:
	;
	F_relation_close(m, v35, int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L22
	} else {
		goto L64
	}
L28:
	;
	if v58&int32(1) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	goto L28
L30:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	if v50 == int32(0) {
		v58 = v47
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v53 != int32(15) {
		v58 = v47
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+13)))
	v58 = v56
	goto L29
L33:
	;
	v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	v65 = v63 - int32(3)
	v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+4)))
	if v65 != v66 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
	if v145 != 0 {
		goto L5
	} else {
		goto L52
	}
L36:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v74 = F_MemoryContextAllocZero(m, v69, v65*int32(28)+int32(144))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L22
	} else {
		goto L37
	}
L37:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v74
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v29
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v81
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+4)))
	v85 = v83 << (uint(int32(2)) % 32)
	if v85 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	base.MemoryCopy(m, v79+int32(16), v88, v85)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+4)))
	if v90 <= int32(0) {
		v229 = v79
		goto L27
	} else {
		goto L41
	}
L41:
	;
	v98 = int32(0)
	goto L42
L42:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v111 = F_get_fn_expr_argtype(m, v108, v98+int32(3))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L22
	} else {
		goto L44
	}
L43:
	;
	v229 = v79
	goto L27
L44:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v98<<(uint(int32(2))%32))))
	if v111 != v117 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v119 = F_IsBinaryCoercible(m, v111, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L22
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v124 = v98 * int32(28)
	v125 = v79 + int32(144) + v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v127 = v126 + v124
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+20))
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v127)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v125)+16)) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+24)) = v132
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v127)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v125)+8)) = v134
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v127)))
	*(*int64)(unsafe.Add(mBase, uint32(v125))) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v125)+20)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v125)+16)) = int32(0)
	goto L50
L48:
	;
	if v119 == int32(0) {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v142 = v98 + int32(1)
	v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+4)))
	if v142 < v143 {
		v98 = v142
		goto L42
	} else {
		goto L51
	}
L51:
	;
	goto L43
L52:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v147 = F_pg_detoast_datum(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L22
	} else {
		goto L53
	}
L53:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	v152 = F_MemoryContextAllocZero(m, v150, int32(172))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L22
	} else {
		goto L54
	}
L54:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+16)) = v152
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v29
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+4)) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+8)) = v161
	F_get_typlenbyvalalign(m, v161, v157+int32(12), v157+int32(14), v157+int32(15))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L22
	} else {
		goto L55
	}
L55:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+16)) = v172
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+4)))
	if int32(0) < v174 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v182 = int32(0)
	goto L59
L57:
	;
	goto L58
L58:
	;
	v213 = v157 + int32(144)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+20))
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v214)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v213)+16)) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v214)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v213)+24)) = v219
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v214)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v213)+8)) = v221
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v214)))
	*(*int64)(unsafe.Add(mBase, uint32(v213))) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v213)+20)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v213)+16)) = int32(0)
	goto L63
L59:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v178+v182<<(uint(int32(2))%32))))
	if v195 != v177 {
		goto L6
	} else {
		goto L61
	}
L60:
	;
	goto L58
L61:
	;
	v198 = v182 + int32(1)
	if v198 != v174 {
		v182 = v198
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v229 = v157
	goto L27
L64:
	;
	v244 = v229
	goto L17
L65:
	;
	v410 = int64(2147483647)
	v412 = base.I64_rem_u_s(v407, v20&v410)
	v578 = base.I64_extend_i32_u(base.B2i32(v412 == v24&v410))
	goto L1
L66:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v258 <= int32(0) {
		v407 = v10
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
	if v310 == int32(1) {
		v578 = v10
		goto L1
	} else {
		goto L77
	}
L69:
	;
	v268 = int32(0)
	v277 = v10
	goto L70
L70:
	;
	v282 = l0 + int32(24) + v268<<(uint(int32(4))%32)
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+56)))
	if v283 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v407 = v306
	goto L65
L72:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v244+int32(16)+v268<<(uint(int32(2))%32))))
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v282)+48))
	v295 = F_FunctionCall2Coll(m, v244+int32(144)+v268*int32(28), v292, v293, int64(8816678312871386365))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L22
	} else {
		goto L75
	}
L73:
	;
	v306 = v277
	goto L74
L74:
	;
	v308 = v268 + int32(1)
	if v308 != v258 {
		v268 = v308
		v277 = v306
		goto L70
	} else {
		goto L76
	}
L75:
	;
	v306 = v295 + (v277<<(uint(int64(54))%64) + int64(base.Ui64(v277)>>(uint(int64(7))%64))) + int64(5305509591434766563) ^ v277
	goto L74
L76:
	;
	goto L71
L77:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v314 = F_pg_detoast_datum(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L22
	} else {
		goto L78
	}
L78:
	;
	v317 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244)+12)))
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+14)))
	v319 = int32(*(*int8)(unsafe.Add(mBase, uint32(v244)+15)))
	F_deconstruct_array(m, v314, v317, v318, v319, v15+int32(88), v15+int32(84), v15+int32(92))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L22
	} else {
		goto L79
	}
L79:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v328 == v329 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v357 = int32(0)
	v359 = v328
	v366 = v10
	goto L89
L81:
	;
	if int32(0) < v328 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L22
	} else {
		goto L85
	}
L84:
	;
	v407 = v10
	goto L65
L85:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L22
	} else {
		goto L86
	}
L86:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v340
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v342
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_0), v15+int32(16))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L22
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_2), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L22
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v15)+84))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369+v357))))
	if v371 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v407 = v394
	goto L65
L91:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v375+v357<<(uint(int32(3))%32))))
	v381 = F_FunctionCall2Coll(m, v244+int32(144), v374, v379, int64(8816678312871386365))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L22
	} else {
		goto L94
	}
L92:
	;
	v393 = v359
	v394 = v366
	goto L93
L93:
	;
	v396 = v357 + int32(1)
	if v396 < v393 {
		v357 = v396
		v359 = v393
		v366 = v394
		goto L89
	} else {
		goto L95
	}
L94:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	v393 = v392
	v394 = v381 + (v366<<(uint(int64(54))%64) + int64(base.Ui64(v366)>>(uint(int64(7))%64))) + int64(5305509591434766563) ^ v366
	goto L93
L95:
	;
	goto L90
L96:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L22
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_4), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L22
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_5), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L22
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L22
	} else {
		goto L101
	}
L101:
	;
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_6), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L22
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_7), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L22
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L22
	} else {
		goto L105
	}
L105:
	;
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_8), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L22
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_9), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L22
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L22
	} else {
		goto L109
	}
L109:
	;
	v472 = F_get_rel_name(m, v29)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L22
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v472
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_10), v15)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L22
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_11), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L22
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L22
	} else {
		goto L114
	}
L114:
	;
	v490 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v490
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_0), v15-int32(-64))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L22
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_12), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L22
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L22
	} else {
		goto L118
	}
L118:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v510+v98<<(uint(int32(2))%32))))
	v515 = F_format_type_be(m, v514)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L22
	} else {
		goto L119
	}
L119:
	;
	v517 = F_format_type_be(m, v111)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L22
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v98 + int32(1)
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_13), v15+int32(48))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L22
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_14), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L22
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L22
	} else {
		goto L124
	}
L124:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v541+v182<<(uint(int32(2))%32))))
	v546 = F_format_type_be(m, v545)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L22
	} else {
		goto L125
	}
L125:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v549 = F_format_type_be(m, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L22
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v549
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v182 + int32(1)
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_15), v15+int32(32))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L22
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_16), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L22
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	v578 = v10
	goto L1
}
func F_saveNodeLink(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	v2 = l1
	v3 = l2
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14+v15<<(uint(int32(2))%32))+20))
	v22 = v14 + v19&int32(_a_F_saveNodeLink_0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v27 = int32(base.Ui32(v23)>>(uint(int32(3))%32)) & int32(_a_F_saveNodeLink_1)
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)) = uint16(v3)
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+2)) = uint16(v2)
	v74 = int32(base.Ui32(v2) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v37))) = uint16(v74)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_MarkBufferDirty(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L9
	} else {
		goto L13
	}
L2:
	;
	v37 = v22 + int32(base.Ui32(v23)>>(uint(int32(16))%32)) + int32(8)
	v38 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	if v38 == v13 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L4
L7:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+6)))
	v48 = v38 + int32(1)
	if v48 != v27 {
		v37 = v37 + v43&int32(_a_F_saveNodeLink_1)
		v38 = v48
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
	F_errmsg_internal(m, int32(_a_F_saveNodeLink_2), v11)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_saveNodeLink_3), int32(68), int32(_a_F_saveNodeLink_4))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	m.G0 = v11 + int32(16)
	return
}
func F_scalararraysel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v30 int32
	_ = v30
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v69 int32
	_ = v69
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 float64
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 float64
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 float64
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int64
	_ = v240
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 float64
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int64
	_ = v260
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 float64
	_ = v278
	var v279 int32
	_ = v279
	var v282 float64
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v298 float64
	_ = v298
	var v299 int32
	_ = v299
	var v300 float64
	_ = v300
	var v304 float32
	_ = v304
	var v310 float64
	_ = v310
	var v315 float64
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 float64
	_ = v323
	var v331 float64
	_ = v331
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int64
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 float64
	_ = v386
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int64
	_ = v413
	var v414 int64
	_ = v414
	var v424 float64
	_ = v424
	var v425 float64
	_ = v425
	var v429 int32
	_ = v429
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int64
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int64
	_ = v471
	var v472 int32
	_ = v472
	var v475 int64
	_ = v475
	var v476 int32
	_ = v476
	var v479 int64
	_ = v479
	var v480 int32
	_ = v480
	var v481 int64
	_ = v481
	var v482 float64
	_ = v482
	var v489 float64
	_ = v489
	var v490 float64
	_ = v490
	var v496 float64
	_ = v496
	var v497 float64
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v508 float64
	_ = v508
	var v509 float64
	_ = v509
	var v537 float64
	_ = v537
	var v540 float64
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 float64
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v564 int64
	_ = v564
	var v565 int64
	_ = v565
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v577 float64
	_ = v577
	var v578 float64
	_ = v578
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int64
	_ = v619
	var v620 int32
	_ = v620
	var v623 int64
	_ = v623
	var v624 int32
	_ = v624
	var v627 int64
	_ = v627
	var v628 int32
	_ = v628
	var v629 int64
	_ = v629
	var v630 float64
	_ = v630
	var v635 float64
	_ = v635
	var v636 float64
	_ = v636
	var v642 float64
	_ = v642
	var v643 float64
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v654 float64
	_ = v654
	var v655 float64
	_ = v655
	var v681 float64
	_ = v681
	var v684 float64
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v701 int64
	_ = v701
	var v702 int64
	_ = v702
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int64
	_ = v709
	var v710 int32
	_ = v710
	var v716 int64
	_ = v716
	var v717 int32
	_ = v717
	var v721 int64
	_ = v721
	var v722 int32
	_ = v722
	var v723 int64
	_ = v723
	var v724 float64
	_ = v724
	var v742 float64
	_ = v742
	var v745 float64
	_ = v745
	var v748 float64
	_ = v748
	var v751 float64
	_ = v751
	var v754 float64
	_ = v754
	var v757 float64
	_ = v757
	var v760 float64
	_ = v760
	var v763 float64
	_ = v763
	var v766 float64
	_ = v766
	var v776 float64
	_ = v776
	var v799 float64
	_ = v799
	var v814 float64
	_ = v814
	v7 = float64(0)
	v10 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(96)
	m.G0 = v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v41 = F_estimate_expression_value(m, l0, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return float64(0)
L2:
	;
	v45 = F_estimate_expression_value(m, l0, v38)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v32 + int32(96)
	return v814
L4:
	;
	v47 = F_exprType(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v49 = F_get_base_element_type(m, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v49 == int32(0) {
		v814 = float64(0.5)
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v53 = F_exprCollation(m, v45)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v55 = int32(0)
	if v45 == v55 {
		v112 = v55
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v132 = F_lookup_type_cache(m, v49, int32(1))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L19
	}
L10:
	;
	v69 = v45
	goto L11
L11:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	switch v87 - int32(27) {
	case 0:
		goto L13
	default:
		v112 = v69
		goto L9
	case 2:
		goto L14
	}
L12:
	;
	v112 = int32(0)
	goto L9
L13:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v99 != 0 {
		v69 = v99
		goto L11
	} else {
		goto L17
	}
L14:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v91 != int32(27) {
		v112 = v69
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 != int32(34) {
		v112 = v69
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	goto L12
L18:
	;
	v147 = int32(0)
	if l2|base.B2i32(v145|v144 == v147) == v147 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+52))
	if v134 == int32(0) {
		v144 = v10
		v145 = int32(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if v134 == v35 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v144 = int32(1)
	v145 = int32(0)
	goto L18
L22:
	;
	goto L23
L23:
	;
	v140 = F_get_negator(m, v35)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v132)+52))
	v144 = v10
	v145 = base.B2i32(v140 == v142)
	goto L18
L25:
	;
	v154 = m.G0
	v156 = v154 - int32(112)
	m.G0 = v156
	F_examine_variable(m, l0, v112, l3, v156+int32(80))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if l2 != 0 {
		goto L89
	} else {
		goto L90
	}
L28:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v156)+84))
	if v162 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	m.G0 = v156 + int32(112)
	if base.F64_ge(v331, float64(0)) != 0 {
		v814 = v331
		goto L3
	} else {
		goto L87
	}
L30:
	;
	v165 = float64(-1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v156)+88))
	if v166 == int32(0) {
		v331 = v165
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v172 != int32(7) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v156)+92))
	m.T0[v169].(func(*base.Module, int32))(m, v166)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v331 = v165
	goto L29
L35:
	;
	v175 = float64(-1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v156)+88))
	if v176 == int32(0) {
		v331 = v175
		goto L29
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+32)))
	if v182 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v156)+92))
	m.T0[v179].(func(*base.Module, int32))(m, v176)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v331 = v175
	goto L29
L40:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v156)+88))
	if v185 == int32(0) {
		v331 = v7
		goto L29
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v41)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v156)+72)) = v191
	v194 = F_lookup_type_cache(m, v49, int32(64))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v156)+92))
	m.T0[v188].(func(*base.Module, int32))(m, v185)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v331 = v7
	goto L29
L45:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v194)+108))
	if v196 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v199 = float64(-1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v156)+88))
	if v200 == int32(0) {
		v331 = v199
		goto L29
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v206 = v34&int32(1) ^ v144
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v156)+88))
	if v207 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v156)+92))
	m.T0[v203].(func(*base.Module, int32))(m, v200)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v331 = v199
	goto L29
L51:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v156)+88))
	if v316 != 0 {
		goto L78
	} else {
		goto L79
	}
L52:
	;
	if v206 != 0 {
		goto L75
	} else {
		goto L76
	}
L53:
	;
	v212 = F_statistic_proc_security_check(m, v156+int32(80), v196)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v212 == int32(0) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v156)+88))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+22)))
	v225 = F_get_attstatsslot(m, v156+int32(36), v216, int32(4), int32(0), int32(3))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	v304 = *(*float32)(unsafe.Add(mBase, uint32(v217+v218)+8))
	v315 = base.F64_mul(v300, base.F64_sub(float64(1), base.F64_promote_f32(v304)))
	goto L51
L57:
	;
	if v225 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v206 != 0 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	goto L60
L60:
	;
	if v206 != 0 {
		v300 = float64(0.005)
		goto L56
	} else {
		goto L73
	}
L61:
	;
	F_free_attstatsslot(m, v156)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L71
	}
L62:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v156)+48))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v156)+52))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v156)+56))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v156)+60))
	v278 = F_mcelem_array_contained_selec(m, v271, v272, v273, v274, v156+int32(72), int32(1), v269, v270, v194)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L70
	}
L63:
	;
	v258 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v156)+32)) = v258
	v260 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v156)+24)) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v156)+16)) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v156)+8)) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v156))) = v260
	v269 = v227
	v270 = v258
	goto L62
L64:
	;
	v227 = int32(0)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v156)+88))
	v232 = F_get_attstatsslot(m, v156, v228, int32(5), v227, int32(2))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+32)) = int32(0)
	v240 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v156)+24)) = v240
	*(*int64)(unsafe.Add(mBase, uint32(v156)+16)) = v240
	*(*int64)(unsafe.Add(mBase, uint32(v156)+8)) = v240
	*(*int64)(unsafe.Add(mBase, uint32(v156))) = v240
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v156)+48))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v156)+52))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v156)+56))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v156)+60))
	v256 = F_mcelem_array_contain_overlap_selec(m, v248, v249, v250, v251, v156+int32(72), int32(1), int32(2751), v194)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	if v232 == int32(0) {
		goto L63
	} else {
		goto L68
	}
L68:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v156)+24))
	v269 = v236
	v270 = v237
	goto L62
L69:
	;
	v282 = v256
	goto L61
L70:
	;
	v282 = v278
	goto L61
L71:
	;
	F_free_attstatsslot(m, v156+int32(36))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v300 = v282
	goto L56
L73:
	;
	v290 = int32(0)
	v298 = F_mcelem_array_contain_overlap_selec(m, v290, v290, v290, v290, v156+int32(72), int32(1), int32(2751), v194)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v300 = v298
	goto L56
L75:
	;
	v310 = float64(0.005)
	goto L77
L76:
	;
	v310 = float64(0.004999999888241291)
	goto L77
L77:
	;
	v315 = v310
	goto L51
L78:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v156)+92))
	m.T0[v317].(func(*base.Module, int32))(m, v316)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v144 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L80
L82:
	;
	v323 = v315
	goto L84
L83:
	;
	v323 = base.F64_sub(float64(1), v315)
	goto L84
L84:
	;
	if base.F64_lt(v323, float64(0)) != 0 {
		v331 = float64(0)
		goto L29
	} else {
		goto L85
	}
L85:
	;
	if base.F64_gt(v323, float64(1)) == int32(0) {
		v331 = v323
		goto L29
	} else {
		goto L86
	}
L86:
	;
	v331 = float64(1)
	goto L29
L87:
	;
	goto L27
L88:
	;
	if v353 == int32(0) {
		v814 = float64(0.5)
		goto L3
	} else {
		goto L94
	}
L89:
	;
	v349 = F_get_oprjoin(m, v35)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v351 = F_get_oprrest(m, v35)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	v353 = v349
	goto L88
L93:
	;
	v353 = v351
	goto L88
L94:
	;
	F_fmgr_info(m, v353, v32+int32(68))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	switch v353 - int32(101) {
	case 0, 4:
		v366 = int32(1)
		v367 = v145
		goto L96
	case 1, 5:
		goto L98
	default:
		v365 = v145
		goto L97
	}
L96:
	;
	if v112 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v366 = v144
	v367 = v365
	goto L96
L98:
	;
	v365 = int32(1)
	goto L97
L99:
	;
	v799 = float64(0)
	if base.F64_lt(v776, v799) != 0 {
		v814 = v799
		goto L3
	} else {
		goto L209
	}
L100:
	;
	v688 = F_palloc0(m, int32(16))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L198
	}
L101:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v370 != int32(35) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	if v370 != int32(7) {
		goto L100
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+20)))
	if v543 != 0 {
		goto L100
	} else {
		goto L151
	}
L105:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+32)))
	if v375 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v814 = float64(0)
	goto L3
L107:
	;
	v376 = *(*int64)(unsafe.Add(mBase, uint32(v112)+24))
	v378 = F_pg_detoast_datum(m, base.I32_wrap_i64(v376))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v34&int32(1) != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v386 = float64(0)
	goto L111
L110:
	;
	v383 = F_array_contains_nulls(m, v378)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v378)+12))
	F_get_typlenbyvalalign(m, v387, v32+int32(66), v32+int32(65), v32-int32(-64))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	if v383 != 0 {
		goto L106
	} else {
		goto L113
	}
L113:
	;
	v386 = float64(1)
	goto L111
L114:
	;
	v397 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+66)))
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+65)))
	v399 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32)+64)))
	F_deconstruct_array(m, v378, v397, v398, v399, v32+int32(56), v32+int32(52), v32+int32(60))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
	if v408 <= int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v34&int32(1) != 0 {
		goto L140
	} else {
		goto L141
	}
L117:
	;
	v508 = v386
	v509 = v386
	goto L116
L118:
	;
	goto L119
L119:
	;
	v413 = base.I64_extend_i32_u(v35)
	v414 = base.I64_extend_i32_u(l0)
	v424 = v386
	v425 = v386
	v429 = int32(0)
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v41
	v449 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+66)))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v32)+56))
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v450+v429<<(uint(int32(3))%32))))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455+v429))))
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+65)))
	v459 = F_makeConst(m, v49, int32(-1), v53, v449, v454, v457, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	v508 = v496
	v509 = v497
	goto L116
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v459
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v463
	v469 = F_list_make2_impl(m, v32+int32(16), v32+int32(12))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v471 = base.I64_extend_i32_u(v469)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if l2 != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v482 = base.F64_reinterpret_i64(v481)
	if v34&int32(1) != 0 {
		goto L131
	} else {
		goto L132
	}
L125:
	;
	v475 = F_FunctionCall5Coll(m, v32+int32(68), v472, v414, v413, v471, base.I64_extend16_s(base.I64_extend_i32_u(l4)), base.I64_extend_i32_u(l5))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v479 = F_FunctionCall4Coll(m, v32+int32(68), v472, v414, v413, v471, base.I64_extend_i32_s(l3))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L129
	}
L128:
	;
	v481 = v475
	goto L124
L129:
	;
	v481 = v479
	goto L124
L130:
	;
	v499 = v429 + int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
	if v499 < v500 {
		v424 = v496
		v425 = v497
		v429 = v499
		goto L120
	} else {
		goto L138
	}
L131:
	;
	if v366 != 0 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	v490 = base.F64_mul(v424, v482)
	if v367 == int32(0) {
		v496 = v490
		v497 = v425
		goto L130
	} else {
		goto L137
	}
L134:
	;
	v489 = base.F64_add(v425, v482)
	goto L136
L135:
	;
	v489 = v425
	goto L136
L136:
	;
	v496 = base.F64_sub(base.F64_add(v424, v482), base.F64_mul(v424, v482))
	v497 = v489
	goto L130
L137:
	;
	v496 = v490
	v497 = base.F64_add(v425, base.F64_add(v482, float64(-1)))
	goto L130
L138:
	;
	goto L121
L139:
	;
	if base.F64_le(v509, float64(1)) != 0 {
		goto L145
	} else {
		goto L146
	}
L140:
	;
	if v366 != 0 {
		goto L139
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if v367 == int32(0) {
		v776 = v508
		goto L99
	} else {
		goto L144
	}
L143:
	;
	v776 = v508
	goto L99
L144:
	;
	goto L139
L145:
	;
	v537 = v509
	goto L147
L146:
	;
	v537 = v508
	goto L147
L147:
	;
	if base.F64_ge(v509, float64(0)) != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v540 = v537
	goto L150
L149:
	;
	v540 = v508
	goto L150
L150:
	;
	v776 = v540
	goto L99
L151:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	F_get_typlenbyval(m, v544, v32+int32(60), v32+int32(56))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v554 = v34 & int32(1)
	if v554 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v555 = float64(0)
	goto L155
L154:
	;
	v555 = float64(1)
	goto L155
L155:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v556 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v814 = float64(0)
	goto L3
L157:
	;
	if v554 != 0 {
		goto L187
	} else {
		goto L188
	}
L158:
	;
	v654 = v555
	v655 = v555
	goto L157
L159:
	;
	goto L160
L160:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	if v559 <= int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v654 = v555
	v655 = v555
	goto L157
L162:
	;
	goto L163
L163:
	;
	v564 = base.I64_extend_i32_u(v35)
	v565 = base.I64_extend_i32_u(l0)
	v570 = v34 & int32(1)
	v571 = int32(0)
	v577 = v555
	v578 = v555
	goto L164
L164:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v556)+12))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v600+v571<<(uint(int32(2))%32))))
	if v570 != 0 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v654 = v642
	v655 = v643
	goto L157
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v32)+40)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v604
	v617 = F_list_make2_impl(m, v32+int32(24), v32+int32(20))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L170
	}
L167:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v604)))
	if v605 != int32(7) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+32)))
	if v608 != 0 {
		goto L156
	} else {
		goto L169
	}
L169:
	;
	goto L166
L170:
	;
	v619 = base.I64_extend_i32_u(v617)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if l2 != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v630 = base.F64_reinterpret_i64(v629)
	if v570 != 0 {
		goto L178
	} else {
		goto L179
	}
L172:
	;
	v623 = F_FunctionCall5Coll(m, v32+int32(68), v620, v565, v564, v619, base.I64_extend16_s(base.I64_extend_i32_u(l4)), base.I64_extend_i32_u(l5))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v627 = F_FunctionCall4Coll(m, v32+int32(68), v620, v565, v564, v619, base.I64_extend_i32_s(l3))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L176
	}
L175:
	;
	v629 = v623
	goto L171
L176:
	;
	v629 = v627
	goto L171
L177:
	;
	v645 = v571 + int32(1)
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	if v645 < v646 {
		v571 = v645
		v577 = v642
		v578 = v643
		goto L164
	} else {
		goto L185
	}
L178:
	;
	if v366 != 0 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	goto L180
L180:
	;
	v636 = base.F64_mul(v577, v630)
	if v367 == int32(0) {
		v642 = v636
		v643 = v578
		goto L177
	} else {
		goto L184
	}
L181:
	;
	v635 = base.F64_add(v578, v630)
	goto L183
L182:
	;
	v635 = v578
	goto L183
L183:
	;
	v642 = base.F64_sub(base.F64_add(v577, v630), base.F64_mul(v577, v630))
	v643 = v635
	goto L177
L184:
	;
	v642 = v636
	v643 = base.F64_add(v578, base.F64_add(v630, float64(-1)))
	goto L177
L185:
	;
	goto L165
L186:
	;
	if base.F64_le(v655, float64(1)) != 0 {
		goto L192
	} else {
		goto L193
	}
L187:
	;
	if v366 != 0 {
		goto L186
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	if v367 == int32(0) {
		v776 = v654
		goto L99
	} else {
		goto L191
	}
L190:
	;
	v776 = v654
	goto L99
L191:
	;
	goto L186
L192:
	;
	v681 = v655
	goto L194
L193:
	;
	v681 = v654
	goto L194
L194:
	;
	if base.F64_ge(v655, float64(0)) != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v684 = v681
	goto L197
L196:
	;
	v684 = v654
	goto L197
L197:
	;
	v776 = v684
	goto L99
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v688)+8)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v688)+4)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v688))) = int32(34)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v688)+12)) = v695
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v688
	v701 = base.I64_extend_i32_u(v35)
	v702 = base.I64_extend_i32_u(l0)
	v707 = F_list_make2_impl(m, v32+int32(8), v32+int32(4))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	v709 = base.I64_extend_i32_u(v707)
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if l2 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v724 = base.F64_reinterpret_i64(v723)
	if v34&int32(1) == int32(0) {
		goto L206
	} else {
		goto L207
	}
L201:
	;
	v716 = F_FunctionCall5Coll(m, v32+int32(68), v710, v702, v701, v709, base.I64_extend16_s(base.I64_extend_i32_u(l4)), base.I64_extend_i32_u(l5))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v721 = F_FunctionCall4Coll(m, v32+int32(68), v710, v702, v701, v709, base.I64_extend_i32_s(l3))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L205
	}
L204:
	;
	v723 = v716
	goto L200
L205:
	;
	v723 = v721
	goto L200
L206:
	;
	v776 = base.F64_mul(base.F64_mul(base.F64_mul(base.F64_mul(base.F64_mul(base.F64_mul(base.F64_mul(base.F64_mul(base.F64_mul(v724, v724), v724), v724), v724), v724), v724), v724), v724), v724)
	goto L99
L207:
	;
	goto L208
L208:
	;
	v742 = base.F64_add(base.F64_mul(v724, math.Float64frombits(uint64(0x8000000000000000))), base.F64_add(v724, float64(0)))
	v745 = base.F64_sub(base.F64_add(v742, v724), base.F64_mul(v742, v724))
	v748 = base.F64_sub(base.F64_add(v745, v724), base.F64_mul(v745, v724))
	v751 = base.F64_sub(base.F64_add(v748, v724), base.F64_mul(v748, v724))
	v754 = base.F64_sub(base.F64_add(v751, v724), base.F64_mul(v751, v724))
	v757 = base.F64_sub(base.F64_add(v754, v724), base.F64_mul(v754, v724))
	v760 = base.F64_sub(base.F64_add(v757, v724), base.F64_mul(v757, v724))
	v763 = base.F64_sub(base.F64_add(v760, v724), base.F64_mul(v760, v724))
	v766 = base.F64_sub(base.F64_add(v763, v724), base.F64_mul(v763, v724))
	v776 = base.F64_sub(base.F64_add(v766, v724), base.F64_mul(v766, v724))
	goto L99
L209:
	;
	if base.F64_gt(v776, float64(1)) == int32(0) {
		v814 = v776
		goto L3
	} else {
		goto L210
	}
L210:
	;
	v814 = float64(1)
	goto L3
}
func F_scalargtsel(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_scalarineqsel_wrapper(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_scram_SaltedPassword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v310 int32
	_ = v310
	var v332 int32
	_ = v332
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v396 int32
	_ = v396
	v20 = m.G0
	v22 = v20 - int32(80)
	m.G0 = v22
	v24 = F_strlen(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = int32(16777216)
	v27 = F_pg_hmac_create(m, l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v22 + int32(80)
	return v396
L2:
	;
	return int32(0)
L3:
	;
	if v27 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	goto L8
L5:
	;
	goto L6
L6:
	;
	v56 = F_pg_hmac_init(m, v27, l0, v24)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L21
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(_a_F_scram_SaltedPassword_0)
	v396 = int32(-1)
	goto L1
L8:
	;
	goto L7
L20:
	;
	if v27 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L21:
	;
	if v56 < int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if v27 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v74 < int32(0) {
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v74 = int32(-1)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v64 = F_pg_cryptohash_update(m, v63, l3, l4)
	mBase = m.M
	if int32(0) <= v64 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v74 = int32(0)
	goto L23
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = int32(2)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v71 = F_pg_cryptohash_error(m, v70)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v71
	v74 = int32(-1)
	goto L23
L30:
	;
	if v27 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v94 < int32(0) {
		goto L20
	} else {
		goto L38
	}
L32:
	;
	v94 = int32(-1)
	goto L31
L33:
	;
	goto L34
L34:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v84 = F_pg_cryptohash_update(m, v83, v22+int32(76), int32(4))
	mBase = m.M
	if int32(0) <= v84 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v94 = int32(0)
	goto L31
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = int32(2)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v91 = F_pg_cryptohash_error(m, v90)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v91
	v94 = int32(-1)
	goto L31
L38:
	;
	v97 = F_pg_hmac_final(m, v27, v22, l2)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L40
	}
L39:
	;
	v101 = int32(0)
	v102 = base.B2i32(l2 == v101)
	if v102 == v101 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if int32(0) <= v97 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	goto L20
L42:
	;
	base.MemoryCopy(m, l6, v22, l2)
	goto L44
L43:
	;
	goto L44
L44:
	;
	if int32(2) <= l5 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L20
L46:
	;
	v111 = l2 & int32(3)
	v129 = int32(1)
	goto L49
L47:
	;
	goto L48
L48:
	;
	F_pg_hmac_free(m, v27)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L2
	} else {
		goto L83
	}
L49:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_scram_SaltedPassword[0]))
	if v135 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L48
L51:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v138 = F_pg_hmac_init(m, v27, l0, v24)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	if v138 < int32(0) {
		goto L45
	} else {
		goto L56
	}
L56:
	;
	if v27 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v156 < int32(0) {
		goto L45
	} else {
		goto L64
	}
L58:
	;
	v156 = int32(-1)
	goto L57
L59:
	;
	goto L60
L60:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v146 = F_pg_cryptohash_update(m, v145, v22, l2)
	mBase = m.M
	if int32(0) <= v146 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v156 = int32(0)
	goto L57
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = int32(2)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v153 = F_pg_cryptohash_error(m, v152)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v153
	v156 = int32(-1)
	goto L57
L64:
	;
	v161 = F_pg_hmac_final(m, v27, v22+int32(32), l2)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	if v161 < int32(0) {
		goto L45
	} else {
		goto L66
	}
L66:
	;
	if l2 <= int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if v102 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L68:
	;
	v167 = int32(0)
	if base.B2i32(base.Ui32(l2) < base.Ui32(int32(4))) == v167 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v173 = v167
	v176 = v167
	goto L72
L70:
	;
	v238 = v167
	goto L71
L71:
	;
	v256 = v167
	v257 = v238
	goto L76
L72:
	;
	v191 = v173 + l6
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	v194 = v22 + int32(32)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v173))))
	v197 = v192 ^ v196
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v197)
	v200 = v173 | int32(1)
	v201 = l6 + v200
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200+v194))))
	v205 = v202 ^ v204
	*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v205)
	v208 = v173 | int32(2)
	v209 = l6 + v208
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v208))))
	v215 = v210 ^ v214
	*(*uint8)(unsafe.Add(mBase, uint32(v209))) = uint8(v215)
	v218 = v173 | int32(3)
	v219 = l6 + v218
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v218))))
	v225 = v220 ^ v224
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v225)
	v227 = int32(4)
	v228 = v173 + v227
	v230 = v176 + v227
	if v230 != l2&int32(2147483644) {
		v173 = v228
		v176 = v230
		goto L72
	} else {
		goto L74
	}
L73:
	;
	if v111 == int32(0) {
		goto L67
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	v238 = v228
	goto L71
L76:
	;
	v272 = v257 + l6
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(32)+v257))))
	v278 = v273 ^ v277
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v278)
	v280 = int32(1)
	v283 = v256 + v280
	if v283 != v111 {
		v256 = v283
		v257 = v257 + v280
		goto L76
	} else {
		goto L78
	}
L77:
	;
	goto L67
L78:
	;
	goto L77
L79:
	;
	base.MemoryCopy(m, v22, v22+int32(32), l2)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v310 = v129 + int32(1)
	if v310 != l5 {
		v129 = v310
		goto L49
	} else {
		goto L82
	}
L82:
	;
	goto L50
L83:
	;
	v396 = int32(0)
	goto L1
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v372
	F_pg_hmac_free(m, v27)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L2
	} else {
		goto L97
	}
L85:
	;
	v372 = int32(_a_F_scram_SaltedPassword_0)
	goto L84
L86:
	;
	goto L87
L87:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v357 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v369 = v357
	goto L90
L89:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v361 == int32(2) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v372 = v369
	goto L84
L91:
	;
	v364 = int32(_a_F_scram_SaltedPassword_1)
	goto L93
L92:
	;
	v364 = int32(_a_F_scram_SaltedPassword_2)
	goto L93
L93:
	;
	if v361 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v367 = int32(_a_F_scram_SaltedPassword_0)
	goto L96
L95:
	;
	v367 = v364
	goto L96
L96:
	;
	v369 = v367
	goto L90
L97:
	;
	v396 = int32(-1)
	goto L1
}
func F_scram_ServerKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	v6 = F_pg_hmac_create(m, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(_a_F_scram_ServerKey_0)
			return int32(-1)
		} else {
			v36 = F_pg_hmac_init(m, v6, l0, l2)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				if v36 < int32(0) {
					if v6 == int32(0) {
						v82 = int32(_a_F_scram_ServerKey_0)
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						if v67 != 0 {
							v79 = v67
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
							if v71 == int32(2) {
								v74 = int32(_a_F_scram_ServerKey_1)
							} else {
								v74 = int32(_a_F_scram_ServerKey_2)
							}
							if v71 == int32(1) {
								v77 = int32(_a_F_scram_ServerKey_0)
							} else {
								v77 = v74
							}
							v79 = v77
						}
						v82 = v79
					}
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v82
					F_pg_hmac_free(m, v6)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						return int32(-1)
					}
				} else {
					if v6 == int32(0) {
						v56 = int32(-1)
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						v46 = F_pg_cryptohash_update(m, v45, int32(_a_F_scram_ServerKey_3), int32(10))
						mBase = m.M
						if int32(0) <= v46 {
							v56 = int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(2)
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
							v53 = F_pg_cryptohash_error(m, v52)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v53
							v56 = int32(-1)
						}
					}
					if v56 < int32(0) {
						if v6 == int32(0) {
							v82 = int32(_a_F_scram_ServerKey_0)
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
							if v67 != 0 {
								v79 = v67
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
								if v71 == int32(2) {
									v74 = int32(_a_F_scram_ServerKey_1)
								} else {
									v74 = int32(_a_F_scram_ServerKey_2)
								}
								if v71 == int32(1) {
									v77 = int32(_a_F_scram_ServerKey_0)
								} else {
									v77 = v74
								}
								v79 = v77
							}
							v82 = v79
						}
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v82
						F_pg_hmac_free(m, v6)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							return int32(-1)
						}
					} else {
						v59 = F_pg_hmac_final(m, v6, l3, l2)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v59 {
								F_pg_hmac_free(m, v6)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							} else {
								if v6 == int32(0) {
									v82 = int32(_a_F_scram_ServerKey_0)
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
									if v67 != 0 {
										v79 = v67
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
										if v71 == int32(2) {
											v74 = int32(_a_F_scram_ServerKey_1)
										} else {
											v74 = int32(_a_F_scram_ServerKey_2)
										}
										if v71 == int32(1) {
											v77 = int32(_a_F_scram_ServerKey_0)
										} else {
											v77 = v74
										}
										v79 = v77
									}
									v82 = v79
								}
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v82
								F_pg_hmac_free(m, v6)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									return int32(-1)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_scram_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v276 int32
	_ = v276
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = F_palloc0(m, int32(196))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
	v18 = int32(_a_F_scram_init_0)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_scram_init[0])))
	if base.B2i32(v21 == v15)|base.B2i32(v21 != v24) != 0 {
		v42 = v21
		v43 = v24
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L67
	}
L4:
	;
	F_pg_cryptohash_free(m, v104)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L63
	}
L5:
	;
	if v42-v43 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	v27 = l1
	v28 = v18
	goto L8
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	if v32 == int32(0) {
		v42 = v32
		v43 = v31
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v42 = v32
	v43 = v31
	goto L6
L10:
	;
	v35 = int32(1)
	if v32 == v31 {
		v27 = v27 + v35
		v28 = v28 + v35
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v47 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v47)
	if l2 == v47 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L59
	}
L15:
	;
	m.G0 = v8 + int32(32)
	return v11
L16:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+364))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+12)) = int64(137438953475)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_scram_init[1]))
	v104 = F_pg_cryptohash_create(m, int32(3))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L29
	}
L17:
	;
	v51 = F_get_password_type(m, l2)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v51 == int32(2) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v67 = F_parse_scram_secret(m, l2, v11+int32(20), v11+int32(12), v11+int32(16), v11+int32(24), v11+int32(60), v11+int32(92))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v87
	v92 = F_psprintf(m, int32(_a_F_scram_init_1), v8+int32(16))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L28
	}
L22:
	;
	if v67 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v71 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v71 == int32(0) {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v76
	F_errmsg(m, int32(_a_F_scram_init_2), v8)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_scram_init_3), int32(290), int32(_a_F_scram_init_4))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L16
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v92
	goto L16
L29:
	;
	v106 = F_pg_cryptohash_init(m, v104)
	mBase = m.M
	if v106 < int32(0) {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v109 = F_strlen(m, v96)
	mBase = m.M
	v110 = F_pg_cryptohash_update(m, v104, v96, v109)
	mBase = m.M
	if v110 < int32(0) {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v114 = F_pg_cryptohash_update(m, v104, v100+int32(273), int32(32))
	mBase = m.M
	if v114 < int32(0) {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v119 = F_pg_cryptohash_final(m, v104, int32(_a_F_scram_init_5), int32(32))
	mBase = m.M
	if v119 < int32(0) {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_pg_cryptohash_free(m, v104)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v130 = base.I32_div_s(int32(18), int32(3))
	v132 = v130 << (uint(int32(2)) % 32)
	goto L35
L35:
	;
	v135 = F_palloc(m, v132+int32(1))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L40
L37:
	;
	if v248 < int32(0) {
		goto L3
	} else {
		goto L58
	}
L38:
	;
	if v132 != 0 {
		goto L55
	} else {
		goto L56
	}
L39:
	;
	if v132 < v190-v135+int32(4) {
		goto L38
	} else {
		goto L51
	}
L40:
	;
	v145 = int32(_a_F_scram_init_5)
	v146 = int32(0)
	v149 = v135
	v150 = int32(2)
	goto L43
L42:
	;
	v248 = v190 - v135
	goto L37
L43:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v156 = v152<<(uint(v150<<(uint(int32(3))%32))%32) | v146
	if int32(0) < v150 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v191 != int32(2) {
		goto L39
	} else {
		goto L50
	}
L45:
	;
	v189 = v156
	v190 = v149
	v191 = v150 - int32(1)
	goto L47
L46:
	;
	if v132 < v149-v135+int32(4) {
		goto L38
	} else {
		goto L48
	}
L47:
	;
	v193 = v145 + int32(1)
	if base.Ui32(v193) < base.Ui32(int32(_a_F_scram_init_6)) {
		v145 = v193
		v146 = v189
		v149 = v190
		v150 = v191
		goto L43
	} else {
		goto L49
	}
L48:
	;
	v165 = int32(63)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156&v165)+uint32(_c_F_scram_init[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+3)) = uint8(v167)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v156)>>(uint(int32(18))%32)))+uint32(_c_F_scram_init[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v171)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v156)>>(uint(int32(6))%32))&v165)+uint32(_c_F_scram_init[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+2)) = uint8(v177)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v156)>>(uint(int32(12))%32))&v165)+uint32(_c_F_scram_init[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)) = uint8(v183)
	v189 = int32(0)
	v190 = v149 + int32(4)
	v191 = int32(2)
	goto L47
L49:
	;
	goto L44
L50:
	;
	goto L42
L51:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v189)>>(uint(int32(18))%32)))+uint32(_c_F_scram_init[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v190))) = uint8(v211)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v189)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_scram_init[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)) = uint8(v217)
	if v191 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v189)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_scram_init[2]))))
	v227 = v226
	goto L54
L53:
	;
	v227 = int32(61)
	goto L54
L54:
	;
	v228 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+3)) = uint8(v228)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+2)) = uint8(v227)
	v248 = v190 + int32(4) - v135
	goto L37
L55:
	;
	base.MemoryFill(m, v135, int32(0), v132)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v248 = int32(-1)
	goto L37
L58:
	;
	v252 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v248))) = uint8(v252)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v135
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_scram_init[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v256
	v259 = v11 + int32(60)
	v260 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v259)+56)) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v259)+48)) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v259)+40)) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v259)+32)) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v259)+24)) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v259)+16)) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v259)+8)) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v259))) = v260
	v276 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+188)) = uint8(v276)
	goto L15
L59:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(_a_F_scram_init_7), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_scram_init_3), int32(265), int32(_a_F_scram_init_4))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errmsg_internal(m, int32(_a_F_scram_init_8), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_scram_init_3), int32(717), int32(_a_F_scram_init_9))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errmsg_internal(m, int32(_a_F_scram_init_8), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_scram_init_3), int32(726), int32(_a_F_scram_init_9))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_secure_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_ProcessClientWriteInterrupt(m, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = F_pgmem_send(m, v15, l1, l2)
	mBase = m.M
	if int32(0) <= v16 {
		v69 = v16
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	F_ProcessClientWriteInterrupt(m, int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L20
	}
L5:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v19 != 0 {
		v69 = v16
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v23 = v16
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_secure_write[0]))
	if v26 != int32(6) {
		v69 = v23
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v69 = v60
	goto L4
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_secure_write[1]))
	v31 = int32(0)
	F_ModifyWaitEvent(m, v30, v31, int32(4), v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_secure_write[1]))
	v41 = F_WaitEventSetWait(m, v37, int32(-1), v8, int32(1), int32(100663297))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v43&int32(16) != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	if v43&int32(1) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_secure_write[2]))
	v50 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v50
	v55 = base.AtomicRmwOr32(m, v50, int32(_a_F_secure_write_0), v50)
	goto L16
L14:
	;
	goto L15
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = F_pgmem_send(m, v59, l1, l2)
	mBase = m.M
	if int32(0) <= v60 {
		v69 = v60
		goto L4
	} else {
		goto L18
	}
L16:
	;
	F_ProcessClientWriteInterrupt(m, int32(1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v63 == int32(0) {
		v23 = v60
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L8
L20:
	;
	m.G0 = v8 + int32(16)
	return v69
L21:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(_a_F_secure_write_1), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_secure_write_2), int32(354), int32(_a_F_secure_write_3))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sendDir(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v265 int64
	_ = v265
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int64
	_ = v364
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v729 int32
	_ = v729
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v778 int32
	_ = v778
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v827 int32
	_ = v827
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v900 int32
	_ = v900
	var v907 int32
	_ = v907
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v930 int32
	_ = v930
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v982 int32
	_ = v982
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1060 int32
	_ = v1060
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1122 int32
	_ = v1122
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1199 int32
	_ = v1199
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1230 int32
	_ = v1230
	var v1268 int32
	_ = v1268
	var v1297 int32
	_ = v1297
	var v1348 int32
	_ = v1348
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1439 int32
	_ = v1439
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int64
	_ = v1446
	var v1449 int64
	_ = v1449
	var v1455 int64
	_ = v1455
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1630 int32
	_ = v1630
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1671 int64
	_ = v1671
	var v1672 int64
	_ = v1672
	var v1678 int32
	_ = v1678
	var v1684 int32
	_ = v1684
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1700 int32
	_ = v1700
	var v1704 int32
	_ = v1704
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1720 int32
	_ = v1720
	var v1725 int32
	_ = v1725
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1735 int32
	_ = v1735
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1747 int32
	_ = v1747
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1804 int32
	_ = v1804
	var v1809 int32
	_ = v1809
	var v1820 int32
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1843 int32
	_ = v1843
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1915 int32
	_ = v1915
	var v1931 int32
	_ = v1931
	var v1964 int64
	_ = v1964
	var v1967 int64
	_ = v1967
	var v1970 int64
	_ = v1970
	var v1981 int64
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2037 int32
	_ = v2037
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2059 int32
	_ = v2059
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2070 int32
	_ = v2070
	var v2074 int32
	_ = v2074
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2100 int32
	_ = v2100
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2159 int32
	_ = v2159
	var v2161 int32
	_ = v2161
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2172 int32
	_ = v2172
	var v2176 int32
	_ = v2176
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2202 int32
	_ = v2202
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2275 int32
	_ = v2275
	var v2279 int32
	_ = v2279
	var v2283 int32
	_ = v2283
	var v2287 int32
	_ = v2287
	var v2291 int32
	_ = v2291
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2302 int32
	_ = v2302
	var v2312 int32
	_ = v2312
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2358 int32
	_ = v2358
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2369 int32
	_ = v2369
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2383 int32
	_ = v2383
	var v2387 int32
	_ = v2387
	var v2435 int32
	_ = v2435
	var v2438 int32
	_ = v2438
	var v2442 int32
	_ = v2442
	var v2449 int32
	_ = v2449
	var v2451 int32
	_ = v2451
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2491 int32
	_ = v2491
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2511 int32
	_ = v2511
	var v2513 int32
	_ = v2513
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2524 int32
	_ = v2524
	var v2528 int32
	_ = v2528
	var v2532 int32
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2543 int32
	_ = v2543
	var v2546 int32
	_ = v2546
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2554 int32
	_ = v2554
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2563 int32
	_ = v2563
	var v2564 int32
	_ = v2564
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2571 int32
	_ = v2571
	var v2573 int32
	_ = v2573
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2588 int32
	_ = v2588
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2600 int32
	_ = v2600
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2606 int32
	_ = v2606
	var v2607 int32
	_ = v2607
	var v2608 int32
	_ = v2608
	var v2610 int32
	_ = v2610
	var v2611 int32
	_ = v2611
	var v2613 int32
	_ = v2613
	var v2615 int32
	_ = v2615
	var v2619 int32
	_ = v2619
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2626 int32
	_ = v2626
	var v2630 int32
	_ = v2630
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2645 int32
	_ = v2645
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2656 int32
	_ = v2656
	var v2660 int32
	_ = v2660
	var v2661 int32
	_ = v2661
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2681 int32
	_ = v2681
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2700 int32
	_ = v2700
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2718 int32
	_ = v2718
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2723 int32
	_ = v2723
	var v2725 int32
	_ = v2725
	var v2729 int32
	_ = v2729
	var v2733 int32
	_ = v2733
	var v2737 int32
	_ = v2737
	var v2741 int32
	_ = v2741
	var v2745 int32
	_ = v2745
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2756 int32
	_ = v2756
	var v2768 int32
	_ = v2768
	var v2808 int32
	_ = v2808
	var v2811 int32
	_ = v2811
	var v2814 int32
	_ = v2814
	var v2817 int32
	_ = v2817
	var v2818 int32
	_ = v2818
	var v2821 int32
	_ = v2821
	var v2822 int32
	_ = v2822
	var v2825 int32
	_ = v2825
	var v2832 int32
	_ = v2832
	var v2833 int32
	_ = v2833
	var v2839 int32
	_ = v2839
	var v2843 int32
	_ = v2843
	var v2890 int32
	_ = v2890
	var v2894 int32
	_ = v2894
	var v2896 int32
	_ = v2896
	var v2899 int32
	_ = v2899
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2905 int32
	_ = v2905
	var v2907 int32
	_ = v2907
	var v2909 int64
	_ = v2909
	var v2911 int32
	_ = v2911
	var v2914 int32
	_ = v2914
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2924 int64
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2932 int64
	_ = v2932
	var v2934 int64
	_ = v2934
	var v2935 int64
	_ = v2935
	var v2942 int32
	_ = v2942
	var v2945 int32
	_ = v2945
	var v2946 int32
	_ = v2946
	var v2947 int32
	_ = v2947
	var v2958 int32
	_ = v2958
	var v2968 int32
	_ = v2968
	var v2970 int32
	_ = v2970
	var v2974 int32
	_ = v2974
	var v2975 int32
	_ = v2975
	var v2976 int32
	_ = v2976
	var v2978 int32
	_ = v2978
	var v2980 int64
	_ = v2980
	var v2982 int32
	_ = v2982
	var v2985 int32
	_ = v2985
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2988 int32
	_ = v2988
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2995 int64
	_ = v2995
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v3003 int64
	_ = v3003
	var v3005 int64
	_ = v3005
	var v3006 int64
	_ = v3006
	var v3013 int32
	_ = v3013
	var v3016 int32
	_ = v3016
	var v3017 int32
	_ = v3017
	var v3018 int32
	_ = v3018
	var v3029 int32
	_ = v3029
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3050 int32
	_ = v3050
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3061 int32
	_ = v3061
	var v3062 int32
	_ = v3062
	var v3064 int32
	_ = v3064
	var v3079 int32
	_ = v3079
	var v3082 int32
	_ = v3082
	var v3116 int32
	_ = v3116
	var v3120 int32
	_ = v3120
	var v3125 int32
	_ = v3125
	var v3128 int32
	_ = v3128
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3136 int32
	_ = v3136
	var v3152 int32
	_ = v3152
	var v3171 int32
	_ = v3171
	var v3189 int32
	_ = v3189
	var v3201 int32
	_ = v3201
	var v3204 int32
	_ = v3204
	var v3206 int32
	_ = v3206
	var v3221 int32
	_ = v3221
	var v3224 int32
	_ = v3224
	var v3260 int32
	_ = v3260
	var v3272 int32
	_ = v3272
	var v3275 int32
	_ = v3275
	var v3277 int32
	_ = v3277
	var v3291 int32
	_ = v3291
	var v3326 int32
	_ = v3326
	var v3340 int32
	_ = v3340
	var v3384 int32
	_ = v3384
	var v3385 int32
	_ = v3385
	var v3391 int32
	_ = v3391
	var v3392 int32
	_ = v3392
	var v3413 int32
	_ = v3413
	var v3420 int32
	_ = v3420
	var v3446 int32
	_ = v3446
	var v3447 int32
	_ = v3447
	var v3450 int32
	_ = v3450
	var v3453 int32
	_ = v3453
	var v3456 int32
	_ = v3456
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3462 int32
	_ = v3462
	var v3481 int32
	_ = v3481
	var v3527 int32
	_ = v3527
	var v3530 int32
	_ = v3530
	var v3560 int32
	_ = v3560
	var v3561 int32
	_ = v3561
	var v3564 int32
	_ = v3564
	var v3567 int32
	_ = v3567
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3622 int32
	_ = v3622
	var v3624 int32
	_ = v3624
	var v3627 int32
	_ = v3627
	var v3640 int32
	_ = v3640
	var v3701 int32
	_ = v3701
	var v3732 int32
	_ = v3732
	var v3735 int32
	_ = v3735
	var v3740 int32
	_ = v3740
	var v3745 int32
	_ = v3745
	var v3748 int32
	_ = v3748
	var v3750 int32
	_ = v3750
	var v3752 int32
	_ = v3752
	var v3756 int32
	_ = v3756
	var v3762 int32
	_ = v3762
	var v3772 int32
	_ = v3772
	var v3777 int32
	_ = v3777
	var v3778 int32
	_ = v3778
	var v3783 int32
	_ = v3783
	var v3784 int32
	_ = v3784
	var v3786 int32
	_ = v3786
	var v3807 int32
	_ = v3807
	var v3810 int32
	_ = v3810
	var v3840 int32
	_ = v3840
	var v3841 int32
	_ = v3841
	var v3842 int32
	_ = v3842
	var v3843 int32
	_ = v3843
	var v3844 int32
	_ = v3844
	var v3845 int32
	_ = v3845
	var v3848 int64
	_ = v3848
	var v3862 int32
	_ = v3862
	var v3863 int32
	_ = v3863
	var v3873 int32
	_ = v3873
	var v3878 int32
	_ = v3878
	var v3923 int64
	_ = v3923
	var v3925 int32
	_ = v3925
	var v3926 int32
	_ = v3926
	var v3937 int32
	_ = v3937
	var v3946 int32
	_ = v3946
	var v3959 int32
	_ = v3959
	var v3971 int64
	_ = v3971
	var v3974 int32
	_ = v3974
	var v3976 int32
	_ = v3976
	var v3984 int32
	_ = v3984
	var v3987 int32
	_ = v3987
	var v3991 int32
	_ = v3991
	var v3995 int32
	_ = v3995
	var v4000 int32
	_ = v4000
	var v4004 int32
	_ = v4004
	var v4006 int32
	_ = v4006
	var v4014 int32
	_ = v4014
	var v4019 int32
	_ = v4019
	var v4023 int32
	_ = v4023
	var v4026 int32
	_ = v4026
	var v4034 int32
	_ = v4034
	var v4039 int32
	_ = v4039
	v10 = int32(0)
	v45 = int64(0)
	v47 = m.G0
	v49 = v47 - int32(_a_F_sendDir_0)
	m.G0 = v49
	if l8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v53 = F_palloc_mul(m, int32(4), int32(_a_F_sendDir_1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v58 = int32(0)
	goto L3
L3:
	;
	v61 = l1
	v63 = int32(0)
	goto L9
L4:
	;
	return int64(0)
L5:
	;
	v58 = v53
	goto L3
L6:
	;
	v306 = F_AllocateDir(m, l1)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L84
	}
L7:
	;
	v270 = int32(_a_F_sendDir_2)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[0])))
	if base.B2i32(v273 == int32(0))|base.B2i32(v273 != v276) != 0 {
		v294 = v273
		v295 = v276
		goto L72
	} else {
		goto L73
	}
L8:
	;
	if v63 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L9:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v64 != int32(47) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v61 = v61 + int32(1)
	v63 = v67
	goto L9
L12:
	;
	if v64 != 0 {
		v67 = v63
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v67 = v61
	goto L11
L15:
	;
	goto L8
L16:
	;
	v73 = v63 + int32(1)
	v74 = int32(_a_F_sendDir_3)
	v78 = m.G0
	v80 = v78 - int32(32)
	v81 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v80)+24)) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v80)+16)) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v80))) = v81
	v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[1])))
	if v89 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v158 = F_strlen(m, v73)
	mBase = m.M
	if v157 != v158 {
		goto L7
	} else {
		goto L36
	}
L18:
	;
	v157 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[2])))
	if v93 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v97 = v73
	goto L24
L22:
	;
	goto L23
L23:
	;
	v107 = v74
	v108 = v89
	goto L27
L24:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v103 == v89 {
		v97 = v97 + int32(1)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v157 = v97 - v73
	goto L17
L26:
	;
	goto L25
L27:
	;
	v115 = v80 + int32(base.Ui32(v108)>>(uint(int32(3))%32))&int32(28)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v117 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v116 | v117<<(uint(v108)%32)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	if v121 != 0 {
		v107 = v107 + v117
		v108 = v121
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v124 == int32(0) {
		v147 = v73
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v157 = v147 - v73
	goto L17
L31:
	;
	v128 = v73
	v129 = v124
	goto L32
L32:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(base.Ui32(v129)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v137)>>(uint(v129)%32))&int32(1) == int32(0) {
		v147 = v128
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v147 = v145
	goto L30
L34:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	v145 = v128 + int32(1)
	if v143 != 0 {
		v128 = v145
		v129 = v143
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v160 = int32(_a_F_sendDir_4)
	v161 = v63 - l1
	if v161 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v206 != 0 {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v206 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v167 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v168 = l1
	v169 = v160
	v170 = v161
	v171 = v167
	goto L45
L42:
	;
	v194 = v160
	v198 = int32(0)
	goto L43
L43:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	v206 = v198 - v199
	goto L37
L44:
	;
	v194 = v189
	v198 = v191
	goto L43
L45:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if base.B2i32(v171 != v173)|base.B2i32(v173 == int32(0)) != 0 {
		v189 = v169
		v191 = v171
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v189 = v183
	v191 = int32(0)
	goto L44
L47:
	;
	v179 = v170 - int32(1)
	if v179 == int32(0) {
		v189 = v169
		v191 = v171
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v182 = int32(1)
	v183 = v169 + v182
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	if v184 != 0 {
		v168 = v168 + v182
		v169 = v183
		v170 = v179
		v171 = v184
		goto L45
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	v207 = int32(1663)
	if base.Ui32(v161) < base.Ui32(int32(15)) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v265 = F_strtox_2(m, v73, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L70
L53:
	;
	v303 = v10
	v304 = v10
	v305 = v207
	goto L6
L54:
	;
	goto L55
L55:
	;
	v210 = int32(15)
	v211 = v63 - v210
	v212 = int32(_a_F_sendDir_5)
	goto L58
L56:
	;
	if v250-v251 != 0 {
		v303 = v10
		v304 = v10
		v305 = v207
		goto L6
	} else {
		goto L69
	}
L58:
	;
	goto L59
L59:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v219 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v220 = v211
	v221 = v212
	v222 = v210
	v223 = v219
	goto L64
L61:
	;
	v246 = v212
	v250 = int32(0)
	goto L62
L62:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	goto L56
L63:
	;
	v246 = v241
	v250 = v243
	goto L62
L64:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if base.B2i32(v223 != v225)|base.B2i32(v225 == int32(0)) != 0 {
		v241 = v221
		v243 = v223
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v241 = v235
	v243 = int32(0)
	goto L63
L66:
	;
	v231 = v222 - int32(1)
	if v231 == int32(0) {
		v241 = v221
		v243 = v223
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v234 = int32(1)
	v235 = v221 + v234
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+1)))
	if v236 != 0 {
		v220 = v220 + v234
		v221 = v235
		v222 = v231
		v223 = v236
		goto L64
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	goto L52
L70:
	;
	v303 = int32(1)
	v304 = base.I32_wrap_i64(v265)
	v305 = int32(1663)
	goto L6
L71:
	;
	if v296 != 0 {
		goto L78
	} else {
		goto L79
	}
L72:
	;
	v296 = v294 - v295
	goto L71
L73:
	;
	v279 = l1
	v280 = v270
	goto L74
L74:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	if v284 == int32(0) {
		v294 = v284
		v295 = v283
		goto L72
	} else {
		goto L76
	}
L75:
	;
	v294 = v284
	v295 = v283
	goto L72
L76:
	;
	v287 = int32(1)
	if v284 == v283 {
		v279 = v279 + v287
		v280 = v280 + v287
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v297 = int32(1663)
	goto L80
L79:
	;
	v297 = int32(1664)
	goto L80
L80:
	;
	v303 = base.B2i32(v296 == int32(0))
	v304 = v10
	v305 = v297
	goto L6
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4023 = m.ExcPending
	if v4023 != 0 {
		goto L4
	} else {
		goto L701
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4004 = m.ExcPending
	if v4004 != 0 {
		goto L4
	} else {
		goto L697
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3984 = m.ExcPending
	if v3984 != 0 {
		goto L4
	} else {
		goto L692
	}
L84:
	;
	v308 = F_ReadDir(m, v306, l1)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	if v308 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v320 = l0
	v321 = l1
	v322 = l2
	v323 = l3
	v324 = l4
	v325 = l5
	v326 = l6
	v327 = l7
	v328 = l8
	v329 = v308
	v330 = v49
	v339 = v58
	v341 = l2 + v49 + int32(2369)
	v347 = v303
	v348 = v304
	v350 = v305
	v352 = v306
	v360 = l1 + l2 + int32(1)
	v361 = v49 + int32(2368) | int32(2)
	v364 = v45
	goto L89
L87:
	;
	v3937 = v49
	v3946 = v58
	v3959 = v306
	v3971 = v45
	goto L88
L88:
	;
	if v3946 != 0 {
		goto L687
	} else {
		goto L688
	}
L89:
	;
	v366 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2268)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2264)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2260)) = v366
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+19)))
	if v372 != int32(46) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v3937 = v330
	v3946 = v339
	v3959 = v352
	v3971 = v3923
	goto L88
L91:
	;
	v3925 = F_ReadDir(m, v352, v321)
	mBase = m.M
	v3926 = m.ExcPending
	if v3926 != 0 {
		goto L4
	} else {
		goto L685
	}
L92:
	;
	v385 = v329 + int32(19)
	v386 = int32(_a_F_sendDir_6)
	goto L99
L93:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+20)))
	if v375 == int32(0) {
		v3923 = v364
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+20)))
	if v378 != int32(46) {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+21)))
	if v381 == int32(0) {
		v3923 = v364
		goto L91
	} else {
		goto L96
	}
L96:
	;
	goto L92
L97:
	;
	if v424-v425 == int32(0) {
		v3923 = v364
		goto L91
	} else {
		goto L110
	}
L99:
	;
	goto L100
L100:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v393 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v394 = v385
	v395 = v386
	v396 = int32(9)
	v397 = v393
	goto L105
L102:
	;
	v420 = v386
	v424 = int32(0)
	goto L103
L103:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	goto L97
L104:
	;
	v420 = v415
	v424 = v417
	goto L103
L105:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	if base.B2i32(v397 != v399)|base.B2i32(v399 == int32(0)) != 0 {
		v415 = v395
		v417 = v397
		goto L104
	} else {
		goto L107
	}
L106:
	;
	v415 = v409
	v417 = int32(0)
	goto L104
L107:
	;
	v405 = v396 - int32(1)
	if v405 == int32(0) {
		v415 = v395
		v417 = v397
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v408 = int32(1)
	v409 = v395 + v408
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+1)))
	if v410 != 0 {
		v394 = v394 + v408
		v395 = v409
		v396 = v405
		v397 = v410
		goto L105
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	v435 = int32(_a_F_sendDir_7)
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v441 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[3])))
	if base.B2i32(v438 == int32(0))|base.B2i32(v438 != v441) != 0 {
		v459 = v438
		v460 = v441
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v459-v460 == int32(0) {
		v3923 = v364
		goto L91
	} else {
		goto L118
	}
L112:
	;
	goto L111
L113:
	;
	v444 = v385
	v445 = v435
	goto L114
L114:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+1)))
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+1)))
	if v449 == int32(0) {
		v459 = v449
		v460 = v448
		goto L112
	} else {
		goto L116
	}
L115:
	;
	v459 = v449
	v460 = v448
	goto L112
L116:
	;
	v452 = int32(1)
	if v449 == v448 {
		v444 = v444 + v452
		v445 = v445 + v452
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[4]))
	if v465 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[5])))
	if v470 == int32(1) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L121
L123:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[6])))
	if v480 != v482 {
		goto L83
	} else {
		goto L127
	}
L124:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[7]))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+308))
	v478 = base.B2i32(v476 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[5])) = uint8(v478)
	v480 = v478
	goto L126
L125:
	;
	v480 = int32(0)
	goto L126
L126:
	;
	goto L123
L127:
	;
	v484 = int32(_a_F_sendDir_8)
	goto L133
L128:
	;
	if v348 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L129:
	;
	v1041 = int32(1)
	goto L128
L130:
	;
	v1026 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L4
	} else {
		goto L270
	}
L131:
	;
	if v522-v523 == int32(0) {
		goto L130
	} else {
		goto L144
	}
L133:
	;
	goto L134
L134:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v491 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v492 = v385
	v493 = v484
	v494 = int32(25)
	v495 = v491
	goto L139
L136:
	;
	v518 = v484
	v522 = int32(0)
	goto L137
L137:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	goto L131
L138:
	;
	v518 = v513
	v522 = v515
	goto L137
L139:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
	if base.B2i32(v495 != v497)|base.B2i32(v497 == int32(0)) != 0 {
		v513 = v493
		v515 = v495
		goto L138
	} else {
		goto L141
	}
L140:
	;
	v513 = v507
	v515 = int32(0)
	goto L138
L141:
	;
	v503 = v494 - int32(1)
	if v503 == int32(0) {
		v513 = v493
		v515 = v495
		goto L138
	} else {
		goto L142
	}
L142:
	;
	v506 = int32(1)
	v507 = v493 + v506
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+1)))
	if v508 != 0 {
		v492 = v492 + v506
		v493 = v507
		v494 = v503
		v495 = v508
		goto L139
	} else {
		goto L143
	}
L143:
	;
	goto L140
L144:
	;
	v533 = int32(_a_F_sendDir_9)
	goto L147
L145:
	;
	if v571-v572 == int32(0) {
		goto L130
	} else {
		goto L158
	}
L147:
	;
	goto L148
L148:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v540 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v541 = v385
	v542 = v533
	v543 = int32(21)
	v544 = v540
	goto L153
L150:
	;
	v567 = v533
	v571 = int32(0)
	goto L151
L151:
	;
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567))))
	goto L145
L152:
	;
	v567 = v562
	v571 = v564
	goto L151
L153:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542))))
	if base.B2i32(v544 != v546)|base.B2i32(v546 == int32(0)) != 0 {
		v562 = v542
		v564 = v544
		goto L152
	} else {
		goto L155
	}
L154:
	;
	v562 = v556
	v564 = int32(0)
	goto L152
L155:
	;
	v552 = v543 - int32(1)
	if v552 == int32(0) {
		v562 = v542
		v564 = v544
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v555 = int32(1)
	v556 = v542 + v555
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+1)))
	if v557 != 0 {
		v541 = v541 + v555
		v542 = v556
		v543 = v552
		v544 = v557
		goto L153
	} else {
		goto L157
	}
L157:
	;
	goto L154
L158:
	;
	v582 = int32(_a_F_sendDir_10)
	goto L161
L159:
	;
	if v620-v621 == int32(0) {
		goto L130
	} else {
		goto L172
	}
L161:
	;
	goto L162
L162:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v589 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v590 = v385
	v591 = v582
	v592 = int32(16)
	v593 = v589
	goto L167
L164:
	;
	v616 = v582
	v620 = int32(0)
	goto L165
L165:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616))))
	goto L159
L166:
	;
	v616 = v611
	v620 = v613
	goto L165
L167:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591))))
	if base.B2i32(v593 != v595)|base.B2i32(v595 == int32(0)) != 0 {
		v611 = v591
		v613 = v593
		goto L166
	} else {
		goto L169
	}
L168:
	;
	v611 = v605
	v613 = int32(0)
	goto L166
L169:
	;
	v601 = v592 - int32(1)
	if v601 == int32(0) {
		v611 = v591
		v613 = v593
		goto L166
	} else {
		goto L170
	}
L170:
	;
	v604 = int32(1)
	v605 = v591 + v604
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590)+1)))
	if v606 != 0 {
		v590 = v590 + v604
		v591 = v605
		v592 = v601
		v593 = v606
		goto L167
	} else {
		goto L171
	}
L171:
	;
	goto L168
L172:
	;
	v631 = int32(_a_F_sendDir_11)
	goto L175
L173:
	;
	if v669-v670 == int32(0) {
		goto L130
	} else {
		goto L186
	}
L175:
	;
	goto L176
L176:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v638 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v639 = v385
	v640 = v631
	v641 = int32(13)
	v642 = v638
	goto L181
L178:
	;
	v665 = v631
	v669 = int32(0)
	goto L179
L179:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	goto L173
L180:
	;
	v665 = v660
	v669 = v662
	goto L179
L181:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	if base.B2i32(v642 != v644)|base.B2i32(v644 == int32(0)) != 0 {
		v660 = v640
		v662 = v642
		goto L180
	} else {
		goto L183
	}
L182:
	;
	v660 = v654
	v662 = int32(0)
	goto L180
L183:
	;
	v650 = v641 - int32(1)
	if v650 == int32(0) {
		v660 = v640
		v662 = v642
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v653 = int32(1)
	v654 = v640 + v653
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639)+1)))
	if v655 != 0 {
		v639 = v639 + v653
		v640 = v654
		v641 = v650
		v642 = v655
		goto L181
	} else {
		goto L185
	}
L185:
	;
	goto L182
L186:
	;
	v680 = int32(_a_F_sendDir_12)
	goto L189
L187:
	;
	if v718-v719 == int32(0) {
		goto L130
	} else {
		goto L200
	}
L189:
	;
	goto L190
L190:
	;
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v687 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v688 = v385
	v689 = v680
	v690 = int32(15)
	v691 = v687
	goto L195
L192:
	;
	v714 = v680
	v718 = int32(0)
	goto L193
L193:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714))))
	goto L187
L194:
	;
	v714 = v709
	v718 = v711
	goto L193
L195:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689))))
	if base.B2i32(v691 != v693)|base.B2i32(v693 == int32(0)) != 0 {
		v709 = v689
		v711 = v691
		goto L194
	} else {
		goto L197
	}
L196:
	;
	v709 = v703
	v711 = int32(0)
	goto L194
L197:
	;
	v699 = v690 - int32(1)
	if v699 == int32(0) {
		v709 = v689
		v711 = v691
		goto L194
	} else {
		goto L198
	}
L198:
	;
	v702 = int32(1)
	v703 = v689 + v702
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688)+1)))
	if v704 != 0 {
		v688 = v688 + v702
		v689 = v703
		v690 = v699
		v691 = v704
		goto L195
	} else {
		goto L199
	}
L199:
	;
	goto L196
L200:
	;
	v729 = int32(_a_F_sendDir_13)
	goto L203
L201:
	;
	if v767-v768 == int32(0) {
		goto L130
	} else {
		goto L214
	}
L203:
	;
	goto L204
L204:
	;
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v736 != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v737 = v385
	v738 = v729
	v739 = int32(16)
	v740 = v736
	goto L209
L206:
	;
	v763 = v729
	v767 = int32(0)
	goto L207
L207:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v763))))
	goto L201
L208:
	;
	v763 = v758
	v767 = v760
	goto L207
L209:
	;
	v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738))))
	if base.B2i32(v740 != v742)|base.B2i32(v742 == int32(0)) != 0 {
		v758 = v738
		v760 = v740
		goto L208
	} else {
		goto L211
	}
L210:
	;
	v758 = v752
	v760 = int32(0)
	goto L208
L211:
	;
	v748 = v739 - int32(1)
	if v748 == int32(0) {
		v758 = v738
		v760 = v740
		goto L208
	} else {
		goto L212
	}
L212:
	;
	v751 = int32(1)
	v752 = v738 + v751
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737)+1)))
	if v753 != 0 {
		v737 = v737 + v751
		v738 = v752
		v739 = v748
		v740 = v753
		goto L209
	} else {
		goto L213
	}
L213:
	;
	goto L210
L214:
	;
	v778 = int32(_a_F_sendDir_14)
	goto L217
L215:
	;
	if v816-v817 == int32(0) {
		goto L130
	} else {
		goto L228
	}
L217:
	;
	goto L218
L218:
	;
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v785 != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v786 = v385
	v787 = v778
	v788 = int32(15)
	v789 = v785
	goto L223
L220:
	;
	v812 = v778
	v816 = int32(0)
	goto L221
L221:
	;
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812))))
	goto L215
L222:
	;
	v812 = v807
	v816 = v809
	goto L221
L223:
	;
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v787))))
	if base.B2i32(v789 != v791)|base.B2i32(v791 == int32(0)) != 0 {
		v807 = v787
		v809 = v789
		goto L222
	} else {
		goto L225
	}
L224:
	;
	v807 = v801
	v809 = int32(0)
	goto L222
L225:
	;
	v797 = v788 - int32(1)
	if v797 == int32(0) {
		v807 = v787
		v809 = v789
		goto L222
	} else {
		goto L226
	}
L226:
	;
	v800 = int32(1)
	v801 = v787 + v800
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786)+1)))
	if v802 != 0 {
		v786 = v786 + v800
		v787 = v801
		v788 = v797
		v789 = v802
		goto L223
	} else {
		goto L227
	}
L227:
	;
	goto L224
L228:
	;
	v827 = int32(_a_F_sendDir_15)
	goto L231
L229:
	;
	if v865-v866 == int32(0) {
		goto L130
	} else {
		goto L242
	}
L231:
	;
	goto L232
L232:
	;
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v834 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v835 = v385
	v836 = v827
	v837 = int32(16)
	v838 = v834
	goto L237
L234:
	;
	v861 = v827
	v865 = int32(0)
	goto L235
L235:
	;
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861))))
	goto L229
L236:
	;
	v861 = v856
	v865 = v858
	goto L235
L237:
	;
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836))))
	if base.B2i32(v838 != v840)|base.B2i32(v840 == int32(0)) != 0 {
		v856 = v836
		v858 = v838
		goto L236
	} else {
		goto L239
	}
L238:
	;
	v856 = v850
	v858 = int32(0)
	goto L236
L239:
	;
	v846 = v837 - int32(1)
	if v846 == int32(0) {
		v856 = v836
		v858 = v838
		goto L236
	} else {
		goto L240
	}
L240:
	;
	v849 = int32(1)
	v850 = v836 + v849
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835)+1)))
	if v851 != 0 {
		v835 = v835 + v849
		v836 = v850
		v837 = v846
		v838 = v851
		goto L237
	} else {
		goto L241
	}
L241:
	;
	goto L238
L242:
	;
	v876 = int32(0)
	if v347 == v876 {
		v1041 = v876
		goto L128
	} else {
		goto L243
	}
L243:
	;
	v880 = v330 + int32(2268)
	v882 = v330 + int32(2264)
	v884 = v330 + int32(2260)
	v885 = int32(0)
	v890 = m.G0
	v892 = v890 - int32(16)
	m.G0 = v892
	*(*int32)(unsafe.Add(mBase, uint32(v880))) = v885
	*(*int32)(unsafe.Add(mBase, uint32(v882))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v884))) = v885
	v900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if base.Ui32((v900-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v982 = v885
		goto L245
	} else {
		goto L246
	}
L244:
	;
	if v982 == int32(0) {
		v1041 = v982
		goto L128
	} else {
		goto L261
	}
L245:
	;
	m.G0 = v892 + int32(16)
	goto L244
L246:
	;
	v907 = int32(_a_F_sendDir_16)
	*(*int32)(unsafe.Add(mBase, _c_F_sendDir[8])) = int32(0)
	v913 = F_strtoul(m, v385, v892+int32(8), int32(10))
	mBase = m.M
	v915 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[8]))
	if v915 != 0 {
		v982 = v885
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v892)+8))
	if base.B2i32(v913 == int32(0))|base.B2i32(v385 == v918) != 0 {
		v982 = v885
		goto L245
	} else {
		goto L248
	}
L248:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v918))))
	if v921 != int32(95) {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	if v937&int32(255) == int32(46) {
		goto L254
	} else {
		goto L255
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v892)+12)) = int32(0)
	v937 = v921
	v938 = v918
	goto L249
L251:
	;
	goto L252
L252:
	;
	v930 = F_forkname_chars(m, v918+int32(1), v892+int32(12))
	mBase = m.M
	if v930 <= int32(0) {
		v982 = v885
		goto L245
	} else {
		goto L253
	}
L253:
	;
	v935 = v930 + v918 + int32(1)
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935))))
	v937 = v936
	v938 = v935
	goto L249
L254:
	;
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v938)+1)))
	if base.Ui32((v943-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v982 = v885
		goto L245
	} else {
		goto L257
	}
L255:
	;
	v969 = v885
	v970 = v937
	goto L256
L256:
	;
	if v970&int32(255) != 0 {
		v982 = v885
		goto L245
	} else {
		goto L260
	}
L257:
	;
	v950 = int32(_a_F_sendDir_16)
	*(*int32)(unsafe.Add(mBase, _c_F_sendDir[8])) = int32(0)
	v954 = v938 + int32(1)
	v958 = F_strtoul(m, v954, v892+int32(8), int32(10))
	mBase = m.M
	v960 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[8]))
	if v960 != 0 {
		v982 = v885
		goto L245
	} else {
		goto L258
	}
L258:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v892)+8))
	if base.B2i32(v958 == int32(0))|base.B2i32(v954 == v963) != 0 {
		v982 = v885
		goto L245
	} else {
		goto L259
	}
L259:
	;
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963))))
	v969 = v958
	v970 = v966
	goto L256
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v880))) = v913
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v892)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v882))) = v974
	*(*int32)(unsafe.Add(mBase, uint32(v884))) = v969
	v982 = int32(1)
	goto L245
L261:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2264))
	if v988 == int32(3) {
		v1041 = v982
		goto L128
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+176)) = v321
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2268))
	*(*int32)(unsafe.Add(mBase, uint32(v330)+180)) = v992
	v995 = v330 + int32(192)
	v1000 = F_pg_snprintf(m, v995, int32(1024), int32(_a_F_sendDir_17), v330+int32(176))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L4
	} else {
		goto L263
	}
L263:
	;
	v1006 = F___fstatat(m, int32(-100), v995, v330+int32(2272), int32(256))
	mBase = m.M
	goto L264
L264:
	;
	if v1006 != 0 {
		goto L129
	} else {
		goto L265
	}
L265:
	;
	v1009 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L4
	} else {
		goto L266
	}
L266:
	;
	if v1009 == int32(0) {
		v3923 = v364
		goto L91
	} else {
		goto L267
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+16)) = v385
	F_errmsg_internal(m, int32(_a_F_sendDir_18), v330+int32(16))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L4
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1332), int32(_a_F_sendDir_20))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L4
	} else {
		goto L269
	}
L269:
	;
	v3923 = v364
	goto L91
L270:
	;
	if v1026 == int32(0) {
		v3923 = v364
		goto L91
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330))) = v385
	F_errmsg_internal(m, int32(_a_F_sendDir_21), v330)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L4
	} else {
		goto L272
	}
L272:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1297), int32(_a_F_sendDir_20))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	v3923 = v364
	goto L91
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+148)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v330)+144)) = v321
	v1439 = v330 + int32(2368)
	v1444 = F_pg_snprintf(m, v1439, int32(2048), int32(_a_F_sendDir_22), v330+int32(144))
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L4
	} else {
		goto L316
	}
L275:
	;
	v1044 = int32(0)
	v1045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v1045 != int32(116) {
		v1348 = v1044
		goto L276
	} else {
		goto L277
	}
L276:
	;
	if v1348 == int32(0) {
		goto L274
	} else {
		goto L311
	}
L277:
	;
	v1060 = int32(1)
	goto L278
L278:
	;
	v1096 = v1060 + int32(1)
	v1098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1060+v385))))
	if base.Ui32((v1098-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v1060 = v1096
		goto L278
	} else {
		goto L280
	}
L279:
	;
	if base.B2i32(v1060 == int32(1))|base.B2i32(v1098 != int32(95)) != 0 {
		v1348 = v1044
		goto L276
	} else {
		goto L281
	}
L280:
	;
	goto L279
L281:
	;
	v1122 = v1096
	goto L282
L282:
	;
	v1157 = v1122 + int32(1)
	v1158 = v1122 + v385
	v1159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1158))))
	if base.Ui32((v1159-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v1122 = v1157
		goto L282
	} else {
		goto L284
	}
L283:
	;
	if v1122 == v1096 {
		v1348 = v1044
		goto L276
	} else {
		goto L285
	}
L284:
	;
	goto L283
L285:
	;
	if v1159 == int32(95) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1170 = v1158 + int32(1)
	v1174 = int32(3)
	v1177 = F_strncmp(m, int32(_a_F_sendDir_23), v1170, v1174)
	mBase = m.M
	if v1177 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L287:
	;
	v1213 = v1122
	v1214 = v1159
	goto L288
L288:
	;
	if v1214 == int32(46) {
		goto L304
	} else {
		goto L305
	}
L289:
	;
	if v1206 <= int32(0) {
		v1348 = v1044
		goto L276
	} else {
		goto L303
	}
L290:
	;
	goto L289
L292:
	;
	v1206 = v1199
	goto L290
L293:
	;
	v1199 = v1174
	goto L292
L294:
	;
	goto L295
L295:
	;
	v1181 = int32(2)
	v1185 = F_strncmp(m, int32(_a_F_sendDir_24), v1170, v1181)
	mBase = m.M
	if v1185 == int32(0) {
		v1199 = v1181
		goto L292
	} else {
		goto L296
	}
L296:
	;
	v1188 = int32(4)
	v1191 = F_strncmp(m, int32(_a_F_sendDir_25), v1170, v1188)
	mBase = m.M
	if v1191 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	goto L300
L298:
	;
	goto L299
L299:
	;
	v1206 = int32(0)
	goto L290
L300:
	;
	v1206 = v1188
	goto L290
L303:
	;
	v1210 = v1206 + v1157
	v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385+v1210))))
	v1213 = v1210
	v1214 = v1212
	goto L288
L304:
	;
	v1230 = int32(1)
	goto L307
L305:
	;
	v1297 = v1214
	goto L306
L306:
	;
	v1348 = base.B2i32(v1297 == int32(0))
	goto L276
L307:
	;
	v1268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1230+(v1213+v385)))))
	if base.Ui32((v1268-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v1230 = v1230 + int32(1)
		goto L307
	} else {
		goto L309
	}
L308:
	;
	if v1230 < int32(2) {
		v1348 = v1044
		goto L276
	} else {
		goto L310
	}
L309:
	;
	goto L308
L310:
	;
	v1297 = v1268
	goto L306
L311:
	;
	v1375 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L4
	} else {
		goto L312
	}
L312:
	;
	if v1375 == int32(0) {
		v3923 = v364
		goto L91
	} else {
		goto L313
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+160)) = v385
	F_errmsg_internal(m, int32(_a_F_sendDir_26), v330+int32(160))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L4
	} else {
		goto L314
	}
L314:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1343), int32(_a_F_sendDir_20))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L4
	} else {
		goto L315
	}
L315:
	;
	v3923 = v364
	goto L91
L316:
	;
	v1446 = *(*int64)(unsafe.Add(mBase, uint32(v330)+2368))
	v1449 = *(*int64)(unsafe.Add(mBase, uint32(v330)+2376))
	v1455 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v330+int32(2384)))))
	if v1446^int64(7809631459536744238)|(v1449^int64(8389765628430872623))|(v1455^int64(7106418)) == int64(0) {
		v3923 = v364
		goto L91
	} else {
		goto L317
	}
L317:
	;
	v1465 = F___fstatat(m, int32(-100), v1439, v330+int32(2272), int32(256))
	mBase = m.M
	goto L320
L318:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2276))
	v1760 = v1758 & int32(_a_F_sendDir_27)
	v1761 = int32(_a_F_sendDir_28)
	v1764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	v1767 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[9])))
	if base.B2i32(v1764 == int32(0))|base.B2i32(v1764 != v1767) != 0 {
		v1785 = v1764
		v1786 = v1767
		goto L403
	} else {
		goto L404
	}
L319:
	;
	v1728 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L4
	} else {
		goto L392
	}
L320:
	;
	if v1465 == int32(0) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1468 = int32(_a_F_sendDir_29)
	v1471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v1474 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[10])))
	if base.B2i32(v1471 == int32(0))|base.B2i32(v1471 != v1474) != 0 {
		v1492 = v1471
		v1493 = v1474
		goto L325
	} else {
		goto L326
	}
L322:
	;
	goto L323
L323:
	;
	v1704 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[8]))
	if v1704 == int32(44) {
		v3923 = v364
		goto L91
	} else {
		goto L387
	}
L324:
	;
	if v1492-v1493 == int32(0) {
		goto L319
	} else {
		goto L331
	}
L325:
	;
	goto L324
L326:
	;
	v1477 = v385
	v1478 = v1468
	goto L327
L327:
	;
	v1481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1478)+1)))
	v1482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1477)+1)))
	if v1482 == int32(0) {
		v1492 = v1482
		v1493 = v1481
		goto L325
	} else {
		goto L329
	}
L328:
	;
	v1492 = v1482
	v1493 = v1481
	goto L325
L329:
	;
	v1485 = int32(1)
	if v1482 == v1481 {
		v1477 = v1477 + v1485
		v1478 = v1478 + v1485
		goto L327
	} else {
		goto L330
	}
L330:
	;
	goto L328
L331:
	;
	v1497 = int32(_a_F_sendDir_30)
	v1500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v1503 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[11])))
	if base.B2i32(v1500 == int32(0))|base.B2i32(v1500 != v1503) != 0 {
		v1521 = v1500
		v1522 = v1503
		goto L333
	} else {
		goto L334
	}
L332:
	;
	if v1521-v1522 == int32(0) {
		goto L319
	} else {
		goto L339
	}
L333:
	;
	goto L332
L334:
	;
	v1506 = v385
	v1507 = v1497
	goto L335
L335:
	;
	v1510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1507)+1)))
	v1511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1506)+1)))
	if v1511 == int32(0) {
		v1521 = v1511
		v1522 = v1510
		goto L333
	} else {
		goto L337
	}
L336:
	;
	v1521 = v1511
	v1522 = v1510
	goto L333
L337:
	;
	v1514 = int32(1)
	if v1511 == v1510 {
		v1506 = v1506 + v1514
		v1507 = v1507 + v1514
		goto L335
	} else {
		goto L338
	}
L338:
	;
	goto L336
L339:
	;
	v1526 = int32(_a_F_sendDir_31)
	v1529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v1532 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[12])))
	if base.B2i32(v1529 == int32(0))|base.B2i32(v1529 != v1532) != 0 {
		v1550 = v1529
		v1551 = v1532
		goto L341
	} else {
		goto L342
	}
L340:
	;
	if v1550-v1551 == int32(0) {
		goto L319
	} else {
		goto L347
	}
L341:
	;
	goto L340
L342:
	;
	v1535 = v385
	v1536 = v1526
	goto L343
L343:
	;
	v1539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1536)+1)))
	v1540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1535)+1)))
	if v1540 == int32(0) {
		v1550 = v1540
		v1551 = v1539
		goto L341
	} else {
		goto L345
	}
L344:
	;
	v1550 = v1540
	v1551 = v1539
	goto L341
L345:
	;
	v1543 = int32(1)
	if v1540 == v1539 {
		v1535 = v1535 + v1543
		v1536 = v1536 + v1543
		goto L343
	} else {
		goto L346
	}
L346:
	;
	goto L344
L347:
	;
	v1555 = int32(_a_F_sendDir_32)
	v1558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v1561 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[13])))
	if base.B2i32(v1558 == int32(0))|base.B2i32(v1558 != v1561) != 0 {
		v1579 = v1558
		v1580 = v1561
		goto L349
	} else {
		goto L350
	}
L348:
	;
	if v1579-v1580 == int32(0) {
		goto L319
	} else {
		goto L355
	}
L349:
	;
	goto L348
L350:
	;
	v1564 = v385
	v1565 = v1555
	goto L351
L351:
	;
	v1568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1565)+1)))
	v1569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1564)+1)))
	if v1569 == int32(0) {
		v1579 = v1569
		v1580 = v1568
		goto L349
	} else {
		goto L353
	}
L352:
	;
	v1579 = v1569
	v1580 = v1568
	goto L349
L353:
	;
	v1572 = int32(1)
	if v1569 == v1568 {
		v1564 = v1564 + v1572
		v1565 = v1565 + v1572
		goto L351
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	v1584 = int32(_a_F_sendDir_33)
	v1587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[14])))
	if base.B2i32(v1587 == int32(0))|base.B2i32(v1587 != v1590) != 0 {
		v1608 = v1587
		v1609 = v1590
		goto L357
	} else {
		goto L358
	}
L356:
	;
	if v1608-v1609 == int32(0) {
		goto L319
	} else {
		goto L363
	}
L357:
	;
	goto L356
L358:
	;
	v1593 = v385
	v1594 = v1584
	goto L359
L359:
	;
	v1597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1594)+1)))
	v1598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593)+1)))
	if v1598 == int32(0) {
		v1608 = v1598
		v1609 = v1597
		goto L357
	} else {
		goto L361
	}
L360:
	;
	v1608 = v1598
	v1609 = v1597
	goto L357
L361:
	;
	v1601 = int32(1)
	if v1598 == v1597 {
		v1593 = v1593 + v1601
		v1594 = v1594 + v1601
		goto L359
	} else {
		goto L362
	}
L362:
	;
	goto L360
L363:
	;
	v1613 = int32(_a_F_sendDir_34)
	v1616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v1619 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[15])))
	if base.B2i32(v1616 == int32(0))|base.B2i32(v1616 != v1619) != 0 {
		v1637 = v1616
		v1638 = v1619
		goto L365
	} else {
		goto L366
	}
L364:
	;
	if v1637-v1638 == int32(0) {
		goto L319
	} else {
		goto L371
	}
L365:
	;
	goto L364
L366:
	;
	v1622 = v385
	v1623 = v1613
	goto L367
L367:
	;
	v1626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1623)+1)))
	v1627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1622)+1)))
	if v1627 == int32(0) {
		v1637 = v1627
		v1638 = v1626
		goto L365
	} else {
		goto L369
	}
L368:
	;
	v1637 = v1627
	v1638 = v1626
	goto L365
L369:
	;
	v1630 = int32(1)
	if v1627 == v1626 {
		v1622 = v1622 + v1630
		v1623 = v1623 + v1630
		goto L367
	} else {
		goto L370
	}
L370:
	;
	goto L368
L371:
	;
	v1642 = int32(_a_F_sendDir_35)
	v1645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v1648 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[16])))
	if base.B2i32(v1645 == int32(0))|base.B2i32(v1645 != v1648) != 0 {
		v1666 = v1645
		v1667 = v1648
		goto L373
	} else {
		goto L374
	}
L372:
	;
	if v1666-v1667 == int32(0) {
		goto L319
	} else {
		goto L379
	}
L373:
	;
	goto L372
L374:
	;
	v1651 = v385
	v1652 = v1642
	goto L375
L375:
	;
	v1655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1652)+1)))
	v1656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1651)+1)))
	if v1656 == int32(0) {
		v1666 = v1656
		v1667 = v1655
		goto L373
	} else {
		goto L377
	}
L376:
	;
	v1666 = v1656
	v1667 = v1655
	goto L373
L377:
	;
	v1659 = int32(1)
	if v1656 == v1655 {
		v1651 = v1651 + v1659
		v1652 = v1652 + v1659
		goto L375
	} else {
		goto L378
	}
L378:
	;
	goto L376
L379:
	;
	v1671 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v330)+2376)))
	v1672 = *(*int64)(unsafe.Add(mBase, uint32(v330)+2368))
	if v1671|(v1672^int64(7809654480478154542)) != int64(0) {
		goto L318
	} else {
		goto L380
	}
L380:
	;
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2276))
	if v1678&int32(_a_F_sendDir_27) == int32(_a_F_sendDir_36) {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2276)) = v1684 | int32(_a_F_sendDir_37)
	goto L383
L382:
	;
	goto L383
L383:
	;
	v1690 = v330 + int32(2272)
	F__tarWriteHeader(m, v320, v341, int32(0), v1690, v323)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L4
	} else {
		goto L384
	}
L384:
	;
	F__tarWriteHeader(m, v320, int32(_a_F_sendDir_38), int32(0), v1690, v323)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L4
	} else {
		goto L385
	}
L385:
	;
	F__tarWriteHeader(m, v320, int32(_a_F_sendDir_39), int32(0), v1690, v323)
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L4
	} else {
		goto L386
	}
L386:
	;
	v3923 = v364 + int64(1536)
	goto L91
L387:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L4
	} else {
		goto L388
	}
L388:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L4
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+128)) = v330 + int32(2368)
	F_errmsg(m, int32(_a_F_sendDir_40), v330+int32(128))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L4
	} else {
		goto L390
	}
L390:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1360), int32(_a_F_sendDir_20))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L4
	} else {
		goto L391
	}
L391:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L392:
	;
	if v1728 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+32)) = v385
	F_errmsg_internal(m, int32(_a_F_sendDir_41), v330+int32(32))
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L4
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2276))
	if v1741&int32(_a_F_sendDir_27) == int32(_a_F_sendDir_36) {
		goto L398
	} else {
		goto L399
	}
L396:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1372), int32(_a_F_sendDir_20))
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L4
	} else {
		goto L397
	}
L397:
	;
	goto L395
L398:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2276)) = v1747 | int32(_a_F_sendDir_37)
	goto L400
L399:
	;
	goto L400
L400:
	;
	F__tarWriteHeader(m, v320, v341, int32(0), v330+int32(2272), v323)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L4
	} else {
		goto L401
	}
L401:
	;
	v3923 = v364 + int64(512)
	goto L91
L402:
	;
	if v1785-v1786|base.B2i32(v1760 != int32(_a_F_sendDir_36)) == int32(0) {
		goto L409
	} else {
		goto L410
	}
L403:
	;
	goto L402
L404:
	;
	v1770 = v321
	v1771 = v1761
	goto L405
L405:
	;
	v1774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1771)+1)))
	v1775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1770)+1)))
	if v1775 == int32(0) {
		v1785 = v1775
		v1786 = v1774
		goto L403
	} else {
		goto L407
	}
L406:
	;
	v1785 = v1775
	v1786 = v1774
	goto L403
L407:
	;
	v1778 = int32(1)
	if v1775 == v1774 {
		v1770 = v1770 + v1778
		v1771 = v1771 + v1778
		goto L405
	} else {
		goto L408
	}
L408:
	;
	goto L406
L409:
	;
	v1796 = v330 + int32(192)
	v1798 = F_readlink(m, v330+int32(2368), v1796, int32(1024))
	mBase = m.M
	if v1798 < int32(0) {
		goto L82
	} else {
		goto L412
	}
L410:
	;
	goto L411
L411:
	;
	if v1760 != int32(_a_F_sendDir_42) {
		goto L416
	} else {
		goto L417
	}
L412:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v1798) {
		goto L81
	} else {
		goto L413
	}
L413:
	;
	v1804 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1798+v1796))) = uint8(v1804)
	F__tarWriteHeader(m, v320, v341, v1796, v330+int32(2272), v323)
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L4
	} else {
		goto L414
	}
L414:
	;
	v3923 = v364 + int64(512)
	goto L91
L415:
	;
	v3862 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v3863 = m.ExcPending
	if v3863 != 0 {
		goto L4
	} else {
		goto L681
	}
L416:
	;
	if v1760 != int32(_a_F_sendDir_37) {
		goto L415
	} else {
		goto L419
	}
L417:
	;
	goto L418
L418:
	;
	v1984 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2256)) = v1984
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2252)) = v1984
	if base.B2i32(v328 == v1984)|(v1041^int32(1)) == v1984 {
		goto L445
	} else {
		goto L446
	}
L419:
	;
	F__tarWriteHeader(m, v320, v341, int32(0), v330+int32(2272), v323)
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L4
	} else {
		goto L420
	}
L420:
	;
	if v324 == int32(0) {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	v1964 = v364 + int64(512)
	if v1931 == int32(0) {
		v3923 = v1964
		goto L91
	} else {
		goto L442
	}
L422:
	;
	v1931 = int32(1)
	goto L421
L423:
	;
	goto L424
L424:
	;
	v1824 = int32(1)
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	if v1825 <= int32(0) {
		v1931 = v1824
		goto L421
	} else {
		goto L425
	}
L425:
	;
	v1828 = int32(0)
	if v1828 < v1825 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1831 = v1825
	goto L428
L427:
	;
	v1831 = v1828
	goto L428
L428:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v324)+12))
	v1843 = int32(0)
	goto L429
L429:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1832+v1843<<(uint(int32(2))%32))))
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1883)+8))
	if v1884 == int32(0) {
		goto L431
	} else {
		goto L432
	}
L430:
	;
	v1931 = v1824
	goto L421
L431:
	;
	v1915 = v1843 + int32(1)
	if v1915 != v1831 {
		v1843 = v1915
		goto L429
	} else {
		goto L441
	}
L432:
	;
	v1889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1884))))
	v1892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	if base.B2i32(v1889 == int32(0))|base.B2i32(v1889 != v1892) != 0 {
		v1910 = v1889
		v1911 = v1892
		goto L434
	} else {
		goto L435
	}
L433:
	;
	if v1910-v1911 != 0 {
		goto L431
	} else {
		goto L440
	}
L434:
	;
	goto L433
L435:
	;
	v1895 = v1884
	v1896 = v361
	goto L436
L436:
	;
	v1899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896)+1)))
	v1900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1895)+1)))
	if v1900 == int32(0) {
		v1910 = v1900
		v1911 = v1899
		goto L434
	} else {
		goto L438
	}
L437:
	;
	v1910 = v1900
	v1911 = v1899
	goto L434
L438:
	;
	v1903 = int32(1)
	if v1900 == v1899 {
		v1895 = v1895 + v1903
		v1896 = v1896 + v1903
		goto L436
	} else {
		goto L439
	}
L439:
	;
	goto L437
L440:
	;
	v1931 = int32(0)
	goto L421
L441:
	;
	goto L430
L442:
	;
	v1967 = *(*int64)(unsafe.Add(mBase, uint32(v330)+2368))
	v1970 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v330)+2376)))
	if v325|base.B2i32(v1967^int64(7809932656919981870)|(v1970^int64(6516851)) != int64(0)) == int32(0) {
		v3923 = v1964
		goto L91
	} else {
		goto L443
	}
L443:
	;
	v1981 = F_sendDir(m, v320, v330+int32(2368), v322, v323, v324, v325, v326, v327, v328)
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L4
	} else {
		goto L444
	}
L444:
	;
	v3923 = v1981 + v1964
	goto L91
L445:
	;
	if v327 != 0 {
		goto L449
	} else {
		goto L450
	}
L446:
	;
	v3807 = v1984
	v3810 = v341
	goto L447
L447:
	;
	if v323 == int32(0) {
		goto L676
	} else {
		goto L677
	}
L448:
	;
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2268))
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2264))
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2260))
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2296))
	v2014 = v330 + int32(2256)
	v2016 = v330 + int32(2252)
	v2017 = int32(0)
	v2018 = m.G0
	v2020 = v2018 - int32(128)
	m.G0 = v2020
	if v2012&int32(_a_F_sendDir_43)|base.B2i32(v2010 == int32(1))|base.B2i32(base.Ui32(int32(1073741824)) < base.Ui32(v2012)) != 0 {
		v3701 = v2017
		goto L456
	} else {
		goto L457
	}
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+120)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v330)+116)) = v327
	*(*int32)(unsafe.Add(mBase, uint32(v330)+112)) = int32(_a_F_sendDir_44)
	v2003 = F_psprintf(m, int32(_a_F_sendDir_45), v330+int32(112))
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		goto L4
	} else {
		goto L452
	}
L450:
	;
	goto L451
L451:
	;
	v2005 = F_pstrdup(m, v341)
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L4
	} else {
		goto L453
	}
L452:
	;
	v2007 = v2003
	v2008 = v327
	goto L448
L453:
	;
	v2007 = v2005
	v2008 = v350
	goto L448
L454:
	;
	if v3701 == int32(1) {
		goto L668
	} else {
		goto L669
	}
L455:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3732 = m.ExcPending
	if v3732 != 0 {
		goto L4
	} else {
		goto L664
	}
L456:
	;
	m.G0 = v2020 + int32(128)
	goto L454
L457:
	;
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	v2031 = F_strlen(m, v2007)
	mBase = m.M
	v2037 = v2031 - int32(1636608432)
	if v2007&int32(3) != 0 {
		goto L462
	} else {
		goto L463
	}
L458:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2030)+20))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v2030)+12))
	v2298 = (v2291 ^ v2283 - base.I32_rotl(v2291, int32(24))) & v2297
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(v2296+v2298<<(uint(int32(4))%32))))
	if v2302 != 0 {
		goto L499
	} else {
		goto L500
	}
L459:
	;
	v2269 = int32(14)
	v2271 = v2265 ^ v2266 - base.I32_rotl(v2265, v2269)
	v2275 = v2271 ^ v2264 - base.I32_rotl(v2271, int32(11))
	v2279 = v2275 ^ v2265 - base.I32_rotl(v2275, int32(25))
	v2283 = v2279 ^ v2271 - base.I32_rotl(v2279, int32(16))
	v2287 = v2283 ^ v2275 - base.I32_rotl(v2283, int32(4))
	v2291 = v2287 ^ v2279 - base.I32_rotl(v2287, v2269)
	goto L458
L460:
	;
	switch v2195 - int32(1) {
	case 0:
		v2257 = v2196
		v2258 = v2197
		v2259 = v2198
		goto L487
	case 1:
		v2250 = v2196
		v2251 = v2197
		v2252 = v2198
		goto L488
	case 2:
		v2243 = v2196
		v2244 = v2197
		v2245 = v2198
		goto L489
	case 3:
		v2237 = v2197
		v2238 = v2198
		goto L490
	case 4:
		v2233 = v2197
		v2234 = v2198
		goto L491
	case 5:
		v2227 = v2197
		v2228 = v2198
		goto L492
	case 6:
		v2221 = v2197
		v2222 = v2198
		goto L493
	case 7:
		v2216 = v2198
		goto L494
	case 8:
		v2211 = v2198
		goto L495
	case 9:
		v2206 = v2198
		goto L496
	case 10:
		goto L497
	default:
		v2264 = v2196
		v2265 = v2197
		v2266 = v2198
		goto L459
	}
L461:
	;
	v2146 = v2007
	v2147 = v2031
	v2148 = v2037
	v2149 = v2037
	v2150 = v2037
	goto L484
L462:
	;
	if base.Ui32(int32(11)) < base.Ui32(v2031) {
		goto L461
	} else {
		goto L465
	}
L463:
	;
	goto L464
L464:
	;
	if base.Ui32(v2031) < base.Ui32(int32(12)) {
		goto L467
	} else {
		goto L468
	}
L465:
	;
	v2194 = v2007
	v2195 = v2031
	v2196 = v2037
	v2197 = v2037
	v2198 = v2037
	goto L460
L466:
	;
	switch v2093 - int32(1) {
	case 0:
		v2143 = v2094
		goto L473
	case 1:
		v2138 = v2094
		goto L474
	case 2:
		goto L475
	case 3:
		v2131 = v2095
		goto L476
	case 4:
		v2128 = v2095
		goto L477
	case 5:
		v2123 = v2095
		goto L478
	case 6:
		goto L479
	case 7:
		v2114 = v2096
		goto L480
	case 8:
		v2109 = v2096
		goto L481
	case 9:
		v2104 = v2096
		goto L482
	case 10:
		goto L483
	default:
		v2264 = v2094
		v2265 = v2095
		v2266 = v2096
		goto L459
	}
L467:
	;
	v2092 = v2007
	v2093 = v2031
	v2094 = v2037
	v2095 = v2037
	v2096 = v2037
	goto L466
L468:
	;
	goto L469
L469:
	;
	v2044 = v2007
	v2045 = v2031
	v2046 = v2037
	v2047 = v2037
	v2048 = v2037
	goto L470
L470:
	;
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(v2044)+4))
	v2051 = v2050 + v2047
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v2044)))
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v2044)+8))
	v2055 = v2054 + v2048
	v2057 = int32(4)
	v2059 = v2052 + v2046 - v2055 ^ base.I32_rotl(v2055, v2057)
	v2063 = v2051 - v2059 ^ base.I32_rotl(v2059, int32(6))
	v2064 = v2055 + v2051
	v2065 = v2059 + v2064
	v2066 = v2063 + v2065
	v2070 = v2064 - v2063 ^ base.I32_rotl(v2063, int32(8))
	v2074 = v2065 - v2070 ^ base.I32_rotl(v2070, int32(16))
	v2078 = v2066 - v2074 ^ base.I32_rotl(v2074, int32(19))
	v2079 = v2070 + v2066
	v2080 = v2074 + v2079
	v2081 = v2078 + v2080
	v2085 = v2079 - v2078 ^ base.I32_rotl(v2078, v2057)
	v2086 = int32(12)
	v2087 = v2044 + v2086
	v2089 = v2045 - v2086
	if base.Ui32(int32(11)) < base.Ui32(v2089) {
		v2044 = v2087
		v2045 = v2089
		v2046 = v2080
		v2047 = v2081
		v2048 = v2085
		goto L470
	} else {
		goto L472
	}
L471:
	;
	v2092 = v2087
	v2093 = v2089
	v2094 = v2080
	v2095 = v2081
	v2096 = v2085
	goto L466
L472:
	;
	goto L471
L473:
	;
	v2144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2092))))
	v2264 = v2143 + v2144
	v2265 = v2095
	v2266 = v2096
	goto L459
L474:
	;
	v2139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2092)+1)))
	v2143 = v2139<<(uint(int32(8))%32) + v2138
	goto L473
L475:
	;
	v2134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2092)+2)))
	v2138 = v2134<<(uint(int32(16))%32) + v2094
	goto L474
L476:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v2092)))
	v2264 = v2132 + v2094
	v2265 = v2131
	v2266 = v2096
	goto L459
L477:
	;
	v2129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2092)+4)))
	v2131 = v2128 + v2129
	goto L476
L478:
	;
	v2124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2092)+5)))
	v2128 = v2124<<(uint(int32(8))%32) + v2123
	goto L477
L479:
	;
	v2119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2092)+6)))
	v2123 = v2119<<(uint(int32(16))%32) + v2095
	goto L478
L480:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v2092)))
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+4))
	v2264 = v2115 + v2094
	v2265 = v2117 + v2095
	v2266 = v2114
	goto L459
L481:
	;
	v2110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2092)+8)))
	v2114 = v2110<<(uint(int32(8))%32) + v2109
	goto L480
L482:
	;
	v2105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2092)+9)))
	v2109 = v2105<<(uint(int32(16))%32) + v2104
	goto L481
L483:
	;
	v2100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2092)+10)))
	v2104 = v2100<<(uint(int32(24))%32) + v2096
	goto L482
L484:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v2146)+4))
	v2153 = v2152 + v2149
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v2146)))
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v2146)+8))
	v2157 = v2156 + v2150
	v2159 = int32(4)
	v2161 = v2154 + v2148 - v2157 ^ base.I32_rotl(v2157, v2159)
	v2165 = v2153 - v2161 ^ base.I32_rotl(v2161, int32(6))
	v2166 = v2157 + v2153
	v2167 = v2161 + v2166
	v2168 = v2165 + v2167
	v2172 = v2166 - v2165 ^ base.I32_rotl(v2165, int32(8))
	v2176 = v2167 - v2172 ^ base.I32_rotl(v2172, int32(16))
	v2180 = v2168 - v2176 ^ base.I32_rotl(v2176, int32(19))
	v2181 = v2172 + v2168
	v2182 = v2176 + v2181
	v2183 = v2180 + v2182
	v2187 = v2181 - v2180 ^ base.I32_rotl(v2180, v2159)
	v2188 = int32(12)
	v2189 = v2146 + v2188
	v2191 = v2147 - v2188
	if base.Ui32(int32(11)) < base.Ui32(v2191) {
		v2146 = v2189
		v2147 = v2191
		v2148 = v2182
		v2149 = v2183
		v2150 = v2187
		goto L484
	} else {
		goto L486
	}
L485:
	;
	v2194 = v2189
	v2195 = v2191
	v2196 = v2182
	v2197 = v2183
	v2198 = v2187
	goto L460
L486:
	;
	goto L485
L487:
	;
	v2260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194))))
	v2264 = v2257 + v2260
	v2265 = v2258
	v2266 = v2259
	goto L459
L488:
	;
	v2253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194)+1)))
	v2257 = v2253<<(uint(int32(8))%32) + v2250
	v2258 = v2251
	v2259 = v2252
	goto L487
L489:
	;
	v2246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194)+2)))
	v2250 = v2246<<(uint(int32(16))%32) + v2243
	v2251 = v2244
	v2252 = v2245
	goto L488
L490:
	;
	v2239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194)+3)))
	v2243 = v2239<<(uint(int32(24))%32) + v2196
	v2244 = v2237
	v2245 = v2238
	goto L489
L491:
	;
	v2235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194)+4)))
	v2237 = v2233 + v2235
	v2238 = v2234
	goto L490
L492:
	;
	v2229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194)+5)))
	v2233 = v2229<<(uint(int32(8))%32) + v2227
	v2234 = v2228
	goto L491
L493:
	;
	v2223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194)+6)))
	v2227 = v2223<<(uint(int32(16))%32) + v2221
	v2228 = v2222
	goto L492
L494:
	;
	v2217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194)+7)))
	v2221 = v2217<<(uint(int32(24))%32) + v2197
	v2222 = v2216
	goto L493
L495:
	;
	v2212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194)+8)))
	v2216 = v2212<<(uint(int32(8))%32) + v2211
	goto L494
L496:
	;
	v2207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194)+9)))
	v2211 = v2207<<(uint(int32(16))%32) + v2206
	goto L495
L497:
	;
	v2202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194)+10)))
	v2206 = v2202<<(uint(int32(24))%32) + v2198
	goto L496
L498:
	;
	v2890 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+64)) = v2890
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+60)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+56)) = v2008
	v2894 = *(*int32)(unsafe.Add(mBase, uint32(v328)+28))
	v2896 = v2020 + int32(56)
	v2899 = v2020 + int32(52)
	v2903 = m.G0
	v2904 = int32(16)
	v2905 = v2903 - v2904
	m.G0 = v2905
	v2907 = *(*int32)(unsafe.Add(mBase, uint32(v2896)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2905)+8)) = v2907
	v2909 = *(*int64)(unsafe.Add(mBase, uint32(v2896)))
	*(*int64)(unsafe.Add(mBase, uint32(v2905))) = v2909
	v2911 = *(*int32)(unsafe.Add(mBase, uint32(v2894)))
	*(*int32)(unsafe.Add(mBase, uint32(v2905)+12)) = v2890
	v2914 = F_hash_bytes(m, v2905, v2904)
	mBase = m.M
	v2915 = *(*int32)(unsafe.Add(mBase, uint32(v2911)+20))
	v2916 = *(*int32)(unsafe.Add(mBase, uint32(v2911)+12))
	v2917 = v2914 & v2916
	v2920 = v2915 + v2917*int32(40)
	v2921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2920)+20)))
	if v2921 == v2890 {
		goto L580
	} else {
		goto L581
	}
L499:
	;
	v2312 = v2298
	goto L502
L500:
	;
	goto L501
L501:
	;
	v2435 = v2020 + int32(56)
	F_GetRelationPath(m, v2435, v348, v2008, v2009, int32(-1), v2010)
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L4
	} else {
		goto L513
	}
L502:
	;
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(v2296+v2312<<(uint(int32(4))%32))+4))
	v2355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2352))))
	v2358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2007))))
	if base.B2i32(v2355 == int32(0))|base.B2i32(v2355 != v2358) != 0 {
		v2376 = v2355
		v2377 = v2358
		goto L505
	} else {
		goto L506
	}
L503:
	;
	goto L501
L504:
	;
	if v2376-v2377 == int32(0) {
		goto L498
	} else {
		goto L511
	}
L505:
	;
	goto L504
L506:
	;
	v2361 = v2352
	v2362 = v2007
	goto L507
L507:
	;
	v2365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2362)+1)))
	v2366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2361)+1)))
	if v2366 == int32(0) {
		v2376 = v2366
		v2377 = v2365
		goto L505
	} else {
		goto L509
	}
L508:
	;
	v2376 = v2366
	v2377 = v2365
	goto L505
L509:
	;
	v2369 = int32(1)
	if v2366 == v2365 {
		v2361 = v2361 + v2369
		v2362 = v2362 + v2369
		goto L507
	} else {
		goto L510
	}
L510:
	;
	goto L508
L511:
	;
	v2383 = (v2312 + int32(1)) & v2297
	v2387 = *(*int32)(unsafe.Add(mBase, uint32(v2296+v2383<<(uint(int32(4))%32))))
	if v2387 != 0 {
		v2312 = v2383
		goto L502
	} else {
		goto L512
	}
L512:
	;
	goto L503
L513:
	;
	v2442 = F_strlen(m, v2435)
	mBase = m.M
	v2449 = v2442 + int32(1)
	goto L516
L514:
	;
	v2462 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2461))) = uint8(v2462)
	v2465 = v2461 + int32(1)
	if v2011 != 0 {
		goto L521
	} else {
		goto L522
	}
L515:
	;
	goto L514
L516:
	;
	v2451 = int32(0)
	if v2449 == v2451 {
		v2461 = v2451
		goto L515
	} else {
		goto L518
	}
L517:
	;
	v2461 = v2456
	goto L515
L518:
	;
	v2455 = v2449 - int32(1)
	v2456 = v2435 + v2455
	v2457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2456))))
	if v2457 != int32(47) {
		v2449 = v2455
		goto L516
	} else {
		goto L519
	}
L519:
	;
	goto L517
L520:
	;
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	v2485 = F_strlen(m, v2483)
	mBase = m.M
	v2491 = v2485 - int32(1636608432)
	if v2483&int32(3) != 0 {
		goto L530
	} else {
		goto L531
	}
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+40)) = v2011
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+36)) = v2465
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+32)) = v2435
	v2472 = F_psprintf(m, int32(_a_F_sendDir_46), v2020+int32(32))
	mBase = m.M
	v2473 = m.ExcPending
	if v2473 != 0 {
		goto L4
	} else {
		goto L524
	}
L522:
	;
	goto L523
L523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+20)) = v2465
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+16)) = v2020 + int32(56)
	v2481 = F_psprintf(m, int32(_a_F_sendDir_47), v2020+int32(16))
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L4
	} else {
		goto L525
	}
L524:
	;
	v2483 = v2472
	goto L520
L525:
	;
	v2483 = v2481
	goto L520
L526:
	;
	v2750 = *(*int32)(unsafe.Add(mBase, uint32(v2484)+20))
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2484)+12))
	v2752 = (v2745 ^ v2737 - base.I32_rotl(v2745, int32(24))) & v2751
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v2750+v2752<<(uint(int32(4))%32))))
	if v2756 == int32(0) {
		v3701 = v2017
		goto L456
	} else {
		goto L566
	}
L527:
	;
	v2723 = int32(14)
	v2725 = v2719 ^ v2720 - base.I32_rotl(v2719, v2723)
	v2729 = v2725 ^ v2718 - base.I32_rotl(v2725, int32(11))
	v2733 = v2729 ^ v2719 - base.I32_rotl(v2729, int32(25))
	v2737 = v2733 ^ v2725 - base.I32_rotl(v2733, int32(16))
	v2741 = v2737 ^ v2729 - base.I32_rotl(v2737, int32(4))
	v2745 = v2741 ^ v2733 - base.I32_rotl(v2741, v2723)
	goto L526
L528:
	;
	switch v2649 - int32(1) {
	case 0:
		v2711 = v2650
		v2712 = v2651
		v2713 = v2652
		goto L555
	case 1:
		v2704 = v2650
		v2705 = v2651
		v2706 = v2652
		goto L556
	case 2:
		v2697 = v2650
		v2698 = v2651
		v2699 = v2652
		goto L557
	case 3:
		v2691 = v2651
		v2692 = v2652
		goto L558
	case 4:
		v2687 = v2651
		v2688 = v2652
		goto L559
	case 5:
		v2681 = v2651
		v2682 = v2652
		goto L560
	case 6:
		v2675 = v2651
		v2676 = v2652
		goto L561
	case 7:
		v2670 = v2652
		goto L562
	case 8:
		v2665 = v2652
		goto L563
	case 9:
		v2660 = v2652
		goto L564
	case 10:
		goto L565
	default:
		v2718 = v2650
		v2719 = v2651
		v2720 = v2652
		goto L527
	}
L529:
	;
	v2600 = v2483
	v2601 = v2485
	v2602 = v2491
	v2603 = v2491
	v2604 = v2491
	goto L552
L530:
	;
	if base.Ui32(int32(11)) < base.Ui32(v2485) {
		goto L529
	} else {
		goto L533
	}
L531:
	;
	goto L532
L532:
	;
	if base.Ui32(v2485) < base.Ui32(int32(12)) {
		goto L535
	} else {
		goto L536
	}
L533:
	;
	v2648 = v2483
	v2649 = v2485
	v2650 = v2491
	v2651 = v2491
	v2652 = v2491
	goto L528
L534:
	;
	switch v2547 - int32(1) {
	case 0:
		v2597 = v2548
		goto L541
	case 1:
		v2592 = v2548
		goto L542
	case 2:
		goto L543
	case 3:
		v2585 = v2549
		goto L544
	case 4:
		v2582 = v2549
		goto L545
	case 5:
		v2577 = v2549
		goto L546
	case 6:
		goto L547
	case 7:
		v2568 = v2550
		goto L548
	case 8:
		v2563 = v2550
		goto L549
	case 9:
		v2558 = v2550
		goto L550
	case 10:
		goto L551
	default:
		v2718 = v2548
		v2719 = v2549
		v2720 = v2550
		goto L527
	}
L535:
	;
	v2546 = v2483
	v2547 = v2485
	v2548 = v2491
	v2549 = v2491
	v2550 = v2491
	goto L534
L536:
	;
	goto L537
L537:
	;
	v2498 = v2483
	v2499 = v2485
	v2500 = v2491
	v2501 = v2491
	v2502 = v2491
	goto L538
L538:
	;
	v2504 = *(*int32)(unsafe.Add(mBase, uint32(v2498)+4))
	v2505 = v2504 + v2501
	v2506 = *(*int32)(unsafe.Add(mBase, uint32(v2498)))
	v2508 = *(*int32)(unsafe.Add(mBase, uint32(v2498)+8))
	v2509 = v2508 + v2502
	v2511 = int32(4)
	v2513 = v2506 + v2500 - v2509 ^ base.I32_rotl(v2509, v2511)
	v2517 = v2505 - v2513 ^ base.I32_rotl(v2513, int32(6))
	v2518 = v2509 + v2505
	v2519 = v2513 + v2518
	v2520 = v2517 + v2519
	v2524 = v2518 - v2517 ^ base.I32_rotl(v2517, int32(8))
	v2528 = v2519 - v2524 ^ base.I32_rotl(v2524, int32(16))
	v2532 = v2520 - v2528 ^ base.I32_rotl(v2528, int32(19))
	v2533 = v2524 + v2520
	v2534 = v2528 + v2533
	v2535 = v2532 + v2534
	v2539 = v2533 - v2532 ^ base.I32_rotl(v2532, v2511)
	v2540 = int32(12)
	v2541 = v2498 + v2540
	v2543 = v2499 - v2540
	if base.Ui32(int32(11)) < base.Ui32(v2543) {
		v2498 = v2541
		v2499 = v2543
		v2500 = v2534
		v2501 = v2535
		v2502 = v2539
		goto L538
	} else {
		goto L540
	}
L539:
	;
	v2546 = v2541
	v2547 = v2543
	v2548 = v2534
	v2549 = v2535
	v2550 = v2539
	goto L534
L540:
	;
	goto L539
L541:
	;
	v2598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2546))))
	v2718 = v2597 + v2598
	v2719 = v2549
	v2720 = v2550
	goto L527
L542:
	;
	v2593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2546)+1)))
	v2597 = v2593<<(uint(int32(8))%32) + v2592
	goto L541
L543:
	;
	v2588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2546)+2)))
	v2592 = v2588<<(uint(int32(16))%32) + v2548
	goto L542
L544:
	;
	v2586 = *(*int32)(unsafe.Add(mBase, uint32(v2546)))
	v2718 = v2586 + v2548
	v2719 = v2585
	v2720 = v2550
	goto L527
L545:
	;
	v2583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2546)+4)))
	v2585 = v2582 + v2583
	goto L544
L546:
	;
	v2578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2546)+5)))
	v2582 = v2578<<(uint(int32(8))%32) + v2577
	goto L545
L547:
	;
	v2573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2546)+6)))
	v2577 = v2573<<(uint(int32(16))%32) + v2549
	goto L546
L548:
	;
	v2569 = *(*int32)(unsafe.Add(mBase, uint32(v2546)))
	v2571 = *(*int32)(unsafe.Add(mBase, uint32(v2546)+4))
	v2718 = v2569 + v2548
	v2719 = v2571 + v2549
	v2720 = v2568
	goto L527
L549:
	;
	v2564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2546)+8)))
	v2568 = v2564<<(uint(int32(8))%32) + v2563
	goto L548
L550:
	;
	v2559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2546)+9)))
	v2563 = v2559<<(uint(int32(16))%32) + v2558
	goto L549
L551:
	;
	v2554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2546)+10)))
	v2558 = v2554<<(uint(int32(24))%32) + v2550
	goto L550
L552:
	;
	v2606 = *(*int32)(unsafe.Add(mBase, uint32(v2600)+4))
	v2607 = v2606 + v2603
	v2608 = *(*int32)(unsafe.Add(mBase, uint32(v2600)))
	v2610 = *(*int32)(unsafe.Add(mBase, uint32(v2600)+8))
	v2611 = v2610 + v2604
	v2613 = int32(4)
	v2615 = v2608 + v2602 - v2611 ^ base.I32_rotl(v2611, v2613)
	v2619 = v2607 - v2615 ^ base.I32_rotl(v2615, int32(6))
	v2620 = v2611 + v2607
	v2621 = v2615 + v2620
	v2622 = v2619 + v2621
	v2626 = v2620 - v2619 ^ base.I32_rotl(v2619, int32(8))
	v2630 = v2621 - v2626 ^ base.I32_rotl(v2626, int32(16))
	v2634 = v2622 - v2630 ^ base.I32_rotl(v2630, int32(19))
	v2635 = v2626 + v2622
	v2636 = v2630 + v2635
	v2637 = v2634 + v2636
	v2641 = v2635 - v2634 ^ base.I32_rotl(v2634, v2613)
	v2642 = int32(12)
	v2643 = v2600 + v2642
	v2645 = v2601 - v2642
	if base.Ui32(int32(11)) < base.Ui32(v2645) {
		v2600 = v2643
		v2601 = v2645
		v2602 = v2636
		v2603 = v2637
		v2604 = v2641
		goto L552
	} else {
		goto L554
	}
L553:
	;
	v2648 = v2643
	v2649 = v2645
	v2650 = v2636
	v2651 = v2637
	v2652 = v2641
	goto L528
L554:
	;
	goto L553
L555:
	;
	v2714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2648))))
	v2718 = v2711 + v2714
	v2719 = v2712
	v2720 = v2713
	goto L527
L556:
	;
	v2707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2648)+1)))
	v2711 = v2707<<(uint(int32(8))%32) + v2704
	v2712 = v2705
	v2713 = v2706
	goto L555
L557:
	;
	v2700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2648)+2)))
	v2704 = v2700<<(uint(int32(16))%32) + v2697
	v2705 = v2698
	v2706 = v2699
	goto L556
L558:
	;
	v2693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2648)+3)))
	v2697 = v2693<<(uint(int32(24))%32) + v2650
	v2698 = v2691
	v2699 = v2692
	goto L557
L559:
	;
	v2689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2648)+4)))
	v2691 = v2687 + v2689
	v2692 = v2688
	goto L558
L560:
	;
	v2683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2648)+5)))
	v2687 = v2683<<(uint(int32(8))%32) + v2681
	v2688 = v2682
	goto L559
L561:
	;
	v2677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2648)+6)))
	v2681 = v2677<<(uint(int32(16))%32) + v2675
	v2682 = v2676
	goto L560
L562:
	;
	v2671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2648)+7)))
	v2675 = v2671<<(uint(int32(24))%32) + v2651
	v2676 = v2670
	goto L561
L563:
	;
	v2666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2648)+8)))
	v2670 = v2666<<(uint(int32(8))%32) + v2665
	goto L562
L564:
	;
	v2661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2648)+9)))
	v2665 = v2661<<(uint(int32(16))%32) + v2660
	goto L563
L565:
	;
	v2656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2648)+10)))
	v2660 = v2656<<(uint(int32(24))%32) + v2652
	goto L564
L566:
	;
	v2768 = v2752
	goto L567
L567:
	;
	v2808 = *(*int32)(unsafe.Add(mBase, uint32(v2750+v2768<<(uint(int32(4))%32))+4))
	v2811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2808))))
	v2814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2483))))
	if base.B2i32(v2811 == int32(0))|base.B2i32(v2811 != v2814) != 0 {
		v2832 = v2811
		v2833 = v2814
		goto L570
	} else {
		goto L571
	}
L568:
	;
	v3701 = v2017
	goto L456
L569:
	;
	if v2832-v2833 == int32(0) {
		goto L498
	} else {
		goto L576
	}
L570:
	;
	goto L569
L571:
	;
	v2817 = v2808
	v2818 = v2483
	goto L572
L572:
	;
	v2821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2818)+1)))
	v2822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2817)+1)))
	if v2822 == int32(0) {
		v2832 = v2822
		v2833 = v2821
		goto L570
	} else {
		goto L574
	}
L573:
	;
	v2832 = v2822
	v2833 = v2821
	goto L570
L574:
	;
	v2825 = int32(1)
	if v2822 == v2821 {
		v2817 = v2817 + v2825
		v2818 = v2818 + v2825
		goto L572
	} else {
		goto L575
	}
L575:
	;
	goto L573
L576:
	;
	v2839 = (v2768 + int32(1)) & v2751
	v2843 = *(*int32)(unsafe.Add(mBase, uint32(v2750+v2839<<(uint(int32(4))%32))))
	if v2843 != 0 {
		v2768 = v2839
		goto L567
	} else {
		goto L577
	}
L577:
	;
	goto L568
L578:
	;
	if v2958 != 0 {
		v3701 = v2017
		goto L456
	} else {
		goto L588
	}
L579:
	;
	m.G0 = v2905 + int32(16)
	goto L578
L580:
	;
	v2958 = int32(0)
	goto L579
L581:
	;
	v2924 = *(*int64)(unsafe.Add(mBase, uint32(v2905)))
	v2925 = v2917
	v2926 = v2920
	goto L582
L582:
	;
	v2932 = *(*int64)(unsafe.Add(mBase, uint32(v2926)))
	v2934 = *(*int64)(unsafe.Add(mBase, uint32(v2926)+8))
	v2935 = *(*int64)(unsafe.Add(mBase, uint32(v2905)+8))
	if v2932^v2924|(v2934^v2935) != int64(0) {
		goto L584
	} else {
		goto L585
	}
L583:
	;
	v2947 = *(*int32)(unsafe.Add(mBase, uint32(v2926)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2899))) = v2947
	v2958 = v2926
	goto L579
L584:
	;
	v2942 = (v2925 + int32(1)) & v2916
	v2945 = v2915 + v2942*int32(40)
	v2946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2945)+20)))
	if v2946 != 0 {
		v2925 = v2942
		v2926 = v2945
		goto L582
	} else {
		goto L587
	}
L585:
	;
	goto L586
L586:
	;
	goto L583
L587:
	;
	goto L580
L588:
	;
	v2968 = int32(base.Ui32(v2012) >> (uint(int32(13)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+64)) = v2009
	v2970 = *(*int32)(unsafe.Add(mBase, uint32(v328)+28))
	v2974 = m.G0
	v2975 = int32(16)
	v2976 = v2974 - v2975
	m.G0 = v2976
	v2978 = *(*int32)(unsafe.Add(mBase, uint32(v2896)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2976)+8)) = v2978
	v2980 = *(*int64)(unsafe.Add(mBase, uint32(v2896)))
	*(*int64)(unsafe.Add(mBase, uint32(v2976))) = v2980
	v2982 = *(*int32)(unsafe.Add(mBase, uint32(v2970)))
	*(*int32)(unsafe.Add(mBase, uint32(v2976)+12)) = v2010
	v2985 = F_hash_bytes(m, v2976, v2975)
	mBase = m.M
	v2986 = *(*int32)(unsafe.Add(mBase, uint32(v2982)+20))
	v2987 = *(*int32)(unsafe.Add(mBase, uint32(v2982)+12))
	v2988 = v2985 & v2987
	v2991 = v2986 + v2988*int32(40)
	v2992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2991)+20)))
	if v2992 == int32(0) {
		goto L592
	} else {
		goto L593
	}
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2016))) = v3640
	v3701 = int32(1)
	goto L456
L590:
	;
	if v3029 == int32(0) {
		goto L600
	} else {
		goto L601
	}
L591:
	;
	m.G0 = v2976 + int32(16)
	goto L590
L592:
	;
	v3029 = int32(0)
	goto L591
L593:
	;
	v2995 = *(*int64)(unsafe.Add(mBase, uint32(v2976)))
	v2996 = v2988
	v2997 = v2991
	goto L594
L594:
	;
	v3003 = *(*int64)(unsafe.Add(mBase, uint32(v2997)))
	v3005 = *(*int64)(unsafe.Add(mBase, uint32(v2997)+8))
	v3006 = *(*int64)(unsafe.Add(mBase, uint32(v2976)+8))
	if v3003^v2995|(v3005^v3006) != int64(0) {
		goto L596
	} else {
		goto L597
	}
L595:
	;
	v3018 = *(*int32)(unsafe.Add(mBase, uint32(v2997)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2899))) = v3018
	v3029 = v2997
	goto L591
L596:
	;
	v3013 = (v2996 + int32(1)) & v2987
	v3016 = v2986 + v3013*int32(40)
	v3017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3016)+20)))
	if v3017 != 0 {
		v2996 = v3013
		v2997 = v3016
		goto L594
	} else {
		goto L599
	}
L597:
	;
	goto L598
L598:
	;
	goto L595
L599:
	;
	goto L592
L600:
	;
	if v2012 == int32(0) {
		v3701 = v2017
		goto L456
	} else {
		goto L603
	}
L601:
	;
	goto L602
L602:
	;
	v3045 = v2011 << (uint(int32(17)) % 32)
	v3046 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+52))
	if base.Ui32(v3046) <= base.Ui32(v3045) {
		v3701 = v2017
		goto L456
	} else {
		goto L604
	}
L603:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2014))) = int32(0)
	v3640 = v2968
	goto L589
L604:
	;
	if base.Ui32(int32(_a_F_sendDir_48)) < base.Ui32(v2011) {
		goto L455
	} else {
		goto L605
	}
L605:
	;
	v3050 = v2968 + v3045
	if base.Ui32(v3050) < base.Ui32(v2968) {
		goto L455
	} else {
		goto L606
	}
L606:
	;
	v3052 = int32(0)
	v3053 = int32(16)
	v3054 = int32(base.Ui32(v3045) >> (uint(v3053) % 32))
	v3061 = int32(base.Ui32(v3050)>>(uint(v3053)%32)) + base.B2i32(v3050&int32(_a_F_sendDir_49) != v3052)
	v3062 = *(*int32)(unsafe.Add(mBase, uint32(v3029)+24))
	if base.Ui32(v3061) < base.Ui32(v3062) {
		goto L608
	} else {
		goto L609
	}
L607:
	;
	if base.F64_gt(base.F64_convert_i32_u(v3340<<(uint(int32(13))%32)), base.F64_mul(base.F64_convert_i32_u(v2012), float64(0.9))) != 0 {
		v3701 = v2017
		goto L456
	} else {
		goto L641
	}
L608:
	;
	v3064 = v3061
	goto L610
L609:
	;
	v3064 = v3062
	goto L610
L610:
	;
	if base.Ui32(v3064) <= base.Ui32(v3054) {
		v3340 = v3052
		goto L607
	} else {
		goto L611
	}
L611:
	;
	v3079 = v3054
	v3082 = v3052
	goto L612
L612:
	;
	v3116 = *(*int32)(unsafe.Add(mBase, uint32(v3029)+32))
	v3120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3116+v3079<<(uint(int32(1))%32)))))
	if v3120 == int32(0) {
		v3291 = v3082
		goto L614
	} else {
		goto L615
	}
L613:
	;
	v3340 = v3291
	goto L607
L614:
	;
	v3326 = v3079 + int32(1)
	if v3326 != v3064 {
		v3079 = v3326
		v3082 = v3291
		goto L612
	} else {
		goto L640
	}
L615:
	;
	v3125 = v3079 << (uint(int32(16)) % 32)
	if v3079 != v3064-int32(1) {
		goto L616
	} else {
		goto L617
	}
L616:
	;
	v3128 = int32(_a_F_sendDir_50)
	goto L618
L617:
	;
	v3128 = v3050 - v3125
	goto L618
L618:
	;
	if v3079 == v3054 {
		goto L619
	} else {
		goto L620
	}
L619:
	;
	v3131 = v3045 & int32(_a_F_sendDir_49)
	goto L621
L620:
	;
	v3131 = int32(0)
	goto L621
L621:
	;
	v3132 = *(*int32)(unsafe.Add(mBase, uint32(v3029)+36))
	v3136 = *(*int32)(unsafe.Add(mBase, uint32(v3132+v3079<<(uint(int32(2))%32))))
	if v3120 != int32(_a_F_sendDir_51) {
		goto L622
	} else {
		goto L623
	}
L622:
	;
	v3152 = v3082
	v3171 = int32(0)
	goto L625
L623:
	;
	goto L624
L624:
	;
	if base.Ui32(v3128) <= base.Ui32(v3131) {
		v3291 = v3082
		goto L614
	} else {
		goto L632
	}
L625:
	;
	v3189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3136+v3171<<(uint(int32(1))%32)))))
	if base.B2i32(base.Ui32(v3189) < base.Ui32(v3131))|base.B2i32(base.Ui32(v3128) <= base.Ui32(v3189)) == int32(0) {
		goto L627
	} else {
		goto L628
	}
L626:
	;
	v3291 = v3204
	goto L614
L627:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v339+v3152<<(uint(int32(2))%32)))) = v3125 | v3189
	v3201 = v3152 + int32(1)
	if v3201 == int32(_a_F_sendDir_1) {
		v3340 = v3201
		goto L607
	} else {
		goto L630
	}
L628:
	;
	v3204 = v3152
	goto L629
L629:
	;
	v3206 = v3171 + int32(1)
	if v3206 != v3120 {
		v3152 = v3204
		v3171 = v3206
		goto L625
	} else {
		goto L631
	}
L630:
	;
	v3204 = v3201
	goto L629
L631:
	;
	goto L626
L632:
	;
	v3221 = v3082
	v3224 = v3131
	goto L633
L633:
	;
	v3260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3136+int32(base.Ui32(v3224)>>(uint(int32(3))%32))&int32(536870910)))))
	if int32(base.Ui32(v3260)>>(uint(v3224&int32(15))%32))&int32(1) != 0 {
		goto L635
	} else {
		goto L636
	}
L634:
	;
	v3291 = v3275
	goto L614
L635:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v339+v3221<<(uint(int32(2))%32)))) = v3224 + v3125
	v3272 = v3221 + int32(1)
	if v3272 == int32(_a_F_sendDir_1) {
		v3340 = v3272
		goto L607
	} else {
		goto L638
	}
L636:
	;
	v3275 = v3221
	goto L637
L637:
	;
	v3277 = v3224 + int32(1)
	if v3277 != v3128 {
		v3221 = v3275
		v3224 = v3277
		goto L633
	} else {
		goto L639
	}
L638:
	;
	v3275 = v3272
	goto L637
L639:
	;
	goto L634
L640:
	;
	goto L613
L641:
	;
	F_pg_qsort(m, v339, v3340, int32(4), int32(468))
	mBase = m.M
	v3384 = m.ExcPending
	if v3384 != 0 {
		goto L4
	} else {
		goto L642
	}
L642:
	;
	v3385 = int32(0)
	if base.B2i32(v3045 == v3385)|base.B2i32(v3340 == v3385) != 0 {
		goto L643
	} else {
		goto L644
	}
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2014))) = v3340
	*(*int32)(unsafe.Add(mBase, uint32(v2016))) = v2968
	v3617 = int32(1)
	v3618 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+52))
	if v3618 == int32(-1) {
		v3701 = v3617
		goto L456
	} else {
		goto L655
	}
L644:
	;
	v3391 = v3340 & int32(3)
	v3392 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v3340) {
		goto L645
	} else {
		goto L646
	}
L645:
	;
	v3413 = v3392
	v3420 = v2017
	goto L648
L646:
	;
	v3481 = v3392
	goto L647
L647:
	;
	v3527 = v3481
	v3530 = v3392
	goto L652
L648:
	;
	v3446 = v339 + v3413<<(uint(int32(2))%32)
	v3447 = *(*int32)(unsafe.Add(mBase, uint32(v3446)))
	*(*int32)(unsafe.Add(mBase, uint32(v3446))) = v3447 - v3045
	v3450 = *(*int32)(unsafe.Add(mBase, uint32(v3446)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3446)+4)) = v3450 - v3045
	v3453 = *(*int32)(unsafe.Add(mBase, uint32(v3446)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3446)+8)) = v3453 - v3045
	v3456 = *(*int32)(unsafe.Add(mBase, uint32(v3446)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3446)+12)) = v3456 - v3045
	v3459 = int32(4)
	v3460 = v3413 + v3459
	v3462 = v3420 + v3459
	if v3462 != v3340&int32(-4) {
		v3413 = v3460
		v3420 = v3462
		goto L648
	} else {
		goto L650
	}
L649:
	;
	if v3391 == int32(0) {
		goto L643
	} else {
		goto L651
	}
L650:
	;
	goto L649
L651:
	;
	v3481 = v3460
	goto L647
L652:
	;
	v3560 = v339 + v3527<<(uint(int32(2))%32)
	v3561 = *(*int32)(unsafe.Add(mBase, uint32(v3560)))
	*(*int32)(unsafe.Add(mBase, uint32(v3560))) = v3561 - v3045
	v3564 = int32(1)
	v3567 = v3530 + v3564
	if v3567 != v3391 {
		v3527 = v3527 + v3564
		v3530 = v3567
		goto L652
	} else {
		goto L654
	}
L653:
	;
	goto L643
L654:
	;
	goto L653
L655:
	;
	v3622 = v3618 - v3045
	if base.Ui32(v3622) < base.Ui32(v2968) {
		goto L656
	} else {
		goto L657
	}
L656:
	;
	v3624 = v2968
	goto L658
L657:
	;
	v3624 = v3622
	goto L658
L658:
	;
	if base.Ui32(int32(_a_F_sendDir_1)) <= base.Ui32(v3624) {
		goto L659
	} else {
		goto L660
	}
L659:
	;
	v3627 = int32(_a_F_sendDir_1)
	goto L661
L660:
	;
	v3627 = v3624
	goto L661
L661:
	;
	if base.Ui32(v2968) < base.Ui32(v3622) {
		v3640 = v3627
		goto L589
	} else {
		goto L662
	}
L662:
	;
	if base.Ui32(v3624) < base.Ui32(int32(_a_F_sendDir_52)) {
		v3701 = v3617
		goto L456
	} else {
		goto L663
	}
L663:
	;
	v3640 = v3627
	goto L589
L664:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v3735 = m.ExcPending
	if v3735 != 0 {
		goto L4
	} else {
		goto L665
	}
L665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+4)) = v2012
	*(*int32)(unsafe.Add(mBase, uint32(v2020))) = v2011
	F_errmsg_internal(m, int32(_a_F_sendDir_53), v2020)
	mBase = m.M
	v3740 = m.ExcPending
	if v3740 != 0 {
		goto L4
	} else {
		goto L666
	}
L666:
	;
	F_errfinish(m, int32(_a_F_sendDir_54), int32(794), int32(_a_F_sendDir_55))
	mBase = m.M
	v3745 = m.ExcPending
	if v3745 != 0 {
		goto L4
	} else {
		goto L667
	}
L667:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L668:
	;
	v3748 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2256))
	v3750 = v3748 << (uint(int32(2)) % 32)
	v3752 = v3750 + int32(12)
	if v3748 == int32(0) {
		v3762 = v3752
		goto L671
	} else {
		goto L672
	}
L669:
	;
	v3783 = v1984
	v3784 = v341
	goto L670
L670:
	;
	F_pfree(m, v2007)
	mBase = m.M
	v3786 = m.ExcPending
	if v3786 != 0 {
		goto L4
	} else {
		goto L675
	}
L671:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+96)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v330)+100)) = v385
	*(*int64)(unsafe.Add(mBase, uint32(v330)+2296)) = base.I64_extend_i32_u(v3762 + v3748<<(uint(int32(13))%32))
	v3772 = v330 + int32(192)
	v3777 = F_pg_snprintf(m, v3772, int32(2048), int32(_a_F_sendDir_47), v330+int32(96))
	mBase = m.M
	v3778 = m.ExcPending
	if v3778 != 0 {
		goto L4
	} else {
		goto L674
	}
L672:
	;
	v3756 = v3752 & int32(_a_F_sendDir_56)
	if v3756 == int32(0) {
		v3762 = v3752
		goto L671
	} else {
		goto L673
	}
L673:
	;
	v3762 = v3750 - v3756 + int32(_a_F_sendDir_57)
	goto L671
L674:
	;
	v3783 = v339
	v3784 = v3772
	goto L670
L675:
	;
	v3807 = v3783
	v3810 = v3784
	goto L447
L676:
	;
	v3840 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2268))
	v3841 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2260))
	v3842 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2256))
	v3843 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2252))
	v3844 = F_sendFile(m, v320, v330+int32(2368), v3810, v330+int32(2272), int32(1), v348, v327, v3840, v3841, v326, v3842, v3807, v3843)
	mBase = m.M
	v3845 = m.ExcPending
	if v3845 != 0 {
		goto L4
	} else {
		goto L679
	}
L677:
	;
	goto L678
L678:
	;
	v3848 = *(*int64)(unsafe.Add(mBase, uint32(v330)+2296))
	v3923 = v364 + v3848 + ((v3848+int64(511))&int64(4294966784)-v3848)&int64(4294967295) + int64(512)
	goto L91
L679:
	;
	if v3844 == int32(0) {
		v3923 = v364
		goto L91
	} else {
		goto L680
	}
L680:
	;
	goto L678
L681:
	;
	if v3862 == int32(0) {
		v3923 = v364
		goto L91
	} else {
		goto L682
	}
L682:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+80)) = v330 + int32(2368)
	F_errmsg(m, int32(_a_F_sendDir_58), v330+int32(80))
	mBase = m.M
	v3873 = m.ExcPending
	if v3873 != 0 {
		goto L4
	} else {
		goto L683
	}
L683:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1545), int32(_a_F_sendDir_20))
	mBase = m.M
	v3878 = m.ExcPending
	if v3878 != 0 {
		goto L4
	} else {
		goto L684
	}
L684:
	;
	v3923 = v364
	goto L91
L685:
	;
	if v3925 != 0 {
		v329 = v3925
		v364 = v3923
		goto L89
	} else {
		goto L686
	}
L686:
	;
	goto L90
L687:
	;
	F_pfree(m, v3946)
	mBase = m.M
	v3974 = m.ExcPending
	if v3974 != 0 {
		goto L4
	} else {
		goto L690
	}
L688:
	;
	goto L689
L689:
	;
	F_FreeDir(m, v3959)
	mBase = m.M
	v3976 = m.ExcPending
	if v3976 != 0 {
		goto L4
	} else {
		goto L691
	}
L690:
	;
	goto L689
L691:
	;
	m.G0 = v3937 + int32(_a_F_sendDir_0)
	return v3971
L692:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3987 = m.ExcPending
	if v3987 != 0 {
		goto L4
	} else {
		goto L693
	}
L693:
	;
	F_errmsg(m, int32(_a_F_sendDir_59), int32(0))
	mBase = m.M
	v3991 = m.ExcPending
	if v3991 != 0 {
		goto L4
	} else {
		goto L694
	}
L694:
	;
	F_errhint(m, int32(_a_F_sendDir_60), int32(0))
	mBase = m.M
	v3995 = m.ExcPending
	if v3995 != 0 {
		goto L4
	} else {
		goto L695
	}
L695:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1285), int32(_a_F_sendDir_20))
	mBase = m.M
	v4000 = m.ExcPending
	if v4000 != 0 {
		goto L4
	} else {
		goto L696
	}
L696:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L697:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4006 = m.ExcPending
	if v4006 != 0 {
		goto L4
	} else {
		goto L698
	}
L698:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+48)) = v330 + int32(2368)
	F_errmsg(m, int32(_a_F_sendDir_61), v330+int32(48))
	mBase = m.M
	v4014 = m.ExcPending
	if v4014 != 0 {
		goto L4
	} else {
		goto L699
	}
L699:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1419), int32(_a_F_sendDir_20))
	mBase = m.M
	v4019 = m.ExcPending
	if v4019 != 0 {
		goto L4
	} else {
		goto L700
	}
L700:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L701:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v4026 = m.ExcPending
	if v4026 != 0 {
		goto L4
	} else {
		goto L702
	}
L702:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+64)) = v330 + int32(2368)
	F_errmsg(m, int32(_a_F_sendDir_62), v330-int32(-64))
	mBase = m.M
	v4034 = m.ExcPending
	if v4034 != 0 {
		goto L4
	} else {
		goto L703
	}
L703:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1424), int32(_a_F_sendDir_20))
	mBase = m.M
	v4039 = m.ExcPending
	if v4039 != 0 {
		goto L4
	} else {
		goto L704
	}
L704:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sendto(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	v4 = int32(0)
	v7 = m.Env.X__syscall_sendto(m, l0, l1, l2, v4, v4, v4)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v7) {
		*(*int32)(unsafe.Add(mBase, _c_F_sendto[0])) = int32(0) - v7
		v15 = int32(-1)
	} else {
		v15 = v7
	}
	return v15
}
func F_set_apply_error_context_origin(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_set_apply_error_context_origin[0]))
	v5 = F_MemoryContextStrdup(m, v4, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_set_apply_error_context_origin[1])) = v5
		return
	}
}
func F_set_baserel_size_estimates(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 float64
	_ = v14
	var v15 int32
	_ = v15
	var v19 float64
	_ = v19
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v30 float64
	_ = v30
	var v34 float64
	_ = v34
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*float64)(unsafe.Add(mBase, uint32(l1)+128))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v19 = F_clauselist_selectivity(m, l0, v15, v3, v3, v3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v37 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v37
	v43 = v11 + int32(16)
	if v36 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	return
L3:
	;
	v21 = base.F64_mul(v14, v19)
	if base.F64_gt(v21, float64(1e+100))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v21)&int64(9223372036854775807))) != 0 {
		v34 = float64(1e+100)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = float64(1)
	if base.F64_le(v21, v30) != 0 {
		v34 = v30
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = base.F64_nearest(v21)
	goto L1
L6:
	;
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+216)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+208)) = v80
	F_set_rel_width(m, l0, l1)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v46 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v54 = v3
	goto L9
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v54<<(uint(int32(2))%32))))
	v64 = F_cost_qual_eval_walker(m, v61, v11+int32(8))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	v67 = v54 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v67 < v68 {
		v54 = v67
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	m.G0 = v11 + int32(32)
	return
}
func F_set_cte_size_estimates(m *base.Module, l0 int32, l1 int32, l2 float64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v37 float64
	_ = v37
	var v41 float64
	_ = v41
	var v45 int32
	_ = v45
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v6 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
		v20 = v6 + v7<<(uint(int32(2))%32)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
		v20 = v13 + v14<<(uint(int32(2))%32) - int32(4)
	}
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+92)))
	if v22 != int32(1) {
		v41 = l2
	} else {
		v26 = *(*float64)(unsafe.Add(mBase, _c_F_set_cte_size_estimates[0]))
		v27 = base.F64_mul(l2, v26)
		v28 = float64(1e+100)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v27)&int64(9223372036854775807)))|base.F64_gt(v27, v28) != 0 {
			v41 = v28
		} else {
			v37 = float64(1)
			if base.F64_le(v27, v37) != 0 {
				v41 = v37
			} else {
				v41 = base.F64_nearest(v27)
			}
		}
	}
	*(*float64)(unsafe.Add(mBase, uint32(l1)+128)) = v41
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		return
	} else {
		return
	}
}
func F_set_dummy_tlist_references(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
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
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	v3 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v10 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
	return
L2:
	;
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if int32(0) < v15 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = v3
	v25 = v3
	goto L7
L5:
	;
	v82 = v3
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v82
	return
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v24<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v33 != int32(7) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v82 = v69
	goto L6
L9:
	;
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31)+8)))
	v38 = F_exprType(m, v32)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v66 = v31
	goto L11
L11:
	;
	v69 = F_lappend(m, v25, v66)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L22
	}
L12:
	;
	return
L13:
	;
	v40 = F_exprTypmod(m, v32)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v42 = F_exprCollation(m, v32)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v45 = F_makeVar(m, int32(-2), v37, v38, v40, v42, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v47 != int32(6) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+40)) = uint16(v61)
	v63 = F_flatCopyTargetEntry(m, v31)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L21
	}
L18:
	;
	v57 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+36)) = v57
	v61 = v57
	goto L17
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
	if v50 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+36)) = l1 + v50
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+40)))
	v61 = v55
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v45
	v66 = v63
	goto L11
L22:
	;
	v72 = v24 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v72 < v73 {
		v24 = v72
		v25 = v69
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L8
}
func F_set_opfuncid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2 == int32(0) {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v6 = F_get_opcode(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v6
			return
		}
	} else {
		return
	}
}
func F_set_stack_value(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v5 {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		goto L1
	}
L1:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v49
	if base.B2i32(v48 == int32(0))|base.B2i32(v49 == v48) != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
	goto L1
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
	if v15 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v13
	goto L1
L5:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v10
	goto L1
L6:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v7)
	goto L1
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v15 == v22 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v15 == v24 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v15 == v26 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v33 = l0 + int32(56)
	goto L11
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	F_pfree(m, v15)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	if v15 == v35 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	if v15 != v37 {
		v33 = v34
		goto L11
	} else {
		goto L17
	}
L17:
	;
	goto L1
L18:
	;
	return
L19:
	;
	goto L1
L20:
	;
	return
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v48 == v55 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v59 = l0 + int32(56)
	goto L23
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v63 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_pfree(m, v48)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L30
	}
L25:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+40))
	if v48 == v64 {
		goto L20
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+56))
	if v48 != v66 {
		v59 = v63
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L20
L30:
	;
	goto L20
}
func F_setseed(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 float64
	_ = v10
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v64 int64
	_ = v64
	var v69 int64
	_ = v69
	var v74 int64
	_ = v74
	var v80 int32
	_ = v80
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = base.F64_reinterpret_i64(v9)
	if base.B2i32(base.F64_gt(base.F64_abs(v10), float64(1)) == v2)&base.B2i32(base.Ui64(v9&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) == v2 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int64(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v7))) = v10
				F_errmsg(m, int32(_a_F_setseed_0), v7)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_setseed_1), int32(83), int32(_a_F_setseed_2))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
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
		v41 = int32(_a_F_setseed_3)
		v46 = base.I64_trunc_sat_f64_s(base.F64_mul(v10, float64(4.503599627370495e+15)))
		v48 = v46 + int64(4354685564936845354)
		v49 = int64(30)
		v52 = int64(-4658895280553007687)
		v53 = (int64(base.Ui64(v48)>>(uint(v49)%64)) ^ v48) * v52
		v54 = int64(27)
		v57 = int64(-7723592293110705685)
		v58 = (int64(base.Ui64(v53)>>(uint(v54)%64)) ^ v53) * v57
		v59 = int64(31)
		*(*int64)(unsafe.Add(mBase, _c_F_setseed[0])) = int64(base.Ui64(v58)>>(uint(v59)%64)) ^ v58
		v64 = v46 - int64(7046029254386353131)
		v69 = (int64(base.Ui64(v64)>>(uint(v49)%64)) ^ v64) * v52
		v74 = (int64(base.Ui64(v69)>>(uint(v54)%64)) ^ v69) * v57
		*(*int64)(unsafe.Add(mBase, _c_F_setseed[1])) = int64(base.Ui64(v74)>>(uint(v59)%64)) ^ v74
		v80 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_setseed[2])) = uint8(v80)
		m.G0 = v7 + int32(16)
		return int64(0)
	}
}
func F_setvbuf(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(-1)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v4 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(10)
	} else {
	}
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v7 | int32(64)
	return
}
func F_sha384_bytea(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = F_cryptohash_internal(m, int32(4), v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_shdepChangeDep(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v63 int32
	_ = v63
	var v92 int32
	_ = v92
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int64
	_ = v103
	var v105 int32
	_ = v105
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int64
	_ = v119
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	v15 = m.G0
	v17 = v15 - int32(320)
	m.G0 = v17
	v21 = int32(1)
	if l1 <= int32(3591) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v94 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_shdepChangeDep[0])))
	F_shdepLockAndCheckObject(m, l3, l4)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	goto L1
L3:
	;
	v92 = int32(0)
	goto L2
L4:
	;
	if base.B2i32(base.Ui32(l1-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l1-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v92 = v21
		goto L2
	} else {
		goto L25
	}
L5:
	;
	if l1 <= int32(2670) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	if l1 <= int32(_a_F_shdepChangeDep_0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	switch l1 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v92 = v21
		goto L2
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L3
	default:
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v33 = l1 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v33))|base.B2i32(int32(1)<<(uint(v33)%32)&int32(226492515) == int32(0)) != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l1-int32(2396)) {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v92 = v21
	goto L2
L13:
	;
	v92 = v21
	goto L2
L14:
	;
	if base.Ui32(l1-int32(3592)) < base.Ui32(int32(2)) {
		v92 = v21
		goto L2
	} else {
		goto L23
	}
L15:
	;
	v46 = l1 - int32(_a_F_shdepChangeDep_1)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v46))|base.B2i32(int32(1)<<(uint(v46)%32)&int32(963) == int32(0)) != 0 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	switch l1 - int32(_a_F_shdepChangeDep_2) {
	case 0, 1, 2, 3, 4, 59, 60:
		v92 = v21
		goto L2
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L3
	default:
		goto L19
	}
L18:
	;
	v92 = v21
	goto L2
L19:
	;
	if base.Ui32(l1-int32(_a_F_shdepChangeDep_3)) < base.Ui32(int32(3)) {
		v92 = v21
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v63 = l1 - int32(_a_F_shdepChangeDep_4)
	if base.Ui32(int32(15)) < base.Ui32(v63) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if int32(1)<<(uint(v63)%32)&int32(_a_F_shdepChangeDep_5) != 0 {
		v92 = v21
		goto L2
	} else {
		goto L22
	}
L22:
	;
	goto L3
L23:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l1-int32(4060)) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v92 = v21
	goto L2
L25:
	;
	goto L3
L26:
	;
	return
L27:
	;
	v98 = v17 + int32(96)
	if v92 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v103 = int64(0)
	goto L30
L29:
	;
	v103 = v94
	goto L30
L30:
	;
	F_ScanKeyInit(m, v98, int32(1), int32(3), int32(184), v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v111 = base.I64_extend_i32_u(l1)
	F_ScanKeyInit(m, v17+int32(152), int32(2), int32(3), int32(184), v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v116 = int32(3)
	v119 = base.I64_extend_i32_u(l2)
	F_ScanKeyInit(m, v17+int32(208), v116, v116, int32(184), v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	F_ScanKeyInit(m, v17+int32(264), int32(4), int32(3), int32(65), int64(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v134 = F_systable_beginscan(m, l0, int32(1232), int32(1), int32(0), int32(4), v98)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v143 = int32(0)
	goto L37
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L26
	} else {
		goto L63
	}
L37:
	;
	v150 = F_systable_getnext(m, v134)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L26
	} else {
		goto L39
	}
L38:
	;
	F_systable_endscan(m, v134)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L26
	} else {
		goto L46
	}
L39:
	;
	if v150 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+22)))
	v155 = int32(*(*int8)(unsafe.Add(mBase, uint32(v152+v153)+24)))
	if l5 != v155 {
		goto L37
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L38
L43:
	;
	if v143 != 0 {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v157 = F_heap_copytuple(m, v150)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	v143 = v157
	goto L37
L46:
	;
	goto L49
L47:
	;
	m.G0 = v17 + int32(320)
	return
L48:
	;
	F_pfree(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L26
	} else {
		goto L62
	}
L49:
	;
	if base.B2i32(base.B2i32(l3 == int32(2613))|base.B2i32(base.Ui32(int32(_a_F_shdepChangeDep_6)) < base.Ui32(l4)) == int32(0))&((base.B2i32(l3 != int32(2615))|base.B2i32(l4 != int32(2200)))&base.B2i32(l3 != int32(1262))) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v143 == int32(0) {
		goto L47
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v143 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	F_simple_heap_delete(m, l0, v143+int32(4))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L26
	} else {
		goto L54
	}
L54:
	;
	v220 = v143
	goto L48
L55:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+22)))
	v186 = v184 + v185
	*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v186)+16)) = l3
	F_CatalogTupleUpdate(m, l0, v143+int32(4), v143)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L26
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v103
	v198 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+27)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v198
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = base.I64_extend_i32_u(l5)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+72)) = base.I64_extend_i32_u(l4)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = base.I64_extend_i32_u(l3)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v213 = F_heap_form_tuple(m, v208, v17+int32(32), v17+int32(24))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L26
	} else {
		goto L59
	}
L58:
	;
	v220 = v143
	goto L48
L59:
	;
	F_CatalogTupleInsert(m, l0, v213)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L26
	} else {
		goto L60
	}
L60:
	;
	if v213 == int32(0) {
		goto L47
	} else {
		goto L61
	}
L61:
	;
	v220 = v213
	goto L48
L62:
	;
	goto L47
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l1
	F_errmsg_internal(m, int32(_a_F_shdepChangeDep_7), v17)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L26
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_shdepChangeDep_8), int32(255), int32(_a_F_shdepChangeDep_9))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L26
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_shdepLockAndCheckObject(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	F_LockSharedObject(m, l0, l1, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		switch l0 - int32(1260) {
		case 0:
			v15 = int64(0)
			v18 = F_SearchSysCacheExists(m, int32(11), base.I64_extend_i32_u(l1), v15, v15, v15)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				if v18 != 0 {
					m.G0 = v6 - int32(-64)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l1
							F_errmsg(m, int32(_a_F_shdepLockAndCheckObject_0), v4+int32(-48))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_shdepLockAndCheckObject_1), int32(1222), int32(_a_F_shdepLockAndCheckObject_2))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_shdepLockAndCheckObject_3), v6)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_shdepLockAndCheckObject_1), int32(1255), int32(_a_F_shdepLockAndCheckObject_2))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 2:
			v38 = F_get_database_name(m, l1)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				if v38 != 0 {
					v93 = v38
					F_pfree(m, v93)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						m.G0 = v6 - int32(-64)
						return
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = l1
							F_errmsg(m, int32(_a_F_shdepLockAndCheckObject_4), v4+int32(-16))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_shdepLockAndCheckObject_1), int32(1248), int32(_a_F_shdepLockAndCheckObject_2))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
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
		default:
			if l0 == int32(1213) {
				v73 = F_get_tablespace_name(m, l1)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					if v73 != 0 {
						v93 = v73
						F_pfree(m, v93)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return
						} else {
							m.G0 = v6 - int32(-64)
							return
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l1
								F_errmsg(m, int32(_a_F_shdepLockAndCheckObject_5), v4+int32(-32))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_shdepLockAndCheckObject_1), int32(1234), int32(_a_F_shdepLockAndCheckObject_2))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
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
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg_internal(m, int32(_a_F_shdepLockAndCheckObject_3), v6)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_shdepLockAndCheckObject_1), int32(1255), int32(_a_F_shdepLockAndCheckObject_2))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
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
func F_shell_out(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			F_errmsg(m, int32(_a_F_shell_out_0), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_shell_out_1), int32(317), int32(_a_F_shell_out_2))
				v19 = m.ExcPending
				if v19 != 0 {
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
func F_signconsistent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = F_execute(m, l0+v5<<(uint(int32(3))%32), l1, l2, l3, int32(_a_F_signconsistent_0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_sjis_to_utf8(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14231(m, l0, int32(35), v3, v3, v3, int32(_a_F_sjis_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_slice_from_s(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	v12 = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v13 < int32(0) {
		v82 = v12
		return v82
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v16 < v13 {
			v82 = v12
			return v82
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v18 < v16 {
				v82 = v12
				return v82
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v20-int32(4))))
				if v23 < v18 {
					v82 = v12
					return v82
				} else {
					v26 = v13 - v16 + l1
					if v26 == int32(0) {
						v67 = v20
						if l1 != 0 {
							base.MemoryCopy(m, v67+v13, l2, l1)
						} else {
						}
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v74 + l1
						v82 = int32(0)
						return v82
					} else {
						v29 = v26 + v23
						v31 = v20 - int32(8)
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
						if v32 < v29 {
							v36 = F_repalloc(m, v31, v29+int32(29))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								if v36 == int32(0) {
									v82 = v12
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v36))) = v29 + int32(20)
									v46 = v36 + int32(8)
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v46
									v48 = v46
									v49 = v23 - v16
									if v49 != 0 {
										v50 = v48 + v16
										base.MemoryCopy(m, v50+v26, v50, v49)
									} else {
									}
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v54-int32(4)))) = v29
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v58 + v26
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v16 <= v61 {
										v65 = v26 + v61
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v65
										v67 = v54
									} else {
										if v61 <= v13 {
											v67 = v54
										} else {
											v65 = v13
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v65
											v67 = v54
										}
									}
									if l1 != 0 {
										base.MemoryCopy(m, v67+v13, l2, l1)
									} else {
									}
									v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v74 + l1
									v82 = int32(0)
								}
								return v82
							}
						} else {
							v48 = v20
							v49 = v23 - v16
							if v49 != 0 {
								v50 = v48 + v16
								base.MemoryCopy(m, v50+v26, v50, v49)
							} else {
							}
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v54-int32(4)))) = v29
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v58 + v26
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v16 <= v61 {
								v65 = v26 + v61
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v65
								v67 = v54
							} else {
								if v61 <= v13 {
									v67 = v54
								} else {
									v65 = v13
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v65
									v67 = v54
								}
							}
							if l1 != 0 {
								base.MemoryCopy(m, v67+v13, l2, l1)
							} else {
							}
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v74 + l1
							v82 = int32(0)
							return v82
						}
					}
				}
			}
		}
	}
}
func F_smgrnblocks(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrnblocks[0])))
	if v5 == int32(1) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+20))
		if v11 != int32(-1) {
			v35 = v11
			return v35
		} else {
			v15 = int32(_a_F_smgrnblocks_0)
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1])) = v17 + int32(1)
			v24 = F_mdnblocks(m, l0, l1)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+20)) = v24
				v29 = int32(_a_F_smgrnblocks_0)
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1])) = v31 - int32(1)
				v35 = v24
				return v35
			}
		}
	} else {
		v15 = int32(_a_F_smgrnblocks_0)
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1])) = v17 + int32(1)
		v24 = F_mdnblocks(m, l0, l1)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+20)) = v24
			v29 = int32(_a_F_smgrnblocks_0)
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1])) = v31 - int32(1)
			v35 = v24
			return v35
		}
	}
}
func F_smgrreadv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int64
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v206 int32
	_ = v206
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int64
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	v5 = int32(0)
	v14 = int32(_a_F_smgrreadv_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_smgrreadv[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrreadv[0])) = v16 + int32(1)
	v20 = m.G0
	v22 = v20 - int32(1072)
	m.G0 = v22
	v26 = F__mdfd_getseg(m, l0, l1, l2, v5, int32(9))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = int32(1)
	v32 = l2 & int32(_a_F_smgrreadv_1)
	v33 = int32(_a_F_smgrreadv_2) - v32
	if base.Ui32(v29) < base.Ui32(v33) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v22 + int32(1072)
	v382 = int32(_a_F_smgrreadv_0)
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_smgrreadv[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrreadv[0])) = v384 - int32(1)
	return
L4:
	;
	v36 = v29
	goto L6
L5:
	;
	v36 = v33
	goto L6
L6:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v36) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v39 = int32(128)
	goto L9
L8:
	;
	v39 = v36
	goto L9
L9:
	;
	if v39 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = base.I64_extend_i32_u(v32 << (uint(int32(13)) % 32))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = int32(_a_F_smgrreadv_3)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v45
	v49 = int32(1)
	if base.Ui32(int32(2)) <= base.Ui32(v36) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L74
	}
L13:
	;
	v55 = v45
	v56 = v22 + int32(48)
	v58 = v49
	v63 = int32(1)
	v66 = v5
	goto L16
L14:
	;
	v112 = v49
	goto L15
L15:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v126 = F_FileReadV(m, v122, v22+int32(48), v112, v44, int32(167772183))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L27
	}
L16:
	;
	v70 = l3 + v63<<(uint(int32(2))%32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v71 == v55+v72 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v112 = v104
	goto L15
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v88 != v85+v89 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v72 - int32(-8192)
	v85 = v55
	v86 = v56
	v87 = v58
	goto L18
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = int32(_a_F_smgrreadv_3)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v71
	v85 = v71
	v86 = v56 + int32(8)
	v87 = v58 + int32(1)
	goto L18
L22:
	;
	v105 = int32(2)
	v108 = v66 + v105
	if v108 != 0 {
		v55 = v102
		v56 = v103
		v58 = v104
		v63 = v63 + v105
		v66 = v108
		goto L16
	} else {
		goto L26
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = int32(_a_F_smgrreadv_3)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = v88
	v102 = v88
	v103 = v86 + int32(8)
	v104 = v87 + int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v89 - int32(-8192)
	v102 = v85
	v103 = v86
	v104 = v87
	goto L22
L26:
	;
	goto L17
L27:
	;
	if int32(0) <= v126 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v131 = int32(0)
	v132 = v126
	v134 = v112
	v143 = v44
	goto L31
L29:
	;
	goto L30
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L69
	}
L31:
	;
	if v132 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrreadv[1])))
	if v147 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v271 = v131 + v132
	if v271 == int32(_a_F_smgrreadv_3) {
		goto L3
	} else {
		goto L57
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L52
	}
L37:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrreadv[2])))
	if v151&int32(1) == int32(0) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v157 = int32(base.Ui32(v131) >> (uint(int32(13)) % 32))
	if v157 != 0 {
		goto L3
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v161 = (int32(1) - v157) & int32(3)
	if v161 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v164 = v157
	v166 = int32(0)
	goto L45
L43:
	;
	v189 = v157
	goto L44
L44:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v157-int32(1)) {
		goto L3
	} else {
		goto L48
	}
L45:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l3+v164<<(uint(int32(2))%32))))
	base.MemoryFill(m, v179, int32(0), int32(_a_F_smgrreadv_3))
	v183 = int32(1)
	v184 = v164 + v183
	v186 = v166 + v183
	if v186 != v161 {
		v164 = v184
		v166 = v186
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v189 = v184
	goto L44
L47:
	;
	goto L46
L48:
	;
	v206 = v189
	goto L49
L49:
	;
	v220 = l3 + v206<<(uint(int32(2))%32)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v222 = int32(0)
	v223 = int32(_a_F_smgrreadv_3)
	base.MemoryFill(m, v221, v222, v223)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	base.MemoryFill(m, v225, v222, v223)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	base.MemoryFill(m, v229, v222, v223)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	base.MemoryFill(m, v233, v222, v223)
	v238 = v206 + int32(4)
	if v238 != int32(1) {
		v206 = v238
		goto L49
	} else {
		goto L51
	}
L50:
	;
	goto L3
L51:
	;
	goto L50
L52:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_smgrreadv[3]))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v248*int32(48))+32))
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = int32(_a_F_smgrreadv_3)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = l2
	F_errmsg(m, int32(_a_F_smgrreadv_4), v22+int32(16))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_smgrreadv_5), int32(973), int32(_a_F_smgrreadv_6))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	v275 = v22 + int32(48)
	v278 = v275
	v279 = v134
	v280 = v132
	goto L60
L58:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v310 = v143 + base.I64_extend_i32_u(v132)
	v312 = F_FileReadV(m, v308, v275, v307, v310, int32(167772183))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L67
	}
L59:
	;
	if v275 == v278 {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if base.Ui32(v280) < base.Ui32(v282) {
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v307 = int32(0)
	goto L58
L62:
	;
	v288 = v279 - int32(1)
	if v288 != 0 {
		v278 = v278 + int32(8)
		v279 = v288
		v280 = v280 - v282
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v297 + v280
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v275)+4)) = v300 - v280
	v307 = v279
	goto L58
L65:
	;
	v292 = v279 << (uint(int32(3)) % 32)
	if v292 == int32(0) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	base.MemoryCopy(m, v275, v278, v292)
	goto L64
L67:
	;
	if int32(0) <= v312 {
		v131 = v271
		v132 = v312
		v134 = v307
		v143 = v310
		goto L31
	} else {
		goto L68
	}
L68:
	;
	goto L32
L69:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_smgrreadv[3]))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337+v335*int32(48))+32))
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = l2
	F_errmsg(m, int32(_a_F_smgrreadv_7), v22)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_smgrreadv_5), int32(924), int32(_a_F_smgrreadv_6))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errmsg_internal(m, int32(_a_F_smgrreadv_8), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_smgrreadv_5), int32(886), int32(_a_F_smgrreadv_6))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_smgrshutdown(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = int32(_a_F_smgrshutdown_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_smgrshutdown[0]))
	v6 = int32(1)
	v7 = v5 + v6
	*(*int32)(unsafe.Add(mBase, _c_F_smgrshutdown[0])) = v7
	*(*int32)(unsafe.Add(mBase, _c_F_smgrshutdown[0])) = v7 - v6
	return
}
func F_smgrunpin(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v6 = v4 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v6
	if v6 == int32(0) {
		v11 = l0 + int32(76)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_smgrunpin[0]))
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_smgrunpin[1]))
			v20 = v15
		} else {
			v17 = int32(_a_F_smgrunpin_0)
			*(*int32)(unsafe.Add(mBase, _c_F_smgrunpin[0])) = v17
			v20 = v17
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v20
		v22 = int32(_a_F_smgrunpin_0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v11
		*(*int32)(unsafe.Add(mBase, _c_F_smgrunpin[1])) = v11
	} else {
	}
	return
}
func F_spcache_insert(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int64
	_ = v72
	var v77 int64
	_ = v77
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v98 int64
	_ = v98
	var v105 int64
	_ = v105
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v116 int64
	_ = v116
	var v121 int64
	_ = v121
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v168 int64
	_ = v168
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v191 int64
	_ = v191
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int64
	_ = v283
	var v287 int32
	_ = v287
	var v294 int64
	_ = v294
	var v299 int64
	_ = v299
	var v303 int64
	_ = v303
	var v306 int64
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int64
	_ = v313
	var v316 int32
	_ = v316
	var v320 int64
	_ = v320
	var v327 int64
	_ = v327
	var v332 int32
	_ = v332
	var v333 int64
	_ = v333
	var v338 int64
	_ = v338
	var v343 int64
	_ = v343
	var v348 int32
	_ = v348
	var v350 int64
	_ = v350
	var v354 int64
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int64
	_ = v361
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v369 int64
	_ = v369
	var v374 int32
	_ = v374
	var v375 int64
	_ = v375
	var v378 int32
	_ = v378
	var v379 int64
	_ = v379
	var v384 int32
	_ = v384
	var v385 int64
	_ = v385
	var v390 int64
	_ = v390
	var v396 int64
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int64
	_ = v401
	var v413 int64
	_ = v413
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v446 int64
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int64
	_ = v455
	var v457 int64
	_ = v457
	var v460 int64
	_ = v460
	var v461 int64
	_ = v461
	var v471 int64
	_ = v471
	var v476 int32
	_ = v476
	var v477 int64
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int64
	_ = v486
	var v496 int64
	_ = v496
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int64
	_ = v539
	var v542 int32
	_ = v542
	var v549 int64
	_ = v549
	var v554 int64
	_ = v554
	var v558 int64
	_ = v558
	var v561 int64
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v568 int64
	_ = v568
	var v571 int32
	_ = v571
	var v575 int64
	_ = v575
	var v582 int64
	_ = v582
	var v587 int32
	_ = v587
	var v588 int64
	_ = v588
	var v593 int64
	_ = v593
	var v598 int64
	_ = v598
	var v603 int32
	_ = v603
	var v605 int64
	_ = v605
	var v609 int64
	_ = v609
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v616 int64
	_ = v616
	var v622 int32
	_ = v622
	var v623 int64
	_ = v623
	var v624 int64
	_ = v624
	var v629 int32
	_ = v629
	var v630 int64
	_ = v630
	var v633 int32
	_ = v633
	var v634 int64
	_ = v634
	var v639 int32
	_ = v639
	var v640 int64
	_ = v640
	var v645 int64
	_ = v645
	var v651 int64
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int64
	_ = v656
	var v668 int64
	_ = v668
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int64
	_ = v706
	var v714 int64
	_ = v714
	var v719 int64
	_ = v719
	var v723 int64
	_ = v723
	var v726 int64
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v733 int64
	_ = v733
	var v736 int32
	_ = v736
	var v740 int64
	_ = v740
	var v747 int64
	_ = v747
	var v752 int32
	_ = v752
	var v753 int64
	_ = v753
	var v758 int64
	_ = v758
	var v763 int64
	_ = v763
	var v768 int32
	_ = v768
	var v770 int64
	_ = v770
	var v774 int64
	_ = v774
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v781 int64
	_ = v781
	var v787 int32
	_ = v787
	var v788 int64
	_ = v788
	var v789 int64
	_ = v789
	var v794 int32
	_ = v794
	var v795 int64
	_ = v795
	var v798 int32
	_ = v798
	var v799 int64
	_ = v799
	var v804 int32
	_ = v804
	var v805 int64
	_ = v805
	var v810 int64
	_ = v810
	var v816 int64
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int64
	_ = v821
	var v833 int64
	_ = v833
	var v841 int32
	_ = v841
	var v850 int32
	_ = v850
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int64
	_ = v865
	var v867 int64
	_ = v867
	var v869 int64
	_ = v869
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v976 int64
	_ = v976
	var v979 int32
	_ = v979
	var v986 int64
	_ = v986
	var v991 int64
	_ = v991
	var v995 int64
	_ = v995
	var v998 int64
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1005 int64
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1012 int64
	_ = v1012
	var v1019 int64
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1025 int64
	_ = v1025
	var v1030 int64
	_ = v1030
	var v1035 int64
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1042 int64
	_ = v1042
	var v1046 int64
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1053 int64
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1060 int64
	_ = v1060
	var v1061 int64
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1067 int64
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1071 int64
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1077 int64
	_ = v1077
	var v1082 int64
	_ = v1082
	var v1088 int64
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int64
	_ = v1093
	var v1105 int64
	_ = v1105
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1150 int64
	_ = v1150
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1202 int64
	_ = v1202
	var v1204 int64
	_ = v1204
	var v1206 int64
	_ = v1206
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1231 int64
	_ = v1231
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1280 int32
	_ = v1280
	var v1294 int64
	_ = v1294
	var v1314 int32
	_ = v1314
	var v1319 int32
	_ = v1319
	var v1337 int32
	_ = v1337
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1371 int32
	_ = v1371
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_insert[0]))
	if v22 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L67
	} else {
		goto L303
	}
L2:
	;
	m.G0 = v19 + int32(48)
	return v1337
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l0
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v19)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v58
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_insert[1]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v65 = v19 + int32(24)
	v72 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v65)+4)))
	v77 = (int64(base.Ui64(v72)>>(uint(int64(23))%64)) ^ v72) * int64(2388976653695081527)
	v81 = int64(-8645972361240307355)
	v84 = (v77 ^ int64(base.Ui64(v77)>>(uint(int64(47))%64)) ^ v81) * v81
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v86 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v25 != l1 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v30 == int32(0))|base.B2i32(v30 != v33) != 0 {
		v51 = v30
		v52 = v33
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v51-v52 == int32(0) {
		v1337 = v22
		goto L2
	} else {
		goto L13
	}
L7:
	;
	goto L6
L8:
	;
	v36 = v27
	v37 = l0
	goto L9
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v41 == int32(0) {
		v51 = v41
		v52 = v40
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v51 = v41
	v52 = v40
	goto L7
L11:
	;
	v44 = int32(1)
	if v41 == v40 {
		v36 = v36 + v44
		v37 = v37 + v44
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	goto L3
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spcache_insert[0])) = v1319
	v1337 = v1319
	goto L2
L15:
	;
	v199 = v63 & base.I32_wrap_i64(int64(base.Ui64(v191)>>(uint(int64(47))%64))^v191-int64(base.Ui64(v191)>>(uint(int64(32))%64)))
	v202 = v62 + v199*int32(24)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+22)))
	if v203 != 0 {
		goto L50
	} else {
		goto L51
	}
L16:
	;
	v191 = (base.I64_extend_i32_s(v177-v85) + int64(base.Ui64(v179)>>(uint(int64(23))%64)) ^ v179) * int64(2388976653695081527)
	goto L15
L17:
	;
	v177 = v85
	v179 = v84
	goto L16
L18:
	;
	goto L19
L19:
	;
	v89 = v85
	v91 = v84
	v94 = v86
	goto L20
L20:
	;
	v98 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v98 == int64(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v177 = v175
	v179 = v174
	goto L16
L22:
	;
	v168 = (v163 ^ int64(base.Ui64(v163)>>(uint(int64(23))%64))) * int64(2388976653695081527)
	v174 = (v91 ^ int64(base.Ui64(v168)>>(uint(int64(47))%64)) ^ v168) * int64(-8645972361240307355)
	v175 = v89 + v162
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v176 != 0 {
		v89 = v175
		v91 = v174
		v94 = v176
		goto L20
	} else {
		goto L49
	}
L23:
	;
	v162 = v156
	v163 = base.I64_extend8_s(base.I64_extend_i32_u(v94)) | v157
	goto L22
L24:
	;
	v156 = int32(1)
	v157 = int64(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v105 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+2)))
	if v105 == int64(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v156 = v152
	v157 = v98<<(uint(int64(8))%64) | v153
	goto L23
L28:
	;
	v152 = int32(2)
	v153 = int64(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+3)))
	if v110 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v111 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v111 == int64(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	v152 = int32(3)
	v153 = v105 << (uint(int64(16)) % 64)
	goto L27
L34:
	;
	v147 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v89))))
	v162 = v145
	v163 = v146 | v147
	goto L22
L35:
	;
	v145 = int32(4)
	v146 = int64(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v116 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+5)))
	if v116 == int64(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v145 = v138
	v146 = v139 | v111<<(uint(int64(32))%64)
	goto L34
L39:
	;
	v138 = int32(5)
	v139 = int64(0)
	goto L38
L40:
	;
	goto L41
L41:
	;
	v121 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+6)))
	if v121 == int64(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v138 = v133
	v139 = v132 | v116<<(uint(int64(40))%64)
	goto L38
L43:
	;
	v132 = int64(0)
	v133 = int32(6)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+7)))
	if v126 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
	v162 = int32(8)
	v163 = v128
	goto L22
L47:
	;
	goto L48
L48:
	;
	v132 = v121 << (uint(int64(48)) % 64)
	v133 = int32(7)
	goto L42
L49:
	;
	goto L21
L50:
	;
	v206 = v202
	v208 = v199
	goto L53
L51:
	;
	goto L52
L52:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_insert[2]))
	v276 = F_MemoryContextStrdup(m, v275, l0)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L67
	} else {
		goto L68
	}
L53:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if l1 == v220 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L52
L55:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v225 == int32(0))|base.B2i32(v225 != v228) != 0 {
		v246 = v225
		v247 = v228
		goto L59
	} else {
		goto L60
	}
L56:
	;
	goto L57
L57:
	;
	v253 = (v208 + int32(1)) & v63
	v256 = v62 + v253*int32(24)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+22)))
	if v257 != 0 {
		v206 = v256
		v208 = v253
		goto L53
	} else {
		goto L66
	}
L58:
	;
	if v246-v247 == int32(0) {
		v1319 = v206
		goto L14
	} else {
		goto L65
	}
L59:
	;
	goto L58
L60:
	;
	v231 = v222
	v232 = l0
	goto L61
L61:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+1)))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
	if v236 == int32(0) {
		v246 = v236
		v247 = v235
		goto L59
	} else {
		goto L63
	}
L62:
	;
	v246 = v236
	v247 = v235
	goto L59
L63:
	;
	v239 = int32(1)
	if v236 == v235 {
		v231 = v231 + v239
		v232 = v232 + v239
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L57
L66:
	;
	goto L54
L67:
	;
	return int32(0)
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v276
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_insert[1]))
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v19)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v283
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v283
	v287 = v19 + int32(16)
	v294 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v287)+4)))
	v299 = (int64(base.Ui64(v294)>>(uint(int64(23))%64)) ^ v294) * int64(2388976653695081527)
	v303 = int64(-8645972361240307355)
	v306 = (v299 ^ int64(base.Ui64(v299)>>(uint(int64(47))%64)) ^ v303) * v303
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	if v308 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v423 = base.I32_wrap_i64(int64(base.Ui64(v283) >> (uint(int64(32)) % 64)))
	v424 = base.I32_wrap_i64(v283)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v282)+16))
	v430 = base.B2i32(base.Ui32(v425) < base.Ui32(v426))
	goto L104
L70:
	;
	v413 = (base.I64_extend_i32_s(v399-v307) + int64(base.Ui64(v401)>>(uint(int64(23))%64)) ^ v401) * int64(2388976653695081527)
	goto L69
L71:
	;
	v399 = v307
	v401 = v306
	goto L70
L72:
	;
	goto L73
L73:
	;
	v311 = v307
	v313 = v306
	v316 = v308
	goto L74
L74:
	;
	v320 = int64(*(*int8)(unsafe.Add(mBase, uint32(v311)+1)))
	if v320 == int64(0) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v399 = v397
	v401 = v396
	goto L70
L76:
	;
	v390 = (v385 ^ int64(base.Ui64(v385)>>(uint(int64(23))%64))) * int64(2388976653695081527)
	v396 = (v313 ^ int64(base.Ui64(v390)>>(uint(int64(47))%64)) ^ v390) * int64(-8645972361240307355)
	v397 = v311 + v384
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397))))
	if v398 != 0 {
		v311 = v397
		v313 = v396
		v316 = v398
		goto L74
	} else {
		goto L103
	}
L77:
	;
	v384 = v378
	v385 = base.I64_extend8_s(base.I64_extend_i32_u(v316)) | v379
	goto L76
L78:
	;
	v378 = int32(1)
	v379 = int64(0)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v327 = int64(*(*int8)(unsafe.Add(mBase, uint32(v311)+2)))
	if v327 == int64(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v378 = v374
	v379 = v320<<(uint(int64(8))%64) | v375
	goto L77
L82:
	;
	v374 = int32(2)
	v375 = int64(0)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+3)))
	if v332 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v333 = int64(*(*int8)(unsafe.Add(mBase, uint32(v311)+4)))
	if v333 == int64(0) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	goto L87
L87:
	;
	v374 = int32(3)
	v375 = v327 << (uint(int64(16)) % 64)
	goto L81
L88:
	;
	v369 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v311))))
	v384 = v367
	v385 = v368 | v369
	goto L76
L89:
	;
	v367 = int32(4)
	v368 = int64(0)
	goto L88
L90:
	;
	goto L91
L91:
	;
	v338 = int64(*(*int8)(unsafe.Add(mBase, uint32(v311)+5)))
	if v338 == int64(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v367 = v360
	v368 = v361 | v333<<(uint(int64(32))%64)
	goto L88
L93:
	;
	v360 = int32(5)
	v361 = int64(0)
	goto L92
L94:
	;
	goto L95
L95:
	;
	v343 = int64(*(*int8)(unsafe.Add(mBase, uint32(v311)+6)))
	if v343 == int64(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v360 = v355
	v361 = v354 | v338<<(uint(int64(40))%64)
	goto L92
L97:
	;
	v354 = int64(0)
	v355 = int32(6)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+7)))
	if v348 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v311)))
	v384 = int32(8)
	v385 = v350
	goto L76
L101:
	;
	goto L102
L102:
	;
	v354 = v343 << (uint(int64(48)) % 64)
	v355 = int32(7)
	goto L96
L103:
	;
	goto L75
L104:
	;
	if v430 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	v1314 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+16)) = v1314
	v430 = v1314
	goto L104
L107:
	;
	v1294 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1280)+14)) = v1294
	*(*int64)(unsafe.Add(mBase, uint32(v1280)+8)) = v1294
	v1319 = v1280
	goto L14
L108:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v1271 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+8)) = v1270 + v1271
	*(*uint8)(unsafe.Add(mBase, uint32(v1256)+22)) = uint8(v1271)
	*(*int32)(unsafe.Add(mBase, uint32(v1256)+4)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v1256))) = v424
	v1280 = v1256
	goto L107
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L67
	} else {
		goto L300
	}
L110:
	;
	v446 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
	if v446 == int64(4294967296) {
		goto L109
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v282)+20))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v921 = v920 & base.I32_wrap_i64(int64(base.Ui64(v413)>>(uint(int64(47))%64))^v413-int64(base.Ui64(v413)>>(uint(int64(32))%64)))
	v924 = v919 + v921*int32(24)
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924)+22)))
	if v925 == int32(0) {
		v1256 = v924
		goto L108
	} else {
		goto L224
	}
L113:
	;
	v449 = int32(0)
	v451 = m.G0
	v453 = v451 - int32(16)
	m.G0 = v453
	v455 = int64(2)
	v457 = v446 << (uint(int64(1)) % 64)
	if base.Ui64(v457) <= base.Ui64(v455) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v430 = int32(1)
	goto L104
L115:
	;
	v460 = v455
	goto L117
L116:
	;
	v460 = v457
	goto L117
L117:
	;
	v461 = int64(1)
	if v460&(v460-v461) == int64(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v471 = v460
	goto L120
L119:
	;
	v471 = v461 << (uint(int64(64)-base.I64_clz(v460)) % 64)
	goto L120
L120:
	;
	if base.Ui64(v471*int64(24)) < base.Ui64(int64(2147483647)) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v282)+20))
	v477 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v282)+24))
	v483 = F_MemoryContextAllocExtended(m, v478, base.I32_wrap_i64(v471)*int32(24), int32(5))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L67
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	goto L1
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+20)) = v483
	v486 = int64(1)
	if v471&(v471-v486) == int64(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v496 = v471
	goto L127
L126:
	;
	v496 = v486 << (uint(int64(64)-base.I64_clz(v471)) % 64)
	goto L127
L127:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v496*int64(24)) {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v282))) = v496
	v504 = base.I32_wrap_i64(v496) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+12)) = v504
	if v496 == int64(4294967296) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v513 = int32(-85899346)
	goto L131
L130:
	;
	v513 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v496), float64(0.9)))
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+16)) = v513
	if v477 != int64(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v518 = v449
	goto L136
L133:
	;
	goto L134
L134:
	;
	F_pfree(m, v476)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L67
	} else {
		goto L223
	}
L135:
	;
	v685 = v683
	v689 = v449
	goto L176
L136:
	;
	v535 = v476 + v518*int32(24)
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535)+22)))
	if v536 != int32(1) {
		v683 = v518
		goto L135
	} else {
		goto L138
	}
L137:
	;
	v683 = int32(0)
	goto L135
L138:
	;
	v539 = *(*int64)(unsafe.Add(mBase, uint32(v535)))
	*(*int64)(unsafe.Add(mBase, uint32(v453)+8)) = v539
	v542 = v453 + int32(8)
	v549 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v542)+4)))
	v554 = (int64(base.Ui64(v549)>>(uint(int64(23))%64)) ^ v549) * int64(2388976653695081527)
	v558 = int64(-8645972361240307355)
	v561 = (v554 ^ int64(base.Ui64(v554)>>(uint(int64(47))%64)) ^ v558) * v558
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v542)))
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562))))
	if v563 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	if base.I32_wrap_i64(int64(base.Ui64(v668)>>(uint(int64(47))%64))^v668-int64(base.Ui64(v668)>>(uint(int64(32))%64)))&v504 == v518 {
		v683 = v518
		goto L135
	} else {
		goto L174
	}
L140:
	;
	v668 = (base.I64_extend_i32_s(v654-v562) + int64(base.Ui64(v656)>>(uint(int64(23))%64)) ^ v656) * int64(2388976653695081527)
	goto L139
L141:
	;
	v654 = v562
	v656 = v561
	goto L140
L142:
	;
	goto L143
L143:
	;
	v566 = v562
	v568 = v561
	v571 = v563
	goto L144
L144:
	;
	v575 = int64(*(*int8)(unsafe.Add(mBase, uint32(v566)+1)))
	if v575 == int64(0) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v654 = v652
	v656 = v651
	goto L140
L146:
	;
	v645 = (v640 ^ int64(base.Ui64(v640)>>(uint(int64(23))%64))) * int64(2388976653695081527)
	v651 = (v568 ^ int64(base.Ui64(v645)>>(uint(int64(47))%64)) ^ v645) * int64(-8645972361240307355)
	v652 = v566 + v639
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652))))
	if v653 != 0 {
		v566 = v652
		v568 = v651
		v571 = v653
		goto L144
	} else {
		goto L173
	}
L147:
	;
	v639 = v633
	v640 = base.I64_extend8_s(base.I64_extend_i32_u(v571)) | v634
	goto L146
L148:
	;
	v633 = int32(1)
	v634 = int64(0)
	goto L147
L149:
	;
	goto L150
L150:
	;
	v582 = int64(*(*int8)(unsafe.Add(mBase, uint32(v566)+2)))
	if v582 == int64(0) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v633 = v629
	v634 = v575<<(uint(int64(8))%64) | v630
	goto L147
L152:
	;
	v629 = int32(2)
	v630 = int64(0)
	goto L151
L153:
	;
	goto L154
L154:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566)+3)))
	if v587 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v588 = int64(*(*int8)(unsafe.Add(mBase, uint32(v566)+4)))
	if v588 == int64(0) {
		goto L159
	} else {
		goto L160
	}
L156:
	;
	goto L157
L157:
	;
	v629 = int32(3)
	v630 = v582 << (uint(int64(16)) % 64)
	goto L151
L158:
	;
	v624 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v566))))
	v639 = v622
	v640 = v623 | v624
	goto L146
L159:
	;
	v622 = int32(4)
	v623 = int64(0)
	goto L158
L160:
	;
	goto L161
L161:
	;
	v593 = int64(*(*int8)(unsafe.Add(mBase, uint32(v566)+5)))
	if v593 == int64(0) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v622 = v615
	v623 = v616 | v588<<(uint(int64(32))%64)
	goto L158
L163:
	;
	v615 = int32(5)
	v616 = int64(0)
	goto L162
L164:
	;
	goto L165
L165:
	;
	v598 = int64(*(*int8)(unsafe.Add(mBase, uint32(v566)+6)))
	if v598 == int64(0) {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v615 = v610
	v616 = v609 | v593<<(uint(int64(40))%64)
	goto L162
L167:
	;
	v609 = int64(0)
	v610 = int32(6)
	goto L166
L168:
	;
	goto L169
L169:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566)+7)))
	if v603 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v605 = *(*int64)(unsafe.Add(mBase, uint32(v566)))
	v639 = int32(8)
	v640 = v605
	goto L146
L171:
	;
	goto L172
L172:
	;
	v609 = v598 << (uint(int64(48)) % 64)
	v610 = int32(7)
	goto L166
L173:
	;
	goto L145
L174:
	;
	v679 = v518 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v679)) < base.Ui64(v477) {
		v518 = v679
		goto L136
	} else {
		goto L175
	}
L175:
	;
	goto L137
L176:
	;
	v702 = v476 + v685*int32(24)
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702)+22)))
	if v703 == int32(1) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	goto L134
L178:
	;
	v706 = *(*int64)(unsafe.Add(mBase, uint32(v702)))
	*(*int64)(unsafe.Add(mBase, uint32(v453))) = v706
	v714 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v453)+4)))
	v719 = (int64(base.Ui64(v714)>>(uint(int64(23))%64)) ^ v714) * int64(2388976653695081527)
	v723 = int64(-8645972361240307355)
	v726 = (v719 ^ int64(base.Ui64(v719)>>(uint(int64(47))%64)) ^ v723) * v723
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727))))
	if v728 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L179:
	;
	goto L180
L180:
	;
	v888 = v685 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v888)) < base.Ui64(v477) {
		goto L219
	} else {
		goto L220
	}
L181:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v850 = base.I32_wrap_i64(int64(base.Ui64(v833)>>(uint(int64(47))%64)) ^ v833 - int64(base.Ui64(v833)>>(uint(int64(32))%64)))
	goto L216
L182:
	;
	v833 = (base.I64_extend_i32_s(v819-v727) + int64(base.Ui64(v821)>>(uint(int64(23))%64)) ^ v821) * int64(2388976653695081527)
	goto L181
L183:
	;
	v819 = v727
	v821 = v726
	goto L182
L184:
	;
	goto L185
L185:
	;
	v731 = v727
	v733 = v726
	v736 = v728
	goto L186
L186:
	;
	v740 = int64(*(*int8)(unsafe.Add(mBase, uint32(v731)+1)))
	if v740 == int64(0) {
		goto L190
	} else {
		goto L191
	}
L187:
	;
	v819 = v817
	v821 = v816
	goto L182
L188:
	;
	v810 = (v805 ^ int64(base.Ui64(v805)>>(uint(int64(23))%64))) * int64(2388976653695081527)
	v816 = (v733 ^ int64(base.Ui64(v810)>>(uint(int64(47))%64)) ^ v810) * int64(-8645972361240307355)
	v817 = v731 + v804
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817))))
	if v818 != 0 {
		v731 = v817
		v733 = v816
		v736 = v818
		goto L186
	} else {
		goto L215
	}
L189:
	;
	v804 = v798
	v805 = base.I64_extend8_s(base.I64_extend_i32_u(v736)) | v799
	goto L188
L190:
	;
	v798 = int32(1)
	v799 = int64(0)
	goto L189
L191:
	;
	goto L192
L192:
	;
	v747 = int64(*(*int8)(unsafe.Add(mBase, uint32(v731)+2)))
	if v747 == int64(0) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v798 = v794
	v799 = v740<<(uint(int64(8))%64) | v795
	goto L189
L194:
	;
	v794 = int32(2)
	v795 = int64(0)
	goto L193
L195:
	;
	goto L196
L196:
	;
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+3)))
	if v752 != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v753 = int64(*(*int8)(unsafe.Add(mBase, uint32(v731)+4)))
	if v753 == int64(0) {
		goto L201
	} else {
		goto L202
	}
L198:
	;
	goto L199
L199:
	;
	v794 = int32(3)
	v795 = v747 << (uint(int64(16)) % 64)
	goto L193
L200:
	;
	v789 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v731))))
	v804 = v787
	v805 = v788 | v789
	goto L188
L201:
	;
	v787 = int32(4)
	v788 = int64(0)
	goto L200
L202:
	;
	goto L203
L203:
	;
	v758 = int64(*(*int8)(unsafe.Add(mBase, uint32(v731)+5)))
	if v758 == int64(0) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v787 = v780
	v788 = v781 | v753<<(uint(int64(32))%64)
	goto L200
L205:
	;
	v780 = int32(5)
	v781 = int64(0)
	goto L204
L206:
	;
	goto L207
L207:
	;
	v763 = int64(*(*int8)(unsafe.Add(mBase, uint32(v731)+6)))
	if v763 == int64(0) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v780 = v775
	v781 = v774 | v758<<(uint(int64(40))%64)
	goto L204
L209:
	;
	v774 = int64(0)
	v775 = int32(6)
	goto L208
L210:
	;
	goto L211
L211:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+7)))
	if v768 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v770 = *(*int64)(unsafe.Add(mBase, uint32(v731)))
	v804 = int32(8)
	v805 = v770
	goto L188
L213:
	;
	goto L214
L214:
	;
	v774 = v763 << (uint(int64(48)) % 64)
	v775 = int32(7)
	goto L208
L215:
	;
	goto L187
L216:
	;
	v858 = v850 & v841
	v863 = v483 + v858*int32(24)
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863)+22)))
	if v864 != 0 {
		v850 = v858 + int32(1)
		goto L216
	} else {
		goto L218
	}
L217:
	;
	v865 = *(*int64)(unsafe.Add(mBase, uint32(v702)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v863)+16)) = v865
	v867 = *(*int64)(unsafe.Add(mBase, uint32(v702)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v863)+8)) = v867
	v869 = *(*int64)(unsafe.Add(mBase, uint32(v702)))
	*(*int64)(unsafe.Add(mBase, uint32(v863))) = v869
	goto L180
L218:
	;
	goto L217
L219:
	;
	v892 = v888
	goto L221
L220:
	;
	v892 = int32(0)
	goto L221
L221:
	;
	v894 = v689 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v894)) < base.Ui64(v477) {
		v685 = v892
		v689 = v894
		goto L176
	} else {
		goto L222
	}
L222:
	;
	goto L177
L223:
	;
	m.G0 = v453 + int32(16)
	goto L114
L224:
	;
	v930 = int32(0)
	v931 = v924
	v933 = v921
	goto L225
L225:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v931)+4))
	if v423 == v945 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v1256 = v1239
	goto L108
L227:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v931)))
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947))))
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424))))
	if base.B2i32(v950 == int32(0))|base.B2i32(v950 != v953) != 0 {
		v971 = v950
		v972 = v953
		goto L231
	} else {
		goto L232
	}
L228:
	;
	goto L229
L229:
	;
	v976 = *(*int64)(unsafe.Add(mBase, uint32(v931)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v976
	v979 = v19 + int32(8)
	v986 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v979)+4)))
	v991 = (int64(base.Ui64(v986)>>(uint(int64(23))%64)) ^ v986) * int64(2388976653695081527)
	v995 = int64(-8645972361240307355)
	v998 = (v991 ^ int64(base.Ui64(v991)>>(uint(int64(47))%64)) ^ v995) * v995
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v979)))
	v1000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v999))))
	if v1000 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L230:
	;
	if v971-v972 == int32(0) {
		v1280 = v931
		goto L107
	} else {
		goto L237
	}
L231:
	;
	goto L230
L232:
	;
	v956 = v947
	v957 = v424
	goto L233
L233:
	;
	v960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v957)+1)))
	v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956)+1)))
	if v961 == int32(0) {
		v971 = v961
		v972 = v960
		goto L231
	} else {
		goto L235
	}
L234:
	;
	v971 = v961
	v972 = v960
	goto L231
L235:
	;
	v964 = int32(1)
	if v961 == v960 {
		v956 = v956 + v964
		v957 = v957 + v964
		goto L233
	} else {
		goto L236
	}
L236:
	;
	goto L234
L237:
	;
	goto L229
L238:
	;
	v1113 = base.I32_wrap_i64(int64(base.Ui64(v1105)>>(uint(int64(47))%64))^v1105-int64(base.Ui64(v1105)>>(uint(int64(32))%64))) & v920
	if base.Ui32(v933) < base.Ui32(v1113) {
		goto L273
	} else {
		goto L274
	}
L239:
	;
	v1105 = (base.I64_extend_i32_s(v1091-v999) + int64(base.Ui64(v1093)>>(uint(int64(23))%64)) ^ v1093) * int64(2388976653695081527)
	goto L238
L240:
	;
	v1091 = v999
	v1093 = v998
	goto L239
L241:
	;
	goto L242
L242:
	;
	v1003 = v999
	v1005 = v998
	v1008 = v1000
	goto L243
L243:
	;
	v1012 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1003)+1)))
	if v1012 == int64(0) {
		goto L247
	} else {
		goto L248
	}
L244:
	;
	v1091 = v1089
	v1093 = v1088
	goto L239
L245:
	;
	v1082 = (v1077 ^ int64(base.Ui64(v1077)>>(uint(int64(23))%64))) * int64(2388976653695081527)
	v1088 = (v1005 ^ int64(base.Ui64(v1082)>>(uint(int64(47))%64)) ^ v1082) * int64(-8645972361240307355)
	v1089 = v1003 + v1076
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1089))))
	if v1090 != 0 {
		v1003 = v1089
		v1005 = v1088
		v1008 = v1090
		goto L243
	} else {
		goto L272
	}
L246:
	;
	v1076 = v1070
	v1077 = base.I64_extend8_s(base.I64_extend_i32_u(v1008)) | v1071
	goto L245
L247:
	;
	v1070 = int32(1)
	v1071 = int64(0)
	goto L246
L248:
	;
	goto L249
L249:
	;
	v1019 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1003)+2)))
	if v1019 == int64(0) {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v1070 = v1066
	v1071 = v1012<<(uint(int64(8))%64) | v1067
	goto L246
L251:
	;
	v1066 = int32(2)
	v1067 = int64(0)
	goto L250
L252:
	;
	goto L253
L253:
	;
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1003)+3)))
	if v1024 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1025 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1003)+4)))
	if v1025 == int64(0) {
		goto L258
	} else {
		goto L259
	}
L255:
	;
	goto L256
L256:
	;
	v1066 = int32(3)
	v1067 = v1019 << (uint(int64(16)) % 64)
	goto L250
L257:
	;
	v1061 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1003))))
	v1076 = v1059
	v1077 = v1060 | v1061
	goto L245
L258:
	;
	v1059 = int32(4)
	v1060 = int64(0)
	goto L257
L259:
	;
	goto L260
L260:
	;
	v1030 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1003)+5)))
	if v1030 == int64(0) {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v1059 = v1052
	v1060 = v1053 | v1025<<(uint(int64(32))%64)
	goto L257
L262:
	;
	v1052 = int32(5)
	v1053 = int64(0)
	goto L261
L263:
	;
	goto L264
L264:
	;
	v1035 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1003)+6)))
	if v1035 == int64(0) {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	v1052 = v1047
	v1053 = v1046 | v1030<<(uint(int64(40))%64)
	goto L261
L266:
	;
	v1046 = int64(0)
	v1047 = int32(6)
	goto L265
L267:
	;
	goto L268
L268:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1003)+7)))
	if v1040 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1042 = *(*int64)(unsafe.Add(mBase, uint32(v1003)))
	v1076 = int32(8)
	v1077 = v1042
	goto L245
L270:
	;
	goto L271
L271:
	;
	v1046 = v1035 << (uint(int64(48)) % 64)
	v1047 = int32(7)
	goto L265
L272:
	;
	goto L244
L273:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v1117 = v933 + v1115
	goto L275
L274:
	;
	v1117 = v933
	goto L275
L275:
	;
	v1119 = v933 + int32(1)
	if base.Ui32(v1117-v1113) < base.Ui32(v930) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1122 = v1119 & v920
	v1125 = v919 + v1122*int32(24)
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125)+22)))
	if v1126 != 0 {
		goto L279
	} else {
		goto L280
	}
L277:
	;
	goto L278
L278:
	;
	v1226 = v930 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v1226) {
		goto L295
	} else {
		goto L296
	}
L279:
	;
	v1129 = v1122
	v1133 = int32(0)
	goto L282
L280:
	;
	v1162 = v1125
	v1163 = v1122
	goto L281
L281:
	;
	if v1163 != v933 {
		goto L289
	} else {
		goto L290
	}
L282:
	;
	v1145 = v1133 + int32(1)
	if int32(151) <= v1145 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	v1162 = v1160
	v1163 = v1157
	goto L281
L284:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v1150 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1148), base.F64_convert_i64_u(v1150)), float64(0.1)) != 0 {
		goto L106
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1157 = (v1129 + int32(1)) & v920
	v1160 = v919 + v1157*int32(24)
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1160)+22)))
	if v1161 != 0 {
		v1129 = v1157
		v1133 = v1145
		goto L282
	} else {
		goto L288
	}
L287:
	;
	goto L286
L288:
	;
	goto L283
L289:
	;
	v1179 = v1162
	v1180 = v1163
	goto L292
L290:
	;
	goto L291
L291:
	;
	v1256 = v931
	goto L108
L292:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v1198 = v1195 & (v1180 - int32(1))
	v1201 = v919 + v1198*int32(24)
	v1202 = *(*int64)(unsafe.Add(mBase, uint32(v1201)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1179)+16)) = v1202
	v1204 = *(*int64)(unsafe.Add(mBase, uint32(v1201)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1179)+8)) = v1204
	v1206 = *(*int64)(unsafe.Add(mBase, uint32(v1201)))
	*(*int64)(unsafe.Add(mBase, uint32(v1179))) = v1206
	if v1198 != v933 {
		v1179 = v1201
		v1180 = v1198
		goto L292
	} else {
		goto L294
	}
L293:
	;
	goto L291
L294:
	;
	goto L293
L295:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v1231 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1229), base.F64_convert_i64_u(v1231)), float64(0.1)) != 0 {
		goto L106
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	v1236 = v1119 & v920
	v1239 = v919 + v1236*int32(24)
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+22)))
	if v1240 != 0 {
		v930 = v1226
		v931 = v1239
		v933 = v1236
		goto L225
	} else {
		goto L299
	}
L298:
	;
	goto L297
L299:
	;
	goto L226
L300:
	;
	F_errmsg_internal(m, int32(_a_F_spcache_insert_0), int32(0))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L67
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(_a_F_spcache_insert_1), int32(635), int32(_a_F_spcache_insert_2))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L67
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L303:
	;
	F_errmsg_internal(m, int32(_a_F_spcache_insert_3), int32(0))
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L67
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(_a_F_spcache_insert_1), int32(332), int32(_a_F_spcache_insert_4))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L67
	} else {
		goto L305
	}
L305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_specialcolors(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v5 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v9 = F_newcolor(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
			if v12 != 0 {
				v27 = int32(_a_F_specialcolors_0)
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
				v17 = v14 + v9*int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(2)
				*(*int64)(unsafe.Add(mBase, uint32(v17)+12)) = int64(0)
				v22 = int32(_a_F_specialcolors_0)
				*(*uint16)(unsafe.Add(mBase, uint32(v17)+8)) = uint16(v22)
				*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(4294967296)
				v27 = v9
			}
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)) = uint16(v27)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v30 = F_newcolor(m, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
				if v33 != 0 {
					v48 = int32(_a_F_specialcolors_0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
					v38 = v35 + v30*int32(24)
					*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(2)
					*(*int64)(unsafe.Add(mBase, uint32(v38)+12)) = int64(0)
					v43 = int32(_a_F_specialcolors_0)
					*(*uint16)(unsafe.Add(mBase, uint32(v38)+8)) = uint16(v43)
					*(*int64)(unsafe.Add(mBase, uint32(v38))) = int64(4294967296)
					v48 = v30
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)) = uint16(v48)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v51 = F_newcolor(m, v50)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
					if v55 != 0 {
						v70 = int32(_a_F_specialcolors_0)
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
						v60 = v57 + v51*int32(24)
						*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = int32(2)
						*(*int64)(unsafe.Add(mBase, uint32(v60)+12)) = int64(0)
						v65 = int32(_a_F_specialcolors_0)
						*(*uint16)(unsafe.Add(mBase, uint32(v60)+8)) = uint16(v65)
						*(*int64)(unsafe.Add(mBase, uint32(v60))) = int64(4294967296)
						v70 = v51
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)) = uint16(v70)
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v73 = F_newcolor(m, v72)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
						if v76 != 0 {
							v99 = int32(_a_F_specialcolors_0)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)) = uint16(v99)
							return
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
							v80 = v77 + v73*int32(24)
							*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = int32(2)
							*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = int64(0)
							v85 = int32(_a_F_specialcolors_0)
							*(*uint16)(unsafe.Add(mBase, uint32(v80)+8)) = uint16(v85)
							*(*int64)(unsafe.Add(mBase, uint32(v80))) = int64(4294967296)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)) = uint16(v73)
							return
						}
					}
				}
			}
		}
	} else {
		v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+56)))
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)) = uint16(v90)
		v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+58)))
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)) = uint16(v92)
		v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+60)))
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)) = uint16(v94)
		v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+62)))
		v99 = v96
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)) = uint16(v99)
		return
	}
}
func F_spgvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v9 != 0 {
		v33 = l1
		m.G0 = v7 + int32(112)
		return v33
	} else {
		if l1 == int32(0) {
			v13 = F_palloc0(m, int32(40))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(285)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v13
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_spgvacuumscan(m, v7)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = v13
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
					if v26 != 0 {
						v33 = v25
					} else {
						v27 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
						v28 = *(*float64)(unsafe.Add(mBase, uint32(v25)+8))
						if base.F64_lt(v27, v28) == int32(0) {
							v33 = v25
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v25)+8)) = v27
							v33 = v25
						}
					}
					m.G0 = v7 + int32(112)
					return v33
				}
			}
		} else {
			v25 = l1
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
			if v26 != 0 {
				v33 = v25
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v25)+8))
				if base.F64_lt(v27, v28) == int32(0) {
					v33 = v25
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v25)+8)) = v27
					v33 = v25
				}
			}
			m.G0 = v7 + int32(112)
			return v33
		}
	}
}
func F_ssup_datum_int32_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v5 = base.I32_wrap_i64(l0)
	v6 = base.I32_wrap_i64(l1)
	return base.B2i32(v6 < v5) - base.B2i32(v5 < v6)
}
func F_start_apply(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
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
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int64
	_ = v171
	var v173 int64
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v230 int64
	_ = v230
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int64
	_ = v251
	var v252 int32
	_ = v252
	var v254 int64
	_ = v254
	var v255 int32
	_ = v255
	var v257 int64
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int64
	_ = v263
	var v265 int64
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int64
	_ = v301
	var v302 int32
	_ = v302
	var v304 int64
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int64
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int64
	_ = v336
	var v337 int64
	_ = v337
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v354 int64
	_ = v354
	var v355 int32
	_ = v355
	var v358 int64
	_ = v358
	var v359 int32
	_ = v359
	var v362 int64
	_ = v362
	var v363 int32
	_ = v363
	var v366 int64
	_ = v366
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int64
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int64
	_ = v407
	var v408 int64
	_ = v408
	var v418 int32
	_ = v418
	var v423 int64
	_ = v423
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v454 int64
	_ = v454
	var v456 int64
	_ = v456
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int64
	_ = v581
	var v582 int64
	_ = v582
	var v590 int64
	_ = v590
	var v592 int32
	_ = v592
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v679 int64
	_ = v679
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v721 int32
	_ = v721
	var v722 int64
	_ = v722
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(320)
	m.G0 = v17
	v22 = v2
	v23 = int32(-1)
	v24 = v2
	v27 = v2
	v28 = v2
	goto L1
L1:
	;
	if v23 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[0]))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[1]))
	v42 = v17 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v17 + int32(12)
	goto L6
L4:
	;
	v48 = v24
	v49 = v27
	v50 = v28
	goto L5
L5:
	;
	goto L8
L6:
	;
	v48 = int32(0)
	v49 = v38
	v50 = v40
	goto L5
L7:
	;
	goto L2
L8:
	;
	if v48 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v721 = int32(m.ExcTag)
	v722 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v721 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v22
	v55 = int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[1])) = v17 + v55
	v61 = m.G0
	v63 = v61 - v55
	m.G0 = v63
	F_gettimeofday(m, v63)
	mBase = m.M
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
	v67 = int64(*(*int32)(unsafe.Add(mBase, uint32(v63)+8)))
	m.G0 = v63 + v55
	goto L14
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[0])) = v49
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[1])) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v22
	v679 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_start_apply[2])) = v679
	*(*int64)(unsafe.Add(mBase, _c_F_start_apply[3])) = v679
	v685 = int32(0)
	*(*uint16)(unsafe.Add(mBase, _c_F_start_apply[4])) = uint16(v685)
	goto L149
L14:
	;
	v78 = int32(0)
	base.MemoryFill(m, v17+int32(200), v78, int32(96))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v22
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[5]))
	v89 = F_AllocSetContextCreateInternal(m, v84, int32(_a_F_start_apply_0), v78, int32(_a_F_start_apply_1), int32(_a_F_start_apply_2))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[6])) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v22
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[5]))
	v100 = F_AllocSetContextCreateInternal(m, v95, int32(_a_F_start_apply_3), int32(0), int32(_a_F_start_apply_1), int32(_a_F_start_apply_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[7])) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v22
	v105 = int32(0)
	F_pgstat_report_activity(m, int32(2), v105)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v17)+304)) = int32(1052)
	v112 = v17 + int32(300)
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[8])) = v112
	v114 = int32(_a_F_start_apply_4)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[0])) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v17)+300)) = v115
	v121 = v22
	v122 = v105
	v129 = l0
	v131 = v67 + v66*int64(1000000) - int64(946684800000000)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[9]))
	if v138 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v17)+300))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[0])) = v652
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[8])) = v652
	v657 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[10]))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[11]))
	m.T0[v658].(func(*base.Module, int32, int32))(m, v661, v17+int32(312))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L10
	} else {
		goto L148
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v121
	F_ProcessInterrupts(m)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[12])) = v144
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[10]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v121
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[11]))
	v156 = m.T0[v148].(func(*base.Module, int32, int32, int32) int32)(m, v151, v17+int32(192), v17+int32(196))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L10
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v158 = int32(0)
	if v156 == v158 {
		v446 = v121
		v447 = v122
		v449 = v158
		v454 = v129
		v456 = v131
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	v459 = int32(0)
	F_send_feedback(m, v454, v459, v459)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L10
	} else {
		goto L85
	}
L25:
	;
	v163 = v121
	v164 = v122
	v165 = v156
	v171 = v129
	v173 = v131
	goto L26
L26:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[9]))
	if v176 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	F_ProcessInterrupts(m)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L10
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v165 == int32(0) {
		v446 = v163
		v447 = v164
		v449 = v158
		v454 = v171
		v456 = v173
		goto L24
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	if v165 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v187 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L10
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[13]))
	if v204 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v189 = int32(1)
	if v187 == int32(0) {
		v446 = v163
		v447 = v164
		v449 = v189
		v454 = v171
		v456 = v173
		goto L24
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	F_errmsg(m, int32(_a_F_start_apply_5), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	F_errfinish(m, int32(_a_F_start_apply_6), int32(4080), int32(_a_F_start_apply_7))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v446 = v163
	v447 = v164
	v449 = v189
	v454 = v171
	v456 = v173
	goto L24
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[13])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L10
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v216 = m.G0
	v217 = int32(16)
	v218 = v216 - v217
	m.G0 = v218
	F_gettimeofday(m, v218)
	mBase = m.M
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v218)))
	v222 = int64(*(*int32)(unsafe.Add(mBase, uint32(v218)+8)))
	m.G0 = v218 + v217
	v230 = v222 + v221*int64(1000000) - int64(946684800000000)
	goto L44
L43:
	;
	goto L42
L44:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[12])) = v233
	*(*int64)(unsafe.Add(mBase, uint32(v17)+272)) = v230
	*(*int64)(unsafe.Add(mBase, uint32(v17)+184)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+180)) = v165
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v17)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+176)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v244 = F_pq_getmsgbyte(m, v17+int32(176))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L10
	} else {
		goto L49
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[6]))
	F_MemoryContextReset(m, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L10
	} else {
		goto L83
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v353 = v17 + int32(176)
	v354 = F_pq_getmsgint64(m, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L10
	} else {
		goto L71
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v300 = v17 + int32(176)
	v301 = F_pq_getmsgint64(m, v300)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L10
	} else {
		goto L62
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v250 = v17 + int32(176)
	v251 = F_pq_getmsgint64(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L10
	} else {
		goto L50
	}
L49:
	;
	switch v244 - int32(107) {
	case 0:
		goto L47
	default:
		v423 = v171
		goto L45
	case 8:
		goto L46
	case 12:
		goto L48
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v254 = F_pq_getmsgint64(m, v250)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v257 = F_pq_getmsgint64(m, v250)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v260)+88)) = v257
	if base.Ui64(v251) < base.Ui64(v171) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v263 = v171
	goto L55
L54:
	;
	v263 = v251
	goto L55
L55:
	;
	if base.Ui64(v254) < base.Ui64(v263) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v265 = v263
	goto L58
L57:
	;
	v265 = v254
	goto L58
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v260)+80)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v271 = m.G0
	v272 = int32(16)
	v273 = v271 - v272
	m.G0 = v273
	F_gettimeofday(m, v273)
	mBase = m.M
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v273)))
	v277 = int64(*(*int32)(unsafe.Add(mBase, uint32(v273)+8)))
	m.G0 = v273 + v272
	goto L59
L59:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v287)+96)) = v277 + v276*int64(1000000) - int64(946684800000000)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	F_apply_dispatch(m, v250)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	F_maybe_advance_nonremovable_xid(m, v17+int32(200), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	v423 = v265
	goto L45
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v304 = F_pq_getmsgint64(m, v300)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v307 = F_pq_getmsgbyte(m, v300)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	if base.Ui64(v301) < base.Ui64(v171) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v311 = v171
	goto L67
L66:
	;
	v311 = v301
	goto L67
L67:
	;
	v312 = int32(0)
	F_send_feedback(m, v311, base.B2i32(v307 != v312), v312)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	F_maybe_advance_nonremovable_xid(m, v17+int32(200), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v324)+88)) = v304
	*(*int64)(unsafe.Add(mBase, uint32(v324)+80)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v331 = m.G0
	v332 = int32(16)
	v333 = v331 - v332
	m.G0 = v333
	F_gettimeofday(m, v333)
	mBase = m.M
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v333)))
	v337 = int64(*(*int32)(unsafe.Add(mBase, uint32(v333)+8)))
	m.G0 = v333 + v332
	goto L70
L70:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v347)+112)) = v304
	*(*int64)(unsafe.Add(mBase, uint32(v347)+104)) = v311
	*(*int64)(unsafe.Add(mBase, uint32(v347)+96)) = v337 + v336*int64(1000000) - int64(946684800000000)
	v423 = v311
	goto L45
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	*(*int64)(unsafe.Add(mBase, uint32(v17)+208)) = v354
	v358 = F_pq_getmsgint64(m, v353)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	*(*int64)(unsafe.Add(mBase, uint32(v17)+216)) = v358
	v362 = F_pq_getmsgint64(m, v353)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	*(*int64)(unsafe.Add(mBase, uint32(v17)+224)) = v362
	v366 = F_pq_getmsgint64(m, v353)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+232)) = v366
	if v354 == int64(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L10
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	F_maybe_advance_nonremovable_xid(m, v17+int32(200), int32(1))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L10
	} else {
		goto L81
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	F_errmsg_internal(m, int32(_a_F_start_apply_8), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	F_errfinish(m, int32(_a_F_start_apply_6), int32(_a_F_start_apply_9), int32(_a_F_start_apply_7))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	goto L7
L81:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[14]))
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v17)+232))
	*(*int64)(unsafe.Add(mBase, uint32(v394)+88)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v394)+80)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v402 = m.G0
	v403 = int32(16)
	v404 = v402 - v403
	m.G0 = v404
	F_gettimeofday(m, v404)
	mBase = m.M
	v407 = *(*int64)(unsafe.Add(mBase, uint32(v404)))
	v408 = int64(*(*int32)(unsafe.Add(mBase, uint32(v404)+8)))
	m.G0 = v404 + v403
	goto L82
L82:
	;
	v418 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v418)+96)) = v408 + v407*int64(1000000) - int64(946684800000000)
	v423 = v171
	goto L45
L83:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[10]))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v163
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[11]))
	v442 = m.T0[v434].(func(*base.Module, int32, int32, int32) int32)(m, v437, v17+int32(192), v17+int32(196))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v163 = v442
	v164 = int32(0)
	v165 = v442
	v171 = v423
	v173 = v230
	goto L26
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	*(*int64)(unsafe.Add(mBase, uint32(v17)+272)) = int64(0)
	F_maybe_advance_nonremovable_xid(m, v17+int32(200), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_apply[15])))
	if v472 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	v488 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[6]))
	F_MemoryContextReset(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L10
	} else {
		goto L93
	}
L88:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_apply[16])))
	if v474&int32(1) != 0 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	F_maybe_reread_subscription(m)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	F_ProcessSyncingRelations(m, v454)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	goto L87
L93:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[17]))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[12])) = v493
	if v449 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[18]))
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[19]))
	if v501 == int32(_a_F_start_apply_10) {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	goto L96
L96:
	;
	goto L18
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	v529 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[20]))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v17)+196))
	v532 = F_WaitLatchOrSocket(m, v529, v530, v526, int32(83886087))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L10
	} else {
		goto L118
	}
L98:
	;
	v504 = int32(1000)
	goto L100
L99:
	;
	v504 = v499
	goto L100
L100:
	;
	if v501 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v506 = v504
	goto L103
L102:
	;
	v506 = int32(1000)
	goto L103
L103:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[21]))
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+48)))
	if v509 != int32(1) {
		v526 = v506
		goto L97
	} else {
		goto L104
	}
L104:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v17)+200))
	if v512 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v508)+44))
	if v506 < v519 {
		goto L111
	} else {
		goto L112
	}
L106:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v17)+288))
	if v513 == int32(0) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	if v506 < v513 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v517 = v506
	goto L110
L109:
	;
	v517 = v513
	goto L110
L110:
	;
	v526 = v517
	goto L97
L111:
	;
	v521 = v506
	goto L113
L112:
	;
	v521 = v519
	goto L113
L113:
	;
	if int32(0) < v519 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v524 = v521
	goto L116
L115:
	;
	v524 = v506
	goto L116
L116:
	;
	v526 = v524
	goto L97
L117:
	;
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[13]))
	if v555 != 0 {
		goto L123
	} else {
		goto L124
	}
L118:
	;
	if v532&int32(1) == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[20]))
	v541 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v540))) = v541
	v546 = base.AtomicRmwOr32(m, v541, int32(_a_F_start_apply_11), v541)
	goto L120
L120:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[9]))
	if v548 == int32(0) {
		goto L117
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	F_ProcessInterrupts(m)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L10
	} else {
		goto L122
	}
L122:
	;
	goto L117
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[13])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L10
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if v532&int32(8) == int32(0) {
		v121 = v446
		v122 = v447
		v129 = v454
		v131 = v456
		goto L17
	} else {
		goto L127
	}
L126:
	;
	goto L125
L127:
	;
	v567 = int32(0)
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[22]))
	if v569 <= v567 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	F_send_feedback(m, v454, v628, v628)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L10
	} else {
		goto L141
	}
L129:
	;
	v628 = v567
	v629 = v447
	goto L128
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	v576 = m.G0
	v577 = int32(16)
	v578 = v576 - v577
	m.G0 = v578
	F_gettimeofday(m, v578)
	mBase = m.M
	v581 = *(*int64)(unsafe.Add(mBase, uint32(v578)))
	v582 = int64(*(*int32)(unsafe.Add(mBase, uint32(v578)+8)))
	m.G0 = v578 + v577
	v590 = v582 + v581*int64(1000000) - int64(946684800000000)
	goto L132
L132:
	;
	v592 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[22]))
	if base.I64_extend_i32_s(v592)*int64(1000)+v456 <= v590 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L10
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v618 = int32(1)
	if v447&v618 != 0 {
		v628 = v567
		v629 = v618
		goto L128
	} else {
		goto L140
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	F_errcode(m, int32(100663808))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L10
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	F_errmsg(m, int32(_a_F_start_apply_12), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L10
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	F_errfinish(m, int32(_a_F_start_apply_6), int32(_a_F_start_apply_13), int32(_a_F_start_apply_7))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L10
	} else {
		goto L139
	}
L139:
	;
	goto L7
L140:
	;
	v622 = base.I32_div_s(v592, int32(2))
	v627 = base.B2i32(base.I64_extend_i32_s(v622)*int64(1000)+v456 <= v590)
	v628 = v627
	v629 = v627
	goto L128
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	F_maybe_advance_nonremovable_xid(m, v17+int32(200), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L10
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	v643 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[23]))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v643)+20))
	goto L143
L143:
	;
	if v644 == int32(2) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v121 = v446
	v122 = v629
	v129 = v454
	v131 = v456
	goto L17
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v446
	v649 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L10
	} else {
		goto L147
	}
L147:
	;
	v121 = v446
	v122 = v629
	v129 = v454
	v131 = v456
	goto L17
L148:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[1])) = v50
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[0])) = v49
	m.G0 = v17 + int32(320)
	return
L149:
	;
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[21]))
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688)+37)))
	if v689 == int32(1) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v22
	F_DisableSubscriptionAndExit(m)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L10
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v22
	F_AbortOutOfAnyTransaction(m)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L10
	} else {
		goto L154
	}
L153:
	;
	goto L7
L154:
	;
	v699 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[21]))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v699)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v22
	F_pgstat_report_subscription_error(m, v700)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L10
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+316)) = v22
	F_pg_re_throw(m)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L10
	} else {
		goto L156
	}
L156:
	;
	goto L9
L157:
	;
	v726 = int32(v722)
	m.G0 = v17
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v726)+4))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	if v17+int32(12) == v732 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	m.ExcPending = 1
	goto L166
L159:
	;
	if v736 != 0 {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v729)+4))
	v736 = v734
	goto L162
L161:
	;
	v736 = int32(0)
	goto L162
L162:
	;
	goto L159
L163:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v17)+316))
	v22 = v737
	v23 = v736
	v24 = v728
	v27 = v49
	v28 = v50
	goto L1
L164:
	;
	goto L165
L165:
	;
	F___wasm_longjmp(m, v729, v728)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	return
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_storeProcedures(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v26 int64
	_ = v26
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v19 = F_table_open(m, int32(2603), int32(3))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l2 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L51
	}
L4:
	;
	F_relation_close(m, v19, int32(3))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L50
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v23 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v26 = base.I64_extend_i32_u(l1)
	v35 = int32(0)
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v35<<(uint(int32(2))%32))))
	if l3 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	v45 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+12)))
	v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+16)))
	v47 = int64(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v48 = F_SearchSysCacheExists(m, int32(5), v26, v45, v46, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v50 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v50
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v62
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+44)) = uint16(v62)
	v68 = F_GetNewOidWithIndex(m, v19, int32(2757), int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	if v48 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = base.I64_extend_i32_u(v68)
	v73 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v73
	v75 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+16)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v75
	v77 = int64(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v77
	v79 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v86 = F_heap_form_tuple(m, v81, v15+int32(48), v15+int32(40))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_CatalogTupleInsert(m, v19, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_pfree(m, v86)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v92 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = int32(2603)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(1255)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v99
	v104 = v15 + int32(28)
	v106 = v15 + int32(16)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)))
	if v109 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v110 = int32(110)
	goto L20
L19:
	;
	v110 = int32(97)
	goto L20
L20:
	;
	F_recordDependencyOn(m, v104, v106, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+25)))
	if v115 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v116 = int32(2753)
	goto L24
L23:
	;
	v116 = int32(2616)
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v118
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)))
	if v124 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v125 = int32(105)
	goto L27
L26:
	;
	v125 = int32(97)
	goto L27
L27:
	;
	F_recordDependencyOn(m, v104, v106, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v129 = F_typeDepNeeded(m, v128, v43)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v129 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(1247)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v133
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)))
	if v139 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	if v144 == v145 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v140 = int32(110)
	goto L35
L34:
	;
	v140 = int32(97)
	goto L35
L35:
	;
	F_recordDependencyOn(m, v104, v106, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_storeProcedures[0]))
	if v169 != 0 {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v147 = F_typeDepNeeded(m, v144, v43)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v147 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(1247)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v153
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)))
	if v163 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v164 = int32(110)
	goto L43
L42:
	;
	v164 = int32(97)
	goto L43
L43:
	;
	F_recordDependencyOn(m, v15+int32(28), v15+int32(16), v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L37
L45:
	;
	v171 = int32(0)
	F_RunObjectPostCreateHook(m, int32(2603), v68, v171, v171)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v176 = v35 + int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v176 < v177 {
		v35 = v176
		goto L7
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	goto L8
L50:
	;
	m.G0 = v15 + int32(96)
	return
L51:
	;
	F_errcode(m, int32(_a_F_storeProcedures_0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v206 = F_format_type_be(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v209 = F_format_type_be(m, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v211 = F_NameListToString(m, l0)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v204
	F_errmsg(m, int32(_a_F_storeProcedures_1), v15)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_storeProcedures_2), int32(1636), int32(_a_F_storeProcedures_3))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_str_udeescape(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v119 int32
	_ = v119
	var v133 int32
	_ = v133
	var v147 int32
	_ = v147
	var v161 int32
	_ = v161
	var v183 int32
	_ = v183
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v227 int32
	_ = v227
	var v248 int32
	_ = v248
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v294 int32
	_ = v294
	var v308 int32
	_ = v308
	var v322 int32
	_ = v322
	var v336 int32
	_ = v336
	var v350 int32
	_ = v350
	var v364 int32
	_ = v364
	var v378 int32
	_ = v378
	var v400 int32
	_ = v400
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v444 int32
	_ = v444
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v488 int32
	_ = v488
	var v509 int32
	_ = v509
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v552 int32
	_ = v552
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v688 int32
	_ = v688
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	v2 = l1
	v23 = m.G0
	v25 = v23 - int32(32)
	m.G0 = v25
	v27 = F_strlen(m, l0)
	mBase = m.M
	v29 = v27 + int32(17)
	v30 = F_palloc(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v34 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L137
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L134
	}
L5:
	;
	v688 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v672))) = uint8(v688)
	m.G0 = v25 + int32(32)
	return v678
L6:
	;
	v672 = v30
	v678 = v30
	goto L5
L7:
	;
	goto L8
L8:
	;
	v37 = l2 - l0
	v41 = l0
	v43 = int32(0)
	v45 = v34
	v47 = v30
	v53 = v30
	v56 = v29
	goto L11
L9:
	;
	goto L3
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L129
	}
L11:
	;
	v63 = v47 - v53
	if base.Ui32(v56-int32(17)) < base.Ui32(v63) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v618&int32(_a_F_str_udeescape_0) == int32(0) {
		v672 = v607
		v678 = v75
		goto L5
	} else {
		goto L128
	}
L13:
	;
	v68 = v56 << (uint(int32(1)) % 32)
	v69 = F_repalloc(m, v53, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v73 = v45
	v74 = v47
	v75 = v53
	v76 = v56
	goto L15
L15:
	;
	v77 = int32(255)
	v78 = v2 & v77
	if v78 == v73&v77 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v73 = v72
	v74 = v69 + v63
	v75 = v69
	v76 = v68
	goto L15
L17:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(12))+8))
	*(*int32)(unsafe.Add(mBase, _c_F_str_udeescape[0])) = v622
	goto L126
L18:
	;
	F_pg_unicode_to_server(m, v283, v74)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L125
	}
L19:
	;
	v83 = v25 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = int32(533)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v41 + (v37 + int32(3))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v83
	v90 = int32(_a_F_str_udeescape_1)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_str_udeescape[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v91
	*(*int32)(unsafe.Add(mBase, _c_F_str_udeescape[0])) = v25 + int32(20)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v78 == v97 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	if v43&int32(_a_F_str_udeescape_0) != 0 {
		v630 = v41
		goto L10
	} else {
		goto L123
	}
L22:
	;
	if v43&int32(_a_F_str_udeescape_0) != 0 {
		v630 = v41
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L27
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v2)
	v606 = int32(2)
	v607 = v74 + int32(1)
	v618 = int32(0)
	goto L17
L26:
	;
	if v97 != int32(43) {
		goto L63
	} else {
		goto L64
	}
L27:
	;
	if base.B2i32(base.Ui32(v97-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v97|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)))
	goto L29
L29:
	;
	if base.B2i32(base.Ui32(v119-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v119|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+3)))
	goto L31
L31:
	;
	if base.B2i32(base.Ui32(v133-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v133|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	goto L33
L33:
	;
	if base.B2i32(base.Ui32(v147-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v147|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v161 = int32(-48)
	if base.Ui32((v97-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v183 = v161
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if base.Ui32((v119-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v204 = v161
		goto L39
	} else {
		goto L40
	}
L36:
	;
	if base.Ui32((v97-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v183 = int32(-87)
		goto L35
	} else {
		goto L37
	}
L37:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v97-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v183 = int32(-55)
	goto L35
L39:
	;
	v205 = int32(-48)
	if base.Ui32((v133-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v227 = v205
		goto L45
	} else {
		goto L46
	}
L40:
	;
	if base.Ui32((v119-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v204 = int32(-87)
	goto L39
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v119-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v204 = int32(-55)
	goto L39
L45:
	;
	if base.Ui32((v147-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v248 = v205
		goto L49
	} else {
		goto L50
	}
L46:
	;
	if base.Ui32((v133-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v227 = int32(-87)
		goto L45
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v133-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v227 = int32(-55)
	goto L45
L49:
	;
	v261 = v248 + v147 + ((v119+v204)<<(uint(int32(8))%32) + (v97+v183)<<(uint(int32(12))%32) + (v133+v227)<<(uint(int32(4))%32))
	if base.Ui32(int32(_a_F_str_udeescape_2)) <= base.Ui32(v261-int32(1)) {
		goto L3
	} else {
		goto L55
	}
L50:
	;
	if base.Ui32((v147-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v248 = int32(-87)
	goto L49
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v147-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v248 = int32(-55)
	goto L49
L55:
	;
	v267 = v261 & int32(_a_F_str_udeescape_3)
	if v43&int32(_a_F_str_udeescape_0) != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v283&int32(16776192) != int32(_a_F_str_udeescape_4) {
		goto L18
	} else {
		goto L62
	}
L57:
	;
	if v267 != int32(_a_F_str_udeescape_5) {
		v630 = v41
		goto L10
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v267 == int32(_a_F_str_udeescape_5) {
		v630 = v41
		goto L10
	} else {
		goto L61
	}
L60:
	;
	v272 = int32(1023)
	v283 = v261&v272 | v43&v272<<(uint(int32(10))%32) + int32(_a_F_str_udeescape_6)
	goto L56
L61:
	;
	v283 = v261
	goto L56
L62:
	;
	v606 = int32(5)
	v607 = v74
	v618 = v283
	goto L17
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L118
	}
L64:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)))
	goto L65
L65:
	;
	if base.B2i32(base.Ui32(v294-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v294|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+3)))
	goto L67
L67:
	;
	if base.B2i32(base.Ui32(v308-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v308|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L63
	} else {
		goto L68
	}
L68:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	goto L69
L69:
	;
	if base.B2i32(base.Ui32(v322-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v322|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L63
	} else {
		goto L70
	}
L70:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+5)))
	goto L71
L71:
	;
	if base.B2i32(base.Ui32(v336-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v336|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+6)))
	goto L73
L73:
	;
	if base.B2i32(base.Ui32(v350-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v350|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L63
	} else {
		goto L74
	}
L74:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
	goto L75
L75:
	;
	if base.B2i32(base.Ui32(v364-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v364|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L63
	} else {
		goto L76
	}
L76:
	;
	v378 = int32(-48)
	if base.Ui32((v294-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v400 = v378
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if base.Ui32((v308-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v421 = v378
		goto L81
	} else {
		goto L82
	}
L78:
	;
	if base.Ui32((v294-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v400 = int32(-87)
		goto L77
	} else {
		goto L79
	}
L79:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v294-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	v400 = int32(-55)
	goto L77
L81:
	;
	v422 = int32(-48)
	if base.Ui32((v322-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v444 = v422
		goto L87
	} else {
		goto L88
	}
L82:
	;
	if base.Ui32((v308-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v421 = int32(-87)
	goto L81
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v308-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v421 = int32(-55)
	goto L81
L87:
	;
	if base.Ui32((v336-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v465 = v422
		goto L91
	} else {
		goto L92
	}
L88:
	;
	if base.Ui32((v322-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v444 = int32(-87)
		goto L87
	} else {
		goto L89
	}
L89:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v322-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	v444 = int32(-55)
	goto L87
L91:
	;
	v466 = int32(-48)
	if base.Ui32((v350-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v488 = v466
		goto L97
	} else {
		goto L98
	}
L92:
	;
	if base.Ui32((v336-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v465 = int32(-87)
	goto L91
L94:
	;
	goto L95
L95:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v336-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	v465 = int32(-55)
	goto L91
L97:
	;
	if base.Ui32((v364-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v509 = v466
		goto L101
	} else {
		goto L102
	}
L98:
	;
	if base.Ui32((v350-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v488 = int32(-87)
		goto L97
	} else {
		goto L99
	}
L99:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v350-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	v488 = int32(-55)
	goto L97
L101:
	;
	v530 = v364 + v509 + ((v421+v308)<<(uint(int32(16))%32) + (v294+v400)<<(uint(int32(20))%32) + (v322+v444)<<(uint(int32(12))%32) + (v336+v465)<<(uint(int32(8))%32) + (v350+v488)<<(uint(int32(4))%32))
	if base.Ui32(int32(_a_F_str_udeescape_2)) <= base.Ui32(v530-int32(1)) {
		goto L9
	} else {
		goto L107
	}
L102:
	;
	if base.Ui32((v364-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v509 = int32(-87)
	goto L101
L104:
	;
	goto L105
L105:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v364-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	v509 = int32(-55)
	goto L101
L107:
	;
	v536 = v530 & int32(_a_F_str_udeescape_3)
	if v43&int32(_a_F_str_udeescape_0) != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	if v552&int32(16776192) == int32(_a_F_str_udeescape_4) {
		goto L114
	} else {
		goto L115
	}
L109:
	;
	if v536 != int32(_a_F_str_udeescape_5) {
		v630 = v41
		goto L10
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	if v536 == int32(_a_F_str_udeescape_5) {
		v630 = v41
		goto L10
	} else {
		goto L113
	}
L112:
	;
	v541 = int32(1023)
	v552 = v530&v541 | v43&v541<<(uint(int32(10))%32) + int32(_a_F_str_udeescape_6)
	goto L108
L113:
	;
	v552 = v530
	goto L108
L114:
	;
	v606 = int32(8)
	v607 = v74
	v618 = v552
	goto L17
L115:
	;
	goto L116
L116:
	;
	F_pg_unicode_to_server(m, v552, v74)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v560 = F_strlen(m, v74)
	mBase = m.M
	v606 = int32(8)
	v607 = v560 + v74
	v618 = int32(0)
	goto L17
L118:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(_a_F_str_udeescape_7), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errhint(m, int32(_a_F_str_udeescape_8), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(495), int32(_a_F_str_udeescape_10))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v73)
	v593 = int32(1)
	v594 = v74 + v593
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v595 != 0 {
		v41 = v41 + v593
		v43 = int32(0)
		v45 = v595
		v47 = v594
		v53 = v75
		v56 = v76
		goto L11
	} else {
		goto L124
	}
L124:
	;
	v672 = v594
	v678 = v75
	goto L5
L125:
	;
	v601 = F_strlen(m, v74)
	mBase = m.M
	v606 = int32(5)
	v607 = v601 + v74
	v618 = int32(0)
	goto L17
L126:
	;
	v624 = v41 + v606
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624))))
	if v625 != 0 {
		v41 = v624
		v43 = v618
		v45 = v625
		v47 = v607
		v53 = v75
		v56 = v76
		goto L11
	} else {
		goto L127
	}
L127:
	;
	goto L12
L128:
	;
	v630 = v624
	goto L10
L129:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(_a_F_str_udeescape_11), int32(0))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_scanner_errposition(m, v630+v37+int32(3), l3)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(525), int32(_a_F_str_udeescape_10))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errmsg_internal(m, int32(_a_F_str_udeescape_12), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(336), int32(_a_F_str_udeescape_13))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errmsg(m, int32(_a_F_str_udeescape_14), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(347), int32(_a_F_str_udeescape_15))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_strcat(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v3 = F_strlen(m, l0)
	mBase = m.M
	v4 = v3 + l0
	if (l1^v4)&int32(3) != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return l0
L2:
	;
	goto L1
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v58)
	if v58&int32(255) == int32(0) {
		goto L2
	} else {
		goto L18
	}
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v57 = l1
	v58 = v10
	v59 = v4
	goto L3
L5:
	;
	goto L6
L6:
	;
	if l1&int32(3) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v14 = l1
	v16 = v4
	goto L10
L8:
	;
	v28 = l1
	v30 = v4
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v35 = int32(-2139062144)
	if (int32(16843008)-v32|v32)&v35 != v35 {
		v57 = v28
		v58 = v32
		v59 = v30
		goto L3
	} else {
		goto L14
	}
L10:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v17)
	if v17 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v28 = v24
	v30 = v22
	goto L9
L12:
	;
	v21 = int32(1)
	v22 = v16 + v21
	v24 = v14 + v21
	if v24&int32(3) != 0 {
		v14 = v24
		v16 = v22
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v40 = v28
	v41 = v32
	v42 = v30
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v41
	v44 = int32(4)
	v45 = v42 + v44
	v47 = v40 + v44
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v52 = int32(-2139062144)
	if (int32(16843008)-v49|v49)&v52 == v52 {
		v40 = v47
		v41 = v49
		v42 = v45
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v57 = v47
	v58 = v49
	v59 = v45
	goto L3
L17:
	;
	goto L16
L18:
	;
	v66 = v57
	v68 = v59
	goto L19
L19:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)) = uint8(v69)
	v71 = int32(1)
	if v69 != 0 {
		v66 = v66 + v71
		v68 = v68 + v71
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L2
L21:
	;
	goto L20
}
func F_strcmp(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v5 == int32(0))|base.B2i32(v5 != v8) != 0 {
		v26 = v5
		v27 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v26 - v27
L2:
	;
	v11 = l0
	v12 = l1
	goto L3
L3:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v16 == int32(0) {
		v26 = v16
		v27 = v15
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v26 = v16
	v27 = v15
	goto L1
L5:
	;
	v19 = int32(1)
	if v16 == v15 {
		v11 = v11 + v19
		v12 = v12 + v19
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
}
func F_strcpy(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	if (l1^l0)&int32(3) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return l0
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v56)
	if v56&int32(255) == int32(0) {
		goto L1
	} else {
		goto L17
	}
L3:
	;
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v55 = l1
	v56 = v8
	v57 = l0
	goto L2
L4:
	;
	goto L5
L5:
	;
	if l1&int32(3) != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v12 = l1
	v14 = l0
	goto L9
L7:
	;
	v26 = l1
	v28 = l0
	goto L8
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v33 = int32(-2139062144)
	if (int32(16843008)-v30|v30)&v33 != v33 {
		v55 = v26
		v56 = v30
		v57 = v28
		goto L2
	} else {
		goto L13
	}
L9:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v15)
	if v15 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v26 = v22
	v28 = v20
	goto L8
L11:
	;
	v19 = int32(1)
	v20 = v14 + v19
	v22 = v12 + v19
	if v22&int32(3) != 0 {
		v12 = v22
		v14 = v20
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v38 = v26
	v39 = v30
	v40 = v28
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v39
	v42 = int32(4)
	v43 = v40 + v42
	v45 = v38 + v42
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v50 = int32(-2139062144)
	if (int32(16843008)-v47|v47)&v50 == v50 {
		v38 = v45
		v39 = v47
		v40 = v43
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v55 = v45
	v56 = v47
	v57 = v43
	goto L2
L16:
	;
	goto L15
L17:
	;
	v64 = v55
	v66 = v57
	goto L18
L18:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)) = uint8(v67)
	v69 = int32(1)
	if v67 != 0 {
		v64 = v64 + v69
		v66 = v66 + v69
		goto L18
	} else {
		goto L20
	}
L19:
	;
	goto L1
L20:
	;
	goto L19
}
func F_string2ean(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v111 int32
	_ = v111
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v156 int32
	_ = v156
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v746 int32
	_ = v746
	var v755 int32
	_ = v755
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v842 int32
	_ = v842
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v876 int64
	_ = v876
	var v889 int64
	_ = v889
	var v890 int32
	_ = v890
	var v916 int64
	_ = v916
	var v926 int32
	_ = v926
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v948 int64
	_ = v948
	var v961 int64
	_ = v961
	var v962 int32
	_ = v962
	var v965 int64
	_ = v965
	var v990 int64
	_ = v990
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1031 int32
	_ = v1031
	var v1037 int32
	_ = v1037
	var v1043 int32
	_ = v1043
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	v5 = int32(0)
	v20 = int64(0)
	v21 = m.G0
	v23 = v21 - int32(112)
	m.G0 = v23
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string2ean[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+96)) = uint8(v26)
	v29 = *(*int64)(unsafe.Add(mBase, _c_F_string2ean[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+88)) = v29
	v32 = *(*int64)(unsafe.Add(mBase, _c_F_string2ean[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+80)) = v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v34 == v5 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v23 + int32(112)
	return v1054
L2:
	;
	v746 = v23 + int32(80)
	goto L185
L3:
	;
	v731 = int32(0)
	v732 = int32(-1)
	goto L2
L4:
	;
	v702 = int32(0)
	v703 = F_errsave_start(m, l1)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L103
	} else {
		goto L180
	}
L5:
	;
	v679 = int32(0)
	v680 = F_errsave_start(m, l1)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L103
	} else {
		goto L175
	}
L6:
	;
	v41 = v23 + int32(80) | int32(3)
	v47 = v34
	v48 = v5
	v49 = v41
	v50 = l0
	v51 = v5
	v54 = int32(1)
	v56 = v5
	goto L7
L7:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	v64 = base.B2i32(v62 == int32(33))
	v67 = v64 | base.B2i32(v62 == int32(0))
	v68 = int32(255)
	v69 = v47 & v68
	v72 = v67 & base.B2i32(v69 == int32(63))
	v79 = v72 | base.B2i32(base.Ui32((v47-int32(48))&v68) < base.Ui32(int32(10)))
	v80 = v72 | v56
	switch v48 {
	case 0:
		goto L17
	default:
		goto L14
	case 7:
		goto L16
	case 9:
		goto L15
	}
L8:
	;
	v182 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v182)
	switch v174 - int32(8) {
	case 0:
		goto L61
	default:
		goto L5
	case 2:
		goto L62
	case 4:
		goto L63
	case 5:
		goto L64
	}
L9:
	;
	v180 = v50 + int32(1)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v181 != 0 {
		v47 = v181
		v48 = v174
		v49 = v175
		v50 = v180
		v51 = v176
		v54 = v177
		v56 = v178
		goto L7
	} else {
		goto L52
	}
L10:
	;
	if v79 == int32(0) {
		goto L5
	} else {
		goto L50
	}
L11:
	;
	if v62 != 0 {
		goto L10
	} else {
		goto L49
	}
L12:
	;
	switch v69 - int32(32) {
	case 0, 13:
		v174 = v48
		v175 = v49
		v176 = v51
		v177 = v54
		v178 = v80
		goto L9
	case 1:
		goto L11
	default:
		goto L10
	}
L13:
	;
	if v79 == int32(0) {
		goto L5
	} else {
		goto L48
	}
L14:
	;
	if v67&(base.B2i32(v48 == int32(11))&v79) == int32(0) {
		goto L12
	} else {
		goto L46
	}
L15:
	;
	if base.B2i32(v69 == int32(88))|v79 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L16:
	;
	if base.B2i32(v69 == int32(88))|v79 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	switch v69 - int32(32) {
	case 0, 13:
		v174 = int32(0)
		v175 = v49
		v176 = v51
		v177 = v54
		v178 = v80
		goto L9
	case 1:
		goto L11
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		goto L13
	default:
		goto L18
	}
L18:
	;
	if base.B2i32(v69 != int32(109))&base.B2i32(v69 != int32(77)) != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	if v51 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v89 = int32(77)
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v89)
	v91 = int32(1)
	v174 = v91
	v175 = v49 + v91
	v176 = int32(4)
	v177 = v54
	v178 = v80
	goto L9
L21:
	;
	if v51 != 0 {
		goto L5
	} else {
		goto L28
	}
L22:
	;
	if base.B2i32(v69 == int32(120))&v67 != 0 {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v62 == int32(33) {
		goto L21
	} else {
		goto L26
	}
L25:
	;
	goto L12
L26:
	;
	if v62 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	if base.Ui32((v47-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v111 = v47 - int32(32)
	goto L31
L30:
	;
	v111 = v47
	goto L31
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v111)
	v174 = int32(8)
	v175 = v49 + int32(1)
	v176 = int32(5)
	v177 = v54
	v178 = v80
	goto L9
L32:
	;
	if v51&int32(-5) != 0 {
		goto L5
	} else {
		goto L39
	}
L33:
	;
	if base.B2i32(v69 == int32(120))&v67 != 0 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v62 == int32(33) {
		goto L32
	} else {
		goto L37
	}
L36:
	;
	goto L12
L37:
	;
	if v62 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	goto L32
L39:
	;
	if base.Ui32((v47-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v135 = v47 - int32(32)
	goto L42
L41:
	;
	v135 = v47
	goto L42
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v135)
	if v51 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v140 = v51
	goto L45
L44:
	;
	v140 = int32(3)
	goto L45
L45:
	;
	v174 = int32(10)
	v175 = v49 + int32(1)
	v176 = v140
	v177 = v54
	v178 = v80
	goto L9
L46:
	;
	if v51 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v47)
	v174 = int32(12)
	v175 = v49 + int32(1)
	v176 = int32(6)
	v177 = v54
	v178 = v80
	goto L9
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v47)
	v156 = int32(1)
	v174 = v156
	v175 = v49 + v156
	v176 = v51
	v177 = v54
	v178 = v80
	goto L9
L49:
	;
	v174 = v48
	v175 = v49
	v176 = v51
	v177 = v54 & v80
	v178 = int32(1)
	goto L9
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v47)
	if base.Ui32(int32(12)) < base.Ui32(v48) {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v170 = int32(1)
	v174 = v48 + v170
	v175 = v49 + v170
	v176 = v51
	v177 = v54
	v178 = v80
	goto L9
L52:
	;
	goto L8
L53:
	;
	v578 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+82)) = uint8(v578)
	if v177&int32(1) == int32(0) {
		goto L3
	} else {
		goto L155
	}
L54:
	;
	v514 = int32(_a_F_string2ean_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+90)) = uint16(v514)
	v517 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_string2ean[3])))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+80)) = uint16(v517)
	v520 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string2ean[4])))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+82)) = uint8(v520)
	if v177&int32(1) == int32(0) {
		goto L3
	} else {
		goto L142
	}
L55:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string2ean[5])))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+82)) = uint8(v453)
	v456 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_string2ean[6])))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+80)) = uint16(v456)
	if v177&int32(1) == int32(0) {
		goto L3
	} else {
		goto L129
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = int32(809056057)
	if v177&int32(1) == int32(0) {
		goto L3
	} else {
		goto L109
	}
L57:
	;
	v340 = int32(0)
	v341 = F_errsave_start(m, l1)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L103
	} else {
		goto L104
	}
L58:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+83)))
	if v301 == int32(48) {
		v333 = int32(6)
		goto L94
	} else {
		goto L95
	}
L59:
	;
	v222 = int32(0)
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v232 = base.B2i32(v230 == int32(77))
	if v230 == int32(77) {
		goto L76
	} else {
		goto L77
	}
L60:
	;
	if base.B2i32(l3 == int32(2))|base.B2i32(l3 != v214) != 0 {
		v337 = v214
		goto L57
	} else {
		goto L74
	}
L61:
	;
	switch v176 {
	case 0, 5:
		goto L72
	default:
		goto L5
	}
L62:
	;
	if base.Ui32(v176-int32(5)) < base.Ui32(int32(-2)) {
		goto L5
	} else {
		goto L68
	}
L63:
	;
	v190 = int32(6)
	if v176 != v190 {
		goto L5
	} else {
		goto L67
	}
L64:
	;
	if v176 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	if v177&int32(1) != 0 {
		goto L59
	} else {
		goto L66
	}
L66:
	;
	v296 = int32(-1)
	v297 = int32(0)
	goto L58
L67:
	;
	v193 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+94)))
	v214 = v190
	v215 = v193 - int32(48)
	goto L60
L68:
	;
	v200 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+92)))
	if v200 == int32(88) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v214 = v176
	v215 = int32(10)
	goto L60
L70:
	;
	goto L71
L71:
	;
	v214 = v176
	v215 = v200 - int32(48)
	goto L60
L72:
	;
	v206 = int32(5)
	v208 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+90)))
	if v208 == int32(88) {
		v214 = v206
		v215 = int32(10)
		goto L60
	} else {
		goto L73
	}
L73:
	;
	v214 = v206
	v215 = v208 - int32(48)
	goto L60
L74:
	;
	switch l3 - int32(4) {
	case 0:
		goto L56
	case 1:
		goto L54
	case 2:
		goto L53
	default:
		goto L55
	}
L75:
	;
	v291 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+95)))
	v296 = v290
	v297 = base.B2i32(v290 == v291-int32(48)) | v178
	goto L58
L76:
	;
	v233 = int32(3)
	goto L78
L77:
	;
	v233 = v222
	goto L78
L78:
	;
	if v230 == int32(0) {
		v278 = v233
		v280 = v222
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v282 = int32(10)
	v287 = base.I32_rem_u_s(v278*int32(3)+v280, v282)
	if v287 != 0 {
		goto L91
	} else {
		goto L92
	}
L80:
	;
	v237 = v41
	v238 = v230
	v239 = v232
	v240 = v233
	v241 = int32(13)
	v242 = v222
	goto L81
L81:
	;
	v247 = (v238 - int32(48)) & int32(255)
	if base.Ui32(v247) <= base.Ui32(int32(9)) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v278 = v264
	v280 = v266
	goto L79
L83:
	;
	v252 = v239 & int32(1)
	if v252 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v263 = v239
	v264 = v240
	v265 = v241
	v266 = v242
	goto L85
L85:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+1)))
	if v268 == int32(0) {
		v278 = v264
		v280 = v266
		goto L79
	} else {
		goto L89
	}
L86:
	;
	v253 = int32(0)
	goto L88
L87:
	;
	v253 = v247
	goto L88
L88:
	;
	v259 = int32(1)
	v263 = v239 + v259
	v264 = (int32(0)-v252)&v247 + v240
	v265 = v241 - v259
	v266 = v253 + v242
	goto L85
L89:
	;
	v271 = int32(1)
	if base.Ui32(v271) < base.Ui32(v265) {
		v237 = v237 + v271
		v238 = v268
		v239 = v263
		v240 = v264
		v241 = v265
		v242 = v266
		goto L81
	} else {
		goto L90
	}
L90:
	;
	goto L82
L91:
	;
	v290 = v282 - v287
	goto L93
L92:
	;
	v290 = int32(0)
	goto L93
L93:
	;
	goto L75
L94:
	;
	if base.B2i32(l3 == int32(2))|base.B2i32(v333 == l3) != 0 {
		v731 = v297
		v732 = v296
		goto L2
	} else {
		goto L102
	}
L95:
	;
	v305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+83)))
	v307 = v305 ^ int32(_a_F_string2ean_1)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+85)))
	if v307|(v308^int32(55)) == int32(0) {
		v333 = int32(5)
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+85)))
	if v315^int32(56)|v307 == int32(0) {
		v333 = int32(3)
		goto L94
	} else {
		goto L97
	}
L97:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v23)+83))
	if v322 == int32(809056057) {
		v333 = int32(4)
		goto L94
	} else {
		goto L98
	}
L98:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+85)))
	if v327^int32(57)|v307 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v331 = int32(2)
	goto L101
L100:
	;
	v331 = int32(3)
	goto L101
L101:
	;
	v333 = v331
	goto L94
L102:
	;
	v337 = v333
	goto L57
L103:
	;
	return int32(0)
L104:
	;
	if v341 == int32(0) {
		v1054 = v340
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = l0
	v351 = int32(2)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(v351)%32))+uint32(_c_F_string2ean[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v355
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v337<<(uint(v351)%32))+uint32(_c_F_string2ean[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v361
	F_errmsg(m, int32(_a_F_string2ean_2), v23+int32(48))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	F_errsave_finish(m, l1, int32(_a_F_string2ean_3), int32(891), int32(_a_F_string2ean_4))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L103
	} else {
		goto L108
	}
L108:
	;
	v1054 = v340
	goto L1
L109:
	;
	v380 = v23 + int32(80)
	v381 = int32(0)
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	v391 = base.B2i32(v389 == int32(77))
	if v389 == int32(77) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v731 = base.B2i32(v449 == v215) | v178
	v732 = v449
	goto L2
L111:
	;
	v392 = int32(3)
	goto L113
L112:
	;
	v392 = v381
	goto L113
L113:
	;
	if v389 == int32(0) {
		v437 = v392
		v439 = v381
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v441 = int32(10)
	v446 = base.I32_rem_u_s(v437*int32(3)+v439, v441)
	if v446 != 0 {
		goto L126
	} else {
		goto L127
	}
L115:
	;
	v396 = v380
	v397 = v389
	v398 = v391
	v399 = v392
	v400 = int32(13)
	v401 = v381
	goto L116
L116:
	;
	v406 = (v397 - int32(48)) & int32(255)
	if base.Ui32(v406) <= base.Ui32(int32(9)) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v437 = v423
	v439 = v425
	goto L114
L118:
	;
	v411 = v398 & int32(1)
	if v411 != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v422 = v398
	v423 = v399
	v424 = v400
	v425 = v401
	goto L120
L120:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+1)))
	if v427 == int32(0) {
		v437 = v423
		v439 = v425
		goto L114
	} else {
		goto L124
	}
L121:
	;
	v412 = int32(0)
	goto L123
L122:
	;
	v412 = v406
	goto L123
L123:
	;
	v418 = int32(1)
	v422 = v398 + v418
	v423 = (int32(0)-v411)&v406 + v399
	v424 = v400 - v418
	v425 = v412 + v401
	goto L120
L124:
	;
	v430 = int32(1)
	if base.Ui32(v430) < base.Ui32(v424) {
		v396 = v396 + v430
		v397 = v427
		v398 = v422
		v399 = v423
		v400 = v424
		v401 = v425
		goto L116
	} else {
		goto L125
	}
L125:
	;
	goto L117
L126:
	;
	v449 = v441 - v446
	goto L128
L127:
	;
	v449 = int32(0)
	goto L128
L128:
	;
	goto L110
L129:
	;
	v463 = int32(0)
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v466 == v463 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v731 = base.B2i32(v511 == v215) | v178
	v732 = v511
	goto L2
L131:
	;
	v511 = int32(0)
	goto L130
L132:
	;
	v469 = v41
	v470 = int32(10)
	v471 = v466
	v472 = v463
	goto L133
L133:
	;
	v477 = (v471 - int32(48)) & int32(255)
	v481 = base.B2i32(base.Ui32(v477) < base.Ui32(int32(10)))
	if base.Ui32(v477) < base.Ui32(int32(10)) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v495 = base.I32_rem_u_s(v483, int32(11))
	if v495 == int32(0) {
		goto L131
	} else {
		goto L141
	}
L135:
	;
	goto L134
L136:
	;
	v482 = v470 * v477
	goto L138
L137:
	;
	v482 = int32(0)
	goto L138
L138:
	;
	v483 = v482 + v472
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+1)))
	if v484 == int32(0) {
		goto L135
	} else {
		goto L139
	}
L139:
	;
	v487 = int32(1)
	v489 = v470 - v481
	if base.Ui32(v487) < base.Ui32(v489) {
		v469 = v469 + v487
		v470 = v489
		v471 = v484
		v472 = v483
		goto L133
	} else {
		goto L140
	}
L140:
	;
	goto L135
L141:
	;
	v511 = int32(11) - v495
	goto L130
L142:
	;
	v527 = int32(0)
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v530 == v527 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v731 = base.B2i32(v575 == v215) | v178
	v732 = v575
	goto L2
L144:
	;
	v575 = int32(0)
	goto L143
L145:
	;
	v533 = v41
	v534 = int32(8)
	v535 = v530
	v536 = v527
	goto L146
L146:
	;
	v541 = (v535 - int32(48)) & int32(255)
	v545 = base.B2i32(base.Ui32(v541) < base.Ui32(int32(10)))
	if base.Ui32(v541) < base.Ui32(int32(10)) {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	v559 = base.I32_rem_u_s(v547, int32(11))
	if v559 == int32(0) {
		goto L144
	} else {
		goto L154
	}
L148:
	;
	goto L147
L149:
	;
	v546 = v534 * v541
	goto L151
L150:
	;
	v546 = int32(0)
	goto L151
L151:
	;
	v547 = v546 + v536
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+1)))
	if v548 == int32(0) {
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v551 = int32(1)
	v553 = v534 - v545
	if base.Ui32(v551) < base.Ui32(v553) {
		v533 = v533 + v551
		v534 = v553
		v535 = v548
		v536 = v547
		goto L146
	} else {
		goto L153
	}
L153:
	;
	goto L148
L154:
	;
	v575 = int32(11) - v559
	goto L143
L155:
	;
	v587 = v23 + int32(80) | int32(2)
	v588 = int32(0)
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	v598 = base.B2i32(v596 == int32(77))
	if v596 == int32(77) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v731 = base.B2i32(v656 == v215) | v178
	v732 = v656
	goto L2
L157:
	;
	v599 = int32(3)
	goto L159
L158:
	;
	v599 = v588
	goto L159
L159:
	;
	if v596 == int32(0) {
		v644 = v599
		v646 = v588
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v648 = int32(10)
	v653 = base.I32_rem_u_s(v644*int32(3)+v646, v648)
	if v653 != 0 {
		goto L172
	} else {
		goto L173
	}
L161:
	;
	v603 = v587
	v604 = v596
	v605 = v598
	v606 = v599
	v607 = int32(13)
	v608 = v588
	goto L162
L162:
	;
	v613 = (v604 - int32(48)) & int32(255)
	if base.Ui32(v613) <= base.Ui32(int32(9)) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v644 = v630
	v646 = v632
	goto L160
L164:
	;
	v618 = v605 & int32(1)
	if v618 != 0 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v629 = v605
	v630 = v606
	v631 = v607
	v632 = v608
	goto L166
L166:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+1)))
	if v634 == int32(0) {
		v644 = v630
		v646 = v632
		goto L160
	} else {
		goto L170
	}
L167:
	;
	v619 = int32(0)
	goto L169
L168:
	;
	v619 = v613
	goto L169
L169:
	;
	v625 = int32(1)
	v629 = v605 + v625
	v630 = (int32(0)-v618)&v613 + v606
	v631 = v607 - v625
	v632 = v619 + v608
	goto L166
L170:
	;
	v637 = int32(1)
	if base.Ui32(v637) < base.Ui32(v631) {
		v603 = v603 + v637
		v604 = v634
		v605 = v629
		v606 = v630
		v607 = v631
		v608 = v632
		goto L162
	} else {
		goto L171
	}
L171:
	;
	goto L163
L172:
	;
	v656 = v648 - v653
	goto L174
L173:
	;
	v656 = int32(0)
	goto L174
L174:
	;
	goto L156
L175:
	;
	if v680 == int32(0) {
		v1054 = v679
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L103
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = l0
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_string2ean[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v692
	F_errmsg(m, int32(_a_F_string2ean_5), v23)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L103
	} else {
		goto L178
	}
L178:
	;
	F_errsave_finish(m, l1, int32(_a_F_string2ean_3), int32(885), int32(_a_F_string2ean_4))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L103
	} else {
		goto L179
	}
L179:
	;
	v1054 = v679
	goto L1
L180:
	;
	if v703 == int32(0) {
		v1054 = v702
		goto L1
	} else {
		goto L181
	}
L181:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L103
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = l0
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_string2ean[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v715
	F_errmsg(m, int32(_a_F_string2ean_6), v23-int32(-64))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L103
	} else {
		goto L183
	}
L183:
	;
	F_errsave_finish(m, l1, int32(_a_F_string2ean_3), int32(897), int32(_a_F_string2ean_4))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L103
	} else {
		goto L184
	}
L184:
	;
	v1054 = v702
	goto L1
L185:
	;
	v755 = int32(*(*int8)(unsafe.Add(mBase, uint32(v746))))
	if v755 != 0 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v764 = base.B2i32(v755 == int32(77))
	if v755 == int32(77) {
		goto L191
	} else {
		goto L192
	}
L187:
	;
	if v755 < int32(33) {
		v746 = v746 + int32(1)
		goto L185
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	goto L186
L190:
	;
	goto L189
L191:
	;
	v765 = int32(3)
	goto L193
L192:
	;
	v765 = int32(0)
	goto L193
L193:
	;
	if v755 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v842 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v746)+13)) = uint8(v842)
	v849 = base.I32_rem_u_s(v829*int32(3)+v834, int32(10))
	if v849 != 0 {
		goto L208
	} else {
		goto L209
	}
L195:
	;
	v829 = v765
	v834 = int32(0)
	goto L194
L196:
	;
	goto L197
L197:
	;
	v777 = v755
	v778 = v765
	v782 = v746
	v783 = int32(0)
	v784 = v764
	v786 = int32(13)
	goto L198
L198:
	;
	v794 = (v777 - int32(48)) & int32(255)
	if base.Ui32(v794) <= base.Ui32(int32(9)) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v829 = v810
	v834 = v811
	goto L194
L200:
	;
	v799 = v784 & int32(1)
	if v799 != 0 {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	v810 = v778
	v811 = v783
	v812 = v784
	v813 = v786
	goto L202
L202:
	;
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v782)+1)))
	if v815 == int32(0) {
		v829 = v810
		v834 = v811
		goto L194
	} else {
		goto L206
	}
L203:
	;
	v800 = int32(0)
	goto L205
L204:
	;
	v800 = v794
	goto L205
L205:
	;
	v802 = int32(1)
	v810 = (int32(0)-v799)&v794 + v778
	v811 = v800 + v783
	v812 = v784 + v802
	v813 = v786 - v802
	goto L202
L206:
	;
	v818 = int32(1)
	if base.Ui32(v818) < base.Ui32(v813) {
		v777 = v815
		v778 = v810
		v782 = v782 + v818
		v783 = v811
		v784 = v812
		v786 = v813
		goto L198
	} else {
		goto L207
	}
L207:
	;
	goto L199
L208:
	;
	v852 = int32(58) - v849
	goto L210
L209:
	;
	v852 = int32(48)
	goto L210
L210:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v746)+12)) = uint8(v852)
	if (v731|v178)&int32(1) != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	if v755 != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	goto L213
L213:
	;
	v926 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string2ean[8])))
	if v926 == int32(1) {
		goto L223
	} else {
		goto L224
	}
L214:
	;
	v862 = v755
	v865 = v746
	v876 = v20
	goto L217
L215:
	;
	v916 = int64(0)
	goto L216
L216:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v916 | base.I64_extend_i32_u(v731^int32(-1))&int64(1)
	v1054 = int32(1)
	goto L1
L217:
	;
	if base.Ui32((v862-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v916 = v889 << (uint(int64(1)) % 64)
	goto L216
L219:
	;
	v889 = v876*int64(10) + base.I64_extend_i32_u(v862)&int64(15)
	goto L221
L220:
	;
	v889 = v876
	goto L221
L221:
	;
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865)+1)))
	if v890 != 0 {
		v862 = v890
		v865 = v865 + int32(1)
		v876 = v889
		goto L217
	} else {
		goto L222
	}
L222:
	;
	goto L218
L223:
	;
	if v755 != 0 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L225
L225:
	;
	v993 = int32(0)
	v994 = F_errsave_start(m, l1)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L103
	} else {
		goto L235
	}
L226:
	;
	v934 = v755
	v937 = v746
	v948 = v20
	goto L229
L227:
	;
	v990 = int64(1)
	goto L228
L228:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v990
	v1054 = int32(1)
	goto L1
L229:
	;
	if base.Ui32((v934-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	v965 = int64(1)
	v990 = v961<<(uint(v965)%64) | v965
	goto L228
L231:
	;
	v961 = v948*int64(10) + base.I64_extend_i32_u(v934)&int64(15)
	goto L233
L232:
	;
	v961 = v948
	goto L233
L233:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937)+1)))
	if v962 != 0 {
		v934 = v962
		v937 = v937 + int32(1)
		v948 = v961
		goto L229
	} else {
		goto L234
	}
L234:
	;
	goto L230
L235:
	;
	if v732 == int32(-1) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	if v994 == int32(0) {
		v1054 = v993
		goto L1
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	if v994 == int32(0) {
		v1054 = v993
		goto L1
	} else {
		goto L243
	}
L239:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L103
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = l0
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_string2ean[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v1008
	F_errmsg(m, int32(_a_F_string2ean_7), v23+int32(16))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L103
	} else {
		goto L241
	}
L241:
	;
	F_errsave_finish(m, l1, int32(_a_F_string2ean_3), int32(871), int32(_a_F_string2ean_4))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L103
	} else {
		goto L242
	}
L242:
	;
	v1054 = v993
	goto L1
L243:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L103
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = l0
	if v732 == int32(10) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1031 = int32(88)
	goto L247
L246:
	;
	v1031 = v732 + int32(48)
	goto L247
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v1031
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_string2ean[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v1037
	F_errmsg(m, int32(_a_F_string2ean_8), v23+int32(32))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L103
	} else {
		goto L248
	}
L248:
	;
	F_errsave_finish(m, l1, int32(_a_F_string2ean_3), int32(878), int32(_a_F_string2ean_4))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L103
	} else {
		goto L249
	}
L249:
	;
	v1054 = v993
	goto L1
}
func F_strlcat(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	v7 = F_memchr(m, l0, int32(0), l2)
	mBase = m.M
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	if v9 == l2 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v9 = v7 - l0
	goto L5
L4:
	;
	v9 = l2
	goto L5
L5:
	;
	goto L2
L6:
	;
	v11 = F_strlen(m, l1)
	mBase = m.M
	goto L1
L7:
	;
	goto L8
L8:
	;
	v12 = l0 + v9
	v13 = l2 - v9
	if v13 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L1
L10:
	;
	v129 = F_strlen(m, v125)
	mBase = m.M
	goto L9
L11:
	;
	v125 = l1
	goto L10
L12:
	;
	goto L13
L13:
	;
	v19 = v13 - int32(1)
	if (v12^l1)&int32(3) != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v122)
	v125 = v118
	goto L10
L15:
	;
	v103 = v98
	v104 = v99
	v105 = v100
	goto L36
L16:
	;
	if v93 == int32(0) {
		v118 = v91
		v119 = v92
		goto L14
	} else {
		goto L35
	}
L17:
	;
	v91 = l1
	v92 = v12
	v93 = v19
	goto L16
L18:
	;
	goto L19
L19:
	;
	v23 = int32(0)
	if base.B2i32(l1&int32(3) == v23)|base.B2i32(v19 == v23) == v23 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v59 == int32(0) {
		v118 = v56
		v119 = v57
		goto L14
	} else {
		goto L29
	}
L21:
	;
	v35 = l1
	v36 = v12
	v37 = v19
	goto L24
L22:
	;
	goto L23
L23:
	;
	v56 = l1
	v57 = v12
	v58 = v19
	v59 = base.B2i32(v19 != v23)
	goto L20
L24:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v39)
	if v39 == int32(0) {
		v98 = v35
		v99 = v36
		v100 = v37
		goto L15
	} else {
		goto L26
	}
L25:
	;
	v56 = v50
	v57 = v44
	v58 = v46
	v59 = v48
	goto L20
L26:
	;
	v43 = int32(1)
	v44 = v36 + v43
	v46 = v37 - v43
	v47 = int32(0)
	v48 = base.B2i32(v46 != v47)
	v50 = v35 + v43
	if v50&int32(3) == v47 {
		v56 = v50
		v57 = v44
		v58 = v46
		v59 = v48
		goto L20
	} else {
		goto L27
	}
L27:
	;
	if v46 != 0 {
		v35 = v50
		v36 = v44
		v37 = v46
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if base.B2i32(v62 == int32(0))|base.B2i32(base.Ui32(v58) < base.Ui32(int32(4))) != 0 {
		v91 = v56
		v92 = v57
		v93 = v58
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v69 = v56
	v70 = v57
	v71 = v58
	goto L31
L31:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v77 = int32(-2139062144)
	if (int32(16843008)-v74|v74)&v77 != v77 {
		v98 = v69
		v99 = v70
		v100 = v71
		goto L15
	} else {
		goto L33
	}
L32:
	;
	v91 = v85
	v92 = v83
	v93 = v87
	goto L16
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v74
	v82 = int32(4)
	v83 = v70 + v82
	v85 = v69 + v82
	v87 = v71 - v82
	if base.Ui32(int32(3)) < base.Ui32(v87) {
		v69 = v85
		v70 = v83
		v71 = v87
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v98 = v91
	v99 = v92
	v100 = v93
	goto L15
L36:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	*(*uint8)(unsafe.Add(mBase, uint32(v104))) = uint8(v107)
	if v107 == int32(0) {
		v118 = v103
		v119 = v104
		goto L14
	} else {
		goto L38
	}
L37:
	;
	v118 = v114
	v119 = v112
	goto L14
L38:
	;
	v111 = int32(1)
	v112 = v104 + v111
	v114 = v103 + v111
	v116 = v105 - v111
	if v116 != 0 {
		v103 = v114
		v104 = v112
		v105 = v116
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
}
func F_strlower_libc_mb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v408 int32
	_ = v408
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	v9 = l3 + int32(1)
	if base.Ui32(v9) < base.Ui32(int32(536870912)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v14 = F_palloc_mul(m, int32(4), v9)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L135
	}
L4:
	;
	return int32(0)
L5:
	;
	F_char2wchar(m, v14, v9, l2, l3, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v20 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = v20
	v28 = v22
	goto L10
L8:
	;
	v45 = v20
	goto L9
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_strlower_libc_mb[0]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51*int32(28))+uint32(_c_F_strlower_libc_mb[1])))
	goto L14
L10:
	;
	v34 = F_casemap(m, v28, int32(0))
	mBase = m.M
	goto L12
L11:
	;
	v45 = v37
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v26<<(uint(int32(2))%32)))) = v34
	v37 = v26 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v14+v37<<(uint(int32(2))%32))))
	if v41 != 0 {
		v26 = v37
		v28 = v41
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v59 = v56*v45 + int32(1)
	v60 = F_palloc(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if v59 == int32(0) {
		v446 = v20
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if base.Ui32(v446+int32(1)) <= base.Ui32(l1) {
		goto L127
	} else {
		goto L128
	}
L17:
	;
	if v12 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v66 = int32(0)
	v72 = m.G0
	v73 = int32(16)
	v74 = v72 - v73
	m.G0 = v74
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v14
	v78 = v74 + int32(12)
	v79 = m.G0
	v81 = v79 - v73
	m.G0 = v81
	if v60 != 0 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	goto L20
L20:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_strlower_libc_mb[2]))
	if v12 != 0 {
		goto L65
	} else {
		goto L66
	}
L21:
	;
	v446 = v233
	goto L16
L22:
	;
	v237 = int32(16)
	m.G0 = v81 + v237
	m.G0 = v74 + v237
	goto L21
L23:
	;
	v233 = v59 - v218
	goto L22
L24:
	;
	if v155 != 0 {
		goto L49
	} else {
		goto L50
	}
L25:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v114 = v59
	v115 = v60
	v119 = v113
	goto L38
L26:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v59) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 == int32(0) {
		v233 = v66
		goto L22
	} else {
		goto L30
	}
L29:
	;
	v155 = v59
	v156 = v60
	goto L24
L30:
	;
	v89 = v86
	v90 = v85
	v92 = v66
	goto L31
L31:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v89) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v233 = v112
	goto L22
L33:
	;
	v101 = int32(-1)
	v104 = F_wcrtomb(m, v81+int32(12), v89)
	mBase = m.M
	if v104 == v101 {
		v233 = v101
		goto L22
	} else {
		goto L36
	}
L34:
	;
	v107 = int32(1)
	goto L35
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v112 = v107 + v92
	if v109 != 0 {
		v89 = v109
		v90 = v90 + int32(4)
		v92 = v112
		goto L31
	} else {
		goto L37
	}
L36:
	;
	v107 = v104
	goto L35
L37:
	;
	goto L32
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if base.Ui32(v123-int32(128)) <= base.Ui32(int32(-128)) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v155 = v145
	v156 = v148
	goto L24
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v151 = v149 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v151
	if base.Ui32(int32(3)) < base.Ui32(v145) {
		v114 = v145
		v115 = v148
		v119 = v151
		goto L38
	} else {
		goto L48
	}
L41:
	;
	if v123 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v123)
	v141 = int32(1)
	v145 = v114 - v141
	v148 = v115 + v141
	goto L40
L44:
	;
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v130)
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v130
	v218 = v114
	goto L23
L45:
	;
	goto L46
L46:
	;
	v134 = int32(-1)
	v135 = F_wcrtomb(m, v115, v123)
	mBase = m.M
	if v135 == v134 {
		v233 = v134
		goto L22
	} else {
		goto L47
	}
L47:
	;
	v145 = v114 - v135
	v148 = v115 + v135
	goto L40
L48:
	;
	goto L39
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v165 = v155
	v166 = v156
	v168 = v164
	goto L52
L50:
	;
	goto L51
L51:
	;
	v233 = v59
	goto L22
L52:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	if base.Ui32(v174-int32(128)) <= base.Ui32(int32(-128)) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L51
L54:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v207 = v205 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v207
	if v201 != 0 {
		v165 = v201
		v166 = v204
		v168 = v207
		goto L52
	} else {
		goto L63
	}
L55:
	;
	if v174 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v174)
	v197 = int32(1)
	v201 = v165 - v197
	v204 = v166 + v197
	goto L54
L58:
	;
	v181 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v181)
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v181
	v218 = v165
	goto L23
L59:
	;
	goto L60
L60:
	;
	v185 = int32(-1)
	v188 = F_wcrtomb(m, v81+int32(12), v174)
	mBase = m.M
	if v188 == v185 {
		v233 = v185
		goto L22
	} else {
		goto L61
	}
L61:
	;
	if base.Ui32(v165) < base.Ui32(v188) {
		v218 = v165
		goto L23
	} else {
		goto L62
	}
L62:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v193 = F_wcrtomb(m, v166, v192)
	mBase = m.M
	v201 = v165 - v188
	v204 = v166 + v188
	goto L54
L63:
	;
	goto L53
L64:
	;
	v256 = int32(0)
	v262 = m.G0
	v263 = int32(16)
	v264 = v262 - v263
	m.G0 = v264
	*(*int32)(unsafe.Add(mBase, uint32(v264)+12)) = v14
	v268 = v264 + int32(12)
	v269 = m.G0
	v271 = v269 - v263
	m.G0 = v271
	if v60 != 0 {
		goto L79
	} else {
		goto L80
	}
L65:
	;
	if v12 == int32(-1) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if v245 == int32(_a_F_strlower_libc_mb_0) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v250 = int32(_a_F_strlower_libc_mb_0)
	goto L70
L69:
	;
	v250 = v12
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strlower_libc_mb[2])) = v250
	goto L67
L71:
	;
	v255 = int32(-1)
	goto L73
L72:
	;
	v255 = v245
	goto L73
L73:
	;
	goto L64
L74:
	;
	if v255 != 0 {
		goto L118
	} else {
		goto L119
	}
L75:
	;
	v427 = int32(16)
	m.G0 = v271 + v427
	m.G0 = v264 + v427
	goto L74
L76:
	;
	v423 = v59 - v408
	goto L75
L77:
	;
	if v345 != 0 {
		goto L102
	} else {
		goto L103
	}
L78:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v304 = v59
	v305 = v60
	v309 = v303
	goto L91
L79:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v59) {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v276 == int32(0) {
		v423 = v256
		goto L75
	} else {
		goto L83
	}
L82:
	;
	v345 = v59
	v346 = v60
	goto L77
L83:
	;
	v279 = v276
	v280 = v275
	v282 = v256
	goto L84
L84:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v279) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v423 = v302
	goto L75
L86:
	;
	v291 = int32(-1)
	v294 = F_wcrtomb(m, v271+int32(12), v279)
	mBase = m.M
	if v294 == v291 {
		v423 = v291
		goto L75
	} else {
		goto L89
	}
L87:
	;
	v297 = int32(1)
	goto L88
L88:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	v302 = v297 + v282
	if v299 != 0 {
		v279 = v299
		v280 = v280 + int32(4)
		v282 = v302
		goto L84
	} else {
		goto L90
	}
L89:
	;
	v297 = v294
	goto L88
L90:
	;
	goto L85
L91:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	if base.Ui32(v313-int32(128)) <= base.Ui32(int32(-128)) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v345 = v335
	v346 = v338
	goto L77
L93:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v341 = v339 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v341
	if base.Ui32(int32(3)) < base.Ui32(v335) {
		v304 = v335
		v305 = v338
		v309 = v341
		goto L91
	} else {
		goto L101
	}
L94:
	;
	if v313 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v305))) = uint8(v313)
	v331 = int32(1)
	v335 = v304 - v331
	v338 = v305 + v331
	goto L93
L97:
	;
	v320 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v305))) = uint8(v320)
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v320
	v408 = v304
	goto L76
L98:
	;
	goto L99
L99:
	;
	v324 = int32(-1)
	v325 = F_wcrtomb(m, v305, v313)
	mBase = m.M
	if v325 == v324 {
		v423 = v324
		goto L75
	} else {
		goto L100
	}
L100:
	;
	v335 = v304 - v325
	v338 = v305 + v325
	goto L93
L101:
	;
	goto L92
L102:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v355 = v345
	v356 = v346
	v358 = v354
	goto L105
L103:
	;
	goto L104
L104:
	;
	v423 = v59
	goto L75
L105:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	if base.Ui32(v364-int32(128)) <= base.Ui32(int32(-128)) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L104
L107:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v397 = v395 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v397
	if v391 != 0 {
		v355 = v391
		v356 = v394
		v358 = v397
		goto L105
	} else {
		goto L116
	}
L108:
	;
	if v364 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	goto L110
L110:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v356))) = uint8(v364)
	v387 = int32(1)
	v391 = v355 - v387
	v394 = v356 + v387
	goto L107
L111:
	;
	v371 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v356))) = uint8(v371)
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v371
	v408 = v355
	goto L76
L112:
	;
	goto L113
L113:
	;
	v375 = int32(-1)
	v378 = F_wcrtomb(m, v271+int32(12), v364)
	mBase = m.M
	if v378 == v375 {
		v423 = v375
		goto L75
	} else {
		goto L114
	}
L114:
	;
	if base.Ui32(v355) < base.Ui32(v378) {
		v408 = v355
		goto L76
	} else {
		goto L115
	}
L115:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	v383 = F_wcrtomb(m, v356, v382)
	mBase = m.M
	v391 = v355 - v378
	v394 = v356 + v378
	goto L107
L116:
	;
	goto L106
L117:
	;
	v446 = v423
	goto L16
L118:
	;
	if v255 == int32(-1) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	goto L124
L121:
	;
	v440 = int32(_a_F_strlower_libc_mb_0)
	goto L123
L122:
	;
	v440 = v255
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strlower_libc_mb[2])) = v440
	goto L120
L124:
	;
	goto L126
L126:
	;
	goto L117
L127:
	;
	if v446 != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	F_pfree(m, v14)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L4
	} else {
		goto L133
	}
L130:
	;
	base.MemoryCopy(m, l0, v60, v446)
	goto L132
L131:
	;
	goto L132
L132:
	;
	v452 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v446))) = uint8(v452)
	goto L129
L133:
	;
	F_pfree(m, v60)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	return v446
L135:
	;
	F_errcode(m, int32(_a_F_strlower_libc_mb_1))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_strlower_libc_mb_2), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_strlower_libc_mb_3), int32(551), int32(_a_F_strlower_libc_mb_4))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_strlower_libc_sb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	if base.Ui32(l1) < base.Ui32(l3+int32(1)) {
	} else {
		if l3 != 0 {
			base.MemoryCopy(m, l0, l2, l3)
		} else {
		}
		v12 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0+l3))) = uint8(v12)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v14 == v12 {
		} else {
			v17 = l0
			v19 = v14
			for {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
				if v22 == int32(1) {
					if base.Ui32((v19-int32(65))&int32(255)) <= base.Ui32(int32(25)) {
						v61 = v19 | int32(32)
						*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v61)
					} else {
						if int32(0) <= base.I32_extend8_s(v19) {
						} else {
							if base.B2i32(base.Ui32(v19&int32(255)-int32(65)) < base.Ui32(int32(26))) == int32(0) {
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
								if base.Ui32(v44-int32(65)) < base.Ui32(int32(26)) {
									v51 = v44 | int32(32)
								} else {
									v51 = v44
								}
								v61 = v51
								*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v61)
							}
						}
					}
				} else {
					v53 = v19 & int32(255)
					if base.Ui32(v53-int32(65)) < base.Ui32(int32(26)) {
						v60 = v53 | int32(32)
					} else {
						v60 = v53
					}
					v61 = v60
					*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v61)
				}
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
				if v65 != 0 {
					v17 = v17 + int32(1)
					v19 = v65
					continue
				} else {
					break
				}
				break
			}
		}
	}
	return l3
}
func F_strncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v11 = l0
	v12 = l1
	v13 = l2
	v14 = v10
	goto L8
L5:
	;
	v37 = l1
	v41 = int32(0)
	goto L6
L6:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	return v41 - v42
L7:
	;
	v37 = v32
	v41 = v34
	goto L6
L8:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if base.B2i32(v14 != v16)|base.B2i32(v16 == int32(0)) != 0 {
		v32 = v12
		v34 = v14
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v32 = v26
	v34 = int32(0)
	goto L7
L10:
	;
	v22 = v13 - int32(1)
	if v22 == int32(0) {
		v32 = v12
		v34 = v14
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v25 = int32(1)
	v26 = v12 + v25
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v27 != 0 {
		v11 = v11 + v25
		v12 = v26
		v13 = v22
		v14 = v27
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
}
func F_strncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int64
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	if (l1^l0)&int32(3) != 0 {
		v74 = l1
		v75 = l2
		v76 = l0
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v111 = int32(0)
	if v108 == v111 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	v108 = int32(0)
	v109 = v103
	goto L1
L3:
	;
	v86 = v81
	v87 = v82
	v88 = v83
	goto L21
L4:
	;
	if v75 == int32(0) {
		v103 = v76
		goto L2
	} else {
		goto L20
	}
L5:
	;
	v9 = int32(0)
	if base.B2i32(l1&int32(3) == v9)|base.B2i32(l2 == v9) != 0 {
		v40 = l1
		v41 = l2
		v42 = l0
		v43 = base.B2i32(l2 != v9)
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v43 == int32(0) {
		v103 = v42
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v19 = l1
	v20 = l2
	v21 = l0
	goto L8
L8:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v23)
	if v23 == int32(0) {
		v108 = v20
		v109 = v21
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v40 = v34
	v41 = v30
	v42 = v28
	v43 = v32
	goto L6
L10:
	;
	v27 = int32(1)
	v28 = v21 + v27
	v30 = v20 - v27
	v31 = int32(0)
	v32 = base.B2i32(v30 != v31)
	v34 = v19 + v27
	if v34&int32(3) == v31 {
		v40 = v34
		v41 = v30
		v42 = v28
		v43 = v32
		goto L6
	} else {
		goto L11
	}
L11:
	;
	if v30 != 0 {
		v19 = v34
		v20 = v30
		v21 = v28
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v46 == int32(0) {
		v108 = v41
		v109 = v42
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(v41) < base.Ui32(int32(4)) {
		v74 = v40
		v75 = v41
		v76 = v42
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v52 = v40
	v53 = v41
	v54 = v42
	goto L16
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v60 = int32(-2139062144)
	if (int32(16843008)-v57|v57)&v60 != v60 {
		v81 = v52
		v82 = v53
		v83 = v54
		goto L3
	} else {
		goto L18
	}
L17:
	;
	v74 = v68
	v75 = v70
	v76 = v66
	goto L4
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v57
	v65 = int32(4)
	v66 = v54 + v65
	v68 = v52 + v65
	v70 = v53 - v65
	if base.Ui32(int32(3)) < base.Ui32(v70) {
		v52 = v68
		v53 = v70
		v54 = v66
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v81 = v74
	v82 = v75
	v83 = v76
	goto L3
L21:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v90)
	if v90 == int32(0) {
		v108 = v87
		v109 = v88
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v103 = v95
	goto L2
L23:
	;
	v94 = int32(1)
	v95 = v88 + v94
	v99 = v87 - v94
	if v99 != 0 {
		v86 = v86 + v94
		v87 = v99
		v88 = v95
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	return l0
L26:
	;
	goto L25
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v111)
	v118 = v109 + v108
	*(*uint8)(unsafe.Add(mBase, uint32(v118-int32(1)))) = uint8(v111)
	if base.Ui32(v108) < base.Ui32(int32(3)) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+2)) = uint8(v111)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)) = uint8(v111)
	*(*uint8)(unsafe.Add(mBase, uint32(v118-int32(3)))) = uint8(v111)
	*(*uint8)(unsafe.Add(mBase, uint32(v118-int32(2)))) = uint8(v111)
	if base.Ui32(v108) < base.Ui32(int32(7)) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+3)) = uint8(v111)
	*(*uint8)(unsafe.Add(mBase, uint32(v118-int32(4)))) = uint8(v111)
	if base.Ui32(v108) < base.Ui32(int32(9)) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v140 = int32(0)
	v143 = (v140 - v109) & int32(3)
	v144 = v109 + v143
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v140
	v152 = (v108 - v143) & int32(-4)
	v153 = v144 + v152
	*(*int32)(unsafe.Add(mBase, uint32(v153-int32(4)))) = v140
	if base.Ui32(v152) < base.Ui32(int32(9)) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+8)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v153-int32(8)))) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v153-int32(12)))) = v140
	if base.Ui32(v152) < base.Ui32(int32(25)) {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v144)+20)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v144)+16)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v144)+12)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v153-int32(16)))) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v153-int32(20)))) = v140
	v179 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v153-v179))) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v153-int32(28)))) = v140
	v188 = v144&int32(4) | v179
	v189 = v152 - v188
	if base.Ui32(v189) < base.Ui32(int32(32)) {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v194 = base.I64_extend_i32_u(v140) * int64(4294967297)
	v197 = v188 + v144
	v198 = v189
	goto L34
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v197)+24)) = v194
	*(*int64)(unsafe.Add(mBase, uint32(v197)+16)) = v194
	*(*int64)(unsafe.Add(mBase, uint32(v197)+8)) = v194
	*(*int64)(unsafe.Add(mBase, uint32(v197))) = v194
	v206 = int32(32)
	v209 = v198 - v206
	if base.Ui32(int32(31)) < base.Ui32(v209) {
		v197 = v197 + v206
		v198 = v209
		goto L34
	} else {
		goto L36
	}
L35:
	;
	goto L26
L36:
	;
	goto L35
}
func F_strsep(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if v12 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L3
L3:
	;
	return v4
L4:
	;
	v71 = v63 - v4 + v4
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v72 != 0 {
		goto L22
	} else {
		goto L23
	}
L5:
	;
	m.G0 = v10 + int32(32)
	goto L4
L6:
	;
	F___memset(m, v10, int32(0), int32(32))
	mBase = m.M
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v18 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v13 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v14 = F___strchrnul(m, v4, v12)
	mBase = m.M
	v63 = v14
	goto L5
L10:
	;
	goto L9
L11:
	;
	v20 = l1
	v21 = v18
	goto L14
L12:
	;
	goto L13
L13:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v42 == int32(0) {
		v63 = v4
		goto L5
	} else {
		goto L17
	}
L14:
	;
	v28 = v10 + int32(base.Ui32(v21)>>(uint(int32(3))%32))&int32(28)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v30 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v29 | v30<<(uint(v21)%32)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v34 != 0 {
		v20 = v20 + v30
		v21 = v34
		goto L14
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	goto L15
L17:
	;
	v46 = v4
	v47 = v42
	goto L18
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(base.Ui32(v47)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v55)>>(uint(v47)%32))&int32(1) != 0 {
		v63 = v46
		goto L5
	} else {
		goto L20
	}
L19:
	;
	v63 = v61
	goto L5
L20:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
	v61 = v46 + int32(1)
	if v59 != 0 {
		v46 = v61
		v47 = v59
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v73 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v73)
	v78 = v71 + int32(1)
	goto L24
L23:
	;
	v78 = int32(0)
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v78
	goto L3
}
func F_strtitle_builtin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v6)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+13)) = uint16(v6)
	v20 = int32(1)
	v21 = v15 ^ v20
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)) = uint8(v21)
	v25 = F_convert_case(m, l0, l1, l2, l3, v20, v15, int32(1615), v9)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		m.G0 = v9 + int32(16)
		return v25
	}
}
func F_strtoul(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(4294967295))
	return base.I32_wrap_i64(v5)
}
func F_strxfrm_libc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v7 = F_strlen(m, l2)
	if base.Ui32(v7) < base.Ui32(l1) {
		v9 = F_strcpy(m, l0, l2)
	} else {
	}
	return v7
}
func F_subre(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	v2 = l1
	v3 = l2
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = m.T0[v10].(func(*base.Module) int32)(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v17 != 0 {
				v19 = v17
			} else {
				v19 = int32(19)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v19
			return int32(0)
		} else {
			if v7 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v23
				v42 = v7
				v43 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v42)+4)) = v43
				v45 = int32(255)
				*(*uint8)(unsafe.Add(mBase, uint32(v42)+2)) = uint8(v45)
				*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)) = uint8(v3)
				*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v2)
				*(*int32)(unsafe.Add(mBase, uint32(v42)+36)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v42)+32)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v42)+28)) = l3
				*(*int64)(unsafe.Add(mBase, uint32(v42)+20)) = v43
				*(*int64)(unsafe.Add(mBase, uint32(v42)+12)) = int64(281479271677952)
				return v42
			} else {
				v27 = F_palloc_extended(m, int32(88), int32(2))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v33 != 0 {
							v35 = v33
						} else {
							v35 = int32(12)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v35
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+84)) = v39
						*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v27
						v42 = v27
						v43 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v42)+4)) = v43
						v45 = int32(255)
						*(*uint8)(unsafe.Add(mBase, uint32(v42)+2)) = uint8(v45)
						*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)) = uint8(v3)
						*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v2)
						*(*int32)(unsafe.Add(mBase, uint32(v42)+36)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v42)+32)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v42)+28)) = l3
						*(*int64)(unsafe.Add(mBase, uint32(v42)+20)) = v43
						*(*int64)(unsafe.Add(mBase, uint32(v42)+12)) = int64(281479271677952)
						return v42
					}
				}
			}
		}
	}
}
func F_substitute_actual_parameters_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 == v3 {
		v46 = v3
		m.G0 = v7 + int32(32)
		return v46
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v12 == int32(8) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v15 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v55
					F_errmsg_internal(m, int32(_a_F_substitute_actual_parameters_mutator_0), v7+int32(16))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_substitute_actual_parameters_mutator_1), int32(_a_F_substitute_actual_parameters_mutator_2), int32(_a_F_substitute_actual_parameters_mutator_3))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v16 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v71
						F_errmsg_internal(m, int32(_a_F_substitute_actual_parameters_mutator_4), v7)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_substitute_actual_parameters_mutator_1), int32(_a_F_substitute_actual_parameters_mutator_5), int32(_a_F_substitute_actual_parameters_mutator_3))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v19 < v16 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v71
							F_errmsg_internal(m, int32(_a_F_substitute_actual_parameters_mutator_4), v7)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_substitute_actual_parameters_mutator_1), int32(_a_F_substitute_actual_parameters_mutator_5), int32(_a_F_substitute_actual_parameters_mutator_3))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v22 = int32(2)
						v25 = int32(4)
						v26 = v21 + v16<<(uint(v22)%32) - v25
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27 + int32(1)
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v32+v33<<(uint(v22)%32)-v25)))
						v46 = v39
						m.G0 = v7 + int32(32)
						return v46
					}
				}
			}
		} else {
			v41 = F_expression_tree_mutator_impl(m, l0, int32(927), l1)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v46 = v41
				m.G0 = v7 + int32(32)
				return v46
			}
		}
	}
}
func F_summarizer_read_local_xlog_page(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v31 int64
	_ = v31
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v125 int32
	_ = v125
	var v129 int64
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int32(_a_F_summarizer_read_local_xlog_page_0)
	v22 = l1 - int64(-8192)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v23)+8))
	if base.Ui64(v22) <= base.Ui64(v24) {
		v152 = v20
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v14 - int32(-64)
	return v178
L4:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v158 = v12 + int32(-40)
	v159 = F_WALRead(m, l0, l4, l1, v152, v156, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L42
	}
L5:
	;
	v31 = v24
	goto L6
L6:
	;
	if base.Ui64(v31) < base.Ui64(l1+base.I64_extend_i32_s(l2)) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v152 = base.I32_wrap_i64(v31 - l1)
	goto L4
L8:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
	if v40 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	v43 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+16)) = uint8(v43)
	v178 = int32(-1)
	goto L3
L12:
	;
	goto L13
L13:
	;
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[0]))
	if v50 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_pgstat_report_wal(m, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L27
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[1])) = v73
	goto L15
L17:
	;
	v53 = int32(150)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[1]))
	v57 = v55 << (uint(int32(1)) % 32)
	if v53 <= v57 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	if v50 < int32(2) {
		goto L15
	} else {
		goto L23
	}
L20:
	;
	v60 = v53
	goto L22
L21:
	;
	v60 = v57
	goto L22
L22:
	;
	v73 = v60
	goto L16
L23:
	;
	v63 = int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[1]))
	if v65-v63 < v50 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v70 = v63
	goto L26
L25:
	;
	v70 = v65 - v50
	goto L26
L26:
	;
	v73 = v70
	goto L16
L27:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[2]))
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[1]))
	v88 = F_WaitLatch(m, v81, int32(41), v84*int32(200), int32(83886096))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[2]))
	v92 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v92
	v97 = base.AtomicRmwOr32(m, v92, int32(_a_F_summarizer_read_local_xlog_page_1), v92)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[0])) = int32(0)
	v103 = F_GetLatestLSN(m, v12+int32(-40))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v105 == v106 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if base.Ui64(v141) < base.Ui64(v22) {
		v31 = v141
		goto L6
	} else {
		goto L41
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v103
	v141 = v103
	goto L31
L33:
	;
	goto L34
L34:
	;
	v109 = F_readTimeLineHistory(m, v105)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v111 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)) = uint8(v111)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v115 = F_tliSwitchPoint(m, v113, v109, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v115
	v120 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v23)+8))
	if v120 == int32(0) {
		v141 = v122
		goto L31
	} else {
		goto L38
	}
L38:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v125
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)) = uint32(v122)
	v129 = int64(base.Ui64(v122) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+4)) = uint32(v129)
	F_errmsg_internal(m, int32(_a_F_summarizer_read_local_xlog_page_2), v14)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_summarizer_read_local_xlog_page_3), int32(1738), int32(_a_F_summarizer_read_local_xlog_page_4))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v23)+8))
	v141 = v139
	goto L31
L41:
	;
	v152 = v20
	goto L4
L42:
	;
	if v159 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_WALReadRaiseError(m, v158)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v165 = int32(_a_F_summarizer_read_local_xlog_page_5)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[0])) = v167 + int32(1)
	v178 = v152
	goto L3
L46:
	;
	goto L45
}
func F_symlink(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	v4 = m.Env.X__syscall_symlinkat(m, l0, int32(-100), l1)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v4) {
		*(*int32)(unsafe.Add(mBase, _c_F_symlink[0])) = int32(0) - v4
		v12 = int32(-1)
	} else {
		v12 = v4
	}
	return v12
}
func F_syslog(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v219 int32
	_ = v219
	v11 = m.G0
	v12 = int32(16)
	v13 = v11 - v12
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l2
	v16 = m.G0
	v18 = v16 - v12
	m.G0 = v18
	if base.Ui32(int32(1023)) < base.Ui32(l0) {
		v219 = int32(16)
		m.G0 = v18 + v219
		m.G0 = v13 + v219
		return
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[0]))
		if v23&(int32(1)<<(uint(l0&int32(7))%32)) == int32(0) {
			v219 = int32(16)
			m.G0 = v18 + v219
			m.G0 = v13 + v219
			return
		} else {
			v31 = m.G0
			v33 = v31 - int32(1168)
			m.G0 = v33
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[1]))
			v38 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[2]))
			if v38 < int32(0) {
				v41 = int32(0)
				v46 = F_socket(m, int32(1), int32(_a_F_syslog_0), v41)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, _c_F_syslog[2])) = v46
				if v41 <= v46 {
					v50 = F_connect(m, v46)
					mBase = m.M
				} else {
				}
			} else {
			}
			v52 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[3]))
			v53 = F_time(m)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, uint32(v33)+1144)) = v53
			v58 = v33 + int32(1100)
			v59 = F___gmtime_r(m, v33+int32(1144), v58)
			mBase = m.M
			v61 = v33 + int32(1152)
			v65 = F___strftime_l(m, v61, int32(16), int32(_a_F_syslog_1), v58, int32(_a_F_syslog_2))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
				if v70&int32(1) != 0 {
					v73 = int32(42)
				} else {
					v73 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = v73
				v75 = int32(0)
				v76 = base.B2i32(v73 == v75)
				*(*int32)(unsafe.Add(mBase, uint32(v33)+56)) = v76 + int32(_a_F_syslog_3)
				*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v76 + int32(_a_F_syslog_4)
				*(*int32)(unsafe.Add(mBase, uint32(v33)+44)) = int32(_a_F_syslog_5)
				if l0&int32(1016) != 0 {
					v88 = v75
				} else {
					v88 = v52
				}
				*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v88 | l0
				*(*int32)(unsafe.Add(mBase, uint32(v33)+40)) = v33 + int32(60)
				*(*int32)(unsafe.Add(mBase, uint32(v33)+36)) = v61
				v96 = v33 - int32(-64)
				v101 = F_snprintf(m, v96, int32(1024), int32(_a_F_syslog_6), v33+int32(32))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_syslog[1])) = v36
					v107 = int32(1024) - v101
					v108 = F_vsnprintf(m, v101+v96, v107, l1, l2)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						if v108 < int32(0) {
							m.G0 = v33 + int32(1168)
							v219 = int32(16)
							m.G0 = v18 + v219
							m.G0 = v13 + v219
							return
						} else {
							if base.Ui32(v107) <= base.Ui32(v108) {
								v115 = int32(1023)
							} else {
								v115 = v101 + v108
							}
							v116 = v96 + v115
							v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116-int32(1)))))
							if v119 != int32(10) {
								v122 = int32(10)
								*(*uint8)(unsafe.Add(mBase, uint32(v116))) = uint8(v122)
								v126 = v115 + int32(1)
							} else {
								v126 = v115
							}
							v128 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[2]))
							v131 = F_sendto(m, v128, v33-int32(-64), v126)
							mBase = m.M
							if int32(0) <= v131 {
								v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
								if v189&int32(32) == int32(0) {
									m.G0 = v33 + int32(1168)
									v219 = int32(16)
									m.G0 = v18 + v219
									m.G0 = v13 + v219
									return
								} else {
									v194 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
									*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v194
									*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v194 + (v33 - int32(-64))
									F_dprintf(m, int32(2), v33)
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return
									} else {
										m.G0 = v33 + int32(1168)
										v219 = int32(16)
										m.G0 = v18 + v219
										m.G0 = v13 + v219
										return
									}
								}
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[1]))
								if base.B2i32(base.Ui32(v136-int32(14)) < base.Ui32(int32(2)))|base.B2i32(v136 == int32(53)) != 0 {
									v147 = int32(1)
								} else {
									v147 = base.B2i32(v136 == int32(64))
								}
								if v147 == int32(0) {
									v163 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
									if v163&int32(2) == int32(0) {
										v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
										if v189&int32(32) == int32(0) {
											m.G0 = v33 + int32(1168)
											v219 = int32(16)
											m.G0 = v18 + v219
											m.G0 = v13 + v219
											return
										} else {
											v194 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v194
											*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v194 + (v33 - int32(-64))
											F_dprintf(m, int32(2), v33)
											mBase = m.M
											v203 = m.ExcPending
											if v203 != 0 {
												return
											} else {
												m.G0 = v33 + int32(1168)
												v219 = int32(16)
												m.G0 = v18 + v219
												m.G0 = v13 + v219
												return
											}
										}
									} else {
										v170 = int32(0)
										v171 = F_open(m, int32(_a_F_syslog_7), int32(_a_F_syslog_8), v170)
										mBase = m.M
										if v171 < v170 {
											v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
											if v189&int32(32) == int32(0) {
												m.G0 = v33 + int32(1168)
												v219 = int32(16)
												m.G0 = v18 + v219
												m.G0 = v13 + v219
												return
											} else {
												v194 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v194
												*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v194 + (v33 - int32(-64))
												F_dprintf(m, int32(2), v33)
												mBase = m.M
												v203 = m.ExcPending
												if v203 != 0 {
													return
												} else {
													m.G0 = v33 + int32(1168)
													v219 = int32(16)
													m.G0 = v18 + v219
													m.G0 = v13 + v219
													return
												}
											}
										} else {
											v174 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v126 - v174
											*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v174 + (v33 - int32(-64))
											F_dprintf(m, v171, v33+int32(16))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return
											} else {
												v185 = F_close(m, v171)
												mBase = m.M
												v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
												if v189&int32(32) == int32(0) {
													m.G0 = v33 + int32(1168)
													v219 = int32(16)
													m.G0 = v18 + v219
													m.G0 = v13 + v219
													return
												} else {
													v194 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v194
													*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v194 + (v33 - int32(-64))
													F_dprintf(m, int32(2), v33)
													mBase = m.M
													v203 = m.ExcPending
													if v203 != 0 {
														return
													} else {
														m.G0 = v33 + int32(1168)
														v219 = int32(16)
														m.G0 = v18 + v219
														m.G0 = v13 + v219
														return
													}
												}
											}
										}
									}
								} else {
									v151 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[2]))
									v152 = F_connect(m, v151)
									mBase = m.M
									if v152 < int32(0) {
										v163 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
										if v163&int32(2) == int32(0) {
											v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
											if v189&int32(32) == int32(0) {
												m.G0 = v33 + int32(1168)
												v219 = int32(16)
												m.G0 = v18 + v219
												m.G0 = v13 + v219
												return
											} else {
												v194 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v194
												*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v194 + (v33 - int32(-64))
												F_dprintf(m, int32(2), v33)
												mBase = m.M
												v203 = m.ExcPending
												if v203 != 0 {
													return
												} else {
													m.G0 = v33 + int32(1168)
													v219 = int32(16)
													m.G0 = v18 + v219
													m.G0 = v13 + v219
													return
												}
											}
										} else {
											v170 = int32(0)
											v171 = F_open(m, int32(_a_F_syslog_7), int32(_a_F_syslog_8), v170)
											mBase = m.M
											if v171 < v170 {
												v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
												if v189&int32(32) == int32(0) {
													m.G0 = v33 + int32(1168)
													v219 = int32(16)
													m.G0 = v18 + v219
													m.G0 = v13 + v219
													return
												} else {
													v194 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v194
													*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v194 + (v33 - int32(-64))
													F_dprintf(m, int32(2), v33)
													mBase = m.M
													v203 = m.ExcPending
													if v203 != 0 {
														return
													} else {
														m.G0 = v33 + int32(1168)
														v219 = int32(16)
														m.G0 = v18 + v219
														m.G0 = v13 + v219
														return
													}
												}
											} else {
												v174 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v126 - v174
												*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v174 + (v33 - int32(-64))
												F_dprintf(m, v171, v33+int32(16))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return
												} else {
													v185 = F_close(m, v171)
													mBase = m.M
													v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
													if v189&int32(32) == int32(0) {
														m.G0 = v33 + int32(1168)
														v219 = int32(16)
														m.G0 = v18 + v219
														m.G0 = v13 + v219
														return
													} else {
														v194 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
														*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v194
														*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v194 + (v33 - int32(-64))
														F_dprintf(m, int32(2), v33)
														mBase = m.M
														v203 = m.ExcPending
														if v203 != 0 {
															return
														} else {
															m.G0 = v33 + int32(1168)
															v219 = int32(16)
															m.G0 = v18 + v219
															m.G0 = v13 + v219
															return
														}
													}
												}
											}
										}
									} else {
										v156 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[2]))
										v159 = F_sendto(m, v156, v33-int32(-64), v126)
										mBase = m.M
										if int32(0) <= v159 {
											v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
											if v189&int32(32) == int32(0) {
												m.G0 = v33 + int32(1168)
												v219 = int32(16)
												m.G0 = v18 + v219
												m.G0 = v13 + v219
												return
											} else {
												v194 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v194
												*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v194 + (v33 - int32(-64))
												F_dprintf(m, int32(2), v33)
												mBase = m.M
												v203 = m.ExcPending
												if v203 != 0 {
													return
												} else {
													m.G0 = v33 + int32(1168)
													v219 = int32(16)
													m.G0 = v18 + v219
													m.G0 = v13 + v219
													return
												}
											}
										} else {
											v163 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
											if v163&int32(2) == int32(0) {
												v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
												if v189&int32(32) == int32(0) {
													m.G0 = v33 + int32(1168)
													v219 = int32(16)
													m.G0 = v18 + v219
													m.G0 = v13 + v219
													return
												} else {
													v194 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v194
													*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v194 + (v33 - int32(-64))
													F_dprintf(m, int32(2), v33)
													mBase = m.M
													v203 = m.ExcPending
													if v203 != 0 {
														return
													} else {
														m.G0 = v33 + int32(1168)
														v219 = int32(16)
														m.G0 = v18 + v219
														m.G0 = v13 + v219
														return
													}
												}
											} else {
												v170 = int32(0)
												v171 = F_open(m, int32(_a_F_syslog_7), int32(_a_F_syslog_8), v170)
												mBase = m.M
												if v171 < v170 {
													v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
													if v189&int32(32) == int32(0) {
														m.G0 = v33 + int32(1168)
														v219 = int32(16)
														m.G0 = v18 + v219
														m.G0 = v13 + v219
														return
													} else {
														v194 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
														*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v194
														*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v194 + (v33 - int32(-64))
														F_dprintf(m, int32(2), v33)
														mBase = m.M
														v203 = m.ExcPending
														if v203 != 0 {
															return
														} else {
															m.G0 = v33 + int32(1168)
															v219 = int32(16)
															m.G0 = v18 + v219
															m.G0 = v13 + v219
															return
														}
													}
												} else {
													v174 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v126 - v174
													*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v174 + (v33 - int32(-64))
													F_dprintf(m, v171, v33+int32(16))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return
													} else {
														v185 = F_close(m, v171)
														mBase = m.M
														v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
														if v189&int32(32) == int32(0) {
															m.G0 = v33 + int32(1168)
															v219 = int32(16)
															m.G0 = v18 + v219
															m.G0 = v13 + v219
															return
														} else {
															v194 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
															*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v194
															*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v194 + (v33 - int32(-64))
															F_dprintf(m, int32(2), v33)
															mBase = m.M
															v203 = m.ExcPending
															if v203 != 0 {
																return
															} else {
																m.G0 = v33 + int32(1168)
																v219 = int32(16)
																m.G0 = v18 + v219
																m.G0 = v13 + v219
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
