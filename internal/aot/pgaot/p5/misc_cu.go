package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cursor_to_xmlschema(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_text_to_cstring(m, v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v19 = F_pg_detoast_datum_packed(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = F_text_to_cstring(m, v19)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_SPI_connect_ext(m, int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_GetPortalByName(m, v15)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							if v26 != 0 {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
								if v28 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(258))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v15
											F_errmsg(m, int32(_a_F_cursor_to_xmlschema_0), v8+int32(16))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_cursor_to_xmlschema_1), int32(3111), int32(_a_F_cursor_to_xmlschema_2))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v31 = int32(0)
									v34 = F_map_sql_table_to_xmlschema(m, v28, v31, base.B2i32(v17 != v31), v21)
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int32(0)
									} else {
										v36 = F_strlen(m, v34)
										mBase = m.M
										v38 = v36 + int32(1)
										v39 = F_SPI_palloc(m, v38)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											if v38 != 0 {
												v41 = F__emscripten_memcpy_bulkmem(m, v39, v34, v38)
												mBase = m.M
												v42 = v41
											} else {
												v42 = v39
											}
											v43 = F_SPI_finish(m)
											mBase = m.M
											v44 = m.ExcPending
											if v44 != 0 {
												return int32(0)
											} else {
												v45 = F_cstring_to_text(m, v42)
												mBase = m.M
												v46 = m.ExcPending
												if v46 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(32)
													return v45
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(259))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
										F_errmsg(m, int32(_a_F_cursor_to_xmlschema_3), v8)
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_cursor_to_xmlschema_1), int32(3107), int32(_a_F_cursor_to_xmlschema_2))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
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
}
