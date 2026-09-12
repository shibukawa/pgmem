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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
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
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
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
					v96 = v38
					v97 = v37
					v99 = v3
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[3])))
					v51 = int32(0)
					v54 = v38
					v55 = v37
					v57 = v3
					v60 = v37 - int32(2048)
					for {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v48+v51<<(uint(int32(2))%32))))
						v70 = v20 + int32(_a_F_SICleanupQueue_0) + v67<<(uint(int32(4))%32)
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)))
						if v71 != 0 {
							v84 = v54
							v85 = v55
							v87 = v57
							v88 = v60
						} else {
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+11)))
							if v72 != 0 {
								v84 = v54
								v85 = v55
								v87 = v57
								v88 = v60
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
								if v73 < l1+v37-int32(_a_F_SICleanupQueue_1) {
									v75 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)) = uint8(v75)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[2])))
									v84 = v77
									v85 = v55
									v87 = v57
									v88 = v60
								} else {
									if v73 < v55 {
										v79 = v73
									} else {
										v79 = v55
									}
									if v60 <= v73 {
										v84 = v54
										v85 = v79
										v87 = v57
										v88 = v60
									} else {
										v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+9)))
										if v81 != 0 {
											v82 = v60
										} else {
											v82 = v73
										}
										if v81 != 0 {
											v83 = v57
										} else {
											v83 = v70
										}
										v84 = v54
										v85 = v79
										v87 = v83
										v88 = v82
									}
								}
							}
						}
						v90 = v51 + int32(1)
						if v90 < v84 {
							v51 = v90
							v54 = v84
							v55 = v85
							v57 = v87
							v60 = v88
							continue
						} else {
							break
						}
						break
					}
					v96 = v84
					v97 = v85
					v99 = v87
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v97
				if v97 < int32(1073741824) {
					v198 = v97
					v201 = v37
				} else {
					v109 = int32(1073741824)
					v110 = v37 - v109
					*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v110
					v113 = v97 - v109
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v113
					if v96 <= int32(0) {
						v198 = v113
						v201 = v110
					} else {
						v117 = int32(1)
						v120 = v20 + int32(_a_F_SICleanupQueue_2)
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[3])))
						v122 = int32(0)
						if v96 != v117 {
							v129 = v122
							v132 = int32(0)
							for {
								v142 = int32(2)
								v144 = v121 + v129<<(uint(v142)%32)
								v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
								v146 = int32(4)
								v148 = v120 + v145<<(uint(v146)%32)
								v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
								v150 = int32(1073741824)
								*(*int32)(unsafe.Add(mBase, uint32(v148))) = v149 - v150
								v153 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
								v156 = v120 + v153<<(uint(v146)%32)
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
								*(*int32)(unsafe.Add(mBase, uint32(v156))) = v157 - v150
								v162 = v129 + v142
								v164 = v132 + v142
								if v164 != v96&int32(2147483646) {
									v129 = v162
									v132 = v164
									continue
								} else {
									break
								}
								break
							}
							v167 = v162
						} else {
							v167 = v122
						}
						if v96&v117 == int32(0) {
							v198 = v113
							v201 = v110
						} else {
							v185 = *(*int32)(unsafe.Add(mBase, uint32(v121+v167<<(uint(int32(2))%32))))
							v188 = v120 + v185<<(uint(int32(4))%32)
							v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
							*(*int32)(unsafe.Add(mBase, uint32(v188))) = v189 - int32(1073741824)
							v198 = v113
							v201 = v110
						}
					}
				}
				v207 = int32(2048)
				v208 = v201 - v198
				if v208 < v207 {
					v215 = v207
				} else {
					v215 = v208&int32(2147483392) + int32(256)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v215
				if v99 != 0 {
					v217 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v99)+9)) = uint8(v217)
					v219 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
					v221 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
					F_LWLockRelease(m, v221+int32(640))
					mBase = m.M
					v225 = m.ExcPending
					if v225 != 0 {
						return
					} else {
						v227 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
						F_LWLockRelease(m, v227+int32(768))
						mBase = m.M
						v231 = m.ExcPending
						if v231 != 0 {
							return
						} else {
							v239 = F_errstart(m, int32(11), int32(0))
							mBase = m.M
							v240 = m.ExcPending
							if v240 != 0 {
								return
							} else {
								if v239 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v17))) = v219
									F_errmsg_internal(m, int32(_a_F_SICleanupQueue_3), v17)
									mBase = m.M
									v244 = m.ExcPending
									if v244 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SICleanupQueue_4), int32(673), int32(_a_F_SICleanupQueue_5))
										mBase = m.M
										v249 = m.ExcPending
										if v249 != 0 {
											return
										} else {
											v251 = F_SendProcSignal(m, v219, int32(0), (v99-v20-int32(_a_F_SICleanupQueue_0))>>(uint(int32(4))%32))
											mBase = m.M
											v252 = m.ExcPending
											if v252 != 0 {
												return
											} else {
												if l0 == int32(0) {
													m.G0 = v17 + int32(16)
													return
												} else {
													v256 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
													v260 = F_LWLockAcquire(m, v256+int32(768), int32(0))
													mBase = m.M
													v261 = m.ExcPending
													if v261 != 0 {
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
									v251 = F_SendProcSignal(m, v219, int32(0), (v99-v20-int32(_a_F_SICleanupQueue_0))>>(uint(int32(4))%32))
									mBase = m.M
									v252 = m.ExcPending
									if v252 != 0 {
										return
									} else {
										if l0 == int32(0) {
											m.G0 = v17 + int32(16)
											return
										} else {
											v256 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
											v260 = F_LWLockAcquire(m, v256+int32(768), int32(0))
											mBase = m.M
											v261 = m.ExcPending
											if v261 != 0 {
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
					v263 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
					F_LWLockRelease(m, v263+int32(640))
					mBase = m.M
					v267 = m.ExcPending
					if v267 != 0 {
						return
					} else {
						if l0 != 0 {
							m.G0 = v17 + int32(16)
							return
						} else {
							v269 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
							F_LWLockRelease(m, v269+int32(768))
							mBase = m.M
							v273 = m.ExcPending
							if v273 != 0 {
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
				v96 = v38
				v97 = v37
				v99 = v3
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[3])))
				v51 = int32(0)
				v54 = v38
				v55 = v37
				v57 = v3
				v60 = v37 - int32(2048)
				for {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v48+v51<<(uint(int32(2))%32))))
					v70 = v20 + int32(_a_F_SICleanupQueue_0) + v67<<(uint(int32(4))%32)
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)))
					if v71 != 0 {
						v84 = v54
						v85 = v55
						v87 = v57
						v88 = v60
					} else {
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+11)))
						if v72 != 0 {
							v84 = v54
							v85 = v55
							v87 = v57
							v88 = v60
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
							if v73 < l1+v37-int32(_a_F_SICleanupQueue_1) {
								v75 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)) = uint8(v75)
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[2])))
								v84 = v77
								v85 = v55
								v87 = v57
								v88 = v60
							} else {
								if v73 < v55 {
									v79 = v73
								} else {
									v79 = v55
								}
								if v60 <= v73 {
									v84 = v54
									v85 = v79
									v87 = v57
									v88 = v60
								} else {
									v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+9)))
									if v81 != 0 {
										v82 = v60
									} else {
										v82 = v73
									}
									if v81 != 0 {
										v83 = v57
									} else {
										v83 = v70
									}
									v84 = v54
									v85 = v79
									v87 = v83
									v88 = v82
								}
							}
						}
					}
					v90 = v51 + int32(1)
					if v90 < v84 {
						v51 = v90
						v54 = v84
						v55 = v85
						v57 = v87
						v60 = v88
						continue
					} else {
						break
					}
					break
				}
				v96 = v84
				v97 = v85
				v99 = v87
			}
			*(*int32)(unsafe.Add(mBase, uint32(v20))) = v97
			if v97 < int32(1073741824) {
				v198 = v97
				v201 = v37
			} else {
				v109 = int32(1073741824)
				v110 = v37 - v109
				*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v110
				v113 = v97 - v109
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v113
				if v96 <= int32(0) {
					v198 = v113
					v201 = v110
				} else {
					v117 = int32(1)
					v120 = v20 + int32(_a_F_SICleanupQueue_2)
					v121 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_SICleanupQueue[3])))
					v122 = int32(0)
					if v96 != v117 {
						v129 = v122
						v132 = int32(0)
						for {
							v142 = int32(2)
							v144 = v121 + v129<<(uint(v142)%32)
							v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
							v146 = int32(4)
							v148 = v120 + v145<<(uint(v146)%32)
							v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
							v150 = int32(1073741824)
							*(*int32)(unsafe.Add(mBase, uint32(v148))) = v149 - v150
							v153 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
							v156 = v120 + v153<<(uint(v146)%32)
							v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
							*(*int32)(unsafe.Add(mBase, uint32(v156))) = v157 - v150
							v162 = v129 + v142
							v164 = v132 + v142
							if v164 != v96&int32(2147483646) {
								v129 = v162
								v132 = v164
								continue
							} else {
								break
							}
							break
						}
						v167 = v162
					} else {
						v167 = v122
					}
					if v96&v117 == int32(0) {
						v198 = v113
						v201 = v110
					} else {
						v185 = *(*int32)(unsafe.Add(mBase, uint32(v121+v167<<(uint(int32(2))%32))))
						v188 = v120 + v185<<(uint(int32(4))%32)
						v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
						*(*int32)(unsafe.Add(mBase, uint32(v188))) = v189 - int32(1073741824)
						v198 = v113
						v201 = v110
					}
				}
			}
			v207 = int32(2048)
			v208 = v201 - v198
			if v208 < v207 {
				v215 = v207
			} else {
				v215 = v208&int32(2147483392) + int32(256)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v215
			if v99 != 0 {
				v217 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v99)+9)) = uint8(v217)
				v219 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
				v221 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
				F_LWLockRelease(m, v221+int32(640))
				mBase = m.M
				v225 = m.ExcPending
				if v225 != 0 {
					return
				} else {
					v227 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
					F_LWLockRelease(m, v227+int32(768))
					mBase = m.M
					v231 = m.ExcPending
					if v231 != 0 {
						return
					} else {
						v239 = F_errstart(m, int32(11), int32(0))
						mBase = m.M
						v240 = m.ExcPending
						if v240 != 0 {
							return
						} else {
							if v239 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v17))) = v219
								F_errmsg_internal(m, int32(_a_F_SICleanupQueue_3), v17)
								mBase = m.M
								v244 = m.ExcPending
								if v244 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SICleanupQueue_4), int32(673), int32(_a_F_SICleanupQueue_5))
									mBase = m.M
									v249 = m.ExcPending
									if v249 != 0 {
										return
									} else {
										v251 = F_SendProcSignal(m, v219, int32(0), (v99-v20-int32(_a_F_SICleanupQueue_0))>>(uint(int32(4))%32))
										mBase = m.M
										v252 = m.ExcPending
										if v252 != 0 {
											return
										} else {
											if l0 == int32(0) {
												m.G0 = v17 + int32(16)
												return
											} else {
												v256 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
												v260 = F_LWLockAcquire(m, v256+int32(768), int32(0))
												mBase = m.M
												v261 = m.ExcPending
												if v261 != 0 {
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
								v251 = F_SendProcSignal(m, v219, int32(0), (v99-v20-int32(_a_F_SICleanupQueue_0))>>(uint(int32(4))%32))
								mBase = m.M
								v252 = m.ExcPending
								if v252 != 0 {
									return
								} else {
									if l0 == int32(0) {
										m.G0 = v17 + int32(16)
										return
									} else {
										v256 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
										v260 = F_LWLockAcquire(m, v256+int32(768), int32(0))
										mBase = m.M
										v261 = m.ExcPending
										if v261 != 0 {
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
				v263 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
				F_LWLockRelease(m, v263+int32(640))
				mBase = m.M
				v267 = m.ExcPending
				if v267 != 0 {
					return
				} else {
					if l0 != 0 {
						m.G0 = v17 + int32(16)
						return
					} else {
						v269 = *(*int32)(unsafe.Add(mBase, _c_F_SICleanupQueue[1]))
						F_LWLockRelease(m, v269+int32(768))
						mBase = m.M
						v273 = m.ExcPending
						if v273 != 0 {
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
func F_SetQuitSignalReason(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_SetQuitSignalReason[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = l0
	return
}
func F_SetRemoteDestReceiverParams(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
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
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	v78 = m.ExcPending
	if v78 != 0 {
		goto L16
	} else {
		goto L24
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
		goto L1
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
		goto L2
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(_a_F_SplitColQualList_0), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	F_scanner_errposition(m, v67, l3)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_SplitColQualList_1), int32(_a_F_SplitColQualList_2), int32(_a_F_SplitColQualList_3))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v79
	F_errmsg_internal(m, int32(_a_F_SplitColQualList_4), v11)
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
	var v6 int32
	_ = v6
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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyFreeBuffer[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(1)
	if v6 != 0 {
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
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(0)
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
		*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(0)
		return
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
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
	v24 = v13 + int32(16)
	v25 = int32(2)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v32 = m.Wasi_snapshot_preview1.Fd_write(m, v26, v24, v25, v13+int32(12))
	mBase = m.M
	if v32 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v148
L2:
	;
	v125 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v125
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v129 | int32(32)
	if v122 == int32(2) {
		v148 = v125
		goto L1
	} else {
		goto L30
	}
L3:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v109
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v109 + v112
	v148 = l2
	goto L1
L4:
	;
	if v93 != int32(-1) {
		v120 = v92
		v122 = v94
		goto L2
	} else {
		goto L29
	}
L5:
	;
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v39 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___stdio_write[0])) = v32
	v39 = int32(-1)
	goto L5
L9:
	;
	v92 = v24
	v93 = v22
	v94 = v25
	goto L4
L10:
	;
	goto L11
L11:
	;
	v44 = v24
	v46 = v22
	v47 = v25
	goto L12
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v46 == v50 {
		goto L3
	} else {
		goto L14
	}
L13:
	;
	v92 = v58
	v93 = v72
	v94 = v74
	goto L4
L14:
	;
	if v50 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v120 = v44
	v122 = v47
	goto L2
L16:
	;
	goto L17
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v55 = base.B2i32(base.Ui32(v54) < base.Ui32(v50))
	v58 = v44 + v55<<(uint(int32(3))%32)
	if base.Ui32(v54) < base.Ui32(v50) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = v54
	goto L20
L19:
	;
	v60 = int32(0)
	goto L20
L20:
	;
	v61 = v50 - v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v61 + v62
	if base.Ui32(v54) < base.Ui32(v50) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v67 = int32(12)
	goto L23
L22:
	;
	v67 = int32(4)
	goto L23
L23:
	;
	v68 = v44 + v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v69 - v61
	v72 = v46 - v50
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v74 = v47 - v55
	v77 = m.Wasi_snapshot_preview1.Fd_write(m, v73, v58, v74, v13+int32(12))
	mBase = m.M
	if v77 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v84 == int32(0) {
		v44 = v58
		v46 = v72
		v47 = v74
		goto L12
	} else {
		goto L28
	}
L25:
	;
	v84 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___stdio_write[0])) = v77
	v84 = int32(-1)
	goto L24
L28:
	;
	goto L13
L29:
	;
	goto L3
L30:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v148 = l2 - v136
	goto L1
}
func F___strftime_l(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
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
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v355 int64
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
	var v598 int64
	_ = v598
	var v602 int64
	_ = v602
	var v606 int64
	_ = v606
	var v608 int64
	_ = v608
	var v611 int64
	_ = v611
	var v613 int64
	_ = v613
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v665 int64
	_ = v665
	var v669 int32
	_ = v669
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v796 int64
	_ = v796
	var v800 int64
	_ = v800
	var v803 int32
	_ = v803
	var v810 int32
	_ = v810
	var v812 int64
	_ = v812
	var v816 int32
	_ = v816
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v1018 int32
	_ = v1018
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1171 int32
	_ = v1171
	var v1184 int32
	_ = v1184
	var v1204 int32
	_ = v1204
	var v1215 int32
	_ = v1215
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1243 int32
	_ = v1243
	v6 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(128)
	m.G0 = v29
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = l2
	v40 = v6
	goto L5
L2:
	;
	v1243 = v6
	goto L3
L3:
	;
	m.G0 = v29 + int32(128)
	return v1243
L4:
	;
	v1234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v1215))) = uint8(v1234)
	v1243 = v1232
	goto L3
L5:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v57 != int32(37) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	if l1 == v1184 {
		goto L328
	} else {
		goto L329
	}
L7:
	;
	goto L6
L8:
	;
	if base.Ui32(v1171) < base.Ui32(l1) {
		v33 = v1147 + int32(1)
		v40 = v1171
		goto L5
	} else {
		goto L327
	}
L9:
	;
	v81 = v76 & int32(255)
	v84 = v33 + v77 + base.B2i32(v81 == int32(43))
	v85 = int32(*(*int8)(unsafe.Add(mBase, uint32(v84))))
	if base.Ui32(v85-int32(48)) <= base.Ui32(int32(9)) {
		goto L20
	} else {
		goto L21
	}
L10:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+2)))
	v76 = v74
	v77 = int32(2)
	v78 = v62
	goto L9
L11:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v40))) = uint8(v57)
	v1147 = v33
	v1171 = v40 + int32(1)
	goto L8
L12:
	;
	if v57 != 0 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v60 = int32(0)
	v61 = int32(1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	switch v62 - int32(45) {
	case 0, 3:
		goto L10
	case 1, 2:
		v76 = v62
		v77 = v61
		v78 = v60
		goto L9
	default:
		goto L16
	}
L15:
	;
	v1215 = v40
	v1232 = v40
	goto L4
L16:
	;
	if v62 == int32(95) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	if v62 != 0 {
		v76 = v62
		v77 = v61
		v78 = v60
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	v101 = int32(0)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v104 = v102 - int32(67)
	if base.Ui32(int32(22)) < base.Ui32(v104) {
		v114 = v101
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v94 = F_strtox_2(m, v84, v29+int32(12), int32(10), int64(4294967295))
	mBase = m.M
	goto L23
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v84
	v99 = int32(0)
	v100 = v84
	goto L19
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v99 = base.I32_wrap_i64(v94)
	v100 = v96
	goto L19
L24:
	;
	if v102 == int32(79) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	if int32(1)<<(uint(v104)%32)&int32(_a_F___strftime_l_0) == int32(0) {
		v114 = v101
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if v99 != 0 {
		v114 = v99
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v114 = base.B2i32(v100 != v84)
	goto L24
L28:
	;
	v125 = v29 + int32(16)
	v127 = m.G0
	v129 = v127 - int32(80)
	m.G0 = v129
	v132 = int32(48)
	v135 = v29 + int32(124)
	v136 = base.I32_extend8_s(v122)
	switch v136 - int32(37) {
	case 0:
		goto L44
	default:
		v908 = int32(0)
		goto L32
	case 28:
		goto L76
	case 29:
		goto L74
	case 30:
		goto L73
	case 31:
		v786 = int32(_a_F___strftime_l_1)
		goto L38
	case 33:
		goto L70
	case 34, 66:
		goto L69
	case 35:
		goto L68
	case 36:
		goto L67
	case 40:
		goto L64
	case 45:
		goto L61
	case 46:
		goto L59
	case 47:
		goto L57
	case 48:
		goto L55
	case 49:
		goto L53
	case 50:
		goto L54
	case 51:
		goto L49
	case 52:
		goto L47
	case 53:
		goto L45
	case 60:
		goto L77
	case 61, 67:
		goto L75
	case 62:
		v728 = int32(_a_F___strftime_l_2)
		goto L39
	case 63:
		v163 = v132
		goto L71
	case 64:
		goto L72
	case 69:
		goto L66
	case 72:
		goto L65
	case 73:
		goto L63
	case 75:
		goto L62
	case 77:
		goto L40
	case 78:
		goto L60
	case 79:
		goto L58
	case 80:
		goto L56
	case 82:
		goto L52
	case 83:
		goto L50
	case 84:
		goto L48
	case 85:
		goto L46
	}
L29:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	v122 = v119
	v123 = v100 + int32(1)
	goto L28
L30:
	;
	if v102 == int32(69) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v122 = v102
	v123 = v100
	goto L28
L32:
	;
	m.G0 = v129 + int32(80)
	if v908 == int32(0) {
		v1184 = v40
		goto L7
	} else {
		goto L285
	}
L33:
	;
	if v849&int32(3) == int32(0) {
		v873 = v849
		goto L270
	} else {
		goto L271
	}
L34:
	;
	v849 = int32(_a_F___strftime_l_3)
	goto L33
L35:
	;
	if v78 != 0 {
		goto L258
	} else {
		goto L259
	}
L36:
	;
	v803 = int32(4)
	v810 = v132
	v812 = v800
	goto L35
L37:
	;
	v803 = int32(2)
	v810 = v795
	v812 = v796
	goto L35
L38:
	;
	v789 = F___strftime_l(m, v125, int32(100), v786, l3, l4)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L188
	} else {
		goto L253
	}
L39:
	;
	if v728 == int32(14) {
		goto L226
	} else {
		goto L227
	}
L40:
	;
	v728 = int32(_a_F___strftime_l_4)
	goto L39
L41:
	;
	if v669 == int32(14) {
		goto L198
	} else {
		goto L199
	}
L42:
	;
	v669 = v139 | int32(_a_F___strftime_l_5)
	goto L41
L43:
	;
	v665 = base.I64_rem_s(v290, int64(100))
	v795 = v132
	v796 = v665
	goto L37
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(1)
	v908 = int32(_a_F___strftime_l_6)
	goto L32
L45:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v653 < int32(0) {
		goto L194
	} else {
		goto L195
	}
L46:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v626 < int32(0) {
		goto L190
	} else {
		goto L191
	}
L47:
	;
	v611 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v613 = v611 + int64(1900)
	if v611 < int64(8100) {
		v800 = v613
		goto L36
	} else {
		goto L187
	}
L48:
	;
	v602 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v606 = base.I64_rem_s(v602+int64(1900), int64(100))
	v608 = v606 >> (uint(int64(63)) % 64)
	v795 = v132
	v796 = v606 ^ v608 - v608
	goto L37
L49:
	;
	v728 = int32(_a_F___strftime_l_7)
	goto L39
L50:
	;
	v728 = int32(_a_F___strftime_l_8)
	goto L39
L51:
	;
	v803 = int32(1)
	v810 = v132
	v812 = v598
	goto L35
L52:
	;
	v586 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+24)))
	v598 = v586
	goto L51
L53:
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
		goto L177
	} else {
		goto L178
	}
L54:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v524 = int32(7)
	v525 = base.I32_rem_u_s(v521+int32(6), v524)
	v530 = base.I32_div_u_s(v520-v525+v524, v524)
	v795 = v132
	v796 = base.I64_extend_i32_u(v530)
	goto L37
L55:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v515 = int32(7)
	v518 = base.I32_div_u_s(v512-v513+v515, v515)
	v795 = v132
	v796 = base.I64_extend_i32_u(v518)
	goto L37
L56:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v508 != 0 {
		goto L172
	} else {
		goto L173
	}
L57:
	;
	v786 = int32(_a_F___strftime_l_9)
	goto L38
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(1)
	v908 = int32(_a_F___strftime_l_10)
	goto L32
L59:
	;
	v503 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3))))
	v795 = v132
	v796 = v503
	goto L37
L60:
	;
	v325 = int32(0)
	v327 = m.G0
	v329 = v327 - int32(16)
	m.G0 = v329
	v331 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v332 = int32(12)
	v333 = v329 + v332
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui32(v332) <= base.Ui32(v334) {
		goto L128
	} else {
		goto L129
	}
L61:
	;
	v786 = int32(_a_F___strftime_l_11)
	goto L38
L62:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if int32(11) < v320 {
		goto L124
	} else {
		goto L125
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(1)
	v908 = int32(_a_F___strftime_l_12)
	goto L32
L64:
	;
	v314 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+4)))
	v795 = v132
	v796 = v314
	goto L37
L65:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v795 = v132
	v796 = base.I64_extend_i32_s(v310 + int32(1))
	goto L37
L66:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v803 = int32(3)
	v810 = v132
	v812 = base.I64_extend_i32_s(v305 + int32(1))
	goto L35
L67:
	;
	v294 = int32(2)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v295 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L68:
	;
	v293 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+8)))
	v795 = v132
	v796 = v293
	goto L37
L69:
	;
	v166 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v168 = v166 + int64(1900)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	if v169 <= int32(2) {
		goto L83
	} else {
		goto L84
	}
L70:
	;
	v786 = int32(_a_F___strftime_l_13)
	goto L38
L71:
	;
	v164 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+12)))
	v795 = v163
	v796 = v164
	goto L37
L72:
	;
	v163 = int32(95)
	goto L71
L73:
	;
	v157 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v161 = base.I64_div_s(v157+int64(1900), int64(100))
	v795 = v132
	v796 = v161
	goto L37
L74:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui32(int32(11)) < base.Ui32(v152) {
		goto L34
	} else {
		goto L81
	}
L75:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui32(int32(11)) < base.Ui32(v147) {
		goto L34
	} else {
		goto L80
	}
L76:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if base.Ui32(int32(6)) < base.Ui32(v142) {
		goto L34
	} else {
		goto L79
	}
L77:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if base.Ui32(v139) <= base.Ui32(int32(6)) {
		goto L42
	} else {
		goto L78
	}
L78:
	;
	goto L34
L79:
	;
	v669 = v142 + int32(_a_F___strftime_l_14)
	goto L41
L80:
	;
	v669 = v147 + int32(_a_F___strftime_l_15)
	goto L41
L81:
	;
	v669 = v152 + int32(_a_F___strftime_l_16)
	goto L41
L82:
	;
	if v136 == int32(103) {
		goto L43
	} else {
		goto L117
	}
L83:
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
		goto L88
	} else {
		goto L89
	}
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(v169) < base.Ui32(int32(361)) {
		v290 = v168
		goto L82
	} else {
		goto L101
	}
L86:
	;
	if v226 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L87:
	;
	v226 = v224
	goto L86
L88:
	;
	if v196 != 0 {
		v224 = v196
		goto L87
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v218 = base.I32_rem_u_s(v189+int32(371), int32(7))
	switch v218 - int32(3) {
	case 0:
		goto L96
	case 1:
		v224 = v177
		goto L87
	default:
		goto L95
	}
L91:
	;
	v199 = int32(52)
	v203 = base.I32_rem_u_s(v189+int32(6), int32(7))
	switch v203 - int32(4) {
	case 0:
		goto L92
	case 1:
		goto L93
	default:
		v224 = v199
		goto L87
	}
L92:
	;
	v226 = int32(53)
	goto L86
L93:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v208 = base.I32_rem_s(v206, int32(400))
	v211 = F_is_leap(m, v208-int32(1))
	mBase = m.M
	if v211 == int32(0) {
		v224 = v199
		goto L87
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v224 = int32(1)
	goto L87
L96:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v222 = F_is_leap(m, v221)
	mBase = m.M
	if v222 != 0 {
		v224 = v177
		goto L87
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v229 = v168
	goto L100
L99:
	;
	v229 = v166 + int64(1899)
	goto L100
L100:
	;
	v290 = v229
	goto L82
L101:
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
		goto L104
	} else {
		goto L105
	}
L102:
	;
	if v286 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L103:
	;
	v286 = v284
	goto L102
L104:
	;
	if v256 != 0 {
		v284 = v256
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v278 = base.I32_rem_u_s(v249+int32(371), int32(7))
	switch v278 - int32(3) {
	case 0:
		goto L112
	case 1:
		v284 = v237
		goto L103
	default:
		goto L111
	}
L107:
	;
	v259 = int32(52)
	v263 = base.I32_rem_u_s(v249+int32(6), int32(7))
	switch v263 - int32(4) {
	case 0:
		goto L108
	case 1:
		goto L109
	default:
		v284 = v259
		goto L103
	}
L108:
	;
	v286 = int32(53)
	goto L102
L109:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v268 = base.I32_rem_s(v266, int32(400))
	v271 = F_is_leap(m, v268-int32(1))
	mBase = m.M
	if v271 == int32(0) {
		v284 = v259
		goto L103
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v284 = int32(1)
	goto L103
L112:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v282 = F_is_leap(m, v281)
	mBase = m.M
	if v282 != 0 {
		v284 = v237
		goto L103
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v289 = v166 + int64(1901)
	goto L116
L115:
	;
	v289 = v168
	goto L116
L116:
	;
	v290 = v289
	goto L82
L117:
	;
	v800 = v290
	goto L36
L118:
	;
	v803 = v294
	v810 = v132
	v812 = int64(12)
	goto L35
L119:
	;
	goto L120
L120:
	;
	v299 = base.I64_extend_i32_s(v295)
	if int32(12) < v295 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v304 = v299 - int64(12)
	goto L123
L122:
	;
	v304 = v299
	goto L123
L123:
	;
	v803 = v294
	v810 = v132
	v812 = v304
	goto L35
L124:
	;
	v323 = int32(_a_F___strftime_l_17)
	goto L126
L125:
	;
	v323 = int32(_a_F___strftime_l_18)
	goto L126
L126:
	;
	v669 = v323
	goto L41
L127:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v354<<(uint(int32(2))%32))+uint32(_c_F___strftime_l[0])))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	if v474 != 0 {
		goto L166
	} else {
		goto L167
	}
L128:
	;
	v337 = int32(12)
	v338 = base.I32_div_s(v334, v337)
	v341 = v334 - v338*v337
	if v341 < int32(0) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	v354 = v334
	v355 = v331
	goto L130
L130:
	;
	if base.Ui64(v355-int64(2)) <= base.Ui64(int64(136)) {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	v346 = v341 + v337
	goto L133
L132:
	;
	v346 = v341
	goto L133
L133:
	;
	v354 = v346
	v355 = base.I64_extend_i32_s(v338+v341>>(uint(int32(31))%32)) + v331
	goto L130
L134:
	;
	v360 = base.I32_wrap_i64(v355)
	v364 = (v360 - int32(68)) >> (uint(int32(2)) % 32)
	if v360&int32(3) == int32(0) {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	goto L136
L136:
	;
	v390 = v355 - int64(100)
	v391 = int64(400)
	v392 = base.I64_div_s(v390, v391)
	v395 = v390 - v392*v391
	v401 = base.I32_wrap_i64(v395)
	if v395 < int64(0) {
		goto L147
	} else {
		goto L148
	}
L137:
	;
	v466 = base.I64_extend_i32_s(v360*int32(31536000) + v380*int32(_a_F___strftime_l_19) + int32(2087447296))
	goto L127
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = v378
	v380 = v377
	goto L137
L139:
	;
	v370 = v364 - int32(1)
	if v333 == int32(0) {
		v380 = v370
		goto L137
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	if v333 == int32(0) {
		v380 = v364
		goto L137
	} else {
		goto L143
	}
L142:
	;
	v377 = v370
	v378 = int32(1)
	goto L138
L143:
	;
	v377 = v364
	v378 = int32(0)
	goto L138
L144:
	;
	v466 = v390*int64(31536000) + base.I64_extend_i32_s(v443+(v442*int32(24)+(base.I32_wrap_i64(v395>>(uint(int64(63))%64))+base.I32_wrap_i64(v392))*int32(97))-v441)*int64(86400) + int64(946771200)
	goto L127
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = v437
	v441 = v437
	v442 = v438
	v443 = v439
	goto L144
L146:
	;
	v430 = int32(base.Ui32(v423) >> (uint(int32(2)) % 32))
	v433 = int32(0)
	v434 = base.B2i32(v423&int32(3) == v433)
	if v333 == v433 {
		v441 = v434
		v442 = v422
		v443 = v430
		goto L144
	} else {
		goto L165
	}
L147:
	;
	v406 = v401 + int32(400)
	goto L149
L148:
	;
	v406 = v401
	goto L149
L149:
	;
	if v406 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if int32(200) <= v406 {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	v427 = v325
	v428 = int32(1)
	goto L152
L152:
	;
	if v333 != 0 {
		v437 = v428
		v438 = v427
		v439 = v325
		goto L145
	} else {
		goto L164
	}
L153:
	;
	if v423 != 0 {
		goto L146
	} else {
		goto L163
	}
L154:
	;
	if base.Ui32(int32(300)) <= base.Ui32(v406) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L156
L156:
	;
	v420 = base.B2i32(int32(99) < v406)
	if int32(99) < v406 {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	v422 = int32(3)
	v423 = v406 - int32(300)
	goto L153
L158:
	;
	goto L159
L159:
	;
	v422 = int32(2)
	v423 = v406 - int32(200)
	goto L153
L160:
	;
	v421 = v406 - int32(100)
	goto L162
L161:
	;
	v421 = v406
	goto L162
L162:
	;
	v422 = v420
	v423 = v421
	goto L153
L163:
	;
	v427 = v422
	v428 = int32(0)
	goto L152
L164:
	;
	v441 = v428
	v442 = v427
	v443 = v325
	goto L144
L165:
	;
	v437 = v434
	v438 = v422
	v439 = v430
	goto L145
L166:
	;
	v475 = v471 + int32(_a_F___strftime_l_19)
	goto L168
L167:
	;
	v475 = v471
	goto L168
L168:
	;
	if int32(1) < v354 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v478 = v475
	goto L171
L170:
	;
	v478 = v471
	goto L171
L171:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v480 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+8)))
	v481 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+4)))
	v482 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3))))
	m.G0 = v329 + int32(16)
	v501 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+36)))
	v598 = v482 + (v466 + base.I64_extend_i32_s(v478) + base.I64_extend_i32_s(v479-int32(1))*int64(86400) + v480*int64(3600) + v481*int64(60)) - v501
	goto L51
L172:
	;
	v510 = v508
	goto L174
L173:
	;
	v510 = int32(7)
	goto L174
L174:
	;
	v598 = base.I64_extend_i32_s(v510)
	goto L51
L175:
	;
	v795 = v132
	v796 = base.I64_extend_i32_u(v584)
	goto L37
L176:
	;
	v584 = v582
	goto L175
L177:
	;
	if v554 != 0 {
		v582 = v554
		goto L176
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v576 = base.I32_rem_u_s(v547+int32(371), int32(7))
	switch v576 - int32(3) {
	case 0:
		goto L185
	case 1:
		v582 = v535
		goto L176
	default:
		goto L184
	}
L180:
	;
	v557 = int32(52)
	v561 = base.I32_rem_u_s(v547+int32(6), int32(7))
	switch v561 - int32(4) {
	case 0:
		goto L181
	case 1:
		goto L182
	default:
		v582 = v557
		goto L176
	}
L181:
	;
	v584 = int32(53)
	goto L175
L182:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v566 = base.I32_rem_s(v564, int32(400))
	v569 = F_is_leap(m, v566-int32(1))
	mBase = m.M
	if v569 == int32(0) {
		v582 = v557
		goto L176
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v582 = int32(1)
	goto L176
L185:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v580 = F_is_leap(m, v579)
	mBase = m.M
	if v580 != 0 {
		v582 = v535
		goto L176
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v129)+48)) = v613
	v621 = F_snprintf(m, v125, int32(100), int32(_a_F___strftime_l_20), v129+int32(48))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	return int32(0)
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v621
	v908 = v125
	goto L32
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(0)
	v908 = int32(_a_F___strftime_l_21)
	goto L32
L191:
	;
	goto L192
L192:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	v633 = int32(3600)
	v634 = base.I32_div_s(v632, v633)
	v635 = int32(100)
	v642 = base.I32_div_s(base.I32_extend16_s(v632-v634*v633), int32(60))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+64)) = v634*v635 + base.I32_extend16_s(v642)
	v650 = F_snprintf(m, v125, v635, int32(_a_F___strftime_l_22), v129-int32(-64))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L188
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v650
	v908 = v125
	goto L32
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(0)
	v908 = int32(_a_F___strftime_l_21)
	goto L32
L195:
	;
	goto L196
L196:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	F_do_tzset(m)
	mBase = m.M
	v849 = v659
	goto L33
L197:
	;
	v849 = v726
	goto L33
L198:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v676 != 0 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	goto L200
L200:
	;
	v679 = v669 >> (uint(int32(16)) % 32)
	v680 = int32(_a_F___strftime_l_23)
	v681 = v669 & v680
	if v681 != v680 {
		goto L204
	} else {
		goto L205
	}
L201:
	;
	v677 = int32(_a_F___strftime_l_24)
	goto L203
L202:
	;
	v677 = int32(_a_F___strftime_l_25)
	goto L203
L203:
	;
	v726 = v677
	goto L197
L204:
	;
	v694 = int32(_a_F___strftime_l_21)
	switch v679 - int32(1) {
	case 0:
		goto L214
	case 1:
		goto L213
	default:
		v718 = v694
		goto L210
	case 4:
		goto L212
	}
L205:
	;
	if int32(5) < v679 {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l4+v679<<(uint(int32(2))%32))))
	if v689 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v693 = v689 + int32(8)
	goto L209
L208:
	;
	v693 = int32(_a_F___strftime_l_26)
	goto L209
L209:
	;
	v726 = v693
	goto L197
L210:
	;
	v726 = v718
	goto L197
L211:
	;
	if v681 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L212:
	;
	if base.Ui32(int32(3)) < base.Ui32(v681) {
		v718 = v694
		goto L210
	} else {
		goto L217
	}
L213:
	;
	if base.Ui32(int32(49)) < base.Ui32(v681) {
		v718 = v694
		goto L210
	} else {
		goto L216
	}
L214:
	;
	if base.Ui32(int32(1)) < base.Ui32(v681) {
		v718 = v694
		goto L210
	} else {
		goto L215
	}
L215:
	;
	v706 = int32(_a_F___strftime_l_27)
	goto L211
L216:
	;
	v706 = int32(_a_F___strftime_l_28)
	goto L211
L217:
	;
	v706 = int32(_a_F___strftime_l_29)
	goto L211
L218:
	;
	v726 = v706
	goto L197
L219:
	;
	goto L220
L220:
	;
	v709 = v706
	v711 = v681
	goto L221
L221:
	;
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v709))))
	v715 = v709 + int32(1)
	if v713 != 0 {
		v709 = v715
		goto L221
	} else {
		goto L223
	}
L222:
	;
	v718 = v715
	goto L210
L223:
	;
	v717 = v711 - int32(1)
	if v717 != 0 {
		v709 = v715
		v711 = v717
		goto L221
	} else {
		goto L224
	}
L224:
	;
	goto L222
L225:
	;
	v786 = v785
	goto L38
L226:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v735 != 0 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	goto L228
L228:
	;
	v738 = v728 >> (uint(int32(16)) % 32)
	v739 = int32(_a_F___strftime_l_23)
	v740 = v728 & v739
	if v740 != v739 {
		goto L232
	} else {
		goto L233
	}
L229:
	;
	v736 = int32(_a_F___strftime_l_24)
	goto L231
L230:
	;
	v736 = int32(_a_F___strftime_l_25)
	goto L231
L231:
	;
	v785 = v736
	goto L225
L232:
	;
	v753 = int32(_a_F___strftime_l_21)
	switch v738 - int32(1) {
	case 0:
		goto L242
	case 1:
		goto L241
	default:
		v777 = v753
		goto L238
	case 4:
		goto L240
	}
L233:
	;
	if int32(5) < v738 {
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l4+v738<<(uint(int32(2))%32))))
	if v748 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v752 = v748 + int32(8)
	goto L237
L236:
	;
	v752 = int32(_a_F___strftime_l_26)
	goto L237
L237:
	;
	v785 = v752
	goto L225
L238:
	;
	v785 = v777
	goto L225
L239:
	;
	if v740 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L240:
	;
	if base.Ui32(int32(3)) < base.Ui32(v740) {
		v777 = v753
		goto L238
	} else {
		goto L245
	}
L241:
	;
	if base.Ui32(int32(49)) < base.Ui32(v740) {
		v777 = v753
		goto L238
	} else {
		goto L244
	}
L242:
	;
	if base.Ui32(int32(1)) < base.Ui32(v740) {
		v777 = v753
		goto L238
	} else {
		goto L243
	}
L243:
	;
	v765 = int32(_a_F___strftime_l_27)
	goto L239
L244:
	;
	v765 = int32(_a_F___strftime_l_28)
	goto L239
L245:
	;
	v765 = int32(_a_F___strftime_l_29)
	goto L239
L246:
	;
	v785 = v765
	goto L225
L247:
	;
	goto L248
L248:
	;
	v768 = v765
	v770 = v740
	goto L249
L249:
	;
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768))))
	v774 = v768 + int32(1)
	if v772 != 0 {
		v768 = v774
		goto L249
	} else {
		goto L251
	}
L250:
	;
	v777 = v774
	goto L238
L251:
	;
	v776 = v770 - int32(1)
	if v776 != 0 {
		v768 = v774
		v770 = v776
		goto L249
	} else {
		goto L252
	}
L252:
	;
	goto L250
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v789
	if v789 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v793 = v125
	goto L256
L255:
	;
	v793 = int32(0)
	goto L256
L256:
	;
	v908 = v793
	goto L32
L257:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v129)+8)) = v812
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v803
	v842 = F_snprintf(m, v125, int32(100), int32(_a_F___strftime_l_30), v129)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L188
	} else {
		goto L267
	}
L258:
	;
	v816 = v78
	goto L260
L259:
	;
	v816 = v810
	goto L260
L260:
	;
	if v816 != int32(95) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	if v816 != int32(45) {
		goto L257
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v129)+40)) = v812
	*(*int32)(unsafe.Add(mBase, uint32(v129)+32)) = v803
	v835 = F_snprintf(m, v125, int32(100), int32(_a_F___strftime_l_31), v129+int32(32))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L188
	} else {
		goto L266
	}
L264:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v129)+16)) = v812
	v826 = F_snprintf(m, v125, int32(100), int32(_a_F___strftime_l_32), v129+int32(16))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L188
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v826
	v908 = v125
	goto L32
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v835
	v908 = v125
	goto L32
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v842
	v908 = v125
	goto L32
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v906
	v908 = v849
	goto L32
L269:
	;
	v906 = v898 - v849
	goto L268
L270:
	;
	v877 = v873
	goto L279
L271:
	;
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849))))
	if v857 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v906 = int32(0)
	goto L268
L273:
	;
	goto L274
L274:
	;
	v862 = v849
	goto L275
L275:
	;
	v866 = v862 + int32(1)
	if v866&int32(3) == int32(0) {
		v873 = v866
		goto L270
	} else {
		goto L277
	}
L276:
	;
	v898 = v866
	goto L269
L277:
	;
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866))))
	if v871 != 0 {
		v862 = v866
		goto L275
	} else {
		goto L278
	}
L278:
	;
	goto L276
L279:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v877)))
	v886 = int32(-2139062144)
	if (int32(16843008)-v883|v883)&v886 == v886 {
		v877 = v877 + int32(4)
		goto L279
	} else {
		goto L281
	}
L280:
	;
	v892 = v877
	goto L282
L281:
	;
	goto L280
L282:
	;
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v892))))
	if v896 != 0 {
		v892 = v892 + int32(1)
		goto L282
	} else {
		goto L284
	}
L283:
	;
	v898 = v892
	goto L269
L284:
	;
	goto L283
L285:
	;
	if v114 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	v1136 = l1 - v1119
	if base.Ui32(v1118) < base.Ui32(v1136) {
		goto L320
	} else {
		goto L321
	}
L287:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	v1118 = v928
	v1119 = v40
	v1121 = v908
	goto L286
L288:
	;
	goto L289
L289:
	;
	v929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908))))
	switch v929 - int32(43) {
	case 0, 2:
		goto L291
	default:
		goto L292
	}
L290:
	;
	if v939&int32(255) != int32(48) {
		v991 = v941
		v994 = v940
		goto L293
	} else {
		goto L294
	}
L291:
	;
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908)+1)))
	v934 = int32(1)
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	v939 = v933
	v940 = v908 + v934
	v941 = v936 - v934
	goto L290
L292:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	v939 = v929
	v940 = v908
	v941 = v932
	goto L290
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+124)) = v991
	v1018 = int32(0)
	goto L299
L294:
	;
	v954 = v941
	v957 = v940
	goto L295
L295:
	;
	v972 = int32(*(*int8)(unsafe.Add(mBase, uint32(v957)+1)))
	if base.Ui32(int32(9)) < base.Ui32(v972-int32(48)) {
		v991 = v954
		v994 = v957
		goto L293
	} else {
		goto L297
	}
L296:
	;
	v991 = v980
	v994 = v978
	goto L293
L297:
	;
	v977 = int32(1)
	v978 = v957 + v977
	v980 = v954 - v977
	if v972 == int32(48) {
		v954 = v980
		v957 = v978
		goto L295
	} else {
		goto L298
	}
L298:
	;
	goto L296
L299:
	;
	v1040 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1018+v994))))
	if base.Ui32(v1040-int32(48)) < base.Ui32(int32(10)) {
		v1018 = v1018 + int32(1)
		goto L299
	} else {
		goto L301
	}
L300:
	;
	if base.Ui32(v991) < base.Ui32(v114) {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	goto L300
L302:
	;
	v1046 = v114
	goto L304
L303:
	;
	v1046 = v991
	goto L304
L304:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v1048 < int32(-1900) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	if base.Ui32(v1071) <= base.Ui32(v991) {
		v1118 = v991
		v1119 = v1072
		v1121 = v994
		goto L286
	} else {
		goto L314
	}
L306:
	;
	v1065 = int32(45)
	goto L308
L307:
	;
	if v81 != int32(43) {
		v1071 = v1046
		v1072 = v40
		goto L305
	} else {
		goto L309
	}
L308:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v40))) = uint8(v1065)
	v1067 = int32(1)
	v1071 = v1046 - v1067
	v1072 = v40 + v1067
	goto L305
L309:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v1059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1058))))
	if v1059 == int32(67) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1062 = int32(3)
	goto L312
L311:
	;
	v1062 = int32(5)
	goto L312
L312:
	;
	if base.Ui32(v1046-v991+v1018) < base.Ui32(v1062) {
		v1071 = v1046
		v1072 = v40
		goto L305
	} else {
		goto L313
	}
L313:
	;
	v1065 = int32(43)
	goto L308
L314:
	;
	if base.Ui32(l1) <= base.Ui32(v1072) {
		v1118 = v991
		v1119 = v1072
		v1121 = v994
		goto L286
	} else {
		goto L315
	}
L315:
	;
	v1082 = v1071
	v1084 = v1072
	goto L316
L316:
	;
	v1102 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v1084))) = uint8(v1102)
	v1104 = int32(1)
	v1105 = v1084 + v1104
	v1107 = v1082 - v1104
	if base.Ui32(v1107) <= base.Ui32(v991) {
		v1118 = v991
		v1119 = v1105
		v1121 = v994
		goto L286
	} else {
		goto L318
	}
L317:
	;
	v1118 = v991
	v1119 = v1105
	v1121 = v994
	goto L286
L318:
	;
	if base.Ui32(v1105) < base.Ui32(l1) {
		v1082 = v1107
		v1084 = v1105
		goto L316
	} else {
		goto L319
	}
L319:
	;
	goto L317
L320:
	;
	v1138 = v1118
	goto L322
L321:
	;
	v1138 = v1136
	goto L322
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+124)) = v1138
	if v1138 != 0 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	v1147 = v123
	v1171 = v1143 + v1119
	goto L8
L324:
	;
	v1141 = F__emscripten_memcpy_bulkmem(m, l0+v1119, v1121, v1138)
	mBase = m.M
	goto L326
L325:
	;
	goto L326
L326:
	;
	goto L323
L327:
	;
	v1184 = v1171
	goto L7
L328:
	;
	v1204 = l1 - int32(1)
	goto L330
L329:
	;
	v1204 = v1184
	goto L330
L330:
	;
	v1215 = v1204
	v1232 = int32(0)
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
func F_s_lock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v74 float64
	_ = v74
	var v77 float64
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v134 int32
	_ = v134
	v5 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_s_lock[0])) = v134
	goto L1
L3:
	;
	if int32(999) < v112 {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	F_s_lock_stuck(m, l1, l2, l3)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L17
	} else {
		goto L30
	}
L5:
	;
	v19 = int32(0)
	v21 = v5
	v22 = v5
	goto L8
L6:
	;
	goto L7
L7:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[0]))
	v112 = v105
	goto L3
L8:
	;
	v25 = v19 + int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[0]))
	if v27 <= v25 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[0]))
	if v90 == int32(0) {
		v112 = v97
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v30 = v22 + int32(1)
	if int32(1001) <= v30 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	v88 = v25
	v90 = v21
	v91 = v22
	goto L12
L12:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v93 != 0 {
		v19 = v88
		v21 = v90
		v22 = v91
		goto L8
	} else {
		goto L27
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(150994950)
	if v21 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v39 = v21
	goto L16
L15:
	;
	v39 = int32(1000)
	goto L16
L16:
	;
	F_pg_usleep(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_s_lock[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(0)
	v48 = int32(_a_F_s_lock_0)
	v51 = *(*int64)(unsafe.Add(mBase, _c_F_s_lock[2]))
	v52 = *(*int64)(unsafe.Add(mBase, _c_F_s_lock[3]))
	v53 = v51 ^ v52
	*(*int64)(unsafe.Add(mBase, _c_F_s_lock[3])) = base.I64_rotl(v53, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_s_lock[2])) = v53<<(uint(int64(16))%64) ^ base.I64_rotl(v51, int64(24)) ^ v53
	v74 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v51*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L20
L19:
	;
	v84 = v83 + v39
	if int32(_a_F_s_lock_1) < v84 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v77 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v39), v74), float64(0.5))
	if base.F64_lt(base.F64_abs(v77), float64(2.147483648e+09)) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v81 = base.I32_trunc_f64_s(v77)
	v83 = v81
	goto L19
L22:
	;
	goto L23
L23:
	;
	v83 = int32(-2147483648)
	goto L19
L24:
	;
	v87 = int32(1000)
	goto L26
L25:
	;
	v87 = v84
	goto L26
L26:
	;
	v88 = int32(0)
	v90 = v87
	v91 = v30
	goto L12
L27:
	;
	goto L9
L28:
	;
	if v97 < int32(11) {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v134 = v97 - int32(1)
	goto L2
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v119 = int32(900)
	if v119 <= v112 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v122 = v119
	goto L34
L33:
	;
	v122 = v112
	goto L34
L34:
	;
	v134 = v122 + int32(100)
	goto L2
}
func F_satisfies_hash_partition(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v252 int32
	_ = v252
	var v254 int64
	_ = v254
	var v260 int32
	_ = v260
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v328 int64
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int64
	_ = v345
	var v355 int64
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v420 int64
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int64
	_ = v434
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int64
	_ = v446
	var v448 int32
	_ = v448
	var v467 int64
	_ = v467
	var v470 int64
	_ = v470
	var v474 int32
	_ = v474
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	v2 = int32(0)
	v18 = int64(0)
	v19 = m.G0
	v21 = v19 - int32(96)
	m.G0 = v21
	v24 = F_Int64GetDatum(m, int64(8816678312871386365))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v28 != 0 {
		v474 = v2
		goto L10
	} else {
		goto L11
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L123
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L117
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L113
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L108
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L104
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L100
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L96
	}
L10:
	;
	m.G0 = v21 + int32(96)
	return v474
L11:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v29 != 0 {
		v474 = v2
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v30 != 0 {
		v474 = v2
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v31 <= int32(0) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v34 < int32(0) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if base.Ui32(v31) <= base.Ui32(v34) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v40 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v281)+8))
	if v298 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v41 == v38 {
		v281 = v40
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v44 = F_relation_open(m, v38, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v46 = F_RelationGetPartitionKey(m, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v46 == int32(0) {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v50 != int32(104) {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v54 = int32(0)
	if v53 == v54 {
		v65 = v54
		goto L28
	} else {
		goto L29
	}
L26:
	;
	F_relation_close(m, v44, int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L63
	}
L27:
	;
	if v65&int32(1) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	goto L27
L29:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	if v57 == int32(0) {
		v65 = v54
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v60 != int32(15) {
		v65 = v54
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+13)))
	v65 = v63
	goto L28
L32:
	;
	v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	v72 = v70 - int32(3)
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+4)))
	if v72 != v73 {
		goto L5
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v163 = F_pg_detoast_datum(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L52
	}
L35:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	v81 = F_MemoryContextAllocZero(m, v76, v72*int32(28)+int32(144))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v81
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v38
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v88
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+4)))
	v95 = v93 << (uint(int32(2)) % 32)
	if v95 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+4)))
	if v98 <= int32(0) {
		v260 = v86
		goto L26
	} else {
		goto L41
	}
L38:
	;
	v96 = F__emscripten_memcpy_bulkmem(m, v86+int32(16), v92, v95)
	mBase = m.M
	goto L40
L39:
	;
	goto L40
L40:
	;
	goto L37
L41:
	;
	v106 = int32(0)
	goto L42
L42:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v125 = F_get_fn_expr_argtype(m, v122, v106+int32(3))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	v260 = v86
	goto L26
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v46)+32))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v106<<(uint(int32(2))%32))))
	if v125 != v131 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v133 = F_IsBinaryCoercible(m, v125, v131)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v138 = v106 * int32(28)
	v139 = v86 + int32(144) + v138
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	v141 = v140 + v138
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+20))
	v146 = v139 + int32(16)
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v141)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v146))) = v147
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
	*(*int64)(unsafe.Add(mBase, uint32(v139))) = v149
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v141)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v151
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v141)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v139)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = int32(0)
	goto L50
L48:
	;
	if v133 == int32(0) {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v159 = v106 + int32(1)
	v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+4)))
	if v159 < v160 {
		v106 = v159
		goto L42
	} else {
		goto L51
	}
L51:
	;
	goto L43
L52:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v168 = F_MemoryContextAllocZero(m, v166, int32(172))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v170)+16)) = v168
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v38
	v175 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = v177
	F_get_typlenbyvalalign(m, v177, v173+int32(12), v173+int32(14), v173+int32(15))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+16)) = v188
	v190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+4)))
	if int32(0) < v190 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v46)+32))
	v198 = int32(0)
	goto L58
L56:
	;
	goto L57
L57:
	;
	v241 = v173 + int32(144)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+20))
	v247 = v173 + int32(160)
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v242)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v247))) = v248
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v242)))
	*(*int64)(unsafe.Add(mBase, uint32(v241))) = v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v241)+24)) = v252
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v242)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = int32(0)
	goto L62
L58:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v194+v198<<(uint(int32(2))%32))))
	if v217 != v193 {
		goto L3
	} else {
		goto L60
	}
L59:
	;
	goto L57
L60:
	;
	v220 = v198 + int32(1)
	if v220 != v190 {
		v198 = v220
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v260 = v173
	goto L26
L63:
	;
	v281 = v260
	goto L17
L64:
	;
	v470 = base.I64_rem_u_s(v467, base.I64_extend_i32_u(v31))
	v474 = base.B2i32(base.I64_extend_i32_u(v34) == v470)
	goto L10
L65:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	if v301 <= int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v360 = F_pg_detoast_datum(m, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L78
	}
L68:
	;
	v467 = v18
	goto L64
L69:
	;
	goto L70
L70:
	;
	v311 = int32(0)
	v328 = v18
	goto L71
L71:
	;
	v331 = v311<<(uint(int32(3))%32) + (l0 + int32(20))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+28)))
	if v332 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v467 = v355
	goto L64
L73:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v281+int32(16)+v311<<(uint(int32(2))%32))))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v331)+24))
	v343 = F_FunctionCall2Coll(m, v281+int32(144)+v311*int32(28), v341, v342, v24)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	v355 = v328
	goto L75
L75:
	;
	v357 = v311 + int32(1)
	if v357 != v301 {
		v311 = v357
		v328 = v355
		goto L71
	} else {
		goto L77
	}
L76:
	;
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v343)))
	v355 = v345 + (v328<<(uint(int64(54))%64) + int64(base.Ui64(v328)>>(uint(int64(7))%64))) + int64(5305509591434766563) ^ v328
	goto L75
L77:
	;
	goto L72
L78:
	;
	v363 = int32(*(*int16)(unsafe.Add(mBase, uint32(v281)+12)))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+14)))
	v365 = int32(*(*int8)(unsafe.Add(mBase, uint32(v281)+15)))
	F_deconstruct_array(m, v360, v363, v364, v365, v21+int32(88), v21+int32(84), v21+int32(92))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v21)+92))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	if v374 == v375 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v403 = int32(0)
	v405 = v374
	v420 = v18
	goto L89
L81:
	;
	if int32(0) < v374 {
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
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	v467 = v18
	goto L64
L85:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v21)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v388
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_0), v21+int32(16))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_2), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
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
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421+v403))))
	if v423 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v467 = v446
	goto L64
L91:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v281)+16))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v21)+88))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v427+v403<<(uint(int32(2))%32))))
	v432 = F_FunctionCall2Coll(m, v281+int32(144), v426, v431, v24)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	v445 = v405
	v446 = v420
	goto L93
L93:
	;
	v448 = v403 + int32(1)
	if v448 < v445 {
		v403 = v448
		v405 = v445
		v420 = v446
		goto L89
	} else {
		goto L95
	}
L94:
	;
	v434 = *(*int64)(unsafe.Add(mBase, uint32(v432)))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v21)+92))
	v445 = v444
	v446 = v434 + (v420<<(uint(int64(54))%64) + int64(base.Ui64(v420)>>(uint(int64(7))%64))) + int64(5305509591434766563) ^ v420
	goto L93
L95:
	;
	goto L90
L96:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_4), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_5), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
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
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_6), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_7), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
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
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_8), int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_9), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
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
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v549 = F_get_rel_name(m, v38)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v549
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_10), v21)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_11), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
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
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v567 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v567
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_0), v21-int32(-64))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_12), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
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
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v46)+32))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v587+v106<<(uint(int32(2))%32))))
	v592 = F_format_type_be(m, v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v594 = F_format_type_be(m, v125)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v594
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v592
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v106 + int32(1)
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_13), v21+int32(48))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_14), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v46)+32))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v618+v198<<(uint(int32(2))%32))))
	v623 = F_format_type_be(m, v622)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	v626 = F_format_type_be(m, v625)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v626
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v623
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v198 + int32(1)
	F_errmsg(m, int32(_a_F_satisfies_hash_partition_15), v21+int32(32))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_satisfies_hash_partition_1), int32(_a_F_satisfies_hash_partition_16), int32(_a_F_satisfies_hash_partition_3))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
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
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 float64
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 float64
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 float64
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
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
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int64
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 float64
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int64
	_ = v239
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 float64
	_ = v257
	var v258 int32
	_ = v258
	var v261 float64
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v277 float64
	_ = v277
	var v278 int32
	_ = v278
	var v279 float64
	_ = v279
	var v283 float32
	_ = v283
	var v289 float64
	_ = v289
	var v294 float64
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 float64
	_ = v302
	var v310 float64
	_ = v310
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v387 float64
	_ = v387
	var v388 int32
	_ = v388
	var v399 float64
	_ = v399
	var v401 float64
	_ = v401
	var v404 int32
	_ = v404
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
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
	var v449 int32
	_ = v449
	var v450 float64
	_ = v450
	var v457 float64
	_ = v457
	var v458 float64
	_ = v458
	var v464 float64
	_ = v464
	var v465 float64
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v476 float64
	_ = v476
	var v478 float64
	_ = v478
	var v499 float64
	_ = v499
	var v502 float64
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 float64
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v529 float64
	_ = v529
	var v531 float64
	_ = v531
	var v534 int32
	_ = v534
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 float64
	_ = v571
	var v578 float64
	_ = v578
	var v579 float64
	_ = v579
	var v585 float64
	_ = v585
	var v586 float64
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 float64
	_ = v623
	var v641 float64
	_ = v641
	var v644 float64
	_ = v644
	var v647 float64
	_ = v647
	var v650 float64
	_ = v650
	var v653 float64
	_ = v653
	var v656 float64
	_ = v656
	var v659 float64
	_ = v659
	var v662 float64
	_ = v662
	var v665 float64
	_ = v665
	var v675 float64
	_ = v675
	var v677 float64
	_ = v677
	var v696 float64
	_ = v696
	var v699 float64
	_ = v699
	var v706 float64
	_ = v706
	var v723 float64
	_ = v723
	var v738 float64
	_ = v738
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
	return v738
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
		v738 = float64(0.5)
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
	if l2 != 0 {
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
	if l2 != 0 {
		goto L89
	} else {
		goto L90
	}
L26:
	;
	if v127|v126 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v133 = m.G0
	v135 = v133 - int32(112)
	m.G0 = v135
	F_examine_variable(m, l0, v100, l3, v135+int32(80))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135)+84))
	if v141 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	m.G0 = v135 + int32(112)
	if base.F64_ge(v310, float64(0)) != 0 {
		v738 = v310
		goto L3
	} else {
		goto L87
	}
L30:
	;
	v144 = float64(-1)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v135)+88))
	if v145 == int32(0) {
		v310 = v144
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v151 != int32(7) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v135)+92))
	m.T0[v148].(func(*base.Module, int32))(m, v145)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v310 = v144
	goto L29
L35:
	;
	v154 = float64(-1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v135)+88))
	if v155 == int32(0) {
		v310 = v154
		goto L29
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)))
	if v161 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v135)+92))
	m.T0[v158].(func(*base.Module, int32))(m, v155)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v310 = v154
	goto L29
L40:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v135)+88))
	if v164 == int32(0) {
		v310 = v7
		goto L29
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+76)) = v170
	v173 = F_lookup_type_cache(m, v43, int32(64))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v135)+92))
	m.T0[v167].(func(*base.Module, int32))(m, v164)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v310 = v7
	goto L29
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v173)+108))
	if v175 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v178 = float64(-1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v135)+88))
	if v179 == int32(0) {
		v310 = v178
		goto L29
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v185 = v28&int32(1) ^ v126
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v135)+88))
	if v186 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v135)+92))
	m.T0[v182].(func(*base.Module, int32))(m, v179)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v310 = v178
	goto L29
L51:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v135)+88))
	if v295 != 0 {
		goto L78
	} else {
		goto L79
	}
L52:
	;
	if v185 != 0 {
		goto L75
	} else {
		goto L76
	}
L53:
	;
	v191 = F_statistic_proc_security_check(m, v135+int32(80), v175)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v191 == int32(0) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v135)+88))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+16))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+22)))
	v204 = F_get_attstatsslot(m, v135+int32(40), v195, int32(4), int32(0), int32(3))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	v283 = *(*float32)(unsafe.Add(mBase, uint32(v196+v197)+8))
	v294 = base.F64_mul(v279, base.F64_sub(float64(1), base.F64_promote_f32(v283)))
	goto L51
L57:
	;
	if v204 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v185 != 0 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	goto L60
L60:
	;
	if v185 != 0 {
		v279 = float64(0.005)
		goto L56
	} else {
		goto L73
	}
L61:
	;
	F_free_attstatsslot(m, v135)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L71
	}
L62:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v135)+52))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v135)+56))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v135)+60))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v135)+64))
	v257 = F_mcelem_array_contained_selec(m, v250, v251, v252, v253, v135+int32(76), int32(1), v249, v248, v173)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L70
	}
L63:
	;
	v237 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v135)+32)) = v237
	v239 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v135)+24)) = v239
	*(*int64)(unsafe.Add(mBase, uint32(v135)+16)) = v239
	*(*int64)(unsafe.Add(mBase, uint32(v135)+8)) = v239
	*(*int64)(unsafe.Add(mBase, uint32(v135))) = v239
	v248 = v206
	v249 = v237
	goto L62
L64:
	;
	v206 = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v135)+88))
	v211 = F_get_attstatsslot(m, v135, v207, int32(5), v206, int32(2))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+32)) = int32(0)
	v219 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v135)+24)) = v219
	*(*int64)(unsafe.Add(mBase, uint32(v135)+16)) = v219
	*(*int64)(unsafe.Add(mBase, uint32(v135)+8)) = v219
	*(*int64)(unsafe.Add(mBase, uint32(v135))) = v219
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v135)+52))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v135)+56))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v135)+60))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v135)+64))
	v235 = F_mcelem_array_contain_overlap_selec(m, v227, v228, v229, v230, v135+int32(76), int32(1), int32(2751), v173)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	if v211 == int32(0) {
		goto L63
	} else {
		goto L68
	}
L68:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
	v248 = v215
	v249 = v216
	goto L62
L69:
	;
	v261 = v235
	goto L61
L70:
	;
	v261 = v257
	goto L61
L71:
	;
	F_free_attstatsslot(m, v135+int32(40))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v279 = v261
	goto L56
L73:
	;
	v269 = int32(0)
	v277 = F_mcelem_array_contain_overlap_selec(m, v269, v269, v269, v269, v135+int32(76), int32(1), int32(2751), v173)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v279 = v277
	goto L56
L75:
	;
	v289 = float64(0.005)
	goto L77
L76:
	;
	v289 = float64(0.004999999888241291)
	goto L77
L77:
	;
	v294 = v289
	goto L51
L78:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v135)+92))
	m.T0[v296].(func(*base.Module, int32))(m, v295)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
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
	v302 = v294
	goto L84
L83:
	;
	v302 = base.F64_sub(float64(1), v294)
	goto L84
L84:
	;
	if base.F64_lt(v302, float64(0)) != 0 {
		v310 = float64(0)
		goto L29
	} else {
		goto L85
	}
L85:
	;
	if base.F64_gt(v302, float64(1)) == int32(0) {
		v310 = v302
		goto L29
	} else {
		goto L86
	}
L86:
	;
	v310 = float64(1)
	goto L29
L87:
	;
	goto L25
L88:
	;
	if v333 == int32(0) {
		v738 = float64(0.5)
		goto L3
	} else {
		goto L94
	}
L89:
	;
	v329 = F_get_oprjoin(m, v29)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v331 = F_get_oprrest(m, v29)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	v333 = v329
	goto L88
L93:
	;
	v333 = v331
	goto L88
L94:
	;
	F_fmgr_info(m, v333, v26+int32(68))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	switch v333 - int32(101) {
	case 0, 4:
		v346 = int32(1)
		v347 = v127
		goto L96
	case 1, 5:
		goto L98
	default:
		v345 = v127
		goto L97
	}
L96:
	;
	if v100 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L97:
	;
	v346 = v126
	v347 = v345
	goto L96
L98:
	;
	v345 = int32(1)
	goto L97
L99:
	;
	v723 = float64(0)
	if base.F64_lt(v706, v723) != 0 {
		v738 = v723
		goto L3
	} else {
		goto L202
	}
L100:
	;
	if v514 != 0 {
		goto L191
	} else {
		goto L192
	}
L101:
	;
	v593 = F_palloc0(m, int32(16))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L179
	}
L102:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v350 != int32(35) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if v350 != int32(7) {
		goto L101
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+20)))
	if v503 != 0 {
		goto L101
	} else {
		goto L151
	}
L106:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+24)))
	if v355 == int32(1) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v738 = float64(0)
	goto L3
L108:
	;
	goto L109
L109:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	v360 = F_pg_detoast_datum(m, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v360)+12))
	F_get_typlenbyvalalign(m, v362, v26+int32(66), v26+int32(65), v26-int32(-64))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v372 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+66)))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+65)))
	v374 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26)+64)))
	F_deconstruct_array(m, v360, v372, v373, v374, v26+int32(56), v26+int32(52), v26+int32(60))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	if v28&int32(1) != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v387 = float64(0)
	goto L115
L114:
	;
	v387 = float64(1)
	goto L115
L115:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v26)+60))
	if v388 <= int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v28&int32(1) != 0 {
		goto L140
	} else {
		goto L141
	}
L117:
	;
	v476 = v387
	v478 = v387
	goto L116
L118:
	;
	goto L119
L119:
	;
	v399 = v387
	v401 = v387
	v404 = int32(0)
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v35
	v418 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+66)))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v419+v404<<(uint(int32(2))%32))))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424+v404))))
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+65)))
	v428 = F_makeConst(m, v43, int32(-1), v47, v418, v423, v426, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	v476 = v464
	v478 = v465
	goto L116
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v428
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v432
	v438 = F_list_make2_impl(m, v26+int32(16), v26+int32(12))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if l2 != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v450 = *(*float64)(unsafe.Add(mBase, uint32(v449)))
	if v28&int32(1) != 0 {
		goto L131
	} else {
		goto L132
	}
L125:
	;
	v443 = F_FunctionCall5Coll(m, v26+int32(68), v440, l0, v29, v438, base.I32_extend16_s(l4), l5)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v447 = F_FunctionCall4Coll(m, v26+int32(68), v440, l0, v29, v438, l3)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L129
	}
L128:
	;
	v449 = v443
	goto L124
L129:
	;
	v449 = v447
	goto L124
L130:
	;
	v467 = v404 + int32(1)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v26)+60))
	if v467 < v468 {
		v399 = v464
		v401 = v465
		v404 = v467
		goto L120
	} else {
		goto L138
	}
L131:
	;
	if v346 != 0 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	v458 = base.F64_mul(v399, v450)
	if v347 == int32(0) {
		v464 = v458
		v465 = v401
		goto L130
	} else {
		goto L137
	}
L134:
	;
	v457 = base.F64_add(v401, v450)
	goto L136
L135:
	;
	v457 = v401
	goto L136
L136:
	;
	v464 = base.F64_sub(base.F64_add(v399, v450), base.F64_mul(v399, v450))
	v465 = v457
	goto L130
L137:
	;
	v464 = v458
	v465 = base.F64_add(v401, base.F64_add(v450, float64(-1)))
	goto L130
L138:
	;
	goto L121
L139:
	;
	if base.F64_le(v478, float64(1)) != 0 {
		goto L145
	} else {
		goto L146
	}
L140:
	;
	if v346 != 0 {
		goto L139
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if v347 == int32(0) {
		v706 = v476
		goto L99
	} else {
		goto L144
	}
L143:
	;
	v706 = v476
	goto L99
L144:
	;
	goto L139
L145:
	;
	v499 = v478
	goto L147
L146:
	;
	v499 = v476
	goto L147
L147:
	;
	if base.F64_ge(v478, float64(0)) != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v502 = v499
	goto L150
L149:
	;
	v502 = v476
	goto L150
L150:
	;
	v706 = v502
	goto L99
L151:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	F_get_typlenbyval(m, v504, v26+int32(60), v26+int32(56))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v514 = v28 & int32(1)
	if v514 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v515 = float64(0)
	goto L155
L154:
	;
	v515 = float64(1)
	goto L155
L155:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	if v516 != 0 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v529 = v515
	v531 = v515
	v534 = v517
	goto L161
L157:
	;
	v517 = int32(0)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v517 < v518 {
		goto L156
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v675 = v515
	v677 = v515
	goto L100
L160:
	;
	goto L159
L161:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v546+v534<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v550
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v550
	v559 = F_list_make2_impl(m, v26+int32(24), v26+int32(20))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L163
	}
L162:
	;
	v675 = v585
	v677 = v586
	goto L100
L163:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if l2 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v571 = *(*float64)(unsafe.Add(mBase, uint32(v570)))
	if v28&int32(1) != 0 {
		goto L171
	} else {
		goto L172
	}
L165:
	;
	v564 = F_FunctionCall5Coll(m, v26+int32(68), v561, l0, v29, v559, base.I32_extend16_s(l4), l5)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v568 = F_FunctionCall4Coll(m, v26+int32(68), v561, l0, v29, v559, l3)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L169
	}
L168:
	;
	v570 = v564
	goto L164
L169:
	;
	v570 = v568
	goto L164
L170:
	;
	v588 = v534 + int32(1)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v588 < v589 {
		v529 = v585
		v531 = v586
		v534 = v588
		goto L161
	} else {
		goto L178
	}
L171:
	;
	if v346 != 0 {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	goto L173
L173:
	;
	v579 = base.F64_mul(v529, v571)
	if v347 == int32(0) {
		v585 = v579
		v586 = v531
		goto L170
	} else {
		goto L177
	}
L174:
	;
	v578 = base.F64_add(v531, v571)
	goto L176
L175:
	;
	v578 = v531
	goto L176
L176:
	;
	v585 = base.F64_sub(base.F64_add(v529, v571), base.F64_mul(v529, v571))
	v586 = v578
	goto L170
L177:
	;
	v585 = v579
	v586 = base.F64_add(v531, base.F64_add(v571, float64(-1)))
	goto L170
L178:
	;
	goto L162
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v593)+8)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v593)+4)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v593))) = int32(34)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v593)+12)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v593
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v593
	v610 = F_list_make2_impl(m, v26+int32(8), v26+int32(4))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if l2 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v623 = *(*float64)(unsafe.Add(mBase, uint32(v622)))
	if v28&int32(1) == int32(0) {
		goto L187
	} else {
		goto L188
	}
L182:
	;
	v616 = F_FunctionCall5Coll(m, v26+int32(68), v612, l0, v29, v610, base.I32_extend16_s(l4), l5)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v620 = F_FunctionCall4Coll(m, v26+int32(68), v612, l0, v29, v610, l3)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L186
	}
L185:
	;
	v622 = v616
	goto L181
L186:
	;
	v622 = v620
	goto L181
L187:
	;
	v706 = base.F64_mul(v623, base.F64_mul(v623, base.F64_mul(v623, base.F64_mul(v623, base.F64_mul(v623, base.F64_mul(v623, base.F64_mul(v623, base.F64_mul(v623, base.F64_mul(v623, v623)))))))))
	goto L99
L188:
	;
	goto L189
L189:
	;
	v641 = base.F64_add(base.F64_mul(v623, math.Float64frombits(uint64(0x8000000000000000))), base.F64_add(v623, float64(0)))
	v644 = base.F64_sub(base.F64_add(v623, v641), base.F64_mul(v641, v623))
	v647 = base.F64_sub(base.F64_add(v623, v644), base.F64_mul(v644, v623))
	v650 = base.F64_sub(base.F64_add(v623, v647), base.F64_mul(v647, v623))
	v653 = base.F64_sub(base.F64_add(v623, v650), base.F64_mul(v650, v623))
	v656 = base.F64_sub(base.F64_add(v623, v653), base.F64_mul(v653, v623))
	v659 = base.F64_sub(base.F64_add(v623, v656), base.F64_mul(v656, v623))
	v662 = base.F64_sub(base.F64_add(v623, v659), base.F64_mul(v659, v623))
	v665 = base.F64_sub(base.F64_add(v623, v662), base.F64_mul(v662, v623))
	v706 = base.F64_sub(base.F64_add(v623, v665), base.F64_mul(v665, v623))
	goto L99
L190:
	;
	if base.F64_le(v677, float64(1)) != 0 {
		goto L196
	} else {
		goto L197
	}
L191:
	;
	if v346 != 0 {
		goto L190
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	if v347 == int32(0) {
		v706 = v675
		goto L99
	} else {
		goto L195
	}
L194:
	;
	v706 = v675
	goto L99
L195:
	;
	goto L190
L196:
	;
	v696 = v677
	goto L198
L197:
	;
	v696 = v675
	goto L198
L198:
	;
	if base.F64_ge(v677, float64(0)) != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v699 = v696
	goto L201
L200:
	;
	v699 = v675
	goto L201
L201:
	;
	v706 = v699
	goto L99
L202:
	;
	if base.F64_gt(v706, float64(1)) == int32(0) {
		v738 = v706
		goto L3
	} else {
		goto L203
	}
L203:
	;
	v738 = float64(1)
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
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
	var v183 int32
	_ = v183
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v434 int32
	_ = v434
	var v453 int32
	_ = v453
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	if l0&int32(3) == int32(0) {
		v45 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = int32(16777216)
	v81 = F_pg_hmac_create(m, l1)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v78 = v70 - l0
	goto L1
L3:
	;
	v49 = v45
	goto L12
L4:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v29 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v78 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v34 = l0
	goto L8
L8:
	;
	v38 = v34 + int32(1)
	if v38&int32(3) == int32(0) {
		v45 = v38
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v70 = v38
	goto L2
L10:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v43 != 0 {
		v34 = v38
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v58 = int32(-2139062144)
	if (int32(16843008)-v55|v55)&v58 == v58 {
		v49 = v49 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v64 = v49
	goto L15
L14:
	;
	goto L13
L15:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v68 != 0 {
		v64 = v64 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v70 = v64
	goto L2
L17:
	;
	goto L16
L18:
	;
	m.G0 = v20 + int32(80)
	return v453
L19:
	;
	return int32(0)
L20:
	;
	if v81 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	goto L25
L22:
	;
	goto L23
L23:
	;
	v110 = F_pg_hmac_init(m, v81, l0, v78)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L40
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(_a_F_scram_SaltedPassword_0)
	v453 = int32(-1)
	goto L18
L25:
	;
	goto L24
L37:
	;
	F_pg_hmac_free(m, v81)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L19
	} else {
		goto L130
	}
L38:
	;
	if l2 != 0 {
		goto L74
	} else {
		goto L75
	}
L39:
	;
	if v81 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L40:
	;
	if v110 < int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	if v81 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v128 < int32(0) {
		goto L39
	} else {
		goto L49
	}
L43:
	;
	v128 = int32(-1)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v118 = F_pg_cryptohash_update(m, v117, l3, l4)
	mBase = m.M
	if int32(0) <= v118 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v128 = int32(0)
	goto L42
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = int32(2)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v125 = F_pg_cryptohash_error(m, v124)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = v125
	v128 = int32(-1)
	goto L42
L49:
	;
	if v81 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v148 < int32(0) {
		goto L39
	} else {
		goto L57
	}
L51:
	;
	v148 = int32(-1)
	goto L50
L52:
	;
	goto L53
L53:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v138 = F_pg_cryptohash_update(m, v137, v20+int32(76), int32(4))
	mBase = m.M
	if int32(0) <= v138 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v148 = int32(0)
	goto L50
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = int32(2)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v145 = F_pg_cryptohash_error(m, v144)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = v145
	v148 = int32(-1)
	goto L50
L57:
	;
	v151 = F_pg_hmac_final(m, v81, v20, l2)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L19
	} else {
		goto L58
	}
L58:
	;
	if int32(0) <= v151 {
		goto L38
	} else {
		goto L59
	}
L59:
	;
	goto L39
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v174
	goto L37
L61:
	;
	v174 = int32(_a_F_scram_SaltedPassword_0)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v159 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v171 = v159
	goto L66
L65:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v163 == int32(2) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v174 = v171
	goto L60
L67:
	;
	v166 = int32(_a_F_scram_SaltedPassword_1)
	goto L69
L68:
	;
	v166 = int32(_a_F_scram_SaltedPassword_2)
	goto L69
L69:
	;
	if v163 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v169 = int32(_a_F_scram_SaltedPassword_0)
	goto L72
L71:
	;
	v169 = v166
	goto L72
L72:
	;
	v171 = v169
	goto L66
L73:
	;
	if int32(2) <= l5 {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v176 = F__emscripten_memcpy_bulkmem(m, l6, v20, l2)
	mBase = m.M
	v177 = v176
	goto L76
L75:
	;
	v177 = l6
	goto L76
L76:
	;
	goto L73
L77:
	;
	if v81 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L78:
	;
	v183 = l2 & int32(3)
	v201 = int32(1)
	goto L81
L79:
	;
	goto L80
L80:
	;
	F_pg_hmac_free(m, v81)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L19
	} else {
		goto L116
	}
L81:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_scram_SaltedPassword[0]))
	if v205 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L80
L83:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L19
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v208 = F_pg_hmac_init(m, v81, l0, v78)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L19
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	if v208 < int32(0) {
		goto L77
	} else {
		goto L88
	}
L88:
	;
	if v81 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if v226 < int32(0) {
		goto L77
	} else {
		goto L96
	}
L90:
	;
	v226 = int32(-1)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v216 = F_pg_cryptohash_update(m, v215, v20, l2)
	mBase = m.M
	if int32(0) <= v216 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v226 = int32(0)
	goto L89
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = int32(2)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v223 = F_pg_cryptohash_error(m, v222)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = v223
	v226 = int32(-1)
	goto L89
L96:
	;
	v231 = F_pg_hmac_final(m, v81, v20+int32(32), l2)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L19
	} else {
		goto L97
	}
L97:
	;
	if v231 < int32(0) {
		goto L77
	} else {
		goto L98
	}
L98:
	;
	if l2 <= int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if l2 != 0 {
		goto L112
	} else {
		goto L113
	}
L100:
	;
	v237 = int32(0)
	if base.B2i32(base.Ui32(l2) < base.Ui32(int32(4))) == v237 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v243 = v237
	v248 = v237
	goto L104
L102:
	;
	v303 = v237
	goto L103
L103:
	;
	if v183 == int32(0) {
		goto L99
	} else {
		goto L107
	}
L104:
	;
	v259 = v243 + v177
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	v262 = v20 + int32(32)
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262+v243))))
	v265 = v260 ^ v264
	*(*uint8)(unsafe.Add(mBase, uint32(v259))) = uint8(v265)
	v268 = v243 | int32(1)
	v269 = v177 + v268
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262+v268))))
	v275 = v270 ^ v274
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v275)
	v278 = v243 | int32(2)
	v279 = v177 + v278
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262+v278))))
	v285 = v280 ^ v284
	*(*uint8)(unsafe.Add(mBase, uint32(v279))) = uint8(v285)
	v288 = v243 | int32(3)
	v289 = v177 + v288
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262+v288))))
	v295 = v290 ^ v294
	*(*uint8)(unsafe.Add(mBase, uint32(v289))) = uint8(v295)
	v297 = int32(4)
	v298 = v243 + v297
	v300 = v248 + v297
	if v300 != l2&int32(2147483644) {
		v243 = v298
		v248 = v300
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v303 = v298
	goto L103
L106:
	;
	goto L105
L107:
	;
	v322 = v303
	v324 = v237
	goto L108
L108:
	;
	v338 = v322 + v177
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(32)+v322))))
	v344 = v339 ^ v343
	*(*uint8)(unsafe.Add(mBase, uint32(v338))) = uint8(v344)
	v346 = int32(1)
	v349 = v324 + v346
	if v349 != v183 {
		v322 = v322 + v346
		v324 = v349
		goto L108
	} else {
		goto L110
	}
L109:
	;
	goto L99
L110:
	;
	goto L109
L111:
	;
	v373 = v201 + int32(1)
	if v373 != l5 {
		v201 = v373
		goto L81
	} else {
		goto L115
	}
L112:
	;
	v370 = F__emscripten_memcpy_bulkmem(m, v20, v20+int32(32), l2)
	mBase = m.M
	goto L114
L113:
	;
	goto L114
L114:
	;
	goto L111
L115:
	;
	goto L82
L116:
	;
	v453 = int32(0)
	goto L18
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v414
	goto L37
L118:
	;
	v414 = int32(_a_F_scram_SaltedPassword_0)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v399 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v411 = v399
	goto L123
L122:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v403 == int32(2) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v414 = v411
	goto L117
L124:
	;
	v406 = int32(_a_F_scram_SaltedPassword_1)
	goto L126
L125:
	;
	v406 = int32(_a_F_scram_SaltedPassword_2)
	goto L126
L126:
	;
	if v403 == int32(1) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v409 = int32(_a_F_scram_SaltedPassword_0)
	goto L129
L128:
	;
	v409 = v406
	goto L129
L129:
	;
	v411 = v409
	goto L123
L130:
	;
	v453 = int32(-1)
	goto L18
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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v311 int32
	_ = v311
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v347 int32
	_ = v347
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
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
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_scram_init[0])))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v22 == v15 {
		v41 = v21
		v42 = v22
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L82
	}
L4:
	;
	F_pg_cryptohash_free(m, v103)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L78
	}
L5:
	;
	if v42-v41 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	goto L5
L7:
	;
	if v21 != v22 {
		v41 = v21
		v42 = v22
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v26 = l1
	v27 = v18
	goto L9
L9:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v31 == int32(0) {
		v41 = v30
		v42 = v31
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v41 = v30
	v42 = v31
	goto L6
L11:
	;
	v34 = int32(1)
	if v30 == v31 {
		v26 = v26 + v34
		v27 = v27 + v34
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)) = uint8(v46)
	if l2 == v46 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L74
	}
L16:
	;
	m.G0 = v8 + int32(32)
	return v11
L17:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+364))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(137438953475)
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_scram_init[1]))
	v103 = F_pg_cryptohash_create(m, int32(3))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L30
	}
L18:
	;
	v50 = F_get_password_type(m, l2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v50 == int32(2) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v66 = F_parse_scram_secret(m, l2, v11+int32(24), v11+int32(16), v11+int32(20), v11+int32(28), v11-int32(-64), v11+int32(96))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v86
	v91 = F_psprintf(m, int32(_a_F_scram_init_1), v8+int32(16))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L29
	}
L23:
	;
	if v66 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v70 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v70 == int32(0) {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v75
	F_errmsg(m, int32(_a_F_scram_init_2), v8)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_scram_init_3), int32(292), int32(_a_F_scram_init_4))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L17
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v91
	goto L17
L30:
	;
	v105 = F_pg_cryptohash_init(m, v103)
	mBase = m.M
	if v105 < int32(0) {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v95&int32(3) == int32(0) {
		v131 = v95
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v165 = F_pg_cryptohash_update(m, v103, v95, v164)
	mBase = m.M
	if v165 < int32(0) {
		goto L4
	} else {
		goto L49
	}
L33:
	;
	v164 = v156 - v95
	goto L32
L34:
	;
	v135 = v131
	goto L43
L35:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v115 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v164 = int32(0)
	goto L32
L37:
	;
	goto L38
L38:
	;
	v120 = v95
	goto L39
L39:
	;
	v124 = v120 + int32(1)
	if v124&int32(3) == int32(0) {
		v131 = v124
		goto L34
	} else {
		goto L41
	}
L40:
	;
	v156 = v124
	goto L33
L41:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v129 != 0 {
		v120 = v124
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v144 = int32(-2139062144)
	if (int32(16843008)-v141|v141)&v144 == v144 {
		v135 = v135 + int32(4)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v150 = v135
	goto L46
L45:
	;
	goto L44
L46:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if v154 != 0 {
		v150 = v150 + int32(1)
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v156 = v150
	goto L33
L48:
	;
	goto L47
L49:
	;
	v169 = F_pg_cryptohash_update(m, v103, v99+int32(257), int32(32))
	mBase = m.M
	if v169 < int32(0) {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v174 = F_pg_cryptohash_final(m, v103, int32(_a_F_scram_init_5), int32(32))
	mBase = m.M
	if v174 < int32(0) {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	F_pg_cryptohash_free(m, v103)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v185 = base.I32_div_s(int32(18), int32(3))
	v187 = v185 << (uint(int32(2)) % 32)
	goto L53
L53:
	;
	v190 = F_palloc(m, v187+int32(1))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L58
L55:
	;
	if v320 < int32(0) {
		goto L3
	} else {
		goto L73
	}
L56:
	;
	v311 = F___memset(m, v190, int32(0), v187)
	mBase = m.M
	v320 = int32(-1)
	goto L55
L57:
	;
	if v187 < v254-v190+int32(4) {
		goto L56
	} else {
		goto L69
	}
L58:
	;
	v199 = int32(_a_F_scram_init_5)
	v200 = int32(0)
	v203 = v190
	v204 = int32(2)
	goto L61
L60:
	;
	v320 = v254 - v190
	goto L55
L61:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	v210 = v206<<(uint(v204<<(uint(int32(3))%32))%32) | v200
	if int32(0) < v204 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v255 != int32(2) {
		goto L57
	} else {
		goto L68
	}
L63:
	;
	v253 = v210
	v254 = v203
	v255 = v204 - int32(1)
	goto L65
L64:
	;
	if v187 < v203-v190+int32(4) {
		goto L56
	} else {
		goto L66
	}
L65:
	;
	v257 = v199 + int32(1)
	if v257 != int32(_a_F_scram_init_6) {
		v199 = v257
		v200 = v253
		v203 = v254
		v204 = v255
		goto L61
	} else {
		goto L67
	}
L66:
	;
	v219 = int32(63)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210&v219)+uint32(_c_F_scram_init[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+3)) = uint8(v223)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v210)>>(uint(int32(6))%32))&v219)+uint32(_c_F_scram_init[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+2)) = uint8(v231)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v210)>>(uint(int32(12))%32))&v219)+uint32(_c_F_scram_init[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)) = uint8(v239)
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v210)>>(uint(int32(18))%32))&v219)+uint32(_c_F_scram_init[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v247)
	v253 = int32(0)
	v254 = v203 + int32(4)
	v255 = int32(2)
	goto L65
L67:
	;
	goto L62
L68:
	;
	goto L60
L69:
	;
	v275 = int32(63)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v253)>>(uint(int32(12))%32))&v275)+uint32(_c_F_scram_init[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)) = uint8(v279)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v253)>>(uint(int32(18))%32))&v275)+uint32(_c_F_scram_init[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v254))) = uint8(v287)
	if v255 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v253)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_scram_init[2]))))
	v299 = v298
	goto L72
L71:
	;
	v299 = int32(61)
	goto L72
L72:
	;
	v300 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v254)+3)) = uint8(v300)
	*(*uint8)(unsafe.Add(mBase, uint32(v254)+2)) = uint8(v299)
	v320 = v254 + int32(4) - v190
	goto L55
L73:
	;
	v324 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v320+v190))) = uint8(v324)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(_a_F_scram_init_7)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v190
	v330 = v11 - int32(-64)
	v331 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v330)+56)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v330)+48)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v330)+40)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v330)+32)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v330)+24)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v330)+16)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v330)+8)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v330))) = v331
	v347 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+192)) = uint8(v347)
	goto L16
L74:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(_a_F_scram_init_8), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_scram_init_3), int32(267), int32(_a_F_scram_init_4))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errmsg_internal(m, int32(_a_F_scram_init_9), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_scram_init_3), int32(719), int32(_a_F_scram_init_10))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errmsg_internal(m, int32(_a_F_scram_init_9), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_scram_init_3), int32(728), int32(_a_F_scram_init_10))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	v17 = F_pgl_send(m, v15, l1, l2, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
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
		goto L23
	}
L4:
	;
	F_ProcessClientWriteInterrupt(m, int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L22
	}
L5:
	;
	if int32(0) <= v17 {
		v69 = v17
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v21 != 0 {
		v69 = v17
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v25 = v17
	goto L8
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_secure_write[0]))
	if v28 != int32(6) {
		v69 = v25
		goto L4
	} else {
		goto L10
	}
L9:
	;
	v69 = v59
	goto L4
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_secure_write[1]))
	v33 = int32(0)
	F_ModifyWaitEvent(m, v32, v33, int32(4), v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_secure_write[1]))
	v43 = F_WaitEventSetWait(m, v39, int32(-1), v8, int32(1), int32(100663297))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v45&int32(16) != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	if v45&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_secure_write[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = F_pgl_send(m, v57, l1, l2, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	F_ProcessClientWriteInterrupt(m, int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	if int32(0) <= v59 {
		v69 = v59
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v63 == int32(0) {
		v25 = v59
		goto L8
	} else {
		goto L21
	}
L21:
	;
	goto L9
L22:
	;
	m.G0 = v8 + int32(16)
	return v69
L23:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errmsg(m, int32(_a_F_secure_write_0), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_secure_write_1), int32(350), int32(_a_F_secure_write_2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
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
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
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
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v322 int64
	_ = v322
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int64
	_ = v422
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v780 int32
	_ = v780
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
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
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v948 int32
	_ = v948
	var v955 int32
	_ = v955
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v999 int32
	_ = v999
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1028 int32
	_ = v1028
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1108 int32
	_ = v1108
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1173 int32
	_ = v1173
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1231 int32
	_ = v1231
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1249 int32
	_ = v1249
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1282 int32
	_ = v1282
	var v1320 int32
	_ = v1320
	var v1346 int32
	_ = v1346
	var v1397 int32
	_ = v1397
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1566 int32
	_ = v1566
	var v1575 int32
	_ = v1575
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1606 int32
	_ = v1606
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1690 int32
	_ = v1690
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1826 int32
	_ = v1826
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1846 int32
	_ = v1846
	var v1854 int32
	_ = v1854
	var v1860 int32
	_ = v1860
	var v1866 int32
	_ = v1866
	var v1870 int32
	_ = v1870
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1886 int32
	_ = v1886
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1901 int32
	_ = v1901
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1913 int32
	_ = v1913
	var v1921 int32
	_ = v1921
	var v1924 int32
	_ = v1924
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1943 int32
	_ = v1943
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1960 int32
	_ = v1960
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1975 int32
	_ = v1975
	var v1982 int32
	_ = v1982
	var v1987 int32
	_ = v1987
	var v1990 int32
	_ = v1990
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2009 int32
	_ = v2009
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2067 int32
	_ = v2067
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2079 int32
	_ = v2079
	var v2093 int32
	_ = v2093
	var v2130 int64
	_ = v2130
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2178 int32
	_ = v2178
	var v2183 int32
	_ = v2183
	var v2196 int32
	_ = v2196
	var v2202 int64
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2235 int32
	_ = v2235
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2249 int32
	_ = v2249
	var v2257 int32
	_ = v2257
	var v2262 int32
	_ = v2262
	var v2266 int32
	_ = v2266
	var v2271 int32
	_ = v2271
	var v2273 int32
	_ = v2273
	var v2277 int32
	_ = v2277
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2292 int32
	_ = v2292
	var v2296 int32
	_ = v2296
	var v2298 int32
	_ = v2298
	var v2306 int32
	_ = v2306
	var v2312 int32
	_ = v2312
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2332 int32
	_ = v2332
	var v2334 int32
	_ = v2334
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2345 int32
	_ = v2345
	var v2349 int32
	_ = v2349
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2364 int32
	_ = v2364
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2375 int32
	_ = v2375
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2392 int32
	_ = v2392
	var v2394 int32
	_ = v2394
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2409 int32
	_ = v2409
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2436 int32
	_ = v2436
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2447 int32
	_ = v2447
	var v2451 int32
	_ = v2451
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2466 int32
	_ = v2466
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2477 int32
	_ = v2477
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
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
	var v2544 int32
	_ = v2544
	var v2546 int32
	_ = v2546
	var v2550 int32
	_ = v2550
	var v2554 int32
	_ = v2554
	var v2558 int32
	_ = v2558
	var v2562 int32
	_ = v2562
	var v2566 int32
	_ = v2566
	var v2571 int32
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2577 int32
	_ = v2577
	var v2588 int32
	_ = v2588
	var v2628 int32
	_ = v2628
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2644 int32
	_ = v2644
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2658 int32
	_ = v2658
	var v2662 int32
	_ = v2662
	var v2714 int32
	_ = v2714
	var v2716 int32
	_ = v2716
	var v2720 int32
	_ = v2720
	var v2727 int32
	_ = v2727
	var v2729 int32
	_ = v2729
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2743 int32
	_ = v2743
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2772 int32
	_ = v2772
	var v2777 int32
	_ = v2777
	var v2781 int32
	_ = v2781
	var v2786 int32
	_ = v2786
	var v2788 int32
	_ = v2788
	var v2792 int32
	_ = v2792
	var v2798 int32
	_ = v2798
	var v2801 int32
	_ = v2801
	var v2807 int32
	_ = v2807
	var v2811 int32
	_ = v2811
	var v2813 int32
	_ = v2813
	var v2821 int32
	_ = v2821
	var v2827 int32
	_ = v2827
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2840 int32
	_ = v2840
	var v2841 int32
	_ = v2841
	var v2842 int32
	_ = v2842
	var v2844 int32
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2847 int32
	_ = v2847
	var v2849 int32
	_ = v2849
	var v2853 int32
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2860 int32
	_ = v2860
	var v2864 int32
	_ = v2864
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2875 int32
	_ = v2875
	var v2876 int32
	_ = v2876
	var v2877 int32
	_ = v2877
	var v2879 int32
	_ = v2879
	var v2882 int32
	_ = v2882
	var v2883 int32
	_ = v2883
	var v2884 int32
	_ = v2884
	var v2885 int32
	_ = v2885
	var v2886 int32
	_ = v2886
	var v2890 int32
	_ = v2890
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2904 int32
	_ = v2904
	var v2905 int32
	_ = v2905
	var v2907 int32
	_ = v2907
	var v2909 int32
	_ = v2909
	var v2913 int32
	_ = v2913
	var v2914 int32
	_ = v2914
	var v2918 int32
	_ = v2918
	var v2919 int32
	_ = v2919
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2924 int32
	_ = v2924
	var v2928 int32
	_ = v2928
	var v2929 int32
	_ = v2929
	var v2933 int32
	_ = v2933
	var v2934 int32
	_ = v2934
	var v2936 int32
	_ = v2936
	var v2937 int32
	_ = v2937
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2940 int32
	_ = v2940
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2944 int32
	_ = v2944
	var v2946 int32
	_ = v2946
	var v2947 int32
	_ = v2947
	var v2949 int32
	_ = v2949
	var v2951 int32
	_ = v2951
	var v2955 int32
	_ = v2955
	var v2956 int32
	_ = v2956
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2962 int32
	_ = v2962
	var v2966 int32
	_ = v2966
	var v2970 int32
	_ = v2970
	var v2971 int32
	_ = v2971
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2977 int32
	_ = v2977
	var v2978 int32
	_ = v2978
	var v2979 int32
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
	var v2988 int32
	_ = v2988
	var v2992 int32
	_ = v2992
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3013 int32
	_ = v3013
	var v3017 int32
	_ = v3017
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3025 int32
	_ = v3025
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3036 int32
	_ = v3036
	var v3040 int32
	_ = v3040
	var v3041 int32
	_ = v3041
	var v3042 int32
	_ = v3042
	var v3043 int32
	_ = v3043
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3049 int32
	_ = v3049
	var v3050 int32
	_ = v3050
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3056 int32
	_ = v3056
	var v3059 int32
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3065 int32
	_ = v3065
	var v3069 int32
	_ = v3069
	var v3073 int32
	_ = v3073
	var v3077 int32
	_ = v3077
	var v3081 int32
	_ = v3081
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3093 int32
	_ = v3093
	var v3106 int32
	_ = v3106
	var v3146 int32
	_ = v3146
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3154 int32
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3158 int32
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3162 int32
	_ = v3162
	var v3169 int32
	_ = v3169
	var v3170 int32
	_ = v3170
	var v3176 int32
	_ = v3176
	var v3180 int32
	_ = v3180
	var v3228 int32
	_ = v3228
	var v3233 int32
	_ = v3233
	var v3235 int32
	_ = v3235
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3245 int32
	_ = v3245
	var v3247 int64
	_ = v3247
	var v3249 int32
	_ = v3249
	var v3252 int32
	_ = v3252
	var v3253 int32
	_ = v3253
	var v3254 int32
	_ = v3254
	var v3255 int32
	_ = v3255
	var v3258 int32
	_ = v3258
	var v3259 int32
	_ = v3259
	var v3262 int32
	_ = v3262
	var v3263 int32
	_ = v3263
	var v3269 int32
	_ = v3269
	var v3272 int32
	_ = v3272
	var v3275 int32
	_ = v3275
	var v3276 int32
	_ = v3276
	var v3280 int32
	_ = v3280
	var v3290 int32
	_ = v3290
	var v3299 int32
	_ = v3299
	var v3301 int32
	_ = v3301
	var v3303 int32
	_ = v3303
	var v3308 int32
	_ = v3308
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3312 int32
	_ = v3312
	var v3314 int64
	_ = v3314
	var v3316 int32
	_ = v3316
	var v3319 int32
	_ = v3319
	var v3320 int32
	_ = v3320
	var v3321 int32
	_ = v3321
	var v3322 int32
	_ = v3322
	var v3325 int32
	_ = v3325
	var v3326 int32
	_ = v3326
	var v3329 int32
	_ = v3329
	var v3330 int32
	_ = v3330
	var v3336 int32
	_ = v3336
	var v3339 int32
	_ = v3339
	var v3342 int32
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3347 int32
	_ = v3347
	var v3357 int32
	_ = v3357
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3377 int32
	_ = v3377
	var v3379 int32
	_ = v3379
	var v3380 int32
	_ = v3380
	var v3381 int32
	_ = v3381
	var v3388 int32
	_ = v3388
	var v3389 int32
	_ = v3389
	var v3391 int32
	_ = v3391
	var v3417 int32
	_ = v3417
	var v3425 int32
	_ = v3425
	var v3444 int32
	_ = v3444
	var v3448 int32
	_ = v3448
	var v3453 int32
	_ = v3453
	var v3456 int32
	_ = v3456
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3464 int32
	_ = v3464
	var v3496 int32
	_ = v3496
	var v3500 int32
	_ = v3500
	var v3518 int32
	_ = v3518
	var v3527 int32
	_ = v3527
	var v3530 int32
	_ = v3530
	var v3532 int32
	_ = v3532
	var v3560 int32
	_ = v3560
	var v3563 int32
	_ = v3563
	var v3587 int32
	_ = v3587
	var v3599 int32
	_ = v3599
	var v3602 int32
	_ = v3602
	var v3604 int32
	_ = v3604
	var v3634 int32
	_ = v3634
	var v3654 int32
	_ = v3654
	var v3684 int32
	_ = v3684
	var v3713 int32
	_ = v3713
	var v3719 int32
	_ = v3719
	var v3720 int32
	_ = v3720
	var v3736 int32
	_ = v3736
	var v3750 int32
	_ = v3750
	var v3775 int32
	_ = v3775
	var v3776 int32
	_ = v3776
	var v3779 int32
	_ = v3779
	var v3780 int32
	_ = v3780
	var v3781 int32
	_ = v3781
	var v3785 int32
	_ = v3785
	var v3786 int32
	_ = v3786
	var v3790 int32
	_ = v3790
	var v3791 int32
	_ = v3791
	var v3795 int32
	_ = v3795
	var v3797 int32
	_ = v3797
	var v3809 int32
	_ = v3809
	var v3858 int32
	_ = v3858
	var v3871 int32
	_ = v3871
	var v3897 int32
	_ = v3897
	var v3898 int32
	_ = v3898
	var v3901 int32
	_ = v3901
	var v3904 int32
	_ = v3904
	var v3955 int32
	_ = v3955
	var v3956 int32
	_ = v3956
	var v3960 int32
	_ = v3960
	var v3962 int32
	_ = v3962
	var v3965 int32
	_ = v3965
	var v3979 int32
	_ = v3979
	var v4042 int32
	_ = v4042
	var v4072 int32
	_ = v4072
	var v4075 int32
	_ = v4075
	var v4080 int32
	_ = v4080
	var v4085 int32
	_ = v4085
	var v4088 int32
	_ = v4088
	var v4090 int32
	_ = v4090
	var v4092 int32
	_ = v4092
	var v4096 int32
	_ = v4096
	var v4102 int32
	_ = v4102
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4124 int32
	_ = v4124
	var v4125 int32
	_ = v4125
	var v4127 int32
	_ = v4127
	var v4144 int32
	_ = v4144
	var v4158 int32
	_ = v4158
	var v4182 int32
	_ = v4182
	var v4183 int32
	_ = v4183
	var v4184 int32
	_ = v4184
	var v4185 int32
	_ = v4185
	var v4186 int32
	_ = v4186
	var v4187 int32
	_ = v4187
	var v4190 int64
	_ = v4190
	var v4204 int32
	_ = v4204
	var v4205 int32
	_ = v4205
	var v4215 int32
	_ = v4215
	var v4220 int32
	_ = v4220
	var v4266 int64
	_ = v4266
	var v4268 int32
	_ = v4268
	var v4269 int32
	_ = v4269
	var v4279 int32
	_ = v4279
	var v4296 int32
	_ = v4296
	var v4301 int32
	_ = v4301
	var v4315 int64
	_ = v4315
	var v4318 int32
	_ = v4318
	var v4320 int32
	_ = v4320
	var v4328 int32
	_ = v4328
	var v4331 int32
	_ = v4331
	var v4335 int32
	_ = v4335
	var v4339 int32
	_ = v4339
	var v4344 int32
	_ = v4344
	var v4348 int32
	_ = v4348
	var v4350 int32
	_ = v4350
	var v4358 int32
	_ = v4358
	var v4363 int32
	_ = v4363
	var v4367 int32
	_ = v4367
	var v4370 int32
	_ = v4370
	var v4378 int32
	_ = v4378
	var v4383 int32
	_ = v4383
	v10 = int32(0)
	v46 = int64(0)
	v48 = m.G0
	v50 = v48 - int32(_a_F_sendDir_0)
	m.G0 = v50
	if l8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v53 = F_palloc(m, int32(_a_F_sendDir_1))
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
	v62 = int32(0)
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
	v361 = F_AllocateDir(m, l1)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L4
	} else {
		goto L106
	}
L7:
	;
	v327 = int32(_a_F_sendDir_2)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[0])))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v331 == int32(0) {
		v350 = v330
		v351 = v331
		goto L93
	} else {
		goto L94
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
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v64 != int32(47) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v61 = v61 + int32(1)
	v62 = v67
	goto L9
L12:
	;
	if v64 != 0 {
		v67 = v62
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
	v73 = v62 + int32(1)
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
	if v73&int32(3) == int32(0) {
		v181 = v73
		goto L40
	} else {
		goto L41
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
		v149 = v73
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v157 = v149 - v73
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
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v149 = v145
	goto L30
L34:
	;
	v149 = v128
	goto L30
L35:
	;
	goto L36
L36:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	v145 = v128 + int32(1)
	if v143 != 0 {
		v128 = v145
		v129 = v143
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	if v157 != v214 {
		goto L7
	} else {
		goto L55
	}
L39:
	;
	v214 = v206 - v73
	goto L38
L40:
	;
	v185 = v181
	goto L49
L41:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v165 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v214 = int32(0)
	goto L38
L43:
	;
	goto L44
L44:
	;
	v170 = v73
	goto L45
L45:
	;
	v174 = v170 + int32(1)
	if v174&int32(3) == int32(0) {
		v181 = v174
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v206 = v174
	goto L39
L47:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v179 != 0 {
		v170 = v174
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v194 = int32(-2139062144)
	if (int32(16843008)-v191|v191)&v194 == v194 {
		v185 = v185 + int32(4)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v200 = v185
	goto L52
L51:
	;
	goto L50
L52:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v204 != 0 {
		v200 = v200 + int32(1)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v206 = v200
	goto L39
L54:
	;
	goto L53
L55:
	;
	v216 = int32(_a_F_sendDir_4)
	v217 = v62 - l1
	if v217 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v261 != 0 {
		goto L70
	} else {
		goto L71
	}
L57:
	;
	v261 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v223 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v224 = l1
	v225 = v216
	v226 = v217
	v227 = v223
	goto L64
L61:
	;
	v249 = v216
	v253 = int32(0)
	goto L62
L62:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	v261 = v253 - v254
	goto L56
L63:
	;
	v249 = v244
	v253 = v246
	goto L62
L64:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if v227 != v229 {
		v244 = v225
		v246 = v227
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v244 = v238
	v246 = int32(0)
	goto L63
L66:
	;
	if v229 == int32(0) {
		v244 = v225
		v246 = v227
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v234 = v226 - int32(1)
	if v234 == int32(0) {
		v244 = v225
		v246 = v227
		goto L63
	} else {
		goto L68
	}
L68:
	;
	v237 = int32(1)
	v238 = v225 + v237
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+1)))
	if v239 != 0 {
		v224 = v224 + v237
		v225 = v238
		v226 = v234
		v227 = v239
		goto L64
	} else {
		goto L69
	}
L69:
	;
	goto L65
L70:
	;
	v262 = int32(1663)
	if base.Ui32(v217) < base.Ui32(int32(15)) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v322 = F_strtox_2(m, v73, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L91
L73:
	;
	v357 = int32(0)
	v359 = v10
	v360 = v262
	goto L6
L74:
	;
	goto L75
L75:
	;
	v267 = int32(15)
	v268 = v62 - v267
	v269 = int32(_a_F_sendDir_5)
	goto L78
L76:
	;
	if v306-v307 != 0 {
		v357 = int32(0)
		v359 = v10
		v360 = v262
		goto L6
	} else {
		goto L90
	}
L78:
	;
	goto L79
L79:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v276 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v277 = v268
	v278 = v269
	v279 = v267
	v280 = v276
	goto L84
L81:
	;
	v302 = v269
	v306 = int32(0)
	goto L82
L82:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	goto L76
L83:
	;
	v302 = v297
	v306 = v299
	goto L82
L84:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v280 != v282 {
		v297 = v278
		v299 = v280
		goto L83
	} else {
		goto L86
	}
L85:
	;
	v297 = v291
	v299 = int32(0)
	goto L83
L86:
	;
	if v282 == int32(0) {
		v297 = v278
		v299 = v280
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v287 = v279 - int32(1)
	if v287 == int32(0) {
		v297 = v278
		v299 = v280
		goto L83
	} else {
		goto L88
	}
L88:
	;
	v290 = int32(1)
	v291 = v278 + v290
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)))
	if v292 != 0 {
		v277 = v277 + v290
		v278 = v291
		v279 = v287
		v280 = v292
		goto L84
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	goto L72
L91:
	;
	v357 = int32(1)
	v359 = base.I32_wrap_i64(v322)
	v360 = int32(1663)
	goto L6
L92:
	;
	if v352 != 0 {
		goto L100
	} else {
		goto L101
	}
L93:
	;
	v352 = v351 - v350
	goto L92
L94:
	;
	if v330 != v331 {
		v350 = v330
		v351 = v331
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v335 = l1
	v336 = v327
	goto L96
L96:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+1)))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	if v340 == int32(0) {
		v350 = v339
		v351 = v340
		goto L93
	} else {
		goto L98
	}
L97:
	;
	v350 = v339
	v351 = v340
	goto L93
L98:
	;
	v343 = int32(1)
	if v339 == v340 {
		v335 = v335 + v343
		v336 = v336 + v343
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v353 = int32(1663)
	goto L102
L101:
	;
	v353 = int32(1664)
	goto L102
L102:
	;
	v357 = base.B2i32(v352 == int32(0))
	v359 = v10
	v360 = v353
	goto L6
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4367 = m.ExcPending
	if v4367 != 0 {
		goto L4
	} else {
		goto L842
	}
L104:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4348 = m.ExcPending
	if v4348 != 0 {
		goto L4
	} else {
		goto L838
	}
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4328 = m.ExcPending
	if v4328 != 0 {
		goto L4
	} else {
		goto L833
	}
L106:
	;
	v363 = F_ReadDir(m, v361, l1)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	if v363 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v365 = int32(1)
	v377 = l0
	v378 = l1
	v379 = l2
	v380 = l3
	v381 = l4
	v382 = l5
	v383 = l6
	v384 = l7
	v385 = l8
	v386 = v50
	v388 = v363
	v403 = v58
	v404 = l2 + v50 + int32(2369)
	v406 = v357
	v408 = v361
	v410 = v359
	v412 = v360
	v419 = l5 ^ v365
	v420 = l1 + l2 + v365
	v421 = v50 + int32(2368) | int32(2)
	v422 = v46
	goto L111
L109:
	;
	v4279 = v50
	v4296 = v58
	v4301 = v361
	v4315 = v46
	goto L110
L110:
	;
	if v4296 != 0 {
		goto L828
	} else {
		goto L829
	}
L111:
	;
	v424 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v386)+2268)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v386)+2264)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v386)+2260)) = v424
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+19)))
	if v430 != int32(46) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v4279 = v386
	v4296 = v403
	v4301 = v408
	v4315 = v4266
	goto L110
L113:
	;
	v4268 = F_ReadDir(m, v408, v378)
	mBase = m.M
	v4269 = m.ExcPending
	if v4269 != 0 {
		goto L4
	} else {
		goto L826
	}
L114:
	;
	v443 = v388 + int32(19)
	v444 = int32(_a_F_sendDir_6)
	goto L121
L115:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+20)))
	if v433 == int32(0) {
		v4266 = v422
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+20)))
	if v436 != int32(46) {
		goto L114
	} else {
		goto L117
	}
L117:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+21)))
	if v439 == int32(0) {
		v4266 = v422
		goto L113
	} else {
		goto L118
	}
L118:
	;
	goto L114
L119:
	;
	if v481-v482 == int32(0) {
		v4266 = v422
		goto L113
	} else {
		goto L133
	}
L121:
	;
	goto L122
L122:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v451 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v452 = v443
	v453 = v444
	v454 = int32(9)
	v455 = v451
	goto L127
L124:
	;
	v477 = v444
	v481 = int32(0)
	goto L125
L125:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	goto L119
L126:
	;
	v477 = v472
	v481 = v474
	goto L125
L127:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453))))
	if v455 != v457 {
		v472 = v453
		v474 = v455
		goto L126
	} else {
		goto L129
	}
L128:
	;
	v472 = v466
	v474 = int32(0)
	goto L126
L129:
	;
	if v457 == int32(0) {
		v472 = v453
		v474 = v455
		goto L126
	} else {
		goto L130
	}
L130:
	;
	v462 = v454 - int32(1)
	if v462 == int32(0) {
		v472 = v453
		v474 = v455
		goto L126
	} else {
		goto L131
	}
L131:
	;
	v465 = int32(1)
	v466 = v453 + v465
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452)+1)))
	if v467 != 0 {
		v452 = v452 + v465
		v453 = v466
		v454 = v462
		v455 = v467
		goto L127
	} else {
		goto L132
	}
L132:
	;
	goto L128
L133:
	;
	v492 = int32(_a_F_sendDir_7)
	v495 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[3])))
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v496 == int32(0) {
		v515 = v495
		v516 = v496
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v516-v515 == int32(0) {
		v4266 = v422
		goto L113
	} else {
		goto L142
	}
L135:
	;
	goto L134
L136:
	;
	if v495 != v496 {
		v515 = v495
		v516 = v496
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v500 = v443
	v501 = v492
	goto L138
L138:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+1)))
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+1)))
	if v505 == int32(0) {
		v515 = v504
		v516 = v505
		goto L135
	} else {
		goto L140
	}
L139:
	;
	v515 = v504
	v516 = v505
	goto L135
L140:
	;
	v508 = int32(1)
	if v504 == v505 {
		v500 = v500 + v508
		v501 = v501 + v508
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v521 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[4]))
	if v521 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[5])))
	if v526 == int32(1) {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L145
L147:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[6])))
	if v536 != v538 {
		goto L105
	} else {
		goto L151
	}
L148:
	;
	v531 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[7]))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v531)+316))
	v534 = base.B2i32(v532 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[5])) = uint8(v534)
	v536 = v534
	goto L150
L149:
	;
	v536 = int32(0)
	goto L150
L150:
	;
	goto L147
L151:
	;
	v540 = int32(_a_F_sendDir_8)
	goto L157
L152:
	;
	if v410 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L153:
	;
	v1088 = int32(1)
	goto L152
L154:
	;
	v1074 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L4
	} else {
		goto L304
	}
L155:
	;
	if v577-v578 == int32(0) {
		goto L154
	} else {
		goto L169
	}
L157:
	;
	goto L158
L158:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v547 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v548 = v443
	v549 = v540
	v550 = int32(25)
	v551 = v547
	goto L163
L160:
	;
	v573 = v540
	v577 = int32(0)
	goto L161
L161:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573))))
	goto L155
L162:
	;
	v573 = v568
	v577 = v570
	goto L161
L163:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549))))
	if v551 != v553 {
		v568 = v549
		v570 = v551
		goto L162
	} else {
		goto L165
	}
L164:
	;
	v568 = v562
	v570 = int32(0)
	goto L162
L165:
	;
	if v553 == int32(0) {
		v568 = v549
		v570 = v551
		goto L162
	} else {
		goto L166
	}
L166:
	;
	v558 = v550 - int32(1)
	if v558 == int32(0) {
		v568 = v549
		v570 = v551
		goto L162
	} else {
		goto L167
	}
L167:
	;
	v561 = int32(1)
	v562 = v549 + v561
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+1)))
	if v563 != 0 {
		v548 = v548 + v561
		v549 = v562
		v550 = v558
		v551 = v563
		goto L163
	} else {
		goto L168
	}
L168:
	;
	goto L164
L169:
	;
	v588 = int32(_a_F_sendDir_9)
	goto L172
L170:
	;
	if v625-v626 == int32(0) {
		goto L154
	} else {
		goto L184
	}
L172:
	;
	goto L173
L173:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v595 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v596 = v443
	v597 = v588
	v598 = int32(21)
	v599 = v595
	goto L178
L175:
	;
	v621 = v588
	v625 = int32(0)
	goto L176
L176:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621))))
	goto L170
L177:
	;
	v621 = v616
	v625 = v618
	goto L176
L178:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597))))
	if v599 != v601 {
		v616 = v597
		v618 = v599
		goto L177
	} else {
		goto L180
	}
L179:
	;
	v616 = v610
	v618 = int32(0)
	goto L177
L180:
	;
	if v601 == int32(0) {
		v616 = v597
		v618 = v599
		goto L177
	} else {
		goto L181
	}
L181:
	;
	v606 = v598 - int32(1)
	if v606 == int32(0) {
		v616 = v597
		v618 = v599
		goto L177
	} else {
		goto L182
	}
L182:
	;
	v609 = int32(1)
	v610 = v597 + v609
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+1)))
	if v611 != 0 {
		v596 = v596 + v609
		v597 = v610
		v598 = v606
		v599 = v611
		goto L178
	} else {
		goto L183
	}
L183:
	;
	goto L179
L184:
	;
	v636 = int32(_a_F_sendDir_10)
	goto L187
L185:
	;
	if v673-v674 == int32(0) {
		goto L154
	} else {
		goto L199
	}
L187:
	;
	goto L188
L188:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v643 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v644 = v443
	v645 = v636
	v646 = int32(16)
	v647 = v643
	goto L193
L190:
	;
	v669 = v636
	v673 = int32(0)
	goto L191
L191:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669))))
	goto L185
L192:
	;
	v669 = v664
	v673 = v666
	goto L191
L193:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645))))
	if v647 != v649 {
		v664 = v645
		v666 = v647
		goto L192
	} else {
		goto L195
	}
L194:
	;
	v664 = v658
	v666 = int32(0)
	goto L192
L195:
	;
	if v649 == int32(0) {
		v664 = v645
		v666 = v647
		goto L192
	} else {
		goto L196
	}
L196:
	;
	v654 = v646 - int32(1)
	if v654 == int32(0) {
		v664 = v645
		v666 = v647
		goto L192
	} else {
		goto L197
	}
L197:
	;
	v657 = int32(1)
	v658 = v645 + v657
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644)+1)))
	if v659 != 0 {
		v644 = v644 + v657
		v645 = v658
		v646 = v654
		v647 = v659
		goto L193
	} else {
		goto L198
	}
L198:
	;
	goto L194
L199:
	;
	v684 = int32(_a_F_sendDir_11)
	goto L202
L200:
	;
	if v721-v722 == int32(0) {
		goto L154
	} else {
		goto L214
	}
L202:
	;
	goto L203
L203:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v691 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v692 = v443
	v693 = v684
	v694 = int32(13)
	v695 = v691
	goto L208
L205:
	;
	v717 = v684
	v721 = int32(0)
	goto L206
L206:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717))))
	goto L200
L207:
	;
	v717 = v712
	v721 = v714
	goto L206
L208:
	;
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v693))))
	if v695 != v697 {
		v712 = v693
		v714 = v695
		goto L207
	} else {
		goto L210
	}
L209:
	;
	v712 = v706
	v714 = int32(0)
	goto L207
L210:
	;
	if v697 == int32(0) {
		v712 = v693
		v714 = v695
		goto L207
	} else {
		goto L211
	}
L211:
	;
	v702 = v694 - int32(1)
	if v702 == int32(0) {
		v712 = v693
		v714 = v695
		goto L207
	} else {
		goto L212
	}
L212:
	;
	v705 = int32(1)
	v706 = v693 + v705
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692)+1)))
	if v707 != 0 {
		v692 = v692 + v705
		v693 = v706
		v694 = v702
		v695 = v707
		goto L208
	} else {
		goto L213
	}
L213:
	;
	goto L209
L214:
	;
	v732 = int32(_a_F_sendDir_12)
	goto L217
L215:
	;
	if v769-v770 == int32(0) {
		goto L154
	} else {
		goto L229
	}
L217:
	;
	goto L218
L218:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v739 != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v740 = v443
	v741 = v732
	v742 = int32(15)
	v743 = v739
	goto L223
L220:
	;
	v765 = v732
	v769 = int32(0)
	goto L221
L221:
	;
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765))))
	goto L215
L222:
	;
	v765 = v760
	v769 = v762
	goto L221
L223:
	;
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v741))))
	if v743 != v745 {
		v760 = v741
		v762 = v743
		goto L222
	} else {
		goto L225
	}
L224:
	;
	v760 = v754
	v762 = int32(0)
	goto L222
L225:
	;
	if v745 == int32(0) {
		v760 = v741
		v762 = v743
		goto L222
	} else {
		goto L226
	}
L226:
	;
	v750 = v742 - int32(1)
	if v750 == int32(0) {
		v760 = v741
		v762 = v743
		goto L222
	} else {
		goto L227
	}
L227:
	;
	v753 = int32(1)
	v754 = v741 + v753
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740)+1)))
	if v755 != 0 {
		v740 = v740 + v753
		v741 = v754
		v742 = v750
		v743 = v755
		goto L223
	} else {
		goto L228
	}
L228:
	;
	goto L224
L229:
	;
	v780 = int32(_a_F_sendDir_13)
	goto L232
L230:
	;
	if v817-v818 == int32(0) {
		goto L154
	} else {
		goto L244
	}
L232:
	;
	goto L233
L233:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v787 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v788 = v443
	v789 = v780
	v790 = int32(16)
	v791 = v787
	goto L238
L235:
	;
	v813 = v780
	v817 = int32(0)
	goto L236
L236:
	;
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813))))
	goto L230
L237:
	;
	v813 = v808
	v817 = v810
	goto L236
L238:
	;
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v789))))
	if v791 != v793 {
		v808 = v789
		v810 = v791
		goto L237
	} else {
		goto L240
	}
L239:
	;
	v808 = v802
	v810 = int32(0)
	goto L237
L240:
	;
	if v793 == int32(0) {
		v808 = v789
		v810 = v791
		goto L237
	} else {
		goto L241
	}
L241:
	;
	v798 = v790 - int32(1)
	if v798 == int32(0) {
		v808 = v789
		v810 = v791
		goto L237
	} else {
		goto L242
	}
L242:
	;
	v801 = int32(1)
	v802 = v789 + v801
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788)+1)))
	if v803 != 0 {
		v788 = v788 + v801
		v789 = v802
		v790 = v798
		v791 = v803
		goto L238
	} else {
		goto L243
	}
L243:
	;
	goto L239
L244:
	;
	v828 = int32(_a_F_sendDir_14)
	goto L247
L245:
	;
	if v865-v866 == int32(0) {
		goto L154
	} else {
		goto L259
	}
L247:
	;
	goto L248
L248:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v835 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v836 = v443
	v837 = v828
	v838 = int32(15)
	v839 = v835
	goto L253
L250:
	;
	v861 = v828
	v865 = int32(0)
	goto L251
L251:
	;
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861))))
	goto L245
L252:
	;
	v861 = v856
	v865 = v858
	goto L251
L253:
	;
	v841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837))))
	if v839 != v841 {
		v856 = v837
		v858 = v839
		goto L252
	} else {
		goto L255
	}
L254:
	;
	v856 = v850
	v858 = int32(0)
	goto L252
L255:
	;
	if v841 == int32(0) {
		v856 = v837
		v858 = v839
		goto L252
	} else {
		goto L256
	}
L256:
	;
	v846 = v838 - int32(1)
	if v846 == int32(0) {
		v856 = v837
		v858 = v839
		goto L252
	} else {
		goto L257
	}
L257:
	;
	v849 = int32(1)
	v850 = v837 + v849
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836)+1)))
	if v851 != 0 {
		v836 = v836 + v849
		v837 = v850
		v838 = v846
		v839 = v851
		goto L253
	} else {
		goto L258
	}
L258:
	;
	goto L254
L259:
	;
	v876 = int32(_a_F_sendDir_15)
	goto L262
L260:
	;
	if v913-v914 == int32(0) {
		goto L154
	} else {
		goto L274
	}
L262:
	;
	goto L263
L263:
	;
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v883 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v884 = v443
	v885 = v876
	v886 = int32(16)
	v887 = v883
	goto L268
L265:
	;
	v909 = v876
	v913 = int32(0)
	goto L266
L266:
	;
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909))))
	goto L260
L267:
	;
	v909 = v904
	v913 = v906
	goto L266
L268:
	;
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885))))
	if v887 != v889 {
		v904 = v885
		v906 = v887
		goto L267
	} else {
		goto L270
	}
L269:
	;
	v904 = v898
	v906 = int32(0)
	goto L267
L270:
	;
	if v889 == int32(0) {
		v904 = v885
		v906 = v887
		goto L267
	} else {
		goto L271
	}
L271:
	;
	v894 = v886 - int32(1)
	if v894 == int32(0) {
		v904 = v885
		v906 = v887
		goto L267
	} else {
		goto L272
	}
L272:
	;
	v897 = int32(1)
	v898 = v885 + v897
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v884)+1)))
	if v899 != 0 {
		v884 = v884 + v897
		v885 = v898
		v886 = v894
		v887 = v899
		goto L268
	} else {
		goto L273
	}
L273:
	;
	goto L269
L274:
	;
	v924 = int32(0)
	if v406 == v924 {
		v1088 = v924
		goto L152
	} else {
		goto L275
	}
L275:
	;
	v928 = v386 + int32(2268)
	v930 = v386 + int32(2264)
	v932 = v386 + int32(2260)
	v933 = int32(0)
	v938 = m.G0
	v940 = v938 - int32(16)
	m.G0 = v940
	*(*int32)(unsafe.Add(mBase, uint32(v928))) = v933
	*(*int32)(unsafe.Add(mBase, uint32(v930))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v932))) = v933
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if base.Ui32((v948-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v1028 = v933
		goto L277
	} else {
		goto L278
	}
L276:
	;
	if v1028 == int32(0) {
		v1088 = v1028
		goto L152
	} else {
		goto L295
	}
L277:
	;
	m.G0 = v940 + int32(16)
	goto L276
L278:
	;
	v955 = int32(_a_F_sendDir_16)
	*(*int32)(unsafe.Add(mBase, _c_F_sendDir[8])) = int32(0)
	v961 = F_strtoul(m, v443, v940+int32(8), int32(10))
	mBase = m.M
	v963 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[8]))
	if v963 != 0 {
		v1028 = v933
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v940)+8))
	if v443 == v964 {
		v1028 = v933
		goto L277
	} else {
		goto L280
	}
L280:
	;
	if v961 == int32(0) {
		v1028 = v933
		goto L277
	} else {
		goto L281
	}
L281:
	;
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964))))
	if v968 != int32(95) {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	if v984&int32(255) == int32(46) {
		goto L287
	} else {
		goto L288
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v940)+12)) = int32(0)
	v984 = v968
	v985 = v964
	goto L282
L284:
	;
	goto L285
L285:
	;
	v977 = F_forkname_chars(m, v964+int32(1), v940+int32(12))
	mBase = m.M
	if v977 <= int32(0) {
		v1028 = v933
		goto L277
	} else {
		goto L286
	}
L286:
	;
	v982 = v977 + v964 + int32(1)
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v982))))
	v984 = v983
	v985 = v982
	goto L282
L287:
	;
	v991 = v985 + int32(1)
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v991))))
	if base.Ui32((v992-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v1028 = v933
		goto L277
	} else {
		goto L290
	}
L288:
	;
	v1015 = v933
	v1016 = v984
	goto L289
L289:
	;
	if v1016&int32(255) != 0 {
		v1028 = v933
		goto L277
	} else {
		goto L294
	}
L290:
	;
	v999 = int32(_a_F_sendDir_16)
	*(*int32)(unsafe.Add(mBase, _c_F_sendDir[8])) = int32(0)
	v1005 = F_strtoul(m, v991, v940+int32(8), int32(10))
	mBase = m.M
	v1007 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[8]))
	if v1007 != 0 {
		v1028 = v933
		goto L277
	} else {
		goto L291
	}
L291:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v940)+8))
	if v991 == v1008 {
		v1028 = v933
		goto L277
	} else {
		goto L292
	}
L292:
	;
	if v1005 == int32(0) {
		v1028 = v933
		goto L277
	} else {
		goto L293
	}
L293:
	;
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008))))
	v1015 = v1005
	v1016 = v1012
	goto L289
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v928))) = v961
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v940)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v930))) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v932))) = v1015
	v1028 = int32(1)
	goto L277
L295:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2264))
	if v1034 == int32(3) {
		v1088 = v1028
		goto L152
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+176)) = v378
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2268))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+180)) = v1038
	v1046 = F_pg_snprintf(m, v386+int32(192), int32(1024), int32(_a_F_sendDir_17), v386+int32(176))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L4
	} else {
		goto L297
	}
L297:
	;
	v1054 = F___fstatat(m, int32(-100), v386+int32(192), v386+int32(2272), int32(256))
	mBase = m.M
	goto L298
L298:
	;
	if v1054 != 0 {
		goto L153
	} else {
		goto L299
	}
L299:
	;
	v1057 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L4
	} else {
		goto L300
	}
L300:
	;
	if v1057 == int32(0) {
		v4266 = v422
		goto L113
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+16)) = v443
	F_errmsg_internal(m, int32(_a_F_sendDir_18), v386+int32(16))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L4
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1333), int32(_a_F_sendDir_20))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L4
	} else {
		goto L303
	}
L303:
	;
	v4266 = v422
	goto L113
L304:
	;
	if v1074 == int32(0) {
		v4266 = v422
		goto L113
	} else {
		goto L305
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386))) = v443
	F_errmsg_internal(m, int32(_a_F_sendDir_21), v386)
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L4
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1298), int32(_a_F_sendDir_20))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L4
	} else {
		goto L307
	}
L307:
	;
	v4266 = v422
	goto L113
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+148)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v386)+144)) = v378
	v1499 = F_pg_snprintf(m, v386+int32(2368), int32(2048), int32(_a_F_sendDir_22), v386+int32(144))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L4
	} else {
		goto L351
	}
L309:
	;
	v1091 = int32(0)
	v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v1092 != int32(116) {
		v1397 = v1091
		goto L310
	} else {
		goto L311
	}
L310:
	;
	if v1397 == int32(0) {
		goto L308
	} else {
		goto L346
	}
L311:
	;
	v1108 = int32(1)
	goto L312
L312:
	;
	v1144 = v1108 + int32(1)
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108+v443))))
	if base.Ui32((v1146-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v1108 = v1144
		goto L312
	} else {
		goto L314
	}
L313:
	;
	if v1108 == int32(1) {
		v1397 = v1091
		goto L310
	} else {
		goto L315
	}
L314:
	;
	goto L313
L315:
	;
	if v1146&int32(255) != int32(95) {
		v1397 = v1091
		goto L310
	} else {
		goto L316
	}
L316:
	;
	v1173 = v1144
	goto L317
L317:
	;
	v1207 = v1173 + int32(1)
	v1208 = v1173 + v443
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1208))))
	if base.Ui32((v1209-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v1173 = v1207
		goto L317
	} else {
		goto L319
	}
L318:
	;
	if v1173 == v1144 {
		v1397 = v1091
		goto L310
	} else {
		goto L320
	}
L319:
	;
	goto L318
L320:
	;
	if v1209 == int32(95) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1220 = v1208 + int32(1)
	v1224 = int32(3)
	v1227 = F_strncmp(m, int32(_a_F_sendDir_23), v1220, v1224)
	mBase = m.M
	if v1227 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L322:
	;
	v1263 = v1173
	v1265 = v1209
	goto L323
L323:
	;
	if v1265 == int32(46) {
		goto L339
	} else {
		goto L340
	}
L324:
	;
	if v1256 <= int32(0) {
		v1397 = v1091
		goto L310
	} else {
		goto L338
	}
L325:
	;
	goto L324
L327:
	;
	v1256 = v1249
	goto L325
L328:
	;
	v1249 = v1224
	goto L327
L329:
	;
	goto L330
L330:
	;
	v1231 = int32(2)
	v1235 = F_strncmp(m, int32(_a_F_sendDir_24), v1220, v1231)
	mBase = m.M
	if v1235 == int32(0) {
		v1249 = v1231
		goto L327
	} else {
		goto L331
	}
L331:
	;
	v1238 = int32(4)
	v1241 = F_strncmp(m, int32(_a_F_sendDir_25), v1220, v1238)
	mBase = m.M
	if v1241 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	goto L335
L333:
	;
	goto L334
L334:
	;
	v1256 = int32(0)
	goto L325
L335:
	;
	v1256 = v1238
	goto L325
L338:
	;
	v1260 = v1256 + v1207
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443+v1260))))
	v1263 = v1260
	v1265 = v1262
	goto L323
L339:
	;
	v1282 = int32(1)
	goto L342
L340:
	;
	v1346 = v1265
	goto L341
L341:
	;
	v1397 = base.B2i32(v1346 == int32(0))
	goto L310
L342:
	;
	v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1282+(v1263+v443)))))
	if base.Ui32((v1320-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v1282 = v1282 + int32(1)
		goto L342
	} else {
		goto L344
	}
L343:
	;
	if v1282 < int32(2) {
		v1397 = v1091
		goto L310
	} else {
		goto L345
	}
L344:
	;
	goto L343
L345:
	;
	v1346 = v1320
	goto L341
L346:
	;
	v1429 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L4
	} else {
		goto L347
	}
L347:
	;
	if v1429 == int32(0) {
		v4266 = v422
		goto L113
	} else {
		goto L348
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+160)) = v443
	F_errmsg_internal(m, int32(_a_F_sendDir_26), v386+int32(160))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L4
	} else {
		goto L349
	}
L349:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1344), int32(_a_F_sendDir_20))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L4
	} else {
		goto L350
	}
L350:
	;
	v4266 = v422
	goto L113
L351:
	;
	v1502 = v386 + int32(2368)
	v1503 = int32(_a_F_sendDir_27)
	v1504 = int32(20)
	goto L355
L352:
	;
	if v1566 == int32(0) {
		v4266 = v422
		goto L113
	} else {
		goto L370
	}
L353:
	;
	v1566 = int32(0)
	goto L352
L354:
	;
	v1540 = v1535
	v1541 = v1536
	v1542 = v1537
	goto L364
L355:
	;
	if (v1502|v1503)&int32(3) != 0 {
		v1535 = v1502
		v1536 = v1503
		v1537 = v1504
		goto L354
	} else {
		goto L358
	}
L357:
	;
	if v1525 == int32(0) {
		goto L353
	} else {
		goto L363
	}
L358:
	;
	v1512 = v1502
	v1513 = v1503
	v1514 = v1504
	goto L359
L359:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v1512)))
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1513)))
	if v1517 != v1518 {
		v1535 = v1512
		v1536 = v1513
		v1537 = v1514
		goto L354
	} else {
		goto L361
	}
L360:
	;
	goto L357
L361:
	;
	v1520 = int32(4)
	v1521 = v1513 + v1520
	v1523 = v1512 + v1520
	v1525 = v1514 - v1520
	if base.Ui32(int32(3)) < base.Ui32(v1525) {
		v1512 = v1523
		v1513 = v1521
		v1514 = v1525
		goto L359
	} else {
		goto L362
	}
L362:
	;
	goto L360
L363:
	;
	v1535 = v1523
	v1536 = v1521
	v1537 = v1525
	goto L354
L364:
	;
	v1545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1540))))
	v1546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1541))))
	if v1545 == v1546 {
		goto L366
	} else {
		goto L367
	}
L365:
	;
	v1566 = v1545 - v1546
	goto L352
L366:
	;
	v1548 = int32(1)
	v1553 = v1542 - v1548
	if v1553 != 0 {
		v1540 = v1540 + v1548
		v1541 = v1541 + v1548
		v1542 = v1553
		goto L364
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	goto L365
L369:
	;
	goto L353
L370:
	;
	v1575 = F___fstatat(m, int32(-100), v386+int32(2368), v386+int32(2272), int32(256))
	mBase = m.M
	goto L373
L371:
	;
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2276))
	v1926 = v1924 & int32(_a_F_sendDir_28)
	v1927 = int32(_a_F_sendDir_29)
	v1930 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[9])))
	v1931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	if v1931 == int32(0) {
		v1950 = v1930
		v1951 = v1931
		goto L482
	} else {
		goto L483
	}
L372:
	;
	v1894 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L4
	} else {
		goto L470
	}
L373:
	;
	if v1575 == int32(0) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1578 = int32(_a_F_sendDir_30)
	v1581 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[10])))
	v1582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v1582 == int32(0) {
		v1601 = v1581
		v1602 = v1582
		goto L378
	} else {
		goto L379
	}
L375:
	;
	goto L376
L376:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[8]))
	if v1870 == int32(44) {
		v4266 = v422
		goto L113
	} else {
		goto L465
	}
L377:
	;
	if v1602-v1601 == int32(0) {
		goto L372
	} else {
		goto L385
	}
L378:
	;
	goto L377
L379:
	;
	if v1581 != v1582 {
		v1601 = v1581
		v1602 = v1582
		goto L378
	} else {
		goto L380
	}
L380:
	;
	v1586 = v443
	v1587 = v1578
	goto L381
L381:
	;
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1587)+1)))
	v1591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1586)+1)))
	if v1591 == int32(0) {
		v1601 = v1590
		v1602 = v1591
		goto L378
	} else {
		goto L383
	}
L382:
	;
	v1601 = v1590
	v1602 = v1591
	goto L378
L383:
	;
	v1594 = int32(1)
	if v1590 == v1591 {
		v1586 = v1586 + v1594
		v1587 = v1587 + v1594
		goto L381
	} else {
		goto L384
	}
L384:
	;
	goto L382
L385:
	;
	v1606 = int32(_a_F_sendDir_31)
	v1609 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[11])))
	v1610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v1610 == int32(0) {
		v1629 = v1609
		v1630 = v1610
		goto L387
	} else {
		goto L388
	}
L386:
	;
	if v1630-v1629 == int32(0) {
		goto L372
	} else {
		goto L394
	}
L387:
	;
	goto L386
L388:
	;
	if v1609 != v1610 {
		v1629 = v1609
		v1630 = v1610
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v1614 = v443
	v1615 = v1606
	goto L390
L390:
	;
	v1618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1615)+1)))
	v1619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1614)+1)))
	if v1619 == int32(0) {
		v1629 = v1618
		v1630 = v1619
		goto L387
	} else {
		goto L392
	}
L391:
	;
	v1629 = v1618
	v1630 = v1619
	goto L387
L392:
	;
	v1622 = int32(1)
	if v1618 == v1619 {
		v1614 = v1614 + v1622
		v1615 = v1615 + v1622
		goto L390
	} else {
		goto L393
	}
L393:
	;
	goto L391
L394:
	;
	v1634 = int32(_a_F_sendDir_32)
	v1637 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[12])))
	v1638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v1638 == int32(0) {
		v1657 = v1637
		v1658 = v1638
		goto L396
	} else {
		goto L397
	}
L395:
	;
	if v1658-v1657 == int32(0) {
		goto L372
	} else {
		goto L403
	}
L396:
	;
	goto L395
L397:
	;
	if v1637 != v1638 {
		v1657 = v1637
		v1658 = v1638
		goto L396
	} else {
		goto L398
	}
L398:
	;
	v1642 = v443
	v1643 = v1634
	goto L399
L399:
	;
	v1646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1643)+1)))
	v1647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1642)+1)))
	if v1647 == int32(0) {
		v1657 = v1646
		v1658 = v1647
		goto L396
	} else {
		goto L401
	}
L400:
	;
	v1657 = v1646
	v1658 = v1647
	goto L396
L401:
	;
	v1650 = int32(1)
	if v1646 == v1647 {
		v1642 = v1642 + v1650
		v1643 = v1643 + v1650
		goto L399
	} else {
		goto L402
	}
L402:
	;
	goto L400
L403:
	;
	v1662 = int32(_a_F_sendDir_33)
	v1665 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[13])))
	v1666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v1666 == int32(0) {
		v1685 = v1665
		v1686 = v1666
		goto L405
	} else {
		goto L406
	}
L404:
	;
	if v1686-v1685 == int32(0) {
		goto L372
	} else {
		goto L412
	}
L405:
	;
	goto L404
L406:
	;
	if v1665 != v1666 {
		v1685 = v1665
		v1686 = v1666
		goto L405
	} else {
		goto L407
	}
L407:
	;
	v1670 = v443
	v1671 = v1662
	goto L408
L408:
	;
	v1674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1671)+1)))
	v1675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1670)+1)))
	if v1675 == int32(0) {
		v1685 = v1674
		v1686 = v1675
		goto L405
	} else {
		goto L410
	}
L409:
	;
	v1685 = v1674
	v1686 = v1675
	goto L405
L410:
	;
	v1678 = int32(1)
	if v1674 == v1675 {
		v1670 = v1670 + v1678
		v1671 = v1671 + v1678
		goto L408
	} else {
		goto L411
	}
L411:
	;
	goto L409
L412:
	;
	v1690 = int32(_a_F_sendDir_34)
	v1693 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[14])))
	v1694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v1694 == int32(0) {
		v1713 = v1693
		v1714 = v1694
		goto L414
	} else {
		goto L415
	}
L413:
	;
	if v1714-v1713 == int32(0) {
		goto L372
	} else {
		goto L421
	}
L414:
	;
	goto L413
L415:
	;
	if v1693 != v1694 {
		v1713 = v1693
		v1714 = v1694
		goto L414
	} else {
		goto L416
	}
L416:
	;
	v1698 = v443
	v1699 = v1690
	goto L417
L417:
	;
	v1702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1699)+1)))
	v1703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1698)+1)))
	if v1703 == int32(0) {
		v1713 = v1702
		v1714 = v1703
		goto L414
	} else {
		goto L419
	}
L418:
	;
	v1713 = v1702
	v1714 = v1703
	goto L414
L419:
	;
	v1706 = int32(1)
	if v1702 == v1703 {
		v1698 = v1698 + v1706
		v1699 = v1699 + v1706
		goto L417
	} else {
		goto L420
	}
L420:
	;
	goto L418
L421:
	;
	v1718 = int32(_a_F_sendDir_35)
	v1721 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[15])))
	v1722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v1722 == int32(0) {
		v1741 = v1721
		v1742 = v1722
		goto L423
	} else {
		goto L424
	}
L422:
	;
	if v1742-v1741 == int32(0) {
		goto L372
	} else {
		goto L430
	}
L423:
	;
	goto L422
L424:
	;
	if v1721 != v1722 {
		v1741 = v1721
		v1742 = v1722
		goto L423
	} else {
		goto L425
	}
L425:
	;
	v1726 = v443
	v1727 = v1718
	goto L426
L426:
	;
	v1730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1727)+1)))
	v1731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726)+1)))
	if v1731 == int32(0) {
		v1741 = v1730
		v1742 = v1731
		goto L423
	} else {
		goto L428
	}
L427:
	;
	v1741 = v1730
	v1742 = v1731
	goto L423
L428:
	;
	v1734 = int32(1)
	if v1730 == v1731 {
		v1726 = v1726 + v1734
		v1727 = v1727 + v1734
		goto L426
	} else {
		goto L429
	}
L429:
	;
	goto L427
L430:
	;
	v1746 = int32(_a_F_sendDir_36)
	v1749 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendDir[16])))
	v1750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v1750 == int32(0) {
		v1769 = v1749
		v1770 = v1750
		goto L432
	} else {
		goto L433
	}
L431:
	;
	if v1770-v1769 == int32(0) {
		goto L372
	} else {
		goto L439
	}
L432:
	;
	goto L431
L433:
	;
	if v1749 != v1750 {
		v1769 = v1749
		v1770 = v1750
		goto L432
	} else {
		goto L434
	}
L434:
	;
	v1754 = v443
	v1755 = v1746
	goto L435
L435:
	;
	v1758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1755)+1)))
	v1759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1754)+1)))
	if v1759 == int32(0) {
		v1769 = v1758
		v1770 = v1759
		goto L432
	} else {
		goto L437
	}
L436:
	;
	v1769 = v1758
	v1770 = v1759
	goto L432
L437:
	;
	v1762 = int32(1)
	if v1758 == v1759 {
		v1754 = v1754 + v1762
		v1755 = v1755 + v1762
		goto L435
	} else {
		goto L438
	}
L438:
	;
	goto L436
L439:
	;
	v1775 = v386 + int32(2368)
	v1776 = int32(_a_F_sendDir_37)
	v1777 = int32(9)
	goto L443
L440:
	;
	if v1839 != 0 {
		goto L371
	} else {
		goto L458
	}
L441:
	;
	v1839 = int32(0)
	goto L440
L442:
	;
	v1813 = v1808
	v1814 = v1809
	v1815 = v1810
	goto L452
L443:
	;
	if (v1775|v1776)&int32(3) != 0 {
		v1808 = v1775
		v1809 = v1776
		v1810 = v1777
		goto L442
	} else {
		goto L446
	}
L445:
	;
	if v1798 == int32(0) {
		goto L441
	} else {
		goto L451
	}
L446:
	;
	v1785 = v1775
	v1786 = v1776
	v1787 = v1777
	goto L447
L447:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v1785)))
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1786)))
	if v1790 != v1791 {
		v1808 = v1785
		v1809 = v1786
		v1810 = v1787
		goto L442
	} else {
		goto L449
	}
L448:
	;
	goto L445
L449:
	;
	v1793 = int32(4)
	v1794 = v1786 + v1793
	v1796 = v1785 + v1793
	v1798 = v1787 - v1793
	if base.Ui32(int32(3)) < base.Ui32(v1798) {
		v1785 = v1796
		v1786 = v1794
		v1787 = v1798
		goto L447
	} else {
		goto L450
	}
L450:
	;
	goto L448
L451:
	;
	v1808 = v1796
	v1809 = v1794
	v1810 = v1798
	goto L442
L452:
	;
	v1818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1813))))
	v1819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1814))))
	if v1818 == v1819 {
		goto L454
	} else {
		goto L455
	}
L453:
	;
	v1839 = v1818 - v1819
	goto L440
L454:
	;
	v1821 = int32(1)
	v1826 = v1815 - v1821
	if v1826 != 0 {
		v1813 = v1813 + v1821
		v1814 = v1814 + v1821
		v1815 = v1826
		goto L452
	} else {
		goto L457
	}
L455:
	;
	goto L456
L456:
	;
	goto L453
L457:
	;
	goto L441
L458:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2276))
	if v1840&int32(_a_F_sendDir_28) == int32(_a_F_sendDir_38) {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v1846 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+2276)) = v1846 | int32(_a_F_sendDir_39)
	goto L461
L460:
	;
	goto L461
L461:
	;
	F__tarWriteHeader(m, v377, v404, int32(0), v386+int32(2272), v380)
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L4
	} else {
		goto L462
	}
L462:
	;
	F__tarWriteHeader(m, v377, int32(_a_F_sendDir_40), int32(0), v386+int32(2272), v380)
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L4
	} else {
		goto L463
	}
L463:
	;
	F__tarWriteHeader(m, v377, int32(_a_F_sendDir_41), int32(0), v386+int32(2272), v380)
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L4
	} else {
		goto L464
	}
L464:
	;
	v4266 = v422 + int64(1536)
	goto L113
L465:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		goto L4
	} else {
		goto L466
	}
L466:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L4
	} else {
		goto L467
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+128)) = v386 + int32(2368)
	F_errmsg(m, int32(_a_F_sendDir_42), v386+int32(128))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L4
	} else {
		goto L468
	}
L468:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1361), int32(_a_F_sendDir_20))
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L4
	} else {
		goto L469
	}
L469:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L470:
	;
	if v1894 != 0 {
		goto L471
	} else {
		goto L472
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+32)) = v443
	F_errmsg_internal(m, int32(_a_F_sendDir_43), v386+int32(32))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L4
	} else {
		goto L474
	}
L472:
	;
	goto L473
L473:
	;
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2276))
	if v1907&int32(_a_F_sendDir_28) == int32(_a_F_sendDir_38) {
		goto L476
	} else {
		goto L477
	}
L474:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1373), int32(_a_F_sendDir_20))
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L4
	} else {
		goto L475
	}
L475:
	;
	goto L473
L476:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, _c_F_sendDir[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+2276)) = v1913 | int32(_a_F_sendDir_39)
	goto L478
L477:
	;
	goto L478
L478:
	;
	F__tarWriteHeader(m, v377, v404, int32(0), v386+int32(2272), v380)
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L4
	} else {
		goto L479
	}
L479:
	;
	v4266 = v422 + int64(512)
	goto L113
L480:
	;
	if v1926 != int32(_a_F_sendDir_44) {
		goto L495
	} else {
		goto L496
	}
L481:
	;
	if v1951-v1950 != 0 {
		goto L480
	} else {
		goto L489
	}
L482:
	;
	goto L481
L483:
	;
	if v1930 != v1931 {
		v1950 = v1930
		v1951 = v1931
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v1935 = v378
	v1936 = v1927
	goto L485
L485:
	;
	v1939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1936)+1)))
	v1940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1935)+1)))
	if v1940 == int32(0) {
		v1950 = v1939
		v1951 = v1940
		goto L482
	} else {
		goto L487
	}
L486:
	;
	v1950 = v1939
	v1951 = v1940
	goto L482
L487:
	;
	v1943 = int32(1)
	if v1939 == v1940 {
		v1935 = v1935 + v1943
		v1936 = v1936 + v1943
		goto L485
	} else {
		goto L488
	}
L488:
	;
	goto L486
L489:
	;
	if v1926 != int32(_a_F_sendDir_38) {
		goto L480
	} else {
		goto L490
	}
L490:
	;
	v1960 = F_readlink(m, v386+int32(2368), v386+int32(192), int32(1024))
	mBase = m.M
	if v1960 < int32(0) {
		goto L104
	} else {
		goto L491
	}
L491:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v1960) {
		goto L103
	} else {
		goto L492
	}
L492:
	;
	v1966 = v386 + int32(192)
	v1968 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1966+v1960))) = uint8(v1968)
	F__tarWriteHeader(m, v377, v404, v1966, v386+int32(2272), v380)
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L4
	} else {
		goto L493
	}
L493:
	;
	v4266 = v422 + int64(512)
	goto L113
L494:
	;
	v4204 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4205 = m.ExcPending
	if v4205 != 0 {
		goto L4
	} else {
		goto L822
	}
L495:
	;
	if v1926 != int32(_a_F_sendDir_39) {
		goto L494
	} else {
		goto L498
	}
L496:
	;
	goto L497
L497:
	;
	v2205 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v386)+2256)) = v2205
	*(*int32)(unsafe.Add(mBase, uint32(v386)+2252)) = v2205
	if base.B2i32(v385 == v2205)|(v1088^int32(1)) == v2205 {
		goto L541
	} else {
		goto L542
	}
L498:
	;
	v1982 = int32(0)
	F__tarWriteHeader(m, v377, v404, v1982, v386+int32(2272), v380)
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L4
	} else {
		goto L499
	}
L499:
	;
	if v381 == int32(0) {
		v2093 = v1982
		goto L500
	} else {
		goto L501
	}
L500:
	;
	v2130 = v422 + int64(512)
	v2132 = v386 + int32(2368)
	v2133 = int32(_a_F_sendDir_29)
	v2134 = int32(12)
	goto L523
L501:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	if v1990 <= int32(0) {
		v2093 = v1982
		goto L500
	} else {
		goto L502
	}
L502:
	;
	v1993 = int32(0)
	if v1993 < v1990 {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v1996 = v1990
	goto L505
L504:
	;
	v1996 = v1993
	goto L505
L505:
	;
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v381)+12))
	v2009 = v1982
	goto L506
L506:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v1997+v2009<<(uint(int32(2))%32))))
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v2048)+8))
	if v2049 == int32(0) {
		goto L508
	} else {
		goto L509
	}
L507:
	;
	v2093 = int32(0)
	goto L500
L508:
	;
	v2079 = v2009 + int32(1)
	if v2079 != v1996 {
		v2009 = v2079
		goto L506
	} else {
		goto L519
	}
L509:
	;
	v2054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421))))
	v2055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2049))))
	if v2055 == int32(0) {
		v2074 = v2054
		v2075 = v2055
		goto L511
	} else {
		goto L512
	}
L510:
	;
	if v2075-v2074 != 0 {
		goto L508
	} else {
		goto L518
	}
L511:
	;
	goto L510
L512:
	;
	if v2054 != v2055 {
		v2074 = v2054
		v2075 = v2055
		goto L511
	} else {
		goto L513
	}
L513:
	;
	v2059 = v2049
	v2060 = v421
	goto L514
L514:
	;
	v2063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2060)+1)))
	v2064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2059)+1)))
	if v2064 == int32(0) {
		v2074 = v2063
		v2075 = v2064
		goto L511
	} else {
		goto L516
	}
L515:
	;
	v2074 = v2063
	v2075 = v2064
	goto L511
L516:
	;
	v2067 = int32(1)
	if v2063 == v2064 {
		v2059 = v2059 + v2067
		v2060 = v2060 + v2067
		goto L514
	} else {
		goto L517
	}
L517:
	;
	goto L515
L518:
	;
	v2093 = int32(1)
	goto L500
L519:
	;
	goto L507
L520:
	;
	if v2093 != 0 {
		v4266 = v2130
		goto L113
	} else {
		goto L538
	}
L521:
	;
	v2196 = int32(0)
	goto L520
L522:
	;
	v2170 = v2165
	v2171 = v2166
	v2172 = v2167
	goto L532
L523:
	;
	if (v2132|v2133)&int32(3) != 0 {
		v2165 = v2132
		v2166 = v2133
		v2167 = v2134
		goto L522
	} else {
		goto L526
	}
L525:
	;
	if v2155 == int32(0) {
		goto L521
	} else {
		goto L531
	}
L526:
	;
	v2142 = v2132
	v2143 = v2133
	v2144 = v2134
	goto L527
L527:
	;
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v2142)))
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v2143)))
	if v2147 != v2148 {
		v2165 = v2142
		v2166 = v2143
		v2167 = v2144
		goto L522
	} else {
		goto L529
	}
L528:
	;
	goto L525
L529:
	;
	v2150 = int32(4)
	v2151 = v2143 + v2150
	v2153 = v2142 + v2150
	v2155 = v2144 - v2150
	if base.Ui32(int32(3)) < base.Ui32(v2155) {
		v2142 = v2153
		v2143 = v2151
		v2144 = v2155
		goto L527
	} else {
		goto L530
	}
L530:
	;
	goto L528
L531:
	;
	v2165 = v2153
	v2166 = v2151
	v2167 = v2155
	goto L522
L532:
	;
	v2175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2170))))
	v2176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2171))))
	if v2175 == v2176 {
		goto L534
	} else {
		goto L535
	}
L533:
	;
	v2196 = v2175 - v2176
	goto L520
L534:
	;
	v2178 = int32(1)
	v2183 = v2172 - v2178
	if v2183 != 0 {
		v2170 = v2170 + v2178
		v2171 = v2171 + v2178
		v2172 = v2183
		goto L532
	} else {
		goto L537
	}
L535:
	;
	goto L536
L536:
	;
	goto L533
L537:
	;
	goto L521
L538:
	;
	if base.B2i32(v2196 == int32(0))&v419 != 0 {
		v4266 = v2130
		goto L113
	} else {
		goto L539
	}
L539:
	;
	v2202 = F_sendDir(m, v377, v386+int32(2368), v379, v380, v381, v382, v383, v384, v385)
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L4
	} else {
		goto L540
	}
L540:
	;
	v4266 = v2202 + v2130
	goto L113
L541:
	;
	if v384 != 0 {
		goto L545
	} else {
		goto L546
	}
L542:
	;
	v4144 = v2205
	v4158 = v404
	goto L543
L543:
	;
	if v380 == int32(0) {
		goto L817
	} else {
		goto L818
	}
L544:
	;
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2268))
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2264))
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2260))
	v2233 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2296))
	v2235 = v386 + int32(2256)
	v2237 = v386 + int32(2252)
	v2238 = int32(0)
	v2239 = m.G0
	v2241 = v2239 - int32(128)
	m.G0 = v2241
	if v2233&int32(_a_F_sendDir_45) != 0 {
		v4042 = v2238
		goto L552
	} else {
		goto L553
	}
L545:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+120)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v386)+116)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v386)+112)) = int32(_a_F_sendDir_46)
	v2224 = F_psprintf(m, int32(_a_F_sendDir_47), v386+int32(112))
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L4
	} else {
		goto L548
	}
L546:
	;
	goto L547
L547:
	;
	v2226 = F_pstrdup(m, v404)
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L4
	} else {
		goto L549
	}
L548:
	;
	v2228 = v2224
	v2229 = v384
	goto L544
L549:
	;
	v2228 = v2226
	v2229 = v412
	goto L544
L550:
	;
	if v4042 == int32(1) {
		goto L809
	} else {
		goto L810
	}
L551:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4072 = m.ExcPending
	if v4072 != 0 {
		goto L4
	} else {
		goto L805
	}
L552:
	;
	m.G0 = v2241 + int32(128)
	goto L550
L553:
	;
	if v2231 == int32(1) {
		v4042 = v2238
		goto L552
	} else {
		goto L554
	}
L554:
	;
	if base.Ui32(int32(1073741824)) < base.Ui32(v2233) {
		v4042 = v2238
		goto L552
	} else {
		goto L555
	}
L555:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v385)+24))
	if v2228&int32(3) == int32(0) {
		v2273 = v2228
		goto L558
	} else {
		goto L559
	}
L556:
	;
	v2312 = v2306 - int32(1636608432)
	if v2228&int32(3) != 0 {
		goto L577
	} else {
		goto L578
	}
L557:
	;
	v2306 = v2298 - v2228
	goto L556
L558:
	;
	v2277 = v2273
	goto L567
L559:
	;
	v2257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2228))))
	if v2257 == int32(0) {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v2306 = int32(0)
	goto L556
L561:
	;
	goto L562
L562:
	;
	v2262 = v2228
	goto L563
L563:
	;
	v2266 = v2262 + int32(1)
	if v2266&int32(3) == int32(0) {
		v2273 = v2266
		goto L558
	} else {
		goto L565
	}
L564:
	;
	v2298 = v2266
	goto L557
L565:
	;
	v2271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2266))))
	if v2271 != 0 {
		v2262 = v2266
		goto L563
	} else {
		goto L566
	}
L566:
	;
	goto L564
L567:
	;
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v2277)))
	v2286 = int32(-2139062144)
	if (int32(16843008)-v2283|v2283)&v2286 == v2286 {
		v2277 = v2277 + int32(4)
		goto L567
	} else {
		goto L569
	}
L568:
	;
	v2292 = v2277
	goto L570
L569:
	;
	goto L568
L570:
	;
	v2296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2292))))
	if v2296 != 0 {
		v2292 = v2292 + int32(1)
		goto L570
	} else {
		goto L572
	}
L571:
	;
	v2298 = v2292
	goto L557
L572:
	;
	goto L571
L573:
	;
	v2571 = *(*int32)(unsafe.Add(mBase, uint32(v2249)+20))
	v2572 = *(*int32)(unsafe.Add(mBase, uint32(v2249)+12))
	v2573 = (v2566 ^ v2558 - base.I32_rotl(v2566, int32(24))) & v2572
	v2577 = *(*int32)(unsafe.Add(mBase, uint32(v2571+v2573<<(uint(int32(4))%32))))
	if v2577 != 0 {
		goto L614
	} else {
		goto L615
	}
L574:
	;
	v2544 = int32(14)
	v2546 = v2540 ^ v2541 - base.I32_rotl(v2540, v2544)
	v2550 = v2546 ^ v2539 - base.I32_rotl(v2546, int32(11))
	v2554 = v2550 ^ v2540 - base.I32_rotl(v2550, int32(25))
	v2558 = v2554 ^ v2546 - base.I32_rotl(v2554, int32(16))
	v2562 = v2558 ^ v2550 - base.I32_rotl(v2558, int32(4))
	v2566 = v2562 ^ v2554 - base.I32_rotl(v2562, v2544)
	goto L573
L575:
	;
	switch v2470 - int32(1) {
	case 0:
		v2532 = v2471
		v2533 = v2472
		v2534 = v2473
		goto L602
	case 1:
		v2525 = v2471
		v2526 = v2472
		v2527 = v2473
		goto L603
	case 2:
		v2518 = v2471
		v2519 = v2472
		v2520 = v2473
		goto L604
	case 3:
		v2512 = v2472
		v2513 = v2473
		goto L605
	case 4:
		v2508 = v2472
		v2509 = v2473
		goto L606
	case 5:
		v2502 = v2472
		v2503 = v2473
		goto L607
	case 6:
		v2496 = v2472
		v2497 = v2473
		goto L608
	case 7:
		v2491 = v2473
		goto L609
	case 8:
		v2486 = v2473
		goto L610
	case 9:
		v2481 = v2473
		goto L611
	case 10:
		goto L612
	default:
		v2539 = v2471
		v2540 = v2472
		v2541 = v2473
		goto L574
	}
L576:
	;
	v2421 = v2228
	v2422 = v2306
	v2423 = v2312
	v2424 = v2312
	v2425 = v2312
	goto L599
L577:
	;
	if base.Ui32(int32(11)) < base.Ui32(v2306) {
		goto L576
	} else {
		goto L580
	}
L578:
	;
	goto L579
L579:
	;
	if base.Ui32(v2306) < base.Ui32(int32(12)) {
		goto L582
	} else {
		goto L583
	}
L580:
	;
	v2469 = v2228
	v2470 = v2306
	v2471 = v2312
	v2472 = v2312
	v2473 = v2312
	goto L575
L581:
	;
	switch v2368 - int32(1) {
	case 0:
		v2418 = v2369
		goto L588
	case 1:
		v2413 = v2369
		goto L589
	case 2:
		goto L590
	case 3:
		v2406 = v2370
		goto L591
	case 4:
		v2403 = v2370
		goto L592
	case 5:
		v2398 = v2370
		goto L593
	case 6:
		goto L594
	case 7:
		v2389 = v2371
		goto L595
	case 8:
		v2384 = v2371
		goto L596
	case 9:
		v2379 = v2371
		goto L597
	case 10:
		goto L598
	default:
		v2539 = v2369
		v2540 = v2370
		v2541 = v2371
		goto L574
	}
L582:
	;
	v2367 = v2228
	v2368 = v2306
	v2369 = v2312
	v2370 = v2312
	v2371 = v2312
	goto L581
L583:
	;
	goto L584
L584:
	;
	v2319 = v2228
	v2320 = v2306
	v2321 = v2312
	v2322 = v2312
	v2323 = v2312
	goto L585
L585:
	;
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(v2319)+4))
	v2326 = v2325 + v2322
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v2319)))
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(v2319)+8))
	v2330 = v2329 + v2323
	v2332 = int32(4)
	v2334 = v2327 + v2321 - v2330 ^ base.I32_rotl(v2330, v2332)
	v2338 = v2326 - v2334 ^ base.I32_rotl(v2334, int32(6))
	v2339 = v2330 + v2326
	v2340 = v2334 + v2339
	v2341 = v2338 + v2340
	v2345 = v2339 - v2338 ^ base.I32_rotl(v2338, int32(8))
	v2349 = v2340 - v2345 ^ base.I32_rotl(v2345, int32(16))
	v2353 = v2341 - v2349 ^ base.I32_rotl(v2349, int32(19))
	v2354 = v2345 + v2341
	v2355 = v2349 + v2354
	v2356 = v2353 + v2355
	v2360 = v2354 - v2353 ^ base.I32_rotl(v2353, v2332)
	v2361 = int32(12)
	v2362 = v2319 + v2361
	v2364 = v2320 - v2361
	if base.Ui32(int32(11)) < base.Ui32(v2364) {
		v2319 = v2362
		v2320 = v2364
		v2321 = v2355
		v2322 = v2356
		v2323 = v2360
		goto L585
	} else {
		goto L587
	}
L586:
	;
	v2367 = v2362
	v2368 = v2364
	v2369 = v2355
	v2370 = v2356
	v2371 = v2360
	goto L581
L587:
	;
	goto L586
L588:
	;
	v2419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2367))))
	v2539 = v2418 + v2419
	v2540 = v2370
	v2541 = v2371
	goto L574
L589:
	;
	v2414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2367)+1)))
	v2418 = v2414<<(uint(int32(8))%32) + v2413
	goto L588
L590:
	;
	v2409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2367)+2)))
	v2413 = v2409<<(uint(int32(16))%32) + v2369
	goto L589
L591:
	;
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v2367)))
	v2539 = v2407 + v2369
	v2540 = v2406
	v2541 = v2371
	goto L574
L592:
	;
	v2404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2367)+4)))
	v2406 = v2403 + v2404
	goto L591
L593:
	;
	v2399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2367)+5)))
	v2403 = v2399<<(uint(int32(8))%32) + v2398
	goto L592
L594:
	;
	v2394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2367)+6)))
	v2398 = v2394<<(uint(int32(16))%32) + v2370
	goto L593
L595:
	;
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(v2367)))
	v2392 = *(*int32)(unsafe.Add(mBase, uint32(v2367)+4))
	v2539 = v2390 + v2369
	v2540 = v2392 + v2370
	v2541 = v2389
	goto L574
L596:
	;
	v2385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2367)+8)))
	v2389 = v2385<<(uint(int32(8))%32) + v2384
	goto L595
L597:
	;
	v2380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2367)+9)))
	v2384 = v2380<<(uint(int32(16))%32) + v2379
	goto L596
L598:
	;
	v2375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2367)+10)))
	v2379 = v2375<<(uint(int32(24))%32) + v2371
	goto L597
L599:
	;
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(v2421)+4))
	v2428 = v2427 + v2424
	v2429 = *(*int32)(unsafe.Add(mBase, uint32(v2421)))
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(v2421)+8))
	v2432 = v2431 + v2425
	v2434 = int32(4)
	v2436 = v2429 + v2423 - v2432 ^ base.I32_rotl(v2432, v2434)
	v2440 = v2428 - v2436 ^ base.I32_rotl(v2436, int32(6))
	v2441 = v2432 + v2428
	v2442 = v2436 + v2441
	v2443 = v2440 + v2442
	v2447 = v2441 - v2440 ^ base.I32_rotl(v2440, int32(8))
	v2451 = v2442 - v2447 ^ base.I32_rotl(v2447, int32(16))
	v2455 = v2443 - v2451 ^ base.I32_rotl(v2451, int32(19))
	v2456 = v2447 + v2443
	v2457 = v2451 + v2456
	v2458 = v2455 + v2457
	v2462 = v2456 - v2455 ^ base.I32_rotl(v2455, v2434)
	v2463 = int32(12)
	v2464 = v2421 + v2463
	v2466 = v2422 - v2463
	if base.Ui32(int32(11)) < base.Ui32(v2466) {
		v2421 = v2464
		v2422 = v2466
		v2423 = v2457
		v2424 = v2458
		v2425 = v2462
		goto L599
	} else {
		goto L601
	}
L600:
	;
	v2469 = v2464
	v2470 = v2466
	v2471 = v2457
	v2472 = v2458
	v2473 = v2462
	goto L575
L601:
	;
	goto L600
L602:
	;
	v2535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2469))))
	v2539 = v2532 + v2535
	v2540 = v2533
	v2541 = v2534
	goto L574
L603:
	;
	v2528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2469)+1)))
	v2532 = v2528<<(uint(int32(8))%32) + v2525
	v2533 = v2526
	v2534 = v2527
	goto L602
L604:
	;
	v2521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2469)+2)))
	v2525 = v2521<<(uint(int32(16))%32) + v2518
	v2526 = v2519
	v2527 = v2520
	goto L603
L605:
	;
	v2514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2469)+3)))
	v2518 = v2514<<(uint(int32(24))%32) + v2471
	v2519 = v2512
	v2520 = v2513
	goto L604
L606:
	;
	v2510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2469)+4)))
	v2512 = v2508 + v2510
	v2513 = v2509
	goto L605
L607:
	;
	v2504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2469)+5)))
	v2508 = v2504<<(uint(int32(8))%32) + v2502
	v2509 = v2503
	goto L606
L608:
	;
	v2498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2469)+6)))
	v2502 = v2498<<(uint(int32(16))%32) + v2496
	v2503 = v2497
	goto L607
L609:
	;
	v2492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2469)+7)))
	v2496 = v2492<<(uint(int32(24))%32) + v2472
	v2497 = v2491
	goto L608
L610:
	;
	v2487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2469)+8)))
	v2491 = v2487<<(uint(int32(8))%32) + v2486
	goto L609
L611:
	;
	v2482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2469)+9)))
	v2486 = v2482<<(uint(int32(16))%32) + v2481
	goto L610
L612:
	;
	v2477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2469)+10)))
	v2481 = v2477<<(uint(int32(24))%32) + v2473
	goto L611
L613:
	;
	v3228 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2241)+64)) = v3228
	*(*int32)(unsafe.Add(mBase, uint32(v2241)+60)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v2241)+56)) = v2229
	v3233 = *(*int32)(unsafe.Add(mBase, uint32(v385)+28))
	v3235 = v2241 + int32(56)
	v3241 = m.G0
	v3242 = int32(16)
	v3243 = v3241 - v3242
	m.G0 = v3243
	v3245 = *(*int32)(unsafe.Add(mBase, uint32(v3235)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3243)+8)) = v3245
	v3247 = *(*int64)(unsafe.Add(mBase, uint32(v3235)))
	*(*int64)(unsafe.Add(mBase, uint32(v3243))) = v3247
	v3249 = *(*int32)(unsafe.Add(mBase, uint32(v3233)))
	*(*int32)(unsafe.Add(mBase, uint32(v3243)+12)) = v3228
	v3252 = F_hash_bytes(m, v3243, v3242)
	mBase = m.M
	v3253 = *(*int32)(unsafe.Add(mBase, uint32(v3249)+20))
	v3254 = *(*int32)(unsafe.Add(mBase, uint32(v3249)+12))
	v3255 = v3252 & v3254
	v3258 = v3253 + v3255*int32(40)
	v3259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3258)+20)))
	if v3259 == v3228 {
		goto L714
	} else {
		goto L715
	}
L614:
	;
	v2588 = v2573
	goto L617
L615:
	;
	goto L616
L616:
	;
	F_GetRelationPath(m, v2241+int32(56), v410, v2229, v2230, int32(-1), v2231)
	mBase = m.M
	v2714 = m.ExcPending
	if v2714 != 0 {
		goto L4
	} else {
		goto L629
	}
L617:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v2571+v2588<<(uint(int32(4))%32))+4))
	v2631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2228))))
	v2632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2628))))
	if v2632 == int32(0) {
		v2651 = v2631
		v2652 = v2632
		goto L620
	} else {
		goto L621
	}
L618:
	;
	goto L616
L619:
	;
	if v2652-v2651 == int32(0) {
		goto L613
	} else {
		goto L627
	}
L620:
	;
	goto L619
L621:
	;
	if v2631 != v2632 {
		v2651 = v2631
		v2652 = v2632
		goto L620
	} else {
		goto L622
	}
L622:
	;
	v2636 = v2628
	v2637 = v2228
	goto L623
L623:
	;
	v2640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2637)+1)))
	v2641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2636)+1)))
	if v2641 == int32(0) {
		v2651 = v2640
		v2652 = v2641
		goto L620
	} else {
		goto L625
	}
L624:
	;
	v2651 = v2640
	v2652 = v2641
	goto L620
L625:
	;
	v2644 = int32(1)
	if v2640 == v2641 {
		v2636 = v2636 + v2644
		v2637 = v2637 + v2644
		goto L623
	} else {
		goto L626
	}
L626:
	;
	goto L624
L627:
	;
	v2658 = (v2588 + int32(1)) & v2572
	v2662 = *(*int32)(unsafe.Add(mBase, uint32(v2571+v2658<<(uint(int32(4))%32))))
	if v2662 != 0 {
		v2588 = v2658
		goto L617
	} else {
		goto L628
	}
L628:
	;
	goto L618
L629:
	;
	v2716 = v2241 + int32(56)
	v2720 = F_strlen(m, v2716)
	mBase = m.M
	v2727 = v2720 + int32(1)
	goto L632
L630:
	;
	v2740 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2739))) = uint8(v2740)
	v2743 = v2739 + int32(1)
	if v2232 != 0 {
		goto L637
	} else {
		goto L638
	}
L631:
	;
	goto L630
L632:
	;
	v2729 = int32(0)
	if v2727 == v2729 {
		v2739 = v2729
		goto L631
	} else {
		goto L634
	}
L633:
	;
	v2739 = v2734
	goto L631
L634:
	;
	v2733 = v2727 - int32(1)
	v2734 = v2716 + v2733
	v2735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2734))))
	if v2735 != int32(47) {
		v2727 = v2733
		goto L632
	} else {
		goto L635
	}
L635:
	;
	goto L633
L636:
	;
	v2764 = *(*int32)(unsafe.Add(mBase, uint32(v385)+24))
	if v2763&int32(3) == int32(0) {
		v2788 = v2763
		goto L644
	} else {
		goto L645
	}
L637:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2241)+40)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v2241)+36)) = v2743
	*(*int32)(unsafe.Add(mBase, uint32(v2241)+32)) = v2241 + int32(56)
	v2752 = F_psprintf(m, int32(_a_F_sendDir_48), v2241+int32(32))
	mBase = m.M
	v2753 = m.ExcPending
	if v2753 != 0 {
		goto L4
	} else {
		goto L640
	}
L638:
	;
	goto L639
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2241)+20)) = v2743
	*(*int32)(unsafe.Add(mBase, uint32(v2241)+16)) = v2241 + int32(56)
	v2761 = F_psprintf(m, int32(_a_F_sendDir_49), v2241+int32(16))
	mBase = m.M
	v2762 = m.ExcPending
	if v2762 != 0 {
		goto L4
	} else {
		goto L641
	}
L640:
	;
	v2763 = v2752
	goto L636
L641:
	;
	v2763 = v2761
	goto L636
L642:
	;
	v2827 = v2821 - int32(1636608432)
	if v2763&int32(3) != 0 {
		goto L663
	} else {
		goto L664
	}
L643:
	;
	v2821 = v2813 - v2763
	goto L642
L644:
	;
	v2792 = v2788
	goto L653
L645:
	;
	v2772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2763))))
	if v2772 == int32(0) {
		goto L646
	} else {
		goto L647
	}
L646:
	;
	v2821 = int32(0)
	goto L642
L647:
	;
	goto L648
L648:
	;
	v2777 = v2763
	goto L649
L649:
	;
	v2781 = v2777 + int32(1)
	if v2781&int32(3) == int32(0) {
		v2788 = v2781
		goto L644
	} else {
		goto L651
	}
L650:
	;
	v2813 = v2781
	goto L643
L651:
	;
	v2786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2781))))
	if v2786 != 0 {
		v2777 = v2781
		goto L649
	} else {
		goto L652
	}
L652:
	;
	goto L650
L653:
	;
	v2798 = *(*int32)(unsafe.Add(mBase, uint32(v2792)))
	v2801 = int32(-2139062144)
	if (int32(16843008)-v2798|v2798)&v2801 == v2801 {
		v2792 = v2792 + int32(4)
		goto L653
	} else {
		goto L655
	}
L654:
	;
	v2807 = v2792
	goto L656
L655:
	;
	goto L654
L656:
	;
	v2811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2807))))
	if v2811 != 0 {
		v2807 = v2807 + int32(1)
		goto L656
	} else {
		goto L658
	}
L657:
	;
	v2813 = v2807
	goto L643
L658:
	;
	goto L657
L659:
	;
	v3086 = int32(0)
	v3087 = *(*int32)(unsafe.Add(mBase, uint32(v2764)+20))
	v3088 = *(*int32)(unsafe.Add(mBase, uint32(v2764)+12))
	v3089 = (v3081 ^ v3073 - base.I32_rotl(v3081, int32(24))) & v3088
	v3093 = *(*int32)(unsafe.Add(mBase, uint32(v3087+v3089<<(uint(int32(4))%32))))
	if v3093 == v3086 {
		v4042 = v3086
		goto L552
	} else {
		goto L699
	}
L660:
	;
	v3059 = int32(14)
	v3061 = v3055 ^ v3056 - base.I32_rotl(v3055, v3059)
	v3065 = v3061 ^ v3054 - base.I32_rotl(v3061, int32(11))
	v3069 = v3065 ^ v3055 - base.I32_rotl(v3065, int32(25))
	v3073 = v3069 ^ v3061 - base.I32_rotl(v3069, int32(16))
	v3077 = v3073 ^ v3065 - base.I32_rotl(v3073, int32(4))
	v3081 = v3077 ^ v3069 - base.I32_rotl(v3077, v3059)
	goto L659
L661:
	;
	switch v2985 - int32(1) {
	case 0:
		v3047 = v2986
		v3048 = v2987
		v3049 = v2988
		goto L688
	case 1:
		v3040 = v2986
		v3041 = v2987
		v3042 = v2988
		goto L689
	case 2:
		v3033 = v2986
		v3034 = v2987
		v3035 = v2988
		goto L690
	case 3:
		v3027 = v2987
		v3028 = v2988
		goto L691
	case 4:
		v3023 = v2987
		v3024 = v2988
		goto L692
	case 5:
		v3017 = v2987
		v3018 = v2988
		goto L693
	case 6:
		v3011 = v2987
		v3012 = v2988
		goto L694
	case 7:
		v3006 = v2988
		goto L695
	case 8:
		v3001 = v2988
		goto L696
	case 9:
		v2996 = v2988
		goto L697
	case 10:
		goto L698
	default:
		v3054 = v2986
		v3055 = v2987
		v3056 = v2988
		goto L660
	}
L662:
	;
	v2936 = v2763
	v2937 = v2821
	v2938 = v2827
	v2939 = v2827
	v2940 = v2827
	goto L685
L663:
	;
	if base.Ui32(int32(11)) < base.Ui32(v2821) {
		goto L662
	} else {
		goto L666
	}
L664:
	;
	goto L665
L665:
	;
	if base.Ui32(v2821) < base.Ui32(int32(12)) {
		goto L668
	} else {
		goto L669
	}
L666:
	;
	v2984 = v2763
	v2985 = v2821
	v2986 = v2827
	v2987 = v2827
	v2988 = v2827
	goto L661
L667:
	;
	switch v2883 - int32(1) {
	case 0:
		v2933 = v2884
		goto L674
	case 1:
		v2928 = v2884
		goto L675
	case 2:
		goto L676
	case 3:
		v2921 = v2885
		goto L677
	case 4:
		v2918 = v2885
		goto L678
	case 5:
		v2913 = v2885
		goto L679
	case 6:
		goto L680
	case 7:
		v2904 = v2886
		goto L681
	case 8:
		v2899 = v2886
		goto L682
	case 9:
		v2894 = v2886
		goto L683
	case 10:
		goto L684
	default:
		v3054 = v2884
		v3055 = v2885
		v3056 = v2886
		goto L660
	}
L668:
	;
	v2882 = v2763
	v2883 = v2821
	v2884 = v2827
	v2885 = v2827
	v2886 = v2827
	goto L667
L669:
	;
	goto L670
L670:
	;
	v2834 = v2763
	v2835 = v2821
	v2836 = v2827
	v2837 = v2827
	v2838 = v2827
	goto L671
L671:
	;
	v2840 = *(*int32)(unsafe.Add(mBase, uint32(v2834)+4))
	v2841 = v2840 + v2837
	v2842 = *(*int32)(unsafe.Add(mBase, uint32(v2834)))
	v2844 = *(*int32)(unsafe.Add(mBase, uint32(v2834)+8))
	v2845 = v2844 + v2838
	v2847 = int32(4)
	v2849 = v2842 + v2836 - v2845 ^ base.I32_rotl(v2845, v2847)
	v2853 = v2841 - v2849 ^ base.I32_rotl(v2849, int32(6))
	v2854 = v2845 + v2841
	v2855 = v2849 + v2854
	v2856 = v2853 + v2855
	v2860 = v2854 - v2853 ^ base.I32_rotl(v2853, int32(8))
	v2864 = v2855 - v2860 ^ base.I32_rotl(v2860, int32(16))
	v2868 = v2856 - v2864 ^ base.I32_rotl(v2864, int32(19))
	v2869 = v2860 + v2856
	v2870 = v2864 + v2869
	v2871 = v2868 + v2870
	v2875 = v2869 - v2868 ^ base.I32_rotl(v2868, v2847)
	v2876 = int32(12)
	v2877 = v2834 + v2876
	v2879 = v2835 - v2876
	if base.Ui32(int32(11)) < base.Ui32(v2879) {
		v2834 = v2877
		v2835 = v2879
		v2836 = v2870
		v2837 = v2871
		v2838 = v2875
		goto L671
	} else {
		goto L673
	}
L672:
	;
	v2882 = v2877
	v2883 = v2879
	v2884 = v2870
	v2885 = v2871
	v2886 = v2875
	goto L667
L673:
	;
	goto L672
L674:
	;
	v2934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2882))))
	v3054 = v2933 + v2934
	v3055 = v2885
	v3056 = v2886
	goto L660
L675:
	;
	v2929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2882)+1)))
	v2933 = v2929<<(uint(int32(8))%32) + v2928
	goto L674
L676:
	;
	v2924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2882)+2)))
	v2928 = v2924<<(uint(int32(16))%32) + v2884
	goto L675
L677:
	;
	v2922 = *(*int32)(unsafe.Add(mBase, uint32(v2882)))
	v3054 = v2922 + v2884
	v3055 = v2921
	v3056 = v2886
	goto L660
L678:
	;
	v2919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2882)+4)))
	v2921 = v2918 + v2919
	goto L677
L679:
	;
	v2914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2882)+5)))
	v2918 = v2914<<(uint(int32(8))%32) + v2913
	goto L678
L680:
	;
	v2909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2882)+6)))
	v2913 = v2909<<(uint(int32(16))%32) + v2885
	goto L679
L681:
	;
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(v2882)))
	v2907 = *(*int32)(unsafe.Add(mBase, uint32(v2882)+4))
	v3054 = v2905 + v2884
	v3055 = v2907 + v2885
	v3056 = v2904
	goto L660
L682:
	;
	v2900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2882)+8)))
	v2904 = v2900<<(uint(int32(8))%32) + v2899
	goto L681
L683:
	;
	v2895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2882)+9)))
	v2899 = v2895<<(uint(int32(16))%32) + v2894
	goto L682
L684:
	;
	v2890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2882)+10)))
	v2894 = v2890<<(uint(int32(24))%32) + v2886
	goto L683
L685:
	;
	v2942 = *(*int32)(unsafe.Add(mBase, uint32(v2936)+4))
	v2943 = v2942 + v2939
	v2944 = *(*int32)(unsafe.Add(mBase, uint32(v2936)))
	v2946 = *(*int32)(unsafe.Add(mBase, uint32(v2936)+8))
	v2947 = v2946 + v2940
	v2949 = int32(4)
	v2951 = v2944 + v2938 - v2947 ^ base.I32_rotl(v2947, v2949)
	v2955 = v2943 - v2951 ^ base.I32_rotl(v2951, int32(6))
	v2956 = v2947 + v2943
	v2957 = v2951 + v2956
	v2958 = v2955 + v2957
	v2962 = v2956 - v2955 ^ base.I32_rotl(v2955, int32(8))
	v2966 = v2957 - v2962 ^ base.I32_rotl(v2962, int32(16))
	v2970 = v2958 - v2966 ^ base.I32_rotl(v2966, int32(19))
	v2971 = v2962 + v2958
	v2972 = v2966 + v2971
	v2973 = v2970 + v2972
	v2977 = v2971 - v2970 ^ base.I32_rotl(v2970, v2949)
	v2978 = int32(12)
	v2979 = v2936 + v2978
	v2981 = v2937 - v2978
	if base.Ui32(int32(11)) < base.Ui32(v2981) {
		v2936 = v2979
		v2937 = v2981
		v2938 = v2972
		v2939 = v2973
		v2940 = v2977
		goto L685
	} else {
		goto L687
	}
L686:
	;
	v2984 = v2979
	v2985 = v2981
	v2986 = v2972
	v2987 = v2973
	v2988 = v2977
	goto L661
L687:
	;
	goto L686
L688:
	;
	v3050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2984))))
	v3054 = v3047 + v3050
	v3055 = v3048
	v3056 = v3049
	goto L660
L689:
	;
	v3043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2984)+1)))
	v3047 = v3043<<(uint(int32(8))%32) + v3040
	v3048 = v3041
	v3049 = v3042
	goto L688
L690:
	;
	v3036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2984)+2)))
	v3040 = v3036<<(uint(int32(16))%32) + v3033
	v3041 = v3034
	v3042 = v3035
	goto L689
L691:
	;
	v3029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2984)+3)))
	v3033 = v3029<<(uint(int32(24))%32) + v2986
	v3034 = v3027
	v3035 = v3028
	goto L690
L692:
	;
	v3025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2984)+4)))
	v3027 = v3023 + v3025
	v3028 = v3024
	goto L691
L693:
	;
	v3019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2984)+5)))
	v3023 = v3019<<(uint(int32(8))%32) + v3017
	v3024 = v3018
	goto L692
L694:
	;
	v3013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2984)+6)))
	v3017 = v3013<<(uint(int32(16))%32) + v3011
	v3018 = v3012
	goto L693
L695:
	;
	v3007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2984)+7)))
	v3011 = v3007<<(uint(int32(24))%32) + v2987
	v3012 = v3006
	goto L694
L696:
	;
	v3002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2984)+8)))
	v3006 = v3002<<(uint(int32(8))%32) + v3001
	goto L695
L697:
	;
	v2997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2984)+9)))
	v3001 = v2997<<(uint(int32(16))%32) + v2996
	goto L696
L698:
	;
	v2992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2984)+10)))
	v2996 = v2992<<(uint(int32(24))%32) + v2988
	goto L697
L699:
	;
	v3106 = v3089
	goto L700
L700:
	;
	v3146 = *(*int32)(unsafe.Add(mBase, uint32(v3087+v3106<<(uint(int32(4))%32))+4))
	v3149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2763))))
	v3150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3146))))
	if v3150 == int32(0) {
		v3169 = v3149
		v3170 = v3150
		goto L703
	} else {
		goto L704
	}
L701:
	;
	v4042 = v3086
	goto L552
L702:
	;
	if v3170-v3169 == int32(0) {
		goto L613
	} else {
		goto L710
	}
L703:
	;
	goto L702
L704:
	;
	if v3149 != v3150 {
		v3169 = v3149
		v3170 = v3150
		goto L703
	} else {
		goto L705
	}
L705:
	;
	v3154 = v3146
	v3155 = v2763
	goto L706
L706:
	;
	v3158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3155)+1)))
	v3159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3154)+1)))
	if v3159 == int32(0) {
		v3169 = v3158
		v3170 = v3159
		goto L703
	} else {
		goto L708
	}
L707:
	;
	v3169 = v3158
	v3170 = v3159
	goto L703
L708:
	;
	v3162 = int32(1)
	if v3158 == v3159 {
		v3154 = v3154 + v3162
		v3155 = v3155 + v3162
		goto L706
	} else {
		goto L709
	}
L709:
	;
	goto L707
L710:
	;
	v3176 = (v3106 + int32(1)) & v3088
	v3180 = *(*int32)(unsafe.Add(mBase, uint32(v3087+v3176<<(uint(int32(4))%32))))
	if v3180 != 0 {
		v3106 = v3176
		goto L700
	} else {
		goto L711
	}
L711:
	;
	goto L701
L712:
	;
	if v3290 != 0 {
		v4042 = v3228
		goto L552
	} else {
		goto L725
	}
L713:
	;
	m.G0 = v3243 + int32(16)
	goto L712
L714:
	;
	v3290 = int32(0)
	goto L713
L715:
	;
	v3262 = v3255
	v3263 = v3258
	goto L716
L716:
	;
	v3269 = F_memcmp(m, v3263, v3243, int32(16))
	mBase = m.M
	if v3269 != 0 {
		goto L718
	} else {
		goto L719
	}
L717:
	;
	if v3263 == int32(0) {
		goto L722
	} else {
		goto L723
	}
L718:
	;
	v3272 = (v3262 + int32(1)) & v3254
	v3275 = v3253 + v3272*int32(40)
	v3276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3275)+20)))
	if v3276 != 0 {
		v3262 = v3272
		v3263 = v3275
		goto L716
	} else {
		goto L721
	}
L719:
	;
	goto L720
L720:
	;
	goto L717
L721:
	;
	goto L714
L722:
	;
	v3290 = int32(0)
	goto L713
L723:
	;
	goto L724
L724:
	;
	v3280 = *(*int32)(unsafe.Add(mBase, uint32(v3263)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2241+int32(52)))) = v3280
	v3290 = v3263
	goto L713
L725:
	;
	v3299 = int32(base.Ui32(v2233) >> (uint(int32(13)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v2241)+64)) = v2230
	v3301 = *(*int32)(unsafe.Add(mBase, uint32(v385)+28))
	v3303 = v2241 + int32(56)
	v3308 = m.G0
	v3309 = int32(16)
	v3310 = v3308 - v3309
	m.G0 = v3310
	v3312 = *(*int32)(unsafe.Add(mBase, uint32(v3303)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3310)+8)) = v3312
	v3314 = *(*int64)(unsafe.Add(mBase, uint32(v3303)))
	*(*int64)(unsafe.Add(mBase, uint32(v3310))) = v3314
	v3316 = *(*int32)(unsafe.Add(mBase, uint32(v3301)))
	*(*int32)(unsafe.Add(mBase, uint32(v3310)+12)) = v2231
	v3319 = F_hash_bytes(m, v3310, v3309)
	mBase = m.M
	v3320 = *(*int32)(unsafe.Add(mBase, uint32(v3316)+20))
	v3321 = *(*int32)(unsafe.Add(mBase, uint32(v3316)+12))
	v3322 = v3319 & v3321
	v3325 = v3320 + v3322*int32(40)
	v3326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3325)+20)))
	if v3326 == int32(0) {
		goto L729
	} else {
		goto L730
	}
L726:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2237))) = v3979
	v4042 = int32(1)
	goto L552
L727:
	;
	if v3357 == int32(0) {
		goto L740
	} else {
		goto L741
	}
L728:
	;
	m.G0 = v3310 + int32(16)
	goto L727
L729:
	;
	v3357 = int32(0)
	goto L728
L730:
	;
	v3329 = v3322
	v3330 = v3325
	goto L731
L731:
	;
	v3336 = F_memcmp(m, v3330, v3310, int32(16))
	mBase = m.M
	if v3336 != 0 {
		goto L733
	} else {
		goto L734
	}
L732:
	;
	if v3330 == int32(0) {
		goto L737
	} else {
		goto L738
	}
L733:
	;
	v3339 = (v3329 + int32(1)) & v3321
	v3342 = v3320 + v3339*int32(40)
	v3343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3342)+20)))
	if v3343 != 0 {
		v3329 = v3339
		v3330 = v3342
		goto L731
	} else {
		goto L736
	}
L734:
	;
	goto L735
L735:
	;
	goto L732
L736:
	;
	goto L729
L737:
	;
	v3357 = int32(0)
	goto L728
L738:
	;
	goto L739
L739:
	;
	v3347 = *(*int32)(unsafe.Add(mBase, uint32(v3330)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2241+int32(52)))) = v3347
	v3357 = v3330
	goto L728
L740:
	;
	if v2233 == int32(0) {
		v4042 = v3228
		goto L552
	} else {
		goto L743
	}
L741:
	;
	goto L742
L742:
	;
	v3372 = v2232 << (uint(int32(17)) % 32)
	v3373 = *(*int32)(unsafe.Add(mBase, uint32(v2241)+52))
	if base.Ui32(v3373) <= base.Ui32(v3372) {
		v4042 = v3228
		goto L552
	} else {
		goto L744
	}
L743:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2235))) = int32(0)
	v3979 = v3299
	goto L726
L744:
	;
	if base.Ui32(int32(_a_F_sendDir_50)) < base.Ui32(v2232) {
		goto L551
	} else {
		goto L745
	}
L745:
	;
	v3377 = v3299 + v3372
	if base.Ui32(v3377) < base.Ui32(v3299) {
		goto L551
	} else {
		goto L746
	}
L746:
	;
	v3379 = int32(0)
	v3380 = int32(16)
	v3381 = int32(base.Ui32(v3372) >> (uint(v3380) % 32))
	v3388 = int32(base.Ui32(v3377)>>(uint(v3380)%32)) + base.B2i32(v3377&int32(_a_F_sendDir_51) != v3379)
	v3389 = *(*int32)(unsafe.Add(mBase, uint32(v3357)+24))
	if base.Ui32(v3388) < base.Ui32(v3389) {
		goto L748
	} else {
		goto L749
	}
L747:
	;
	if base.F64_gt(base.F64_convert_i32_u(v3684<<(uint(int32(13))%32)), base.F64_mul(base.F64_convert_i32_u(v2233), float64(0.9))) != 0 {
		v4042 = v3228
		goto L552
	} else {
		goto L781
	}
L748:
	;
	v3391 = v3388
	goto L750
L749:
	;
	v3391 = v3389
	goto L750
L750:
	;
	if base.Ui32(v3391) <= base.Ui32(v3381) {
		v3684 = v3379
		goto L747
	} else {
		goto L751
	}
L751:
	;
	v3417 = v3381
	v3425 = v3379
	goto L752
L752:
	;
	v3444 = *(*int32)(unsafe.Add(mBase, uint32(v3357)+32))
	v3448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3444+v3417<<(uint(int32(1))%32)))))
	if v3448 == int32(0) {
		v3634 = v3425
		goto L754
	} else {
		goto L755
	}
L753:
	;
	v3684 = v3634
	goto L747
L754:
	;
	v3654 = v3417 + int32(1)
	if v3654 != v3391 {
		v3417 = v3654
		v3425 = v3634
		goto L752
	} else {
		goto L780
	}
L755:
	;
	v3453 = v3417 << (uint(int32(16)) % 32)
	if v3391-int32(1) != v3417 {
		goto L756
	} else {
		goto L757
	}
L756:
	;
	v3456 = int32(_a_F_sendDir_52)
	goto L758
L757:
	;
	v3456 = v3377 - v3453
	goto L758
L758:
	;
	if v3381 == v3417 {
		goto L759
	} else {
		goto L760
	}
L759:
	;
	v3459 = v3372 & int32(_a_F_sendDir_51)
	goto L761
L760:
	;
	v3459 = int32(0)
	goto L761
L761:
	;
	v3460 = *(*int32)(unsafe.Add(mBase, uint32(v3357)+36))
	v3464 = *(*int32)(unsafe.Add(mBase, uint32(v3460+v3417<<(uint(int32(2))%32))))
	if v3448 != int32(_a_F_sendDir_53) {
		goto L762
	} else {
		goto L763
	}
L762:
	;
	v3496 = v3425
	v3500 = int32(0)
	goto L765
L763:
	;
	goto L764
L764:
	;
	if base.Ui32(v3456) <= base.Ui32(v3459) {
		v3634 = v3425
		goto L754
	} else {
		goto L772
	}
L765:
	;
	v3518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3464+v3500<<(uint(int32(1))%32)))))
	if base.Ui32(v3518) < base.Ui32(v3459) {
		v3530 = v3496
		goto L767
	} else {
		goto L768
	}
L766:
	;
	v3634 = v3530
	goto L754
L767:
	;
	v3532 = v3500 + int32(1)
	if v3532 != v3448 {
		v3496 = v3530
		v3500 = v3532
		goto L765
	} else {
		goto L771
	}
L768:
	;
	if base.Ui32(v3456) <= base.Ui32(v3518) {
		v3530 = v3496
		goto L767
	} else {
		goto L769
	}
L769:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403+v3496<<(uint(int32(2))%32)))) = v3518 | v3453
	v3527 = v3496 + int32(1)
	if v3527 == int32(_a_F_sendDir_54) {
		v3684 = v3527
		goto L747
	} else {
		goto L770
	}
L770:
	;
	v3530 = v3527
	goto L767
L771:
	;
	goto L766
L772:
	;
	v3560 = v3459
	v3563 = v3425
	goto L773
L773:
	;
	v3587 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3464+int32(base.Ui32(v3560)>>(uint(int32(3))%32))&int32(536870910)))))
	if int32(base.Ui32(v3587)>>(uint(v3560&int32(15))%32))&int32(1) != 0 {
		goto L775
	} else {
		goto L776
	}
L774:
	;
	v3634 = v3602
	goto L754
L775:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403+v3563<<(uint(int32(2))%32)))) = v3560 + v3453
	v3599 = v3563 + int32(1)
	if v3599 == int32(_a_F_sendDir_54) {
		v3684 = v3599
		goto L747
	} else {
		goto L778
	}
L776:
	;
	v3602 = v3563
	goto L777
L777:
	;
	v3604 = v3560 + int32(1)
	if v3604 != v3456 {
		v3560 = v3604
		v3563 = v3602
		goto L773
	} else {
		goto L779
	}
L778:
	;
	v3602 = v3599
	goto L777
L779:
	;
	goto L774
L780:
	;
	goto L753
L781:
	;
	F_pg_qsort(m, v403, v3684, int32(4), int32(434))
	mBase = m.M
	v3713 = m.ExcPending
	if v3713 != 0 {
		goto L4
	} else {
		goto L782
	}
L782:
	;
	if v3372 == int32(0) {
		goto L783
	} else {
		goto L784
	}
L783:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2235))) = v3684
	*(*int32)(unsafe.Add(mBase, uint32(v2237))) = v3299
	v3955 = int32(1)
	v3956 = *(*int32)(unsafe.Add(mBase, uint32(v2241)+52))
	if v3956 == int32(-1) {
		v4042 = v3955
		goto L552
	} else {
		goto L796
	}
L784:
	;
	if v3684 == int32(0) {
		goto L783
	} else {
		goto L785
	}
L785:
	;
	v3719 = v3684 & int32(3)
	v3720 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v3684) {
		goto L786
	} else {
		goto L787
	}
L786:
	;
	v3736 = v3720
	v3750 = v3228
	goto L789
L787:
	;
	v3809 = v3720
	goto L788
L788:
	;
	if v3719 == int32(0) {
		goto L783
	} else {
		goto L792
	}
L789:
	;
	v3775 = v403 + v3736<<(uint(int32(2))%32)
	v3776 = *(*int32)(unsafe.Add(mBase, uint32(v3775)))
	*(*int32)(unsafe.Add(mBase, uint32(v3775))) = v3776 - v3372
	v3779 = int32(4)
	v3780 = v3775 + v3779
	v3781 = *(*int32)(unsafe.Add(mBase, uint32(v3780)))
	*(*int32)(unsafe.Add(mBase, uint32(v3780))) = v3781 - v3372
	v3785 = v3775 + int32(8)
	v3786 = *(*int32)(unsafe.Add(mBase, uint32(v3785)))
	*(*int32)(unsafe.Add(mBase, uint32(v3785))) = v3786 - v3372
	v3790 = v3775 + int32(12)
	v3791 = *(*int32)(unsafe.Add(mBase, uint32(v3790)))
	*(*int32)(unsafe.Add(mBase, uint32(v3790))) = v3791 - v3372
	v3795 = v3736 + v3779
	v3797 = v3750 + v3779
	if v3797 != v3684&int32(-4) {
		v3736 = v3795
		v3750 = v3797
		goto L789
	} else {
		goto L791
	}
L790:
	;
	v3809 = v3795
	goto L788
L791:
	;
	goto L790
L792:
	;
	v3858 = v3809
	v3871 = v3720
	goto L793
L793:
	;
	v3897 = v403 + v3858<<(uint(int32(2))%32)
	v3898 = *(*int32)(unsafe.Add(mBase, uint32(v3897)))
	*(*int32)(unsafe.Add(mBase, uint32(v3897))) = v3898 - v3372
	v3901 = int32(1)
	v3904 = v3871 + v3901
	if v3904 != v3719 {
		v3858 = v3858 + v3901
		v3871 = v3904
		goto L793
	} else {
		goto L795
	}
L794:
	;
	goto L783
L795:
	;
	goto L794
L796:
	;
	v3960 = v3956 - v3372
	if base.Ui32(v3960) < base.Ui32(v3299) {
		goto L797
	} else {
		goto L798
	}
L797:
	;
	v3962 = v3299
	goto L799
L798:
	;
	v3962 = v3960
	goto L799
L799:
	;
	if base.Ui32(int32(_a_F_sendDir_54)) <= base.Ui32(v3962) {
		goto L800
	} else {
		goto L801
	}
L800:
	;
	v3965 = int32(_a_F_sendDir_54)
	goto L802
L801:
	;
	v3965 = v3962
	goto L802
L802:
	;
	if base.Ui32(v3299) < base.Ui32(v3960) {
		v3979 = v3965
		goto L726
	} else {
		goto L803
	}
L803:
	;
	if base.Ui32(v3962) < base.Ui32(int32(_a_F_sendDir_55)) {
		v4042 = v3955
		goto L552
	} else {
		goto L804
	}
L804:
	;
	v3979 = v3965
	goto L726
L805:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v4075 = m.ExcPending
	if v4075 != 0 {
		goto L4
	} else {
		goto L806
	}
L806:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2241)+4)) = v2233
	*(*int32)(unsafe.Add(mBase, uint32(v2241))) = v2232
	F_errmsg_internal(m, int32(_a_F_sendDir_56), v2241)
	mBase = m.M
	v4080 = m.ExcPending
	if v4080 != 0 {
		goto L4
	} else {
		goto L807
	}
L807:
	;
	F_errfinish(m, int32(_a_F_sendDir_57), int32(796), int32(_a_F_sendDir_58))
	mBase = m.M
	v4085 = m.ExcPending
	if v4085 != 0 {
		goto L4
	} else {
		goto L808
	}
L808:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L809:
	;
	v4088 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2256))
	v4090 = v4088 << (uint(int32(2)) % 32)
	v4092 = v4090 + int32(12)
	if v4088 == int32(0) {
		v4102 = v4092
		goto L812
	} else {
		goto L813
	}
L810:
	;
	v4124 = v2205
	v4125 = v404
	goto L811
L811:
	;
	F_pfree(m, v2228)
	mBase = m.M
	v4127 = m.ExcPending
	if v4127 != 0 {
		goto L4
	} else {
		goto L816
	}
L812:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+96)) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v386)+100)) = v443
	*(*int64)(unsafe.Add(mBase, uint32(v386)+2296)) = base.I64_extend_i32_u(v4102 + v4088<<(uint(int32(13))%32))
	v4117 = F_pg_snprintf(m, v386+int32(192), int32(2048), int32(_a_F_sendDir_49), v386+int32(96))
	mBase = m.M
	v4118 = m.ExcPending
	if v4118 != 0 {
		goto L4
	} else {
		goto L815
	}
L813:
	;
	v4096 = v4092 & int32(_a_F_sendDir_59)
	if v4096 == int32(0) {
		v4102 = v4092
		goto L812
	} else {
		goto L814
	}
L814:
	;
	v4102 = v4090 - v4096 + int32(_a_F_sendDir_60)
	goto L812
L815:
	;
	v4124 = v403
	v4125 = v386 + int32(192)
	goto L811
L816:
	;
	v4144 = v4124
	v4158 = v4125
	goto L543
L817:
	;
	v4182 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2268))
	v4183 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2260))
	v4184 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2256))
	v4185 = *(*int32)(unsafe.Add(mBase, uint32(v386)+2252))
	v4186 = F_sendFile(m, v377, v386+int32(2368), v4158, v386+int32(2272), int32(1), v410, v384, v4182, v4183, v383, v4184, v4144, v4185)
	mBase = m.M
	v4187 = m.ExcPending
	if v4187 != 0 {
		goto L4
	} else {
		goto L820
	}
L818:
	;
	goto L819
L819:
	;
	v4190 = *(*int64)(unsafe.Add(mBase, uint32(v386)+2296))
	v4266 = v422 + v4190 + ((v4190+int64(511))&int64(4294966784)-v4190)&int64(4294967295) + int64(512)
	goto L113
L820:
	;
	if v4186 == int32(0) {
		v4266 = v422
		goto L113
	} else {
		goto L821
	}
L821:
	;
	goto L819
L822:
	;
	if v4204 == int32(0) {
		v4266 = v422
		goto L113
	} else {
		goto L823
	}
L823:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+80)) = v386 + int32(2368)
	F_errmsg(m, int32(_a_F_sendDir_61), v386+int32(80))
	mBase = m.M
	v4215 = m.ExcPending
	if v4215 != 0 {
		goto L4
	} else {
		goto L824
	}
L824:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1546), int32(_a_F_sendDir_20))
	mBase = m.M
	v4220 = m.ExcPending
	if v4220 != 0 {
		goto L4
	} else {
		goto L825
	}
L825:
	;
	v4266 = v422
	goto L113
L826:
	;
	if v4268 != 0 {
		v388 = v4268
		v422 = v4266
		goto L111
	} else {
		goto L827
	}
L827:
	;
	goto L112
L828:
	;
	F_pfree(m, v4296)
	mBase = m.M
	v4318 = m.ExcPending
	if v4318 != 0 {
		goto L4
	} else {
		goto L831
	}
L829:
	;
	goto L830
L830:
	;
	F_FreeDir(m, v4301)
	mBase = m.M
	v4320 = m.ExcPending
	if v4320 != 0 {
		goto L4
	} else {
		goto L832
	}
L831:
	;
	goto L830
L832:
	;
	m.G0 = v4279 + int32(_a_F_sendDir_0)
	return v4315
L833:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v4331 = m.ExcPending
	if v4331 != 0 {
		goto L4
	} else {
		goto L834
	}
L834:
	;
	F_errmsg(m, int32(_a_F_sendDir_62), int32(0))
	mBase = m.M
	v4335 = m.ExcPending
	if v4335 != 0 {
		goto L4
	} else {
		goto L835
	}
L835:
	;
	F_errhint(m, int32(_a_F_sendDir_63), int32(0))
	mBase = m.M
	v4339 = m.ExcPending
	if v4339 != 0 {
		goto L4
	} else {
		goto L836
	}
L836:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1286), int32(_a_F_sendDir_20))
	mBase = m.M
	v4344 = m.ExcPending
	if v4344 != 0 {
		goto L4
	} else {
		goto L837
	}
L837:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L838:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4350 = m.ExcPending
	if v4350 != 0 {
		goto L4
	} else {
		goto L839
	}
L839:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+48)) = v386 + int32(2368)
	F_errmsg(m, int32(_a_F_sendDir_64), v386+int32(48))
	mBase = m.M
	v4358 = m.ExcPending
	if v4358 != 0 {
		goto L4
	} else {
		goto L840
	}
L840:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1420), int32(_a_F_sendDir_20))
	mBase = m.M
	v4363 = m.ExcPending
	if v4363 != 0 {
		goto L4
	} else {
		goto L841
	}
L841:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L842:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v4370 = m.ExcPending
	if v4370 != 0 {
		goto L4
	} else {
		goto L843
	}
L843:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+64)) = v386 + int32(2368)
	F_errmsg(m, int32(_a_F_sendDir_65), v386-int32(-64))
	mBase = m.M
	v4378 = m.ExcPending
	if v4378 != 0 {
		goto L4
	} else {
		goto L844
	}
L844:
	;
	F_errfinish(m, int32(_a_F_sendDir_19), int32(1425), int32(_a_F_sendDir_20))
	mBase = m.M
	v4383 = m.ExcPending
	if v4383 != 0 {
		goto L4
	} else {
		goto L845
	}
L845:
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
	var v13 float64
	_ = v13
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
	var v29 float64
	_ = v29
	var v33 float64
	_ = v33
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v82 int32
	_ = v82
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = float64(1e+100)
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
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v36 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0
	v42 = v11 + int32(16)
	if v35 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	return
L3:
	;
	v21 = base.F64_mul(v14, v19)
	if base.F64_gt(v21, float64(1e+100)) != 0 {
		v33 = v13
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v21)&int64(9223372036854775807)) {
		v33 = v13
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v29 = float64(1)
	if base.F64_le(v21, v29) != 0 {
		v33 = v29
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v33 = base.F64_nearest(v21)
	goto L1
L7:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+192)) = v77
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+200)) = v79
	F_set_rel_width(m, l0, l1)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L14
	}
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v45 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v52 = v3
	goto L10
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v52<<(uint(int32(2))%32))))
	v63 = F_cost_qual_eval_walker(m, v60, v11+int32(8))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L12
	}
L11:
	;
	goto L7
L12:
	;
	v66 = v52 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v66 < v67 {
		v52 = v66
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
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
	var v36 float64
	_ = v36
	var v40 float64
	_ = v40
	var v44 int32
	_ = v44
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
		v40 = l2
	} else {
		v26 = *(*float64)(unsafe.Add(mBase, _c_F_set_cte_size_estimates[0]))
		v27 = base.F64_mul(l2, v26)
		v28 = float64(1e+100)
		if base.F64_gt(v27, v28) != 0 {
			v40 = v28
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v27)&int64(9223372036854775807)) {
				v40 = v28
			} else {
				v36 = float64(1)
				if base.F64_le(v27, v36) != 0 {
					v40 = v36
				} else {
					v40 = base.F64_nearest(v27)
				}
			}
		}
	}
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v40
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
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
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	v3 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v74
	return
L2:
	;
	v17 = v3
	v20 = v3
	goto L7
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if int32(0) < v11 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v74 = v3
	goto L1
L6:
	;
	goto L5
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v20<<(uint(int32(2))%32))))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v29 != int32(7) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v74 = v65
	goto L1
L9:
	;
	v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+8)))
	v34 = F_exprType(m, v28)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v62 = v27
	goto L11
L11:
	;
	v65 = F_lappend(m, v17, v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L22
	}
L12:
	;
	return
L13:
	;
	v36 = F_exprTypmod(m, v28)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v38 = F_exprCollation(m, v28)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v41 = F_makeVar(m, int32(-2), v33, v34, v36, v38, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v43 != int32(6) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+40)) = uint16(v57)
	v59 = F_flatCopyTargetEntry(m, v27)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L21
	}
L18:
	;
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v53
	v57 = v53
	goto L17
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	if v46 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = l1 + v46
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+40)))
	v57 = v51
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v41
	v62 = v59
	goto L11
L22:
	;
	v68 = v20 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v68 < v69 {
		v17 = v65
		v20 = v68
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
	var v32 int32
	_ = v32
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	if v48 == int32(0) {
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
	v32 = l0 + int32(56)
	goto L11
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
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
		v32 = v34
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
	if v48 == v49 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v54 {
	case 0:
		goto L28
	case 1:
		goto L27
	case 2:
		goto L26
	case 3:
		goto L25
	case 4:
		goto L24
	default:
		goto L23
	}
L23:
	;
	v67 = l0 + int32(56)
	goto L34
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v48 == v63 {
		goto L20
	} else {
		goto L33
	}
L25:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v48 != v61 {
		goto L23
	} else {
		goto L32
	}
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v48 != v59 {
		goto L23
	} else {
		goto L31
	}
L27:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v48 != v57 {
		goto L23
	} else {
		goto L30
	}
L28:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v48 != v55 {
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
L31:
	;
	goto L20
L32:
	;
	goto L20
L33:
	;
	goto L23
L34:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v71 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	F_pfree(m, v48)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L18
	} else {
		goto L41
	}
L36:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+40))
	if v48 == v72 {
		goto L20
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L35
L39:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	if v48 != v74 {
		v67 = v71
		goto L34
	} else {
		goto L40
	}
L40:
	;
	goto L20
L41:
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
	var v45 float64
	_ = v45
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
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
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v69 int64
	_ = v69
	var v74 int64
	_ = v74
	var v79 int64
	_ = v79
	var v92 int32
	_ = v92
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
		v45 = base.F64_mul(v10, float64(4.503599627370495e+15))
		if base.F64_lt(base.F64_abs(v45), float64(9.223372036854776e+18)) != 0 {
			v49 = base.I64_trunc_f64_s(v45)
			v51 = v49
		} else {
			v51 = int64(-9223372036854775807 - 1)
		}
		v53 = v51 + int64(4354685564936845354)
		v54 = int64(30)
		v57 = int64(-4658895280553007687)
		v58 = (int64(base.Ui64(v53)>>(uint(v54)%64)) ^ v53) * v57
		v59 = int64(27)
		v62 = int64(-7723592293110705685)
		v63 = (int64(base.Ui64(v58)>>(uint(v59)%64)) ^ v58) * v62
		v64 = int64(31)
		*(*int64)(unsafe.Add(mBase, _c_F_setseed[0])) = int64(base.Ui64(v63)>>(uint(v64)%64)) ^ v63
		v69 = v51 - int64(7046029254386353131)
		v74 = (int64(base.Ui64(v69)>>(uint(v54)%64)) ^ v69) * v57
		v79 = (int64(base.Ui64(v74)>>(uint(v59)%64)) ^ v74) * v62
		*(*int64)(unsafe.Add(mBase, _c_F_setseed[1])) = int64(base.Ui64(v79)>>(uint(v64)%64)) ^ v79
		if v69|v53 == int64(0) {
			*(*int64)(unsafe.Add(mBase, _c_F_setseed[0])) = int64(1442695040888963407)
			*(*int64)(unsafe.Add(mBase, _c_F_setseed[1])) = int64(6364136223846793005)
		} else {
		}
		v92 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_setseed[2])) = uint8(v92)
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
	var v43 int32
	_ = v43
	var v59 int32
	_ = v59
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
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
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_shdepChangeDep[0]))
	F_shdepLockAndCheckObject(m, l3, l4)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L29
	} else {
		goto L30
	}
L2:
	;
	goto L1
L3:
	;
	v87 = int32(0)
	goto L2
L4:
	;
	if base.Ui32(l1-int32(2964)) < base.Ui32(int32(4)) {
		v87 = v19
		goto L2
	} else {
		goto L27
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
		goto L16
	} else {
		goto L17
	}
L8:
	;
	switch l1 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v87 = v19
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
	if base.Ui32(int32(27)) < base.Ui32(v31) {
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
	v87 = v19
	goto L2
L13:
	;
	if int32(1)<<(uint(v31)%32)&int32(226492515) == int32(0) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v87 = v19
	goto L2
L15:
	;
	if base.Ui32(l1-int32(3592)) < base.Ui32(int32(2)) {
		v87 = v19
		goto L2
	} else {
		goto L25
	}
L16:
	;
	v43 = l1 - int32(_a_F_shdepChangeDep_1)
	if base.Ui32(int32(9)) < base.Ui32(v43) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	switch l1 - int32(_a_F_shdepChangeDep_2) {
	case 0, 1, 2, 3, 4, 59, 60:
		v87 = v19
		goto L2
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L3
	default:
		goto L21
	}
L19:
	;
	if int32(1)<<(uint(v43)%32)&int32(963) == int32(0) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v87 = v19
	goto L2
L21:
	;
	if base.Ui32(l1-int32(_a_F_shdepChangeDep_3)) < base.Ui32(int32(3)) {
		v87 = v19
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v59 = l1 - int32(_a_F_shdepChangeDep_4)
	if base.Ui32(int32(15)) < base.Ui32(v59) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	if int32(1)<<(uint(v59)%32)&int32(_a_F_shdepChangeDep_5) != 0 {
		v87 = v19
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L3
L25:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l1-int32(4060)) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v87 = v19
	goto L2
L27:
	;
	if base.Ui32(l1-int32(2846)) < base.Ui32(int32(2)) {
		v87 = v19
		goto L2
	} else {
		goto L28
	}
L28:
	;
	goto L3
L29:
	;
	return
L30:
	;
	if v87 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v98 = int32(0)
	goto L33
L32:
	;
	v98 = v89
	goto L33
L33:
	;
	F_ScanKeyInit(m, v15-int32(-64), int32(1), int32(3), int32(184), v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	F_ScanKeyInit(m, v15+int32(112), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v110 = int32(3)
	F_ScanKeyInit(m, v15+int32(160), v110, v110, int32(184), l2)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	F_ScanKeyInit(m, v15+int32(208), int32(4), int32(3), int32(65), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	v129 = F_systable_beginscan(m, l0, int32(1232), int32(1), int32(0), int32(4), v15-int32(-64))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	v138 = int32(0)
	goto L40
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L29
	} else {
		goto L69
	}
L40:
	;
	v143 = F_systable_getnext(m, v129)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L29
	} else {
		goto L42
	}
L41:
	;
	F_systable_endscan(m, v129)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L29
	} else {
		goto L49
	}
L42:
	;
	if v143 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+22)))
	v148 = int32(*(*int8)(unsafe.Add(mBase, uint32(v145+v146)+24)))
	if l5 != v148 {
		goto L40
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	goto L41
L46:
	;
	if v138 != 0 {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v150 = F_heap_copytuple(m, v143)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L29
	} else {
		goto L48
	}
L48:
	;
	v138 = v150
	goto L40
L49:
	;
	v154 = int32(0)
	if l3 == int32(2613) {
		v167 = v154
		goto L53
	} else {
		goto L54
	}
L50:
	;
	m.G0 = v15 + int32(256)
	return
L51:
	;
	F_pfree(m, v206)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L29
	} else {
		goto L68
	}
L52:
	;
	if v167 != 0 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L52
L54:
	;
	if base.Ui32(int32(_a_F_shdepChangeDep_6)) < base.Ui32(l4) {
		v167 = v154
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v167 = (base.B2i32(l3 != int32(2615)) | base.B2i32(l4 != int32(2200))) & base.B2i32(l3 != int32(1262))
	goto L53
L56:
	;
	if v138 == int32(0) {
		goto L50
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v138 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	F_CatalogTupleDelete(m, l0, v138+int32(4))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L29
	} else {
		goto L60
	}
L60:
	;
	v206 = v138
	goto L51
L61:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+22)))
	v176 = v174 + v175
	*(*int32)(unsafe.Add(mBase, uint32(v176)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v176)+16)) = l3
	F_CatalogTupleUpdate(m, l0, v138+int32(4), v138)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L29
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l3
	v186 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v15)+27)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v186
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v200 = F_heap_form_tuple(m, v195, v15+int32(32), v15+int32(24))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L29
	} else {
		goto L65
	}
L64:
	;
	v206 = v138
	goto L51
L65:
	;
	F_CatalogTupleInsert(m, l0, v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L29
	} else {
		goto L66
	}
L66:
	;
	if v200 == int32(0) {
		goto L50
	} else {
		goto L67
	}
L67:
	;
	v206 = v200
	goto L51
L68:
	;
	goto L50
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	F_errmsg_internal(m, int32(_a_F_shdepChangeDep_7), v15)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L29
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_shdepChangeDep_8), int32(255), int32(_a_F_shdepChangeDep_9))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L29
	} else {
		goto L71
	}
L71:
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
func F_shim_write(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	v3 = m.Env.Pgmem_send(m, l0, l1)
	return v3
}
func F_shmem_exit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_shmem_exit[0])) = uint8(v9)
	F_LWLockReleaseAll(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v19
	F_errmsg_internal(m, int32(_a_F_shmem_exit_0), v6+int32(16))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v31 = int32(_a_F_shmem_exit_1)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[1]))
	v35 = v33 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[1])) = v35
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	F_errfinish(m, int32(_a_F_shmem_exit_2), int32(248), int32(_a_F_shmem_exit_3))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v40 = v35
	goto L12
L10:
	;
	goto L11
L11:
	;
	v64 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[1])) = v64
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[2]))
	if v67 == v64 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v43 = v40 << (uint(int32(3)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_shmem_exit[3])))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_shmem_exit[4])))
	m.T0[v49].(func(*base.Module, int32, int32))(m, l0, v46)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	v52 = int32(_a_F_shmem_exit_1)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[1]))
	v56 = v54 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[1])) = v56
	if int32(0) <= v56 {
		v40 = v56
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v88 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L24
	}
L17:
	;
	if v67 == int32(_a_F_shmem_exit_4) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v73 = v67
	goto L19
L19:
	;
	F_dsm_detach(m, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L16
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[2]))
	if v78 == int32(0) {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	if v78 != int32(_a_F_shmem_exit_4) {
		v73 = v78
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	if v88 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v92
	F_errmsg_internal(m, int32(_a_F_shmem_exit_5), v6)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v102 = int32(_a_F_shmem_exit_6)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[5]))
	v106 = v104 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[5])) = v106
	if int32(0) <= v106 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	F_errfinish(m, int32(_a_F_shmem_exit_2), int32(281), int32(_a_F_shmem_exit_3))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v111 = v106
	goto L33
L31:
	;
	goto L32
L32:
	;
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_shmem_exit[0])) = uint8(v135)
	*(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[5])) = v135
	m.G0 = v6 + int32(32)
	return
L33:
	;
	v114 = v111 << (uint(int32(3)) % 32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_shmem_exit[6])))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_shmem_exit[7])))
	m.T0[v120].(func(*base.Module, int32, int32))(m, l0, v117)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	v123 = int32(_a_F_shmem_exit_6)
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[5]))
	v127 = v125 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_shmem_exit[5])) = v127
	if int32(0) <= v127 {
		v111 = v127
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
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
func F_sigprocmask(m *base.Module, l0 int32, l1 int32) {
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L5
L2:
	;
	goto L3
L3:
	;
	goto L9
L4:
	;
	goto L3
L5:
	;
	v10 = F__emscripten_memcpy_bulkmem(m, l1, int32(_a_F_sigprocmask_0), int32(128))
	mBase = m.M
	goto L7
L7:
	;
	goto L4
L8:
	;
	goto L14
L9:
	;
	v14 = F__emscripten_memcpy_bulkmem(m, int32(_a_F_sigprocmask_0), l0, int32(128))
	mBase = m.M
	goto L11
L11:
	;
	goto L8
L12:
	;
	goto L18
L14:
	;
	goto L15
L15:
	;
	v37 = int32(_a_F_sigprocmask_0)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_sigprocmask[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_sigprocmask[0])) = v38 & base.I32_rotl(int32(-2), int32(8))
	goto L12
L16:
	;
	v74 = int32(0)
	goto L20
L18:
	;
	goto L19
L19:
	;
	v65 = int32(_a_F_sigprocmask_0)
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_sigprocmask[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_sigprocmask[0])) = v66 & base.I32_rotl(int32(-2), int32(18))
	goto L16
L20:
	;
	v79 = v74 - int32(1)
	if base.Ui32(v79) <= base.Ui32(int32(63)) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	m.G0 = v6 + int32(128)
	goto L40
L22:
	;
	v140 = v74 + int32(1)
	if v140 != int32(65) {
		v74 = v140
		goto L20
	} else {
		goto L39
	}
L23:
	;
	if v91 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v79)>>(uint(int32(3))%32))&int32(536870908))+uint32(_c_F_sigprocmask[1])))
	v91 = int32(base.Ui32(v87)>>(uint(v79)%32)) & int32(1)
	goto L26
L25:
	;
	v91 = int32(0)
	goto L26
L26:
	;
	goto L23
L27:
	;
	v97 = v74 - int32(1)
	if base.Ui32(v97) <= base.Ui32(int32(63)) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v109 != 0 {
		goto L22
	} else {
		goto L32
	}
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v97)>>(uint(int32(3))%32))&int32(536870908))+uint32(_c_F_sigprocmask[0])))
	v109 = int32(base.Ui32(v105)>>(uint(v97)%32)) & int32(1)
	goto L31
L30:
	;
	v109 = int32(0)
	goto L31
L31:
	;
	goto L28
L32:
	;
	v113 = v74 - int32(1)
	if base.B2i32(base.Ui32(v113) <= base.Ui32(int32(63)))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v74-int32(32))) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v137 = F_raise(m, v74)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_sigprocmask[2])) = int32(28)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v129 = int32(base.Ui32(v113)>>(uint(int32(3))%32)) & int32(536870908)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+uint32(_c_F_sigprocmask[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+uint32(_c_F_sigprocmask[1]))) = v131 & base.I32_rotl(int32(-2), v113)
	goto L33
L37:
	;
	return
L38:
	;
	goto L22
L39:
	;
	goto L21
L40:
	;
	return
}
func F_sjis_to_utf8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(35), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_LocalToUtf(m, v6, v10, v5, int32(_a_F_sjis_to_utf8_0), v18, v18, v18, int32(35), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
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
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7
	v11 = v6
	goto L1
L1:
	;
	if base.Ui32(v11) <= base.Ui32(v9) {
		v151 = v9
		v153 = v11
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if base.Ui32(v151) < base.Ui32(v153) {
		goto L52
	} else {
		goto L53
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
	v151 = v147
	v153 = v142
	goto L3
L7:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v141 == int32(0) {
		v151 = v143
		v153 = v142
		goto L3
	} else {
		goto L49
	}
L8:
	;
	if base.Ui32(int32(255)) < base.Ui32(v19) {
		v151 = v15
		v153 = v17
		goto L3
	} else {
		goto L47
	}
L9:
	;
	if v19 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L10:
	;
	if base.Ui32(int32(127)) < base.Ui32(v19) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if base.Ui32(int32(127)) < base.Ui32(v19) {
		v151 = v15
		v153 = v17
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_skip[1]))))
	v141 = int32(base.Ui32(v28) >> (uint(int32(7)) % 32))
	goto L7
L13:
	;
	v141 = v79
	goto L7
L14:
	;
	v39 = int32(0)
	v40 = int32(10)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19<<(uint(int32(1))%32))+uint32(_c_F_skip[2]))))
	v79 = int32(base.Ui32(v69&int32(32)) >> (uint(int32(5)) % 32))
	goto L13
L17:
	;
	v45 = base.I32_div_s(v39+v40, int32(2))
	v47 = v45 << (uint(int32(3)) % 32)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_c_F_skip[3])))
	if base.Ui32(v50) < base.Ui32(v19) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v79 = int32(0)
	goto L13
L19:
	;
	if v61 <= v62 {
		v39 = v61
		v40 = v62
		goto L17
	} else {
		goto L26
	}
L20:
	;
	v61 = v45 + int32(1)
	v62 = v40
	goto L19
L21:
	;
	goto L22
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_c_F_skip[4])))
	if base.Ui32(v56) <= base.Ui32(v19) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v79 = int32(1)
	goto L13
L24:
	;
	goto L25
L25:
	;
	v61 = v39
	v62 = v45 - int32(1)
	goto L19
L26:
	;
	goto L18
L27:
	;
	v141 = v126
	goto L7
L28:
	;
	v126 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	if v19 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v126 = base.B2i32(v119 != int32(0))
	goto L27
L32:
	;
	v92 = int32(_a_F_skip_0)
	goto L35
L33:
	;
	goto L34
L34:
	;
	v102 = int32(_a_F_skip_0)
	goto L44
L35:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v94 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v94 != 0 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	if v19 != v94 {
		v92 = v92 + int32(4)
		goto L35
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	goto L36
L40:
	;
	goto L39
L41:
	;
	v100 = v92
	goto L43
L42:
	;
	v100 = int32(0)
	goto L43
L43:
	;
	v119 = v100
	goto L31
L44:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v108 != 0 {
		v102 = v102 + int32(4)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v109 = int32(_a_F_skip_0)
	v119 = (v102-v109)&int32(-4) + v109
	goto L31
L46:
	;
	goto L45
L47:
	;
	goto L48
L48:
	;
	v141 = base.B2i32(base.B2i32(v19 == int32(32))|base.B2i32(base.Ui32(v19-int32(9)) < base.Ui32(int32(5))) != int32(0))
	goto L7
L49:
	;
	v147 = v143 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v147
	if base.Ui32(v147) < base.Ui32(v142) {
		v15 = v147
		v17 = v142
		goto L5
	} else {
		goto L50
	}
L50:
	;
	goto L6
L51:
	;
	v167 = v151
	goto L59
L52:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	if v156 == int32(35) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v151 != v7 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v161 | int32(128)
	goto L58
L57:
	;
	goto L58
L58:
	;
	return
L59:
	;
	v172 = v167 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v172
	if base.Ui32(v153) <= base.Ui32(v172) {
		v9 = v172
		v11 = v153
		goto L1
	} else {
		goto L61
	}
L60:
	;
	v9 = v172
	v11 = v153
	goto L1
L61:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v175 != int32(10) {
		v167 = v172
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrnblocks[0])))
	if v5 == int32(1) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+20))
		if v11 != int32(-1) {
			v41 = v11
			return v41
		} else {
			v15 = int32(_a_F_smgrnblocks_0)
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1])) = v17 + int32(1)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v24*int32(80))+uint32(_c_F_smgrnblocks[2])))
			v30 = m.T0[v29].(func(*base.Module, int32, int32) int32)(m, l0, l1)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+20)) = v30
				v35 = int32(_a_F_smgrnblocks_0)
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1])) = v37 - int32(1)
				v41 = v30
				return v41
			}
		}
	} else {
		v15 = int32(_a_F_smgrnblocks_0)
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1])) = v17 + int32(1)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v24*int32(80))+uint32(_c_F_smgrnblocks[2])))
		v30 = m.T0[v29].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+20)) = v30
			v35 = int32(_a_F_smgrnblocks_0)
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_smgrnblocks[1])) = v37 - int32(1)
			v41 = v30
			return v41
		}
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v102 int64
	_ = v102
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v115 int64
	_ = v115
	var v120 int64
	_ = v120
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v167 int64
	_ = v167
	var v173 int64
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int64
	_ = v179
	var v190 int64
	_ = v190
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
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
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v298 int64
	_ = v298
	var v303 int64
	_ = v303
	var v307 int64
	_ = v307
	var v310 int64
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int64
	_ = v318
	var v321 int32
	_ = v321
	var v323 int64
	_ = v323
	var v328 int64
	_ = v328
	var v332 int32
	_ = v332
	var v334 int64
	_ = v334
	var v341 int64
	_ = v341
	var v346 int64
	_ = v346
	var v350 int32
	_ = v350
	var v352 int64
	_ = v352
	var v356 int32
	_ = v356
	var v357 int64
	_ = v357
	var v362 int32
	_ = v362
	var v363 int64
	_ = v363
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
	var v380 int32
	_ = v380
	var v381 int64
	_ = v381
	var v387 int32
	_ = v387
	var v388 int64
	_ = v388
	var v393 int64
	_ = v393
	var v399 int64
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int64
	_ = v405
	var v416 int64
	_ = v416
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v444 int64
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int64
	_ = v453
	var v455 int64
	_ = v455
	var v458 int64
	_ = v458
	var v459 int64
	_ = v459
	var v469 int64
	_ = v469
	var v474 int32
	_ = v474
	var v475 int64
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int64
	_ = v484
	var v494 int64
	_ = v494
	var v502 int32
	_ = v502
	var v509 float64
	_ = v509
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int64
	_ = v545
	var v548 int32
	_ = v548
	var v555 int64
	_ = v555
	var v560 int64
	_ = v560
	var v564 int64
	_ = v564
	var v567 int64
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v575 int64
	_ = v575
	var v578 int32
	_ = v578
	var v580 int64
	_ = v580
	var v585 int64
	_ = v585
	var v589 int32
	_ = v589
	var v591 int64
	_ = v591
	var v598 int64
	_ = v598
	var v603 int64
	_ = v603
	var v607 int32
	_ = v607
	var v609 int64
	_ = v609
	var v613 int32
	_ = v613
	var v614 int64
	_ = v614
	var v619 int32
	_ = v619
	var v620 int64
	_ = v620
	var v624 int32
	_ = v624
	var v625 int64
	_ = v625
	var v626 int64
	_ = v626
	var v631 int32
	_ = v631
	var v632 int64
	_ = v632
	var v637 int32
	_ = v637
	var v638 int64
	_ = v638
	var v644 int32
	_ = v644
	var v645 int64
	_ = v645
	var v650 int64
	_ = v650
	var v656 int64
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int64
	_ = v662
	var v673 int64
	_ = v673
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v697 int32
	_ = v697
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int64
	_ = v712
	var v720 int64
	_ = v720
	var v725 int64
	_ = v725
	var v729 int64
	_ = v729
	var v732 int64
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v740 int64
	_ = v740
	var v743 int32
	_ = v743
	var v745 int64
	_ = v745
	var v750 int64
	_ = v750
	var v754 int32
	_ = v754
	var v756 int64
	_ = v756
	var v763 int64
	_ = v763
	var v768 int64
	_ = v768
	var v772 int32
	_ = v772
	var v774 int64
	_ = v774
	var v778 int32
	_ = v778
	var v779 int64
	_ = v779
	var v784 int32
	_ = v784
	var v785 int64
	_ = v785
	var v789 int32
	_ = v789
	var v790 int64
	_ = v790
	var v791 int64
	_ = v791
	var v796 int32
	_ = v796
	var v797 int64
	_ = v797
	var v802 int32
	_ = v802
	var v803 int64
	_ = v803
	var v809 int32
	_ = v809
	var v810 int64
	_ = v810
	var v815 int64
	_ = v815
	var v821 int64
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int64
	_ = v827
	var v838 int64
	_ = v838
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int64
	_ = v871
	var v873 int64
	_ = v873
	var v875 int64
	_ = v875
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v922 int32
	_ = v922
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1026 int64
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1036 int64
	_ = v1036
	var v1041 int64
	_ = v1041
	var v1045 int64
	_ = v1045
	var v1048 int64
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1056 int64
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1061 int64
	_ = v1061
	var v1066 int64
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1072 int64
	_ = v1072
	var v1079 int64
	_ = v1079
	var v1084 int64
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1090 int64
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1095 int64
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int64
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1106 int64
	_ = v1106
	var v1107 int64
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1113 int64
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1119 int64
	_ = v1119
	var v1125 int32
	_ = v1125
	var v1126 int64
	_ = v1126
	var v1131 int64
	_ = v1131
	var v1137 int64
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1143 int64
	_ = v1143
	var v1154 int64
	_ = v1154
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1185 int32
	_ = v1185
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1200 int64
	_ = v1200
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1255 int64
	_ = v1255
	var v1257 int64
	_ = v1257
	var v1259 int64
	_ = v1259
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1291 int64
	_ = v1291
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1340 int32
	_ = v1340
	var v1359 int32
	_ = v1359
	var v1374 int64
	_ = v1374
	var v1381 int32
	_ = v1381
	var v1395 int32
	_ = v1395
	var v1400 int32
	_ = v1400
	var v1419 int32
	_ = v1419
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_insert[0]))
	if v23 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(48)
	return v1419
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = l0
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v20)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v58
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_insert[1]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v65 = v20 + int32(24)
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
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v26 != l1 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v32 == int32(0) {
		v51 = v31
		v52 = v32
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v52-v51 == int32(0) {
		v1419 = v23
		goto L1
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	if v31 != v32 {
		v51 = v31
		v52 = v32
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v36 = v28
	v37 = l0
	goto L9
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v41 == int32(0) {
		v51 = v40
		v52 = v41
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v51 = v40
	v52 = v41
	goto L6
L11:
	;
	v44 = int32(1)
	if v40 == v41 {
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
	goto L2
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spcache_insert[0])) = v1400
	v1419 = v1400
	goto L1
L15:
	;
	v198 = v63 & base.I32_wrap_i64(int64(base.Ui64(v190)>>(uint(int64(47))%64))^v190-int64(base.Ui64(v190)>>(uint(int64(32))%64)))
	v201 = v62 + v198*int32(24)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+22)))
	if v202 != 0 {
		goto L45
	} else {
		goto L46
	}
L16:
	;
	v190 = (base.I64_extend_i32_s(v176-v85) + int64(base.Ui64(v179)>>(uint(int64(23))%64)) ^ v179) * int64(2388976653695081527)
	goto L15
L17:
	;
	v176 = v85
	v179 = v84
	goto L16
L18:
	;
	goto L19
L19:
	;
	v89 = v85
	v92 = v84
	v95 = v86
	goto L20
L20:
	;
	v97 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v97 == int64(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v176 = v174
	v179 = v173
	goto L16
L22:
	;
	v167 = (v162 ^ int64(base.Ui64(v162)>>(uint(int64(23))%64))) * int64(2388976653695081527)
	v173 = (v92 ^ int64(base.Ui64(v167)>>(uint(int64(47))%64)) ^ v167) * int64(-8645972361240307355)
	v174 = v89 + v161
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v175 != 0 {
		v89 = v174
		v92 = v173
		v95 = v175
		goto L20
	} else {
		goto L44
	}
L23:
	;
	v154 = int32(1)
	v155 = int64(0)
	goto L25
L24:
	;
	v102 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+2)))
	if v102 == int64(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v161 = v154
	v162 = v155 | base.I64_extend8_s(base.I64_extend_i32_u(v95))
	goto L22
L26:
	;
	v148 = int32(2)
	v149 = int64(0)
	goto L28
L27:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+3)))
	if v106 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v154 = v148
	v155 = v149 | v97<<(uint(int64(8))%64)
	goto L25
L29:
	;
	v108 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v108 == int64(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v148 = int32(3)
	v149 = v102 << (uint(int64(16)) % 64)
	goto L28
L32:
	;
	v141 = int32(4)
	v142 = int64(0)
	goto L34
L33:
	;
	v115 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+5)))
	if v115 == int64(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v143 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v89))))
	v161 = v141
	v162 = v142 | v143
	goto L22
L35:
	;
	v136 = int32(5)
	v137 = int64(0)
	goto L37
L36:
	;
	v120 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+6)))
	if v120 == int64(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v141 = v136
	v142 = v108<<(uint(int64(32))%64) | v137
	goto L34
L38:
	;
	v130 = int32(6)
	v131 = int64(0)
	goto L40
L39:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+7)))
	if v124 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v136 = v130
	v137 = v131 | v115<<(uint(int64(40))%64)
	goto L37
L41:
	;
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
	v161 = int32(8)
	v162 = v126
	goto L22
L42:
	;
	goto L43
L43:
	;
	v130 = int32(7)
	v131 = v120 << (uint(int64(48)) % 64)
	goto L40
L44:
	;
	goto L21
L45:
	;
	v205 = v201
	v209 = v198
	goto L48
L46:
	;
	goto L47
L47:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_insert[2]))
	v276 = F_MemoryContextStrdup(m, v275, l0)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L63
	} else {
		goto L64
	}
L48:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	if l1 == v220 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L47
L50:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v226 == int32(0) {
		v245 = v225
		v246 = v226
		goto L54
	} else {
		goto L55
	}
L51:
	;
	goto L52
L52:
	;
	v252 = (v209 + int32(1)) & v63
	v255 = v62 + v252*int32(24)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+22)))
	if v256 != 0 {
		v205 = v255
		v209 = v252
		goto L48
	} else {
		goto L62
	}
L53:
	;
	if v246-v245 == int32(0) {
		v1400 = v205
		goto L14
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	if v225 != v226 {
		v245 = v225
		v246 = v226
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v230 = v222
	v231 = l0
	goto L57
L57:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if v235 == int32(0) {
		v245 = v234
		v246 = v235
		goto L54
	} else {
		goto L59
	}
L58:
	;
	v245 = v234
	v246 = v235
	goto L54
L59:
	;
	v238 = int32(1)
	if v234 == v235 {
		v230 = v230 + v238
		v231 = v231 + v238
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	goto L52
L62:
	;
	goto L49
L63:
	;
	return int32(0)
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v276
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_insert[1]))
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v20)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v283
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v283
	v288 = base.I32_wrap_i64(int64(base.Ui64(v283) >> (uint(int64(32)) % 64)))
	v289 = base.I32_wrap_i64(v283)
	v291 = v20 + int32(16)
	v298 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v291)+4)))
	v303 = (int64(base.Ui64(v298)>>(uint(int64(23))%64)) ^ v298) * int64(2388976653695081527)
	v307 = int64(-8645972361240307355)
	v310 = (v303 ^ int64(base.Ui64(v303)>>(uint(int64(47))%64)) ^ v307) * v307
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if v312 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v282)+16))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v428 = v424
	v429 = v425
	goto L95
L66:
	;
	v416 = (base.I64_extend_i32_s(v402-v311) + int64(base.Ui64(v405)>>(uint(int64(23))%64)) ^ v405) * int64(2388976653695081527)
	goto L65
L67:
	;
	v402 = v311
	v405 = v310
	goto L66
L68:
	;
	goto L69
L69:
	;
	v315 = v311
	v318 = v310
	v321 = v312
	goto L70
L70:
	;
	v323 = int64(*(*int8)(unsafe.Add(mBase, uint32(v315)+1)))
	if v323 == int64(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v402 = v400
	v405 = v399
	goto L66
L72:
	;
	v393 = (v388 ^ int64(base.Ui64(v388)>>(uint(int64(23))%64))) * int64(2388976653695081527)
	v399 = (v318 ^ int64(base.Ui64(v393)>>(uint(int64(47))%64)) ^ v393) * int64(-8645972361240307355)
	v400 = v315 + v387
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	if v401 != 0 {
		v315 = v400
		v318 = v399
		v321 = v401
		goto L70
	} else {
		goto L94
	}
L73:
	;
	v380 = int32(1)
	v381 = int64(0)
	goto L75
L74:
	;
	v328 = int64(*(*int8)(unsafe.Add(mBase, uint32(v315)+2)))
	if v328 == int64(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v387 = v380
	v388 = v381 | base.I64_extend8_s(base.I64_extend_i32_u(v321))
	goto L72
L76:
	;
	v374 = int32(2)
	v375 = int64(0)
	goto L78
L77:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+3)))
	if v332 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v380 = v374
	v381 = v375 | v323<<(uint(int64(8))%64)
	goto L75
L79:
	;
	v334 = int64(*(*int8)(unsafe.Add(mBase, uint32(v315)+4)))
	if v334 == int64(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	v374 = int32(3)
	v375 = v328 << (uint(int64(16)) % 64)
	goto L78
L82:
	;
	v367 = int32(4)
	v368 = int64(0)
	goto L84
L83:
	;
	v341 = int64(*(*int8)(unsafe.Add(mBase, uint32(v315)+5)))
	if v341 == int64(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v369 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v315))))
	v387 = v367
	v388 = v368 | v369
	goto L72
L85:
	;
	v362 = int32(5)
	v363 = int64(0)
	goto L87
L86:
	;
	v346 = int64(*(*int8)(unsafe.Add(mBase, uint32(v315)+6)))
	if v346 == int64(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v367 = v362
	v368 = v334<<(uint(int64(32))%64) | v363
	goto L84
L88:
	;
	v356 = int32(6)
	v357 = int64(0)
	goto L90
L89:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+7)))
	if v350 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v362 = v356
	v363 = v357 | v341<<(uint(int64(40))%64)
	goto L87
L91:
	;
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
	v387 = int32(8)
	v388 = v352
	goto L72
L92:
	;
	goto L93
L93:
	;
	v356 = int32(7)
	v357 = v346 << (uint(int64(48)) % 64)
	goto L90
L94:
	;
	goto L71
L95:
	;
	if base.Ui32(v428) <= base.Ui32(v429) {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	v1395 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+16)) = v1395
	v428 = v1395
	v429 = v1381
	goto L95
L98:
	;
	v1374 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1359)+8)) = v1374
	*(*int64)(unsafe.Add(mBase, uint32(v1359)+14)) = v1374
	v1400 = v1359
	goto L14
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1340)+4)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v1340))) = v289
	v1359 = v1340
	goto L98
L100:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v1333 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+8)) = v1332 + v1333
	*(*uint8)(unsafe.Add(mBase, uint32(v1317)+22)) = uint8(v1333)
	v1340 = v1317
	goto L99
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L63
	} else {
		goto L289
	}
L102:
	;
	v444 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
	if v444 == int64(4294967296) {
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v969 = int32(0)
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v282)+20))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v972 = v971 & base.I32_wrap_i64(int64(base.Ui64(v416)>>(uint(int64(47))%64))^v416-int64(base.Ui64(v416)>>(uint(int64(32))%64)))
	v975 = v970 + v972*int32(24)
	v976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v975)+22)))
	if v976 == v969 {
		v1317 = v975
		goto L100
	} else {
		goto L217
	}
L105:
	;
	v447 = int32(0)
	v449 = m.G0
	v451 = v449 - int32(16)
	m.G0 = v451
	v453 = int64(2)
	v455 = v444 << (uint(int64(1)) % 64)
	if base.Ui64(v455) <= base.Ui64(v453) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L104
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L63
	} else {
		goto L214
	}
L108:
	;
	v458 = v453
	goto L110
L109:
	;
	v458 = v455
	goto L110
L110:
	;
	v459 = int64(1)
	if v458&(v458-v459) == int64(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v469 = v458
	goto L113
L112:
	;
	v469 = v459 << (uint(int64(64)-base.I64_clz(v458)) % 64)
	goto L113
L113:
	;
	if base.Ui64(v469*int64(24)) < base.Ui64(int64(2147483647)) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v282)+20))
	v475 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v282)+24))
	v481 = F_MemoryContextAllocExtended(m, v476, base.I32_wrap_i64(v469)*int32(24), int32(5))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L63
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L63
	} else {
		goto L211
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+20)) = v481
	v484 = int64(1)
	if v469&(v469-v484) == int64(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v494 = v469
	goto L120
L119:
	;
	v494 = v484 << (uint(int64(64)-base.I64_clz(v469)) % 64)
	goto L120
L120:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v494*int64(24)) {
		goto L107
	} else {
		goto L121
	}
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v282))) = v494
	v502 = base.I32_wrap_i64(v494) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+12)) = v502
	v509 = base.F64_mul(base.F64_convert_i64_u(v494), float64(0.9))
	if base.F64_lt(v509, float64(4.294967296e+09))&base.F64_ge(v509, float64(0)) != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	if v494 == int64(4294967296) {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v515 = base.I32_trunc_f64_u(v509)
	v517 = v515
	goto L122
L124:
	;
	goto L125
L125:
	;
	v517 = int32(0)
	goto L122
L126:
	;
	v518 = int32(-85899346)
	goto L128
L127:
	;
	v518 = v517
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+16)) = v518
	if v475 != int64(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v523 = v447
	goto L133
L130:
	;
	goto L131
L131:
	;
	F_pfree(m, v474)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L63
	} else {
		goto L210
	}
L132:
	;
	v690 = v688
	v697 = v447
	goto L168
L133:
	;
	v541 = v474 + v523*int32(24)
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+22)))
	if v542 != int32(1) {
		v688 = v523
		goto L132
	} else {
		goto L135
	}
L134:
	;
	v688 = int32(0)
	goto L132
L135:
	;
	v545 = *(*int64)(unsafe.Add(mBase, uint32(v541)))
	*(*int64)(unsafe.Add(mBase, uint32(v451)+8)) = v545
	v548 = v451 + int32(8)
	v555 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v548)+4)))
	v560 = (int64(base.Ui64(v555)>>(uint(int64(23))%64)) ^ v555) * int64(2388976653695081527)
	v564 = int64(-8645972361240307355)
	v567 = (v560 ^ int64(base.Ui64(v560)>>(uint(int64(47))%64)) ^ v564) * v564
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568))))
	if v569 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	if base.I32_wrap_i64(int64(base.Ui64(v673)>>(uint(int64(47))%64))^v673-int64(base.Ui64(v673)>>(uint(int64(32))%64)))&v502 == v523 {
		v688 = v523
		goto L132
	} else {
		goto L166
	}
L137:
	;
	v673 = (base.I64_extend_i32_s(v659-v568) + int64(base.Ui64(v662)>>(uint(int64(23))%64)) ^ v662) * int64(2388976653695081527)
	goto L136
L138:
	;
	v659 = v568
	v662 = v567
	goto L137
L139:
	;
	goto L140
L140:
	;
	v572 = v568
	v575 = v567
	v578 = v569
	goto L141
L141:
	;
	v580 = int64(*(*int8)(unsafe.Add(mBase, uint32(v572)+1)))
	if v580 == int64(0) {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	v659 = v657
	v662 = v656
	goto L137
L143:
	;
	v650 = (v645 ^ int64(base.Ui64(v645)>>(uint(int64(23))%64))) * int64(2388976653695081527)
	v656 = (v575 ^ int64(base.Ui64(v650)>>(uint(int64(47))%64)) ^ v650) * int64(-8645972361240307355)
	v657 = v572 + v644
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v657))))
	if v658 != 0 {
		v572 = v657
		v575 = v656
		v578 = v658
		goto L141
	} else {
		goto L165
	}
L144:
	;
	v637 = int32(1)
	v638 = int64(0)
	goto L146
L145:
	;
	v585 = int64(*(*int8)(unsafe.Add(mBase, uint32(v572)+2)))
	if v585 == int64(0) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v644 = v637
	v645 = v638 | base.I64_extend8_s(base.I64_extend_i32_u(v578))
	goto L143
L147:
	;
	v631 = int32(2)
	v632 = int64(0)
	goto L149
L148:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+3)))
	if v589 != 0 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v637 = v631
	v638 = v632 | v580<<(uint(int64(8))%64)
	goto L146
L150:
	;
	v591 = int64(*(*int8)(unsafe.Add(mBase, uint32(v572)+4)))
	if v591 == int64(0) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	v631 = int32(3)
	v632 = v585 << (uint(int64(16)) % 64)
	goto L149
L153:
	;
	v624 = int32(4)
	v625 = int64(0)
	goto L155
L154:
	;
	v598 = int64(*(*int8)(unsafe.Add(mBase, uint32(v572)+5)))
	if v598 == int64(0) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v626 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v572))))
	v644 = v624
	v645 = v625 | v626
	goto L143
L156:
	;
	v619 = int32(5)
	v620 = int64(0)
	goto L158
L157:
	;
	v603 = int64(*(*int8)(unsafe.Add(mBase, uint32(v572)+6)))
	if v603 == int64(0) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v624 = v619
	v625 = v591<<(uint(int64(32))%64) | v620
	goto L155
L159:
	;
	v613 = int32(6)
	v614 = int64(0)
	goto L161
L160:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+7)))
	if v607 != 0 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v619 = v613
	v620 = v614 | v598<<(uint(int64(40))%64)
	goto L158
L162:
	;
	v609 = *(*int64)(unsafe.Add(mBase, uint32(v572)))
	v644 = int32(8)
	v645 = v609
	goto L143
L163:
	;
	goto L164
L164:
	;
	v613 = int32(7)
	v614 = v603 << (uint(int64(48)) % 64)
	goto L161
L165:
	;
	goto L142
L166:
	;
	v684 = v523 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v684)) < base.Ui64(v475) {
		v523 = v684
		goto L133
	} else {
		goto L167
	}
L167:
	;
	goto L134
L168:
	;
	v708 = v474 + v690*int32(24)
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708)+22)))
	if v709 == int32(1) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L131
L170:
	;
	v712 = *(*int64)(unsafe.Add(mBase, uint32(v708)))
	*(*int64)(unsafe.Add(mBase, uint32(v451))) = v712
	v720 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v451)+4)))
	v725 = (int64(base.Ui64(v720)>>(uint(int64(23))%64)) ^ v720) * int64(2388976653695081527)
	v729 = int64(-8645972361240307355)
	v732 = (v725 ^ int64(base.Ui64(v725)>>(uint(int64(47))%64)) ^ v729) * v729
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733))))
	if v734 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L171:
	;
	goto L172
L172:
	;
	v895 = v690 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v895)) < base.Ui64(v475) {
		goto L206
	} else {
		goto L207
	}
L173:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v847 = base.I32_wrap_i64(int64(base.Ui64(v838)>>(uint(int64(47))%64)) ^ v838 - int64(base.Ui64(v838)>>(uint(int64(32))%64)))
	goto L203
L174:
	;
	v838 = (base.I64_extend_i32_s(v824-v733) + int64(base.Ui64(v827)>>(uint(int64(23))%64)) ^ v827) * int64(2388976653695081527)
	goto L173
L175:
	;
	v824 = v733
	v827 = v732
	goto L174
L176:
	;
	goto L177
L177:
	;
	v737 = v733
	v740 = v732
	v743 = v734
	goto L178
L178:
	;
	v745 = int64(*(*int8)(unsafe.Add(mBase, uint32(v737)+1)))
	if v745 == int64(0) {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v824 = v822
	v827 = v821
	goto L174
L180:
	;
	v815 = (v810 ^ int64(base.Ui64(v810)>>(uint(int64(23))%64))) * int64(2388976653695081527)
	v821 = (v740 ^ int64(base.Ui64(v815)>>(uint(int64(47))%64)) ^ v815) * int64(-8645972361240307355)
	v822 = v737 + v809
	v823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822))))
	if v823 != 0 {
		v737 = v822
		v740 = v821
		v743 = v823
		goto L178
	} else {
		goto L202
	}
L181:
	;
	v802 = int32(1)
	v803 = int64(0)
	goto L183
L182:
	;
	v750 = int64(*(*int8)(unsafe.Add(mBase, uint32(v737)+2)))
	if v750 == int64(0) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v809 = v802
	v810 = v803 | base.I64_extend8_s(base.I64_extend_i32_u(v743))
	goto L180
L184:
	;
	v796 = int32(2)
	v797 = int64(0)
	goto L186
L185:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737)+3)))
	if v754 != 0 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v802 = v796
	v803 = v797 | v745<<(uint(int64(8))%64)
	goto L183
L187:
	;
	v756 = int64(*(*int8)(unsafe.Add(mBase, uint32(v737)+4)))
	if v756 == int64(0) {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L189
L189:
	;
	v796 = int32(3)
	v797 = v750 << (uint(int64(16)) % 64)
	goto L186
L190:
	;
	v789 = int32(4)
	v790 = int64(0)
	goto L192
L191:
	;
	v763 = int64(*(*int8)(unsafe.Add(mBase, uint32(v737)+5)))
	if v763 == int64(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v791 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v737))))
	v809 = v789
	v810 = v790 | v791
	goto L180
L193:
	;
	v784 = int32(5)
	v785 = int64(0)
	goto L195
L194:
	;
	v768 = int64(*(*int8)(unsafe.Add(mBase, uint32(v737)+6)))
	if v768 == int64(0) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v789 = v784
	v790 = v756<<(uint(int64(32))%64) | v785
	goto L192
L196:
	;
	v778 = int32(6)
	v779 = int64(0)
	goto L198
L197:
	;
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737)+7)))
	if v772 != 0 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v784 = v778
	v785 = v779 | v763<<(uint(int64(40))%64)
	goto L195
L199:
	;
	v774 = *(*int64)(unsafe.Add(mBase, uint32(v737)))
	v809 = int32(8)
	v810 = v774
	goto L180
L200:
	;
	goto L201
L201:
	;
	v778 = int32(7)
	v779 = v768 << (uint(int64(48)) % 64)
	goto L198
L202:
	;
	goto L179
L203:
	;
	v864 = v847 & v846
	v869 = v481 + v864*int32(24)
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v869)+22)))
	if v870 != 0 {
		v847 = v864 + int32(1)
		goto L203
	} else {
		goto L205
	}
L204:
	;
	v871 = *(*int64)(unsafe.Add(mBase, uint32(v708)))
	*(*int64)(unsafe.Add(mBase, uint32(v869))) = v871
	v873 = *(*int64)(unsafe.Add(mBase, uint32(v708)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v869)+16)) = v873
	v875 = *(*int64)(unsafe.Add(mBase, uint32(v708)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v869)+8)) = v875
	goto L172
L205:
	;
	goto L204
L206:
	;
	v899 = v895
	goto L208
L207:
	;
	v899 = int32(0)
	goto L208
L208:
	;
	v901 = v697 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v901)) < base.Ui64(v475) {
		v690 = v899
		v697 = v901
		goto L168
	} else {
		goto L209
	}
L209:
	;
	goto L169
L210:
	;
	m.G0 = v451 + int32(16)
	goto L106
L211:
	;
	F_errmsg_internal(m, int32(_a_F_spcache_insert_0), int32(0))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L63
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_spcache_insert_1), int32(327), int32(_a_F_spcache_insert_2))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L63
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_errmsg_internal(m, int32(_a_F_spcache_insert_0), int32(0))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L63
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_spcache_insert_1), int32(327), int32(_a_F_spcache_insert_2))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L63
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	v980 = v969
	v981 = v975
	v985 = v972
	goto L218
L218:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v981)+4))
	if v288 == v996 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v1317 = v1300
	goto L100
L220:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v981)))
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998))))
	if v1002 == int32(0) {
		v1021 = v1001
		v1022 = v1002
		goto L224
	} else {
		goto L225
	}
L221:
	;
	goto L222
L222:
	;
	v1026 = *(*int64)(unsafe.Add(mBase, uint32(v981)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v1026
	v1029 = v20 + int32(8)
	v1036 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1029)+4)))
	v1041 = (int64(base.Ui64(v1036)>>(uint(int64(23))%64)) ^ v1036) * int64(2388976653695081527)
	v1045 = int64(-8645972361240307355)
	v1048 = (v1041 ^ int64(base.Ui64(v1041)>>(uint(int64(47))%64)) ^ v1045) * v1045
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1029)))
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1049))))
	if v1050 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L223:
	;
	if v1022-v1021 == int32(0) {
		v1359 = v981
		goto L98
	} else {
		goto L231
	}
L224:
	;
	goto L223
L225:
	;
	if v1001 != v1002 {
		v1021 = v1001
		v1022 = v1002
		goto L224
	} else {
		goto L226
	}
L226:
	;
	v1006 = v998
	v1007 = v289
	goto L227
L227:
	;
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007)+1)))
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+1)))
	if v1011 == int32(0) {
		v1021 = v1010
		v1022 = v1011
		goto L224
	} else {
		goto L229
	}
L228:
	;
	v1021 = v1010
	v1022 = v1011
	goto L224
L229:
	;
	v1014 = int32(1)
	if v1010 == v1011 {
		v1006 = v1006 + v1014
		v1007 = v1007 + v1014
		goto L227
	} else {
		goto L230
	}
L230:
	;
	goto L228
L231:
	;
	goto L222
L232:
	;
	v1162 = base.I32_wrap_i64(int64(base.Ui64(v1154)>>(uint(int64(47))%64))^v1154-int64(base.Ui64(v1154)>>(uint(int64(32))%64))) & v971
	if base.Ui32(v985) < base.Ui32(v1162) {
		goto L262
	} else {
		goto L263
	}
L233:
	;
	v1154 = (base.I64_extend_i32_s(v1140-v1049) + int64(base.Ui64(v1143)>>(uint(int64(23))%64)) ^ v1143) * int64(2388976653695081527)
	goto L232
L234:
	;
	v1140 = v1049
	v1143 = v1048
	goto L233
L235:
	;
	goto L236
L236:
	;
	v1053 = v1049
	v1056 = v1048
	v1059 = v1050
	goto L237
L237:
	;
	v1061 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1053)+1)))
	if v1061 == int64(0) {
		goto L240
	} else {
		goto L241
	}
L238:
	;
	v1140 = v1138
	v1143 = v1137
	goto L233
L239:
	;
	v1131 = (v1126 ^ int64(base.Ui64(v1126)>>(uint(int64(23))%64))) * int64(2388976653695081527)
	v1137 = (v1056 ^ int64(base.Ui64(v1131)>>(uint(int64(47))%64)) ^ v1131) * int64(-8645972361240307355)
	v1138 = v1053 + v1125
	v1139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138))))
	if v1139 != 0 {
		v1053 = v1138
		v1056 = v1137
		v1059 = v1139
		goto L237
	} else {
		goto L261
	}
L240:
	;
	v1118 = int32(1)
	v1119 = int64(0)
	goto L242
L241:
	;
	v1066 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1053)+2)))
	if v1066 == int64(0) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v1125 = v1118
	v1126 = v1119 | base.I64_extend8_s(base.I64_extend_i32_u(v1059))
	goto L239
L243:
	;
	v1112 = int32(2)
	v1113 = int64(0)
	goto L245
L244:
	;
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053)+3)))
	if v1070 != 0 {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v1118 = v1112
	v1119 = v1113 | v1061<<(uint(int64(8))%64)
	goto L242
L246:
	;
	v1072 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1053)+4)))
	if v1072 == int64(0) {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	goto L248
L248:
	;
	v1112 = int32(3)
	v1113 = v1066 << (uint(int64(16)) % 64)
	goto L245
L249:
	;
	v1105 = int32(4)
	v1106 = int64(0)
	goto L251
L250:
	;
	v1079 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1053)+5)))
	if v1079 == int64(0) {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	v1107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1053))))
	v1125 = v1105
	v1126 = v1106 | v1107
	goto L239
L252:
	;
	v1100 = int32(5)
	v1101 = int64(0)
	goto L254
L253:
	;
	v1084 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1053)+6)))
	if v1084 == int64(0) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1105 = v1100
	v1106 = v1072<<(uint(int64(32))%64) | v1101
	goto L251
L255:
	;
	v1094 = int32(6)
	v1095 = int64(0)
	goto L257
L256:
	;
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053)+7)))
	if v1088 != 0 {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v1100 = v1094
	v1101 = v1095 | v1079<<(uint(int64(40))%64)
	goto L254
L258:
	;
	v1090 = *(*int64)(unsafe.Add(mBase, uint32(v1053)))
	v1125 = int32(8)
	v1126 = v1090
	goto L239
L259:
	;
	goto L260
L260:
	;
	v1094 = int32(7)
	v1095 = v1084 << (uint(int64(48)) % 64)
	goto L257
L261:
	;
	goto L238
L262:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v1166 = v985 + v1164
	goto L264
L263:
	;
	v1166 = v985
	goto L264
L264:
	;
	v1168 = v985 + int32(1)
	if base.Ui32(v1166-v1162) < base.Ui32(v980) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1172 = v971 & v1168
	v1175 = v970 + v1172*int32(24)
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+22)))
	if v1176 != 0 {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	goto L267
L267:
	;
	v1286 = v980 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v1286) {
		goto L284
	} else {
		goto L285
	}
L268:
	;
	v1178 = v1172
	v1185 = int32(0)
	goto L271
L269:
	;
	v1214 = v1172
	v1216 = v1175
	goto L270
L270:
	;
	if v1214 != v985 {
		goto L278
	} else {
		goto L279
	}
L271:
	;
	v1195 = v1185 + int32(1)
	if int32(151) <= v1195 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v1214 = v1208
	v1216 = v1211
	goto L270
L273:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v1200 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1198), base.F64_convert_i64_u(v1200)), float64(0.1)) != 0 {
		v1381 = v1198
		goto L97
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1208 = (v1178 + int32(1)) & v971
	v1211 = v970 + v1208*int32(24)
	v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211)+22)))
	if v1212 != 0 {
		v1178 = v1208
		v1185 = v1195
		goto L271
	} else {
		goto L277
	}
L276:
	;
	goto L275
L277:
	;
	goto L272
L278:
	;
	v1232 = v1214
	v1234 = v1216
	goto L281
L279:
	;
	goto L280
L280:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v1280 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+8)) = v1279 + v1280
	*(*uint8)(unsafe.Add(mBase, uint32(v981)+22)) = uint8(v1280)
	v1340 = v981
	goto L99
L281:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v1251 = v1248 & (v1232 - int32(1))
	v1254 = v970 + v1251*int32(24)
	v1255 = *(*int64)(unsafe.Add(mBase, uint32(v1254)))
	*(*int64)(unsafe.Add(mBase, uint32(v1234))) = v1255
	v1257 = *(*int64)(unsafe.Add(mBase, uint32(v1254)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1234)+16)) = v1257
	v1259 = *(*int64)(unsafe.Add(mBase, uint32(v1254)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1234)+8)) = v1259
	if v1251 != v985 {
		v1232 = v1251
		v1234 = v1254
		goto L281
	} else {
		goto L283
	}
L282:
	;
	goto L280
L283:
	;
	goto L282
L284:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v1291 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1289), base.F64_convert_i64_u(v1291)), float64(0.1)) != 0 {
		v1381 = v1289
		goto L97
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1297 = v971 & v1168
	v1300 = v970 + v1297*int32(24)
	v1301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1300)+22)))
	if v1301 != 0 {
		v980 = v1286
		v981 = v1300
		v985 = v1297
		goto L218
	} else {
		goto L288
	}
L287:
	;
	goto L286
L288:
	;
	goto L219
L289:
	;
	F_errmsg_internal(m, int32(_a_F_spcache_insert_3), int32(0))
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L63
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(_a_F_spcache_insert_1), int32(630), int32(_a_F_spcache_insert_4))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L63
	} else {
		goto L291
	}
L291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_specialcolors(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
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
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v6 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v10 = F_newcolor(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = int32(_a_F_specialcolors_0)
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
			if v15 != 0 {
				v29 = v12
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				v19 = v16 + v10*int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(2)
				*(*int64)(unsafe.Add(mBase, uint32(v19)+12)) = int64(0)
				v24 = int32(_a_F_specialcolors_0)
				*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)) = uint16(v24)
				*(*int64)(unsafe.Add(mBase, uint32(v19))) = int64(4294967296)
				v29 = v10
			}
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)) = uint16(v29)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v32 = F_newcolor(m, v31)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
				if v35 != 0 {
					v49 = v12
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
					v39 = v36 + v32*int32(24)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(2)
					*(*int64)(unsafe.Add(mBase, uint32(v39)+12)) = int64(0)
					v44 = int32(_a_F_specialcolors_0)
					*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)) = uint16(v44)
					*(*int64)(unsafe.Add(mBase, uint32(v39))) = int64(4294967296)
					v49 = v32
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)) = uint16(v49)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v52 = F_newcolor(m, v51)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					v54 = int32(_a_F_specialcolors_0)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
					if v57 != 0 {
						v71 = v54
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
						v61 = v58 + v52*int32(24)
						*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = int32(2)
						*(*int64)(unsafe.Add(mBase, uint32(v61)+12)) = int64(0)
						v66 = int32(_a_F_specialcolors_0)
						*(*uint16)(unsafe.Add(mBase, uint32(v61)+8)) = uint16(v66)
						*(*int64)(unsafe.Add(mBase, uint32(v61))) = int64(4294967296)
						v71 = v52
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)) = uint16(v71)
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v74 = F_newcolor(m, v73)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
						if v77 != 0 {
							v97 = v54
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
							v81 = v78 + v74*int32(24)
							*(*int32)(unsafe.Add(mBase, uint32(v81)+20)) = int32(2)
							*(*int64)(unsafe.Add(mBase, uint32(v81)+12)) = int64(0)
							v86 = int32(_a_F_specialcolors_0)
							*(*uint16)(unsafe.Add(mBase, uint32(v81)+8)) = uint16(v86)
							*(*int64)(unsafe.Add(mBase, uint32(v81))) = int64(4294967296)
							v97 = v74
						}
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)) = uint16(v97)
						return
					}
				}
			}
		}
	} else {
		v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+56)))
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)) = uint16(v90)
		v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+58)))
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)) = uint16(v92)
		v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+60)))
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)) = uint16(v94)
		v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+62)))
		v97 = v96
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)) = uint16(v97)
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
func F_standard_ExplainOneQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	v8 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(288)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
	if v17 == int32(1) {
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[0]))
		v26 = F_AllocSetContextCreateInternal(m, v21, int32(_a_F_standard_ExplainOneQuery_0), int32(0), int32(_a_F_standard_ExplainOneQuery_1), int32(_a_F_standard_ExplainOneQuery_2))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = int32(_a_F_standard_ExplainOneQuery_3)
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[0])) = v26
			v32 = v26
			v33 = v29
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
			if v34 == int32(1) {
				v41 = F__emscripten_memcpy_bulkmem(m, v15+int32(152), int32(_a_F_standard_ExplainOneQuery_4), int32(128))
				mBase = m.M
			} else {
			}
			F___clock_gettime(m, int32(1), v15+int32(24))
			mBase = m.M
			v47 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
			v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+32)))
			v49 = F_pg_plan_query(m, l0, l4, l1, l5)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				v51 = int32(1)
				F___clock_gettime(m, v51, v15+int32(24))
				mBase = m.M
				v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+32)))
				v57 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+280)) = v55 - v48 + (v57-v47)*int64(1000000000)
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
				if v63 == v51 {
					*(*int32)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[0])) = v33
					F_MemoryContextMemConsumed(m, v32, v15+int32(8))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
						if v75 == int32(1) {
							v83 = F__emscripten_memset_bulkmem(m, v15+int32(24), base.I32_extend8_s(int32(0)), int32(128))
							mBase = m.M
							v85 = v15 + int32(24)
							v87 = v15 + int32(152)
							v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							v90 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[1]))
							v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
							*(*int64)(unsafe.Add(mBase, uint32(v85))) = v88 + (v90 - v91)
							v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
							v97 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[2]))
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v95 + (v97 - v98)
							v102 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
							v104 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[3]))
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v102 + (v104 - v105)
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
							v111 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[4]))
							v112 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v109 + (v111 - v112)
							v116 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
							v118 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[5]))
							v119 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v116 + (v118 - v119)
							v123 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
							v125 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[6]))
							v126 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v123 + (v125 - v126)
							v130 = *(*int64)(unsafe.Add(mBase, uint32(v85)+48))
							v132 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[7]))
							v133 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v130 + (v132 - v133)
							v137 = *(*int64)(unsafe.Add(mBase, uint32(v85)+56))
							v139 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[8]))
							v140 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v137 + (v139 - v140)
							v144 = *(*int64)(unsafe.Add(mBase, uint32(v85)+64))
							v146 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[9]))
							v147 = *(*int64)(unsafe.Add(mBase, uint32(v87)+64))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v144 + (v146 - v147)
							v151 = *(*int64)(unsafe.Add(mBase, uint32(v85)+72))
							v153 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[10]))
							v154 = *(*int64)(unsafe.Add(mBase, uint32(v87)+72))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v151 + (v153 - v154)
							v158 = *(*int64)(unsafe.Add(mBase, uint32(v85)+80))
							v160 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[11]))
							v161 = *(*int64)(unsafe.Add(mBase, uint32(v87)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+80)) = v158 + (v160 - v161)
							v165 = *(*int64)(unsafe.Add(mBase, uint32(v85)+88))
							v167 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[12]))
							v168 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+88)) = v165 + (v167 - v168)
							v172 = *(*int64)(unsafe.Add(mBase, uint32(v85)+96))
							v174 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[13]))
							v175 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+96)) = v172 + (v174 - v175)
							v179 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
							v181 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[14]))
							v182 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+104)) = v179 + (v181 - v182)
							v186 = *(*int64)(unsafe.Add(mBase, uint32(v85)+112))
							v188 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[15]))
							v189 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+112)) = v186 + (v188 - v189)
							v193 = *(*int64)(unsafe.Add(mBase, uint32(v85)+120))
							v195 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[16]))
							v196 = *(*int64)(unsafe.Add(mBase, uint32(v87)+120))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+120)) = v193 + (v195 - v196)
							v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
							if v203&int32(1) != 0 {
								v206 = v15 + int32(24)
							} else {
								v206 = int32(0)
							}
							v207 = v206
						} else {
							v207 = int32(0)
						}
						v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
						if v211 != 0 {
							v212 = v15 + int32(8)
						} else {
							v212 = int32(0)
						}
						F_ExplainOnePlan(m, v49, l2, l3, l4, l5, l6, v15+int32(280), v207, v212)
						mBase = m.M
						v214 = m.ExcPending
						if v214 != 0 {
							return
						} else {
							m.G0 = v15 + int32(288)
							return
						}
					}
				} else {
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
					if v75 == int32(1) {
						v83 = F__emscripten_memset_bulkmem(m, v15+int32(24), base.I32_extend8_s(int32(0)), int32(128))
						mBase = m.M
						v85 = v15 + int32(24)
						v87 = v15 + int32(152)
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
						v90 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[1]))
						v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
						*(*int64)(unsafe.Add(mBase, uint32(v85))) = v88 + (v90 - v91)
						v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
						v97 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[2]))
						v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v95 + (v97 - v98)
						v102 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
						v104 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[3]))
						v105 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v102 + (v104 - v105)
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
						v111 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[4]))
						v112 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v109 + (v111 - v112)
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
						v118 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[5]))
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v116 + (v118 - v119)
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
						v125 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[6]))
						v126 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v123 + (v125 - v126)
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v85)+48))
						v132 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[7]))
						v133 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v130 + (v132 - v133)
						v137 = *(*int64)(unsafe.Add(mBase, uint32(v85)+56))
						v139 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[8]))
						v140 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v137 + (v139 - v140)
						v144 = *(*int64)(unsafe.Add(mBase, uint32(v85)+64))
						v146 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[9]))
						v147 = *(*int64)(unsafe.Add(mBase, uint32(v87)+64))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v144 + (v146 - v147)
						v151 = *(*int64)(unsafe.Add(mBase, uint32(v85)+72))
						v153 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[10]))
						v154 = *(*int64)(unsafe.Add(mBase, uint32(v87)+72))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v151 + (v153 - v154)
						v158 = *(*int64)(unsafe.Add(mBase, uint32(v85)+80))
						v160 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[11]))
						v161 = *(*int64)(unsafe.Add(mBase, uint32(v87)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+80)) = v158 + (v160 - v161)
						v165 = *(*int64)(unsafe.Add(mBase, uint32(v85)+88))
						v167 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[12]))
						v168 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+88)) = v165 + (v167 - v168)
						v172 = *(*int64)(unsafe.Add(mBase, uint32(v85)+96))
						v174 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[13]))
						v175 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+96)) = v172 + (v174 - v175)
						v179 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
						v181 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[14]))
						v182 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+104)) = v179 + (v181 - v182)
						v186 = *(*int64)(unsafe.Add(mBase, uint32(v85)+112))
						v188 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[15]))
						v189 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+112)) = v186 + (v188 - v189)
						v193 = *(*int64)(unsafe.Add(mBase, uint32(v85)+120))
						v195 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[16]))
						v196 = *(*int64)(unsafe.Add(mBase, uint32(v87)+120))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+120)) = v193 + (v195 - v196)
						v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
						if v203&int32(1) != 0 {
							v206 = v15 + int32(24)
						} else {
							v206 = int32(0)
						}
						v207 = v206
					} else {
						v207 = int32(0)
					}
					v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
					if v211 != 0 {
						v212 = v15 + int32(8)
					} else {
						v212 = int32(0)
					}
					F_ExplainOnePlan(m, v49, l2, l3, l4, l5, l6, v15+int32(280), v207, v212)
					mBase = m.M
					v214 = m.ExcPending
					if v214 != 0 {
						return
					} else {
						m.G0 = v15 + int32(288)
						return
					}
				}
			}
		}
	} else {
		v32 = v8
		v33 = v8
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
		if v34 == int32(1) {
			v41 = F__emscripten_memcpy_bulkmem(m, v15+int32(152), int32(_a_F_standard_ExplainOneQuery_4), int32(128))
			mBase = m.M
		} else {
		}
		F___clock_gettime(m, int32(1), v15+int32(24))
		mBase = m.M
		v47 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
		v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+32)))
		v49 = F_pg_plan_query(m, l0, l4, l1, l5)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			v51 = int32(1)
			F___clock_gettime(m, v51, v15+int32(24))
			mBase = m.M
			v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+32)))
			v57 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v15)+280)) = v55 - v48 + (v57-v47)*int64(1000000000)
			v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
			if v63 == v51 {
				*(*int32)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[0])) = v33
				F_MemoryContextMemConsumed(m, v32, v15+int32(8))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
					if v75 == int32(1) {
						v83 = F__emscripten_memset_bulkmem(m, v15+int32(24), base.I32_extend8_s(int32(0)), int32(128))
						mBase = m.M
						v85 = v15 + int32(24)
						v87 = v15 + int32(152)
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
						v90 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[1]))
						v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
						*(*int64)(unsafe.Add(mBase, uint32(v85))) = v88 + (v90 - v91)
						v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
						v97 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[2]))
						v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v95 + (v97 - v98)
						v102 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
						v104 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[3]))
						v105 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v102 + (v104 - v105)
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
						v111 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[4]))
						v112 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v109 + (v111 - v112)
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
						v118 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[5]))
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v116 + (v118 - v119)
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
						v125 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[6]))
						v126 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v123 + (v125 - v126)
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v85)+48))
						v132 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[7]))
						v133 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v130 + (v132 - v133)
						v137 = *(*int64)(unsafe.Add(mBase, uint32(v85)+56))
						v139 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[8]))
						v140 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v137 + (v139 - v140)
						v144 = *(*int64)(unsafe.Add(mBase, uint32(v85)+64))
						v146 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[9]))
						v147 = *(*int64)(unsafe.Add(mBase, uint32(v87)+64))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v144 + (v146 - v147)
						v151 = *(*int64)(unsafe.Add(mBase, uint32(v85)+72))
						v153 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[10]))
						v154 = *(*int64)(unsafe.Add(mBase, uint32(v87)+72))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v151 + (v153 - v154)
						v158 = *(*int64)(unsafe.Add(mBase, uint32(v85)+80))
						v160 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[11]))
						v161 = *(*int64)(unsafe.Add(mBase, uint32(v87)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+80)) = v158 + (v160 - v161)
						v165 = *(*int64)(unsafe.Add(mBase, uint32(v85)+88))
						v167 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[12]))
						v168 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+88)) = v165 + (v167 - v168)
						v172 = *(*int64)(unsafe.Add(mBase, uint32(v85)+96))
						v174 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[13]))
						v175 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+96)) = v172 + (v174 - v175)
						v179 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
						v181 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[14]))
						v182 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+104)) = v179 + (v181 - v182)
						v186 = *(*int64)(unsafe.Add(mBase, uint32(v85)+112))
						v188 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[15]))
						v189 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+112)) = v186 + (v188 - v189)
						v193 = *(*int64)(unsafe.Add(mBase, uint32(v85)+120))
						v195 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[16]))
						v196 = *(*int64)(unsafe.Add(mBase, uint32(v87)+120))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+120)) = v193 + (v195 - v196)
						v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
						if v203&int32(1) != 0 {
							v206 = v15 + int32(24)
						} else {
							v206 = int32(0)
						}
						v207 = v206
					} else {
						v207 = int32(0)
					}
					v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
					if v211 != 0 {
						v212 = v15 + int32(8)
					} else {
						v212 = int32(0)
					}
					F_ExplainOnePlan(m, v49, l2, l3, l4, l5, l6, v15+int32(280), v207, v212)
					mBase = m.M
					v214 = m.ExcPending
					if v214 != 0 {
						return
					} else {
						m.G0 = v15 + int32(288)
						return
					}
				}
			} else {
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
				if v75 == int32(1) {
					v83 = F__emscripten_memset_bulkmem(m, v15+int32(24), base.I32_extend8_s(int32(0)), int32(128))
					mBase = m.M
					v85 = v15 + int32(24)
					v87 = v15 + int32(152)
					v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
					v90 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[1]))
					v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
					*(*int64)(unsafe.Add(mBase, uint32(v85))) = v88 + (v90 - v91)
					v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
					v97 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[2]))
					v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v95 + (v97 - v98)
					v102 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
					v104 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[3]))
					v105 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v102 + (v104 - v105)
					v109 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
					v111 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[4]))
					v112 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v109 + (v111 - v112)
					v116 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
					v118 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[5]))
					v119 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v116 + (v118 - v119)
					v123 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
					v125 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[6]))
					v126 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v123 + (v125 - v126)
					v130 = *(*int64)(unsafe.Add(mBase, uint32(v85)+48))
					v132 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[7]))
					v133 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v130 + (v132 - v133)
					v137 = *(*int64)(unsafe.Add(mBase, uint32(v85)+56))
					v139 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[8]))
					v140 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v137 + (v139 - v140)
					v144 = *(*int64)(unsafe.Add(mBase, uint32(v85)+64))
					v146 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[9]))
					v147 = *(*int64)(unsafe.Add(mBase, uint32(v87)+64))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v144 + (v146 - v147)
					v151 = *(*int64)(unsafe.Add(mBase, uint32(v85)+72))
					v153 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[10]))
					v154 = *(*int64)(unsafe.Add(mBase, uint32(v87)+72))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v151 + (v153 - v154)
					v158 = *(*int64)(unsafe.Add(mBase, uint32(v85)+80))
					v160 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[11]))
					v161 = *(*int64)(unsafe.Add(mBase, uint32(v87)+80))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+80)) = v158 + (v160 - v161)
					v165 = *(*int64)(unsafe.Add(mBase, uint32(v85)+88))
					v167 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[12]))
					v168 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+88)) = v165 + (v167 - v168)
					v172 = *(*int64)(unsafe.Add(mBase, uint32(v85)+96))
					v174 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[13]))
					v175 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+96)) = v172 + (v174 - v175)
					v179 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
					v181 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[14]))
					v182 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+104)) = v179 + (v181 - v182)
					v186 = *(*int64)(unsafe.Add(mBase, uint32(v85)+112))
					v188 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[15]))
					v189 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+112)) = v186 + (v188 - v189)
					v193 = *(*int64)(unsafe.Add(mBase, uint32(v85)+120))
					v195 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[16]))
					v196 = *(*int64)(unsafe.Add(mBase, uint32(v87)+120))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+120)) = v193 + (v195 - v196)
					v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
					if v203&int32(1) != 0 {
						v206 = v15 + int32(24)
					} else {
						v206 = int32(0)
					}
					v207 = v206
				} else {
					v207 = int32(0)
				}
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
				if v211 != 0 {
					v212 = v15 + int32(8)
				} else {
					v212 = int32(0)
				}
				F_ExplainOnePlan(m, v49, l2, l3, l4, l5, l6, v15+int32(280), v207, v212)
				mBase = m.M
				v214 = m.ExcPending
				if v214 != 0 {
					return
				} else {
					m.G0 = v15 + int32(288)
					return
				}
			}
		}
	}
}
func F_start_apply(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v163 int64
	_ = v163
	var v166 int64
	_ = v166
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v208 int64
	_ = v208
	var v211 int64
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int64
	_ = v259
	var v260 int64
	_ = v260
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int64
	_ = v284
	var v285 int32
	_ = v285
	var v287 int64
	_ = v287
	var v288 int32
	_ = v288
	var v290 int64
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int64
	_ = v296
	var v298 int64
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int64
	_ = v326
	var v327 int32
	_ = v327
	var v329 int64
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int64
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int64
	_ = v355
	var v356 int64
	_ = v356
	var v366 int32
	_ = v366
	var v371 int64
	_ = v371
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v405 int64
	_ = v405
	var v408 int64
	_ = v408
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int64
	_ = v499
	var v500 int64
	_ = v500
	var v508 int64
	_ = v508
	var v510 int32
	_ = v510
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v588 int64
	_ = v588
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v647 int32
	_ = v647
	var v648 int64
	_ = v648
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	v2 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	v28 = v2
	v30 = v2
	v31 = v2
	v32 = int32(-1)
	v34 = v2
	v35 = v24
	v36 = v2
	v37 = v2
	v38 = v2
	v39 = v2
	v40 = v2
	v41 = v2
	goto L1
L1:
	;
	if v32 != int32(1) {
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
	v50 = int32(16)
	v51 = v35 - v50
	m.G0 = v51
	v54 = v51 - v50
	m.G0 = v54
	v57 = v54 - v50
	m.G0 = v57
	v60 = v57 - v50
	m.G0 = v60
	v63 = v60 - v50
	m.G0 = v63
	v66 = v63 - int32(160)
	m.G0 = v66
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[0]))
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v24 + int32(8)
	goto L6
L4:
	;
	v78 = v30
	v79 = v31
	v80 = v34
	v81 = v35
	v82 = v36
	v83 = v37
	v84 = v38
	v85 = v39
	v86 = v40
	v87 = v41
	goto L5
L5:
	;
	goto L8
L6:
	;
	v78 = int32(0)
	v79 = v63
	v80 = v54
	v81 = v66
	v82 = v57
	v83 = v60
	v84 = v66
	v85 = v70
	v86 = v72
	v87 = v51
	goto L5
L7:
	;
	goto L2
L8:
	;
	if v78 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v647 = int32(m.ExcTag)
	v648 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v647 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[1])) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
	v96 = m.G0
	v97 = int32(16)
	v98 = v96 - v97
	m.G0 = v98
	F___gettimeofday(m, v98)
	mBase = m.M
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v98)))
	v102 = int64(*(*int32)(unsafe.Add(mBase, uint32(v98)+8)))
	m.G0 = v98 + v97
	goto L14
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[0])) = v85
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[1])) = v86
	v588 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_start_apply[2])) = v588
	*(*int64)(unsafe.Add(mBase, _c_F_start_apply[3])) = v588
	v594 = int32(0)
	*(*uint16)(unsafe.Add(mBase, _c_F_start_apply[4])) = uint16(v594)
	v597 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[5]))
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597)+29)))
	if v598 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[6]))
	v119 = F_AllocSetContextCreateInternal(m, v114, int32(_a_F_start_apply_0), int32(0), int32(_a_F_start_apply_1), int32(_a_F_start_apply_2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[7])) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[6]))
	v130 = F_AllocSetContextCreateInternal(m, v125, int32(_a_F_start_apply_3), int32(0), int32(_a_F_start_apply_1), int32(_a_F_start_apply_2))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[8])) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
	v135 = int32(0)
	F_pgstat_report_activity(m, int32(2), v135)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = int32(987)
	v140 = int32(_a_F_start_apply_4)
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v141
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[9])) = v80
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[0])) = v80
	v148 = v28
	v152 = v135
	v163 = l0
	v166 = v102 + v101*int64(1000000) - int64(946684800000000)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[10]))
	if v173 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[9])) = v564
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[0])) = v564
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[11]))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[12]))
	m.T0[v570].(func(*base.Module, int32, int32))(m, v573, v87)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L10
	} else {
		goto L117
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v148
	F_ProcessInterrupts(m)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[13])) = v179
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[11]))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v148
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[12]))
	v187 = m.T0[v183].(func(*base.Module, int32, int32, int32) int32)(m, v186, v83, v82)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L10
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v189 = int32(0)
	if v187 == v189 {
		v390 = v148
		v394 = v152
		v395 = v189
		v405 = v163
		v408 = v166
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	v411 = int32(0)
	F_send_feedback(m, v405, v411, v411)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L10
	} else {
		goto L70
	}
L25:
	;
	v193 = v148
	v195 = v187
	v197 = v152
	v208 = v163
	v211 = v166
	goto L26
L26:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[10]))
	if v214 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	F_ProcessInterrupts(m)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L10
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v195 == int32(0) {
		v390 = v193
		v394 = v197
		v395 = v189
		v405 = v208
		v408 = v211
		goto L24
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	if v195 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v225 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L10
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[14]))
	if v242 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v227 = int32(1)
	if v225 == int32(0) {
		v390 = v193
		v394 = v197
		v395 = v227
		v405 = v208
		v408 = v211
		goto L24
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	F_errmsg(m, int32(_a_F_start_apply_5), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	F_errfinish(m, int32(_a_F_start_apply_6), int32(3639), int32(_a_F_start_apply_7))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v390 = v193
	v394 = v197
	v395 = v227
	v405 = v208
	v408 = v211
	goto L24
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[14])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v254 = m.G0
	v255 = int32(16)
	v256 = v254 - v255
	m.G0 = v256
	F___gettimeofday(m, v256)
	mBase = m.M
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
	v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
	m.G0 = v256 + v255
	goto L44
L43:
	;
	goto L42
L44:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[13])) = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v195
	*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v279 = F_pq_getmsgbyte(m, v79)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L10
	} else {
		goto L48
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[7]))
	F_MemoryContextReset(m, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L10
	} else {
		goto L68
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v326 = F_pq_getmsgint64(m, v79)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L10
	} else {
		goto L60
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v284 = F_pq_getmsgint64(m, v79)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L10
	} else {
		goto L49
	}
L48:
	;
	switch v279 - int32(107) {
	case 0:
		goto L46
	default:
		v371 = v208
		goto L45
	case 12:
		goto L47
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v287 = F_pq_getmsgint64(m, v79)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v290 = F_pq_getmsgint64(m, v79)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v293)+80)) = v290
	if base.Ui64(v284) < base.Ui64(v208) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v296 = v208
	goto L54
L53:
	;
	v296 = v284
	goto L54
L54:
	;
	if base.Ui64(v287) < base.Ui64(v296) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v298 = v296
	goto L57
L56:
	;
	v298 = v287
	goto L57
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v293)+72)) = v298
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v304 = m.G0
	v305 = int32(16)
	v306 = v304 - v305
	m.G0 = v306
	F___gettimeofday(m, v306)
	mBase = m.M
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v306)))
	v310 = int64(*(*int32)(unsafe.Add(mBase, uint32(v306)+8)))
	m.G0 = v306 + v305
	goto L58
L58:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+88)) = v310 + v309*int64(1000000) - int64(946684800000000)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	F_apply_dispatch(m, v79)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	v371 = v298
	goto L45
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v329 = F_pq_getmsgint64(m, v79)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v332 = F_pq_getmsgbyte(m, v79)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	if base.Ui64(v326) < base.Ui64(v208) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v336 = v208
	goto L65
L64:
	;
	v336 = v326
	goto L65
L65:
	;
	v337 = int32(0)
	F_send_feedback(m, v336, base.B2i32(v332 != v337), v337)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v343)+80)) = v329
	*(*int64)(unsafe.Add(mBase, uint32(v343)+72)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v350 = m.G0
	v351 = int32(16)
	v352 = v350 - v351
	m.G0 = v352
	F___gettimeofday(m, v352)
	mBase = m.M
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v352)))
	v356 = int64(*(*int32)(unsafe.Add(mBase, uint32(v352)+8)))
	m.G0 = v352 + v351
	goto L67
L67:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+104)) = v329
	*(*int64)(unsafe.Add(mBase, uint32(v366)+96)) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v366)+88)) = v356 + v355*int64(1000000) - int64(946684800000000)
	v371 = v336
	goto L45
L68:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[11]))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v386 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[12]))
	v387 = m.T0[v383].(func(*base.Module, int32, int32, int32) int32)(m, v386, v83, v82)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	v193 = v387
	v195 = v387
	v197 = int32(0)
	v208 = v371
	v211 = v260 + v259*int64(1000000) - int64(946684800000000)
	goto L26
L70:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_apply[16])))
	if v416 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[7]))
	F_MemoryContextReset(m, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L10
	} else {
		goto L77
	}
L72:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_apply[17])))
	if v418 != 0 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	F_maybe_reread_subscription(m)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	F_process_syncing_tables(m, v405)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	goto L71
L77:
	;
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[18]))
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[13])) = v435
	if v395 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	v442 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[19]))
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[20]))
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[21]))
	if v447 == int32(_a_F_start_apply_8) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L80
L80:
	;
	goto L18
L81:
	;
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[14]))
	if v473 != 0 {
		goto L93
	} else {
		goto L94
	}
L82:
	;
	v450 = int32(1000)
	goto L84
L83:
	;
	v450 = v445
	goto L84
L84:
	;
	if v447 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v452 = v450
	goto L87
L86:
	;
	v452 = int32(1000)
	goto L87
L87:
	;
	v454 = F_WaitLatchOrSocket(m, v442, v439, v452, int32(83886087))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L10
	} else {
		goto L88
	}
L88:
	;
	if v454&int32(1) == int32(0) {
		goto L81
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = int32(0)
	goto L90
L90:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[10]))
	if v466 == int32(0) {
		goto L81
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	F_ProcessInterrupts(m)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	goto L81
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[14])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L10
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v454&int32(8) == int32(0) {
		v148 = v390
		v152 = v394
		v163 = v405
		v166 = v408
		goto L17
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	v485 = int32(0)
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[22]))
	if v487 <= v485 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	F_send_feedback(m, v405, v546, v546)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L10
	} else {
		goto L111
	}
L99:
	;
	v546 = v485
	v547 = v394
	goto L98
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	v494 = m.G0
	v495 = int32(16)
	v496 = v494 - v495
	m.G0 = v496
	F___gettimeofday(m, v496)
	mBase = m.M
	v499 = *(*int64)(unsafe.Add(mBase, uint32(v496)))
	v500 = int64(*(*int32)(unsafe.Add(mBase, uint32(v496)+8)))
	m.G0 = v496 + v495
	v508 = v500 + v499*int64(1000000) - int64(946684800000000)
	goto L102
L102:
	;
	v510 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[22]))
	if base.I64_extend_i32_s(v510)*int64(1000)+v408 <= v508 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L10
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v536 = int32(1)
	if v394&v536 != 0 {
		v546 = v485
		v547 = v536
		goto L98
	} else {
		goto L110
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	F_errcode(m, int32(100663808))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L10
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	F_errmsg(m, int32(_a_F_start_apply_9), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L10
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	F_errfinish(m, int32(_a_F_start_apply_6), int32(3793), int32(_a_F_start_apply_7))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	goto L7
L110:
	;
	v540 = base.I32_div_s(v510, int32(2))
	v545 = base.B2i32(base.I64_extend_i32_s(v540)*int64(1000)+v408 <= v508)
	v546 = v545
	v547 = v545
	goto L98
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[23]))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)+20))
	goto L112
L112:
	;
	if v556 == int32(2) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v148 = v390
	v152 = v547
	v163 = v405
	v166 = v408
	goto L17
L114:
	;
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	v561 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L10
	} else {
		goto L116
	}
L116:
	;
	v148 = v390
	v152 = v547
	v163 = v405
	v166 = v408
	goto L17
L117:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[0])) = v85
	*(*int32)(unsafe.Add(mBase, _c_F_start_apply[1])) = v86
	m.G0 = v24 + int32(16)
	return
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
	F_DisableSubscriptionAndExit(m)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L10
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
	F_AbortOutOfAnyTransaction(m)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L10
	} else {
		goto L122
	}
L121:
	;
	goto L7
L122:
	;
	v607 = int32(1)
	v609 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[5]))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v609)))
	v612 = *(*int32)(unsafe.Add(mBase, _c_F_start_apply[15]))
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+16)))
	if v613 == v607 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	v619 = base.B2i32(v616 != int32(1))
	goto L125
L124:
	;
	v619 = v607
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
	F_pgstat_report_subscription_error(m, v610, v619)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L10
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
	F_pg_re_throw(m)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L10
	} else {
		goto L127
	}
L127:
	;
	goto L9
L128:
	;
	v652 = int32(v648)
	m.G0 = v81
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v652)+4))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v652)))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v655)))
	if v24+int32(8) == v659 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	m.ExcPending = 1
	goto L137
L130:
	;
	if v662 != 0 {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v655)+4))
	v662 = v661
	goto L133
L132:
	;
	v662 = int32(0)
	goto L133
L133:
	;
	goto L130
L134:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v28 = v663
	v30 = v654
	v31 = v79
	v32 = v662
	v34 = v80
	v35 = v81
	v36 = v82
	v37 = v83
	v38 = v84
	v39 = v85
	v40 = v86
	v41 = v87
	goto L1
L135:
	;
	goto L136
L136:
	;
	F___wasm_longjmp(m, v655, v654)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	return
L138:
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
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
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
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
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
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
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
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
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
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L51
	}
L4:
	;
	F_sequence_close(m, v19, int32(3))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
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
	v27 = v15 - int32(-64)
	v40 = int32(0)
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v40<<(uint(int32(2))%32))))
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
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+8)))
	v51 = F_SearchSysCacheExists(m, int32(5), l1, v48, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v53 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = v53
	v55 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(44)))) = uint16(v55)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v55
	v65 = F_GetNewOidWithIndex(m, v19, int32(2757), int32(1))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	if v51 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v65
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v71
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v82 = F_heap_form_tuple(m, v77, v15+int32(48), v15+int32(40))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_CatalogTupleInsert(m, v19, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_pfree(m, v82)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v88 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = int32(2603)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(1255)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v95
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+24)))
	if v105 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v106 = int32(110)
	goto L20
L19:
	;
	v106 = int32(97)
	goto L20
L20:
	;
	F_recordDependencyOn(m, v15+int32(28), v15+int32(16), v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+25)))
	if v111 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v112 = int32(2753)
	goto L24
L23:
	;
	v112 = int32(2616)
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v114
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+24)))
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
	F_recordDependencyOn(m, v15+int32(28), v15+int32(16), v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v129 = F_typeDepNeeded(m, v128, v46)
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
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v133
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+24)))
	if v143 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if v148 == v149 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v144 = int32(110)
	goto L35
L34:
	;
	v144 = int32(97)
	goto L35
L35:
	;
	F_recordDependencyOn(m, v15+int32(28), v15+int32(16), v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_storeProcedures[0]))
	if v173 != 0 {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v151 = F_typeDepNeeded(m, v148, v46)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v151 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(1247)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v157
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+24)))
	if v167 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v168 = int32(110)
	goto L43
L42:
	;
	v168 = int32(97)
	goto L43
L43:
	;
	F_recordDependencyOn(m, v15+int32(28), v15+int32(16), v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L37
L45:
	;
	v175 = int32(0)
	F_RunObjectPostCreateHook(m, int32(2603), v65, v175, v175)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v180 = v40 + int32(1)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v180 < v181 {
		v40 = v180
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
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v210 = F_format_type_be(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v213 = F_format_type_be(m, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v215 = F_NameListToString(m, l0)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v208
	F_errmsg(m, int32(_a_F_storeProcedures_1), v15)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_storeProcedures_2), int32(1618), int32(_a_F_storeProcedures_3))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
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
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
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
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v173 int32
	_ = v173
	var v187 int32
	_ = v187
	var v201 int32
	_ = v201
	var v215 int32
	_ = v215
	var v237 int32
	_ = v237
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v281 int32
	_ = v281
	var v302 int32
	_ = v302
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v346 int32
	_ = v346
	var v360 int32
	_ = v360
	var v374 int32
	_ = v374
	var v388 int32
	_ = v388
	var v402 int32
	_ = v402
	var v416 int32
	_ = v416
	var v430 int32
	_ = v430
	var v452 int32
	_ = v452
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v496 int32
	_ = v496
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v540 int32
	_ = v540
	var v561 int32
	_ = v561
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v666 int32
	_ = v666
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v818 int32
	_ = v818
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	v2 = l1
	v23 = m.G0
	v25 = v23 - int32(32)
	m.G0 = v25
	if l0&int32(3) == int32(0) {
		v50 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v85 = v83 + int32(17)
	v86 = F_palloc(m, v85)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v83 = v75 - l0
	goto L1
L3:
	;
	v54 = v50
	goto L12
L4:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v34 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v83 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v39 = l0
	goto L8
L8:
	;
	v43 = v39 + int32(1)
	if v43&int32(3) == int32(0) {
		v50 = v43
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v75 = v43
	goto L2
L10:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v48 != 0 {
		v39 = v43
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v63 = int32(-2139062144)
	if (int32(16843008)-v60|v60)&v63 == v63 {
		v54 = v54 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v69 = v54
	goto L15
L14:
	;
	goto L13
L15:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v73 != 0 {
		v69 = v69 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v75 = v69
	goto L2
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v90 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L18
	} else {
		goto L226
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L18
	} else {
		goto L222
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L18
	} else {
		goto L219
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L18
	} else {
		goto L216
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L18
	} else {
		goto L213
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L18
	} else {
		goto L210
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L18
	} else {
		goto L207
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L18
	} else {
		goto L204
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L18
	} else {
		goto L200
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L18
	} else {
		goto L197
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L18
	} else {
		goto L194
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L18
	} else {
		goto L191
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L18
	} else {
		goto L188
	}
L33:
	;
	v818 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v802))) = uint8(v818)
	m.G0 = v25 + int32(32)
	return v808
L34:
	;
	v802 = v86
	v808 = v86
	goto L33
L35:
	;
	goto L36
L36:
	;
	v93 = l2 - l0
	v97 = l0
	v99 = int32(0)
	v101 = v90
	v103 = v86
	v109 = v86
	v111 = v85
	goto L37
L37:
	;
	v119 = v103 - v109
	if base.Ui32(v111-int32(17)) < base.Ui32(v119) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v788 != 0 {
		v986 = v794
		goto L20
	} else {
		goto L187
	}
L39:
	;
	v124 = v111 << (uint(int32(1)) % 32)
	v125 = F_repalloc(m, v109, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L18
	} else {
		goto L42
	}
L40:
	;
	v129 = v101
	v130 = v103
	v131 = v109
	v132 = v111
	goto L41
L41:
	;
	v133 = int32(255)
	v134 = v2 & v133
	if v134 == v129&v133 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v129 = v128
	v130 = v125 + v119
	v131 = v125
	v132 = v124
	goto L41
L43:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(12))+8))
	*(*int32)(unsafe.Add(mBase, _c_F_str_udeescape[0])) = v792
	goto L185
L44:
	;
	v776 = v762
	v777 = v773
	v788 = int32(0)
	goto L43
L45:
	;
	F_pg_unicode_to_server(m, v335, v130)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L18
	} else {
		goto L167
	}
L46:
	;
	v139 = v25 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v139)+12)) = int32(499)
	*(*int32)(unsafe.Add(mBase, uint32(v139)+4)) = v97 + (v93 + int32(3))
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v139)+16)) = v139
	v146 = int32(_a_F_str_udeescape_0)
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_str_udeescape[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+8)) = v147
	*(*int32)(unsafe.Add(mBase, _c_F_str_udeescape[0])) = v25 + int32(20)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v134 == v153 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	if v99 != 0 {
		v986 = v97
		goto L20
	} else {
		goto L165
	}
L49:
	;
	if v99 != 0 {
		v986 = v97
		goto L20
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L54
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v2)
	v776 = int32(2)
	v777 = v130 + int32(1)
	v788 = int32(0)
	goto L43
L53:
	;
	if v153 != int32(43) {
		goto L90
	} else {
		goto L91
	}
L54:
	;
	if base.B2i32(base.Ui32(v153-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v153|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
	goto L56
L56:
	;
	if base.B2i32(base.Ui32(v173-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v173|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+3)))
	goto L58
L58:
	;
	if base.B2i32(base.Ui32(v187-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v187|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L53
	} else {
		goto L59
	}
L59:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	goto L60
L60:
	;
	if base.B2i32(base.Ui32(v201-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v201|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v215 = int32(-48)
	if base.Ui32((v153-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v237 = v215
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if base.Ui32((v173-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v258 = v215
		goto L66
	} else {
		goto L67
	}
L63:
	;
	if base.Ui32((v153-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v237 = int32(-87)
		goto L62
	} else {
		goto L64
	}
L64:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v153-int32(65))&int32(255)) {
		goto L32
	} else {
		goto L65
	}
L65:
	;
	v237 = int32(-55)
	goto L62
L66:
	;
	v259 = int32(-48)
	if base.Ui32((v187-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v281 = v259
		goto L72
	} else {
		goto L73
	}
L67:
	;
	if base.Ui32((v173-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v258 = int32(-87)
	goto L66
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v173-int32(65))&int32(255)) {
		goto L31
	} else {
		goto L71
	}
L71:
	;
	v258 = int32(-55)
	goto L66
L72:
	;
	if base.Ui32((v201-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v302 = v259
		goto L76
	} else {
		goto L77
	}
L73:
	;
	if base.Ui32((v187-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v281 = int32(-87)
		goto L72
	} else {
		goto L74
	}
L74:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v187-int32(65))&int32(255)) {
		goto L30
	} else {
		goto L75
	}
L75:
	;
	v281 = int32(-55)
	goto L72
L76:
	;
	v315 = v201 + v302 + ((v173+v258)<<(uint(int32(8))%32) + (v153+v237)<<(uint(int32(12))%32) + (v187+v281)<<(uint(int32(4))%32))
	if base.Ui32(int32(_a_F_str_udeescape_1)) <= base.Ui32(v315-int32(1)) {
		goto L28
	} else {
		goto L82
	}
L77:
	;
	if base.Ui32((v201-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v302 = int32(-87)
	goto L76
L79:
	;
	goto L80
L80:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v201-int32(65))&int32(255)) {
		goto L29
	} else {
		goto L81
	}
L81:
	;
	v302 = int32(-55)
	goto L76
L82:
	;
	v321 = v315 & int32(_a_F_str_udeescape_2)
	if v99 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v336 = int32(5)
	if v335&int32(16776192) != int32(_a_F_str_udeescape_3) {
		goto L45
	} else {
		goto L89
	}
L84:
	;
	if v321 != int32(_a_F_str_udeescape_4) {
		v986 = v97
		goto L20
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if v321 == int32(_a_F_str_udeescape_4) {
		v986 = v97
		goto L20
	} else {
		goto L88
	}
L87:
	;
	v335 = v99<<(uint(int32(10))%32)&int32(_a_F_str_udeescape_5) | v315&int32(1023) + int32(_a_F_str_udeescape_6)
	goto L83
L88:
	;
	v335 = v315
	goto L83
L89:
	;
	v776 = v336
	v777 = v130
	v788 = v335
	goto L43
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L18
	} else {
		goto L160
	}
L91:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
	goto L92
L92:
	;
	if base.B2i32(base.Ui32(v346-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v346|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+3)))
	goto L94
L94:
	;
	if base.B2i32(base.Ui32(v360-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v360|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	goto L96
L96:
	;
	if base.B2i32(base.Ui32(v374-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v374|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L90
	} else {
		goto L97
	}
L97:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)))
	goto L98
L98:
	;
	if base.B2i32(base.Ui32(v388-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v388|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L90
	} else {
		goto L99
	}
L99:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+6)))
	goto L100
L100:
	;
	if base.B2i32(base.Ui32(v402-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v402|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L90
	} else {
		goto L101
	}
L101:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+7)))
	goto L102
L102:
	;
	if base.B2i32(base.Ui32(v416-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v416|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L90
	} else {
		goto L103
	}
L103:
	;
	v430 = int32(-48)
	if base.Ui32((v346-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v452 = v430
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if base.Ui32((v360-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v473 = v430
		goto L108
	} else {
		goto L109
	}
L105:
	;
	if base.Ui32((v346-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v452 = int32(-87)
		goto L104
	} else {
		goto L106
	}
L106:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v346-int32(65))&int32(255)) {
		goto L27
	} else {
		goto L107
	}
L107:
	;
	v452 = int32(-55)
	goto L104
L108:
	;
	v474 = int32(-48)
	if base.Ui32((v374-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v496 = v474
		goto L114
	} else {
		goto L115
	}
L109:
	;
	if base.Ui32((v360-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v473 = int32(-87)
	goto L108
L111:
	;
	goto L112
L112:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v360-int32(65))&int32(255)) {
		goto L26
	} else {
		goto L113
	}
L113:
	;
	v473 = int32(-55)
	goto L108
L114:
	;
	if base.Ui32((v388-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v517 = v474
		goto L118
	} else {
		goto L119
	}
L115:
	;
	if base.Ui32((v374-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v496 = int32(-87)
		goto L114
	} else {
		goto L116
	}
L116:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v374-int32(65))&int32(255)) {
		goto L25
	} else {
		goto L117
	}
L117:
	;
	v496 = int32(-55)
	goto L114
L118:
	;
	v518 = int32(-48)
	if base.Ui32((v402-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v540 = v518
		goto L124
	} else {
		goto L125
	}
L119:
	;
	if base.Ui32((v388-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v517 = int32(-87)
	goto L118
L121:
	;
	goto L122
L122:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v388-int32(65))&int32(255)) {
		goto L24
	} else {
		goto L123
	}
L123:
	;
	v517 = int32(-55)
	goto L118
L124:
	;
	if base.Ui32((v416-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v561 = v518
		goto L128
	} else {
		goto L129
	}
L125:
	;
	if base.Ui32((v402-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v540 = int32(-87)
		goto L124
	} else {
		goto L126
	}
L126:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v402-int32(65))&int32(255)) {
		goto L23
	} else {
		goto L127
	}
L127:
	;
	v540 = int32(-55)
	goto L124
L128:
	;
	v582 = v416 + v561 + ((v360+v473)<<(uint(int32(16))%32) + (v346+v452)<<(uint(int32(20))%32) + (v374+v496)<<(uint(int32(12))%32) + (v388+v517)<<(uint(int32(8))%32) + (v402+v540)<<(uint(int32(4))%32))
	if base.Ui32(int32(_a_F_str_udeescape_1)) <= base.Ui32(v582-int32(1)) {
		goto L21
	} else {
		goto L134
	}
L129:
	;
	if base.Ui32((v416-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v561 = int32(-87)
	goto L128
L131:
	;
	goto L132
L132:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v416-int32(65))&int32(255)) {
		goto L22
	} else {
		goto L133
	}
L133:
	;
	v561 = int32(-55)
	goto L128
L134:
	;
	v588 = v582 & int32(_a_F_str_udeescape_2)
	if v99 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v603 = int32(8)
	if v602&int32(16776192) == int32(_a_F_str_udeescape_3) {
		v776 = v603
		v777 = v130
		v788 = v602
		goto L43
	} else {
		goto L141
	}
L136:
	;
	if v588 != int32(_a_F_str_udeescape_4) {
		v986 = v97
		goto L20
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	if v588 == int32(_a_F_str_udeescape_4) {
		v986 = v97
		goto L20
	} else {
		goto L140
	}
L139:
	;
	v602 = v99<<(uint(int32(10))%32)&int32(_a_F_str_udeescape_5) | v582&int32(1023) + int32(_a_F_str_udeescape_6)
	goto L135
L140:
	;
	v602 = v582
	goto L135
L141:
	;
	F_pg_unicode_to_server(m, v602, v130)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L18
	} else {
		goto L142
	}
L142:
	;
	if v130&int32(3) == int32(0) {
		v633 = v130
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v762 = v603
	v773 = v666 + v130
	goto L44
L144:
	;
	v666 = v658 - v130
	goto L143
L145:
	;
	v637 = v633
	goto L154
L146:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v617 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v666 = int32(0)
	goto L143
L148:
	;
	goto L149
L149:
	;
	v622 = v130
	goto L150
L150:
	;
	v626 = v622 + int32(1)
	if v626&int32(3) == int32(0) {
		v633 = v626
		goto L145
	} else {
		goto L152
	}
L151:
	;
	v658 = v626
	goto L144
L152:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	if v631 != 0 {
		v622 = v626
		goto L150
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v646 = int32(-2139062144)
	if (int32(16843008)-v643|v643)&v646 == v646 {
		v637 = v637 + int32(4)
		goto L154
	} else {
		goto L156
	}
L155:
	;
	v652 = v637
	goto L157
L156:
	;
	goto L155
L157:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652))))
	if v656 != 0 {
		v652 = v652 + int32(1)
		goto L157
	} else {
		goto L159
	}
L158:
	;
	v658 = v652
	goto L144
L159:
	;
	goto L158
L160:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L18
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(_a_F_str_udeescape_7), int32(0))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L18
	} else {
		goto L162
	}
L162:
	;
	F_errhint(m, int32(_a_F_str_udeescape_8), int32(0))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L18
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(495), int32(_a_F_str_udeescape_10))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L18
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v129)
	v695 = int32(1)
	v696 = v130 + v695
	v699 = v97 + v695
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699))))
	if v700 != 0 {
		v97 = v699
		v99 = int32(0)
		v101 = v700
		v103 = v696
		v109 = v131
		v111 = v132
		goto L37
	} else {
		goto L166
	}
L166:
	;
	v802 = v696
	v808 = v131
	goto L33
L167:
	;
	if v130&int32(3) == int32(0) {
		v726 = v130
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v762 = v336
	v773 = v759 + v130
	goto L44
L169:
	;
	v759 = v751 - v130
	goto L168
L170:
	;
	v730 = v726
	goto L179
L171:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v710 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v759 = int32(0)
	goto L168
L173:
	;
	goto L174
L174:
	;
	v715 = v130
	goto L175
L175:
	;
	v719 = v715 + int32(1)
	if v719&int32(3) == int32(0) {
		v726 = v719
		goto L170
	} else {
		goto L177
	}
L176:
	;
	v751 = v719
	goto L169
L177:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v719))))
	if v724 != 0 {
		v715 = v719
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	v739 = int32(-2139062144)
	if (int32(16843008)-v736|v736)&v739 == v739 {
		v730 = v730 + int32(4)
		goto L179
	} else {
		goto L181
	}
L180:
	;
	v745 = v730
	goto L182
L181:
	;
	goto L180
L182:
	;
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v745))))
	if v749 != 0 {
		v745 = v745 + int32(1)
		goto L182
	} else {
		goto L184
	}
L183:
	;
	v751 = v745
	goto L169
L184:
	;
	goto L183
L185:
	;
	v794 = v97 + v776
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794))))
	if v795 != 0 {
		v97 = v794
		v99 = v788
		v101 = v795
		v103 = v777
		v109 = v131
		v111 = v132
		goto L37
	} else {
		goto L186
	}
L186:
	;
	goto L38
L187:
	;
	v802 = v777
	v808 = v131
	goto L33
L188:
	;
	F_errmsg_internal(m, int32(_a_F_str_udeescape_11), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L18
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(336), int32(_a_F_str_udeescape_12))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L18
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	F_errmsg_internal(m, int32(_a_F_str_udeescape_11), int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L18
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(336), int32(_a_F_str_udeescape_12))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L18
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	F_errmsg_internal(m, int32(_a_F_str_udeescape_11), int32(0))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L18
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(336), int32(_a_F_str_udeescape_12))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L18
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	F_errmsg_internal(m, int32(_a_F_str_udeescape_11), int32(0))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L18
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(336), int32(_a_F_str_udeescape_12))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L18
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L18
	} else {
		goto L201
	}
L201:
	;
	F_errmsg(m, int32(_a_F_str_udeescape_13), int32(0))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L18
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(347), int32(_a_F_str_udeescape_14))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L18
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	F_errmsg_internal(m, int32(_a_F_str_udeescape_11), int32(0))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L18
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(336), int32(_a_F_str_udeescape_12))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L18
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	F_errmsg_internal(m, int32(_a_F_str_udeescape_11), int32(0))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L18
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(336), int32(_a_F_str_udeescape_12))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L18
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	F_errmsg_internal(m, int32(_a_F_str_udeescape_11), int32(0))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L18
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(336), int32(_a_F_str_udeescape_12))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L18
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	F_errmsg_internal(m, int32(_a_F_str_udeescape_11), int32(0))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L18
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(336), int32(_a_F_str_udeescape_12))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L18
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	F_errmsg_internal(m, int32(_a_F_str_udeescape_11), int32(0))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L18
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(336), int32(_a_F_str_udeescape_12))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L18
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	F_errmsg_internal(m, int32(_a_F_str_udeescape_11), int32(0))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L18
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(336), int32(_a_F_str_udeescape_12))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L18
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L18
	} else {
		goto L223
	}
L223:
	;
	F_errmsg(m, int32(_a_F_str_udeescape_13), int32(0))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L18
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(347), int32(_a_F_str_udeescape_14))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L18
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L18
	} else {
		goto L227
	}
L227:
	;
	F_errmsg(m, int32(_a_F_str_udeescape_15), int32(0))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L18
	} else {
		goto L228
	}
L228:
	;
	F_scanner_errposition(m, v986+v93+int32(3), l3)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L18
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(_a_F_str_udeescape_9), int32(525), int32(_a_F_str_udeescape_10))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L18
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_strcat(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	if l0&int32(3) == int32(0) {
		v26 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v60 = v59 + l0
	if (l1^v60)&int32(3) != 0 {
		goto L21
	} else {
		goto L22
	}
L2:
	;
	v59 = v51 - l0
	goto L1
L3:
	;
	v30 = v26
	goto L12
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v59 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v15 = l0
	goto L8
L8:
	;
	v19 = v15 + int32(1)
	if v19&int32(3) == int32(0) {
		v26 = v19
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v51 = v19
	goto L2
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v24 != 0 {
		v15 = v19
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v39 = int32(-2139062144)
	if (int32(16843008)-v36|v36)&v39 == v39 {
		v30 = v30 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v45 = v30
	goto L15
L14:
	;
	goto L13
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 != 0 {
		v45 = v45 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v51 = v45
	goto L2
L17:
	;
	goto L16
L18:
	;
	return l0
L19:
	;
	goto L18
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v114)
	if v114&int32(255) == int32(0) {
		goto L19
	} else {
		goto L35
	}
L21:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v113 = l1
	v114 = v66
	v115 = v60
	goto L20
L22:
	;
	goto L23
L23:
	;
	if l1&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v70 = l1
	v72 = v60
	goto L27
L25:
	;
	v84 = l1
	v86 = v60
	goto L26
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v91 = int32(-2139062144)
	if (int32(16843008)-v88|v88)&v91 != v91 {
		v113 = v84
		v114 = v88
		v115 = v86
		goto L20
	} else {
		goto L31
	}
L27:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v73)
	if v73 == int32(0) {
		goto L19
	} else {
		goto L29
	}
L28:
	;
	v84 = v80
	v86 = v78
	goto L26
L29:
	;
	v77 = int32(1)
	v78 = v72 + v77
	v80 = v70 + v77
	if v80&int32(3) != 0 {
		v70 = v80
		v72 = v78
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v96 = v84
	v97 = v88
	v98 = v86
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v97
	v100 = int32(4)
	v101 = v98 + v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v104 = v96 + v100
	v108 = int32(-2139062144)
	if (v102|(int32(16843008)-v102))&v108 == v108 {
		v96 = v104
		v97 = v102
		v98 = v101
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v113 = v104
	v114 = v102
	v115 = v101
	goto L20
L34:
	;
	goto L33
L35:
	;
	v122 = v113
	v124 = v115
	goto L36
L36:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)) = uint8(v125)
	v127 = int32(1)
	if v125 != 0 {
		v122 = v122 + v127
		v124 = v124 + v127
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L19
L38:
	;
	goto L37
}
func F_strcmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v6 == int32(0) {
		v25 = v5
		v26 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v26 - v25
L2:
	;
	if v5 != v6 {
		v25 = v5
		v26 = v6
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v10 = l0
	v11 = l1
	goto L4
L4:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v15 == int32(0) {
		v25 = v14
		v26 = v15
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v25 = v14
	v26 = v15
	goto L1
L6:
	;
	v18 = int32(1)
	if v14 == v15 {
		v10 = v10 + v18
		v11 = v11 + v18
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v46 = v38 + v42
	v50 = int32(-2139062144)
	if (v44|(int32(16843008)-v44))&v50 == v50 {
		v38 = v46
		v39 = v44
		v40 = v43
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v55 = v46
	v56 = v44
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
func F_strerror(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	if base.Ui32(l0) <= base.Ui32(int32(153)) {
		v5 = l0
	} else {
		v5 = int32(0)
	}
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5<<(uint(int32(1))%32))+uint32(_c_F_strerror[0]))))
	return v10 + int32(_a_F_strerror_0)
}
func F_strlcat(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
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
	if l1&int32(3) == int32(0) {
		v34 = l1
		goto L11
	} else {
		goto L12
	}
L7:
	;
	goto L8
L8:
	;
	v68 = l0 + v9
	v69 = l2 - v9
	if v69 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	goto L1
L10:
	;
	goto L9
L11:
	;
	v38 = v34
	goto L20
L12:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v18 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	goto L15
L15:
	;
	v23 = l1
	goto L16
L16:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L11
	} else {
		goto L18
	}
L17:
	;
	goto L10
L18:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v53 = v38
	goto L23
L22:
	;
	goto L21
L23:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	goto L10
L25:
	;
	goto L24
L26:
	;
	goto L1
L27:
	;
	v181 = F_strlen(m, v177)
	mBase = m.M
	goto L26
L28:
	;
	v177 = l1
	goto L27
L29:
	;
	goto L30
L30:
	;
	v75 = v69 - int32(1)
	if (v68^l1)&int32(3) != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v174 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v174)
	v177 = v170
	goto L27
L32:
	;
	v155 = v150
	v156 = v151
	v157 = v152
	goto L54
L33:
	;
	if v145 == int32(0) {
		v170 = v143
		v171 = v144
		goto L31
	} else {
		goto L53
	}
L34:
	;
	v143 = l1
	v144 = v68
	v145 = v75
	goto L33
L35:
	;
	goto L36
L36:
	;
	v79 = int32(0)
	if l1&int32(3) == v79 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v112 == int32(0) {
		v170 = v109
		v171 = v110
		goto L31
	} else {
		goto L46
	}
L38:
	;
	v109 = l1
	v110 = v68
	v111 = v75
	v112 = base.B2i32(v75 != v79)
	goto L37
L39:
	;
	if v75 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v88 = l1
	v89 = v68
	v90 = v75
	goto L41
L41:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v92)
	if v92 == int32(0) {
		v150 = v88
		v151 = v89
		v152 = v90
		goto L32
	} else {
		goto L43
	}
L42:
	;
	v109 = v103
	v110 = v97
	v111 = v99
	v112 = v101
	goto L37
L43:
	;
	v96 = int32(1)
	v97 = v89 + v96
	v99 = v90 - v96
	v100 = int32(0)
	v101 = base.B2i32(v99 != v100)
	v103 = v88 + v96
	if v103&int32(3) == v100 {
		v109 = v103
		v110 = v97
		v111 = v99
		v112 = v101
		goto L37
	} else {
		goto L44
	}
L44:
	;
	if v99 != 0 {
		v88 = v103
		v89 = v97
		v90 = v99
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v115 == int32(0) {
		v143 = v109
		v144 = v110
		v145 = v111
		goto L33
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(v111) < base.Ui32(int32(4)) {
		v143 = v109
		v144 = v110
		v145 = v111
		goto L33
	} else {
		goto L48
	}
L48:
	;
	v121 = v109
	v122 = v110
	v123 = v111
	goto L49
L49:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v129 = int32(-2139062144)
	if (int32(16843008)-v126|v126)&v129 != v129 {
		v150 = v121
		v151 = v122
		v152 = v123
		goto L32
	} else {
		goto L51
	}
L50:
	;
	v143 = v137
	v144 = v135
	v145 = v139
	goto L33
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v126
	v134 = int32(4)
	v135 = v122 + v134
	v137 = v121 + v134
	v139 = v123 - v134
	if base.Ui32(int32(3)) < base.Ui32(v139) {
		v121 = v137
		v122 = v135
		v123 = v139
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v150 = v143
	v151 = v144
	v152 = v145
	goto L32
L54:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v159)
	if v159 == int32(0) {
		v170 = v155
		v171 = v156
		goto L31
	} else {
		goto L56
	}
L55:
	;
	v170 = v166
	v171 = v164
	goto L31
L56:
	;
	v163 = int32(1)
	v164 = v156 + v163
	v166 = v155 + v163
	v168 = v157 - v163
	if v168 != 0 {
		v155 = v166
		v156 = v164
		v157 = v168
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	v36 = l1
	v40 = int32(0)
	goto L6
L6:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	return v40 - v41
L7:
	;
	v36 = v31
	v40 = v33
	goto L6
L8:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 != v16 {
		v31 = v12
		v33 = v14
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v31 = v25
	v33 = int32(0)
	goto L7
L10:
	;
	if v16 == int32(0) {
		v31 = v12
		v33 = v14
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v21 = v13 - int32(1)
	if v21 == int32(0) {
		v31 = v12
		v33 = v14
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v24 = int32(1)
	v25 = v12 + v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v26 != 0 {
		v11 = v11 + v24
		v12 = v25
		v13 = v21
		v14 = v26
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
}
func F_strncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	if (l1^l0)&int32(3) != 0 {
		v73 = l1
		v74 = l2
		v75 = l0
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v112 = F__emscripten_memset_bulkmem(m, v108, base.I32_extend8_s(int32(0)), v107)
	mBase = m.M
	goto L26
L2:
	;
	v107 = int32(0)
	v108 = v102
	goto L1
L3:
	;
	v85 = v80
	v86 = v81
	v87 = v82
	goto L22
L4:
	;
	if v74 == int32(0) {
		v102 = v75
		goto L2
	} else {
		goto L21
	}
L5:
	;
	v9 = int32(0)
	v10 = base.B2i32(l2 != v9)
	if l1&int32(3) == v9 {
		v39 = l1
		v40 = l2
		v41 = l0
		v42 = v10
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v42 == int32(0) {
		v102 = v41
		goto L2
	} else {
		goto L14
	}
L7:
	;
	if l2 == int32(0) {
		v39 = l1
		v40 = l2
		v41 = l0
		v42 = v10
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v18 = l1
	v19 = l2
	v20 = l0
	goto L9
L9:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v22)
	if v22 == int32(0) {
		v107 = v19
		v108 = v20
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v39 = v33
	v40 = v29
	v41 = v27
	v42 = v31
	goto L6
L11:
	;
	v26 = int32(1)
	v27 = v20 + v26
	v29 = v19 - v26
	v30 = int32(0)
	v31 = base.B2i32(v29 != v30)
	v33 = v18 + v26
	if v33&int32(3) == v30 {
		v39 = v33
		v40 = v29
		v41 = v27
		v42 = v31
		goto L6
	} else {
		goto L12
	}
L12:
	;
	if v29 != 0 {
		v18 = v33
		v19 = v29
		v20 = v27
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v45 == int32(0) {
		v107 = v40
		v108 = v41
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if base.Ui32(v40) < base.Ui32(int32(4)) {
		v73 = v39
		v74 = v40
		v75 = v41
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v51 = v39
	v52 = v40
	v53 = v41
	goto L17
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v59 = int32(-2139062144)
	if (int32(16843008)-v56|v56)&v59 != v59 {
		v80 = v51
		v81 = v52
		v82 = v53
		goto L3
	} else {
		goto L19
	}
L18:
	;
	v73 = v67
	v74 = v69
	v75 = v65
	goto L4
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v56
	v64 = int32(4)
	v65 = v53 + v64
	v67 = v51 + v64
	v69 = v52 - v64
	if base.Ui32(int32(3)) < base.Ui32(v69) {
		v51 = v67
		v52 = v69
		v53 = v65
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v80 = v73
	v81 = v74
	v82 = v75
	goto L3
L22:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v89)
	if v89 == int32(0) {
		v107 = v86
		v108 = v87
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v102 = v94
	goto L2
L24:
	;
	v93 = int32(1)
	v94 = v87 + v93
	v98 = v86 - v93
	if v98 != 0 {
		v85 = v85 + v93
		v86 = v98
		v87 = v94
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	return l0
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
	var v17 int32
	_ = v17
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
	var v65 int32
	_ = v65
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
	v71 = v65 - v4 + v4
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v72 != 0 {
		goto L24
	} else {
		goto L25
	}
L5:
	;
	m.G0 = v10 + int32(32)
	goto L4
L6:
	;
	v17 = F___memset(m, v10, int32(0), int32(32))
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
	v65 = v14
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
		v65 = v4
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
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v65 = v61
	goto L5
L20:
	;
	v65 = v46
	goto L5
L21:
	;
	goto L22
L22:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
	v61 = v46 + int32(1)
	if v59 != 0 {
		v46 = v61
		v47 = v59
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v73 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v73)
	v78 = v71 + int32(1)
	goto L26
L25:
	;
	v78 = int32(0)
	goto L26
L26:
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
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v176 int32
	_ = v176
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
	return v176
L4:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v155 = F_WALRead(m, l0, l4, l1, v148, v152, v12+int32(-40))
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
	v176 = int32(-1)
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
	F_WALReadRaiseError(m, v12+int32(-40))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v163 = int32(_a_F_summarizer_read_local_xlog_page_4)
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_summarizer_read_local_xlog_page[0])) = v165 + int32(1)
	v176 = v148
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v224 int32
	_ = v224
	v10 = m.G0
	v11 = int32(16)
	v12 = v10 - v11
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l2
	v15 = m.G0
	v17 = v15 - v11
	m.G0 = v17
	if base.Ui32(int32(1023)) < base.Ui32(l0) {
		v224 = int32(16)
		m.G0 = v17 + v224
		m.G0 = v12 + v224
		return
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[0]))
		if v22&(int32(1)<<(uint(l0&int32(7))%32)) == int32(0) {
			v224 = int32(16)
			m.G0 = v17 + v224
			m.G0 = v12 + v224
			return
		} else {
			v30 = m.G0
			v32 = v30 - int32(1168)
			m.G0 = v32
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[1]))
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[2]))
			if v37 < int32(0) {
				v40 = int32(0)
				v45 = F_socket(m, int32(1), int32(_a_F_syslog_0), v40)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, _c_F_syslog[2])) = v45
				if v40 <= v45 {
					v49 = F_connect(m, v45)
					mBase = m.M
				} else {
				}
			} else {
			}
			v51 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[3]))
			v52 = F___time(m)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, uint32(v32)+1144)) = v52
			v57 = v32 + int32(1100)
			v58 = F___gmtime_r(m, v32+int32(1144), v57)
			mBase = m.M
			v66 = F___strftime_l(m, v32+int32(1152), int32(16), int32(_a_F_syslog_1), v57, int32(_a_F_syslog_2))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				if l0&int32(1016) != 0 {
					v71 = int32(0)
				} else {
					v71 = v51
				}
				v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
				if v75&int32(1) != 0 {
					v78 = int32(42)
				} else {
					v78 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v78
				v81 = base.B2i32(v78 == int32(0))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v81 + int32(_a_F_syslog_3)
				*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v81 + int32(_a_F_syslog_4)
				*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = int32(_a_F_syslog_5)
				*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v71 | l0
				*(*int32)(unsafe.Add(mBase, uint32(v32)+40)) = v32 + int32(60)
				*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v32 + int32(1152)
				v103 = F_snprintf(m, v32-int32(-64), int32(1024), int32(_a_F_syslog_6), v32+int32(32))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_syslog[1])) = v35
					v111 = int32(1024) - v103
					v112 = F_vsnprintf(m, v103+(v32-int32(-64)), v111, l1, l2)
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return
					} else {
						if v112 < int32(0) {
							m.G0 = v32 + int32(1168)
							v224 = int32(16)
							m.G0 = v17 + v224
							m.G0 = v12 + v224
							return
						} else {
							if base.Ui32(v111) <= base.Ui32(v112) {
								v119 = int32(1023)
							} else {
								v119 = v112 + v103
							}
							v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v119)+63)))
							if v121 != int32(10) {
								v127 = int32(10)
								*(*uint8)(unsafe.Add(mBase, uint32(v32-int32(-64)+v119))) = uint8(v127)
								v131 = v119 + int32(1)
							} else {
								v131 = v119
							}
							v133 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[2]))
							v136 = int32(0)
							v138 = F_sendto(m, v133, v32-int32(-64), v131, v136, v136)
							mBase = m.M
							if int32(0) <= v138 {
								v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
								if v196&int32(32) == int32(0) {
									m.G0 = v32 + int32(1168)
									v224 = int32(16)
									m.G0 = v17 + v224
									m.G0 = v12 + v224
									return
								} else {
									v201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
									*(*int32)(unsafe.Add(mBase, uint32(v32))) = v131 - v201
									*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v201 + (v32 - int32(-64))
									F_dprintf(m, int32(2), v32)
									mBase = m.M
									v210 = m.ExcPending
									if v210 != 0 {
										return
									} else {
										m.G0 = v32 + int32(1168)
										v224 = int32(16)
										m.G0 = v17 + v224
										m.G0 = v12 + v224
										return
									}
								}
							} else {
								v141 = int32(1)
								v143 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[1]))
								if base.Ui32(v143-int32(14)) < base.Ui32(int32(2)) {
									v152 = v141
								} else {
									if v143 == int32(53) {
										v152 = v141
									} else {
										v152 = base.B2i32(v143 == int32(64))
									}
								}
								if v152 == int32(0) {
									v170 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
									if v170&int32(2) == int32(0) {
										v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
										if v196&int32(32) == int32(0) {
											m.G0 = v32 + int32(1168)
											v224 = int32(16)
											m.G0 = v17 + v224
											m.G0 = v12 + v224
											return
										} else {
											v201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v32))) = v131 - v201
											*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v201 + (v32 - int32(-64))
											F_dprintf(m, int32(2), v32)
											mBase = m.M
											v210 = m.ExcPending
											if v210 != 0 {
												return
											} else {
												m.G0 = v32 + int32(1168)
												v224 = int32(16)
												m.G0 = v17 + v224
												m.G0 = v12 + v224
												return
											}
										}
									} else {
										v177 = int32(0)
										v178 = F_open(m, int32(_a_F_syslog_7), int32(_a_F_syslog_8), v177)
										mBase = m.M
										if v178 < v177 {
											v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
											if v196&int32(32) == int32(0) {
												m.G0 = v32 + int32(1168)
												v224 = int32(16)
												m.G0 = v17 + v224
												m.G0 = v12 + v224
												return
											} else {
												v201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v32))) = v131 - v201
												*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v201 + (v32 - int32(-64))
												F_dprintf(m, int32(2), v32)
												mBase = m.M
												v210 = m.ExcPending
												if v210 != 0 {
													return
												} else {
													m.G0 = v32 + int32(1168)
													v224 = int32(16)
													m.G0 = v17 + v224
													m.G0 = v12 + v224
													return
												}
											}
										} else {
											v181 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v131 - v181
											*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v181 + (v32 - int32(-64))
											F_dprintf(m, v178, v32+int32(16))
											mBase = m.M
											v191 = m.ExcPending
											if v191 != 0 {
												return
											} else {
												v192 = F_close(m, v178)
												mBase = m.M
												v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
												if v196&int32(32) == int32(0) {
													m.G0 = v32 + int32(1168)
													v224 = int32(16)
													m.G0 = v17 + v224
													m.G0 = v12 + v224
													return
												} else {
													v201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v32))) = v131 - v201
													*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v201 + (v32 - int32(-64))
													F_dprintf(m, int32(2), v32)
													mBase = m.M
													v210 = m.ExcPending
													if v210 != 0 {
														return
													} else {
														m.G0 = v32 + int32(1168)
														v224 = int32(16)
														m.G0 = v17 + v224
														m.G0 = v12 + v224
														return
													}
												}
											}
										}
									}
								} else {
									v156 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[2]))
									v157 = F_connect(m, v156)
									mBase = m.M
									if v157 < int32(0) {
										v170 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
										if v170&int32(2) == int32(0) {
											v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
											if v196&int32(32) == int32(0) {
												m.G0 = v32 + int32(1168)
												v224 = int32(16)
												m.G0 = v17 + v224
												m.G0 = v12 + v224
												return
											} else {
												v201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v32))) = v131 - v201
												*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v201 + (v32 - int32(-64))
												F_dprintf(m, int32(2), v32)
												mBase = m.M
												v210 = m.ExcPending
												if v210 != 0 {
													return
												} else {
													m.G0 = v32 + int32(1168)
													v224 = int32(16)
													m.G0 = v17 + v224
													m.G0 = v12 + v224
													return
												}
											}
										} else {
											v177 = int32(0)
											v178 = F_open(m, int32(_a_F_syslog_7), int32(_a_F_syslog_8), v177)
											mBase = m.M
											if v178 < v177 {
												v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
												if v196&int32(32) == int32(0) {
													m.G0 = v32 + int32(1168)
													v224 = int32(16)
													m.G0 = v17 + v224
													m.G0 = v12 + v224
													return
												} else {
													v201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v32))) = v131 - v201
													*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v201 + (v32 - int32(-64))
													F_dprintf(m, int32(2), v32)
													mBase = m.M
													v210 = m.ExcPending
													if v210 != 0 {
														return
													} else {
														m.G0 = v32 + int32(1168)
														v224 = int32(16)
														m.G0 = v17 + v224
														m.G0 = v12 + v224
														return
													}
												}
											} else {
												v181 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v131 - v181
												*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v181 + (v32 - int32(-64))
												F_dprintf(m, v178, v32+int32(16))
												mBase = m.M
												v191 = m.ExcPending
												if v191 != 0 {
													return
												} else {
													v192 = F_close(m, v178)
													mBase = m.M
													v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
													if v196&int32(32) == int32(0) {
														m.G0 = v32 + int32(1168)
														v224 = int32(16)
														m.G0 = v17 + v224
														m.G0 = v12 + v224
														return
													} else {
														v201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
														*(*int32)(unsafe.Add(mBase, uint32(v32))) = v131 - v201
														*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v201 + (v32 - int32(-64))
														F_dprintf(m, int32(2), v32)
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return
														} else {
															m.G0 = v32 + int32(1168)
															v224 = int32(16)
															m.G0 = v17 + v224
															m.G0 = v12 + v224
															return
														}
													}
												}
											}
										}
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, _c_F_syslog[2]))
										v164 = int32(0)
										v166 = F_sendto(m, v161, v32-int32(-64), v131, v164, v164)
										mBase = m.M
										if int32(0) <= v166 {
											v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
											if v196&int32(32) == int32(0) {
												m.G0 = v32 + int32(1168)
												v224 = int32(16)
												m.G0 = v17 + v224
												m.G0 = v12 + v224
												return
											} else {
												v201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v32))) = v131 - v201
												*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v201 + (v32 - int32(-64))
												F_dprintf(m, int32(2), v32)
												mBase = m.M
												v210 = m.ExcPending
												if v210 != 0 {
													return
												} else {
													m.G0 = v32 + int32(1168)
													v224 = int32(16)
													m.G0 = v17 + v224
													m.G0 = v12 + v224
													return
												}
											}
										} else {
											v170 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
											if v170&int32(2) == int32(0) {
												v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
												if v196&int32(32) == int32(0) {
													m.G0 = v32 + int32(1168)
													v224 = int32(16)
													m.G0 = v17 + v224
													m.G0 = v12 + v224
													return
												} else {
													v201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v32))) = v131 - v201
													*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v201 + (v32 - int32(-64))
													F_dprintf(m, int32(2), v32)
													mBase = m.M
													v210 = m.ExcPending
													if v210 != 0 {
														return
													} else {
														m.G0 = v32 + int32(1168)
														v224 = int32(16)
														m.G0 = v17 + v224
														m.G0 = v12 + v224
														return
													}
												}
											} else {
												v177 = int32(0)
												v178 = F_open(m, int32(_a_F_syslog_7), int32(_a_F_syslog_8), v177)
												mBase = m.M
												if v178 < v177 {
													v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
													if v196&int32(32) == int32(0) {
														m.G0 = v32 + int32(1168)
														v224 = int32(16)
														m.G0 = v17 + v224
														m.G0 = v12 + v224
														return
													} else {
														v201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
														*(*int32)(unsafe.Add(mBase, uint32(v32))) = v131 - v201
														*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v201 + (v32 - int32(-64))
														F_dprintf(m, int32(2), v32)
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return
														} else {
															m.G0 = v32 + int32(1168)
															v224 = int32(16)
															m.G0 = v17 + v224
															m.G0 = v12 + v224
															return
														}
													}
												} else {
													v181 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v131 - v181
													*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v181 + (v32 - int32(-64))
													F_dprintf(m, v178, v32+int32(16))
													mBase = m.M
													v191 = m.ExcPending
													if v191 != 0 {
														return
													} else {
														v192 = F_close(m, v178)
														mBase = m.M
														v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_syslog[4])))
														if v196&int32(32) == int32(0) {
															m.G0 = v32 + int32(1168)
															v224 = int32(16)
															m.G0 = v17 + v224
															m.G0 = v12 + v224
															return
														} else {
															v201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
															*(*int32)(unsafe.Add(mBase, uint32(v32))) = v131 - v201
															*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v201 + (v32 - int32(-64))
															F_dprintf(m, int32(2), v32)
															mBase = m.M
															v210 = m.ExcPending
															if v210 != 0 {
																return
															} else {
																m.G0 = v32 + int32(1168)
																v224 = int32(16)
																m.G0 = v17 + v224
																m.G0 = v12 + v224
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
