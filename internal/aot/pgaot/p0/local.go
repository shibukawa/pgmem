package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LocalBufferAlloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v15
	v20 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	if v20 == int32(0) {
		F_InitLocalBuffers(m)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, _consts[579]))
			F_ResourceOwnerEnlarge(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _consts[581]))
				v35 = int32(0)
				v37 = F_hash_search(m, v32, v9+int32(12), v35, v35)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					if v37 != 0 {
						v40 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
						v44 = v40 + v41<<(uint(int32(6))%32)
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
						v47 = *(*int32)(unsafe.Add(mBase, _consts[580]))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
						v53 = v47 + (int32(-2)-v49)<<(uint(int32(2))%32)
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
						if v54 == int32(0) {
							v57 = int32(4464712)
							v59 = *(*int32)(unsafe.Add(mBase, _consts[582]))
							v60 = int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[582])) = v59 + v60
							v66 = v45 + v60
							if base.Ui32(v66&int32(3932160)) < base.Ui32(int32(1310720)) {
								v71 = v45 + int32(262145)
							} else {
								v71 = v66
							}
							*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v71
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
							v74 = v71
							v75 = v73
						} else {
							v74 = v45
							v75 = v54
						}
						v76 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v53))) = v75 + v76
						v80 = *(*int32)(unsafe.Add(mBase, _consts[579]))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
						F_ResourceOwnerRemember(m, v80, v81+v76, int32(1656256))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							v126 = v44
							v128 = int32(base.Ui32(v74)>>(uint(int32(24))%32)) & int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v128)
							m.G0 = v9 + int32(32)
							return v126
						}
					} else {
						v92 = F_GetLocalVictimBuffer(m)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, _consts[9]))
							v97 = *(*int32)(unsafe.Add(mBase, _consts[581]))
							v103 = F_hash_search(m, v97, v9+int32(12), int32(1), v9+int32(11))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
								if v105 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(462097), int32(0))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(518370), int32(159), int32(507436))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v109 = v92 ^ int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v103)+20)) = v109
									v113 = v95 + v109<<(uint(int32(6))%32)
									v114 = *(*int64)(unsafe.Add(mBase, uint32(v9)+12))
									*(*int64)(unsafe.Add(mBase, uint32(v113))) = v114
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v113)+16)) = v116
									v118 = *(*int64)(unsafe.Add(mBase, uint32(v9)+20))
									*(*int64)(unsafe.Add(mBase, uint32(v113)+8)) = v118
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v113)+24)) = v120&int32(262143) | int32(33816576)
									v126 = v113
									v128 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v128)
									m.G0 = v9 + int32(32)
									return v126
								}
							}
						}
					}
				}
			}
		}
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, _consts[579]))
		F_ResourceOwnerEnlarge(m, v28)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, _consts[581]))
			v35 = int32(0)
			v37 = F_hash_search(m, v32, v9+int32(12), v35, v35)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				if v37 != 0 {
					v40 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
					v44 = v40 + v41<<(uint(int32(6))%32)
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
					v47 = *(*int32)(unsafe.Add(mBase, _consts[580]))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
					v53 = v47 + (int32(-2)-v49)<<(uint(int32(2))%32)
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
					if v54 == int32(0) {
						v57 = int32(4464712)
						v59 = *(*int32)(unsafe.Add(mBase, _consts[582]))
						v60 = int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[582])) = v59 + v60
						v66 = v45 + v60
						if base.Ui32(v66&int32(3932160)) < base.Ui32(int32(1310720)) {
							v71 = v45 + int32(262145)
						} else {
							v71 = v66
						}
						*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v71
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
						v74 = v71
						v75 = v73
					} else {
						v74 = v45
						v75 = v54
					}
					v76 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v53))) = v75 + v76
					v80 = *(*int32)(unsafe.Add(mBase, _consts[579]))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
					F_ResourceOwnerRemember(m, v80, v81+v76, int32(1656256))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						v126 = v44
						v128 = int32(base.Ui32(v74)>>(uint(int32(24))%32)) & int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v128)
						m.G0 = v9 + int32(32)
						return v126
					}
				} else {
					v92 = F_GetLocalVictimBuffer(m)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						v97 = *(*int32)(unsafe.Add(mBase, _consts[581]))
						v103 = F_hash_search(m, v97, v9+int32(12), int32(1), v9+int32(11))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
							if v105 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(462097), int32(0))
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(518370), int32(159), int32(507436))
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v109 = v92 ^ int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v103)+20)) = v109
								v113 = v95 + v109<<(uint(int32(6))%32)
								v114 = *(*int64)(unsafe.Add(mBase, uint32(v9)+12))
								*(*int64)(unsafe.Add(mBase, uint32(v113))) = v114
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v113)+16)) = v116
								v118 = *(*int64)(unsafe.Add(mBase, uint32(v9)+20))
								*(*int64)(unsafe.Add(mBase, uint32(v113)+8)) = v118
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v113)+24)) = v120&int32(262143) | int32(33816576)
								v126 = v113
								v128 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v128)
								m.G0 = v9 + int32(32)
								return v126
							}
						}
					}
				}
			}
		}
	}
}
func F_init_local_reloptions(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	return
}
func F_local_buffer_readv_stage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(15)))) = uint8(v17)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[554]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v25 = v21 + v22<<(uint(int32(3))%32)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[554]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = (l0 - v28) >> (uint(int32(7)) % 32)
	v33 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+52)))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+4)) = uint32(v33)
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+8)) = uint32(v35)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v37 == int32(0) {
	} else {
		v40 = int32(1)
		v42 = int32(0)
		if v37 != v40 {
			v49 = v42
			v52 = int32(0)
			for {
				v57 = int32(4464696)
				v58 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v61 = v25 + v49<<(uint(int32(3))%32)
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
				v63 = int32(-1)
				v65 = int32(6)
				v67 = v58 + (v62^v63)<<(uint(v65)%32)
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
				v69 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
				*(*int64)(unsafe.Add(mBase, uint32(v67)+36)) = v69
				v72 = v13 + int32(8)
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
				*(*int32)(unsafe.Add(mBase, uint32(v67)+44)) = v73
				v75 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v67)+24)) = v68 + v75
				v79 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
				v85 = v79 + (v80^v63)<<(uint(v65)%32)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+24))
				v87 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
				*(*int64)(unsafe.Add(mBase, uint32(v85)+36)) = v87
				v89 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
				*(*int32)(unsafe.Add(mBase, uint32(v85)+44)) = v89
				*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v86 + v75
				v94 = int32(2)
				v95 = v49 + v94
				v97 = v52 + v94
				if v97 != v37&int32(254) {
					v49 = v95
					v52 = v97
					continue
				} else {
					break
				}
				break
			}
			v101 = v95
		} else {
			v101 = v42
		}
		if v37&v40 == int32(0) {
		} else {
			v112 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			v116 = *(*int32)(unsafe.Add(mBase, uint32(v25+v101<<(uint(int32(3))%32))))
			v121 = v112 + (v116^int32(-1))<<(uint(int32(6))%32)
			v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+24))
			v123 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
			*(*int64)(unsafe.Add(mBase, uint32(v121)+36)) = v123
			v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v121)+44)) = v125
			*(*int32)(unsafe.Add(mBase, uint32(v121)+24)) = v122 + int32(1)
		}
	}
	m.G0 = v13 + int32(16)
	return
}
