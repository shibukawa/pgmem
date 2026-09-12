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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	if l0 < v2 {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[15]))
		v16 = l0 ^ int32(-1)
		v21 = *(*int32)(unsafe.Add(mBase, _consts[612]))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v16<<(uint(int32(2))%32))))
		v27 = *(*int32)(unsafe.Add(mBase, _consts[72]))
		v90 = v14 + v16<<(uint(int32(6))%32)
		v91 = v25
		v92 = v27
		v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
		v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
		v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
		v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
		v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
		F_GetRelationPath(m, v9+int32(36), v96, v97, v98, v92, v99)
		mBase = m.M
		v101 = m.ExcPending
		if v101 != 0 {
			return int32(0)
		} else {
			v102 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v91
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93 & int32(262143)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v93 & int32(-4194304)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v102
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(36)
			v116 = F_psprintf(m, int32(661885), v9)
			mBase = m.M
			v117 = m.ExcPending
			if v117 != 0 {
				return int32(0)
			} else {
				m.G0 = v9 + int32(112)
				return v116
			}
		}
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, _consts[16]))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+108)) = l0
		v37 = *(*int32)(unsafe.Add(mBase, _consts[613]))
		if l0 == v37 {
			v84 = int32(4402768)
			v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
			v87 = v85
			v90 = v29 + l0<<(uint(int32(6))%32) + int32(-64)
			v91 = v87
			v92 = int32(-1)
			v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
			v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
			v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
			F_GetRelationPath(m, v9+int32(36), v96, v97, v98, v92, v99)
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return int32(0)
			} else {
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v91
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93 & int32(262143)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v93 & int32(-4194304)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v102
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(36)
				v116 = F_psprintf(m, int32(661885), v9)
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(112)
					return v116
				}
			}
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, _consts[614]))
			if l0 == v41 {
				v84 = int32(4402776)
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
				v87 = v85
				v90 = v29 + l0<<(uint(int32(6))%32) + int32(-64)
				v91 = v87
				v92 = int32(-1)
				v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
				F_GetRelationPath(m, v9+int32(36), v96, v97, v98, v92, v99)
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int32(0)
				} else {
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v91
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93 & int32(262143)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v93 & int32(-4194304)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v102
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(36)
					v116 = F_psprintf(m, int32(661885), v9)
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(112)
						return v116
					}
				}
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, _consts[615]))
				if l0 == v45 {
					v84 = int32(4402784)
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
					v87 = v85
					v90 = v29 + l0<<(uint(int32(6))%32) + int32(-64)
					v91 = v87
					v92 = int32(-1)
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
					F_GetRelationPath(m, v9+int32(36), v96, v97, v98, v92, v99)
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						v102 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v91
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93 & int32(262143)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v93 & int32(-4194304)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v102
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(36)
						v116 = F_psprintf(m, int32(661885), v9)
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(112)
							return v116
						}
					}
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, _consts[616]))
					if l0 == v49 {
						v84 = int32(4402792)
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
						v87 = v85
						v90 = v29 + l0<<(uint(int32(6))%32) + int32(-64)
						v91 = v87
						v92 = int32(-1)
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
						F_GetRelationPath(m, v9+int32(36), v96, v97, v98, v92, v99)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93 & int32(262143)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v93 & int32(-4194304)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v102
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(36)
							v116 = F_psprintf(m, int32(661885), v9)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(112)
								return v116
							}
						}
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, _consts[617]))
						if l0 == v53 {
							v84 = int32(4402800)
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
							v87 = v85
							v90 = v29 + l0<<(uint(int32(6))%32) + int32(-64)
							v91 = v87
							v92 = int32(-1)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
							F_GetRelationPath(m, v9+int32(36), v96, v97, v98, v92, v99)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v91
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93 & int32(262143)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v93 & int32(-4194304)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v102
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(36)
								v116 = F_psprintf(m, int32(661885), v9)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(112)
									return v116
								}
							}
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, _consts[618]))
							if l0 == v57 {
								v84 = int32(4402808)
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
								v87 = v85
								v90 = v29 + l0<<(uint(int32(6))%32) + int32(-64)
								v91 = v87
								v92 = int32(-1)
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
								F_GetRelationPath(m, v9+int32(36), v96, v97, v98, v92, v99)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v91
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93 & int32(262143)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v93 & int32(-4194304)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v102
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(36)
									v116 = F_psprintf(m, int32(661885), v9)
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(112)
										return v116
									}
								}
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, _consts[619]))
								if l0 == v61 {
									v84 = int32(4402816)
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
									v87 = v85
									v90 = v29 + l0<<(uint(int32(6))%32) + int32(-64)
									v91 = v87
									v92 = int32(-1)
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
									F_GetRelationPath(m, v9+int32(36), v96, v97, v98, v92, v99)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v91
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93 & int32(262143)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v93 & int32(-4194304)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v102
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
										*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(36)
										v116 = F_psprintf(m, int32(661885), v9)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(112)
											return v116
										}
									}
								} else {
									v65 = *(*int32)(unsafe.Add(mBase, _consts[620]))
									if l0 == v65 {
										v84 = int32(4402824)
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
										v87 = v85
										v90 = v29 + l0<<(uint(int32(6))%32) + int32(-64)
										v91 = v87
										v92 = int32(-1)
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
										F_GetRelationPath(m, v9+int32(36), v96, v97, v98, v92, v99)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return int32(0)
										} else {
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v91
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93 & int32(262143)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v93 & int32(-4194304)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v102
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
											*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(36)
											v116 = F_psprintf(m, int32(661885), v9)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												m.G0 = v9 + int32(112)
												return v116
											}
										}
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, _consts[604]))
										if v69 == int32(0) {
											v87 = v2
											v90 = v29 + l0<<(uint(int32(6))%32) + int32(-64)
											v91 = v87
											v92 = int32(-1)
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
											v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
											v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
											F_GetRelationPath(m, v9+int32(36), v96, v97, v98, v92, v99)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int32(0)
											} else {
												v102 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v91
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93 & int32(262143)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v93 & int32(-4194304)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v102
												*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
												*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(36)
												v116 = F_psprintf(m, int32(661885), v9)
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int32(0)
												} else {
													m.G0 = v9 + int32(112)
													return v116
												}
											}
										} else {
											v73 = *(*int32)(unsafe.Add(mBase, _consts[621]))
											v76 = int32(0)
											v78 = F_hash_search(m, v73, v9+int32(108), v76, v76)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												if v78 == int32(0) {
													v87 = v2
												} else {
													v84 = v78
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
													v87 = v85
												}
												v90 = v29 + l0<<(uint(int32(6))%32) + int32(-64)
												v91 = v87
												v92 = int32(-1)
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
												v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
												F_GetRelationPath(m, v9+int32(36), v96, v97, v98, v92, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													v102 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v91
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93 & int32(262143)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v93 & int32(-4194304)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v102
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
													*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(36)
													v116 = F_psprintf(m, int32(661885), v9)
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return int32(0)
													} else {
														m.G0 = v9 + int32(112)
														return v116
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
