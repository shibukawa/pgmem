package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ScanPgRelation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	if v12 != 0 {
		F_ScanKeyInit(m, v9, int32(1), int32(3), int32(184), l0)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v22 = F_table_open(m, int32(1259), int32(1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if l2 != 0 {
					v25 = F_GetNonHistoricCatalogSnapshot(m, int32(1259))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = F_RegisterSnapshot(m, v25)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = v27
							v33 = int32(*(*uint8)(unsafe.Add(mBase, _consts[881])))
							v36 = F_systable_beginscan(m, v22, int32(2662), l1&v33, v29, int32(1), v9)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = F_systable_getnext(m, v36)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									if v38 != 0 {
										v40 = F_heap_copytuple(m, v38)
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return int32(0)
										} else {
											v42 = v40
											F_systable_endscan(m, v36)
											mBase = m.M
											v44 = m.ExcPending
											if v44 != 0 {
												return int32(0)
											} else {
												if v29 != 0 {
													F_UnregisterSnapshot(m, v29)
													mBase = m.M
													v46 = m.ExcPending
													if v46 != 0 {
														return int32(0)
													} else {
														F_sequence_close(m, v22, int32(1))
														mBase = m.M
														v49 = m.ExcPending
														if v49 != 0 {
															return int32(0)
														} else {
															m.G0 = v9 + int32(48)
															return v42
														}
													}
												} else {
													F_sequence_close(m, v22, int32(1))
													mBase = m.M
													v49 = m.ExcPending
													if v49 != 0 {
														return int32(0)
													} else {
														m.G0 = v9 + int32(48)
														return v42
													}
												}
											}
										}
									} else {
										v42 = int32(0)
										F_systable_endscan(m, v36)
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return int32(0)
										} else {
											if v29 != 0 {
												F_UnregisterSnapshot(m, v29)
												mBase = m.M
												v46 = m.ExcPending
												if v46 != 0 {
													return int32(0)
												} else {
													F_sequence_close(m, v22, int32(1))
													mBase = m.M
													v49 = m.ExcPending
													if v49 != 0 {
														return int32(0)
													} else {
														m.G0 = v9 + int32(48)
														return v42
													}
												}
											} else {
												F_sequence_close(m, v22, int32(1))
												mBase = m.M
												v49 = m.ExcPending
												if v49 != 0 {
													return int32(0)
												} else {
													m.G0 = v9 + int32(48)
													return v42
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v29 = int32(0)
					v33 = int32(*(*uint8)(unsafe.Add(mBase, _consts[881])))
					v36 = F_systable_beginscan(m, v22, int32(2662), l1&v33, v29, int32(1), v9)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = F_systable_getnext(m, v36)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 != 0 {
								v40 = F_heap_copytuple(m, v38)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v42 = v40
									F_systable_endscan(m, v36)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										if v29 != 0 {
											F_UnregisterSnapshot(m, v29)
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
												return int32(0)
											} else {
												F_sequence_close(m, v22, int32(1))
												mBase = m.M
												v49 = m.ExcPending
												if v49 != 0 {
													return int32(0)
												} else {
													m.G0 = v9 + int32(48)
													return v42
												}
											}
										} else {
											F_sequence_close(m, v22, int32(1))
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												m.G0 = v9 + int32(48)
												return v42
											}
										}
									}
								}
							} else {
								v42 = int32(0)
								F_systable_endscan(m, v36)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									if v29 != 0 {
										F_UnregisterSnapshot(m, v29)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int32(0)
										} else {
											F_sequence_close(m, v22, int32(1))
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												m.G0 = v9 + int32(48)
												return v42
											}
										}
									} else {
										F_sequence_close(m, v22, int32(1))
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(48)
											return v42
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
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(361925), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(498618), int32(355), int32(263954))
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
func F_pg_relation_is_updatable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_relation_is_updatable(m, v2, v3, base.B2i32(v4 != v3), v3)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
