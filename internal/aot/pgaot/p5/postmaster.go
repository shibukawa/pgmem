package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitPostmasterChildSlots(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	*(*int32)(unsafe.Add(mBase, _consts[419])) = int32(32)
	v13 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[420])) = v13
	*(*int32)(unsafe.Add(mBase, _consts[421])) = v13
	*(*int32)(unsafe.Add(mBase, _consts[422])) = v13
	*(*int32)(unsafe.Add(mBase, _consts[423])) = v13
	*(*int32)(unsafe.Add(mBase, _consts[424])) = v13
	*(*int32)(unsafe.Add(mBase, _consts[425])) = v13
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v13
	*(*int32)(unsafe.Add(mBase, _consts[427])) = v13
	*(*int32)(unsafe.Add(mBase, _consts[428])) = v13
	*(*int32)(unsafe.Add(mBase, _consts[429])) = v13
	v44 = *(*int32)(unsafe.Add(mBase, _consts[430]))
	*(*int32)(unsafe.Add(mBase, _consts[431])) = v44
	v48 = *(*int32)(unsafe.Add(mBase, _consts[432]))
	*(*int32)(unsafe.Add(mBase, _consts[433])) = v48
	v52 = *(*int32)(unsafe.Add(mBase, _consts[434]))
	v54 = *(*int32)(unsafe.Add(mBase, _consts[435]))
	v57 = (v52 + v54) << (uint(v13) % 32)
	*(*int32)(unsafe.Add(mBase, _consts[436])) = v57
	v61 = *(*int32)(unsafe.Add(mBase, _consts[437]))
	v63 = *(*int32)(unsafe.Add(mBase, _consts[438]))
	v65 = *(*int32)(unsafe.Add(mBase, _consts[439]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[440]))
	v75 = v61 + (v63 + (v48 + (v44 + (v65 + (v67 + v57))))) + int32(42)
	*(*int32)(unsafe.Add(mBase, _consts[441])) = v75
	v79 = F_palloc(m, v75*int32(28))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		return
	} else {
		v84 = int32(0)
		v86 = int32(0)
		for {
			v91 = v86 << (uint(int32(4)) % 32)
			v95 = v91 + int32(4397592)
			*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[442]))) = v95
			*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[443]))) = v84 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[444]))) = v95
			v103 = int32(0)
			v106 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[440])))
			if v103 < v106 {
				v111 = v84
				v114 = v103
				for {
					v119 = v79 + v111*int32(28)
					*(*int64)(unsafe.Add(mBase, uint32(v119)+8)) = int64(0)
					v123 = v111 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v123
					v125 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v119))) = v125
					*(*uint8)(unsafe.Add(mBase, uint32(v119)+16)) = uint8(v125)
					v129 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[442])))
					if v129 == v125 {
						*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[442]))) = v95
						*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[444]))) = v91 + int32(4397592)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v119)+24)) = v95
					v137 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[444])))
					*(*int32)(unsafe.Add(mBase, uint32(v119)+20)) = v137
					v140 = v119 + int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v137)+4)) = v140
					*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[444]))) = v140
					v144 = v114 + int32(1)
					v145 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[440])))
					if v144 < v145 {
						v111 = v123
						v114 = v144
						continue
					} else {
						break
					}
					break
				}
				v149 = v123
			} else {
				v149 = v84
			}
			v156 = v86 + int32(1)
			if v156 != int32(18) {
				v84 = v149
				v86 = v156
				continue
			} else {
				break
			}
			break
		}
		v160 = int32(4397872)
		*(*int32)(unsafe.Add(mBase, _consts[445])) = v160
		*(*int32)(unsafe.Add(mBase, _consts[446])) = v160
		return
	}
}
func F_ReleasePostmasterChildSlot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v15 == int32(2) {
			if v18 != 0 {
				F_errmsg_internal(m, int32(426566), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499089), int32(241), int32(86207))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v93 = int32(1)
							m.G0 = v8 + int32(32)
							return v93
						}
					}
				}
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v93 = int32(1)
					m.G0 = v8 + int32(32)
					return v93
				}
			}
		} else {
			if v18 != 0 {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v36
				F_errmsg_internal(m, int32(467218), v8+int32(16))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499089), int32(249), int32(86207))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v50 == int32(6) {
							v57 = int32(4397600)
						} else {
							v57 = v50<<(uint(int32(4))%32) + int32(4397584)
						}
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
						if v48 < v58 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v102
								F_errmsg_internal(m, int32(83408), v8)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(499089), int32(262), int32(86207))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
							if v60+v58 <= v48 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v102
									F_errmsg_internal(m, int32(83408), v8)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(499089), int32(262), int32(86207))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v64 = v57 + int32(8)
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
								if v65 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v64))) = v64
									v69 = v64
								} else {
									v69 = v65
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v69
								v73 = l0 + int32(20)
								*(*int32)(unsafe.Add(mBase, uint32(v69))) = v73
								*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v73
								v77 = *(*int32)(unsafe.Add(mBase, _consts[447]))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v83 = v77 + v78<<(uint(int32(2))%32) + int32(44)
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
								*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
								v93 = base.B2i32(v84 == int32(1))
								m.G0 = v8 + int32(32)
								return v93
							}
						}
					}
				}
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v50 == int32(6) {
					v57 = int32(4397600)
				} else {
					v57 = v50<<(uint(int32(4))%32) + int32(4397584)
				}
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
				if v48 < v58 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v102
						F_errmsg_internal(m, int32(83408), v8)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(499089), int32(262), int32(86207))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					if v60+v58 <= v48 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v102
							F_errmsg_internal(m, int32(83408), v8)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(499089), int32(262), int32(86207))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v64 = v57 + int32(8)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
						if v65 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v64))) = v64
							v69 = v64
						} else {
							v69 = v65
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v69
						v73 = l0 + int32(20)
						*(*int32)(unsafe.Add(mBase, uint32(v69))) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v73
						v77 = *(*int32)(unsafe.Add(mBase, _consts[447]))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v83 = v77 + v78<<(uint(int32(2))%32) + int32(44)
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
						*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
						v93 = base.B2i32(v84 == int32(1))
						m.G0 = v8 + int32(32)
						return v93
					}
				}
			}
		}
	}
}
func F_postmaster_child_launch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	switch l0 - int32(1) {
	case 0, 5:
		v15 = m.G0
		v16 = int32(16)
		v17 = v15 - v16
		m.G0 = v17
		F___gettimeofday(m, v17)
		mBase = m.M
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
		v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+8)))
		m.G0 = v17 + v16
		*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v21 + v20*int64(1000000) - int64(946684800000000)
	default:
	}
	v31 = m.G0
	v33 = v31 - int32(144)
	m.G0 = v33
	v36 = F_fflush(m, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int32(0)
	} else {
		F_sigprocmask(m, int32(4396792), v33+int32(16))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(52)
			F_sigprocmask(m, v33+int32(16), int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				m.G0 = v33 + int32(144)
				return int32(-1)
			}
		}
	}
}
