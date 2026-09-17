package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResOwnerPrintBufferPin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
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
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
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
	var v103 int32
	_ = v103
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	if l0 < v2 {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[0]))
		v17 = l0 ^ int32(-1)
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[1]))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v17<<(uint(int32(2))%32))))
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[2]))
		v91 = v15 + v17<<(uint(int32(6))%32)
		v92 = v26
		v93 = v28
		v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
		v96 = v10 + int32(24)
		v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
		v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
		v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
		v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
		F_GetRelationPath(m, v96, v97, v98, v99, v93, v100)
		mBase = m.M
		v102 = m.ExcPending
		if v102 != 0 {
			return int32(0)
		} else {
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v92
			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v94 & int32(_a_F_ResOwnerPrintBufferPin_0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v94 & int32(-4194304)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v103
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96
			v115 = F_psprintf(m, int32(_a_F_ResOwnerPrintBufferPin_1), v10)
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(96)
				return v115
			}
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[3]))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l0
		v38 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[4]))
		if l0 == v38 {
			v85 = int32(_a_F_ResOwnerPrintBufferPin_2)
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
			v88 = v86
			v91 = v30 + l0<<(uint(int32(6))%32) + int32(-64)
			v92 = v88
			v93 = int32(-1)
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
			v96 = v10 + int32(24)
			v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
			v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
			v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
			F_GetRelationPath(m, v96, v97, v98, v99, v93, v100)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v92
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v94 & int32(_a_F_ResOwnerPrintBufferPin_0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v94 & int32(-4194304)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v103
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96
				v115 = F_psprintf(m, int32(_a_F_ResOwnerPrintBufferPin_1), v10)
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					m.G0 = v10 + int32(96)
					return v115
				}
			}
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[5]))
			if l0 == v42 {
				v85 = int32(_a_F_ResOwnerPrintBufferPin_3)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
				v88 = v86
				v91 = v30 + l0<<(uint(int32(6))%32) + int32(-64)
				v92 = v88
				v93 = int32(-1)
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
				v96 = v10 + int32(24)
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
				v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
				F_GetRelationPath(m, v96, v97, v98, v99, v93, v100)
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v92
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v94 & int32(_a_F_ResOwnerPrintBufferPin_0)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v94 & int32(-4194304)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v103
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96
					v115 = F_psprintf(m, int32(_a_F_ResOwnerPrintBufferPin_1), v10)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(96)
						return v115
					}
				}
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[6]))
				if l0 == v46 {
					v85 = int32(_a_F_ResOwnerPrintBufferPin_4)
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
					v88 = v86
					v91 = v30 + l0<<(uint(int32(6))%32) + int32(-64)
					v92 = v88
					v93 = int32(-1)
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
					v96 = v10 + int32(24)
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
					F_GetRelationPath(m, v96, v97, v98, v99, v93, v100)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v92
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v94 & int32(_a_F_ResOwnerPrintBufferPin_0)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v94 & int32(-4194304)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v103
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96
						v115 = F_psprintf(m, int32(_a_F_ResOwnerPrintBufferPin_1), v10)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 + int32(96)
							return v115
						}
					}
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[7]))
					if l0 == v50 {
						v85 = int32(_a_F_ResOwnerPrintBufferPin_5)
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
						v88 = v86
						v91 = v30 + l0<<(uint(int32(6))%32) + int32(-64)
						v92 = v88
						v93 = int32(-1)
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
						v96 = v10 + int32(24)
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
						F_GetRelationPath(m, v96, v97, v98, v99, v93, v100)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v94 & int32(_a_F_ResOwnerPrintBufferPin_0)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v94 & int32(-4194304)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v103
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96
							v115 = F_psprintf(m, int32(_a_F_ResOwnerPrintBufferPin_1), v10)
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(96)
								return v115
							}
						}
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[8]))
						if l0 == v54 {
							v85 = int32(_a_F_ResOwnerPrintBufferPin_6)
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
							v88 = v86
							v91 = v30 + l0<<(uint(int32(6))%32) + int32(-64)
							v92 = v88
							v93 = int32(-1)
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
							v96 = v10 + int32(24)
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
							F_GetRelationPath(m, v96, v97, v98, v99, v93, v100)
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int32(0)
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v92
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v94 & int32(_a_F_ResOwnerPrintBufferPin_0)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v94 & int32(-4194304)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v103
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96
								v115 = F_psprintf(m, int32(_a_F_ResOwnerPrintBufferPin_1), v10)
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(96)
									return v115
								}
							}
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[9]))
							if l0 == v58 {
								v85 = int32(_a_F_ResOwnerPrintBufferPin_7)
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
								v88 = v86
								v91 = v30 + l0<<(uint(int32(6))%32) + int32(-64)
								v92 = v88
								v93 = int32(-1)
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
								v96 = v10 + int32(24)
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
								F_GetRelationPath(m, v96, v97, v98, v99, v93, v100)
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v92
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v94 & int32(_a_F_ResOwnerPrintBufferPin_0)
									*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v94 & int32(-4194304)
									*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v103
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
									*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96
									v115 = F_psprintf(m, int32(_a_F_ResOwnerPrintBufferPin_1), v10)
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(96)
										return v115
									}
								}
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[10]))
								if l0 == v62 {
									v85 = int32(_a_F_ResOwnerPrintBufferPin_8)
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
									v88 = v86
									v91 = v30 + l0<<(uint(int32(6))%32) + int32(-64)
									v92 = v88
									v93 = int32(-1)
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
									v96 = v10 + int32(24)
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
									F_GetRelationPath(m, v96, v97, v98, v99, v93, v100)
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v92
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v94 & int32(_a_F_ResOwnerPrintBufferPin_0)
										*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v94 & int32(-4194304)
										*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v103
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
										*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96
										v115 = F_psprintf(m, int32(_a_F_ResOwnerPrintBufferPin_1), v10)
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(96)
											return v115
										}
									}
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[11]))
									if l0 == v66 {
										v85 = int32(_a_F_ResOwnerPrintBufferPin_9)
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
										v88 = v86
										v91 = v30 + l0<<(uint(int32(6))%32) + int32(-64)
										v92 = v88
										v93 = int32(-1)
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
										v96 = v10 + int32(24)
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
										F_GetRelationPath(m, v96, v97, v98, v99, v93, v100)
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int32(0)
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
											*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v92
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v94 & int32(_a_F_ResOwnerPrintBufferPin_0)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v94 & int32(-4194304)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v103
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
											*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96
											v115 = F_psprintf(m, int32(_a_F_ResOwnerPrintBufferPin_1), v10)
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int32(0)
											} else {
												m.G0 = v10 + int32(96)
												return v115
											}
										}
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[12]))
										if v70 == int32(0) {
											v88 = v2
											v91 = v30 + l0<<(uint(int32(6))%32) + int32(-64)
											v92 = v88
											v93 = int32(-1)
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
											v96 = v10 + int32(24)
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
											v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
											F_GetRelationPath(m, v96, v97, v98, v99, v93, v100)
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int32(0)
											} else {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v92
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v94 & int32(_a_F_ResOwnerPrintBufferPin_0)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v94 & int32(-4194304)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v103
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
												*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96
												v115 = F_psprintf(m, int32(_a_F_ResOwnerPrintBufferPin_1), v10)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													m.G0 = v10 + int32(96)
													return v115
												}
											}
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBufferPin[13]))
											v77 = int32(0)
											v79 = F_hash_search(m, v74, v10+int32(24), v77, v77)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												if v79 == int32(0) {
													v88 = v2
												} else {
													v85 = v79
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
													v88 = v86
												}
												v91 = v30 + l0<<(uint(int32(6))%32) + int32(-64)
												v92 = v88
												v93 = int32(-1)
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
												v96 = v10 + int32(24)
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
												v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
												F_GetRelationPath(m, v96, v97, v98, v99, v93, v100)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int32(0)
												} else {
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v92
													*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v94 & int32(_a_F_ResOwnerPrintBufferPin_0)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v94 & int32(-4194304)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v103
													*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
													*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96
													v115 = F_psprintf(m, int32(_a_F_ResOwnerPrintBufferPin_1), v10)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return int32(0)
													} else {
														m.G0 = v10 + int32(96)
														return v115
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
func F_ResOwnerReleaseCachedPlan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = v3 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v5
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			F_MemoryContextDelete(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ResOwnerReleaseTupleDesc(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = v3 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	if v5 == int32(0) {
		F_FreeTupleDesc(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
