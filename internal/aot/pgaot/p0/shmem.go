package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitializeShmemGUCs(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v4 = m.G0
	v6 = v4 - int32(112)
	m.G0 = v6
	v9 = F_ShmemGetRequestedSize(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = F_add_size(m, int32(_a_F_InitializeShmemGUCs_0), v9)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeShmemGUCs[0]))
			v15 = F_add_size(m, v11, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v21 = F_add_size(m, v15, int32(_a_F_InitializeShmemGUCs_1)-v15&int32(_a_F_InitializeShmemGUCs_2))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v24 = F_add_size(m, v21, int32(_a_F_InitializeShmemGUCs_3))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(base.Ui32(v24) >> (uint(int32(20)) % 32))
						v30 = v6 + int32(48)
						v34 = F_pg_sprintf(m, v30, int32(_a_F_InitializeShmemGUCs_4), v6+int32(32))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_5), v30, int32(0), int32(1))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v42 = v6 + int32(44)
								v45 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeShmemGUCs[1]))
								if v45 != 0 {
									v49 = v45 << (uint(int32(10)) % 32)
								} else {
									v49 = int32(_a_F_InitializeShmemGUCs_6)
								}
								if v49 != 0 {
								} else {
								}
								if v42 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v42))) = v49
								} else {
								}
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v62 != 0 {
									v63 = base.I32_div_u_s(v21, v62)
									if v21 != v62*v63 {
										v67 = F_add_size(m, v63, int32(1))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v69 = v67
											*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v69
											v72 = v6 + int32(48)
											v76 = F_pg_sprintf(m, v72, int32(_a_F_InitializeShmemGUCs_4), v6+int32(16))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_7), v72, int32(0), int32(1))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													v85 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeShmemGUCs[2]))
													*(*int32)(unsafe.Add(mBase, uint32(v6))) = v85 + int32(38)
													v90 = v6 + int32(48)
													v92 = F_pg_sprintf(m, v90, int32(_a_F_InitializeShmemGUCs_8), v6)
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_9), v90, int32(0), int32(1))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															m.G0 = v6 + int32(112)
															return
														}
													}
												}
											}
										}
									} else {
										v69 = v63
										*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v69
										v72 = v6 + int32(48)
										v76 = F_pg_sprintf(m, v72, int32(_a_F_InitializeShmemGUCs_4), v6+int32(16))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_7), v72, int32(0), int32(1))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												v85 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeShmemGUCs[2]))
												*(*int32)(unsafe.Add(mBase, uint32(v6))) = v85 + int32(38)
												v90 = v6 + int32(48)
												v92 = F_pg_sprintf(m, v90, int32(_a_F_InitializeShmemGUCs_8), v6)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_9), v90, int32(0), int32(1))
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return
													} else {
														m.G0 = v6 + int32(112)
														return
													}
												}
											}
										}
									}
								} else {
									v85 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeShmemGUCs[2]))
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v85 + int32(38)
									v90 = v6 + int32(48)
									v92 = F_pg_sprintf(m, v90, int32(_a_F_InitializeShmemGUCs_8), v6)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_9), v90, int32(0), int32(1))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return
										} else {
											m.G0 = v6 + int32(112)
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
