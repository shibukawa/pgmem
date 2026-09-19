package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
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
										F_errfinish(m, int32(_a_F_SICleanupQueue_3), int32(673), int32(_a_F_SICleanupQueue_4))
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
									F_errfinish(m, int32(_a_F_SICleanupQueue_3), int32(673), int32(_a_F_SICleanupQueue_4))
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 float64
	_ = v35
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v53 float64
	_ = v53
	var v74 float64
	_ = v74
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v91 float64
	_ = v91
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v102 float64
	_ = v102
	var v105 float64
	_ = v105
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v111 float64
	_ = v111
	var v112 float64
	_ = v112
	var v114 int32
	_ = v114
	var v115 float64
	_ = v115
	var v119 float64
	_ = v119
	var v131 float64
	_ = v131
	var v136 float64
	_ = v136
	var v137 float64
	_ = v137
	var v138 float64
	_ = v138
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v146 float64
	_ = v146
	var v149 float64
	_ = v149
	var v152 float64
	_ = v152
	var v158 float64
	_ = v158
	var v159 int32
	_ = v159
	var v160 float64
	_ = v160
	var v163 float64
	_ = v163
	var v167 float64
	_ = v167
	var v168 float64
	_ = v168
	var v172 float64
	_ = v172
	var v180 float64
	_ = v180
	var v181 float64
	_ = v181
	var v184 float64
	_ = v184
	var v194 float64
	_ = v194
	var v216 float64
	_ = v216
	var v219 float64
	_ = v219
	var v221 float64
	_ = v221
	var v222 float64
	_ = v222
	var v225 float64
	_ = v225
	var v233 float64
	_ = v233
	var v253 float64
	_ = v253
	var v262 float64
	_ = v262
	var v270 float64
	_ = v270
	var v277 int32
	_ = v277
	var v278 float64
	_ = v278
	var v279 float64
	_ = v279
	var v284 float64
	_ = v284
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
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
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l5)+168))
	F_MemoryContextReset(m, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L63
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l5)+60))
	v23 = F_IvfflatCheckNorm(m, v21, v22, v19)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v27 < v17 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if v23 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v284 = *(*float64)(unsafe.Add(mBase, uint32(l5)+136))
	*(*float64)(unsafe.Add(mBase, uint32(l5)+136)) = base.F64_add(v284, float64(1))
	goto L6
L13:
	;
	F_VectorArraySet(m, v16, v27, v19)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v35 = *(*float64)(unsafe.Add(mBase, uint32(l5)+144))
	if base.F64_lt(v35, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v31 + int32(1)
	goto L12
L17:
	;
	v39 = l5 + int32(112)
	v40 = *(*float64)(unsafe.Add(mBase, uint32(l5)+136))
	v41 = float64(0)
	v53 = base.F64_convert_i32_s(v17)
	if base.F64_ge(base.F64_mul(v53, float64(22)), v40) != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v262 = v35
	goto L19
L19:
	;
	if base.F64_le(v262, float64(0)) != 0 {
		goto L55
	} else {
		goto L56
	}
L20:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5)+144)) = v253
	v262 = v253
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
	v107 = float64(1)
	v108 = base.F64_add(v40, v107)
	v109 = base.F64_sub(v40, v53)
	v111 = base.F64_add(v109, v107)
	v112 = base.F64_div(v108, v111)
	v114 = l5 + int32(120)
	v115 = *(*float64)(unsafe.Add(mBase, uint32(v39)))
	v119 = v115
	goto L32
L25:
	;
	v74 = F_pg_prng_double(m, l5+int32(120))
	mBase = m.M
	if base.F64_eq(v74, float64(0)) != 0 {
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v78 = base.F64_add(v40, float64(1))
	v80 = base.F64_div(base.F64_sub(v78, v53), v78)
	if base.F64_gt(v80, v74) == int32(0) {
		v253 = v41
		goto L21
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v87 = v78
	v88 = v80
	v91 = v41
	goto L29
L29:
	;
	v99 = float64(1)
	v100 = base.F64_add(v91, v99)
	v102 = base.F64_add(v87, v99)
	v105 = base.F64_mul(v88, base.F64_div(base.F64_sub(v102, v53), v102))
	if base.F64_gt(v105, v74) != 0 {
		v87 = v102
		v88 = v105
		v91 = v100
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v253 = v100
	goto L21
L31:
	;
	goto L30
L32:
	;
	v131 = F_pg_prng_double(m, v114)
	mBase = m.M
	if base.F64_eq(v131, float64(0)) != 0 {
		goto L32
	} else {
		goto L34
	}
L33:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v39))) = v233
	v253 = v137
	goto L21
L34:
	;
	v136 = base.F64_mul(v40, base.F64_add(v119, float64(-1)))
	v137 = base.F64_floor(v136)
	v138 = base.F64_add(v111, v137)
	v142 = base.F64_add(v40, v136)
	v144 = F_log(m, base.F64_div(base.F64_mul(v138, base.F64_mul(v112, base.F64_mul(v112, v131))), v142))
	mBase = m.M
	v146 = F_exp(m, base.F64_div(v144, v53))
	mBase = m.M
	v149 = base.F64_div(base.F64_mul(v111, base.F64_div(v142, v138)), v40)
	if base.F64_le(v146, v149) != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v233 = base.F64_div(v149, v146)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v152 = base.F64_add(v40, v137)
	v158 = base.F64_div(base.F64_mul(base.F64_add(v152, float64(1)), base.F64_div(base.F64_mul(v108, v131), v111)), v142)
	v159 = base.F64_lt(v53, v137)
	if v159 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v160 = v138
	goto L41
L40:
	;
	v160 = v108
	goto L41
L41:
	;
	if base.F64_le(v160, v152) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v159 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v194 = v158
	goto L44
L44:
	;
	goto L51
L45:
	;
	v163 = v40
	goto L47
L46:
	;
	v163 = base.F64_add(v109, v137)
	goto L47
L47:
	;
	v167 = v152
	v168 = v163
	v172 = v158
	goto L48
L48:
	;
	v180 = base.F64_mul(v172, base.F64_div(v167, v168))
	v181 = float64(-1)
	v184 = base.F64_add(v167, v181)
	if base.F64_ge(v184, v160) != 0 {
		v167 = v184
		v168 = base.F64_add(v168, v181)
		v172 = v180
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v194 = v180
	goto L44
L50:
	;
	goto L49
L51:
	;
	v216 = F_pg_prng_double(m, v114)
	mBase = m.M
	if base.F64_eq(v216, float64(0)) != 0 {
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v219 = F_log(m, v194)
	mBase = m.M
	v221 = F_exp(m, base.F64_div(v219, v53))
	mBase = m.M
	v222 = F_log(m, v216)
	mBase = m.M
	v225 = F_exp(m, base.F64_div(base.F64_neg(v222), v53))
	mBase = m.M
	if base.F64_le(v221, base.F64_div(v142, v40)) == int32(0) {
		v119 = v225
		goto L32
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v233 = v225
	goto L35
L55:
	;
	goto L59
L56:
	;
	v279 = v262
	goto L57
L57:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5)+144)) = base.F64_add(v279, float64(-1))
	goto L12
L58:
	;
	F_VectorArraySet(m, v16, base.I32_trunc_sat_f64_s(base.F64_mul(v270, base.F64_convert_i32_s(v17))), v19)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L62
	}
L59:
	;
	v270 = F_pg_prng_double(m, l5+int32(120))
	mBase = m.M
	if base.F64_eq(v270, float64(0)) != 0 {
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
	v278 = *(*float64)(unsafe.Add(mBase, uint32(l5)+144))
	v279 = v278
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
	*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = l0
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
func F_StrategyFreeBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyFreeBuffer[0]))
	v8 = base.AtomicRmwXchg32(m, v5, int32(0), int32(1))
	if v8 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyFreeBuffer[0]))
		F_s_lock(m, v10, int32(_a_F_StrategyFreeBuffer_0), int32(365), int32(_a_F_StrategyFreeBuffer_1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyFreeBuffer[0]))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v18 == int32(-2) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v21
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v21 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v23
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v23
			} else {
			}
			v30 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v17))), uint32(v30))
			return
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyFreeBuffer[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		if v18 == int32(-2) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v21 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v23
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v23
		} else {
		}
		v30 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v17))), uint32(v30))
		return
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var __phi55 int32
	_ = __phi55
	var v56 int32
	_ = v56
	var __phi56 int32
	_ = __phi56
	var v58 int32
	_ = v58
	var __phi58 int32
	_ = __phi58
	var v59 int32
	_ = v59
	var __phi59 int32
	_ = __phi59
	var v60 int32
	_ = v60
	var __phi60 int32
	_ = __phi60
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if base.Ui32(v12-int32(97)) < base.Ui32(int32(26)) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v10 = l0
	v12 = v9
	goto L5
L3:
	;
	goto L4
L4:
	;
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v37)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37
	return
L5:
	;
	if base.Ui32(int32(229)) < base.Ui32((v12|int32(32)-int32(123))&int32(255)) {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L4
L7:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v26 != 0 {
		v10 = v10 + int32(1)
		v12 = v26
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v47)
	v49 = int32(1)
	v51 = l1 + v49
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v52 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v47 = v12 & int32(95)
	goto L12
L11:
	;
	v47 = v12
	goto L12
L12:
	;
	goto L9
L13:
	;
	v165 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v165)
	return
L14:
	;
	__phi55 = v10
	__phi56 = v52
	__phi58 = v51
	__phi59 = v49
	__phi60 = v10 + int32(1)
	v55 = __phi55
	v56 = __phi56
	v58 = __phi58
	v59 = __phi59
	v60 = __phi60
	goto L17
L15:
	;
	v147 = v51
	v148 = v49
	goto L16
L16:
	;
	v153 = int32(4) - v148
	if v153 != 0 {
		goto L46
	} else {
		goto L47
	}
L17:
	;
	if base.Ui32(int32(25)) < base.Ui32((v56|int32(32)-int32(97))&int32(255)) {
		v130 = v58
		v131 = v59
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if int32(3) < v131 {
		v160 = v130
		goto L13
	} else {
		goto L45
	}
L19:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	if v135 != 0 {
		goto L41
	} else {
		goto L42
	}
L20:
	;
	v72 = v56 & int32(255)
	if base.Ui32(v72-int32(97)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if base.Ui32(v93-int32(97)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v80 = base.I32_extend8_s(v79)
	v84 = base.B2i32(base.Ui32(int32(25)) < base.Ui32(v80-int32(65)))
	if v84 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v79 = v72 & int32(95)
	goto L25
L24:
	;
	v79 = v72
	goto L25
L25:
	;
	goto L22
L26:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+uint32(_c_F__soundex[0]))))
	v90 = v89
	goto L21
L27:
	;
	goto L28
L28:
	;
	v90 = v79
	goto L21
L29:
	;
	v101 = base.I32_extend8_s(v100)
	if base.Ui32(v101-int32(65)) <= base.Ui32(int32(25)) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v100 = v93 & int32(95)
	goto L32
L31:
	;
	v100 = v93
	goto L32
L32:
	;
	goto L29
L33:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F__soundex[0]))))
	v109 = v108
	goto L35
L34:
	;
	v109 = v100
	goto L35
L35:
	;
	if v90&int32(255) == v109&int32(255) {
		v130 = v58
		v131 = v59
		goto L19
	} else {
		goto L36
	}
L36:
	;
	if v84 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+uint32(_c_F__soundex[0]))))
	v118 = v117
	goto L39
L38:
	;
	v118 = v79
	goto L39
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v118)
	if v118&int32(255) == int32(48) {
		v130 = v58
		v131 = v59
		goto L19
	} else {
		goto L40
	}
L40:
	;
	v124 = int32(1)
	v130 = v58 + v124
	v131 = v59 + v124
	goto L19
L41:
	;
	if v131 < int32(4) {
		__phi55 = v60
		__phi56 = v135
		__phi58 = v130
		__phi59 = v131
		__phi60 = v60 + int32(1)
		v55 = __phi55
		v56 = __phi56
		v58 = __phi58
		v59 = __phi59
		v60 = __phi60
		goto L17
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	goto L18
L44:
	;
	goto L43
L45:
	;
	v147 = v130
	v148 = v131
	goto L16
L46:
	;
	base.MemoryFill(m, v147, int32(48), v153)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v160 = v153 + v147
	goto L13
}
func F_s_lock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v68 float64
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	v5 = int32(0)
	v11 = base.AtomicRmwXchg32(m, l0, v5, int32(1))
	if v11 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_s_lock[0])) = v117
	goto L1
L3:
	;
	if int32(999) < v99 {
		goto L1
	} else {
		goto L26
	}
L4:
	;
	F_s_lock_stuck(m, l1, l2, l3)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L24
	} else {
		goto L25
	}
L5:
	;
	v16 = v5
	v17 = v5
	v18 = v5
	goto L8
L6:
	;
	goto L7
L7:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[0]))
	v99 = v92
	goto L3
L8:
	;
	v20 = v16 + int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[0]))
	if v22 <= v20 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[0]))
	if v78 == int32(0) {
		v99 = v84
		goto L3
	} else {
		goto L22
	}
L10:
	;
	v25 = v18 + int32(1)
	if int32(1001) <= v25 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	v77 = v20
	v78 = v17
	v79 = v18
	goto L12
L12:
	;
	v82 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v82 != 0 {
		v16 = v77
		v17 = v78
		v18 = v79
		goto L8
	} else {
		goto L21
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(150994950)
	if v17 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v34 = v17
	goto L16
L15:
	;
	v34 = int32(1000)
	goto L16
L16:
	;
	F_pg_usleep(m, v34)
	mBase = m.M
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(0)
	v42 = int32(_a_F_s_lock_0)
	v45 = *(*int64)(unsafe.Add(mBase, _c_F_s_lock[2]))
	v46 = *(*int64)(unsafe.Add(mBase, _c_F_s_lock[3]))
	v47 = v45 ^ v46
	*(*int64)(unsafe.Add(mBase, _c_F_s_lock[3])) = base.I64_rotl(v47, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_s_lock[2])) = v47<<(uint(int64(16))%64) ^ base.I64_rotl(v45, int64(24)) ^ v47
	v68 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v45*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L17
L17:
	;
	v73 = v34 + base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v34), v68), float64(0.5)))
	if int32(_a_F_s_lock_1) < v73 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v76 = int32(1000)
	goto L20
L19:
	;
	v76 = v73
	goto L20
L20:
	;
	v77 = int32(0)
	v78 = v76
	v79 = v25
	goto L12
L21:
	;
	goto L9
L22:
	;
	if v84 < int32(11) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v117 = v84 - int32(1)
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
	v104 = int32(900)
	if v104 <= v99 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v107 = v104
	goto L29
L28:
	;
	v107 = v99
	goto L29
L29:
	;
	v117 = v107 + int32(100)
	goto L2
}
func F_satisfies_hash_partition(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int64
	_ = v221
	var v223 int32
	_ = v223
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v233 int32
	_ = v233
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v274 int32
	_ = v274
	var v286 int64
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int64
	_ = v303
	var v313 int64
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v373 int64
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int64
	_ = v387
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int64
	_ = v399
	var v401 int32
	_ = v401
	var v415 int64
	_ = v415
	var v418 int64
	_ = v418
	var v422 int32
	_ = v422
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
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
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	v2 = int32(0)
	v13 = int64(0)
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	v19 = F_Int64GetDatum(m, int64(8816678312871386365))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v23 != 0 {
		v422 = v2
		goto L10
	} else {
		goto L11
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L120
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L114
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L110
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L105
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L101
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L97
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L93
	}
L10:
	;
	m.G0 = v16 + int32(96)
	return v422
L11:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v24 != 0 {
		v422 = v2
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v25 != 0 {
		v422 = v2
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v26 <= int32(0) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v29 < int32(0) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if base.Ui32(v26) <= base.Ui32(v29) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v35 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v249)+8))
	if v261 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 == v33 {
		v249 = v35
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v39 = F_relation_open(m, v33, int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v41 = F_RelationGetPartitionKey(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v41 == int32(0) {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v45 != int32(104) {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = int32(0)
	if v48 == v49 {
		v60 = v49
		goto L28
	} else {
		goto L29
	}
L26:
	;
	F_relation_close(m, v39, int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L62
	}
L27:
	;
	if v60&int32(1) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	goto L27
L29:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v52 == int32(0) {
		v60 = v49
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v55 != int32(15) {
		v60 = v49
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+13)))
	v60 = v58
	goto L28
L32:
	;
	v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	v67 = v65 - int32(3)
	v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	if v67 != v68 {
		goto L5
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v149 = F_pg_detoast_datum(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L51
	}
L35:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v76 = F_MemoryContextAllocZero(m, v71, v67*int32(28)+int32(144))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v76
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v33
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v83
	v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	v87 = v85 << (uint(int32(2)) % 32)
	if v87 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	base.MemoryCopy(m, v81+int32(16), v90, v87)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	if v92 <= int32(0) {
		v233 = v81
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v100 = int32(0)
	goto L41
L41:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v114 = F_get_fn_expr_argtype(m, v111, v100+int32(3))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	v233 = v81
	goto L26
L43:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v41)+32))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+v100<<(uint(int32(2))%32))))
	if v114 != v120 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v122 = F_IsBinaryCoercible(m, v114, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v127 = v100 * int32(28)
	v128 = v81 + int32(144) + v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	v130 = v129 + v127
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v130)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+16)) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+24)) = v135
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v130)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+8)) = v137
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
	*(*int64)(unsafe.Add(mBase, uint32(v128))) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v128)+20)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v128)+16)) = int32(0)
	goto L49
L47:
	;
	if v122 == int32(0) {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v145 = v100 + int32(1)
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	if v145 < v146 {
		v100 = v145
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+20))
	v154 = F_MemoryContextAllocZero(m, v152, int32(172))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v156)+16)) = v154
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v33
	v161 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+8)) = v163
	F_get_typlenbyvalalign(m, v163, v159+int32(12), v159+int32(14), v159+int32(15))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+16)) = v174
	v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	if int32(0) < v176 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v41)+32))
	v184 = int32(0)
	goto L57
L55:
	;
	goto L56
L56:
	;
	v217 = v159 + int32(144)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v218)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v217)+16)) = v221
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v218)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+24)) = v223
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v218)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v217)+8)) = v225
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v218)))
	*(*int64)(unsafe.Add(mBase, uint32(v217))) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v217)+20)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v217)+16)) = int32(0)
	goto L61
L57:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v180+v184<<(uint(int32(2))%32))))
	if v198 != v179 {
		goto L3
	} else {
		goto L59
	}
L58:
	;
	goto L56
L59:
	;
	v201 = v184 + int32(1)
	if v201 != v176 {
		v184 = v201
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v233 = v159
	goto L26
L62:
	;
	v249 = v233
	goto L17
L63:
	;
	v418 = base.I64_rem_u_s(v415, base.I64_extend_i32_u(v26))
	v422 = base.B2i32(base.I64_extend_i32_u(v29) == v418)
	goto L10
L64:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v264 <= int32(0) {
		v415 = v13
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v318 = F_pg_detoast_datum(m, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L75
	}
L67:
	;
	v274 = int32(0)
	v286 = v13
	goto L68
L68:
	;
	v289 = l0 + int32(20) + v274<<(uint(int32(3))%32)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+28)))
	if v290 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v415 = v313
	goto L63
L70:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v249+int32(16)+v274<<(uint(int32(2))%32))))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v289)+24))
	v301 = F_FunctionCall2Coll(m, v249+int32(144)+v274*int32(28), v299, v300, v19)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	v313 = v286
	goto L72
L72:
	;
	v315 = v274 + int32(1)
	if v315 != v264 {
		v274 = v315
		v286 = v313
		goto L68
	} else {
		goto L74
	}
L73:
	;
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v301)))
	v313 = v303 + (v286<<(uint(int64(54))%64) + int64(base.Ui64(v286)>>(uint(int64(7))%64))) + int64(5305509591434766563) ^ v286
	goto L72
L74:
	;
	goto L69
L75:
	;
	v321 = int32(*(*int16)(unsafe.Add(mBase, uint32(v249)+12)))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+14)))
	v323 = int32(*(*int8)(unsafe.Add(mBase, uint32(v249)+15)))
	F_deconstruct_array(m, v318, v321, v322, v323, v16+int32(88), v16+int32(84), v16+int32(92))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v332 == v333 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v361 = int32(0)
	v363 = v332
	v373 = v13
	goto L86
L78:
	;
	if int32(0) < v332 {
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L82
	}
L81:
	;
	v415 = v13
	goto L63
L82:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v344
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v346
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_0), v16+int32(16))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_2), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374+v361))))
	if v376 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v415 = v399
	goto L63
L88:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v249)+16))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v380+v361<<(uint(int32(2))%32))))
	v385 = F_FunctionCall2Coll(m, v249+int32(144), v379, v384, v19)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	v398 = v363
	v399 = v373
	goto L90
L90:
	;
	v401 = v361 + int32(1)
	if v401 < v398 {
		v361 = v401
		v363 = v398
		v373 = v399
		goto L86
	} else {
		goto L92
	}
L91:
	;
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v385)))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v398 = v397
	v399 = v387 + (v373<<(uint(int64(54))%64) + int64(base.Ui64(v373)>>(uint(int64(7))%64))) + int64(5305509591434766563) ^ v373
	goto L90
L92:
	;
	goto L87
L93:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_4), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_5), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_6), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_7), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_8), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_9), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v492 = F_get_rel_name(m, v33)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v492
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_10), v16)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_11), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v510 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v510
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_0), v16-int32(-64))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_12), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v41)+32))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v530+v100<<(uint(int32(2))%32))))
	v535 = F_format_type_be(m, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v537 = F_format_type_be(m, v114)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v535
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v100 + int32(1)
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_13), v16+int32(48))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_14), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v41)+32))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v561+v184<<(uint(int32(2))%32))))
	v566 = F_format_type_be(m, v565)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	v569 = F_format_type_be(m, v568)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v569
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v184 + int32(1)
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_15), v16+int32(32))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_16), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 float64
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 float64
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 float64
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
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
	var v222 int64
	_ = v222
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 float64
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int64
	_ = v242
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 float64
	_ = v260
	var v261 int32
	_ = v261
	var v264 float64
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v280 float64
	_ = v280
	var v281 int32
	_ = v281
	var v282 float64
	_ = v282
	var v286 float32
	_ = v286
	var v292 float64
	_ = v292
	var v297 float64
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 float64
	_ = v305
	var v313 float64
	_ = v313
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v389 float64
	_ = v389
	var v390 int32
	_ = v390
	var v401 float64
	_ = v401
	var v402 float64
	_ = v402
	var v406 int32
	_ = v406
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 float64
	_ = v452
	var v459 float64
	_ = v459
	var v460 float64
	_ = v460
	var v466 float64
	_ = v466
	var v467 float64
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v478 float64
	_ = v478
	var v479 float64
	_ = v479
	var v501 float64
	_ = v501
	var v504 float64
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 float64
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v532 float64
	_ = v532
	var v533 float64
	_ = v533
	var v537 int32
	_ = v537
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 float64
	_ = v574
	var v581 float64
	_ = v581
	var v582 float64
	_ = v582
	var v588 float64
	_ = v588
	var v589 float64
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v600 float64
	_ = v600
	var v601 float64
	_ = v601
	var v621 float64
	_ = v621
	var v624 float64
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 float64
	_ = v657
	var v675 float64
	_ = v675
	var v678 float64
	_ = v678
	var v681 float64
	_ = v681
	var v684 float64
	_ = v684
	var v687 float64
	_ = v687
	var v690 float64
	_ = v690
	var v693 float64
	_ = v693
	var v696 float64
	_ = v696
	var v699 float64
	_ = v699
	var v709 float64
	_ = v709
	var v726 float64
	_ = v726
	var v741 float64
	_ = v741
	v7 = float64(0)
	v10 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(96)
	m.G0 = v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v35 = F_estimate_expression_value(m, l0, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return float64(0)
L2:
	;
	v39 = F_estimate_expression_value(m, l0, v32)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v26 + int32(96)
	return v741
L4:
	;
	v41 = F_exprType(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v43 = F_get_base_element_type(m, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v43 == int32(0) {
		v741 = float64(0.5)
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v47 = F_exprCollation(m, v39)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v49 = int32(0)
	if v39 == v49 {
		v100 = v49
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v114 = F_lookup_type_cache(m, v43, int32(1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L19
	}
L10:
	;
	v63 = v39
	goto L11
L11:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	switch v75 - int32(27) {
	case 0:
		goto L13
	default:
		v100 = v63
		goto L9
	case 2:
		goto L14
	}
L12:
	;
	v100 = int32(0)
	goto L9
L13:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v87 != 0 {
		v63 = v87
		goto L11
	} else {
		goto L17
	}
L14:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 != int32(27) {
		v100 = v63
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v83 != int32(34) {
		v100 = v63
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
	v129 = int32(0)
	if l2|base.B2i32(v127|v126 == v129) == v129 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+52))
	if v116 == int32(0) {
		v126 = v10
		v127 = int32(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if v116 == v29 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v126 = int32(1)
	v127 = int32(0)
	goto L18
L22:
	;
	goto L23
L23:
	;
	v122 = F_get_negator(m, v29)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v114)+52))
	v126 = v10
	v127 = base.B2i32(v122 == v124)
	goto L18
L25:
	;
	v136 = m.G0
	v138 = v136 - int32(112)
	m.G0 = v138
	F_examine_variable(m, l0, v100, l3, v138+int32(80))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
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
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)+84))
	if v144 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	m.G0 = v138 + int32(112)
	if base.F64_ge(v313, float64(0)) != 0 {
		v741 = v313
		goto L3
	} else {
		goto L87
	}
L30:
	;
	v147 = float64(-1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v138)+88))
	if v148 == int32(0) {
		v313 = v147
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v154 != int32(7) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v138)+92))
	m.T0[v151].(func(*base.Module, int32))(m, v148)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v313 = v147
	goto L29
L35:
	;
	v157 = float64(-1)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v138)+88))
	if v158 == int32(0) {
		v313 = v157
		goto L29
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)))
	if v164 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v138)+92))
	m.T0[v161].(func(*base.Module, int32))(m, v158)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v313 = v157
	goto L29
L40:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v138)+88))
	if v167 == int32(0) {
		v313 = v7
		goto L29
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+76)) = v173
	v176 = F_lookup_type_cache(m, v43, int32(64))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v138)+92))
	m.T0[v170].(func(*base.Module, int32))(m, v167)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v313 = v7
	goto L29
L45:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v176)+108))
	if v178 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v181 = float64(-1)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v138)+88))
	if v182 == int32(0) {
		v313 = v181
		goto L29
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v188 = v28&int32(1) ^ v126
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v138)+88))
	if v189 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v138)+92))
	m.T0[v185].(func(*base.Module, int32))(m, v182)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v313 = v181
	goto L29
L51:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v138)+88))
	if v298 != 0 {
		goto L78
	} else {
		goto L79
	}
L52:
	;
	if v188 != 0 {
		goto L75
	} else {
		goto L76
	}
L53:
	;
	v194 = F_statistic_proc_security_check(m, v138+int32(80), v178)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v194 == int32(0) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v138)+88))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+16))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+22)))
	v207 = F_get_attstatsslot(m, v138+int32(40), v198, int32(4), int32(0), int32(3))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	v286 = *(*float32)(unsafe.Add(mBase, uint32(v199+v200)+8))
	v297 = base.F64_mul(v282, base.F64_sub(float64(1), base.F64_promote_f32(v286)))
	goto L51
L57:
	;
	if v207 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v188 != 0 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	goto L60
L60:
	;
	if v188 != 0 {
		v282 = float64(0.005)
		goto L56
	} else {
		goto L73
	}
L61:
	;
	F_free_attstatsslot(m, v138)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L71
	}
L62:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v138)+52))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v138)+56))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v138)+60))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v138)+64))
	v260 = F_mcelem_array_contained_selec(m, v253, v254, v255, v256, v138+int32(76), int32(1), v251, v252, v176)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L70
	}
L63:
	;
	v240 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v138)+32)) = v240
	v242 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v138)+24)) = v242
	*(*int64)(unsafe.Add(mBase, uint32(v138)+16)) = v242
	*(*int64)(unsafe.Add(mBase, uint32(v138)+8)) = v242
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v242
	v251 = v209
	v252 = v240
	goto L62
L64:
	;
	v209 = int32(0)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v138)+88))
	v214 = F_get_attstatsslot(m, v138, v210, int32(5), v209, int32(2))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+32)) = int32(0)
	v222 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v138)+24)) = v222
	*(*int64)(unsafe.Add(mBase, uint32(v138)+16)) = v222
	*(*int64)(unsafe.Add(mBase, uint32(v138)+8)) = v222
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v222
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v138)+52))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v138)+56))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v138)+60))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v138)+64))
	v238 = F_mcelem_array_contain_overlap_selec(m, v230, v231, v232, v233, v138+int32(76), int32(1), int32(2751), v176)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	if v214 == int32(0) {
		goto L63
	} else {
		goto L68
	}
L68:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v138)+24))
	v251 = v218
	v252 = v219
	goto L62
L69:
	;
	v264 = v238
	goto L61
L70:
	;
	v264 = v260
	goto L61
L71:
	;
	F_free_attstatsslot(m, v138+int32(40))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v282 = v264
	goto L56
L73:
	;
	v272 = int32(0)
	v280 = F_mcelem_array_contain_overlap_selec(m, v272, v272, v272, v272, v138+int32(76), int32(1), int32(2751), v176)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v282 = v280
	goto L56
L75:
	;
	v292 = float64(0.005)
	goto L77
L76:
	;
	v292 = float64(0.004999999888241291)
	goto L77
L77:
	;
	v297 = v292
	goto L51
L78:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v138)+92))
	m.T0[v299].(func(*base.Module, int32))(m, v298)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v126 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L80
L82:
	;
	v305 = v297
	goto L84
L83:
	;
	v305 = base.F64_sub(float64(1), v297)
	goto L84
L84:
	;
	if base.F64_lt(v305, float64(0)) != 0 {
		v313 = float64(0)
		goto L29
	} else {
		goto L85
	}
L85:
	;
	if base.F64_gt(v305, float64(1)) == int32(0) {
		v313 = v305
		goto L29
	} else {
		goto L86
	}
L86:
	;
	v313 = float64(1)
	goto L29
L87:
	;
	goto L27
L88:
	;
	if v335 == int32(0) {
		v741 = float64(0.5)
		goto L3
	} else {
		goto L94
	}
L89:
	;
	v331 = F_get_oprjoin(m, v29)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v333 = F_get_oprrest(m, v29)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	v335 = v331
	goto L88
L93:
	;
	v335 = v333
	goto L88
L94:
	;
	F_fmgr_info(m, v335, v26+int32(68))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	switch v335 - int32(101) {
	case 0, 4:
		v348 = int32(1)
		v349 = v127
		goto L96
	case 1, 5:
		goto L98
	default:
		v347 = v127
		goto L97
	}
L96:
	;
	if v100 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v348 = v126
	v349 = v347
	goto L96
L98:
	;
	v347 = int32(1)
	goto L97
L99:
	;
	v726 = float64(0)
	if base.F64_lt(v709, v726) != 0 {
		v741 = v726
		goto L3
	} else {
		goto L203
	}
L100:
	;
	v627 = F_palloc0(m, int32(16))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L192
	}
L101:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v352 != int32(35) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	if v352 != int32(7) {
		goto L100
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+20)))
	if v505 != 0 {
		goto L100
	} else {
		goto L150
	}
L105:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+24)))
	if v357 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v741 = float64(0)
	goto L3
L107:
	;
	goto L108
L108:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	v362 = F_pg_detoast_datum(m, v361)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v362)+12))
	F_get_typlenbyvalalign(m, v364, v26+int32(66), v26+int32(65), v26-int32(-64))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v374 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+66)))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+65)))
	v376 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26)+64)))
	F_deconstruct_array(m, v362, v374, v375, v376, v26+int32(56), v26+int32(52), v26+int32(60))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v28&int32(1) != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v389 = float64(0)
	goto L114
L113:
	;
	v389 = float64(1)
	goto L114
L114:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v26)+60))
	if v390 <= int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v28&int32(1) != 0 {
		goto L139
	} else {
		goto L140
	}
L116:
	;
	v478 = v389
	v479 = v389
	goto L115
L117:
	;
	goto L118
L118:
	;
	v401 = v389
	v402 = v389
	v406 = int32(0)
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v35
	v420 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+66)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v421+v406<<(uint(int32(2))%32))))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426+v406))))
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+65)))
	v430 = F_makeConst(m, v43, int32(-1), v47, v420, v425, v428, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L121
	}
L120:
	;
	v478 = v466
	v479 = v467
	goto L115
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v430
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v430
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v434
	v440 = F_list_make2_impl(m, v26+int32(16), v26+int32(12))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if l2 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v452 = *(*float64)(unsafe.Add(mBase, uint32(v451)))
	if v28&int32(1) != 0 {
		goto L130
	} else {
		goto L131
	}
L124:
	;
	v445 = F_FunctionCall5Coll(m, v26+int32(68), v442, l0, v29, v440, base.I32_extend16_s(l4), l5)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v449 = F_FunctionCall4Coll(m, v26+int32(68), v442, l0, v29, v440, l3)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L128
	}
L127:
	;
	v451 = v445
	goto L123
L128:
	;
	v451 = v449
	goto L123
L129:
	;
	v469 = v406 + int32(1)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v26)+60))
	if v469 < v470 {
		v401 = v466
		v402 = v467
		v406 = v469
		goto L119
	} else {
		goto L137
	}
L130:
	;
	if v348 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	goto L132
L132:
	;
	v460 = base.F64_mul(v401, v452)
	if v349 == int32(0) {
		v466 = v460
		v467 = v402
		goto L129
	} else {
		goto L136
	}
L133:
	;
	v459 = base.F64_add(v402, v452)
	goto L135
L134:
	;
	v459 = v402
	goto L135
L135:
	;
	v466 = base.F64_sub(base.F64_add(v401, v452), base.F64_mul(v401, v452))
	v467 = v459
	goto L129
L136:
	;
	v466 = v460
	v467 = base.F64_add(v402, base.F64_add(v452, float64(-1)))
	goto L129
L137:
	;
	goto L120
L138:
	;
	if base.F64_le(v479, float64(1)) != 0 {
		goto L144
	} else {
		goto L145
	}
L139:
	;
	if v348 != 0 {
		goto L138
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	if v349 == int32(0) {
		v709 = v478
		goto L99
	} else {
		goto L143
	}
L142:
	;
	v709 = v478
	goto L99
L143:
	;
	goto L138
L144:
	;
	v501 = v479
	goto L146
L145:
	;
	v501 = v478
	goto L146
L146:
	;
	if base.F64_ge(v479, float64(0)) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v504 = v501
	goto L149
L148:
	;
	v504 = v478
	goto L149
L149:
	;
	v709 = v504
	goto L99
L150:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	F_get_typlenbyval(m, v506, v26+int32(60), v26+int32(56))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v516 = v28 & int32(1)
	if v516 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v517 = float64(0)
	goto L154
L153:
	;
	v517 = float64(1)
	goto L154
L154:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	if v518 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	if v516 != 0 {
		goto L181
	} else {
		goto L182
	}
L156:
	;
	v600 = v517
	v601 = v517
	goto L155
L157:
	;
	goto L158
L158:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v521 <= int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v600 = v517
	v601 = v517
	goto L155
L160:
	;
	goto L161
L161:
	;
	v532 = v517
	v533 = v517
	v537 = int32(0)
	goto L162
L162:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v549+v537<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v553
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v553
	v562 = F_list_make2_impl(m, v26+int32(24), v26+int32(20))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L164
	}
L163:
	;
	v600 = v588
	v601 = v589
	goto L155
L164:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if l2 != 0 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v574 = *(*float64)(unsafe.Add(mBase, uint32(v573)))
	if v28&int32(1) != 0 {
		goto L172
	} else {
		goto L173
	}
L166:
	;
	v567 = F_FunctionCall5Coll(m, v26+int32(68), v564, l0, v29, v562, base.I32_extend16_s(l4), l5)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v571 = F_FunctionCall4Coll(m, v26+int32(68), v564, l0, v29, v562, l3)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L170
	}
L169:
	;
	v573 = v567
	goto L165
L170:
	;
	v573 = v571
	goto L165
L171:
	;
	v591 = v537 + int32(1)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v591 < v592 {
		v532 = v588
		v533 = v589
		v537 = v591
		goto L162
	} else {
		goto L179
	}
L172:
	;
	if v348 != 0 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	goto L174
L174:
	;
	v582 = base.F64_mul(v532, v574)
	if v349 == int32(0) {
		v588 = v582
		v589 = v533
		goto L171
	} else {
		goto L178
	}
L175:
	;
	v581 = base.F64_add(v533, v574)
	goto L177
L176:
	;
	v581 = v533
	goto L177
L177:
	;
	v588 = base.F64_sub(base.F64_add(v532, v574), base.F64_mul(v532, v574))
	v589 = v581
	goto L171
L178:
	;
	v588 = v582
	v589 = base.F64_add(v533, base.F64_add(v574, float64(-1)))
	goto L171
L179:
	;
	goto L163
L180:
	;
	if base.F64_le(v601, float64(1)) != 0 {
		goto L186
	} else {
		goto L187
	}
L181:
	;
	if v348 != 0 {
		goto L180
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	if v349 == int32(0) {
		v709 = v600
		goto L99
	} else {
		goto L185
	}
L184:
	;
	v709 = v600
	goto L99
L185:
	;
	goto L180
L186:
	;
	v621 = v601
	goto L188
L187:
	;
	v621 = v600
	goto L188
L188:
	;
	if base.F64_ge(v601, float64(0)) != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v624 = v621
	goto L191
L190:
	;
	v624 = v600
	goto L191
L191:
	;
	v709 = v624
	goto L99
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v627)+8)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v627)+4)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v627))) = int32(34)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v627)+12)) = v634
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v627
	v644 = F_list_make2_impl(m, v26+int32(8), v26+int32(4))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if l2 != 0 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v657 = *(*float64)(unsafe.Add(mBase, uint32(v656)))
	if v28&int32(1) == int32(0) {
		goto L200
	} else {
		goto L201
	}
L195:
	;
	v650 = F_FunctionCall5Coll(m, v26+int32(68), v646, l0, v29, v644, base.I32_extend16_s(l4), l5)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v654 = F_FunctionCall4Coll(m, v26+int32(68), v646, l0, v29, v644, l3)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L199
	}
L198:
	;
	v656 = v650
	goto L194
L199:
	;
	v656 = v654
	goto L194
L200:
	;
	v709 = base.F64_mul(v657, base.F64_mul(v657, base.F64_mul(v657, base.F64_mul(v657, base.F64_mul(v657, base.F64_mul(v657, base.F64_mul(v657, base.F64_mul(v657, base.F64_mul(v657, v657)))))))))
	goto L99
L201:
	;
	goto L202
L202:
	;
	v675 = base.F64_add(base.F64_mul(v657, math.Float64frombits(uint64(0x8000000000000000))), base.F64_add(v657, float64(0)))
	v678 = base.F64_sub(base.F64_add(v657, v675), base.F64_mul(v675, v657))
	v681 = base.F64_sub(base.F64_add(v657, v678), base.F64_mul(v678, v657))
	v684 = base.F64_sub(base.F64_add(v657, v681), base.F64_mul(v681, v657))
	v687 = base.F64_sub(base.F64_add(v657, v684), base.F64_mul(v684, v657))
	v690 = base.F64_sub(base.F64_add(v657, v687), base.F64_mul(v687, v657))
	v693 = base.F64_sub(base.F64_add(v657, v690), base.F64_mul(v690, v657))
	v696 = base.F64_sub(base.F64_add(v657, v693), base.F64_mul(v693, v657))
	v699 = base.F64_sub(base.F64_add(v657, v696), base.F64_mul(v696, v657))
	v709 = base.F64_sub(base.F64_add(v657, v699), base.F64_mul(v699, v657))
	goto L99
L203:
	;
	if base.F64_gt(v709, float64(1)) == int32(0) {
		v741 = v709
		goto L3
	} else {
		goto L204
	}
L204:
	;
	v741 = float64(1)
	goto L3
}
func F_scalargtsel(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_scalarineqsel_wrapper(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
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
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v275 int32
	_ = v275
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = F_palloc0(m, int32(200))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0
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
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L67
	}
L4:
	;
	F_pg_cryptohash_free(m, v104)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
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
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)) = uint8(v47)
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
	v287 = m.ExcPending
	if v287 != 0 {
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
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+364))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(137438953475)
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
	v67 = F_parse_scram_secret(m, l2, v11+int32(24), v11+int32(16), v11+int32(20), v11+int32(28), v11-int32(-64), v11+int32(96))
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
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
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
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
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
	F_errfinish(m, int32(_a_F_scram_init_3), int32(292), int32(_a_F_scram_init_4))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v92
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
	v114 = F_pg_cryptohash_update(m, v104, v100+int32(257), int32(32))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(_a_F_scram_init_7)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v135
	v258 = v11 - int32(-64)
	v259 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v258)+56)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v258)+48)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v258)+40)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v258)+32)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v258)+24)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v258)+16)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v258)+8)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v258))) = v259
	v275 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+192)) = uint8(v275)
	goto L15
L59:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(_a_F_scram_init_8), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_scram_init_3), int32(267), int32(_a_F_scram_init_4))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
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
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errmsg_internal(m, int32(_a_F_scram_init_9), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_scram_init_3), int32(719), int32(_a_F_scram_init_10))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_scram_init_9), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_scram_init_3), int32(728), int32(_a_F_scram_init_10))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
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
		v65 = v16
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	F_ProcessClientWriteInterrupt(m, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L5:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v19 != 0 {
		v65 = v16
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
		v65 = v23
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v65 = v56
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
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = F_pgmem_send(m, v55, l1, l2)
	mBase = m.M
	if int32(0) <= v56 {
		v65 = v56
		goto L4
	} else {
		goto L18
	}
L16:
	;
	F_ProcessClientWriteInterrupt(m, int32(1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v59 == int32(0) {
		v23 = v56
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
	return v65
L21:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(_a_F_secure_write_0), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_secure_write_1), int32(350), int32(_a_F_secure_write_2))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v264 int64
	_ = v264
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
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
	var v319 int32
	_ = v319
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
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int64
	_ = v363
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v630 int32
	_ = v630
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v826 int32
	_ = v826
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v899 int32
	_ = v899
	var v906 int32
	_ = v906
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v981 int32
	_ = v981
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1059 int32
	_ = v1059
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1121 int32
	_ = v1121
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1198 int32
	_ = v1198
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1229 int32
	_ = v1229
	var v1267 int32
	_ = v1267
	var v1296 int32
	_ = v1296
	var v1347 int32
	_ = v1347
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1383 int32
	_ = v1383
	var v1388 int32
	_ = v1388
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int64
	_ = v1445
	var v1448 int64
	_ = v1448
	var v1454 int64
	_ = v1454
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1554 int32
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1612 int32
	_ = v1612
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1629 int32
	_ = v1629
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1670 int64
	_ = v1670
	var v1671 int64
	_ = v1671
	var v1677 int32
	_ = v1677
	var v1683 int32
	_ = v1683
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1719 int32
	_ = v1719
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1734 int32
	_ = v1734
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1746 int32
	_ = v1746
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1777 int32
	_ = v1777
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1803 int32
	_ = v1803
	var v1808 int32
	_ = v1808
	var v1819 int32
	_ = v1819
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1842 int32
	_ = v1842
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1914 int32
	_ = v1914
	var v1930 int32
	_ = v1930
	var v1963 int64
	_ = v1963
	var v1966 int64
	_ = v1966
	var v1969 int64
	_ = v1969
	var v1980 int64
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1983 int32
	_ = v1983
	var v2002 int32
	_ = v2002
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
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2036 int32
	_ = v2036
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2073 int32
	_ = v2073
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2099 int32
	_ = v2099
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2133 int32
	_ = v2133
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2171 int32
	_ = v2171
	var v2175 int32
	_ = v2175
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2201 int32
	_ = v2201
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2268 int32
	_ = v2268
	var v2270 int32
	_ = v2270
	var v2274 int32
	_ = v2274
	var v2278 int32
	_ = v2278
	var v2282 int32
	_ = v2282
	var v2286 int32
	_ = v2286
	var v2290 int32
	_ = v2290
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2301 int32
	_ = v2301
	var v2311 int32
	_ = v2311
	var v2351 int32
	_ = v2351
	var v2354 int32
	_ = v2354
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2368 int32
	_ = v2368
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2382 int32
	_ = v2382
	var v2386 int32
	_ = v2386
	var v2434 int32
	_ = v2434
	var v2437 int32
	_ = v2437
	var v2441 int32
	_ = v2441
	var v2448 int32
	_ = v2448
	var v2450 int32
	_ = v2450
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2464 int32
	_ = v2464
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2490 int32
	_ = v2490
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2510 int32
	_ = v2510
	var v2512 int32
	_ = v2512
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2523 int32
	_ = v2523
	var v2527 int32
	_ = v2527
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2542 int32
	_ = v2542
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2553 int32
	_ = v2553
	var v2557 int32
	_ = v2557
	var v2558 int32
	_ = v2558
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2570 int32
	_ = v2570
	var v2572 int32
	_ = v2572
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2587 int32
	_ = v2587
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2605 int32
	_ = v2605
	var v2606 int32
	_ = v2606
	var v2607 int32
	_ = v2607
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2612 int32
	_ = v2612
	var v2614 int32
	_ = v2614
	var v2618 int32
	_ = v2618
	var v2619 int32
	_ = v2619
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2625 int32
	_ = v2625
	var v2629 int32
	_ = v2629
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2655 int32
	_ = v2655
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2674 int32
	_ = v2674
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2680 int32
	_ = v2680
	var v2681 int32
	_ = v2681
	var v2682 int32
	_ = v2682
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2690 int32
	_ = v2690
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2710 int32
	_ = v2710
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2717 int32
	_ = v2717
	var v2718 int32
	_ = v2718
	var v2719 int32
	_ = v2719
	var v2722 int32
	_ = v2722
	var v2724 int32
	_ = v2724
	var v2728 int32
	_ = v2728
	var v2732 int32
	_ = v2732
	var v2736 int32
	_ = v2736
	var v2740 int32
	_ = v2740
	var v2744 int32
	_ = v2744
	var v2749 int32
	_ = v2749
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2755 int32
	_ = v2755
	var v2767 int32
	_ = v2767
	var v2807 int32
	_ = v2807
	var v2810 int32
	_ = v2810
	var v2813 int32
	_ = v2813
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2824 int32
	_ = v2824
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2838 int32
	_ = v2838
	var v2842 int32
	_ = v2842
	var v2889 int32
	_ = v2889
	var v2893 int32
	_ = v2893
	var v2895 int32
	_ = v2895
	var v2898 int32
	_ = v2898
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2906 int32
	_ = v2906
	var v2908 int64
	_ = v2908
	var v2910 int32
	_ = v2910
	var v2913 int32
	_ = v2913
	var v2914 int32
	_ = v2914
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2923 int64
	_ = v2923
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2931 int64
	_ = v2931
	var v2933 int64
	_ = v2933
	var v2934 int64
	_ = v2934
	var v2941 int32
	_ = v2941
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2946 int32
	_ = v2946
	var v2957 int32
	_ = v2957
	var v2967 int32
	_ = v2967
	var v2969 int32
	_ = v2969
	var v2973 int32
	_ = v2973
	var v2974 int32
	_ = v2974
	var v2975 int32
	_ = v2975
	var v2977 int32
	_ = v2977
	var v2979 int64
	_ = v2979
	var v2981 int32
	_ = v2981
	var v2984 int32
	_ = v2984
	var v2985 int32
	_ = v2985
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v2994 int64
	_ = v2994
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v3002 int64
	_ = v3002
	var v3004 int64
	_ = v3004
	var v3005 int64
	_ = v3005
	var v3012 int32
	_ = v3012
	var v3015 int32
	_ = v3015
	var v3016 int32
	_ = v3016
	var v3017 int32
	_ = v3017
	var v3028 int32
	_ = v3028
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3049 int32
	_ = v3049
	var v3051 int32
	_ = v3051
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3060 int32
	_ = v3060
	var v3061 int32
	_ = v3061
	var v3063 int32
	_ = v3063
	var v3078 int32
	_ = v3078
	var v3081 int32
	_ = v3081
	var v3115 int32
	_ = v3115
	var v3119 int32
	_ = v3119
	var v3124 int32
	_ = v3124
	var v3127 int32
	_ = v3127
	var v3130 int32
	_ = v3130
	var v3131 int32
	_ = v3131
	var v3135 int32
	_ = v3135
	var v3151 int32
	_ = v3151
	var v3170 int32
	_ = v3170
	var v3188 int32
	_ = v3188
	var v3200 int32
	_ = v3200
	var v3203 int32
	_ = v3203
	var v3205 int32
	_ = v3205
	var v3220 int32
	_ = v3220
	var v3223 int32
	_ = v3223
	var v3259 int32
	_ = v3259
	var v3271 int32
	_ = v3271
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3290 int32
	_ = v3290
	var v3325 int32
	_ = v3325
	var v3339 int32
	_ = v3339
	var v3383 int32
	_ = v3383
	var v3384 int32
	_ = v3384
	var v3390 int32
	_ = v3390
	var v3391 int32
	_ = v3391
	var v3412 int32
	_ = v3412
	var v3419 int32
	_ = v3419
	var v3445 int32
	_ = v3445
	var v3446 int32
	_ = v3446
	var v3449 int32
	_ = v3449
	var v3452 int32
	_ = v3452
	var v3455 int32
	_ = v3455
	var v3458 int32
	_ = v3458
	var v3459 int32
	_ = v3459
	var v3461 int32
	_ = v3461
	var v3480 int32
	_ = v3480
	var v3526 int32
	_ = v3526
	var v3529 int32
	_ = v3529
	var v3559 int32
	_ = v3559
	var v3560 int32
	_ = v3560
	var v3563 int32
	_ = v3563
	var v3566 int32
	_ = v3566
	var v3616 int32
	_ = v3616
	var v3617 int32
	_ = v3617
	var v3621 int32
	_ = v3621
	var v3623 int32
	_ = v3623
	var v3626 int32
	_ = v3626
	var v3639 int32
	_ = v3639
	var v3700 int32
	_ = v3700
	var v3731 int32
	_ = v3731
	var v3734 int32
	_ = v3734
	var v3739 int32
	_ = v3739
	var v3744 int32
	_ = v3744
	var v3747 int32
	_ = v3747
	var v3749 int32
	_ = v3749
	var v3751 int32
	_ = v3751
	var v3755 int32
	_ = v3755
	var v3761 int32
	_ = v3761
	var v3771 int32
	_ = v3771
	var v3776 int32
	_ = v3776
	var v3777 int32
	_ = v3777
	var v3782 int32
	_ = v3782
	var v3783 int32
	_ = v3783
	var v3785 int32
	_ = v3785
	var v3806 int32
	_ = v3806
	var v3809 int32
	_ = v3809
	var v3839 int32
	_ = v3839
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
	var v3847 int64
	_ = v3847
	var v3861 int32
	_ = v3861
	var v3862 int32
	_ = v3862
	var v3872 int32
	_ = v3872
	var v3877 int32
	_ = v3877
	var v3922 int64
	_ = v3922
	var v3924 int32
	_ = v3924
	var v3925 int32
	_ = v3925
	var v3936 int32
	_ = v3936
	var v3945 int32
	_ = v3945
	var v3958 int32
	_ = v3958
	var v3970 int64
	_ = v3970
	var v3973 int32
	_ = v3973
	var v3975 int32
	_ = v3975
	var v3983 int32
	_ = v3983
	var v3986 int32
	_ = v3986
	var v3990 int32
	_ = v3990
	var v3994 int32
	_ = v3994
	var v3999 int32
	_ = v3999
	var v4003 int32
	_ = v4003
	var v4005 int32
	_ = v4005
	var v4013 int32
	_ = v4013
	var v4018 int32
	_ = v4018
	var v4022 int32
	_ = v4022
	var v4025 int32
	_ = v4025
	var v4033 int32
	_ = v4033
	var v4038 int32
	_ = v4038
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
	v52 = F_palloc(m, int32(_a_F_sendDir_1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v57 = int32(0)
	goto L3
L3:
	;
	v60 = l1
	v62 = int32(0)
	goto L9
L4:
	;
	return int64(0)
L5:
	;
	v57 = v52
	goto L3
L6:
	;
	v305 = F_AllocateDir(m, l1)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L4
	} else {
		goto L84
	}
L7:
	;
	v269 = int32(_a_F_sendDir_2)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v275 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[0])))
	if base.B2i32(v272 == int32(0))|base.B2i32(v272 != v275) != 0 {
		v293 = v272
		v294 = v275
		goto L72
	} else {
		goto L73
	}
L8:
	;
	if v62 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L9:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v63 != int32(47) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v60 = v60 + int32(1)
	v62 = v66
	goto L9
L12:
	;
	if v63 != 0 {
		v66 = v62
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v66 = v60
	goto L11
L15:
	;
	goto L8
L16:
	;
	v72 = v62 + int32(1)
	v73 = int32(_a_F_sendDir_3)
	v77 = m.G0
	v79 = v77 - int32(32)
	v80 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v79)+24)) = v80
	*(*int64)(unsafe.Add(mBase, uint32(v79)+16)) = v80
	*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = v80
	*(*int64)(unsafe.Add(mBase, uint32(v79))) = v80
	v88 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[1])))
	if v88 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v157 = F_strlen(m, v72)
	mBase = m.M
	if v156 != v157 {
		goto L7
	} else {
		goto L36
	}
L18:
	;
	v156 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[2])))
	if v92 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v96 = v72
	goto L24
L22:
	;
	goto L23
L23:
	;
	v106 = v73
	v107 = v88
	goto L27
L24:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v102 == v88 {
		v96 = v96 + int32(1)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v156 = v96 - v72
	goto L17
L26:
	;
	goto L25
L27:
	;
	v114 = v79 + int32(base.Ui32(v107)>>(uint(int32(3))%32))&int32(28)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v116 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v115 | v116<<(uint(v107)%32)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v120 != 0 {
		v106 = v106 + v116
		v107 = v120
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v123 == int32(0) {
		v146 = v72
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v156 = v146 - v72
	goto L17
L31:
	;
	v127 = v72
	v128 = v123
	goto L32
L32:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v79+int32(base.Ui32(v128)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v136)>>(uint(v128)%32))&int32(1) == int32(0) {
		v146 = v127
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v146 = v144
	goto L30
L34:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	v144 = v127 + int32(1)
	if v142 != 0 {
		v127 = v144
		v128 = v142
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v159 = int32(_a_F_sendDir_4)
	v160 = v62 - l1
	if v160 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v205 != 0 {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v205 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v166 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v167 = l1
	v168 = v159
	v169 = v160
	v170 = v166
	goto L45
L42:
	;
	v193 = v159
	v197 = int32(0)
	goto L43
L43:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v205 = v197 - v198
	goto L37
L44:
	;
	v193 = v188
	v197 = v190
	goto L43
L45:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if base.B2i32(v170 != v172)|base.B2i32(v172 == int32(0)) != 0 {
		v188 = v168
		v190 = v170
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v188 = v182
	v190 = int32(0)
	goto L44
L47:
	;
	v178 = v169 - int32(1)
	if v178 == int32(0) {
		v188 = v168
		v190 = v170
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v181 = int32(1)
	v182 = v168 + v181
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	if v183 != 0 {
		v167 = v167 + v181
		v168 = v182
		v169 = v178
		v170 = v183
		goto L45
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	v206 = int32(1663)
	if base.Ui32(v160) < base.Ui32(int32(15)) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v264 = F_strtox_2(m, v72, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L70
L53:
	;
	v302 = v10
	v303 = v10
	v304 = v206
	goto L6
L54:
	;
	goto L55
L55:
	;
	v209 = int32(15)
	v210 = v62 - v209
	v211 = int32(_a_F_sendDir_5)
	goto L58
L56:
	;
	if v249-v250 != 0 {
		v302 = v10
		v303 = v10
		v304 = v206
		goto L6
	} else {
		goto L69
	}
L58:
	;
	goto L59
L59:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if v218 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v219 = v210
	v220 = v211
	v221 = v209
	v222 = v218
	goto L64
L61:
	;
	v245 = v211
	v249 = int32(0)
	goto L62
L62:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	goto L56
L63:
	;
	v245 = v240
	v249 = v242
	goto L62
L64:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	if base.B2i32(v222 != v224)|base.B2i32(v224 == int32(0)) != 0 {
		v240 = v220
		v242 = v222
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v240 = v234
	v242 = int32(0)
	goto L63
L66:
	;
	v230 = v221 - int32(1)
	if v230 == int32(0) {
		v240 = v220
		v242 = v222
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v233 = int32(1)
	v234 = v220 + v233
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	if v235 != 0 {
		v219 = v219 + v233
		v220 = v234
		v221 = v230
		v222 = v235
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
	v302 = int32(1)
	v303 = base.I32_wrap_i64(v264)
	v304 = int32(1663)
	goto L6
L71:
	;
	if v295 != 0 {
		goto L78
	} else {
		goto L79
	}
L72:
	;
	v295 = v293 - v294
	goto L71
L73:
	;
	v278 = l1
	v279 = v269
	goto L74
L74:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+1)))
	if v283 == int32(0) {
		v293 = v283
		v294 = v282
		goto L72
	} else {
		goto L76
	}
L75:
	;
	v293 = v283
	v294 = v282
	goto L72
L76:
	;
	v286 = int32(1)
	if v283 == v282 {
		v278 = v278 + v286
		v279 = v279 + v286
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v296 = int32(1663)
	goto L80
L79:
	;
	v296 = int32(1664)
	goto L80
L80:
	;
	v302 = base.B2i32(v295 == int32(0))
	v303 = v10
	v304 = v296
	goto L6
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4022 = m.ExcPending
	if v4022 != 0 {
		goto L4
	} else {
		goto L701
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4003 = m.ExcPending
	if v4003 != 0 {
		goto L4
	} else {
		goto L697
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3983 = m.ExcPending
	if v3983 != 0 {
		goto L4
	} else {
		goto L692
	}
L84:
	;
	v307 = F_ReadDir(m, v305, l1)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	if v307 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v319 = l0
	v320 = l1
	v321 = l2
	v322 = l3
	v323 = l4
	v324 = l5
	v325 = l6
	v326 = l7
	v327 = l8
	v328 = v307
	v329 = v49
	v338 = v57
	v340 = l2 + v49 + int32(2369)
	v346 = v302
	v347 = v303
	v349 = v304
	v351 = v305
	v359 = l1 + l2 + int32(1)
	v360 = v49 + int32(2368) | int32(2)
	v363 = v45
	goto L89
L87:
	;
	v3936 = v49
	v3945 = v57
	v3958 = v305
	v3970 = v45
	goto L88
L88:
	;
	if v3945 != 0 {
		goto L687
	} else {
		goto L688
	}
L89:
	;
	v365 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v329)+2268)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v329)+2264)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v329)+2260)) = v365
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+19)))
	if v371 != int32(46) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v3936 = v329
	v3945 = v338
	v3958 = v351
	v3970 = v3922
	goto L88
L91:
	;
	v3924 = F_ReadDir(m, v351, v320)
	mBase = m.M
	v3925 = m.ExcPending
	if v3925 != 0 {
		goto L4
	} else {
		goto L685
	}
L92:
	;
	v384 = v328 + int32(19)
	v385 = int32(_a_F_sendDir_6)
	goto L99
L93:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+20)))
	if v374 == int32(0) {
		v3922 = v363
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+20)))
	if v377 != int32(46) {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+21)))
	if v380 == int32(0) {
		v3922 = v363
		goto L91
	} else {
		goto L96
	}
L96:
	;
	goto L92
L97:
	;
	if v423-v424 == int32(0) {
		v3922 = v363
		goto L91
	} else {
		goto L110
	}
L99:
	;
	goto L100
L100:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v392 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v393 = v384
	v394 = v385
	v395 = int32(9)
	v396 = v392
	goto L105
L102:
	;
	v419 = v385
	v423 = int32(0)
	goto L103
L103:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419))))
	goto L97
L104:
	;
	v419 = v414
	v423 = v416
	goto L103
L105:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if base.B2i32(v396 != v398)|base.B2i32(v398 == int32(0)) != 0 {
		v414 = v394
		v416 = v396
		goto L104
	} else {
		goto L107
	}
L106:
	;
	v414 = v408
	v416 = int32(0)
	goto L104
L107:
	;
	v404 = v395 - int32(1)
	if v404 == int32(0) {
		v414 = v394
		v416 = v396
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v407 = int32(1)
	v408 = v394 + v407
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+1)))
	if v409 != 0 {
		v393 = v393 + v407
		v394 = v408
		v395 = v404
		v396 = v409
		goto L105
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	v434 = int32(_a_F_sendDir_7)
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[3])))
	if base.B2i32(v437 == int32(0))|base.B2i32(v437 != v440) != 0 {
		v458 = v437
		v459 = v440
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v458-v459 == int32(0) {
		v3922 = v363
		goto L91
	} else {
		goto L118
	}
L112:
	;
	goto L111
L113:
	;
	v443 = v384
	v444 = v434
	goto L114
L114:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+1)))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+1)))
	if v448 == int32(0) {
		v458 = v448
		v459 = v447
		goto L112
	} else {
		goto L116
	}
L115:
	;
	v458 = v448
	v459 = v447
	goto L112
L116:
	;
	v451 = int32(1)
	if v448 == v447 {
		v443 = v443 + v451
		v444 = v444 + v451
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[4]))
	if v464 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[5])))
	if v469 == int32(1) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L121
L123:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[6])))
	if v479 != v481 {
		goto L83
	} else {
		goto L127
	}
L124:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[7]))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)+316))
	v477 = base.B2i32(v475 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[5])) = uint8(v477)
	v479 = v477
	goto L126
L125:
	;
	v479 = int32(0)
	goto L126
L126:
	;
	goto L123
L127:
	;
	v483 = int32(_a_F_sendDir_8)
	goto L133
L128:
	;
	if v347 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L129:
	;
	v1040 = int32(1)
	goto L128
L130:
	;
	v1025 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L4
	} else {
		goto L270
	}
L131:
	;
	if v521-v522 == int32(0) {
		goto L130
	} else {
		goto L144
	}
L133:
	;
	goto L134
L134:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v490 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v491 = v384
	v492 = v483
	v493 = int32(25)
	v494 = v490
	goto L139
L136:
	;
	v517 = v483
	v521 = int32(0)
	goto L137
L137:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517))))
	goto L131
L138:
	;
	v517 = v512
	v521 = v514
	goto L137
L139:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492))))
	if base.B2i32(v494 != v496)|base.B2i32(v496 == int32(0)) != 0 {
		v512 = v492
		v514 = v494
		goto L138
	} else {
		goto L141
	}
L140:
	;
	v512 = v506
	v514 = int32(0)
	goto L138
L141:
	;
	v502 = v493 - int32(1)
	if v502 == int32(0) {
		v512 = v492
		v514 = v494
		goto L138
	} else {
		goto L142
	}
L142:
	;
	v505 = int32(1)
	v506 = v492 + v505
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+1)))
	if v507 != 0 {
		v491 = v491 + v505
		v492 = v506
		v493 = v502
		v494 = v507
		goto L139
	} else {
		goto L143
	}
L143:
	;
	goto L140
L144:
	;
	v532 = int32(_a_F_sendDir_9)
	goto L147
L145:
	;
	if v570-v571 == int32(0) {
		goto L130
	} else {
		goto L158
	}
L147:
	;
	goto L148
L148:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v539 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v540 = v384
	v541 = v532
	v542 = int32(21)
	v543 = v539
	goto L153
L150:
	;
	v566 = v532
	v570 = int32(0)
	goto L151
L151:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
	goto L145
L152:
	;
	v566 = v561
	v570 = v563
	goto L151
L153:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	if base.B2i32(v543 != v545)|base.B2i32(v545 == int32(0)) != 0 {
		v561 = v541
		v563 = v543
		goto L152
	} else {
		goto L155
	}
L154:
	;
	v561 = v555
	v563 = int32(0)
	goto L152
L155:
	;
	v551 = v542 - int32(1)
	if v551 == int32(0) {
		v561 = v541
		v563 = v543
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v554 = int32(1)
	v555 = v541 + v554
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+1)))
	if v556 != 0 {
		v540 = v540 + v554
		v541 = v555
		v542 = v551
		v543 = v556
		goto L153
	} else {
		goto L157
	}
L157:
	;
	goto L154
L158:
	;
	v581 = int32(_a_F_sendDir_10)
	goto L161
L159:
	;
	if v619-v620 == int32(0) {
		goto L130
	} else {
		goto L172
	}
L161:
	;
	goto L162
L162:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v588 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v589 = v384
	v590 = v581
	v591 = int32(16)
	v592 = v588
	goto L167
L164:
	;
	v615 = v581
	v619 = int32(0)
	goto L165
L165:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615))))
	goto L159
L166:
	;
	v615 = v610
	v619 = v612
	goto L165
L167:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	if base.B2i32(v592 != v594)|base.B2i32(v594 == int32(0)) != 0 {
		v610 = v590
		v612 = v592
		goto L166
	} else {
		goto L169
	}
L168:
	;
	v610 = v604
	v612 = int32(0)
	goto L166
L169:
	;
	v600 = v591 - int32(1)
	if v600 == int32(0) {
		v610 = v590
		v612 = v592
		goto L166
	} else {
		goto L170
	}
L170:
	;
	v603 = int32(1)
	v604 = v590 + v603
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589)+1)))
	if v605 != 0 {
		v589 = v589 + v603
		v590 = v604
		v591 = v600
		v592 = v605
		goto L167
	} else {
		goto L171
	}
L171:
	;
	goto L168
L172:
	;
	v630 = int32(_a_F_sendDir_11)
	goto L175
L173:
	;
	if v668-v669 == int32(0) {
		goto L130
	} else {
		goto L186
	}
L175:
	;
	goto L176
L176:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v637 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v638 = v384
	v639 = v630
	v640 = int32(13)
	v641 = v637
	goto L181
L178:
	;
	v664 = v630
	v668 = int32(0)
	goto L179
L179:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	goto L173
L180:
	;
	v664 = v659
	v668 = v661
	goto L179
L181:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639))))
	if base.B2i32(v641 != v643)|base.B2i32(v643 == int32(0)) != 0 {
		v659 = v639
		v661 = v641
		goto L180
	} else {
		goto L183
	}
L182:
	;
	v659 = v653
	v661 = int32(0)
	goto L180
L183:
	;
	v649 = v640 - int32(1)
	if v649 == int32(0) {
		v659 = v639
		v661 = v641
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v652 = int32(1)
	v653 = v639 + v652
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638)+1)))
	if v654 != 0 {
		v638 = v638 + v652
		v639 = v653
		v640 = v649
		v641 = v654
		goto L181
	} else {
		goto L185
	}
L185:
	;
	goto L182
L186:
	;
	v679 = int32(_a_F_sendDir_12)
	goto L189
L187:
	;
	if v717-v718 == int32(0) {
		goto L130
	} else {
		goto L200
	}
L189:
	;
	goto L190
L190:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v686 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v687 = v384
	v688 = v679
	v689 = int32(15)
	v690 = v686
	goto L195
L192:
	;
	v713 = v679
	v717 = int32(0)
	goto L193
L193:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	goto L187
L194:
	;
	v713 = v708
	v717 = v710
	goto L193
L195:
	;
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688))))
	if base.B2i32(v690 != v692)|base.B2i32(v692 == int32(0)) != 0 {
		v708 = v688
		v710 = v690
		goto L194
	} else {
		goto L197
	}
L196:
	;
	v708 = v702
	v710 = int32(0)
	goto L194
L197:
	;
	v698 = v689 - int32(1)
	if v698 == int32(0) {
		v708 = v688
		v710 = v690
		goto L194
	} else {
		goto L198
	}
L198:
	;
	v701 = int32(1)
	v702 = v688 + v701
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687)+1)))
	if v703 != 0 {
		v687 = v687 + v701
		v688 = v702
		v689 = v698
		v690 = v703
		goto L195
	} else {
		goto L199
	}
L199:
	;
	goto L196
L200:
	;
	v728 = int32(_a_F_sendDir_13)
	goto L203
L201:
	;
	if v766-v767 == int32(0) {
		goto L130
	} else {
		goto L214
	}
L203:
	;
	goto L204
L204:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v735 != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v736 = v384
	v737 = v728
	v738 = int32(16)
	v739 = v735
	goto L209
L206:
	;
	v762 = v728
	v766 = int32(0)
	goto L207
L207:
	;
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762))))
	goto L201
L208:
	;
	v762 = v757
	v766 = v759
	goto L207
L209:
	;
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737))))
	if base.B2i32(v739 != v741)|base.B2i32(v741 == int32(0)) != 0 {
		v757 = v737
		v759 = v739
		goto L208
	} else {
		goto L211
	}
L210:
	;
	v757 = v751
	v759 = int32(0)
	goto L208
L211:
	;
	v747 = v738 - int32(1)
	if v747 == int32(0) {
		v757 = v737
		v759 = v739
		goto L208
	} else {
		goto L212
	}
L212:
	;
	v750 = int32(1)
	v751 = v737 + v750
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736)+1)))
	if v752 != 0 {
		v736 = v736 + v750
		v737 = v751
		v738 = v747
		v739 = v752
		goto L209
	} else {
		goto L213
	}
L213:
	;
	goto L210
L214:
	;
	v777 = int32(_a_F_sendDir_14)
	goto L217
L215:
	;
	if v815-v816 == int32(0) {
		goto L130
	} else {
		goto L228
	}
L217:
	;
	goto L218
L218:
	;
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v784 != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v785 = v384
	v786 = v777
	v787 = int32(15)
	v788 = v784
	goto L223
L220:
	;
	v811 = v777
	v815 = int32(0)
	goto L221
L221:
	;
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811))))
	goto L215
L222:
	;
	v811 = v806
	v815 = v808
	goto L221
L223:
	;
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786))))
	if base.B2i32(v788 != v790)|base.B2i32(v790 == int32(0)) != 0 {
		v806 = v786
		v808 = v788
		goto L222
	} else {
		goto L225
	}
L224:
	;
	v806 = v800
	v808 = int32(0)
	goto L222
L225:
	;
	v796 = v787 - int32(1)
	if v796 == int32(0) {
		v806 = v786
		v808 = v788
		goto L222
	} else {
		goto L226
	}
L226:
	;
	v799 = int32(1)
	v800 = v786 + v799
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v785)+1)))
	if v801 != 0 {
		v785 = v785 + v799
		v786 = v800
		v787 = v796
		v788 = v801
		goto L223
	} else {
		goto L227
	}
L227:
	;
	goto L224
L228:
	;
	v826 = int32(_a_F_sendDir_15)
	goto L231
L229:
	;
	if v864-v865 == int32(0) {
		goto L130
	} else {
		goto L242
	}
L231:
	;
	goto L232
L232:
	;
	v833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v833 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v834 = v384
	v835 = v826
	v836 = int32(16)
	v837 = v833
	goto L237
L234:
	;
	v860 = v826
	v864 = int32(0)
	goto L235
L235:
	;
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860))))
	goto L229
L236:
	;
	v860 = v855
	v864 = v857
	goto L235
L237:
	;
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835))))
	if base.B2i32(v837 != v839)|base.B2i32(v839 == int32(0)) != 0 {
		v855 = v835
		v857 = v837
		goto L236
	} else {
		goto L239
	}
L238:
	;
	v855 = v849
	v857 = int32(0)
	goto L236
L239:
	;
	v845 = v836 - int32(1)
	if v845 == int32(0) {
		v855 = v835
		v857 = v837
		goto L236
	} else {
		goto L240
	}
L240:
	;
	v848 = int32(1)
	v849 = v835 + v848
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834)+1)))
	if v850 != 0 {
		v834 = v834 + v848
		v835 = v849
		v836 = v845
		v837 = v850
		goto L237
	} else {
		goto L241
	}
L241:
	;
	goto L238
L242:
	;
	v875 = int32(0)
	if v346 == v875 {
		v1040 = v875
		goto L128
	} else {
		goto L243
	}
L243:
	;
	v879 = v329 + int32(2268)
	v881 = v329 + int32(2264)
	v883 = v329 + int32(2260)
	v884 = int32(0)
	v889 = m.G0
	v891 = v889 - int32(16)
	m.G0 = v891
	*(*int32)(unsafe.Add(mBase, uint32(v879))) = v884
	*(*int32)(unsafe.Add(mBase, uint32(v881))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v883))) = v884
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if base.Ui32((v899-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v981 = v884
		goto L245
	} else {
		goto L246
	}
L244:
	;
	if v981 == int32(0) {
		v1040 = v981
		goto L128
	} else {
		goto L261
	}
L245:
	;
	m.G0 = v891 + int32(16)
	goto L244
L246:
	;
	v906 = int32(_a_F_sendDir_16)
	*(*int32)(unsafe.Add(mBase, _c_F_sendDir[8])) = int32(0)
	v912 = F_strtoul(m, v384, v891+int32(8), int32(10))
	mBase = m.M
	v914 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[8]))
	if v914 != 0 {
		v981 = v884
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v891)+8))
	if base.B2i32(v912 == int32(0))|base.B2i32(v384 == v917) != 0 {
		v981 = v884
		goto L245
	} else {
		goto L248
	}
L248:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v917))))
	if v920 != int32(95) {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	if v936&int32(255) == int32(46) {
		goto L254
	} else {
		goto L255
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891)+12)) = int32(0)
	v936 = v920
	v937 = v917
	goto L249
L251:
	;
	goto L252
L252:
	;
	v929 = F_forkname_chars(m, v917+int32(1), v891+int32(12))
	mBase = m.M
	if v929 <= int32(0) {
		v981 = v884
		goto L245
	} else {
		goto L253
	}
L253:
	;
	v934 = v929 + v917 + int32(1)
	v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934))))
	v936 = v935
	v937 = v934
	goto L249
L254:
	;
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937)+1)))
	if base.Ui32((v942-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v981 = v884
		goto L245
	} else {
		goto L257
	}
L255:
	;
	v968 = v884
	v969 = v936
	goto L256
L256:
	;
	if v969&int32(255) != 0 {
		v981 = v884
		goto L245
	} else {
		goto L260
	}
L257:
	;
	v949 = int32(_a_F_sendDir_16)
	*(*int32)(unsafe.Add(mBase, _c_F_sendDir[8])) = int32(0)
	v953 = v937 + int32(1)
	v957 = F_strtoul(m, v953, v891+int32(8), int32(10))
	mBase = m.M
	v959 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[8]))
	if v959 != 0 {
		v981 = v884
		goto L245
	} else {
		goto L258
	}
L258:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v891)+8))
	if base.B2i32(v957 == int32(0))|base.B2i32(v953 == v962) != 0 {
		v981 = v884
		goto L245
	} else {
		goto L259
	}
L259:
	;
	v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v962))))
	v968 = v957
	v969 = v965
	goto L256
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v879))) = v912
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v891)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v881))) = v973
	*(*int32)(unsafe.Add(mBase, uint32(v883))) = v968
	v981 = int32(1)
	goto L245
L261:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2264))
	if v987 == int32(3) {
		v1040 = v981
		goto L128
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+176)) = v320
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2268))
	*(*int32)(unsafe.Add(mBase, uint32(v329)+180)) = v991
	v994 = v329 + int32(192)
	v999 = F_pg_snprintf(m, v994, int32(1024), int32(_a_F_sendDir_17), v329+int32(176))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L4
	} else {
		goto L263
	}
L263:
	;
	v1005 = F___fstatat(m, int32(-100), v994, v329+int32(2272), int32(256))
	mBase = m.M
	goto L264
L264:
	;
	if v1005 != 0 {
		goto L129
	} else {
		goto L265
	}
L265:
	;
	v1008 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L4
	} else {
		goto L266
	}
L266:
	;
	if v1008 == int32(0) {
		v3922 = v363
		goto L91
	} else {
		goto L267
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+16)) = v384
	F_errmsg_internal(m, int32(_a_F_sendDir_18), v329+int32(16))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L4
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1333), int32(_a_F_sendDir_20))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L4
	} else {
		goto L269
	}
L269:
	;
	v3922 = v363
	goto L91
L270:
	;
	if v1025 == int32(0) {
		v3922 = v363
		goto L91
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329))) = v384
	F_errmsg_internal(m, int32(_a_F_sendDir_21), v329)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L4
	} else {
		goto L272
	}
L272:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1298), int32(_a_F_sendDir_20))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	v3922 = v363
	goto L91
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+148)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v329)+144)) = v320
	v1438 = v329 + int32(2368)
	v1443 = F_pg_snprintf(m, v1438, int32(2048), int32(_a_F_sendDir_22), v329+int32(144))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L4
	} else {
		goto L316
	}
L275:
	;
	v1043 = int32(0)
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v1044 != int32(116) {
		v1347 = v1043
		goto L276
	} else {
		goto L277
	}
L276:
	;
	if v1347 == int32(0) {
		goto L274
	} else {
		goto L311
	}
L277:
	;
	v1059 = int32(1)
	goto L278
L278:
	;
	v1095 = v1059 + int32(1)
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1059+v384))))
	if base.Ui32((v1097-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v1059 = v1095
		goto L278
	} else {
		goto L280
	}
L279:
	;
	if base.B2i32(v1059 == int32(1))|base.B2i32(v1097 != int32(95)) != 0 {
		v1347 = v1043
		goto L276
	} else {
		goto L281
	}
L280:
	;
	goto L279
L281:
	;
	v1121 = v1095
	goto L282
L282:
	;
	v1156 = v1121 + int32(1)
	v1157 = v1121 + v384
	v1158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157))))
	if base.Ui32((v1158-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v1121 = v1156
		goto L282
	} else {
		goto L284
	}
L283:
	;
	if v1121 == v1095 {
		v1347 = v1043
		goto L276
	} else {
		goto L285
	}
L284:
	;
	goto L283
L285:
	;
	if v1158 == int32(95) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1169 = v1157 + int32(1)
	v1173 = int32(3)
	v1176 = F_strncmp(m, int32(_a_F_sendDir_23), v1169, v1173)
	mBase = m.M
	if v1176 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L287:
	;
	v1212 = v1121
	v1213 = v1158
	goto L288
L288:
	;
	if v1213 == int32(46) {
		goto L304
	} else {
		goto L305
	}
L289:
	;
	if v1205 <= int32(0) {
		v1347 = v1043
		goto L276
	} else {
		goto L303
	}
L290:
	;
	goto L289
L292:
	;
	v1205 = v1198
	goto L290
L293:
	;
	v1198 = v1173
	goto L292
L294:
	;
	goto L295
L295:
	;
	v1180 = int32(2)
	v1184 = F_strncmp(m, int32(_a_F_sendDir_24), v1169, v1180)
	mBase = m.M
	if v1184 == int32(0) {
		v1198 = v1180
		goto L292
	} else {
		goto L296
	}
L296:
	;
	v1187 = int32(4)
	v1190 = F_strncmp(m, int32(_a_F_sendDir_25), v1169, v1187)
	mBase = m.M
	if v1190 == int32(0) {
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
	v1205 = int32(0)
	goto L290
L300:
	;
	v1205 = v1187
	goto L290
L303:
	;
	v1209 = v1205 + v1156
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384+v1209))))
	v1212 = v1209
	v1213 = v1211
	goto L288
L304:
	;
	v1229 = int32(1)
	goto L307
L305:
	;
	v1296 = v1213
	goto L306
L306:
	;
	v1347 = base.B2i32(v1296 == int32(0))
	goto L276
L307:
	;
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1229+(v1212+v384)))))
	if base.Ui32((v1267-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v1229 = v1229 + int32(1)
		goto L307
	} else {
		goto L309
	}
L308:
	;
	if v1229 < int32(2) {
		v1347 = v1043
		goto L276
	} else {
		goto L310
	}
L309:
	;
	goto L308
L310:
	;
	v1296 = v1267
	goto L306
L311:
	;
	v1374 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L4
	} else {
		goto L312
	}
L312:
	;
	if v1374 == int32(0) {
		v3922 = v363
		goto L91
	} else {
		goto L313
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+160)) = v384
	F_errmsg_internal(m, int32(_a_F_sendDir_26), v329+int32(160))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L4
	} else {
		goto L314
	}
L314:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1344), int32(_a_F_sendDir_20))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L4
	} else {
		goto L315
	}
L315:
	;
	v3922 = v363
	goto L91
L316:
	;
	v1445 = *(*int64)(unsafe.Add(mBase, uint32(v329)+2368))
	v1448 = *(*int64)(unsafe.Add(mBase, uint32(v329)+2376))
	v1454 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v329+int32(2384)))))
	if v1445^int64(7809631459536744238)|(v1448^int64(8389765628430872623))|(v1454^int64(7106418)) == int64(0) {
		v3922 = v363
		goto L91
	} else {
		goto L317
	}
L317:
	;
	v1464 = F___fstatat(m, int32(-100), v1438, v329+int32(2272), int32(256))
	mBase = m.M
	goto L320
L318:
	;
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2276))
	v1759 = v1757 & int32(_a_F_sendDir_27)
	v1760 = int32(_a_F_sendDir_28)
	v1763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320))))
	v1766 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[9])))
	if base.B2i32(v1763 == int32(0))|base.B2i32(v1763 != v1766) != 0 {
		v1784 = v1763
		v1785 = v1766
		goto L403
	} else {
		goto L404
	}
L319:
	;
	v1727 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L4
	} else {
		goto L392
	}
L320:
	;
	if v1464 == int32(0) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1467 = int32(_a_F_sendDir_29)
	v1470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	v1473 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[10])))
	if base.B2i32(v1470 == int32(0))|base.B2i32(v1470 != v1473) != 0 {
		v1491 = v1470
		v1492 = v1473
		goto L325
	} else {
		goto L326
	}
L322:
	;
	goto L323
L323:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[8]))
	if v1703 == int32(44) {
		v3922 = v363
		goto L91
	} else {
		goto L387
	}
L324:
	;
	if v1491-v1492 == int32(0) {
		goto L319
	} else {
		goto L331
	}
L325:
	;
	goto L324
L326:
	;
	v1476 = v384
	v1477 = v1467
	goto L327
L327:
	;
	v1480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1477)+1)))
	v1481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1476)+1)))
	if v1481 == int32(0) {
		v1491 = v1481
		v1492 = v1480
		goto L325
	} else {
		goto L329
	}
L328:
	;
	v1491 = v1481
	v1492 = v1480
	goto L325
L329:
	;
	v1484 = int32(1)
	if v1481 == v1480 {
		v1476 = v1476 + v1484
		v1477 = v1477 + v1484
		goto L327
	} else {
		goto L330
	}
L330:
	;
	goto L328
L331:
	;
	v1496 = int32(_a_F_sendDir_30)
	v1499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[11])))
	if base.B2i32(v1499 == int32(0))|base.B2i32(v1499 != v1502) != 0 {
		v1520 = v1499
		v1521 = v1502
		goto L333
	} else {
		goto L334
	}
L332:
	;
	if v1520-v1521 == int32(0) {
		goto L319
	} else {
		goto L339
	}
L333:
	;
	goto L332
L334:
	;
	v1505 = v384
	v1506 = v1496
	goto L335
L335:
	;
	v1509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1506)+1)))
	v1510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1505)+1)))
	if v1510 == int32(0) {
		v1520 = v1510
		v1521 = v1509
		goto L333
	} else {
		goto L337
	}
L336:
	;
	v1520 = v1510
	v1521 = v1509
	goto L333
L337:
	;
	v1513 = int32(1)
	if v1510 == v1509 {
		v1505 = v1505 + v1513
		v1506 = v1506 + v1513
		goto L335
	} else {
		goto L338
	}
L338:
	;
	goto L336
L339:
	;
	v1525 = int32(_a_F_sendDir_31)
	v1528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	v1531 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[12])))
	if base.B2i32(v1528 == int32(0))|base.B2i32(v1528 != v1531) != 0 {
		v1549 = v1528
		v1550 = v1531
		goto L341
	} else {
		goto L342
	}
L340:
	;
	if v1549-v1550 == int32(0) {
		goto L319
	} else {
		goto L347
	}
L341:
	;
	goto L340
L342:
	;
	v1534 = v384
	v1535 = v1525
	goto L343
L343:
	;
	v1538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1535)+1)))
	v1539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1534)+1)))
	if v1539 == int32(0) {
		v1549 = v1539
		v1550 = v1538
		goto L341
	} else {
		goto L345
	}
L344:
	;
	v1549 = v1539
	v1550 = v1538
	goto L341
L345:
	;
	v1542 = int32(1)
	if v1539 == v1538 {
		v1534 = v1534 + v1542
		v1535 = v1535 + v1542
		goto L343
	} else {
		goto L346
	}
L346:
	;
	goto L344
L347:
	;
	v1554 = int32(_a_F_sendDir_32)
	v1557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	v1560 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[13])))
	if base.B2i32(v1557 == int32(0))|base.B2i32(v1557 != v1560) != 0 {
		v1578 = v1557
		v1579 = v1560
		goto L349
	} else {
		goto L350
	}
L348:
	;
	if v1578-v1579 == int32(0) {
		goto L319
	} else {
		goto L355
	}
L349:
	;
	goto L348
L350:
	;
	v1563 = v384
	v1564 = v1554
	goto L351
L351:
	;
	v1567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1564)+1)))
	v1568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1563)+1)))
	if v1568 == int32(0) {
		v1578 = v1568
		v1579 = v1567
		goto L349
	} else {
		goto L353
	}
L352:
	;
	v1578 = v1568
	v1579 = v1567
	goto L349
L353:
	;
	v1571 = int32(1)
	if v1568 == v1567 {
		v1563 = v1563 + v1571
		v1564 = v1564 + v1571
		goto L351
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	v1583 = int32(_a_F_sendDir_33)
	v1586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	v1589 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[14])))
	if base.B2i32(v1586 == int32(0))|base.B2i32(v1586 != v1589) != 0 {
		v1607 = v1586
		v1608 = v1589
		goto L357
	} else {
		goto L358
	}
L356:
	;
	if v1607-v1608 == int32(0) {
		goto L319
	} else {
		goto L363
	}
L357:
	;
	goto L356
L358:
	;
	v1592 = v384
	v1593 = v1583
	goto L359
L359:
	;
	v1596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593)+1)))
	v1597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1592)+1)))
	if v1597 == int32(0) {
		v1607 = v1597
		v1608 = v1596
		goto L357
	} else {
		goto L361
	}
L360:
	;
	v1607 = v1597
	v1608 = v1596
	goto L357
L361:
	;
	v1600 = int32(1)
	if v1597 == v1596 {
		v1592 = v1592 + v1600
		v1593 = v1593 + v1600
		goto L359
	} else {
		goto L362
	}
L362:
	;
	goto L360
L363:
	;
	v1612 = int32(_a_F_sendDir_34)
	v1615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	v1618 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[15])))
	if base.B2i32(v1615 == int32(0))|base.B2i32(v1615 != v1618) != 0 {
		v1636 = v1615
		v1637 = v1618
		goto L365
	} else {
		goto L366
	}
L364:
	;
	if v1636-v1637 == int32(0) {
		goto L319
	} else {
		goto L371
	}
L365:
	;
	goto L364
L366:
	;
	v1621 = v384
	v1622 = v1612
	goto L367
L367:
	;
	v1625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1622)+1)))
	v1626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1621)+1)))
	if v1626 == int32(0) {
		v1636 = v1626
		v1637 = v1625
		goto L365
	} else {
		goto L369
	}
L368:
	;
	v1636 = v1626
	v1637 = v1625
	goto L365
L369:
	;
	v1629 = int32(1)
	if v1626 == v1625 {
		v1621 = v1621 + v1629
		v1622 = v1622 + v1629
		goto L367
	} else {
		goto L370
	}
L370:
	;
	goto L368
L371:
	;
	v1641 = int32(_a_F_sendDir_35)
	v1644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	v1647 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[16])))
	if base.B2i32(v1644 == int32(0))|base.B2i32(v1644 != v1647) != 0 {
		v1665 = v1644
		v1666 = v1647
		goto L373
	} else {
		goto L374
	}
L372:
	;
	if v1665-v1666 == int32(0) {
		goto L319
	} else {
		goto L379
	}
L373:
	;
	goto L372
L374:
	;
	v1650 = v384
	v1651 = v1641
	goto L375
L375:
	;
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1651)+1)))
	v1655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1650)+1)))
	if v1655 == int32(0) {
		v1665 = v1655
		v1666 = v1654
		goto L373
	} else {
		goto L377
	}
L376:
	;
	v1665 = v1655
	v1666 = v1654
	goto L373
L377:
	;
	v1658 = int32(1)
	if v1655 == v1654 {
		v1650 = v1650 + v1658
		v1651 = v1651 + v1658
		goto L375
	} else {
		goto L378
	}
L378:
	;
	goto L376
L379:
	;
	v1670 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v329)+2376)))
	v1671 = *(*int64)(unsafe.Add(mBase, uint32(v329)+2368))
	if v1670|(v1671^int64(7809654480478154542)) != int64(0) {
		goto L318
	} else {
		goto L380
	}
L380:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2276))
	if v1677&int32(_a_F_sendDir_27) == int32(_a_F_sendDir_36) {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v329)+2276)) = v1683 | int32(_a_F_sendDir_37)
	goto L383
L382:
	;
	goto L383
L383:
	;
	v1689 = v329 + int32(2272)
	F__tarWriteHeader(m, v319, v340, int32(0), v1689, v322)
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L4
	} else {
		goto L384
	}
L384:
	;
	F__tarWriteHeader(m, v319, int32(_a_F_sendDir_38), int32(0), v1689, v322)
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L4
	} else {
		goto L385
	}
L385:
	;
	F__tarWriteHeader(m, v319, int32(_a_F_sendDir_39), int32(0), v1689, v322)
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L4
	} else {
		goto L386
	}
L386:
	;
	v3922 = v363 + int64(1536)
	goto L91
L387:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L4
	} else {
		goto L388
	}
L388:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L4
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+128)) = v329 + int32(2368)
	F_errmsg(m, int32(_a_F_sendDir_40), v329+int32(128))
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L4
	} else {
		goto L390
	}
L390:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1361), int32(_a_F_sendDir_20))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
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
	if v1727 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+32)) = v384
	F_errmsg_internal(m, int32(_a_F_sendDir_41), v329+int32(32))
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L4
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2276))
	if v1740&int32(_a_F_sendDir_27) == int32(_a_F_sendDir_36) {
		goto L398
	} else {
		goto L399
	}
L396:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1373), int32(_a_F_sendDir_20))
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L4
	} else {
		goto L397
	}
L397:
	;
	goto L395
L398:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v329)+2276)) = v1746 | int32(_a_F_sendDir_37)
	goto L400
L399:
	;
	goto L400
L400:
	;
	F__tarWriteHeader(m, v319, v340, int32(0), v329+int32(2272), v322)
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L4
	} else {
		goto L401
	}
L401:
	;
	v3922 = v363 + int64(512)
	goto L91
L402:
	;
	if v1784-v1785|base.B2i32(v1759 != int32(_a_F_sendDir_36)) == int32(0) {
		goto L409
	} else {
		goto L410
	}
L403:
	;
	goto L402
L404:
	;
	v1769 = v320
	v1770 = v1760
	goto L405
L405:
	;
	v1773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1770)+1)))
	v1774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1769)+1)))
	if v1774 == int32(0) {
		v1784 = v1774
		v1785 = v1773
		goto L403
	} else {
		goto L407
	}
L406:
	;
	v1784 = v1774
	v1785 = v1773
	goto L403
L407:
	;
	v1777 = int32(1)
	if v1774 == v1773 {
		v1769 = v1769 + v1777
		v1770 = v1770 + v1777
		goto L405
	} else {
		goto L408
	}
L408:
	;
	goto L406
L409:
	;
	v1795 = v329 + int32(192)
	v1797 = F_readlink(m, v329+int32(2368), v1795, int32(1024))
	mBase = m.M
	if v1797 < int32(0) {
		goto L82
	} else {
		goto L412
	}
L410:
	;
	goto L411
L411:
	;
	if v1759 != int32(_a_F_sendDir_42) {
		goto L416
	} else {
		goto L417
	}
L412:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v1797) {
		goto L81
	} else {
		goto L413
	}
L413:
	;
	v1803 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1797+v1795))) = uint8(v1803)
	F__tarWriteHeader(m, v319, v340, v1795, v329+int32(2272), v322)
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L4
	} else {
		goto L414
	}
L414:
	;
	v3922 = v363 + int64(512)
	goto L91
L415:
	;
	v3861 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v3862 = m.ExcPending
	if v3862 != 0 {
		goto L4
	} else {
		goto L681
	}
L416:
	;
	if v1759 != int32(_a_F_sendDir_37) {
		goto L415
	} else {
		goto L419
	}
L417:
	;
	goto L418
L418:
	;
	v1983 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v329)+2256)) = v1983
	*(*int32)(unsafe.Add(mBase, uint32(v329)+2252)) = v1983
	if base.B2i32(v327 == v1983)|(v1040^int32(1)) == v1983 {
		goto L445
	} else {
		goto L446
	}
L419:
	;
	F__tarWriteHeader(m, v319, v340, int32(0), v329+int32(2272), v322)
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L4
	} else {
		goto L420
	}
L420:
	;
	if v323 == int32(0) {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	v1963 = v363 + int64(512)
	if v1930 == int32(0) {
		v3922 = v1963
		goto L91
	} else {
		goto L442
	}
L422:
	;
	v1930 = int32(1)
	goto L421
L423:
	;
	goto L424
L424:
	;
	v1823 = int32(1)
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	if v1824 <= int32(0) {
		v1930 = v1823
		goto L421
	} else {
		goto L425
	}
L425:
	;
	v1827 = int32(0)
	if v1827 < v1824 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1830 = v1824
	goto L428
L427:
	;
	v1830 = v1827
	goto L428
L428:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	v1842 = int32(0)
	goto L429
L429:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v1831+v1842<<(uint(int32(2))%32))))
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+8))
	if v1883 == int32(0) {
		goto L431
	} else {
		goto L432
	}
L430:
	;
	v1930 = v1823
	goto L421
L431:
	;
	v1914 = v1842 + int32(1)
	if v1914 != v1830 {
		v1842 = v1914
		goto L429
	} else {
		goto L441
	}
L432:
	;
	v1888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1883))))
	v1891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	if base.B2i32(v1888 == int32(0))|base.B2i32(v1888 != v1891) != 0 {
		v1909 = v1888
		v1910 = v1891
		goto L434
	} else {
		goto L435
	}
L433:
	;
	if v1909-v1910 != 0 {
		goto L431
	} else {
		goto L440
	}
L434:
	;
	goto L433
L435:
	;
	v1894 = v1883
	v1895 = v360
	goto L436
L436:
	;
	v1898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1895)+1)))
	v1899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1894)+1)))
	if v1899 == int32(0) {
		v1909 = v1899
		v1910 = v1898
		goto L434
	} else {
		goto L438
	}
L437:
	;
	v1909 = v1899
	v1910 = v1898
	goto L434
L438:
	;
	v1902 = int32(1)
	if v1899 == v1898 {
		v1894 = v1894 + v1902
		v1895 = v1895 + v1902
		goto L436
	} else {
		goto L439
	}
L439:
	;
	goto L437
L440:
	;
	v1930 = int32(0)
	goto L421
L441:
	;
	goto L430
L442:
	;
	v1966 = *(*int64)(unsafe.Add(mBase, uint32(v329)+2368))
	v1969 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v329)+2376)))
	if v324|base.B2i32(v1966^int64(7809932656919981870)|(v1969^int64(6516851)) != int64(0)) == int32(0) {
		v3922 = v1963
		goto L91
	} else {
		goto L443
	}
L443:
	;
	v1980 = F_sendDir(m, v319, v329+int32(2368), v321, v322, v323, v324, v325, v326, v327)
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L4
	} else {
		goto L444
	}
L444:
	;
	v3922 = v1980 + v1963
	goto L91
L445:
	;
	if v326 != 0 {
		goto L449
	} else {
		goto L450
	}
L446:
	;
	v3806 = v1983
	v3809 = v340
	goto L447
L447:
	;
	if v322 == int32(0) {
		goto L676
	} else {
		goto L677
	}
L448:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2268))
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2264))
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2260))
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2296))
	v2013 = v329 + int32(2256)
	v2015 = v329 + int32(2252)
	v2016 = int32(0)
	v2017 = m.G0
	v2019 = v2017 - int32(128)
	m.G0 = v2019
	if v2011&int32(_a_F_sendDir_43)|base.B2i32(v2009 == int32(1))|base.B2i32(base.Ui32(int32(1073741824)) < base.Ui32(v2011)) != 0 {
		v3700 = v2016
		goto L456
	} else {
		goto L457
	}
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+120)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v329)+116)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v329)+112)) = int32(_a_F_sendDir_44)
	v2002 = F_psprintf(m, int32(_a_F_sendDir_45), v329+int32(112))
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L4
	} else {
		goto L452
	}
L450:
	;
	goto L451
L451:
	;
	v2004 = F_pstrdup(m, v340)
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L4
	} else {
		goto L453
	}
L452:
	;
	v2006 = v2002
	v2007 = v326
	goto L448
L453:
	;
	v2006 = v2004
	v2007 = v349
	goto L448
L454:
	;
	if v3700 == int32(1) {
		goto L668
	} else {
		goto L669
	}
L455:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3731 = m.ExcPending
	if v3731 != 0 {
		goto L4
	} else {
		goto L664
	}
L456:
	;
	m.G0 = v2019 + int32(128)
	goto L454
L457:
	;
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(v327)+24))
	v2030 = F_strlen(m, v2006)
	mBase = m.M
	v2036 = v2030 - int32(1636608432)
	if v2006&int32(3) != 0 {
		goto L462
	} else {
		goto L463
	}
L458:
	;
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(v2029)+20))
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2029)+12))
	v2297 = (v2290 ^ v2282 - base.I32_rotl(v2290, int32(24))) & v2296
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v2295+v2297<<(uint(int32(4))%32))))
	if v2301 != 0 {
		goto L499
	} else {
		goto L500
	}
L459:
	;
	v2268 = int32(14)
	v2270 = v2264 ^ v2265 - base.I32_rotl(v2264, v2268)
	v2274 = v2270 ^ v2263 - base.I32_rotl(v2270, int32(11))
	v2278 = v2274 ^ v2264 - base.I32_rotl(v2274, int32(25))
	v2282 = v2278 ^ v2270 - base.I32_rotl(v2278, int32(16))
	v2286 = v2282 ^ v2274 - base.I32_rotl(v2282, int32(4))
	v2290 = v2286 ^ v2278 - base.I32_rotl(v2286, v2268)
	goto L458
L460:
	;
	switch v2194 - int32(1) {
	case 0:
		v2256 = v2195
		v2257 = v2196
		v2258 = v2197
		goto L487
	case 1:
		v2249 = v2195
		v2250 = v2196
		v2251 = v2197
		goto L488
	case 2:
		v2242 = v2195
		v2243 = v2196
		v2244 = v2197
		goto L489
	case 3:
		v2236 = v2196
		v2237 = v2197
		goto L490
	case 4:
		v2232 = v2196
		v2233 = v2197
		goto L491
	case 5:
		v2226 = v2196
		v2227 = v2197
		goto L492
	case 6:
		v2220 = v2196
		v2221 = v2197
		goto L493
	case 7:
		v2215 = v2197
		goto L494
	case 8:
		v2210 = v2197
		goto L495
	case 9:
		v2205 = v2197
		goto L496
	case 10:
		goto L497
	default:
		v2263 = v2195
		v2264 = v2196
		v2265 = v2197
		goto L459
	}
L461:
	;
	v2145 = v2006
	v2146 = v2030
	v2147 = v2036
	v2148 = v2036
	v2149 = v2036
	goto L484
L462:
	;
	if base.Ui32(int32(11)) < base.Ui32(v2030) {
		goto L461
	} else {
		goto L465
	}
L463:
	;
	goto L464
L464:
	;
	if base.Ui32(v2030) < base.Ui32(int32(12)) {
		goto L467
	} else {
		goto L468
	}
L465:
	;
	v2193 = v2006
	v2194 = v2030
	v2195 = v2036
	v2196 = v2036
	v2197 = v2036
	goto L460
L466:
	;
	switch v2092 - int32(1) {
	case 0:
		v2142 = v2093
		goto L473
	case 1:
		v2137 = v2093
		goto L474
	case 2:
		goto L475
	case 3:
		v2130 = v2094
		goto L476
	case 4:
		v2127 = v2094
		goto L477
	case 5:
		v2122 = v2094
		goto L478
	case 6:
		goto L479
	case 7:
		v2113 = v2095
		goto L480
	case 8:
		v2108 = v2095
		goto L481
	case 9:
		v2103 = v2095
		goto L482
	case 10:
		goto L483
	default:
		v2263 = v2093
		v2264 = v2094
		v2265 = v2095
		goto L459
	}
L467:
	;
	v2091 = v2006
	v2092 = v2030
	v2093 = v2036
	v2094 = v2036
	v2095 = v2036
	goto L466
L468:
	;
	goto L469
L469:
	;
	v2043 = v2006
	v2044 = v2030
	v2045 = v2036
	v2046 = v2036
	v2047 = v2036
	goto L470
L470:
	;
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v2043)+4))
	v2050 = v2049 + v2046
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v2043)))
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v2043)+8))
	v2054 = v2053 + v2047
	v2056 = int32(4)
	v2058 = v2051 + v2045 - v2054 ^ base.I32_rotl(v2054, v2056)
	v2062 = v2050 - v2058 ^ base.I32_rotl(v2058, int32(6))
	v2063 = v2054 + v2050
	v2064 = v2058 + v2063
	v2065 = v2062 + v2064
	v2069 = v2063 - v2062 ^ base.I32_rotl(v2062, int32(8))
	v2073 = v2064 - v2069 ^ base.I32_rotl(v2069, int32(16))
	v2077 = v2065 - v2073 ^ base.I32_rotl(v2073, int32(19))
	v2078 = v2069 + v2065
	v2079 = v2073 + v2078
	v2080 = v2077 + v2079
	v2084 = v2078 - v2077 ^ base.I32_rotl(v2077, v2056)
	v2085 = int32(12)
	v2086 = v2043 + v2085
	v2088 = v2044 - v2085
	if base.Ui32(int32(11)) < base.Ui32(v2088) {
		v2043 = v2086
		v2044 = v2088
		v2045 = v2079
		v2046 = v2080
		v2047 = v2084
		goto L470
	} else {
		goto L472
	}
L471:
	;
	v2091 = v2086
	v2092 = v2088
	v2093 = v2079
	v2094 = v2080
	v2095 = v2084
	goto L466
L472:
	;
	goto L471
L473:
	;
	v2143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2091))))
	v2263 = v2142 + v2143
	v2264 = v2094
	v2265 = v2095
	goto L459
L474:
	;
	v2138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2091)+1)))
	v2142 = v2138<<(uint(int32(8))%32) + v2137
	goto L473
L475:
	;
	v2133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2091)+2)))
	v2137 = v2133<<(uint(int32(16))%32) + v2093
	goto L474
L476:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2091)))
	v2263 = v2131 + v2093
	v2264 = v2130
	v2265 = v2095
	goto L459
L477:
	;
	v2128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2091)+4)))
	v2130 = v2127 + v2128
	goto L476
L478:
	;
	v2123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2091)+5)))
	v2127 = v2123<<(uint(int32(8))%32) + v2122
	goto L477
L479:
	;
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2091)+6)))
	v2122 = v2118<<(uint(int32(16))%32) + v2094
	goto L478
L480:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v2091)))
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v2091)+4))
	v2263 = v2114 + v2093
	v2264 = v2116 + v2094
	v2265 = v2113
	goto L459
L481:
	;
	v2109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2091)+8)))
	v2113 = v2109<<(uint(int32(8))%32) + v2108
	goto L480
L482:
	;
	v2104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2091)+9)))
	v2108 = v2104<<(uint(int32(16))%32) + v2103
	goto L481
L483:
	;
	v2099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2091)+10)))
	v2103 = v2099<<(uint(int32(24))%32) + v2095
	goto L482
L484:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(v2145)+4))
	v2152 = v2151 + v2148
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v2145)))
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(v2145)+8))
	v2156 = v2155 + v2149
	v2158 = int32(4)
	v2160 = v2153 + v2147 - v2156 ^ base.I32_rotl(v2156, v2158)
	v2164 = v2152 - v2160 ^ base.I32_rotl(v2160, int32(6))
	v2165 = v2156 + v2152
	v2166 = v2160 + v2165
	v2167 = v2164 + v2166
	v2171 = v2165 - v2164 ^ base.I32_rotl(v2164, int32(8))
	v2175 = v2166 - v2171 ^ base.I32_rotl(v2171, int32(16))
	v2179 = v2167 - v2175 ^ base.I32_rotl(v2175, int32(19))
	v2180 = v2171 + v2167
	v2181 = v2175 + v2180
	v2182 = v2179 + v2181
	v2186 = v2180 - v2179 ^ base.I32_rotl(v2179, v2158)
	v2187 = int32(12)
	v2188 = v2145 + v2187
	v2190 = v2146 - v2187
	if base.Ui32(int32(11)) < base.Ui32(v2190) {
		v2145 = v2188
		v2146 = v2190
		v2147 = v2181
		v2148 = v2182
		v2149 = v2186
		goto L484
	} else {
		goto L486
	}
L485:
	;
	v2193 = v2188
	v2194 = v2190
	v2195 = v2181
	v2196 = v2182
	v2197 = v2186
	goto L460
L486:
	;
	goto L485
L487:
	;
	v2259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193))))
	v2263 = v2256 + v2259
	v2264 = v2257
	v2265 = v2258
	goto L459
L488:
	;
	v2252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193)+1)))
	v2256 = v2252<<(uint(int32(8))%32) + v2249
	v2257 = v2250
	v2258 = v2251
	goto L487
L489:
	;
	v2245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193)+2)))
	v2249 = v2245<<(uint(int32(16))%32) + v2242
	v2250 = v2243
	v2251 = v2244
	goto L488
L490:
	;
	v2238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193)+3)))
	v2242 = v2238<<(uint(int32(24))%32) + v2195
	v2243 = v2236
	v2244 = v2237
	goto L489
L491:
	;
	v2234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193)+4)))
	v2236 = v2232 + v2234
	v2237 = v2233
	goto L490
L492:
	;
	v2228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193)+5)))
	v2232 = v2228<<(uint(int32(8))%32) + v2226
	v2233 = v2227
	goto L491
L493:
	;
	v2222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193)+6)))
	v2226 = v2222<<(uint(int32(16))%32) + v2220
	v2227 = v2221
	goto L492
L494:
	;
	v2216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193)+7)))
	v2220 = v2216<<(uint(int32(24))%32) + v2196
	v2221 = v2215
	goto L493
L495:
	;
	v2211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193)+8)))
	v2215 = v2211<<(uint(int32(8))%32) + v2210
	goto L494
L496:
	;
	v2206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193)+9)))
	v2210 = v2206<<(uint(int32(16))%32) + v2205
	goto L495
L497:
	;
	v2201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193)+10)))
	v2205 = v2201<<(uint(int32(24))%32) + v2197
	goto L496
L498:
	;
	v2889 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+64)) = v2889
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+60)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+56)) = v2007
	v2893 = *(*int32)(unsafe.Add(mBase, uint32(v327)+28))
	v2895 = v2019 + int32(56)
	v2898 = v2019 + int32(52)
	v2902 = m.G0
	v2903 = int32(16)
	v2904 = v2902 - v2903
	m.G0 = v2904
	v2906 = *(*int32)(unsafe.Add(mBase, uint32(v2895)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2904)+8)) = v2906
	v2908 = *(*int64)(unsafe.Add(mBase, uint32(v2895)))
	*(*int64)(unsafe.Add(mBase, uint32(v2904))) = v2908
	v2910 = *(*int32)(unsafe.Add(mBase, uint32(v2893)))
	*(*int32)(unsafe.Add(mBase, uint32(v2904)+12)) = v2889
	v2913 = F_hash_bytes(m, v2904, v2903)
	mBase = m.M
	v2914 = *(*int32)(unsafe.Add(mBase, uint32(v2910)+20))
	v2915 = *(*int32)(unsafe.Add(mBase, uint32(v2910)+12))
	v2916 = v2913 & v2915
	v2919 = v2914 + v2916*int32(40)
	v2920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2919)+20)))
	if v2920 == v2889 {
		goto L580
	} else {
		goto L581
	}
L499:
	;
	v2311 = v2297
	goto L502
L500:
	;
	goto L501
L501:
	;
	v2434 = v2019 + int32(56)
	F_GetRelationPath(m, v2434, v347, v2007, v2008, int32(-1), v2009)
	mBase = m.M
	v2437 = m.ExcPending
	if v2437 != 0 {
		goto L4
	} else {
		goto L513
	}
L502:
	;
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(v2295+v2311<<(uint(int32(4))%32))+4))
	v2354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2351))))
	v2357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2006))))
	if base.B2i32(v2354 == int32(0))|base.B2i32(v2354 != v2357) != 0 {
		v2375 = v2354
		v2376 = v2357
		goto L505
	} else {
		goto L506
	}
L503:
	;
	goto L501
L504:
	;
	if v2375-v2376 == int32(0) {
		goto L498
	} else {
		goto L511
	}
L505:
	;
	goto L504
L506:
	;
	v2360 = v2351
	v2361 = v2006
	goto L507
L507:
	;
	v2364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2361)+1)))
	v2365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2360)+1)))
	if v2365 == int32(0) {
		v2375 = v2365
		v2376 = v2364
		goto L505
	} else {
		goto L509
	}
L508:
	;
	v2375 = v2365
	v2376 = v2364
	goto L505
L509:
	;
	v2368 = int32(1)
	if v2365 == v2364 {
		v2360 = v2360 + v2368
		v2361 = v2361 + v2368
		goto L507
	} else {
		goto L510
	}
L510:
	;
	goto L508
L511:
	;
	v2382 = (v2311 + int32(1)) & v2296
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2295+v2382<<(uint(int32(4))%32))))
	if v2386 != 0 {
		v2311 = v2382
		goto L502
	} else {
		goto L512
	}
L512:
	;
	goto L503
L513:
	;
	v2441 = F_strlen(m, v2434)
	mBase = m.M
	v2448 = v2441 + int32(1)
	goto L516
L514:
	;
	v2461 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2460))) = uint8(v2461)
	v2464 = v2460 + int32(1)
	if v2010 != 0 {
		goto L521
	} else {
		goto L522
	}
L515:
	;
	goto L514
L516:
	;
	v2450 = int32(0)
	if v2448 == v2450 {
		v2460 = v2450
		goto L515
	} else {
		goto L518
	}
L517:
	;
	v2460 = v2455
	goto L515
L518:
	;
	v2454 = v2448 - int32(1)
	v2455 = v2434 + v2454
	v2456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2455))))
	if v2456 != int32(47) {
		v2448 = v2454
		goto L516
	} else {
		goto L519
	}
L519:
	;
	goto L517
L520:
	;
	v2483 = *(*int32)(unsafe.Add(mBase, uint32(v327)+24))
	v2484 = F_strlen(m, v2482)
	mBase = m.M
	v2490 = v2484 - int32(1636608432)
	if v2482&int32(3) != 0 {
		goto L530
	} else {
		goto L531
	}
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+40)) = v2010
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+36)) = v2464
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+32)) = v2434
	v2471 = F_psprintf(m, int32(_a_F_sendDir_46), v2019+int32(32))
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L4
	} else {
		goto L524
	}
L522:
	;
	goto L523
L523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+20)) = v2464
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+16)) = v2019 + int32(56)
	v2480 = F_psprintf(m, int32(_a_F_sendDir_47), v2019+int32(16))
	mBase = m.M
	v2481 = m.ExcPending
	if v2481 != 0 {
		goto L4
	} else {
		goto L525
	}
L524:
	;
	v2482 = v2471
	goto L520
L525:
	;
	v2482 = v2480
	goto L520
L526:
	;
	v2749 = *(*int32)(unsafe.Add(mBase, uint32(v2483)+20))
	v2750 = *(*int32)(unsafe.Add(mBase, uint32(v2483)+12))
	v2751 = (v2744 ^ v2736 - base.I32_rotl(v2744, int32(24))) & v2750
	v2755 = *(*int32)(unsafe.Add(mBase, uint32(v2749+v2751<<(uint(int32(4))%32))))
	if v2755 == int32(0) {
		v3700 = v2016
		goto L456
	} else {
		goto L566
	}
L527:
	;
	v2722 = int32(14)
	v2724 = v2718 ^ v2719 - base.I32_rotl(v2718, v2722)
	v2728 = v2724 ^ v2717 - base.I32_rotl(v2724, int32(11))
	v2732 = v2728 ^ v2718 - base.I32_rotl(v2728, int32(25))
	v2736 = v2732 ^ v2724 - base.I32_rotl(v2732, int32(16))
	v2740 = v2736 ^ v2728 - base.I32_rotl(v2736, int32(4))
	v2744 = v2740 ^ v2732 - base.I32_rotl(v2740, v2722)
	goto L526
L528:
	;
	switch v2648 - int32(1) {
	case 0:
		v2710 = v2649
		v2711 = v2650
		v2712 = v2651
		goto L555
	case 1:
		v2703 = v2649
		v2704 = v2650
		v2705 = v2651
		goto L556
	case 2:
		v2696 = v2649
		v2697 = v2650
		v2698 = v2651
		goto L557
	case 3:
		v2690 = v2650
		v2691 = v2651
		goto L558
	case 4:
		v2686 = v2650
		v2687 = v2651
		goto L559
	case 5:
		v2680 = v2650
		v2681 = v2651
		goto L560
	case 6:
		v2674 = v2650
		v2675 = v2651
		goto L561
	case 7:
		v2669 = v2651
		goto L562
	case 8:
		v2664 = v2651
		goto L563
	case 9:
		v2659 = v2651
		goto L564
	case 10:
		goto L565
	default:
		v2717 = v2649
		v2718 = v2650
		v2719 = v2651
		goto L527
	}
L529:
	;
	v2599 = v2482
	v2600 = v2484
	v2601 = v2490
	v2602 = v2490
	v2603 = v2490
	goto L552
L530:
	;
	if base.Ui32(int32(11)) < base.Ui32(v2484) {
		goto L529
	} else {
		goto L533
	}
L531:
	;
	goto L532
L532:
	;
	if base.Ui32(v2484) < base.Ui32(int32(12)) {
		goto L535
	} else {
		goto L536
	}
L533:
	;
	v2647 = v2482
	v2648 = v2484
	v2649 = v2490
	v2650 = v2490
	v2651 = v2490
	goto L528
L534:
	;
	switch v2546 - int32(1) {
	case 0:
		v2596 = v2547
		goto L541
	case 1:
		v2591 = v2547
		goto L542
	case 2:
		goto L543
	case 3:
		v2584 = v2548
		goto L544
	case 4:
		v2581 = v2548
		goto L545
	case 5:
		v2576 = v2548
		goto L546
	case 6:
		goto L547
	case 7:
		v2567 = v2549
		goto L548
	case 8:
		v2562 = v2549
		goto L549
	case 9:
		v2557 = v2549
		goto L550
	case 10:
		goto L551
	default:
		v2717 = v2547
		v2718 = v2548
		v2719 = v2549
		goto L527
	}
L535:
	;
	v2545 = v2482
	v2546 = v2484
	v2547 = v2490
	v2548 = v2490
	v2549 = v2490
	goto L534
L536:
	;
	goto L537
L537:
	;
	v2497 = v2482
	v2498 = v2484
	v2499 = v2490
	v2500 = v2490
	v2501 = v2490
	goto L538
L538:
	;
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(v2497)+4))
	v2504 = v2503 + v2500
	v2505 = *(*int32)(unsafe.Add(mBase, uint32(v2497)))
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v2497)+8))
	v2508 = v2507 + v2501
	v2510 = int32(4)
	v2512 = v2505 + v2499 - v2508 ^ base.I32_rotl(v2508, v2510)
	v2516 = v2504 - v2512 ^ base.I32_rotl(v2512, int32(6))
	v2517 = v2508 + v2504
	v2518 = v2512 + v2517
	v2519 = v2516 + v2518
	v2523 = v2517 - v2516 ^ base.I32_rotl(v2516, int32(8))
	v2527 = v2518 - v2523 ^ base.I32_rotl(v2523, int32(16))
	v2531 = v2519 - v2527 ^ base.I32_rotl(v2527, int32(19))
	v2532 = v2523 + v2519
	v2533 = v2527 + v2532
	v2534 = v2531 + v2533
	v2538 = v2532 - v2531 ^ base.I32_rotl(v2531, v2510)
	v2539 = int32(12)
	v2540 = v2497 + v2539
	v2542 = v2498 - v2539
	if base.Ui32(int32(11)) < base.Ui32(v2542) {
		v2497 = v2540
		v2498 = v2542
		v2499 = v2533
		v2500 = v2534
		v2501 = v2538
		goto L538
	} else {
		goto L540
	}
L539:
	;
	v2545 = v2540
	v2546 = v2542
	v2547 = v2533
	v2548 = v2534
	v2549 = v2538
	goto L534
L540:
	;
	goto L539
L541:
	;
	v2597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2545))))
	v2717 = v2596 + v2597
	v2718 = v2548
	v2719 = v2549
	goto L527
L542:
	;
	v2592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2545)+1)))
	v2596 = v2592<<(uint(int32(8))%32) + v2591
	goto L541
L543:
	;
	v2587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2545)+2)))
	v2591 = v2587<<(uint(int32(16))%32) + v2547
	goto L542
L544:
	;
	v2585 = *(*int32)(unsafe.Add(mBase, uint32(v2545)))
	v2717 = v2585 + v2547
	v2718 = v2584
	v2719 = v2549
	goto L527
L545:
	;
	v2582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2545)+4)))
	v2584 = v2581 + v2582
	goto L544
L546:
	;
	v2577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2545)+5)))
	v2581 = v2577<<(uint(int32(8))%32) + v2576
	goto L545
L547:
	;
	v2572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2545)+6)))
	v2576 = v2572<<(uint(int32(16))%32) + v2548
	goto L546
L548:
	;
	v2568 = *(*int32)(unsafe.Add(mBase, uint32(v2545)))
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v2545)+4))
	v2717 = v2568 + v2547
	v2718 = v2570 + v2548
	v2719 = v2567
	goto L527
L549:
	;
	v2563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2545)+8)))
	v2567 = v2563<<(uint(int32(8))%32) + v2562
	goto L548
L550:
	;
	v2558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2545)+9)))
	v2562 = v2558<<(uint(int32(16))%32) + v2557
	goto L549
L551:
	;
	v2553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2545)+10)))
	v2557 = v2553<<(uint(int32(24))%32) + v2549
	goto L550
L552:
	;
	v2605 = *(*int32)(unsafe.Add(mBase, uint32(v2599)+4))
	v2606 = v2605 + v2602
	v2607 = *(*int32)(unsafe.Add(mBase, uint32(v2599)))
	v2609 = *(*int32)(unsafe.Add(mBase, uint32(v2599)+8))
	v2610 = v2609 + v2603
	v2612 = int32(4)
	v2614 = v2607 + v2601 - v2610 ^ base.I32_rotl(v2610, v2612)
	v2618 = v2606 - v2614 ^ base.I32_rotl(v2614, int32(6))
	v2619 = v2610 + v2606
	v2620 = v2614 + v2619
	v2621 = v2618 + v2620
	v2625 = v2619 - v2618 ^ base.I32_rotl(v2618, int32(8))
	v2629 = v2620 - v2625 ^ base.I32_rotl(v2625, int32(16))
	v2633 = v2621 - v2629 ^ base.I32_rotl(v2629, int32(19))
	v2634 = v2625 + v2621
	v2635 = v2629 + v2634
	v2636 = v2633 + v2635
	v2640 = v2634 - v2633 ^ base.I32_rotl(v2633, v2612)
	v2641 = int32(12)
	v2642 = v2599 + v2641
	v2644 = v2600 - v2641
	if base.Ui32(int32(11)) < base.Ui32(v2644) {
		v2599 = v2642
		v2600 = v2644
		v2601 = v2635
		v2602 = v2636
		v2603 = v2640
		goto L552
	} else {
		goto L554
	}
L553:
	;
	v2647 = v2642
	v2648 = v2644
	v2649 = v2635
	v2650 = v2636
	v2651 = v2640
	goto L528
L554:
	;
	goto L553
L555:
	;
	v2713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647))))
	v2717 = v2710 + v2713
	v2718 = v2711
	v2719 = v2712
	goto L527
L556:
	;
	v2706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+1)))
	v2710 = v2706<<(uint(int32(8))%32) + v2703
	v2711 = v2704
	v2712 = v2705
	goto L555
L557:
	;
	v2699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+2)))
	v2703 = v2699<<(uint(int32(16))%32) + v2696
	v2704 = v2697
	v2705 = v2698
	goto L556
L558:
	;
	v2692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+3)))
	v2696 = v2692<<(uint(int32(24))%32) + v2649
	v2697 = v2690
	v2698 = v2691
	goto L557
L559:
	;
	v2688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+4)))
	v2690 = v2686 + v2688
	v2691 = v2687
	goto L558
L560:
	;
	v2682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+5)))
	v2686 = v2682<<(uint(int32(8))%32) + v2680
	v2687 = v2681
	goto L559
L561:
	;
	v2676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+6)))
	v2680 = v2676<<(uint(int32(16))%32) + v2674
	v2681 = v2675
	goto L560
L562:
	;
	v2670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+7)))
	v2674 = v2670<<(uint(int32(24))%32) + v2650
	v2675 = v2669
	goto L561
L563:
	;
	v2665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+8)))
	v2669 = v2665<<(uint(int32(8))%32) + v2664
	goto L562
L564:
	;
	v2660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+9)))
	v2664 = v2660<<(uint(int32(16))%32) + v2659
	goto L563
L565:
	;
	v2655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+10)))
	v2659 = v2655<<(uint(int32(24))%32) + v2651
	goto L564
L566:
	;
	v2767 = v2751
	goto L567
L567:
	;
	v2807 = *(*int32)(unsafe.Add(mBase, uint32(v2749+v2767<<(uint(int32(4))%32))+4))
	v2810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2807))))
	v2813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2482))))
	if base.B2i32(v2810 == int32(0))|base.B2i32(v2810 != v2813) != 0 {
		v2831 = v2810
		v2832 = v2813
		goto L570
	} else {
		goto L571
	}
L568:
	;
	v3700 = v2016
	goto L456
L569:
	;
	if v2831-v2832 == int32(0) {
		goto L498
	} else {
		goto L576
	}
L570:
	;
	goto L569
L571:
	;
	v2816 = v2807
	v2817 = v2482
	goto L572
L572:
	;
	v2820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2817)+1)))
	v2821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+1)))
	if v2821 == int32(0) {
		v2831 = v2821
		v2832 = v2820
		goto L570
	} else {
		goto L574
	}
L573:
	;
	v2831 = v2821
	v2832 = v2820
	goto L570
L574:
	;
	v2824 = int32(1)
	if v2821 == v2820 {
		v2816 = v2816 + v2824
		v2817 = v2817 + v2824
		goto L572
	} else {
		goto L575
	}
L575:
	;
	goto L573
L576:
	;
	v2838 = (v2767 + int32(1)) & v2750
	v2842 = *(*int32)(unsafe.Add(mBase, uint32(v2749+v2838<<(uint(int32(4))%32))))
	if v2842 != 0 {
		v2767 = v2838
		goto L567
	} else {
		goto L577
	}
L577:
	;
	goto L568
L578:
	;
	if v2957 != 0 {
		v3700 = v2016
		goto L456
	} else {
		goto L588
	}
L579:
	;
	m.G0 = v2904 + int32(16)
	goto L578
L580:
	;
	v2957 = int32(0)
	goto L579
L581:
	;
	v2923 = *(*int64)(unsafe.Add(mBase, uint32(v2904)))
	v2924 = v2916
	v2925 = v2919
	goto L582
L582:
	;
	v2931 = *(*int64)(unsafe.Add(mBase, uint32(v2925)))
	v2933 = *(*int64)(unsafe.Add(mBase, uint32(v2925)+8))
	v2934 = *(*int64)(unsafe.Add(mBase, uint32(v2904)+8))
	if v2931^v2923|(v2933^v2934) != int64(0) {
		goto L584
	} else {
		goto L585
	}
L583:
	;
	v2946 = *(*int32)(unsafe.Add(mBase, uint32(v2925)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2898))) = v2946
	v2957 = v2925
	goto L579
L584:
	;
	v2941 = (v2924 + int32(1)) & v2915
	v2944 = v2914 + v2941*int32(40)
	v2945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2944)+20)))
	if v2945 != 0 {
		v2924 = v2941
		v2925 = v2944
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
	v2967 = int32(base.Ui32(v2011) >> (uint(int32(13)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+64)) = v2008
	v2969 = *(*int32)(unsafe.Add(mBase, uint32(v327)+28))
	v2973 = m.G0
	v2974 = int32(16)
	v2975 = v2973 - v2974
	m.G0 = v2975
	v2977 = *(*int32)(unsafe.Add(mBase, uint32(v2895)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2975)+8)) = v2977
	v2979 = *(*int64)(unsafe.Add(mBase, uint32(v2895)))
	*(*int64)(unsafe.Add(mBase, uint32(v2975))) = v2979
	v2981 = *(*int32)(unsafe.Add(mBase, uint32(v2969)))
	*(*int32)(unsafe.Add(mBase, uint32(v2975)+12)) = v2009
	v2984 = F_hash_bytes(m, v2975, v2974)
	mBase = m.M
	v2985 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+20))
	v2986 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+12))
	v2987 = v2984 & v2986
	v2990 = v2985 + v2987*int32(40)
	v2991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2990)+20)))
	if v2991 == int32(0) {
		goto L592
	} else {
		goto L593
	}
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2015))) = v3639
	v3700 = int32(1)
	goto L456
L590:
	;
	if v3028 == int32(0) {
		goto L600
	} else {
		goto L601
	}
L591:
	;
	m.G0 = v2975 + int32(16)
	goto L590
L592:
	;
	v3028 = int32(0)
	goto L591
L593:
	;
	v2994 = *(*int64)(unsafe.Add(mBase, uint32(v2975)))
	v2995 = v2987
	v2996 = v2990
	goto L594
L594:
	;
	v3002 = *(*int64)(unsafe.Add(mBase, uint32(v2996)))
	v3004 = *(*int64)(unsafe.Add(mBase, uint32(v2996)+8))
	v3005 = *(*int64)(unsafe.Add(mBase, uint32(v2975)+8))
	if v3002^v2994|(v3004^v3005) != int64(0) {
		goto L596
	} else {
		goto L597
	}
L595:
	;
	v3017 = *(*int32)(unsafe.Add(mBase, uint32(v2996)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2898))) = v3017
	v3028 = v2996
	goto L591
L596:
	;
	v3012 = (v2995 + int32(1)) & v2986
	v3015 = v2985 + v3012*int32(40)
	v3016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3015)+20)))
	if v3016 != 0 {
		v2995 = v3012
		v2996 = v3015
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
	if v2011 == int32(0) {
		v3700 = v2016
		goto L456
	} else {
		goto L603
	}
L601:
	;
	goto L602
L602:
	;
	v3044 = v2010 << (uint(int32(17)) % 32)
	v3045 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+52))
	if base.Ui32(v3045) <= base.Ui32(v3044) {
		v3700 = v2016
		goto L456
	} else {
		goto L604
	}
L603:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2013))) = int32(0)
	v3639 = v2967
	goto L589
L604:
	;
	if base.Ui32(int32(_a_F_sendDir_48)) < base.Ui32(v2010) {
		goto L455
	} else {
		goto L605
	}
L605:
	;
	v3049 = v2967 + v3044
	if base.Ui32(v3049) < base.Ui32(v2967) {
		goto L455
	} else {
		goto L606
	}
L606:
	;
	v3051 = int32(0)
	v3052 = int32(16)
	v3053 = int32(base.Ui32(v3044) >> (uint(v3052) % 32))
	v3060 = int32(base.Ui32(v3049)>>(uint(v3052)%32)) + base.B2i32(v3049&int32(_a_F_sendDir_49) != v3051)
	v3061 = *(*int32)(unsafe.Add(mBase, uint32(v3028)+24))
	if base.Ui32(v3060) < base.Ui32(v3061) {
		goto L608
	} else {
		goto L609
	}
L607:
	;
	if base.F64_gt(base.F64_convert_i32_u(v3339<<(uint(int32(13))%32)), base.F64_mul(base.F64_convert_i32_u(v2011), float64(0.9))) != 0 {
		v3700 = v2016
		goto L456
	} else {
		goto L641
	}
L608:
	;
	v3063 = v3060
	goto L610
L609:
	;
	v3063 = v3061
	goto L610
L610:
	;
	if base.Ui32(v3063) <= base.Ui32(v3053) {
		v3339 = v3051
		goto L607
	} else {
		goto L611
	}
L611:
	;
	v3078 = v3053
	v3081 = v3051
	goto L612
L612:
	;
	v3115 = *(*int32)(unsafe.Add(mBase, uint32(v3028)+32))
	v3119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3115+v3078<<(uint(int32(1))%32)))))
	if v3119 == int32(0) {
		v3290 = v3081
		goto L614
	} else {
		goto L615
	}
L613:
	;
	v3339 = v3290
	goto L607
L614:
	;
	v3325 = v3078 + int32(1)
	if v3325 != v3063 {
		v3078 = v3325
		v3081 = v3290
		goto L612
	} else {
		goto L640
	}
L615:
	;
	v3124 = v3078 << (uint(int32(16)) % 32)
	if v3078 != v3063-int32(1) {
		goto L616
	} else {
		goto L617
	}
L616:
	;
	v3127 = int32(_a_F_sendDir_50)
	goto L618
L617:
	;
	v3127 = v3049 - v3124
	goto L618
L618:
	;
	if v3078 == v3053 {
		goto L619
	} else {
		goto L620
	}
L619:
	;
	v3130 = v3044 & int32(_a_F_sendDir_49)
	goto L621
L620:
	;
	v3130 = int32(0)
	goto L621
L621:
	;
	v3131 = *(*int32)(unsafe.Add(mBase, uint32(v3028)+36))
	v3135 = *(*int32)(unsafe.Add(mBase, uint32(v3131+v3078<<(uint(int32(2))%32))))
	if v3119 != int32(_a_F_sendDir_51) {
		goto L622
	} else {
		goto L623
	}
L622:
	;
	v3151 = v3081
	v3170 = int32(0)
	goto L625
L623:
	;
	goto L624
L624:
	;
	if base.Ui32(v3127) <= base.Ui32(v3130) {
		v3290 = v3081
		goto L614
	} else {
		goto L632
	}
L625:
	;
	v3188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3135+v3170<<(uint(int32(1))%32)))))
	if base.B2i32(base.Ui32(v3188) < base.Ui32(v3130))|base.B2i32(base.Ui32(v3127) <= base.Ui32(v3188)) == int32(0) {
		goto L627
	} else {
		goto L628
	}
L626:
	;
	v3290 = v3203
	goto L614
L627:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338+v3151<<(uint(int32(2))%32)))) = v3124 | v3188
	v3200 = v3151 + int32(1)
	if v3200 == int32(_a_F_sendDir_52) {
		v3339 = v3200
		goto L607
	} else {
		goto L630
	}
L628:
	;
	v3203 = v3151
	goto L629
L629:
	;
	v3205 = v3170 + int32(1)
	if v3205 != v3119 {
		v3151 = v3203
		v3170 = v3205
		goto L625
	} else {
		goto L631
	}
L630:
	;
	v3203 = v3200
	goto L629
L631:
	;
	goto L626
L632:
	;
	v3220 = v3081
	v3223 = v3130
	goto L633
L633:
	;
	v3259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3135+int32(base.Ui32(v3223)>>(uint(int32(3))%32))&int32(536870910)))))
	if int32(base.Ui32(v3259)>>(uint(v3223&int32(15))%32))&int32(1) != 0 {
		goto L635
	} else {
		goto L636
	}
L634:
	;
	v3290 = v3274
	goto L614
L635:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338+v3220<<(uint(int32(2))%32)))) = v3223 + v3124
	v3271 = v3220 + int32(1)
	if v3271 == int32(_a_F_sendDir_52) {
		v3339 = v3271
		goto L607
	} else {
		goto L638
	}
L636:
	;
	v3274 = v3220
	goto L637
L637:
	;
	v3276 = v3223 + int32(1)
	if v3276 != v3127 {
		v3220 = v3274
		v3223 = v3276
		goto L633
	} else {
		goto L639
	}
L638:
	;
	v3274 = v3271
	goto L637
L639:
	;
	goto L634
L640:
	;
	goto L613
L641:
	;
	F_pg_qsort(m, v338, v3339, int32(4), int32(434))
	mBase = m.M
	v3383 = m.ExcPending
	if v3383 != 0 {
		goto L4
	} else {
		goto L642
	}
L642:
	;
	v3384 = int32(0)
	if base.B2i32(v3044 == v3384)|base.B2i32(v3339 == v3384) != 0 {
		goto L643
	} else {
		goto L644
	}
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2013))) = v3339
	*(*int32)(unsafe.Add(mBase, uint32(v2015))) = v2967
	v3616 = int32(1)
	v3617 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+52))
	if v3617 == int32(-1) {
		v3700 = v3616
		goto L456
	} else {
		goto L655
	}
L644:
	;
	v3390 = v3339 & int32(3)
	v3391 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v3339) {
		goto L645
	} else {
		goto L646
	}
L645:
	;
	v3412 = v3391
	v3419 = v2016
	goto L648
L646:
	;
	v3480 = v3391
	goto L647
L647:
	;
	v3526 = v3480
	v3529 = v3391
	goto L652
L648:
	;
	v3445 = v338 + v3412<<(uint(int32(2))%32)
	v3446 = *(*int32)(unsafe.Add(mBase, uint32(v3445)))
	*(*int32)(unsafe.Add(mBase, uint32(v3445))) = v3446 - v3044
	v3449 = *(*int32)(unsafe.Add(mBase, uint32(v3445)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3445)+4)) = v3449 - v3044
	v3452 = *(*int32)(unsafe.Add(mBase, uint32(v3445)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3445)+8)) = v3452 - v3044
	v3455 = *(*int32)(unsafe.Add(mBase, uint32(v3445)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3445)+12)) = v3455 - v3044
	v3458 = int32(4)
	v3459 = v3412 + v3458
	v3461 = v3419 + v3458
	if v3461 != v3339&int32(-4) {
		v3412 = v3459
		v3419 = v3461
		goto L648
	} else {
		goto L650
	}
L649:
	;
	if v3390 == int32(0) {
		goto L643
	} else {
		goto L651
	}
L650:
	;
	goto L649
L651:
	;
	v3480 = v3459
	goto L647
L652:
	;
	v3559 = v338 + v3526<<(uint(int32(2))%32)
	v3560 = *(*int32)(unsafe.Add(mBase, uint32(v3559)))
	*(*int32)(unsafe.Add(mBase, uint32(v3559))) = v3560 - v3044
	v3563 = int32(1)
	v3566 = v3529 + v3563
	if v3566 != v3390 {
		v3526 = v3526 + v3563
		v3529 = v3566
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
	v3621 = v3617 - v3044
	if base.Ui32(v3621) < base.Ui32(v2967) {
		goto L656
	} else {
		goto L657
	}
L656:
	;
	v3623 = v2967
	goto L658
L657:
	;
	v3623 = v3621
	goto L658
L658:
	;
	if base.Ui32(int32(_a_F_sendDir_52)) <= base.Ui32(v3623) {
		goto L659
	} else {
		goto L660
	}
L659:
	;
	v3626 = int32(_a_F_sendDir_52)
	goto L661
L660:
	;
	v3626 = v3623
	goto L661
L661:
	;
	if base.Ui32(v2967) < base.Ui32(v3621) {
		v3639 = v3626
		goto L589
	} else {
		goto L662
	}
L662:
	;
	if base.Ui32(v3623) < base.Ui32(int32(_a_F_sendDir_53)) {
		v3700 = v3616
		goto L456
	} else {
		goto L663
	}
L663:
	;
	v3639 = v3626
	goto L589
L664:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v3734 = m.ExcPending
	if v3734 != 0 {
		goto L4
	} else {
		goto L665
	}
L665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+4)) = v2011
	*(*int32)(unsafe.Add(mBase, uint32(v2019))) = v2010
	F_errmsg_internal(m, int32(_a_F_sendDir_54), v2019)
	mBase = m.M
	v3739 = m.ExcPending
	if v3739 != 0 {
		goto L4
	} else {
		goto L666
	}
L666:
	;
	F_errfinish(m, int32(_a_F_sendDir_55), int32(796), int32(_a_F_sendDir_56))
	mBase = m.M
	v3744 = m.ExcPending
	if v3744 != 0 {
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
	v3747 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2256))
	v3749 = v3747 << (uint(int32(2)) % 32)
	v3751 = v3749 + int32(12)
	if v3747 == int32(0) {
		v3761 = v3751
		goto L671
	} else {
		goto L672
	}
L669:
	;
	v3782 = v1983
	v3783 = v340
	goto L670
L670:
	;
	F_pfree(m, v2006)
	mBase = m.M
	v3785 = m.ExcPending
	if v3785 != 0 {
		goto L4
	} else {
		goto L675
	}
L671:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+96)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v329)+100)) = v384
	*(*int64)(unsafe.Add(mBase, uint32(v329)+2296)) = base.I64_extend_i32_u(v3761 + v3747<<(uint(int32(13))%32))
	v3771 = v329 + int32(192)
	v3776 = F_pg_snprintf(m, v3771, int32(2048), int32(_a_F_sendDir_47), v329+int32(96))
	mBase = m.M
	v3777 = m.ExcPending
	if v3777 != 0 {
		goto L4
	} else {
		goto L674
	}
L672:
	;
	v3755 = v3751 & int32(_a_F_sendDir_57)
	if v3755 == int32(0) {
		v3761 = v3751
		goto L671
	} else {
		goto L673
	}
L673:
	;
	v3761 = v3749 - v3755 + int32(_a_F_sendDir_58)
	goto L671
L674:
	;
	v3782 = v338
	v3783 = v3771
	goto L670
L675:
	;
	v3806 = v3782
	v3809 = v3783
	goto L447
L676:
	;
	v3839 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2268))
	v3840 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2260))
	v3841 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2256))
	v3842 = *(*int32)(unsafe.Add(mBase, uint32(v329)+2252))
	v3843 = F_sendFile(m, v319, v329+int32(2368), v3809, v329+int32(2272), int32(1), v347, v326, v3839, v3840, v325, v3841, v3806, v3842)
	mBase = m.M
	v3844 = m.ExcPending
	if v3844 != 0 {
		goto L4
	} else {
		goto L679
	}
L677:
	;
	goto L678
L678:
	;
	v3847 = *(*int64)(unsafe.Add(mBase, uint32(v329)+2296))
	v3922 = v363 + v3847 + ((v3847+int64(511))&int64(4294966784)-v3847)&int64(4294967295) + int64(512)
	goto L91
L679:
	;
	if v3843 == int32(0) {
		v3922 = v363
		goto L91
	} else {
		goto L680
	}
L680:
	;
	goto L678
L681:
	;
	if v3861 == int32(0) {
		v3922 = v363
		goto L91
	} else {
		goto L682
	}
L682:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+80)) = v329 + int32(2368)
	F_errmsg(m, int32(_a_F_sendDir_59), v329+int32(80))
	mBase = m.M
	v3872 = m.ExcPending
	if v3872 != 0 {
		goto L4
	} else {
		goto L683
	}
L683:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1546), int32(_a_F_sendDir_20))
	mBase = m.M
	v3877 = m.ExcPending
	if v3877 != 0 {
		goto L4
	} else {
		goto L684
	}
L684:
	;
	v3922 = v363
	goto L91
L685:
	;
	if v3924 != 0 {
		v328 = v3924
		v363 = v3922
		goto L89
	} else {
		goto L686
	}
L686:
	;
	goto L90
L687:
	;
	F_pfree(m, v3945)
	mBase = m.M
	v3973 = m.ExcPending
	if v3973 != 0 {
		goto L4
	} else {
		goto L690
	}
L688:
	;
	goto L689
L689:
	;
	F_FreeDir(m, v3958)
	mBase = m.M
	v3975 = m.ExcPending
	if v3975 != 0 {
		goto L4
	} else {
		goto L691
	}
L690:
	;
	goto L689
L691:
	;
	m.G0 = v3936 + int32(_a_F_sendDir_0)
	return v3970
L692:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3986 = m.ExcPending
	if v3986 != 0 {
		goto L4
	} else {
		goto L693
	}
L693:
	;
	F_errmsg(m, int32(_a_F_sendDir_60), int32(0))
	mBase = m.M
	v3990 = m.ExcPending
	if v3990 != 0 {
		goto L4
	} else {
		goto L694
	}
L694:
	;
	F_errhint(m, int32(_a_F_sendDir_61), int32(0))
	mBase = m.M
	v3994 = m.ExcPending
	if v3994 != 0 {
		goto L4
	} else {
		goto L695
	}
L695:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1286), int32(_a_F_sendDir_20))
	mBase = m.M
	v3999 = m.ExcPending
	if v3999 != 0 {
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
	v4005 = m.ExcPending
	if v4005 != 0 {
		goto L4
	} else {
		goto L698
	}
L698:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+48)) = v329 + int32(2368)
	F_errmsg(m, int32(_a_F_sendDir_62), v329+int32(48))
	mBase = m.M
	v4013 = m.ExcPending
	if v4013 != 0 {
		goto L4
	} else {
		goto L699
	}
L699:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1420), int32(_a_F_sendDir_20))
	mBase = m.M
	v4018 = m.ExcPending
	if v4018 != 0 {
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
	v4025 = m.ExcPending
	if v4025 != 0 {
		goto L4
	} else {
		goto L702
	}
L702:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+64)) = v329 + int32(2368)
	F_errmsg(m, int32(_a_F_sendDir_63), v329-int32(-64))
	mBase = m.M
	v4033 = m.ExcPending
	if v4033 != 0 {
		goto L4
	} else {
		goto L703
	}
L703:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1425), int32(_a_F_sendDir_20))
	mBase = m.M
	v4038 = m.ExcPending
	if v4038 != 0 {
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
func F_sendto(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	v7 = m.Env.X__syscall_sendto(m, l0, l1, l2, int32(0), l3, l4)
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
	v14 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
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
	*(*int64)(unsafe.Add(mBase, uint32(l1)+200)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+192)) = v80
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v6 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
		v20 = v6 + v7<<(uint(int32(2))%32)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
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
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v41
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
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
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
	goto L1
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
	if v15 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v13
	goto L1
L5:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v10
	goto L1
L6:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v7)
	goto L1
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v15 == v22 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v15 == v24 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
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
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v55 {
	case 0:
		goto L27
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	default:
		goto L22
	}
L22:
	;
	v68 = l0 + int32(56)
	goto L33
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v48 == v64 {
		goto L20
	} else {
		goto L32
	}
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v48 != v62 {
		goto L22
	} else {
		goto L31
	}
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v48 != v60 {
		goto L22
	} else {
		goto L30
	}
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v48 != v58 {
		goto L22
	} else {
		goto L29
	}
L27:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v48 != v56 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	goto L20
L29:
	;
	goto L20
L30:
	;
	goto L20
L31:
	;
	goto L20
L32:
	;
	goto L22
L33:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v72 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_pfree(m, v48)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L40
	}
L35:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+40))
	if v48 == v73 {
		goto L20
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	goto L34
L38:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+56))
	if v48 != v75 {
		v68 = v72
		goto L33
	} else {
		goto L39
	}
L39:
	;
	goto L20
L40:
	;
	goto L20
}
func F_setseed(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v63 int64
	_ = v63
	var v68 int64
	_ = v68
	var v73 int64
	_ = v73
	var v79 int32
	_ = v79
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v11 = base.F64_abs(v10)
	if base.B2i32(base.F64_gt(v11, float64(1)) == v2)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v11)) < base.Ui64(int64(9218868437227405313))) == v2 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v7))) = v10
				F_errmsg(m, int32(_a_F_setseed_0), v7)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_setseed_1), int32(70), int32(_a_F_setseed_2))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
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
		v40 = int32(_a_F_setseed_3)
		v45 = base.I64_trunc_sat_f64_s(base.F64_mul(v10, float64(4.503599627370495e+15)))
		v47 = v45 + int64(4354685564936845354)
		v48 = int64(30)
		v51 = int64(-4658895280553007687)
		v52 = (int64(base.Ui64(v47)>>(uint(v48)%64)) ^ v47) * v51
		v53 = int64(27)
		v56 = int64(-7723592293110705685)
		v57 = (int64(base.Ui64(v52)>>(uint(v53)%64)) ^ v52) * v56
		v58 = int64(31)
		*(*int64)(unsafe.Add(mBase, _c_F_setseed[0])) = int64(base.Ui64(v57)>>(uint(v58)%64)) ^ v57
		v63 = v45 - int64(7046029254386353131)
		v68 = (int64(base.Ui64(v63)>>(uint(v48)%64)) ^ v63) * v51
		v73 = (int64(base.Ui64(v68)>>(uint(v53)%64)) ^ v68) * v56
		*(*int64)(unsafe.Add(mBase, _c_F_setseed[1])) = int64(base.Ui64(v73)>>(uint(v58)%64)) ^ v73
		v79 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_setseed[2])) = uint8(v79)
		m.G0 = v7 + int32(16)
		return int32(0)
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
func F_sha384_bytea(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_cryptohash_internal(m, int32(4), v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_shdepChangeDep(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v44 int32
	_ = v44
	var v61 int32
	_ = v61
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	v13 = m.G0
	v15 = v13 - int32(256)
	m.G0 = v15
	v19 = int32(1)
	if l1 <= int32(3591) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_shdepChangeDep[0]))
	F_shdepLockAndCheckObject(m, l3, l4)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	goto L1
L3:
	;
	v90 = int32(0)
	goto L2
L4:
	;
	if base.B2i32(base.Ui32(l1-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l1-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v90 = v19
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
		v90 = v19
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
	v31 = l1 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v31))|base.B2i32(int32(1)<<(uint(v31)%32)&int32(226492515) == int32(0)) != 0 {
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
	v90 = v19
	goto L2
L13:
	;
	v90 = v19
	goto L2
L14:
	;
	if base.Ui32(l1-int32(3592)) < base.Ui32(int32(2)) {
		v90 = v19
		goto L2
	} else {
		goto L23
	}
L15:
	;
	v44 = l1 - int32(_a_F_shdepChangeDep_1)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v44))|base.B2i32(int32(1)<<(uint(v44)%32)&int32(963) == int32(0)) != 0 {
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
		v90 = v19
		goto L2
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L3
	default:
		goto L19
	}
L18:
	;
	v90 = v19
	goto L2
L19:
	;
	if base.Ui32(l1-int32(_a_F_shdepChangeDep_3)) < base.Ui32(int32(3)) {
		v90 = v19
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v61 = l1 - int32(_a_F_shdepChangeDep_4)
	if base.Ui32(int32(15)) < base.Ui32(v61) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if int32(1)<<(uint(v61)%32)&int32(_a_F_shdepChangeDep_5) != 0 {
		v90 = v19
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
	v90 = v19
	goto L2
L25:
	;
	goto L3
L26:
	;
	return
L27:
	;
	v96 = v15 - int32(-64)
	if v90 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v101 = int32(0)
	goto L30
L29:
	;
	v101 = v92
	goto L30
L30:
	;
	F_ScanKeyInit(m, v96, int32(1), int32(3), int32(184), v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	F_ScanKeyInit(m, v15+int32(112), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v113 = int32(3)
	F_ScanKeyInit(m, v15+int32(160), v113, v113, int32(184), l2)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	F_ScanKeyInit(m, v15+int32(208), int32(4), int32(3), int32(65), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v130 = F_systable_beginscan(m, l0, int32(1232), int32(1), int32(0), int32(4), v96)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v139 = int32(0)
	goto L37
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L26
	} else {
		goto L63
	}
L37:
	;
	v144 = F_systable_getnext(m, v130)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L26
	} else {
		goto L39
	}
L38:
	;
	F_systable_endscan(m, v130)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L26
	} else {
		goto L46
	}
L39:
	;
	if v144 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v144)+16))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+22)))
	v149 = int32(*(*int8)(unsafe.Add(mBase, uint32(v146+v147)+24)))
	if l5 != v149 {
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
	if v139 != 0 {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v151 = F_heap_copytuple(m, v144)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	v139 = v151
	goto L37
L46:
	;
	goto L49
L47:
	;
	m.G0 = v15 + int32(256)
	return
L48:
	;
	F_pfree(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
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
	if v139 == int32(0) {
		goto L47
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v139 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	F_simple_heap_delete(m, l0, v139+int32(4))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L26
	} else {
		goto L54
	}
L54:
	;
	v211 = v139
	goto L48
L55:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+22)))
	v180 = v178 + v179
	*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v180)+16)) = l3
	F_CatalogTupleUpdate(m, l0, v139+int32(4), v139)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L26
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l3
	v190 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v15)+27)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v190
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v204 = F_heap_form_tuple(m, v199, v15+int32(32), v15+int32(24))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L26
	} else {
		goto L59
	}
L58:
	;
	v211 = v139
	goto L48
L59:
	;
	F_CatalogTupleInsert(m, l0, v204)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L26
	} else {
		goto L60
	}
L60:
	;
	if v204 == int32(0) {
		goto L47
	} else {
		goto L61
	}
L61:
	;
	v211 = v204
	goto L48
L62:
	;
	goto L47
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	F_errmsg_internal(m, int32(_a_F_shdepChangeDep_7), v15)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L26
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_shdepChangeDep_8), int32(255), int32(_a_F_shdepChangeDep_9))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
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
			v14 = int32(0)
			v17 = F_SearchSysCacheExists(m, int32(11), l1, v14, v14, v14)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				if v17 != 0 {
					m.G0 = v6 - int32(-64)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l1
							F_errmsg(m, int32(_a_F_shdepLockAndCheckObject_0), v4+int32(-48))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_shdepLockAndCheckObject_1), int32(1223), int32(_a_F_shdepLockAndCheckObject_2))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
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
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_shdepLockAndCheckObject_3), v6)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_shdepLockAndCheckObject_1), int32(1256), int32(_a_F_shdepLockAndCheckObject_2))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 2:
			v37 = F_get_database_name(m, l1)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				if v37 != 0 {
					v92 = v37
					F_pfree(m, v92)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						m.G0 = v6 - int32(-64)
						return
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = l1
							F_errmsg(m, int32(_a_F_shdepLockAndCheckObject_4), v4+int32(-16))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_shdepLockAndCheckObject_1), int32(1249), int32(_a_F_shdepLockAndCheckObject_2))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
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
				v72 = F_get_tablespace_name(m, l1)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					if v72 != 0 {
						v92 = v72
						F_pfree(m, v92)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							m.G0 = v6 - int32(-64)
							return
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l1
								F_errmsg(m, int32(_a_F_shdepLockAndCheckObject_5), v4+int32(-32))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_shdepLockAndCheckObject_1), int32(1235), int32(_a_F_shdepLockAndCheckObject_2))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
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
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg_internal(m, int32(_a_F_shdepLockAndCheckObject_3), v6)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_shdepLockAndCheckObject_1), int32(1256), int32(_a_F_shdepLockAndCheckObject_2))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
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
func F_shell_out(m *base.Module, l0 int32) int32 {
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
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(_a_F_shell_out_0), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_shell_out_1), int32(317), int32(_a_F_shell_out_2))
				v19 = m.ExcPending
				if v19 != 0 {
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
func F_sjis_to_utf8(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn13870(m, l0, int32(35), v3, v3, v3, int32(_a_F_sjis_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_skip(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7
	v11 = v6
	goto L1
L1:
	;
	if base.Ui32(v11) <= base.Ui32(v9) {
		v98 = v9
		v100 = v11
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if base.Ui32(v98) < base.Ui32(v100) {
		goto L36
	} else {
		goto L37
	}
L4:
	;
	v15 = v9
	v17 = v11
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_skip[0]))
	switch v21 - int32(1) {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	default:
		goto L11
	}
L6:
	;
	v98 = v94
	v100 = v89
	goto L3
L7:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v88 == int32(0) {
		v98 = v90
		v100 = v89
		goto L3
	} else {
		goto L33
	}
L8:
	;
	if base.Ui32(int32(255)) < base.Ui32(v19) {
		v98 = v15
		v100 = v17
		goto L3
	} else {
		goto L31
	}
L9:
	;
	if v19 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v34 = Fn13991(m, v19, int32(5), int32(32), int32(_a_F_skip_0), int32(_a_F_skip_1), int32(10))
	mBase = m.M
	goto L13
L11:
	;
	if base.Ui32(int32(127)) < base.Ui32(v19) {
		v98 = v15
		v100 = v17
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_skip[1]))))
	v88 = int32(base.Ui32(v26) >> (uint(int32(7)) % 32))
	goto L7
L13:
	;
	v88 = v34
	goto L7
L14:
	;
	v88 = v73
	goto L7
L15:
	;
	v73 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	if v19 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v73 = base.B2i32(v66 != int32(0))
	goto L14
L19:
	;
	v46 = int32(_a_F_skip_2)
	goto L22
L20:
	;
	goto L21
L21:
	;
	v56 = int32(_a_F_skip_2)
	v57 = F_wcslen(m, v56)
	mBase = m.M
	v66 = v57<<(uint(int32(2))%32) + v56
	goto L18
L22:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v49 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v49 != 0 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if v19 != v49 {
		v46 = v46 + int32(4)
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L26
L28:
	;
	v55 = v46
	goto L30
L29:
	;
	v55 = int32(0)
	goto L30
L30:
	;
	v66 = v55
	goto L18
L31:
	;
	goto L32
L32:
	;
	v88 = base.B2i32(base.B2i32(v19 == int32(32))|base.B2i32(base.Ui32(v19-int32(9)) < base.Ui32(int32(5))) != int32(0))
	goto L7
L33:
	;
	v94 = v90 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v94
	if base.Ui32(v94) < base.Ui32(v89) {
		v15 = v94
		v17 = v89
		goto L5
	} else {
		goto L34
	}
L34:
	;
	goto L6
L35:
	;
	v114 = v98
	goto L43
L36:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v103 == int32(35) {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v98 != v7 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v108 | int32(128)
	goto L42
L41:
	;
	goto L42
L42:
	;
	return
L43:
	;
	v119 = v114 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v119
	if base.Ui32(v100) <= base.Ui32(v119) {
		v9 = v119
		v11 = v100
		goto L1
	} else {
		goto L45
	}
L44:
	;
	v9 = v119
	v11 = v100
	goto L1
L45:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v122 != int32(10) {
		v114 = v119
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
}
func F_slice_from_s(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v9 = int32(-1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v10 < int32(0) {
		v29 = v9
		return v29
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v13 < v10 {
			v29 = v9
			return v29
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v15 < v13 {
				v29 = v9
				return v29
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v17 == int32(0) {
					v29 = v9
					return v29
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v17-int32(4))))
					if v22 < v15 {
						v29 = v9
						return v29
					} else {
						v25 = F_replace_s(m, l0, v10, v13, l1, l2, int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = v25
							return v29
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
	v126 = F_FileReadV(m, v122, v22+int32(48), v112, v44, int32(167772181))
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
	F_errfinish(m, int32(_a_F_smgrreadv_5), int32(962), int32(_a_F_smgrreadv_6))
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
	v312 = F_FileReadV(m, v308, v275, v307, v310, int32(167772181))
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
	F_errfinish(m, int32(_a_F_smgrreadv_5), int32(913), int32(_a_F_smgrreadv_6))
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
	F_errfinish(m, int32(_a_F_smgrreadv_5), int32(875), int32(_a_F_smgrreadv_6))
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
func F_smgrshutdown(m *base.Module, l0 int32, l1 int32) {
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
	F_errfinish(m, int32(_a_F_spcache_insert_1), int32(630), int32(_a_F_spcache_insert_2))
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
	F_errfinish(m, int32(_a_F_spcache_insert_1), int32(327), int32(_a_F_spcache_insert_4))
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
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(277)
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
func F_ssup_datum_int32_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return base.B2i32(l1 < l0) - base.B2i32(l0 < l1)
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int64
	_ = v123
	var v126 int64
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int64
	_ = v245
	var v246 int32
	_ = v246
	var v248 int64
	_ = v248
	var v249 int32
	_ = v249
	var v251 int64
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int64
	_ = v270
	var v271 int64
	_ = v271
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int64
	_ = v289
	var v290 int32
	_ = v290
	var v292 int64
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int64
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v329 int32
	_ = v329
	var v334 int64
	_ = v334
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int64
	_ = v365
	var v368 int64
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int64
	_ = v461
	var v462 int64
	_ = v462
	var v470 int64
	_ = v470
	var v472 int32
	_ = v472
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v552 int64
	_ = v552
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v604 int32
	_ = v604
	var v605 int64
	_ = v605
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(224)
	m.G0 = v17
	v22 = v2
	v23 = int32(-1)
	v24 = v2
	v26 = v2
	v27 = v2
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
	v49 = v26
	v50 = v27
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
	v604 = int32(m.ExcTag)
	v605 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v604 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v22
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
	v552 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_start_apply[2])) = v552
	*(*int64)(unsafe.Add(mBase, _c_F_start_apply[3])) = v552
	v558 = int32(0)
	*(*uint16)(unsafe.Add(mBase, _c_F_start_apply[4])) = uint16(v558)
	v561 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[5]))
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561)+29)))
	if v562 == int32(1) {
		goto L120
	} else {
		goto L121
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v22
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[6]))
	v84 = F_AllocSetContextCreateInternal(m, v79, int32(_a_F_start_apply_0), int32(0), int32(_a_F_start_apply_1), int32(_a_F_start_apply_2))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[7])) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v22
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[6]))
	v95 = F_AllocSetContextCreateInternal(m, v90, int32(_a_F_start_apply_3), int32(0), int32(_a_F_start_apply_1), int32(_a_F_start_apply_2))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[8])) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v22
	v100 = int32(0)
	F_pgstat_report_activity(m, int32(2), v100)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = int32(987)
	v107 = v17 + int32(204)
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[9])) = v107
	v109 = int32(_a_F_start_apply_4)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[0])) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v110
	v116 = v22
	v117 = v100
	v123 = l0
	v126 = v67 + v66*int64(1000000) - int64(946684800000000)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[10]))
	if v133 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[0])) = v526
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[9])) = v526
	v531 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[11]))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v531)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	v535 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[12]))
	m.T0[v532].(func(*base.Module, int32, int32))(m, v535, v17+int32(216))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L10
	} else {
		goto L119
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v116
	F_ProcessInterrupts(m)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[13])) = v139
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[11]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v116
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[12]))
	v151 = m.T0[v143].(func(*base.Module, int32, int32, int32) int32)(m, v146, v17+int32(196), v17+int32(200))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v153 = int32(0)
	if v151 == v153 {
		v358 = v116
		v359 = v117
		v361 = v153
		v365 = v123
		v368 = v126
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	v371 = int32(0)
	F_send_feedback(m, v365, v371, v371)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L10
	} else {
		goto L72
	}
L25:
	;
	v158 = v116
	v159 = v117
	v160 = v151
	v165 = v123
	v168 = v126
	goto L26
L26:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[10]))
	if v171 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	F_ProcessInterrupts(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L10
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v160 == int32(0) {
		v358 = v158
		v359 = v159
		v361 = v153
		v365 = v165
		v368 = v168
		goto L24
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	if v160 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v182 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L10
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[14]))
	if v199 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v184 = int32(1)
	if v182 == int32(0) {
		v358 = v158
		v359 = v159
		v361 = v184
		v365 = v165
		v368 = v168
		goto L24
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	F_errmsg(m, int32(_a_F_start_apply_5), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	F_errfinish(m, int32(_a_F_start_apply_6), int32(3639), int32(_a_F_start_apply_7))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v358 = v158
	v359 = v159
	v361 = v184
	v365 = v165
	v368 = v168
	goto L24
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[14])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L10
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v211 = m.G0
	v212 = int32(16)
	v213 = v211 - v212
	m.G0 = v213
	F_gettimeofday(m, v213)
	mBase = m.M
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v213)))
	v217 = int64(*(*int32)(unsafe.Add(mBase, uint32(v213)+8)))
	m.G0 = v213 + v212
	goto L44
L43:
	;
	goto L42
L44:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[13])) = v228
	*(*int64)(unsafe.Add(mBase, uint32(v17)+188)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v160
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v17)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+180)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v237 = v17 + int32(180)
	v238 = F_pq_getmsgbyte(m, v237)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L10
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[7]))
	F_MemoryContextReset(m, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L10
	} else {
		goto L70
	}
L46:
	;
	v241 = v238 - int32(107)
	if v241 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if v241 != int32(12) {
		v334 = v165
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v288 = v17 + int32(180)
	v289 = F_pq_getmsgint64(m, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L10
	} else {
		goto L62
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v245 = F_pq_getmsgint64(m, v237)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v248 = F_pq_getmsgint64(m, v237)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v251 = F_pq_getmsgint64(m, v237)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v254)+80)) = v251
	if base.Ui64(v245) < base.Ui64(v165) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v257 = v165
	goto L56
L55:
	;
	v257 = v245
	goto L56
L56:
	;
	if base.Ui64(v248) < base.Ui64(v257) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v259 = v257
	goto L59
L58:
	;
	v259 = v248
	goto L59
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v254)+72)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v265 = m.G0
	v266 = int32(16)
	v267 = v265 - v266
	m.G0 = v267
	F_gettimeofday(m, v267)
	mBase = m.M
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v267)))
	v271 = int64(*(*int32)(unsafe.Add(mBase, uint32(v267)+8)))
	m.G0 = v267 + v266
	goto L60
L60:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+88)) = v271 + v270*int64(1000000) - int64(946684800000000)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	F_apply_dispatch(m, v237)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	v334 = v259
	goto L45
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v292 = F_pq_getmsgint64(m, v288)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v295 = F_pq_getmsgbyte(m, v288)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	if base.Ui64(v289) < base.Ui64(v165) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v299 = v165
	goto L67
L66:
	;
	v299 = v289
	goto L67
L67:
	;
	v300 = int32(0)
	F_send_feedback(m, v299, base.B2i32(v295 != v300), v300)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v306)+80)) = v292
	*(*int64)(unsafe.Add(mBase, uint32(v306)+72)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v313 = m.G0
	v314 = int32(16)
	v315 = v313 - v314
	m.G0 = v315
	F_gettimeofday(m, v315)
	mBase = m.M
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
	v319 = int64(*(*int32)(unsafe.Add(mBase, uint32(v315)+8)))
	m.G0 = v315 + v314
	goto L69
L69:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v329)+104)) = v292
	*(*int64)(unsafe.Add(mBase, uint32(v329)+96)) = v299
	*(*int64)(unsafe.Add(mBase, uint32(v329)+88)) = v319 + v318*int64(1000000) - int64(946684800000000)
	v334 = v299
	goto L45
L70:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[11]))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v158
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[12]))
	v354 = m.T0[v346].(func(*base.Module, int32, int32, int32) int32)(m, v349, v17+int32(196), v17+int32(200))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L10
	} else {
		goto L71
	}
L71:
	;
	v158 = v354
	v159 = int32(0)
	v160 = v354
	v165 = v334
	v168 = v217 + v216*int64(1000000) - int64(946684800000000)
	goto L26
L72:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_apply[16])))
	if v376 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[7]))
	F_MemoryContextReset(m, v392)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L10
	} else {
		goto L79
	}
L74:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_apply[17])))
	if v378&int32(1) != 0 {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	F_maybe_reread_subscription(m)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L10
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	F_process_syncing_tables(m, v365)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	goto L73
L79:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[18]))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[13])) = v397
	if v361 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	v403 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[19]))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v17)+200))
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[20]))
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[21]))
	if v409 == int32(_a_F_start_apply_8) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	goto L82
L82:
	;
	goto L18
L83:
	;
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[14]))
	if v435 != 0 {
		goto L95
	} else {
		goto L96
	}
L84:
	;
	v412 = int32(1000)
	goto L86
L85:
	;
	v412 = v407
	goto L86
L86:
	;
	if v409 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v414 = v412
	goto L89
L88:
	;
	v414 = int32(1000)
	goto L89
L89:
	;
	v416 = F_WaitLatchOrSocket(m, v403, v404, v414, int32(83886087))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	if v416&int32(1) == int32(0) {
		goto L83
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v424))) = int32(0)
	goto L92
L92:
	;
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[10]))
	if v428 == int32(0) {
		goto L83
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	F_ProcessInterrupts(m)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L10
	} else {
		goto L94
	}
L94:
	;
	goto L83
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[14])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L10
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if v416&int32(8) == int32(0) {
		v116 = v358
		v117 = v359
		v123 = v365
		v126 = v368
		goto L17
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	v447 = int32(0)
	v449 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[22]))
	if v449 <= v447 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	F_send_feedback(m, v365, v508, v508)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L10
	} else {
		goto L113
	}
L101:
	;
	v508 = v447
	v509 = v359
	goto L100
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	v456 = m.G0
	v457 = int32(16)
	v458 = v456 - v457
	m.G0 = v458
	F_gettimeofday(m, v458)
	mBase = m.M
	v461 = *(*int64)(unsafe.Add(mBase, uint32(v458)))
	v462 = int64(*(*int32)(unsafe.Add(mBase, uint32(v458)+8)))
	m.G0 = v458 + v457
	v470 = v462 + v461*int64(1000000) - int64(946684800000000)
	goto L104
L104:
	;
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[22]))
	if base.I64_extend_i32_s(v472)*int64(1000)+v368 <= v470 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L10
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v498 = int32(1)
	if v359&v498 != 0 {
		v508 = v447
		v509 = v498
		goto L100
	} else {
		goto L112
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	F_errcode(m, int32(100663808))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	F_errmsg(m, int32(_a_F_start_apply_9), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	F_errfinish(m, int32(_a_F_start_apply_6), int32(3793), int32(_a_F_start_apply_7))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L10
	} else {
		goto L111
	}
L111:
	;
	goto L7
L112:
	;
	v502 = base.I32_div_s(v472, int32(2))
	v507 = base.B2i32(base.I64_extend_i32_s(v502)*int64(1000)+v368 <= v470)
	v508 = v507
	v509 = v507
	goto L100
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	v517 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[23]))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)+20))
	goto L114
L114:
	;
	if v518 == int32(2) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v116 = v358
	v117 = v509
	v123 = v365
	v126 = v368
	goto L17
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v358
	v523 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L10
	} else {
		goto L118
	}
L118:
	;
	v116 = v358
	v117 = v509
	v123 = v365
	v126 = v368
	goto L17
L119:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[1])) = v50
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[0])) = v49
	m.G0 = v17 + int32(224)
	return
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v22
	F_DisableSubscriptionAndExit(m)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L10
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v22
	F_AbortOutOfAnyTransaction(m)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L10
	} else {
		goto L124
	}
L123:
	;
	goto L7
L124:
	;
	v571 = int32(1)
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[5]))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[15]))
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+16)))
	if v577 == v571 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	v583 = base.B2i32(v580 != int32(1))
	goto L127
L126:
	;
	v583 = v571
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v22
	F_pgstat_report_subscription_error(m, v574, v583)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L10
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v22
	F_pg_re_throw(m)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L10
	} else {
		goto L129
	}
L129:
	;
	goto L9
L130:
	;
	v609 = int32(v605)
	m.G0 = v17
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v609)+4))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v609)))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	if v17+int32(12) == v615 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	m.ExcPending = 1
	goto L139
L132:
	;
	if v619 != 0 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v612)+4))
	v619 = v617
	goto L135
L134:
	;
	v619 = int32(0)
	goto L135
L135:
	;
	goto L132
L136:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v17)+220))
	v22 = v620
	v23 = v619
	v24 = v611
	v26 = v49
	v27 = v50
	goto L1
L137:
	;
	goto L138
L138:
	;
	F___wasm_longjmp(m, v612, v611)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	return
L140:
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
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
	var v49 int64
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v185 int32
	_ = v185
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
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	v13 = m.G0
	v15 = v13 - int32(80)
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
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L51
	}
L4:
	;
	F_relation_close(m, v19, int32(3))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
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
	v35 = int32(0)
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v35<<(uint(int32(2))%32))))
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+8)))
	v47 = F_SearchSysCacheExists(m, int32(5), l1, v44, v45, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v49 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v49
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+44)) = uint16(v55)
	v61 = F_GetNewOidWithIndex(m, v19, int32(2757), int32(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	if v47 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v61
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v67
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v78 = F_heap_form_tuple(m, v73, v15+int32(48), v15+int32(40))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_CatalogTupleInsert(m, v19, v78)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_pfree(m, v78)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = int32(2603)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(1255)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v91
	v96 = v15 + int32(28)
	v98 = v15 + int32(16)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
	if v101 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v102 = int32(110)
	goto L20
L19:
	;
	v102 = int32(97)
	goto L20
L20:
	;
	F_recordDependencyOn(m, v96, v98, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+25)))
	if v107 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v108 = int32(2753)
	goto L24
L23:
	;
	v108 = int32(2616)
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v110
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
	if v116 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v117 = int32(105)
	goto L27
L26:
	;
	v117 = int32(97)
	goto L27
L27:
	;
	F_recordDependencyOn(m, v96, v98, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v121 = F_typeDepNeeded(m, v120, v42)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v121 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(1247)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v125
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
	if v131 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v136 == v137 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v132 = int32(110)
	goto L35
L34:
	;
	v132 = int32(97)
	goto L35
L35:
	;
	F_recordDependencyOn(m, v96, v98, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_storeProcedures[0]))
	if v161 != 0 {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v139 = F_typeDepNeeded(m, v136, v42)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v139 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(1247)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v145
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
	if v155 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v156 = int32(110)
	goto L43
L42:
	;
	v156 = int32(97)
	goto L43
L43:
	;
	F_recordDependencyOn(m, v15+int32(28), v15+int32(16), v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L37
L45:
	;
	v163 = int32(0)
	F_RunObjectPostCreateHook(m, int32(2603), v61, v163, v163)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v168 = v35 + int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v168 < v169 {
		v35 = v168
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
	m.G0 = v15 + int32(80)
	return
L51:
	;
	F_errcode(m, int32(_a_F_storeProcedures_0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v198 = F_format_type_be(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v201 = F_format_type_be(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v203 = F_NameListToString(m, l0)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v196
	F_errmsg(m, int32(_a_F_storeProcedures_1), v15)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_storeProcedures_2), int32(1618), int32(_a_F_storeProcedures_3))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
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
	var v117 int32
	_ = v117
	var v131 int32
	_ = v131
	var v145 int32
	_ = v145
	var v159 int32
	_ = v159
	var v181 int32
	_ = v181
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v225 int32
	_ = v225
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v290 int32
	_ = v290
	var v304 int32
	_ = v304
	var v318 int32
	_ = v318
	var v332 int32
	_ = v332
	var v346 int32
	_ = v346
	var v360 int32
	_ = v360
	var v374 int32
	_ = v374
	var v396 int32
	_ = v396
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v440 int32
	_ = v440
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v484 int32
	_ = v484
	var v505 int32
	_ = v505
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v683 int32
	_ = v683
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
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
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L134
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L131
	}
L5:
	;
	v683 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v667))) = uint8(v683)
	m.G0 = v25 + int32(32)
	return v673
L6:
	;
	v667 = v30
	v673 = v30
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
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L126
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
	if v615 == int32(0) {
		v667 = v604
		v673 = v75
		goto L5
	} else {
		goto L125
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
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(12))+8))
	*(*int32)(unsafe.Add(mBase, _c_F_str_udeescape[0])) = v619
	goto L123
L18:
	;
	F_pg_unicode_to_server(m, v585, v74)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L122
	}
L19:
	;
	v83 = v25 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = int32(499)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v41 + (v37 + int32(3))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v83
	v90 = int32(_a_F_str_udeescape_0)
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
	if v43 != 0 {
		v625 = v41
		goto L10
	} else {
		goto L120
	}
L22:
	;
	if v43 != 0 {
		v625 = v41
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
	v603 = int32(2)
	v604 = v74 + int32(1)
	v615 = int32(0)
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
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)))
	goto L29
L29:
	;
	if base.B2i32(base.Ui32(v117-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v117|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+3)))
	goto L31
L31:
	;
	if base.B2i32(base.Ui32(v131-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v131|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	goto L33
L33:
	;
	if base.B2i32(base.Ui32(v145-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v145|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v159 = int32(-48)
	if base.Ui32((v97-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v181 = v159
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if base.Ui32((v117-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v202 = v159
		goto L39
	} else {
		goto L40
	}
L36:
	;
	if base.Ui32((v97-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v181 = int32(-87)
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
	v181 = int32(-55)
	goto L35
L39:
	;
	v203 = int32(-48)
	if base.Ui32((v131-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v225 = v203
		goto L45
	} else {
		goto L46
	}
L40:
	;
	if base.Ui32((v117-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v202 = int32(-87)
	goto L39
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v117-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v202 = int32(-55)
	goto L39
L45:
	;
	if base.Ui32((v145-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v246 = v203
		goto L49
	} else {
		goto L50
	}
L46:
	;
	if base.Ui32((v131-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v225 = int32(-87)
		goto L45
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v131-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v225 = int32(-55)
	goto L45
L49:
	;
	v259 = v246 + v145 + ((v117+v202)<<(uint(int32(8))%32) + (v97+v181)<<(uint(int32(12))%32) + (v131+v225)<<(uint(int32(4))%32))
	if base.Ui32(int32(_a_F_str_udeescape_1)) <= base.Ui32(v259-int32(1)) {
		goto L3
	} else {
		goto L55
	}
L50:
	;
	if base.Ui32((v145-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v246 = int32(-87)
	goto L49
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v145-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v246 = int32(-55)
	goto L49
L55:
	;
	v265 = v259 & int32(_a_F_str_udeescape_2)
	if v43 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v280 = int32(5)
	if v279&int32(16776192) != int32(_a_F_str_udeescape_3) {
		v585 = v279
		v586 = v280
		goto L18
	} else {
		goto L62
	}
L57:
	;
	if v265 != int32(_a_F_str_udeescape_4) {
		v625 = v41
		goto L10
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v265 == int32(_a_F_str_udeescape_4) {
		v625 = v41
		goto L10
	} else {
		goto L61
	}
L60:
	;
	v279 = v43<<(uint(int32(10))%32)&int32(_a_F_str_udeescape_5) | v259&int32(1023) + int32(_a_F_str_udeescape_6)
	goto L56
L61:
	;
	v279 = v259
	goto L56
L62:
	;
	v603 = v280
	v604 = v74
	v615 = v279
	goto L17
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L115
	}
L64:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)))
	goto L65
L65:
	;
	if base.B2i32(base.Ui32(v290-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v290|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+3)))
	goto L67
L67:
	;
	if base.B2i32(base.Ui32(v304-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v304|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L63
	} else {
		goto L68
	}
L68:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	goto L69
L69:
	;
	if base.B2i32(base.Ui32(v318-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v318|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L63
	} else {
		goto L70
	}
L70:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+5)))
	goto L71
L71:
	;
	if base.B2i32(base.Ui32(v332-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v332|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+6)))
	goto L73
L73:
	;
	if base.B2i32(base.Ui32(v346-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v346|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L63
	} else {
		goto L74
	}
L74:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
	goto L75
L75:
	;
	if base.B2i32(base.Ui32(v360-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v360|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L63
	} else {
		goto L76
	}
L76:
	;
	v374 = int32(-48)
	if base.Ui32((v290-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v396 = v374
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if base.Ui32((v304-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v417 = v374
		goto L81
	} else {
		goto L82
	}
L78:
	;
	if base.Ui32((v290-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v396 = int32(-87)
		goto L77
	} else {
		goto L79
	}
L79:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v290-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	v396 = int32(-55)
	goto L77
L81:
	;
	v418 = int32(-48)
	if base.Ui32((v318-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v440 = v418
		goto L87
	} else {
		goto L88
	}
L82:
	;
	if base.Ui32((v304-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v417 = int32(-87)
	goto L81
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v304-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v417 = int32(-55)
	goto L81
L87:
	;
	if base.Ui32((v332-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v461 = v418
		goto L91
	} else {
		goto L92
	}
L88:
	;
	if base.Ui32((v318-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v440 = int32(-87)
		goto L87
	} else {
		goto L89
	}
L89:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v318-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	v440 = int32(-55)
	goto L87
L91:
	;
	v462 = int32(-48)
	if base.Ui32((v346-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v484 = v462
		goto L97
	} else {
		goto L98
	}
L92:
	;
	if base.Ui32((v332-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v461 = int32(-87)
	goto L91
L94:
	;
	goto L95
L95:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v332-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	v461 = int32(-55)
	goto L91
L97:
	;
	if base.Ui32((v360-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v505 = v462
		goto L101
	} else {
		goto L102
	}
L98:
	;
	if base.Ui32((v346-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v484 = int32(-87)
		goto L97
	} else {
		goto L99
	}
L99:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v346-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	v484 = int32(-55)
	goto L97
L101:
	;
	v526 = v360 + v505 + ((v417+v304)<<(uint(int32(16))%32) + (v290+v396)<<(uint(int32(20))%32) + (v318+v440)<<(uint(int32(12))%32) + (v332+v461)<<(uint(int32(8))%32) + (v346+v484)<<(uint(int32(4))%32))
	if base.Ui32(int32(_a_F_str_udeescape_1)) <= base.Ui32(v526-int32(1)) {
		goto L9
	} else {
		goto L107
	}
L102:
	;
	if base.Ui32((v360-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v505 = int32(-87)
	goto L101
L104:
	;
	goto L105
L105:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v360-int32(65))&int32(255)) {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	v505 = int32(-55)
	goto L101
L107:
	;
	v532 = v526 & int32(_a_F_str_udeescape_2)
	if v43 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v547 = int32(8)
	if v546&int32(16776192) == int32(_a_F_str_udeescape_3) {
		v603 = v547
		v604 = v74
		v615 = v546
		goto L17
	} else {
		goto L114
	}
L109:
	;
	if v532 != int32(_a_F_str_udeescape_4) {
		v625 = v41
		goto L10
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	if v532 == int32(_a_F_str_udeescape_4) {
		v625 = v41
		goto L10
	} else {
		goto L113
	}
L112:
	;
	v546 = v43<<(uint(int32(10))%32)&int32(_a_F_str_udeescape_5) | v526&int32(1023) + int32(_a_F_str_udeescape_6)
	goto L108
L113:
	;
	v546 = v526
	goto L108
L114:
	;
	v585 = v546
	v586 = v547
	goto L18
L115:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(_a_F_str_udeescape_7), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errhint(m, int32(_a_F_str_udeescape_8), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(495), int32(_a_F_str_udeescape_10))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v73)
	v579 = int32(1)
	v580 = v74 + v579
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v581 != 0 {
		v41 = v41 + v579
		v43 = int32(0)
		v45 = v581
		v47 = v580
		v53 = v75
		v56 = v76
		goto L11
	} else {
		goto L121
	}
L121:
	;
	v667 = v580
	v673 = v75
	goto L5
L122:
	;
	v599 = F_strlen(m, v74)
	mBase = m.M
	v603 = v586
	v604 = v599 + v74
	v615 = int32(0)
	goto L17
L123:
	;
	v621 = v41 + v603
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621))))
	if v622 != 0 {
		v41 = v621
		v43 = v615
		v45 = v622
		v47 = v604
		v53 = v75
		v56 = v76
		goto L11
	} else {
		goto L124
	}
L124:
	;
	goto L12
L125:
	;
	v625 = v621
	goto L10
L126:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errmsg(m, int32(_a_F_str_udeescape_11), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_scanner_errposition(m, v625+v37+int32(3), l3)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(525), int32(_a_F_str_udeescape_10))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errmsg_internal(m, int32(_a_F_str_udeescape_12), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(336), int32(_a_F_str_udeescape_13))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errmsg(m, int32(_a_F_str_udeescape_14), int32(0))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(347), int32(_a_F_str_udeescape_15))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
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
	var v109 int32
	_ = v109
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v152 int32
	_ = v152
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
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
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v742 int32
	_ = v742
	var v751 int32
	_ = v751
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v872 int64
	_ = v872
	var v885 int64
	_ = v885
	var v886 int32
	_ = v886
	var v912 int64
	_ = v912
	var v922 int32
	_ = v922
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v944 int64
	_ = v944
	var v957 int64
	_ = v957
	var v958 int32
	_ = v958
	var v961 int64
	_ = v961
	var v986 int64
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v998 int32
	_ = v998
	var v1004 int32
	_ = v1004
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
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
	return v1050
L2:
	;
	v742 = v23 + int32(80)
	goto L187
L3:
	;
	v727 = int32(0)
	v728 = int32(-1)
	goto L2
L4:
	;
	v698 = int32(0)
	v699 = F_errsave_start(m, l1)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L105
	} else {
		goto L182
	}
L5:
	;
	v675 = int32(0)
	v676 = F_errsave_start(m, l1)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L105
	} else {
		goto L177
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
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v178)
	switch v170 - int32(8) {
	case 0:
		goto L63
	default:
		goto L5
	case 2:
		goto L64
	case 4:
		goto L65
	case 5:
		goto L66
	}
L9:
	;
	v176 = v50 + int32(1)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v177 != 0 {
		v47 = v177
		v48 = v170
		v49 = v171
		v50 = v176
		v51 = v172
		v54 = v173
		v56 = v174
		goto L7
	} else {
		goto L54
	}
L10:
	;
	if v79 == int32(0) {
		goto L5
	} else {
		goto L52
	}
L11:
	;
	if v62 != 0 {
		goto L10
	} else {
		goto L51
	}
L12:
	;
	switch v69 - int32(32) {
	case 0, 13:
		v170 = v48
		v171 = v49
		v172 = v51
		v173 = v54
		v174 = v80
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
		goto L50
	}
L14:
	;
	if v67&(base.B2i32(v48 == int32(11))&v79) == int32(0) {
		goto L12
	} else {
		goto L48
	}
L15:
	;
	if base.B2i32(v69 == int32(88))|v79 == int32(0) {
		goto L34
	} else {
		goto L35
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
		v170 = int32(0)
		v171 = v49
		v172 = v51
		v173 = v54
		v174 = v80
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
	v170 = v91
	v171 = v49 + v91
	v172 = int32(4)
	v173 = v54
	v174 = v80
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
	if base.Ui32(v69-int32(97)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v109)
	v170 = int32(8)
	v171 = v49 + int32(1)
	v172 = int32(5)
	v173 = v54
	v174 = v80
	goto L9
L30:
	;
	v109 = v69 & int32(95)
	goto L32
L31:
	;
	v109 = v69
	goto L32
L32:
	;
	goto L29
L33:
	;
	if v51&int32(-5) != 0 {
		goto L5
	} else {
		goto L40
	}
L34:
	;
	if base.B2i32(v69 == int32(120))&v67 != 0 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v62 == int32(33) {
		goto L33
	} else {
		goto L38
	}
L37:
	;
	goto L12
L38:
	;
	if v62 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	goto L33
L40:
	;
	if base.Ui32(v69-int32(97)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v131)
	if v51 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v131 = v69 & int32(95)
	goto L44
L43:
	;
	v131 = v69
	goto L44
L44:
	;
	goto L41
L45:
	;
	v136 = v51
	goto L47
L46:
	;
	v136 = int32(3)
	goto L47
L47:
	;
	v170 = int32(10)
	v171 = v49 + int32(1)
	v172 = v136
	v173 = v54
	v174 = v80
	goto L9
L48:
	;
	if v51 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v47)
	v170 = int32(12)
	v171 = v49 + int32(1)
	v172 = int32(6)
	v173 = v54
	v174 = v80
	goto L9
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v47)
	v152 = int32(1)
	v170 = v152
	v171 = v49 + v152
	v172 = v51
	v173 = v54
	v174 = v80
	goto L9
L51:
	;
	v170 = v48
	v171 = v49
	v172 = v51
	v173 = v54 & v80
	v174 = int32(1)
	goto L9
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v47)
	if base.Ui32(int32(12)) < base.Ui32(v48) {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v166 = int32(1)
	v170 = v48 + v166
	v171 = v49 + v166
	v172 = v51
	v173 = v54
	v174 = v80
	goto L9
L54:
	;
	goto L8
L55:
	;
	v574 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+82)) = uint8(v574)
	if v173&int32(1) == int32(0) {
		goto L3
	} else {
		goto L157
	}
L56:
	;
	v510 = int32(_a_F_string2ean_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+90)) = uint16(v510)
	v513 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_string2ean[3])))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+80)) = uint16(v513)
	v516 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string2ean[4])))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+82)) = uint8(v516)
	if v173&int32(1) == int32(0) {
		goto L3
	} else {
		goto L144
	}
L57:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string2ean[5])))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+82)) = uint8(v449)
	v452 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_string2ean[6])))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+80)) = uint16(v452)
	if v173&int32(1) == int32(0) {
		goto L3
	} else {
		goto L131
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = int32(809056057)
	if v173&int32(1) == int32(0) {
		goto L3
	} else {
		goto L111
	}
L59:
	;
	v336 = int32(0)
	v337 = F_errsave_start(m, l1)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L105
	} else {
		goto L106
	}
L60:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+83)))
	if v297 == int32(48) {
		v329 = int32(6)
		goto L96
	} else {
		goto L97
	}
L61:
	;
	v218 = int32(0)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v228 = base.B2i32(v226 == int32(77))
	if v226 == int32(77) {
		goto L78
	} else {
		goto L79
	}
L62:
	;
	if base.B2i32(l3 == int32(2))|base.B2i32(l3 != v210) != 0 {
		v333 = v210
		goto L59
	} else {
		goto L76
	}
L63:
	;
	switch v172 {
	case 0, 5:
		goto L74
	default:
		goto L5
	}
L64:
	;
	if base.Ui32(v172-int32(5)) < base.Ui32(int32(-2)) {
		goto L5
	} else {
		goto L70
	}
L65:
	;
	v186 = int32(6)
	if v172 != v186 {
		goto L5
	} else {
		goto L69
	}
L66:
	;
	if v172 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	if v173&int32(1) != 0 {
		goto L61
	} else {
		goto L68
	}
L68:
	;
	v292 = int32(-1)
	v293 = int32(0)
	goto L60
L69:
	;
	v189 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+94)))
	v210 = v186
	v211 = v189 - int32(48)
	goto L62
L70:
	;
	v196 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+92)))
	if v196 == int32(88) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v210 = v172
	v211 = int32(10)
	goto L62
L72:
	;
	goto L73
L73:
	;
	v210 = v172
	v211 = v196 - int32(48)
	goto L62
L74:
	;
	v202 = int32(5)
	v204 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+90)))
	if v204 == int32(88) {
		v210 = v202
		v211 = int32(10)
		goto L62
	} else {
		goto L75
	}
L75:
	;
	v210 = v202
	v211 = v204 - int32(48)
	goto L62
L76:
	;
	switch l3 - int32(4) {
	case 0:
		goto L58
	case 1:
		goto L56
	case 2:
		goto L55
	default:
		goto L57
	}
L77:
	;
	v287 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+95)))
	v292 = v286
	v293 = base.B2i32(v286 == v287-int32(48)) | v174
	goto L60
L78:
	;
	v229 = int32(3)
	goto L80
L79:
	;
	v229 = v218
	goto L80
L80:
	;
	if v226 == int32(0) {
		v274 = v229
		v276 = v218
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v278 = int32(10)
	v283 = base.I32_rem_u_s(v274*int32(3)+v276, v278)
	if v283 != 0 {
		goto L93
	} else {
		goto L94
	}
L82:
	;
	v233 = v41
	v234 = v226
	v235 = v228
	v236 = v229
	v237 = int32(13)
	v238 = v218
	goto L83
L83:
	;
	v243 = (v234 - int32(48)) & int32(255)
	if base.Ui32(v243) <= base.Ui32(int32(9)) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v274 = v260
	v276 = v262
	goto L81
L85:
	;
	v248 = v235 & int32(1)
	if v248 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v259 = v235
	v260 = v236
	v261 = v237
	v262 = v238
	goto L87
L87:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+1)))
	if v264 == int32(0) {
		v274 = v260
		v276 = v262
		goto L81
	} else {
		goto L91
	}
L88:
	;
	v249 = int32(0)
	goto L90
L89:
	;
	v249 = v243
	goto L90
L90:
	;
	v255 = int32(1)
	v259 = v235 + v255
	v260 = (int32(0)-v248)&v243 + v236
	v261 = v237 - v255
	v262 = v249 + v238
	goto L87
L91:
	;
	v267 = int32(1)
	if base.Ui32(v267) < base.Ui32(v261) {
		v233 = v233 + v267
		v234 = v264
		v235 = v259
		v236 = v260
		v237 = v261
		v238 = v262
		goto L83
	} else {
		goto L92
	}
L92:
	;
	goto L84
L93:
	;
	v286 = v278 - v283
	goto L95
L94:
	;
	v286 = int32(0)
	goto L95
L95:
	;
	goto L77
L96:
	;
	if base.B2i32(l3 == int32(2))|base.B2i32(v329 == l3) != 0 {
		v727 = v293
		v728 = v292
		goto L2
	} else {
		goto L104
	}
L97:
	;
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+83)))
	v303 = v301 ^ int32(_a_F_string2ean_1)
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+85)))
	if v303|(v304^int32(55)) == int32(0) {
		v329 = int32(5)
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+85)))
	if v311^int32(56)|v303 == int32(0) {
		v329 = int32(3)
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v23)+83))
	if v318 == int32(809056057) {
		v329 = int32(4)
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+85)))
	if v323^int32(57)|v303 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v327 = int32(2)
	goto L103
L102:
	;
	v327 = int32(3)
	goto L103
L103:
	;
	v329 = v327
	goto L96
L104:
	;
	v333 = v329
	goto L59
L105:
	;
	return int32(0)
L106:
	;
	if v337 == int32(0) {
		v1050 = v336
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L105
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = l0
	v347 = int32(2)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(v347)%32))+uint32(_c_F_string2ean[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v351
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v333<<(uint(v347)%32))+uint32(_c_F_string2ean[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v357
	F_errmsg(m, int32(_a_F_string2ean_2), v23+int32(48))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L105
	} else {
		goto L109
	}
L109:
	;
	F_errsave_finish(m, l1, int32(_a_F_string2ean_3), int32(908), int32(_a_F_string2ean_4))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L105
	} else {
		goto L110
	}
L110:
	;
	v1050 = v336
	goto L1
L111:
	;
	v376 = v23 + int32(80)
	v377 = int32(0)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	v387 = base.B2i32(v385 == int32(77))
	if v385 == int32(77) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v727 = base.B2i32(v445 == v211) | v174
	v728 = v445
	goto L2
L113:
	;
	v388 = int32(3)
	goto L115
L114:
	;
	v388 = v377
	goto L115
L115:
	;
	if v385 == int32(0) {
		v433 = v388
		v435 = v377
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v437 = int32(10)
	v442 = base.I32_rem_u_s(v433*int32(3)+v435, v437)
	if v442 != 0 {
		goto L128
	} else {
		goto L129
	}
L117:
	;
	v392 = v376
	v393 = v385
	v394 = v387
	v395 = v388
	v396 = int32(13)
	v397 = v377
	goto L118
L118:
	;
	v402 = (v393 - int32(48)) & int32(255)
	if base.Ui32(v402) <= base.Ui32(int32(9)) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v433 = v419
	v435 = v421
	goto L116
L120:
	;
	v407 = v394 & int32(1)
	if v407 != 0 {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v418 = v394
	v419 = v395
	v420 = v396
	v421 = v397
	goto L122
L122:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+1)))
	if v423 == int32(0) {
		v433 = v419
		v435 = v421
		goto L116
	} else {
		goto L126
	}
L123:
	;
	v408 = int32(0)
	goto L125
L124:
	;
	v408 = v402
	goto L125
L125:
	;
	v414 = int32(1)
	v418 = v394 + v414
	v419 = (int32(0)-v407)&v402 + v395
	v420 = v396 - v414
	v421 = v408 + v397
	goto L122
L126:
	;
	v426 = int32(1)
	if base.Ui32(v426) < base.Ui32(v420) {
		v392 = v392 + v426
		v393 = v423
		v394 = v418
		v395 = v419
		v396 = v420
		v397 = v421
		goto L118
	} else {
		goto L127
	}
L127:
	;
	goto L119
L128:
	;
	v445 = v437 - v442
	goto L130
L129:
	;
	v445 = int32(0)
	goto L130
L130:
	;
	goto L112
L131:
	;
	v459 = int32(0)
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v462 == v459 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v727 = base.B2i32(v507 == v211) | v174
	v728 = v507
	goto L2
L133:
	;
	v507 = int32(0)
	goto L132
L134:
	;
	v465 = v41
	v466 = int32(10)
	v467 = v462
	v468 = v459
	goto L135
L135:
	;
	v473 = (v467 - int32(48)) & int32(255)
	v477 = base.B2i32(base.Ui32(v473) < base.Ui32(int32(10)))
	if base.Ui32(v473) < base.Ui32(int32(10)) {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v491 = base.I32_rem_u_s(v479, int32(11))
	if v491 == int32(0) {
		goto L133
	} else {
		goto L143
	}
L137:
	;
	goto L136
L138:
	;
	v478 = v466 * v473
	goto L140
L139:
	;
	v478 = int32(0)
	goto L140
L140:
	;
	v479 = v478 + v468
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+1)))
	if v480 == int32(0) {
		goto L137
	} else {
		goto L141
	}
L141:
	;
	v483 = int32(1)
	v485 = v466 - v477
	if base.Ui32(v483) < base.Ui32(v485) {
		v465 = v465 + v483
		v466 = v485
		v467 = v480
		v468 = v479
		goto L135
	} else {
		goto L142
	}
L142:
	;
	goto L137
L143:
	;
	v507 = int32(11) - v491
	goto L132
L144:
	;
	v523 = int32(0)
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v526 == v523 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v727 = base.B2i32(v571 == v211) | v174
	v728 = v571
	goto L2
L146:
	;
	v571 = int32(0)
	goto L145
L147:
	;
	v529 = v41
	v530 = int32(8)
	v531 = v526
	v532 = v523
	goto L148
L148:
	;
	v537 = (v531 - int32(48)) & int32(255)
	v541 = base.B2i32(base.Ui32(v537) < base.Ui32(int32(10)))
	if base.Ui32(v537) < base.Ui32(int32(10)) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v555 = base.I32_rem_u_s(v543, int32(11))
	if v555 == int32(0) {
		goto L146
	} else {
		goto L156
	}
L150:
	;
	goto L149
L151:
	;
	v542 = v530 * v537
	goto L153
L152:
	;
	v542 = int32(0)
	goto L153
L153:
	;
	v543 = v542 + v532
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529)+1)))
	if v544 == int32(0) {
		goto L150
	} else {
		goto L154
	}
L154:
	;
	v547 = int32(1)
	v549 = v530 - v541
	if base.Ui32(v547) < base.Ui32(v549) {
		v529 = v529 + v547
		v530 = v549
		v531 = v544
		v532 = v543
		goto L148
	} else {
		goto L155
	}
L155:
	;
	goto L150
L156:
	;
	v571 = int32(11) - v555
	goto L145
L157:
	;
	v583 = v23 + int32(80) | int32(2)
	v584 = int32(0)
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583))))
	v594 = base.B2i32(v592 == int32(77))
	if v592 == int32(77) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v727 = base.B2i32(v652 == v211) | v174
	v728 = v652
	goto L2
L159:
	;
	v595 = int32(3)
	goto L161
L160:
	;
	v595 = v584
	goto L161
L161:
	;
	if v592 == int32(0) {
		v640 = v595
		v642 = v584
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v644 = int32(10)
	v649 = base.I32_rem_u_s(v640*int32(3)+v642, v644)
	if v649 != 0 {
		goto L174
	} else {
		goto L175
	}
L163:
	;
	v599 = v583
	v600 = v592
	v601 = v594
	v602 = v595
	v603 = int32(13)
	v604 = v584
	goto L164
L164:
	;
	v609 = (v600 - int32(48)) & int32(255)
	if base.Ui32(v609) <= base.Ui32(int32(9)) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v640 = v626
	v642 = v628
	goto L162
L166:
	;
	v614 = v601 & int32(1)
	if v614 != 0 {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	v625 = v601
	v626 = v602
	v627 = v603
	v628 = v604
	goto L168
L168:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599)+1)))
	if v630 == int32(0) {
		v640 = v626
		v642 = v628
		goto L162
	} else {
		goto L172
	}
L169:
	;
	v615 = int32(0)
	goto L171
L170:
	;
	v615 = v609
	goto L171
L171:
	;
	v621 = int32(1)
	v625 = v601 + v621
	v626 = (int32(0)-v614)&v609 + v602
	v627 = v603 - v621
	v628 = v615 + v604
	goto L168
L172:
	;
	v633 = int32(1)
	if base.Ui32(v633) < base.Ui32(v627) {
		v599 = v599 + v633
		v600 = v630
		v601 = v625
		v602 = v626
		v603 = v627
		v604 = v628
		goto L164
	} else {
		goto L173
	}
L173:
	;
	goto L165
L174:
	;
	v652 = v644 - v649
	goto L176
L175:
	;
	v652 = int32(0)
	goto L176
L176:
	;
	goto L158
L177:
	;
	if v676 == int32(0) {
		v1050 = v675
		goto L1
	} else {
		goto L178
	}
L178:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L105
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = l0
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_string2ean[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v688
	F_errmsg(m, int32(_a_F_string2ean_5), v23)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L105
	} else {
		goto L180
	}
L180:
	;
	F_errsave_finish(m, l1, int32(_a_F_string2ean_3), int32(902), int32(_a_F_string2ean_4))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L105
	} else {
		goto L181
	}
L181:
	;
	v1050 = v675
	goto L1
L182:
	;
	if v699 == int32(0) {
		v1050 = v698
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L105
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = l0
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_string2ean[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v711
	F_errmsg(m, int32(_a_F_string2ean_6), v23-int32(-64))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L105
	} else {
		goto L185
	}
L185:
	;
	F_errsave_finish(m, l1, int32(_a_F_string2ean_3), int32(914), int32(_a_F_string2ean_4))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L105
	} else {
		goto L186
	}
L186:
	;
	v1050 = v698
	goto L1
L187:
	;
	v751 = int32(*(*int8)(unsafe.Add(mBase, uint32(v742))))
	if v751 != 0 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v760 = base.B2i32(v751 == int32(77))
	if v751 == int32(77) {
		goto L193
	} else {
		goto L194
	}
L189:
	;
	if v751 < int32(33) {
		v742 = v742 + int32(1)
		goto L187
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	goto L188
L192:
	;
	goto L191
L193:
	;
	v761 = int32(3)
	goto L195
L194:
	;
	v761 = int32(0)
	goto L195
L195:
	;
	if v751 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v838 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v742)+13)) = uint8(v838)
	v845 = base.I32_rem_u_s(v825*int32(3)+v830, int32(10))
	if v845 != 0 {
		goto L210
	} else {
		goto L211
	}
L197:
	;
	v825 = v761
	v830 = int32(0)
	goto L196
L198:
	;
	goto L199
L199:
	;
	v773 = v751
	v774 = v761
	v778 = v742
	v779 = int32(0)
	v780 = v760
	v782 = int32(13)
	goto L200
L200:
	;
	v790 = (v773 - int32(48)) & int32(255)
	if base.Ui32(v790) <= base.Ui32(int32(9)) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v825 = v806
	v830 = v807
	goto L196
L202:
	;
	v795 = v780 & int32(1)
	if v795 != 0 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	v806 = v774
	v807 = v779
	v808 = v780
	v809 = v782
	goto L204
L204:
	;
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778)+1)))
	if v811 == int32(0) {
		v825 = v806
		v830 = v807
		goto L196
	} else {
		goto L208
	}
L205:
	;
	v796 = int32(0)
	goto L207
L206:
	;
	v796 = v790
	goto L207
L207:
	;
	v798 = int32(1)
	v806 = (int32(0)-v795)&v790 + v774
	v807 = v796 + v779
	v808 = v780 + v798
	v809 = v782 - v798
	goto L204
L208:
	;
	v814 = int32(1)
	if base.Ui32(v814) < base.Ui32(v809) {
		v773 = v811
		v774 = v806
		v778 = v778 + v814
		v779 = v807
		v780 = v808
		v782 = v809
		goto L200
	} else {
		goto L209
	}
L209:
	;
	goto L201
L210:
	;
	v848 = int32(58) - v845
	goto L212
L211:
	;
	v848 = int32(48)
	goto L212
L212:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v742)+12)) = uint8(v848)
	if (v727|v174)&int32(1) != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	if v751 != 0 {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	goto L215
L215:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_string2ean[8])))
	if v922 == int32(1) {
		goto L225
	} else {
		goto L226
	}
L216:
	;
	v858 = v751
	v861 = v742
	v872 = v20
	goto L219
L217:
	;
	v912 = int64(0)
	goto L218
L218:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v912 | base.I64_extend_i32_u(v727^int32(-1))&int64(1)
	v1050 = int32(1)
	goto L1
L219:
	;
	if base.Ui32((v858-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v912 = v885 << (uint(int64(1)) % 64)
	goto L218
L221:
	;
	v885 = v872*int64(10) + base.I64_extend_i32_u(v858)&int64(15)
	goto L223
L222:
	;
	v885 = v872
	goto L223
L223:
	;
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+1)))
	if v886 != 0 {
		v858 = v886
		v861 = v861 + int32(1)
		v872 = v885
		goto L219
	} else {
		goto L224
	}
L224:
	;
	goto L220
L225:
	;
	if v751 != 0 {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	goto L227
L227:
	;
	v989 = int32(0)
	v990 = F_errsave_start(m, l1)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L105
	} else {
		goto L237
	}
L228:
	;
	v930 = v751
	v933 = v742
	v944 = v20
	goto L231
L229:
	;
	v986 = int64(1)
	goto L230
L230:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v986
	v1050 = int32(1)
	goto L1
L231:
	;
	if base.Ui32((v930-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v961 = int64(1)
	v986 = v957<<(uint(v961)%64) | v961
	goto L230
L233:
	;
	v957 = v944*int64(10) + base.I64_extend_i32_u(v930)&int64(15)
	goto L235
L234:
	;
	v957 = v944
	goto L235
L235:
	;
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933)+1)))
	if v958 != 0 {
		v930 = v958
		v933 = v933 + int32(1)
		v944 = v957
		goto L231
	} else {
		goto L236
	}
L236:
	;
	goto L232
L237:
	;
	if v728 == int32(-1) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	if v990 == int32(0) {
		v1050 = v989
		goto L1
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	if v990 == int32(0) {
		v1050 = v989
		goto L1
	} else {
		goto L245
	}
L241:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L105
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = l0
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_string2ean[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v1004
	F_errmsg(m, int32(_a_F_string2ean_7), v23+int32(16))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L105
	} else {
		goto L243
	}
L243:
	;
	F_errsave_finish(m, l1, int32(_a_F_string2ean_3), int32(888), int32(_a_F_string2ean_4))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L105
	} else {
		goto L244
	}
L244:
	;
	v1050 = v989
	goto L1
L245:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L105
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = l0
	if v728 == int32(10) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1027 = int32(88)
	goto L249
L248:
	;
	v1027 = v728 + int32(48)
	goto L249
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v1027
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_string2ean[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v1033
	F_errmsg(m, int32(_a_F_string2ean_8), v23+int32(32))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L105
	} else {
		goto L250
	}
L250:
	;
	F_errsave_finish(m, l1, int32(_a_F_string2ean_3), int32(895), int32(_a_F_string2ean_4))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L105
	} else {
		goto L251
	}
L251:
	;
	v1050 = v989
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
func F_strtoul(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(4294967295))
	return base.I32_wrap_i64(v5)
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
			v41 = F_expression_tree_mutator_impl(m, l0, int32(878), l1)
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
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
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
		v148 = v20
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v14 - int32(-64)
	return v174
L4:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v154 = v12 + int32(-40)
	v155 = F_WALRead(m, l0, l4, l1, v148, v152, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
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
	v148 = base.I32_wrap_i64(v31 - l1)
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
	v174 = int32(-1)
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
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[0])) = int32(0)
	v99 = F_GetLatestLSN(m, v12+int32(-40))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v101 == v102 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if base.Ui64(v137) < base.Ui64(v22) {
		v31 = v137
		goto L6
	} else {
		goto L41
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v99
	v137 = v99
	goto L31
L33:
	;
	goto L34
L34:
	;
	v105 = F_readTimeLineHistory(m, v101)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)) = uint8(v107)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v111 = F_tliSwitchPoint(m, v109, v105, int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v111
	v116 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v23)+8))
	if v116 == int32(0) {
		v137 = v118
		goto L31
	} else {
		goto L38
	}
L38:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v121
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)) = uint32(v118)
	v125 = int64(base.Ui64(v118) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+4)) = uint32(v125)
	F_errmsg_internal(m, int32(_a_F_summarizer_read_local_xlog_page_1), v14)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_summarizer_read_local_xlog_page_2), int32(1584), int32(_a_F_summarizer_read_local_xlog_page_3))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v23)+8))
	v137 = v135
	goto L31
L41:
	;
	v148 = v20
	goto L4
L42:
	;
	if v155 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_WALReadRaiseError(m, v154)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v161 = int32(_a_F_summarizer_read_local_xlog_page_4)
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[0])) = v163 + int32(1)
	v174 = v148
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
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v223 int32
	_ = v223
	v11 = m.G0
	v12 = int32(16)
	v13 = v11 - v12
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l2
	v16 = m.G0
	v18 = v16 - v12
	m.G0 = v18
	if base.Ui32(int32(1023)) < base.Ui32(l0) {
		v223 = int32(16)
		m.G0 = v18 + v223
		m.G0 = v13 + v223
		return
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[0]))
		if v23&(int32(1)<<(uint(l0&int32(7))%32)) == int32(0) {
			v223 = int32(16)
			m.G0 = v18 + v223
			m.G0 = v13 + v223
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
							v223 = int32(16)
							m.G0 = v18 + v223
							m.G0 = v13 + v223
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
							v131 = int32(0)
							v133 = F_sendto(m, v128, v33-int32(-64), v126, v131, v131)
							mBase = m.M
							if int32(0) <= v133 {
								v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
								if v193&int32(32) == int32(0) {
									m.G0 = v33 + int32(1168)
									v223 = int32(16)
									m.G0 = v18 + v223
									m.G0 = v13 + v223
									return
								} else {
									v198 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
									*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v198
									*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v198 + (v33 - int32(-64))
									F_dprintf(m, int32(2), v33)
									mBase = m.M
									v207 = m.ExcPending
									if v207 != 0 {
										return
									} else {
										m.G0 = v33 + int32(1168)
										v223 = int32(16)
										m.G0 = v18 + v223
										m.G0 = v13 + v223
										return
									}
								}
							} else {
								v138 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[1]))
								if base.B2i32(base.Ui32(v138-int32(14)) < base.Ui32(int32(2)))|base.B2i32(v138 == int32(53)) != 0 {
									v149 = int32(1)
								} else {
									v149 = base.B2i32(v138 == int32(64))
								}
								if v149 == int32(0) {
									v167 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
									if v167&int32(2) == int32(0) {
										v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
										if v193&int32(32) == int32(0) {
											m.G0 = v33 + int32(1168)
											v223 = int32(16)
											m.G0 = v18 + v223
											m.G0 = v13 + v223
											return
										} else {
											v198 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v198
											*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v198 + (v33 - int32(-64))
											F_dprintf(m, int32(2), v33)
											mBase = m.M
											v207 = m.ExcPending
											if v207 != 0 {
												return
											} else {
												m.G0 = v33 + int32(1168)
												v223 = int32(16)
												m.G0 = v18 + v223
												m.G0 = v13 + v223
												return
											}
										}
									} else {
										v174 = int32(0)
										v175 = F_open(m, int32(_a_F_syslog_7), int32(_a_F_syslog_8), v174)
										mBase = m.M
										if v175 < v174 {
											v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
											if v193&int32(32) == int32(0) {
												m.G0 = v33 + int32(1168)
												v223 = int32(16)
												m.G0 = v18 + v223
												m.G0 = v13 + v223
												return
											} else {
												v198 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v198
												*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v198 + (v33 - int32(-64))
												F_dprintf(m, int32(2), v33)
												mBase = m.M
												v207 = m.ExcPending
												if v207 != 0 {
													return
												} else {
													m.G0 = v33 + int32(1168)
													v223 = int32(16)
													m.G0 = v18 + v223
													m.G0 = v13 + v223
													return
												}
											}
										} else {
											v178 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v126 - v178
											*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v178 + (v33 - int32(-64))
											F_dprintf(m, v175, v33+int32(16))
											mBase = m.M
											v188 = m.ExcPending
											if v188 != 0 {
												return
											} else {
												v189 = F_close(m, v175)
												mBase = m.M
												v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
												if v193&int32(32) == int32(0) {
													m.G0 = v33 + int32(1168)
													v223 = int32(16)
													m.G0 = v18 + v223
													m.G0 = v13 + v223
													return
												} else {
													v198 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v198
													*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v198 + (v33 - int32(-64))
													F_dprintf(m, int32(2), v33)
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return
													} else {
														m.G0 = v33 + int32(1168)
														v223 = int32(16)
														m.G0 = v18 + v223
														m.G0 = v13 + v223
														return
													}
												}
											}
										}
									}
								} else {
									v153 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[2]))
									v154 = F_connect(m, v153)
									mBase = m.M
									if v154 < int32(0) {
										v167 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
										if v167&int32(2) == int32(0) {
											v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
											if v193&int32(32) == int32(0) {
												m.G0 = v33 + int32(1168)
												v223 = int32(16)
												m.G0 = v18 + v223
												m.G0 = v13 + v223
												return
											} else {
												v198 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v198
												*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v198 + (v33 - int32(-64))
												F_dprintf(m, int32(2), v33)
												mBase = m.M
												v207 = m.ExcPending
												if v207 != 0 {
													return
												} else {
													m.G0 = v33 + int32(1168)
													v223 = int32(16)
													m.G0 = v18 + v223
													m.G0 = v13 + v223
													return
												}
											}
										} else {
											v174 = int32(0)
											v175 = F_open(m, int32(_a_F_syslog_7), int32(_a_F_syslog_8), v174)
											mBase = m.M
											if v175 < v174 {
												v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
												if v193&int32(32) == int32(0) {
													m.G0 = v33 + int32(1168)
													v223 = int32(16)
													m.G0 = v18 + v223
													m.G0 = v13 + v223
													return
												} else {
													v198 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v198
													*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v198 + (v33 - int32(-64))
													F_dprintf(m, int32(2), v33)
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return
													} else {
														m.G0 = v33 + int32(1168)
														v223 = int32(16)
														m.G0 = v18 + v223
														m.G0 = v13 + v223
														return
													}
												}
											} else {
												v178 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v126 - v178
												*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v178 + (v33 - int32(-64))
												F_dprintf(m, v175, v33+int32(16))
												mBase = m.M
												v188 = m.ExcPending
												if v188 != 0 {
													return
												} else {
													v189 = F_close(m, v175)
													mBase = m.M
													v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
													if v193&int32(32) == int32(0) {
														m.G0 = v33 + int32(1168)
														v223 = int32(16)
														m.G0 = v18 + v223
														m.G0 = v13 + v223
														return
													} else {
														v198 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
														*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v198
														*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v198 + (v33 - int32(-64))
														F_dprintf(m, int32(2), v33)
														mBase = m.M
														v207 = m.ExcPending
														if v207 != 0 {
															return
														} else {
															m.G0 = v33 + int32(1168)
															v223 = int32(16)
															m.G0 = v18 + v223
															m.G0 = v13 + v223
															return
														}
													}
												}
											}
										}
									} else {
										v158 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[2]))
										v161 = int32(0)
										v163 = F_sendto(m, v158, v33-int32(-64), v126, v161, v161)
										mBase = m.M
										if int32(0) <= v163 {
											v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
											if v193&int32(32) == int32(0) {
												m.G0 = v33 + int32(1168)
												v223 = int32(16)
												m.G0 = v18 + v223
												m.G0 = v13 + v223
												return
											} else {
												v198 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v198
												*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v198 + (v33 - int32(-64))
												F_dprintf(m, int32(2), v33)
												mBase = m.M
												v207 = m.ExcPending
												if v207 != 0 {
													return
												} else {
													m.G0 = v33 + int32(1168)
													v223 = int32(16)
													m.G0 = v18 + v223
													m.G0 = v13 + v223
													return
												}
											}
										} else {
											v167 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
											if v167&int32(2) == int32(0) {
												v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
												if v193&int32(32) == int32(0) {
													m.G0 = v33 + int32(1168)
													v223 = int32(16)
													m.G0 = v18 + v223
													m.G0 = v13 + v223
													return
												} else {
													v198 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v198
													*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v198 + (v33 - int32(-64))
													F_dprintf(m, int32(2), v33)
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return
													} else {
														m.G0 = v33 + int32(1168)
														v223 = int32(16)
														m.G0 = v18 + v223
														m.G0 = v13 + v223
														return
													}
												}
											} else {
												v174 = int32(0)
												v175 = F_open(m, int32(_a_F_syslog_7), int32(_a_F_syslog_8), v174)
												mBase = m.M
												if v175 < v174 {
													v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
													if v193&int32(32) == int32(0) {
														m.G0 = v33 + int32(1168)
														v223 = int32(16)
														m.G0 = v18 + v223
														m.G0 = v13 + v223
														return
													} else {
														v198 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
														*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v198
														*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v198 + (v33 - int32(-64))
														F_dprintf(m, int32(2), v33)
														mBase = m.M
														v207 = m.ExcPending
														if v207 != 0 {
															return
														} else {
															m.G0 = v33 + int32(1168)
															v223 = int32(16)
															m.G0 = v18 + v223
															m.G0 = v13 + v223
															return
														}
													}
												} else {
													v178 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v126 - v178
													*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v178 + (v33 - int32(-64))
													F_dprintf(m, v175, v33+int32(16))
													mBase = m.M
													v188 = m.ExcPending
													if v188 != 0 {
														return
													} else {
														v189 = F_close(m, v175)
														mBase = m.M
														v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
														if v193&int32(32) == int32(0) {
															m.G0 = v33 + int32(1168)
															v223 = int32(16)
															m.G0 = v18 + v223
															m.G0 = v13 + v223
															return
														} else {
															v198 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
															*(*int32)(unsafe.Add(mBase, uint32(v33))) = v126 - v198
															*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v198 + (v33 - int32(-64))
															F_dprintf(m, int32(2), v33)
															mBase = m.M
															v207 = m.ExcPending
															if v207 != 0 {
																return
															} else {
																m.G0 = v33 + int32(1168)
																v223 = int32(16)
																m.G0 = v18 + v223
																m.G0 = v13 + v223
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
