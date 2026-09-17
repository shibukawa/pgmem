package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetColumnDefCollation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = F_get_typcollation(m, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
			if l0 == int32(0) {
				v23 = F_get_collation_oid(m, v18, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v48 = v19
					v49 = v23
					v52 = int32(0)
					if base.B2i32(v49 == v52)|v13 == v52 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67141764))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v64 = F_format_type_be(m, l2)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v64
									F_errmsg(m, int32(_a_F_GetColumnDefCollation_0), v11)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, l0, v48)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_GetColumnDefCollation_1), int32(570), int32(_a_F_GetColumnDefCollation_2))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
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
					} else {
						m.G0 = v11 + int32(32)
						return v49
					}
				}
			} else {
				v26 = v11 + int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(489)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v26))) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v26
				v32 = int32(_a_F_GetColumnDefCollation_3)
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_GetColumnDefCollation[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v33
				*(*int32)(unsafe.Add(mBase, _c_F_GetColumnDefCollation[0])) = v11 + int32(20)
				v40 = F_get_collation_oid(m, v18, int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
					*(*int32)(unsafe.Add(mBase, _c_F_GetColumnDefCollation[0])) = v43
					v48 = v19
					v49 = v40
					v52 = int32(0)
					if base.B2i32(v49 == v52)|v13 == v52 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67141764))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v64 = F_format_type_be(m, l2)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v64
									F_errmsg(m, int32(_a_F_GetColumnDefCollation_0), v11)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, l0, v48)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_GetColumnDefCollation_1), int32(570), int32(_a_F_GetColumnDefCollation_2))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
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
					} else {
						m.G0 = v11 + int32(32)
						return v49
					}
				}
			}
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
			if v45 != 0 {
				v46 = v45
			} else {
				v46 = v13
			}
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
			v48 = v47
			v49 = v46
			v52 = int32(0)
			if base.B2i32(v49 == v52)|v13 == v52 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67141764))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = F_format_type_be(m, l2)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v64
							F_errmsg(m, int32(_a_F_GetColumnDefCollation_0), v11)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								F_parser_errposition(m, l0, v48)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_GetColumnDefCollation_1), int32(570), int32(_a_F_GetColumnDefCollation_2))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
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
			} else {
				m.G0 = v11 + int32(32)
				return v49
			}
		}
	}
}
