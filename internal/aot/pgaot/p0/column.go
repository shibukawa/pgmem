package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetColumnDefCollation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = F_get_typcollation(m, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
		if v16 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
			if l0 == int32(0) {
				v22 = F_get_collation_oid(m, v17, int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v49 = v18
					v50 = v22
					if v50 == int32(0) {
						m.G0 = v10 + int32(32)
						return v50
					} else {
						if v12 != 0 {
							m.G0 = v10 + int32(32)
							return v50
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v61 = F_format_type_be(m, l2)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v61
										F_errmsg(m, int32(178188), v10)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											F_parser_errposition(m, l0, v49)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(475707), int32(570), int32(250388))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
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
							}
						}
					}
				}
			} else {
				v25 = v10 + int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(489)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v25
				v31 = int32(4435896)
				v32 = *(*int32)(unsafe.Add(mBase, _consts[77]))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v32
				*(*int32)(unsafe.Add(mBase, _consts[77])) = v10 + int32(20)
				v39 = F_get_collation_oid(m, v17, int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(12))+8))
					*(*int32)(unsafe.Add(mBase, _consts[77])) = v44
					v49 = v18
					v50 = v39
					if v50 == int32(0) {
						m.G0 = v10 + int32(32)
						return v50
					} else {
						if v12 != 0 {
							m.G0 = v10 + int32(32)
							return v50
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v61 = F_format_type_be(m, l2)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v61
										F_errmsg(m, int32(178188), v10)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											F_parser_errposition(m, l0, v49)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(475707), int32(570), int32(250388))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
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
							}
						}
					}
				}
			}
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
			if v46 != 0 {
				v47 = v46
			} else {
				v47 = v12
			}
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
			v49 = v48
			v50 = v47
			if v50 == int32(0) {
				m.G0 = v10 + int32(32)
				return v50
			} else {
				if v12 != 0 {
					m.G0 = v10 + int32(32)
					return v50
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							v61 = F_format_type_be(m, l2)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v61
								F_errmsg(m, int32(178188), v10)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									F_parser_errposition(m, l0, v49)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(475707), int32(570), int32(250388))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
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
					}
				}
			}
		}
	}
}
