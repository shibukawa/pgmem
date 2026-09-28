package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LocalBufferAlloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int64
	_ = v70
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v16
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[0]))
	if v21 == int32(0) {
		F_InitLocalBuffers(m)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[1]))
			F_ResourceOwnerEnlarge(m, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[0]))
				v36 = int32(0)
				v38 = F_hash_search(m, v33, v10+int32(12), v36, v36)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					if v38 != 0 {
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[2]))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
						v45 = v41 + v42*int32(56)
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
						v47 = int64(0)
						v50 = base.AtomicRmwCmpxchg64(m, v45, int32(24), v47, v47)
						v52 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[3]))
						v57 = v52 + (int32(-2)-v46)<<(uint(int32(2))%32)
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						if v58 == int32(0) {
							v61 = int32(_a_F_LocalBufferAlloc_0)
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[4]))
							*(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[4])) = v63 + int32(1)
							v70 = v50 + int64(1)
							if base.Ui64(v70&int64(3932160)) < base.Ui64(int64(1310720)) {
								v75 = v50 + int64(262145)
							} else {
								v75 = v70
							}
							*(*int64)(unsafe.Add(mBase, uint32(v45)+24)) = v75
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
							v78 = v77
							v79 = v75
						} else {
							v78 = v58
							v79 = v50
						}
						v80 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v57))) = v78 + v80
						v84 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[1]))
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
						F_ResourceOwnerRemember(m, v84, base.I64_extend_i32_s(v85+v80), int32(_a_F_LocalBufferAlloc_1))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v135 = v45
							v137 = int32(base.Ui32(base.I32_wrap_i64(v79))>>(uint(int32(24))%32)) & int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v137)
							m.G0 = v10 + int32(32)
							return v135
						}
					} else {
						v98 = F_GetLocalVictimBuffer(m)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v101 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[2]))
							v103 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[0]))
							v109 = F_hash_search(m, v103, v10+int32(12), int32(1), v10+int32(11))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int32(0)
							} else {
								v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
								if v111 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_LocalBufferAlloc_2), int32(0))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_LocalBufferAlloc_3), int32(160), int32(_a_F_LocalBufferAlloc_4))
											mBase = m.M
											v156 = m.ExcPending
											if v156 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v115 = v98 ^ int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v115
									v119 = v101 + v115*int32(56)
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v119)+16)) = v120
									v122 = *(*int64)(unsafe.Add(mBase, uint32(v10)+20))
									*(*int64)(unsafe.Add(mBase, uint32(v119)+8)) = v122
									v124 = *(*int64)(unsafe.Add(mBase, uint32(v10)+12))
									*(*int64)(unsafe.Add(mBase, uint32(v119))) = v124
									v126 = int64(0)
									v129 = base.AtomicRmwCmpxchg64(m, v119, int32(24), v126, v126)
									*(*int64)(unsafe.Add(mBase, uint32(v119)+24)) = v129&int64(-17179607041) | int64(33816576)
									v135 = v119
									v137 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v137)
									m.G0 = v10 + int32(32)
									return v135
								}
							}
						}
					}
				}
			}
		}
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[1]))
		F_ResourceOwnerEnlarge(m, v29)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[0]))
			v36 = int32(0)
			v38 = F_hash_search(m, v33, v10+int32(12), v36, v36)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				if v38 != 0 {
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[2]))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
					v45 = v41 + v42*int32(56)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
					v47 = int64(0)
					v50 = base.AtomicRmwCmpxchg64(m, v45, int32(24), v47, v47)
					v52 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[3]))
					v57 = v52 + (int32(-2)-v46)<<(uint(int32(2))%32)
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					if v58 == int32(0) {
						v61 = int32(_a_F_LocalBufferAlloc_0)
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[4]))
						*(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[4])) = v63 + int32(1)
						v70 = v50 + int64(1)
						if base.Ui64(v70&int64(3932160)) < base.Ui64(int64(1310720)) {
							v75 = v50 + int64(262145)
						} else {
							v75 = v70
						}
						*(*int64)(unsafe.Add(mBase, uint32(v45)+24)) = v75
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						v78 = v77
						v79 = v75
					} else {
						v78 = v58
						v79 = v50
					}
					v80 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v57))) = v78 + v80
					v84 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[1]))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
					F_ResourceOwnerRemember(m, v84, base.I64_extend_i32_s(v85+v80), int32(_a_F_LocalBufferAlloc_1))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v135 = v45
						v137 = int32(base.Ui32(base.I32_wrap_i64(v79))>>(uint(int32(24))%32)) & int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v137)
						m.G0 = v10 + int32(32)
						return v135
					}
				} else {
					v98 = F_GetLocalVictimBuffer(m)
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						v101 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[2]))
						v103 = *(*int32)(unsafe.Add(mBase, _c_F_LocalBufferAlloc[0]))
						v109 = F_hash_search(m, v103, v10+int32(12), int32(1), v10+int32(11))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
							if v111 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_LocalBufferAlloc_2), int32(0))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_LocalBufferAlloc_3), int32(160), int32(_a_F_LocalBufferAlloc_4))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v115 = v98 ^ int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v115
								v119 = v101 + v115*int32(56)
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v119)+16)) = v120
								v122 = *(*int64)(unsafe.Add(mBase, uint32(v10)+20))
								*(*int64)(unsafe.Add(mBase, uint32(v119)+8)) = v122
								v124 = *(*int64)(unsafe.Add(mBase, uint32(v10)+12))
								*(*int64)(unsafe.Add(mBase, uint32(v119))) = v124
								v126 = int64(0)
								v129 = base.AtomicRmwCmpxchg64(m, v119, int32(24), v126, v126)
								*(*int64)(unsafe.Add(mBase, uint32(v119)+24)) = v129&int64(-17179607041) | int64(33816576)
								v135 = v119
								v137 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v137)
								m.G0 = v10 + int32(32)
								return v135
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(15)))) = uint8(v12)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_readv_stage[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_readv_stage[0]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = (l0 - v23) >> (uint(int32(7)) % 32)
	v28 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+52)))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v28)
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v30)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
	if v32 != 0 {
		v35 = int32(0)
		for {
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_readv_stage[1]))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v16+v17<<(uint(int32(3))%32)+v35<<(uint(int32(3))%32))))
			v48 = v39 + (v43^int32(-1))*int32(56)
			v49 = int64(0)
			v52 = base.AtomicRmwCmpxchg64(m, v48, int32(24), v49, v49)
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v48)+44)) = v53
			v55 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
			*(*int64)(unsafe.Add(mBase, uint32(v48)+36)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = v52 + int64(1)
			v61 = v35 + int32(1)
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
			if base.Ui32(v61) < base.Ui32(v62) {
				v35 = v61
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	m.G0 = v8 + int32(16)
	return
}
