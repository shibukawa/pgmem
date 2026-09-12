package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bbsink_copystream_archive_contents(m *base.Module, l0 int32, l1 int32) {
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v11 == int32(1) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v19 = *(*int32)(unsafe.Add(mBase, _consts[218]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
		v21 = m.T0[v20].(func(*base.Module, int32, int32, int32) int32)(m, int32(100), v15, l1+int32(1))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
			v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
			if base.Ui64(v23) < base.Ui64(v24+int64(65536)) {
				m.G0 = v8 + int32(16)
				return
			} else {
				v31 = m.G0
				v32 = int32(16)
				v33 = v31 - v32
				m.G0 = v33
				F___gettimeofday(m, v33)
				mBase = m.M
				v36 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
				v37 = int64(*(*int32)(unsafe.Add(mBase, uint32(v33)+8)))
				m.G0 = v33 + v32
				v45 = v37 + v36*int64(1000000) - int64(946684800000000)
				v46 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v46
				v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
				if v45 <= v48 {
					v65 = int32(0)
				} else {
					v51 = int32(2147483647)
					v54 = v45 - v48
					if base.B2i32(int64(0) < v48)^base.B2i32(v54 < v45) != 0 {
						v65 = v51
					} else {
						if int64(2147483646000) < v54 {
							v65 = v51
						} else {
							v62 = base.I64_div_s(v54+int64(999), int64(1000))
							v65 = base.I32_wrap_i64(v62)
						}
					}
				}
				if v65 <= int32(999) {
					v68 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
					if v68 <= v45 {
						m.G0 = v8 + int32(16)
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v45
						F_pq_beginmessage(m, v8, int32(100))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_enlargeStringInfo(m, v8, int32(1))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
								v80 = int32(112)
								*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))) = uint8(v80)
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v77 + int32(1)
								v85 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
								F_enlargeStringInfo(m, v8, int32(8))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
									v92 = int64(56)
									v94 = int64(65280)
									v96 = int64(40)
									v99 = int64(16711680)
									v101 = int64(24)
									v103 = int64(4278190080)
									v105 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v89+v90))) = v85<<(uint(v92)%64) | v85&v94<<(uint(v96)%64) | (v85&v99<<(uint(v101)%64) | v85&v103<<(uint(v105)%64)) | (int64(base.Ui64(v85)>>(uint(v105)%64))&v103 | int64(base.Ui64(v85)>>(uint(v101)%64))&v99 | (int64(base.Ui64(v85)>>(uint(v96)%64))&v94 | int64(base.Ui64(v85)>>(uint(v92)%64))))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v89 + int32(8)
									F_pq_endmessage(m, v8)
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return
									} else {
										v134 = *(*int32)(unsafe.Add(mBase, _consts[218]))
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
										v136 = m.T0[v135].(func(*base.Module) int32)(m)
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v45
					F_pq_beginmessage(m, v8, int32(100))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						F_enlargeStringInfo(m, v8, int32(1))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v80 = int32(112)
							*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))) = uint8(v80)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v77 + int32(1)
							v85 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
							F_enlargeStringInfo(m, v8, int32(8))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
								v92 = int64(56)
								v94 = int64(65280)
								v96 = int64(40)
								v99 = int64(16711680)
								v101 = int64(24)
								v103 = int64(4278190080)
								v105 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v89+v90))) = v85<<(uint(v92)%64) | v85&v94<<(uint(v96)%64) | (v85&v99<<(uint(v101)%64) | v85&v103<<(uint(v105)%64)) | (int64(base.Ui64(v85)>>(uint(v105)%64))&v103 | int64(base.Ui64(v85)>>(uint(v101)%64))&v99 | (int64(base.Ui64(v85)>>(uint(v96)%64))&v94 | int64(base.Ui64(v85)>>(uint(v92)%64))))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v89 + int32(8)
								F_pq_endmessage(m, v8)
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return
								} else {
									v134 = *(*int32)(unsafe.Add(mBase, _consts[218]))
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
									v136 = m.T0[v135].(func(*base.Module) int32)(m)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
		v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		if base.Ui64(v23) < base.Ui64(v24+int64(65536)) {
			m.G0 = v8 + int32(16)
			return
		} else {
			v31 = m.G0
			v32 = int32(16)
			v33 = v31 - v32
			m.G0 = v33
			F___gettimeofday(m, v33)
			mBase = m.M
			v36 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
			v37 = int64(*(*int32)(unsafe.Add(mBase, uint32(v33)+8)))
			m.G0 = v33 + v32
			v45 = v37 + v36*int64(1000000) - int64(946684800000000)
			v46 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v46
			v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
			if v45 <= v48 {
				v65 = int32(0)
			} else {
				v51 = int32(2147483647)
				v54 = v45 - v48
				if base.B2i32(int64(0) < v48)^base.B2i32(v54 < v45) != 0 {
					v65 = v51
				} else {
					if int64(2147483646000) < v54 {
						v65 = v51
					} else {
						v62 = base.I64_div_s(v54+int64(999), int64(1000))
						v65 = base.I32_wrap_i64(v62)
					}
				}
			}
			if v65 <= int32(999) {
				v68 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
				if v68 <= v45 {
					m.G0 = v8 + int32(16)
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v45
					F_pq_beginmessage(m, v8, int32(100))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						F_enlargeStringInfo(m, v8, int32(1))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v80 = int32(112)
							*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))) = uint8(v80)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v77 + int32(1)
							v85 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
							F_enlargeStringInfo(m, v8, int32(8))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
								v92 = int64(56)
								v94 = int64(65280)
								v96 = int64(40)
								v99 = int64(16711680)
								v101 = int64(24)
								v103 = int64(4278190080)
								v105 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v89+v90))) = v85<<(uint(v92)%64) | v85&v94<<(uint(v96)%64) | (v85&v99<<(uint(v101)%64) | v85&v103<<(uint(v105)%64)) | (int64(base.Ui64(v85)>>(uint(v105)%64))&v103 | int64(base.Ui64(v85)>>(uint(v101)%64))&v99 | (int64(base.Ui64(v85)>>(uint(v96)%64))&v94 | int64(base.Ui64(v85)>>(uint(v92)%64))))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v89 + int32(8)
								F_pq_endmessage(m, v8)
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return
								} else {
									v134 = *(*int32)(unsafe.Add(mBase, _consts[218]))
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
									v136 = m.T0[v135].(func(*base.Module) int32)(m)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v45
				F_pq_beginmessage(m, v8, int32(100))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					F_enlargeStringInfo(m, v8, int32(1))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v80 = int32(112)
						*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))) = uint8(v80)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v77 + int32(1)
						v85 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
						F_enlargeStringInfo(m, v8, int32(8))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v92 = int64(56)
							v94 = int64(65280)
							v96 = int64(40)
							v99 = int64(16711680)
							v101 = int64(24)
							v103 = int64(4278190080)
							v105 = int64(8)
							*(*int64)(unsafe.Add(mBase, uint32(v89+v90))) = v85<<(uint(v92)%64) | v85&v94<<(uint(v96)%64) | (v85&v99<<(uint(v101)%64) | v85&v103<<(uint(v105)%64)) | (int64(base.Ui64(v85)>>(uint(v105)%64))&v103 | int64(base.Ui64(v85)>>(uint(v101)%64))&v99 | (int64(base.Ui64(v85)>>(uint(v96)%64))&v94 | int64(base.Ui64(v85)>>(uint(v92)%64))))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v89 + int32(8)
							F_pq_endmessage(m, v8)
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return
							} else {
								v134 = *(*int32)(unsafe.Add(mBase, _consts[218]))
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
								v136 = m.T0[v135].(func(*base.Module) int32)(m)
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
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
func F_bbsink_forward_archive_contents(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	m.T0[v5].(func(*base.Module, int32, int32))(m, v3, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_bbsink_forward_begin_manifest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
	m.T0[v4].(func(*base.Module, int32))(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_bbsink_throttle_begin_backup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	F_bbsink_forward_begin_backup(m, l0)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		v7 = m.G0
		v8 = int32(16)
		v9 = v7 - v8
		m.G0 = v9
		F___gettimeofday(m, v9)
		mBase = m.M
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
		v13 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+8)))
		m.G0 = v9 + v8
		*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v13 + v12*int64(1000000) - int64(946684800000000)
		return
	}
}
