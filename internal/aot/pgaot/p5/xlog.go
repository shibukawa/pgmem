package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XLOGShmemSize(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	if v9 != int32(-1) {
		v58 = F_mul_size(m, int32(128), int32(9))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return int32(0)
		} else {
			v60 = F_add_size(m, int32(448), v58)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				v64 = *(*int32)(unsafe.Add(mBase, _consts[129]))
				v65 = F_mul_size(m, int32(8), v64)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					v67 = F_add_size(m, v60, v65)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v70 = F_add_size(m, v67, int32(8192))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, _consts[129]))
							v75 = F_mul_size(m, int32(8192), v74)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								v77 = F_add_size(m, v70, v75)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									m.G0 = v6 + int32(48)
									return v77
								}
							}
						}
					}
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[34]))
		v16 = base.I32_div_s(v14, int32(32))
		v18 = *(*int32)(unsafe.Add(mBase, _consts[116]))
		v20 = base.I32_div_s(v18, int32(8192))
		if v16 < v20 {
			v22 = v16
		} else {
			v22 = v20
		}
		if v22 <= int32(8) {
			v25 = int32(8)
		} else {
			v25 = v22
		}
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v25
		v31 = F_pg_snprintf(m, v6+int32(16), int32(32), int32(486504), v6)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v38 = int32(1)
			F_SetConfigOption(m, int32(134805), v6+int32(16), v38, v38)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, _consts[129]))
				if v43 != int32(-1) {
					v58 = F_mul_size(m, int32(128), int32(9))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						v60 = F_add_size(m, int32(448), v58)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, _consts[129]))
							v65 = F_mul_size(m, int32(8), v64)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = F_add_size(m, v60, v65)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v70 = F_add_size(m, v67, int32(8192))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, _consts[129]))
										v75 = F_mul_size(m, int32(8192), v74)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v77 = F_add_size(m, v70, v75)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												m.G0 = v6 + int32(48)
												return v77
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_SetConfigOption(m, int32(134805), v6+int32(16), int32(1), int32(10))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v58 = F_mul_size(m, int32(128), int32(9))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = F_add_size(m, int32(448), v58)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, _consts[129]))
								v65 = F_mul_size(m, int32(8), v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = F_add_size(m, v60, v65)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v70 = F_add_size(m, v67, int32(8192))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, _consts[129]))
											v75 = F_mul_size(m, int32(8192), v74)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												v77 = F_add_size(m, v70, v75)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													m.G0 = v6 + int32(48)
													return v77
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
func F_xlog_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = int32(base.Ui32(l0) >> (uint(int32(4)) % 32))
	if v4 != int32(15) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4<<(uint(int32(2))%32))+uint32(_consts[63])))
		v12 = v11
	} else {
		v12 = int32(0)
	}
	return v12
}
