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
	v20 = *(*int32)(unsafe.Add(mBase, _consts[846]))
	if l0 == v3 {
		v24 = *(*int32)(unsafe.Add(mBase, _consts[44]))
		v28 = F_LWLockAcquire(m, v24+int32(768), int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, _consts[44]))
			v35 = F_LWLockAcquire(m, v31+int32(640), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[848])))
				if v38 <= int32(0) {
					v96 = v38
					v97 = v37
					v99 = v3
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[849])))
					v51 = int32(0)
					v54 = v38
					v55 = v37
					v57 = v3
					v60 = v37 - int32(2048)
					for {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v48+v51<<(uint(int32(2))%32))))
						v70 = v20 + int32(65560) + v67<<(uint(int32(4))%32)
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
								if v73 < l1+v37-int32(4096) {
									v75 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)) = uint8(v75)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[848])))
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
						v120 = v20 + int32(65564)
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[849])))
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
					v221 = *(*int32)(unsafe.Add(mBase, _consts[44]))
					F_LWLockRelease(m, v221+int32(640))
					mBase = m.M
					v225 = m.ExcPending
					if v225 != 0 {
						return
					} else {
						v227 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
									F_errmsg_internal(m, int32(480094), v17)
									mBase = m.M
									v244 = m.ExcPending
									if v244 != 0 {
										return
									} else {
										F_errfinish(m, int32(493780), int32(673), int32(348760))
										mBase = m.M
										v249 = m.ExcPending
										if v249 != 0 {
											return
										} else {
											v251 = F_SendProcSignal(m, v219, int32(0), (v99-v20-int32(65560))>>(uint(int32(4))%32))
											mBase = m.M
											v252 = m.ExcPending
											if v252 != 0 {
												return
											} else {
												if l0 == int32(0) {
													m.G0 = v17 + int32(16)
													return
												} else {
													v256 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
									v251 = F_SendProcSignal(m, v219, int32(0), (v99-v20-int32(65560))>>(uint(int32(4))%32))
									mBase = m.M
									v252 = m.ExcPending
									if v252 != 0 {
										return
									} else {
										if l0 == int32(0) {
											m.G0 = v17 + int32(16)
											return
										} else {
											v256 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
					v263 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
							v269 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
		v31 = *(*int32)(unsafe.Add(mBase, _consts[44]))
		v35 = F_LWLockAcquire(m, v31+int32(640), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[848])))
			if v38 <= int32(0) {
				v96 = v38
				v97 = v37
				v99 = v3
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[849])))
				v51 = int32(0)
				v54 = v38
				v55 = v37
				v57 = v3
				v60 = v37 - int32(2048)
				for {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v48+v51<<(uint(int32(2))%32))))
					v70 = v20 + int32(65560) + v67<<(uint(int32(4))%32)
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
							if v73 < l1+v37-int32(4096) {
								v75 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)) = uint8(v75)
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[848])))
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
					v120 = v20 + int32(65564)
					v121 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[849])))
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
				v221 = *(*int32)(unsafe.Add(mBase, _consts[44]))
				F_LWLockRelease(m, v221+int32(640))
				mBase = m.M
				v225 = m.ExcPending
				if v225 != 0 {
					return
				} else {
					v227 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
								F_errmsg_internal(m, int32(480094), v17)
								mBase = m.M
								v244 = m.ExcPending
								if v244 != 0 {
									return
								} else {
									F_errfinish(m, int32(493780), int32(673), int32(348760))
									mBase = m.M
									v249 = m.ExcPending
									if v249 != 0 {
										return
									} else {
										v251 = F_SendProcSignal(m, v219, int32(0), (v99-v20-int32(65560))>>(uint(int32(4))%32))
										mBase = m.M
										v252 = m.ExcPending
										if v252 != 0 {
											return
										} else {
											if l0 == int32(0) {
												m.G0 = v17 + int32(16)
												return
											} else {
												v256 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
								v251 = F_SendProcSignal(m, v219, int32(0), (v99-v20-int32(65560))>>(uint(int32(4))%32))
								mBase = m.M
								v252 = m.ExcPending
								if v252 != 0 {
									return
								} else {
									if l0 == int32(0) {
										m.G0 = v17 + int32(16)
										return
									} else {
										v256 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
				v263 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
						v269 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[655]))
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
	*(*uint8)(unsafe.Add(mBase, _consts[133])) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, _consts[130])) = l0
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[132])))
	if v8 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[3])) = l0
		*(*int32)(unsafe.Add(mBase, _consts[131])) = l0
		if v2 != 0 {
			v18 = int32(273318)
		} else {
			v18 = int32(339246)
		}
		F_SetConfigOption(m, int32(217591), v18, int32(0), int32(1))
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
	F_errmsg(m, int32(439779), int32(0))
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
	F_errfinish(m, int32(26994), int32(19430), int32(76263))
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
	F_errmsg_internal(m, int32(476953), v11)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(26994), int32(19434), int32(76263))
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[812]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(1)
	if v6 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[812]))
		F_s_lock(m, v10, int32(492990), int32(365), int32(226860))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[812]))
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
		v17 = *(*int32)(unsafe.Add(mBase, _consts[812]))
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
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v32
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
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v77
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
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v962 int32
	_ = v962
	var v984 int32
	_ = v984
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1115 int32
	_ = v1115
	var v1128 int32
	_ = v1128
	var v1148 int32
	_ = v1148
	var v1159 int32
	_ = v1159
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1187 int32
	_ = v1187
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
	v1187 = v6
	goto L3
L3:
	;
	m.G0 = v29 + int32(128)
	return v1187
L4:
	;
	v1178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v1159))) = uint8(v1178)
	v1187 = v1176
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
	if l1 == v1128 {
		goto L311
	} else {
		goto L312
	}
L7:
	;
	goto L6
L8:
	;
	if base.Ui32(v1115) < base.Ui32(l1) {
		v33 = v1091 + int32(1)
		v40 = v1115
		goto L5
	} else {
		goto L310
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
	v1091 = v33
	v1115 = v40 + int32(1)
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
	v1159 = v40
	v1176 = v40
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
	if int32(1)<<(uint(v104)%32)&int32(4194329) == int32(0) {
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
		v852 = int32(0)
		goto L32
	case 28:
		goto L76
	case 29:
		goto L74
	case 30:
		goto L73
	case 31:
		v786 = int32(27036)
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
		v728 = int32(131112)
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
	if v852 == int32(0) {
		v1128 = v40
		goto L7
	} else {
		goto L268
	}
L33:
	;
	v850 = F_strlen(m, v849)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v850
	v852 = v849
	goto L32
L34:
	;
	v849 = int32(670315)
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
	v728 = int32(131115)
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
	v669 = v139 | int32(131072)
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
	v852 = int32(690252)
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
	v728 = int32(131114)
	goto L39
L50:
	;
	v728 = int32(131113)
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
	v786 = int32(524704)
	goto L38
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(1)
	v852 = int32(757755)
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
	v786 = int32(532466)
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
	v852 = int32(757606)
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
	v786 = int32(467237)
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
	v669 = v142 + int32(131079)
	goto L41
L80:
	;
	v669 = v147 + int32(131086)
	goto L41
L81:
	;
	v669 = v152 + int32(131098)
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
	v323 = int32(131111)
	goto L126
L125:
	;
	v323 = int32(131110)
	goto L126
L126:
	;
	v669 = v323
	goto L41
L127:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v354<<(uint(int32(2))%32))+uint32(_consts[1384])))
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
	v466 = base.I64_extend_i32_s(v360*int32(31536000) + v380*int32(86400) + int32(2087447296))
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
	v475 = v471 + int32(86400)
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
	v621 = F_snprintf(m, v125, int32(100), int32(430423), v129+int32(48))
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
	v852 = v125
	goto L32
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(0)
	v852 = int32(757756)
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
	v650 = F_snprintf(m, v125, v635, int32(432420), v129-int32(-64))
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
	v852 = v125
	goto L32
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(0)
	v852 = int32(757756)
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
	v680 = int32(65535)
	v681 = v669 & v680
	if v681 != v680 {
		goto L204
	} else {
		goto L205
	}
L201:
	;
	v677 = int32(556686)
	goto L203
L202:
	;
	v677 = int32(535009)
	goto L203
L203:
	;
	v726 = v677
	goto L197
L204:
	;
	v694 = int32(757756)
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
	v693 = int32(545070)
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
	v706 = int32(4092248)
	goto L211
L216:
	;
	v706 = int32(4092256)
	goto L211
L217:
	;
	v706 = int32(4092576)
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
	v739 = int32(65535)
	v740 = v728 & v739
	if v740 != v739 {
		goto L232
	} else {
		goto L233
	}
L229:
	;
	v736 = int32(556686)
	goto L231
L230:
	;
	v736 = int32(535009)
	goto L231
L231:
	;
	v785 = v736
	goto L225
L232:
	;
	v753 = int32(757756)
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
	v752 = int32(545070)
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
	v765 = int32(4092248)
	goto L239
L244:
	;
	v765 = int32(4092256)
	goto L239
L245:
	;
	v765 = int32(4092576)
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
	v852 = v793
	goto L32
L257:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v129)+8)) = v812
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v803
	v842 = F_snprintf(m, v125, int32(100), int32(430259), v129)
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
	v835 = F_snprintf(m, v125, int32(100), int32(430266), v129+int32(32))
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
	v826 = F_snprintf(m, v125, int32(100), int32(430874), v129+int32(16))
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
	v852 = v125
	goto L32
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v835
	v852 = v125
	goto L32
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v842
	v852 = v125
	goto L32
L268:
	;
	if v114 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v1080 = l1 - v1063
	if base.Ui32(v1062) < base.Ui32(v1080) {
		goto L303
	} else {
		goto L304
	}
L270:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	v1062 = v872
	v1063 = v40
	v1065 = v852
	goto L269
L271:
	;
	goto L272
L272:
	;
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852))))
	switch v873 - int32(43) {
	case 0, 2:
		goto L274
	default:
		goto L275
	}
L273:
	;
	if v883&int32(255) != int32(48) {
		v935 = v885
		v938 = v884
		goto L276
	} else {
		goto L277
	}
L274:
	;
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+1)))
	v878 = int32(1)
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	v883 = v877
	v884 = v852 + v878
	v885 = v880 - v878
	goto L273
L275:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	v883 = v873
	v884 = v852
	v885 = v876
	goto L273
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+124)) = v935
	v962 = int32(0)
	goto L282
L277:
	;
	v898 = v885
	v901 = v884
	goto L278
L278:
	;
	v916 = int32(*(*int8)(unsafe.Add(mBase, uint32(v901)+1)))
	if base.Ui32(int32(9)) < base.Ui32(v916-int32(48)) {
		v935 = v898
		v938 = v901
		goto L276
	} else {
		goto L280
	}
L279:
	;
	v935 = v924
	v938 = v922
	goto L276
L280:
	;
	v921 = int32(1)
	v922 = v901 + v921
	v924 = v898 - v921
	if v916 == int32(48) {
		v898 = v924
		v901 = v922
		goto L278
	} else {
		goto L281
	}
L281:
	;
	goto L279
L282:
	;
	v984 = int32(*(*int8)(unsafe.Add(mBase, uint32(v962+v938))))
	if base.Ui32(v984-int32(48)) < base.Ui32(int32(10)) {
		v962 = v962 + int32(1)
		goto L282
	} else {
		goto L284
	}
L283:
	;
	if base.Ui32(v935) < base.Ui32(v114) {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	goto L283
L285:
	;
	v990 = v114
	goto L287
L286:
	;
	v990 = v935
	goto L287
L287:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v992 < int32(-1900) {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	if base.Ui32(v1015) <= base.Ui32(v935) {
		v1062 = v935
		v1063 = v1016
		v1065 = v938
		goto L269
	} else {
		goto L297
	}
L289:
	;
	v1009 = int32(45)
	goto L291
L290:
	;
	if v81 != int32(43) {
		v1015 = v990
		v1016 = v40
		goto L288
	} else {
		goto L292
	}
L291:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v40))) = uint8(v1009)
	v1011 = int32(1)
	v1015 = v990 - v1011
	v1016 = v40 + v1011
	goto L288
L292:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1002))))
	if v1003 == int32(67) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1006 = int32(3)
	goto L295
L294:
	;
	v1006 = int32(5)
	goto L295
L295:
	;
	if base.Ui32(v990-v935+v962) < base.Ui32(v1006) {
		v1015 = v990
		v1016 = v40
		goto L288
	} else {
		goto L296
	}
L296:
	;
	v1009 = int32(43)
	goto L291
L297:
	;
	if base.Ui32(l1) <= base.Ui32(v1016) {
		v1062 = v935
		v1063 = v1016
		v1065 = v938
		goto L269
	} else {
		goto L298
	}
L298:
	;
	v1026 = v1015
	v1028 = v1016
	goto L299
L299:
	;
	v1046 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v1028))) = uint8(v1046)
	v1048 = int32(1)
	v1049 = v1028 + v1048
	v1051 = v1026 - v1048
	if base.Ui32(v1051) <= base.Ui32(v935) {
		v1062 = v935
		v1063 = v1049
		v1065 = v938
		goto L269
	} else {
		goto L301
	}
L300:
	;
	v1062 = v935
	v1063 = v1049
	v1065 = v938
	goto L269
L301:
	;
	if base.Ui32(v1049) < base.Ui32(l1) {
		v1026 = v1051
		v1028 = v1049
		goto L299
	} else {
		goto L302
	}
L302:
	;
	goto L300
L303:
	;
	v1082 = v1062
	goto L305
L304:
	;
	v1082 = v1080
	goto L305
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+124)) = v1082
	if v1082 != 0 {
		goto L307
	} else {
		goto L308
	}
L306:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	v1091 = v123
	v1115 = v1087 + v1063
	goto L8
L307:
	;
	v1085 = F__emscripten_memcpy_bulkmem(m, l0+v1063, v1065, v1082)
	mBase = m.M
	goto L309
L308:
	;
	goto L309
L309:
	;
	goto L306
L310:
	;
	v1128 = v1115
	goto L7
L311:
	;
	v1148 = l1 - int32(1)
	goto L313
L312:
	;
	v1148 = v1128
	goto L313
L313:
	;
	v1159 = v1148
	v1176 = int32(0)
	goto L4
}
func F___syscall_ret(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	if base.Ui32(int32(-4095)) <= base.Ui32(l0) {
		*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0) - l0
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
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var __phi57 int32
	_ = __phi57
	var v58 int32
	_ = v58
	var __phi58 int32
	_ = __phi58
	var v60 int32
	_ = v60
	var __phi60 int32
	_ = __phi60
	var v61 int32
	_ = v61
	var __phi61 int32
	_ = __phi61
	var v62 int32
	_ = v62
	var __phi62 int32
	_ = __phi62
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v42 = v12 & int32(255)
	if base.Ui32(v42-int32(97)) < base.Ui32(int32(26)) {
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
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v37)
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
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v49)
	v51 = int32(1)
	v53 = l1 + v51
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v54 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v49 = v42 & int32(95)
	goto L12
L11:
	;
	v49 = v42
	goto L12
L12:
	;
	goto L9
L13:
	;
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v163))) = uint8(v168)
	return
L14:
	;
	__phi57 = v10
	__phi58 = v54
	__phi60 = v53
	__phi61 = v51
	__phi62 = v10 + int32(1)
	v57 = __phi57
	v58 = __phi58
	v60 = __phi60
	v61 = __phi61
	v62 = __phi62
	goto L17
L15:
	;
	v149 = v53
	v150 = v51
	goto L16
L16:
	;
	v156 = int32(4) - v150
	v158 = F__emscripten_memset_bulkmem(m, v149, base.I32_extend8_s(int32(48)), v156)
	mBase = m.M
	goto L46
L17:
	;
	if base.Ui32(int32(25)) < base.Ui32((v58|int32(32)-int32(97))&int32(255)) {
		v132 = v60
		v133 = v61
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if int32(3) < v133 {
		v163 = v132
		goto L13
	} else {
		goto L45
	}
L19:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v137 != 0 {
		goto L41
	} else {
		goto L42
	}
L20:
	;
	v74 = v58 & int32(255)
	if base.Ui32(v74-int32(97)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if base.Ui32(v95-int32(97)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v82 = base.I32_extend8_s(v81)
	v86 = base.B2i32(base.Ui32(int32(25)) < base.Ui32(v82-int32(65)))
	if v86 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v81 = v74 & int32(95)
	goto L25
L24:
	;
	v81 = v74
	goto L25
L25:
	;
	goto L22
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1320]))))
	v92 = v91
	goto L21
L27:
	;
	goto L28
L28:
	;
	v92 = v81
	goto L21
L29:
	;
	v103 = base.I32_extend8_s(v102)
	if base.Ui32(v103-int32(65)) <= base.Ui32(int32(25)) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v102 = v95 & int32(95)
	goto L32
L31:
	;
	v102 = v95
	goto L32
L32:
	;
	goto L29
L33:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+uint32(_consts[1320]))))
	v111 = v110
	goto L35
L34:
	;
	v111 = v102
	goto L35
L35:
	;
	if v92&int32(255) == v111&int32(255) {
		v132 = v60
		v133 = v61
		goto L19
	} else {
		goto L36
	}
L36:
	;
	if v86 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1320]))))
	v120 = v119
	goto L39
L38:
	;
	v120 = v81
	goto L39
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v120)
	if v120&int32(255) == int32(48) {
		v132 = v60
		v133 = v61
		goto L19
	} else {
		goto L40
	}
L40:
	;
	v126 = int32(1)
	v132 = v60 + v126
	v133 = v61 + v126
	goto L19
L41:
	;
	if v133 < int32(4) {
		__phi57 = v62
		__phi58 = v137
		__phi60 = v132
		__phi61 = v133
		__phi62 = v62 + int32(1)
		v57 = __phi57
		v58 = __phi58
		v60 = __phi60
		v61 = __phi61
		v62 = __phi62
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
	v149 = v132
	v150 = v133
	goto L16
L46:
	;
	v163 = v158 + v156
	goto L13
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
	*(*int32)(unsafe.Add(mBase, _consts[333])) = v134
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
	v105 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	v112 = v105
	goto L3
L8:
	;
	v25 = v19 + int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	if v27 <= v25 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[333]))
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
	v35 = *(*int32)(unsafe.Add(mBase, _consts[39]))
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
	v43 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(0)
	v48 = int32(4599216)
	v51 = *(*int64)(unsafe.Add(mBase, _consts[819]))
	v52 = *(*int64)(unsafe.Add(mBase, _consts[818]))
	v53 = v51 ^ v52
	*(*int64)(unsafe.Add(mBase, _consts[818])) = base.I64_rotl(v53, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[819])) = v53<<(uint(int64(16))%64) ^ base.I64_rotl(v51, int64(24)) ^ v53
	v74 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v51*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L20
L19:
	;
	v84 = v83 + v39
	if int32(1000000) < v84 {
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
	F_errmsg(m, int32(680352), v21+int32(16))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(494635), int32(4959), int32(248535))
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
	F_errmsg(m, int32(239874), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(494635), int32(4802), int32(248535))
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
	F_errmsg(m, int32(239790), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(494635), int32(4806), int32(248535))
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
	F_errmsg(m, int32(115188), int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(494635), int32(4810), int32(248535))
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
	F_errmsg(m, int32(394935), v21)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(494635), int32(4831), int32(248535))
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
	F_errmsg(m, int32(680352), v21-int32(-64))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(494635), int32(4842), int32(248535))
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
	F_errmsg(m, int32(192369), v21+int32(48))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(494635), int32(4864), int32(248535))
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
	F_errmsg(m, int32(714571), v21+int32(32))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(494635), int32(4898), int32(248535))
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
	v22 = v14 + v19&int32(32767)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v27 = int32(base.Ui32(v23)>>(uint(int32(3))%32)) & int32(8191)
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
		v37 = v37 + v43&int32(8191)
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
	F_errmsg_internal(m, int32(384495), v11)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(493341), int32(68), int32(316051))
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v378 int32
	_ = v378
	var v397 int32
	_ = v397
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	v22 = F_strlen(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = int32(16777216)
	v25 = F_pg_hmac_create(m, l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(80)
	return v397
L2:
	;
	return int32(0)
L3:
	;
	if v25 == int32(0) {
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
	v54 = F_pg_hmac_init(m, v25, l0, v22)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L23
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(13904)
	v397 = int32(-1)
	goto L1
L8:
	;
	goto L7
L20:
	;
	F_pg_hmac_free(m, v25)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L2
	} else {
		goto L113
	}
L21:
	;
	if l2 != 0 {
		goto L57
	} else {
		goto L58
	}
L22:
	;
	if v25 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L23:
	;
	if v54 < int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v25 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v72 < int32(0) {
		goto L22
	} else {
		goto L32
	}
L26:
	;
	v72 = int32(-1)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v62 = F_pg_cryptohash_update(m, v61, l3, l4)
	mBase = m.M
	if int32(0) <= v62 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v72 = int32(0)
	goto L25
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(2)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v69 = F_pg_cryptohash_error(m, v68)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v69
	v72 = int32(-1)
	goto L25
L32:
	;
	if v25 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v92 < int32(0) {
		goto L22
	} else {
		goto L40
	}
L34:
	;
	v92 = int32(-1)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v82 = F_pg_cryptohash_update(m, v81, v20+int32(76), int32(4))
	mBase = m.M
	if int32(0) <= v82 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v92 = int32(0)
	goto L33
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(2)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v89 = F_pg_cryptohash_error(m, v88)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v89
	v92 = int32(-1)
	goto L33
L40:
	;
	v95 = F_pg_hmac_final(m, v25, v20, l2)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	if int32(0) <= v95 {
		goto L21
	} else {
		goto L42
	}
L42:
	;
	goto L22
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v118
	goto L20
L44:
	;
	v118 = int32(13904)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v103 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v115 = v103
	goto L49
L48:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v107 == int32(2) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v118 = v115
	goto L43
L50:
	;
	v110 = int32(212449)
	goto L52
L51:
	;
	v110 = int32(130268)
	goto L52
L52:
	;
	if v107 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v113 = int32(13904)
	goto L55
L54:
	;
	v113 = v110
	goto L55
L55:
	;
	v115 = v113
	goto L49
L56:
	;
	if int32(2) <= l5 {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v120 = F__emscripten_memcpy_bulkmem(m, l6, v20, l2)
	mBase = m.M
	v121 = v120
	goto L59
L58:
	;
	v121 = l6
	goto L59
L59:
	;
	goto L56
L60:
	;
	if v25 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L61:
	;
	v127 = l2 & int32(3)
	v145 = int32(1)
	goto L64
L62:
	;
	goto L63
L63:
	;
	F_pg_hmac_free(m, v25)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L2
	} else {
		goto L99
	}
L64:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v149 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L63
L66:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v152 = F_pg_hmac_init(m, v25, l0, v22)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v152 < int32(0) {
		goto L60
	} else {
		goto L71
	}
L71:
	;
	if v25 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v170 < int32(0) {
		goto L60
	} else {
		goto L79
	}
L73:
	;
	v170 = int32(-1)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v160 = F_pg_cryptohash_update(m, v159, v20, l2)
	mBase = m.M
	if int32(0) <= v160 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v170 = int32(0)
	goto L72
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(2)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v167 = F_pg_cryptohash_error(m, v166)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v167
	v170 = int32(-1)
	goto L72
L79:
	;
	v175 = F_pg_hmac_final(m, v25, v20+int32(32), l2)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	if v175 < int32(0) {
		goto L60
	} else {
		goto L81
	}
L81:
	;
	if l2 <= int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if l2 != 0 {
		goto L95
	} else {
		goto L96
	}
L83:
	;
	v181 = int32(0)
	if base.B2i32(base.Ui32(l2) < base.Ui32(int32(4))) == v181 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v187 = v181
	v192 = v181
	goto L87
L85:
	;
	v247 = v181
	goto L86
L86:
	;
	if v127 == int32(0) {
		goto L82
	} else {
		goto L90
	}
L87:
	;
	v203 = v187 + v121
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v206 = v20 + int32(32)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206+v187))))
	v209 = v204 ^ v208
	*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v209)
	v212 = v187 | int32(1)
	v213 = v121 + v212
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206+v212))))
	v219 = v214 ^ v218
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v219)
	v222 = v187 | int32(2)
	v223 = v121 + v222
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206+v222))))
	v229 = v224 ^ v228
	*(*uint8)(unsafe.Add(mBase, uint32(v223))) = uint8(v229)
	v232 = v187 | int32(3)
	v233 = v121 + v232
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206+v232))))
	v239 = v234 ^ v238
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v239)
	v241 = int32(4)
	v242 = v187 + v241
	v244 = v192 + v241
	if v244 != l2&int32(2147483644) {
		v187 = v242
		v192 = v244
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v247 = v242
	goto L86
L89:
	;
	goto L88
L90:
	;
	v266 = v247
	v268 = v181
	goto L91
L91:
	;
	v282 = v266 + v121
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(32)+v266))))
	v288 = v283 ^ v287
	*(*uint8)(unsafe.Add(mBase, uint32(v282))) = uint8(v288)
	v290 = int32(1)
	v293 = v268 + v290
	if v293 != v127 {
		v266 = v266 + v290
		v268 = v293
		goto L91
	} else {
		goto L93
	}
L92:
	;
	goto L82
L93:
	;
	goto L92
L94:
	;
	v317 = v145 + int32(1)
	if v317 != l5 {
		v145 = v317
		goto L64
	} else {
		goto L98
	}
L95:
	;
	v314 = F__emscripten_memcpy_bulkmem(m, v20, v20+int32(32), l2)
	mBase = m.M
	goto L97
L96:
	;
	goto L97
L97:
	;
	goto L94
L98:
	;
	goto L65
L99:
	;
	v397 = int32(0)
	goto L1
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v358
	goto L20
L101:
	;
	v358 = int32(13904)
	goto L100
L102:
	;
	goto L103
L103:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v343 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v355 = v343
	goto L106
L105:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v347 == int32(2) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v358 = v355
	goto L100
L107:
	;
	v350 = int32(212449)
	goto L109
L108:
	;
	v350 = int32(130268)
	goto L109
L109:
	;
	if v347 == int32(1) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v353 = int32(13904)
	goto L112
L111:
	;
	v353 = v350
	goto L112
L112:
	;
	v355 = v353
	goto L106
L113:
	;
	v397 = int32(-1)
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
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(13904)
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
						v82 = int32(13904)
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						if v67 != 0 {
							v79 = v67
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
							if v71 == int32(2) {
								v74 = int32(212449)
							} else {
								v74 = int32(130268)
							}
							if v71 == int32(1) {
								v77 = int32(13904)
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
						v46 = F_pg_cryptohash_update(m, v45, int32(22757), int32(10))
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
							v82 = int32(13904)
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
							if v67 != 0 {
								v79 = v67
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
								if v71 == int32(2) {
									v74 = int32(212449)
								} else {
									v74 = int32(130268)
								}
								if v71 == int32(1) {
									v77 = int32(13904)
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
									v82 = int32(13904)
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
									if v67 != 0 {
										v79 = v67
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
										if v71 == int32(2) {
											v74 = int32(212449)
										} else {
											v74 = int32(130268)
										}
										if v71 == int32(1) {
											v77 = int32(13904)
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int64
	_ = v275
	var v291 int32
	_ = v291
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
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
	v18 = int32(557250)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[562])))
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
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L65
	}
L4:
	;
	F_pg_cryptohash_free(m, v103)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L61
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
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L57
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
	v99 = *(*int32)(unsafe.Add(mBase, _consts[36]))
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
	v91 = F_psprintf(m, int32(582525), v8+int32(16))
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
	F_errmsg(m, int32(701473), v8)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(497634), int32(292), int32(100155))
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
	v108 = F_strlen(m, v95)
	mBase = m.M
	v109 = F_pg_cryptohash_update(m, v103, v95, v108)
	mBase = m.M
	if v109 < int32(0) {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v113 = F_pg_cryptohash_update(m, v103, v99+int32(257), int32(32))
	mBase = m.M
	if v113 < int32(0) {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v118 = F_pg_cryptohash_final(m, v103, int32(4414432), int32(32))
	mBase = m.M
	if v118 < int32(0) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	F_pg_cryptohash_free(m, v103)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v129 = base.I32_div_s(int32(18), int32(3))
	v131 = v129 << (uint(int32(2)) % 32)
	goto L36
L36:
	;
	v134 = F_palloc(m, v131+int32(1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L41
L38:
	;
	if v264 < int32(0) {
		goto L3
	} else {
		goto L56
	}
L39:
	;
	v255 = F___memset(m, v134, int32(0), v131)
	mBase = m.M
	v264 = int32(-1)
	goto L38
L40:
	;
	if v131 < v198-v134+int32(4) {
		goto L39
	} else {
		goto L52
	}
L41:
	;
	v143 = int32(4414432)
	v144 = int32(0)
	v147 = v134
	v148 = int32(2)
	goto L44
L43:
	;
	v264 = v198 - v134
	goto L38
L44:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v154 = v150<<(uint(v148<<(uint(int32(3))%32))%32) | v144
	if int32(0) < v148 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v199 != int32(2) {
		goto L40
	} else {
		goto L51
	}
L46:
	;
	v197 = v154
	v198 = v147
	v199 = v148 - int32(1)
	goto L48
L47:
	;
	if v131 < v147-v134+int32(4) {
		goto L39
	} else {
		goto L49
	}
L48:
	;
	v201 = v143 + int32(1)
	if v201 != int32(4414448) {
		v143 = v201
		v144 = v197
		v147 = v198
		v148 = v199
		goto L44
	} else {
		goto L50
	}
L49:
	;
	v163 = int32(63)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154&v163)+uint32(_consts[563]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v147)+3)) = uint8(v167)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v154)>>(uint(int32(6))%32))&v163)+uint32(_consts[563]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v147)+2)) = uint8(v175)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v154)>>(uint(int32(12))%32))&v163)+uint32(_consts[563]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)) = uint8(v183)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v154)>>(uint(int32(18))%32))&v163)+uint32(_consts[563]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v191)
	v197 = int32(0)
	v198 = v147 + int32(4)
	v199 = int32(2)
	goto L48
L50:
	;
	goto L45
L51:
	;
	goto L43
L52:
	;
	v219 = int32(63)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v197)>>(uint(int32(12))%32))&v219)+uint32(_consts[563]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)) = uint8(v223)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v197)>>(uint(int32(18))%32))&v219)+uint32(_consts[563]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v231)
	if v199 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v197)>>(uint(int32(6))%32))&int32(63))+uint32(_consts[563]))))
	v243 = v242
	goto L55
L54:
	;
	v243 = int32(61)
	goto L55
L55:
	;
	v244 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v198)+3)) = uint8(v244)
	*(*uint8)(unsafe.Add(mBase, uint32(v198)+2)) = uint8(v243)
	v264 = v198 + int32(4) - v134
	goto L38
L56:
	;
	v268 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v264+v134))) = uint8(v268)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(4096)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v134
	v274 = v11 - int32(-64)
	v275 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v274)+56)) = v275
	*(*int64)(unsafe.Add(mBase, uint32(v274)+48)) = v275
	*(*int64)(unsafe.Add(mBase, uint32(v274)+40)) = v275
	*(*int64)(unsafe.Add(mBase, uint32(v274)+32)) = v275
	*(*int64)(unsafe.Add(mBase, uint32(v274)+24)) = v275
	*(*int64)(unsafe.Add(mBase, uint32(v274)+16)) = v275
	*(*int64)(unsafe.Add(mBase, uint32(v274)+8)) = v275
	*(*int64)(unsafe.Add(mBase, uint32(v274))) = v275
	v291 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+192)) = uint8(v291)
	goto L16
L57:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(287881), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(497634), int32(267), int32(100155))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errmsg_internal(m, int32(98811), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(497634), int32(719), int32(107387))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errmsg_internal(m, int32(98811), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(497634), int32(728), int32(107387))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
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
	v28 = *(*int32)(unsafe.Add(mBase, _consts[40]))
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
	v32 = *(*int32)(unsafe.Add(mBase, _consts[569]))
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
	v39 = *(*int32)(unsafe.Add(mBase, _consts[569]))
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
	v51 = *(*int32)(unsafe.Add(mBase, _consts[311]))
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
	F_errmsg(m, int32(99504), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(499173), int32(350), int32(350084))
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
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
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
	var v266 int64
	_ = v266
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
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
	var v301 int32
	_ = v301
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
	var v332 int32
	_ = v332
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
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
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v820 int32
	_ = v820
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v892 int32
	_ = v892
	var v899 int32
	_ = v899
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v943 int32
	_ = v943
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1052 int32
	_ = v1052
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1117 int32
	_ = v1117
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1193 int32
	_ = v1193
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1226 int32
	_ = v1226
	var v1264 int32
	_ = v1264
	var v1290 int32
	_ = v1290
	var v1341 int32
	_ = v1341
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1382 int32
	_ = v1382
	var v1387 int32
	_ = v1387
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1497 int32
	_ = v1497
	var v1510 int32
	_ = v1510
	var v1519 int32
	_ = v1519
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1566 int32
	_ = v1566
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
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
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1770 int32
	_ = v1770
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1790 int32
	_ = v1790
	var v1798 int32
	_ = v1798
	var v1804 int32
	_ = v1804
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1830 int32
	_ = v1830
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1845 int32
	_ = v1845
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1857 int32
	_ = v1857
	var v1865 int32
	_ = v1865
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1904 int32
	_ = v1904
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1919 int32
	_ = v1919
	var v1926 int32
	_ = v1926
	var v1931 int32
	_ = v1931
	var v1934 int32
	_ = v1934
	var v1937 int32
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1953 int32
	_ = v1953
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2011 int32
	_ = v2011
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2023 int32
	_ = v2023
	var v2037 int32
	_ = v2037
	var v2074 int64
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2122 int32
	_ = v2122
	var v2127 int32
	_ = v2127
	var v2140 int32
	_ = v2140
	var v2146 int64
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2149 int32
	_ = v2149
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2179 int32
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2200 int32
	_ = v2200
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2222 int32
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2233 int32
	_ = v2233
	var v2237 int32
	_ = v2237
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
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
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2280 int32
	_ = v2280
	var v2282 int32
	_ = v2282
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2297 int32
	_ = v2297
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2322 int32
	_ = v2322
	var v2324 int32
	_ = v2324
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2335 int32
	_ = v2335
	var v2339 int32
	_ = v2339
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2354 int32
	_ = v2354
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2365 int32
	_ = v2365
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2374 int32
	_ = v2374
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
	var v2386 int32
	_ = v2386
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2438 int32
	_ = v2438
	var v2442 int32
	_ = v2442
	var v2446 int32
	_ = v2446
	var v2450 int32
	_ = v2450
	var v2454 int32
	_ = v2454
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2465 int32
	_ = v2465
	var v2476 int32
	_ = v2476
	var v2516 int32
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2532 int32
	_ = v2532
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2546 int32
	_ = v2546
	var v2550 int32
	_ = v2550
	var v2602 int32
	_ = v2602
	var v2604 int32
	_ = v2604
	var v2608 int32
	_ = v2608
	var v2615 int32
	_ = v2615
	var v2617 int32
	_ = v2617
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2631 int32
	_ = v2631
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2659 int32
	_ = v2659
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2679 int32
	_ = v2679
	var v2681 int32
	_ = v2681
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2692 int32
	_ = v2692
	var v2696 int32
	_ = v2696
	var v2700 int32
	_ = v2700
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2707 int32
	_ = v2707
	var v2708 int32
	_ = v2708
	var v2709 int32
	_ = v2709
	var v2711 int32
	_ = v2711
	var v2714 int32
	_ = v2714
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2718 int32
	_ = v2718
	var v2722 int32
	_ = v2722
	var v2726 int32
	_ = v2726
	var v2727 int32
	_ = v2727
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2739 int32
	_ = v2739
	var v2741 int32
	_ = v2741
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2756 int32
	_ = v2756
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2765 int32
	_ = v2765
	var v2766 int32
	_ = v2766
	var v2768 int32
	_ = v2768
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2772 int32
	_ = v2772
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2781 int32
	_ = v2781
	var v2783 int32
	_ = v2783
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2794 int32
	_ = v2794
	var v2798 int32
	_ = v2798
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2804 int32
	_ = v2804
	var v2805 int32
	_ = v2805
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2813 int32
	_ = v2813
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2818 int32
	_ = v2818
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2824 int32
	_ = v2824
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2843 int32
	_ = v2843
	var v2844 int32
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2849 int32
	_ = v2849
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2857 int32
	_ = v2857
	var v2859 int32
	_ = v2859
	var v2860 int32
	_ = v2860
	var v2861 int32
	_ = v2861
	var v2865 int32
	_ = v2865
	var v2866 int32
	_ = v2866
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2879 int32
	_ = v2879
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2886 int32
	_ = v2886
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2897 int32
	_ = v2897
	var v2901 int32
	_ = v2901
	var v2905 int32
	_ = v2905
	var v2909 int32
	_ = v2909
	var v2913 int32
	_ = v2913
	var v2918 int32
	_ = v2918
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2925 int32
	_ = v2925
	var v2938 int32
	_ = v2938
	var v2978 int32
	_ = v2978
	var v2981 int32
	_ = v2981
	var v2982 int32
	_ = v2982
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v2994 int32
	_ = v2994
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3008 int32
	_ = v3008
	var v3012 int32
	_ = v3012
	var v3060 int32
	_ = v3060
	var v3065 int32
	_ = v3065
	var v3067 int32
	_ = v3067
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3077 int32
	_ = v3077
	var v3079 int64
	_ = v3079
	var v3081 int32
	_ = v3081
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3090 int32
	_ = v3090
	var v3091 int32
	_ = v3091
	var v3094 int32
	_ = v3094
	var v3095 int32
	_ = v3095
	var v3101 int32
	_ = v3101
	var v3104 int32
	_ = v3104
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3112 int32
	_ = v3112
	var v3122 int32
	_ = v3122
	var v3131 int32
	_ = v3131
	var v3133 int32
	_ = v3133
	var v3135 int32
	_ = v3135
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3144 int32
	_ = v3144
	var v3146 int64
	_ = v3146
	var v3148 int32
	_ = v3148
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3153 int32
	_ = v3153
	var v3154 int32
	_ = v3154
	var v3157 int32
	_ = v3157
	var v3158 int32
	_ = v3158
	var v3161 int32
	_ = v3161
	var v3162 int32
	_ = v3162
	var v3168 int32
	_ = v3168
	var v3171 int32
	_ = v3171
	var v3174 int32
	_ = v3174
	var v3175 int32
	_ = v3175
	var v3179 int32
	_ = v3179
	var v3189 int32
	_ = v3189
	var v3204 int32
	_ = v3204
	var v3205 int32
	_ = v3205
	var v3209 int32
	_ = v3209
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3213 int32
	_ = v3213
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3223 int32
	_ = v3223
	var v3249 int32
	_ = v3249
	var v3257 int32
	_ = v3257
	var v3276 int32
	_ = v3276
	var v3280 int32
	_ = v3280
	var v3285 int32
	_ = v3285
	var v3288 int32
	_ = v3288
	var v3291 int32
	_ = v3291
	var v3292 int32
	_ = v3292
	var v3296 int32
	_ = v3296
	var v3328 int32
	_ = v3328
	var v3332 int32
	_ = v3332
	var v3350 int32
	_ = v3350
	var v3359 int32
	_ = v3359
	var v3362 int32
	_ = v3362
	var v3364 int32
	_ = v3364
	var v3392 int32
	_ = v3392
	var v3395 int32
	_ = v3395
	var v3419 int32
	_ = v3419
	var v3431 int32
	_ = v3431
	var v3434 int32
	_ = v3434
	var v3436 int32
	_ = v3436
	var v3466 int32
	_ = v3466
	var v3486 int32
	_ = v3486
	var v3516 int32
	_ = v3516
	var v3545 int32
	_ = v3545
	var v3551 int32
	_ = v3551
	var v3552 int32
	_ = v3552
	var v3568 int32
	_ = v3568
	var v3582 int32
	_ = v3582
	var v3607 int32
	_ = v3607
	var v3608 int32
	_ = v3608
	var v3611 int32
	_ = v3611
	var v3612 int32
	_ = v3612
	var v3613 int32
	_ = v3613
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3627 int32
	_ = v3627
	var v3629 int32
	_ = v3629
	var v3641 int32
	_ = v3641
	var v3690 int32
	_ = v3690
	var v3703 int32
	_ = v3703
	var v3729 int32
	_ = v3729
	var v3730 int32
	_ = v3730
	var v3733 int32
	_ = v3733
	var v3736 int32
	_ = v3736
	var v3787 int32
	_ = v3787
	var v3788 int32
	_ = v3788
	var v3792 int32
	_ = v3792
	var v3794 int32
	_ = v3794
	var v3797 int32
	_ = v3797
	var v3811 int32
	_ = v3811
	var v3874 int32
	_ = v3874
	var v3904 int32
	_ = v3904
	var v3907 int32
	_ = v3907
	var v3912 int32
	_ = v3912
	var v3917 int32
	_ = v3917
	var v3920 int32
	_ = v3920
	var v3922 int32
	_ = v3922
	var v3924 int32
	_ = v3924
	var v3928 int32
	_ = v3928
	var v3934 int32
	_ = v3934
	var v3949 int32
	_ = v3949
	var v3950 int32
	_ = v3950
	var v3956 int32
	_ = v3956
	var v3957 int32
	_ = v3957
	var v3959 int32
	_ = v3959
	var v3976 int32
	_ = v3976
	var v3990 int32
	_ = v3990
	var v4014 int32
	_ = v4014
	var v4015 int32
	_ = v4015
	var v4016 int32
	_ = v4016
	var v4017 int32
	_ = v4017
	var v4018 int32
	_ = v4018
	var v4019 int32
	_ = v4019
	var v4022 int64
	_ = v4022
	var v4036 int32
	_ = v4036
	var v4037 int32
	_ = v4037
	var v4047 int32
	_ = v4047
	var v4052 int32
	_ = v4052
	var v4098 int64
	_ = v4098
	var v4100 int32
	_ = v4100
	var v4101 int32
	_ = v4101
	var v4111 int32
	_ = v4111
	var v4128 int32
	_ = v4128
	var v4133 int32
	_ = v4133
	var v4147 int64
	_ = v4147
	var v4150 int32
	_ = v4150
	var v4152 int32
	_ = v4152
	var v4160 int32
	_ = v4160
	var v4163 int32
	_ = v4163
	var v4167 int32
	_ = v4167
	var v4171 int32
	_ = v4171
	var v4176 int32
	_ = v4176
	var v4180 int32
	_ = v4180
	var v4182 int32
	_ = v4182
	var v4190 int32
	_ = v4190
	var v4195 int32
	_ = v4195
	var v4199 int32
	_ = v4199
	var v4202 int32
	_ = v4202
	var v4210 int32
	_ = v4210
	var v4215 int32
	_ = v4215
	v10 = int32(0)
	v46 = int64(0)
	v48 = m.G0
	v50 = v48 - int32(4416)
	m.G0 = v50
	if l8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v53 = F_palloc(m, int32(524288))
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
	v305 = F_AllocateDir(m, l1)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L4
	} else {
		goto L89
	}
L7:
	;
	v271 = int32(314636)
	v274 = int32(*(*uint8)(unsafe.Add(mBase, _consts[410])))
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v275 == int32(0) {
		v294 = v274
		v295 = v275
		goto L76
	} else {
		goto L77
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
	v74 = int32(552432)
	v78 = m.G0
	v80 = v78 - int32(32)
	v81 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v80)+24)) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v80)+16)) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v80))) = v81
	v89 = int32(*(*uint8)(unsafe.Add(mBase, _consts[264])))
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
		goto L38
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
	v93 = int32(*(*uint8)(unsafe.Add(mBase, _consts[265])))
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
	v160 = int32(362921)
	v161 = v62 - l1
	if v161 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v205 != 0 {
		goto L53
	} else {
		goto L54
	}
L40:
	;
	v205 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v167 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v168 = l1
	v169 = v160
	v170 = v161
	v171 = v167
	goto L47
L44:
	;
	v193 = v160
	v197 = int32(0)
	goto L45
L45:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v205 = v197 - v198
	goto L39
L46:
	;
	v193 = v188
	v197 = v190
	goto L45
L47:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v171 != v173 {
		v188 = v169
		v190 = v171
		goto L46
	} else {
		goto L49
	}
L48:
	;
	v188 = v182
	v190 = int32(0)
	goto L46
L49:
	;
	if v173 == int32(0) {
		v188 = v169
		v190 = v171
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v178 = v170 - int32(1)
	if v178 == int32(0) {
		v188 = v169
		v190 = v171
		goto L46
	} else {
		goto L51
	}
L51:
	;
	v181 = int32(1)
	v182 = v169 + v181
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	if v183 != 0 {
		v168 = v168 + v181
		v169 = v182
		v170 = v178
		v171 = v183
		goto L47
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	v206 = int32(1663)
	if base.Ui32(v161) < base.Ui32(int32(15)) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v266 = F_strtox_2(m, v73, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L74
L56:
	;
	v301 = int32(0)
	v303 = v10
	v304 = v206
	goto L6
L57:
	;
	goto L58
L58:
	;
	v211 = int32(15)
	v212 = v62 - v211
	v213 = int32(562055)
	goto L61
L59:
	;
	if v250-v251 != 0 {
		v301 = int32(0)
		v303 = v10
		v304 = v206
		goto L6
	} else {
		goto L73
	}
L61:
	;
	goto L62
L62:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	if v220 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v221 = v212
	v222 = v213
	v223 = v211
	v224 = v220
	goto L67
L64:
	;
	v246 = v213
	v250 = int32(0)
	goto L65
L65:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	goto L59
L66:
	;
	v246 = v241
	v250 = v243
	goto L65
L67:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v224 != v226 {
		v241 = v222
		v243 = v224
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v241 = v235
	v243 = int32(0)
	goto L66
L69:
	;
	if v226 == int32(0) {
		v241 = v222
		v243 = v224
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v231 = v223 - int32(1)
	if v231 == int32(0) {
		v241 = v222
		v243 = v224
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v234 = int32(1)
	v235 = v222 + v234
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	if v236 != 0 {
		v221 = v221 + v234
		v222 = v235
		v223 = v231
		v224 = v236
		goto L67
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	goto L55
L74:
	;
	v301 = int32(1)
	v303 = base.I32_wrap_i64(v266)
	v304 = int32(1663)
	goto L6
L75:
	;
	if v296 != 0 {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	v296 = v295 - v294
	goto L75
L77:
	;
	if v274 != v275 {
		v294 = v274
		v295 = v275
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v279 = l1
	v280 = v271
	goto L79
L79:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	if v284 == int32(0) {
		v294 = v283
		v295 = v284
		goto L76
	} else {
		goto L81
	}
L80:
	;
	v294 = v283
	v295 = v284
	goto L76
L81:
	;
	v287 = int32(1)
	if v283 == v284 {
		v279 = v279 + v287
		v280 = v280 + v287
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v297 = int32(1663)
	goto L85
L84:
	;
	v297 = int32(1664)
	goto L85
L85:
	;
	v301 = base.B2i32(v296 == int32(0))
	v303 = v10
	v304 = v297
	goto L6
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4199 = m.ExcPending
	if v4199 != 0 {
		goto L4
	} else {
		goto L791
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4180 = m.ExcPending
	if v4180 != 0 {
		goto L4
	} else {
		goto L787
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4160 = m.ExcPending
	if v4160 != 0 {
		goto L4
	} else {
		goto L782
	}
L89:
	;
	v307 = F_ReadDir(m, v305, l1)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	if v307 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v309 = int32(1)
	v321 = l0
	v322 = l1
	v323 = l2
	v324 = l3
	v325 = l4
	v326 = l5
	v327 = l6
	v328 = l7
	v329 = l8
	v330 = v50
	v332 = v307
	v347 = v58
	v348 = l2 + v50 + int32(2369)
	v350 = v301
	v352 = v305
	v354 = v303
	v356 = v304
	v363 = l5 ^ v309
	v364 = l1 + l2 + v309
	v365 = v50 + int32(2368) | int32(2)
	v366 = v46
	goto L94
L92:
	;
	v4111 = v50
	v4128 = v58
	v4133 = v305
	v4147 = v46
	goto L93
L93:
	;
	if v4128 != 0 {
		goto L777
	} else {
		goto L778
	}
L94:
	;
	v368 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2268)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2264)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2260)) = v368
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+19)))
	if v374 != int32(46) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v4111 = v330
	v4128 = v347
	v4133 = v352
	v4147 = v4098
	goto L93
L96:
	;
	v4100 = F_ReadDir(m, v352, v322)
	mBase = m.M
	v4101 = m.ExcPending
	if v4101 != 0 {
		goto L4
	} else {
		goto L775
	}
L97:
	;
	v387 = v332 + int32(19)
	v388 = int32(235497)
	goto L104
L98:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+20)))
	if v377 == int32(0) {
		v4098 = v366
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+20)))
	if v380 != int32(46) {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+21)))
	if v383 == int32(0) {
		v4098 = v366
		goto L96
	} else {
		goto L101
	}
L101:
	;
	goto L97
L102:
	;
	if v425-v426 == int32(0) {
		v4098 = v366
		goto L96
	} else {
		goto L116
	}
L104:
	;
	goto L105
L105:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v395 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v396 = v387
	v397 = v388
	v398 = int32(9)
	v399 = v395
	goto L110
L107:
	;
	v421 = v388
	v425 = int32(0)
	goto L108
L108:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421))))
	goto L102
L109:
	;
	v421 = v416
	v425 = v418
	goto L108
L110:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397))))
	if v399 != v401 {
		v416 = v397
		v418 = v399
		goto L109
	} else {
		goto L112
	}
L111:
	;
	v416 = v410
	v418 = int32(0)
	goto L109
L112:
	;
	if v401 == int32(0) {
		v416 = v397
		v418 = v399
		goto L109
	} else {
		goto L113
	}
L113:
	;
	v406 = v398 - int32(1)
	if v406 == int32(0) {
		v416 = v397
		v418 = v399
		goto L109
	} else {
		goto L114
	}
L114:
	;
	v409 = int32(1)
	v410 = v397 + v409
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+1)))
	if v411 != 0 {
		v396 = v396 + v409
		v397 = v410
		v398 = v406
		v399 = v411
		goto L110
	} else {
		goto L115
	}
L115:
	;
	goto L111
L116:
	;
	v436 = int32(364900)
	v439 = int32(*(*uint8)(unsafe.Add(mBase, _consts[411])))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v440 == int32(0) {
		v459 = v439
		v460 = v440
		goto L118
	} else {
		goto L119
	}
L117:
	;
	if v460-v459 == int32(0) {
		v4098 = v366
		goto L96
	} else {
		goto L125
	}
L118:
	;
	goto L117
L119:
	;
	if v439 != v440 {
		v459 = v439
		v460 = v440
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v444 = v387
	v445 = v436
	goto L121
L121:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+1)))
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+1)))
	if v449 == int32(0) {
		v459 = v448
		v460 = v449
		goto L118
	} else {
		goto L123
	}
L122:
	;
	v459 = v448
	v460 = v449
	goto L118
L123:
	;
	v452 = int32(1)
	if v448 == v449 {
		v444 = v444 + v452
		v445 = v445 + v452
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v465 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v470 == int32(1) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L128
L130:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, _consts[412])))
	if v480 != v482 {
		goto L88
	} else {
		goto L134
	}
L131:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+316))
	v478 = base.B2i32(v476 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v478)
	v480 = v478
	goto L133
L132:
	;
	v480 = int32(0)
	goto L133
L133:
	;
	goto L130
L134:
	;
	v484 = int32(235642)
	goto L140
L135:
	;
	if v354 == int32(0) {
		goto L291
	} else {
		goto L292
	}
L136:
	;
	v1032 = int32(1)
	goto L135
L137:
	;
	v1018 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L4
	} else {
		goto L287
	}
L138:
	;
	if v521-v522 == int32(0) {
		goto L137
	} else {
		goto L152
	}
L140:
	;
	goto L141
L141:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v491 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v492 = v387
	v493 = v484
	v494 = int32(25)
	v495 = v491
	goto L146
L143:
	;
	v517 = v484
	v521 = int32(0)
	goto L144
L144:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517))))
	goto L138
L145:
	;
	v517 = v512
	v521 = v514
	goto L144
L146:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
	if v495 != v497 {
		v512 = v493
		v514 = v495
		goto L145
	} else {
		goto L148
	}
L147:
	;
	v512 = v506
	v514 = int32(0)
	goto L145
L148:
	;
	if v497 == int32(0) {
		v512 = v493
		v514 = v495
		goto L145
	} else {
		goto L149
	}
L149:
	;
	v502 = v494 - int32(1)
	if v502 == int32(0) {
		v512 = v493
		v514 = v495
		goto L145
	} else {
		goto L150
	}
L150:
	;
	v505 = int32(1)
	v506 = v493 + v505
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+1)))
	if v507 != 0 {
		v492 = v492 + v505
		v493 = v506
		v494 = v502
		v495 = v507
		goto L146
	} else {
		goto L151
	}
L151:
	;
	goto L147
L152:
	;
	v532 = int32(235591)
	goto L155
L153:
	;
	if v569-v570 == int32(0) {
		goto L137
	} else {
		goto L167
	}
L155:
	;
	goto L156
L156:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v539 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v540 = v387
	v541 = v532
	v542 = int32(21)
	v543 = v539
	goto L161
L158:
	;
	v565 = v532
	v569 = int32(0)
	goto L159
L159:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565))))
	goto L153
L160:
	;
	v565 = v560
	v569 = v562
	goto L159
L161:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	if v543 != v545 {
		v560 = v541
		v562 = v543
		goto L160
	} else {
		goto L163
	}
L162:
	;
	v560 = v554
	v562 = int32(0)
	goto L160
L163:
	;
	if v545 == int32(0) {
		v560 = v541
		v562 = v543
		goto L160
	} else {
		goto L164
	}
L164:
	;
	v550 = v542 - int32(1)
	if v550 == int32(0) {
		v560 = v541
		v562 = v543
		goto L160
	} else {
		goto L165
	}
L165:
	;
	v553 = int32(1)
	v554 = v541 + v553
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+1)))
	if v555 != 0 {
		v540 = v540 + v553
		v541 = v554
		v542 = v550
		v543 = v555
		goto L161
	} else {
		goto L166
	}
L166:
	;
	goto L162
L167:
	;
	v580 = int32(100399)
	goto L170
L168:
	;
	if v617-v618 == int32(0) {
		goto L137
	} else {
		goto L182
	}
L170:
	;
	goto L171
L171:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v587 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v588 = v387
	v589 = v580
	v590 = int32(16)
	v591 = v587
	goto L176
L173:
	;
	v613 = v580
	v617 = int32(0)
	goto L174
L174:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	goto L168
L175:
	;
	v613 = v608
	v617 = v610
	goto L174
L176:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	if v591 != v593 {
		v608 = v589
		v610 = v591
		goto L175
	} else {
		goto L178
	}
L177:
	;
	v608 = v602
	v610 = int32(0)
	goto L175
L178:
	;
	if v593 == int32(0) {
		v608 = v589
		v610 = v591
		goto L175
	} else {
		goto L179
	}
L179:
	;
	v598 = v590 - int32(1)
	if v598 == int32(0) {
		v608 = v589
		v610 = v591
		goto L175
	} else {
		goto L180
	}
L180:
	;
	v601 = int32(1)
	v602 = v589 + v601
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588)+1)))
	if v603 != 0 {
		v588 = v588 + v601
		v589 = v602
		v590 = v598
		v591 = v603
		goto L176
	} else {
		goto L181
	}
L181:
	;
	goto L177
L182:
	;
	v628 = int32(308363)
	goto L185
L183:
	;
	if v665-v666 == int32(0) {
		goto L137
	} else {
		goto L197
	}
L185:
	;
	goto L186
L186:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v635 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v636 = v387
	v637 = v628
	v638 = int32(13)
	v639 = v635
	goto L191
L188:
	;
	v661 = v628
	v665 = int32(0)
	goto L189
L189:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661))))
	goto L183
L190:
	;
	v661 = v656
	v665 = v658
	goto L189
L191:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637))))
	if v639 != v641 {
		v656 = v637
		v658 = v639
		goto L190
	} else {
		goto L193
	}
L192:
	;
	v656 = v650
	v658 = int32(0)
	goto L190
L193:
	;
	if v641 == int32(0) {
		v656 = v637
		v658 = v639
		goto L190
	} else {
		goto L194
	}
L194:
	;
	v646 = v638 - int32(1)
	if v646 == int32(0) {
		v656 = v637
		v658 = v639
		goto L190
	} else {
		goto L195
	}
L195:
	;
	v649 = int32(1)
	v650 = v637 + v649
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636)+1)))
	if v651 != 0 {
		v636 = v636 + v649
		v637 = v650
		v638 = v646
		v639 = v651
		goto L191
	} else {
		goto L196
	}
L196:
	;
	goto L192
L197:
	;
	v676 = int32(238548)
	goto L200
L198:
	;
	if v713-v714 == int32(0) {
		goto L137
	} else {
		goto L212
	}
L200:
	;
	goto L201
L201:
	;
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v683 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v684 = v387
	v685 = v676
	v686 = int32(15)
	v687 = v683
	goto L206
L203:
	;
	v709 = v676
	v713 = int32(0)
	goto L204
L204:
	;
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v709))))
	goto L198
L205:
	;
	v709 = v704
	v713 = v706
	goto L204
L206:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685))))
	if v687 != v689 {
		v704 = v685
		v706 = v687
		goto L205
	} else {
		goto L208
	}
L207:
	;
	v704 = v698
	v706 = int32(0)
	goto L205
L208:
	;
	if v689 == int32(0) {
		v704 = v685
		v706 = v687
		goto L205
	} else {
		goto L209
	}
L209:
	;
	v694 = v686 - int32(1)
	if v694 == int32(0) {
		v704 = v685
		v706 = v687
		goto L205
	} else {
		goto L210
	}
L210:
	;
	v697 = int32(1)
	v698 = v685 + v697
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v684)+1)))
	if v699 != 0 {
		v684 = v684 + v697
		v685 = v698
		v686 = v694
		v687 = v699
		goto L206
	} else {
		goto L211
	}
L211:
	;
	goto L207
L212:
	;
	v724 = int32(77521)
	goto L215
L213:
	;
	if v761-v762 == int32(0) {
		goto L137
	} else {
		goto L227
	}
L215:
	;
	goto L216
L216:
	;
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v731 != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v732 = v387
	v733 = v724
	v734 = int32(16)
	v735 = v731
	goto L221
L218:
	;
	v757 = v724
	v761 = int32(0)
	goto L219
L219:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757))))
	goto L213
L220:
	;
	v757 = v752
	v761 = v754
	goto L219
L221:
	;
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733))))
	if v735 != v737 {
		v752 = v733
		v754 = v735
		goto L220
	} else {
		goto L223
	}
L222:
	;
	v752 = v746
	v754 = int32(0)
	goto L220
L223:
	;
	if v737 == int32(0) {
		v752 = v733
		v754 = v735
		goto L220
	} else {
		goto L224
	}
L224:
	;
	v742 = v734 - int32(1)
	if v742 == int32(0) {
		v752 = v733
		v754 = v735
		goto L220
	} else {
		goto L225
	}
L225:
	;
	v745 = int32(1)
	v746 = v733 + v745
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732)+1)))
	if v747 != 0 {
		v732 = v732 + v745
		v733 = v746
		v734 = v742
		v735 = v747
		goto L221
	} else {
		goto L226
	}
L226:
	;
	goto L222
L227:
	;
	v772 = int32(433770)
	goto L230
L228:
	;
	if v809-v810 == int32(0) {
		goto L137
	} else {
		goto L242
	}
L230:
	;
	goto L231
L231:
	;
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v779 != 0 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v780 = v387
	v781 = v772
	v782 = int32(15)
	v783 = v779
	goto L236
L233:
	;
	v805 = v772
	v809 = int32(0)
	goto L234
L234:
	;
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v805))))
	goto L228
L235:
	;
	v805 = v800
	v809 = v802
	goto L234
L236:
	;
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781))))
	if v783 != v785 {
		v800 = v781
		v802 = v783
		goto L235
	} else {
		goto L238
	}
L237:
	;
	v800 = v794
	v802 = int32(0)
	goto L235
L238:
	;
	if v785 == int32(0) {
		v800 = v781
		v802 = v783
		goto L235
	} else {
		goto L239
	}
L239:
	;
	v790 = v782 - int32(1)
	if v790 == int32(0) {
		v800 = v781
		v802 = v783
		goto L235
	} else {
		goto L240
	}
L240:
	;
	v793 = int32(1)
	v794 = v781 + v793
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780)+1)))
	if v795 != 0 {
		v780 = v780 + v793
		v781 = v794
		v782 = v790
		v783 = v795
		goto L236
	} else {
		goto L241
	}
L241:
	;
	goto L237
L242:
	;
	v820 = int32(118257)
	goto L245
L243:
	;
	if v857-v858 == int32(0) {
		goto L137
	} else {
		goto L257
	}
L245:
	;
	goto L246
L246:
	;
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v827 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v828 = v387
	v829 = v820
	v830 = int32(16)
	v831 = v827
	goto L251
L248:
	;
	v853 = v820
	v857 = int32(0)
	goto L249
L249:
	;
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v853))))
	goto L243
L250:
	;
	v853 = v848
	v857 = v850
	goto L249
L251:
	;
	v833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829))))
	if v831 != v833 {
		v848 = v829
		v850 = v831
		goto L250
	} else {
		goto L253
	}
L252:
	;
	v848 = v842
	v850 = int32(0)
	goto L250
L253:
	;
	if v833 == int32(0) {
		v848 = v829
		v850 = v831
		goto L250
	} else {
		goto L254
	}
L254:
	;
	v838 = v830 - int32(1)
	if v838 == int32(0) {
		v848 = v829
		v850 = v831
		goto L250
	} else {
		goto L255
	}
L255:
	;
	v841 = int32(1)
	v842 = v829 + v841
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v828)+1)))
	if v843 != 0 {
		v828 = v828 + v841
		v829 = v842
		v830 = v838
		v831 = v843
		goto L251
	} else {
		goto L256
	}
L256:
	;
	goto L252
L257:
	;
	v868 = int32(0)
	if v350 == v868 {
		v1032 = v868
		goto L135
	} else {
		goto L258
	}
L258:
	;
	v872 = v330 + int32(2268)
	v874 = v330 + int32(2264)
	v876 = v330 + int32(2260)
	v877 = int32(0)
	v882 = m.G0
	v884 = v882 - int32(16)
	m.G0 = v884
	*(*int32)(unsafe.Add(mBase, uint32(v872))) = v877
	*(*int32)(unsafe.Add(mBase, uint32(v874))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v876))) = v877
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if base.Ui32((v892-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v972 = v877
		goto L260
	} else {
		goto L261
	}
L259:
	;
	if v972 == int32(0) {
		v1032 = v972
		goto L135
	} else {
		goto L278
	}
L260:
	;
	m.G0 = v884 + int32(16)
	goto L259
L261:
	;
	v899 = int32(4680308)
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
	v905 = F_strtoul(m, v387, v884+int32(8), int32(10))
	mBase = m.M
	v907 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v907 != 0 {
		v972 = v877
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v884)+8))
	if v387 == v908 {
		v972 = v877
		goto L260
	} else {
		goto L263
	}
L263:
	;
	if v905 == int32(0) {
		v972 = v877
		goto L260
	} else {
		goto L264
	}
L264:
	;
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908))))
	if v912 != int32(95) {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	if v928&int32(255) == int32(46) {
		goto L270
	} else {
		goto L271
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v884)+12)) = int32(0)
	v928 = v912
	v929 = v908
	goto L265
L267:
	;
	goto L268
L268:
	;
	v921 = F_forkname_chars(m, v908+int32(1), v884+int32(12))
	mBase = m.M
	if v921 <= int32(0) {
		v972 = v877
		goto L260
	} else {
		goto L269
	}
L269:
	;
	v926 = v921 + v908 + int32(1)
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926))))
	v928 = v927
	v929 = v926
	goto L265
L270:
	;
	v935 = v929 + int32(1)
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935))))
	if base.Ui32((v936-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v972 = v877
		goto L260
	} else {
		goto L273
	}
L271:
	;
	v959 = v877
	v960 = v928
	goto L272
L272:
	;
	if v960&int32(255) != 0 {
		v972 = v877
		goto L260
	} else {
		goto L277
	}
L273:
	;
	v943 = int32(4680308)
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
	v949 = F_strtoul(m, v935, v884+int32(8), int32(10))
	mBase = m.M
	v951 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v951 != 0 {
		v972 = v877
		goto L260
	} else {
		goto L274
	}
L274:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v884)+8))
	if v935 == v952 {
		v972 = v877
		goto L260
	} else {
		goto L275
	}
L275:
	;
	if v949 == int32(0) {
		v972 = v877
		goto L260
	} else {
		goto L276
	}
L276:
	;
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v952))))
	v959 = v949
	v960 = v956
	goto L272
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v872))) = v905
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v884)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v874))) = v964
	*(*int32)(unsafe.Add(mBase, uint32(v876))) = v959
	v972 = int32(1)
	goto L260
L278:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2264))
	if v978 == int32(3) {
		v1032 = v972
		goto L135
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+176)) = v322
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2268))
	*(*int32)(unsafe.Add(mBase, uint32(v330)+180)) = v982
	v990 = F_pg_snprintf(m, v330+int32(192), int32(1024), int32(99858), v330+int32(176))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L4
	} else {
		goto L280
	}
L280:
	;
	v998 = F___fstatat(m, int32(-100), v330+int32(192), v330+int32(2272), int32(256))
	mBase = m.M
	goto L281
L281:
	;
	if v998 != 0 {
		goto L136
	} else {
		goto L282
	}
L282:
	;
	v1001 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L4
	} else {
		goto L283
	}
L283:
	;
	if v1001 == int32(0) {
		v4098 = v366
		goto L96
	} else {
		goto L284
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+16)) = v387
	F_errmsg_internal(m, int32(233418), v330+int32(16))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L4
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(496115), int32(1333), int32(213662))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L4
	} else {
		goto L286
	}
L286:
	;
	v4098 = v366
	goto L96
L287:
	;
	if v1018 == int32(0) {
		v4098 = v366
		goto L96
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330))) = v387
	F_errmsg_internal(m, int32(233436), v330)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L4
	} else {
		goto L289
	}
L289:
	;
	F_errfinish(m, int32(496115), int32(1298), int32(213662))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L4
	} else {
		goto L290
	}
L290:
	;
	v4098 = v366
	goto L96
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+148)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v330)+144)) = v322
	v1443 = F_pg_snprintf(m, v330+int32(2368), int32(2048), int32(177111), v330+int32(144))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L4
	} else {
		goto L334
	}
L292:
	;
	v1035 = int32(0)
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v1036 != int32(116) {
		v1341 = v1035
		goto L293
	} else {
		goto L294
	}
L293:
	;
	if v1341 == int32(0) {
		goto L291
	} else {
		goto L329
	}
L294:
	;
	v1052 = int32(1)
	goto L295
L295:
	;
	v1088 = v1052 + int32(1)
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052+v387))))
	if base.Ui32((v1090-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v1052 = v1088
		goto L295
	} else {
		goto L297
	}
L296:
	;
	if v1052 == int32(1) {
		v1341 = v1035
		goto L293
	} else {
		goto L298
	}
L297:
	;
	goto L296
L298:
	;
	if v1090&int32(255) != int32(95) {
		v1341 = v1035
		goto L293
	} else {
		goto L299
	}
L299:
	;
	v1117 = v1088
	goto L300
L300:
	;
	v1151 = v1117 + int32(1)
	v1152 = v1117 + v387
	v1153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1152))))
	if base.Ui32((v1153-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v1117 = v1151
		goto L300
	} else {
		goto L302
	}
L301:
	;
	if v1117 == v1088 {
		v1341 = v1035
		goto L293
	} else {
		goto L303
	}
L302:
	;
	goto L301
L303:
	;
	if v1153 == int32(95) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1164 = v1152 + int32(1)
	v1168 = int32(3)
	v1171 = F_strncmp(m, int32(287938), v1164, v1168)
	mBase = m.M
	if v1171 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L305:
	;
	v1207 = v1117
	v1209 = v1153
	goto L306
L306:
	;
	if v1209 == int32(46) {
		goto L322
	} else {
		goto L323
	}
L307:
	;
	if v1200 <= int32(0) {
		v1341 = v1035
		goto L293
	} else {
		goto L321
	}
L308:
	;
	goto L307
L310:
	;
	v1200 = v1193
	goto L308
L311:
	;
	v1193 = v1168
	goto L310
L312:
	;
	goto L313
L313:
	;
	v1175 = int32(2)
	v1179 = F_strncmp(m, int32(285952), v1164, v1175)
	mBase = m.M
	if v1179 == int32(0) {
		v1193 = v1175
		goto L310
	} else {
		goto L314
	}
L314:
	;
	v1182 = int32(4)
	v1185 = F_strncmp(m, int32(100467), v1164, v1182)
	mBase = m.M
	if v1185 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	goto L318
L316:
	;
	goto L317
L317:
	;
	v1200 = int32(0)
	goto L308
L318:
	;
	v1200 = v1182
	goto L308
L321:
	;
	v1204 = v1200 + v1151
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387+v1204))))
	v1207 = v1204
	v1209 = v1206
	goto L306
L322:
	;
	v1226 = int32(1)
	goto L325
L323:
	;
	v1290 = v1209
	goto L324
L324:
	;
	v1341 = base.B2i32(v1290 == int32(0))
	goto L293
L325:
	;
	v1264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226+(v1207+v387)))))
	if base.Ui32((v1264-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v1226 = v1226 + int32(1)
		goto L325
	} else {
		goto L327
	}
L326:
	;
	if v1226 < int32(2) {
		v1341 = v1035
		goto L293
	} else {
		goto L328
	}
L327:
	;
	goto L326
L328:
	;
	v1290 = v1264
	goto L324
L329:
	;
	v1373 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L4
	} else {
		goto L330
	}
L330:
	;
	if v1373 == int32(0) {
		v4098 = v366
		goto L96
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+160)) = v387
	F_errmsg_internal(m, int32(233368), v330+int32(160))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L4
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(496115), int32(1344), int32(213662))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L4
	} else {
		goto L333
	}
L333:
	;
	v4098 = v366
	goto L96
L334:
	;
	v1446 = v330 + int32(2368)
	v1447 = int32(301640)
	v1448 = int32(20)
	goto L338
L335:
	;
	if v1510 == int32(0) {
		v4098 = v366
		goto L96
	} else {
		goto L353
	}
L336:
	;
	v1510 = int32(0)
	goto L335
L337:
	;
	v1484 = v1479
	v1485 = v1480
	v1486 = v1481
	goto L347
L338:
	;
	if (v1446|v1447)&int32(3) != 0 {
		v1479 = v1446
		v1480 = v1447
		v1481 = v1448
		goto L337
	} else {
		goto L341
	}
L340:
	;
	if v1469 == int32(0) {
		goto L336
	} else {
		goto L346
	}
L341:
	;
	v1456 = v1446
	v1457 = v1447
	v1458 = v1448
	goto L342
L342:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1456)))
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1457)))
	if v1461 != v1462 {
		v1479 = v1456
		v1480 = v1457
		v1481 = v1458
		goto L337
	} else {
		goto L344
	}
L343:
	;
	goto L340
L344:
	;
	v1464 = int32(4)
	v1465 = v1457 + v1464
	v1467 = v1456 + v1464
	v1469 = v1458 - v1464
	if base.Ui32(int32(3)) < base.Ui32(v1469) {
		v1456 = v1467
		v1457 = v1465
		v1458 = v1469
		goto L342
	} else {
		goto L345
	}
L345:
	;
	goto L343
L346:
	;
	v1479 = v1467
	v1480 = v1465
	v1481 = v1469
	goto L337
L347:
	;
	v1489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1484))))
	v1490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1485))))
	if v1489 == v1490 {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	v1510 = v1489 - v1490
	goto L335
L349:
	;
	v1492 = int32(1)
	v1497 = v1486 - v1492
	if v1497 != 0 {
		v1484 = v1484 + v1492
		v1485 = v1485 + v1492
		v1486 = v1497
		goto L347
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	goto L348
L352:
	;
	goto L336
L353:
	;
	v1519 = F___fstatat(m, int32(-100), v330+int32(2368), v330+int32(2272), int32(256))
	mBase = m.M
	goto L356
L354:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2276))
	v1870 = v1868 & int32(61440)
	v1871 = int32(489143)
	v1874 = int32(*(*uint8)(unsafe.Add(mBase, _consts[413])))
	v1875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	if v1875 == int32(0) {
		v1894 = v1874
		v1895 = v1875
		goto L465
	} else {
		goto L466
	}
L355:
	;
	v1838 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L4
	} else {
		goto L453
	}
L356:
	;
	if v1519 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1522 = int32(235485)
	v1525 = int32(*(*uint8)(unsafe.Add(mBase, _consts[414])))
	v1526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v1526 == int32(0) {
		v1545 = v1525
		v1546 = v1526
		goto L361
	} else {
		goto L362
	}
L358:
	;
	goto L359
L359:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v1814 == int32(44) {
		v4098 = v366
		goto L96
	} else {
		goto L448
	}
L360:
	;
	if v1546-v1545 == int32(0) {
		goto L355
	} else {
		goto L368
	}
L361:
	;
	goto L360
L362:
	;
	if v1525 != v1526 {
		v1545 = v1525
		v1546 = v1526
		goto L361
	} else {
		goto L363
	}
L363:
	;
	v1530 = v387
	v1531 = v1522
	goto L364
L364:
	;
	v1534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1531)+1)))
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1530)+1)))
	if v1535 == int32(0) {
		v1545 = v1534
		v1546 = v1535
		goto L361
	} else {
		goto L366
	}
L365:
	;
	v1545 = v1534
	v1546 = v1535
	goto L361
L366:
	;
	v1538 = int32(1)
	if v1534 == v1535 {
		v1530 = v1530 + v1538
		v1531 = v1531 + v1538
		goto L364
	} else {
		goto L367
	}
L367:
	;
	goto L365
L368:
	;
	v1550 = int32(84710)
	v1553 = int32(*(*uint8)(unsafe.Add(mBase, _consts[415])))
	v1554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v1554 == int32(0) {
		v1573 = v1553
		v1574 = v1554
		goto L370
	} else {
		goto L371
	}
L369:
	;
	if v1574-v1573 == int32(0) {
		goto L355
	} else {
		goto L377
	}
L370:
	;
	goto L369
L371:
	;
	if v1553 != v1554 {
		v1573 = v1553
		v1574 = v1554
		goto L370
	} else {
		goto L372
	}
L372:
	;
	v1558 = v387
	v1559 = v1550
	goto L373
L373:
	;
	v1562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1559)+1)))
	v1563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1558)+1)))
	if v1563 == int32(0) {
		v1573 = v1562
		v1574 = v1563
		goto L370
	} else {
		goto L375
	}
L374:
	;
	v1573 = v1562
	v1574 = v1563
	goto L370
L375:
	;
	v1566 = int32(1)
	if v1562 == v1563 {
		v1558 = v1558 + v1566
		v1559 = v1559 + v1566
		goto L373
	} else {
		goto L376
	}
L376:
	;
	goto L374
L377:
	;
	v1578 = int32(291661)
	v1581 = int32(*(*uint8)(unsafe.Add(mBase, _consts[416])))
	v1582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v1582 == int32(0) {
		v1601 = v1581
		v1602 = v1582
		goto L379
	} else {
		goto L380
	}
L378:
	;
	if v1602-v1601 == int32(0) {
		goto L355
	} else {
		goto L386
	}
L379:
	;
	goto L378
L380:
	;
	if v1581 != v1582 {
		v1601 = v1581
		v1602 = v1582
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v1586 = v387
	v1587 = v1578
	goto L382
L382:
	;
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1587)+1)))
	v1591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1586)+1)))
	if v1591 == int32(0) {
		v1601 = v1590
		v1602 = v1591
		goto L379
	} else {
		goto L384
	}
L383:
	;
	v1601 = v1590
	v1602 = v1591
	goto L379
L384:
	;
	v1594 = int32(1)
	if v1590 == v1591 {
		v1586 = v1586 + v1594
		v1587 = v1587 + v1594
		goto L382
	} else {
		goto L385
	}
L385:
	;
	goto L383
L386:
	;
	v1606 = int32(20719)
	v1609 = int32(*(*uint8)(unsafe.Add(mBase, _consts[417])))
	v1610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v1610 == int32(0) {
		v1629 = v1609
		v1630 = v1610
		goto L388
	} else {
		goto L389
	}
L387:
	;
	if v1630-v1629 == int32(0) {
		goto L355
	} else {
		goto L395
	}
L388:
	;
	goto L387
L389:
	;
	if v1609 != v1610 {
		v1629 = v1609
		v1630 = v1610
		goto L388
	} else {
		goto L390
	}
L390:
	;
	v1614 = v387
	v1615 = v1606
	goto L391
L391:
	;
	v1618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1615)+1)))
	v1619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1614)+1)))
	if v1619 == int32(0) {
		v1629 = v1618
		v1630 = v1619
		goto L388
	} else {
		goto L393
	}
L392:
	;
	v1629 = v1618
	v1630 = v1619
	goto L388
L393:
	;
	v1622 = int32(1)
	if v1618 == v1619 {
		v1614 = v1614 + v1622
		v1615 = v1615 + v1622
		goto L391
	} else {
		goto L394
	}
L394:
	;
	goto L392
L395:
	;
	v1634 = int32(314258)
	v1637 = int32(*(*uint8)(unsafe.Add(mBase, _consts[418])))
	v1638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v1638 == int32(0) {
		v1657 = v1637
		v1658 = v1638
		goto L397
	} else {
		goto L398
	}
L396:
	;
	if v1658-v1657 == int32(0) {
		goto L355
	} else {
		goto L404
	}
L397:
	;
	goto L396
L398:
	;
	if v1637 != v1638 {
		v1657 = v1637
		v1658 = v1638
		goto L397
	} else {
		goto L399
	}
L399:
	;
	v1642 = v387
	v1643 = v1634
	goto L400
L400:
	;
	v1646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1643)+1)))
	v1647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1642)+1)))
	if v1647 == int32(0) {
		v1657 = v1646
		v1658 = v1647
		goto L397
	} else {
		goto L402
	}
L401:
	;
	v1657 = v1646
	v1658 = v1647
	goto L397
L402:
	;
	v1650 = int32(1)
	if v1646 == v1647 {
		v1642 = v1642 + v1650
		v1643 = v1643 + v1650
		goto L400
	} else {
		goto L403
	}
L403:
	;
	goto L401
L404:
	;
	v1662 = int32(119116)
	v1665 = int32(*(*uint8)(unsafe.Add(mBase, _consts[419])))
	v1666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v1666 == int32(0) {
		v1685 = v1665
		v1686 = v1666
		goto L406
	} else {
		goto L407
	}
L405:
	;
	if v1686-v1685 == int32(0) {
		goto L355
	} else {
		goto L413
	}
L406:
	;
	goto L405
L407:
	;
	if v1665 != v1666 {
		v1685 = v1665
		v1686 = v1666
		goto L406
	} else {
		goto L408
	}
L408:
	;
	v1670 = v387
	v1671 = v1662
	goto L409
L409:
	;
	v1674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1671)+1)))
	v1675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1670)+1)))
	if v1675 == int32(0) {
		v1685 = v1674
		v1686 = v1675
		goto L406
	} else {
		goto L411
	}
L410:
	;
	v1685 = v1674
	v1686 = v1675
	goto L406
L411:
	;
	v1678 = int32(1)
	if v1674 == v1675 {
		v1670 = v1670 + v1678
		v1671 = v1671 + v1678
		goto L409
	} else {
		goto L412
	}
L412:
	;
	goto L410
L413:
	;
	v1690 = int32(149938)
	v1693 = int32(*(*uint8)(unsafe.Add(mBase, _consts[420])))
	v1694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v1694 == int32(0) {
		v1713 = v1693
		v1714 = v1694
		goto L415
	} else {
		goto L416
	}
L414:
	;
	if v1714-v1713 == int32(0) {
		goto L355
	} else {
		goto L422
	}
L415:
	;
	goto L414
L416:
	;
	if v1693 != v1694 {
		v1713 = v1693
		v1714 = v1694
		goto L415
	} else {
		goto L417
	}
L417:
	;
	v1698 = v387
	v1699 = v1690
	goto L418
L418:
	;
	v1702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1699)+1)))
	v1703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1698)+1)))
	if v1703 == int32(0) {
		v1713 = v1702
		v1714 = v1703
		goto L415
	} else {
		goto L420
	}
L419:
	;
	v1713 = v1702
	v1714 = v1703
	goto L415
L420:
	;
	v1706 = int32(1)
	if v1702 == v1703 {
		v1698 = v1698 + v1706
		v1699 = v1699 + v1706
		goto L418
	} else {
		goto L421
	}
L421:
	;
	goto L419
L422:
	;
	v1719 = v330 + int32(2368)
	v1720 = int32(308816)
	v1721 = int32(9)
	goto L426
L423:
	;
	if v1783 != 0 {
		goto L354
	} else {
		goto L441
	}
L424:
	;
	v1783 = int32(0)
	goto L423
L425:
	;
	v1757 = v1752
	v1758 = v1753
	v1759 = v1754
	goto L435
L426:
	;
	if (v1719|v1720)&int32(3) != 0 {
		v1752 = v1719
		v1753 = v1720
		v1754 = v1721
		goto L425
	} else {
		goto L429
	}
L428:
	;
	if v1742 == int32(0) {
		goto L424
	} else {
		goto L434
	}
L429:
	;
	v1729 = v1719
	v1730 = v1720
	v1731 = v1721
	goto L430
L430:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1729)))
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1730)))
	if v1734 != v1735 {
		v1752 = v1729
		v1753 = v1730
		v1754 = v1731
		goto L425
	} else {
		goto L432
	}
L431:
	;
	goto L428
L432:
	;
	v1737 = int32(4)
	v1738 = v1730 + v1737
	v1740 = v1729 + v1737
	v1742 = v1731 - v1737
	if base.Ui32(int32(3)) < base.Ui32(v1742) {
		v1729 = v1740
		v1730 = v1738
		v1731 = v1742
		goto L430
	} else {
		goto L433
	}
L433:
	;
	goto L431
L434:
	;
	v1752 = v1740
	v1753 = v1738
	v1754 = v1742
	goto L425
L435:
	;
	v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1757))))
	v1763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1758))))
	if v1762 == v1763 {
		goto L437
	} else {
		goto L438
	}
L436:
	;
	v1783 = v1762 - v1763
	goto L423
L437:
	;
	v1765 = int32(1)
	v1770 = v1759 - v1765
	if v1770 != 0 {
		v1757 = v1757 + v1765
		v1758 = v1758 + v1765
		v1759 = v1770
		goto L435
	} else {
		goto L440
	}
L438:
	;
	goto L439
L439:
	;
	goto L436
L440:
	;
	goto L424
L441:
	;
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2276))
	if v1784&int32(61440) == int32(40960) {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2276)) = v1790 | int32(16384)
	goto L444
L443:
	;
	goto L444
L444:
	;
	F__tarWriteHeader(m, v321, v348, int32(0), v330+int32(2272), v324)
	mBase = m.M
	v1798 = m.ExcPending
	if v1798 != 0 {
		goto L4
	} else {
		goto L445
	}
L445:
	;
	F__tarWriteHeader(m, v321, int32(114715), int32(0), v330+int32(2272), v324)
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L4
	} else {
		goto L446
	}
L446:
	;
	F__tarWriteHeader(m, v321, int32(168365), int32(0), v330+int32(2272), v324)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L4
	} else {
		goto L447
	}
L447:
	;
	v4098 = v366 + int64(1536)
	goto L96
L448:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L4
	} else {
		goto L449
	}
L449:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L4
	} else {
		goto L450
	}
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+128)) = v330 + int32(2368)
	F_errmsg(m, int32(296611), v330+int32(128))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L4
	} else {
		goto L451
	}
L451:
	;
	F_errfinish(m, int32(496115), int32(1361), int32(213662))
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L4
	} else {
		goto L452
	}
L452:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L453:
	;
	if v1838 != 0 {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+32)) = v387
	F_errmsg_internal(m, int32(233320), v330+int32(32))
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L4
	} else {
		goto L457
	}
L455:
	;
	goto L456
L456:
	;
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2276))
	if v1851&int32(61440) == int32(40960) {
		goto L459
	} else {
		goto L460
	}
L457:
	;
	F_errfinish(m, int32(496115), int32(1373), int32(213662))
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L4
	} else {
		goto L458
	}
L458:
	;
	goto L456
L459:
	;
	v1857 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2276)) = v1857 | int32(16384)
	goto L461
L460:
	;
	goto L461
L461:
	;
	F__tarWriteHeader(m, v321, v348, int32(0), v330+int32(2272), v324)
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L4
	} else {
		goto L462
	}
L462:
	;
	v4098 = v366 + int64(512)
	goto L96
L463:
	;
	if v1870 != int32(32768) {
		goto L478
	} else {
		goto L479
	}
L464:
	;
	if v1895-v1894 != 0 {
		goto L463
	} else {
		goto L472
	}
L465:
	;
	goto L464
L466:
	;
	if v1874 != v1875 {
		v1894 = v1874
		v1895 = v1875
		goto L465
	} else {
		goto L467
	}
L467:
	;
	v1879 = v322
	v1880 = v1871
	goto L468
L468:
	;
	v1883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1880)+1)))
	v1884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1879)+1)))
	if v1884 == int32(0) {
		v1894 = v1883
		v1895 = v1884
		goto L465
	} else {
		goto L470
	}
L469:
	;
	v1894 = v1883
	v1895 = v1884
	goto L465
L470:
	;
	v1887 = int32(1)
	if v1883 == v1884 {
		v1879 = v1879 + v1887
		v1880 = v1880 + v1887
		goto L468
	} else {
		goto L471
	}
L471:
	;
	goto L469
L472:
	;
	if v1870 != int32(40960) {
		goto L463
	} else {
		goto L473
	}
L473:
	;
	v1904 = F_readlink(m, v330+int32(2368), v330+int32(192), int32(1024))
	mBase = m.M
	if v1904 < int32(0) {
		goto L87
	} else {
		goto L474
	}
L474:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v1904) {
		goto L86
	} else {
		goto L475
	}
L475:
	;
	v1910 = v330 + int32(192)
	v1912 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1910+v1904))) = uint8(v1912)
	F__tarWriteHeader(m, v321, v348, v1910, v330+int32(2272), v324)
	mBase = m.M
	v1919 = m.ExcPending
	if v1919 != 0 {
		goto L4
	} else {
		goto L476
	}
L476:
	;
	v4098 = v366 + int64(512)
	goto L96
L477:
	;
	v4036 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4037 = m.ExcPending
	if v4037 != 0 {
		goto L4
	} else {
		goto L771
	}
L478:
	;
	if v1870 != int32(16384) {
		goto L477
	} else {
		goto L481
	}
L479:
	;
	goto L480
L480:
	;
	v2149 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2256)) = v2149
	*(*int32)(unsafe.Add(mBase, uint32(v330)+2252)) = v2149
	if base.B2i32(v329 == v2149)|(v1032^int32(1)) == v2149 {
		goto L524
	} else {
		goto L525
	}
L481:
	;
	v1926 = int32(0)
	F__tarWriteHeader(m, v321, v348, v1926, v330+int32(2272), v324)
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L4
	} else {
		goto L482
	}
L482:
	;
	if v325 == int32(0) {
		v2037 = v1926
		goto L483
	} else {
		goto L484
	}
L483:
	;
	v2074 = v366 + int64(512)
	v2076 = v330 + int32(2368)
	v2077 = int32(489143)
	v2078 = int32(12)
	goto L506
L484:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	if v1934 <= int32(0) {
		v2037 = v1926
		goto L483
	} else {
		goto L485
	}
L485:
	;
	v1937 = int32(0)
	if v1937 < v1934 {
		goto L486
	} else {
		goto L487
	}
L486:
	;
	v1940 = v1934
	goto L488
L487:
	;
	v1940 = v1937
	goto L488
L488:
	;
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	v1953 = v1926
	goto L489
L489:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v1941+v1953<<(uint(int32(2))%32))))
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+8))
	if v1993 == int32(0) {
		goto L491
	} else {
		goto L492
	}
L490:
	;
	v2037 = int32(0)
	goto L483
L491:
	;
	v2023 = v1953 + int32(1)
	if v2023 != v1940 {
		v1953 = v2023
		goto L489
	} else {
		goto L502
	}
L492:
	;
	v1998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	v1999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1993))))
	if v1999 == int32(0) {
		v2018 = v1998
		v2019 = v1999
		goto L494
	} else {
		goto L495
	}
L493:
	;
	if v2019-v2018 != 0 {
		goto L491
	} else {
		goto L501
	}
L494:
	;
	goto L493
L495:
	;
	if v1998 != v1999 {
		v2018 = v1998
		v2019 = v1999
		goto L494
	} else {
		goto L496
	}
L496:
	;
	v2003 = v1993
	v2004 = v365
	goto L497
L497:
	;
	v2007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2004)+1)))
	v2008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2003)+1)))
	if v2008 == int32(0) {
		v2018 = v2007
		v2019 = v2008
		goto L494
	} else {
		goto L499
	}
L498:
	;
	v2018 = v2007
	v2019 = v2008
	goto L494
L499:
	;
	v2011 = int32(1)
	if v2007 == v2008 {
		v2003 = v2003 + v2011
		v2004 = v2004 + v2011
		goto L497
	} else {
		goto L500
	}
L500:
	;
	goto L498
L501:
	;
	v2037 = int32(1)
	goto L483
L502:
	;
	goto L490
L503:
	;
	if v2037 != 0 {
		v4098 = v2074
		goto L96
	} else {
		goto L521
	}
L504:
	;
	v2140 = int32(0)
	goto L503
L505:
	;
	v2114 = v2109
	v2115 = v2110
	v2116 = v2111
	goto L515
L506:
	;
	if (v2076|v2077)&int32(3) != 0 {
		v2109 = v2076
		v2110 = v2077
		v2111 = v2078
		goto L505
	} else {
		goto L509
	}
L508:
	;
	if v2099 == int32(0) {
		goto L504
	} else {
		goto L514
	}
L509:
	;
	v2086 = v2076
	v2087 = v2077
	v2088 = v2078
	goto L510
L510:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v2086)))
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v2087)))
	if v2091 != v2092 {
		v2109 = v2086
		v2110 = v2087
		v2111 = v2088
		goto L505
	} else {
		goto L512
	}
L511:
	;
	goto L508
L512:
	;
	v2094 = int32(4)
	v2095 = v2087 + v2094
	v2097 = v2086 + v2094
	v2099 = v2088 - v2094
	if base.Ui32(int32(3)) < base.Ui32(v2099) {
		v2086 = v2097
		v2087 = v2095
		v2088 = v2099
		goto L510
	} else {
		goto L513
	}
L513:
	;
	goto L511
L514:
	;
	v2109 = v2097
	v2110 = v2095
	v2111 = v2099
	goto L505
L515:
	;
	v2119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2114))))
	v2120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2115))))
	if v2119 == v2120 {
		goto L517
	} else {
		goto L518
	}
L516:
	;
	v2140 = v2119 - v2120
	goto L503
L517:
	;
	v2122 = int32(1)
	v2127 = v2116 - v2122
	if v2127 != 0 {
		v2114 = v2114 + v2122
		v2115 = v2115 + v2122
		v2116 = v2127
		goto L515
	} else {
		goto L520
	}
L518:
	;
	goto L519
L519:
	;
	goto L516
L520:
	;
	goto L504
L521:
	;
	if base.B2i32(v2140 == int32(0))&v363 != 0 {
		v4098 = v2074
		goto L96
	} else {
		goto L522
	}
L522:
	;
	v2146 = F_sendDir(m, v321, v330+int32(2368), v323, v324, v325, v326, v327, v328, v329)
	mBase = m.M
	v2147 = m.ExcPending
	if v2147 != 0 {
		goto L4
	} else {
		goto L523
	}
L523:
	;
	v4098 = v2146 + v2074
	goto L96
L524:
	;
	if v328 != 0 {
		goto L528
	} else {
		goto L529
	}
L525:
	;
	v3976 = v2149
	v3990 = v348
	goto L526
L526:
	;
	if v324 == int32(0) {
		goto L766
	} else {
		goto L767
	}
L527:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2268))
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2264))
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2260))
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2296))
	v2179 = v330 + int32(2256)
	v2181 = v330 + int32(2252)
	v2182 = int32(0)
	v2183 = m.G0
	v2185 = v2183 - int32(128)
	m.G0 = v2185
	if v2177&int32(8191) != 0 {
		v3874 = v2182
		goto L535
	} else {
		goto L536
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+120)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v330)+116)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v330)+112)) = int32(489145)
	v2168 = F_psprintf(m, int32(176985), v330+int32(112))
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L4
	} else {
		goto L531
	}
L529:
	;
	goto L530
L530:
	;
	v2170 = F_pstrdup(m, v348)
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L4
	} else {
		goto L532
	}
L531:
	;
	v2172 = v2168
	v2173 = v328
	goto L527
L532:
	;
	v2172 = v2170
	v2173 = v356
	goto L527
L533:
	;
	if v3874 == int32(1) {
		goto L758
	} else {
		goto L759
	}
L534:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3904 = m.ExcPending
	if v3904 != 0 {
		goto L4
	} else {
		goto L754
	}
L535:
	;
	m.G0 = v2185 + int32(128)
	goto L533
L536:
	;
	if v2175 == int32(1) {
		v3874 = v2182
		goto L535
	} else {
		goto L537
	}
L537:
	;
	if base.Ui32(int32(1073741824)) < base.Ui32(v2177) {
		v3874 = v2182
		goto L535
	} else {
		goto L538
	}
L538:
	;
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v329)+24))
	v2194 = F_strlen(m, v2172)
	mBase = m.M
	v2200 = v2194 - int32(1636608432)
	if v2172&int32(3) != 0 {
		goto L543
	} else {
		goto L544
	}
L539:
	;
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v2193)+20))
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(v2193)+12))
	v2461 = (v2454 ^ v2446 - base.I32_rotl(v2454, int32(24))) & v2460
	v2465 = *(*int32)(unsafe.Add(mBase, uint32(v2459+v2461<<(uint(int32(4))%32))))
	if v2465 != 0 {
		goto L580
	} else {
		goto L581
	}
L540:
	;
	v2432 = int32(14)
	v2434 = v2428 ^ v2429 - base.I32_rotl(v2428, v2432)
	v2438 = v2434 ^ v2427 - base.I32_rotl(v2434, int32(11))
	v2442 = v2438 ^ v2428 - base.I32_rotl(v2438, int32(25))
	v2446 = v2442 ^ v2434 - base.I32_rotl(v2442, int32(16))
	v2450 = v2446 ^ v2438 - base.I32_rotl(v2446, int32(4))
	v2454 = v2450 ^ v2442 - base.I32_rotl(v2450, v2432)
	goto L539
L541:
	;
	switch v2358 - int32(1) {
	case 0:
		v2420 = v2359
		v2421 = v2360
		v2422 = v2361
		goto L568
	case 1:
		v2413 = v2359
		v2414 = v2360
		v2415 = v2361
		goto L569
	case 2:
		v2406 = v2359
		v2407 = v2360
		v2408 = v2361
		goto L570
	case 3:
		v2400 = v2360
		v2401 = v2361
		goto L571
	case 4:
		v2396 = v2360
		v2397 = v2361
		goto L572
	case 5:
		v2390 = v2360
		v2391 = v2361
		goto L573
	case 6:
		v2384 = v2360
		v2385 = v2361
		goto L574
	case 7:
		v2379 = v2361
		goto L575
	case 8:
		v2374 = v2361
		goto L576
	case 9:
		v2369 = v2361
		goto L577
	case 10:
		goto L578
	default:
		v2427 = v2359
		v2428 = v2360
		v2429 = v2361
		goto L540
	}
L542:
	;
	v2309 = v2172
	v2310 = v2194
	v2311 = v2200
	v2312 = v2200
	v2313 = v2200
	goto L565
L543:
	;
	if base.Ui32(int32(11)) < base.Ui32(v2194) {
		goto L542
	} else {
		goto L546
	}
L544:
	;
	goto L545
L545:
	;
	if base.Ui32(v2194) < base.Ui32(int32(12)) {
		goto L548
	} else {
		goto L549
	}
L546:
	;
	v2357 = v2172
	v2358 = v2194
	v2359 = v2200
	v2360 = v2200
	v2361 = v2200
	goto L541
L547:
	;
	switch v2256 - int32(1) {
	case 0:
		v2306 = v2257
		goto L554
	case 1:
		v2301 = v2257
		goto L555
	case 2:
		goto L556
	case 3:
		v2294 = v2258
		goto L557
	case 4:
		v2291 = v2258
		goto L558
	case 5:
		v2286 = v2258
		goto L559
	case 6:
		goto L560
	case 7:
		v2277 = v2259
		goto L561
	case 8:
		v2272 = v2259
		goto L562
	case 9:
		v2267 = v2259
		goto L563
	case 10:
		goto L564
	default:
		v2427 = v2257
		v2428 = v2258
		v2429 = v2259
		goto L540
	}
L548:
	;
	v2255 = v2172
	v2256 = v2194
	v2257 = v2200
	v2258 = v2200
	v2259 = v2200
	goto L547
L549:
	;
	goto L550
L550:
	;
	v2207 = v2172
	v2208 = v2194
	v2209 = v2200
	v2210 = v2200
	v2211 = v2200
	goto L551
L551:
	;
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v2207)+4))
	v2214 = v2213 + v2210
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v2207)))
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v2207)+8))
	v2218 = v2217 + v2211
	v2220 = int32(4)
	v2222 = v2215 + v2209 - v2218 ^ base.I32_rotl(v2218, v2220)
	v2226 = v2214 - v2222 ^ base.I32_rotl(v2222, int32(6))
	v2227 = v2218 + v2214
	v2228 = v2222 + v2227
	v2229 = v2226 + v2228
	v2233 = v2227 - v2226 ^ base.I32_rotl(v2226, int32(8))
	v2237 = v2228 - v2233 ^ base.I32_rotl(v2233, int32(16))
	v2241 = v2229 - v2237 ^ base.I32_rotl(v2237, int32(19))
	v2242 = v2233 + v2229
	v2243 = v2237 + v2242
	v2244 = v2241 + v2243
	v2248 = v2242 - v2241 ^ base.I32_rotl(v2241, v2220)
	v2249 = int32(12)
	v2250 = v2207 + v2249
	v2252 = v2208 - v2249
	if base.Ui32(int32(11)) < base.Ui32(v2252) {
		v2207 = v2250
		v2208 = v2252
		v2209 = v2243
		v2210 = v2244
		v2211 = v2248
		goto L551
	} else {
		goto L553
	}
L552:
	;
	v2255 = v2250
	v2256 = v2252
	v2257 = v2243
	v2258 = v2244
	v2259 = v2248
	goto L547
L553:
	;
	goto L552
L554:
	;
	v2307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255))))
	v2427 = v2306 + v2307
	v2428 = v2258
	v2429 = v2259
	goto L540
L555:
	;
	v2302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255)+1)))
	v2306 = v2302<<(uint(int32(8))%32) + v2301
	goto L554
L556:
	;
	v2297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255)+2)))
	v2301 = v2297<<(uint(int32(16))%32) + v2257
	goto L555
L557:
	;
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(v2255)))
	v2427 = v2295 + v2257
	v2428 = v2294
	v2429 = v2259
	goto L540
L558:
	;
	v2292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255)+4)))
	v2294 = v2291 + v2292
	goto L557
L559:
	;
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255)+5)))
	v2291 = v2287<<(uint(int32(8))%32) + v2286
	goto L558
L560:
	;
	v2282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255)+6)))
	v2286 = v2282<<(uint(int32(16))%32) + v2258
	goto L559
L561:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2255)))
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(v2255)+4))
	v2427 = v2278 + v2257
	v2428 = v2280 + v2258
	v2429 = v2277
	goto L540
L562:
	;
	v2273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255)+8)))
	v2277 = v2273<<(uint(int32(8))%32) + v2272
	goto L561
L563:
	;
	v2268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255)+9)))
	v2272 = v2268<<(uint(int32(16))%32) + v2267
	goto L562
L564:
	;
	v2263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255)+10)))
	v2267 = v2263<<(uint(int32(24))%32) + v2259
	goto L563
L565:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v2309)+4))
	v2316 = v2315 + v2312
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(v2309)))
	v2319 = *(*int32)(unsafe.Add(mBase, uint32(v2309)+8))
	v2320 = v2319 + v2313
	v2322 = int32(4)
	v2324 = v2317 + v2311 - v2320 ^ base.I32_rotl(v2320, v2322)
	v2328 = v2316 - v2324 ^ base.I32_rotl(v2324, int32(6))
	v2329 = v2320 + v2316
	v2330 = v2324 + v2329
	v2331 = v2328 + v2330
	v2335 = v2329 - v2328 ^ base.I32_rotl(v2328, int32(8))
	v2339 = v2330 - v2335 ^ base.I32_rotl(v2335, int32(16))
	v2343 = v2331 - v2339 ^ base.I32_rotl(v2339, int32(19))
	v2344 = v2335 + v2331
	v2345 = v2339 + v2344
	v2346 = v2343 + v2345
	v2350 = v2344 - v2343 ^ base.I32_rotl(v2343, v2322)
	v2351 = int32(12)
	v2352 = v2309 + v2351
	v2354 = v2310 - v2351
	if base.Ui32(int32(11)) < base.Ui32(v2354) {
		v2309 = v2352
		v2310 = v2354
		v2311 = v2345
		v2312 = v2346
		v2313 = v2350
		goto L565
	} else {
		goto L567
	}
L566:
	;
	v2357 = v2352
	v2358 = v2354
	v2359 = v2345
	v2360 = v2346
	v2361 = v2350
	goto L541
L567:
	;
	goto L566
L568:
	;
	v2423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357))))
	v2427 = v2420 + v2423
	v2428 = v2421
	v2429 = v2422
	goto L540
L569:
	;
	v2416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357)+1)))
	v2420 = v2416<<(uint(int32(8))%32) + v2413
	v2421 = v2414
	v2422 = v2415
	goto L568
L570:
	;
	v2409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357)+2)))
	v2413 = v2409<<(uint(int32(16))%32) + v2406
	v2414 = v2407
	v2415 = v2408
	goto L569
L571:
	;
	v2402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357)+3)))
	v2406 = v2402<<(uint(int32(24))%32) + v2359
	v2407 = v2400
	v2408 = v2401
	goto L570
L572:
	;
	v2398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357)+4)))
	v2400 = v2396 + v2398
	v2401 = v2397
	goto L571
L573:
	;
	v2392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357)+5)))
	v2396 = v2392<<(uint(int32(8))%32) + v2390
	v2397 = v2391
	goto L572
L574:
	;
	v2386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357)+6)))
	v2390 = v2386<<(uint(int32(16))%32) + v2384
	v2391 = v2385
	goto L573
L575:
	;
	v2380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357)+7)))
	v2384 = v2380<<(uint(int32(24))%32) + v2360
	v2385 = v2379
	goto L574
L576:
	;
	v2375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357)+8)))
	v2379 = v2375<<(uint(int32(8))%32) + v2374
	goto L575
L577:
	;
	v2370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357)+9)))
	v2374 = v2370<<(uint(int32(16))%32) + v2369
	goto L576
L578:
	;
	v2365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357)+10)))
	v2369 = v2365<<(uint(int32(24))%32) + v2361
	goto L577
L579:
	;
	v3060 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2185)+64)) = v3060
	*(*int32)(unsafe.Add(mBase, uint32(v2185)+60)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v2185)+56)) = v2173
	v3065 = *(*int32)(unsafe.Add(mBase, uint32(v329)+28))
	v3067 = v2185 + int32(56)
	v3073 = m.G0
	v3074 = int32(16)
	v3075 = v3073 - v3074
	m.G0 = v3075
	v3077 = *(*int32)(unsafe.Add(mBase, uint32(v3067)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3075)+8)) = v3077
	v3079 = *(*int64)(unsafe.Add(mBase, uint32(v3067)))
	*(*int64)(unsafe.Add(mBase, uint32(v3075))) = v3079
	v3081 = *(*int32)(unsafe.Add(mBase, uint32(v3065)))
	*(*int32)(unsafe.Add(mBase, uint32(v3075)+12)) = v3060
	v3084 = F_hash_bytes(m, v3075, v3074)
	mBase = m.M
	v3085 = *(*int32)(unsafe.Add(mBase, uint32(v3081)+20))
	v3086 = *(*int32)(unsafe.Add(mBase, uint32(v3081)+12))
	v3087 = v3084 & v3086
	v3090 = v3085 + v3087*int32(40)
	v3091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3090)+20)))
	if v3091 == v3060 {
		goto L663
	} else {
		goto L664
	}
L580:
	;
	v2476 = v2461
	goto L583
L581:
	;
	goto L582
L582:
	;
	F_GetRelationPath(m, v2185+int32(56), v354, v2173, v2174, int32(-1), v2175)
	mBase = m.M
	v2602 = m.ExcPending
	if v2602 != 0 {
		goto L4
	} else {
		goto L595
	}
L583:
	;
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(v2459+v2476<<(uint(int32(4))%32))+4))
	v2519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2172))))
	v2520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2516))))
	if v2520 == int32(0) {
		v2539 = v2519
		v2540 = v2520
		goto L586
	} else {
		goto L587
	}
L584:
	;
	goto L582
L585:
	;
	if v2540-v2539 == int32(0) {
		goto L579
	} else {
		goto L593
	}
L586:
	;
	goto L585
L587:
	;
	if v2519 != v2520 {
		v2539 = v2519
		v2540 = v2520
		goto L586
	} else {
		goto L588
	}
L588:
	;
	v2524 = v2516
	v2525 = v2172
	goto L589
L589:
	;
	v2528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2525)+1)))
	v2529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2524)+1)))
	if v2529 == int32(0) {
		v2539 = v2528
		v2540 = v2529
		goto L586
	} else {
		goto L591
	}
L590:
	;
	v2539 = v2528
	v2540 = v2529
	goto L586
L591:
	;
	v2532 = int32(1)
	if v2528 == v2529 {
		v2524 = v2524 + v2532
		v2525 = v2525 + v2532
		goto L589
	} else {
		goto L592
	}
L592:
	;
	goto L590
L593:
	;
	v2546 = (v2476 + int32(1)) & v2460
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v2459+v2546<<(uint(int32(4))%32))))
	if v2550 != 0 {
		v2476 = v2546
		goto L583
	} else {
		goto L594
	}
L594:
	;
	goto L584
L595:
	;
	v2604 = v2185 + int32(56)
	v2608 = F_strlen(m, v2604)
	mBase = m.M
	v2615 = v2608 + int32(1)
	goto L598
L596:
	;
	v2628 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2627))) = uint8(v2628)
	v2631 = v2627 + int32(1)
	if v2176 != 0 {
		goto L603
	} else {
		goto L604
	}
L597:
	;
	goto L596
L598:
	;
	v2617 = int32(0)
	if v2615 == v2617 {
		v2627 = v2617
		goto L597
	} else {
		goto L600
	}
L599:
	;
	v2627 = v2622
	goto L597
L600:
	;
	v2621 = v2615 - int32(1)
	v2622 = v2604 + v2621
	v2623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2622))))
	if v2623 != int32(47) {
		v2615 = v2621
		goto L598
	} else {
		goto L601
	}
L601:
	;
	goto L599
L602:
	;
	v2652 = *(*int32)(unsafe.Add(mBase, uint32(v329)+24))
	v2653 = F_strlen(m, v2651)
	mBase = m.M
	v2659 = v2653 - int32(1636608432)
	if v2651&int32(3) != 0 {
		goto L612
	} else {
		goto L613
	}
L603:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2185)+40)) = v2176
	*(*int32)(unsafe.Add(mBase, uint32(v2185)+36)) = v2631
	*(*int32)(unsafe.Add(mBase, uint32(v2185)+32)) = v2185 + int32(56)
	v2640 = F_psprintf(m, int32(39289), v2185+int32(32))
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		goto L4
	} else {
		goto L606
	}
L604:
	;
	goto L605
L605:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2185)+20)) = v2631
	*(*int32)(unsafe.Add(mBase, uint32(v2185)+16)) = v2185 + int32(56)
	v2649 = F_psprintf(m, int32(177301), v2185+int32(16))
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		goto L4
	} else {
		goto L607
	}
L606:
	;
	v2651 = v2640
	goto L602
L607:
	;
	v2651 = v2649
	goto L602
L608:
	;
	v2918 = int32(0)
	v2919 = *(*int32)(unsafe.Add(mBase, uint32(v2652)+20))
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(v2652)+12))
	v2921 = (v2913 ^ v2905 - base.I32_rotl(v2913, int32(24))) & v2920
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v2919+v2921<<(uint(int32(4))%32))))
	if v2925 == v2918 {
		v3874 = v2918
		goto L535
	} else {
		goto L648
	}
L609:
	;
	v2891 = int32(14)
	v2893 = v2887 ^ v2888 - base.I32_rotl(v2887, v2891)
	v2897 = v2893 ^ v2886 - base.I32_rotl(v2893, int32(11))
	v2901 = v2897 ^ v2887 - base.I32_rotl(v2897, int32(25))
	v2905 = v2901 ^ v2893 - base.I32_rotl(v2901, int32(16))
	v2909 = v2905 ^ v2897 - base.I32_rotl(v2905, int32(4))
	v2913 = v2909 ^ v2901 - base.I32_rotl(v2909, v2891)
	goto L608
L610:
	;
	switch v2817 - int32(1) {
	case 0:
		v2879 = v2818
		v2880 = v2819
		v2881 = v2820
		goto L637
	case 1:
		v2872 = v2818
		v2873 = v2819
		v2874 = v2820
		goto L638
	case 2:
		v2865 = v2818
		v2866 = v2819
		v2867 = v2820
		goto L639
	case 3:
		v2859 = v2819
		v2860 = v2820
		goto L640
	case 4:
		v2855 = v2819
		v2856 = v2820
		goto L641
	case 5:
		v2849 = v2819
		v2850 = v2820
		goto L642
	case 6:
		v2843 = v2819
		v2844 = v2820
		goto L643
	case 7:
		v2838 = v2820
		goto L644
	case 8:
		v2833 = v2820
		goto L645
	case 9:
		v2828 = v2820
		goto L646
	case 10:
		goto L647
	default:
		v2886 = v2818
		v2887 = v2819
		v2888 = v2820
		goto L609
	}
L611:
	;
	v2768 = v2651
	v2769 = v2653
	v2770 = v2659
	v2771 = v2659
	v2772 = v2659
	goto L634
L612:
	;
	if base.Ui32(int32(11)) < base.Ui32(v2653) {
		goto L611
	} else {
		goto L615
	}
L613:
	;
	goto L614
L614:
	;
	if base.Ui32(v2653) < base.Ui32(int32(12)) {
		goto L617
	} else {
		goto L618
	}
L615:
	;
	v2816 = v2651
	v2817 = v2653
	v2818 = v2659
	v2819 = v2659
	v2820 = v2659
	goto L610
L616:
	;
	switch v2715 - int32(1) {
	case 0:
		v2765 = v2716
		goto L623
	case 1:
		v2760 = v2716
		goto L624
	case 2:
		goto L625
	case 3:
		v2753 = v2717
		goto L626
	case 4:
		v2750 = v2717
		goto L627
	case 5:
		v2745 = v2717
		goto L628
	case 6:
		goto L629
	case 7:
		v2736 = v2718
		goto L630
	case 8:
		v2731 = v2718
		goto L631
	case 9:
		v2726 = v2718
		goto L632
	case 10:
		goto L633
	default:
		v2886 = v2716
		v2887 = v2717
		v2888 = v2718
		goto L609
	}
L617:
	;
	v2714 = v2651
	v2715 = v2653
	v2716 = v2659
	v2717 = v2659
	v2718 = v2659
	goto L616
L618:
	;
	goto L619
L619:
	;
	v2666 = v2651
	v2667 = v2653
	v2668 = v2659
	v2669 = v2659
	v2670 = v2659
	goto L620
L620:
	;
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v2666)+4))
	v2673 = v2672 + v2669
	v2674 = *(*int32)(unsafe.Add(mBase, uint32(v2666)))
	v2676 = *(*int32)(unsafe.Add(mBase, uint32(v2666)+8))
	v2677 = v2676 + v2670
	v2679 = int32(4)
	v2681 = v2674 + v2668 - v2677 ^ base.I32_rotl(v2677, v2679)
	v2685 = v2673 - v2681 ^ base.I32_rotl(v2681, int32(6))
	v2686 = v2677 + v2673
	v2687 = v2681 + v2686
	v2688 = v2685 + v2687
	v2692 = v2686 - v2685 ^ base.I32_rotl(v2685, int32(8))
	v2696 = v2687 - v2692 ^ base.I32_rotl(v2692, int32(16))
	v2700 = v2688 - v2696 ^ base.I32_rotl(v2696, int32(19))
	v2701 = v2692 + v2688
	v2702 = v2696 + v2701
	v2703 = v2700 + v2702
	v2707 = v2701 - v2700 ^ base.I32_rotl(v2700, v2679)
	v2708 = int32(12)
	v2709 = v2666 + v2708
	v2711 = v2667 - v2708
	if base.Ui32(int32(11)) < base.Ui32(v2711) {
		v2666 = v2709
		v2667 = v2711
		v2668 = v2702
		v2669 = v2703
		v2670 = v2707
		goto L620
	} else {
		goto L622
	}
L621:
	;
	v2714 = v2709
	v2715 = v2711
	v2716 = v2702
	v2717 = v2703
	v2718 = v2707
	goto L616
L622:
	;
	goto L621
L623:
	;
	v2766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2714))))
	v2886 = v2765 + v2766
	v2887 = v2717
	v2888 = v2718
	goto L609
L624:
	;
	v2761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2714)+1)))
	v2765 = v2761<<(uint(int32(8))%32) + v2760
	goto L623
L625:
	;
	v2756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2714)+2)))
	v2760 = v2756<<(uint(int32(16))%32) + v2716
	goto L624
L626:
	;
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(v2714)))
	v2886 = v2754 + v2716
	v2887 = v2753
	v2888 = v2718
	goto L609
L627:
	;
	v2751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2714)+4)))
	v2753 = v2750 + v2751
	goto L626
L628:
	;
	v2746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2714)+5)))
	v2750 = v2746<<(uint(int32(8))%32) + v2745
	goto L627
L629:
	;
	v2741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2714)+6)))
	v2745 = v2741<<(uint(int32(16))%32) + v2717
	goto L628
L630:
	;
	v2737 = *(*int32)(unsafe.Add(mBase, uint32(v2714)))
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(v2714)+4))
	v2886 = v2737 + v2716
	v2887 = v2739 + v2717
	v2888 = v2736
	goto L609
L631:
	;
	v2732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2714)+8)))
	v2736 = v2732<<(uint(int32(8))%32) + v2731
	goto L630
L632:
	;
	v2727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2714)+9)))
	v2731 = v2727<<(uint(int32(16))%32) + v2726
	goto L631
L633:
	;
	v2722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2714)+10)))
	v2726 = v2722<<(uint(int32(24))%32) + v2718
	goto L632
L634:
	;
	v2774 = *(*int32)(unsafe.Add(mBase, uint32(v2768)+4))
	v2775 = v2774 + v2771
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(v2768)))
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(v2768)+8))
	v2779 = v2778 + v2772
	v2781 = int32(4)
	v2783 = v2776 + v2770 - v2779 ^ base.I32_rotl(v2779, v2781)
	v2787 = v2775 - v2783 ^ base.I32_rotl(v2783, int32(6))
	v2788 = v2779 + v2775
	v2789 = v2783 + v2788
	v2790 = v2787 + v2789
	v2794 = v2788 - v2787 ^ base.I32_rotl(v2787, int32(8))
	v2798 = v2789 - v2794 ^ base.I32_rotl(v2794, int32(16))
	v2802 = v2790 - v2798 ^ base.I32_rotl(v2798, int32(19))
	v2803 = v2794 + v2790
	v2804 = v2798 + v2803
	v2805 = v2802 + v2804
	v2809 = v2803 - v2802 ^ base.I32_rotl(v2802, v2781)
	v2810 = int32(12)
	v2811 = v2768 + v2810
	v2813 = v2769 - v2810
	if base.Ui32(int32(11)) < base.Ui32(v2813) {
		v2768 = v2811
		v2769 = v2813
		v2770 = v2804
		v2771 = v2805
		v2772 = v2809
		goto L634
	} else {
		goto L636
	}
L635:
	;
	v2816 = v2811
	v2817 = v2813
	v2818 = v2804
	v2819 = v2805
	v2820 = v2809
	goto L610
L636:
	;
	goto L635
L637:
	;
	v2882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816))))
	v2886 = v2879 + v2882
	v2887 = v2880
	v2888 = v2881
	goto L609
L638:
	;
	v2875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+1)))
	v2879 = v2875<<(uint(int32(8))%32) + v2872
	v2880 = v2873
	v2881 = v2874
	goto L637
L639:
	;
	v2868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+2)))
	v2872 = v2868<<(uint(int32(16))%32) + v2865
	v2873 = v2866
	v2874 = v2867
	goto L638
L640:
	;
	v2861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+3)))
	v2865 = v2861<<(uint(int32(24))%32) + v2818
	v2866 = v2859
	v2867 = v2860
	goto L639
L641:
	;
	v2857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+4)))
	v2859 = v2855 + v2857
	v2860 = v2856
	goto L640
L642:
	;
	v2851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+5)))
	v2855 = v2851<<(uint(int32(8))%32) + v2849
	v2856 = v2850
	goto L641
L643:
	;
	v2845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+6)))
	v2849 = v2845<<(uint(int32(16))%32) + v2843
	v2850 = v2844
	goto L642
L644:
	;
	v2839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+7)))
	v2843 = v2839<<(uint(int32(24))%32) + v2819
	v2844 = v2838
	goto L643
L645:
	;
	v2834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+8)))
	v2838 = v2834<<(uint(int32(8))%32) + v2833
	goto L644
L646:
	;
	v2829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+9)))
	v2833 = v2829<<(uint(int32(16))%32) + v2828
	goto L645
L647:
	;
	v2824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+10)))
	v2828 = v2824<<(uint(int32(24))%32) + v2820
	goto L646
L648:
	;
	v2938 = v2921
	goto L649
L649:
	;
	v2978 = *(*int32)(unsafe.Add(mBase, uint32(v2919+v2938<<(uint(int32(4))%32))+4))
	v2981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2651))))
	v2982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2978))))
	if v2982 == int32(0) {
		v3001 = v2981
		v3002 = v2982
		goto L652
	} else {
		goto L653
	}
L650:
	;
	v3874 = v2918
	goto L535
L651:
	;
	if v3002-v3001 == int32(0) {
		goto L579
	} else {
		goto L659
	}
L652:
	;
	goto L651
L653:
	;
	if v2981 != v2982 {
		v3001 = v2981
		v3002 = v2982
		goto L652
	} else {
		goto L654
	}
L654:
	;
	v2986 = v2978
	v2987 = v2651
	goto L655
L655:
	;
	v2990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2987)+1)))
	v2991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2986)+1)))
	if v2991 == int32(0) {
		v3001 = v2990
		v3002 = v2991
		goto L652
	} else {
		goto L657
	}
L656:
	;
	v3001 = v2990
	v3002 = v2991
	goto L652
L657:
	;
	v2994 = int32(1)
	if v2990 == v2991 {
		v2986 = v2986 + v2994
		v2987 = v2987 + v2994
		goto L655
	} else {
		goto L658
	}
L658:
	;
	goto L656
L659:
	;
	v3008 = (v2938 + int32(1)) & v2920
	v3012 = *(*int32)(unsafe.Add(mBase, uint32(v2919+v3008<<(uint(int32(4))%32))))
	if v3012 != 0 {
		v2938 = v3008
		goto L649
	} else {
		goto L660
	}
L660:
	;
	goto L650
L661:
	;
	if v3122 != 0 {
		v3874 = v3060
		goto L535
	} else {
		goto L674
	}
L662:
	;
	m.G0 = v3075 + int32(16)
	goto L661
L663:
	;
	v3122 = int32(0)
	goto L662
L664:
	;
	v3094 = v3087
	v3095 = v3090
	goto L665
L665:
	;
	v3101 = F_memcmp(m, v3095, v3075, int32(16))
	mBase = m.M
	if v3101 != 0 {
		goto L667
	} else {
		goto L668
	}
L666:
	;
	if v3095 == int32(0) {
		goto L671
	} else {
		goto L672
	}
L667:
	;
	v3104 = (v3094 + int32(1)) & v3086
	v3107 = v3085 + v3104*int32(40)
	v3108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3107)+20)))
	if v3108 != 0 {
		v3094 = v3104
		v3095 = v3107
		goto L665
	} else {
		goto L670
	}
L668:
	;
	goto L669
L669:
	;
	goto L666
L670:
	;
	goto L663
L671:
	;
	v3122 = int32(0)
	goto L662
L672:
	;
	goto L673
L673:
	;
	v3112 = *(*int32)(unsafe.Add(mBase, uint32(v3095)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2185+int32(52)))) = v3112
	v3122 = v3095
	goto L662
L674:
	;
	v3131 = int32(base.Ui32(v2177) >> (uint(int32(13)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v2185)+64)) = v2174
	v3133 = *(*int32)(unsafe.Add(mBase, uint32(v329)+28))
	v3135 = v2185 + int32(56)
	v3140 = m.G0
	v3141 = int32(16)
	v3142 = v3140 - v3141
	m.G0 = v3142
	v3144 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3142)+8)) = v3144
	v3146 = *(*int64)(unsafe.Add(mBase, uint32(v3135)))
	*(*int64)(unsafe.Add(mBase, uint32(v3142))) = v3146
	v3148 = *(*int32)(unsafe.Add(mBase, uint32(v3133)))
	*(*int32)(unsafe.Add(mBase, uint32(v3142)+12)) = v2175
	v3151 = F_hash_bytes(m, v3142, v3141)
	mBase = m.M
	v3152 = *(*int32)(unsafe.Add(mBase, uint32(v3148)+20))
	v3153 = *(*int32)(unsafe.Add(mBase, uint32(v3148)+12))
	v3154 = v3151 & v3153
	v3157 = v3152 + v3154*int32(40)
	v3158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3157)+20)))
	if v3158 == int32(0) {
		goto L678
	} else {
		goto L679
	}
L675:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2181))) = v3811
	v3874 = int32(1)
	goto L535
L676:
	;
	if v3189 == int32(0) {
		goto L689
	} else {
		goto L690
	}
L677:
	;
	m.G0 = v3142 + int32(16)
	goto L676
L678:
	;
	v3189 = int32(0)
	goto L677
L679:
	;
	v3161 = v3154
	v3162 = v3157
	goto L680
L680:
	;
	v3168 = F_memcmp(m, v3162, v3142, int32(16))
	mBase = m.M
	if v3168 != 0 {
		goto L682
	} else {
		goto L683
	}
L681:
	;
	if v3162 == int32(0) {
		goto L686
	} else {
		goto L687
	}
L682:
	;
	v3171 = (v3161 + int32(1)) & v3153
	v3174 = v3152 + v3171*int32(40)
	v3175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3174)+20)))
	if v3175 != 0 {
		v3161 = v3171
		v3162 = v3174
		goto L680
	} else {
		goto L685
	}
L683:
	;
	goto L684
L684:
	;
	goto L681
L685:
	;
	goto L678
L686:
	;
	v3189 = int32(0)
	goto L677
L687:
	;
	goto L688
L688:
	;
	v3179 = *(*int32)(unsafe.Add(mBase, uint32(v3162)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2185+int32(52)))) = v3179
	v3189 = v3162
	goto L677
L689:
	;
	if v2177 == int32(0) {
		v3874 = v3060
		goto L535
	} else {
		goto L692
	}
L690:
	;
	goto L691
L691:
	;
	v3204 = v2176 << (uint(int32(17)) % 32)
	v3205 = *(*int32)(unsafe.Add(mBase, uint32(v2185)+52))
	if base.Ui32(v3205) <= base.Ui32(v3204) {
		v3874 = v3060
		goto L535
	} else {
		goto L693
	}
L692:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2179))) = int32(0)
	v3811 = v3131
	goto L675
L693:
	;
	if base.Ui32(int32(32767)) < base.Ui32(v2176) {
		goto L534
	} else {
		goto L694
	}
L694:
	;
	v3209 = v3131 + v3204
	if base.Ui32(v3209) < base.Ui32(v3131) {
		goto L534
	} else {
		goto L695
	}
L695:
	;
	v3211 = int32(0)
	v3212 = int32(16)
	v3213 = int32(base.Ui32(v3204) >> (uint(v3212) % 32))
	v3220 = int32(base.Ui32(v3209)>>(uint(v3212)%32)) + base.B2i32(v3209&int32(65535) != v3211)
	v3221 = *(*int32)(unsafe.Add(mBase, uint32(v3189)+24))
	if base.Ui32(v3220) < base.Ui32(v3221) {
		goto L697
	} else {
		goto L698
	}
L696:
	;
	if base.F64_gt(base.F64_convert_i32_u(v3516<<(uint(int32(13))%32)), base.F64_mul(base.F64_convert_i32_u(v2177), float64(0.9))) != 0 {
		v3874 = v3060
		goto L535
	} else {
		goto L730
	}
L697:
	;
	v3223 = v3220
	goto L699
L698:
	;
	v3223 = v3221
	goto L699
L699:
	;
	if base.Ui32(v3223) <= base.Ui32(v3213) {
		v3516 = v3211
		goto L696
	} else {
		goto L700
	}
L700:
	;
	v3249 = v3213
	v3257 = v3211
	goto L701
L701:
	;
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(v3189)+32))
	v3280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3276+v3249<<(uint(int32(1))%32)))))
	if v3280 == int32(0) {
		v3466 = v3257
		goto L703
	} else {
		goto L704
	}
L702:
	;
	v3516 = v3466
	goto L696
L703:
	;
	v3486 = v3249 + int32(1)
	if v3486 != v3223 {
		v3249 = v3486
		v3257 = v3466
		goto L701
	} else {
		goto L729
	}
L704:
	;
	v3285 = v3249 << (uint(int32(16)) % 32)
	if v3223-int32(1) != v3249 {
		goto L705
	} else {
		goto L706
	}
L705:
	;
	v3288 = int32(65536)
	goto L707
L706:
	;
	v3288 = v3209 - v3285
	goto L707
L707:
	;
	if v3213 == v3249 {
		goto L708
	} else {
		goto L709
	}
L708:
	;
	v3291 = v3204 & int32(65535)
	goto L710
L709:
	;
	v3291 = int32(0)
	goto L710
L710:
	;
	v3292 = *(*int32)(unsafe.Add(mBase, uint32(v3189)+36))
	v3296 = *(*int32)(unsafe.Add(mBase, uint32(v3292+v3249<<(uint(int32(2))%32))))
	if v3280 != int32(4096) {
		goto L711
	} else {
		goto L712
	}
L711:
	;
	v3328 = v3257
	v3332 = int32(0)
	goto L714
L712:
	;
	goto L713
L713:
	;
	if base.Ui32(v3288) <= base.Ui32(v3291) {
		v3466 = v3257
		goto L703
	} else {
		goto L721
	}
L714:
	;
	v3350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3296+v3332<<(uint(int32(1))%32)))))
	if base.Ui32(v3350) < base.Ui32(v3291) {
		v3362 = v3328
		goto L716
	} else {
		goto L717
	}
L715:
	;
	v3466 = v3362
	goto L703
L716:
	;
	v3364 = v3332 + int32(1)
	if v3364 != v3280 {
		v3328 = v3362
		v3332 = v3364
		goto L714
	} else {
		goto L720
	}
L717:
	;
	if base.Ui32(v3288) <= base.Ui32(v3350) {
		v3362 = v3328
		goto L716
	} else {
		goto L718
	}
L718:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347+v3328<<(uint(int32(2))%32)))) = v3350 | v3285
	v3359 = v3328 + int32(1)
	if v3359 == int32(131072) {
		v3516 = v3359
		goto L696
	} else {
		goto L719
	}
L719:
	;
	v3362 = v3359
	goto L716
L720:
	;
	goto L715
L721:
	;
	v3392 = v3291
	v3395 = v3257
	goto L722
L722:
	;
	v3419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3296+int32(base.Ui32(v3392)>>(uint(int32(3))%32))&int32(536870910)))))
	if int32(base.Ui32(v3419)>>(uint(v3392&int32(15))%32))&int32(1) != 0 {
		goto L724
	} else {
		goto L725
	}
L723:
	;
	v3466 = v3434
	goto L703
L724:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347+v3395<<(uint(int32(2))%32)))) = v3392 + v3285
	v3431 = v3395 + int32(1)
	if v3431 == int32(131072) {
		v3516 = v3431
		goto L696
	} else {
		goto L727
	}
L725:
	;
	v3434 = v3395
	goto L726
L726:
	;
	v3436 = v3392 + int32(1)
	if v3436 != v3288 {
		v3392 = v3436
		v3395 = v3434
		goto L722
	} else {
		goto L728
	}
L727:
	;
	v3434 = v3431
	goto L726
L728:
	;
	goto L723
L729:
	;
	goto L702
L730:
	;
	F_pg_qsort(m, v347, v3516, int32(4), int32(434))
	mBase = m.M
	v3545 = m.ExcPending
	if v3545 != 0 {
		goto L4
	} else {
		goto L731
	}
L731:
	;
	if v3204 == int32(0) {
		goto L732
	} else {
		goto L733
	}
L732:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2179))) = v3516
	*(*int32)(unsafe.Add(mBase, uint32(v2181))) = v3131
	v3787 = int32(1)
	v3788 = *(*int32)(unsafe.Add(mBase, uint32(v2185)+52))
	if v3788 == int32(-1) {
		v3874 = v3787
		goto L535
	} else {
		goto L745
	}
L733:
	;
	if v3516 == int32(0) {
		goto L732
	} else {
		goto L734
	}
L734:
	;
	v3551 = v3516 & int32(3)
	v3552 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v3516) {
		goto L735
	} else {
		goto L736
	}
L735:
	;
	v3568 = v3552
	v3582 = v3060
	goto L738
L736:
	;
	v3641 = v3552
	goto L737
L737:
	;
	if v3551 == int32(0) {
		goto L732
	} else {
		goto L741
	}
L738:
	;
	v3607 = v347 + v3568<<(uint(int32(2))%32)
	v3608 = *(*int32)(unsafe.Add(mBase, uint32(v3607)))
	*(*int32)(unsafe.Add(mBase, uint32(v3607))) = v3608 - v3204
	v3611 = int32(4)
	v3612 = v3607 + v3611
	v3613 = *(*int32)(unsafe.Add(mBase, uint32(v3612)))
	*(*int32)(unsafe.Add(mBase, uint32(v3612))) = v3613 - v3204
	v3617 = v3607 + int32(8)
	v3618 = *(*int32)(unsafe.Add(mBase, uint32(v3617)))
	*(*int32)(unsafe.Add(mBase, uint32(v3617))) = v3618 - v3204
	v3622 = v3607 + int32(12)
	v3623 = *(*int32)(unsafe.Add(mBase, uint32(v3622)))
	*(*int32)(unsafe.Add(mBase, uint32(v3622))) = v3623 - v3204
	v3627 = v3568 + v3611
	v3629 = v3582 + v3611
	if v3629 != v3516&int32(-4) {
		v3568 = v3627
		v3582 = v3629
		goto L738
	} else {
		goto L740
	}
L739:
	;
	v3641 = v3627
	goto L737
L740:
	;
	goto L739
L741:
	;
	v3690 = v3641
	v3703 = v3552
	goto L742
L742:
	;
	v3729 = v347 + v3690<<(uint(int32(2))%32)
	v3730 = *(*int32)(unsafe.Add(mBase, uint32(v3729)))
	*(*int32)(unsafe.Add(mBase, uint32(v3729))) = v3730 - v3204
	v3733 = int32(1)
	v3736 = v3703 + v3733
	if v3736 != v3551 {
		v3690 = v3690 + v3733
		v3703 = v3736
		goto L742
	} else {
		goto L744
	}
L743:
	;
	goto L732
L744:
	;
	goto L743
L745:
	;
	v3792 = v3788 - v3204
	if base.Ui32(v3792) < base.Ui32(v3131) {
		goto L746
	} else {
		goto L747
	}
L746:
	;
	v3794 = v3131
	goto L748
L747:
	;
	v3794 = v3792
	goto L748
L748:
	;
	if base.Ui32(int32(131072)) <= base.Ui32(v3794) {
		goto L749
	} else {
		goto L750
	}
L749:
	;
	v3797 = int32(131072)
	goto L751
L750:
	;
	v3797 = v3794
	goto L751
L751:
	;
	if base.Ui32(v3131) < base.Ui32(v3792) {
		v3811 = v3797
		goto L675
	} else {
		goto L752
	}
L752:
	;
	if base.Ui32(v3794) < base.Ui32(int32(131073)) {
		v3874 = v3787
		goto L535
	} else {
		goto L753
	}
L753:
	;
	v3811 = v3797
	goto L675
L754:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v3907 = m.ExcPending
	if v3907 != 0 {
		goto L4
	} else {
		goto L755
	}
L755:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2185)+4)) = v2177
	*(*int32)(unsafe.Add(mBase, uint32(v2185))) = v2176
	F_errmsg_internal(m, int32(37492), v2185)
	mBase = m.M
	v3912 = m.ExcPending
	if v3912 != 0 {
		goto L4
	} else {
		goto L756
	}
L756:
	;
	F_errfinish(m, int32(497962), int32(796), int32(423447))
	mBase = m.M
	v3917 = m.ExcPending
	if v3917 != 0 {
		goto L4
	} else {
		goto L757
	}
L757:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L758:
	;
	v3920 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2256))
	v3922 = v3920 << (uint(int32(2)) % 32)
	v3924 = v3922 + int32(12)
	if v3920 == int32(0) {
		v3934 = v3924
		goto L761
	} else {
		goto L762
	}
L759:
	;
	v3956 = v2149
	v3957 = v348
	goto L760
L760:
	;
	F_pfree(m, v2172)
	mBase = m.M
	v3959 = m.ExcPending
	if v3959 != 0 {
		goto L4
	} else {
		goto L765
	}
L761:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+96)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v330)+100)) = v387
	*(*int64)(unsafe.Add(mBase, uint32(v330)+2296)) = base.I64_extend_i32_u(v3934 + v3920<<(uint(int32(13))%32))
	v3949 = F_pg_snprintf(m, v330+int32(192), int32(2048), int32(177301), v330+int32(96))
	mBase = m.M
	v3950 = m.ExcPending
	if v3950 != 0 {
		goto L4
	} else {
		goto L764
	}
L762:
	;
	v3928 = v3924 & int32(8188)
	if v3928 == int32(0) {
		v3934 = v3924
		goto L761
	} else {
		goto L763
	}
L763:
	;
	v3934 = v3922 - v3928 + int32(8204)
	goto L761
L764:
	;
	v3956 = v347
	v3957 = v330 + int32(192)
	goto L760
L765:
	;
	v3976 = v3956
	v3990 = v3957
	goto L526
L766:
	;
	v4014 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2268))
	v4015 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2260))
	v4016 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2256))
	v4017 = *(*int32)(unsafe.Add(mBase, uint32(v330)+2252))
	v4018 = F_sendFile(m, v321, v330+int32(2368), v3990, v330+int32(2272), int32(1), v354, v328, v4014, v4015, v327, v4016, v3976, v4017)
	mBase = m.M
	v4019 = m.ExcPending
	if v4019 != 0 {
		goto L4
	} else {
		goto L769
	}
L767:
	;
	goto L768
L768:
	;
	v4022 = *(*int64)(unsafe.Add(mBase, uint32(v330)+2296))
	v4098 = v366 + v4022 + ((v4022+int64(511))&int64(4294966784)-v4022)&int64(4294967295) + int64(512)
	goto L96
L769:
	;
	if v4018 == int32(0) {
		v4098 = v366
		goto L96
	} else {
		goto L770
	}
L770:
	;
	goto L768
L771:
	;
	if v4036 == int32(0) {
		v4098 = v366
		goto L96
	} else {
		goto L772
	}
L772:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+80)) = v330 + int32(2368)
	F_errmsg(m, int32(717152), v330+int32(80))
	mBase = m.M
	v4047 = m.ExcPending
	if v4047 != 0 {
		goto L4
	} else {
		goto L773
	}
L773:
	;
	F_errfinish(m, int32(496115), int32(1546), int32(213662))
	mBase = m.M
	v4052 = m.ExcPending
	if v4052 != 0 {
		goto L4
	} else {
		goto L774
	}
L774:
	;
	v4098 = v366
	goto L96
L775:
	;
	if v4100 != 0 {
		v332 = v4100
		v366 = v4098
		goto L94
	} else {
		goto L776
	}
L776:
	;
	goto L95
L777:
	;
	F_pfree(m, v4128)
	mBase = m.M
	v4150 = m.ExcPending
	if v4150 != 0 {
		goto L4
	} else {
		goto L780
	}
L778:
	;
	goto L779
L779:
	;
	F_FreeDir(m, v4133)
	mBase = m.M
	v4152 = m.ExcPending
	if v4152 != 0 {
		goto L4
	} else {
		goto L781
	}
L780:
	;
	goto L779
L781:
	;
	m.G0 = v4111 + int32(4416)
	return v4147
L782:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v4163 = m.ExcPending
	if v4163 != 0 {
		goto L4
	} else {
		goto L783
	}
L783:
	;
	F_errmsg(m, int32(233748), int32(0))
	mBase = m.M
	v4167 = m.ExcPending
	if v4167 != 0 {
		goto L4
	} else {
		goto L784
	}
L784:
	;
	F_errhint(m, int32(612088), int32(0))
	mBase = m.M
	v4171 = m.ExcPending
	if v4171 != 0 {
		goto L4
	} else {
		goto L785
	}
L785:
	;
	F_errfinish(m, int32(496115), int32(1286), int32(213662))
	mBase = m.M
	v4176 = m.ExcPending
	if v4176 != 0 {
		goto L4
	} else {
		goto L786
	}
L786:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L787:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4182 = m.ExcPending
	if v4182 != 0 {
		goto L4
	} else {
		goto L788
	}
L788:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+48)) = v330 + int32(2368)
	F_errmsg(m, int32(297749), v330+int32(48))
	mBase = m.M
	v4190 = m.ExcPending
	if v4190 != 0 {
		goto L4
	} else {
		goto L789
	}
L789:
	;
	F_errfinish(m, int32(496115), int32(1420), int32(213662))
	mBase = m.M
	v4195 = m.ExcPending
	if v4195 != 0 {
		goto L4
	} else {
		goto L790
	}
L790:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L791:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v4202 = m.ExcPending
	if v4202 != 0 {
		goto L4
	} else {
		goto L792
	}
L792:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+64)) = v330 + int32(2368)
	F_errmsg(m, int32(327990), v330-int32(-64))
	mBase = m.M
	v4210 = m.ExcPending
	if v4210 != 0 {
		goto L4
	} else {
		goto L793
	}
L793:
	;
	F_errfinish(m, int32(496115), int32(1425), int32(213662))
	mBase = m.M
	v4215 = m.ExcPending
	if v4215 != 0 {
		goto L4
	} else {
		goto L794
	}
L794:
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
		*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0) - v7
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[689]))
	v5 = F_MemoryContextStrdup(m, v4, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[706])) = v5
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
		v26 = *(*float64)(unsafe.Add(mBase, _consts[601]))
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
				F_errmsg(m, int32(508167), v7)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495107), int32(70), int32(460342))
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
		v40 = int32(4500560)
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
		*(*int64)(unsafe.Add(mBase, _consts[1030])) = int64(base.Ui64(v63)>>(uint(v64)%64)) ^ v63
		v69 = v51 - int64(7046029254386353131)
		v74 = (int64(base.Ui64(v69)>>(uint(v54)%64)) ^ v69) * v57
		v79 = (int64(base.Ui64(v74)>>(uint(v59)%64)) ^ v74) * v62
		*(*int64)(unsafe.Add(mBase, _consts[1031])) = int64(base.Ui64(v79)>>(uint(v64)%64)) ^ v79
		if v69|v53 == int64(0) {
			*(*int64)(unsafe.Add(mBase, _consts[1030])) = int64(1442695040888963407)
			*(*int64)(unsafe.Add(mBase, _consts[1031])) = int64(6364136223846793005)
		} else {
		}
		v92 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[1032])) = uint8(v92)
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
	v89 = *(*int32)(unsafe.Add(mBase, _consts[128]))
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
	if l1 <= int32(5999) {
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
	v43 = l1 - int32(4177)
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
	switch l1 - int32(6243) {
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
	if base.Ui32(l1-int32(6000)) < base.Ui32(int32(3)) {
		v87 = v19
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v59 = l1 - int32(6100)
	if base.Ui32(int32(15)) < base.Ui32(v59) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	if int32(1)<<(uint(v59)%32)&int32(49153) != 0 {
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
	if base.Ui32(int32(11999)) < base.Ui32(l4) {
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
	F_errmsg_internal(m, int32(501194), v15)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L29
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(500056), int32(255), int32(238079))
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
							F_errmsg(m, int32(452018), v4+int32(-48))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_errfinish(m, int32(500056), int32(1223), int32(111565))
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
				F_errmsg_internal(m, int32(59001), v6)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					F_errfinish(m, int32(500056), int32(1256), int32(111565))
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
							F_errmsg(m, int32(451981), v4+int32(-16))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								F_errfinish(m, int32(500056), int32(1249), int32(111565))
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
								F_errmsg(m, int32(452051), v4+int32(-32))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									F_errfinish(m, int32(500056), int32(1235), int32(111565))
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
					F_errmsg_internal(m, int32(59001), v6)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errfinish(m, int32(500056), int32(1256), int32(111565))
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
			F_errmsg(m, int32(369561), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(494551), int32(317), int32(67158))
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
	*(*uint8)(unsafe.Add(mBase, _consts[829])) = uint8(v9)
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
	v19 = *(*int32)(unsafe.Add(mBase, _consts[830]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v19
	F_errmsg_internal(m, int32(399092), v6+int32(16))
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
	v31 = int32(4431800)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[830]))
	v35 = v33 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[830])) = v35
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	F_errfinish(m, int32(500241), int32(248), int32(99458))
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
	*(*int32)(unsafe.Add(mBase, _consts[830])) = v64
	v67 = *(*int32)(unsafe.Add(mBase, _consts[826]))
	if v67 == v64 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v43 = v40 << (uint(int32(3)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[831])))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[832])))
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
	v52 = int32(4431800)
	v54 = *(*int32)(unsafe.Add(mBase, _consts[830]))
	v56 = v54 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[830])) = v56
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
	if v67 == int32(4122148) {
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
	v78 = *(*int32)(unsafe.Add(mBase, _consts[826]))
	if v78 == int32(0) {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	if v78 != int32(4122148) {
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
	v92 = *(*int32)(unsafe.Add(mBase, _consts[833]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v92
	F_errmsg_internal(m, int32(399041), v6)
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
	v102 = int32(4431968)
	v104 = *(*int32)(unsafe.Add(mBase, _consts[833]))
	v106 = v104 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[833])) = v106
	if int32(0) <= v106 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	F_errfinish(m, int32(500241), int32(281), int32(99458))
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
	*(*uint8)(unsafe.Add(mBase, _consts[829])) = uint8(v135)
	*(*int32)(unsafe.Add(mBase, _consts[833])) = v135
	m.G0 = v6 + int32(32)
	return
L33:
	;
	v114 = v111 << (uint(int32(3)) % 32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+uint32(_consts[834])))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114)+uint32(_consts[835])))
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
	v123 = int32(4431968)
	v125 = *(*int32)(unsafe.Add(mBase, _consts[833]))
	v127 = v125 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[833])) = v127
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
	v10 = F_execute(m, l0+v5<<(uint(int32(3))%32), l1, l2, l3, int32(7115))
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
	v10 = F__emscripten_memcpy_bulkmem(m, l1, int32(4680692), int32(128))
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
	v14 = F__emscripten_memcpy_bulkmem(m, int32(4680692), l0, int32(128))
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
	v37 = int32(4680692)
	v38 = *(*int32)(unsafe.Add(mBase, _consts[1381]))
	*(*int32)(unsafe.Add(mBase, _consts[1381])) = v38 & base.I32_rotl(int32(-2), int32(8))
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
	v65 = int32(4680692)
	v66 = *(*int32)(unsafe.Add(mBase, _consts[1381]))
	*(*int32)(unsafe.Add(mBase, _consts[1381])) = v66 & base.I32_rotl(int32(-2), int32(18))
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
	v87 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v79)>>(uint(int32(3))%32))&int32(536870908))+uint32(_consts[1382])))
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
	v105 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v97)>>(uint(int32(3))%32))&int32(536870908))+uint32(_consts[1381])))
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
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(28)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v129 = int32(base.Ui32(v113)>>(uint(int32(3))%32)) & int32(536870908)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+uint32(_consts[1382])))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+uint32(_consts[1382]))) = v131 & base.I32_rotl(int32(-2), v113)
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
		v24 = F_LocalToUtf(m, v6, v10, v5, int32(4392252), v18, v18, v18, int32(35), base.B2i32(v7 != v18))
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
	v21 = *(*int32)(unsafe.Add(mBase, _consts[661]))
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
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[662]))))
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
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19<<(uint(int32(1))%32))+uint32(_consts[664]))))
	v79 = int32(base.Ui32(v69&int32(32)) >> (uint(int32(5)) % 32))
	goto L13
L17:
	;
	v45 = base.I32_div_s(v39+v40, int32(2))
	v47 = v45 << (uint(int32(3)) % 32)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[665])))
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
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[666])))
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
	v92 = int32(4092160)
	goto L35
L33:
	;
	goto L34
L34:
	;
	v102 = int32(4092160)
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
	v109 = int32(4092160)
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
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _consts[82])))
	if v5 == int32(1) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+20))
		if v11 != int32(-1) {
			v41 = v11
			return v41
		} else {
			v15 = int32(4510364)
			v17 = *(*int32)(unsafe.Add(mBase, _consts[115]))
			*(*int32)(unsafe.Add(mBase, _consts[115])) = v17 + int32(1)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v24*int32(80))+uint32(_consts[880])))
			v30 = m.T0[v29].(func(*base.Module, int32, int32) int32)(m, l0, l1)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+20)) = v30
				v35 = int32(4510364)
				v37 = *(*int32)(unsafe.Add(mBase, _consts[115]))
				*(*int32)(unsafe.Add(mBase, _consts[115])) = v37 - int32(1)
				v41 = v30
				return v41
			}
		}
	} else {
		v15 = int32(4510364)
		v17 = *(*int32)(unsafe.Add(mBase, _consts[115]))
		*(*int32)(unsafe.Add(mBase, _consts[115])) = v17 + int32(1)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v24*int32(80))+uint32(_consts[880])))
		v30 = m.T0[v29].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+20)) = v30
			v35 = int32(4510364)
			v37 = *(*int32)(unsafe.Add(mBase, _consts[115]))
			*(*int32)(unsafe.Add(mBase, _consts[115])) = v37 - int32(1)
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
	v3 = int32(4510364)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	v6 = int32(1)
	v7 = v5 + v6
	*(*int32)(unsafe.Add(mBase, _consts[115])) = v7
	*(*int32)(unsafe.Add(mBase, _consts[115])) = v7 - v6
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
		v13 = *(*int32)(unsafe.Add(mBase, _consts[878]))
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, _consts[879]))
			v20 = v15
		} else {
			v17 = int32(4439192)
			*(*int32)(unsafe.Add(mBase, _consts[878])) = v17
			v20 = v17
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v20
		v22 = int32(4439192)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v11
		*(*int32)(unsafe.Add(mBase, _consts[879])) = v11
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
	v23 = *(*int32)(unsafe.Add(mBase, _consts[434]))
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
	v61 = *(*int32)(unsafe.Add(mBase, _consts[435]))
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
	*(*int32)(unsafe.Add(mBase, _consts[434])) = v1400
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
	v275 = *(*int32)(unsafe.Add(mBase, _consts[436]))
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
	v282 = *(*int32)(unsafe.Add(mBase, _consts[435]))
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
	F_errmsg_internal(m, int32(401023), int32(0))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L63
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(326930), int32(327), int32(341377))
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
	F_errmsg_internal(m, int32(401023), int32(0))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L63
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(326930), int32(327), int32(341377))
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
	F_errmsg_internal(m, int32(462923), int32(0))
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L63
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(326930), int32(630), int32(311968))
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
			v12 = int32(65535)
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
			if v15 != 0 {
				v29 = v12
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				v19 = v16 + v10*int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(2)
				*(*int64)(unsafe.Add(mBase, uint32(v19)+12)) = int64(0)
				v24 = int32(65535)
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
					v44 = int32(65535)
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
					v54 = int32(65535)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
					if v57 != 0 {
						v71 = v54
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
						v61 = v58 + v52*int32(24)
						*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = int32(2)
						*(*int64)(unsafe.Add(mBase, uint32(v61)+12)) = int64(0)
						v66 = int32(65535)
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
							v86 = int32(65535)
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
		v21 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v26 = F_AllocSetContextCreateInternal(m, v21, int32(60580), int32(0), int32(8192), int32(8388608))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = int32(4515712)
			v29 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
			v32 = v26
			v33 = v29
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
			if v34 == int32(1) {
				v41 = F__emscripten_memcpy_bulkmem(m, v15+int32(152), int32(4413976), int32(128))
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
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v33
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
							v90 = *(*int64)(unsafe.Add(mBase, _consts[52]))
							v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
							*(*int64)(unsafe.Add(mBase, uint32(v85))) = v88 + (v90 - v91)
							v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
							v97 = *(*int64)(unsafe.Add(mBase, _consts[53]))
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v95 + (v97 - v98)
							v102 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
							v104 = *(*int64)(unsafe.Add(mBase, _consts[54]))
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v102 + (v104 - v105)
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
							v111 = *(*int64)(unsafe.Add(mBase, _consts[55]))
							v112 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v109 + (v111 - v112)
							v116 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
							v118 = *(*int64)(unsafe.Add(mBase, _consts[56]))
							v119 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v116 + (v118 - v119)
							v123 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
							v125 = *(*int64)(unsafe.Add(mBase, _consts[57]))
							v126 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v123 + (v125 - v126)
							v130 = *(*int64)(unsafe.Add(mBase, uint32(v85)+48))
							v132 = *(*int64)(unsafe.Add(mBase, _consts[58]))
							v133 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v130 + (v132 - v133)
							v137 = *(*int64)(unsafe.Add(mBase, uint32(v85)+56))
							v139 = *(*int64)(unsafe.Add(mBase, _consts[59]))
							v140 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v137 + (v139 - v140)
							v144 = *(*int64)(unsafe.Add(mBase, uint32(v85)+64))
							v146 = *(*int64)(unsafe.Add(mBase, _consts[60]))
							v147 = *(*int64)(unsafe.Add(mBase, uint32(v87)+64))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v144 + (v146 - v147)
							v151 = *(*int64)(unsafe.Add(mBase, uint32(v85)+72))
							v153 = *(*int64)(unsafe.Add(mBase, _consts[61]))
							v154 = *(*int64)(unsafe.Add(mBase, uint32(v87)+72))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v151 + (v153 - v154)
							v158 = *(*int64)(unsafe.Add(mBase, uint32(v85)+80))
							v160 = *(*int64)(unsafe.Add(mBase, _consts[62]))
							v161 = *(*int64)(unsafe.Add(mBase, uint32(v87)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+80)) = v158 + (v160 - v161)
							v165 = *(*int64)(unsafe.Add(mBase, uint32(v85)+88))
							v167 = *(*int64)(unsafe.Add(mBase, _consts[63]))
							v168 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+88)) = v165 + (v167 - v168)
							v172 = *(*int64)(unsafe.Add(mBase, uint32(v85)+96))
							v174 = *(*int64)(unsafe.Add(mBase, _consts[64]))
							v175 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+96)) = v172 + (v174 - v175)
							v179 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
							v181 = *(*int64)(unsafe.Add(mBase, _consts[65]))
							v182 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+104)) = v179 + (v181 - v182)
							v186 = *(*int64)(unsafe.Add(mBase, uint32(v85)+112))
							v188 = *(*int64)(unsafe.Add(mBase, _consts[66]))
							v189 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+112)) = v186 + (v188 - v189)
							v193 = *(*int64)(unsafe.Add(mBase, uint32(v85)+120))
							v195 = *(*int64)(unsafe.Add(mBase, _consts[67]))
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
						v90 = *(*int64)(unsafe.Add(mBase, _consts[52]))
						v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
						*(*int64)(unsafe.Add(mBase, uint32(v85))) = v88 + (v90 - v91)
						v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
						v97 = *(*int64)(unsafe.Add(mBase, _consts[53]))
						v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v95 + (v97 - v98)
						v102 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
						v104 = *(*int64)(unsafe.Add(mBase, _consts[54]))
						v105 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v102 + (v104 - v105)
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
						v111 = *(*int64)(unsafe.Add(mBase, _consts[55]))
						v112 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v109 + (v111 - v112)
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
						v118 = *(*int64)(unsafe.Add(mBase, _consts[56]))
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v116 + (v118 - v119)
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
						v125 = *(*int64)(unsafe.Add(mBase, _consts[57]))
						v126 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v123 + (v125 - v126)
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v85)+48))
						v132 = *(*int64)(unsafe.Add(mBase, _consts[58]))
						v133 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v130 + (v132 - v133)
						v137 = *(*int64)(unsafe.Add(mBase, uint32(v85)+56))
						v139 = *(*int64)(unsafe.Add(mBase, _consts[59]))
						v140 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v137 + (v139 - v140)
						v144 = *(*int64)(unsafe.Add(mBase, uint32(v85)+64))
						v146 = *(*int64)(unsafe.Add(mBase, _consts[60]))
						v147 = *(*int64)(unsafe.Add(mBase, uint32(v87)+64))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v144 + (v146 - v147)
						v151 = *(*int64)(unsafe.Add(mBase, uint32(v85)+72))
						v153 = *(*int64)(unsafe.Add(mBase, _consts[61]))
						v154 = *(*int64)(unsafe.Add(mBase, uint32(v87)+72))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v151 + (v153 - v154)
						v158 = *(*int64)(unsafe.Add(mBase, uint32(v85)+80))
						v160 = *(*int64)(unsafe.Add(mBase, _consts[62]))
						v161 = *(*int64)(unsafe.Add(mBase, uint32(v87)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+80)) = v158 + (v160 - v161)
						v165 = *(*int64)(unsafe.Add(mBase, uint32(v85)+88))
						v167 = *(*int64)(unsafe.Add(mBase, _consts[63]))
						v168 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+88)) = v165 + (v167 - v168)
						v172 = *(*int64)(unsafe.Add(mBase, uint32(v85)+96))
						v174 = *(*int64)(unsafe.Add(mBase, _consts[64]))
						v175 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+96)) = v172 + (v174 - v175)
						v179 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
						v181 = *(*int64)(unsafe.Add(mBase, _consts[65]))
						v182 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+104)) = v179 + (v181 - v182)
						v186 = *(*int64)(unsafe.Add(mBase, uint32(v85)+112))
						v188 = *(*int64)(unsafe.Add(mBase, _consts[66]))
						v189 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+112)) = v186 + (v188 - v189)
						v193 = *(*int64)(unsafe.Add(mBase, uint32(v85)+120))
						v195 = *(*int64)(unsafe.Add(mBase, _consts[67]))
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
			v41 = F__emscripten_memcpy_bulkmem(m, v15+int32(152), int32(4413976), int32(128))
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
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v33
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
						v90 = *(*int64)(unsafe.Add(mBase, _consts[52]))
						v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
						*(*int64)(unsafe.Add(mBase, uint32(v85))) = v88 + (v90 - v91)
						v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
						v97 = *(*int64)(unsafe.Add(mBase, _consts[53]))
						v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v95 + (v97 - v98)
						v102 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
						v104 = *(*int64)(unsafe.Add(mBase, _consts[54]))
						v105 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v102 + (v104 - v105)
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
						v111 = *(*int64)(unsafe.Add(mBase, _consts[55]))
						v112 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v109 + (v111 - v112)
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
						v118 = *(*int64)(unsafe.Add(mBase, _consts[56]))
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v116 + (v118 - v119)
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
						v125 = *(*int64)(unsafe.Add(mBase, _consts[57]))
						v126 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v123 + (v125 - v126)
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v85)+48))
						v132 = *(*int64)(unsafe.Add(mBase, _consts[58]))
						v133 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v130 + (v132 - v133)
						v137 = *(*int64)(unsafe.Add(mBase, uint32(v85)+56))
						v139 = *(*int64)(unsafe.Add(mBase, _consts[59]))
						v140 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v137 + (v139 - v140)
						v144 = *(*int64)(unsafe.Add(mBase, uint32(v85)+64))
						v146 = *(*int64)(unsafe.Add(mBase, _consts[60]))
						v147 = *(*int64)(unsafe.Add(mBase, uint32(v87)+64))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v144 + (v146 - v147)
						v151 = *(*int64)(unsafe.Add(mBase, uint32(v85)+72))
						v153 = *(*int64)(unsafe.Add(mBase, _consts[61]))
						v154 = *(*int64)(unsafe.Add(mBase, uint32(v87)+72))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v151 + (v153 - v154)
						v158 = *(*int64)(unsafe.Add(mBase, uint32(v85)+80))
						v160 = *(*int64)(unsafe.Add(mBase, _consts[62]))
						v161 = *(*int64)(unsafe.Add(mBase, uint32(v87)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+80)) = v158 + (v160 - v161)
						v165 = *(*int64)(unsafe.Add(mBase, uint32(v85)+88))
						v167 = *(*int64)(unsafe.Add(mBase, _consts[63]))
						v168 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+88)) = v165 + (v167 - v168)
						v172 = *(*int64)(unsafe.Add(mBase, uint32(v85)+96))
						v174 = *(*int64)(unsafe.Add(mBase, _consts[64]))
						v175 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+96)) = v172 + (v174 - v175)
						v179 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
						v181 = *(*int64)(unsafe.Add(mBase, _consts[65]))
						v182 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+104)) = v179 + (v181 - v182)
						v186 = *(*int64)(unsafe.Add(mBase, uint32(v85)+112))
						v188 = *(*int64)(unsafe.Add(mBase, _consts[66]))
						v189 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+112)) = v186 + (v188 - v189)
						v193 = *(*int64)(unsafe.Add(mBase, uint32(v85)+120))
						v195 = *(*int64)(unsafe.Add(mBase, _consts[67]))
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
					v90 = *(*int64)(unsafe.Add(mBase, _consts[52]))
					v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
					*(*int64)(unsafe.Add(mBase, uint32(v85))) = v88 + (v90 - v91)
					v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
					v97 = *(*int64)(unsafe.Add(mBase, _consts[53]))
					v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v95 + (v97 - v98)
					v102 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
					v104 = *(*int64)(unsafe.Add(mBase, _consts[54]))
					v105 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v102 + (v104 - v105)
					v109 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
					v111 = *(*int64)(unsafe.Add(mBase, _consts[55]))
					v112 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v109 + (v111 - v112)
					v116 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
					v118 = *(*int64)(unsafe.Add(mBase, _consts[56]))
					v119 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v116 + (v118 - v119)
					v123 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
					v125 = *(*int64)(unsafe.Add(mBase, _consts[57]))
					v126 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v123 + (v125 - v126)
					v130 = *(*int64)(unsafe.Add(mBase, uint32(v85)+48))
					v132 = *(*int64)(unsafe.Add(mBase, _consts[58]))
					v133 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v130 + (v132 - v133)
					v137 = *(*int64)(unsafe.Add(mBase, uint32(v85)+56))
					v139 = *(*int64)(unsafe.Add(mBase, _consts[59]))
					v140 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v137 + (v139 - v140)
					v144 = *(*int64)(unsafe.Add(mBase, uint32(v85)+64))
					v146 = *(*int64)(unsafe.Add(mBase, _consts[60]))
					v147 = *(*int64)(unsafe.Add(mBase, uint32(v87)+64))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v144 + (v146 - v147)
					v151 = *(*int64)(unsafe.Add(mBase, uint32(v85)+72))
					v153 = *(*int64)(unsafe.Add(mBase, _consts[61]))
					v154 = *(*int64)(unsafe.Add(mBase, uint32(v87)+72))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v151 + (v153 - v154)
					v158 = *(*int64)(unsafe.Add(mBase, uint32(v85)+80))
					v160 = *(*int64)(unsafe.Add(mBase, _consts[62]))
					v161 = *(*int64)(unsafe.Add(mBase, uint32(v87)+80))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+80)) = v158 + (v160 - v161)
					v165 = *(*int64)(unsafe.Add(mBase, uint32(v85)+88))
					v167 = *(*int64)(unsafe.Add(mBase, _consts[63]))
					v168 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+88)) = v165 + (v167 - v168)
					v172 = *(*int64)(unsafe.Add(mBase, uint32(v85)+96))
					v174 = *(*int64)(unsafe.Add(mBase, _consts[64]))
					v175 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+96)) = v172 + (v174 - v175)
					v179 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
					v181 = *(*int64)(unsafe.Add(mBase, _consts[65]))
					v182 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+104)) = v179 + (v181 - v182)
					v186 = *(*int64)(unsafe.Add(mBase, uint32(v85)+112))
					v188 = *(*int64)(unsafe.Add(mBase, _consts[66]))
					v189 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+112)) = v186 + (v188 - v189)
					v193 = *(*int64)(unsafe.Add(mBase, uint32(v85)+120))
					v195 = *(*int64)(unsafe.Add(mBase, _consts[67]))
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
	v70 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	v72 = *(*int32)(unsafe.Add(mBase, _consts[50]))
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
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v84
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
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v85
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v86
	v588 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[685])) = v588
	*(*int64)(unsafe.Add(mBase, _consts[686])) = v588
	v594 = int32(0)
	*(*uint16)(unsafe.Add(mBase, _consts[379])) = uint16(v594)
	v597 = *(*int32)(unsafe.Add(mBase, _consts[676]))
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597)+29)))
	if v598 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
	v114 = *(*int32)(unsafe.Add(mBase, _consts[689]))
	v119 = F_AllocSetContextCreateInternal(m, v114, int32(62476), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[692])) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
	v125 = *(*int32)(unsafe.Add(mBase, _consts[689]))
	v130 = F_AllocSetContextCreateInternal(m, v125, int32(62339), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[698])) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
	v135 = int32(0)
	F_pgstat_report_activity(m, int32(2), v135)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = int32(987)
	v140 = int32(4508616)
	v141 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v141
	*(*int32)(unsafe.Add(mBase, _consts[703])) = v80
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v80
	v148 = v28
	v152 = v135
	v163 = l0
	v166 = v102 + v101*int64(1000000) - int64(946684800000000)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v173 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, _consts[703])) = v564
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v564
	v569 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	v573 = *(*int32)(unsafe.Add(mBase, _consts[683]))
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
	v179 = *(*int32)(unsafe.Add(mBase, _consts[692]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v179
	v182 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v148
	v186 = *(*int32)(unsafe.Add(mBase, _consts[683]))
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
	v214 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	v242 = *(*int32)(unsafe.Add(mBase, _consts[346]))
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
	F_errmsg(m, int32(462071), int32(0))
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
	F_errfinish(m, int32(495845), int32(3639), int32(234866))
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
	*(*int32)(unsafe.Add(mBase, _consts[346])) = int32(0)
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
	v271 = *(*int32)(unsafe.Add(mBase, _consts[692]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v271
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
	v377 = *(*int32)(unsafe.Add(mBase, _consts[692]))
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
	v293 = *(*int32)(unsafe.Add(mBase, _consts[589]))
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
	v320 = *(*int32)(unsafe.Add(mBase, _consts[589]))
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
	v343 = *(*int32)(unsafe.Add(mBase, _consts[589]))
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
	v366 = *(*int32)(unsafe.Add(mBase, _consts[589]))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+104)) = v329
	*(*int64)(unsafe.Add(mBase, uint32(v366)+96)) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v366)+88)) = v356 + v355*int64(1000000) - int64(946684800000000)
	v371 = v336
	goto L45
L68:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	v386 = *(*int32)(unsafe.Add(mBase, _consts[683]))
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
	v416 = int32(*(*uint8)(unsafe.Add(mBase, _consts[688])))
	if v416 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	v430 = *(*int32)(unsafe.Add(mBase, _consts[692]))
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
	v418 = int32(*(*uint8)(unsafe.Add(mBase, _consts[694])))
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
	v435 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v435
	if v395 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v390
	v442 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	v445 = *(*int32)(unsafe.Add(mBase, _consts[704]))
	v447 = *(*int32)(unsafe.Add(mBase, _consts[690]))
	if v447 == int32(4121768) {
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
	v473 = *(*int32)(unsafe.Add(mBase, _consts[346]))
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
	v462 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = int32(0)
	goto L90
L90:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	*(*int32)(unsafe.Add(mBase, _consts[346])) = int32(0)
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
	v487 = *(*int32)(unsafe.Add(mBase, _consts[705]))
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
	v510 = *(*int32)(unsafe.Add(mBase, _consts[705]))
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
	F_errmsg(m, int32(65835), int32(0))
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
	F_errfinish(m, int32(495845), int32(3793), int32(234866))
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
	v555 = *(*int32)(unsafe.Add(mBase, _consts[72]))
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
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v85
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v86
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
	v609 = *(*int32)(unsafe.Add(mBase, _consts[676]))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v609)))
	v612 = *(*int32)(unsafe.Add(mBase, _consts[589]))
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
	v173 = *(*int32)(unsafe.Add(mBase, _consts[442]))
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
	F_errcode(m, int32(290948))
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
	F_errmsg(m, int32(692996), v15)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(494685), int32(1618), int32(161683))
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
	var v55 int32
	_ = v55
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
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v650 int32
	_ = v650
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
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
		goto L17
	} else {
		goto L18
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L175
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L171
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L168
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L165
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L162
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L159
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L156
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L153
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L149
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L146
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L143
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L140
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L137
	}
L16:
	;
	v650 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v634))) = uint8(v650)
	m.G0 = v25 + int32(32)
	return v640
L17:
	;
	v634 = v30
	v640 = v30
	goto L16
L18:
	;
	goto L19
L19:
	;
	v37 = l2 - l0
	v41 = l0
	v43 = int32(0)
	v45 = v34
	v47 = v30
	v53 = v30
	v55 = v29
	goto L20
L20:
	;
	v63 = v47 - v53
	if base.Ui32(v55-int32(17)) < base.Ui32(v63) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v620 != 0 {
		v818 = v626
		goto L3
	} else {
		goto L136
	}
L22:
	;
	v68 = v55 << (uint(int32(1)) % 32)
	v69 = F_repalloc(m, v53, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v73 = v45
	v74 = v47
	v75 = v53
	v76 = v55
	goto L24
L24:
	;
	v77 = int32(255)
	v78 = v2 & v77
	if v78 == v73&v77 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v73 = v72
	v74 = v69 + v63
	v75 = v69
	v76 = v68
	goto L24
L26:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(12))+8))
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v624
	goto L134
L27:
	;
	v608 = v594
	v609 = v605
	v620 = int32(0)
	goto L26
L28:
	;
	F_pg_unicode_to_server(m, v279, v74)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L133
	}
L29:
	;
	v83 = v25 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = int32(499)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v41 + (v37 + int32(3))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v83
	v90 = int32(4508616)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v91
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v25 + int32(20)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v78 == v97 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if v43 != 0 {
		v818 = v41
		goto L3
	} else {
		goto L131
	}
L32:
	;
	if v43 != 0 {
		v818 = v41
		goto L3
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L37
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v2)
	v608 = int32(2)
	v609 = v74 + int32(1)
	v620 = int32(0)
	goto L26
L36:
	;
	if v97 != int32(43) {
		goto L73
	} else {
		goto L74
	}
L37:
	;
	if base.B2i32(base.Ui32(v97-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v97|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)))
	goto L39
L39:
	;
	if base.B2i32(base.Ui32(v117-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v117|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+3)))
	goto L41
L41:
	;
	if base.B2i32(base.Ui32(v131-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v131|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	goto L43
L43:
	;
	if base.B2i32(base.Ui32(v145-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v145|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v159 = int32(-48)
	if base.Ui32((v97-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v181 = v159
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if base.Ui32((v117-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v202 = v159
		goto L49
	} else {
		goto L50
	}
L46:
	;
	if base.Ui32((v97-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v181 = int32(-87)
		goto L45
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v97-int32(65))&int32(255)) {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	v181 = int32(-55)
	goto L45
L49:
	;
	v203 = int32(-48)
	if base.Ui32((v131-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v225 = v203
		goto L55
	} else {
		goto L56
	}
L50:
	;
	if base.Ui32((v117-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v202 = int32(-87)
	goto L49
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v117-int32(65))&int32(255)) {
		goto L14
	} else {
		goto L54
	}
L54:
	;
	v202 = int32(-55)
	goto L49
L55:
	;
	if base.Ui32((v145-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v246 = v203
		goto L59
	} else {
		goto L60
	}
L56:
	;
	if base.Ui32((v131-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v225 = int32(-87)
		goto L55
	} else {
		goto L57
	}
L57:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v131-int32(65))&int32(255)) {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	v225 = int32(-55)
	goto L55
L59:
	;
	v259 = v145 + v246 + ((v117+v202)<<(uint(int32(8))%32) + (v97+v181)<<(uint(int32(12))%32) + (v131+v225)<<(uint(int32(4))%32))
	if base.Ui32(int32(1114111)) <= base.Ui32(v259-int32(1)) {
		goto L11
	} else {
		goto L65
	}
L60:
	;
	if base.Ui32((v145-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v246 = int32(-87)
	goto L59
L62:
	;
	goto L63
L63:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v145-int32(65))&int32(255)) {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	v246 = int32(-55)
	goto L59
L65:
	;
	v265 = v259 & int32(2096128)
	if v43 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v280 = int32(5)
	if v279&int32(16776192) != int32(55296) {
		goto L28
	} else {
		goto L72
	}
L67:
	;
	if v265 != int32(56320) {
		v818 = v41
		goto L3
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v265 == int32(56320) {
		v818 = v41
		goto L3
	} else {
		goto L71
	}
L70:
	;
	v279 = v43<<(uint(int32(10))%32)&int32(1047552) | v259&int32(1023) + int32(65536)
	goto L66
L71:
	;
	v279 = v259
	goto L66
L72:
	;
	v608 = v280
	v609 = v74
	v620 = v279
	goto L26
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L126
	}
L74:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)))
	goto L75
L75:
	;
	if base.B2i32(base.Ui32(v290-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v290|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+3)))
	goto L77
L77:
	;
	if base.B2i32(base.Ui32(v304-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v304|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	goto L79
L79:
	;
	if base.B2i32(base.Ui32(v318-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v318|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L73
	} else {
		goto L80
	}
L80:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+5)))
	goto L81
L81:
	;
	if base.B2i32(base.Ui32(v332-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v332|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L73
	} else {
		goto L82
	}
L82:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+6)))
	goto L83
L83:
	;
	if base.B2i32(base.Ui32(v346-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v346|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L73
	} else {
		goto L84
	}
L84:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
	goto L85
L85:
	;
	if base.B2i32(base.Ui32(v360-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v360|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L73
	} else {
		goto L86
	}
L86:
	;
	v374 = int32(-48)
	if base.Ui32((v290-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v396 = v374
		goto L87
	} else {
		goto L88
	}
L87:
	;
	if base.Ui32((v304-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v417 = v374
		goto L91
	} else {
		goto L92
	}
L88:
	;
	if base.Ui32((v290-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v396 = int32(-87)
		goto L87
	} else {
		goto L89
	}
L89:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v290-int32(65))&int32(255)) {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	v396 = int32(-55)
	goto L87
L91:
	;
	v418 = int32(-48)
	if base.Ui32((v318-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v440 = v418
		goto L97
	} else {
		goto L98
	}
L92:
	;
	if base.Ui32((v304-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v417 = int32(-87)
	goto L91
L94:
	;
	goto L95
L95:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v304-int32(65))&int32(255)) {
		goto L9
	} else {
		goto L96
	}
L96:
	;
	v417 = int32(-55)
	goto L91
L97:
	;
	if base.Ui32((v332-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v461 = v418
		goto L101
	} else {
		goto L102
	}
L98:
	;
	if base.Ui32((v318-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v440 = int32(-87)
		goto L97
	} else {
		goto L99
	}
L99:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v318-int32(65))&int32(255)) {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	v440 = int32(-55)
	goto L97
L101:
	;
	v462 = int32(-48)
	if base.Ui32((v346-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v484 = v462
		goto L107
	} else {
		goto L108
	}
L102:
	;
	if base.Ui32((v332-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v461 = int32(-87)
	goto L101
L104:
	;
	goto L105
L105:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v332-int32(65))&int32(255)) {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	v461 = int32(-55)
	goto L101
L107:
	;
	if base.Ui32((v360-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v505 = v462
		goto L111
	} else {
		goto L112
	}
L108:
	;
	if base.Ui32((v346-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v484 = int32(-87)
		goto L107
	} else {
		goto L109
	}
L109:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v346-int32(65))&int32(255)) {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	v484 = int32(-55)
	goto L107
L111:
	;
	v526 = v360 + v505 + ((v304+v417)<<(uint(int32(16))%32) + (v290+v396)<<(uint(int32(20))%32) + (v318+v440)<<(uint(int32(12))%32) + (v332+v461)<<(uint(int32(8))%32) + (v346+v484)<<(uint(int32(4))%32))
	if base.Ui32(int32(1114111)) <= base.Ui32(v526-int32(1)) {
		goto L4
	} else {
		goto L117
	}
L112:
	;
	if base.Ui32((v360-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v505 = int32(-87)
	goto L111
L114:
	;
	goto L115
L115:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v360-int32(65))&int32(255)) {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v505 = int32(-55)
	goto L111
L117:
	;
	v532 = v526 & int32(2096128)
	if v43 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v547 = int32(8)
	if v546&int32(16776192) == int32(55296) {
		v608 = v547
		v609 = v74
		v620 = v546
		goto L26
	} else {
		goto L124
	}
L119:
	;
	if v532 != int32(56320) {
		v818 = v41
		goto L3
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	if v532 == int32(56320) {
		v818 = v41
		goto L3
	} else {
		goto L123
	}
L122:
	;
	v546 = v43<<(uint(int32(10))%32)&int32(1047552) | v526&int32(1023) + int32(65536)
	goto L118
L123:
	;
	v546 = v526
	goto L118
L124:
	;
	F_pg_unicode_to_server(m, v546, v74)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v554 = F_strlen(m, v74)
	mBase = m.M
	v594 = v547
	v605 = v554 + v74
	goto L27
L126:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errmsg(m, int32(372395), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errhint(m, int32(656605), int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(495690), int32(495), int32(372312))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
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
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v73)
	v583 = int32(1)
	v584 = v74 + v583
	v587 = v41 + v583
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	if v588 != 0 {
		v41 = v587
		v43 = int32(0)
		v45 = v588
		v47 = v584
		v53 = v75
		v55 = v76
		goto L20
	} else {
		goto L132
	}
L132:
	;
	v634 = v584
	v640 = v75
	goto L16
L133:
	;
	v591 = F_strlen(m, v74)
	mBase = m.M
	v594 = v280
	v605 = v591 + v74
	goto L27
L134:
	;
	v626 = v41 + v608
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	if v627 != 0 {
		v41 = v626
		v43 = v620
		v45 = v627
		v47 = v609
		v53 = v75
		v55 = v76
		goto L20
	} else {
		goto L135
	}
L135:
	;
	goto L21
L136:
	;
	v634 = v609
	v640 = v75
	goto L16
L137:
	;
	F_errmsg_internal(m, int32(102919), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(495690), int32(336), int32(308946))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errmsg_internal(m, int32(102919), int32(0))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(495690), int32(336), int32(308946))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errmsg_internal(m, int32(102919), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(495690), int32(336), int32(308946))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errmsg_internal(m, int32(102919), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(495690), int32(336), int32(308946))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(346448), int32(0))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(495690), int32(347), int32(345641))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errmsg_internal(m, int32(102919), int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(495690), int32(336), int32(308946))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errmsg_internal(m, int32(102919), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(495690), int32(336), int32(308946))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errmsg_internal(m, int32(102919), int32(0))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(495690), int32(336), int32(308946))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	F_errmsg_internal(m, int32(102919), int32(0))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(495690), int32(336), int32(308946))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(102919), int32(0))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(495690), int32(336), int32(308946))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	F_errmsg_internal(m, int32(102919), int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(495690), int32(336), int32(308946))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errmsg(m, int32(346448), int32(0))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(495690), int32(347), int32(345641))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_errmsg(m, int32(213426), int32(0))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_scanner_errposition(m, v818+v37+int32(3), l3)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(495690), int32(525), int32(372312))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
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
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v48 = v40 + v44
	v52 = int32(-2139062144)
	if (v46|(int32(16843008)-v46))&v52 == v52 {
		v40 = v48
		v41 = v46
		v42 = v45
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v57 = v48
	v58 = v46
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
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5<<(uint(int32(1))%32))+uint32(_consts[1383]))))
	return v10 + int32(4101764)
}
func F_string2ean(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v161 int32
	_ = v161
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
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v368 int32
	_ = v368
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
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
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v775 int32
	_ = v775
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
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v900 int32
	_ = v900
	var v907 int32
	_ = v907
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v932 int32
	_ = v932
	var v943 int32
	_ = v943
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1016 int32
	_ = v1016
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1061 int64
	_ = v1061
	var v1074 int64
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1099 int64
	_ = v1099
	var v1109 int32
	_ = v1109
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1130 int64
	_ = v1130
	var v1143 int64
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int64
	_ = v1147
	var v1171 int64
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1183 int32
	_ = v1183
	var v1191 int32
	_ = v1191
	var v1198 int32
	_ = v1198
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1217 int32
	_ = v1217
	var v1225 int32
	_ = v1225
	var v1232 int32
	_ = v1232
	var v1239 int32
	_ = v1239
	var v1245 int32
	_ = v1245
	v5 = int32(0)
	v19 = int64(0)
	v20 = m.G0
	v22 = v20 - int32(112)
	m.G0 = v22
	v24 = int32(4078816)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1352])))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+96)) = uint8(v25)
	v27 = *(*int64)(unsafe.Add(mBase, _consts[1353]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v27
	v29 = *(*int64)(unsafe.Add(mBase, _consts[1354]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v31 == v5 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v22 + int32(112)
	return v1245
L2:
	;
	v932 = v22 + int32(80)
	goto L272
L3:
	;
	v920 = int32(0)
	v921 = int32(-1)
	goto L2
L4:
	;
	v885 = int32(0)
	v886 = F_errsave_start(m, l1)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L186
	} else {
		goto L267
	}
L5:
	;
	v857 = int32(0)
	v858 = F_errsave_start(m, l1)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L186
	} else {
		goto L262
	}
L6:
	;
	v38 = v22 + int32(80) | int32(3)
	v43 = l0
	v46 = v38
	v47 = v31
	v48 = v5
	v50 = int32(1)
	v51 = v5
	v53 = v5
	goto L7
L7:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	v61 = int32(255)
	v66 = v47 & v61
	if v66 != int32(63) {
		v74 = v53
		v75 = base.B2i32(base.Ui32((v47-int32(48))&v61) < base.Ui32(int32(10)))
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v185 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v185)
	switch v174 - int32(8) {
	case 0:
		goto L80
	default:
		goto L5
	case 2:
		goto L81
	case 4:
		goto L82
	case 5:
		goto L83
	}
L9:
	;
	switch v51 {
	case 0:
		goto L23
	default:
		goto L19
	case 7:
		goto L22
	case 9:
		goto L21
	case 11:
		goto L20
	}
L10:
	;
	if v58 != int32(33) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v58 != 0 {
		v74 = v53
		v75 = int32(0)
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v72 = int32(1)
	v74 = v72
	v75 = v72
	goto L9
L14:
	;
	goto L13
L15:
	;
	v180 = v43 + int32(1)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v181 != 0 {
		goto L68
	} else {
		goto L69
	}
L16:
	;
	if v75 == int32(0) {
		goto L5
	} else {
		goto L66
	}
L17:
	;
	if v75 == int32(0) {
		goto L5
	} else {
		goto L65
	}
L18:
	;
	if v58 != 0 {
		goto L16
	} else {
		goto L64
	}
L19:
	;
	switch v66 - int32(32) {
	case 0, 13:
		v174 = v51
		v175 = v46
		v176 = v48
		v177 = v50
		v178 = v74
		goto L15
	case 1:
		goto L18
	default:
		goto L16
	}
L20:
	;
	if v75 == int32(0) {
		goto L19
	} else {
		goto L58
	}
L21:
	;
	if v75 != 0 {
		goto L41
	} else {
		goto L42
	}
L22:
	;
	if v75 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	switch v66 - int32(32) {
	case 0, 13:
		v174 = int32(0)
		v175 = v46
		v176 = v48
		v177 = v50
		v178 = v74
		goto L15
	case 1:
		goto L18
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		goto L17
	default:
		goto L24
	}
L24:
	;
	if base.B2i32(v66 != int32(109))&base.B2i32(v66 != int32(77)) != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	if v48 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v84 = int32(77)
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v84)
	v86 = int32(1)
	v174 = v86
	v175 = v46 + v86
	v176 = int32(4)
	v177 = v50
	v178 = v74
	goto L15
L27:
	;
	if v58 != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	switch v66 - int32(32) {
	case 0, 13:
		v174 = int32(7)
		v175 = v46
		v176 = v48
		v177 = v50
		v178 = v74
		goto L15
	case 1:
		goto L18
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		goto L5
	default:
		goto L29
	}
L29:
	;
	if v66 == int32(120) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	if v66 != int32(88) {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	v101 = base.B2i32(v58 != int32(33))
	goto L34
L33:
	;
	v101 = int32(0)
	goto L34
L34:
	;
	if v101 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	if v48 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v66-int32(97)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v108)
	v174 = int32(8)
	v175 = v46 + int32(1)
	v176 = int32(5)
	v177 = v50
	v178 = v74
	goto L15
L38:
	;
	v108 = v66 & int32(95)
	goto L40
L39:
	;
	v108 = v66
	goto L40
L40:
	;
	goto L37
L41:
	;
	if v58 != 0 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	switch v66 - int32(32) {
	case 0, 13:
		v174 = int32(9)
		v175 = v46
		v176 = v48
		v177 = v50
		v178 = v74
		goto L15
	case 1:
		goto L18
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		goto L5
	default:
		goto L43
	}
L43:
	;
	if v66 == int32(120) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	if v66 != int32(88) {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	v125 = base.B2i32(v58 != int32(33))
	goto L48
L47:
	;
	v125 = int32(0)
	goto L48
L48:
	;
	if v125 != 0 {
		goto L19
	} else {
		goto L49
	}
L49:
	;
	if v48&int32(-5) != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	if base.Ui32(v66-int32(97)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v134)
	if v48 != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v134 = v66 & int32(95)
	goto L54
L53:
	;
	v134 = v66
	goto L54
L54:
	;
	goto L51
L55:
	;
	v139 = v48
	goto L57
L56:
	;
	v139 = int32(3)
	goto L57
L57:
	;
	v174 = int32(10)
	v175 = v46 + int32(1)
	v176 = v139
	v177 = v50
	v178 = v74
	goto L15
L58:
	;
	if v58 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v146 = base.B2i32(v58 != int32(33))
	goto L61
L60:
	;
	v146 = int32(0)
	goto L61
L61:
	;
	if v146 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	if v48 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v47)
	v174 = int32(12)
	v175 = v46 + int32(1)
	v176 = int32(6)
	v177 = v50
	v178 = v74
	goto L15
L64:
	;
	v174 = v51
	v175 = v46
	v176 = v48
	v177 = v50 & v74
	v178 = int32(1)
	goto L15
L65:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v47)
	v161 = int32(1)
	v174 = v161
	v175 = v46 + v161
	v176 = v48
	v177 = v50
	v178 = v74
	goto L15
L66:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v47)
	if base.Ui32(int32(12)) < base.Ui32(v51) {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v170 = int32(1)
	v174 = v51 + v170
	v175 = v46 + v170
	v176 = v48
	v177 = v50
	v178 = v74
	goto L15
L68:
	;
	if base.Ui32(v174) < base.Ui32(int32(14)) {
		v43 = v180
		v46 = v175
		v47 = v181
		v48 = v176
		v50 = v177
		v51 = v174
		v53 = v178
		goto L7
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	goto L8
L71:
	;
	goto L70
L72:
	;
	v757 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+82)) = uint8(v757)
	if v177&int32(1) == int32(0) {
		goto L3
	} else {
		goto L240
	}
L73:
	;
	v694 = int32(12336)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+90)) = uint16(v694)
	v696 = int32(556810)
	v697 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1355])))
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+80)) = uint16(v697)
	v699 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1356])))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+82)) = uint8(v699)
	if v177&int32(1) == int32(0) {
		goto L3
	} else {
		goto L227
	}
L74:
	;
	v633 = int32(556533)
	v634 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1357])))
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+80)) = uint16(v634)
	v636 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1358])))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+82)) = uint8(v636)
	if v177&int32(1) == int32(0) {
		goto L3
	} else {
		goto L214
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = int32(809056057)
	if v177&int32(1) == int32(0) {
		goto L3
	} else {
		goto L192
	}
L76:
	;
	v517 = int32(0)
	v518 = F_errsave_start(m, l1)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L186
	} else {
		goto L187
	}
L77:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+83)))
	if v300 == int32(48) {
		goto L117
	} else {
		goto L118
	}
L78:
	;
	v224 = int32(0)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v234 = base.B2i32(v232 == int32(77))
	if v232 == int32(77) {
		goto L96
	} else {
		goto L97
	}
L79:
	;
	if l3 == int32(2) {
		v513 = v216
		goto L76
	} else {
		goto L93
	}
L80:
	;
	switch v176 {
	case 0, 5:
		goto L91
	default:
		goto L5
	}
L81:
	;
	if base.Ui32(v176-int32(5)) < base.Ui32(int32(-2)) {
		goto L5
	} else {
		goto L87
	}
L82:
	;
	v193 = int32(6)
	if v176 != v193 {
		goto L5
	} else {
		goto L86
	}
L83:
	;
	if v176 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	if v177&int32(1) != 0 {
		goto L78
	} else {
		goto L85
	}
L85:
	;
	v298 = int32(-1)
	v299 = int32(0)
	goto L77
L86:
	;
	v196 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+94)))
	v216 = v193
	v218 = v196 - int32(48)
	goto L79
L87:
	;
	v203 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+92)))
	if v203 == int32(88) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v216 = v176
	v218 = int32(10)
	goto L79
L89:
	;
	goto L90
L90:
	;
	v216 = v176
	v218 = v203 - int32(48)
	goto L79
L91:
	;
	v209 = int32(5)
	v211 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+90)))
	if v211 == int32(88) {
		v216 = v209
		v218 = int32(10)
		goto L79
	} else {
		goto L92
	}
L92:
	;
	v216 = v209
	v218 = v211 - int32(48)
	goto L79
L93:
	;
	if l3 != v216 {
		v513 = v216
		goto L76
	} else {
		goto L94
	}
L94:
	;
	switch l3 - int32(4) {
	case 0:
		goto L75
	case 1:
		goto L73
	case 2:
		goto L72
	default:
		goto L74
	}
L95:
	;
	v293 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+95)))
	v298 = v292
	v299 = base.B2i32(v292 == v293-int32(48)) | v178
	goto L77
L96:
	;
	v235 = int32(3)
	goto L98
L97:
	;
	v235 = v224
	goto L98
L98:
	;
	if v232 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v284 = int32(10)
	v289 = base.I32_rem_u_s(v281*int32(3)+v280, v284)
	if v289 != 0 {
		goto L113
	} else {
		goto L114
	}
L100:
	;
	v280 = v224
	v281 = v235
	goto L99
L101:
	;
	goto L102
L102:
	;
	v239 = v38
	v240 = v232
	v241 = v234
	v242 = v224
	v243 = v235
	v244 = int32(13)
	goto L103
L103:
	;
	v249 = (v240 - int32(48)) & int32(255)
	if base.Ui32(v249) <= base.Ui32(int32(9)) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v280 = v266
	v281 = v267
	goto L99
L105:
	;
	v254 = v241 & int32(1)
	if v254 != 0 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v265 = v241
	v266 = v242
	v267 = v243
	v268 = v244
	goto L107
L107:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+1)))
	if v270 == int32(0) {
		v280 = v266
		v281 = v267
		goto L99
	} else {
		goto L111
	}
L108:
	;
	v255 = int32(0)
	goto L110
L109:
	;
	v255 = v249
	goto L110
L110:
	;
	v261 = int32(1)
	v265 = v241 + v261
	v266 = v255 + v242
	v267 = (int32(0)-v254)&v249 + v243
	v268 = v244 - v261
	goto L107
L111:
	;
	v273 = int32(1)
	if base.Ui32(v273) < base.Ui32(v268) {
		v239 = v239 + v273
		v240 = v270
		v241 = v265
		v242 = v266
		v243 = v267
		v244 = v268
		goto L103
	} else {
		goto L112
	}
L112:
	;
	goto L104
L113:
	;
	v292 = v284 - v289
	goto L115
L114:
	;
	v292 = int32(0)
	goto L115
L115:
	;
	goto L95
L116:
	;
	if l3 == int32(2) {
		v920 = v299
		v921 = v298
		goto L2
	} else {
		goto L184
	}
L117:
	;
	v509 = int32(6)
	goto L116
L118:
	;
	goto L119
L119:
	;
	v304 = int32(3)
	goto L124
L120:
	;
	if v368 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L121:
	;
	v368 = int32(0)
	goto L120
L122:
	;
	v342 = int32(556810)
	v343 = v38
	v344 = v304
	goto L132
L124:
	;
	goto L125
L125:
	;
	goto L131
L131:
	;
	goto L122
L132:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343))))
	if v347 == v348 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v368 = v347 - v348
	goto L120
L134:
	;
	v350 = int32(1)
	v355 = v344 - v350
	if v355 != 0 {
		v342 = v342 + v350
		v343 = v343 + v350
		v344 = v355
		goto L132
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	goto L133
L137:
	;
	goto L121
L138:
	;
	v509 = int32(5)
	goto L116
L139:
	;
	goto L140
L140:
	;
	goto L145
L141:
	;
	if v435 == int32(0) {
		v509 = v304
		goto L116
	} else {
		goto L159
	}
L142:
	;
	v435 = int32(0)
	goto L141
L143:
	;
	v409 = int32(556533)
	v410 = v38
	v411 = int32(3)
	goto L153
L145:
	;
	goto L146
L146:
	;
	goto L152
L152:
	;
	goto L143
L153:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	if v414 == v415 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v435 = v414 - v415
	goto L141
L155:
	;
	v417 = int32(1)
	v422 = v411 - v417
	if v422 != 0 {
		v409 = v409 + v417
		v410 = v410 + v417
		v411 = v422
		goto L153
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	goto L154
L158:
	;
	goto L142
L159:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v438 == int32(809056057) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v509 = int32(4)
	goto L116
L161:
	;
	goto L162
L162:
	;
	v443 = int32(3)
	goto L167
L163:
	;
	if v507 != 0 {
		goto L181
	} else {
		goto L182
	}
L164:
	;
	v507 = int32(0)
	goto L163
L165:
	;
	v481 = int32(552687)
	v482 = v38
	v483 = v443
	goto L175
L167:
	;
	goto L168
L168:
	;
	goto L174
L174:
	;
	goto L165
L175:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481))))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	if v486 == v487 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v507 = v486 - v487
	goto L163
L177:
	;
	v489 = int32(1)
	v494 = v483 - v489
	if v494 != 0 {
		v481 = v481 + v489
		v482 = v482 + v489
		v483 = v494
		goto L175
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	goto L176
L180:
	;
	goto L164
L181:
	;
	v508 = int32(2)
	goto L183
L182:
	;
	v508 = v443
	goto L183
L183:
	;
	v509 = v508
	goto L116
L184:
	;
	if l3 == v509 {
		v920 = v299
		v921 = v298
		goto L2
	} else {
		goto L185
	}
L185:
	;
	v513 = v509
	goto L76
L186:
	;
	return int32(0)
L187:
	;
	if v518 == int32(0) {
		v1245 = v517
		goto L1
	} else {
		goto L188
	}
L188:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L186
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = l0
	v531 = int32(2)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(v531)%32))+uint32(_consts[1359])))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v534
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v513<<(uint(v531)%32))+uint32(_consts[1359])))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v539
	F_errmsg(m, int32(726442), v22-int32(-64))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L186
	} else {
		goto L190
	}
L190:
	;
	F_errsave_finish(m, l1, int32(496621), int32(908), int32(284464))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L186
	} else {
		goto L191
	}
L191:
	;
	v1245 = v517
	goto L1
L192:
	;
	v561 = v22 + int32(80)
	v562 = int32(0)
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561))))
	v572 = base.B2i32(v570 == int32(77))
	if v570 == int32(77) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v920 = base.B2i32(v630 == v218) | v178
	v921 = v630
	goto L2
L194:
	;
	v573 = int32(3)
	goto L196
L195:
	;
	v573 = v562
	goto L196
L196:
	;
	if v570 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v622 = int32(10)
	v627 = base.I32_rem_u_s(v619*int32(3)+v618, v622)
	if v627 != 0 {
		goto L211
	} else {
		goto L212
	}
L198:
	;
	v618 = v562
	v619 = v573
	goto L197
L199:
	;
	goto L200
L200:
	;
	v577 = v561
	v578 = v570
	v579 = v572
	v580 = v562
	v581 = v573
	v582 = int32(13)
	goto L201
L201:
	;
	v587 = (v578 - int32(48)) & int32(255)
	if base.Ui32(v587) <= base.Ui32(int32(9)) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v618 = v604
	v619 = v605
	goto L197
L203:
	;
	v592 = v579 & int32(1)
	if v592 != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	v603 = v579
	v604 = v580
	v605 = v581
	v606 = v582
	goto L205
L205:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577)+1)))
	if v608 == int32(0) {
		v618 = v604
		v619 = v605
		goto L197
	} else {
		goto L209
	}
L206:
	;
	v593 = int32(0)
	goto L208
L207:
	;
	v593 = v587
	goto L208
L208:
	;
	v599 = int32(1)
	v603 = v579 + v599
	v604 = v593 + v580
	v605 = (int32(0)-v592)&v587 + v581
	v606 = v582 - v599
	goto L205
L209:
	;
	v611 = int32(1)
	if base.Ui32(v611) < base.Ui32(v606) {
		v577 = v577 + v611
		v578 = v608
		v579 = v603
		v580 = v604
		v581 = v605
		v582 = v606
		goto L201
	} else {
		goto L210
	}
L210:
	;
	goto L202
L211:
	;
	v630 = v622 - v627
	goto L213
L212:
	;
	v630 = int32(0)
	goto L213
L213:
	;
	goto L193
L214:
	;
	v643 = int32(0)
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v646 == v643 {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v920 = base.B2i32(v691 == v218) | v178
	v921 = v691
	goto L2
L216:
	;
	v691 = int32(0)
	goto L215
L217:
	;
	v649 = v38
	v650 = int32(10)
	v651 = v646
	v652 = v643
	goto L218
L218:
	;
	v657 = (v651 - int32(48)) & int32(255)
	v661 = base.B2i32(base.Ui32(v657) < base.Ui32(int32(10)))
	if base.Ui32(v657) < base.Ui32(int32(10)) {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	v675 = base.I32_rem_u_s(v663, int32(11))
	if v675 == int32(0) {
		goto L216
	} else {
		goto L226
	}
L220:
	;
	goto L219
L221:
	;
	v662 = v650 * v657
	goto L223
L222:
	;
	v662 = int32(0)
	goto L223
L223:
	;
	v663 = v662 + v652
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+1)))
	if v664 == int32(0) {
		goto L220
	} else {
		goto L224
	}
L224:
	;
	v667 = int32(1)
	v669 = v650 - v661
	if base.Ui32(v667) < base.Ui32(v669) {
		v649 = v649 + v667
		v650 = v669
		v651 = v664
		v652 = v663
		goto L218
	} else {
		goto L225
	}
L225:
	;
	goto L220
L226:
	;
	v691 = int32(11) - v675
	goto L215
L227:
	;
	v706 = int32(0)
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v709 == v706 {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	v920 = base.B2i32(v754 == v218) | v178
	v921 = v754
	goto L2
L229:
	;
	v754 = int32(0)
	goto L228
L230:
	;
	v712 = v38
	v713 = int32(8)
	v714 = v709
	v715 = v706
	goto L231
L231:
	;
	v720 = (v714 - int32(48)) & int32(255)
	v724 = base.B2i32(base.Ui32(v720) < base.Ui32(int32(10)))
	if base.Ui32(v720) < base.Ui32(int32(10)) {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	v738 = base.I32_rem_u_s(v726, int32(11))
	if v738 == int32(0) {
		goto L229
	} else {
		goto L239
	}
L233:
	;
	goto L232
L234:
	;
	v725 = v713 * v720
	goto L236
L235:
	;
	v725 = int32(0)
	goto L236
L236:
	;
	v726 = v725 + v715
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712)+1)))
	if v727 == int32(0) {
		goto L233
	} else {
		goto L237
	}
L237:
	;
	v730 = int32(1)
	v732 = v713 - v724
	if base.Ui32(v730) < base.Ui32(v732) {
		v712 = v712 + v730
		v713 = v732
		v714 = v727
		v715 = v726
		goto L231
	} else {
		goto L238
	}
L238:
	;
	goto L233
L239:
	;
	v754 = int32(11) - v738
	goto L228
L240:
	;
	v766 = v22 + int32(80) | int32(2)
	v767 = int32(0)
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766))))
	v777 = base.B2i32(v775 == int32(77))
	if v775 == int32(77) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v920 = base.B2i32(v835 == v218) | v178
	v921 = v835
	goto L2
L242:
	;
	v778 = int32(3)
	goto L244
L243:
	;
	v778 = v767
	goto L244
L244:
	;
	if v775 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v827 = int32(10)
	v832 = base.I32_rem_u_s(v824*int32(3)+v823, v827)
	if v832 != 0 {
		goto L259
	} else {
		goto L260
	}
L246:
	;
	v823 = v767
	v824 = v778
	goto L245
L247:
	;
	goto L248
L248:
	;
	v782 = v766
	v783 = v775
	v784 = v777
	v785 = v767
	v786 = v778
	v787 = int32(13)
	goto L249
L249:
	;
	v792 = (v783 - int32(48)) & int32(255)
	if base.Ui32(v792) <= base.Ui32(int32(9)) {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v823 = v809
	v824 = v810
	goto L245
L251:
	;
	v797 = v784 & int32(1)
	if v797 != 0 {
		goto L254
	} else {
		goto L255
	}
L252:
	;
	v808 = v784
	v809 = v785
	v810 = v786
	v811 = v787
	goto L253
L253:
	;
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v782)+1)))
	if v813 == int32(0) {
		v823 = v809
		v824 = v810
		goto L245
	} else {
		goto L257
	}
L254:
	;
	v798 = int32(0)
	goto L256
L255:
	;
	v798 = v792
	goto L256
L256:
	;
	v804 = int32(1)
	v808 = v784 + v804
	v809 = v798 + v785
	v810 = (int32(0)-v797)&v792 + v786
	v811 = v787 - v804
	goto L253
L257:
	;
	v816 = int32(1)
	if base.Ui32(v816) < base.Ui32(v811) {
		v782 = v782 + v816
		v783 = v813
		v784 = v808
		v785 = v809
		v786 = v810
		v787 = v811
		goto L249
	} else {
		goto L258
	}
L258:
	;
	goto L250
L259:
	;
	v835 = v827 - v832
	goto L261
L260:
	;
	v835 = int32(0)
	goto L261
L261:
	;
	goto L241
L262:
	;
	if v858 == int32(0) {
		v1245 = v857
		goto L1
	} else {
		goto L263
	}
L263:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L186
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = l0
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_consts[1359])))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v872
	F_errmsg(m, int32(726377), v22)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L186
	} else {
		goto L265
	}
L265:
	;
	F_errsave_finish(m, l1, int32(496621), int32(902), int32(284464))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L186
	} else {
		goto L266
	}
L266:
	;
	v1245 = v857
	goto L1
L267:
	;
	if v886 == int32(0) {
		v1245 = v885
		goto L1
	} else {
		goto L268
	}
L268:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L186
	} else {
		goto L269
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l0
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_consts[1359])))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v900
	F_errmsg(m, int32(368436), v22+int32(16))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L186
	} else {
		goto L270
	}
L270:
	;
	F_errsave_finish(m, l1, int32(496621), int32(914), int32(284464))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L186
	} else {
		goto L271
	}
L271:
	;
	v1245 = v885
	goto L1
L272:
	;
	v943 = int32(*(*int8)(unsafe.Add(mBase, uint32(v932))))
	if v943 != 0 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	v952 = base.B2i32(v943 == int32(77))
	if v943 == int32(77) {
		goto L278
	} else {
		goto L279
	}
L274:
	;
	if v943 < int32(33) {
		v932 = v932 + int32(1)
		goto L272
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	goto L273
L277:
	;
	goto L276
L278:
	;
	v953 = int32(3)
	goto L280
L279:
	;
	v953 = int32(0)
	goto L280
L280:
	;
	if v943 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	v1028 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v932)+13)) = uint8(v1028)
	v1035 = base.I32_rem_u_s(v1024*int32(3)+v1016, int32(10))
	if v1035 != 0 {
		goto L295
	} else {
		goto L296
	}
L282:
	;
	v1016 = int32(0)
	v1024 = v953
	goto L281
L283:
	;
	goto L284
L284:
	;
	v966 = int32(0)
	v967 = v932
	v969 = v952
	v971 = v943
	v972 = int32(13)
	v974 = v953
	goto L285
L285:
	;
	v981 = (v971 - int32(48)) & int32(255)
	if base.Ui32(v981) <= base.Ui32(int32(9)) {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	v1016 = v997
	v1024 = v1001
	goto L281
L287:
	;
	v986 = v969 & int32(1)
	if v986 != 0 {
		goto L290
	} else {
		goto L291
	}
L288:
	;
	v997 = v966
	v999 = v969
	v1000 = v972
	v1001 = v974
	goto L289
L289:
	;
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v967)+1)))
	if v1002 == int32(0) {
		v1016 = v997
		v1024 = v1001
		goto L281
	} else {
		goto L293
	}
L290:
	;
	v987 = int32(0)
	goto L292
L291:
	;
	v987 = v981
	goto L292
L292:
	;
	v993 = int32(1)
	v997 = v987 + v966
	v999 = v969 + v993
	v1000 = v972 - v993
	v1001 = (int32(0)-v986)&v981 + v974
	goto L289
L293:
	;
	v1005 = int32(1)
	if base.Ui32(v1005) < base.Ui32(v1000) {
		v966 = v997
		v967 = v967 + v1005
		v969 = v999
		v971 = v1002
		v972 = v1000
		v974 = v1001
		goto L285
	} else {
		goto L294
	}
L294:
	;
	goto L286
L295:
	;
	v1038 = int32(58) - v1035
	goto L297
L296:
	;
	v1038 = int32(48)
	goto L297
L297:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v932)+12)) = uint8(v1038)
	if (v920|v178)&int32(1) != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	if v943 != 0 {
		goto L301
	} else {
		goto L302
	}
L299:
	;
	goto L300
L300:
	;
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1360])))
	if v1109 == int32(1) {
		goto L310
	} else {
		goto L311
	}
L301:
	;
	v1047 = v932
	v1048 = v943
	v1061 = v19
	goto L304
L302:
	;
	v1099 = v19
	goto L303
L303:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v1099 | base.I64_extend_i32_u(v920^int32(-1))&int64(1)
	v1245 = int32(1)
	goto L1
L304:
	;
	if base.Ui32((v1048-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	v1099 = v1074 << (uint(int64(1)) % 64)
	goto L303
L306:
	;
	v1074 = v1061*int64(10) + base.I64_extend_i32_u(v1048)&int64(15)
	goto L308
L307:
	;
	v1074 = v1061
	goto L308
L308:
	;
	v1075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1)))
	if v1075 != 0 {
		v1047 = v1047 + int32(1)
		v1048 = v1075
		v1061 = v1074
		goto L304
	} else {
		goto L309
	}
L309:
	;
	goto L305
L310:
	;
	if v943 != 0 {
		goto L313
	} else {
		goto L314
	}
L311:
	;
	goto L312
L312:
	;
	v1174 = int32(0)
	v1175 = F_errsave_start(m, l1)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L186
	} else {
		goto L322
	}
L313:
	;
	v1116 = v932
	v1117 = v943
	v1130 = v19
	goto L316
L314:
	;
	v1171 = int64(1)
	goto L315
L315:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v1171
	v1245 = int32(1)
	goto L1
L316:
	;
	if base.Ui32((v1117-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	v1147 = int64(1)
	v1171 = v1143<<(uint(v1147)%64) | v1147
	goto L315
L318:
	;
	v1143 = v1130*int64(10) + base.I64_extend_i32_u(v1117)&int64(15)
	goto L320
L319:
	;
	v1143 = v1130
	goto L320
L320:
	;
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116)+1)))
	if v1144 != 0 {
		v1116 = v1116 + int32(1)
		v1117 = v1144
		v1130 = v1143
		goto L316
	} else {
		goto L321
	}
L321:
	;
	goto L317
L322:
	;
	if v921 == int32(-1) {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	if v1175 == int32(0) {
		v1245 = v1174
		goto L1
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	if v1175 == int32(0) {
		v1245 = v1174
		goto L1
	} else {
		goto L330
	}
L326:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L186
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = l0
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_consts[1359])))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v1191
	F_errmsg(m, int32(726418), v22+int32(32))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L186
	} else {
		goto L328
	}
L328:
	;
	F_errsave_finish(m, l1, int32(496621), int32(888), int32(284464))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L186
	} else {
		goto L329
	}
L329:
	;
	v1245 = v1174
	goto L1
L330:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L186
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = l0
	if v921 == int32(10) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1217 = int32(88)
	goto L334
L333:
	;
	v1217 = v921 + int32(48)
	goto L334
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v1217
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_consts[1359])))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v1225
	F_errmsg(m, int32(501984), v22+int32(48))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L186
	} else {
		goto L335
	}
L335:
	;
	F_errsave_finish(m, l1, int32(496621), int32(895), int32(284464))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L186
	} else {
		goto L336
	}
L336:
	;
	v1245 = v1174
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
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
	v125 = F_strlen(m, v121)
	mBase = m.M
	goto L9
L11:
	;
	v121 = l1
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
	v118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v118)
	v121 = v114
	goto L10
L15:
	;
	v99 = v94
	v100 = v95
	v101 = v96
	goto L37
L16:
	;
	if v89 == int32(0) {
		v114 = v87
		v115 = v88
		goto L14
	} else {
		goto L36
	}
L17:
	;
	v87 = l1
	v88 = v12
	v89 = v19
	goto L16
L18:
	;
	goto L19
L19:
	;
	v23 = int32(0)
	if l1&int32(3) == v23 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v56 == int32(0) {
		v114 = v53
		v115 = v54
		goto L14
	} else {
		goto L29
	}
L21:
	;
	v53 = l1
	v54 = v12
	v55 = v19
	v56 = base.B2i32(v19 != v23)
	goto L20
L22:
	;
	if v19 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v32 = l1
	v33 = v12
	v34 = v19
	goto L24
L24:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v36)
	if v36 == int32(0) {
		v94 = v32
		v95 = v33
		v96 = v34
		goto L15
	} else {
		goto L26
	}
L25:
	;
	v53 = v47
	v54 = v41
	v55 = v43
	v56 = v45
	goto L20
L26:
	;
	v40 = int32(1)
	v41 = v33 + v40
	v43 = v34 - v40
	v44 = int32(0)
	v45 = base.B2i32(v43 != v44)
	v47 = v32 + v40
	if v47&int32(3) == v44 {
		v53 = v47
		v54 = v41
		v55 = v43
		v56 = v45
		goto L20
	} else {
		goto L27
	}
L27:
	;
	if v43 != 0 {
		v32 = v47
		v33 = v41
		v34 = v43
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v59 == int32(0) {
		v87 = v53
		v88 = v54
		v89 = v55
		goto L16
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(v55) < base.Ui32(int32(4)) {
		v87 = v53
		v88 = v54
		v89 = v55
		goto L16
	} else {
		goto L31
	}
L31:
	;
	v65 = v53
	v66 = v54
	v67 = v55
	goto L32
L32:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v73 = int32(-2139062144)
	if (int32(16843008)-v70|v70)&v73 != v73 {
		v94 = v65
		v95 = v66
		v96 = v67
		goto L15
	} else {
		goto L34
	}
L33:
	;
	v87 = v81
	v88 = v79
	v89 = v83
	goto L16
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v70
	v78 = int32(4)
	v79 = v66 + v78
	v81 = v65 + v78
	v83 = v67 - v78
	if base.Ui32(int32(3)) < base.Ui32(v83) {
		v65 = v81
		v66 = v79
		v67 = v83
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v94 = v87
	v95 = v88
	v96 = v89
	goto L15
L37:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
	if v103 == int32(0) {
		v114 = v99
		v115 = v100
		goto L14
	} else {
		goto L39
	}
L38:
	;
	v114 = v110
	v115 = v108
	goto L14
L39:
	;
	v107 = int32(1)
	v108 = v100 + v107
	v110 = v99 + v107
	v112 = v101 - v107
	if v112 != 0 {
		v99 = v110
		v100 = v108
		v101 = v112
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
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
					F_errmsg_internal(m, int32(487642), v7+int32(16))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(494519), int32(4933), int32(208730))
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
						F_errmsg_internal(m, int32(487924), v7)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494519), int32(4935), int32(208730))
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
							F_errmsg_internal(m, int32(487924), v7)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(494519), int32(4935), int32(208730))
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
	v20 = int32(8192)
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
	v50 = *(*int32)(unsafe.Add(mBase, _consts[659]))
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
	*(*int32)(unsafe.Add(mBase, _consts[660])) = v73
	goto L15
L17:
	;
	v53 = int32(150)
	v55 = *(*int32)(unsafe.Add(mBase, _consts[660]))
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
	v65 = *(*int32)(unsafe.Add(mBase, _consts[660]))
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
	v81 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	v84 = *(*int32)(unsafe.Add(mBase, _consts[660]))
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
	v91 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[659])) = int32(0)
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
	F_errmsg_internal(m, int32(515020), v14)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(495568), int32(1584), int32(407483))
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
	v163 = int32(4425592)
	v165 = *(*int32)(unsafe.Add(mBase, _consts[659]))
	*(*int32)(unsafe.Add(mBase, _consts[659])) = v165 + int32(1)
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
		*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0) - v4
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
		v22 = *(*int32)(unsafe.Add(mBase, _consts[1392]))
		if v22&(int32(1)<<(uint(l0&int32(7))%32)) == int32(0) {
			v224 = int32(16)
			m.G0 = v17 + v224
			m.G0 = v12 + v224
			return
		} else {
			v30 = m.G0
			v32 = v30 - int32(1168)
			m.G0 = v32
			v35 = *(*int32)(unsafe.Add(mBase, _consts[40]))
			v37 = *(*int32)(unsafe.Add(mBase, _consts[1111]))
			if v37 < int32(0) {
				v40 = int32(0)
				v45 = F_socket(m, int32(1), int32(524290), v40)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, _consts[1111])) = v45
				if v40 <= v45 {
					v49 = F_connect(m, v45)
					mBase = m.M
				} else {
				}
			} else {
			}
			v51 = *(*int32)(unsafe.Add(mBase, _consts[1393]))
			v52 = F___time(m)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, uint32(v32)+1144)) = v52
			v57 = v32 + int32(1100)
			v58 = F___gmtime_r(m, v32+int32(1144), v57)
			mBase = m.M
			v66 = F___strftime_l(m, v32+int32(1152), int32(16), int32(522697), v57, int32(4097160))
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
				v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
				if v75&int32(1) != 0 {
					v78 = int32(42)
				} else {
					v78 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v78
				v81 = base.B2i32(v78 == int32(0))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v81 + int32(508230)
				*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v81 + int32(508390)
				*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = int32(4692384)
				*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v71 | l0
				*(*int32)(unsafe.Add(mBase, uint32(v32)+40)) = v32 + int32(60)
				*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v32 + int32(1152)
				v103 = F_snprintf(m, v32-int32(-64), int32(1024), int32(745888), v32+int32(32))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[40])) = v35
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
							v133 = *(*int32)(unsafe.Add(mBase, _consts[1111]))
							v136 = int32(0)
							v138 = F_sendto(m, v133, v32-int32(-64), v131, v136, v136)
							mBase = m.M
							if int32(0) <= v138 {
								v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
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
								v143 = *(*int32)(unsafe.Add(mBase, _consts[40]))
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
									v170 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
									if v170&int32(2) == int32(0) {
										v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
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
										v178 = F_open(m, int32(385681), int32(524545), v177)
										mBase = m.M
										if v178 < v177 {
											v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
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
												v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
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
									v156 = *(*int32)(unsafe.Add(mBase, _consts[1111]))
									v157 = F_connect(m, v156)
									mBase = m.M
									if v157 < int32(0) {
										v170 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
										if v170&int32(2) == int32(0) {
											v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
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
											v178 = F_open(m, int32(385681), int32(524545), v177)
											mBase = m.M
											if v178 < v177 {
												v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
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
													v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
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
										v161 = *(*int32)(unsafe.Add(mBase, _consts[1111]))
										v164 = int32(0)
										v166 = F_sendto(m, v161, v32-int32(-64), v131, v164, v164)
										mBase = m.M
										if int32(0) <= v166 {
											v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
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
											v170 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
											if v170&int32(2) == int32(0) {
												v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
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
												v178 = F_open(m, int32(385681), int32(524545), v177)
												mBase = m.M
												if v178 < v177 {
													v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
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
														v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
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
