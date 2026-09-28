package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_XLOGShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v239 int32
	_ = v239
	var v243 int64
	_ = v243
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[0]))
	base.MemoryFill(m, v11, v2, int32(448))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[1]))
	if v16 != 0 {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[2]))
		base.MemoryCopy(m, v18, v16, int32(312))
		F_pfree(m, v16)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[1])) = int32(0)
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[0]))
			v28 = v27
			v30 = v28 + int32(448)
			*(*int32)(unsafe.Add(mBase, uint32(v28)+292)) = v30
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[3]))
			if v33 <= int32(0) {
			} else {
				v37 = v33 & int32(3)
				v38 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v33) {
					v43 = v38
					v50 = v2
					for {
						v53 = v43 << (uint(int32(3)) % 32)
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
						v56 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v53+v54))) = v56
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
						*(*int64)(unsafe.Add(mBase, uint32(v58+v53)+8)) = v56
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
						*(*int64)(unsafe.Add(mBase, uint32(v62+v53)+16)) = v56
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
						*(*int64)(unsafe.Add(mBase, uint32(v66+v53)+24)) = v56
						v70 = int32(4)
						v71 = v43 + v70
						v73 = v50 + v70
						if v73 != v33&int32(-4) {
							v43 = v71
							v50 = v73
							continue
						} else {
							break
						}
						break
					}
					if v37 == int32(0) {
					} else {
						v77 = v71
						v86 = v77
						v94 = v2
						for {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
							*(*int64)(unsafe.Add(mBase, uint32(v95+v86<<(uint(int32(3))%32)))) = int64(0)
							v101 = int32(1)
							v104 = v94 + v101
							if v104 != v37 {
								v86 = v86 + v101
								v94 = v104
								continue
							} else {
								break
							}
							break
						}
					}
				} else {
					v77 = v38
					v86 = v77
					v94 = v2
					for {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
						*(*int64)(unsafe.Add(mBase, uint32(v95+v86<<(uint(int32(3))%32)))) = int64(0)
						v101 = int32(1)
						v104 = v94 + v101
						if v104 != v37 {
							v86 = v86 + v101
							v94 = v104
							continue
						} else {
							break
						}
						break
					}
				}
			}
			v119 = (v30 + v33<<(uint(int32(3))%32)) & int32(-128)
			v121 = v119 + int32(128)
			*(*int32)(unsafe.Add(mBase, uint32(v28)+176)) = v121
			*(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4])) = v121
			F_LWLockInitialize(m, v121, int32(66))
			mBase = m.M
			v127 = m.ExcPending
			if v127 != 0 {
				return
			} else {
				v129 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
				v130 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v129)+24)) = v130
				*(*int64)(unsafe.Add(mBase, uint32(v129)+16)) = v130
				F_LWLockInitialize(m, v129+int32(128), int32(66))
				mBase = m.M
				v138 = m.ExcPending
				if v138 != 0 {
					return
				} else {
					v140 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
					v141 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v140)+152)) = v141
					*(*int64)(unsafe.Add(mBase, uint32(v140)+144)) = v141
					F_LWLockInitialize(m, v140+int32(256), int32(66))
					mBase = m.M
					v149 = m.ExcPending
					if v149 != 0 {
						return
					} else {
						v151 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
						v152 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v151)+280)) = v152
						*(*int64)(unsafe.Add(mBase, uint32(v151)+272)) = v152
						F_LWLockInitialize(m, v151+int32(384), int32(66))
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return
						} else {
							v162 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
							v163 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v162)+408)) = v163
							*(*int64)(unsafe.Add(mBase, uint32(v162)+400)) = v163
							F_LWLockInitialize(m, v162+int32(512), int32(66))
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return
							} else {
								v173 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
								v174 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v173)+536)) = v174
								*(*int64)(unsafe.Add(mBase, uint32(v173)+528)) = v174
								F_LWLockInitialize(m, v173+int32(640), int32(66))
								mBase = m.M
								v182 = m.ExcPending
								if v182 != 0 {
									return
								} else {
									v184 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
									v185 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v184)+664)) = v185
									*(*int64)(unsafe.Add(mBase, uint32(v184)+656)) = v185
									F_LWLockInitialize(m, v184+int32(768), int32(66))
									mBase = m.M
									v193 = m.ExcPending
									if v193 != 0 {
										return
									} else {
										v195 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
										v196 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v195)+792)) = v196
										*(*int64)(unsafe.Add(mBase, uint32(v195)+784)) = v196
										F_LWLockInitialize(m, v195+int32(896), int32(66))
										mBase = m.M
										v204 = m.ExcPending
										if v204 != 0 {
											return
										} else {
											v206 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
											v207 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v206)+920)) = v207
											*(*int64)(unsafe.Add(mBase, uint32(v206)+912)) = v207
											v212 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[0]))
											v216 = (v119 + int32(_a_F_XLOGShmemInit_0)) & int32(-8192)
											*(*int32)(unsafe.Add(mBase, uint32(v212)+288)) = v216
											v219 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[3]))
											v221 = v219 << (uint(int32(13)) % 32)
											if v221 != 0 {
												base.MemoryFill(m, v216, int32(0), v221)
											} else {
											}
											v225 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[3]))
											v226 = int32(_a_F_XLOGShmemInit_1)
											v227 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[0]))
											v228 = int32(0)
											*(*uint16)(unsafe.Add(mBase, uint32(v227)+312)) = uint16(v228)
											*(*int32)(unsafe.Add(mBase, uint32(v227)+308)) = v228
											*(*int32)(unsafe.Add(mBase, uint32(v227)+296)) = v225 - int32(1)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v227))), uint32(v228))
											v239 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[0]))
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v239)+440)), uint32(v228))
											v243 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v239)+256)) = v243
											*(*int64)(unsafe.Add(mBase, uint32(v239)+264)) = v243
											*(*int64)(unsafe.Add(mBase, uint32(v239)+272)) = v243
											*(*int64)(unsafe.Add(mBase, uint32(v239)+232)) = v243
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
	} else {
		v28 = v11
		v30 = v28 + int32(448)
		*(*int32)(unsafe.Add(mBase, uint32(v28)+292)) = v30
		v33 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[3]))
		if v33 <= int32(0) {
		} else {
			v37 = v33 & int32(3)
			v38 = int32(0)
			if base.Ui32(int32(4)) <= base.Ui32(v33) {
				v43 = v38
				v50 = v2
				for {
					v53 = v43 << (uint(int32(3)) % 32)
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
					v56 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v53+v54))) = v56
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
					*(*int64)(unsafe.Add(mBase, uint32(v58+v53)+8)) = v56
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
					*(*int64)(unsafe.Add(mBase, uint32(v62+v53)+16)) = v56
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
					*(*int64)(unsafe.Add(mBase, uint32(v66+v53)+24)) = v56
					v70 = int32(4)
					v71 = v43 + v70
					v73 = v50 + v70
					if v73 != v33&int32(-4) {
						v43 = v71
						v50 = v73
						continue
					} else {
						break
					}
					break
				}
				if v37 == int32(0) {
				} else {
					v77 = v71
					v86 = v77
					v94 = v2
					for {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
						*(*int64)(unsafe.Add(mBase, uint32(v95+v86<<(uint(int32(3))%32)))) = int64(0)
						v101 = int32(1)
						v104 = v94 + v101
						if v104 != v37 {
							v86 = v86 + v101
							v94 = v104
							continue
						} else {
							break
						}
						break
					}
				}
			} else {
				v77 = v38
				v86 = v77
				v94 = v2
				for {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
					*(*int64)(unsafe.Add(mBase, uint32(v95+v86<<(uint(int32(3))%32)))) = int64(0)
					v101 = int32(1)
					v104 = v94 + v101
					if v104 != v37 {
						v86 = v86 + v101
						v94 = v104
						continue
					} else {
						break
					}
					break
				}
			}
		}
		v119 = (v30 + v33<<(uint(int32(3))%32)) & int32(-128)
		v121 = v119 + int32(128)
		*(*int32)(unsafe.Add(mBase, uint32(v28)+176)) = v121
		*(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4])) = v121
		F_LWLockInitialize(m, v121, int32(66))
		mBase = m.M
		v127 = m.ExcPending
		if v127 != 0 {
			return
		} else {
			v129 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
			v130 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v129)+24)) = v130
			*(*int64)(unsafe.Add(mBase, uint32(v129)+16)) = v130
			F_LWLockInitialize(m, v129+int32(128), int32(66))
			mBase = m.M
			v138 = m.ExcPending
			if v138 != 0 {
				return
			} else {
				v140 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
				v141 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v140)+152)) = v141
				*(*int64)(unsafe.Add(mBase, uint32(v140)+144)) = v141
				F_LWLockInitialize(m, v140+int32(256), int32(66))
				mBase = m.M
				v149 = m.ExcPending
				if v149 != 0 {
					return
				} else {
					v151 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
					v152 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v151)+280)) = v152
					*(*int64)(unsafe.Add(mBase, uint32(v151)+272)) = v152
					F_LWLockInitialize(m, v151+int32(384), int32(66))
					mBase = m.M
					v160 = m.ExcPending
					if v160 != 0 {
						return
					} else {
						v162 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
						v163 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v162)+408)) = v163
						*(*int64)(unsafe.Add(mBase, uint32(v162)+400)) = v163
						F_LWLockInitialize(m, v162+int32(512), int32(66))
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
							v174 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v173)+536)) = v174
							*(*int64)(unsafe.Add(mBase, uint32(v173)+528)) = v174
							F_LWLockInitialize(m, v173+int32(640), int32(66))
							mBase = m.M
							v182 = m.ExcPending
							if v182 != 0 {
								return
							} else {
								v184 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
								v185 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v184)+664)) = v185
								*(*int64)(unsafe.Add(mBase, uint32(v184)+656)) = v185
								F_LWLockInitialize(m, v184+int32(768), int32(66))
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return
								} else {
									v195 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
									v196 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v195)+792)) = v196
									*(*int64)(unsafe.Add(mBase, uint32(v195)+784)) = v196
									F_LWLockInitialize(m, v195+int32(896), int32(66))
									mBase = m.M
									v204 = m.ExcPending
									if v204 != 0 {
										return
									} else {
										v206 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[4]))
										v207 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v206)+920)) = v207
										*(*int64)(unsafe.Add(mBase, uint32(v206)+912)) = v207
										v212 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[0]))
										v216 = (v119 + int32(_a_F_XLOGShmemInit_0)) & int32(-8192)
										*(*int32)(unsafe.Add(mBase, uint32(v212)+288)) = v216
										v219 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[3]))
										v221 = v219 << (uint(int32(13)) % 32)
										if v221 != 0 {
											base.MemoryFill(m, v216, int32(0), v221)
										} else {
										}
										v225 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[3]))
										v226 = int32(_a_F_XLOGShmemInit_1)
										v227 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[0]))
										v228 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v227)+312)) = uint16(v228)
										*(*int32)(unsafe.Add(mBase, uint32(v227)+308)) = v228
										*(*int32)(unsafe.Add(mBase, uint32(v227)+296)) = v225 - int32(1)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v227))), uint32(v228))
										v239 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemInit[0]))
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v239)+440)), uint32(v228))
										v243 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v239)+256)) = v243
										*(*int64)(unsafe.Add(mBase, uint32(v239)+264)) = v243
										*(*int64)(unsafe.Add(mBase, uint32(v239)+272)) = v243
										*(*int64)(unsafe.Add(mBase, uint32(v239)+232)) = v243
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
